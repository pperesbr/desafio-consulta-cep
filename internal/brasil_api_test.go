package internal

import "testing"

func TestRequestBrasilApi(t *testing.T) {

	r := &RequestBrasilApi{}
	message, err := r.Do("87560000")
	if err != nil {
		t.Errorf("RequestBrasilApi.Do() error = %v", err)
		return
	}
	if message.Cep != "87560000" {
		t.Errorf("RequestBrasilApi.Do() = %v, want %v", message.Cep, "87560000")
	}
}
