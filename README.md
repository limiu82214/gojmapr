# GetJSON

![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/limiu82214/gojmapr?label=version) [![Go Reference](https://pkg.go.dev/badge/github.com/limiu82214/gojmapr.svg)](https://pkg.go.dev/github.com/limiu82214/gojmapr) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![codecov](https://codecov.io/gh/limiu82214/gojmapr/branch/master/graph/badge.svg?token=QX59JZM663)](https://codecov.io/gh/limiu82214/gojmapr) [![Go Report Card](https://goreportcard.com/badge/github.com/limiu82214/gojmapr)](https://goreportcard.com/report/github.com/limiu82214/gojmapr) ![github actions workflow](https://github.com/limiu82214/gojmapr/actions/workflows/go.yml/badge.svg)  

[中文版文檔](./README_ZH.md)

GetJSON is a Golang library that allows for quick extraction of specified properties from complex JSON strings and converts them into corresponding Go structures.

With GetJSON, you don't need to declare a complete Go structure that corresponds to the entire JSON structure, just provide the required properties.

This makes GetJSON particularly useful for extracting specific data when accessing third-party resources, making your code more concise and readable.

## Features

Easy to use: Easily extract required properties from JSON strings by adding a few tags.
Supports nested properties: Easily extract required properties from JSON strings with multiple nested levels.

## Installation

To use GetJSON, first add it to your Golang project:

```go
go get github.com/limiu82214/getjson
```

## Usage

To use the GetJSON library, simply import it into your code and follow these steps:

Define a struct that corresponds to the JSON string you want to parse.
Add the getjson tag to each property in the struct to specify the path to extract the property from the JSON string (reference jpath).
Use the getjson.Unmarshal function to parse the JSON string into a struct object.

## Example

Here's a simple example that shows how to use the GetJSON library to extract properties from a JSON string.

```go
package main

import (
    "fmt"

    "github.com/limiu82214/getjson"
)

func main() {
    jsonString := `{
        "user": {
            "name": "John",
            "email": "john@example.com"
        },
        "cart": {
            "items": [
                {
                    "product": {
                        "id": "123",
                        "name": "Product A",
                        "description": "Product A description",
                        "price": 10.99
                    },
                    "quantity": 2
                },
                {
                    "product": {
                        "id": "456",
                        "name": "Product B",
                        "description": "Product B description",
                        "price": 5.99
                    },
                    "quantity": 1
                }
            ],
            "total": 27.97
        },
        "shipping": {
            "method": "standard",
            "address": {
                "street": "123 Main St",
                "city": "Anytown",
                "state": "CA",
                "zip": "12345"
            },
            "fee": 5.99
        },
        "create_at": "2020-01-01T00:00:00Z"
    }`

    type tmpStruct struct {
        Name string `getjson:"user.name"`
    }

    var s tmpStruct
    err := getjson.Unmarshal([]byte(jsonString), &s)
    if err != nil {
        panic(err)
    }

    fmt.Println(s.Name) // Output: John

    type tmpStruct2 struct {
        ID    string  `getjson:"$.cart.items[0].product.id"`
        Price float64 `getjson:"$.cart.items.0.product.price"`
    }

    var s2 tmpStruct
    err := Unmarshal([]byte(jsonString), &s2)
    if err != nil {
        panic(err)
    }

    fmt.Println(s2.ID) // Output: 123
    fmt.Println(s2.Price) // Output: 10.99
}
```

More examples of usage can be found in the test code in the project.

## Use other Unmarshal package

```go

import jsoniter "github.com/json-iterator/go"

type tmpStruct struct {
    RequestID string `getjson:"$.request_id"`
}

SetUnmarshalFunc(jsoniter.Unmarshal) // You can use other Unmarshal module ex: json-iterator

var s tmpStruct
err := Unmarshal([]byte(jsonString), &s)
ex.Assert().Nil(err)
ex.Assert().Equal(ex.anserStruct.RequestID, s.RequestID)
```

gojmapr can use other Unmarshal package ex: json-iterator.  

## Testing

GetJSON uses the testify package for testing. To run the tests, use the following command:

```bash
go test -v ./...
```

## Dependency

* github.com/limiu82214/gojpath

## Other

If you encounter any issues during use, please feel free to raise an issue on the GitHub project or contact me via email. If you find this project helpful, please consider giving it a star.

## LICENSE

[MIT License](./LICENSE)
