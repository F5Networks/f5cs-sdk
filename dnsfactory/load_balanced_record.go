package factory

import ()

// LoadBalancedRecord : configuration.gslb_service.load_balanced_records[]
type LoadBalancedRecord struct {
	// Field 0
	DisplayName *string `draft_validate:"required,max=64" json:"display_name,omitempty"`
	// Field 1
	Remark *string `draft_validate:"omitempty,max=255" json:"remark,omitempty"`
	// Field 2
	Enable *bool `draft_validate:"omitempty" json:"enable,omitempty"`

	// Field 3
	RRType *string `draft_validate:"required,oneof=AAAA A MX CNAME" json:"rr_type,omitempty"`

	// Field 4
	Aliases []string `draft_validate:"required,gte=1,lte=256,unique" json:"aliases,omitempty" conform:"lower"`
	// Field 5
	ProximityRules []ProximityRule `draft_validate:"required,gte=0,lte=64,unique,dive" json:"proximity_rules,omitempty"`

	// Field 6
	Persistence *bool `json:"persistence,omitempty"`

	// Field 7
	PersistCidrIpv4 *int `draft_validate:"omitempty,min=0,max=32" json:"persist_cidr_ipv4,omitempty"`
	// Field 8
	PersistCidrIpv6 *int `draft_validate:"omitempty,min=0,max=128" json:"persist_cidr_ipv6,omitempty"`
	// Field 9
	PersistTtl *int `draft_validate:"omitempty,min=0,max=86400" json:"persistence_ttl,omitempty"`
}

// ProximityRule : configuration.gslb_service.load_balanced_records[].proximity_rules[]
type ProximityRule struct {
	// Field 0
	// Region: "Either 'global' or name of declared region"
	Region *string `draft_validate:"required" json:"region,omitempty"`

	// Field 1
	Pool *string `draft_validate:"required" json:"pool,omitempty"`

	// Field 2
	Score *int `draft_validate:"required,min=1,max=32767" json:"score,omitempty"`
}

func (lbr *LoadBalancedRecord) Enabled() bool {
	if lbr.Enable == nil || !*lbr.Enable {
		return false
	}
	return true
}
