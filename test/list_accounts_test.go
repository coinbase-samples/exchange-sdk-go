package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/accounts"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
)

func TestListAccounts(t *testing.T) {
	credentials, err := credentials.ReadEnvCredentials("EXCHANGE_CREDENTIALS")
	if err != nil {
		panic(fmt.Sprintf("unable to read exchange credentials: %v", err))
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		t.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(credentials, httpClient)
	accountsSvc := accounts.NewAccountsService(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := accountsSvc.ListAccounts(ctx, &accounts.ListAccountsRequest{})
	if err != nil {
		t.Fatalf("error listing accounts: %v", err)
	}

	if response == nil || len(response.Accounts) == 0 {
		t.Fatal("expected at least one account in list accounts response")
	}

	firstAccount := response.Accounts[0]
	if firstAccount.Id == "" || firstAccount.Currency == "" || firstAccount.Balance == "" || firstAccount.ProfileId == "" {
		t.Fatal("expected all necessary fields to be set in first account")
	}
}
