/*
 * account-service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package account
// V1ValidateSignupTokenRequest struct for V1ValidateSignupTokenRequest
type V1ValidateSignupTokenRequest struct {
	SignupProvider string `json:"signup_provider,omitempty"`
	SignupToken string `json:"signup_token,omitempty"`
}
