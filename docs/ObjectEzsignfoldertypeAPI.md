# eZmaxAPI\ObjectEzsignfoldertypeAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignfoldertypeCreateObjectV3**](ObjectEzsignfoldertypeAPI.md#EzsignfoldertypeCreateObjectV3) | **Post** /3/object/ezsignfoldertype | Create a new Ezsignfoldertype
[**EzsignfoldertypeEditObjectV3**](ObjectEzsignfoldertypeAPI.md#EzsignfoldertypeEditObjectV3) | **Put** /3/object/ezsignfoldertype/{pkiEzsignfoldertypeID} | Edit an existing Ezsignfoldertype
[**EzsignfoldertypeGetAutocompleteV2**](ObjectEzsignfoldertypeAPI.md#EzsignfoldertypeGetAutocompleteV2) | **Get** /2/object/ezsignfoldertype/getAutocomplete/{sSelector} | Retrieve Ezsignfoldertypes and IDs
[**EzsignfoldertypeGetListV1**](ObjectEzsignfoldertypeAPI.md#EzsignfoldertypeGetListV1) | **Get** /1/object/ezsignfoldertype/getList | Retrieve Ezsignfoldertype list
[**EzsignfoldertypeGetObjectV2**](ObjectEzsignfoldertypeAPI.md#EzsignfoldertypeGetObjectV2) | **Get** /2/object/ezsignfoldertype/{pkiEzsignfoldertypeID} | Retrieve an existing Ezsignfoldertype
[**EzsignfoldertypeGetObjectV4**](ObjectEzsignfoldertypeAPI.md#EzsignfoldertypeGetObjectV4) | **Get** /4/object/ezsignfoldertype/{pkiEzsignfoldertypeID} | Retrieve an existing Ezsignfoldertype



## EzsignfoldertypeCreateObjectV3

> EzsignfoldertypeCreateObjectV3Response EzsignfoldertypeCreateObjectV3(ctx).EzsignfoldertypeCreateObjectV3Request(ezsignfoldertypeCreateObjectV3Request).Execute()

Create a new Ezsignfoldertype



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
	ezsignfoldertypeCreateObjectV3Request := *openapiclient.NewEzsignfoldertypeCreateObjectV3Request([]openapiclient.EzsignfoldertypeRequestCompoundV3{*openapiclient.NewEzsignfoldertypeRequestCompoundV3(*openapiclient.NewMultilingualEzsignfoldertypeName(), int32(78), []int32{int32(2)}, openapiclient.Field-eEzsignfoldertypePrivacylevel("User"), int32(30), openapiclient.Field-eEzsignfoldertypeDisposal("No"), openapiclient.Field-eEzsignfoldertypeCompletion("PerEzsigndocument"), int32(5), false, false, false, false, false, false, true)}) // EzsignfoldertypeCreateObjectV3Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignfoldertypeAPI.EzsignfoldertypeCreateObjectV3(context.Background()).EzsignfoldertypeCreateObjectV3Request(ezsignfoldertypeCreateObjectV3Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldertypeAPI.EzsignfoldertypeCreateObjectV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldertypeCreateObjectV3`: EzsignfoldertypeCreateObjectV3Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldertypeAPI.EzsignfoldertypeCreateObjectV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldertypeCreateObjectV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignfoldertypeCreateObjectV3Request** | [**EzsignfoldertypeCreateObjectV3Request**](EzsignfoldertypeCreateObjectV3Request.md) |  | 

### Return type

[**EzsignfoldertypeCreateObjectV3Response**](EzsignfoldertypeCreateObjectV3Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldertypeEditObjectV3

> CommonResponse EzsignfoldertypeEditObjectV3(ctx, pkiEzsignfoldertypeID).EzsignfoldertypeEditObjectV3Request(ezsignfoldertypeEditObjectV3Request).Execute()

Edit an existing Ezsignfoldertype



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
	pkiEzsignfoldertypeID := int32(56) // int32 | 
	ezsignfoldertypeEditObjectV3Request := *openapiclient.NewEzsignfoldertypeEditObjectV3Request(*openapiclient.NewEzsignfoldertypeRequestCompoundV3(*openapiclient.NewMultilingualEzsignfoldertypeName(), int32(78), []int32{int32(2)}, openapiclient.Field-eEzsignfoldertypePrivacylevel("User"), int32(30), openapiclient.Field-eEzsignfoldertypeDisposal("No"), openapiclient.Field-eEzsignfoldertypeCompletion("PerEzsigndocument"), int32(5), false, false, false, false, false, false, true)) // EzsignfoldertypeEditObjectV3Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignfoldertypeAPI.EzsignfoldertypeEditObjectV3(context.Background(), pkiEzsignfoldertypeID).EzsignfoldertypeEditObjectV3Request(ezsignfoldertypeEditObjectV3Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldertypeAPI.EzsignfoldertypeEditObjectV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldertypeEditObjectV3`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldertypeAPI.EzsignfoldertypeEditObjectV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfoldertypeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldertypeEditObjectV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfoldertypeEditObjectV3Request** | [**EzsignfoldertypeEditObjectV3Request**](EzsignfoldertypeEditObjectV3Request.md) |  | 

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


## EzsignfoldertypeGetAutocompleteV2

> EzsignfoldertypeGetAutocompleteV2Response EzsignfoldertypeGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Ezsignfoldertypes and IDs



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
	sSelector := "sSelector_example" // string | The type of Ezsignfoldertypes to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignfoldertypeAPI.EzsignfoldertypeGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldertypeAPI.EzsignfoldertypeGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldertypeGetAutocompleteV2`: EzsignfoldertypeGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldertypeAPI.EzsignfoldertypeGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Ezsignfoldertypes to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldertypeGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**EzsignfoldertypeGetAutocompleteV2Response**](EzsignfoldertypeGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldertypeGetListV1

> EzsignfoldertypeGetListV1Response EzsignfoldertypeGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Ezsignfoldertype list



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
	resp, r, err := apiClient.ObjectEzsignfoldertypeAPI.EzsignfoldertypeGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldertypeAPI.EzsignfoldertypeGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldertypeGetListV1`: EzsignfoldertypeGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldertypeAPI.EzsignfoldertypeGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldertypeGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**EzsignfoldertypeGetListV1Response**](EzsignfoldertypeGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldertypeGetObjectV2

> EzsignfoldertypeGetObjectV2Response EzsignfoldertypeGetObjectV2(ctx, pkiEzsignfoldertypeID).Execute()

Retrieve an existing Ezsignfoldertype



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
	pkiEzsignfoldertypeID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignfoldertypeAPI.EzsignfoldertypeGetObjectV2(context.Background(), pkiEzsignfoldertypeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldertypeAPI.EzsignfoldertypeGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldertypeGetObjectV2`: EzsignfoldertypeGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldertypeAPI.EzsignfoldertypeGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfoldertypeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldertypeGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfoldertypeGetObjectV2Response**](EzsignfoldertypeGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldertypeGetObjectV4

> EzsignfoldertypeGetObjectV4Response EzsignfoldertypeGetObjectV4(ctx, pkiEzsignfoldertypeID).Execute()

Retrieve an existing Ezsignfoldertype



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
	pkiEzsignfoldertypeID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignfoldertypeAPI.EzsignfoldertypeGetObjectV4(context.Background(), pkiEzsignfoldertypeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldertypeAPI.EzsignfoldertypeGetObjectV4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignfoldertypeGetObjectV4`: EzsignfoldertypeGetObjectV4Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldertypeAPI.EzsignfoldertypeGetObjectV4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfoldertypeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldertypeGetObjectV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfoldertypeGetObjectV4Response**](EzsignfoldertypeGetObjectV4Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

