# eZmaxAPI\ObjectInscriptionnotauthenticatedAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InscriptionnotauthenticatedGetCommunicationListV1**](ObjectInscriptionnotauthenticatedAPI.md#InscriptionnotauthenticatedGetCommunicationListV1) | **Get** /1/object/inscriptionnotauthenticated/{pkiInscriptionnotauthenticatedID}/getCommunicationList | Retrieve Communication list



## InscriptionnotauthenticatedGetCommunicationListV1

> InscriptionnotauthenticatedGetCommunicationListV1Response InscriptionnotauthenticatedGetCommunicationListV1(ctx, pkiInscriptionnotauthenticatedID).Execute()

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
    pkiInscriptionnotauthenticatedID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetCommunicationListV1(context.Background(), pkiInscriptionnotauthenticatedID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetCommunicationListV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InscriptionnotauthenticatedGetCommunicationListV1`: InscriptionnotauthenticatedGetCommunicationListV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionnotauthenticatedAPI.InscriptionnotauthenticatedGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionnotauthenticatedID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionnotauthenticatedGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptionnotauthenticatedGetCommunicationListV1Response**](InscriptionnotauthenticatedGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

