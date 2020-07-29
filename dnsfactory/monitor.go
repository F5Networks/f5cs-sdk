package factory

const (
	MonitorTypeStandardHTTP  = "http_standard"
	MonitorTypeAdvancedHTTP  = "http_advanced"
	MonitorTypeAdvancedHTTPS = "https_advanced"
	MonitorTypeStandardTCP   = "tcp_standard"
	MonitorTypeAdvancedTCP   = "tcp_advanced"
	MonitorTypeAdvancedUDP   = "udp_advanced"
	MonitorTypeStandardICMP  = "icmp_standard"
)

var (
	noReceiveMonitorMap = map[string]bool{
		MonitorTypeStandardHTTP: true,
		MonitorTypeStandardTCP:  true,
		MonitorTypeStandardICMP: true,
	}
	noSendMonitorMap = map[string]bool{
		MonitorTypeStandardHTTP: true,
		MonitorTypeStandardTCP:  true,
		MonitorTypeStandardICMP: true,
	}
	sendRequiredMonitorMap = map[string]bool{
		MonitorTypeAdvancedHTTPS: true,
		MonitorTypeAdvancedHTTP:  true,
		MonitorTypeAdvancedUDP:   true,
		MonitorTypeAdvancedTCP:   true,
	}
	receiveRequiredMonitorMap = map[string]bool{
		MonitorTypeAdvancedHTTPS: true,
		MonitorTypeAdvancedHTTP:  true,
		MonitorTypeAdvancedUDP:   true,
		MonitorTypeAdvancedTCP:   true,
	}
)

// Monitor : configuration.gslb_service.monitors[]
type Monitor struct {
	// Field 0: DisplayName
	DisplayName *string `draft_validate:"required,max=64" json:"display_name,omitempty"`
	// Field 1: Remark
	Remark *string `draft_validate:"omitempty,max=128" json:"remark,omitempty"`

	// Field 2: MonitorType
	MonitorType string `draft_validate:"required,oneof=http_standard http_advanced https_advanced icmp_standard tcp_standard udp_advanced tcp_advanced" json:"monitor_type,omitempty"`

	// Field 3: TargetPort
	TargetPort *int `draft_validate:"omitempty,min=0,max=65535" json:"target_port,omitempty"`
	// Field 4: Send
	Send *string `draft_validate:"omitempty,min=0,max=2048" json:"send,omitempty"`
	// Field 5: Receive
	Receive *string `draft_validate:"omitempty,ReceiveRegex,min=0,max=2048" json:"receive,omitempty"`
}

func (m *Monitor) IsSendForbidden() bool {
	return noSendMonitorMap[m.MonitorType]
}
func (m *Monitor) IsReceiveForbidden() bool {
	return noReceiveMonitorMap[m.MonitorType]
}
func (m *Monitor) IsSendRequired() bool {
	return sendRequiredMonitorMap[m.MonitorType]
}
func (m *Monitor) IsReceiveRequired() bool {
	return receiveRequiredMonitorMap[m.MonitorType]
}

// IsAdvanced : check if monitor type is advanced.
func (m *Monitor) IsAdvanced() bool {
	switch m.MonitorType {
	// TODO: update when standard/advanced monitor_type updates are merged !145.
	case MonitorTypeAdvancedHTTPS, MonitorTypeAdvancedHTTP, MonitorTypeAdvancedTCP, MonitorTypeAdvancedUDP:
		return true
	default:
		return false
	}
}
