# proovl-go
Proovl Go

# Proovl Go Package

This package provides a simple way to send SMS messages using the Proovl API in Go. 

## Installation

To install the package, run: 

go get github.com/aarnhub/proovl-go/sms

## Usage

You can use the `SendSMS` and `BulkSMS` functions to send SMS messages. Here are examples of how to use them:

### SendSMS Example

```go
package main

import (
	"fmt"
	"github.com/aarnhub/proovl-go/sms"
)

func main() {
	user := "your_user"
	token := "your_token"
	route := "1"
	from := "your_from"
	to := "destination_number"
	text := "your_message_text"

	response, err := proovl.SendSMS(user, token, route, from, to, text)
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Response:", response)
	}
}
```

### BulkSMS Example

```go
package main

import (
	"fmt"
	"github.com/aarnhub/proovl-go/sms"
)

func main() {
	user := "your_user"
	token := "your_token"
	route := "1"
	from := "your_from"
	to := "destination_number1,destination_number2,destination_number3"
	text := "your_message_text"

	response, err := proovl.BulkSMS(user, token, route, from, to, text)
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Response:", response)
	}
}
```

Replace your_user, your_token, your_route, your_from, destination_number, and your_message_text with your Proovl API credentials and desired message details.

Make sure to import the package by adding import "github.com/aarnhub/proovl-go" to your code.
