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
	"github.com/coinbase-samples/exchange-sdk-go/client"
)

type ReportsService interface {
	ListReports(ctx context.Context, request *ListReportsRequest) (*ListReportsResponse, error)
	GetReport(ctx context.Context, request *GetReportRequest) (*GetReportResponse, error)
	CreateReport(ctx context.Context, request *CreateReportRequest) (*CreateReportResponse, error)
}

func NewReportsService(c client.RestClient) ReportsService {
	return &reportsServiceImpl{client: c}
}

type reportsServiceImpl struct {
	client client.RestClient
}
