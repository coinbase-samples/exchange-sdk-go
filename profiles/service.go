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
	"github.com/coinbase-samples/exchange-sdk-go/client"
)

type ProfilesService interface {
	ListProfiles(ctx context.Context, request *ListProfilesRequest) (*ListProfilesResponse, error)
	GetProfile(ctx context.Context, request *GetProfileRequest) (*GetProfileResponse, error)
	CreateProfile(ctx context.Context, request *CreateProfileRequest) (*CreateProfileResponse, error)
	RenameProfile(ctx context.Context, request *RenameProfileRequest) (*RenameProfileResponse, error)
	DeleteProfile(ctx context.Context, request *DeleteProfileRequest) (*DeleteProfileResponse, error)
	TransferFundsBetweenProfiles(ctx context.Context, request *TransferFundsBetweenProfilesRequest) (*TransferFundsBetweenProfilesResponse, error)
}

func NewProfilesService(c client.RestClient) ProfilesService {
	return &profilesServiceImpl{client: c}
}

type profilesServiceImpl struct {
	client client.RestClient
}