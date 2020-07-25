package factory

// Configuration : struct representation of common_service_request.configuration
type Configuration struct {
	// Field 0
	GslbService *GSLBService `draft_validate:"required,dive" activation_validate:"dive" json:"gslb_service,omitempty"`

	// Field 1
	NameServers *[]NameServer `json:"nameservers,omitempty"`

	// Field 2
	Details *ConfigDetails `json:"details,omitempty"`
}

func (c *Configuration) GetNameServers() []NameServer {
	if c.NameServers == nil {
		return []NameServer{}
	}
	return *c.NameServers
}

type NameServer struct {
	Ipv4 string `json:"ipv4"`
	Ipv6 string `json:"ipv6"`
	Name string `json:"name"`
}

type ConfigDetails struct {
	LBRMetadata *LBRMetadata              `json:"lbr_metadata,omitempty"`
	PoolsHealth *map[string]MembersHealth `json:"pools_health,omitempty"`
}

type LBRMetadata struct {
	ActiveLBRCount                       int `json:"active_lbr_count,omitempty"`
	ActivePoolMemberAdvancedMonitorCount int `json:"active_pool_member_advanced_monitor_count,omitempty"`
	ActivePoolMemberStandardMonitorCount int `json:"active_pool_member_standard_monitor_count,omitempty"`
}

type MembersHealth struct {
	MembersHealth *[]MemberHealth `json:"members_health,omitempty"`
}

type MemberHealth struct {
	Details       []MemberDetails `json:"details,omitempty"`
	Health        string          `json:"health,omitempty"`
	Monitor       string          `json:"monitor,omitempty"`
	VirtualServer string          `json:"virtual_server,omitempty"`
}

type MemberDetails struct {
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}
