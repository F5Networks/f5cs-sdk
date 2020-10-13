# \DashboardServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountServices**](DashboardServiceApi.md#GetAccountServices) | **Get** /v1/svc-dashboard/dashboard/accounts/{account_id} | 
[**GetAccountSummary**](DashboardServiceApi.md#GetAccountSummary) | **Get** /v1/svc-dashboard/dashboard/accounts/{account_id}/summary | 
[**GetServiceDetails**](DashboardServiceApi.md#GetServiceDetails) | **Get** /v1/svc-dashboard/dashboard/accounts/{account_id}/services/{service_id} | 



## GetAccountServices

> DashboardServiceGetAccountServicesReply GetAccountServices(ctx, accountId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***GetAccountServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAccountServicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duration** | **optional.String**|  | 
 **type_** | **optional.String**|  | 

### Return type

[**DashboardServiceGetAccountServicesReply**](dashboard_serviceGetAccountServicesReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountSummary

> DashboardServiceGetAccountSummaryReply GetAccountSummary(ctx, accountId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***GetAccountSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAccountSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duration** | **optional.String**|  | 
 **type_** | **optional.String**|  | 

### Return type

[**DashboardServiceGetAccountSummaryReply**](dashboard_serviceGetAccountSummaryReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceDetails

> DashboardServiceGetServiceDetailsReply GetServiceDetails(ctx, accountId, serviceId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**serviceId** | **string**|  | 

### Return type

[**DashboardServiceGetServiceDetailsReply**](dashboard_serviceGetServiceDetailsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

