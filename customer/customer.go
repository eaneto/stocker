package customer

import (
	"fmt"
	"time"
)

type CustomerNotFoundError struct {
	ID uint
}

func (e CustomerNotFoundError) Error() string {
	return fmt.Sprintf("Customer not found. id=%d", e.ID)
}

const CUSTOMER = "customer"

type Customer struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CustomerEntity struct {
	ID        uint
	Name      string
	CreatedAt time.Time
}
