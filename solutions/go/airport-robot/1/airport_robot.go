package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

type Italian struct {}

type Portuguese struct {}

func SayHello(n string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(n))
}

func (Italian) LanguageName() string {
	return "Italian"
}

func (Italian) Greet(n string) string {
	return fmt.Sprintf("Ciao %s!", n)
}

func (Portuguese) LanguageName() string {
	return "Portuguese"
}

func (Portuguese) Greet(n string) string {
	return fmt.Sprintf("Olá %s!", n)
}