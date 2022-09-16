package model

type Server struct {
	Id       uint16 `json:"id,omitempty"`
	Name     string `json:"name"`
	Ip       string `json:"ip"`
	Port     uint16 `json:"port"`
	Status   bool   `json:"status"`
	Password string `json:"password"`
}
