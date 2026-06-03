# eZmaxAPI\ObjectEzmaxpartnerAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzmaxpartnerGetCustomDeveloppersV1**](ObjectEzmaxpartnerAPI.md#EzmaxpartnerGetCustomDeveloppersV1) | **Get** /1/object/ezmaxpartner/getCustomDeveloppers | Retrieve Ezmaxpartner custom developpers list
[**EzmaxpartnerGetObjectV2**](ObjectEzmaxpartnerAPI.md#EzmaxpartnerGetObjectV2) | **Get** /2/object/ezmaxpartner/{pkiEzmaxpartnerID} | Retrieve an existing Ezmaxpartner



## EzmaxpartnerGetCustomDeveloppersV1

> EzmaxpartnerGetCustomDeveloppersV1Response EzmaxpartnerGetCustomDeveloppersV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Ezmaxpartner custom developpers list

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
	resp, r, err := apiClient.ObjectEzmaxpartnerAPI.EzmaxpartnerGetCustomDeveloppersV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzmaxpartnerAPI.EzmaxpartnerGetCustomDeveloppersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxpartnerGetCustomDeveloppersV1`: EzmaxpartnerGetCustomDeveloppersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzmaxpartnerAPI.EzmaxpartnerGetCustomDeveloppersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxpartnerGetCustomDeveloppersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**EzmaxpartnerGetCustomDeveloppersV1Response**](EzmaxpartnerGetCustomDeveloppersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzmaxpartnerGetObjectV2

> EzmaxpartnerGetObjectV2Response EzmaxpartnerGetObjectV2(ctx, pkiEzmaxpartnerID).Execute()

Retrieve an existing Ezmaxpartner



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
	pkiEzmaxpartnerID := int32(56) // int32 | The unique ID of the Ezmaxpartner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzmaxpartnerAPI.EzmaxpartnerGetObjectV2(context.Background(), pkiEzmaxpartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzmaxpartnerAPI.EzmaxpartnerGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxpartnerGetObjectV2`: EzmaxpartnerGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzmaxpartnerAPI.EzmaxpartnerGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzmaxpartnerID** | **int32** | The unique ID of the Ezmaxpartner | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxpartnerGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzmaxpartnerGetObjectV2Response**](EzmaxpartnerGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

