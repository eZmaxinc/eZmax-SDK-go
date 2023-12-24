# eZmaxAPI\ObjectUsergroupdelegationAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsergroupdelegationCreateObjectV1**](ObjectUsergroupdelegationAPI.md#UsergroupdelegationCreateObjectV1) | **Post** /1/object/usergroupdelegation | Create a new Usergroupdelegation
[**UsergroupdelegationDeleteObjectV1**](ObjectUsergroupdelegationAPI.md#UsergroupdelegationDeleteObjectV1) | **Delete** /1/object/usergroupdelegation/{pkiUsergroupdelegationID} | Delete an existing Usergroupdelegation
[**UsergroupdelegationEditObjectV1**](ObjectUsergroupdelegationAPI.md#UsergroupdelegationEditObjectV1) | **Put** /1/object/usergroupdelegation/{pkiUsergroupdelegationID} | Edit an existing Usergroupdelegation
[**UsergroupdelegationGetObjectV2**](ObjectUsergroupdelegationAPI.md#UsergroupdelegationGetObjectV2) | **Get** /2/object/usergroupdelegation/{pkiUsergroupdelegationID} | Retrieve an existing Usergroupdelegation



## UsergroupdelegationCreateObjectV1

> UsergroupdelegationCreateObjectV1Response UsergroupdelegationCreateObjectV1(ctx).UsergroupdelegationCreateObjectV1Request(usergroupdelegationCreateObjectV1Request).Execute()

Create a new Usergroupdelegation



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
	usergroupdelegationCreateObjectV1Request := *openapiclient.NewUsergroupdelegationCreateObjectV1Request([]openapiclient.UsergroupdelegationRequestCompound{*openapiclient.NewUsergroupdelegationRequestCompound(int32(2), int32(70))}) // UsergroupdelegationCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUsergroupdelegationAPI.UsergroupdelegationCreateObjectV1(context.Background()).UsergroupdelegationCreateObjectV1Request(usergroupdelegationCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupdelegationAPI.UsergroupdelegationCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupdelegationCreateObjectV1`: UsergroupdelegationCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupdelegationAPI.UsergroupdelegationCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupdelegationCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usergroupdelegationCreateObjectV1Request** | [**UsergroupdelegationCreateObjectV1Request**](UsergroupdelegationCreateObjectV1Request.md) |  | 

### Return type

[**UsergroupdelegationCreateObjectV1Response**](UsergroupdelegationCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupdelegationDeleteObjectV1

> UsergroupdelegationDeleteObjectV1Response UsergroupdelegationDeleteObjectV1(ctx, pkiUsergroupdelegationID).Execute()

Delete an existing Usergroupdelegation



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
	pkiUsergroupdelegationID := int32(56) // int32 | The unique ID of the Usergroupdelegation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUsergroupdelegationAPI.UsergroupdelegationDeleteObjectV1(context.Background(), pkiUsergroupdelegationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupdelegationAPI.UsergroupdelegationDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupdelegationDeleteObjectV1`: UsergroupdelegationDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupdelegationAPI.UsergroupdelegationDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupdelegationID** | **int32** | The unique ID of the Usergroupdelegation | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupdelegationDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsergroupdelegationDeleteObjectV1Response**](UsergroupdelegationDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupdelegationEditObjectV1

> UsergroupdelegationEditObjectV1Response UsergroupdelegationEditObjectV1(ctx, pkiUsergroupdelegationID).UsergroupdelegationEditObjectV1Request(usergroupdelegationEditObjectV1Request).Execute()

Edit an existing Usergroupdelegation



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
	pkiUsergroupdelegationID := int32(56) // int32 | The unique ID of the Usergroupdelegation
	usergroupdelegationEditObjectV1Request := *openapiclient.NewUsergroupdelegationEditObjectV1Request(*openapiclient.NewUsergroupdelegationRequestCompound(int32(2), int32(70))) // UsergroupdelegationEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUsergroupdelegationAPI.UsergroupdelegationEditObjectV1(context.Background(), pkiUsergroupdelegationID).UsergroupdelegationEditObjectV1Request(usergroupdelegationEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupdelegationAPI.UsergroupdelegationEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupdelegationEditObjectV1`: UsergroupdelegationEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupdelegationAPI.UsergroupdelegationEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupdelegationID** | **int32** | The unique ID of the Usergroupdelegation | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupdelegationEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usergroupdelegationEditObjectV1Request** | [**UsergroupdelegationEditObjectV1Request**](UsergroupdelegationEditObjectV1Request.md) |  | 

### Return type

[**UsergroupdelegationEditObjectV1Response**](UsergroupdelegationEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupdelegationGetObjectV2

> UsergroupdelegationGetObjectV2Response UsergroupdelegationGetObjectV2(ctx, pkiUsergroupdelegationID).Execute()

Retrieve an existing Usergroupdelegation



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
	pkiUsergroupdelegationID := int32(56) // int32 | The unique ID of the Usergroupdelegation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUsergroupdelegationAPI.UsergroupdelegationGetObjectV2(context.Background(), pkiUsergroupdelegationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupdelegationAPI.UsergroupdelegationGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupdelegationGetObjectV2`: UsergroupdelegationGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupdelegationAPI.UsergroupdelegationGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupdelegationID** | **int32** | The unique ID of the Usergroupdelegation | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupdelegationGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsergroupdelegationGetObjectV2Response**](UsergroupdelegationGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

