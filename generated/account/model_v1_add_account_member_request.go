/*
 * account-service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package account
// V1AddAccountMemberRequest struct for V1AddAccountMemberRequest
type V1AddAccountMemberRequest struct {
	AccountId string `json:"account_id,omitempty"`
	UserId string `json:"user_id,omitempty"`
	RoleId string `json:"role_id,omitempty"`
	InviteToken string `json:"invite_token,omitempty"`
}
