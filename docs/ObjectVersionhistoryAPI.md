# eZmaxAPI\ObjectVersionhistoryAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VersionhistoryGetObjectV2**](ObjectVersionhistoryAPI.md#VersionhistoryGetObjectV2) | **Get** /2/object/versionhistory/{pkiVersionhistoryID} | Retrieve an existing Versionhistory



## VersionhistoryGetObjectV2

> VersionhistoryGetObjectV2Response VersionhistoryGetObjectV2(ctx, pkiVersionhistoryID).Execute()

Retrieve an existing Versionhistory



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
	pkiVersionhistoryID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectVersionhistoryAPI.VersionhistoryGetObjectV2(context.Background(), pkiVersionhistoryID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectVersionhistoryAPI.VersionhistoryGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VersionhistoryGetObjectV2`: VersionhistoryGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectVersionhistoryAPI.VersionhistoryGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiVersionhistoryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVersionhistoryGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VersionhistoryGetObjectV2Response**](VersionhistoryGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

