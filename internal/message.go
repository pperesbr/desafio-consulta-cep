package internal

type Message struct {
	Cep   string `json:"cep"`
	State string `json:"state"`
	City  string `json:"city"`
	Api   string `json:"api"`
}
