# eZmaxAPI\ObjectApikeyApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApikeyCreateObjectV1**](ObjectApikeyApi.md#ApikeyCreateObjectV1) | **Post** /1/object/apikey | Create a new Apikey



## ApikeyCreateObjectV1

> ApikeyCreateObjectV1Response ApikeyCreateObjectV1(ctx).ApikeyCreateObjectV1Request(apikeyCreateObjectV1Request).Execute()

Create a new Apikey



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
    apikeyCreateObjectV1Request := []openapiclient.ApikeyCreateObjectV1Request{*openapiclient.NewApikeyCreateObjectV1Request()} // []ApikeyCreateObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectApikeyApi.ApikeyCreateObjectV1(context.Background()).ApikeyCreateObjectV1Request(apikeyCreateObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectApikeyApi.ApikeyCreateObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApikeyCreateObjectV1`: ApikeyCreateObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectApikeyApi.ApikeyCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApikeyCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apikeyCreateObjectV1Request** | [**[]ApikeyCreateObjectV1Request**](apikey-createObject-v1-Request.md) |  | 

### Return type

[**ApikeyCreateObjectV1Response**](apikey-createObject-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

