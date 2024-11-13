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
	"fmt"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

type GetConversionRequest struct {
	ConversionId string `json:"conversion_id"`
	ProfileId    string `json:"start_date,omitempty"`
}

type GetConversionResponse struct {
	Conversion []*model.Conversion `json:"conversion"`
}

func (s *conversionsServiceImpl) GetConversion(
	ctx context.Context,
	request *GetConversionRequest,
) (*GetConversionResponse, error) {

	path := fmt.Sprintf("/conversions/%s", request.ConversionId)

	var queryParams string
	if len(request.ProfileId) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "profile_id", request.ProfileId)
	}

	var conversion []*model.Conversion

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&conversion,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &GetConversionResponse{Conversion: conversion}, nil
}
