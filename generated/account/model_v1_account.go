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
// V1Account struct for V1Account
type V1Account struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	// Optional parent account ID. If not empty, this account is a sub-account.
	ParentAccountId string `json:"parent_account_id,omitempty"`
	Status string `json:"status,omitempty"`
	Level string `json:"level,omitempty"`
	SignupProvider string `json:"signup_provider,omitempty"`
	Address V1Address `json:"address,omitempty"`
	Phone string `json:"phone,omitempty"`
	Compliant bool `json:"compliant,omitempty"`
	CatalogItems []V1AccountCatalogItem `json:"catalog_items,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
	DeleteTime time.Time `json:"delete_time,omitempty"`
}