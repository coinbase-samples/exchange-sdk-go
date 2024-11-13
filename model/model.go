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

package model

import "time"

type ErrorMessage struct {
	Value string `json:"message"`
}

type PaginationParams struct {
	Before string `json:"before,omitempty"`
	After  string `json:"after,omitempty"`
	Limit  string `json:"limit,omitempty"`
}

type Account struct {
	Id             string `json:"id"`
	Currency       string `json:"currency"`
	Balance        string `json:"balance"`
	Hold           string `json:"hold"`
	Available      string `json:"available"`
	ProfileId      string `json:"profile_id"`
	TradingEnabled bool   `json:"trading_enabled"`
	PendingDeposit string `json:"pending_deposit"`
	DisplayName    string `json:"display_name"`
}

type AccountHold struct {
	CreatedAt string `json:"created_at"`
	Id        string `json:"id"`
	Amount    string `json:"amount"`
	Type      string `json:"type"`
	Ref       string `json:"ref"`
}

type AccountLedger struct {
	Id        string  `json:"id"`
	Amount    string  `json:"amount"`
	CreatedAt string  `json:"created_at"`
	Balance   string  `json:"balance"`
	Type      string  `json:"type"`
	Details   Details `json:"details"`
}

type Details struct {
	To                string `json:"to"`
	From              string `json:"from"`
	ProfileTransferId string `json:"profile_transfer_id"`
}

type AccountTransfer struct {
	Id          string             `json:"id"`
	Type        string             `json:"type"`
	CreatedAt   string             `json:"created_at"`
	CompletedAt string             `json:"completed_at"`
	CanceledAt  *string            `json:"canceled_at,omitempty"`
	ProcessedAt *string            `json:"processed_at,omitempty"`
	Amount      string             `json:"amount"`
	Details     TransactionDetails `json:"details"`
	UserNonce   *int64             `json:"user_nonce,omitempty"`
	Currency    string             `json:"currency"`
}

type TransactionDetails struct {
	CoinbaseAccountId       string `json:"coinbase_account_id"`
	CoinbaseTransactionId   string `json:"coinbase_transaction_id"`
	CoinbasePaymentMethodId string `json:"coinbase_payment_method_id"`
}

type AddressBook struct {
	Id                         string  `json:"id"`
	Address                    string  `json:"address"`
	Currency                   string  `json:"currency"`
	Label                      string  `json:"label"`
	LastUsed                   *string `json:"last_used,omitempty"`
	AddressBookAddedAt         string  `json:"address_book_added_at"`
	DestinationTag             *string `json:"destination_tag,omitempty"`
	IsVerifiedSelfHostedWallet bool    `json:"is_verified_self_hosted_wallet"`
	VaspId                     *string `json:"vasp_id,omitempty"`
}

type Address struct {
	Currency                   string  `json:"currency"`
	To                         To      `json:"to"`
	Label                      string  `json:"label"`
	IsVerifiedSelfHostedWallet bool    `json:"is_verified_self_hosted_wallet"`
	VaspId                     *string `json:"vasp_id,omitempty"`
}

type To struct {
	Address        string  `json:"address"`
	DestinationTag *string `json:"destination_tag,omitempty"`
}

type AddressBookResponse struct {
	Id                  string      `json:"id"`
	Address             string      `json:"address"`
	AddressInfo         AddressInfo `json:"address_info"`
	DisplayAddress      string      `json:"display_address"`
	Trusted             bool        `json:"trusted"`
	AddressBooked       bool        `json:"address_booked"`
	AddressBookAddedAt  string      `json:"address_book_added_at"`
	LastUsed            *string     `json:"last_used,omitempty"`
	Label               string      `json:"label"`
	PreferLegacyAddress bool        `json:"prefer_legacy_address"`
	Currency            string      `json:"currency"`
}

type AddressInfo struct {
	Address        string  `json:"address"`
	DisplayAddress string  `json:"display_address"`
	DestinationTag *string `json:"destination_tag,omitempty"`
}

type CoinbaseWallet struct {
	Id                  string  `json:"id"`
	Name                string  `json:"name"`
	Balance             string  `json:"balance"`
	Currency            string  `json:"currency"`
	Type                string  `json:"type"`
	Primary             bool    `json:"primary"`
	Active              bool    `json:"active"`
	AvailableOnConsumer bool    `json:"available_on_consumer"`
	HoldBalance         string  `json:"hold_balance"`
	HoldCurrency        string  `json:"hold_currency"`
	DestinationTagName  *string `json:"destination_tag_name"`
	DestinationTagRegex *string `json:"destination_tag_regex"`
}

type AddressSubset struct {
	Address string `json:"address"`
}

type Warning struct {
	Title    string `json:"title"`
	Details  string `json:"details"`
	ImageUrl string `json:"image_url"`
}

type AddressResponse struct {
	Id                     string      `json:"id"`
	Address                string      `json:"address"`
	AddressInfo            AddressInfo `json:"address_info"`
	Name                   string      `json:"name"`
	CreatedAt              string      `json:"created_at"`
	UpdatedAt              string      `json:"updated_at"`
	Network                string      `json:"network"`
	UriScheme              string      `json:"uri_scheme"`
	Resource               string      `json:"resource"`
	ResourcePath           string      `json:"resource_path"`
	Warnings               []Warning   `json:"warnings"`
	DepositUri             string      `json:"deposit_uri"`
	CallbackUrl            *string     `json:"callback_url"`
	ExchangeDepositAddress bool        `json:"exchange_deposit_address"`
}

type Conversion struct {
	Id            string `json:"id"`
	Amount        string `json:"amount"`
	FromAccountId string `json:"from_account_id"`
	ToAccountId   string `json:"to_account_id"`
	From          string `json:"from"`
	To            string `json:"to"`
	FeeAmount     string `json:"fee_amount"`
}

type FeeRate struct {
	FromCurrency    string `json:"from_currency"`
	ToCurrency      string `json:"to_currency"`
	FeeRate         string `json:"fee_rate"`
	ThirtyDayVolume string `json:"thirty_day_volume"`
}

type Currency struct {
	Id              string  `json:"id"`
	Name            string  `json:"name"`
	MinSize         string  `json:"min_size"`
	MaxPrecision    string  `json:"max_precision"`
	Status          string  `json:"status"`
	CurrencyDetails Details `json:"details"`
}

type CurrencyDetails struct {
	Type               string   `json:"type"`
	Symbol             string   `json:"symbol"`
	SortOrder          string   `json:"sort_order"`
	PushPaymentMethods []string `json:"push_payment_methods"`
	DisplayName        string   `json:"display_name"`
	GroupTypes         []string `json:"group_types"`
}

type PaymentMethod struct {
	Id                 string            `json:"id"`
	Type               string            `json:"type"`
	Name               string            `json:"name"`
	Currency           string            `json:"currency"`
	PrimaryBuy         bool              `json:"primary_buy"`
	PrimarySell        bool              `json:"primary_sell"`
	InstantBuy         bool              `json:"instant_buy"`
	InstantSell        bool              `json:"instant_sell"`
	CreatedAt          string            `json:"created_at"`
	UpdatedAt          string            `json:"updated_at"`
	Resource           string            `json:"resource"`
	ResourcePath       string            `json:"resource_path"`
	Limits             Limits            `json:"limits"`
	AllowBuy           bool              `json:"allow_buy"`
	AllowSell          bool              `json:"allow_sell"`
	AllowDeposit       bool              `json:"allow_deposit"`
	AllowWithdraw      bool              `json:"allow_withdraw"`
	FiatAccount        FiatAccount       `json:"fiat_account"`
	CryptoAccount      CryptoAccount     `json:"crypto_account,omitempty"`
	RecurringOptions   []RecurringOption `json:"recurring_options,omitempty"`
	AvailableBalance   AvailableBalance  `json:"available_balance,omitempty"`
	PickerData         PickerData        `json:"picker_data"`
	HoldBusinessDays   int               `json:"hold_business_days"`
	HoldDays           int               `json:"hold_days"`
	VerificationMethod string            `json:"verificationMethod,omitempty"`
	CdvStatus          string            `json:"cdvStatus,omitempty"`
}

type Limits struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type FiatAccount struct {
	Id           string `json:"id"`
	Resource     string `json:"resource"`
	ResourcePath string `json:"resource_path"`
}

type CryptoAccount struct {
	Id           string `json:"id"`
	Resource     string `json:"resource"`
	ResourcePath string `json:"resource_path"`
}

type RecurringOption struct {
	Period string `json:"period"`
	Label  string `json:"label"`
}

type AvailableBalance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Scale    string `json:"scale,omitempty"`
}

type PickerData struct {
	Symbol                string  `json:"symbol"`
	CustomerName          string  `json:"customer_name,omitempty"`
	AccountName           string  `json:"account_name,omitempty"`
	AccountNumber         string  `json:"account_number,omitempty"`
	AccountType           string  `json:"account_type,omitempty"`
	InstitutionCode       string  `json:"institution_code,omitempty"`
	InstitutionName       string  `json:"institution_name,omitempty"`
	Iban                  string  `json:"iban,omitempty"`
	Swift                 string  `json:"swift,omitempty"`
	PaypalEmail           string  `json:"paypal_email,omitempty"`
	PaypalOwner           string  `json:"paypal_owner,omitempty"`
	RoutingNumber         string  `json:"routing_number,omitempty"`
	InstitutionIdentifier string  `json:"institution_identifier,omitempty"`
	BankName              string  `json:"bank_name,omitempty"`
	BranchName            string  `json:"branch_name,omitempty"`
	IconUrl               string  `json:"icon_url,omitempty"`
	Balance               Balance `json:"balance"`
}

type Balance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Transfer struct {
	Id          string          `json:"id"`
	Type        string          `json:"type"`
	CreatedAt   string          `json:"created_at"`
	CompletedAt string          `json:"completed_at"`
	CanceledAt  *string         `json:"canceled_at,omitempty"`
	ProcessedAt *string         `json:"processed_at,omitempty"`
	UserNonce   *int64          `json:"user_nonce,omitempty"`
	Amount      string          `json:"amount"`
	Details     TransferDetails `json:"details"`
	Currency    string          `json:"currency"`
}

type TransferDetails struct {
	CoinbaseAccountId       string `json:"coinbase_account_id"`
	CoinbaseTransactionId   string `json:"coinbase_transaction_id"`
	CoinbasePaymentMethodId string `json:"coinbase_payment_method_id"`
}

type FeeEstimate struct {
	Fee              string `json:"fee"`
	FeeBeforeSubsidy string `json:"fee_before_subsidy"`
}

type Fees struct {
	MakerFeeRate string `json:"maker_fee_rate"`
	TakerFeeRate string `json:"taker_fee_rate"`
	UsdVolume    string `json:"usd_volume"`
}

type Fill struct {
	CreatedAt       time.Time `json:"created_at"`
	TradeId         int       `json:"trade_id"`
	ProductId       string    `json:"product_id"`
	OrderId         string    `json:"order_id"`
	UserId          string    `json:"user_id"`
	ProfileId       string    `json:"profile_id"`
	Liquidity       string    `json:"liquidity"`
	Price           string    `json:"price"`
	Size            string    `json:"size"`
	Fee             string    `json:"fee"`
	Side            string    `json:"side"`
	Settled         bool      `json:"settled"`
	UsdVolume       string    `json:"usd_volume"`
	FundingCurrency string    `json:"funding_currency"`
}

type Order struct {
	Id            string    `json:"id"`
	Price         string    `json:"price"`
	Size          string    `json:"size"`
	ProductId     string    `json:"product_id"`
	ProfileId     string    `json:"profile_id"`
	Side          string    `json:"side"`
	Type          string    `json:"type"`
	TimeInForce   string    `json:"time_in_force"`
	PostOnly      bool      `json:"post_only"`
	MaxFloor      string    `json:"max_floor"`
	CreatedAt     time.Time `json:"created_at"`
	FillFees      string    `json:"fill_fees"`
	FilledSize    string    `json:"filled_size"`
	ExecutedValue string    `json:"executed_value"`
	Status        string    `json:"status"`
	Settled       bool      `json:"settled"`
}

type SignedPrice struct {
	Timestamp  string        `json:"timestamp"`
	Messages   []interface{} `json:"messages"`
	Signatures []interface{} `json:"signatures"`
	Prices     interface{}   `json:"prices"`
}

type WrappedAsset struct {
	Id                string `json:"id"`
	CirculatingSupply string `json:"circulating_supply"`
	TotalSupply       string `json:"total_supply"`
	ConversionRate    string `json:"conversion_rate"`
	Apy               string `json:"apy"`
}

type Stakewrap struct {
	Id             string    `json:"id"`
	FromAmount     string    `json:"from_amount"`
	ToAmount       string    `json:"to_amount"`
	FromAccountId  string    `json:"from_account_id"`
	ToAccountId    string    `json:"to_account_id"`
	FromCurrency   string    `json:"from_currency"`
	ToCurrency     string    `json:"to_currency"`
	Status         string    `json:"status"`
	ConversionRate string    `json:"conversion_rate"`
	CreatedAt      time.Time `json:"created_at"`
	CompletedAt    time.Time `json:"completed_at"`
	CanceledAt     time.Time `json:"canceled_at"`
}

type Amount struct {
	Amount string `json:"amount"`
}

type ExchangeLimit struct {
	LimitCurrency  string                           `json:"limit_currency"`
	TransferLimits map[string]map[string]AssetLimit `json:"transfer_limits"`
}

type AssetLimit struct {
	Max          string `json:"max"`
	Remaining    string `json:"remaining"`
	PeriodInDays int    `json:"period_in_days"`
}

type ActivityMetrics struct {
	StartDate                                     string `json:"start_date"`
	EndDate                                       string `json:"end_date"`
	MakerVolumeNotionalUSD                        string `json:"maker_volume_notional_usd"`
	MakerVolumeRelativePercentage                 string `json:"maker_volume_relative_percentage"`
	TakerVolumeNotionalUSD                        string `json:"taker_volume_notional_usd"`
	TakerVolumeRelativePercentage                 string `json:"taker_volume_relative_percentage"`
	TotalExchangeVolumeNotionalUSD                string `json:"total_exchange_volume_notional_usd"`
	TotalExchangeVolumeRelativePercentage         string `json:"total_exchange_volume_relative_percentage"`
	MakerRank                                     string `json:"maker_rank"`
	TakerRank                                     string `json:"taker_rank"`
	TotalExchangeVolumeRank                       string `json:"total_exchange_volume_rank"`
	AdjustedMakerVolumeNotionalUSD                string `json:"adjusted_maker_volume_notional_usd"`
	AdjustedMakerVolumeRelativePercentage         string `json:"adjusted_maker_volume_relative_percentage"`
	AdjustedTotalExchangeVolumeNotionalUSD        string `json:"adjusted_total_exchange_volume_notional_usd"`
	AdjustedTotalExchangeVolumeRelativePercentage string `json:"adjusted_total_exchange_volume_relative_percentage"`
	AdjustedMakerVolumeRank                       string `json:"adjusted_maker_volume_rank"`
	AdjustedTotalExchangeVolumeRank               string `json:"adjusted_total_exchange_volume_rank"`
	LiquidityProgramTier                          string `json:"liquidity_program_tier"`
	NextLiquidityProgramTier                      string `json:"next_liquidity_program_tier"`
}

type AggregatedData struct {
	ActivityMetrics ActivityMetrics `json:"activity_metrics"`
}

type IndividualData struct {
	Email           string          `json:"email"`
	ActivityMetrics ActivityMetrics `json:"activity_metrics"`
}

type TradingVolume struct {
	AggregatedData AggregatedData `json:"aggregated_data"`
	IndividualData IndividualData `json:"individual_data"`
}

type TravelRuleResponse struct {
	Id                TravelRuleDetail   `json:"id"`
	CreatedAt         TravelRuleCreation `json:"created_at"`
	Address           TravelRuleDetail   `json:"address"`
	OriginatorName    TravelRuleDetail   `json:"originator_name"`
	OriginatorCountry TravelRuleDetail   `json:"originator_country"`
}

type TravelRuleDetail struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type TravelRuleCreation struct {
	Type        string `json:"type"`
	Format      string `json:"format"`
	Description string `json:"description"`
}

type Preferences struct {
	PreferredMarket              string    `json:"preferred_market"`
	MarginTermsCompletedInUtc    time.Time `json:"margin_terms_completed_in_utc"`
	MarginTutorialCompletedInUtc time.Time `json:"margin_tutorial_completed_in_utc"`
}

type User struct {
	CreatedAt               time.Time              `json:"created_at"`
	ActiveAt                time.Time              `json:"active_at"`
	Id                      string                 `json:"id"`
	Name                    string                 `json:"name"`
	Email                   string                 `json:"email"`
	Roles                   []interface{}          `json:"roles"`
	IsBanned                bool                   `json:"is_banned"`
	UserType                string                 `json:"user_type"`
	FulfillsNewRequirements bool                   `json:"fulfills_new_requirements"`
	Flags                   map[string]string      `json:"flags"`
	Details                 map[string]interface{} `json:"details"`
	OauthClient             string                 `json:"oauth_client"`
	Preferences             Preferences            `json:"preferences"`
	HasDefault              bool                   `json:"has_default"`
}

type Params struct {
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	Format       string    `json:"format"`
	ProductId    string    `json:"product_id"`
	AccountId    string    `json:"account_id"`
	ProfileId    string    `json:"profile_id"`
	Email        string    `json:"email"`
	User         User      `json:"user"`
	NewYorkState bool      `json:"new_york_state"`
}

type Report struct {
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	Id          string    `json:"id"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	UserId      string    `json:"user_id"`
	FileUrl     string    `json:"file_url"`
	Params      Params    `json:"params"`
	FileCount   *uint64   `json:"file_count"`
}

type Profile struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	IsDefault bool      `json:"is_default"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	Id             string `json:"id"`
	BaseCurrency   string `json:"base_currency"`
	QuoteCurrency  string `json:"quote_currency"`
	QuoteIncrement string `json:"quote_increment"`
	BaseIncrement  string `json:"base_increment"`
	DisplayName    string `json:"display_name"`
	MinMarketFunds string `json:"min_market_funds"`
	MarginEnabled  bool   `json:"margin_enabled"`
	PostOnly       bool   `json:"post_only"`
	LimitOnly      bool   `json:"limit_only"`
	CancelOnly     bool   `json:"cancel_only"`
	Status         string `json:"status"`
	StatusMessage  string `json:"status_message"`
	AuctionMode    bool   `json:"auction_mode"`
}

type ProductVolume struct {
	Id                     string   `json:"id"`
	BaseCurrency           string   `json:"base_currency"`
	QuoteCurrency          string   `json:"quote_currency"`
	DisplayName            string   `json:"display_name"`
	MarketTypes            []string `json:"market_types"`
	SpotVolume24Hour       string   `json:"spot_volume_24hour"`
	SpotVolume30Day        string   `json:"spot_volume_30day"`
	RfqVolume24Hour        string   `json:"rfq_volume_24hour"`
	RfqVolume30Day         string   `json:"rfq_volume_30day"`
	ConversionVolume24Hour string   `json:"conversion_volume_24hour"`
	ConversionVolume30Day  string   `json:"conversion_volume_30day"`
}

type ProductBook struct {
	Sequence int64           `json:"sequence"`
	Bids     [][]interface{} `json:"bids"`
	Asks     [][]interface{} `json:"asks"`
	Time     time.Time       `json:"time"`
}

type ProductStats struct {
	Open                    string `json:"open"`
	High                    string `json:"high"`
	Low                     string `json:"low"`
	Volume                  string `json:"volume"`
	Last                    string `json:"last"`
	Volume30Day             string `json:"volume_30day"`
	RfqVolume24Hour         string `json:"rfq_volume_24hour"`
	ConversionsVolume24Hour string `json:"conversions_volume_24hour"`
	RfqVolume30Day          string `json:"rfq_volume_30day"`
	ConversionsVolume30Day  string `json:"conversions_volume_30day"`
}

type ProductTicker struct {
	TradeId           int       `json:"trade_id"`
	Price             string    `json:"price"`
	Size              string    `json:"size"`
	Time              time.Time `json:"time"`
	Bid               string    `json:"bid"`
	Ask               string    `json:"ask"`
	Volume            string    `json:"volume"`
	RfqVolume         string    `json:"rfq_volume"`
	ConversionsVolume string    `json:"conversions_volume"`
}

type ProductTrades struct {
	Time    time.Time `json:"time"`
	TradeId int       `json:"trade_id"`
	Price   string    `json:"price"`
	Size    string    `json:"size"`
	Side    string    `json:"side"`
}

type Loan struct {
	Id                         string `json:"id"`
	Currency                   string `json:"currency"`
	PrincipalAmount            string `json:"principal_amount"`
	OutstandingPrincipalAmount string `json:"outstanding_principal_amount"`
	InterestRate               string `json:"interest_rate"`
	InterestCurrency           string `json:"interest_currency"`
	Status                     string `json:"status"`
	EffectiveAt                string `json:"effective_at"`
	ClosedAt                   string `json:"closed_at,omitempty"`
	TermStartDate              string `json:"term_start_date"`
	TermEndDate                string `json:"term_end_date"`
}

type LoanAsset struct {
	CollateralAssets     interface{} `json:"collateral_assets"`
	DiversificationRatio string      `json:"diversification_ratio"`
	BorrowableAssets     []string    `json:"borrowable_assets"`
}

type InterestSummary struct {
	Currency               string `json:"currency"`
	CurrentOwed            string `json:"current_owed"`
	LastPaymentDate        string `json:"last_payment_date"`
	PaymentStatus          string `json:"payment_status"`
	LastPaymentAmount      string `json:"last_payment_amount"`
	PriorPeriodOverdue     string `json:"prior_period_overdue"`
	CurrentInterestDueDate string `json:"current_interest_due_date"`
}

type RateHistory struct {
	InterestRate string `json:"interest_rate"`
	EffectiveAt  string `json:"effective_at"`
}

type InterestCharge struct {
	Date            string `json:"date"`
	Currency        string `json:"currency"`
	PrincipalAmount string `json:"principal_amount"`
	InterestRate    string `json:"interest_rate"`
	InterestAccrued string `json:"interest_accrued"`
}

type Overview struct {
	OpenLoanValue                       string      `json:"open_loan_value"`
	CollateralValue                     string      `json:"collateral_value"`
	CollateralizationPercentage         string      `json:"collateralization_percentage"`
	AvailableToBorrow                   string      `json:"available_to_borrow"`
	AvailablePerAsset                   interface{} `json:"available_per_asset"`
	WithdrawalRestricted                string      `json:"withdrawal_restricted"`
	CreditLimitValue                    string      `json:"credit_limit_value"`
	AvailableCreditValue                string      `json:"available_credit_value"`
	CollateralizationPercentageOpenOnly string      `json:"collateralization_percentage_open_only"`
	PendingLoanValue                    string      `json:"pending_loan_value"`
	InitialMarginPercentage             string      `json:"initial_margin_percentage"`
	MinimumMarginPercentage             string      `json:"minimum_margin_percentage"`
	UnlockMarginPercentage              string      `json:"unlock_margin_percentage"`
}

type LendingOverview struct {
	Overview Overview `json:"overview"`
	Loans    Loan     `json:"loans"`
}

type LoanPreview struct {
	Before Before `json:"before"`
	After  After  `json:"after"`
}

type Before struct {
	OpenLoanValue                       string      `json:"open_loan_value"`
	CollateralValue                     string      `json:"collateral_value"`
	CollateralizationPercentage         string      `json:"collateralization_percentage"`
	AvailableToBorrow                   string      `json:"available_to_borrow"`
	AvailablePerAsset                   interface{} `json:"available_per_asset"`
	WithdrawalRestricted                string      `json:"withdrawal_restricted"`
	CreditLimitValue                    string      `json:"credit_limit_value"`
	AvailableCreditValue                string      `json:"available_credit_value"`
	CollateralizationPercentageOpenOnly string      `json:"collateralization_percentage_open_only"`
	PendingLoanValue                    string      `json:"pending_loan_value"`
	InitialMarginPercentage             string      `json:"initial_margin_percentage"`
	MinimumMarginPercentage             string      `json:"minimum_margin_percentage"`
	UnlockMarginPercentage              string      `json:"unlock_margin_percentage"`
}

type After struct {
	OpenLoanValue                       string      `json:"open_loan_value"`
	CollateralValue                     string      `json:"collateral_value"`
	CollateralizationPercentage         string      `json:"collateralization_percentage"`
	AvailableToBorrow                   string      `json:"available_to_borrow"`
	AvailablePerAsset                   interface{} `json:"available_per_asset"`
	WithdrawalRestricted                string      `json:"withdrawal_restricted"`
	CreditLimitValue                    string      `json:"credit_limit_value"`
	AvailableCreditValue                string      `json:"available_credit_value"`
	CollateralizationPercentageOpenOnly string      `json:"collateralization_percentage_open_only"`
	PendingLoanValue                    string      `json:"pending_loan_value"`
	InitialMarginPercentage             string      `json:"initial_margin_percentage"`
	MinimumMarginPercentage             string      `json:"minimum_margin_percentage"`
	UnlockMarginPercentage              string      `json:"unlock_margin_percentage"`
}

type MaxPrincipalAmount struct {
	Native   string `json:"native"`
	Notional string `json:"notional"`
}

type LoanOption struct {
	Currency           string             `json:"currency"`
	MaxPrincipalAmount MaxPrincipalAmount `json:"max_principal_amount"`
	InterestRate       string             `json:"interest_rate"`
}

type Transaction struct {
	Id       string `json:"id"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	PayoutAt string `json:"payout_at"`
	Fee      string `json:"fee"`
	Subtotal string `json:"subtotal"`
}

type Message struct {
	Message string `json:"message"`
}

type Description struct {
	Description string `json:"description"`
}

type WithdrawalInformation struct {
	ProfileId                       string          `json:"profile_id"`
	Amount                          string          `json:"amount"`
	Currency                        string          `json:"currency"`
	CryptoAddress                   string          `json:"crypto_address"`
	DestinationTag                  string          `json:"destination_tag,omitempty"`
	NoDestinationTag                bool            `json:"no_destination_tag"`
	Nonce                           int32           `json:"nonce"`
	Network                         string          `json:"network"`
	AddNetworkFeeToTotal            bool            `json:"add_network_fee_to_total,omitempty"`
	IsIntermediary                  bool            `json:"is_intermediary,omitempty"`
	IntermediaryJurisdiction        string          `json:"intermediary_jurisdiction,omitempty"`
	TravelRuleData                  *TravelRuleData `json:"travel_rule_data,omitempty"`
	TransferPurpose                 string          `json:"transfer_purpose,omitempty"`
	BeneficiaryName                 string          `json:"beneficiary_name,omitempty"`
	BeneficiaryAddress              *Address        `json:"beneficiary_address,omitempty"`
	BeneficiaryTelephoneNumber      string          `json:"beneficiary_telephone_number,omitempty"`
	BeneficiaryAccountLocation      string          `json:"beneficiary_account_location,omitempty"`
	BeneficiaryDateOfBirth          *DateOfBirth    `json:"beneficiary_date_of_birth,omitempty"`
	BeneficiaryFinancialInstitution string          `json:"beneficiary_financial_institution,omitempty"`
	IsSelf                          bool            `json:"is_self"`
	OriginatorWalletAddress         string          `json:"originator_wallet_address,omitempty"`
}

type TravelRuleData struct {
	OriginatorName                   string             `json:"originator_name"`
	OriginatorNaturalName            *NaturalName       `json:"originator_natural_name,omitempty"`
	OriginatorAddress                *OriginatorAddress `json:"originator_address,omitempty"`
	OriginatorTelephoneNumber        string             `json:"originator_telephone_number,omitempty"`
	OriginatorAccount                string             `json:"originator_account,omitempty"`
	OriginatorAccountLocationCountry string             `json:"originator_account_location_country_code,omitempty"`
	OriginatorDateOfBirth            *DateOfBirth       `json:"originator_date_of_birth,omitempty"`
	OriginatorAccountNumber          string             `json:"originator_account_number,omitempty"`
}

type NaturalName struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type OriginatorAddress struct {
	Address1   string `json:"address_1"`
	Address2   string `json:"address_2,omitempty"`
	Address3   string `json:"address_3,omitempty"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

type DateOfBirth struct {
	Year  int32 `json:"year"`
	Month int32 `json:"month"`
	Day   int32 `json:"day"`
}

type Repayment struct {
	Id           string `json:"id"`
	NativeAmount string `json:"native_amount"`
	Status       string `json:"status"`
	Type         string `json:"type"`
}

type BalanceParams struct {
	Datetime string `json:"datetime,omitempty"`
}

type FillsParams struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	ProductId string `json:"product_id,omitempty"`
}

type AccountParams struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	AccountId string `json:"account_id,omitempty"`
}

type OtcFillsParams struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	ProductId string `json:"product_id,omitempty"`
}

type TaxInvoiceParams struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	ProductId string `json:"product_id,omitempty"`
}

type RfqFillsParams struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	ProductId string `json:"product_id,omitempty"`
}

type ReportResponse struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

type SettlementPreference struct {
	SettlementPreference string `json:"settlement_preference"`
}
