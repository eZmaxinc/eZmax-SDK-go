# eZmaxAPI\ObjectActivesessionApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivesessionGetCurrentV1**](ObjectActivesessionApi.md#ActivesessionGetCurrentV1) | **Get** /1/object/activesession/getCurrent | Get Current Activesession



## ActivesessionGetCurrentV1

> ActivesessionGetCurrentV1Response ActivesessionGetCurrentV1(ctx).Execute()

Get Current Activesession



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectActivesessionApi.ActivesessionGetCurrentV1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectActivesessionApi.ActivesessionGetCurrentV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivesessionGetCurrentV1`: ActivesessionGetCurrentV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectActivesessionApi.ActivesessionGetCurrentV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActivesessionGetCurrentV1Request struct via the builder pattern


### Return type

[**ActivesessionGetCurrentV1Response**](ActivesessionGetCurrentV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

