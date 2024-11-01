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
	"strings"
)

type ListOrdersRequest struct {
	ProfileId  string                  `json:"profile_id"`
	ProductId  string                  `json:"product_id"`
	SortedBy   string                  `json:"sorted_by,omitempty"`
	Sorting    string                  `json:"sorting,omitempty"`
	StartDate  string                  `json:"start_date,omitempty"`
	EndDate    string                  `json:"end_date,omitempty"`
	Status     []string                `json:"status"`
	MarketType string                  `json:"market_type,omitempty"`
	Pagination *model.PaginationParams `json:"pagination,omitempty"`
}

type ListOrdersResponse []*model.Order

func (s *ordersServiceImpl) ListOrders(
	ctx context.Context,
	request *ListOrdersRequest,
) (*ListOrdersResponse, error) {

	path := "/orders"

	var queryParams string
	if len(request.ProfileId) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "profile_id", request.ProfileId)
	}
	if len(request.ProductId) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "product_id", request.ProductId)
	}
	if len(request.SortedBy) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "sorted_by", request.SortedBy)
	}
	if len(request.Sorting) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "sorting", request.Sorting)
	}
	if len(request.StartDate) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "start_date", request.StartDate)
	}
	if len(request.EndDate) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "end_date", request.EndDate)
	}
	if len(request.Status) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "status", strings.Join(request.Status, ","))
	}
	if len(request.MarketType) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "market_type", request.MarketType)
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &ListOrdersResponse{}

	if err := core.HttpGet(ctx, s.client, path, queryParams, client.DefaultSuccessHttpStatusCodes, request, response, s.client.HeadersFunc()); err != nil {
		return nil, err
	}

	return response, nil
}
