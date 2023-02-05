package main

import (
	"fmt"
	"log"
	"urlShortner/pkg/routes"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("!!!--- Error while reading config file %s ---!!!", err)
	}

	fmt.Println("!!!--- Loader environment secrets ---!!!")
	routes.AllRoutes()
	fmt.Println("working")
}
