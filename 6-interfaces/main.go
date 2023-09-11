package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string { //if you dont use the passed object you can deleted
	return "Hi, I'm a english bot"
}

func (spanishBot) getGreeting() string {
	return "Hola soy un robot español"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting()) //this is called method call
}

/*
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}*/

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
