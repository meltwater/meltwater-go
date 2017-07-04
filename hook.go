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

// Creates a hook for one of your predefined searches.  Delivers search results for the query referenced by the `search_id` to the `target_url` via HTTP POST. Note that a `hook_id` will be returned on successful creation, this id is needed to delete the hook.   We are also returning the header: `X-Hook-Secret`, a shared secret used to sign the document body pushed to your `target_url`.    Requires an OAuth access token.
type Hook struct {

	// Search id
	SearchId int32 `json:"search_id,omitempty"`

	// The URL that results from the search will be posted to
	TargetUrl string `json:"target_url,omitempty"`

	Updated string `json:"updated,omitempty"`

	// The type of search the hook is for
	SearchType string `json:"search_type,omitempty"`

	HookId string `json:"hook_id,omitempty"`
}
