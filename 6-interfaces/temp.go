package main

type user struct{
	name string
}

type newbot interface{
	getGreeting(string, int)(string, error)
	getBotVersion() float64
	respondToUser(user) string
}