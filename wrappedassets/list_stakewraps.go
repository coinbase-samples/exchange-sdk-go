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

package wrappedassets

import (
	"context"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
	"github.com/coinbase-samples/exchange-sdk-go/utils"
)

type ListStakewrapsRequest struct {
	From       string                  `json:"from"`
	To         string                  `json:"to"`
	Pagination *model.PaginationParams `json:"pagination,omitempty"`
}

type ListStakewrapsResponse []*model.StakeWrap

func (s *wrappedAssetsServiceImpl) ListStakewraps(
	ctx context.Context,
	request *ListStakewrapsRequest,
) (*ListStakewrapsResponse, error) {

	path := "/wrapped-assets/stake-wrap"

	var queryParams string
	if len(request.From) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "from", request.From)
	}

	if len(request.To) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "to", request.To)
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &ListStakewrapsResponse{}

	if err := core.HttpGet(ctx, s.client, path, queryParams, client.DefaultSuccessHttpStatusCodes, request, response, s.client.HeadersFunc()); err != nil {
		return nil, err
	}

	return response, nil
}
