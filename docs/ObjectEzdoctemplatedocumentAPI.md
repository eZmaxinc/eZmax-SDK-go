# eZmaxAPI\ObjectEzdoctemplatedocumentAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzdoctemplatedocumentCreateObjectV1**](ObjectEzdoctemplatedocumentAPI.md#EzdoctemplatedocumentCreateObjectV1) | **Post** /1/object/ezdoctemplatedocument | Create a new Ezdoctemplatedocument
[**EzdoctemplatedocumentDownloadV1**](ObjectEzdoctemplatedocumentAPI.md#EzdoctemplatedocumentDownloadV1) | **Get** /1/object/ezdoctemplatedocument/{pkiEzdoctemplatedocumentID}/download | Retrieve the content
[**EzdoctemplatedocumentEditObjectV1**](ObjectEzdoctemplatedocumentAPI.md#EzdoctemplatedocumentEditObjectV1) | **Put** /1/object/ezdoctemplatedocument/{pkiEzdoctemplatedocumentID} | Edit an existing Ezdoctemplatedocument
[**EzdoctemplatedocumentGetAutocompleteV2**](ObjectEzdoctemplatedocumentAPI.md#EzdoctemplatedocumentGetAutocompleteV2) | **Get** /2/object/ezdoctemplatedocument/getAutocomplete/{sSelector} | Retrieve Ezdoctemplatedocuments and IDs
[**EzdoctemplatedocumentGetListV1**](ObjectEzdoctemplatedocumentAPI.md#EzdoctemplatedocumentGetListV1) | **Get** /1/object/ezdoctemplatedocument/getList | Retrieve Ezdoctemplatedocument list
[**EzdoctemplatedocumentGetObjectV2**](ObjectEzdoctemplatedocumentAPI.md#EzdoctemplatedocumentGetObjectV2) | **Get** /2/object/ezdoctemplatedocument/{pkiEzdoctemplatedocumentID} | Retrieve an existing Ezdoctemplatedocument
[**EzdoctemplatedocumentPatchObjectV1**](ObjectEzdoctemplatedocumentAPI.md#EzdoctemplatedocumentPatchObjectV1) | **Patch** /1/object/ezdoctemplatedocument/{pkiEzdoctemplatedocumentID} | Patch an existing Ezdoctemplatedocument



## EzdoctemplatedocumentCreateObjectV1

> EzdoctemplatedocumentCreateObjectV1Response EzdoctemplatedocumentCreateObjectV1(ctx).EzdoctemplatedocumentCreateObjectV1Request(ezdoctemplatedocumentCreateObjectV1Request).Execute()

Create a new Ezdoctemplatedocument



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
	ezdoctemplatedocumentCreateObjectV1Request := *openapiclient.NewEzdoctemplatedocumentCreateObjectV1Request([]openapiclient.EzdoctemplatedocumentRequestCompound{*openapiclient.NewEzdoctemplatedocumentRequestCompound(int32(2), int32(7), int32(4), true, *openapiclient.NewMultilingualEzdoctemplatedocumentName())}) // EzdoctemplatedocumentCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentCreateObjectV1(context.Background()).EzdoctemplatedocumentCreateObjectV1Request(ezdoctemplatedocumentCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzdoctemplatedocumentCreateObjectV1`: EzdoctemplatedocumentCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzdoctemplatedocumentCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezdoctemplatedocumentCreateObjectV1Request** | [**EzdoctemplatedocumentCreateObjectV1Request**](EzdoctemplatedocumentCreateObjectV1Request.md) |  | 

### Return type

[**EzdoctemplatedocumentCreateObjectV1Response**](EzdoctemplatedocumentCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzdoctemplatedocumentDownloadV1

> EzdoctemplatedocumentDownloadV1(ctx, pkiEzdoctemplatedocumentID).Execute()

Retrieve the content



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
	pkiEzdoctemplatedocumentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentDownloadV1(context.Background(), pkiEzdoctemplatedocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzdoctemplatedocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzdoctemplatedocumentDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization), [Presigned](../README.md#Presigned)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzdoctemplatedocumentEditObjectV1

> CommonResponse EzdoctemplatedocumentEditObjectV1(ctx, pkiEzdoctemplatedocumentID).EzdoctemplatedocumentEditObjectV1Request(ezdoctemplatedocumentEditObjectV1Request).Execute()

Edit an existing Ezdoctemplatedocument



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
	pkiEzdoctemplatedocumentID := int32(56) // int32 | The unique ID of the Ezdoctemplatedocument
	ezdoctemplatedocumentEditObjectV1Request := *openapiclient.NewEzdoctemplatedocumentEditObjectV1Request(*openapiclient.NewEzdoctemplatedocumentRequestCompound(int32(2), int32(7), int32(4), true, *openapiclient.NewMultilingualEzdoctemplatedocumentName())) // EzdoctemplatedocumentEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentEditObjectV1(context.Background(), pkiEzdoctemplatedocumentID).EzdoctemplatedocumentEditObjectV1Request(ezdoctemplatedocumentEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzdoctemplatedocumentEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzdoctemplatedocumentID** | **int32** | The unique ID of the Ezdoctemplatedocument | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzdoctemplatedocumentEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezdoctemplatedocumentEditObjectV1Request** | [**EzdoctemplatedocumentEditObjectV1Request**](EzdoctemplatedocumentEditObjectV1Request.md) |  | 

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


## EzdoctemplatedocumentGetAutocompleteV2

> EzdoctemplatedocumentGetAutocompleteV2Response EzdoctemplatedocumentGetAutocompleteV2(ctx, sSelector).EType(eType).FkiEzsignfoldertypeID(fkiEzsignfoldertypeID).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Ezdoctemplatedocuments and IDs



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
	sSelector := "sSelector_example" // string | The type of Ezdoctemplatedocuments to return
	eType := "eType_example" // string | The type of Ezdoctemplatedocument (default to "CompanyEzsignfoldertype")
	fkiEzsignfoldertypeID := "fkiEzsignfoldertypeID_example" // string | Specify which fkiEzsignfoldertypeID we want to display. only used when eType = Ezsignfoldertype (optional)
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentGetAutocompleteV2(context.Background(), sSelector).EType(eType).FkiEzsignfoldertypeID(fkiEzsignfoldertypeID).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzdoctemplatedocumentGetAutocompleteV2`: EzdoctemplatedocumentGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Ezdoctemplatedocuments to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzdoctemplatedocumentGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eType** | **string** | The type of Ezdoctemplatedocument | [default to &quot;CompanyEzsignfoldertype&quot;]
 **fkiEzsignfoldertypeID** | **string** | Specify which fkiEzsignfoldertypeID we want to display. only used when eType &#x3D; Ezsignfoldertype | 
 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**EzdoctemplatedocumentGetAutocompleteV2Response**](EzdoctemplatedocumentGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzdoctemplatedocumentGetListV1

> EzdoctemplatedocumentGetListV1Response EzdoctemplatedocumentGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Ezdoctemplatedocument list



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
	resp, r, err := apiClient.ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzdoctemplatedocumentGetListV1`: EzdoctemplatedocumentGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzdoctemplatedocumentGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**EzdoctemplatedocumentGetListV1Response**](EzdoctemplatedocumentGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzdoctemplatedocumentGetObjectV2

> EzdoctemplatedocumentGetObjectV2Response EzdoctemplatedocumentGetObjectV2(ctx, pkiEzdoctemplatedocumentID).Execute()

Retrieve an existing Ezdoctemplatedocument



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
	pkiEzdoctemplatedocumentID := int32(56) // int32 | The unique ID of the Ezdoctemplatedocument

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentGetObjectV2(context.Background(), pkiEzdoctemplatedocumentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzdoctemplatedocumentGetObjectV2`: EzdoctemplatedocumentGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzdoctemplatedocumentID** | **int32** | The unique ID of the Ezdoctemplatedocument | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzdoctemplatedocumentGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzdoctemplatedocumentGetObjectV2Response**](EzdoctemplatedocumentGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzdoctemplatedocumentPatchObjectV1

> CommonResponse EzdoctemplatedocumentPatchObjectV1(ctx, pkiEzdoctemplatedocumentID).EzdoctemplatedocumentPatchObjectV1Request(ezdoctemplatedocumentPatchObjectV1Request).Execute()

Patch an existing Ezdoctemplatedocument



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
	pkiEzdoctemplatedocumentID := int32(56) // int32 | The unique ID of the Ezdoctemplatedocument
	ezdoctemplatedocumentPatchObjectV1Request := *openapiclient.NewEzdoctemplatedocumentPatchObjectV1Request(*openapiclient.NewEzdoctemplatedocumentRequestPatch()) // EzdoctemplatedocumentPatchObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentPatchObjectV1(context.Background(), pkiEzdoctemplatedocumentID).EzdoctemplatedocumentPatchObjectV1Request(ezdoctemplatedocumentPatchObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentPatchObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzdoctemplatedocumentPatchObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzdoctemplatedocumentAPI.EzdoctemplatedocumentPatchObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzdoctemplatedocumentID** | **int32** | The unique ID of the Ezdoctemplatedocument | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzdoctemplatedocumentPatchObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezdoctemplatedocumentPatchObjectV1Request** | [**EzdoctemplatedocumentPatchObjectV1Request**](EzdoctemplatedocumentPatchObjectV1Request.md) |  | 

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

