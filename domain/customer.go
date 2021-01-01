package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CustomerNotFoundError struct {
	Code uuid.UUID
}

func (e CustomerNotFoundError) Error() string {
	return fmt.Sprintf("Customer not found. code=%s", e.Code)
}

const CUSTOMER = "customer"

type Customer struct {
	Code uuid.UUID `json:"code"`
	Name string    `json:"name"`
}

type CustomerEntity struct {
	ID        uint
	Code      uuid.UUID
	Name      string
	CreatedAt time.Time
}
