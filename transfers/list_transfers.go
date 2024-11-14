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

package transfers

import (
	"context"
	"fmt"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
	"github.com/coinbase-samples/exchange-sdk-go/utils"
)

type ListTransfersRequest struct {
	ProfileId      string `json:"profile_id"`
	Before         string `json:"before,omitempty"`
	After          string `json:"after,omitempty"`
	Limit          int64  `json:"limit,omitempty"`
	Type           string `json:"type,omitempty"`
	CurrencyType   string `json:"currency_type,omitempty"`
	TransferReason string `json:"transfer_reason,omitempty"`
	Currency       string `json:"currency,omitempty"`
}

type ListTransfersResponse []*model.CoinbaseWallet

func (s *transfersServiceImpl) ListTransfers(
	ctx context.Context,
	request *ListTransfersRequest,
) (*ListTransfersResponse, error) {

	path := "/transfers"

	var queryParams string
	if len(request.ProfileId) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "profile_id", request.ProfileId)
	}
	if len(request.Before) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "before", request.Before)
	}
	if len(request.After) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "after", request.After)
	}
	if request.Limit > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "limit", fmt.Sprintf("%d", request.Limit))
	}
	if len(request.Type) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "type", request.Type)
	}
	if len(request.CurrencyType) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "currency_type", request.CurrencyType)
	}
	if len(request.TransferReason) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "transfer_reason", request.TransferReason)
	}
	if len(request.Currency) > 0 {
		queryParams = utils.AppendQueryParam(queryParams, "currency", request.Currency)
	}

	response := &ListTransfersResponse{}

	if err := core.HttpGet(ctx, s.client, path, queryParams, client.DefaultSuccessHttpStatusCodes, request, response, s.client.HeadersFunc()); err != nil {
		return nil, err
	}

	return response, nil
}
