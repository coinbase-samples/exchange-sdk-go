package test

import (
	"context"
	"encoding/json"
	"github.com/coinbase-samples/exchange-sdk-go/accounts"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/coinbaseaccounts"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
	"github.com/coinbase-samples/exchange-sdk-go/profiles"
	"os"
	"testing"
	"time"

	"github.com/coinbase-samples/core-go"
)

func TestCreateCryptoAddress(t *testing.T) {
	creds := &credentials.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("EXCHANGE_CREDENTIALS")), creds); err != nil {
		t.Fatalf("unable to deserialize exchange credentials JSON: %v", err)
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		t.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(creds, httpClient)
	profilesSvc := profiles.NewProfilesService(client)
	accountsSvc := accounts.NewAccountsService(client)
	coinbaseAccountsSvc := coinbaseaccounts.NewCoinbaseAccountsService(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	profilesResponse, err := profilesSvc.ListProfiles(ctx, &profiles.ListProfilesRequest{Active: "true"})
	if err != nil {
		t.Fatalf("error fetching profiles: %v", err)
	}
	if profilesResponse == nil || len(*profilesResponse) == 0 {
		t.Fatal("expected non-empty profiles response")
	}
	profileId := (*profilesResponse)[0].Id

	accountsResponse, err := accountsSvc.ListAccounts(ctx, &accounts.ListAccountsRequest{})
	if err != nil {
		t.Fatalf("error fetching accounts: %v", err)
	}
	var accountId string
	for _, account := range *accountsResponse {
		if account.Currency == "BTC" {
			accountId = account.Id
			break
		}
	}
	if accountId == "" {
		t.Fatal("BTC account not found")
	}

	request := &coinbaseaccounts.CreateCryptoAddressRequest{
		AccountId: accountId,
		ProfileId: profileId,
		Network:   "BTC",
	}
	response, err := coinbaseAccountsSvc.CreateCryptoAddress(ctx, request)
	if err != nil {
		t.Fatalf("error creating crypto address: %v", err)
	}

	if response == nil {
		t.Fatal("expected non-nil response")
	}

	if response.Id == "" {
		t.Fatal("expected crypto address ID to be set")
	}

	if response.Address == "" {
		t.Fatal("expected crypto address field to be set")
	}

	if response.Network != "BTC" {
		t.Fatalf("expected network to be BTC, got %v", response.Network)
	}
}
