# \Oauth2Api

All URIs are relative to *https://api.meltwater.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](Oauth2Api.md#CreateToken) | **Post** /oauth2/token | Create an access token


# **CreateToken**
> OAuth2Token CreateToken($userKey, $authorization, $grantType, $scope)

Create an access token

Create an OAuth2 access token based on the provided `client_id` and `client_secret`.  #### Appendix    The Base64-encoded `client_id`:`client_secret` string can be generated in a  terminal with following command:        $ echo -n \"your_client_id:your_client_secret\" | base64    <i>You will need `base64` installed.</i>


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userKey** | **string**| The &#x60;user_key&#x60; from [developer.meltwater.com](https://developer.meltwater.com/admin/applications/). | 
 **authorization** | **string**| &#x60;client_id:client_secret&#x60;  Basic Auth (RFC2617) credentials. Must contain the realm &#x60;Basic&#x60; followed by a Base64-encoded &#x60;client_id&#x60;:&#x60;client_secret&#x60; pair.   #### Example:      Basic aAlfbb1haWxDSXhhDXxxZWKJAyZXQ&#x3D; | 
 **grantType** | **string**| OAuth2 grant type, use &#x60;client_credentials&#x60; | 
 **scope** | **string**| OAuth2 scope, use &#x60;search&#x60; | 

### Return type

[**OAuth2Token**](OAuth2Token.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

