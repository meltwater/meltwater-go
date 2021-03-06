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

// Create an OAuth2 access token based on the provided `client_id` and `client_secret`.  #### Appendix    The Base64-encoded `client_id`:`client_secret` string can be generated in a  terminal with following command:        $ echo -n \"your_client_id:your_client_secret\" | base64    <i>You will need `base64` installed.</i>
type OAuth2Token struct {

	// The oauth2 access token to use for subsequent API calls
	AccessToken string `json:"access_token,omitempty"`

	// Time until the token expires in seconds
	ExpiresIn int32 `json:"expires_in,omitempty"`

	// The type of oauth2 scope the token is valid for
	Scope string `json:"scope,omitempty"`

	// The type/realm of the access token
	TokenType string `json:"token_type,omitempty"`
}
