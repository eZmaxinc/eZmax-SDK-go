# eZmaxAPI\ObjectPermissionAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermissionCreateObjectV1**](ObjectPermissionAPI.md#PermissionCreateObjectV1) | **Post** /1/object/permission | Create a new Permission
[**PermissionDeleteObjectV1**](ObjectPermissionAPI.md#PermissionDeleteObjectV1) | **Delete** /1/object/permission/{pkiPermissionID} | Delete an existing Permission
[**PermissionEditObjectV1**](ObjectPermissionAPI.md#PermissionEditObjectV1) | **Put** /1/object/permission/{pkiPermissionID} | Edit an existing Permission
[**PermissionGetObjectV2**](ObjectPermissionAPI.md#PermissionGetObjectV2) | **Get** /2/object/permission/{pkiPermissionID} | Retrieve an existing Permission



## PermissionCreateObjectV1

> PermissionCreateObjectV1Response PermissionCreateObjectV1(ctx).PermissionCreateObjectV1Request(permissionCreateObjectV1Request).Execute()

Create a new Permission



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
	permissionCreateObjectV1Request := *openapiclient.NewPermissionCreateObjectV1Request([]openapiclient.PermissionRequestCompound{*openapiclient.NewPermissionRequest(int32(53))}) // PermissionCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPermissionAPI.PermissionCreateObjectV1(context.Background()).PermissionCreateObjectV1Request(permissionCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPermissionAPI.PermissionCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionCreateObjectV1`: PermissionCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPermissionAPI.PermissionCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissionCreateObjectV1Request** | [**PermissionCreateObjectV1Request**](PermissionCreateObjectV1Request.md) |  | 

### Return type

[**PermissionCreateObjectV1Response**](PermissionCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionDeleteObjectV1

> CommonResponse PermissionDeleteObjectV1(ctx, pkiPermissionID).Execute()

Delete an existing Permission



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
	pkiPermissionID := int32(56) // int32 | The unique ID of the Permission

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPermissionAPI.PermissionDeleteObjectV1(context.Background(), pkiPermissionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPermissionAPI.PermissionDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionDeleteObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectPermissionAPI.PermissionDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPermissionID** | **int32** | The unique ID of the Permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionEditObjectV1

> CommonResponse PermissionEditObjectV1(ctx, pkiPermissionID).PermissionEditObjectV1Request(permissionEditObjectV1Request).Execute()

Edit an existing Permission



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
	pkiPermissionID := int32(56) // int32 | The unique ID of the Permission
	permissionEditObjectV1Request := *openapiclient.NewPermissionEditObjectV1Request(*openapiclient.NewPermissionRequest(int32(53))) // PermissionEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPermissionAPI.PermissionEditObjectV1(context.Background(), pkiPermissionID).PermissionEditObjectV1Request(permissionEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPermissionAPI.PermissionEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectPermissionAPI.PermissionEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPermissionID** | **int32** | The unique ID of the Permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permissionEditObjectV1Request** | [**PermissionEditObjectV1Request**](PermissionEditObjectV1Request.md) |  | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionGetObjectV2

> PermissionGetObjectV2Response PermissionGetObjectV2(ctx, pkiPermissionID).Execute()

Retrieve an existing Permission



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
	pkiPermissionID := int32(56) // int32 | The unique ID of the Permission

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPermissionAPI.PermissionGetObjectV2(context.Background(), pkiPermissionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPermissionAPI.PermissionGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionGetObjectV2`: PermissionGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPermissionAPI.PermissionGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPermissionID** | **int32** | The unique ID of the Permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermissionGetObjectV2Response**](PermissionGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

