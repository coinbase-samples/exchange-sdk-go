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
	"github.com/coinbase-samples/exchange-sdk-go/utils"
)

type ListFillsRequest struct {
	OrderId    string                  `json:"order_id"`
	ProductId  string                  `json:"product_id"`
	MarketType string                  `json:"market_type,omitempty"`
	StartDate  string                  `json:"start_date,omitempty"`
	EndDate    string                  `json:"end_date,omitempty"`
	Pagination *model.PaginationParams `json:"pagination,omitempty"`
}

type ListFillsResponse []*model.Fill

func (s *ordersServiceImpl) ListFills(
	ctx context.Context,
	request *ListFillsRequest,
) (*ListFillsResponse, error) {

	path := "/fills"

	var queryParams string
	if len(request.OrderId) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "order_id", request.OrderId)
	}

	if len(request.ProductId) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "product_id", request.ProductId)
	}

	if len(request.MarketType) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "market_type", request.MarketType)
	}

	if len(request.StartDate) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "start_date", request.StartDate)
	}

	if len(request.EndDate) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "end_date", request.EndDate)
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &ListFillsResponse{}

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
