# Exchange Go SDK README

[![GoDoc](https://godoc.org/github.com/coinbase-samples/exchange-sdk-go?status.svg)](https://godoc.org/github.com/coinbase-samples/exchange-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/coinbase-samples/exchange-sdk-go)](https://goreportcard.com/report/coinbase-samples/exchange-sdk-go)

## Overview

The *Exchange Go SDK* is a sample library that demonstrates the structure of a [Coinbase Exchange](https://exchange.coinbase.com/) driver for
the [REST APIs](https://docs.cdp.coinbase.com/exchange/reference).

## License

The *Exchange Go SDK* sample library is free and open source and released under the [Apache License, Version 2.0](LICENSE).

The application and code are only available for demonstration purposes.

## Usage

To use the *Exchange Go SDK*, initialize the [Credentials](credentials/credentials.go) struct and create a new client. The Credentials struct is JSON
enabled. Ensure that Exchange API credentials are stored in a secure manner.

```
credentials := &credentials.Credentials{}
if err := json.Unmarshal([]byte(os.Getenv("EXCHANGE_CREDENTIALS")), credentials); err != nil {
    panic(fmt.Sprintf("unable to deserialize exchange credentials JSON: %v", err))
}

httpClient, err := core.DefaultHttpClient()
if err != nil {
    panic(fmt.Sprintf("unable to load default http client: %v", err))
}

client := client.NewRestClient(credentials, httpClient)
```

There are convenience functions to read the credentials as an environment variable (exchange.ReadEnvCredentials) and to deserialize the JSON structure (exchange.UnmarshalCredentials) if pulled from a different source. The JSON format expected by both is:

```
export EXCHANGE_CREDENTIALS='{
    "apiKey":"",
    "passphrase":"",
    "signingKey":""
}'
```

This can be set from the command 

Coinbase Exchange API credentials can be created in the Exchange web console under Settings -> APIs.

## Accessing the API

Once the client is initialized, make the desired call. For example, to [list accounts](accounts/list_accounts.go),
pass in the request object, check for an error, and if nil, process the response.


```
accountsSvc := accounts.NewAccountsService(client)
request := &accounts.ListAccountsRequest{}

response, err := accountsSvc.ListAccounts(context.Background(), request)
if err != nil {
    panic(fmt.Sprintf("unable to list accounts: %v", err))
}

// Print the JSON-formatted response
jsonResponse, err := json.MarshalIndent(response, "", "  ")
if err != nil {
    panic(fmt.Sprintf("error marshaling response to JSON: %v", err))
}
fmt.Println(string(jsonResponse))
```

In the example above, ListAccounts is accessed through an accounts service, which uses the client to manage requests. This modular approach allows for additional services to be added and structured as separate modules with their respective implementations.


## Build

To build the sample library, ensure that [Go](https://go.dev/) 1.19+ is installed and then run:

```bash
go build ./...
```
