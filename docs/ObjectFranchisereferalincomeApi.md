# eZmaxAPI\ObjectFranchisereferalincomeApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FranchisereferalincomeCreateObjectV1**](ObjectFranchisereferalincomeApi.md#FranchisereferalincomeCreateObjectV1) | **Post** /1/object/franchisereferalincome | Create a new Franchisereferalincome



## FranchisereferalincomeCreateObjectV1

> FranchisereferalincomeCreateObjectV1Response FranchisereferalincomeCreateObjectV1(ctx).FranchisereferalincomeCreateObjectV1Request(franchisereferalincomeCreateObjectV1Request).Execute()

Create a new Franchisereferalincome



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
    franchisereferalincomeCreateObjectV1Request := []openapiclient.FranchisereferalincomeCreateObjectV1Request{*openapiclient.NewFranchisereferalincomeCreateObjectV1Request()} // []FranchisereferalincomeCreateObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectFranchisereferalincomeApi.FranchisereferalincomeCreateObjectV1(context.Background()).FranchisereferalincomeCreateObjectV1Request(franchisereferalincomeCreateObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectFranchisereferalincomeApi.FranchisereferalincomeCreateObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FranchisereferalincomeCreateObjectV1`: FranchisereferalincomeCreateObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectFranchisereferalincomeApi.FranchisereferalincomeCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFranchisereferalincomeCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **franchisereferalincomeCreateObjectV1Request** | [**[]FranchisereferalincomeCreateObjectV1Request**](franchisereferalincome-createObject-v1-Request.md) |  | 

### Return type

[**FranchisereferalincomeCreateObjectV1Response**](franchisereferalincome-createObject-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

