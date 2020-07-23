# \SubscriptionServiceApi

All URIs are relative to *https://api.cloudservices.f5.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSubscription**](SubscriptionServiceApi.md#ActivateSubscription) | **Post** /v1/svc-subscription/subscriptions/{subscription_id}/activate | Activate a subscription
[**BatchActivateSubscription**](SubscriptionServiceApi.md#BatchActivateSubscription) | **Post** /v1/svc-subscription/subscriptions/batch-activate | Activate multiple subscriptions
[**BatchRetireSubscription**](SubscriptionServiceApi.md#BatchRetireSubscription) | **Post** /v1/svc-subscription/subscriptions/batch-retire | Retire multiple subscriptions
[**BatchSubscriptionStatus**](SubscriptionServiceApi.md#BatchSubscriptionStatus) | **Post** /v1/svc-subscription/subscriptions/batch-status | Get the status of multiple subscriptions
[**BatchSuspendSubscription**](SubscriptionServiceApi.md#BatchSuspendSubscription) | **Post** /v1/svc-subscription/subscriptions/batch-suspend | Suspend multiple subscriptions
[**BatchUnretireSubscription**](SubscriptionServiceApi.md#BatchUnretireSubscription) | **Post** /v1/svc-subscription/subscriptions/batch-unretire | Un-retire multiple subscriptions
[**CreateSubscription**](SubscriptionServiceApi.md#CreateSubscription) | **Post** /v1/svc-subscription/subscriptions | Create a subscription
[**GetSubscription**](SubscriptionServiceApi.md#GetSubscription) | **Get** /v1/svc-subscription/subscriptions/{subscription_id} | Get a subscription
[**ListSubscriptions**](SubscriptionServiceApi.md#ListSubscriptions) | **Get** /v1/svc-subscription/subscriptions | List subscriptions in an account
[**RetireSubscription**](SubscriptionServiceApi.md#RetireSubscription) | **Post** /v1/svc-subscription/subscriptions/{subscription_id}/retire | Retire a subscription
[**SubscriptionStatus**](SubscriptionServiceApi.md#SubscriptionStatus) | **Get** /v1/svc-subscription/subscriptions/{subscription_id}/status | Get a subscription&#39;s status
[**SuspendSubscription**](SubscriptionServiceApi.md#SuspendSubscription) | **Post** /v1/svc-subscription/subscriptions/{subscription_id}/suspend | Suspend a subscription
[**TestSubscription**](SubscriptionServiceApi.md#TestSubscription) | **Post** /v1/svc-subscription/subscriptions/{subscription_id}/test | Test a subscription
[**UnretireSubscription**](SubscriptionServiceApi.md#UnretireSubscription) | **Post** /v1/svc-subscription/subscriptions/{subscription_id}/unretire | Un-retire a subscription
[**UpdateSubscription**](SubscriptionServiceApi.md#UpdateSubscription) | **Put** /v1/svc-subscription/subscriptions/{subscription_id} | Update a subscription



## ActivateSubscription

> V1SubscriptionStatusResponse ActivateSubscription(ctx, subscriptionId)

Activate a subscription

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**| subscription identifier | 

### Return type

[**V1SubscriptionStatusResponse**](v1SubscriptionStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchActivateSubscription

> V1BatchSubscriptionStatusResponse BatchActivateSubscription(ctx, v1BatchSubscriptionIdRequest)

Activate multiple subscriptions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1BatchSubscriptionIdRequest** | [**V1BatchSubscriptionIdRequest**](V1BatchSubscriptionIdRequest.md)|  | 

### Return type

[**V1BatchSubscriptionStatusResponse**](v1BatchSubscriptionStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchRetireSubscription

> V1BatchSubscriptionStatusResponse BatchRetireSubscription(ctx, v1BatchSubscriptionIdRequest)

Retire multiple subscriptions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1BatchSubscriptionIdRequest** | [**V1BatchSubscriptionIdRequest**](V1BatchSubscriptionIdRequest.md)|  | 

### Return type

[**V1BatchSubscriptionStatusResponse**](v1BatchSubscriptionStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchSubscriptionStatus

> V1BatchSubscriptionStatusResponse BatchSubscriptionStatus(ctx, v1BatchSubscriptionIdRequest)

Get the status of multiple subscriptions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1BatchSubscriptionIdRequest** | [**V1BatchSubscriptionIdRequest**](V1BatchSubscriptionIdRequest.md)|  | 

### Return type

[**V1BatchSubscriptionStatusResponse**](v1BatchSubscriptionStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchSuspendSubscription

> V1BatchSubscriptionStatusResponse BatchSuspendSubscription(ctx, v1BatchSubscriptionIdRequest)

Suspend multiple subscriptions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1BatchSubscriptionIdRequest** | [**V1BatchSubscriptionIdRequest**](V1BatchSubscriptionIdRequest.md)|  | 

### Return type

[**V1BatchSubscriptionStatusResponse**](v1BatchSubscriptionStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchUnretireSubscription

> V1BatchSubscriptionStatusResponse BatchUnretireSubscription(ctx, v1BatchSubscriptionIdRequest)

Un-retire multiple subscriptions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1BatchSubscriptionIdRequest** | [**V1BatchSubscriptionIdRequest**](V1BatchSubscriptionIdRequest.md)|  | 

### Return type

[**V1BatchSubscriptionStatusResponse**](v1BatchSubscriptionStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscription

> V1Subscription CreateSubscription(ctx, v1CreateSubscriptionRequest)

Create a subscription

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1CreateSubscriptionRequest** | [**V1CreateSubscriptionRequest**](V1CreateSubscriptionRequest.md)|  | 

### Return type

[**V1Subscription**](v1Subscription.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscription

> V1Subscription GetSubscription(ctx, subscriptionId, optional)

Get a subscription

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**| subscription identifier | 
 **optional** | ***GetSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **omitConfig** | **optional.Bool**|  | [default to false]

### Return type

[**V1Subscription**](v1Subscription.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptions

> V1Subscriptions ListSubscriptions(ctx, accountId, optional)

List subscriptions in an account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| filter by account identifier. | 
 **optional** | ***ListSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSubscriptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **optional.String**| filter by status. | [default to _allStatusFilter]
 **serviceInstanceId** | **optional.String**| filter by service instance identifier. | 
 **catalogId** | **optional.String**| filter by catalog identifier. | 
 **deleted** | **optional.Bool**| false (default) to filter deleted subscriptions. | 
 **serviceType** | **optional.String**| filter by service_type. | 
 **omitConfig** | **optional.Bool**| choose whether to include configurations in response. | [default to false]
 **pageSize** | **optional.Int32**| The maximum number of items to return, capped at 500. | 
 **pageNumber** | **optional.Int32**| One-indexed page number, the default of 0 means to follow the page_token page. | 
 **pageToken** | **optional.String**| The page_token to use returned in Subscriptions.page_token. | 

### Return type

[**V1Subscriptions**](v1Subscriptions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetireSubscription

> V1SubscriptionStatusResponse RetireSubscription(ctx, subscriptionId)

Retire a subscription

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**| subscription identifier | 

### Return type

[**V1SubscriptionStatusResponse**](v1SubscriptionStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionStatus

> V1SubscriptionStatusResponse SubscriptionStatus(ctx, subscriptionId)

Get a subscription's status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**| subscription identifier | 

### Return type

[**V1SubscriptionStatusResponse**](v1SubscriptionStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendSubscription

> V1SubscriptionStatusResponse SuspendSubscription(ctx, subscriptionId)

Suspend a subscription

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**| subscription identifier | 

### Return type

[**V1SubscriptionStatusResponse**](v1SubscriptionStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSubscription

> V1TestSubscriptionResponse TestSubscription(ctx, subscriptionId)

Test a subscription

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**| subscription identifier | 

### Return type

[**V1TestSubscriptionResponse**](v1TestSubscriptionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnretireSubscription

> V1SubscriptionStatusResponse UnretireSubscription(ctx, subscriptionId)

Un-retire a subscription

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**| subscription identifier | 

### Return type

[**V1SubscriptionStatusResponse**](v1SubscriptionStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> V1Subscription UpdateSubscription(ctx, subscriptionId, v1UpdateSubscriptionRequest)

Update a subscription

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**| subscription identifier | 
**v1UpdateSubscriptionRequest** | [**V1UpdateSubscriptionRequest**](V1UpdateSubscriptionRequest.md)|  | 

### Return type

[**V1Subscription**](v1Subscription.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

