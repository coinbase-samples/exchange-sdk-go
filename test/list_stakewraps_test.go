package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
	"github.com/coinbase-samples/exchange-sdk-go/wrappedassets"
)

func TestListStakewraps(t *testing.T) {
	credentials, err := credentials.ReadEnvCredentials("EXCHANGE_CREDENTIALS")
	if err != nil {
		panic(fmt.Sprintf("unable to read exchange credentials: %v", err))
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		t.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(credentials, httpClient)
	wrappedAssetsSvc := wrappedassets.NewWrappedAssetsService(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := &wrappedassets.ListStakewrapsRequest{}

	response, err := wrappedAssetsSvc.ListStakewraps(ctx, request)
	if err != nil {
		t.Fatalf("error listing stakewraps: %v", err)
	}

	if response == nil {
		t.Fatal("expected non-nil response")
	}
	if len(response.Stakewraps) == 0 {
		t.Fatal("expected at least one stakewrap in response")
	}

	firstStakewrap := response.Stakewraps[0]
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
