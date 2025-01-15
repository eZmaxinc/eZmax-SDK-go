# eZmaxAPI\ObjectEzsigntemplatepackagemembershipAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigntemplatepackagemembershipCreateObjectV1**](ObjectEzsigntemplatepackagemembershipAPI.md#EzsigntemplatepackagemembershipCreateObjectV1) | **Post** /1/object/ezsigntemplatepackagemembership | Create a new Ezsigntemplatepackagemembership
[**EzsigntemplatepackagemembershipDeleteObjectV1**](ObjectEzsigntemplatepackagemembershipAPI.md#EzsigntemplatepackagemembershipDeleteObjectV1) | **Delete** /1/object/ezsigntemplatepackagemembership/{pkiEzsigntemplatepackagemembershipID} | Delete an existing Ezsigntemplatepackagemembership
[**EzsigntemplatepackagemembershipGetObjectV2**](ObjectEzsigntemplatepackagemembershipAPI.md#EzsigntemplatepackagemembershipGetObjectV2) | **Get** /2/object/ezsigntemplatepackagemembership/{pkiEzsigntemplatepackagemembershipID} | Retrieve an existing Ezsigntemplatepackagemembership



## EzsigntemplatepackagemembershipCreateObjectV1

> EzsigntemplatepackagemembershipCreateObjectV1Response EzsigntemplatepackagemembershipCreateObjectV1(ctx).EzsigntemplatepackagemembershipCreateObjectV1Request(ezsigntemplatepackagemembershipCreateObjectV1Request).Execute()

Create a new Ezsigntemplatepackagemembership



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
	ezsigntemplatepackagemembershipCreateObjectV1Request := *openapiclient.NewEzsigntemplatepackagemembershipCreateObjectV1Request([]openapiclient.EzsigntemplatepackagemembershipRequestCompound{*openapiclient.NewEzsigntemplatepackagemembershipRequestCompound(int32(99), int32(36))}) // EzsigntemplatepackagemembershipCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepackagemembershipAPI.EzsigntemplatepackagemembershipCreateObjectV1(context.Background()).EzsigntemplatepackagemembershipCreateObjectV1Request(ezsigntemplatepackagemembershipCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackagemembershipAPI.EzsigntemplatepackagemembershipCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepackagemembershipCreateObjectV1`: EzsigntemplatepackagemembershipCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackagemembershipAPI.EzsigntemplatepackagemembershipCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackagemembershipCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplatepackagemembershipCreateObjectV1Request** | [**EzsigntemplatepackagemembershipCreateObjectV1Request**](EzsigntemplatepackagemembershipCreateObjectV1Request.md) |  | 

### Return type

[**EzsigntemplatepackagemembershipCreateObjectV1Response**](EzsigntemplatepackagemembershipCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepackagemembershipDeleteObjectV1

> CommonResponse EzsigntemplatepackagemembershipDeleteObjectV1(ctx, pkiEzsigntemplatepackagemembershipID).Execute()

Delete an existing Ezsigntemplatepackagemembership



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
	pkiEzsigntemplatepackagemembershipID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepackagemembershipAPI.EzsigntemplatepackagemembershipDeleteObjectV1(context.Background(), pkiEzsigntemplatepackagemembershipID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackagemembershipAPI.EzsigntemplatepackagemembershipDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepackagemembershipDeleteObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackagemembershipAPI.EzsigntemplatepackagemembershipDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepackagemembershipID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackagemembershipDeleteObjectV1Request struct via the builder pattern


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


## EzsigntemplatepackagemembershipGetObjectV2

> EzsigntemplatepackagemembershipGetObjectV2Response EzsigntemplatepackagemembershipGetObjectV2(ctx, pkiEzsigntemplatepackagemembershipID).Execute()

Retrieve an existing Ezsigntemplatepackagemembership



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
	pkiEzsigntemplatepackagemembershipID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepackagemembershipAPI.EzsigntemplatepackagemembershipGetObjectV2(context.Background(), pkiEzsigntemplatepackagemembershipID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackagemembershipAPI.EzsigntemplatepackagemembershipGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepackagemembershipGetObjectV2`: EzsigntemplatepackagemembershipGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackagemembershipAPI.EzsigntemplatepackagemembershipGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepackagemembershipID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackagemembershipGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatepackagemembershipGetObjectV2Response**](EzsigntemplatepackagemembershipGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

