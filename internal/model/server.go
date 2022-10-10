package model

type Server struct {
	ID     uint16 `json:"id,omitempty"`
	Name   string `json:"name"`
	Ip     string `json:"ip"`
	Port   uint16 `json:"port"`
	Status bool   `json:"status"`
}

// Constructor
func (s *Server) Init(name string, ip string, port uint16, status bool) {
	s.Name = name
	s.Ip = ip
	s.Port = port
	s.Status = status
}
