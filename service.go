package main

// Service does inspector gadgety stuff.
type Service interface {
	Inspect(string) string
}
