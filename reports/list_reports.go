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
	"github.com/coinbase-samples/exchange-sdk-go/utils"
)

type ListReportsRequest struct {
	ProfileId     string                  `json:"profile_id,omitempty"`
	Type          string                  `json:"type,omitempty"`
	IgnoreExpired string                  `json:"ignore_expired"`
	Pagination    *model.PaginationParams `json:"pagination_params,omitempty"`
}

type ListReportsResponse struct {
	Reports []*model.Report `json:"reports"`
}

func (s *reportsServiceImpl) ListReports(
	ctx context.Context,
	request *ListReportsRequest,
) (*ListReportsResponse, error) {

	path := "/reports"

	var queryParams string
	if len(request.ProfileId) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "profile_id", request.ProfileId)
	}

	if len(request.Type) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "type", request.Type)
	}

	if len(request.IgnoreExpired) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "ignore_expired", request.IgnoreExpired)
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	var reports []*model.Report

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&reports,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &ListReportsResponse{Reports: reports}, nil
}
