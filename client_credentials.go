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

// Register new client             Creates a new pair of client credentials (`client_id`/`client_secret`          pair). Requires your Meltwater credentials (`email`:`password`)          to authenticate.           #### Appendix    The Base64-encoded `email`:`password` string can be generated in a terminal  with following command:        $ echo -n \"your_email@your_domain.com:your_secret_password\" | base64    <i>You will need `base64` installed.</i>
type ClientCredentials struct {

	// Id of the client
	ClientId string `json:"client_id,omitempty"`

	// The secret of the client
	ClientSecret string `json:"client_secret,omitempty"`
}
