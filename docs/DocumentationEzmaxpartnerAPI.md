# eZmaxAPI\DocumentationEzmaxpartnerAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocumentationSubscribeV1**](DocumentationEzmaxpartnerAPI.md#DocumentationSubscribeV1) | **Post** /1/documentation/subscribe | Subscribe to an Ezmaxparnerproductstage



## DocumentationSubscribeV1

> DocumentationSubscribeV1Response DocumentationSubscribeV1(ctx).DocumentationSubscribeV1Request(documentationSubscribeV1Request).Execute()

Subscribe to an Ezmaxparnerproductstage



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
	documentationSubscribeV1Request := *openapiclient.NewDocumentationSubscribeV1Request() // DocumentationSubscribeV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentationEzmaxpartnerAPI.DocumentationSubscribeV1(context.Background()).DocumentationSubscribeV1Request(documentationSubscribeV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentationEzmaxpartnerAPI.DocumentationSubscribeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentationSubscribeV1`: DocumentationSubscribeV1Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentationEzmaxpartnerAPI.DocumentationSubscribeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentationSubscribeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentationSubscribeV1Request** | [**DocumentationSubscribeV1Request**](DocumentationSubscribeV1Request.md) |  | 

### Return type

[**DocumentationSubscribeV1Response**](DocumentationSubscribeV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

