package factory

import (
	"encoding/json"
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

func (d *DNSService) MarshalJSON() ([]byte, error) {
	type AliasGslb DNSService
	if d.Records != nil {
		tempSlice := make(map[string][]RRSet)
		for n, rrWrap := range *d.Records {
			tempSlice[n] = rrWrap.RRSets
		}

		return json.Marshal(&struct {
			*AliasGslb
			Records *map[string][]RRSet `json:"records,omitempty"`
		}{
			AliasGslb: (*AliasGslb)(d),
			Records:   &tempSlice,
		})
	}

	return json.Marshal(&struct {
		*AliasGslb
	}{
		AliasGslb: (*AliasGslb)(d),
	})

}

func (g *DNSService) MarshalToString() string {
	b, err := json.MarshalIndent(g, " ", " ")
	if err != nil {
		return ""
	}
	return string(b)
}
