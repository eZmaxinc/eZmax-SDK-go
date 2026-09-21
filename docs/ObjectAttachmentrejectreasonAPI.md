# eZmaxAPI\ObjectAttachmentrejectreasonAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachmentrejectreasonGetAutocompleteV2**](ObjectAttachmentrejectreasonAPI.md#AttachmentrejectreasonGetAutocompleteV2) | **Get** /2/object/attachmentrejectreason/getAutocomplete/{sSelector} | Retrieve Attachmentrejectreasons and IDs



## AttachmentrejectreasonGetAutocompleteV2

> AttachmentrejectreasonGetAutocompleteV2Response AttachmentrejectreasonGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Attachmentrejectreasons and IDs



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
	sSelector := "sSelector_example" // string | The type of Attachmentrejectreasons to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAttachmentrejectreasonAPI.AttachmentrejectreasonGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAttachmentrejectreasonAPI.AttachmentrejectreasonGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentrejectreasonGetAutocompleteV2`: AttachmentrejectreasonGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAttachmentrejectreasonAPI.AttachmentrejectreasonGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Attachmentrejectreasons to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentrejectreasonGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**AttachmentrejectreasonGetAutocompleteV2Response**](AttachmentrejectreasonGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

