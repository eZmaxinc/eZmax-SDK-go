# eZmaxAPI\ObjectEzsignbulksendsignermappingAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignbulksendsignermappingCreateObjectV1**](ObjectEzsignbulksendsignermappingAPI.md#EzsignbulksendsignermappingCreateObjectV1) | **Post** /1/object/ezsignbulksendsignermapping | Create a new Ezsignbulksendsignermapping
[**EzsignbulksendsignermappingDeleteObjectV1**](ObjectEzsignbulksendsignermappingAPI.md#EzsignbulksendsignermappingDeleteObjectV1) | **Delete** /1/object/ezsignbulksendsignermapping/{pkiEzsignbulksendsignermappingID} | Delete an existing Ezsignbulksendsignermapping
[**EzsignbulksendsignermappingGetObjectV2**](ObjectEzsignbulksendsignermappingAPI.md#EzsignbulksendsignermappingGetObjectV2) | **Get** /2/object/ezsignbulksendsignermapping/{pkiEzsignbulksendsignermappingID} | Retrieve an existing Ezsignbulksendsignermapping



## EzsignbulksendsignermappingCreateObjectV1

> EzsignbulksendsignermappingCreateObjectV1Response EzsignbulksendsignermappingCreateObjectV1(ctx).EzsignbulksendsignermappingCreateObjectV1Request(ezsignbulksendsignermappingCreateObjectV1Request).Execute()

Create a new Ezsignbulksendsignermapping



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
    ezsignbulksendsignermappingCreateObjectV1Request := *openapiclient.NewEzsignbulksendsignermappingCreateObjectV1Request([]openapiclient.EzsignbulksendsignermappingRequestCompound{*openapiclient.NewEzsignbulksendsignermappingRequestCompound(int32(8), "Supervisor")}) // EzsignbulksendsignermappingCreateObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignbulksendsignermappingAPI.EzsignbulksendsignermappingCreateObjectV1(context.Background()).EzsignbulksendsignermappingCreateObjectV1Request(ezsignbulksendsignermappingCreateObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendsignermappingAPI.EzsignbulksendsignermappingCreateObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignbulksendsignermappingCreateObjectV1`: EzsignbulksendsignermappingCreateObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendsignermappingAPI.EzsignbulksendsignermappingCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendsignermappingCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignbulksendsignermappingCreateObjectV1Request** | [**EzsignbulksendsignermappingCreateObjectV1Request**](EzsignbulksendsignermappingCreateObjectV1Request.md) |  | 

### Return type

[**EzsignbulksendsignermappingCreateObjectV1Response**](EzsignbulksendsignermappingCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendsignermappingDeleteObjectV1

> EzsignbulksendsignermappingDeleteObjectV1Response EzsignbulksendsignermappingDeleteObjectV1(ctx, pkiEzsignbulksendsignermappingID).Execute()

Delete an existing Ezsignbulksendsignermapping



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
    pkiEzsignbulksendsignermappingID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignbulksendsignermappingAPI.EzsignbulksendsignermappingDeleteObjectV1(context.Background(), pkiEzsignbulksendsignermappingID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendsignermappingAPI.EzsignbulksendsignermappingDeleteObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignbulksendsignermappingDeleteObjectV1`: EzsignbulksendsignermappingDeleteObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendsignermappingAPI.EzsignbulksendsignermappingDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendsignermappingID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendsignermappingDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignbulksendsignermappingDeleteObjectV1Response**](EzsignbulksendsignermappingDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendsignermappingGetObjectV2

> EzsignbulksendsignermappingGetObjectV2Response EzsignbulksendsignermappingGetObjectV2(ctx, pkiEzsignbulksendsignermappingID).Execute()

Retrieve an existing Ezsignbulksendsignermapping



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
    pkiEzsignbulksendsignermappingID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignbulksendsignermappingAPI.EzsignbulksendsignermappingGetObjectV2(context.Background(), pkiEzsignbulksendsignermappingID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendsignermappingAPI.EzsignbulksendsignermappingGetObjectV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignbulksendsignermappingGetObjectV2`: EzsignbulksendsignermappingGetObjectV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendsignermappingAPI.EzsignbulksendsignermappingGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendsignermappingID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendsignermappingGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignbulksendsignermappingGetObjectV2Response**](EzsignbulksendsignermappingGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

