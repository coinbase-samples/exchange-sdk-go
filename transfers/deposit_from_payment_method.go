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
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

type DepositFromPaymentMethodRequest struct {
	ProfileId       string `json:"profile_id"`
	Amount          string `json:"amount"`
	PaymentMethodId string `json:"payment_method_id"`
	Currency        string `json:"currency"`
}

type DepositFromPaymentMethodResponse struct {
	Transaction model.Transaction `json:"transaction"`
}

func (s *transfersServiceImpl) DepositFromPaymentMethod(
	ctx context.Context,
	request *DepositFromPaymentMethodRequest,
) (*DepositFromPaymentMethodResponse, error) {

	path := "/deposits/payment-method"

	var transaction model.Transaction

	if err := core.HttpPost(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&transaction,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &DepositFromPaymentMethodResponse{Transaction: transaction}, nil
}
