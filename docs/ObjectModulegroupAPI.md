# eZmaxAPI\ObjectModulegroupAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModulegroupGetAllV1**](ObjectModulegroupAPI.md#ModulegroupGetAllV1) | **Get** /1/object/modulegroup/getAll/{eContext} | Retrieve all Modulegroups



## ModulegroupGetAllV1

> ModulegroupGetAllV1Response ModulegroupGetAllV1(ctx, eContext).Execute()

Retrieve all Modulegroups

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
	eContext := "eContext_example" // string | The context of the Modulegroup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectModulegroupAPI.ModulegroupGetAllV1(context.Background(), eContext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectModulegroupAPI.ModulegroupGetAllV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModulegroupGetAllV1`: ModulegroupGetAllV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectModulegroupAPI.ModulegroupGetAllV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eContext** | **string** | The context of the Modulegroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiModulegroupGetAllV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModulegroupGetAllV1Response**](ModulegroupGetAllV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

