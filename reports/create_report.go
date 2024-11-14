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

package reports

import (
	"context"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

type CreateReportRequest struct {
	Type           string                  `json:"type"`
	Year           string                  `json:"year,omitempty"`
	Format         string                  `json:"format,omitempty"`
	Email          string                  `json:"email,omitempty"`
	ProfileId      string                  `json:"profile_id,omitempty"`
	Balance        *model.BalanceParams    `json:"balance,omitempty"`
	GroupByProfile bool                    `json:"group_by_profile,omitempty"`
	Fills          *model.FillsParams      `json:"fills,omitempty"`
	Account        *model.AccountParams    `json:"account,omitempty"`
	OtcFills       *model.OtcFillsParams   `json:"otc-fills,omitempty"`
	TaxInvoice     *model.TaxInvoiceParams `json:"tax-invoice,omitempty"`
	RfqFills       *model.RfqFillsParams   `json:"rfq-fills,omitempty"`
}

type CreateReportResponse model.ReportResponse

func (s *reportsServiceImpl) CreateReport(
	ctx context.Context,
	request *CreateReportRequest,
) (*CreateReportResponse, error) {

	path := "/reports"

	response := &CreateReportResponse{}

	if err := core.HttpPost(ctx, s.client, path, core.EmptyQueryParams, client.DefaultSuccessHttpStatusCodes, request, response, s.client.HeadersFunc()); err != nil {
		return nil, err
	}

	return response, nil
}
