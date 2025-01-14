# eZmaxAPI\ObjectEzsigntemplateAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigntemplateCopyV1**](ObjectEzsigntemplateAPI.md#EzsigntemplateCopyV1) | **Post** /1/object/ezsigntemplate/{pkiEzsigntemplateID}/copy | Copy the Ezsigntemplate
[**EzsigntemplateCreateObjectV3**](ObjectEzsigntemplateAPI.md#EzsigntemplateCreateObjectV3) | **Post** /3/object/ezsigntemplate | Create a new Ezsigntemplate
[**EzsigntemplateDeleteObjectV1**](ObjectEzsigntemplateAPI.md#EzsigntemplateDeleteObjectV1) | **Delete** /1/object/ezsigntemplate/{pkiEzsigntemplateID} | Delete an existing Ezsigntemplate
[**EzsigntemplateEditObjectV3**](ObjectEzsigntemplateAPI.md#EzsigntemplateEditObjectV3) | **Put** /3/object/ezsigntemplate/{pkiEzsigntemplateID} | Edit an existing Ezsigntemplate
[**EzsigntemplateGetAutocompleteV2**](ObjectEzsigntemplateAPI.md#EzsigntemplateGetAutocompleteV2) | **Get** /2/object/ezsigntemplate/getAutocomplete/{sSelector} | Retrieve Ezsigntemplates and IDs
[**EzsigntemplateGetListV1**](ObjectEzsigntemplateAPI.md#EzsigntemplateGetListV1) | **Get** /1/object/ezsigntemplate/getList | Retrieve Ezsigntemplate list
[**EzsigntemplateGetObjectV3**](ObjectEzsigntemplateAPI.md#EzsigntemplateGetObjectV3) | **Get** /3/object/ezsigntemplate/{pkiEzsigntemplateID} | Retrieve an existing Ezsigntemplate



## EzsigntemplateCopyV1

> EzsigntemplateCopyV1Response EzsigntemplateCopyV1(ctx, pkiEzsigntemplateID).EzsigntemplateCopyV1Request(ezsigntemplateCopyV1Request).Execute()

Copy the Ezsigntemplate



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
	pkiEzsigntemplateID := int32(56) // int32 | 
	ezsigntemplateCopyV1Request := *openapiclient.NewEzsigntemplateCopyV1Request() // EzsigntemplateCopyV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplateAPI.EzsigntemplateCopyV1(context.Background(), pkiEzsigntemplateID).EzsigntemplateCopyV1Request(ezsigntemplateCopyV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateAPI.EzsigntemplateCopyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateCopyV1`: EzsigntemplateCopyV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateAPI.EzsigntemplateCopyV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplateID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateCopyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplateCopyV1Request** | [**EzsigntemplateCopyV1Request**](EzsigntemplateCopyV1Request.md) |  | 

### Return type

[**EzsigntemplateCopyV1Response**](EzsigntemplateCopyV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplateCreateObjectV3

> EzsigntemplateCreateObjectV3Response EzsigntemplateCreateObjectV3(ctx).EzsigntemplateCreateObjectV3Request(ezsigntemplateCreateObjectV3Request).Execute()

Create a new Ezsigntemplate



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
	ezsigntemplateCreateObjectV3Request := *openapiclient.NewEzsigntemplateCreateObjectV3Request([]openapiclient.EzsigntemplateRequestCompoundV3{*openapiclient.NewEzsigntemplateRequestCompoundV3(int32(2), "Standard Contract", false, openapiclient.Field-eEzsigntemplateType("User"))}) // EzsigntemplateCreateObjectV3Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplateAPI.EzsigntemplateCreateObjectV3(context.Background()).EzsigntemplateCreateObjectV3Request(ezsigntemplateCreateObjectV3Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateAPI.EzsigntemplateCreateObjectV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateCreateObjectV3`: EzsigntemplateCreateObjectV3Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateAPI.EzsigntemplateCreateObjectV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateCreateObjectV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplateCreateObjectV3Request** | [**EzsigntemplateCreateObjectV3Request**](EzsigntemplateCreateObjectV3Request.md) |  | 

### Return type

[**EzsigntemplateCreateObjectV3Response**](EzsigntemplateCreateObjectV3Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplateDeleteObjectV1

> EzsigntemplateDeleteObjectV1Response EzsigntemplateDeleteObjectV1(ctx, pkiEzsigntemplateID).Execute()

Delete an existing Ezsigntemplate



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
	pkiEzsigntemplateID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplateAPI.EzsigntemplateDeleteObjectV1(context.Background(), pkiEzsigntemplateID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateAPI.EzsigntemplateDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateDeleteObjectV1`: EzsigntemplateDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateAPI.EzsigntemplateDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplateID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplateDeleteObjectV1Response**](EzsigntemplateDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplateEditObjectV3

> EzsigntemplateEditObjectV3Response EzsigntemplateEditObjectV3(ctx, pkiEzsigntemplateID).EzsigntemplateEditObjectV3Request(ezsigntemplateEditObjectV3Request).Execute()

Edit an existing Ezsigntemplate



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
	pkiEzsigntemplateID := int32(56) // int32 | 
	ezsigntemplateEditObjectV3Request := *openapiclient.NewEzsigntemplateEditObjectV3Request(*openapiclient.NewEzsigntemplateRequestCompoundV3(int32(2), "Standard Contract", false, openapiclient.Field-eEzsigntemplateType("User"))) // EzsigntemplateEditObjectV3Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplateAPI.EzsigntemplateEditObjectV3(context.Background(), pkiEzsigntemplateID).EzsigntemplateEditObjectV3Request(ezsigntemplateEditObjectV3Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateAPI.EzsigntemplateEditObjectV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateEditObjectV3`: EzsigntemplateEditObjectV3Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateAPI.EzsigntemplateEditObjectV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplateID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateEditObjectV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplateEditObjectV3Request** | [**EzsigntemplateEditObjectV3Request**](EzsigntemplateEditObjectV3Request.md) |  | 

### Return type

[**EzsigntemplateEditObjectV3Response**](EzsigntemplateEditObjectV3Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplateGetAutocompleteV2

> EzsigntemplateGetAutocompleteV2Response EzsigntemplateGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).FkiEzsignfoldertypeID(fkiEzsignfoldertypeID).Execute()

Retrieve Ezsigntemplates and IDs



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
	sSelector := "sSelector_example" // string | The type of Ezsigntemplates to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)
	fkiEzsignfoldertypeID := int32(56) // int32 | The fkiEzsignfoldertypeID to use with the selector Ezsigntemplatepublic (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplateAPI.EzsigntemplateGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).FkiEzsignfoldertypeID(fkiEzsignfoldertypeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateAPI.EzsigntemplateGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateGetAutocompleteV2`: EzsigntemplateGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateAPI.EzsigntemplateGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Ezsigntemplates to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **fkiEzsignfoldertypeID** | **int32** | The fkiEzsignfoldertypeID to use with the selector Ezsigntemplatepublic | 

### Return type

[**EzsigntemplateGetAutocompleteV2Response**](EzsigntemplateGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplateGetListV1

> EzsigntemplateGetListV1Response EzsigntemplateGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Ezsigntemplate list



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
	resp, r, err := apiClient.ObjectEzsigntemplateAPI.EzsigntemplateGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateAPI.EzsigntemplateGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateGetListV1`: EzsigntemplateGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateAPI.EzsigntemplateGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**EzsigntemplateGetListV1Response**](EzsigntemplateGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplateGetObjectV3

> EzsigntemplateGetObjectV3Response EzsigntemplateGetObjectV3(ctx, pkiEzsigntemplateID).Execute()

Retrieve an existing Ezsigntemplate



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
	pkiEzsigntemplateID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplateAPI.EzsigntemplateGetObjectV3(context.Background(), pkiEzsigntemplateID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplateAPI.EzsigntemplateGetObjectV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplateGetObjectV3`: EzsigntemplateGetObjectV3Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplateAPI.EzsigntemplateGetObjectV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplateID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplateGetObjectV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplateGetObjectV3Response**](EzsigntemplateGetObjectV3Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

