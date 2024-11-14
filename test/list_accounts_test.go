package test

import (
	"context"
	"encoding/json"
	"github.com/coinbase-samples/exchange-sdk-go/accounts"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
	"os"
	"testing"
	"time"

	"github.com/coinbase-samples/core-go"
)

func TestListAccounts(t *testing.T) {
	creds := &credentials.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("EXCHANGE_CREDENTIALS")), creds); err != nil {
		t.Fatalf("unable to deserialize exchange credentials JSON: %v", err)
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		t.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(creds, httpClient)
	accountsSvc := accounts.NewAccountsService(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := accountsSvc.ListAccounts(ctx, &accounts.ListAccountsRequest{})
	if err != nil {
		t.Fatalf("error listing accounts: %v", err)
	}

	if response == nil {
		t.Fatal("expected non-nil response")
	}

	if len(*response) == 0 {
		t.Fatal("expected at least one account in list accounts response")
	}

	firstAccount := (*response)[0]
	if firstAccount.Id == "" {
		t.Fatal("expected account ID to be set")
	}
	if firstAccount.Currency == "" {
		t.Fatal("expected currency to be set")
	}
	if firstAccount.Balance == "" {
		t.Fatal("expected balance to be set")
	}
	if firstAccount.ProfileId == "" {
		t.Fatal("expected profile ID to be set")
	}
}
