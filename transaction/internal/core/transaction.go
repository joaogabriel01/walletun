package core

import "github.com/google/uuid"

type Transaction struct {
	ID              string    `json:"id"`
	Amount          float64   `json:"amount"`
	SourceAccountID uuid.UUID `json:"source_account_id"`
	TargetAccountID uuid.UUID `json:"target_account_id"`
}
