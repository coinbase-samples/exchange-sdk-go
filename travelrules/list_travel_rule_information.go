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

package travelrules

import (
	"context"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
	"github.com/coinbase-samples/exchange-sdk-go/utils"
)

type ListTravelRuleInformationRequest struct {
	Address    string                  `json:"address,omitempty"`
	Pagination *model.PaginationParams `json:"pagination_params,omitempty"`
}

type ListTravelRuleInformationResponse struct {
	TravelRules []*model.TravelRule `json:"travel_rules"`
}

func (s *travelRulesServiceImpl) ListTravelRuleInformation(
	ctx context.Context,
	request *ListTravelRuleInformationRequest,
) (*ListTravelRuleInformationResponse, error) {

	path := "/travel-rules"

	var queryParams string
	if len(request.Address) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "address", request.Address)
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	var travelRules []*model.TravelRule

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&travelRules,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &ListTravelRuleInformationResponse{TravelRules: travelRules}, nil
}
