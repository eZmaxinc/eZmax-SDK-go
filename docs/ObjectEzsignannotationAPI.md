# eZmaxAPI\ObjectEzsignannotationAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignannotationCreateObjectV1**](ObjectEzsignannotationAPI.md#EzsignannotationCreateObjectV1) | **Post** /1/object/ezsignannotation | Create a new Ezsignannotation
[**EzsignannotationDeleteObjectV1**](ObjectEzsignannotationAPI.md#EzsignannotationDeleteObjectV1) | **Delete** /1/object/ezsignannotation/{pkiEzsignannotationID} | Delete an existing Ezsignannotation
[**EzsignannotationEditObjectV1**](ObjectEzsignannotationAPI.md#EzsignannotationEditObjectV1) | **Put** /1/object/ezsignannotation/{pkiEzsignannotationID} | Edit an existing Ezsignannotation
[**EzsignannotationGetObjectV2**](ObjectEzsignannotationAPI.md#EzsignannotationGetObjectV2) | **Get** /2/object/ezsignannotation/{pkiEzsignannotationID} | Retrieve an existing Ezsignannotation



## EzsignannotationCreateObjectV1

> EzsignannotationCreateObjectV1Response EzsignannotationCreateObjectV1(ctx).EzsignannotationCreateObjectV1Request(ezsignannotationCreateObjectV1Request).Execute()

Create a new Ezsignannotation



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
	ezsignannotationCreateObjectV1Request := *openapiclient.NewEzsignannotationCreateObjectV1Request([]openapiclient.EzsignannotationRequestCompound{*openapiclient.NewEzsignannotationRequestCompound(int32(97), openapiclient.Field-eEzsignannotationType("StrikethroughBlock"), int32(50), int32(50), int32(1))}) // EzsignannotationCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignannotationAPI.EzsignannotationCreateObjectV1(context.Background()).EzsignannotationCreateObjectV1Request(ezsignannotationCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignannotationAPI.EzsignannotationCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignannotationCreateObjectV1`: EzsignannotationCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignannotationAPI.EzsignannotationCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignannotationCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignannotationCreateObjectV1Request** | [**EzsignannotationCreateObjectV1Request**](EzsignannotationCreateObjectV1Request.md) |  | 

### Return type

[**EzsignannotationCreateObjectV1Response**](EzsignannotationCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignannotationDeleteObjectV1

> CommonResponse EzsignannotationDeleteObjectV1(ctx, pkiEzsignannotationID).Execute()

Delete an existing Ezsignannotation



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
	pkiEzsignannotationID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignannotationAPI.EzsignannotationDeleteObjectV1(context.Background(), pkiEzsignannotationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignannotationAPI.EzsignannotationDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignannotationDeleteObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignannotationAPI.EzsignannotationDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignannotationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignannotationDeleteObjectV1Request struct via the builder pattern


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


## EzsignannotationEditObjectV1

> CommonResponse EzsignannotationEditObjectV1(ctx, pkiEzsignannotationID).EzsignannotationEditObjectV1Request(ezsignannotationEditObjectV1Request).Execute()

Edit an existing Ezsignannotation



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
	pkiEzsignannotationID := int32(56) // int32 | 
	ezsignannotationEditObjectV1Request := *openapiclient.NewEzsignannotationEditObjectV1Request(*openapiclient.NewEzsignannotationRequestCompound(int32(97), openapiclient.Field-eEzsignannotationType("StrikethroughBlock"), int32(50), int32(50), int32(1))) // EzsignannotationEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignannotationAPI.EzsignannotationEditObjectV1(context.Background(), pkiEzsignannotationID).EzsignannotationEditObjectV1Request(ezsignannotationEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignannotationAPI.EzsignannotationEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignannotationEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignannotationAPI.EzsignannotationEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignannotationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignannotationEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignannotationEditObjectV1Request** | [**EzsignannotationEditObjectV1Request**](EzsignannotationEditObjectV1Request.md) |  | 

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


## EzsignannotationGetObjectV2

> EzsignannotationGetObjectV2Response EzsignannotationGetObjectV2(ctx, pkiEzsignannotationID).Execute()

Retrieve an existing Ezsignannotation



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
	pkiEzsignannotationID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignannotationAPI.EzsignannotationGetObjectV2(context.Background(), pkiEzsignannotationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignannotationAPI.EzsignannotationGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignannotationGetObjectV2`: EzsignannotationGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignannotationAPI.EzsignannotationGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignannotationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignannotationGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignannotationGetObjectV2Response**](EzsignannotationGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

