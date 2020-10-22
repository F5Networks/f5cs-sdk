package factory

import (
	"encoding/json"
	"strings"
)

// DNSService is a struct representation of primaryDNS
type DNSService struct {
	// Field 0: For Primary DNS
	Remark *string `draft_validate:"omitempty,max=128" json:"remark,omitempty"`

	// Field 1: For Primary DNS
	Zone *string `draft_validate:"required" json:"zone,omitempty" conform:"trim,lower,rtrimdot"`

	// Field 2: For Primary DNS
	Records *map[string]RRSetsWrapper `draft_validate:"omitempty,dive" activation_validate:"omitempty,dive" json:"records,omitempty"`

	// Field 3: For Primary DNS
	TTL *int `draft_validate:"omitempty,min=0" json:"ttl,omitempty"`

	// Field 4: For Primary DNS
	NameServer *string `draft_validate:"omitempty,hostname" json:"primary_nameserver,omitempty"`

	// Field 5: For Primary DNS
	Admin *string `draft_validate:"omitempty,email" json:"admin,omitempty"` // dns-admin@f5cloudservices.com

	// Field 6: For Primary DNS
	Serial *int `draft_validate:"omitempty,min=0" json:"serial,omitempty"`

	// Field 7: For Primary DNS
	Refresh *int `draft_validate:"omitempty,min=0" json:"refresh,omitempty"` // default 24 hours

	// Field 8: For Primary DNS
	Retry *int `draft_validate:"omitempty,min=0" json:"retry,omitempty"` // default 2 hours

	// Field 9: For Primary DNS
	Expire *int `draft_validate:"omitempty,min=0" json:"expire,omitempty"` // default 100 hours

	// Field 10: For Primary DNS
	NegativeTTL *int `draft_validate:"omitempty,min=0" json:"negative_ttl,omitempty"` // 6
}

type RRSetsWrapper struct {
	RRSets []RRSet `draft_validate:"omitempty,dive" activation_validate:"omitempty,dive" json:"-"`
}

func (d *DNSService) GetZone() string {
	if d.Zone == nil {
		return ""
	}
	return *d.Zone
}

func (d *DNSService) GetServiceType() string {
	return "dns"
}

func (d *DNSService) GetEmptyService() Service {
	return &DNSService{}
}

func (rrWrapper *RRSetsWrapper) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		rrset := RRSet{}
		err := json.Unmarshal(data, &rrset)
		if err != nil {
			return err
		}
		rrWrapper.RRSets = []RRSet{rrset}
		return err

	} else {
		return json.Unmarshal(data, &rrWrapper.RRSets)
	}
}

func (rrsetsWrapper RRSetsWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(rrsetsWrapper.RRSets)
}

func (d *DNSService) MarshalToString() string {
	b, err := json.MarshalIndent(d, " ", " ")
	if err != nil {
		return ""
	}
	return string(b)
}

func (d *DNSService) AddRecord(recordName, recordType string, values ...string) *DNSService {

	rrSet := RRSet{Type: newStringPointer(recordType)}
	if len(values) > 0 {
		switch recordType {
		case RRTypeA:
			for i := 0; i < len(values); i++ {
				rrSet.Values = append(rrSet.Values, ARRSetValue{Address: newStringPointer(values[i])})
			}
		case RRTypeAAAA:
			for i := 0; i < len(values); i++ {
				rrSet.Values = append(rrSet.Values, AAAARRSetValue{Address: newStringPointer(values[i])})
			}
		case RRTypeNS:
			for i := 0; i < len(values); i++ {
				rrSet.Values = append(rrSet.Values, NSRRSetValue{NameServer: newStringPointer(values[i])})
			}
		case RRTypeMX:
			for i := 0; i < len(values); i++ {
				rrSet.Values = append(rrSet.Values, MXRRSetValue{Domain: newStringPointer(values[0])})
			}
		case RRTypeCNAME:
			rrSet.Value = CNAMERRSetValue{Domain: newStringPointer(values[0])}
		case RRTypeTXT:
			for i := 0; i < len(values); i++ {
				rrSet.Values = append(rrSet.Values, TXTRRSetValue{Text: newStringPointer(values[0])})
			}
		default:
			rrSet = RRSet{}
		}
	}

	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	rrSetsWrapper := RRSetsWrapper{}
	if currentRRSets, ok := (*d.Records)[recordName]; ok {
		rrSetsWrapper.RRSets = append(currentRRSets.RRSets, rrSet)
	} else {
		rrSetsWrapper.RRSets = []RRSet{rrSet}
	}

	(*d.Records)[recordName] = rrSetsWrapper
	return d
}

func (d *DNSService) GetRecord(recordName string) RRSetsWrapper {
	if d.Records == nil {
		return RRSetsWrapper{}
	}
	return (*d.Records)[recordName]
}

func (d *DNSService) AddARecord(recordName string, address ...string) *DNSService {
	rrSet := RRSet{Type: newStringPointer(RRTypeA)}

	for i := 0; i < len(address); i++ {
		rrSet.Values = append(rrSet.Values, ARRSetValue{Address: newStringPointer(address[i])})
	}

	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	currentSet := d.GetARecordSet(recordName)
	currentSet.Values = append(currentSet.Values, rrSet.Values...)
	d.SetARecordSet(recordName, currentSet)
	return d
}

func (d *DNSService) AddAAAARecord(recordName string, address ...string) *DNSService {
	rrSet := RRSet{Type: newStringPointer(RRTypeAAAA)}
	for i := 0; i < len(address); i++ {
		rrSet.Values = append(rrSet.Values, AAAARRSetValue{Address: newStringPointer(address[i])})
	}
	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	currentSet := d.GetAAAARecordSet(recordName)
	currentSet.Values = append(currentSet.Values, rrSet.Values...)
	d.SetAAAARecordSet(recordName, currentSet)
	return d
}

func (d *DNSService) AddNSRecord(recordName string, nameserver ...string) *DNSService {
	rrSet := RRSet{Type: newStringPointer(RRTypeNS)}
	for i := 0; i < len(nameserver); i++ {
		rrSet.Values = append(rrSet.Values, NSRRSetValue{NameServer: newStringPointer(nameserver[i])})
	}
	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	currentSet := d.GetNSRecordSet(recordName)
	currentSet.Values = append(currentSet.Values, rrSet.Values...)
	d.SetNSRecordSet(recordName, currentSet)
	return d
}

func (d *DNSService) AddCNAMERecord(recordName string, domain ...string) *DNSService {
	rrSet := RRSet{Type: newStringPointer(RRTypeCNAME)}
	for i := 0; i < len(domain); i++ {
		rrSet.Values = append(rrSet.Values, CNAMERRSetValue{Domain: newStringPointer(domain[i])})
	}

	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	currentSet := d.GetCNAMERecordSet(recordName)
	currentSet.Values = append(currentSet.Values, rrSet.Values...)
	d.SetCNAMERecordSet(recordName, currentSet)
	return d
}

func (d *DNSService) AddMXRecord(recordName string, mxSetValues ...MXRRSetValue) *DNSService {
	rrSet := RRSet{
		Type: newStringPointer(RRTypeMX),
	}

	for _, setValue := range mxSetValues {
		rrSet.Values = append(rrSet.Values, setValue)
	}

	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	currentSet := d.GetMXRecordSet(recordName)
	currentSet.Values = append(currentSet.Values, rrSet.Values...)
	d.SetMXRecordSet(recordName, currentSet)
	return d
}

func (d *DNSService) AddTXTRecord(recordName string, text ...string) *DNSService {
	rrSet := RRSet{
		Type: newStringPointer(RRTypeTXT),
	}
	for i := 0; i < len(text); i++ {
		rrSet.Values = append(rrSet.Values, TXTRRSetValue{Text: newStringPointer(text[i])})
	}
	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	currentSet := d.GetTXTRecordSet(recordName)
	currentSet.Values = append(currentSet.Values, rrSet.Values...)
	d.SetTXTRecordSet(recordName, currentSet)
	return d
}

func (d *DNSService) RemoveRecord(recordName string, recordType ...string) *DNSService {
	if d.Records == nil {
		return d
	}
	var rrSets []RRSet
	if rrSetWrapper, ok := (*d.Records)[recordName]; ok {
		rrSets = rrSetWrapper.RRSets
	} else {
		return d
	}

	if len(recordType) == 0 {
		// remove all
		delete(*d.Records, recordName)
	} else {
		// remove specified types
		newRRSets := []RRSet{}
		removeTypes := strings.Join(recordType, ",")
		for _, rrSet := range rrSets {
			if rrSet.Type == nil {
				continue
			}
			if !strings.Contains(removeTypes, *rrSet.Type) {
				newRRSets = append(newRRSets, rrSet)
			}
		}
		(*d.Records)[recordName] = RRSetsWrapper{newRRSets}
	}

	return d
}

func (d *DNSService) GetARecordSet(recordName string) RRSet {
	if record, ok := (*d.Records)[recordName]; ok {
		for _, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeA {
				return rrset
			}
		}
	}
	return RRSet{Type: newStringPointer(RRTypeA)}
}

func (d *DNSService) SetARecordSet(recordName string, newRrSet RRSet) *DNSService {
	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	if record, ok := (*d.Records)[recordName]; ok {
		for n, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeA {
				(*d.Records)[recordName].RRSets[n] = newRrSet
				return d
			}
		}
	}
	rrSetsWrapper := (*d.Records)[recordName]
	rrSetsWrapper.RRSets = append(rrSetsWrapper.RRSets, newRrSet)
	(*d.Records)[recordName] = rrSetsWrapper
	return d
}

func (d *DNSService) GetAAAARecordSet(recordName string) RRSet {
	if record, ok := (*d.Records)[recordName]; ok {
		for _, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeAAAA {
				return rrset
			}
		}
	}
	return RRSet{Type: newStringPointer(RRTypeAAAA)}
}

func (d *DNSService) SetAAAARecordSet(recordName string, newRrSet RRSet) *DNSService {
	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	if record, ok := (*d.Records)[recordName]; ok {
		for n, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeAAAA {
				(*d.Records)[recordName].RRSets[n] = newRrSet
				return d
			}
		}
	}
	rrSetsWrapper := (*d.Records)[recordName]
	rrSetsWrapper.RRSets = append(rrSetsWrapper.RRSets, newRrSet)
	(*d.Records)[recordName] = rrSetsWrapper
	return d
}

func (d *DNSService) GetMXRecordSet(recordName string) RRSet {
	if record, ok := (*d.Records)[recordName]; ok {
		for _, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeMX {
				return rrset
			}
		}
	}
	return RRSet{Type: newStringPointer(RRTypeMX)}
}

func (d *DNSService) SetMXRecordSet(recordName string, newRrSet RRSet) *DNSService {
	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	if record, ok := (*d.Records)[recordName]; ok {
		for n, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeMX {
				(*d.Records)[recordName].RRSets[n] = newRrSet
				return d
			}
		}
	}
	rrSetsWrapper := (*d.Records)[recordName]
	rrSetsWrapper.RRSets = append(rrSetsWrapper.RRSets, newRrSet)
	(*d.Records)[recordName] = rrSetsWrapper
	return d
}

func (d *DNSService) GetNSRecordSet(recordName string) RRSet {
	if record, ok := (*d.Records)[recordName]; ok {
		for _, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeNS {
				return rrset
			}
		}
	}
	return RRSet{Type: newStringPointer(RRTypeNS)}
}

func (d *DNSService) SetNSRecordSet(recordName string, rrSet RRSet) *DNSService {
	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	if record, ok := (*d.Records)[recordName]; ok {
		for n, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeNS {
				(*d.Records)[recordName].RRSets[n] = rrSet
				return d
			}
		}
	}
	rrSetsWrapper := (*d.Records)[recordName]
	rrSetsWrapper.RRSets = append(rrSetsWrapper.RRSets, rrSet)
	(*d.Records)[recordName] = rrSetsWrapper
	return d
}

func (d *DNSService) GetCNAMERecordSet(recordName string) RRSet {
	if record, ok := (*d.Records)[recordName]; ok {
		for _, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeCNAME {
				return rrset
			}
		}
	}
	return RRSet{Type: newStringPointer(RRTypeCNAME)}
}

func (d *DNSService) SetCNAMERecordSet(recordName string, newRrSet RRSet) *DNSService {
	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	if record, ok := (*d.Records)[recordName]; ok {
		for n, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeCNAME {
				(*d.Records)[recordName].RRSets[n] = newRrSet
				return d
			}
		}
	}
	rrSetsWrapper := (*d.Records)[recordName]
	rrSetsWrapper.RRSets = append(rrSetsWrapper.RRSets, newRrSet)
	(*d.Records)[recordName] = rrSetsWrapper
	return d
}

func (d *DNSService) GetTXTRecordSet(recordName string) RRSet {
	if record, ok := (*d.Records)[recordName]; ok {
		for _, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeTXT {
				return rrset
			}
		}
	}
	return RRSet{Type: newStringPointer(RRTypeTXT)}
}

func (d *DNSService) SetTXTRecordSet(recordName string, newRrSet RRSet) *DNSService {
	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	if record, ok := (*d.Records)[recordName]; ok {
		for n, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeTXT {
				(*d.Records)[recordName].RRSets[n] = newRrSet
				return d
			}
		}
	}
	rrSetsWrapper := (*d.Records)[recordName]
	rrSetsWrapper.RRSets = append(rrSetsWrapper.RRSets, newRrSet)
	(*d.Records)[recordName] = rrSetsWrapper
	return d
}
