package main

import (
	"fmt"
	"time"

	"github.com/ianberinger/stockfighter/api"
)

const (
	apiKey string = ""
)

func main() {
	//start a new API instance with test values, use api.NewInstance() for real use
	i := api.NewTestInstance()
	i.SetAPIKey(apiKey) //set your API-Key if you want to test the authorization-only API calls.

	//see if the API is up
	prettyPrint("API is up:", i.Heartbeat())

	//see if the venue is up
	prettyPrint("venue is up:", i.VenueHeartbeat())

	//available symbols at the current venue
	prettyPrint("available stocks:", i.AvailableStocks())

	//getting the current orderbook
	prettyPrint("current orderbook:", i.Orderbook())

	//getting a quote
	quote := i.Quote()
	prettyPrint("quote:", quote)

	//websocket based calls
	/*
		ticker := i.QuotesForVenue()
		for tick := range ticker {
			prettyPrint("tick:", ticker)
		}
	*/

	/// the following calls need a valid api key

	//make an order
	order := i.NewOrder(quote.LastPrice, 100, api.Buy, api.Limit)
	if i.Err() != nil {
		prettyPrint("we got an error:", i.Err())
	} else {
		prettyPrint("created order:", order)
		//see status of order
		fmt.Println("waiting for 5 seconds before querying order status")
		time.Sleep(5 * time.Second)
		prettyPrint("status of order:", i.OrderStatus(order.ID))

		//cancel order
		prettyPrint("canceled order:", i.CancelOrder(order.ID))
	}
}

func prettyPrint(description, v interface{}) {
	fmt.Printf("%s %+v\n", description, v)
}
