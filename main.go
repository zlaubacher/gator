package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {
	cfg1, err := config.Read()
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	err = cfg1.SetUser("Zach")
	if err != nil {
		fmt.Println("error setting user:", err)
		return
	}

	cfg2, err := config.Read()
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	fmt.Println(cfg2)
}
