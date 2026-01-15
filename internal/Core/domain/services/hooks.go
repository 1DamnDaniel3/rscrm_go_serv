package services

type BeforeCreate interface {
	BeforeCreate() error
}

type BeforeUpdate interface {
	BeforeUpdate() error
}
