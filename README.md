# MiningRigRentals.com for Go #

## Installation ##

    # install the library:
    go get github.com/bitbandi/go-miningrigrentals-api

    // Use in your .go code:
    import (
        "github.com/bitbandi/go-miningrigrentals-api"
    )

## API Documentation ##

Full godoc output from the latest code in master is available here:

http://godoc.org/github.com/bitbandi/go-miningrigrentals-api

## Quickstart ##

```go
package main

import (
    "github.com/bitbandi/go-miningrigrentals-api"
    "log"
)

func main() {
    client := miningrigrentals.NewClient("YOUR_KEY", "YOUR_SECRET")
	balance, err := client.GetBalance()
    if err != nil {
        log.Fatalln("Unable to query balance: ", err)
    }

    log.Printf("Your confirmed balance: %f btc\n", balance.Confirmed)
}
```