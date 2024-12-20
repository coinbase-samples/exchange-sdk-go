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

type GetNewLoanPreviewRequest struct {
	Currency     string `json:"currency,omitempty"`
	NativeAmount string `json:"native_amount,omitempty"`
}

type GetNewLoanPreviewResponse struct {
	LoanPreview model.LoanPreview `json:"loan_preview"`
}

func (s *loansServiceImpl) GetNewLoanPreview(
	ctx context.Context,
	request *GetNewLoanPreviewRequest,
) (*GetNewLoanPreviewResponse, error) {

	path := "/loans/loan-preview"

	var queryParams string
	if len(request.Currency) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "currency", request.Currency)
	}

	if len(request.NativeAmount) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "native_amount", request.NativeAmount)
	}

	var loanPreview model.LoanPreview

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&loanPreview,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &GetNewLoanPreviewResponse{}, nil
}
