/**
 * Copyright 2024-present Coinbase Global, Inc.
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

package orders

import (
	"context"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

type CreateOrderRequest struct {
	ProfileId      string `json:"profile_id,omitempty"`
	Type           string `json:"type"`
	Side           string `json:"side"`
	ProductId      string `json:"product_id"`
	Stp            string `json:"stp,omitempty"`
	Stop           string `json:"stop,omitempty"`
	StopPrice      string `json:"stop_price,omitempty"`
	Price          string `json:"price,omitempty"`
	Size           string `json:"size,omitempty"`
	Funds          string `json:"funds,omitempty"`
	TimeInForce    string `json:"time_in_force,omitempty"`
	CancelAfter    string `json:"cancel_after,omitempty"`
	PostOnly       bool   `json:"post_only,omitempty"`
	ClientOid      string `json:"client_oid,omitempty"`
	MaxFloor       string `json:"max_floor,omitempty"`
	StopLimitPrice string `json:"stop_limit_price,omitempty"`
}

type CreateOrderResponse struct {
	Order model.Order `json:"order"`
}

func (s *ordersServiceImpl) CreateOrder(
	ctx context.Context,
	request *CreateOrderRequest,
) (*CreateOrderResponse, error) {

	path := "/orders"

	var order model.Order

	if err := core.HttpPost(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&order,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &CreateOrderResponse{Order: order}, nil
}
