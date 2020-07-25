package factory

import "encoding/json"

var (
// aliasRegexpFunc func(alias string) bool
)

// GSLBService : struct representation of common_service_request.configuration.gslb_service
type GSLBService struct {
	// Field 0
	Remark *string `draft_validate:"omitempty,max=128" json:"remark,omitempty"`

	// Field 1
	Zone *string `draft_validate:"required" json:"zone,omitempty" conform:"trim,lower,rtrimdot"`

	// Field 2
	Pools *map[string]Pool `draft_validate:"omitempty,dive" activation_validate:"omitempty,dive" json:"pools,omitempty"`

	// Field 3
	LoadBalancedRecords *map[string]LoadBalancedRecord `draft_validate:"omitempty,gte=1,lte=50,dive" activation_validate:"omitempty,dive" json:"load_balanced_records,omitempty"`

	// Field 4
	VirtualServers *map[string]VirtualServer `draft_validate:"omitempty,dive" activation_validate:"omitempty,dive" json:"virtual_servers,omitempty"`

	// Field 5
	Members *[]Member `draft_validate:"omitempty,dive" activation_validate:"omitempty,dive" json:"members,omitempty"`

	// Field 6
	Regions *map[string]Region `draft_validate:"omitempty,dive" activation_validate:"omitempty,dive" json:"regions,omitempty"`

	// Field 7
	Monitors *map[string]Monitor `draft_validate:"omitempty,dive" activation_validate:"omitempty,dive" json:"monitors,omitempty"`
}

// GetMonitors : Return the GSLBService.Monitors map. If nil, return an empty map.
func (g *GSLBService) GetMonitors() map[string]Monitor {
	if g.Monitors == nil {
		return map[string]Monitor{}
	}
	return *g.Monitors
}

func (g *GSLBService) GetMonitor(name string) *Monitor {
	if g.Monitors == nil {
		return nil
	}
	if m, ok := (*g.Monitors)[name]; ok {
		return &m
	}
	return nil
}

func (g *GSLBService) SetMonitor(monitorName string, m Monitor) {
	if monitorName == "" {
		return
	}
	if g.Monitors == nil {
		g.Monitors = &map[string]Monitor{}
	}
	(*g.Monitors)[monitorName] = m
}

// GetVirtualServers : Return the GSLBService.VirtualServers map. If nil, return an empty map.
func (g *GSLBService) GetVirtualServers() map[string]VirtualServer {
	if g.VirtualServers == nil {
		return map[string]VirtualServer{}
	}
	return *g.VirtualServers
}

func (g *GSLBService) GetVirtualServer(name string) *VirtualServer {
	if g.VirtualServers == nil {
		return nil
	}
	if vs, ok := (*g.VirtualServers)[name]; ok {
		return &vs
	}
	return nil
}

func (g *GSLBService) SetVirtualServer(vsName string, vs VirtualServer) {
	if vsName == "" {
		return
	}
	if g.VirtualServers == nil {
		g.VirtualServers = &map[string]VirtualServer{}
	}
	(*g.VirtualServers)[vsName] = vs
}

func (g *GSLBService) SetPool(poolName string, pool Pool) {
	if poolName == "" {
		return
	}
	if g.Pools == nil {
		g.Pools = &map[string]Pool{}
	}
	(*g.Pools)[poolName] = pool
}

func (g *GSLBService) GetPool(name string) *Pool {
	if g.Pools == nil {
		return nil
	}
	if vs, ok := (*g.Pools)[name]; ok {
		return &vs
	}
	return nil
}

func (g *GSLBService) GetBigIpLtmVirtualServers() map[string]VirtualServer {
	bigIpLtmVirtualServers := map[string]VirtualServer{}
	for vsName, vs := range g.GetVirtualServers() {
		if vs.IsBigIPLTM() {
			bigIpLtmVirtualServers[vsName] = vs
		}
	}
	return bigIpLtmVirtualServers
}

func (g *GSLBService) GetVipIDs() []string {
	bigIpLtmVirtualServers := g.GetBigIpLtmVirtualServers()
	vidIDs := make([]string, len(bigIpLtmVirtualServers))
	n := 0
	for _, vs := range bigIpLtmVirtualServers {
		if vs.VipID != nil {
			vidIDs[n] = *vs.VipID
		}
	}
	return vidIDs
}

func (g *GSLBService) GetVipIDMap() map[string]bool {
	vipIDMap := map[string]bool{}
	for _, vipID := range g.GetVipIDs() {
		vipIDMap[vipID] = true
	}
	return vipIDMap
}

// GetPools : Return the GSLBService.Pools map. If nil, return an empty map.
func (g *GSLBService) GetPools() map[string]Pool {
	if g.Pools == nil {
		return map[string]Pool{}
	}
	return *g.Pools
}

// GetRegions : Return the GSLBService.Regions map. If nil, return an empty map.
func (g *GSLBService) GetRegions() map[string]Region {
	if g.Regions == nil {
		return map[string]Region{}
	}
	return *g.Regions
}

// GetLoadBalancedRecords : Return the GSLBService.LoadBalancedRecords map. If nil, return an empty map.
func (g *GSLBService) GetLoadBalancedRecords() map[string]LoadBalancedRecord {
	if g.LoadBalancedRecords == nil {
		return map[string]LoadBalancedRecord{}
	}
	return *g.LoadBalancedRecords
}

// GetZone : Return the GSLBService.Zone value. If nil, return an empty string.
func (g *GSLBService) GetZone() string {
	if g.Zone == nil {
		return ""
	}
	return *g.Zone
}

// SetZone : Set the GSLBService.Zone value.
func (g *GSLBService) SetZone(zone string) {
	g.Zone = newStringPointer(zone)
}

// RemoveZone : Set the GSLBService.Zone to nil.
func (g *GSLBService) RemoveZone() {
	g.Zone = nil
}

// AddLoadBalancedRecord : Add a LoadBalancedRecord object to the map of GSLBService.LoadBalancedRecords.
// If GSLBService.LoadBalancedRecords is nil, a new map will be created before inserting the input LBR.
// If an LBR already exists with the input provided name, the existing LBR will be overridden with the input LBR.
func (g *GSLBService) AddLoadBalancedRecord(name string, lbr LoadBalancedRecord) {
	if g.LoadBalancedRecords == nil {
		g.LoadBalancedRecords = &map[string]LoadBalancedRecord{}
	}
	(*g.LoadBalancedRecords)[name] = lbr
}

// RemoveLoadBalancedRecord : Remove the map entry for the provided key 'name'.
// No effect if 'name' does not exist as a key.
func (g *GSLBService) RemoveLoadBalancedRecord(name string) {
	if g.LoadBalancedRecords == nil {
		return
	}
	delete((*g.LoadBalancedRecords), name)
}

// EmptyLoadBalancedRecords : Set the GSLBService.LoadBalancedRecords to an empty map[string]LoadBalancedRecords{}.
func (g *GSLBService) EmptyLoadBalancedRecords() {
	g.LoadBalancedRecords = &map[string]LoadBalancedRecord{}
}

// RemoveLoadBalancedRecords : Set GSLBService.LoadBalancedRecords to nil.
func (g *GSLBService) RemoveLoadBalancedRecords() {
	g.LoadBalancedRecords = nil
}

// AddPool : Add a Pool object to the map of GSLBService.Pools.
// If GSLBService.Pools is nil, a new map will be created before inserting the input Pool.
// If a Pool already exists with the input provided name, the existing Pool will be overridden with the input Pool.
func (g *GSLBService) AddPool(name string, pool Pool) {
	if g.Pools == nil {
		g.Pools = &map[string]Pool{}
	}
	(*g.Pools)[name] = pool
}

// RemovePool : Remove the map entry for the provided key 'name'.
// No effect if 'name' does not exist as a key.
func (g *GSLBService) RemovePool(name string) {
	if g.Pools == nil {
		return
	}
	delete(*g.Pools, name)
}

// EmptyPools : Set the GSLBService.Pools to an empty map[string]Pool{}.
func (g *GSLBService) EmptyPools() {
	g.Pools = &map[string]Pool{}
}

// RemovePools : Set GSLBService.Pools to nil.
func (g *GSLBService) RemovePools() {
	g.Pools = nil
}

// AddRegion : Add a Region object to the map of GSLBService.Regions.
// If GSLBService.Regions is nil, a new map will be created before inserting the input Region.
// If a Region already exists with the input provided name, the existing Region will be overridden with the input Region.
func (g *GSLBService) AddRegion(name string, region Region) {
	if g.Regions == nil {
		g.Regions = &map[string]Region{}
	}
	(*g.Regions)[name] = region
}

// RemoveRegion : Remove the map entry for the provided key 'name'.
// No effect if 'name' does not exist as a key.
func (g *GSLBService) RemoveRegion(name string) {
	if g.Regions == nil {
		return
	}
	delete(*g.Regions, name)
}

// EmptyRegions : Set the GSLBService.Regions to an empty map[string]Region{}.
func (g *GSLBService) EmptyRegions() {
	g.Regions = &map[string]Region{}
}

// RemoveRegions : Set GSLBService.Regions to nil.
func (g *GSLBService) RemoveRegions() {
	g.Regions = nil
}

// AddMonitor : Add a Monitor object to the map of GSLBService.Monitor.
// If GSLBService.Monitors is nil, a new map will be created before inserting the input Monitor.
// If a Monitor already exists with the input provided name, the existing Monitor will be overridden with the input Monitor.
func (g *GSLBService) AddMonitor(name string, monitor Monitor) {
	if g.Monitors == nil {
		g.Monitors = &map[string]Monitor{}
	}
	(*g.Monitors)[name] = monitor
}

// RemoveMonitor : Remove the map entry for the provided key 'name'.
// No effect if 'name' does not exist as a key.
func (g *GSLBService) RemoveMonitor(name string) {
	if g.Monitors == nil {
		return
	}
	delete(*g.Monitors, name)
}

// EmptyMonitors : Set the GSLBService.Monitors to an empty map[string]Monitor{}.
func (g *GSLBService) EmptyMonitors() {
	g.Monitors = &map[string]Monitor{}
}

// RemoveMonitors : Set GSLBService.Monitors to nil.
func (g *GSLBService) RemoveMonitors() {
	g.Monitors = nil
}

// AddVirtualServer : Add a VirtualServer object to the map of GSLBService.VirtualServer.
// If GSLBService.VirtualServers is nil, a new map will be created before inserting the input VirtualServer.
// If a VirtualServer already exists with the input provided name, the existing VirtualServer will be overridden with the input VirtualServer.
func (g *GSLBService) AddVirtualServer(name string, virtualServer VirtualServer) {
	if g.VirtualServers == nil {
		g.VirtualServers = &map[string]VirtualServer{}
	}
	(*g.VirtualServers)[name] = virtualServer
}

// RemoveVirtualServer : Remove the map entry for the provided key 'name'.
// No effect if 'name' does not exist as a key.
func (g *GSLBService) RemoveVirtualServer(name string) {
	if g.VirtualServers == nil {
		return
	}
	delete(*g.VirtualServers, name)
}

// EmptyVirtualServers : Set the GSLBService.VirtualServers to an empty map[string]VirtualServer{}.
func (g *GSLBService) EmptyVirtualServers() {
	g.VirtualServers = &map[string]VirtualServer{}
}

// RemoveVirtualServers : Set GSLBService.VirtualServers to nil.
func (g *GSLBService) RemoveVirtualServers() {
	g.VirtualServers = nil
}

// Trim : some fields require conditional trimming of fields.
func (g *GSLBService) Trim() {
	// Trim virtual_servers based on virtual_server_type
	if g.VirtualServers != nil {
		for vsName, vs := range *g.VirtualServers {
			vs.Trim()
			(*g.VirtualServers)[vsName] = vs
		}
	}

	// Trim Members
	if g.Pools != nil {
		for poolName, pool := range *g.Pools {
			pool.TrimMembers()
			(*g.Pools)[poolName] = pool
		}
	}

	// remove monitors for pool members which reference bigip-ltm virtual-servers
	bigipLtmVs := g.GetBigIpLtmVirtualServers()
	for poolName, pool := range g.GetPools() {
		updatedPool := false
		for n := range pool.Members {
			if _, ok := bigipLtmVs[ssp(pool.Members[n].VirtualServer)]; ok && pool.Members[n].Monitor != nil {
				pool.Members[n].Monitor = nil
				updatedPool = true
			}
		}
		if updatedPool {
			g.SetPool(poolName, pool)
		}
	}
}

// MarshalToString Helps debug process by transforming VS object into string
func (g *GSLBService) MarshalToString() string {
	b, err := json.MarshalIndent(g, " ", " ")
	if err != nil {
		return ""
	}
	return string(b)

}

func LBRFromJSON(lbr string, alias []string) *LoadBalancedRecord {
	var lbrObj LoadBalancedRecord
	err := json.Unmarshal([]byte(lbr), &lbrObj)
	if err == nil {
		lbrObj.Aliases = alias
		return &lbrObj
	}
	return nil
}
