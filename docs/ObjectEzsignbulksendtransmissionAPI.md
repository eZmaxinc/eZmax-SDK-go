# eZmaxAPI\ObjectEzsignbulksendtransmissionAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignbulksendtransmissionGetCsvErrorsV1**](ObjectEzsignbulksendtransmissionAPI.md#EzsignbulksendtransmissionGetCsvErrorsV1) | **Get** /1/object/ezsignbulksendtransmission/{pkiEzsignbulksendtransmissionID}/getCsvErrors | Retrieve an existing Ezsignbulksendtransmission&#39;s Csv containing errors
[**EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1**](ObjectEzsignbulksendtransmissionAPI.md#EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1) | **Get** /1/object/ezsignbulksendtransmission/{pkiEzsignbulksendtransmissionID}/getEzsignsignaturesAutomatic | Retrieve an existing Ezsignbulksendtransmission&#39;s automatic Ezsignsignatures
[**EzsignbulksendtransmissionGetFormsDataV1**](ObjectEzsignbulksendtransmissionAPI.md#EzsignbulksendtransmissionGetFormsDataV1) | **Get** /1/object/ezsignbulksendtransmission/{pkiEzsignbulksendtransmissionID}/getFormsData | Retrieve an existing Ezsignbulksendtransmission&#39;s forms data
[**EzsignbulksendtransmissionGetObjectV2**](ObjectEzsignbulksendtransmissionAPI.md#EzsignbulksendtransmissionGetObjectV2) | **Get** /2/object/ezsignbulksendtransmission/{pkiEzsignbulksendtransmissionID} | Retrieve an existing Ezsignbulksendtransmission



## EzsignbulksendtransmissionGetCsvErrorsV1

> string EzsignbulksendtransmissionGetCsvErrorsV1(ctx, pkiEzsignbulksendtransmissionID).Execute()

Retrieve an existing Ezsignbulksendtransmission's Csv containing errors



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
    pkiEzsignbulksendtransmissionID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignbulksendtransmissionAPI.EzsignbulksendtransmissionGetCsvErrorsV1(context.Background(), pkiEzsignbulksendtransmissionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendtransmissionAPI.EzsignbulksendtransmissionGetCsvErrorsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignbulksendtransmissionGetCsvErrorsV1`: string
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendtransmissionAPI.EzsignbulksendtransmissionGetCsvErrorsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendtransmissionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendtransmissionGetCsvErrorsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1

> EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1(ctx, pkiEzsignbulksendtransmissionID).Execute()

Retrieve an existing Ezsignbulksendtransmission's automatic Ezsignsignatures



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
    pkiEzsignbulksendtransmissionID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignbulksendtransmissionAPI.EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1(context.Background(), pkiEzsignbulksendtransmissionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendtransmissionAPI.EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1`: EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendtransmissionAPI.EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendtransmissionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response**](EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendtransmissionGetFormsDataV1

> EzsignbulksendtransmissionGetFormsDataV1Response EzsignbulksendtransmissionGetFormsDataV1(ctx, pkiEzsignbulksendtransmissionID).Execute()

Retrieve an existing Ezsignbulksendtransmission's forms data



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
    pkiEzsignbulksendtransmissionID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignbulksendtransmissionAPI.EzsignbulksendtransmissionGetFormsDataV1(context.Background(), pkiEzsignbulksendtransmissionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendtransmissionAPI.EzsignbulksendtransmissionGetFormsDataV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignbulksendtransmissionGetFormsDataV1`: EzsignbulksendtransmissionGetFormsDataV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendtransmissionAPI.EzsignbulksendtransmissionGetFormsDataV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendtransmissionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendtransmissionGetFormsDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignbulksendtransmissionGetFormsDataV1Response**](EzsignbulksendtransmissionGetFormsDataV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendtransmissionGetObjectV2

> EzsignbulksendtransmissionGetObjectV2Response EzsignbulksendtransmissionGetObjectV2(ctx, pkiEzsignbulksendtransmissionID).Execute()

Retrieve an existing Ezsignbulksendtransmission



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
    pkiEzsignbulksendtransmissionID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignbulksendtransmissionAPI.EzsignbulksendtransmissionGetObjectV2(context.Background(), pkiEzsignbulksendtransmissionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendtransmissionAPI.EzsignbulksendtransmissionGetObjectV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignbulksendtransmissionGetObjectV2`: EzsignbulksendtransmissionGetObjectV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendtransmissionAPI.EzsignbulksendtransmissionGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendtransmissionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendtransmissionGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignbulksendtransmissionGetObjectV2Response**](EzsignbulksendtransmissionGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

