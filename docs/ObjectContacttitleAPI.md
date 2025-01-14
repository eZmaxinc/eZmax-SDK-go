# eZmaxAPI\ObjectContacttitleAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContacttitleGetAutocompleteV2**](ObjectContacttitleAPI.md#ContacttitleGetAutocompleteV2) | **Get** /2/object/contacttitle/getAutocomplete/{sSelector} | Retrieve Contacttitles and IDs



## ContacttitleGetAutocompleteV2

> ContacttitleGetAutocompleteV2Response ContacttitleGetAutocompleteV2(ctx, sSelector).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Contacttitles and IDs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	sSelector := "sSelector_example" // string | The type of Contacttitles to return
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectContacttitleAPI.ContacttitleGetAutocompleteV2(context.Background(), sSelector).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectContacttitleAPI.ContacttitleGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContacttitleGetAutocompleteV2`: ContacttitleGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectContacttitleAPI.ContacttitleGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Contacttitles to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiContacttitleGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**ContacttitleGetAutocompleteV2Response**](ContacttitleGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

