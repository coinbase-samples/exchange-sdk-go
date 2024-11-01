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

package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var defaultV1ApiBaseUrl = "https://api.exchange.coinbase.com"
var defaultHeadersFunc = AddExchangeHeaders
var DefaultSuccessHttpStatusCodes = []int{http.StatusOK}

type RestClient interface {
	HttpBaseUrl() string
	SetBaseUrl(u string) RestClient

	HttpClient() *http.Client
	SetHeadersFunc(hf core.HttpHeaderFunc) RestClient
	HeadersFunc() core.HttpHeaderFunc

	Credentials() *credentials.Credentials
}

type restClientImpl struct {
	httpClient http.Client
	baseUrl    string

	headersFunc core.HttpHeaderFunc
	credentials *credentials.Credentials
}

func (c *restClientImpl) HttpBaseUrl() string {
	return c.baseUrl
}

func (c *restClientImpl) SetBaseUrl(u string) RestClient {
	c.baseUrl = u
	return c
}

func (c *restClientImpl) HttpClient() *http.Client {
	return &c.httpClient
}

func (c *restClientImpl) Credentials() *credentials.Credentials {
	return c.credentials
}

func (c *restClientImpl) SetHeadersFunc(hf core.HttpHeaderFunc) RestClient {
	c.headersFunc = hf
	return c
}

func (c *restClientImpl) HeadersFunc() core.HttpHeaderFunc {
	return c.headersFunc
}

func NewRestClient(credentials *credentials.Credentials, httpClient http.Client) RestClient {
	httpBaseUrl := os.Getenv("EXCHANGE_BASE_URL")
	if httpBaseUrl == "" {
		httpBaseUrl = defaultV1ApiBaseUrl
	}
	return &restClientImpl{
		baseUrl:     httpBaseUrl,
		credentials: credentials,
		httpClient:  httpClient,
		headersFunc: defaultHeadersFunc,
	}
}

func AddExchangeHeaders(req *http.Request, path string, body []byte, client core.RestClient, t time.Time) {
	c := client.(*restClientImpl)
	signature := sign(req.Method, path, t.Unix(), c.Credentials().SigningKey, body)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("CB-ACCESS-KEY", c.Credentials().ApiKey)
	req.Header.Add("CB-ACCESS-PASSPHRASE", c.Credentials().Passphrase)
	req.Header.Add("CB-ACCESS-SIGN", signature)
	req.Header.Add("CB-ACCESS-TIMESTAMP", strconv.FormatInt(t.Unix(), 10))
	req.Header.Add("Content-Type", "application/json")

}

func sign(method, path string, t int64, signingKey string, body []byte) string {
	key, err := base64.StdEncoding.DecodeString(signingKey)
	if err != nil {
		log.Fatalf("Error decoding signing key: %v", err)
	}

	message := fmt.Sprintf("%d%s%s", t, method, path)
	if len(body) > 0 {
		message += string(body)
	}

	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
