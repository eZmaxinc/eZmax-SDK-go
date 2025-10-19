# eZmaxAPI\ObjectEzmaxcustomerAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzmaxcustomerPatchObjectV1**](ObjectEzmaxcustomerAPI.md#EzmaxcustomerPatchObjectV1) | **Patch** /1/object/ezmaxcustomer/{pkiEzmaxcustomerID} | Patch an existing Ezmaxcustomer



## EzmaxcustomerPatchObjectV1

> EzmaxcustomerPatchObjectV1Response EzmaxcustomerPatchObjectV1(ctx, pkiEzmaxcustomerID).EzmaxcustomerPatchObjectV1Request(ezmaxcustomerPatchObjectV1Request).Execute()

Patch an existing Ezmaxcustomer



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
	pkiEzmaxcustomerID := int32(56) // int32 | The unique ID of the Ezmaxcustomer
	ezmaxcustomerPatchObjectV1Request := *openapiclient.NewEzmaxcustomerPatchObjectV1Request(*openapiclient.NewEzmaxcustomerRequestPatch()) // EzmaxcustomerPatchObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzmaxcustomerAPI.EzmaxcustomerPatchObjectV1(context.Background(), pkiEzmaxcustomerID).EzmaxcustomerPatchObjectV1Request(ezmaxcustomerPatchObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzmaxcustomerAPI.EzmaxcustomerPatchObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxcustomerPatchObjectV1`: EzmaxcustomerPatchObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzmaxcustomerAPI.EzmaxcustomerPatchObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzmaxcustomerID** | **int32** | The unique ID of the Ezmaxcustomer | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxcustomerPatchObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezmaxcustomerPatchObjectV1Request** | [**EzmaxcustomerPatchObjectV1Request**](EzmaxcustomerPatchObjectV1Request.md) |  | 

### Return type

[**EzmaxcustomerPatchObjectV1Response**](EzmaxcustomerPatchObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

