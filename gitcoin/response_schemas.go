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
	Grants []GrantResponse `json:"grants"`
	// Sensitive info as such shouldn't be out
	Credentials struct {
		IsStaff         bool `json:"is_staff"`
		IsAuthenticated bool `json:"is_authenticated"`
	} `json:"-"`
	// TODO figure MetaData out
	Metadata struct {
		ClaimStartDate interface{} `json:"claim_start_date"`
		ClaimEndDate   interface{} `json:"claim_end_date"`
		StartDate      interface{} `json:"start_date"`
		EndDate        interface{} `json:"end_date"`
	} `json:"-"`
}

type GrantResponse struct {
	Id                            int    `json:"id"`
	Active                        bool   `json:"active"`
	DetailsUrl                    string `json:"details_url"`
	Title                         string `json:"title"`
	AmountReceivedInRound         string `json:"amount_received_in_round"`
	AmountReceived                string `json:"amount_received"`
	PositiveRoundContributorCount int    `json:"positive_round_contributor_count"`
	Slug                          string `json:"slug"`
	Url                           string `json:"url"`
	Verified                      bool   `json:"verified"`
	TwitterHandle1                string `json:"twitter_handle_1"`
	ReferenceUrl                  string `json:"reference_url"`
	GithubProjectUrl              string `json:"github_project_url"`
	HasExternalFunding            string `json:"has_external_funding"`
	IsHidden                      bool   `json:"is_hidden"`
}
