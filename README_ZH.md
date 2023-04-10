# gojmapr

![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/limiu82214/gojmapr?label=version) [![Go Reference](https://pkg.go.dev/badge/github.com/limiu82214/gojmapr.svg)](https://pkg.go.dev/github.com/limiu82214/gojmapr) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![codecov](https://codecov.io/gh/limiu82214/gojmapr/branch/master/graph/badge.svg?token=0XAK9BB5WL)](https://codecov.io/gh/limiu82214/gojmapr) [![Go Report Card](https://goreportcard.com/badge/github.com/limiu82214/gojmapr)](https://goreportcard.com/report/github.com/limiu82214/gojmapr) ![github actions workflow](https://github.com/limiu82214/gojmapr/actions/workflows/lint.yml/badge.svg)  

gojmapr是一個Golang庫，可以從複雜的JSON字符串中快速提取指定的屬性並轉換為Go結構。

使用gojmapr，您不需要宣告完整對應JSON的Go結構，只需要提供需要的屬性即可。

這使得gojmapr非常適合在存取第三方資源時提取指定資料使用，讓您的程式碼更加簡潔易讀。

## 特點

簡單易用：只需要添加幾個標籤就可以輕鬆地從JSON字符串中提取所需的屬性。
支持嵌套屬性：可以輕鬆地從多層嵌套的JSON字符串中提取所需的屬性。

## 安裝

要使用gojmapr，首先需要將其添加到你的Golang項目中：

```go
go get github.com/limiu82214/gojmapr
```

## 使用方式

要使用gojmapr庫，只需要將其導入到你的代碼中，然後按照以下步驟進行操作：

1. 定義一個與你要解析的JSON字符串對應的結構體。
2. 為結構體中的每個屬性添加gojmapr標籤，以指定從JSON字符串中提取該屬性的路徑。(參考jpath)
3. 使用gojmapr.Unmarshal函數將JSON字符串解析為結構體對象。

## 使用方法

下面是一個簡單的示例，展示了如何使用gojmapr庫從JSON字符串中提取屬性。

```go
package main

import (
    "fmt"

    "github.com/limiu82214/gojmapr"
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
        Name string `gojmapr:"user.name"`
    }

    var s tmpStruct
    err := gojmapr.Unmarshal([]byte(jsonString), &s)
    if err != nil {
        panic(err)
    }

    fmt.Println(s.Name) // Output: John

    type tmpStruct2 struct {
        ID    string  `gojmapr:"$.cart.items[0].product.id"`
        Price float64 `gojmapr:"$.cart.items.0.product.price"`
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

更多的使用示例可以在項目中的測試代碼中找到。

## 使用其他 Unmarshal 套件

```go

import jsoniter "github.com/json-iterator/go"

type tmpStruct struct {
    RequestID string `gojmapr:"$.request_id"`
}

SetUnmarshalFunc(jsoniter.Unmarshal) // 您可以使用其他解析套件，例如json-iterator

var s tmpStruct
err := Unmarshal([]byte(jsonString), &s)
ex.Assert().Nil(err)
ex.Assert().Equal(ex.anserStruct.RequestID, s.RequestID)
```

gojmapr可以使用其他解析套件，例如json-iterator。

## 測試

gojmapr使用testify套件進行測試。  
要運行測試，請使用以下命令：

```bash
go test -v ./...
```

## 依賴

* [github.com/limiu82214/gojpath](http://github.com/limiu82214/gojpath)

## 其他

如果您在使用過程中有任何問題，歡迎在 GitHub 專案上發起一個 issue，或是透過 email 與我聯繫。  
如果您認為這個專案對您有所幫助，也請不吝給予一個 star。

## 授權

[MIT License](./LICENSE)
