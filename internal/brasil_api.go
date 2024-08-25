package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type RequestBrasilApi struct{}

type ResponseBrasilApi struct {
	Cep   string `json:"cep"`
	City  string `json:"city"`
	State string `json:"state"`
}

func (r *RequestBrasilApi) Do(cep string) (*Message, error) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	var response ResponseBrasilApi
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	var message Message
	message.Api = "BrasilAPI"
	message.Cep = strings.Replace(response.Cep, "-", "", -1)
	message.City = response.City
	message.State = response.State

	return &message, nil
}
