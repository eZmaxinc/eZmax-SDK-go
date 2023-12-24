# eZmaxAPI\ObjectEzsigntemplatesignatureAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigntemplatesignatureCreateObjectV1**](ObjectEzsigntemplatesignatureAPI.md#EzsigntemplatesignatureCreateObjectV1) | **Post** /1/object/ezsigntemplatesignature | Create a new Ezsigntemplatesignature
[**EzsigntemplatesignatureDeleteObjectV1**](ObjectEzsigntemplatesignatureAPI.md#EzsigntemplatesignatureDeleteObjectV1) | **Delete** /1/object/ezsigntemplatesignature/{pkiEzsigntemplatesignatureID} | Delete an existing Ezsigntemplatesignature
[**EzsigntemplatesignatureEditObjectV1**](ObjectEzsigntemplatesignatureAPI.md#EzsigntemplatesignatureEditObjectV1) | **Put** /1/object/ezsigntemplatesignature/{pkiEzsigntemplatesignatureID} | Edit an existing Ezsigntemplatesignature
[**EzsigntemplatesignatureGetObjectV2**](ObjectEzsigntemplatesignatureAPI.md#EzsigntemplatesignatureGetObjectV2) | **Get** /2/object/ezsigntemplatesignature/{pkiEzsigntemplatesignatureID} | Retrieve an existing Ezsigntemplatesignature



## EzsigntemplatesignatureCreateObjectV1

> EzsigntemplatesignatureCreateObjectV1Response EzsigntemplatesignatureCreateObjectV1(ctx).EzsigntemplatesignatureCreateObjectV1Request(ezsigntemplatesignatureCreateObjectV1Request).Execute()

Create a new Ezsigntemplatesignature



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
	ezsigntemplatesignatureCreateObjectV1Request := *openapiclient.NewEzsigntemplatesignatureCreateObjectV1Request([]openapiclient.EzsigntemplatesignatureRequestCompound{*openapiclient.NewEzsigntemplatesignatureRequestCompound(int32(133), int32(9), int32(1), int32(200), int32(300), int32(1), openapiclient.Field-eEzsigntemplatesignatureType("Acknowledgement"))}) // EzsigntemplatesignatureCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureCreateObjectV1(context.Background()).EzsigntemplatesignatureCreateObjectV1Request(ezsigntemplatesignatureCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatesignatureCreateObjectV1`: EzsigntemplatesignatureCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatesignatureCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplatesignatureCreateObjectV1Request** | [**EzsigntemplatesignatureCreateObjectV1Request**](EzsigntemplatesignatureCreateObjectV1Request.md) |  | 

### Return type

[**EzsigntemplatesignatureCreateObjectV1Response**](EzsigntemplatesignatureCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatesignatureDeleteObjectV1

> EzsigntemplatesignatureDeleteObjectV1Response EzsigntemplatesignatureDeleteObjectV1(ctx, pkiEzsigntemplatesignatureID).Execute()

Delete an existing Ezsigntemplatesignature



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
	pkiEzsigntemplatesignatureID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureDeleteObjectV1(context.Background(), pkiEzsigntemplatesignatureID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatesignatureDeleteObjectV1`: EzsigntemplatesignatureDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatesignatureID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatesignatureDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatesignatureDeleteObjectV1Response**](EzsigntemplatesignatureDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatesignatureEditObjectV1

> EzsigntemplatesignatureEditObjectV1Response EzsigntemplatesignatureEditObjectV1(ctx, pkiEzsigntemplatesignatureID).EzsigntemplatesignatureEditObjectV1Request(ezsigntemplatesignatureEditObjectV1Request).Execute()

Edit an existing Ezsigntemplatesignature



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
	pkiEzsigntemplatesignatureID := int32(56) // int32 | 
	ezsigntemplatesignatureEditObjectV1Request := *openapiclient.NewEzsigntemplatesignatureEditObjectV1Request(*openapiclient.NewEzsigntemplatesignatureRequestCompound(int32(133), int32(9), int32(1), int32(200), int32(300), int32(1), openapiclient.Field-eEzsigntemplatesignatureType("Acknowledgement"))) // EzsigntemplatesignatureEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureEditObjectV1(context.Background(), pkiEzsigntemplatesignatureID).EzsigntemplatesignatureEditObjectV1Request(ezsigntemplatesignatureEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatesignatureEditObjectV1`: EzsigntemplatesignatureEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatesignatureID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatesignatureEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplatesignatureEditObjectV1Request** | [**EzsigntemplatesignatureEditObjectV1Request**](EzsigntemplatesignatureEditObjectV1Request.md) |  | 

### Return type

[**EzsigntemplatesignatureEditObjectV1Response**](EzsigntemplatesignatureEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatesignatureGetObjectV2

> EzsigntemplatesignatureGetObjectV2Response EzsigntemplatesignatureGetObjectV2(ctx, pkiEzsigntemplatesignatureID).Execute()

Retrieve an existing Ezsigntemplatesignature



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
	pkiEzsigntemplatesignatureID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureGetObjectV2(context.Background(), pkiEzsigntemplatesignatureID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatesignatureGetObjectV2`: EzsigntemplatesignatureGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatesignatureID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatesignatureGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatesignatureGetObjectV2Response**](EzsigntemplatesignatureGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

