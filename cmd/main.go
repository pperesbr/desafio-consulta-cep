package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pperesbr/desafio-consulta-cep/internal"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please, provide a CEP")
		os.Exit(1)
	}

	cep := os.Args[1]

	brasilApiChanel := make(chan internal.Message)
	viaCepApiChanel := make(chan internal.Message)

	var requestBrasilApi internal.Request = &internal.RequestBrasilApi{}
	var requestViaCep internal.Request = &internal.RequestViaCEP{}

	go func() {
		message, err := requestBrasilApi.Do(cep)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		brasilApiChanel <- *message
	}()

	go func() {
		message, err := requestViaCep.Do(cep)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viaCepApiChanel <- *message
	}()

	select {
	case brasilApiMessage := <-brasilApiChanel:
		fmt.Printf("API: %s, CEP: %s, State: %s, City: %s\n", brasilApiMessage.Api, brasilApiMessage.Cep, brasilApiMessage.State, brasilApiMessage.City)
	case viaCepApiMessage := <-viaCepApiChanel:
		fmt.Printf("API: %s, CEP: %s, State: %s, City: %s\n", viaCepApiMessage.Api, viaCepApiMessage.Cep, viaCepApiMessage.State, viaCepApiMessage.City)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}
