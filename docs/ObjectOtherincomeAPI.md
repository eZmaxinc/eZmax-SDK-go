# eZmaxAPI\ObjectOtherincomeAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OtherincomeGetCommunicationListV1**](ObjectOtherincomeAPI.md#OtherincomeGetCommunicationListV1) | **Get** /1/object/otherincome/{pkiOtherincomeID}/getCommunicationList | Retrieve Communication list



## OtherincomeGetCommunicationListV1

> OtherincomeGetCommunicationListV1Response OtherincomeGetCommunicationListV1(ctx, pkiOtherincomeID).Execute()

Retrieve Communication list



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
	pkiOtherincomeID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectOtherincomeAPI.OtherincomeGetCommunicationListV1(context.Background(), pkiOtherincomeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectOtherincomeAPI.OtherincomeGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OtherincomeGetCommunicationListV1`: OtherincomeGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectOtherincomeAPI.OtherincomeGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiOtherincomeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOtherincomeGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OtherincomeGetCommunicationListV1Response**](OtherincomeGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

