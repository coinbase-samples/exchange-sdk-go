package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/accounts"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
	"log"
	"os"
)

func main() {
	credentials := &credentials.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("EXCHANGE_CREDENTIALS")), credentials); err != nil {
		log.Fatalf("unable to deserialize exchange credentials JSON: %v", err)
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		log.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(credentials, httpClient)

	accountsSvc := accounts.NewAccountsService(client)

	request := &accounts.ListAccountsRequest{}

	response, err := accountsSvc.ListAccounts(context.Background(), request)
	if err != nil {
		log.Fatalf("unable to list accounts: %v", err)
	}

	jsonResponse, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling response to JSON:", err)
		return
	}
	fmt.Println(string(jsonResponse))
}
