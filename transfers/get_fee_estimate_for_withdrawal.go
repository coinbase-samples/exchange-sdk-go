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
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

type GetFeeEstimateForWithdrawalRequest struct {
	Currency      string `json:"currency"`
	CryptoAddress string `json:"crypto_address"`
	Network       string `json:"network"`
}

type GetFeeEstimateForWithdrawalResponse struct {
	FeeEstimate model.FeeEstimate `json:"fee_estimate"`
}

func (s *transfersServiceImpl) GetFeeEstimateForWithdrawal(
	ctx context.Context,
	request *GetFeeEstimateForWithdrawalRequest,
) (*GetFeeEstimateForWithdrawalResponse, error) {

	path := "/withdrawals/fee-estimate"

	var queryParams string
	if len(request.Currency) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "currency", request.Currency)
	}

	if len(request.CryptoAddress) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "crypto_address", request.CryptoAddress)
	}

	if len(request.Network) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "network", request.Network)
	}

	var feeEstimate model.FeeEstimate

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&feeEstimate,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &GetFeeEstimateForWithdrawalResponse{FeeEstimate: feeEstimate}, nil
}
