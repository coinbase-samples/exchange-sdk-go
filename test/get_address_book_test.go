package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/addressbook"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
)

func TestGetAddressBook(t *testing.T) {
	credentials, err := credentials.ReadEnvCredentials("EXCHANGE_CREDENTIALS")
	if err != nil {
		panic(fmt.Sprintf("unable to read exchange credentials: %v", err))
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		t.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(credentials, httpClient)
	addressBookSvc := addressbook.NewAddressBookService(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := &addressbook.GetAddressBookRequest{}
	response, err := addressBookSvc.GetAddressBook(ctx, request)
	if err != nil {
		t.Fatalf("error fetching address book: %v", err)
	}

	if response == nil {
		t.Fatal("expected non-nil response")
	}

	if len(response.AddressBooks) == 0 {
		t.Fatal("expected at least one address book entry in response")
	}

	firstEntry := response.AddressBooks[0]
	if firstEntry.Id == "" {
		t.Fatal("expected address book entry ID to be set")
	}
	if firstEntry.Address == "" {
		t.Fatal("expected address field to be set")
	}
	if firstEntry.Currency == "" {
		t.Fatal("expected currency field to be set")
	}
}
