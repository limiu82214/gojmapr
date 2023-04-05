[中文](./README_ZH.md)
# GetJSON

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
}
```

More examples of usage can be found in the test code in the project.

## Testing

GetJSON uses the testify package for testing. To run the tests, use the following command:

```bash
go test -v ./...
```

## TODO

* [] Write an interface that supports multiple different third-party JSON packages to enhance the versatility of GetJSON.
* [] Provide a native interface that allows users to use the original JSON methods through GetJSON, making it more familiar and convenient for users.
