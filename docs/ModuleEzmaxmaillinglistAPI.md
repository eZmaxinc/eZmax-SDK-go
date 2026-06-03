# eZmaxAPI\ModuleEzmaxmaillinglistAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzmaxmaillinglistSubscribeV1**](ModuleEzmaxmaillinglistAPI.md#EzmaxmaillinglistSubscribeV1) | **Post** /1/module/ezmaxmaillinglist/subscribe | Subscribe to specific Ezmaxmaillinglist



## EzmaxmaillinglistSubscribeV1

> EzmaxmaillinglistSubscribeV1Response EzmaxmaillinglistSubscribeV1(ctx).EzmaxmaillinglistSubscribeV1Request(ezmaxmaillinglistSubscribeV1Request).Execute()

Subscribe to specific Ezmaxmaillinglist



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
	ezmaxmaillinglistSubscribeV1Request := *openapiclient.NewEzmaxmaillinglistSubscribeV1Request([]int32{int32(102)}) // EzmaxmaillinglistSubscribeV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModuleEzmaxmaillinglistAPI.EzmaxmaillinglistSubscribeV1(context.Background()).EzmaxmaillinglistSubscribeV1Request(ezmaxmaillinglistSubscribeV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModuleEzmaxmaillinglistAPI.EzmaxmaillinglistSubscribeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxmaillinglistSubscribeV1`: EzmaxmaillinglistSubscribeV1Response
	fmt.Fprintf(os.Stdout, "Response from `ModuleEzmaxmaillinglistAPI.EzmaxmaillinglistSubscribeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxmaillinglistSubscribeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezmaxmaillinglistSubscribeV1Request** | [**EzmaxmaillinglistSubscribeV1Request**](EzmaxmaillinglistSubscribeV1Request.md) |  | 

### Return type

[**EzmaxmaillinglistSubscribeV1Response**](EzmaxmaillinglistSubscribeV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

