[![Go Reference](https://pkg.go.dev/badge/github.com/qba73/geonames.svg)](https://pkg.go.dev/github.com/qba73/geonames)
![Go](https://github.com/qba73/geonames/workflows/Go/badge.svg)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/qba73/geonames)
[![Maintainability](https://api.codeclimate.com/v1/badges/b4cb743c9bb7f5c405ee/maintainability)](https://codeclimate.com/github/qba73/geonames/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/b4cb743c9bb7f5c405ee/test_coverage)](https://codeclimate.com/github/qba73/geonames/test_coverage)
![GitHub](https://img.shields.io/github/license/qba73/geonames)



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

	city, err := geo.Wikipedia.Get("Dublin", "IE", 1)
	if err != nil {
		// handle error
	}
}
```
