# eZmaxAPI\ObjectEzsignsignatureAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignsignatureCreateObjectV1**](ObjectEzsignsignatureAPI.md#EzsignsignatureCreateObjectV1) | **Post** /1/object/ezsignsignature | Create a new Ezsignsignature
[**EzsignsignatureCreateObjectV2**](ObjectEzsignsignatureAPI.md#EzsignsignatureCreateObjectV2) | **Post** /2/object/ezsignsignature | Create a new Ezsignsignature
[**EzsignsignatureCreateObjectV3**](ObjectEzsignsignatureAPI.md#EzsignsignatureCreateObjectV3) | **Post** /3/object/ezsignsignature | Create a new Ezsignsignature
[**EzsignsignatureDeleteObjectV1**](ObjectEzsignsignatureAPI.md#EzsignsignatureDeleteObjectV1) | **Delete** /1/object/ezsignsignature/{pkiEzsignsignatureID} | Delete an existing Ezsignsignature
[**EzsignsignatureEditObjectV2**](ObjectEzsignsignatureAPI.md#EzsignsignatureEditObjectV2) | **Put** /2/object/ezsignsignature/{pkiEzsignsignatureID} | Edit an existing Ezsignsignature
[**EzsignsignatureGetEzsignsignatureattachmentV1**](ObjectEzsignsignatureAPI.md#EzsignsignatureGetEzsignsignatureattachmentV1) | **Get** /1/object/ezsignsignature/{pkiEzsignsignatureID}/getEzsignsignatureattachment | Retrieve an existing Ezsignsignature&#39;s Ezsignsignatureattachments
[**EzsignsignatureGetEzsignsignaturesAutomaticV1**](ObjectEzsignsignatureAPI.md#EzsignsignatureGetEzsignsignaturesAutomaticV1) | **Get** /1/object/ezsignsignature/getEzsignsignaturesAutomatic | Retrieve all automatic Ezsignsignatures
[**EzsignsignatureGetObjectV3**](ObjectEzsignsignatureAPI.md#EzsignsignatureGetObjectV3) | **Get** /3/object/ezsignsignature/{pkiEzsignsignatureID} | Retrieve an existing Ezsignsignature
[**EzsignsignatureSignV1**](ObjectEzsignsignatureAPI.md#EzsignsignatureSignV1) | **Post** /1/object/ezsignsignature/{pkiEzsignsignatureID}/sign | Sign the Ezsignsignature



## EzsignsignatureCreateObjectV1

> EzsignsignatureCreateObjectV1Response EzsignsignatureCreateObjectV1(ctx).EzsignsignatureCreateObjectV1Request(ezsignsignatureCreateObjectV1Request).Execute()

Create a new Ezsignsignature



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
	ezsignsignatureCreateObjectV1Request := []openapiclient.EzsignsignatureCreateObjectV1Request{*openapiclient.NewEzsignsignatureCreateObjectV1Request()} // []EzsignsignatureCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureCreateObjectV1(context.Background()).EzsignsignatureCreateObjectV1Request(ezsignsignatureCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignatureAPI.EzsignsignatureCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsignatureCreateObjectV1`: EzsignsignatureCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignatureAPI.EzsignsignatureCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignatureCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignsignatureCreateObjectV1Request** | [**[]EzsignsignatureCreateObjectV1Request**](EzsignsignatureCreateObjectV1Request.md) |  | 

### Return type

[**EzsignsignatureCreateObjectV1Response**](EzsignsignatureCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignatureCreateObjectV2

> EzsignsignatureCreateObjectV2Response EzsignsignatureCreateObjectV2(ctx).EzsignsignatureCreateObjectV2Request(ezsignsignatureCreateObjectV2Request).Execute()

Create a new Ezsignsignature



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
	ezsignsignatureCreateObjectV2Request := *openapiclient.NewEzsignsignatureCreateObjectV2Request([]openapiclient.EzsignsignatureRequestCompound{*openapiclient.NewEzsignsignatureRequestCompound(int32(20), int32(1), int32(200), int32(300), int32(1), openapiclient.Field-eEzsignsignatureType("Acknowledgement"), int32(97))}) // EzsignsignatureCreateObjectV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureCreateObjectV2(context.Background()).EzsignsignatureCreateObjectV2Request(ezsignsignatureCreateObjectV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignatureAPI.EzsignsignatureCreateObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsignatureCreateObjectV2`: EzsignsignatureCreateObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignatureAPI.EzsignsignatureCreateObjectV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignatureCreateObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignsignatureCreateObjectV2Request** | [**EzsignsignatureCreateObjectV2Request**](EzsignsignatureCreateObjectV2Request.md) |  | 

### Return type

[**EzsignsignatureCreateObjectV2Response**](EzsignsignatureCreateObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignatureCreateObjectV3

> EzsignsignatureCreateObjectV3Response EzsignsignatureCreateObjectV3(ctx).EzsignsignatureCreateObjectV3Request(ezsignsignatureCreateObjectV3Request).Execute()

Create a new Ezsignsignature



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
	ezsignsignatureCreateObjectV3Request := *openapiclient.NewEzsignsignatureCreateObjectV3Request([]openapiclient.EzsignsignatureRequestCompoundV2{*openapiclient.NewEzsignsignatureRequestCompoundV2(int32(20), int32(1), int32(200), int32(300), int32(1), openapiclient.Field-eEzsignsignatureType("Acknowledgement"), int32(97))}) // EzsignsignatureCreateObjectV3Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureCreateObjectV3(context.Background()).EzsignsignatureCreateObjectV3Request(ezsignsignatureCreateObjectV3Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignatureAPI.EzsignsignatureCreateObjectV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsignatureCreateObjectV3`: EzsignsignatureCreateObjectV3Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignatureAPI.EzsignsignatureCreateObjectV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignatureCreateObjectV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignsignatureCreateObjectV3Request** | [**EzsignsignatureCreateObjectV3Request**](EzsignsignatureCreateObjectV3Request.md) |  | 

### Return type

[**EzsignsignatureCreateObjectV3Response**](EzsignsignatureCreateObjectV3Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignatureDeleteObjectV1

> CommonResponse EzsignsignatureDeleteObjectV1(ctx, pkiEzsignsignatureID).Execute()

Delete an existing Ezsignsignature



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
	pkiEzsignsignatureID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureDeleteObjectV1(context.Background(), pkiEzsignsignatureID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignatureAPI.EzsignsignatureDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsignatureDeleteObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignatureAPI.EzsignsignatureDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignatureID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignatureDeleteObjectV1Request struct via the builder pattern


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


## EzsignsignatureEditObjectV2

> CommonResponse EzsignsignatureEditObjectV2(ctx, pkiEzsignsignatureID).EzsignsignatureEditObjectV2Request(ezsignsignatureEditObjectV2Request).Execute()

Edit an existing Ezsignsignature



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
	pkiEzsignsignatureID := int32(56) // int32 | 
	ezsignsignatureEditObjectV2Request := *openapiclient.NewEzsignsignatureEditObjectV2Request(*openapiclient.NewEzsignsignatureRequestCompoundV2(int32(20), int32(1), int32(200), int32(300), int32(1), openapiclient.Field-eEzsignsignatureType("Acknowledgement"), int32(97))) // EzsignsignatureEditObjectV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureEditObjectV2(context.Background(), pkiEzsignsignatureID).EzsignsignatureEditObjectV2Request(ezsignsignatureEditObjectV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignatureAPI.EzsignsignatureEditObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsignatureEditObjectV2`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignatureAPI.EzsignsignatureEditObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignatureID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignatureEditObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignsignatureEditObjectV2Request** | [**EzsignsignatureEditObjectV2Request**](EzsignsignatureEditObjectV2Request.md) |  | 

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


## EzsignsignatureGetEzsignsignatureattachmentV1

> EzsignsignatureGetEzsignsignatureattachmentV1Response EzsignsignatureGetEzsignsignatureattachmentV1(ctx, pkiEzsignsignatureID).Execute()

Retrieve an existing Ezsignsignature's Ezsignsignatureattachments

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
	pkiEzsignsignatureID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureGetEzsignsignatureattachmentV1(context.Background(), pkiEzsignsignatureID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignatureAPI.EzsignsignatureGetEzsignsignatureattachmentV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsignatureGetEzsignsignatureattachmentV1`: EzsignsignatureGetEzsignsignatureattachmentV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignatureAPI.EzsignsignatureGetEzsignsignatureattachmentV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignatureID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignatureGetEzsignsignatureattachmentV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignsignatureGetEzsignsignatureattachmentV1Response**](EzsignsignatureGetEzsignsignatureattachmentV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignatureGetEzsignsignaturesAutomaticV1

> EzsignsignatureGetEzsignsignaturesAutomaticV1Response EzsignsignatureGetEzsignsignaturesAutomaticV1(ctx).Execute()

Retrieve all automatic Ezsignsignatures



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureGetEzsignsignaturesAutomaticV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignatureAPI.EzsignsignatureGetEzsignsignaturesAutomaticV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsignatureGetEzsignsignaturesAutomaticV1`: EzsignsignatureGetEzsignsignaturesAutomaticV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignatureAPI.EzsignsignatureGetEzsignsignaturesAutomaticV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignatureGetEzsignsignaturesAutomaticV1Request struct via the builder pattern


### Return type

[**EzsignsignatureGetEzsignsignaturesAutomaticV1Response**](EzsignsignatureGetEzsignsignaturesAutomaticV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignatureGetObjectV3

> EzsignsignatureGetObjectV3Response EzsignsignatureGetObjectV3(ctx, pkiEzsignsignatureID).Execute()

Retrieve an existing Ezsignsignature



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
	pkiEzsignsignatureID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureGetObjectV3(context.Background(), pkiEzsignsignatureID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignatureAPI.EzsignsignatureGetObjectV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsignatureGetObjectV3`: EzsignsignatureGetObjectV3Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignatureAPI.EzsignsignatureGetObjectV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignatureID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignatureGetObjectV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignsignatureGetObjectV3Response**](EzsignsignatureGetObjectV3Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignatureSignV1

> EzsignsignatureSignV1Response EzsignsignatureSignV1(ctx, pkiEzsignsignatureID).EzsignsignatureSignV1Request(ezsignsignatureSignV1Request).Execute()

Sign the Ezsignsignature



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
	pkiEzsignsignatureID := int32(56) // int32 | 
	ezsignsignatureSignV1Request := *openapiclient.NewEzsignsignatureSignV1Request(false) // EzsignsignatureSignV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsignatureAPI.EzsignsignatureSignV1(context.Background(), pkiEzsignsignatureID).EzsignsignatureSignV1Request(ezsignsignatureSignV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignatureAPI.EzsignsignatureSignV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsignatureSignV1`: EzsignsignatureSignV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignatureAPI.EzsignsignatureSignV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignatureID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignatureSignV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignsignatureSignV1Request** | [**EzsignsignatureSignV1Request**](EzsignsignatureSignV1Request.md) |  | 

### Return type

[**EzsignsignatureSignV1Response**](EzsignsignatureSignV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

