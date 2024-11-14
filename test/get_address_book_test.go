package test

import (
	"context"
	"encoding/json"
	"github.com/coinbase-samples/exchange-sdk-go/addressbook"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
	"os"
	"testing"
	"time"

	"github.com/coinbase-samples/core-go"
)

func TestGetAddressBook(t *testing.T) {
	creds := &credentials.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("EXCHANGE_CREDENTIALS")), creds); err != nil {
		t.Fatalf("unable to deserialize exchange credentials JSON: %v", err)
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		t.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(creds, httpClient)
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

	if len(*response) == 0 {
		t.Fatal("expected addresses in get address book")
	}

	firstEntry := (*response)[0]
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
