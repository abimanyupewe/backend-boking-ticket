package entity

// Response adalah format standar untuk balasan API
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"pesan"`
	Data    interface{} `json:"data,omitempty"`
}
