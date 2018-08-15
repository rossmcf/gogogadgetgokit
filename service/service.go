package service

import "strings"

type Service struct {
}

func (s Service) Inspect(string1 string) string {
	return "GO GO GADGET " + strings.ToUpper(string1)
}
