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

type GetProductBookRequest struct {
	ProductId string `json:"product_id"`
	Level     string `json:"level,omitempty"`
}

type GetProductBookResponse model.ProductBook

func (s *productsServiceImpl) GetProductBook(
	ctx context.Context,
	request *GetProductBookRequest,
) (*GetProductBookResponse, error) {

	path := fmt.Sprintf("/products/%s/book", request.ProductId)

	var queryParams string
	if len(request.Level) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "level", request.Level)
	}

	response := &GetProductBookResponse{}

	if err := core.HttpGet(ctx, s.client, path, queryParams, client.DefaultSuccessHttpStatusCodes, request, response, s.client.HeadersFunc()); err != nil {
		return nil, err
	}

	return response, nil
}
