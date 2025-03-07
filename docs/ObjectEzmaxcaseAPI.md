# eZmaxAPI\ObjectEzmaxcaseAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzmaxcasePatchObjectV1**](ObjectEzmaxcaseAPI.md#EzmaxcasePatchObjectV1) | **Patch** /1/object/ezmaxcase/{pkiEzmaxcaseID} | Patch an existing Ezmaxcase



## EzmaxcasePatchObjectV1

> EzmaxcasePatchObjectV1Response EzmaxcasePatchObjectV1(ctx, pkiEzmaxcaseID).EzmaxcasePatchObjectV1Request(ezmaxcasePatchObjectV1Request).Execute()

Patch an existing Ezmaxcase



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
	pkiEzmaxcaseID := int32(56) // int32 | The unique ID of the Ezmaxcase
	ezmaxcasePatchObjectV1Request := *openapiclient.NewEzmaxcasePatchObjectV1Request(*openapiclient.NewEzmaxcaseRequestPatch()) // EzmaxcasePatchObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzmaxcaseAPI.EzmaxcasePatchObjectV1(context.Background(), pkiEzmaxcaseID).EzmaxcasePatchObjectV1Request(ezmaxcasePatchObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzmaxcaseAPI.EzmaxcasePatchObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxcasePatchObjectV1`: EzmaxcasePatchObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzmaxcaseAPI.EzmaxcasePatchObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzmaxcaseID** | **int32** | The unique ID of the Ezmaxcase | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxcasePatchObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezmaxcasePatchObjectV1Request** | [**EzmaxcasePatchObjectV1Request**](EzmaxcasePatchObjectV1Request.md) |  | 

### Return type

[**EzmaxcasePatchObjectV1Response**](EzmaxcasePatchObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

