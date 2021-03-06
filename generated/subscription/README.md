# Go API client for subscription

Manages subscriptions tied to a specific service in F5 Cloud Services.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./subscription"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.cloudservices.f5.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*SubscriptionServiceApi* | [**ActivateSubscription**](docs/SubscriptionServiceApi.md#activatesubscription) | **Post** /v1/svc-subscription/subscriptions/{subscription_id}/activate | Activate a subscription
*SubscriptionServiceApi* | [**BatchActivateSubscription**](docs/SubscriptionServiceApi.md#batchactivatesubscription) | **Post** /v1/svc-subscription/subscriptions/batch-activate | Activate multiple subscriptions
*SubscriptionServiceApi* | [**BatchRetireSubscription**](docs/SubscriptionServiceApi.md#batchretiresubscription) | **Post** /v1/svc-subscription/subscriptions/batch-retire | Retire multiple subscriptions
*SubscriptionServiceApi* | [**BatchSubscriptionStatus**](docs/SubscriptionServiceApi.md#batchsubscriptionstatus) | **Post** /v1/svc-subscription/subscriptions/batch-status | Get the status of multiple subscriptions
*SubscriptionServiceApi* | [**BatchSuspendSubscription**](docs/SubscriptionServiceApi.md#batchsuspendsubscription) | **Post** /v1/svc-subscription/subscriptions/batch-suspend | Suspend multiple subscriptions
*SubscriptionServiceApi* | [**BatchUnretireSubscription**](docs/SubscriptionServiceApi.md#batchunretiresubscription) | **Post** /v1/svc-subscription/subscriptions/batch-unretire | Un-retire multiple subscriptions
*SubscriptionServiceApi* | [**CreateSubscription**](docs/SubscriptionServiceApi.md#createsubscription) | **Post** /v1/svc-subscription/subscriptions | Create a subscription
*SubscriptionServiceApi* | [**GetSubscription**](docs/SubscriptionServiceApi.md#getsubscription) | **Get** /v1/svc-subscription/subscriptions/{subscription_id} | Get a subscription
*SubscriptionServiceApi* | [**ListSubscriptions**](docs/SubscriptionServiceApi.md#listsubscriptions) | **Get** /v1/svc-subscription/subscriptions | List subscriptions in an account
*SubscriptionServiceApi* | [**RetireSubscription**](docs/SubscriptionServiceApi.md#retiresubscription) | **Post** /v1/svc-subscription/subscriptions/{subscription_id}/retire | Retire a subscription
*SubscriptionServiceApi* | [**SubscriptionStatus**](docs/SubscriptionServiceApi.md#subscriptionstatus) | **Get** /v1/svc-subscription/subscriptions/{subscription_id}/status | Get a subscription&#39;s status
*SubscriptionServiceApi* | [**SuspendSubscription**](docs/SubscriptionServiceApi.md#suspendsubscription) | **Post** /v1/svc-subscription/subscriptions/{subscription_id}/suspend | Suspend a subscription
*SubscriptionServiceApi* | [**TestSubscription**](docs/SubscriptionServiceApi.md#testsubscription) | **Post** /v1/svc-subscription/subscriptions/{subscription_id}/test | Test a subscription
*SubscriptionServiceApi* | [**UnretireSubscription**](docs/SubscriptionServiceApi.md#unretiresubscription) | **Post** /v1/svc-subscription/subscriptions/{subscription_id}/unretire | Un-retire a subscription
*SubscriptionServiceApi* | [**UpdateSubscription**](docs/SubscriptionServiceApi.md#updatesubscription) | **Put** /v1/svc-subscription/subscriptions/{subscription_id} | Update a subscription


## Documentation For Models

 - [ProtobufNullValue](docs/ProtobufNullValue.md)
 - [V1BatchSubscriptionIdRequest](docs/V1BatchSubscriptionIdRequest.md)
 - [V1BatchSubscriptionStatusResponse](docs/V1BatchSubscriptionStatusResponse.md)
 - [V1CheckSubscriptionAccessResponse](docs/V1CheckSubscriptionAccessResponse.md)
 - [V1CreateSubscriptionRequest](docs/V1CreateSubscriptionRequest.md)
 - [V1ServiceState](docs/V1ServiceState.md)
 - [V1Status](docs/V1Status.md)
 - [V1Subscription](docs/V1Subscription.md)
 - [V1SubscriptionStatusResponse](docs/V1SubscriptionStatusResponse.md)
 - [V1Subscriptions](docs/V1Subscriptions.md)
 - [V1TestSubscriptionResponse](docs/V1TestSubscriptionResponse.md)
 - [V1UpdateSubscriptionRequest](docs/V1UpdateSubscriptionRequest.md)


## Documentation For Authorization



## OAuth2


- **Type**: OAuth
- **Flow**: password
- **Authorization URL**: 
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), subscription.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, subscription.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```



## Author



