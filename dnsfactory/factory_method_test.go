package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFactory(t *testing.T) {
	f := NewFactory(Opts{JSONServiceConfig: serviceRequestJSON})
	// println(f.sr.ToJSON())

	f = NewFactory(Opts{JSONServiceConfig: configurationJSON})
	// println(f.sr.ToJSON())
	assert.NotEmpty(t, f.sr.SubscriptionId)
	assert.NotEmpty(t, f.sr.ServiceInstanceId)
	assert.NotEmpty(t, f.sr.ServiceInstanceName)
	assert.NotEmpty(t, f.sr.AccountId)

	f = NewFactory(Opts{JSONServiceConfig: gslbServiceJSON})
	// println(f.sr.ToJSON())
	assert.NotEmpty(t, f.sr.SubscriptionId)
	assert.NotEmpty(t, f.sr.ServiceInstanceId)
	assert.NotEmpty(t, f.sr.ServiceInstanceName)
	assert.NotEmpty(t, f.sr.AccountId)

	f = NewFactory(Opts{JSONServiceConfig: partialGslbServiceJSON})
	// println(f.sr.ToJSON())
	assert.NotEmpty(t, f.sr.SubscriptionId)
	assert.NotEmpty(t, f.sr.ServiceInstanceId)
	assert.NotEmpty(t, f.sr.ServiceInstanceName)
	assert.NotEmpty(t, f.sr.AccountId)
}

func TestAddMethods(t *testing.T) {
	f := NewFactory()
	gslbObj := f.GetGslbService()

	lbrLen := len(f.GetGslbService().GetLoadBalancedRecords())
	lbrName := "mylbr"
	gslbObj.AddLoadBalancedRecord(lbrName, LoadBalancedRecord{})
	assert.Len(t, gslbObj.GetLoadBalancedRecords(), lbrLen+1)
	gslbObj.RemoveLoadBalancedRecord(lbrName)
	assert.Len(t, gslbObj.GetLoadBalancedRecords(), lbrLen)
	gslbObj.EmptyLoadBalancedRecords()
	assert.Len(t, gslbObj.GetLoadBalancedRecords(), 0)
	gslbObj.RemoveLoadBalancedRecords()
	assert.Nil(t, gslbObj.LoadBalancedRecords)

	poolLen := len(gslbObj.GetPools())
	poolName := "mypool"
	gslbObj.AddPool(poolName, Pool{})
	assert.Len(t, gslbObj.GetPools(), poolLen+1)
	gslbObj.RemovePool(poolName)
	assert.Len(t, gslbObj.GetPools(), poolLen)
	gslbObj.EmptyPools()
	assert.Len(t, gslbObj.GetPools(), 0)
	gslbObj.RemovePools()
	assert.Nil(t, gslbObj.Pools)

	monitorLen := len(gslbObj.GetMonitors())
	monitorName := "mymonitor"
	gslbObj.AddMonitor(monitorName, Monitor{})
	assert.Len(t, gslbObj.GetMonitors(), monitorLen+1)
	gslbObj.RemoveMonitor(monitorName)
	assert.Len(t, gslbObj.GetMonitors(), monitorLen)
	gslbObj.EmptyMonitors()
	assert.Len(t, gslbObj.GetMonitors(), 0)
	gslbObj.RemoveMonitors()
	assert.Nil(t, gslbObj.Monitors)

	regionLen := len(gslbObj.GetRegions())
	regionName := "myregion"
	gslbObj.AddRegion(regionName, Region{})
	assert.Len(t, gslbObj.GetRegions(), regionLen+1)
	gslbObj.RemoveRegion(regionName)
	assert.Len(t, gslbObj.GetRegions(), regionLen)
	gslbObj.EmptyRegions()
	assert.Len(t, gslbObj.GetRegions(), 0)
	gslbObj.RemoveRegions()
	assert.Nil(t, gslbObj.Regions)

	virtualServerLen := len(gslbObj.GetVirtualServers())
	virtualServerName := "myvirtualServer"
	gslbObj.AddVirtualServer(virtualServerName, VirtualServer{})
	assert.Len(t, gslbObj.GetVirtualServers(), virtualServerLen+1)
	gslbObj.RemoveVirtualServer(virtualServerName)
	assert.Len(t, gslbObj.GetVirtualServers(), virtualServerLen)
	gslbObj.EmptyVirtualServers()
	assert.Len(t, gslbObj.GetVirtualServers(), 0)
	gslbObj.RemoveVirtualServers()
	assert.Nil(t, gslbObj.VirtualServers)

}

var serviceRequestJSON = `{
	"subscription_id": "s-aaVORdXTHP",
	"account_id": "a-aa4DHvEfZ_",
	"user_id": "u-aaHCrtHVT6",
	"catalog_id": "c-aaQnOrPjGu",
	"service_instance_id": "gslb-aa9qy8-6R8",
	"status": "DISABLED",
	"service_instance_name": "gslb-go-client-instance-1572042084443805000",
	"deleted": false,
	"service_type": "gslb",
	"configuration": {
		"gslb_service": {
			"load_balanced_records": {
				"lbr1": {
					"aliases": ["www"],
					"display_name": "lbr1",
					"enable": true,
					"persist_cidr_ipv4": 24,
					"persist_cidr_ipv6": 64,
					"persistence": false,
					"persistence_ttl": 3600,
					"proximity_rules": [{
						"pool": "pool1",
						"region": "us",
						"score": 30
					}],
					"remark": "lbr1",
					"rr_type": "A"
				}
			},
			"monitors": {
				"monitor1": {
					"display_name": "monitor1",
					"monitor_type": "http_advanced",
					"receive": "HTTP/1.",
					"remark": "Monitor1",
					"send": "HEAD / HTTP/1.0\\r\\n\\r\\n",
					"target_port": 80
				}
			},
			"pools": {
				"pool1": {
					"display_name": "Pool1",
					"enable": true,
					"load_balancing_mode": "ratio-member",
					"max_answers": 1,
					"members": [{
						"final": null,
						"monitor": "basic",
						"ratio": 80,
						"virtual_server": "vs1"
					}],
					"remark": "Pool1_sample",
					"rr_type": "A",
					"ttl": 600
				}
			},
			"regions": {
				"us": {
					"display_name": "us",
					"remark": "us",
					"sectors": [{
						"code": "US",
						"scale": "country"
					}]
				}
			},
			"remark": "sample",
			"virtual_servers": {
				"vs1": {
					"address": "13.56.180.35",
					"display_name": "vs1",
					"monitor": "monitor1",
					"port": 80,
					"remark": "sample vs"
				}
			},
			"zone": "goclient.gslb1572042084443333000.example.com"
		},
		"nameservers": [{
			"ipv4": "107.162.158.143",
			"ipv6": "2604:e180:1021::ffff:6ba2:9e8f",
			"name": "ns1.dns.dev.f5aas.com"
		}, {
			"ipv4": "107.162.158.144",
			"ipv6": "2604:e180:1021::ffff:6ba2:9e90",
			"name": "ns2.dns.dev.f5aas.com"
		}]
	},
	"create_time": "2019-10-25T22:21:24.792305Z",
	"update_time": "2019-10-25T22:21:24.792305Z",
	"cancel_time": null,
	"end_time": null
}`
var configurationJSON = `{
	"gslb_service": {
		"load_balanced_records": {
			"lbr1": {
				"aliases": ["www"],
				"display_name": "lbr1",
				"enable": true,
				"persist_cidr_ipv4": 24,
				"persist_cidr_ipv6": 64,
				"persistence": false,
				"persistence_ttl": 3600,
				"proximity_rules": [{
					"pool": "pool1",
					"region": "us",
					"score": 30
				}],
				"remark": "lbr1",
				"rr_type": "A"
			}
		},
		"monitors": {
			"monitor1": {
				"display_name": "monitor1",
				"monitor_type": "http_advanced",
				"receive": "HTTP/1.",
				"remark": "Monitor1",
				"send": "HEAD / HTTP/1.0\\r\\n\\r\\n",
				"target_port": 80
			}
		},
		"pools": {
			"pool1": {
				"display_name": "Pool1",
				"enable": true,
				"load_balancing_mode": "ratio-member",
				"max_answers": 1,
				"members": [{
					"final": null,
					"monitor": "basic",
					"ratio": 80,
					"virtual_server": "vs1"
				}],
				"remark": "Pool1_sample",
				"rr_type": "A",
				"ttl": 600
			}
		},
		"regions": {
			"us": {
				"display_name": "us",
				"remark": "us",
				"sectors": [{
					"code": "US",
					"scale": "country"
				}]
			}
		},
		"remark": "sample",
		"virtual_servers": {
			"vs1": {
				"address": "13.56.180.35",
				"display_name": "vs1",
				"monitor": "monitor1",
				"port": 80,
				"remark": "sample vs"
			}
		},
		"zone": "goclient.gslb1572042084443333000.example.com"
	},
	"nameservers": [{
		"ipv4": "107.162.158.143",
		"ipv6": "2604:e180:1021::ffff:6ba2:9e8f",
		"name": "ns1.dns.dev.f5aas.com"
	}, {
		"ipv4": "107.162.158.144",
		"ipv6": "2604:e180:1021::ffff:6ba2:9e90",
		"name": "ns2.dns.dev.f5aas.com"
	}]
}`
var gslbServiceJSON = `{
	"load_balanced_records": {
		"lbr1": {
			"aliases": ["www"],
			"display_name": "lbr1",
			"enable": true,
			"persist_cidr_ipv4": 24,
			"persist_cidr_ipv6": 64,
			"persistence": false,
			"persistence_ttl": 3600,
			"proximity_rules": [{
				"pool": "pool1",
				"region": "us",
				"score": 30
			}],
			"remark": "lbr1",
			"rr_type": "A"
		}
	},
	"monitors": {
		"monitor1": {
			"display_name": "monitor1",
			"monitor_type": "http_advanced",
			"receive": "HTTP/1.",
			"remark": "Monitor1",
			"send": "HEAD / HTTP/1.0\\r\\n\\r\\n",
			"target_port": 80
		}
	},
	"pools": {
		"pool1": {
			"display_name": "Pool1",
			"enable": true,
			"load_balancing_mode": "ratio-member",
			"max_answers": 1,
			"members": [{
				"final": null,
				"monitor": "basic",
				"ratio": 80,
				"virtual_server": "vs1"
			}],
			"remark": "Pool1_sample",
			"rr_type": "A",
			"ttl": 600
		}
	},
	"regions": {
		"us": {
			"display_name": "us",
			"remark": "us",
			"sectors": [{
				"code": "US",
				"scale": "country"
			}]
		}
	},
	"remark": "sample",
	"virtual_servers": {
		"vs1": {
			"address": "13.56.180.35",
			"display_name": "vs1",
			"monitor": "monitor1",
			"port": 80,
			"remark": "sample vs"
		}
	},
	"zone": "goclient.gslb1572042084443333000.example.com"
}`

var partialGslbServiceJSON = `{
	"pools": {
		"pool1": {
			"display_name": "Pool1",
			"enable": true,
			"load_balancing_mode": "ratio-member",
			"max_answers": 1,
			"members": [{
				"final": null,
				"monitor": "basic",
				"ratio": 80,
				"virtual_server": "vs1"
			}],
			"remark": "Pool1_sample",
			"rr_type": "A",
			"ttl": 600
		}
	},
	"zone": "goclient.gslb1572042084443333000.example.com"
}`
