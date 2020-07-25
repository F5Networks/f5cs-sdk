package factory

import (
	"encoding/json"

	loglib "log"
)

type ServiceRequests struct {
	Subscriptions []ServiceRequest `json:"subscriptions"`
}

// This struct exists for unmarshalling purposes only. No Validation is performed on this struct.
type ServiceRequest struct {
	// Field 0
	SubscriptionId string `json:"subscription_id,omitempty"`

	// Field 1
	ServiceInstanceId string `json:"service_instance_id,omitempty"`

	// Field 2
	ServiceInstanceName string `json:"service_instance_name,omitempty"`

	// Field 3
	Configuration *Configuration `json:"configuration,omitempty"`

	//Field 4
	AccountId string `json:"account_id,omitempty"`

	// The following fields are used in testing.
	UserID string `json:"user_id,omitempty"`

	CatalogID string `json:"catalog_id,omitempty"`

	Deleted bool `json:"deleted,omitempty"`

	ServiceType string `json:"service_type,omitempty"`

	CreateTime string `json:"create_time,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`
	CancelTime string `json:"cancel_time,omitempty"`
	EndTime    string `json:"end_time,omitempty"`
}

func (s ServiceRequest) ToJSON() string {
	acctID := s.AccountId
	s.AccountId = ""
	jbytes, err := json.MarshalIndent(s, " ", " ")
	if err != nil {
		loglib.Printf("failed to marshal service request to string err %v", err.Error())
		return ""
	}
	s.AccountId = acctID
	return string(jbytes)
}
