package main

import (
	"log"
)

func main() {
	log.Println("TODO: build-in functions, lambda functions")

	log.Println(ExecuteString("[0, 2, 1].map(function var_a -> var_a * 2);", DefaultTimeout))
}
