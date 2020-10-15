package factory

import (
	"encoding/json"
	"strings"
)

const (
	// Putting here for reference but A, AAAA, MX, CNAME
	// are declared in pool.go in the same package
	// RRTypeA     = "A"
	// RRTypeAAAA  = "AAAA"
	// RRTypeMX    = "MX"
	// RRTypeCNAME = "CNAME"

	// RRTypeSOA => SOA Record will be on its own under DNSService
	RRTypeNS  = "NS"
	RRTypeTXT = "TXT"
)

type RRSet struct {
	// Field 0
	Type *string `draft_validate:"required,oneof=AAAA A MX CNAME NS TXT" json:"type,omitempty"`

	// Field 1
	TTL *int `draft_validate:"omitempty,min=0" json:"ttl,omitempty"`

	// Field 2
	Value RRSetValue `draft_validate:"omitempty,dive" json:"value,omitempty"` //0-255 characters

	// Field 3
	// Depending on the type, value will be array of strings, or array of structs
	Values []RRSetValue `draft_validate:"omitempty,dive" json:"values,omitempty"`
}

type RRSetValue interface {
	GetRRType() string
}

type NSRRSetValue struct {
	NameServer *string `draft_validate:"omitempty,hostname" json:",omitempty"`
}

func (ns NSRRSetValue) GetRRType() string {
	return RRTypeNS
}

type CNAMERRSetValue struct {
	Domain *string `draft_validate:"omitempty,hostname" json:",omitempty" conform:"trim,lower,rtrimdot"`
}
type MXRRSetValue struct {
	Domain   *string `draft_validate:"omitempty,hostname" json:"domain,omitempty" conform:"trim,lower,rtrimdot"`
	Priority *int    `draft_validate:"omitempty" json:"priority,omitempty"` // 0-255
}

// A or AAAA record
type AorAAAARRSetValue struct {
	Address *string `draft_validate:"omitempty,ip,IPv6MappedIPv4" activation_validate:"omitempty,IPv6MappedIPv4" json:",omitempty"`
}

type TXTRRSetValue struct {
	Text *string `draft_validate:"omitempty,min=1,max=255" json:",omitempty"`
}

func (r *RRSet) MarshalToString() string {
	b, err := json.MarshalIndent(r, " ", " ")
	if err != nil {
		return ""
	}
	return string(b)

}

func (rrset *RRSet) UnmarshalJSON(data []byte) error {
	// get a map of fields in the incoming data
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(data, &objMap)

	// map them all out to their corresponging struct keys

	if objMap["type"] != nil {
		err = json.Unmarshal(*objMap["type"], &rrset.Type)
		if err != nil {
			return err
		}
	}

	if objMap["ttl"] != nil {
		err = json.Unmarshal(*objMap["ttl"], &rrset.TTL)
		if err != nil {
			return err
		}
	} else {
		rrset.TTL = newIntPointer(86400)
	}

	switch strings.ToUpper(ssp(rrset.Type)) {
	case RRTypeNS:
		var NS []string
		var NSRRSet []RRSetValue
		if objMap["values"] != nil {
			err = json.Unmarshal(*objMap["values"], &NS)
			if err != nil {
				return err
			}
			for i := 0; i < len(NS); i++ {
				nsVal := NSRRSetValue{NameServer: newStringPointer(NS[i])}
				NSRRSet = append(NSRRSet, nsVal)
			}
			rrset.Values = NSRRSet
		}
	}
	return err
}

func (rrset *RRSet) MarshalJSON() ([]byte, error) {
	type AliasRRSet RRSet
	switch strings.ToUpper(ssp(rrset.Type)) {
	case RRTypeNS:
		var nsArray []string
		castingOK := true
		for _, values := range rrset.Values {
			if nsVal, ok := values.(NSRRSetValue); ok {
				nsArray = append(nsArray, ssp(nsVal.NameServer))
			} else {
				castingOK = false
				break
			}
		}
		// if there was casting problem, fall back to json.Marshaler
		if !castingOK {
			break
		}
		tempStr := &struct {
			*AliasRRSet
			Values []string `json:"values,omitempty"`
		}{
			AliasRRSet: (*AliasRRSet)(rrset),
			Values:     nsArray,
		}
		return json.Marshal(tempStr)
	}

	return json.Marshal(&struct {
		*AliasRRSet
	}{
		AliasRRSet: (*AliasRRSet)(rrset),
	})
}
