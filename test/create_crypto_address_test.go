package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/coinbaseaccounts"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
	"github.com/coinbase-samples/exchange-sdk-go/profiles"
)

func TestCreateCryptoAddress(t *testing.T) {
	credentials, err := credentials.ReadEnvCredentials("EXCHANGE_CREDENTIALS")
	if err != nil {
		panic(fmt.Sprintf("unable to read exchange credentials: %v", err))
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		t.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(credentials, httpClient)
	profilesSvc := profiles.NewProfilesService(client)
	coinbaseAccountsSvc := coinbaseaccounts.NewCoinbaseAccountsService(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	profilesResponse, err := profilesSvc.ListProfiles(ctx, &profiles.ListProfilesRequest{})
	if err != nil {
		t.Fatalf("error fetching profiles: %v", err)
	}
	if profilesResponse == nil || len(profilesResponse.Profiles) == 0 {
		t.Fatal("expected non-empty profiles response")
	}
	profileId := profilesResponse.Profiles[0].Id

	coinbaseAccountsResponse, err := coinbaseAccountsSvc.ListCoinbaseWallets(ctx, &coinbaseaccounts.ListCoinbaseWalletsRequest{})
	if err != nil {
		t.Fatalf("error fetching accounts: %v", err)
	}
	var coinbaseWalletId string
	for _, coinbase_wallets := range coinbaseAccountsResponse.CoinbaseWallets {
		if coinbase_wallets.Name == "BTC Wallet" {
			coinbaseWalletId = coinbase_wallets.Id
			break
		}
	}
	if coinbaseWalletId == "" {
		t.Fatal("BTC account not found")
	}

	request := &coinbaseaccounts.CreateCryptoAddressRequest{
		AccountId: coinbaseWalletId,
		ProfileId: profileId,
		Network:   "bitcoin",
	}
	response, err := coinbaseAccountsSvc.CreateCryptoAddress(ctx, request)
	if err != nil {
		t.Fatalf("error creating crypto address: %v", err)
	}

	if response == nil {
		t.Fatal("expected non-nil response")
	}

	cryptoAddress := response.CryptoAddress
	if cryptoAddress.Id == "" {
		t.Fatal("expected crypto address ID to be set")
	}

	if cryptoAddress.Address == "" {
		t.Fatal("expected crypto address field to be set")
	}

	if cryptoAddress.Network != "bitcoin" {
		t.Fatalf("expected network to be bitcoin, got %v", cryptoAddress.Network)
	}
}
