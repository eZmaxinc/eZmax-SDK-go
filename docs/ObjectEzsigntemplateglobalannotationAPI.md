# eZmaxAPI\ObjectEzsigntemplateglobalannotationAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigntemplateglobalannotationGetObjectV2**](ObjectEzsigntemplateglobalannotationAPI.md#EzsigntemplateglobalannotationGetObjectV2) | **Get** /2/object/ezsigntemplateglobalannotation/{pkiEzsigntemplateglobalannotationID} | Retrieve an existing Ezsigntemplateglobalannotation



## EzsigntemplateglobalannotationGetObjectV2

> EzsigntemplateglobalannotationGetObjectV2Response EzsigntemplateglobalannotationGetObjectV2(ctx, pkiEzsigntemplateglobalannotationID).Execute()

Retrieve an existing Ezsigntemplateglobalannotation



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
	pkiEzsigntemplateglobalannotationID := int32(56) // int32 | The unique ID of the Ezsigntemplateglobalannotation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplateglobalannotationAPI.EzsigntemplateglobalannotationGetObjectV2(context.Background(), pkiEzsigntemplateglobalannotationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateglobalannotationAPI.EzsigntemplateglobalannotationGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateglobalannotationGetObjectV2`: EzsigntemplateglobalannotationGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateglobalannotationAPI.EzsigntemplateglobalannotationGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplateglobalannotationID** | **int32** | The unique ID of the Ezsigntemplateglobalannotation | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateglobalannotationGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplateglobalannotationGetObjectV2Response**](EzsigntemplateglobalannotationGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

