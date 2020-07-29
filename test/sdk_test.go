package test

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	factory "gitswarm.f5net.com/f5aas/f5cs-sdk/dnsfactory"
	authentication "gitswarm.f5net.com/f5aas/f5cs-sdk/generated/authentication"
	subscription "gitswarm.f5net.com/f5aas/f5cs-sdk/generated/subscription"
	"testing"
)

const (
	USERNAME = "f5aas.qe+exploratory_07@gmail.com"
	PASSWORD = "F5aaStest!"
	ACCOUNT  = "a-aanS3R3GQY"
)

func TestClient(t *testing.T) {
	authCfg := authentication.NewConfiguration()
	authClient := authentication.NewAPIClient(authCfg)
	login, _, err := authClient.AuthenticationServiceApi.Login(context.Background(), authentication.AuthenticationServiceLoginRequest{
		Username: USERNAME,
		Password: PASSWORD,
	})
	assert.NoError(t, err)
	cfg := subscription.NewConfiguration()
	c := subscription.NewAPIClient(cfg)
	auth := context.WithValue(context.Background(), subscription.ContextAccessToken, login.AccessToken)
	subs, _, err := c.SubscriptionServiceApi.ListSubscriptions(auth, ACCOUNT, &subscription.ListSubscriptionsOpts{})
	assert.NoError(t, err)
	_, err = json.Marshal(subs)
	assert.NoError(t, err)
}

func TestWithFactory(t *testing.T) {
	authCfg := authentication.NewConfiguration()
	authClient := authentication.NewAPIClient(authCfg)
	login, _, err := authClient.AuthenticationServiceApi.Login(context.Background(), authentication.AuthenticationServiceLoginRequest{
		Username: USERNAME,
		Password: PASSWORD,
	})
	assert.NoError(t, err)
	cfg := subscription.NewConfiguration()
	c := subscription.NewAPIClient(cfg)
	auth := context.WithValue(context.Background(), subscription.ContextAccessToken, login.AccessToken)
	subs, _, err := c.SubscriptionServiceApi.ListSubscriptions(auth, ACCOUNT, &subscription.ListSubscriptionsOpts{})
	assert.NoError(t, err)
	b, err := json.Marshal(subs)
	assert.NoError(t, err)
	var subscriptions factory.ServiceRequests
	err = json.Unmarshal(b, &subscriptions)
	assert.NoError(t, err)
}

func TestSubscriptionOperations(t *testing.T) {
	// Login
	authCfg := authentication.NewConfiguration()
	authClient := authentication.NewAPIClient(authCfg)
	login, _, err := authClient.AuthenticationServiceApi.Login(context.Background(), authentication.AuthenticationServiceLoginRequest{
		Username: USERNAME,
		Password: PASSWORD,
	})
	assert.NoError(t, err)
	cfg := subscription.NewConfiguration()
	c := subscription.NewAPIClient(cfg)
	auth := context.WithValue(context.Background(), subscription.ContextAccessToken, login.AccessToken)
	// Create Subscription with 1 LBR having 1 pool and 1 member and no monitor.
	createSubRequest, err := BuildSimpleSubscription()
	assert.NoError(t, err)
	sub, resp, err := c.SubscriptionServiceApi.CreateSubscription(auth, *createSubRequest)
	assert.NoError(t, err)
	t.Logf("Error %v, REsp %v", err, resp)
	assert.NotNil(t, sub)
	// Get Subscription s-aaN6xqwSSn. If not create
	gSubId := sub.SubscriptionId
	//gSubId := "s-aaN6xqwSSn"
	sub, _, err = c.SubscriptionServiceApi.GetSubscription(auth, gSubId, nil)
	assert.NoError(t, err)
	b, err := json.Marshal(sub)
	assert.NoError(t, err)
	t.Logf("Initial Subscription %v", string(b))
	// Modify subscription to add LBR
	var sr factory.ServiceRequest
	err = json.Unmarshal(b, &sr)
	assert.NoError(t, err)
	gslbSvc := sr.Configuration.GslbService
	gslbSvc = BuildLBRWithPoolAndVS("pool2", "104.67.19.23", "vs2", "lbr2", gslbSvc)
	svcCfg := factory.Configuration{
		GslbService: gslbSvc,
	}
	var iface map[string]interface{}
	inrec, err := json.Marshal(svcCfg)
	assert.NoError(t, err)
	t.Logf("Updated Subscription %v", string(inrec))
	json.Unmarshal(inrec, &iface)
	updateReq := subscription.V1UpdateSubscriptionRequest{
		SubscriptionId:      gSubId,
		Configuration:       iface,
		ServiceInstanceName: sr.ServiceInstanceName,
	}
	sub, _, err = c.SubscriptionServiceApi.UpdateSubscription(auth, gSubId, updateReq)
	assert.NoError(t, err)
	// Activate Subscription
	_, _, err = c.SubscriptionServiceApi.ActivateSubscription(auth, gSubId)
	assert.NoError(t, err)
	// Delete Lbr2 and Update Subscription
	b, err = json.Marshal(sub)
	assert.NoError(t, err)
	t.Logf("Activated Subscription %v", string(b))
	err = json.Unmarshal(b, &sr)
	assert.NoError(t, err)
	gslbSvc = sr.Configuration.GslbService
	gslbSvc.RemoveLoadBalancedRecord("lbr2")
	svcCfg = factory.Configuration{
		GslbService: gslbSvc,
	}
	inrec, err = json.Marshal(svcCfg)
	assert.NoError(t, err)
	t.Logf("Updated Subscription %v", string(inrec))
	json.Unmarshal(inrec, &iface)
	updateReq = subscription.V1UpdateSubscriptionRequest{
		SubscriptionId:      gSubId,
		Configuration:       iface,
		ServiceInstanceName: sr.ServiceInstanceName,
	}
	sub, _, err = c.SubscriptionServiceApi.UpdateSubscription(auth, gSubId, updateReq)
	assert.NoError(t, err)
	// Retire Subscription
	_, _, err = c.SubscriptionServiceApi.RetireSubscription(auth, gSubId)
	assert.NoError(t, err)
}

func BuildSimpleSubscription() (*subscription.V1CreateSubscriptionRequest, error) {
	cfg, err := BuildGslbServiceConfig()
	if err != nil {
		return nil, err
	}
	createSubRequest := subscription.V1CreateSubscriptionRequest{
		AccountId:           ACCOUNT,
		CatalogId:           "c-aaQnOrPjGu",
		ServiceType:         "gslb",
		ServiceInstanceName: "extdnslbrinst",
		Configuration:       cfg,
	}
	return &createSubRequest, nil
}

func BuildLBRWithPoolAndVS(poolName string, vsAddress string, vsName string, alias string, gslbSvc *factory.GSLBService) *factory.GSLBService {
	f := factory.NewFactory()
	pRule := factory.ProximityRule{
		Region: f.NewStringPointer("global"),
		Score:  f.NewIntPointer(100),
		Pool:   f.NewStringPointer(poolName),
	}
	vs := factory.VirtualServer{
		Monitor:           f.NewStringPointer("none"),
		Address:           f.NewStringPointer(vsAddress),
		DisplayName:       f.NewStringPointer(vsName),
		Port:              f.NewIntPointer(80),
		VirtualServerType: f.NewStringPointer(factory.VIRTUAL_SERVER_TYPE_CLOUD),
	}
	pm := factory.Member{
		VirtualServer: f.NewStringPointer(vsName),
		Final:         f.NewBoolPointer(true),
		Monitor:       f.NewStringPointer("basic"),
	}
	pool := factory.Pool{
		Enable:            f.NewBoolPointer(true),
		DisplayName:       f.NewStringPointer(poolName),
		Remark:            f.NewStringPointer(poolName),
		RRType:            factory.RRTypeA,
		TTL:               f.NewIntPointer(30),
		LoadBalancingMode: f.NewStringPointer("round-robin"),
		MaxAnswers:        f.NewIntPointer(1),
		Members: []factory.Member{
			pm,
		},
	}
	lbr1 := factory.LoadBalancedRecord{
		DisplayName:    f.NewStringPointer(alias),
		Remark:         f.NewStringPointer(alias),
		RRType:         f.NewStringPointer("A"),
		Enable:         f.NewBoolPointer(true),
		Aliases:        []string{alias},
		ProximityRules: []factory.ProximityRule{pRule},
		Persistence:    f.NewBoolPointer(false),
	}
	gslbSvc.AddLoadBalancedRecord(alias, lbr1)
	gslbSvc.AddVirtualServer(vsName, vs)
	gslbSvc.AddPool(poolName, pool)
	return gslbSvc
}

func BuildGslbServiceConfig() (map[string]interface{}, error) {
	gslbSvc := &factory.GSLBService{}
	gslbSvc.SetZone("extdnsint.net")
	gslbSvc = BuildLBRWithPoolAndVS("extdns1pool", "145.12.1.1", "extdnsvs1", "lbr1", gslbSvc)
	cfg := factory.Configuration{
		GslbService: gslbSvc,
	}
	var iface map[string]interface{}
	inrec, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(inrec, &iface)
	return iface, nil
}
