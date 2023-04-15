package main

import (
	"fmt"

	"github.com/ivanov-slk/tma-data-generator/internal/mock"
)

func main() {
	ouat, err := mock.JsonizeStaticOutput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", ouat)
}
