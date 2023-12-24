# eZmaxAPI\ObjectEzsigntemplatepackageAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigntemplatepackageCreateObjectV1**](ObjectEzsigntemplatepackageAPI.md#EzsigntemplatepackageCreateObjectV1) | **Post** /1/object/ezsigntemplatepackage | Create a new Ezsigntemplatepackage
[**EzsigntemplatepackageDeleteObjectV1**](ObjectEzsigntemplatepackageAPI.md#EzsigntemplatepackageDeleteObjectV1) | **Delete** /1/object/ezsigntemplatepackage/{pkiEzsigntemplatepackageID} | Delete an existing Ezsigntemplatepackage
[**EzsigntemplatepackageEditEzsigntemplatepackagesignersV1**](ObjectEzsigntemplatepackageAPI.md#EzsigntemplatepackageEditEzsigntemplatepackagesignersV1) | **Put** /1/object/ezsigntemplatepackage/{pkiEzsigntemplatepackageID}/editEzsigntemplatepackagesigners | Edit multiple Ezsigntemplatepackagesigners
[**EzsigntemplatepackageEditObjectV1**](ObjectEzsigntemplatepackageAPI.md#EzsigntemplatepackageEditObjectV1) | **Put** /1/object/ezsigntemplatepackage/{pkiEzsigntemplatepackageID} | Edit an existing Ezsigntemplatepackage
[**EzsigntemplatepackageGetAutocompleteV2**](ObjectEzsigntemplatepackageAPI.md#EzsigntemplatepackageGetAutocompleteV2) | **Get** /2/object/ezsigntemplatepackage/getAutocomplete/{sSelector} | Retrieve Ezsigntemplatepackages and IDs
[**EzsigntemplatepackageGetListV1**](ObjectEzsigntemplatepackageAPI.md#EzsigntemplatepackageGetListV1) | **Get** /1/object/ezsigntemplatepackage/getList | Retrieve Ezsigntemplatepackage list
[**EzsigntemplatepackageGetObjectV2**](ObjectEzsigntemplatepackageAPI.md#EzsigntemplatepackageGetObjectV2) | **Get** /2/object/ezsigntemplatepackage/{pkiEzsigntemplatepackageID} | Retrieve an existing Ezsigntemplatepackage



## EzsigntemplatepackageCreateObjectV1

> EzsigntemplatepackageCreateObjectV1Response EzsigntemplatepackageCreateObjectV1(ctx).EzsigntemplatepackageCreateObjectV1Request(ezsigntemplatepackageCreateObjectV1Request).Execute()

Create a new Ezsigntemplatepackage



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
	ezsigntemplatepackageCreateObjectV1Request := *openapiclient.NewEzsigntemplatepackageCreateObjectV1Request([]openapiclient.EzsigntemplatepackageRequestCompound{*openapiclient.NewEzsigntemplatepackageRequestCompound(int32(5), int32(2), "Package for new clients", false, true)}) // EzsigntemplatepackageCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageCreateObjectV1(context.Background()).EzsigntemplatepackageCreateObjectV1Request(ezsigntemplatepackageCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepackageCreateObjectV1`: EzsigntemplatepackageCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackageCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplatepackageCreateObjectV1Request** | [**EzsigntemplatepackageCreateObjectV1Request**](EzsigntemplatepackageCreateObjectV1Request.md) |  | 

### Return type

[**EzsigntemplatepackageCreateObjectV1Response**](EzsigntemplatepackageCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepackageDeleteObjectV1

> EzsigntemplatepackageDeleteObjectV1Response EzsigntemplatepackageDeleteObjectV1(ctx, pkiEzsigntemplatepackageID).Execute()

Delete an existing Ezsigntemplatepackage



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
	pkiEzsigntemplatepackageID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageDeleteObjectV1(context.Background(), pkiEzsigntemplatepackageID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepackageDeleteObjectV1`: EzsigntemplatepackageDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepackageID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackageDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatepackageDeleteObjectV1Response**](EzsigntemplatepackageDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepackageEditEzsigntemplatepackagesignersV1

> EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Response EzsigntemplatepackageEditEzsigntemplatepackagesignersV1(ctx, pkiEzsigntemplatepackageID).EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request(ezsigntemplatepackageEditEzsigntemplatepackagesignersV1Request).Execute()

Edit multiple Ezsigntemplatepackagesigners



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
	pkiEzsigntemplatepackageID := int32(56) // int32 | 
	ezsigntemplatepackageEditEzsigntemplatepackagesignersV1Request := *openapiclient.NewEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request([]openapiclient.EzsigntemplatepackagesignerRequestCompound{*openapiclient.NewEzsigntemplatepackagesignerRequestCompound(int32(99), "Customer")}) // EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageEditEzsigntemplatepackagesignersV1(context.Background(), pkiEzsigntemplatepackageID).EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request(ezsigntemplatepackageEditEzsigntemplatepackagesignersV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageEditEzsigntemplatepackagesignersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepackageEditEzsigntemplatepackagesignersV1`: EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageEditEzsigntemplatepackagesignersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepackageID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplatepackageEditEzsigntemplatepackagesignersV1Request** | [**EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request**](EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Request.md) |  | 

### Return type

[**EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Response**](EzsigntemplatepackageEditEzsigntemplatepackagesignersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepackageEditObjectV1

> EzsigntemplatepackageEditObjectV1Response EzsigntemplatepackageEditObjectV1(ctx, pkiEzsigntemplatepackageID).EzsigntemplatepackageEditObjectV1Request(ezsigntemplatepackageEditObjectV1Request).Execute()

Edit an existing Ezsigntemplatepackage



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
	pkiEzsigntemplatepackageID := int32(56) // int32 | 
	ezsigntemplatepackageEditObjectV1Request := *openapiclient.NewEzsigntemplatepackageEditObjectV1Request(*openapiclient.NewEzsigntemplatepackageRequestCompound(int32(5), int32(2), "Package for new clients", false, true)) // EzsigntemplatepackageEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageEditObjectV1(context.Background(), pkiEzsigntemplatepackageID).EzsigntemplatepackageEditObjectV1Request(ezsigntemplatepackageEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepackageEditObjectV1`: EzsigntemplatepackageEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepackageID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackageEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplatepackageEditObjectV1Request** | [**EzsigntemplatepackageEditObjectV1Request**](EzsigntemplatepackageEditObjectV1Request.md) |  | 

### Return type

[**EzsigntemplatepackageEditObjectV1Response**](EzsigntemplatepackageEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepackageGetAutocompleteV2

> EzsigntemplatepackageGetAutocompleteV2Response EzsigntemplatepackageGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Ezsigntemplatepackages and IDs



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
	sSelector := "sSelector_example" // string | The type of Ezsigntemplatepackages to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepackageGetAutocompleteV2`: EzsigntemplatepackageGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Ezsigntemplatepackages to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackageGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**EzsigntemplatepackageGetAutocompleteV2Response**](EzsigntemplatepackageGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepackageGetListV1

> EzsigntemplatepackageGetListV1Response EzsigntemplatepackageGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Ezsigntemplatepackage list



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
	resp, r, err := apiClient.ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepackageGetListV1`: EzsigntemplatepackageGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackageGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**EzsigntemplatepackageGetListV1Response**](EzsigntemplatepackageGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepackageGetObjectV2

> EzsigntemplatepackageGetObjectV2Response EzsigntemplatepackageGetObjectV2(ctx, pkiEzsigntemplatepackageID).Execute()

Retrieve an existing Ezsigntemplatepackage



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
	pkiEzsigntemplatepackageID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageGetObjectV2(context.Background(), pkiEzsigntemplatepackageID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepackageGetObjectV2`: EzsigntemplatepackageGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepackageAPI.EzsigntemplatepackageGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepackageID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepackageGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatepackageGetObjectV2Response**](EzsigntemplatepackageGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

