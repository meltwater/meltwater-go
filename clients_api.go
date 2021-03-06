/* 
 * Meltwater Streaming API v2
 *
 * The Meltwater Streaming API provides the needed resources for Meltwater clients to create & delete REST Hooks and stream Meltwater search results to your specified destination.
 *
 * OpenAPI spec version: 2.0.0
 * Contact: support@api.meltwater.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

import (
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
)

type ClientsApi struct {
	Configuration *Configuration
}

func NewClientsApi() *ClientsApi {
	configuration := NewConfiguration()
	return &ClientsApi{
		Configuration: configuration,
	}
}

func NewClientsApiWithBasePath(basePath string) *ClientsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &ClientsApi{
		Configuration: configuration,
	}
}

/**
 * Register new client
 * Register new client             Creates a new pair of client credentials (&#x60;client_id&#x60;/&#x60;client_secret&#x60;          pair). Requires your Meltwater credentials (&#x60;email&#x60;:&#x60;password&#x60;)          to authenticate.           #### Appendix    The Base64-encoded &#x60;email&#x60;:&#x60;password&#x60; string can be generated in a terminal  with following command:        $ echo -n \&quot;your_email@your_domain.com:your_secret_password\&quot; | base64    &lt;i&gt;You will need &#x60;base64&#x60; installed.&lt;/i&gt;
 *
 * @param userKey The &#x60;user_key&#x60; from [developer.meltwater.com](https://developer.meltwater.com/admin/applications/).
 * @param authorization &#x60;email&#x60;:&#x60;password&#x60;    Basic Auth (RFC2617) credentials. Must contain the realm &#x60;Basic&#x60; followed by a  Base64-encoded &#x60;email&#x60;:&#x60;password&#x60; pair using your Meltwater credentials.    #### Example:        Basic bXlfZW1haWxAZXhhbXJzZWNyZXQ&#x3D;
 * @return *ClientCredentials
 */
func (a ClientsApi) CreateClientCredentials(userKey string, authorization string) (*ClientCredentials, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v2/clients"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "user-key"
	localVarHeaderParams["user-key"] = a.Configuration.APIClient.ParameterToString(userKey, "")
	// header params "Authorization"
	localVarHeaderParams["Authorization"] = a.Configuration.APIClient.ParameterToString(authorization, "")
	var successPayload = new(ClientCredentials)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateClientCredentials", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Delete client.
 * Delete client.    Deletes your current client credentials consisting                 of client_id and client_secret. After calling this resource,                  you will not be able to use the Meltwater API unless you create                  a new set of client credentials! Requires your Meltwater                  credentials (&#x60;email&#x60;:&#x60;password&#x60;) to authenticate.           #### Appendix    The Base64-encoded &#x60;email&#x60;:&#x60;password&#x60; string can be generated in a terminal  with following command:        $ echo -n \&quot;your_email@your_domain.com:your_secret_password\&quot; | base64    &lt;i&gt;You will need &#x60;base64&#x60; installed.&lt;/i&gt;
 *
 * @param userKey The &#x60;user_key&#x60; from [developer.meltwater.com](https://developer.meltwater.com/admin/applications/).
 * @param authorization &#x60;email&#x60;:&#x60;password&#x60;    Basic Auth (RFC2617) credentials. Must contain the realm &#x60;Basic&#x60; followed by a  Base64-encoded &#x60;email&#x60;:&#x60;password&#x60; pair using your Meltwater credentials.    #### Example:        Basic bXlfZW1haWxAZXhhbXJzZWNyZXQ&#x3D;
 * @param clientId Client ID
 * @return void
 */
func (a ClientsApi) DeleteClientCredentials(userKey string, authorization string, clientId string) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v2/clients/{client_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"client_id"+"}", fmt.Sprintf("%v", clientId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "user-key"
	localVarHeaderParams["user-key"] = a.Configuration.APIClient.ParameterToString(userKey, "")
	// header params "Authorization"
	localVarHeaderParams["Authorization"] = a.Configuration.APIClient.ParameterToString(authorization, "")
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteClientCredentials", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return localVarAPIResponse, err
	}
	return localVarAPIResponse, err
}

