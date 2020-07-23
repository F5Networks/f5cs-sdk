# \AuthenticationServiceApi

All URIs are relative to *https://api.cloudservices.f5.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Login**](AuthenticationServiceApi.md#Login) | **Post** /v1/svc-auth/login | Login to F5 Cloud Services
[**Logout**](AuthenticationServiceApi.md#Logout) | **Post** /v1/svc-auth/logout | Logout from F5 Cloud Services
[**ReLogin**](AuthenticationServiceApi.md#ReLogin) | **Post** /v1/svc-auth/relogin | Login again to F5 Cloud Services



## Login

> AuthenticationServiceLoginReply Login(ctx, authenticationServiceLoginRequest)

Login to F5 Cloud Services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticationServiceLoginRequest** | [**AuthenticationServiceLoginRequest**](AuthenticationServiceLoginRequest.md)|  | 

### Return type

[**AuthenticationServiceLoginReply**](authentication_serviceLoginReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> map[string]interface{} Logout(ctx, authenticationServiceLogoutRequest)

Logout from F5 Cloud Services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticationServiceLogoutRequest** | [**AuthenticationServiceLogoutRequest**](AuthenticationServiceLogoutRequest.md)|  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReLogin

> AuthenticationServiceReLoginReply ReLogin(ctx, authenticationServiceReLoginRequest)

Login again to F5 Cloud Services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticationServiceReLoginRequest** | [**AuthenticationServiceReLoginRequest**](AuthenticationServiceReLoginRequest.md)|  | 

### Return type

[**AuthenticationServiceReLoginReply**](authentication_serviceReLoginReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

