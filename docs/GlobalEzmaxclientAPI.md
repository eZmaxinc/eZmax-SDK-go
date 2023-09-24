# eZmaxAPI\GlobalEzmaxclientAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalEzmaxclientVersionV1**](GlobalEzmaxclientAPI.md#GlobalEzmaxclientVersionV1) | **Get** /1/ezmaxclient/{pksEzmaxclientOs}/version | Retrieve the latest version of the Ezmaxclient



## GlobalEzmaxclientVersionV1

> GlobalEzmaxclientVersionV1Response GlobalEzmaxclientVersionV1(ctx, pksEzmaxclientOs).Execute()

Retrieve the latest version of the Ezmaxclient



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
    pksEzmaxclientOs := openapiclient.Field-pksEzmaxclientOs("iOS") // FieldPksEzmaxclientOs | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalEzmaxclientAPI.GlobalEzmaxclientVersionV1(context.Background(), pksEzmaxclientOs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalEzmaxclientAPI.GlobalEzmaxclientVersionV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GlobalEzmaxclientVersionV1`: GlobalEzmaxclientVersionV1Response
    fmt.Fprintf(os.Stdout, "Response from `GlobalEzmaxclientAPI.GlobalEzmaxclientVersionV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pksEzmaxclientOs** | [**FieldPksEzmaxclientOs**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalEzmaxclientVersionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GlobalEzmaxclientVersionV1Response**](GlobalEzmaxclientVersionV1Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

