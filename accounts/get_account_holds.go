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

package accounts

import (
	"context"
	"fmt"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
	"github.com/coinbase-samples/exchange-sdk-go/utils"
)

type GetAccountHoldsRequest struct {
	AccountId  string                  `json:"account_id"`
	Pagination *model.PaginationParams `json:"pagination_params"`
}

type GetAccountHoldsResponse []*model.AccountHold

func (s *accountsServiceImpl) GetAccountHolds(
	ctx context.Context,
	request *GetAccountHoldsRequest,
) (*GetAccountHoldsResponse, error) {

	path := fmt.Sprintf("/accounts/%s/holds", request.AccountId)

	queryParams := utils.AppendPaginationParams(core.EmptyQueryParams, request.Pagination)

	response := &GetAccountHoldsResponse{}

	if err := core.HttpGet(ctx, s.client, path, queryParams, client.DefaultSuccessHttpStatusCodes, request, response, s.client.HeadersFunc()); err != nil {
		return nil, err
	}

	return response, nil
}