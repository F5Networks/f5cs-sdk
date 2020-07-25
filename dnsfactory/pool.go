package factory

import (
	"encoding/json"
	"strings"
)

const (
	// minMaxAnswers = 1
	// maxMaxAnswers = 32
	// minPriority   = 0
	// maxPriority   = 255

	RRTypeA     = "A"
	RRTypeAAAA  = "AAAA"
	RRTypeMX    = "MX"
	RRTypeCNAME = "CNAME"

	LoadBalancingModeRatio       = "ratio-member"
	LoadBalancingModeRoundRobin  = "round-robin"
	LoadBalancingModeStatic      = "static-persist"
	LoadBalancingModeFewestConn  = "fewest-connections"
	LoadBalancingModeVIPCapacity = "vip-capacity"
	LoadBalancingModeLeastPacket = "least-packet-rate"
	LoadBalancingModePriority    = "priority"
)

type Member struct {
	// Field 0
	Domain *string `draft_validate:"omitempty,hostname" json:"domain,omitempty" conform:"trim,lower,rtrimdot"`

	// Field 1
	Final *bool `json:"final,omitempty"`

	// Field 2
	Ratio *int `draft_validate:"omitempty,min=0,max=100" json:"ratio,omitempty"`

	// Field 3
	Priority *int `draft_validate:"omitempty" json:"priority,omitempty"` //min=0,max=255

	// Field 4
	VirtualServer *string `json:"virtual_server,omitempty"` // regexp ^[A-Za-z][0-9A-Za-z_]{0,47}$

	// Field 5
	Monitor *string `json:"monitor,omitempty"` // regexp ^[A-Za-z][0-9A-Za-z_]{0,47}$
}

type Pool struct {
	// Field 0
	DisplayName *string `draft_validate:"required,max=64" json:"display_name,omitempty"`

	// Field 1
	Remark *string `draft_validate:"omitempty,max=128" json:"remark,omitempty"`

	// Field 2
	Enable *bool `draft_validate:"omitempty" json:"enable,omitempty"`

	// Field 3
	RRType string `draft_validate:"required,oneof=AAAA A MX CNAME" json:"rr_type,omitempty"`

	// Field 4
	TTL *int `draft_validate:"omitempty,min=0,max=86400" json:"ttl,omitempty"`

	// json:omitempty will drop an empty members array "members":[] when the proto is scrubbed
	// and cause validation to fail.
	// Field 5
	Members []Member `draft_validate:"required,gte=0,lte=50,dive" json:"members"`

	// Field 6
	MaxAnswers *int `draft_validate:"omitempty" json:"max_answers,omitempty"` //min=1,max=32

	// Field 7
	LoadBalancingMode *string `draft_validate:"omitempty,oneof=ratio-member round-robin static-persist fewest-connections vip-capacity least-packet-rate priority" json:"load_balancing_mode,omitempty"`
}

// EnablePool : Set the pool enable attribute to 'true'.
func (p *Pool) EnablePool() {
	p.Enable = newBoolPointer(true)
}

// DisablePool : Set the pool enable attribute to 'false'.
func (p *Pool) DisablePool() {
	p.Enable = newBoolPointer(false)
}

// IsEnabled : Return the value of the Pool.Enable attribute.
func (p *Pool) IsEnabled() bool {
	return p.Enable != nil && *p.Enable
}

// AddMember : add Member to Pool.Member list
func (p *Pool) AddMember(m *Member) {
	if m == nil {
		return
	}
	p.Members = append(p.Members, *m)
}

func (p *Pool) MarshalToString() string {
	b, err := json.MarshalIndent(p, " ", " ")
	if err != nil {
		return ""
	}
	return string(b)

}

func (mmbr *Member) MarshalToString() string {
	b, err := json.MarshalIndent(mmbr, " ", " ")
	if err != nil {
		return ""
	}
	return string(b)

}

func getRatioOrDefault(r *int) *int {
	if r != nil {
		return r
	} else {
		return newIntPointer(1)
	}
}

// UnmarshalJSON : custom logic to unmarshal a pool, including setting its members
// this is called by the json.Unmarshal function internally, and overrides the default behavior
// it can be expected to return a valid Pool struct based on the JSON provided,
// with defaults added where appropriate
func (p *Pool) UnmarshalJSON(data []byte) error {
	// get a map of the fields in the incoming data
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(data, &objMap)

	// map them all out to their corresponding struct keys
	if objMap["display_name"] != nil {
		err = json.Unmarshal(*objMap["display_name"], &p.DisplayName)
		if err != nil {
			return err
		}
	}

	if objMap["remark"] != nil {
		err = json.Unmarshal(*objMap["remark"], &p.Remark)
		if err != nil {
			return err
		}
	}

	if objMap["enable"] != nil {
		err = json.Unmarshal(*objMap["enable"], &p.Enable)
		if err != nil {
			return err
		}
	} else {
		p.Enable = newBoolPointer(true)
	}

	if objMap["rr_type"] != nil {
		err = json.Unmarshal(*objMap["rr_type"], &p.RRType)
		if err != nil {
			return err
		}
	}

	if objMap["ttl"] != nil {
		err = json.Unmarshal(*objMap["ttl"], &p.TTL)
		if err != nil {
			return err
		}
	} else {
		p.TTL = newIntPointer(30)
	}

	if objMap["max_answers"] != nil {
		err = json.Unmarshal(*objMap["max_answers"], &p.MaxAnswers)
		if err != nil {
			return err
		}
	} else if p.RRType != RRTypeCNAME {
		p.MaxAnswers = newIntPointer(1)
	}

	if objMap["load_balancing_mode"] != nil {
		err = json.Unmarshal(*objMap["load_balancing_mode"], &p.LoadBalancingMode)
		if err != nil {
			return err
		}
	} else {
		p.LoadBalancingMode = newStringPointer("round-robin")
	}

	if objMap["members"] != nil {
		var members []Member
		err = json.Unmarshal(*objMap["members"], &members)
		if err != nil {
			return err
		}

		// fill the struct's members with a slice of the correct length
		p.Members = make([]Member, len(members))

		usingRatioBalancing := strings.ToLower(*p.LoadBalancingMode) == strings.ToLower(LoadBalancingModeRatio)
		usingPriorityBalancing := strings.ToLower(*p.LoadBalancingMode) == strings.ToLower(LoadBalancingModePriority)

		// get the correct fields according to the RR type we set earlier
		switch strings.ToUpper(p.RRType) {
		case RRTypeA, RRTypeAAAA:
			for i := 0; i < len(members); i++ {
				p.Members[i] = Member{
					VirtualServer: members[i].VirtualServer,
					Monitor:       members[i].Monitor,
				}

				if usingRatioBalancing {
					p.Members[i].Ratio = getRatioOrDefault(members[i].Ratio)
				}
				if usingPriorityBalancing {
					p.Members[i].Priority = members[i].Priority
				}
			}
		case RRTypeCNAME:
			for i := 0; i < len(members); i++ {
				p.Members[i] = Member{
					Domain: members[i].Domain,
					Final:  members[i].Final,
				}
				if usingRatioBalancing {
					p.Members[i].Ratio = getRatioOrDefault(members[i].Ratio)
				}
			}
		case RRTypeMX:
			for i := 0; i < len(members); i++ {
				p.Members[i] = Member{
					Domain:   members[i].Domain,
					Priority: members[i].Priority,
				}
				if usingRatioBalancing {
					p.Members[i].Ratio = getRatioOrDefault(members[i].Ratio)
				}
			}
		default:
			// if we can't find the RR type just unmarshal all of the possible member data, since we don't know what we need
			err = json.Unmarshal(*objMap["members"], &p.Members)
			if err != nil {
				return err
			}
		}
	}

	return err
}

// UnmarshalString : helper function to unmarshal a string and
// mask unnecessary member attributes if provided.
func (p *Pool) UnmarshalString(jstr string) {
	err := json.Unmarshal([]byte(jstr), p)
	if err != nil {
		p = &Pool{}
		return
	}
	p.TrimMembers()

}

func (p *Pool) TrimMembers() {
	switch p.RRType {
	case RRTypeA, RRTypeAAAA:
		p.trimAddrMember()
	case RRTypeMX:
		p.trimMXMember()
	case RRTypeCNAME:
		p.trimCNAMEMember()
	}
}

func (p *Pool) trimAddrMember() {
	for n := range p.Members {
		p.Members[n].Domain = nil
		if *p.LoadBalancingMode != LoadBalancingModePriority {
			p.Members[n].Priority = nil
		}
		p.Members[n].Final = nil
	}
}

func (p *Pool) trimMXMember() {
	for n := range p.Members {
		p.Members[n].Final = nil
		p.Members[n].VirtualServer = nil
		p.Members[n].Monitor = nil
	}
}

func (p *Pool) trimCNAMEMember() {
	p.MaxAnswers = nil
	for n := range p.Members {
		// p.Members[n].MaxAnswers = nil
		p.Members[n].Priority = nil
		p.Members[n].VirtualServer = nil
		p.Members[n].Monitor = nil
	}
}

func (p *Pool) IsPoolTypeAOrAAAA() bool {
	if p.RRType == RRTypeA || p.RRType == RRTypeAAAA {
		return true
	}
	return false
}

func (m1 *Member) IsEqual(m2 *Member) bool {
	if m1 == nil || m2 == nil {
		return false
	}
	if !m1.IsSameMemberType(m2) {
		return false
	}
	if m1.IsMXMember() || m1.IsCNAMEMember() {
		// mx, cname -> domain defines uniqueness
		if *m1.Domain == *m2.Domain {
			return true
		}
	} else if m1.IsAddressMember() {
		// A, AAAA -> virtual_server defines uniqueness
		if *m1.VirtualServer == *m2.VirtualServer {
			return true
		}
	}
	return false
}

func (m *Member) IsAddressMember() bool {
	if m.VirtualServer != nil && m.Domain == nil {
		return true
	}
	return false
}

func (m *Member) IsMXMember() bool {
	if m.Domain != nil && m.Final == nil && m.Monitor == nil && m.VirtualServer == nil {
		return true
	}
	return false
}
func (m *Member) IsCNAMEMember() bool {
	if m.Domain != nil && m.Priority == nil && m.Monitor == nil && m.VirtualServer == nil {
		return true
	}
	return false
}

func (m1 *Member) IsSameMemberType(m2 *Member) bool {
	if m1 == nil || m2 == nil {
		return false
	}
	if m1.IsAddressMember() && m2.IsAddressMember() {
		return true
	} else if m1.IsMXMember() && m2.IsMXMember() {
		return true
	} else if m1.IsCNAMEMember() && m2.IsCNAMEMember() {
		return true
	}
	return false
}
