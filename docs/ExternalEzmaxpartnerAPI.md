# eZmaxAPI\ExternalEzmaxpartnerAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzmaxpartnerSubscribeV1**](ExternalEzmaxpartnerAPI.md#EzmaxpartnerSubscribeV1) | **Post** /1/external/ezmaxpartner/subscribe | Subscribe to an Ezmaxparnerproductstage



## EzmaxpartnerSubscribeV1

> EzmaxpartnerSubscribeV1Response EzmaxpartnerSubscribeV1(ctx).EzmaxpartnerSubscribeV1Request(ezmaxpartnerSubscribeV1Request).Execute()

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
	ezmaxpartnerSubscribeV1Request := *openapiclient.NewEzmaxpartnerSubscribeV1Request() // EzmaxpartnerSubscribeV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalEzmaxpartnerAPI.EzmaxpartnerSubscribeV1(context.Background()).EzmaxpartnerSubscribeV1Request(ezmaxpartnerSubscribeV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEzmaxpartnerAPI.EzmaxpartnerSubscribeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxpartnerSubscribeV1`: EzmaxpartnerSubscribeV1Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalEzmaxpartnerAPI.EzmaxpartnerSubscribeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxpartnerSubscribeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezmaxpartnerSubscribeV1Request** | [**EzmaxpartnerSubscribeV1Request**](EzmaxpartnerSubscribeV1Request.md) |  | 

### Return type

[**EzmaxpartnerSubscribeV1Response**](EzmaxpartnerSubscribeV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

