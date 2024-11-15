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

package profiles

import (
	"context"
	"fmt"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

type GetProfileRequest struct {
	ProfileId string `json:"profile_id"`
	Active    string `json:"active,omitempty"`
}

type GetProfileResponse struct {
	Profile model.Profile `json:"profile"`
}

func (s *profilesServiceImpl) GetProfile(
	ctx context.Context,
	request *GetProfileRequest,
) (*GetProfileResponse, error) {

	path := fmt.Sprintf("/profiles/%s", request.ProfileId)

	var queryParams string
	if len(request.Active) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "active", request.Active)
	}

	var profile model.Profile

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&profile,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &GetProfileResponse{Profile: profile}, nil
}
