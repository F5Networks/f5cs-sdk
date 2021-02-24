package factory

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	APEXRecordName = ""
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
	RRSets []RRSet `draft_validate:"omitempty,dive" activation_validate:"omitempty,dive"`
}

// GetTypes returns the set of rrtypes contained in the record.
func (rrw *RRSetsWrapper) GetTypes() []string {
	rrtypes := []string{}
	rrtypesmap := map[string]bool{}
	for _, rrset := range rrw.RRSets {
		if rrsettype := ssp(rrset.Type); rrsettype != "" {
			rrtypesmap[rrsettype] = true
		}
	}
	for rrsettype := range rrtypesmap {
		rrtypes = append(rrtypes, rrsettype)
	}
	return rrtypes
}

func (d *DNSService) GetRemark() string {
	if d.Remark == nil {
		return ""
	}
	return *d.Remark
}

func (d *DNSService) SetRemark(remark string) *DNSService {
	d.Remark = &remark
	return d
}

func (d *DNSService) GetZone() string {
	if d.Zone == nil {
		return ""
	}
	return *d.Zone
}

func (d *DNSService) SetZone(zone string) *DNSService {
	d.Zone = &zone
	return d
}

func (d *DNSService) GetNameserver() string {
	if d.NameServer == nil {
		return ""
	}
	return *d.NameServer
}

func (d *DNSService) SetNameserver(nameserver string) *DNSService {
	d.NameServer = &nameserver
	return d
}

func (d *DNSService) GetAdmin() string {
	if d.Admin == nil {
		return ""
	}
	return *d.Admin
}

func (d *DNSService) SetAdmin(admin string) *DNSService {
	d.Admin = &admin
	return d
}

func (d *DNSService) GetSerial() int {
	if d.Serial == nil {
		return 0
	}
	return *d.Serial
}

func (d *DNSService) SetSerial(serial int) *DNSService {
	d.Serial = &serial
	return d
}

func (d *DNSService) GetRefresh() int {
	if d.Refresh == nil {
		return 0
	}
	return *d.Refresh
}

func (d *DNSService) SetRefresh(refresh int) *DNSService {
	d.Refresh = &refresh
	return d
}

func (d *DNSService) GetRetry() int {
	if d.Retry == nil {
		return 0
	}
	return *d.Retry
}

func (d *DNSService) SetRetry(retry int) *DNSService {
	d.Retry = &retry
	return d
}

func (d *DNSService) GetExpire() int {
	if d.Expire == nil {
		return 0
	}
	return *d.Expire
}

func (d *DNSService) SetExpire(expire int) *DNSService {
	d.Expire = &expire
	return d
}

func (d *DNSService) GetTTL() int {
	if d.Zone == nil {
		return 0
	}
	return *d.TTL
}

func (d *DNSService) SetTTL(ttl int) *DNSService {
	d.TTL = &ttl
	return d
}

func (d *DNSService) GetNegativeTTL() int {
	if d.Zone == nil {
		return 0
	}
	return *d.TTL
}

func (d *DNSService) SetNegativeTTL(negativeTTL int) *DNSService {
	d.TTL = &negativeTTL
	return d
}

func (d *DNSService) GetServiceType() string {
	return "dns"
}

func (d *DNSService) GetEmptyService() Service {
	return &DNSService{}
}

func (d *DNSService) GetRecordsCount() int {
	totalCount := 0

	if d.Records != nil {
		for name, rrwrapper := range *d.Records {
			for _, rrset := range rrwrapper.RRSets {
				if name == APEXRecordName && ssp(rrset.Type) == RRTypeNS {
					// don't count APEX NS record
					continue
				}
				if rrset.Value != nil {
					totalCount++
				}
				totalCount = totalCount + len(rrset.Values)
			}
		}
	}
	return totalCount
}

func (d *DNSService) GetRecordsNames() map[string]bool {
	recordNames := make(map[string]bool, 0)
	if d.Records != nil {
		for name := range *d.Records {
			recordNames[name] = true
		}
	}
	return recordNames
}

func (rrsetsWrapper *RRSetsWrapper) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		rrset := RRSet{}
		err := json.Unmarshal(data, &rrset)
		if err != nil {
			return err
		}
		rrsetsWrapper.RRSets = []RRSet{rrset}
		return err

	} else {
		return json.Unmarshal(data, &rrsetsWrapper.RRSets)
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
		case RRTypeSRV:
			for i := 0; i < len(values); i++ {
				rrSet.Values = append(rrSet.Values, SRVRRSetValue{Target: newStringPointer(values[0])})
			}
			// case RRTypeCAA:
		case RRTypeALIAS:
			rrSet.Value = ALIASRRSetValue{Domain: newStringPointer(values[0])}
		case RRTypePTR:
			for i := 0; i < len(values); i++ {
				rrSet.Values = append(rrSet.Values, PTRRRSetValue{Domain: newStringPointer(values[0])})
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

func (d *DNSService) GetCnameRecords() map[string]RRSetsWrapper {
	records := map[string]RRSetsWrapper{}
	if d.Records == nil {
		return records
	}
	for recordName, record := range *d.Records {
		cnameRRsets := []RRSet{}
		for _, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeCNAME {
				cnameRRsets = append(cnameRRsets,
					RRSet{
						Type:  rrset.Type,
						TTL:   rrset.TTL,
						Value: rrset.Value,
					})
			}
		}
		if len(cnameRRsets) > 0 {
			records[recordName] = RRSetsWrapper{cnameRRsets}
		}
	}
	return records
}

type Matches struct {
	exactMatches    map[string]bool
	wildcardMatches map[string]bool
}

func (m Matches) IsCnameExactMatch() bool {
	return m.exactMatches[RRTypeCNAME]
}

func (m Matches) IsCnameMatch() bool {
	return m.IsCnameExactMatch() || m.IsCnameWildcardMatch()
}

func (m Matches) IsCnameWildcardMatch() bool {
	return m.wildcardMatches[RRTypeCNAME]
}

func (m Matches) IsNonCnameExactMatch() bool {
	for rrtype := range m.exactMatches {
		if rrtype != RRTypeCNAME {
			return true
		}
	}
	return false
}

func (d *DNSService) MatchRecord(domain string) (Matches, error) {
	matches := Matches{
		exactMatches:    map[string]bool{},
		wildcardMatches: map[string]bool{},
	}
	if d.Records == nil {
		return Matches{}, nil
	}
	zone := d.GetZone()
	if zone == "" {
		return Matches{}, nil
	}
	lowerDomain := strings.ToLower(domain)
	domainPhrase, err := newPhrase(lowerDomain)
	if err != nil {
		return Matches{}, err
	}
	for recordName, record := range *d.Records {
		if recordName == APEXRecordName {
			continue
		}
		checkDomain := fmt.Sprintf("%s.%s", recordName, zone)
		checkDomainPhrase, err := newPhrase(checkDomain)
		if err != nil {
			return Matches{}, err
		}
		if intersect, err := wildcardIntersect(checkDomainPhrase, domainPhrase); err != nil {
			return Matches{}, err
		} else if intersect {
			exactMatch := checkDomain == lowerDomain
			recordTypes := record.GetTypes()
			for _, rtype := range recordTypes {
				if exactMatch {
					matches.exactMatches[rtype] = true
				} else {
					matches.wildcardMatches[rtype] = true
				}
			}
		}
	}
	return matches, nil
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

func (d *DNSService) AddCNAMERecord(recordName string, domain string) *DNSService {
	rrSet := RRSet{Type: newStringPointer(RRTypeCNAME)}
	rrSet.Value = CNAMERRSetValue{Domain: newStringPointer(domain)}

	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	currentSet := d.GetCNAMERecordSet(recordName)
	currentSet.Value = rrSet.Value
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

func (d *DNSService) AddSRVRecord(recordName string, srvSetValues ...SRVRRSetValue) *DNSService {
	rrSet := RRSet{
		Type: newStringPointer(RRTypeSRV),
	}

	for _, setValue := range srvSetValues {
		rrSet.Values = append(rrSet.Values, setValue)
	}

	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	currentSet := d.GetSRVRecordSet(recordName)
	currentSet.Values = append(currentSet.Values, rrSet.Values...)
	d.SetSRVRecordSet(recordName, currentSet)
	return d
}

func (d *DNSService) AddCAARecord(recordName string, caaSetValues ...CAARRSetValue) *DNSService {
	rrSet := RRSet{
		Type: newStringPointer(RRTypeCAA),
	}

	for _, setValue := range caaSetValues {
		rrSet.Values = append(rrSet.Values, setValue)
	}

	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	currentSet := d.GetCAARecordSet(recordName)
	currentSet.Values = append(currentSet.Values, rrSet.Values...)
	d.SetCAARecordSet(recordName, currentSet)
	return d
}

func (d *DNSService) AddALIASRecord(recordName string, domain string) *DNSService {
	rrSet := RRSet{Type: newStringPointer(RRTypeALIAS)}
	rrSet.Value = ALIASRRSetValue{Domain: newStringPointer(domain)}

	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	currentSet := d.GetALIASRecordSet(recordName)
	currentSet.Value = rrSet.Value
	d.SetALIASRecordSet(recordName, currentSet)
	return d
}

func (d *DNSService) AddPTRRecord(recordName string, domain ...string) *DNSService {
	rrSet := RRSet{Type: newStringPointer(RRTypePTR)}

	for i := 0; i < len(domain); i++ {
		rrSet.Values = append(rrSet.Values, PTRRRSetValue{Domain: newStringPointer(domain[i])})
	}

	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	currentSet := d.GetPTRRecordSet(recordName)
	currentSet.Values = append(currentSet.Values, rrSet.Values...)
	d.SetPTRRecordSet(recordName, currentSet)
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

func (d *DNSService) GetSRVRecordSet(recordName string) RRSet {
	if record, ok := (*d.Records)[recordName]; ok {
		for _, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeSRV {
				return rrset
			}
		}
	}
	return RRSet{Type: newStringPointer(RRTypeSRV)}
}

func (d *DNSService) SetSRVRecordSet(recordName string, newRrSet RRSet) *DNSService {
	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	if record, ok := (*d.Records)[recordName]; ok {
		for n, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeSRV {
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

func (d *DNSService) GetCAARecordSet(recordName string) RRSet {
	if record, ok := (*d.Records)[recordName]; ok {
		for _, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeCAA {
				return rrset
			}
		}
	}
	return RRSet{Type: newStringPointer(RRTypeCAA)}
}

func (d *DNSService) SetCAARecordSet(recordName string, newRrSet RRSet) *DNSService {
	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	if record, ok := (*d.Records)[recordName]; ok {
		for n, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeCAA {
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

func (d *DNSService) GetALIASRecordSet(recordName string) RRSet {
	if record, ok := (*d.Records)[recordName]; ok {
		for _, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeALIAS {
				return rrset
			}
		}
	}
	return RRSet{Type: newStringPointer(RRTypeALIAS)}
}

func (d *DNSService) SetALIASRecordSet(recordName string, newRrSet RRSet) *DNSService {
	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	if record, ok := (*d.Records)[recordName]; ok {
		for n, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypeALIAS {
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

func (d *DNSService) GetPTRRecordSet(recordName string) RRSet {
	if record, ok := (*d.Records)[recordName]; ok {
		for _, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypePTR {
				return rrset
			}
		}
	}
	return RRSet{Type: newStringPointer(RRTypePTR)}
}

func (d *DNSService) SetPTRRecordSet(recordName string, newRrSet RRSet) *DNSService {
	if d.Records == nil {
		d.Records = &map[string]RRSetsWrapper{}
	}
	if record, ok := (*d.Records)[recordName]; ok {
		for n, rrset := range record.RRSets {
			if ssp(rrset.Type) == RRTypePTR {
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
