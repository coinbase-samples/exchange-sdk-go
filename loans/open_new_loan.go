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

type OpenNewLoanRequest struct {
	LoanId        string `json:"loan_id,omitempty"`
	Currency      string `json:"currency"`
	NativeAmount  string `json:"native_amount"`
	InterestRate  string `json:"interest_rate"`
	TermStartDate string `json:"term_start_date"`
	TermEndDate   string `json:"term_end_date"`
	ProfileId     string `json:"profile_id"`
}

type OpenNewLoanResponse struct {
	Loan model.Loan `json:"loan"`
}

func (s *loansServiceImpl) OpenNewLoan(
	ctx context.Context,
	request *OpenNewLoanRequest,
) (*OpenNewLoanResponse, error) {

	path := "/loans/open"

	var loan model.Loan

	if err := core.HttpPost(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&loan,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &OpenNewLoanResponse{Loan: loan}, nil
}
