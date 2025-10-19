# eZmaxAPI\ObjectEzmaxcustomeruserAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzmaxcustomeruserPatchObjectV1**](ObjectEzmaxcustomeruserAPI.md#EzmaxcustomeruserPatchObjectV1) | **Patch** /1/object/ezmaxcustomeruser/{pkiEzmaxcustomeruserID} | Patch an existing Ezmaxcustomeruser



## EzmaxcustomeruserPatchObjectV1

> EzmaxcustomeruserPatchObjectV1Response EzmaxcustomeruserPatchObjectV1(ctx, pkiEzmaxcustomeruserID).EzmaxcustomeruserPatchObjectV1Request(ezmaxcustomeruserPatchObjectV1Request).Execute()

Patch an existing Ezmaxcustomeruser



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
	pkiEzmaxcustomeruserID := int32(56) // int32 | The unique ID of the Ezmaxcustomeruser
	ezmaxcustomeruserPatchObjectV1Request := *openapiclient.NewEzmaxcustomeruserPatchObjectV1Request(*openapiclient.NewEzmaxcustomeruserRequestPatch()) // EzmaxcustomeruserPatchObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzmaxcustomeruserAPI.EzmaxcustomeruserPatchObjectV1(context.Background(), pkiEzmaxcustomeruserID).EzmaxcustomeruserPatchObjectV1Request(ezmaxcustomeruserPatchObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzmaxcustomeruserAPI.EzmaxcustomeruserPatchObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxcustomeruserPatchObjectV1`: EzmaxcustomeruserPatchObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzmaxcustomeruserAPI.EzmaxcustomeruserPatchObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzmaxcustomeruserID** | **int32** | The unique ID of the Ezmaxcustomeruser | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxcustomeruserPatchObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezmaxcustomeruserPatchObjectV1Request** | [**EzmaxcustomeruserPatchObjectV1Request**](EzmaxcustomeruserPatchObjectV1Request.md) |  | 

### Return type

[**EzmaxcustomeruserPatchObjectV1Response**](EzmaxcustomeruserPatchObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

