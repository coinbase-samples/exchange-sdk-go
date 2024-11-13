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

package accounts

import (
	"context"
	"github.com/coinbase-samples/exchange-sdk-go/client"
)

type AccountsService interface {
	GetAccount(ctx context.Context, request *GetAccountRequest) (*GetAccountResponse, error)
	GetAccountHolds(ctx context.Context, request *GetAccountHoldsRequest) (*GetAccountHoldsResponse, error)
	GetAccountLedger(ctx context.Context, request *GetAccountLedgerRequest) (*GetAccountLedgerResponse, error)
	GetAccountTransfers(ctx context.Context, request *GetAccountTransfersRequest) (*GetAccountTransfersResponse, error)
	ListAccounts(ctx context.Context, request *ListAccountsRequest) (*ListAccountsResponse, error)
}

func NewAccountsService(c client.RestClient) AccountsService {
	return &accountsServiceImpl{client: c}
}

type accountsServiceImpl struct {
	client client.RestClient
}
