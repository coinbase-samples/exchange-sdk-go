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

package loans

import (
	"context"
	"fmt"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

type GetInterestRateHistoryRequest struct {
	LoanId string `json:"loan_id"`
}

type GetInterestRateHistoryResponse struct {
	RateHistories []*model.RateHistory `json:"rate_histories"`
}

func (s *loansServiceImpl) GetInterestRateHistory(
	ctx context.Context,
	request *GetInterestRateHistoryRequest,
) (*GetInterestRateHistoryResponse, error) {

	path := fmt.Sprintf("/loans/interest/history/%s", request.LoanId)

	var rateHistories []*model.RateHistory

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&rateHistories,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &GetInterestRateHistoryResponse{RateHistories: rateHistories}, nil
}
