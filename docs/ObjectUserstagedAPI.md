# eZmaxAPI\ObjectUserstagedAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserstagedCreateUserV1**](ObjectUserstagedAPI.md#UserstagedCreateUserV1) | **Post** /1/object/userstaged/{pkiUserstagedID}/createUser | Create a User from a Userstaged and then map it
[**UserstagedDeleteObjectV1**](ObjectUserstagedAPI.md#UserstagedDeleteObjectV1) | **Delete** /1/object/userstaged/{pkiUserstagedID} | Delete an existing Userstaged
[**UserstagedGetListV1**](ObjectUserstagedAPI.md#UserstagedGetListV1) | **Get** /1/object/userstaged/getList | Retrieve Userstaged list
[**UserstagedGetObjectV2**](ObjectUserstagedAPI.md#UserstagedGetObjectV2) | **Get** /2/object/userstaged/{pkiUserstagedID} | Retrieve an existing Userstaged
[**UserstagedMapV1**](ObjectUserstagedAPI.md#UserstagedMapV1) | **Post** /1/object/userstaged/{pkiUserstagedID}/map | Map the Userstaged to an existing user



## UserstagedCreateUserV1

> UserstagedCreateUserV1Response UserstagedCreateUserV1(ctx, pkiUserstagedID).Body(body).Execute()

Create a User from a Userstaged and then map it



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
	pkiUserstagedID := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUserstagedAPI.UserstagedCreateUserV1(context.Background(), pkiUserstagedID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserstagedAPI.UserstagedCreateUserV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserstagedCreateUserV1`: UserstagedCreateUserV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUserstagedAPI.UserstagedCreateUserV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUserstagedID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserstagedCreateUserV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**UserstagedCreateUserV1Response**](UserstagedCreateUserV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserstagedDeleteObjectV1

> UserstagedDeleteObjectV1Response UserstagedDeleteObjectV1(ctx, pkiUserstagedID).Execute()

Delete an existing Userstaged



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
	pkiUserstagedID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUserstagedAPI.UserstagedDeleteObjectV1(context.Background(), pkiUserstagedID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserstagedAPI.UserstagedDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserstagedDeleteObjectV1`: UserstagedDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUserstagedAPI.UserstagedDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUserstagedID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserstagedDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserstagedDeleteObjectV1Response**](UserstagedDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserstagedGetListV1

> UserstagedGetListV1Response UserstagedGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Userstaged list



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
	eOrderBy := "eOrderBy_example" // string | Specify how you want the results to be sorted (optional)
	iRowMax := int32(56) // int32 |  (optional)
	iRowOffset := int32(56) // int32 |  (optional) (default to 0)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)
	sFilter := "sFilter_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUserstagedAPI.UserstagedGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserstagedAPI.UserstagedGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserstagedGetListV1`: UserstagedGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUserstagedAPI.UserstagedGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserstagedGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**UserstagedGetListV1Response**](UserstagedGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserstagedGetObjectV2

> UserstagedGetObjectV2Response UserstagedGetObjectV2(ctx, pkiUserstagedID).Execute()

Retrieve an existing Userstaged



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
	pkiUserstagedID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUserstagedAPI.UserstagedGetObjectV2(context.Background(), pkiUserstagedID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserstagedAPI.UserstagedGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserstagedGetObjectV2`: UserstagedGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUserstagedAPI.UserstagedGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUserstagedID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserstagedGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserstagedGetObjectV2Response**](UserstagedGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserstagedMapV1

> UserstagedMapV1Response UserstagedMapV1(ctx, pkiUserstagedID).UserstagedMapV1Request(userstagedMapV1Request).Execute()

Map the Userstaged to an existing user



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
	pkiUserstagedID := int32(56) // int32 | 
	userstagedMapV1Request := *openapiclient.NewUserstagedMapV1Request(int32(70)) // UserstagedMapV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUserstagedAPI.UserstagedMapV1(context.Background(), pkiUserstagedID).UserstagedMapV1Request(userstagedMapV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUserstagedAPI.UserstagedMapV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserstagedMapV1`: UserstagedMapV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUserstagedAPI.UserstagedMapV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUserstagedID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserstagedMapV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userstagedMapV1Request** | [**UserstagedMapV1Request**](UserstagedMapV1Request.md) |  | 

### Return type

[**UserstagedMapV1Response**](UserstagedMapV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

