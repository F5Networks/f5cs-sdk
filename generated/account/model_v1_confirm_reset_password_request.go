/*
 * account-service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package account
// V1ConfirmResetPasswordRequest struct for V1ConfirmResetPasswordRequest
type V1ConfirmResetPasswordRequest struct {
	PasswordResetToken string `json:"password_reset_token,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
	Email string `json:"email,omitempty"`
	Hash string `json:"hash,omitempty"`
}
