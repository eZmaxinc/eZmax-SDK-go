# eZmaxAPI\ObjectEzsignpageAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignpageConsultV1**](ObjectEzsignpageAPI.md#EzsignpageConsultV1) | **Post** /1/object/ezsignpage/{pkiEzsignpageID}/consult | Consult an Ezsignpage



## EzsignpageConsultV1

> CommonResponse EzsignpageConsultV1(ctx, pkiEzsignpageID).Body(body).Execute()

Consult an Ezsignpage

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
	pkiEzsignpageID := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignpageAPI.EzsignpageConsultV1(context.Background(), pkiEzsignpageID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignpageAPI.EzsignpageConsultV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignpageConsultV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignpageAPI.EzsignpageConsultV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignpageID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignpageConsultV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

