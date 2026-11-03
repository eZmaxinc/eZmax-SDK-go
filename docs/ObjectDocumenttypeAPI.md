# eZmaxAPI\ObjectDocumenttypeAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocumenttypeCreateObjectV1**](ObjectDocumenttypeAPI.md#DocumenttypeCreateObjectV1) | **Post** /1/object/documenttype | Create a new Documenttype
[**DocumenttypeEditObjectV1**](ObjectDocumenttypeAPI.md#DocumenttypeEditObjectV1) | **Put** /1/object/documenttype/{pkiDocumenttypeID} | Edit an existing Documenttype
[**DocumenttypeGetAutocompleteV2**](ObjectDocumenttypeAPI.md#DocumenttypeGetAutocompleteV2) | **Get** /2/object/documenttype/getAutocomplete/{sSelector} | Retrieve Documenttypes and IDs
[**DocumenttypeGetListV1**](ObjectDocumenttypeAPI.md#DocumenttypeGetListV1) | **Get** /1/object/documenttype/getList | Retrieve Documenttype list
[**DocumenttypeGetObjectV2**](ObjectDocumenttypeAPI.md#DocumenttypeGetObjectV2) | **Get** /2/object/documenttype/{pkiDocumenttypeID} | Retrieve an existing Documenttype



## DocumenttypeCreateObjectV1

> DocumenttypeCreateObjectV1Response DocumenttypeCreateObjectV1(ctx).DocumenttypeCreateObjectV1Request(documenttypeCreateObjectV1Request).Execute()

Create a new Documenttype



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
	documenttypeCreateObjectV1Request := *openapiclient.NewDocumenttypeCreateObjectV1Request([]openapiclient.DocumenttypeRequestCompound{*openapiclient.NewDocumenttypeRequestCompound(*openapiclient.NewMultilingualDocumenttypeName(), openapiclient.Computed-eDocumenttypeStep("Agentbroker"), "invoice.pdf", true, true, true, true, true, true, true, true, true, true, true, true, true, openapiclient.Field-eDocumenttypePrivacydefault("Administration"), openapiclient.Field-eDocumenttypeStatusdefault("Not required"), openapiclient.Field-eDocumenttypeSalebyowner("All"), openapiclient.Field-eDocumenttypeSaledisplaybyrole("InscriptorSeller"), openapiclient.Field-eDocumenttypeTransactiontype("All"), openapiclient.Field-eDocumenttypeRepossession("All"), []int32{int32(51)})}) // DocumenttypeCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDocumenttypeAPI.DocumenttypeCreateObjectV1(context.Background()).DocumenttypeCreateObjectV1Request(documenttypeCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDocumenttypeAPI.DocumenttypeCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumenttypeCreateObjectV1`: DocumenttypeCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDocumenttypeAPI.DocumenttypeCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumenttypeCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documenttypeCreateObjectV1Request** | [**DocumenttypeCreateObjectV1Request**](DocumenttypeCreateObjectV1Request.md) |  | 

### Return type

[**DocumenttypeCreateObjectV1Response**](DocumenttypeCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumenttypeEditObjectV1

> DocumenttypeEditObjectV1Response DocumenttypeEditObjectV1(ctx, pkiDocumenttypeID).DocumenttypeEditObjectV1Request(documenttypeEditObjectV1Request).Execute()

Edit an existing Documenttype



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
	pkiDocumenttypeID := int32(56) // int32 | The unique ID of the Documenttype
	documenttypeEditObjectV1Request := *openapiclient.NewDocumenttypeEditObjectV1Request(*openapiclient.NewDocumenttypeRequestCompound(*openapiclient.NewMultilingualDocumenttypeName(), openapiclient.Computed-eDocumenttypeStep("Agentbroker"), "invoice.pdf", true, true, true, true, true, true, true, true, true, true, true, true, true, openapiclient.Field-eDocumenttypePrivacydefault("Administration"), openapiclient.Field-eDocumenttypeStatusdefault("Not required"), openapiclient.Field-eDocumenttypeSalebyowner("All"), openapiclient.Field-eDocumenttypeSaledisplaybyrole("InscriptorSeller"), openapiclient.Field-eDocumenttypeTransactiontype("All"), openapiclient.Field-eDocumenttypeRepossession("All"), []int32{int32(51)})) // DocumenttypeEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDocumenttypeAPI.DocumenttypeEditObjectV1(context.Background(), pkiDocumenttypeID).DocumenttypeEditObjectV1Request(documenttypeEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDocumenttypeAPI.DocumenttypeEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumenttypeEditObjectV1`: DocumenttypeEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDocumenttypeAPI.DocumenttypeEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDocumenttypeID** | **int32** | The unique ID of the Documenttype | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocumenttypeEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **documenttypeEditObjectV1Request** | [**DocumenttypeEditObjectV1Request**](DocumenttypeEditObjectV1Request.md) |  | 

### Return type

[**DocumenttypeEditObjectV1Response**](DocumenttypeEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumenttypeGetAutocompleteV2

> DocumenttypeGetAutocompleteV2Response DocumenttypeGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Documenttypes and IDs



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
	sSelector := "sSelector_example" // string | The type of Documenttypes to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDocumenttypeAPI.DocumenttypeGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDocumenttypeAPI.DocumenttypeGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumenttypeGetAutocompleteV2`: DocumenttypeGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDocumenttypeAPI.DocumenttypeGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Documenttypes to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocumenttypeGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**DocumenttypeGetAutocompleteV2Response**](DocumenttypeGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumenttypeGetListV1

> DocumenttypeGetListV1Response DocumenttypeGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Documenttype list



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
	resp, r, err := apiClient.ObjectDocumenttypeAPI.DocumenttypeGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDocumenttypeAPI.DocumenttypeGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumenttypeGetListV1`: DocumenttypeGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDocumenttypeAPI.DocumenttypeGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumenttypeGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**DocumenttypeGetListV1Response**](DocumenttypeGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumenttypeGetObjectV2

> DocumenttypeGetObjectV2Response DocumenttypeGetObjectV2(ctx, pkiDocumenttypeID).Execute()

Retrieve an existing Documenttype



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
	pkiDocumenttypeID := int32(56) // int32 | The unique ID of the Documenttype

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDocumenttypeAPI.DocumenttypeGetObjectV2(context.Background(), pkiDocumenttypeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDocumenttypeAPI.DocumenttypeGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumenttypeGetObjectV2`: DocumenttypeGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDocumenttypeAPI.DocumenttypeGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDocumenttypeID** | **int32** | The unique ID of the Documenttype | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocumenttypeGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DocumenttypeGetObjectV2Response**](DocumenttypeGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

