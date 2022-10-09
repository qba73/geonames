package main

import (
	"fmt"
	"log"
	"os"

	"github.com/qba73/geonames"
)

func main() {
	// We exported valid "GEONAMES_USER" env var
	geo, err := geonames.NewClient(os.Getenv("GEONAMES_USER"))
	if err != nil {
		panic(err)
	}

	codes, err := geo.GetPostCode("Fort William", "UK")
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range codes {
		fmt.Printf("%+v\n", c)
	}
}
