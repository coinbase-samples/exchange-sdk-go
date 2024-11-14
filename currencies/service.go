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
	"github.com/coinbase-samples/exchange-sdk-go/client"
)

type CurrenciesService interface {
	ListCurrencies(ctx context.Context, request *ListCurrenciesRequest) (*ListCurrenciesResponse, error)
	GetCurrency(ctx context.Context, request *GetCurrencyRequest) (*GetCurrencyResponse, error)
}

func NewCurrenciesService(c client.RestClient) CurrenciesService {
	return &currenciesServiceImpl{client: c}
}

type currenciesServiceImpl struct {
	client client.RestClient
}
