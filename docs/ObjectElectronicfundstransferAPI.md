# eZmaxAPI\ObjectElectronicfundstransferAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ElectronicfundstransferGetCommunicationListV1**](ObjectElectronicfundstransferAPI.md#ElectronicfundstransferGetCommunicationListV1) | **Get** /1/object/electronicfundstransfer/{pkiElectronicfundstransferID}/getCommunicationList | Retrieve Communication list



## ElectronicfundstransferGetCommunicationListV1

> ElectronicfundstransferGetCommunicationListV1Response ElectronicfundstransferGetCommunicationListV1(ctx, pkiElectronicfundstransferID).Execute()

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
	pkiElectronicfundstransferID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectElectronicfundstransferAPI.ElectronicfundstransferGetCommunicationListV1(context.Background(), pkiElectronicfundstransferID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectElectronicfundstransferAPI.ElectronicfundstransferGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ElectronicfundstransferGetCommunicationListV1`: ElectronicfundstransferGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectElectronicfundstransferAPI.ElectronicfundstransferGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiElectronicfundstransferID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiElectronicfundstransferGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ElectronicfundstransferGetCommunicationListV1Response**](ElectronicfundstransferGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

