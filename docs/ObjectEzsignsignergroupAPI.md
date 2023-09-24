# eZmaxAPI\ObjectEzsignsignergroupAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignsignergroupCreateObjectV1**](ObjectEzsignsignergroupAPI.md#EzsignsignergroupCreateObjectV1) | **Post** /1/object/ezsignsignergroup | Create a new Ezsignsignergroup
[**EzsignsignergroupDeleteObjectV1**](ObjectEzsignsignergroupAPI.md#EzsignsignergroupDeleteObjectV1) | **Delete** /1/object/ezsignsignergroup/{pkiEzsignsignergroupID} | Delete an existing Ezsignsignergroup
[**EzsignsignergroupEditEzsignsignergroupmembershipsV1**](ObjectEzsignsignergroupAPI.md#EzsignsignergroupEditEzsignsignergroupmembershipsV1) | **Put** /1/object/ezsignsignergroup/{pkiEzsignsignergroupID}/editEzsignsignergroupmemberships | Edit multiple Ezsignsignergroupmemberships
[**EzsignsignergroupEditObjectV1**](ObjectEzsignsignergroupAPI.md#EzsignsignergroupEditObjectV1) | **Put** /1/object/ezsignsignergroup/{pkiEzsignsignergroupID} | Edit an existing Ezsignsignergroup
[**EzsignsignergroupGetEzsignsignergroupmembershipsV1**](ObjectEzsignsignergroupAPI.md#EzsignsignergroupGetEzsignsignergroupmembershipsV1) | **Get** /1/object/ezsignsignergroup/{pkiEzsignsignergroupID}/getEzsignsignergroupmemberships | Retrieve an existing Ezsignsignergroup&#39;s Ezsignsignergroupmemberships
[**EzsignsignergroupGetObjectV2**](ObjectEzsignsignergroupAPI.md#EzsignsignergroupGetObjectV2) | **Get** /2/object/ezsignsignergroup/{pkiEzsignsignergroupID} | Retrieve an existing Ezsignsignergroup



## EzsignsignergroupCreateObjectV1

> EzsignsignergroupCreateObjectV1Response EzsignsignergroupCreateObjectV1(ctx).EzsignsignergroupCreateObjectV1Request(ezsignsignergroupCreateObjectV1Request).Execute()

Create a new Ezsignsignergroup



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
    ezsignsignergroupCreateObjectV1Request := *openapiclient.NewEzsignsignergroupCreateObjectV1Request([]openapiclient.EzsignsignergroupRequestCompound{*openapiclient.NewEzsignsignergroupRequestCompound(int32(33), *openapiclient.NewMultilingualEzsignsignergroupDescription())}) // EzsignsignergroupCreateObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignsignergroupAPI.EzsignsignergroupCreateObjectV1(context.Background()).EzsignsignergroupCreateObjectV1Request(ezsignsignergroupCreateObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignergroupAPI.EzsignsignergroupCreateObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignsignergroupCreateObjectV1`: EzsignsignergroupCreateObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignergroupAPI.EzsignsignergroupCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignergroupCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignsignergroupCreateObjectV1Request** | [**EzsignsignergroupCreateObjectV1Request**](EzsignsignergroupCreateObjectV1Request.md) |  | 

### Return type

[**EzsignsignergroupCreateObjectV1Response**](EzsignsignergroupCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignergroupDeleteObjectV1

> EzsignsignergroupDeleteObjectV1Response EzsignsignergroupDeleteObjectV1(ctx, pkiEzsignsignergroupID).Execute()

Delete an existing Ezsignsignergroup



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
    pkiEzsignsignergroupID := int32(56) // int32 | The unique ID of the Ezsignsignergroup

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignsignergroupAPI.EzsignsignergroupDeleteObjectV1(context.Background(), pkiEzsignsignergroupID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignergroupAPI.EzsignsignergroupDeleteObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignsignergroupDeleteObjectV1`: EzsignsignergroupDeleteObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignergroupAPI.EzsignsignergroupDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignergroupID** | **int32** | The unique ID of the Ezsignsignergroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignergroupDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignsignergroupDeleteObjectV1Response**](EzsignsignergroupDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignergroupEditEzsignsignergroupmembershipsV1

> EzsignsignergroupEditEzsignsignergroupmembershipsV1Response EzsignsignergroupEditEzsignsignergroupmembershipsV1(ctx, pkiEzsignsignergroupID).EzsignsignergroupEditEzsignsignergroupmembershipsV1Request(ezsignsignergroupEditEzsignsignergroupmembershipsV1Request).Execute()

Edit multiple Ezsignsignergroupmemberships



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
    pkiEzsignsignergroupID := int32(56) // int32 | 
    ezsignsignergroupEditEzsignsignergroupmembershipsV1Request := *openapiclient.NewEzsignsignergroupEditEzsignsignergroupmembershipsV1Request([]openapiclient.EzsignsignergroupmembershipRequestCompound{*openapiclient.NewEzsignsignergroupmembershipRequestCompound(int32(27))}) // EzsignsignergroupEditEzsignsignergroupmembershipsV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignsignergroupAPI.EzsignsignergroupEditEzsignsignergroupmembershipsV1(context.Background(), pkiEzsignsignergroupID).EzsignsignergroupEditEzsignsignergroupmembershipsV1Request(ezsignsignergroupEditEzsignsignergroupmembershipsV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignergroupAPI.EzsignsignergroupEditEzsignsignergroupmembershipsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignsignergroupEditEzsignsignergroupmembershipsV1`: EzsignsignergroupEditEzsignsignergroupmembershipsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignergroupAPI.EzsignsignergroupEditEzsignsignergroupmembershipsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignergroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignergroupEditEzsignsignergroupmembershipsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignsignergroupEditEzsignsignergroupmembershipsV1Request** | [**EzsignsignergroupEditEzsignsignergroupmembershipsV1Request**](EzsignsignergroupEditEzsignsignergroupmembershipsV1Request.md) |  | 

### Return type

[**EzsignsignergroupEditEzsignsignergroupmembershipsV1Response**](EzsignsignergroupEditEzsignsignergroupmembershipsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignergroupEditObjectV1

> EzsignsignergroupEditObjectV1Response EzsignsignergroupEditObjectV1(ctx, pkiEzsignsignergroupID).EzsignsignergroupEditObjectV1Request(ezsignsignergroupEditObjectV1Request).Execute()

Edit an existing Ezsignsignergroup



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
    pkiEzsignsignergroupID := int32(56) // int32 | The unique ID of the Ezsignsignergroup
    ezsignsignergroupEditObjectV1Request := *openapiclient.NewEzsignsignergroupEditObjectV1Request(*openapiclient.NewEzsignsignergroupRequestCompound(int32(33), *openapiclient.NewMultilingualEzsignsignergroupDescription())) // EzsignsignergroupEditObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignsignergroupAPI.EzsignsignergroupEditObjectV1(context.Background(), pkiEzsignsignergroupID).EzsignsignergroupEditObjectV1Request(ezsignsignergroupEditObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignergroupAPI.EzsignsignergroupEditObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignsignergroupEditObjectV1`: EzsignsignergroupEditObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignergroupAPI.EzsignsignergroupEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignergroupID** | **int32** | The unique ID of the Ezsignsignergroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignergroupEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignsignergroupEditObjectV1Request** | [**EzsignsignergroupEditObjectV1Request**](EzsignsignergroupEditObjectV1Request.md) |  | 

### Return type

[**EzsignsignergroupEditObjectV1Response**](EzsignsignergroupEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignergroupGetEzsignsignergroupmembershipsV1

> EzsignsignergroupGetEzsignsignergroupmembershipsV1Response EzsignsignergroupGetEzsignsignergroupmembershipsV1(ctx, pkiEzsignsignergroupID).Execute()

Retrieve an existing Ezsignsignergroup's Ezsignsignergroupmemberships

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
    pkiEzsignsignergroupID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignsignergroupAPI.EzsignsignergroupGetEzsignsignergroupmembershipsV1(context.Background(), pkiEzsignsignergroupID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignergroupAPI.EzsignsignergroupGetEzsignsignergroupmembershipsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignsignergroupGetEzsignsignergroupmembershipsV1`: EzsignsignergroupGetEzsignsignergroupmembershipsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignergroupAPI.EzsignsignergroupGetEzsignsignergroupmembershipsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignergroupID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignergroupGetEzsignsignergroupmembershipsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignsignergroupGetEzsignsignergroupmembershipsV1Response**](EzsignsignergroupGetEzsignsignergroupmembershipsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignergroupGetObjectV2

> EzsignsignergroupGetObjectV2Response EzsignsignergroupGetObjectV2(ctx, pkiEzsignsignergroupID).Execute()

Retrieve an existing Ezsignsignergroup



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
    pkiEzsignsignergroupID := int32(56) // int32 | The unique ID of the Ezsignsignergroup

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignsignergroupAPI.EzsignsignergroupGetObjectV2(context.Background(), pkiEzsignsignergroupID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignergroupAPI.EzsignsignergroupGetObjectV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignsignergroupGetObjectV2`: EzsignsignergroupGetObjectV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignergroupAPI.EzsignsignergroupGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignergroupID** | **int32** | The unique ID of the Ezsignsignergroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignergroupGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignsignergroupGetObjectV2Response**](EzsignsignergroupGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

