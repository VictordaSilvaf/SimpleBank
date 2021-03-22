// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Conta struct {
	ID           int64     `json:"id"`
	Proprietario string    `json:"proprietario"`
	Saldo        int64     `json:"saldo"`
	Moeda        string    `json:"moeda"`
	CreatedAt    time.Time `json:"created_at"`
}

type Entrada struct {
	ID         int64     `json:"id"`
	IDConta    int64     `json:"id_conta"`
	Quantidade string    `json:"quantidade"`
	CreatedAt  time.Time `json:"created_at"`
}

type Transferencia struct {
	ID              int64     `json:"id"`
	EmissorIDConta  int64     `json:"emissor_id_conta"`
	ReceptorIDConta int64     `json:"receptor_id_conta"`
	Quantidade      string    `json:"quantidade"`
	CreatedAt       time.Time `json:"created_at"`
}