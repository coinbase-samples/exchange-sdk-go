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
	"github.com/coinbase-samples/exchange-sdk-go/conversions"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
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

	if len(os.Args) < 5 {
		log.Fatalf("usage: %s <profile_id> <from_currency> <to_currency> <amount>", os.Args[0])
	}

	profileId := os.Args[1]
	fromCurrency := os.Args[2]
	toCurrency := os.Args[3]
	amount := os.Args[4]

	nonce := fmt.Sprintf("%d", time.Now().UnixNano())

	conversionsSvc := conversions.NewConversionsService(client)

	// Converts funds from from currency to to currency. Funds are converted on the from account in the profile_id profile.
	request := &conversions.CreateConversionRequest{
		ProfileId: profileId,
		From:      fromCurrency,
		To:        toCurrency,
		Amount:    amount,
		Nonce:     nonce,
	}

	response, err := conversionsSvc.CreateConversion(context.Background(), request)
	if err != nil {
		log.Fatalf("unable to create conversion: %v", err)
	}

	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling response to JSON: %v", err)
	}
	fmt.Println(string(output))
}
