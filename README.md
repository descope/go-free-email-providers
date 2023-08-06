go-free-email-providers
----
[![Go Reference](https://pkg.go.dev/badge/github.com/lindell/go-burner-email-providers/burner.svg)](https://pkg.go.dev/github.com/lindell/go-burner-email-providers/burner)
[![Daily list sync](https://github.com/lindell/go-burner-email-providers/workflows/Daily%20list%20sync/badge.svg)](https://github.com/lindell/go-burner-email-providers/actions?query=workflow%3A%22Daily+list+sync%22)
[![Go Report Card](https://goreportcard.com/badge/github.com/lindell/go-burner-email-providers)](https://goreportcard.com/report/github.com/lindell/go-burner-email-providers)


Go package that detects free email providers 

## Installation

```
go get github.com/descope/go-free-email-providers
```

## Usage

```go
import (
    "github.com/descope/go-free-email-providers/free"
)

func main() {
	isBurnerEmail := burner.IsFreeEmail("test@gmail.org")
	fmt.Println(isBurnerEmail) // true

	isBurnerEmail = burner.IsFreeEmail("johan@descope.com")
	fmt.Println(isBurnerEmail) // false

	isBurnerDomain := burner.IsFreeDomain("gmail.org")
	fmt.Println(isBurnerDomain) // true

	isBurnerEmail = burner.IsFreeDomain("descope.com")
	fmt.Println(isBurnerEmail) // false
}
```

