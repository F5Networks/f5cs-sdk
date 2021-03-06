/*
 * account-service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package account
import (
	"time"
)
// V1AccountMember struct for V1AccountMember
type V1AccountMember struct {
	AccountId string `json:"account_id,omitempty"`
	UserId string `json:"user_id,omitempty"`
	RoleId string `json:"role_id,omitempty"`
	User V1AccountMemberUser `json:"user,omitempty"`
	AccountName string `json:"account_name,omitempty"`
	RoleName string `json:"role_name,omitempty"`
	Level string `json:"level,omitempty"`
	SignupProvider string `json:"signup_provider,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
}
