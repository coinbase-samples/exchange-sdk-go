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

package loans

import (
	"context"
	"fmt"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

type GetInterestChargesRequest struct {
	LoanId string `json:"loan_id"`
}

type GetInterestChargesResponse struct {
	InterestCharges []*model.InterestCharge `json:"interest_charges"`
}

func (s *loansServiceImpl) GetInterestCharges(
	ctx context.Context,
	request *GetInterestChargesRequest,
) (*GetInterestChargesResponse, error) {

	path := fmt.Sprintf("/loans/interest/%s", request.LoanId)

	var interestCharges []*model.InterestCharge

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&interestCharges,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return &GetInterestChargesResponse{InterestCharges: interestCharges}, nil
}
