# eZmaxAPI\ObjectInscriptiontempAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InscriptiontempGetCommunicationListV1**](ObjectInscriptiontempAPI.md#InscriptiontempGetCommunicationListV1) | **Get** /1/object/inscriptiontemp/{pkiInscriptiontempID}/getCommunicationList | Retrieve Communication list



## InscriptiontempGetCommunicationListV1

> InscriptiontempGetCommunicationListV1Response InscriptiontempGetCommunicationListV1(ctx, pkiInscriptiontempID).Execute()

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
    pkiInscriptiontempID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectInscriptiontempAPI.InscriptiontempGetCommunicationListV1(context.Background(), pkiInscriptiontempID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptiontempAPI.InscriptiontempGetCommunicationListV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InscriptiontempGetCommunicationListV1`: InscriptiontempGetCommunicationListV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptiontempAPI.InscriptiontempGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptiontempID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptiontempGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptiontempGetCommunicationListV1Response**](InscriptiontempGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

