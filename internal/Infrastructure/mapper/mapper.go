package mapper

import "reflect"

func MapDomainToDTO[D, R any](domain *D) *R {
	dto := new(R)
	domainVal := reflect.ValueOf(domain).Elem()
	dtoVal := reflect.ValueOf(dto).Elem()

	for i := 0; i < dtoVal.NumField(); i++ {
		field := dtoVal.Type().Field(i)
		if f := domainVal.FieldByName(field.Name); f.IsValid() {
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
			domainVal.Field(i).Set(f)
		}
	}

	return domain
}
