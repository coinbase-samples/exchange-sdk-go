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
)

type DeleteProfileRequest struct {
	ProfileId string `json:"profile_id"`
	To        string `json:"to"`
}

type DeleteProfileResponse struct {
	Response string `json:"response,omitempty"`
}

func (s *profilesServiceImpl) DeleteProfile(
	ctx context.Context,
	request *DeleteProfileRequest,
) (*DeleteProfileResponse, error) {

	path := fmt.Sprintf("/profiles/%s/deactivate", request.ProfileId)

	var response string

	if err := core.HttpPut(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &DeleteProfileResponse{Response: response}, nil
}
