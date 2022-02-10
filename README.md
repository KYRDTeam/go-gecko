# CoinGecko API Client for Go

[![Build Status](https://travis-ci.com/superoo7/go-gecko.svg?branch=master)](https://travis-ci.com/superoo7/go-gecko) [![GoDoc](https://godoc.org/github.com/superoo7/go-gecko?status.svg)](https://godoc.org/github.com/superoo7/go-gecko)

Simple API Client for CoinGecko written in Go

<p align="center">
    <img src="gogecko.png" alt="gogecko" height="200" />
</p>

gopher resources from [free-gophers-pack](https://github.com/MariaLetta/free-gophers-pack)

## Available endpoint

[Refer to CoinGecko official API](https://www.coingecko.com/api)

|            Endpoint             | Status | Testing |            Function            |
| :-----------------------------: | :----: | :-----: | :----------------------------: |
|              /ping              |  [/]   |   [/]   |              Ping              |
|          /simple/price          |  [/]   |   [/]   | SimpleSinglePrice, SimplePrice |
| /simple/supported_vs_currencies |  [/]   |   [/]   |  SimpleSupportedVSCurrencies   |
|           /coins/list           |  [/]   |   [/]   |           CoinsList            |
|          /coins/market          |  [/]   |   [/]   |          CoinsMarket           |
|           /coins/{id}           |  [/]   |         |            CoinsID             |
|       /coins/{id}/history       |  [/]   |         |         CoinsIDHistory         |
|    /coins/{id}/market_chart     |  [/]   |         |       CoinsIDMarketChart       |
|        /events/countries        |  [/]   |         |        EventsCountries         |
|          /events/types          |  [/]   |         |           EventsType           |
|         /exchange_rates         |  [/]   |         |          ExchangeRate          |
|             /global             |  [/]   |         |             Global             |

## Usage

Installation:

Add the following line to your project `go.mod` file.

```
replace github.com/superoo7/go-gecko => github.com/KYRDTeam/go-gecko latest
```

and run 

```
go mod tidy
```

to evaluate `latest` to a concrete tag.

For usage, checkout [Example folder for v3](/_example/v3)

For production, you might need to set time out for httpClient, here's a sample code:

```go
package main

import (
	"net/http"
	"time"

	coingecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG := coingecko.NewClient(httpClient, nil)
	c, err := CG.CoinsList(false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}
```

## Convention

refer to https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c

## License

MIT
