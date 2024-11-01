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
	"github.com/coinbase-samples/exchange-sdk-go/utils"
)

type GetProductCandlesRequest struct {
	ProductId   string `json:"product_id"`
	Granularity string `json:"granularity,omitempty"`
	Start       string `json:"start"`
	End         string `json:"end"`
}

type GetProductCandlesResponse [][]float64

func (s *productsServiceImpl) GetProductCandles(
	ctx context.Context,
	request *GetProductCandlesRequest,
) (*GetProductCandlesResponse, error) {

	path := fmt.Sprintf("/products/%s/candles", request.ProductId)

	var queryParams string
	if len(request.Granularity) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "granularity", request.Granularity)
	}

	if len(request.Start) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "start", request.Start)
	}

	if len(request.End) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "end", request.End)
	}

	response := &GetProductCandlesResponse{}

	if err := core.HttpGet(ctx, s.client, path, queryParams, client.DefaultSuccessHttpStatusCodes, request, response, s.client.HeadersFunc()); err != nil {
		return nil, err
	}

	return response, nil
}
