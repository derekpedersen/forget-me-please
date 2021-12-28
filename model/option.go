package model

type Option struct {
	Key         string
	Value       string
	Description string
	Action      func()
}
