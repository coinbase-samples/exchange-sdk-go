package test

import (
	"context"
	"encoding/json"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
	"github.com/coinbase-samples/exchange-sdk-go/products"
	"os"
	"testing"
	"time"

	"github.com/coinbase-samples/core-go"
)

func TestGetProductTicker(t *testing.T) {
	creds := &credentials.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("EXCHANGE_CREDENTIALS")), creds); err != nil {
		t.Fatalf("unable to deserialize exchange credentials JSON: %v", err)
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		t.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(creds, httpClient)
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

	if response.TradeId == 0 {
		t.Fatal("expected trade ID to be set")
	}

	if response.Price == "" {
		t.Fatal("expected price field to be set")
	}

	if response.Bid == "" {
		t.Fatal("expected bid field to be set")
	}

	if response.Ask == "" {
		t.Fatal("expected ask field to be set")
	}
}
