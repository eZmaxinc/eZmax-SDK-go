# eZmaxAPI\ObjectEzsigntemplatedocumentpagerecognitionAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigntemplatedocumentpagerecognitionCreateObjectV1**](ObjectEzsigntemplatedocumentpagerecognitionAPI.md#EzsigntemplatedocumentpagerecognitionCreateObjectV1) | **Post** /1/object/ezsigntemplatedocumentpagerecognition | Create a new Ezsigntemplatedocumentpagerecognition
[**EzsigntemplatedocumentpagerecognitionDeleteObjectV1**](ObjectEzsigntemplatedocumentpagerecognitionAPI.md#EzsigntemplatedocumentpagerecognitionDeleteObjectV1) | **Delete** /1/object/ezsigntemplatedocumentpagerecognition/{pkiEzsigntemplatedocumentpagerecognitionID} | Delete an existing Ezsigntemplatedocumentpagerecognition
[**EzsigntemplatedocumentpagerecognitionEditObjectV1**](ObjectEzsigntemplatedocumentpagerecognitionAPI.md#EzsigntemplatedocumentpagerecognitionEditObjectV1) | **Put** /1/object/ezsigntemplatedocumentpagerecognition/{pkiEzsigntemplatedocumentpagerecognitionID} | Edit an existing Ezsigntemplatedocumentpagerecognition
[**EzsigntemplatedocumentpagerecognitionGetObjectV2**](ObjectEzsigntemplatedocumentpagerecognitionAPI.md#EzsigntemplatedocumentpagerecognitionGetObjectV2) | **Get** /2/object/ezsigntemplatedocumentpagerecognition/{pkiEzsigntemplatedocumentpagerecognitionID} | Retrieve an existing Ezsigntemplatedocumentpagerecognition



## EzsigntemplatedocumentpagerecognitionCreateObjectV1

> EzsigntemplatedocumentpagerecognitionCreateObjectV1Response EzsigntemplatedocumentpagerecognitionCreateObjectV1(ctx).EzsigntemplatedocumentpagerecognitionCreateObjectV1Request(ezsigntemplatedocumentpagerecognitionCreateObjectV1Request).Execute()

Create a new Ezsigntemplatedocumentpagerecognition



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
	ezsigntemplatedocumentpagerecognitionCreateObjectV1Request := *openapiclient.NewEzsigntemplatedocumentpagerecognitionCreateObjectV1Request([]openapiclient.EzsigntemplatedocumentpagerecognitionRequestCompound{*openapiclient.NewEzsigntemplatedocumentpagerecognitionRequestCompound(int32(85), openapiclient.Field-eEzsigntemplatedocumentpagerecognitionOperator("eq"), openapiclient.Field-eEzsigntemplatedocumentpagerecognitionSection("FirstLine"), "Contract")}) // EzsigntemplatedocumentpagerecognitionCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionCreateObjectV1(context.Background()).EzsigntemplatedocumentpagerecognitionCreateObjectV1Request(ezsigntemplatedocumentpagerecognitionCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatedocumentpagerecognitionCreateObjectV1`: EzsigntemplatedocumentpagerecognitionCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatedocumentpagerecognitionCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplatedocumentpagerecognitionCreateObjectV1Request** | [**EzsigntemplatedocumentpagerecognitionCreateObjectV1Request**](EzsigntemplatedocumentpagerecognitionCreateObjectV1Request.md) |  | 

### Return type

[**EzsigntemplatedocumentpagerecognitionCreateObjectV1Response**](EzsigntemplatedocumentpagerecognitionCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatedocumentpagerecognitionDeleteObjectV1

> CommonResponse EzsigntemplatedocumentpagerecognitionDeleteObjectV1(ctx, pkiEzsigntemplatedocumentpagerecognitionID).Execute()

Delete an existing Ezsigntemplatedocumentpagerecognition



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
	pkiEzsigntemplatedocumentpagerecognitionID := int32(56) // int32 | The unique ID of the Ezsigntemplatedocumentpagerecognition

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionDeleteObjectV1(context.Background(), pkiEzsigntemplatedocumentpagerecognitionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatedocumentpagerecognitionDeleteObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatedocumentpagerecognitionID** | **int32** | The unique ID of the Ezsigntemplatedocumentpagerecognition | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatedocumentpagerecognitionDeleteObjectV1Request struct via the builder pattern


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


## EzsigntemplatedocumentpagerecognitionEditObjectV1

> CommonResponse EzsigntemplatedocumentpagerecognitionEditObjectV1(ctx, pkiEzsigntemplatedocumentpagerecognitionID).EzsigntemplatedocumentpagerecognitionEditObjectV1Request(ezsigntemplatedocumentpagerecognitionEditObjectV1Request).Execute()

Edit an existing Ezsigntemplatedocumentpagerecognition



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
	pkiEzsigntemplatedocumentpagerecognitionID := int32(56) // int32 | The unique ID of the Ezsigntemplatedocumentpagerecognition
	ezsigntemplatedocumentpagerecognitionEditObjectV1Request := *openapiclient.NewEzsigntemplatedocumentpagerecognitionEditObjectV1Request(*openapiclient.NewEzsigntemplatedocumentpagerecognitionRequestCompound(int32(85), openapiclient.Field-eEzsigntemplatedocumentpagerecognitionOperator("eq"), openapiclient.Field-eEzsigntemplatedocumentpagerecognitionSection("FirstLine"), "Contract")) // EzsigntemplatedocumentpagerecognitionEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionEditObjectV1(context.Background(), pkiEzsigntemplatedocumentpagerecognitionID).EzsigntemplatedocumentpagerecognitionEditObjectV1Request(ezsigntemplatedocumentpagerecognitionEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatedocumentpagerecognitionEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatedocumentpagerecognitionID** | **int32** | The unique ID of the Ezsigntemplatedocumentpagerecognition | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatedocumentpagerecognitionEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplatedocumentpagerecognitionEditObjectV1Request** | [**EzsigntemplatedocumentpagerecognitionEditObjectV1Request**](EzsigntemplatedocumentpagerecognitionEditObjectV1Request.md) |  | 

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


## EzsigntemplatedocumentpagerecognitionGetObjectV2

> EzsigntemplatedocumentpagerecognitionGetObjectV2Response EzsigntemplatedocumentpagerecognitionGetObjectV2(ctx, pkiEzsigntemplatedocumentpagerecognitionID).Execute()

Retrieve an existing Ezsigntemplatedocumentpagerecognition



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
	pkiEzsigntemplatedocumentpagerecognitionID := int32(56) // int32 | The unique ID of the Ezsigntemplatedocumentpagerecognition

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionGetObjectV2(context.Background(), pkiEzsigntemplatedocumentpagerecognitionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatedocumentpagerecognitionGetObjectV2`: EzsigntemplatedocumentpagerecognitionGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatedocumentpagerecognitionAPI.EzsigntemplatedocumentpagerecognitionGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatedocumentpagerecognitionID** | **int32** | The unique ID of the Ezsigntemplatedocumentpagerecognition | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatedocumentpagerecognitionGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatedocumentpagerecognitionGetObjectV2Response**](EzsigntemplatedocumentpagerecognitionGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

