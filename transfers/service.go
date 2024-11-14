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
	"github.com/coinbase-samples/exchange-sdk-go/client"
)

type TransfersService interface {
	ListPaymentMethods(ctx context.Context, request *ListPaymentMethodsRequest) (*ListPaymentMethodsResponse, error)
	ListTransfers(ctx context.Context, request *ListTransfersRequest) (*ListTransfersResponse, error)
	GetTransfer(ctx context.Context, request *GetTransferRequest) (*GetTransferResponse, error)
	GetFeeEstimateForWithdrawal(ctx context.Context, request *GetFeeEstimateForWithdrawalRequest) (*GetFeeEstimateForWithdrawalResponse, error)
	DepositFromCoinbaseAccount(ctx context.Context, request *DepositFromCoinbaseAccountRequest) (*DepositFromCoinbaseAccountResponse, error)
	DepositFromPaymentMethod(ctx context.Context, request *DepositFromPaymentMethodRequest) (*DepositFromPaymentMethodResponse, error)
	SubmitTravelInformationForTransfer(ctx context.Context, request *SubmitTravelInformationForTransferRequest) (*SubmitTravelInformationForTransferResponse, error)
	WithdrawToCoinbaseAccount(ctx context.Context, request *WithdrawToCoinbaseAccountRequest) (*WithdrawToCoinbaseAccountResponse, error)
	WithdrawToCryptoAddress(ctx context.Context, request *WithdrawToCryptoAddressRequest) (*WithdrawToCryptoAddressResponse, error)
	WithdrawToPaymentMethod(ctx context.Context, request *WithdrawToPaymentMethodRequest) (*WithdrawToPaymentMethodResponse, error)
}

func NewTransfersService(c client.RestClient) TransfersService {
	return &transfersServiceImpl{client: c}
}

type transfersServiceImpl struct {
	client client.RestClient
}
