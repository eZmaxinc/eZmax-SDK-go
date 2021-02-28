# eZmaxAPI\GlobalCustomerApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalCustomerGetEndpointV1**](GlobalCustomerApi.md#GlobalCustomerGetEndpointV1) | **Get** /1/customer/{pksCustomerCode}/endpoint | Get customer endpoint



## GlobalCustomerGetEndpointV1

> GlobalCustomerGetEndpointV1Response GlobalCustomerGetEndpointV1(ctx, pksCustomerCode).SInfrastructureproductCode(sInfrastructureproductCode).Execute()

Get customer endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pksCustomerCode := "pksCustomerCode_example" // string | The customer code assigned to your account
    sInfrastructureproductCode := "sInfrastructureproductCode_example" // string | The infrastructure product Code  If undefined, \"appcluster01\" is assumed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GlobalCustomerApi.GlobalCustomerGetEndpointV1(context.Background(), pksCustomerCode).SInfrastructureproductCode(sInfrastructureproductCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalCustomerApi.GlobalCustomerGetEndpointV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GlobalCustomerGetEndpointV1`: GlobalCustomerGetEndpointV1Response
    fmt.Fprintf(os.Stdout, "Response from `GlobalCustomerApi.GlobalCustomerGetEndpointV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pksCustomerCode** | **string** | The customer code assigned to your account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalCustomerGetEndpointV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sInfrastructureproductCode** | **string** | The infrastructure product Code  If undefined, \&quot;appcluster01\&quot; is assumed | 

### Return type

[**GlobalCustomerGetEndpointV1Response**](GlobalCustomerGetEndpointV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

