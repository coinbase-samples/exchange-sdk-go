/**
 * Copyright 2025-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
	"github.com/coinbase-samples/exchange-sdk-go/transfers"
)

func main() {
	credentials, err := credentials.ReadEnvCredentials("EXCHANGE_CREDENTIALS")
	if err != nil {
		log.Fatalf("unable to read credentials from environment: %v", err)
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		log.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(credentials, httpClient)

	if len(os.Args) < 6 {
		log.Fatalf("usage: %s <profile_id> <amount> <currency> <crypto_address> <network> [destination_tag]", os.Args[0])
	}

	profileId := os.Args[1]
	amount := os.Args[2]
	currency := os.Args[3]
	cryptoAddress := os.Args[4]
	network := os.Args[5]

	nonce := int32(time.Now().UnixNano())

	transfersSvc := transfers.NewTransfersService(client)

	request := &transfers.WithdrawToCryptoAddressRequest{
		ProfileId:     profileId,
		Amount:        amount,
		Currency:      currency,
		CryptoAddress: cryptoAddress,
		Nonce:         nonce,
		Network:       network,
	}

	if len(os.Args) > 6 {
		request.DestinationTag = os.Args[6]
	}

	response, err := transfersSvc.WithdrawToCryptoAddress(context.Background(), request)
	if err != nil {
		log.Fatalf("unable to withdraw to crypto address: %v", err)
	}

	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling response to JSON: %v", err)
	}
	fmt.Println(string(output))
}
