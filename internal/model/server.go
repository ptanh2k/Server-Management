package model

type Server struct {
	Id       uint16 `json:"id,omitempty"`
	Name     string `json:"name"`
	Ip       string `json:"ip"`
	Port     uint16 `json:"port"`
	Status   bool   `json:"status"`
	Password string `json:"password"`
}

// Constructor
func (s *Server) Init(name string, ip string, port uint16, status bool, password string) {
	s.Name = name
	s.Ip = ip
	s.Port = port
	s.Status = status
	s.Password = password
}
