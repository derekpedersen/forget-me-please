package model

type Option struct {
	Key         string
	Value       string
	Display     string
	Description string
	Action      func() error
}
