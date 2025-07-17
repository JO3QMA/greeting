package main

import (
	"fmt"
	"time"
)

func Greet() string {
	var current_hour int = time.Now().Hour()
	if (current_hour >= 4) && (current_hour < 10) {
		return "おはよう"
	} else if (current_hour >= 10) && (current_hour < 17) {
		return "こんにちは"
	} else if (current_hour >= 17) || (current_hour < 4) {
		return "こんばんは"
	}
	return ""
}

func main() {
	fmt.Println(Greet())
}
