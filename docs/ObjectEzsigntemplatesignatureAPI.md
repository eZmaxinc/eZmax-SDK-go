# eZmaxAPI\ObjectEzsigntemplatesignatureAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigntemplatesignatureCreateObjectV2**](ObjectEzsigntemplatesignatureAPI.md#EzsigntemplatesignatureCreateObjectV2) | **Post** /2/object/ezsigntemplatesignature | Create a new Ezsigntemplatesignature
[**EzsigntemplatesignatureDeleteObjectV1**](ObjectEzsigntemplatesignatureAPI.md#EzsigntemplatesignatureDeleteObjectV1) | **Delete** /1/object/ezsigntemplatesignature/{pkiEzsigntemplatesignatureID} | Delete an existing Ezsigntemplatesignature
[**EzsigntemplatesignatureEditObjectV2**](ObjectEzsigntemplatesignatureAPI.md#EzsigntemplatesignatureEditObjectV2) | **Put** /2/object/ezsigntemplatesignature/{pkiEzsigntemplatesignatureID} | Edit an existing Ezsigntemplatesignature
[**EzsigntemplatesignatureGetObjectV3**](ObjectEzsigntemplatesignatureAPI.md#EzsigntemplatesignatureGetObjectV3) | **Get** /3/object/ezsigntemplatesignature/{pkiEzsigntemplatesignatureID} | Retrieve an existing Ezsigntemplatesignature



## EzsigntemplatesignatureCreateObjectV2

> EzsigntemplatesignatureCreateObjectV2Response EzsigntemplatesignatureCreateObjectV2(ctx).EzsigntemplatesignatureCreateObjectV2Request(ezsigntemplatesignatureCreateObjectV2Request).Execute()

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
	ezsigntemplatesignatureCreateObjectV2Request := *openapiclient.NewEzsigntemplatesignatureCreateObjectV2Request([]openapiclient.EzsigntemplatesignatureRequestCompoundV2{*openapiclient.NewEzsigntemplatesignatureRequestCompoundV2(int32(133), int32(9), int32(1), int32(1), openapiclient.Field-eEzsigntemplatesignatureType("Acknowledgement"))}) // EzsigntemplatesignatureCreateObjectV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureCreateObjectV2(context.Background()).EzsigntemplatesignatureCreateObjectV2Request(ezsigntemplatesignatureCreateObjectV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureCreateObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatesignatureCreateObjectV2`: EzsigntemplatesignatureCreateObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureCreateObjectV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatesignatureCreateObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplatesignatureCreateObjectV2Request** | [**EzsigntemplatesignatureCreateObjectV2Request**](EzsigntemplatesignatureCreateObjectV2Request.md) |  | 

### Return type

[**EzsigntemplatesignatureCreateObjectV2Response**](EzsigntemplatesignatureCreateObjectV2Response.md)

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


## EzsigntemplatesignatureEditObjectV2

> EzsigntemplatesignatureEditObjectV2Response EzsigntemplatesignatureEditObjectV2(ctx, pkiEzsigntemplatesignatureID).EzsigntemplatesignatureEditObjectV2Request(ezsigntemplatesignatureEditObjectV2Request).Execute()

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
	ezsigntemplatesignatureEditObjectV2Request := *openapiclient.NewEzsigntemplatesignatureEditObjectV2Request(*openapiclient.NewEzsigntemplatesignatureRequestCompoundV2(int32(133), int32(9), int32(1), int32(1), openapiclient.Field-eEzsigntemplatesignatureType("Acknowledgement"))) // EzsigntemplatesignatureEditObjectV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureEditObjectV2(context.Background(), pkiEzsigntemplatesignatureID).EzsigntemplatesignatureEditObjectV2Request(ezsigntemplatesignatureEditObjectV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureEditObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatesignatureEditObjectV2`: EzsigntemplatesignatureEditObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureEditObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatesignatureID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatesignatureEditObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplatesignatureEditObjectV2Request** | [**EzsigntemplatesignatureEditObjectV2Request**](EzsigntemplatesignatureEditObjectV2Request.md) |  | 

### Return type

[**EzsigntemplatesignatureEditObjectV2Response**](EzsigntemplatesignatureEditObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatesignatureGetObjectV3

> EzsigntemplatesignatureGetObjectV3Response EzsigntemplatesignatureGetObjectV3(ctx, pkiEzsigntemplatesignatureID).Execute()

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
	resp, r, err := apiClient.ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureGetObjectV3(context.Background(), pkiEzsigntemplatesignatureID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureGetObjectV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatesignatureGetObjectV3`: EzsigntemplatesignatureGetObjectV3Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatesignatureAPI.EzsigntemplatesignatureGetObjectV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatesignatureID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatesignatureGetObjectV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatesignatureGetObjectV3Response**](EzsigntemplatesignatureGetObjectV3Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

