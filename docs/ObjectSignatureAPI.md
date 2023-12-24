# eZmaxAPI\ObjectSignatureAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SignatureCreateObjectV1**](ObjectSignatureAPI.md#SignatureCreateObjectV1) | **Post** /1/object/signature | Create a new Signature
[**SignatureDeleteObjectV1**](ObjectSignatureAPI.md#SignatureDeleteObjectV1) | **Delete** /1/object/signature/{pkiSignatureID} | Delete an existing Signature
[**SignatureEditObjectV1**](ObjectSignatureAPI.md#SignatureEditObjectV1) | **Put** /1/object/signature/{pkiSignatureID} | Edit an existing Signature
[**SignatureGetObjectV2**](ObjectSignatureAPI.md#SignatureGetObjectV2) | **Get** /2/object/signature/{pkiSignatureID} | Retrieve an existing Signature



## SignatureCreateObjectV1

> SignatureCreateObjectV1Response SignatureCreateObjectV1(ctx).SignatureCreateObjectV1Request(signatureCreateObjectV1Request).Execute()

Create a new Signature



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
	signatureCreateObjectV1Request := *openapiclient.NewSignatureCreateObjectV1Request([]openapiclient.SignatureRequestCompound{*openapiclient.NewSignatureRequestCompound("{"$ref":"#/components/examples/Svg/value"}")}) // SignatureCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSignatureAPI.SignatureCreateObjectV1(context.Background()).SignatureCreateObjectV1Request(signatureCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSignatureAPI.SignatureCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignatureCreateObjectV1`: SignatureCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSignatureAPI.SignatureCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignatureCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signatureCreateObjectV1Request** | [**SignatureCreateObjectV1Request**](SignatureCreateObjectV1Request.md) |  | 

### Return type

[**SignatureCreateObjectV1Response**](SignatureCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignatureDeleteObjectV1

> SignatureDeleteObjectV1Response SignatureDeleteObjectV1(ctx, pkiSignatureID).Execute()

Delete an existing Signature



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
	pkiSignatureID := int32(56) // int32 | The unique ID of the Signature

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSignatureAPI.SignatureDeleteObjectV1(context.Background(), pkiSignatureID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSignatureAPI.SignatureDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignatureDeleteObjectV1`: SignatureDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSignatureAPI.SignatureDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSignatureID** | **int32** | The unique ID of the Signature | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignatureDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SignatureDeleteObjectV1Response**](SignatureDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignatureEditObjectV1

> SignatureEditObjectV1Response SignatureEditObjectV1(ctx, pkiSignatureID).SignatureEditObjectV1Request(signatureEditObjectV1Request).Execute()

Edit an existing Signature



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
	pkiSignatureID := int32(56) // int32 | The unique ID of the Signature
	signatureEditObjectV1Request := *openapiclient.NewSignatureEditObjectV1Request(*openapiclient.NewSignatureRequestCompound("{"$ref":"#/components/examples/Svg/value"}")) // SignatureEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSignatureAPI.SignatureEditObjectV1(context.Background(), pkiSignatureID).SignatureEditObjectV1Request(signatureEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSignatureAPI.SignatureEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignatureEditObjectV1`: SignatureEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSignatureAPI.SignatureEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSignatureID** | **int32** | The unique ID of the Signature | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignatureEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signatureEditObjectV1Request** | [**SignatureEditObjectV1Request**](SignatureEditObjectV1Request.md) |  | 

### Return type

[**SignatureEditObjectV1Response**](SignatureEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignatureGetObjectV2

> SignatureGetObjectV2Response SignatureGetObjectV2(ctx, pkiSignatureID).Execute()

Retrieve an existing Signature



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
	pkiSignatureID := int32(56) // int32 | The unique ID of the Signature

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSignatureAPI.SignatureGetObjectV2(context.Background(), pkiSignatureID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSignatureAPI.SignatureGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignatureGetObjectV2`: SignatureGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSignatureAPI.SignatureGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSignatureID** | **int32** | The unique ID of the Signature | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignatureGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SignatureGetObjectV2Response**](SignatureGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

