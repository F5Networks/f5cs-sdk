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
}

type NSRRSetValue struct {
	NameServer *string `draft_validate:"omitempty,hostname"`
}

func (rrsetValue NSRRSetValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(rrsetValue.NameServer)
}

type CNAMERRSetValue struct {
	Domain *string `draft_validate:"omitempty,hostname" conform:"trim,lower,rtrimdot"`
}

func (rrsetValue CNAMERRSetValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(rrsetValue.Domain)
}

type MXRRSetValue struct {
	Domain   *string `draft_validate:"omitempty,hostname" json:"domain,omitempty" conform:"trim,lower,rtrimdot"`
	Priority *int    `draft_validate:"omitempty" json:"priority,omitempty"` // 0-255
}

type ARRSetValue struct {
	Address *string `draft_validate:"omitempty,ipv4"`
}

func (rrsetValue ARRSetValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(rrsetValue.Address)
}

type AAAARRSetValue struct {
	Address *string `draft_validate:"omitempty,ipv6"`
}

func (rrsetValue AAAARRSetValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(rrsetValue.Address)
}

type TXTRRSetValue struct {
	Text *string `draft_validate:"omitempty,min=0,max=512"`
}

func (rrsetValue TXTRRSetValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(rrsetValue.Text)
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

	// map them all out to their corresponding struct keys
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
		rrset.TTL = newIntPointer(3600)
	}

	setValue := func(callback func(*string) RRSetValue) error {
		if objMap["value"] == nil {
			return nil
		}
		var value string
		err := json.Unmarshal(*objMap["value"], &value)
		if err != nil {
			return err
		}
		rrset.Value = callback(&value)
		return nil
	}

	setValues := func(callback func(*string) RRSetValue) error {
		if objMap["values"] == nil {
			return nil
		}
		var values []string
		err := json.Unmarshal(*objMap["values"], &values)
		if err != nil {
			return err
		}
		for _, value := range values {
			val := value
			rrset.Values = append(rrset.Values, callback(&val))
		}
		return nil
	}

	setValueAndValues := func(callback func(*string) RRSetValue) error {
		err := setValue(callback)
		if err != nil {
			return err
		}
		err = setValues(callback)
		if err != nil {
			return err
		}
		return nil
	}

	switch strings.ToUpper(ssp(rrset.Type)) {
	case RRTypeNS:
		err = setValueAndValues(func(value *string) RRSetValue {
			return NSRRSetValue{NameServer: value}
		})
	case RRTypeA:
		err = setValueAndValues(func(value *string) RRSetValue {
			return ARRSetValue{Address: value}
		})
	case RRTypeAAAA:
		err = setValueAndValues(func(value *string) RRSetValue {
			return AAAARRSetValue{Address: value}
		})
	case RRTypeCNAME:
		err = setValue(func(value *string) RRSetValue {
			return CNAMERRSetValue{Domain: value}
		})
	case RRTypeMX:
		if objMap["value"] != nil {
			var mxStruct MXRRSetValue
			err = json.Unmarshal(*objMap["value"], &mxStruct)
			if err != nil {
				return err
			}
			rrset.Value = mxStruct
		}

		if objMap["values"] != nil {
			var mxArray []MXRRSetValue
			var MXRRSet []RRSetValue
			err = json.Unmarshal(*objMap["values"], &mxArray)
			if err != nil {
				return err
			}
			for i := 0; i < len(mxArray); i++ {
				MXRRSet = append(MXRRSet, mxArray[i])
			}
			rrset.Values = MXRRSet
		}
	case RRTypeTXT:
		err = setValueAndValues(func(value *string) RRSetValue {
			return TXTRRSetValue{Text: value}
		})
	}
	return err
}

func (rrset RRSet) MarshalJSON() ([]byte, error) {
	var value RRSetValue
	var values []RRSetValue

	// if not CNAME, marshal value inside the values field
	if rrset.Type != nil {
		if *rrset.Type == RRTypeCNAME {
			value = rrset.Value
			values = nil
		} else {
			values = rrset.Values
			if rrset.Value != nil {
				values = append(values, rrset.Value)
			}
			value = nil
		}
	}

	return json.Marshal(&struct {
		Type   *string      `json:"type,omitempty"`
		TTL    *int         `json:"ttl,omitempty"`
		Value  RRSetValue   `json:"value,omitempty"`
		Values []RRSetValue `json:"values,omitempty"`
	}{
		Type:   rrset.Type,
		TTL:    rrset.TTL,
		Value:  value,
		Values: values,
	})
}
