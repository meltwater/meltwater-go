/* 
 * The Meltwater API
 *
 * The Meltwater API provides the needed resources for Meltwater clients to create & delete REST Hooks and stream Meltwater search results to your specified destination.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: support@api.meltwater.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

// Create an OAuth2 access token based on the provided `client_id` and `client_secret`
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
