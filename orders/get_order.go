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
	"fmt"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

type GetOrderRequest struct {
	OrderId    string `json:"order_id"`
	MarketType string `json:"market_type,omitempty"`
}

type GetOrderResponse struct {
	Order model.Order `json:"order"`
}

func (s *ordersServiceImpl) GetOrder(
	ctx context.Context,
	request *GetOrderRequest,
) (*GetOrderResponse, error) {

	path := fmt.Sprintf("/orders/%s", request.OrderId)

	var queryParams string
	if len(request.MarketType) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "market_type", request.MarketType)
	}

	var order model.Order

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&order,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &GetOrderResponse{Order: order}, nil
}
