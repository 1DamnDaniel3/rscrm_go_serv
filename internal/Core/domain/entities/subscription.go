package entities

import (
	"time"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type Subscription struct {
	ID          int64
	Name        string
	Price       valuetypes.Money
	Visit_limit int
	Active_from time.Time
	Active_to   time.Time
	Is_archived bool
	School_id   string
}
