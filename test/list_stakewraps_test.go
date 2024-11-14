package test

import (
	"context"
	"encoding/json"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
	"github.com/coinbase-samples/exchange-sdk-go/wrappedassets"
	"os"
	"testing"
	"time"

	"github.com/coinbase-samples/core-go"
)

func TestListStakewraps(t *testing.T) {
	creds := &credentials.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("EXCHANGE_CREDENTIALS")), creds); err != nil {
		t.Fatalf("unable to deserialize exchange credentials JSON: %v", err)
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		t.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(creds, httpClient)
	wrappedAssetsSvc := wrappedassets.NewWrappedAssetsService(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := &wrappedassets.ListStakewrapsRequest{
		From: "2024-01-01T00:00:00Z",
		To:   "2024-12-31T23:59:59Z",
	}

	response, err := wrappedAssetsSvc.ListStakewraps(ctx, request)
	if err != nil {
		t.Fatalf("error listing stakewraps: %v", err)
	}

	if response == nil {
		t.Fatal("expected non-nil response")
	}

	if len(*response) == 0 {
		t.Fatal("expected at least one stakewrap in response")
	}

	firstStakewrap := (*response)[0]
	if firstStakewrap.Id == "" {
		t.Fatal("expected stakewrap ID to be set")
	}
	if firstStakewrap.FromAmount == "" {
		t.Fatal("expected from_amount field to be set")
	}
	if firstStakewrap.ToAmount == "" {
		t.Fatal("expected to_amount field to be set")
	}
	if firstStakewrap.FromCurrency == "" {
		t.Fatal("expected from_currency field to be set")
	}
	if firstStakewrap.ToCurrency == "" {
		t.Fatal("expected to_currency field to be set")
	}
}
