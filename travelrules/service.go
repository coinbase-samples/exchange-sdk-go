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
	"github.com/coinbase-samples/exchange-sdk-go/client"
)

type TravelRulesService interface {
	ListTravelRuleInformation(ctx context.Context, request *ListTravelRuleInformationRequest) (*ListTravelRuleInformationResponse, error)
	CreateTravelRuleEntry(ctx context.Context, request *CreateTravelRuleEntryRequest) (*CreateTravelRuleEntryResponse, error)
	DeleteTravelRuleEntry(ctx context.Context, request *DeleteTravelRuleEntryRequest) (*DeleteTravelRuleEntryResponse, error)
}

func NewTravelRulesService(c client.RestClient) TravelRulesService {
	return &travelRulesServiceImpl{client: c}
}

type travelRulesServiceImpl struct {
	client client.RestClient
}
