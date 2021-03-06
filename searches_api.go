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
)

type SearchesApi struct {
	Configuration *Configuration
}

func NewSearchesApi() *SearchesApi {
	configuration := NewConfiguration()
	return &SearchesApi{
		Configuration: configuration,
	}
}

func NewSearchesApiWithBasePath(basePath string) *SearchesApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &SearchesApi{
		Configuration: configuration,
	}
}

/**
 * List your saved searches.
 * List your saved searches.
 *
 * @param userKey The &#x60;user_key&#x60; from [developer.meltwater.com](https://developer.meltwater.com/admin/applications/).
 * @param authorization &#x60;Oauth Access Token&#x60;    OAuth access token (RFC 6749). Must contain the access token type &#x60;Bearer&#x60;  followed by an OAuth access token.    #### Example:        Bearer KKwmfHwxsEoeMDTMAfxOpO...
 * @return *SearchesCollection
 */
func (a SearchesApi) GetAllSearches(userKey string, authorization string) (*SearchesCollection, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v2/searches"

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
	var successPayload = new(SearchesCollection)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetAllSearches", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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

