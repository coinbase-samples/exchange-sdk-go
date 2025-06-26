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
	"github.com/coinbase-samples/exchange-sdk-go/orders"
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

	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <order_id> OR %s <product_id>", os.Args[0], os.Args[0])
		log.Fatalf("note: provide either order_id or product_id, but not both")
	}

	// Check if we have exactly one argument
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s <order_id> OR %s <product_id>", os.Args[0], os.Args[0])
		log.Fatalf("note: provide either order_id or product_id, but not both")
	}

	identifier := os.Args[1]

	ordersSvc := orders.NewOrdersService(client)

	request := &orders.ListFillsRequest{}

	// Determine if the identifier is an order_id or product_id
	// This is a simple heuristic - order IDs typically contain hyphens and are longer
	// Product IDs are typically in format like "BTC-USD", "ETH-USD", etc.
	if len(identifier) > 20 && contains(identifier, "-") {
		// Likely an order ID
		request.OrderId = identifier
	} else {
		// Likely a product ID
		request.ProductId = identifier
	}

	response, err := ordersSvc.ListFills(context.Background(), request)
	if err != nil {
		log.Fatalf("unable to list fills: %v", err)
	}

	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling response to JSON: %v", err)
	}
	fmt.Println(string(output))
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr))
}
