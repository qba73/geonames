package main

import (
	"fmt"
	"log"

	"github.com/qba73/geonames"
)

func main() {
	// We exported valid "GEONAMES_USER" env var
	codes, err := geonames.GetPostCode("Fort William", "UK")
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range codes {
		fmt.Printf("%+v\n", c)
	}
}
