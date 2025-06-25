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
	"strings"

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

	if len(os.Args) < 7 {
		log.Fatalf("usage: %s <type> <side> <product_id> <stp> <price> <size> <time_in_force>", os.Args[0])
		log.Fatalf("note: price is only required when type is not 'market'")
	}

	orderType := os.Args[1]
	side := os.Args[2]
	productId := os.Args[3]
	stp := os.Args[4]
	price := os.Args[5]
	size := os.Args[6]
	timeInForce := os.Args[7]

	if strings.ToLower(orderType) != "market" && price == "" {
		log.Fatalf("price is required when order type is not 'market'")
	}

	ordersSvc := orders.NewOrdersService(client)

	request := &orders.CreateOrderRequest{
		Type:        orderType,
		Side:        side,
		ProductId:   productId,
		Stp:         stp,
		Size:        size,
		TimeInForce: timeInForce,
	}

	if price != "" {
		request.Price = price
	}

	response, err := ordersSvc.CreateOrder(context.Background(), request)
	if err != nil {
		log.Fatalf("unable to create order: %v", err)
	}

	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling response to JSON: %v", err)
	}
	fmt.Println(string(output))
}
