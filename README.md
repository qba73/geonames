![Go](https://github.com/qba73/geonames/workflows/Tests/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/geonames)](https://goreportcard.com/report/github.com/qba73/geonames)
![GitHub](https://img.shields.io/github/license/qba73/geonames)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/qba73/geonames)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/qba73/geonames)
[![Go Reference](https://pkg.go.dev/badge/github.com/qba73/geonames.svg)](https://pkg.go.dev/github.com/qba73/geonames)
[![CodeQL](https://github.com/qba73/geonames/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/qba73/geonames/actions/workflows/github-code-scanning/codeql)

# geonames

`geonames` is a Go library for [Geonames Web Services](http://www.geonames.org) (geonames.org). The GeoNames geographical database covers all countries and contains over eleven million placenames.

## Setting your username

To use this client library with your geonames account, you will need a unique username. Go to the [geonames login](https://www.geonames.org/login) and register.

## Using the Go library

```go
import github.com/qba73/geonames
```

## Creating a client

Export ENV Var ```GEONAMES_USER```

```go
client, err := geonames.NewClient(os.Getenv("GEONAMES_USER"))
if err != nil {
    // handle error
}
```

Provide username directly:

```go
client, err := geonames.NewClient("dummy_user")
if err != nil {
    // handle error
}
```

## Complete example programs

You can see complete example programs which retrive coordinates and postal codes in the [examples](examples/) folder.

- postal codes lookup example: [examples/postal](examples/postal/main.go)
- coordinates lookup example: [examples/wikipedia](examples/wikipedia/main.go)

## Bugs and feature request

If you find a bug in the ```geonames``` client library, please [open an issue](https://github.com/qba73/geonames/issues). If you'd like a feature added or improved, let me know via an issue.

The project is under development, and not all the functionality of the GeoNames Web Services is implemented yet.

Pull requests welcome!
