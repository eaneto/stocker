package domain

import (
	"time"

	"github.com/google/uuid"
)

const CUSTOMER = "customer"

type CustomerEntity struct {
	ID        uint
	Code      uuid.UUID
	Name      string
	CreatedAt time.Time
}
