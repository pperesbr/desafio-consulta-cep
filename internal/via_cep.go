package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type RequestViaCEP struct{}

type ResponseViaCep struct {
	Cep   string `json:"cep"`
	City  string `json:"localidade"`
	State string `json:"uf"`
}

func (r *RequestViaCEP) Do(cep string) (*Message, error) {

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	var response ResponseViaCep
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	var message Message
	message.Api = "ViaCEP"
	message.Cep = strings.Replace(response.Cep, "-", "", -1)
	message.City = response.City
	message.State = response.State

	return &message, nil
}
