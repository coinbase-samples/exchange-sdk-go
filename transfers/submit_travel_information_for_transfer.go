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

package transfers

import (
	"context"
	"fmt"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

type SubmitTravelInformationForTransferRequest struct {
	TransferId        string `json:"transfer_id"`
	OriginatorName    string `json:"originator_name"`
	CoinbaseAccountId string `json:"coinbase_account_id"`
	Currency          string `json:"currency"`
}

type SubmitTravelInformationForTransferResponse struct {
	Message model.Message `json:"message"`
}

func (s *transfersServiceImpl) SubmitTravelInformationForTransfer(
	ctx context.Context,
	request *SubmitTravelInformationForTransferRequest,
) (*SubmitTravelInformationForTransferResponse, error) {

	path := fmt.Sprintf("/transfers/%s/travel-rules", request.TransferId)

	var message model.Message

	if err := core.HttpPost(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&message,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &SubmitTravelInformationForTransferResponse{Message: message}, nil
}
