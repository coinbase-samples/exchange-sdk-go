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
)

type CreateTravelRuleEntryRequest struct {
	Address           string `json:"address"`
	OriginatorName    string `json:"originator_name"`
	OriginatorCountry string `json:"originator_country"`
}

type CreateTravelRuleEntryResponse struct {
	TravelRule model.TravelRule `json:"travel_rule"`
}

func (s *travelRulesServiceImpl) CreateTravelRuleEntry(
	ctx context.Context,
	request *CreateTravelRuleEntryRequest,
) (*CreateTravelRuleEntryResponse, error) {

	path := "/travel-rules"

	var travelRule model.TravelRule

	if err := core.HttpPost(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&travelRule,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &CreateTravelRuleEntryResponse{TravelRule: travelRule}, nil
}
