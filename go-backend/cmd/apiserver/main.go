package main

import (
	"github.com/go-next-pizza/internal/app/apiserver"
	_ "github.com/go-next-pizza/internal/app/storage"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	
	if err != nil {
		panic(err)
	}
}

func main() {
	Init()

	config := apiserver.NewConfig()

	if err := apiserver.Start(config); err != nil {
		panic(err)
	}
	
}
