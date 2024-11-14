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
)

type TransferFundsBetweenProfilesRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type TransferFundsBetweenProfilesResponse struct {
	Response string `json:"response,omitempty"`
}

func (s *profilesServiceImpl) TransferFundsBetweenProfiles(
	ctx context.Context,
	request *TransferFundsBetweenProfilesRequest,
) (*TransferFundsBetweenProfilesResponse, error) {

	path := "/profiles/transfer"

	var response string

	if err := core.HttpPost(
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

	return &TransferFundsBetweenProfilesResponse{Response: response}, nil
}
