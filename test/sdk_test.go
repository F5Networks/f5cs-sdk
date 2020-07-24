package test

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	authentication "gitswarm.f5net.com/f5aas/f5cs-sdk/generated/authentication"
	subscription "gitswarm.f5net.com/f5aas/f5cs-sdk/generated/subscription"
	"testing"
)

func TestClient(t *testing.T) {
	authCfg := authentication.NewConfiguration()
	authClient := authentication.NewAPIClient(authCfg)
	login, _, err := authClient.AuthenticationServiceApi.Login(context.Background(), authentication.AuthenticationServiceLoginRequest{
		Username: "f5aas.qe+exploratory_07@gmail.com",
		Password: "F5aaStest!",
	})
	assert.NoError(t, err)
	cfg := subscription.NewConfiguration()
	c := subscription.NewAPIClient(cfg)
	auth := context.WithValue(context.Background(), subscription.ContextAccessToken, login.AccessToken)
	subs, _, err := c.SubscriptionServiceApi.ListSubscriptions(auth, "a-aanS3R3GQY", &subscription.ListSubscriptionsOpts{})
	assert.NoError(t, err)
	_, err = json.Marshal(subs)
	assert.NoError(t, err)
}
