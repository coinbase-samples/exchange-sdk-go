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

type ListPaymentMethodsRequest struct{}

type ListPaymentMethodsResponse struct {
	PaymentMethods []*model.PaymentMethod `json:"payment_methods"`
}

func (s *transfersServiceImpl) ListPaymentMethods(
	ctx context.Context,
	request *ListPaymentMethodsRequest,
) (*ListPaymentMethodsResponse, error) {

	path := "/payment-methods"

	var paymentMethods []*model.PaymentMethod

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&paymentMethods,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &ListPaymentMethodsResponse{PaymentMethods: paymentMethods}, nil
}
