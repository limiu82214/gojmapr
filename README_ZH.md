# GetJSON

GetJSON是一個Golang庫，可以從複雜的JSON字符串中快速提取指定的屬性並轉換為Go結構。

使用GetJSON，您不需要宣告完整對應JSON的Go結構，只需要提供需要的屬性即可。

這使得GetJSON非常適合在存取第三方資源時提取指定資料使用，讓您的程式碼更加簡潔易讀。

## 特點

簡單易用：只需要添加幾個標籤就可以輕鬆地從JSON字符串中提取所需的屬性。
支持嵌套屬性：可以輕鬆地從多層嵌套的JSON字符串中提取所需的屬性。

## 安裝

要使用GetJSON，首先需要將其添加到你的Golang項目中：

```go
go get github.com/limiu82214/getjson
```

## 使用方式

要使用GetJSON庫，只需要將其導入到你的代碼中，然後按照以下步驟進行操作：

1. 定義一個與你要解析的JSON字符串對應的結構體。
2. 為結構體中的每個屬性添加getjson標籤，以指定從JSON字符串中提取該屬性的路徑。(參考jpath)
3. 使用getjson.Unmarshal函數將JSON字符串解析為結構體對象。

## 使用方法

下面是一個簡單的示例，展示了如何使用GetJSON庫從JSON字符串中提取屬性。

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

更多的使用示例可以在項目中的測試代碼中找到。

## 測試

GetJSON使用testify套件進行測試。要運行測試，請使用以下命令：

```bash
go test -v ./...
```

## TODO

* [] 編寫介面，支援多種不同的第三方json套件，以提高GetJSON的適用性。
* [] 提供一個原生的介面，使得用戶可以透過GetJSON使用原本json的方法，從而讓使用者更加熟悉和方便使用。
