![Go](https://github.com/qba73/geonames/workflows/Go/badge.svg)
[![Maintainability](https://api.codeclimate.com/v1/badges/b4cb743c9bb7f5c405ee/maintainability)](https://codeclimate.com/github/qba73/geonames/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/b4cb743c9bb7f5c405ee/test_coverage)](https://codeclimate.com/github/qba73/geonames/test_coverage)
![GitHub](https://img.shields.io/github/license/qba73/geonames)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/qba73/geonames)


# geonames
Go client library for Geonames Web Services (geonames.org)


## Usage

- Register your username at [geonames.org](https://www.geonames.org/login)
- Export ENV Var or read the username from your app config

### Example

```go
package main

import (
	"fmt"

	"github.com/qba73/geonames"
)

func main() {
	user := os.Getenv("GEO_USER")
	geo := geonames.NewClient(user)

	res, err := geo.Wikipedia.Get("Dublin", "IE", 1)
	if err != nil {
		fmt.Println(err)
	}

	city := struct {
		Title       string
		GeoNameID   int
		Feature     string
		Lat         float64
		Lng         float64
		CountryCode string
		Language    string
	}{
		Title:       res.Geonames[0].Title,
		GeoNameID:   res.Geonames[0].GeoNameID,
		Feature:     res.Geonames[0].Feature,
		Lat:         res.Geonames[0].Lat,
		Lng:         res.Geonames[0].Lng,
		CountryCode: res.Geonames[0].CountryCode,
		Language:    res.Geonames[0].Lang,
	}

	fmt.Printf("%+v\n", city)

	// Prints:
	// {Title:Dublin GeoNameID:2964574 Feature:city Lat:53.343418 Lng:-6.267612 CountryCode:IE Language:en}
}
```


