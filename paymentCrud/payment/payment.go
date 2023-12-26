package payment

import (
	"github.com/google/uuid"
	"time"
)

type Payment struct {
	Id           uuid.UUID
	Payment_type string
	Date         time.Time
	Amount       int
}
