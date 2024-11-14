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

package wrappedassets

import (
	"context"
	"github.com/coinbase-samples/exchange-sdk-go/client"
)

type WrappedAssetsService interface {
	GetStakeWrap(ctx context.Context, request *GetStakeWrapRequest) (*GetStakeWrapResponse, error)
	ListStakewraps(ctx context.Context, request *ListStakewrapsRequest) (*ListStakewrapsResponse, error)
	ListWrappedAssets(ctx context.Context, request *ListWrappedAssetsRequest) (*ListWrappedAssetsResponse, error)
	GetWrappedAsset(ctx context.Context, request *GetWrappedAssetRequest) (*GetWrappedAssetResponse, error)
	GetWrappedAssetConversionRate(ctx context.Context, request *GetWrappedAssetConversionRateRequest) (*GetWrappedAssetConversionRateResponse, error)
	CreateStakewrap(ctx context.Context, request *CreateStakewrapRequest) (*CreateStakewrapResponse, error)
}

func NewWrappedAssetsService(c client.RestClient) WrappedAssetsService {
	return &wrappedAssetsServiceImpl{client: c}
}

type wrappedAssetsServiceImpl struct {
	client client.RestClient
}
