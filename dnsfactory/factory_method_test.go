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

	f = NewFactory(Opts{JSONServiceConfig: configurationDnsServiceJSON})
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
var configurationDnsServiceJSON = `{
	"dns_service": { 
		"zone": "default.useastexample.com",
		"ttl": 3000,
		"refresh": 6666,
		"retry": 7777,
		"expire": 88888,
		"negative_ttl": 3333,
		"primary_nameserver": "ns1.f5cloudservices.com",
		"records":{
			"":[
				{
					"ttl": 1123,
					"type": "NS",
					"values":["ns1","ns2"]
				}
			]
		}
	}
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

func TestDnsAddARecord(t *testing.T) {
	recordName := "www"
	d := &DNSService{}
	d.AddARecord(recordName, "1.2.2.2", "2.3.4.5")
	assert.Len(t, d.GetARecordSet(recordName).Values, 2)
}

func TestDnsAddAAAARecord(t *testing.T) {
	recordName := "www"
	d := &DNSService{}
	d.AddAAAARecord(recordName, "1234::")
	assert.Len(t, d.GetAAAARecordSet(recordName).Values, 1)

	d.AddAAAARecord(recordName, "2345::", "3456::")
	assert.Len(t, d.GetAAAARecordSet(recordName).Values, 3)
}

func TestDnsAddRecord(t *testing.T) {
	recordName := "www"
	d := &DNSService{}
	d.AddARecord(recordName, "1.2.2.2", "2.3.4.5").
		AddAAAARecord(recordName, "1234::")

	assert.Len(t, d.GetAAAARecordSet(recordName).Values, 1)
	assert.Len(t, d.GetARecordSet(recordName).Values, 2)
	assert.Nil(t, d.GetCNAMERecordSet(recordName).Value)
	assert.Len(t, d.GetMXRecordSet(recordName).Values, 0)
	assert.Len(t, d.GetTXTRecordSet(recordName).Values, 0)

	d.AddAAAARecord(recordName, "2345::", "3456::")
	assert.Len(t, d.GetAAAARecordSet(recordName).Values, 3)
	assert.Len(t, d.GetARecordSet(recordName).Values, 2)
	assert.Nil(t, d.GetCNAMERecordSet(recordName).Value)
	assert.Len(t, d.GetMXRecordSet(recordName).Values, 0)
	assert.Len(t, d.GetTXTRecordSet(recordName).Values, 0)

	d.AddCNAMERecord(recordName, "my.cname")
	assert.Len(t, d.GetAAAARecordSet(recordName).Values, 3)
	assert.Len(t, d.GetARecordSet(recordName).Values, 2)
	assert.NotNil(t, d.GetCNAMERecordSet(recordName).Value)
	assert.Len(t, d.GetMXRecordSet(recordName).Values, 0)
	assert.Len(t, d.GetTXTRecordSet(recordName).Values, 0)

	d.AddMXRecord(recordName, []MXRRSetValue{
		{Domain: newStringPointer("my.domain"), Priority: newIntPointer(1)},
		{Domain: newStringPointer("my.domain2"), Priority: newIntPointer(1)},
	}...)
	assert.Len(t, d.GetAAAARecordSet(recordName).Values, 3)
	assert.Len(t, d.GetARecordSet(recordName).Values, 2)
	assert.NotNil(t, d.GetCNAMERecordSet(recordName).Value)
	assert.Len(t, d.GetMXRecordSet(recordName).Values, 2)
	assert.Len(t, d.GetTXTRecordSet(recordName).Values, 0)

	d.AddTXTRecord(recordName, "my txt 1", "my txt 2")
	assert.Len(t, d.GetAAAARecordSet(recordName).Values, 3)
	assert.Len(t, d.GetARecordSet(recordName).Values, 2)
	assert.NotNil(t, d.GetCNAMERecordSet(recordName).Value)
	assert.Len(t, d.GetMXRecordSet(recordName).Values, 2)
	assert.Len(t, d.GetTXTRecordSet(recordName).Values, 2)

	println(d.MarshalToString())
}

func TestDnsRemoveRecord(t *testing.T) {
	recordName := "www"
	d := &DNSService{}
	d.AddRecord(recordName, "A", "1.2.2.2", "2.3.4.5").
		AddRecord(recordName, "AAAA", "1234::").
		AddRecord(recordName, "AAAA", "2345::", "3456::")
	assert.Len(t, d.GetRecord(recordName).RRSets, 3)

	// remove all record types
	d.RemoveRecord(recordName)
	assert.Len(t, d.GetRecord(recordName).RRSets, 0)

	// repopulate
	d.AddRecord(recordName, "A", "1.2.2.2", "2.3.4.5").
		AddRecord(recordName, "AAAA", "1234::").
		AddRecord(recordName, "AAAA", "2345::", "3456::")

	// remove A record type
	d.RemoveRecord(recordName, RRTypeA)
	assert.Len(t, d.GetRecord(recordName).RRSets, 2)

	// remove unused record type
	d.RemoveRecord(recordName, RRTypeNS)
	assert.Len(t, d.GetRecord(recordName).RRSets, 2)
}
