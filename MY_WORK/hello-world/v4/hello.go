package main

import f"fmt"

const englishHelloPrefix string = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}

func main(){
	f.Println(Hello("Chris"))
}

