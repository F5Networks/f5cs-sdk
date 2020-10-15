# \AccountServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptInvite**](AccountServiceApi.md#AcceptInvite) | **Post** /v1/svc-account/invites/{invite_token}/accept | Accept an invite
[**AcceptSignup**](AccountServiceApi.md#AcceptSignup) | **Post** /v1/svc-account/signup/accept | Accept signup
[**AddAccountMember**](AccountServiceApi.md#AddAccountMember) | **Post** /v1/svc-account/accounts/{account_id}/members | Add an account member
[**BatchCreateInvites**](AccountServiceApi.md#BatchCreateInvites) | **Post** /v1/svc-account/invites/batch-create | Batch create multiple invites
[**BatchGetAccountMembers**](AccountServiceApi.md#BatchGetAccountMembers) | **Post** /v1/svc-account/accounts/-/members/batch-get | Batch get members from multiple accounts by account IDs
[**BatchGetAccounts**](AccountServiceApi.md#BatchGetAccounts) | **Post** /v1/svc-account/accounts/batch-get | Batch get multiple accounts by ids
[**ConfirmEmail**](AccountServiceApi.md#ConfirmEmail) | **Post** /v1/svc-account/users/confirm-email | Confirm email
[**ConfirmResetPassword**](AccountServiceApi.md#ConfirmResetPassword) | **Post** /v1/svc-account/users/confirm-reset-password | Confirm password reset
[**CreateAccount**](AccountServiceApi.md#CreateAccount) | **Post** /v1/svc-account/accounts | Create a new account
[**CreateInvite**](AccountServiceApi.md#CreateInvite) | **Post** /v1/svc-account/invites | Create a new invite into an account
[**CreateSubAccount**](AccountServiceApi.md#CreateSubAccount) | **Post** /v1/svc-account/accounts/{parent_account_id}/subaccounts | Create a sub-account
[**CreateUser**](AccountServiceApi.md#CreateUser) | **Post** /v1/svc-account/users | Create a new user
[**CreateUserTermsAcceptance**](AccountServiceApi.md#CreateUserTermsAcceptance) | **Post** /v1/svc-account/users/{user_id}/terms | Create terms of service acceptance
[**DeleteAccount**](AccountServiceApi.md#DeleteAccount) | **Delete** /v1/svc-account/accounts/{id} | Delete an account
[**DeleteAccountMember**](AccountServiceApi.md#DeleteAccountMember) | **Delete** /v1/svc-account/accounts/{account_id}/members/{user_id} | Delete an account member
[**DeleteInvite**](AccountServiceApi.md#DeleteInvite) | **Delete** /v1/svc-account/invites/{invite_id} | Delete an invite
[**DeleteUser**](AccountServiceApi.md#DeleteUser) | **Delete** /v1/svc-account/users/{id} | Delete a user
[**DisableCatalogItemForAccount**](AccountServiceApi.md#DisableCatalogItemForAccount) | **Delete** /v1/svc-account/accounts/{account_id}/catalogs/{catalog_id} | Disable catalog item for an account
[**EnableCatalogItemForAccount**](AccountServiceApi.md#EnableCatalogItemForAccount) | **Post** /v1/svc-account/accounts/{account_id}/catalogs | Enable catalog item for an account
[**GetAccount**](AccountServiceApi.md#GetAccount) | **Get** /v1/svc-account/accounts/{id} | Get account
[**GetAccountMember**](AccountServiceApi.md#GetAccountMember) | **Get** /v1/svc-account/accounts/{account_id}/members/{user_id} | Get an account member
[**GetCurrentUser**](AccountServiceApi.md#GetCurrentUser) | **Get** /v1/svc-account/user | Get currently authenticated user
[**GetRole**](AccountServiceApi.md#GetRole) | **Get** /v1/svc-account/roles/{id} | Get role
[**GetUser**](AccountServiceApi.md#GetUser) | **Get** /v1/svc-account/users/{id} | Get a user
[**GetUserHistory**](AccountServiceApi.md#GetUserHistory) | **Get** /v1/svc-account/users/{id}/history | Get user history
[**GetUserTermsAcceptance**](AccountServiceApi.md#GetUserTermsAcceptance) | **Get** /v1/svc-account/users/{user_id}/terms/{name} | Get terms of service acceptance
[**ListAccountMembers**](AccountServiceApi.md#ListAccountMembers) | **Get** /v1/svc-account/accounts/{account_id}/members | List account members
[**ListCatalogItemsForAccount**](AccountServiceApi.md#ListCatalogItemsForAccount) | **Get** /v1/svc-account/accounts/{account_id}/catalogs | List catalog items for an account
[**ListInvites**](AccountServiceApi.md#ListInvites) | **Get** /v1/svc-account/invites | List invites
[**ListRoles**](AccountServiceApi.md#ListRoles) | **Get** /v1/svc-account/roles | List roles
[**ListSubAccounts**](AccountServiceApi.md#ListSubAccounts) | **Get** /v1/svc-account/accounts/{parent_account_id}/subaccounts | List sub-accounts
[**ListUserMemberships**](AccountServiceApi.md#ListUserMemberships) | **Get** /v1/svc-account/users/{user_id}/memberships | List user memberships
[**ListUserTermsAcceptance**](AccountServiceApi.md#ListUserTermsAcceptance) | **Get** /v1/svc-account/users/{user_id}/terms | List terms of service acceptance
[**MigratePreviewAccount**](AccountServiceApi.md#MigratePreviewAccount) | **Post** /v1/svc-account/accounts/{id}/migrate | Migrate preview account
[**ResetPassword**](AccountServiceApi.md#ResetPassword) | **Post** /v1/svc-account/users/reset-password | Reset password
[**SendConfirmEmail**](AccountServiceApi.md#SendConfirmEmail) | **Post** /v1/svc-account/users/send-confirm-email | Send confirmation email
[**UpdateAccount**](AccountServiceApi.md#UpdateAccount) | **Put** /v1/svc-account/accounts/{id} | Update an account
[**UpdateAccountMember**](AccountServiceApi.md#UpdateAccountMember) | **Put** /v1/svc-account/accounts/{account_id}/members/{user_id} | Update an account member
[**UpdateUser**](AccountServiceApi.md#UpdateUser) | **Put** /v1/svc-account/users/{id} | Update a user
[**UpdateUserTermsAcceptance**](AccountServiceApi.md#UpdateUserTermsAcceptance) | **Put** /v1/svc-account/users/{user_id}/terms/{name} | Update terms of service acceptance
[**ValidateSignupToken**](AccountServiceApi.md#ValidateSignupToken) | **Post** /v1/svc-account/signup/validate | Validate signup token



## AcceptInvite

> V1AcceptInviteResponse AcceptInvite(ctx, inviteToken, v1InviteeRequest)

Accept an invite

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteToken** | **string**| Invite token | 
**v1InviteeRequest** | [**V1InviteeRequest**](V1InviteeRequest.md)|  | 

### Return type

[**V1AcceptInviteResponse**](v1AcceptInviteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcceptSignup

> V1AcceptSignupResponse AcceptSignup(ctx, v1AcceptSignupRequest)

Accept signup

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1AcceptSignupRequest** | [**V1AcceptSignupRequest**](V1AcceptSignupRequest.md)|  | 

### Return type

[**V1AcceptSignupResponse**](v1AcceptSignupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAccountMember

> V1AccountMember AddAccountMember(ctx, accountId, v1AddAccountMemberRequest)

Add an account member

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**v1AddAccountMemberRequest** | [**V1AddAccountMemberRequest**](V1AddAccountMemberRequest.md)|  | 

### Return type

[**V1AccountMember**](v1AccountMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCreateInvites

> V1Invites BatchCreateInvites(ctx, v1BatchCreateInviteRequest)

Batch create multiple invites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1BatchCreateInviteRequest** | [**V1BatchCreateInviteRequest**](V1BatchCreateInviteRequest.md)|  | 

### Return type

[**V1Invites**](v1Invites.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchGetAccountMembers

> V1BatchGetAccountMembersResponse BatchGetAccountMembers(ctx, v1BatchGetAccountMembersRequest)

Batch get members from multiple accounts by account IDs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1BatchGetAccountMembersRequest** | [**V1BatchGetAccountMembersRequest**](V1BatchGetAccountMembersRequest.md)|  | 

### Return type

[**V1BatchGetAccountMembersResponse**](v1BatchGetAccountMembersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchGetAccounts

> V1BatchGetAccountsResponse BatchGetAccounts(ctx, v1BatchGetAccountsRequest)

Batch get multiple accounts by ids

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1BatchGetAccountsRequest** | [**V1BatchGetAccountsRequest**](V1BatchGetAccountsRequest.md)|  | 

### Return type

[**V1BatchGetAccountsResponse**](v1BatchGetAccountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmEmail

> map[string]interface{} ConfirmEmail(ctx, v1ConfirmEmailRequest)

Confirm email

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1ConfirmEmailRequest** | [**V1ConfirmEmailRequest**](V1ConfirmEmailRequest.md)|  | 

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


## ConfirmResetPassword

> map[string]interface{} ConfirmResetPassword(ctx, v1ConfirmResetPasswordRequest)

Confirm password reset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1ConfirmResetPasswordRequest** | [**V1ConfirmResetPasswordRequest**](V1ConfirmResetPasswordRequest.md)|  | 

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


## CreateAccount

> V1Account CreateAccount(ctx, v1CreateAccountRequest)

Create a new account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1CreateAccountRequest** | [**V1CreateAccountRequest**](V1CreateAccountRequest.md)|  | 

### Return type

[**V1Account**](v1Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvite

> V1Invite CreateInvite(ctx, v1CreateInviteRequest)

Create a new invite into an account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1CreateInviteRequest** | [**V1CreateInviteRequest**](V1CreateInviteRequest.md)|  | 

### Return type

[**V1Invite**](v1Invite.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubAccount

> V1Account CreateSubAccount(ctx, parentAccountId, v1CreateAccountRequest)

Create a sub-account

Create a new account with a specified parent account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentAccountId** | **string**| Parent account ID  If parent account ID is specified, this operation will create a sub-account under parent account. | 
**v1CreateAccountRequest** | [**V1CreateAccountRequest**](V1CreateAccountRequest.md)|  | 

### Return type

[**V1Account**](v1Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> V1User CreateUser(ctx, v1CreateUserRequest)

Create a new user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1CreateUserRequest** | [**V1CreateUserRequest**](V1CreateUserRequest.md)|  | 

### Return type

[**V1User**](v1User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserTermsAcceptance

> V1TermsAcceptance CreateUserTermsAcceptance(ctx, userId, v1TermsAcceptanceRequest)

Create terms of service acceptance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User ID for terms of service acceptance. | 
**v1TermsAcceptanceRequest** | [**V1TermsAcceptanceRequest**](V1TermsAcceptanceRequest.md)|  | 

### Return type

[**V1TermsAcceptance**](v1TermsAcceptance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> map[string]interface{} DeleteAccount(ctx, id, optional)

Delete an account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***DeleteAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **optional.Bool**|  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountMember

> map[string]interface{} DeleteAccountMember(ctx, accountId, userId)

Delete an account member

Removes a member from an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**userId** | **string**|  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInvite

> map[string]interface{} DeleteInvite(ctx, inviteId)

Delete an invite

Rescind an invitation to an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteId** | **string**|  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> map[string]interface{} DeleteUser(ctx, id)

Delete a user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableCatalogItemForAccount

> map[string]interface{} DisableCatalogItemForAccount(ctx, accountId, catalogId)

Disable catalog item for an account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**catalogId** | **string**|  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableCatalogItemForAccount

> V1AccountCatalogItem EnableCatalogItemForAccount(ctx, accountId, v1EnableCatalogItemForAccountRequest)

Enable catalog item for an account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**v1EnableCatalogItemForAccountRequest** | [**V1EnableCatalogItemForAccountRequest**](V1EnableCatalogItemForAccountRequest.md)|  | 

### Return type

[**V1AccountCatalogItem**](v1AccountCatalogItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> V1Account GetAccount(ctx, id)

Get account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**V1Account**](v1Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountMember

> V1AccountMember GetAccountMember(ctx, accountId, userId)

Get an account member

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**userId** | **string**|  | 

### Return type

[**V1AccountMember**](v1AccountMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUser

> V1User GetCurrentUser(ctx, )

Get currently authenticated user

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1User**](v1User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> V1Role GetRole(ctx, id)

Get role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**V1Role**](v1Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> V1User GetUser(ctx, id)

Get a user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**V1User**](v1User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserHistory

> V1User GetUserHistory(ctx, id)

Get user history

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**V1User**](v1User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTermsAcceptance

> V1TermsAcceptance GetUserTermsAcceptance(ctx, userId, name)

Get terms of service acceptance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 
**name** | **string**|  | 

### Return type

[**V1TermsAcceptance**](v1TermsAcceptance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountMembers

> V1ListAccountMembersResponse ListAccountMembers(ctx, accountId)

List account members

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

### Return type

[**V1ListAccountMembersResponse**](v1ListAccountMembersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCatalogItemsForAccount

> V1ListCatalogItemsForAccountResponse ListCatalogItemsForAccount(ctx, accountId)

List catalog items for an account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

### Return type

[**V1ListCatalogItemsForAccountResponse**](v1ListCatalogItemsForAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvites

> V1Invites ListInvites(ctx, optional)

List invites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListInvitesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListInvitesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteeEmail** | **optional.String**|  | 

### Return type

[**V1Invites**](v1Invites.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> V1ListRolesResponse ListRoles(ctx, )

List roles

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1ListRolesResponse**](v1ListRolesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubAccounts

> V1ListAccountsResponse ListSubAccounts(ctx, parentAccountId, optional)

List sub-accounts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentAccountId** | **string**|  | 
 **optional** | ***ListSubAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSubAccountsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **optional.String**|  | 

### Return type

[**V1ListAccountsResponse**](v1ListAccountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserMemberships

> V1ListUserMembershipsResponse ListUserMemberships(ctx, userId)

List user memberships

List accounts user has access to.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

### Return type

[**V1ListUserMembershipsResponse**](v1ListUserMembershipsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserTermsAcceptance

> V1TermsAcceptanceList ListUserTermsAcceptance(ctx, userId)

List terms of service acceptance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

### Return type

[**V1TermsAcceptanceList**](v1TermsAcceptanceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigratePreviewAccount

> map[string]interface{} MigratePreviewAccount(ctx, id, v1AccountIdRequest)

Migrate preview account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**v1AccountIdRequest** | [**V1AccountIdRequest**](V1AccountIdRequest.md)|  | 

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


## ResetPassword

> map[string]interface{} ResetPassword(ctx, v1ResetPasswordRequest)

Reset password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1ResetPasswordRequest** | [**V1ResetPasswordRequest**](V1ResetPasswordRequest.md)|  | 

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


## SendConfirmEmail

> map[string]interface{} SendConfirmEmail(ctx, v1SendConfirmEmailRequest)

Send confirmation email

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1SendConfirmEmailRequest** | [**V1SendConfirmEmailRequest**](V1SendConfirmEmailRequest.md)|  | 

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


## UpdateAccount

> V1Account UpdateAccount(ctx, id, v1Account)

Update an account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| Account ID | 
**v1Account** | [**V1Account**](V1Account.md)|  | 

### Return type

[**V1Account**](v1Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountMember

> V1AccountMember UpdateAccountMember(ctx, accountId, userId, v1UpdateAccountMemberRequest)

Update an account member

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**userId** | **string**|  | 
**v1UpdateAccountMemberRequest** | [**V1UpdateAccountMemberRequest**](V1UpdateAccountMemberRequest.md)|  | 

### Return type

[**V1AccountMember**](v1AccountMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> V1User UpdateUser(ctx, id, v1User)

Update a user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**v1User** | [**V1User**](V1User.md)|  | 

### Return type

[**V1User**](v1User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserTermsAcceptance

> V1TermsAcceptance UpdateUserTermsAcceptance(ctx, userId, name, v1TermsAcceptanceRequest)

Update terms of service acceptance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| User ID for terms of service acceptance. | 
**name** | **string**| Terms of service name  Current values: &#x60;cloud_services&#x60;, &#x60;business_entity&#x60;. | 
**v1TermsAcceptanceRequest** | [**V1TermsAcceptanceRequest**](V1TermsAcceptanceRequest.md)|  | 

### Return type

[**V1TermsAcceptance**](v1TermsAcceptance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateSignupToken

> map[string]interface{} ValidateSignupToken(ctx, v1ValidateSignupTokenRequest)

Validate signup token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v1ValidateSignupTokenRequest** | [**V1ValidateSignupTokenRequest**](V1ValidateSignupTokenRequest.md)|  | 

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

