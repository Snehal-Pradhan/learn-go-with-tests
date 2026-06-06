package main

import f "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string { 
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main(){
	f.Println(Hello("Chris"))
}