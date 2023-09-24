# eZmaxAPI\ObjectEzsigntemplatepackagesignermembershipAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigntemplatepackagesignermembershipCreateObjectV1**](ObjectEzsigntemplatepackagesignermembershipAPI.md#EzsigntemplatepackagesignermembershipCreateObjectV1) | **Post** /1/object/ezsigntemplatepackagesignermembership | Create a new Ezsigntemplatepackagesignermembership
[**EzsigntemplatepackagesignermembershipDeleteObjectV1**](ObjectEzsigntemplatepackagesignermembershipAPI.md#EzsigntemplatepackagesignermembershipDeleteObjectV1) | **Delete** /1/object/ezsigntemplatepackagesignermembership/{pkiEzsigntemplatepackagesignermembershipID} | Delete an existing Ezsigntemplatepackagesignermembership
[**EzsigntemplatepackagesignermembershipGetObjectV2**](ObjectEzsigntemplatepackagesignermembershipAPI.md#EzsigntemplatepackagesignermembershipGetObjectV2) | **Get** /2/object/ezsigntemplatepackagesignermembership/{pkiEzsigntemplatepackagesignermembershipID} | Retrieve an existing Ezsigntemplatepackagesignermembership



## EzsigntemplatepackagesignermembershipCreateObjectV1

> EzsigntemplatepackagesignermembershipCreateObjectV1Response EzsigntemplatepackagesignermembershipCreateObjectV1(ctx).EzsigntemplatepackagesignermembershipCreateObjectV1Request(ezsigntemplatepackagesignermembershipCreateObjectV1Request).Execute()

Create a new Ezsigntemplatepackagesignermembership



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
    ezsigntemplatepackagesignermembershipCreateObjectV1Request := *openapiclient.NewEzsigntemplatepackagesignermembershipCreateObjectV1Request([]openapiclient.EzsigntemplatepackagesignermembershipRequestCompound{*openapiclient.NewEzsigntemplatepackagesignermembershipRequestCompound(int32(194), int32(174), int32(9))}) // EzsigntemplatepackagesignermembershipCreateObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigntemplatepackagesignermembershipAPI.EzsigntemplatepackagesignermembershipCreateObjectV1(context.Background()).EzsigntemplatepackagesignermembershipCreateObjectV1Request(ezsigntemplatepackagesignermembershipCreateObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackagesignermembershipAPI.EzsigntemplatepackagesignermembershipCreateObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigntemplatepackagesignermembershipCreateObjectV1`: EzsigntemplatepackagesignermembershipCreateObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackagesignermembershipAPI.EzsigntemplatepackagesignermembershipCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackagesignermembershipCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplatepackagesignermembershipCreateObjectV1Request** | [**EzsigntemplatepackagesignermembershipCreateObjectV1Request**](EzsigntemplatepackagesignermembershipCreateObjectV1Request.md) |  | 

### Return type

[**EzsigntemplatepackagesignermembershipCreateObjectV1Response**](EzsigntemplatepackagesignermembershipCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepackagesignermembershipDeleteObjectV1

> EzsigntemplatepackagesignermembershipDeleteObjectV1Response EzsigntemplatepackagesignermembershipDeleteObjectV1(ctx, pkiEzsigntemplatepackagesignermembershipID).Execute()

Delete an existing Ezsigntemplatepackagesignermembership



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
    pkiEzsigntemplatepackagesignermembershipID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigntemplatepackagesignermembershipAPI.EzsigntemplatepackagesignermembershipDeleteObjectV1(context.Background(), pkiEzsigntemplatepackagesignermembershipID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackagesignermembershipAPI.EzsigntemplatepackagesignermembershipDeleteObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigntemplatepackagesignermembershipDeleteObjectV1`: EzsigntemplatepackagesignermembershipDeleteObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackagesignermembershipAPI.EzsigntemplatepackagesignermembershipDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepackagesignermembershipID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackagesignermembershipDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatepackagesignermembershipDeleteObjectV1Response**](EzsigntemplatepackagesignermembershipDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepackagesignermembershipGetObjectV2

> EzsigntemplatepackagesignermembershipGetObjectV2Response EzsigntemplatepackagesignermembershipGetObjectV2(ctx, pkiEzsigntemplatepackagesignermembershipID).Execute()

Retrieve an existing Ezsigntemplatepackagesignermembership



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
    pkiEzsigntemplatepackagesignermembershipID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigntemplatepackagesignermembershipAPI.EzsigntemplatepackagesignermembershipGetObjectV2(context.Background(), pkiEzsigntemplatepackagesignermembershipID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackagesignermembershipAPI.EzsigntemplatepackagesignermembershipGetObjectV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigntemplatepackagesignermembershipGetObjectV2`: EzsigntemplatepackagesignermembershipGetObjectV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackagesignermembershipAPI.EzsigntemplatepackagesignermembershipGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepackagesignermembershipID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackagesignermembershipGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatepackagesignermembershipGetObjectV2Response**](EzsigntemplatepackagesignermembershipGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

