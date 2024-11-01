package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
	"github.com/coinbase-samples/exchange-sdk-go/products"
)

func TestGetProductTicker(t *testing.T) {
	credentials, err := credentials.ReadEnvCredentials("EXCHANGE_CREDENTIALS")
	if err != nil {
		panic(fmt.Sprintf("unable to read exchange credentials: %v", err))
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		t.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(credentials, httpClient)
	productsSvc := products.NewProductsService(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := &products.GetProductTickerRequest{
		ProductId: "BTC-USD",
	}

	response, err := productsSvc.GetProductTicker(ctx, request)
	if err != nil {
		t.Fatalf("error fetching product ticker: %v", err)
	}

	if response == nil {
		t.Fatal("expected non-nil response")
	}

	productTicker := response.ProductTicker
	if productTicker.TradeId == 0 {
		t.Fatal("expected trade ID to be set")
	}
	if productTicker.Price == "" {
		t.Fatal("expected price field to be set")
	}
	if productTicker.Bid == "" {
		t.Fatal("expected bid field to be set")
	}
	if productTicker.Ask == "" {
		t.Fatal("expected ask field to be set")
	}
}
