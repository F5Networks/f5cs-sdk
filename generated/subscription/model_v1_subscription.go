/*
 * Subscription Service
 *
 * Manages subscriptions tied to a specific service in F5 Cloud Services.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package subscription
import (
	"time"
)
// V1Subscription struct for V1Subscription
type V1Subscription struct {
	SubscriptionId string `json:"subscription_id,omitempty"`
	AccountId string `json:"account_id,omitempty"`
	UserId string `json:"user_id,omitempty"`
	CatalogId string `json:"catalog_id,omitempty"`
	ServiceInstanceId string `json:"service_instance_id,omitempty"`
	Status V1Status `json:"status,omitempty"`
	ServiceInstanceName string `json:"service_instance_name,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	ServiceType string `json:"service_type,omitempty"`
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
	CancelTime time.Time `json:"cancel_time,omitempty"`
	EndTime time.Time `json:"end_time,omitempty"`
}
