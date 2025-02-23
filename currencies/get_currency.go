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
	"fmt"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

type GetCurrencyRequest struct {
	CurrencyId string `json:"currency_id"`
}

type GetCurrencyResponse struct {
	Currency model.Currency `json:"currency"`
}

func (s *currenciesServiceImpl) GetCurrency(
	ctx context.Context,
	request *GetCurrencyRequest,
) (*GetCurrencyResponse, error) {

	path := fmt.Sprintf("/currencies/%s", request.CurrencyId)

	var currency model.Currency

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&currency,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &GetCurrencyResponse{Currency: currency}, nil
}
