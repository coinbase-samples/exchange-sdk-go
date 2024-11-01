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

package products

import (
	"context"
	"fmt"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
	"github.com/coinbase-samples/exchange-sdk-go/utils"
)

type GetProductTradesRequest struct {
	ProductId  string                  `json:"product_id"`
	Pagination *model.PaginationParams `json:"pagination_params"`
}

type GetProductTradesResponse struct {
	ProductTrades []*model.ProductTrades `json:"product_trades"`
}

func (s *productsServiceImpl) GetProductTrades(
	ctx context.Context,
	request *GetProductTradesRequest,
) (*GetProductTradesResponse, error) {

	path := fmt.Sprintf("/products/%s/trades", request.ProductId)

	queryParams := utils.AppendPaginationParams(core.EmptyQueryParams, request.Pagination)

	var productTrades []*model.ProductTrades

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&productTrades,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &GetProductTradesResponse{}, nil
}
