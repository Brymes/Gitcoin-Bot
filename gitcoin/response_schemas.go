package gitcoin

import "time"

type BountiesResponse struct {
	Url             string    `json:"url"`
	Title           string    `json:"title"`
	GithubUrl       string    `json:"github_url"`
	ExperienceLevel string    `json:"experience_level"`
	Status          string    `json:"status"`
	Web3Created     time.Time `json:"web3_created"`
	Network         string    `json:"network"`
	InterestedCount int       `json:"interested_count"`
	TokenName       string    `json:"token_name"`
	ValueInUsdt     string    `json:"value_in_usdt"`
	Keywords        string    `json:"keywords"`
	ValueInToken    string    `json:"value_in_token"`
	IsOpen          bool      `json:"is_open"`
	ExpiresDate     time.Time `json:"expires_date"`
	ValueTrueUsd    string    `json:"value_true_usd"`
}

type GrantResponseWrapper struct {
	Count   int             `json:"count"`
	Next    string          `json:"next"`
	Results []GrantResponse `json:"results"`
}

type GrantResponse struct {
	CustomerName          string    `json:"customer_name"`
	DisplayText           string    `json:"display_text"`
	GrantClrPercentageCap string    `json:"grant_clr_percentage_cap"`
	IsActive              bool      `json:"is_active"`
	StartDate             time.Time `json:"start_date"`
	EndDate               time.Time `json:"end_date"`
	ClaimStartDate        time.Time `json:"claim_start_date"`
	ClaimEndDate          time.Time `json:"claim_end_date"`
	VerifiedThreshold     string    `json:"verified_threshold"`
	UnverifiedThreshold   string    `json:"unverified_threshold"`
	BannerText            string    `json:"banner_text"`
}
