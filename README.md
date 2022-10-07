[![Go Reference](https://pkg.go.dev/badge/github.com/qba73/geonames.svg)](https://pkg.go.dev/github.com/qba73/geonames)
![Go](https://github.com/qba73/geonames/workflows/Go/badge.svg)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/qba73/geonames)
[![Maintainability](https://api.codeclimate.com/v1/badges/b4cb743c9bb7f5c405ee/maintainability)](https://codeclimate.com/github/qba73/geonames/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/b4cb743c9bb7f5c405ee/test_coverage)](https://codeclimate.com/github/qba73/geonames/test_coverage)
![GitHub](https://img.shields.io/github/license/qba73/geonames)



# geonames
`geonames` is a Go library for [Geonames Web Services](http://www.geonames.org) (geonames.org). The GeoNames geographical database covers all countries and contains over eleven million placenames.

## Setting your username

To use this client library with your geonames account, you will need a unique username. Go to the [geonames login](https://www.geonames.org/login) and register.

## Using the Go library

```go
import github.com/qba73/geonames
```

## Creating a client

Export ENV Var or read the username from your app config

```go
user := os.Getenv("GEO_USER")
client := geonames.NewClient(user)
```

## Bugs and feature request

If you find a bug in the ```geonames``` client library, please [open an issue](https://github.com/qba73/geonames/issues). If you'd like a feature added or improved, let me know via an issue.

The project is under development, and not all the functionality of the GeoNames Web Services is implemented yet.

Pull requests welcome!
