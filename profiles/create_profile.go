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
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

type CreateProfileRequest struct {
	Name string `json:"name"`
}

type CreateProfileResponse struct {
	Profile model.Profile `json:"profile"`
}

func (s *profilesServiceImpl) CreateProfile(
	ctx context.Context,
	request *CreateProfileRequest,
) (*CreateProfileResponse, error) {

	path := "/profiles"

	var profile model.Profile

	if err := core.HttpPost(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&profile,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &CreateProfileResponse{Profile: profile}, nil
}
