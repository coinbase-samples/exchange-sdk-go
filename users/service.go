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

package users

import (
	"context"
	"github.com/coinbase-samples/exchange-sdk-go/client"
)

type UsersService interface {
	GetUserExchangeLimits(ctx context.Context, request *GetUserExchangeLimitsRequest) (*GetUserExchangeLimitsResponse, error)
	GetUserTradingVolume(ctx context.Context, request *GetUserTradingVolumeRequest) (*GetUserTradingVolumeResponse, error)
	UpdateSettlementPreference(ctx context.Context, request *UpdateSettlementPreferenceRequest) (*UpdateSettlementPreferenceResponse, error)
}

func NewUsersService(c client.RestClient) UsersService {
	return &usersServiceImpl{client: c}
}

type usersServiceImpl struct {
	client client.RestClient
}
