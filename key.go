package chef

import (
)

type ChefKey struct {
	Name           string `json:"name"`
	PublicKey      string `json:"public_key"`
	ExpirationDate string `json:"expiration_date"`
	Uri            string `json:"uri"`
	PrivateKey     string `json:"private_key"`
}

type AccessKey struct {
	KeyName        string `json:"name,omitempty"`
	PublicKey      string `json:"public_key,omitempty"`
	ExpirationDate string `json:"expiration_date,omitempty"`
}

type KeyItem struct {
	KeyName string `json:"name,omitempty"`
	Uri     string `json:"uri,omitempty"`
	Expired bool   `json:"expired,omitempty"`
}
