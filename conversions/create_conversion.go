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

package conversions

import (
	"context"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

type CreateConversionRequest struct {
	ProfileId string `json:"profile_id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Amount    string `json:"amount"`
	Nonce     string `json:"nonce"`
}

type CreateConversionResponse struct {
	Conversion model.Conversion `json:"conversion"`
}

func (s *conversionsServiceImpl) CreateConversion(
	ctx context.Context,
	request *CreateConversionRequest,
) (*CreateConversionResponse, error) {

	path := "/conversions"

	var conversion model.Conversion

	if err := core.HttpPost(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&conversion,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &CreateConversionResponse{Conversion: conversion}, nil
}
