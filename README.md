# aaapi

Go Twitter Account Acticity API

## Installation

```bash
$ go get github.com/dqn/aaapi
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/dqn/aaapi"
)

func main() {
	a := aaapi.NewPremium(
		"CONSUMER_KEY",
		"CONSUMER_SECRET",
		"ACCESS_TOKEN",
		"ACCESS_TOKEN_SECRET",
		"ENV_NAME",
	)
	r, err := a.PostWebhooks("https://example.com/webhook")
	if err != nil {
		// handle error
	}

	fmt.Println(r.ID) // => webhook id
}
```

## Features

- [x] Premium AAAPI
- [ ] Enterprise AAAPI

### Premium AAAPI

- [x] Post webhooks
- [x] Get webhooks
- [x] Get webhooks with env name
- [x] Put webhooks
- [x] Post subscriptions
- [x] Get subscriptions count
- [x] Get subscriptions
- [x] Get subscriptions list
- [x] Delete webhooks
- [x] Delete webhooks
- [x] Delete subscriptions user
