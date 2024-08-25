package internal

import "testing"

func TestRequestViaCep(t *testing.T) {
	r := &RequestViaCEP{}
	message, err := r.Do("87560000")
	if err != nil {
		t.Errorf("RequestViaCEP.Do() error = %v", err)
		return
	}
	if message.Cep != "87560000" {
		t.Errorf("RequestViaCEP.Do() = %v, want %v", message.Cep, "87560000")
	}
}
