package model

type Product struct {
    Id      uint   `json:"id"`
    Name    string  `json:"name"`
    Photo []uint8   `json:"photo"`
}