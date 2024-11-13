package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/accounts"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
)

func main() {
	credentials, err := credentials.ReadEnvCredentials("EXCHANGE_CREDENTIALS")
	if err != nil {
		panic(fmt.Sprintf("unable to read exchange credentials: %v", err))
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		panic(fmt.Sprintf("unable to load default http client: %v", err))
	}

	client := client.NewRestClient(credentials, httpClient)

	accountsSvc := accounts.NewAccountsService(client)

	request := &accounts.ListAccountsRequest{}

	response, err := accountsSvc.ListAccounts(context.Background(), request)
	if err != nil {
		panic(fmt.Sprintf("unable to list accounts: %v", err))
	}

	jsonResponse, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("error marshaling response to JSON: %v", err))
	}
	fmt.Println(string(jsonResponse))
}
