package model

type Configuration struct {
	ServerPort uint16   `json:"server_port"`
	Database   Database `json:"database"`
}

type Database struct {
	Engine   string `json:"engine"`
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     uint   `json:"port"`
	Name     string `json:"name"`
	SSLMode  string `json:"ssl_mode"`
}
