package getjson

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ExampleSuite struct {
	suite.Suite
	complexJSONString string
	anserStruct       ExampleStruct
}

func (ex *ExampleSuite) TestSimpleJPath() {
	jsonString := ex.complexJSONString

	type tmpStruct struct {
		RequestID string `getjson:"$.request_id"`
	}

	var s tmpStruct
	err := Unmarshal([]byte(jsonString), &s)
	ex.Assert().Nil(err)
	ex.Assert().Equal(ex.anserStruct.RequestID, s.RequestID)
}

func (ex *ExampleSuite) TestSimpleJPathWithTime() {
	jsonString := ex.complexJSONString

	type tmpStruct struct {
		CreateAt time.Time `getjson:"$.create_at"`
	}

	var s tmpStruct
	err := Unmarshal([]byte(jsonString), &s)
	ex.Assert().Nil(err)
	ex.Assert().Equal(ex.anserStruct.CreateAt, s.CreateAt)
}
func (ex *ExampleSuite) TestNestedJPath() {
	jsonString := ex.complexJSONString

	type tmpStruct struct {
		Name  string `getjson:"$.user.name"`
		Email string `getjson:"$.user.email"`
	}

	var s tmpStruct
	err := Unmarshal([]byte(jsonString), &s)
	ex.Assert().Nil(err)
	ex.Assert().Equal(ex.anserStruct.User.Name, s.Name)
	ex.Assert().Equal(ex.anserStruct.User.Email, s.Email)
}

func (ex *ExampleSuite) TestNested2JPath() {
	jsonString := ex.complexJSONString

	type tmpStruct struct {
		ID    string  `getjson:"$.cart.items[0].product.id"`
		Price float64 `getjson:"$.cart.items.0.product.price"`
	}

	var s tmpStruct
	err := Unmarshal([]byte(jsonString), &s)
	ex.Assert().Nil(err)
	ex.Assert().Equal(ex.anserStruct.Cart.Items[0].Product.ID, s.ID)
	ex.Assert().Equal(ex.anserStruct.Cart.Items[0].Product.Price, s.Price)
}

func (ex *ExampleSuite) TestNestedStructJPath() {
	jsonString := ex.complexJSONString

	type tmpStruct struct {
		User struct {
			Name  string `getjson:"$.cart.items[0].product.name"`
			Email string `getjson:"$.user.email"`
		}
		ID string `getjson:"$.cart.items[0].product.id"`
	}

	s := tmpStruct{}
	err := Unmarshal([]byte(jsonString), &s)
	ex.Assert().Nil(err)
	ex.Assert().Equal(ex.anserStruct.Cart.Items[0].Product.Name, s.User.Name)
	ex.Assert().Equal(ex.anserStruct.User.Email, s.User.Email)
	ex.Assert().Equal(ex.anserStruct.Cart.Items[0].Product.ID, s.ID)
}

func TestExampleSuite(t *testing.T) {
	suite.Run(t, new(ExampleSuite))
}

type ExampleStruct struct {
	User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"user"`
	Cart struct {
		Items []struct {
			Product struct {
				ID          string  `json:"id"`
				Name        string  `json:"name"`
				Description string  `json:"description"`
				Price       float64 `json:"price"`
			} `json:"product"`
			Quantity int `json:"quantity"`
		} `json:"items"`
		Total float64 `json:"total"`
	} `json:"cart"`
	Shipping struct {
		Method  string `json:"method"`
		Address struct {
			Street string `json:"street"`
			City   string `json:"city"`
			State  string `json:"state"`
			Zip    string `json:"zip"`
		} `json:"address"`
		Fee float64 `json:"fee"`
	} `json:"shipping"`
	CreateAt  time.Time `json:"create_at"`
	RequestID string    `json:"request_id"`
}

func (ex *ExampleSuite) SetupTest() {
	ex.complexJSONString = `{
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
		"create_at": "2020-02-14T00:00:00Z",
		"request_id": "omg9487"
	}
	`

	err := json.Unmarshal([]byte(ex.complexJSONString), &ex.anserStruct)
	ex.Assert().Nil(err)
}
