package valuetypes

import (
	"database/sql/driver"
	"fmt"

	"github.com/shopspring/decimal"
)

type Money struct {
	Amount decimal.Decimal
}

func NewMoneyFromString(s string) (Money, error) {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return Money{}, fmt.Errorf("invalid money format: %w", err)
	}
	return Money{Amount: d}, nil
}

func (m Money) Add(o Money) Money {
	return Money{Amount: m.Amount.Add(o.Amount)}
}

func (m Money) Sub(o Money) Money {
	return Money{Amount: m.Amount.Sub(o.Amount)}
}

func (m Money) MulPercent(p decimal.Decimal) Money {
	factor := p.Div(decimal.NewFromInt(100))
	return Money{Amount: m.Amount.Mul(factor).Round(2)}
}

func (m Money) String() string {
	return m.Amount.StringFixed(2)
}

// -=== GORM

func (m Money) Value() (driver.Value, error) {
	return m.Amount.StringFixed(2), nil
}

func (m *Money) Scan(value interface{}) error {
	switch v := value.(type) {
	case float64:
		m.Amount = decimal.NewFromFloat(v)
	case string:
		d, err := decimal.NewFromString(v)
		if err != nil {
			return err
		}
		m.Amount = d
	case []byte:
		d, err := decimal.NewFromString(string(v))
		if err != nil {
			return err
		}
		m.Amount = d
	default:
		return fmt.Errorf("cannot scan %T into Money", value)
	}
	return nil
}
