package factory

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	loglib "log"

	"github.com/icza/dyno"
)

const (
	// initial resource domain - all IDs generated by this version will have this prefix
	resourceIDDomain = "aa"
	// resource ID random component length: 6 bytes (48 bits), 281,474,976,710,655 values
	resourceIDByteLength = 6
	// resource ID template in the following regex format: ([a-z]{1,4})-([A-z0-9]{2})([A-z0-9-_]{8})
	resourceIDTemplate = "%s-%s%s"
)

// NewFactory : Generate a new factory. Calling without Opts will generate default service configuration.
// Use Opts.JSONServiceConfig to initialize factory to a json config at the ServiceRequest, Configuration, or GslbService level. Partial configs supported.
// If there is an error unmarshalling the Opts.JSONServiceConfig, NewFactory will return an empty ServiceRequest struct.
func NewFactory(opts ...Opts) Factory {
	f := Factory{}
	if len(opts) > 0 && opts[0].JSONServiceConfig != "" {
		f.sr = f.serviceRequestFromJSON(opts[0].JSONServiceConfig)
	} else {
		f.sr = f.ServiceRequestDefault()
	}
	return f
}

type Opts struct {
	// Json configuration string.
	JSONServiceConfig string

	// Whether to fill in missing service fields with default values.
	FillWithDefaults bool
}
type Factory struct {
	sr ServiceRequest
}

func (f *Factory) NewIntPointer(i int) *int {
	return newIntPointer(i)
}

func newIntPointer(i int) *int {
	return &i
}

func (f *Factory) NewStringPointer(s string) *string {
	return newStringPointer(s)
}

func newStringPointer(s string) *string {
	return &s
}

func (f *Factory) NewBoolPointer(b bool) *bool {
	return newBoolPointer(b)
}

func newBoolPointer(b bool) *bool {
	return &b
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (f *Factory) RandStringRunes(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (f *Factory) RandIPv4() string {
	rand.Seed(time.Now().UTC().UnixNano())
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

func (f *Factory) PoolDefault() Pool {
	return Pool{
		Enable:            f.NewBoolPointer(true),
		DisplayName:       f.NewStringPointer("this display_name"),
		Remark:            f.NewStringPointer("this remark"),
		RRType:            RRTypeA,
		TTL:               f.NewIntPointer(1),
		LoadBalancingMode: f.NewStringPointer("round-robin"),
		MaxAnswers:        f.NewIntPointer(1),
		Members: []Member{
			f.PoolMemberDefault(),
		},
	}
}

func (f *Factory) PoolMemberDefault() Member {
	return *f.PoolMemberDefaultasPtr()
}

func (f *Factory) PoolMemberDefaultasPtr() *Member {
	return &Member{
		Domain:        f.NewStringPointer("this.domain.com"),
		Final:         f.NewBoolPointer(false),
		Ratio:         f.NewIntPointer(1),
		Priority:      f.NewIntPointer(1),
		VirtualServer: f.NewStringPointer("vs1"),
		Monitor:       f.NewStringPointer("m1"),
	}
}

func (f *Factory) PoolADefault() Pool {
	p := f.PoolDefault()
	p.RRType = RRTypeA
	return p
}

func (f *Factory) PoolAAAADefault() Pool {
	p := f.PoolDefault()
	p.RRType = RRTypeAAAA
	return p
}

func (f *Factory) PoolsDefault() *map[string]Pool {
	return &map[string]Pool{
		"p1": f.PoolDefault(),
	}
}

func (f *Factory) VirtualServerDefault() VirtualServer {
	return VirtualServer{
		DisplayName:        f.NewStringPointer("this display_name"),
		Remark:             f.NewStringPointer("this remark"),
		Port:               f.NewIntPointer(1),
		Monitor:            f.NewStringPointer("m1"),
		Address:            f.NewStringPointer("1.2.3.4"),
		VirtualServerType:  f.NewStringPointer("cloud"),
		VipID:              f.NewStringPointer("abcdefghij"),
		TranslationAddress: f.NewStringPointer("210.012.120.021"),
	}
}

func (f *Factory) VirtualServersDefault() *map[string]VirtualServer {
	return &map[string]VirtualServer{
		"vs1": f.VirtualServerDefault(),
	}
}

func (f *Factory) AddressMemberDefault() Member {
	return Member{
		VirtualServer: f.NewStringPointer("vs1"),
	}
}
func (f *Factory) MxMemberDefault() Member {
	return Member{
		Domain:   f.NewStringPointer("member.domain.com"),
		Priority: f.NewIntPointer(1),
	}
}
func (f *Factory) CnameMemberDefault() Member {
	return Member{
		Domain: f.NewStringPointer("member.domain.com"),
		Final:  f.NewBoolPointer(true),
	}
}

func (f *Factory) LoadBalancedRecordDefault() LoadBalancedRecord {
	return LoadBalancedRecord{
		DisplayName:     f.NewStringPointer("this display_name"),
		Remark:          f.NewStringPointer("this remark"),
		RRType:          f.NewStringPointer("A"),
		Enable:          f.NewBoolPointer(true),
		Aliases:         []string{"alias*com"},
		ProximityRules:  []ProximityRule{f.ProximityRuleDefault()},
		Persistence:     f.NewBoolPointer(true),
		PersistCidrIpv4: f.NewIntPointer(1),
		PersistCidrIpv6: f.NewIntPointer(1),
		PersistTtl:      f.NewIntPointer(1),
	}
}

func (f *Factory) LoadBalancedRecordsDefault() *map[string]LoadBalancedRecord {
	return &map[string]LoadBalancedRecord{
		"lbrName1": f.LoadBalancedRecordDefault(),
	}
}

func (f *Factory) MonitorDefault() Monitor {
	return Monitor{
		DisplayName: f.NewStringPointer("this display_name"),
		Remark:      f.NewStringPointer("this remark"),
		MonitorType: "http_advanced",
		TargetPort:  f.NewIntPointer(1),
		Send:        f.NewStringPointer("this send string"),
		Receive:     f.NewStringPointer("this receive string"),
	}
}

func (f *Factory) MonitorsDefault() *map[string]Monitor {
	return &map[string]Monitor{
		"m1": f.MonitorDefault(),
	}
}

func (f *Factory) SectorDefault() Sector {
	return Sector{
		Code:  "Washington",
		Scale: "state",
	}
}
func (f *Factory) SectorsDefault() []Sector {
	return []Sector{f.SectorDefault()}
}
func (f *Factory) RegionDefault() Region {
	return Region{
		DisplayName: f.NewStringPointer("this display_name"),
		Remark:      f.NewStringPointer("this remark"),
		Sectors:     f.SectorsDefault(),
	}
}

func (f *Factory) RegionsDefault() *map[string]Region {
	return &map[string]Region{
		"regionName1": f.RegionDefault(),
	}
}

func (f *Factory) ProximityRuleDefault() ProximityRule {
	return ProximityRule{
		Region: f.NewStringPointer("global"),
		Pool:   f.NewStringPointer("p1"),
		Score:  f.NewIntPointer(1),
	}
}

func (f *Factory) GslbServiceDefault() *GSLBService {
	return &GSLBService{
		Remark:              f.NewStringPointer("this remark"),
		Zone:                f.NewStringPointer(f.GenerateZoneName("")),
		Pools:               f.PoolsDefault(),
		LoadBalancedRecords: f.LoadBalancedRecordsDefault(),
		VirtualServers:      f.VirtualServersDefault(),
		Monitors:            f.MonitorsDefault(),
		Regions:             f.RegionsDefault(),
	}
}

func (f *Factory) ConfigurationDefault() Configuration {
	return Configuration{
		GslbService: f.GslbServiceDefault(),
	}
}

func (f *Factory) ServiceRequestDefault() ServiceRequest {
	sid, err := NewResourceID("gslb")
	if err != nil {
		loglib.Fatal(`NewResourceID("gslb")`, err)
	}
	subid, err := NewResourceID("s")
	if err != nil {
		loglib.Fatal(`NewResourceID("s")`, err)
	}
	acctid, err := NewResourceID("acct")
	if err != nil {
		loglib.Fatal(`NewResourceID("acct")`, err)
	}
	sname := f.RandStringRunes(15)
	c := f.ConfigurationDefault()
	return ServiceRequest{
		ServiceInstanceId:   sid,
		SubscriptionId:      subid,
		ServiceInstanceName: sname,
		Configuration:       &c,
		AccountId:           acctid,
	}
}

func (f *Factory) ServiceRequestJSON() string {
	sr := f.ServiceRequestDefault()
	b, err := json.MarshalIndent(sr, " ", " ")
	if err != nil {
		return ""
	}
	return string(b)
}

func (f *Factory) GenerateZoneName(prefix string) string {
	sid, err := NewResourceID("gslb")
	if err != nil {
		return ""
	}
	sid = strings.Replace(sid, "_", "", -1)
	sid = strings.Replace(sid, "-", "", -1)
	prefix = strings.Trim(prefix, ". -_")
	if prefix == "" {
		prefix = "gslb"
	}
	return strings.Join([]string{prefix, sid, "com"}, ".")
}

func (f *Factory) serviceRequestFromJSON(jsonStr string) ServiceRequest {

	var v interface{}
	err := json.Unmarshal([]byte(jsonStr), &v)
	if err != nil {
		return ServiceRequest{}
	}
	// if configuration exists -> unmarshal into service request
	// if gslb_service or dns_service exists -> unmarshal into configuration
	// if lbr, pools, monitors, virtual_servers, regions, zone exist -> unmarshal into GSLBService
	_, err = dyno.Get(v, "configuration")
	if err == nil {
		serviceRequest := ServiceRequest{}
		err := json.Unmarshal([]byte(jsonStr), &serviceRequest)
		if err != nil {
			return ServiceRequest{}
		}
		return serviceRequest
	}
	_, gslbErr := dyno.Get(v, "gslb_service")
	_, dnsErr := dyno.Get(v, "dns_service")
	if gslbErr == nil || dnsErr == nil {
		serviceRequest := f.ServiceRequestDefault()
		configuration := Configuration{}
		err := json.Unmarshal([]byte(jsonStr), &configuration)
		if err != nil {
			return ServiceRequest{}
		}
		serviceRequest.Configuration = &configuration
		return serviceRequest
	}
	for _, path := range []string{"load_balanced_records", "pools", "virtual_servers", "regions", "monitors", "zone"} {
		_, err = dyno.Get(v, path)
		if err == nil {
			serviceRequest := f.ServiceRequestDefault()
			configuration := &Configuration{GslbService: &GSLBService{}}
			err := json.Unmarshal([]byte(jsonStr), configuration.GslbService)
			if err != nil {
				return f.ServiceRequestDefault()
			}
			serviceRequest.Configuration = configuration
			return serviceRequest
		}
	}
	return f.ServiceRequestDefault()
}

// ssp : safe string pointer
func ssp(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func NewResourceID(prefix string) (string, error) {
	if len(prefix) == 0 {
		return "", fmt.Errorf("prefix is required")
	}
	if len(prefix) > 4 {
		return "", fmt.Errorf("prefix is too long")
	}

	lowerPrefix := strings.ToLower(prefix)
	s, err := RandomString(resourceIDByteLength)
	if err != nil {
		return "", err
	}

	id := fmt.Sprintf(resourceIDTemplate, lowerPrefix, resourceIDDomain, s)

	return id, nil
}

func RandomString(bytesLen uint32) (string, error) {
	bytes, err := RandomBytes(bytesLen)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(bytes), nil
}

func RandomBytes(length uint32) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
