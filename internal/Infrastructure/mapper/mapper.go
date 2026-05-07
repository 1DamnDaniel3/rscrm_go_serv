package mapper

import (
	"log"
	"reflect"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
	"github.com/shopspring/decimal"
)

// Автоматически маппит Domain -> DTO
func MapDomainToDTO[D, R any](domain *D) *R {
	dto := new(R)
	domainVal := reflect.ValueOf(domain).Elem()
	dtoVal := reflect.ValueOf(dto).Elem()

	for i := 0; i < dtoVal.NumField(); i++ {
		field := dtoVal.Type().Field(i)
		if f := domainVal.FieldByName(field.Name); f.IsValid() {
			// ДЛЯ MAPPING денежных значений
			if f.Type() == reflect.TypeOf(valuetypes.Money{}) && field.Type.Kind() == reflect.String {
				dtoVal.Field(i).SetString(f.Interface().(valuetypes.Money).String())
				continue
			}
			dtoVal.Field(i).Set(f)
		}
	}
	return dto

}

// MapDTOToDomain копирует все совпадающие поля из DTO в Domain
func MapDTOToDomain[D, R any](dto *D) *R {
	domain := new(R)
	dtoVal := reflect.ValueOf(dto).Elem()
	domainVal := reflect.ValueOf(domain).Elem()

	for i := 0; i < domainVal.NumField(); i++ {
		field := domainVal.Type().Field(i)
		if f := dtoVal.FieldByName(field.Name); f.IsValid() {

			if field.Type == reflect.TypeOf(valuetypes.Money{}) && f.Kind() == reflect.String {
				m, err := valuetypes.NewMoneyFromString(f.String())
				if err != nil {
					log.Printf("Mapper warning: cannot convert field %s=%q to Money, defaulting to 0", field.Name, f.String())
					m = valuetypes.Money{Amount: decimal.Zero}
				}
				domainVal.Field(i).Set(reflect.ValueOf(m))
				continue
			}
			domainVal.Field(i).Set(f)
		}
	}

	return domain
}
