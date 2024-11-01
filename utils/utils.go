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

package utils

import (
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/model"
)

func AppendQueryParam(queryParams, key, value string) string {
	if queryParams == "" {
		return "?" + key + "=" + value
	} else {
		return queryParams + "&" + key + "=" + value
	}
}

func AppendPaginationParams(v string, p *model.PaginationParams) string {

	if p == nil {
		return v
	}

	if len(p.Before) > 0 {
		v = core.AppendHttpQueryParam(v, "before", p.Before)
	}

	if len(p.Limit) > 0 {
		v = core.AppendHttpQueryParam(v, "limit", p.Limit)
	}

	if len(p.After) > 0 {
		v = core.AppendHttpQueryParam(v, "after", p.After)
	}

	return v
}