package main

import (
	"fmt"
	"time"
)

func main() {
	ninjas := []string{"Taji", "Zume", "Akara"}

	for _, ninja := range ninjas {
		go attack(ninja)
	}

	time.Sleep(time.Microsecond * 200)
}

func attack(ninja string) {
	fmt.Println("Throw an kunai to", ninja)
	time.Sleep(time.Microsecond * 500)
}
