package tickets

import (
	"github.com/google/uuid"
	"time"
)

type Ticket struct {
	ID          uuid.UUID
	Departure   string
	Destination string
	Date        time.Time
}
