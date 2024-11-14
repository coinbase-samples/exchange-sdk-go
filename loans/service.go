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
	"github.com/coinbase-samples/exchange-sdk-go/client"
)

type LoansService interface {
	ListLoans(ctx context.Context, request *ListLoansRequest) (*ListLoansResponse, error)
	ListLoanAssets(ctx context.Context, request *ListLoanAssetsRequest) (*ListLoanAssetsResponse, error)
	ListNewLoanOptions(ctx context.Context, request *ListNewLoanOptionsRequest) (*ListNewLoanOptionsResponse, error)
	ListInterestSummaries(ctx context.Context, request *ListInterestSummariesRequest) (*ListInterestSummariesResponse, error)
	GetInterestRateHistory(ctx context.Context, request *GetInterestRateHistoryRequest) (*GetInterestRateHistoryResponse, error)
	GetInterestCharges(ctx context.Context, request *GetInterestChargesRequest) (*GetInterestChargesResponse, error)
	GetLendingOverview(ctx context.Context, request *GetLendingOverviewRequest) (*GetLendingOverviewResponse, error)
	GetNewLoanPreview(ctx context.Context, request *GetNewLoanPreviewRequest) (*GetNewLoanPreviewResponse, error)
	OpenNewLoan(ctx context.Context, request *OpenNewLoanRequest) (*OpenNewLoanResponse, error)
	RepayLoanInterest(ctx context.Context, request *RepayLoanInterestRequest) (*RepayLoanInterestResponse, error)
	RepayLoanPrincipal(ctx context.Context, request *RepayLoanPrincipalRequest) (*RepayLoanPrincipalResponse, error)
	GetPrincipalRepaymentPreview(ctx context.Context, request *GetPrincipalRepaymentPreviewRequest) (*GetPrincipalRepaymentPreviewResponse, error)
}

func NewLoansService(c client.RestClient) LoansService {
	return &loansServiceImpl{client: c}
}

type loansServiceImpl struct {
	client client.RestClient
}
