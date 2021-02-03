# eZmaxAPI\ModuleUserApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserCreateEzsignuserV1**](ModuleUserApi.md#UserCreateEzsignuserV1) | **Post** /1/module/user/createezsignuser | Create a new User of type Ezsignuser



## UserCreateEzsignuserV1

> UserCreateEzsignuserV1Response UserCreateEzsignuserV1(ctx).UserCreateEzsignuserV1Request(userCreateEzsignuserV1Request).Execute()

Create a new User of type Ezsignuser



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
    userCreateEzsignuserV1Request := []openapiclient.UserCreateEzsignuserV1Request{*openapiclient.NewUserCreateEzsignuserV1Request(int32(2), "John", "Doe", "example@domain.com", "514", "990", "1516")} // []UserCreateEzsignuserV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ModuleUserApi.UserCreateEzsignuserV1(context.Background()).UserCreateEzsignuserV1Request(userCreateEzsignuserV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModuleUserApi.UserCreateEzsignuserV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserCreateEzsignuserV1`: UserCreateEzsignuserV1Response
    fmt.Fprintf(os.Stdout, "Response from `ModuleUserApi.UserCreateEzsignuserV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserCreateEzsignuserV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userCreateEzsignuserV1Request** | [**[]UserCreateEzsignuserV1Request**](user-createEzsignuser-v1-Request.md) |  | 

### Return type

[**UserCreateEzsignuserV1Response**](user-createEzsignuser-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

