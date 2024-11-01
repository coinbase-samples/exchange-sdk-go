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
	"github.com/coinbase-samples/exchange-sdk-go/utils"
)

type CancelOrderRequest struct {
	OrderId   string `json:"order_id"`
	ProfileId string `json:"profile_id,omitempty"`
	ProductId string `json:"product_id,omitempty"`
}

type CancelOrderResponse model.Description

func (s *ordersServiceImpl) CancelOrder(
	ctx context.Context,
	request *CancelOrderRequest,
) (*CancelOrderResponse, error) {

	path := fmt.Sprintf("/orders/%s", request.OrderId)

	var queryParams string
	if len(request.ProfileId) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "profile_id", request.ProfileId)
	}
	if len(request.ProductId) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "product_id", request.ProductId)
	}

	response := &CancelOrderResponse{}

	if err := core.HttpDelete(ctx, s.client, path, queryParams, client.DefaultSuccessHttpStatusCodes, request, response, s.client.HeadersFunc()); err != nil {
		return nil, err
	}

	return response, nil
}
