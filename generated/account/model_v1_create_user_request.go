/*
 * account-service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package account
// V1CreateUserRequest struct for V1CreateUserRequest
type V1CreateUserRequest struct {
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Phone string `json:"phone,omitempty"`
	TimeZone string `json:"time_zone,omitempty"`
	PreferredLanguage string `json:"preferred_language,omitempty"`
	InviteToken string `json:"invite_token,omitempty"`
	Password string `json:"password"`
	Account V1CreateAccountRequest `json:"account"`
	// Signup passphrase. Not case-sensitive.
	SignupPassphrase string `json:"signup_passphrase,omitempty"`
	// This should be a valid signup provider, set to 'aws' if using AWS Marketplace. If signup_provider is specified, signup_token also must be supplied.
	SignupProvider string `json:"signup_provider,omitempty"`
	// This should be a valid token, set to value of x-amzn-marketplace-token for AWS Marketplace.
	SignupToken string `json:"signup_token,omitempty"`
	Terms V1TermsAcceptanceRequest `json:"terms,omitempty"`
	TermsOfService []V1CreateTermsAcceptanceRequest `json:"terms_of_service"`
}
