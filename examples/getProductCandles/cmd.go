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

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
	"github.com/coinbase-samples/exchange-sdk-go/products"
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

	if len(os.Args) < 3 {
		log.Fatalf("usage: %s <product_id> <granularity> [start] [end]", os.Args[0])
		log.Fatalf("granularity must be one of: 60, 300, 900, 3600, 21600, 86400")
	}

	productId := os.Args[1]
	granularity := os.Args[2]

	validGranularities := map[string]bool{
		"60":    true, // 1 minute
		"300":   true, // 5 minutes
		"900":   true, // 15 minutes
		"3600":  true, // 1 hour
		"21600": true, // 6 hours
		"86400": true, // 1 day
	}

	if !validGranularities[granularity] {
		log.Fatalf("invalid granularity: %s. Must be one of: 60, 300, 900, 3600, 21600, 86400", granularity)
	}

	productsSvc := products.NewProductsService(client)

	request := &products.GetProductCandlesRequest{
		ProductId:   productId,
		Granularity: granularity,
	}

	if len(os.Args) > 3 {
		request.Start = os.Args[3]
	}
	if len(os.Args) > 4 {
		request.End = os.Args[4]
	}

	response, err := productsSvc.GetProductCandles(context.Background(), request)
	if err != nil {
		log.Fatalf("unable to get product candles: %v", err)
	}

	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling response to JSON: %v", err)
	}
	fmt.Println(string(output))
}
