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

package currencies

import (
	"context"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

type ListCurrenciesRequest struct{}

type ListCurrenciesResponse struct {
	Currencies []*model.Currency `json:"currencies"`
}

func (s *currenciesServiceImpl) ListCurrencies(
	ctx context.Context,
	request *ListCurrenciesRequest,
) (*ListCurrenciesResponse, error) {

	path := "/currencies"

	var currencies []*model.Currency

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&currencies,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &ListCurrenciesResponse{Currencies: currencies}, nil
}
