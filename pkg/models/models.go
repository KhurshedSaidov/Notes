package models

type ServerConfig struct {
	Ip   string `json: "ip"`
	Port string `json: "port`
}

type Note struct {
	Id          int    `json: "id"`
	Content     string `json: "content"`
	CreatedData string `json: "createddata"`
}
