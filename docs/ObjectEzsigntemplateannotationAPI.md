# eZmaxAPI\ObjectEzsigntemplateannotationAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigntemplateannotationCreateObjectV1**](ObjectEzsigntemplateannotationAPI.md#EzsigntemplateannotationCreateObjectV1) | **Post** /1/object/ezsigntemplateannotation | Create a new Ezsigntemplateannotation
[**EzsigntemplateannotationDeleteObjectV1**](ObjectEzsigntemplateannotationAPI.md#EzsigntemplateannotationDeleteObjectV1) | **Delete** /1/object/ezsigntemplateannotation/{pkiEzsigntemplateannotationID} | Delete an existing Ezsigntemplateannotation
[**EzsigntemplateannotationEditObjectV1**](ObjectEzsigntemplateannotationAPI.md#EzsigntemplateannotationEditObjectV1) | **Put** /1/object/ezsigntemplateannotation/{pkiEzsigntemplateannotationID} | Edit an existing Ezsigntemplateannotation
[**EzsigntemplateannotationGetObjectV2**](ObjectEzsigntemplateannotationAPI.md#EzsigntemplateannotationGetObjectV2) | **Get** /2/object/ezsigntemplateannotation/{pkiEzsigntemplateannotationID} | Retrieve an existing Ezsigntemplateannotation



## EzsigntemplateannotationCreateObjectV1

> EzsigntemplateannotationCreateObjectV1Response EzsigntemplateannotationCreateObjectV1(ctx).EzsigntemplateannotationCreateObjectV1Request(ezsigntemplateannotationCreateObjectV1Request).Execute()

Create a new Ezsigntemplateannotation



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
	ezsigntemplateannotationCreateObjectV1Request := *openapiclient.NewEzsigntemplateannotationCreateObjectV1Request([]openapiclient.EzsigntemplateannotationRequestCompound{*openapiclient.NewEzsigntemplateannotationRequestCompound(int32(85), int32(216), openapiclient.Field-eEzsigntemplateannotationHorizontalalignment("Center"), openapiclient.Field-eEzsigntemplateannotationVerticalalignment("Bottom"), openapiclient.Field-eEzsigntemplateannotationType("Dropdown"), int32(17864), int32(23342), int32(51755), int32(58213), "Name", "John Doe", "{"Accepts","Refuses"}")}) // EzsigntemplateannotationCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplateannotationAPI.EzsigntemplateannotationCreateObjectV1(context.Background()).EzsigntemplateannotationCreateObjectV1Request(ezsigntemplateannotationCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateannotationAPI.EzsigntemplateannotationCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateannotationCreateObjectV1`: EzsigntemplateannotationCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateannotationAPI.EzsigntemplateannotationCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateannotationCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplateannotationCreateObjectV1Request** | [**EzsigntemplateannotationCreateObjectV1Request**](EzsigntemplateannotationCreateObjectV1Request.md) |  | 

### Return type

[**EzsigntemplateannotationCreateObjectV1Response**](EzsigntemplateannotationCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplateannotationDeleteObjectV1

> EzsigntemplateannotationDeleteObjectV1Response EzsigntemplateannotationDeleteObjectV1(ctx, pkiEzsigntemplateannotationID).Execute()

Delete an existing Ezsigntemplateannotation



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
	pkiEzsigntemplateannotationID := int32(56) // int32 | The unique ID of the Ezsigntemplateannotation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplateannotationAPI.EzsigntemplateannotationDeleteObjectV1(context.Background(), pkiEzsigntemplateannotationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateannotationAPI.EzsigntemplateannotationDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateannotationDeleteObjectV1`: EzsigntemplateannotationDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateannotationAPI.EzsigntemplateannotationDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplateannotationID** | **int32** | The unique ID of the Ezsigntemplateannotation | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateannotationDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplateannotationDeleteObjectV1Response**](EzsigntemplateannotationDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplateannotationEditObjectV1

> EzsigntemplateannotationEditObjectV1Response EzsigntemplateannotationEditObjectV1(ctx, pkiEzsigntemplateannotationID).EzsigntemplateannotationEditObjectV1Request(ezsigntemplateannotationEditObjectV1Request).Execute()

Edit an existing Ezsigntemplateannotation



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
	pkiEzsigntemplateannotationID := int32(56) // int32 | The unique ID of the Ezsigntemplateannotation
	ezsigntemplateannotationEditObjectV1Request := *openapiclient.NewEzsigntemplateannotationEditObjectV1Request(*openapiclient.NewEzsigntemplateannotationRequestCompound(int32(85), int32(216), openapiclient.Field-eEzsigntemplateannotationHorizontalalignment("Center"), openapiclient.Field-eEzsigntemplateannotationVerticalalignment("Bottom"), openapiclient.Field-eEzsigntemplateannotationType("Dropdown"), int32(17864), int32(23342), int32(51755), int32(58213), "Name", "John Doe", "{"Accepts","Refuses"}")) // EzsigntemplateannotationEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplateannotationAPI.EzsigntemplateannotationEditObjectV1(context.Background(), pkiEzsigntemplateannotationID).EzsigntemplateannotationEditObjectV1Request(ezsigntemplateannotationEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateannotationAPI.EzsigntemplateannotationEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateannotationEditObjectV1`: EzsigntemplateannotationEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateannotationAPI.EzsigntemplateannotationEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplateannotationID** | **int32** | The unique ID of the Ezsigntemplateannotation | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateannotationEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplateannotationEditObjectV1Request** | [**EzsigntemplateannotationEditObjectV1Request**](EzsigntemplateannotationEditObjectV1Request.md) |  | 

### Return type

[**EzsigntemplateannotationEditObjectV1Response**](EzsigntemplateannotationEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplateannotationGetObjectV2

> EzsigntemplateannotationGetObjectV2Response EzsigntemplateannotationGetObjectV2(ctx, pkiEzsigntemplateannotationID).Execute()

Retrieve an existing Ezsigntemplateannotation



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
	pkiEzsigntemplateannotationID := int32(56) // int32 | The unique ID of the Ezsigntemplateannotation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplateannotationAPI.EzsigntemplateannotationGetObjectV2(context.Background(), pkiEzsigntemplateannotationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateannotationAPI.EzsigntemplateannotationGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateannotationGetObjectV2`: EzsigntemplateannotationGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateannotationAPI.EzsigntemplateannotationGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplateannotationID** | **int32** | The unique ID of the Ezsigntemplateannotation | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateannotationGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplateannotationGetObjectV2Response**](EzsigntemplateannotationGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

