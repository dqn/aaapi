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
	r, err := a.PostWebhooks("https://example.com/crc")
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
- [ ] Post subscriptions
- [ ] Get subscriptions count
- [ ] Get subscriptions
- [ ] Get subscriptions list
- [ ] Delete webhooks
- [ ] Delete webhooks
- [ ] Delete subscriptions user
