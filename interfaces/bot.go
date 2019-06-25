package main

import "fmt"

type bot interface {
	getGreetings() string
}
type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreetings() string {
	// VERY custom logic for generating an english greeting
	return "Hello there!"
}

func (spanishBot) getGreetings() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreetings())
}
