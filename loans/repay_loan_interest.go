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
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

type RepayLoanInterestRequest struct {
	Idem          string `json:"idem,omitempty"`
	FromProfileId string `json:"from_profile_id"`
	Currency      string `json:"currency"`
	NativeAmount  string `json:"native_amount"`
}

type RepayLoanInterestResponse struct {
	Repayment model.Repayment `json:"repayment"`
}

func (s *loansServiceImpl) RepayLoanInterest(
	ctx context.Context,
	request *RepayLoanInterestRequest,
) (*RepayLoanInterestResponse, error) {

	path := "/loans/repay-interest"

	var repayment model.Repayment

	if err := core.HttpPost(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&repayment,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &RepayLoanInterestResponse{Repayment: repayment}, nil
}
