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

type GetAccountLedgerRequest struct {
	AccountId  string                  `json:"account_id"`
	StartDate  string                  `json:"start_date,omitempty"`
	EndDate    string                  `json:"end_date,omitempty"`
	Pagination *model.PaginationParams `json:"pagination_params"`
}

type GetAccountLedgerResponse struct {
	AccountLedgers []*model.AccountLedger `json:"account_ledgers"`
}

func (s *accountsServiceImpl) GetAccountLedger(
	ctx context.Context,
	request *GetAccountLedgerRequest,
) (*GetAccountLedgerResponse, error) {

	path := fmt.Sprintf("/accounts/%s/ledger", request.AccountId)

	var queryParams string

	if request.StartDate != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "start_date", request.StartDate)
	}
	if request.EndDate != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "end_date", request.EndDate)
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	var accountLedgers []*model.AccountLedger

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&accountLedgers,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &GetAccountLedgerResponse{AccountLedgers: accountLedgers}, nil
}
