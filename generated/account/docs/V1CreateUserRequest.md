# V1CreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Phone** | **string** |  | [optional] 
**TimeZone** | **string** |  | [optional] 
**PreferredLanguage** | **string** |  | [optional] 
**InviteToken** | **string** |  | [optional] 
**Password** | **string** |  | 
**Account** | [**V1CreateAccountRequest**](v1CreateAccountRequest.md) |  | 
**SignupPassphrase** | **string** | Signup passphrase. Not case-sensitive. | [optional] 
**SignupProvider** | **string** | This should be a valid signup provider, set to &#39;aws&#39; if using AWS Marketplace. If signup_provider is specified, signup_token also must be supplied. | [optional] 
**SignupToken** | **string** | This should be a valid token, set to value of x-amzn-marketplace-token for AWS Marketplace. | [optional] 
**Terms** | [**V1TermsAcceptanceRequest**](v1TermsAcceptanceRequest.md) |  | [optional] 
**TermsOfService** | [**[]V1CreateTermsAcceptanceRequest**](v1CreateTermsAcceptanceRequest.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


