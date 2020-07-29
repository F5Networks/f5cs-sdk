package factory

import "encoding/json"

const (
	VIRTUAL_SERVER_TYPE_CLOUD     = "cloud"
	VIRTUAL_SERVER_TYPE_BIGIP_LTM = "bigip-ltm"
)

type VirtualServer struct {
	// Field 0
	DisplayName *string `draft_validate:"required,max=64" json:"display_name,omitempty"`

	// Field 1
	Remark *string `draft_validate:"omitempty,max=128" json:"remark,omitempty"`

	// Field 2
	Port *int `draft_validate:"omitempty,min=1,max=65535" json:"port,omitempty"` //TODO check range?

	// Field 3
	Monitor *string `draft_validate:"omitempty" json:"monitor,omitempty"`

	// Field 4
	Address *string `draft_validate:"omitempty,ip,IPv6MappedIPv4,RestrictedSubnet" activation_validate:"omitempty,IPv6MappedIPv4,RestrictedSubnet" json:"address,omitempty"`

	// Field 5
	VirtualServerType *string `draft_validate:"required,oneof=cloud bigip-ltm" json:"virtual_server_type,omitempty"`

	// Field 6
	VipID *string `draft_validate:"omitempty" json:"vip_id,omitempty"`

	// Field 7
	TranslationAddress *string `draft_validate:"omitempty,ip,IPv6MappedIPv4,RestrictedSubnetVIP" json:"translation_address,omitempty"`
}

// GetAddress : Return Virtual Server address. Return empty string if address is nil
func (v *VirtualServer) GetAddress() string {
	if v.Address == nil {
		return ""
	}
	return *v.Address
}

// SetAddress :
func (v *VirtualServer) SetAddress(addr string) {
	v.Address = newStringPointer(addr)
}

// GetPort :  Return Virtual Server port. Return 0 if port is nil
func (v *VirtualServer) GetPort() int {
	if v.Port == nil {
		return 0
	}
	return *v.Port
}

// SetPort :
func (v *VirtualServer) SetPort(port int) {
	v.Port = newIntPointer(port)
}

// GetMonitor : Return Virtual Server Monitor. Return empty string if Monitor is nil
func (v *VirtualServer) GetMonitor() string {
	if v.Monitor == nil {
		return ""
	}
	return *v.Monitor
}

// SetMonitor :
func (v *VirtualServer) SetMonitor(monitor string) {
	v.Monitor = newStringPointer(monitor)
}

func (v *VirtualServer) Trim() {
	if v.VirtualServerType == nil {
		return
	}
	switch *v.VirtualServerType {
	case VIRTUAL_SERVER_TYPE_CLOUD:
		v.VipID = nil
		v.TranslationAddress = nil
	case VIRTUAL_SERVER_TYPE_BIGIP_LTM:
		v.Address = nil
		v.Port = nil
		v.Monitor = nil
	}
}

func (v *VirtualServer) IsBigIPLTM() bool {
	if v.VirtualServerType == nil {
		return false
	}
	return *v.VirtualServerType == VIRTUAL_SERVER_TYPE_BIGIP_LTM
}

// GetVipID : Return Virtual Server VipID. Return empty string if VipID is nil
func (v *VirtualServer) GetVipID() string {
	if v.VipID == nil {
		return ""
	}
	return *v.VipID
}

// SetVipID :
func (v *VirtualServer) SetVipID(vipID string) {
	v.VipID = newStringPointer(vipID)
}

// GetVirtualServerType : Return Virtual Server address. Return empty string if address is nil
func (v *VirtualServer) GetVirtualServerType() string {
	if v.VirtualServerType == nil {
		return ""
	}
	return *v.VirtualServerType
}

// SetVirtualServerType :
func (v *VirtualServer) SetVirtualServerType(vsType string) {
	v.VirtualServerType = newStringPointer(vsType)
}

// GetTranslationAddress : Return Translation address. Return empty string if Translation address is nil
func (v *VirtualServer) GetTranslationAddress() string {
	if v.TranslationAddress == nil {
		return ""
	}
	return *v.TranslationAddress
}

// SetTranslationAddress :
func (v *VirtualServer) SetTranslationAddress(translationAddress string) {
	v.TranslationAddress = newStringPointer(translationAddress)
}

// MarshalToString Helps debug process by transforming VS object into string
func (v *VirtualServer) MarshalToString() string {
	b, err := json.MarshalIndent(v, " ", " ")
	if err != nil {
		return ""
	}
	return string(b)
}

// GetAddressOrTranslationAddress returns the address used by the virtual server
// regardless of the VirtualServerType
func (v *VirtualServer) GetAddressOrTranslationAddress() string {
	if v.GetVirtualServerType() == VIRTUAL_SERVER_TYPE_CLOUD {
		return v.GetAddress()
	} else if v.GetVirtualServerType() == VIRTUAL_SERVER_TYPE_BIGIP_LTM {
		return v.GetTranslationAddress()
	} else {
		return ""
	}
}

// SetDisplayName :
func (v *VirtualServer) SetDisplayName(displayName string) {
	v.DisplayName = newStringPointer(displayName)
}
