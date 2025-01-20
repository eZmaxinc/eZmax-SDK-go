# eZmaxAPI\ObjectEzsignbulksendAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignbulksendCreateEzsignbulksendtransmissionV2**](ObjectEzsignbulksendAPI.md#EzsignbulksendCreateEzsignbulksendtransmissionV2) | **Post** /2/object/ezsignbulksend/{pkiEzsignbulksendID}/createEzsignbulksendtransmission | Create a new Ezsignbulksendtransmission in the Ezsignbulksend
[**EzsignbulksendCreateObjectV1**](ObjectEzsignbulksendAPI.md#EzsignbulksendCreateObjectV1) | **Post** /1/object/ezsignbulksend | Create a new Ezsignbulksend
[**EzsignbulksendDeleteObjectV1**](ObjectEzsignbulksendAPI.md#EzsignbulksendDeleteObjectV1) | **Delete** /1/object/ezsignbulksend/{pkiEzsignbulksendID} | Delete an existing Ezsignbulksend
[**EzsignbulksendEditObjectV1**](ObjectEzsignbulksendAPI.md#EzsignbulksendEditObjectV1) | **Put** /1/object/ezsignbulksend/{pkiEzsignbulksendID} | Edit an existing Ezsignbulksend
[**EzsignbulksendGetCsvTemplateV1**](ObjectEzsignbulksendAPI.md#EzsignbulksendGetCsvTemplateV1) | **Get** /1/object/ezsignbulksend/{pkiEzsignbulksendID}/getCsvTemplate | Retrieve an existing Ezsignbulksend&#39;s empty Csv template
[**EzsignbulksendGetEzsignbulksendtransmissionsV1**](ObjectEzsignbulksendAPI.md#EzsignbulksendGetEzsignbulksendtransmissionsV1) | **Get** /1/object/ezsignbulksend/{pkiEzsignbulksendID}/getEzsignbulksendtransmissions | Retrieve an existing Ezsignbulksend&#39;s Ezsignbulksendtransmissions
[**EzsignbulksendGetEzsignsignaturesAutomaticV1**](ObjectEzsignbulksendAPI.md#EzsignbulksendGetEzsignsignaturesAutomaticV1) | **Get** /1/object/ezsignbulksend/{pkiEzsignbulksendID}/getEzsignsignaturesAutomatic | Retrieve an existing Ezsignbulksend&#39;s automatic Ezsignsignatures
[**EzsignbulksendGetFormsDataV1**](ObjectEzsignbulksendAPI.md#EzsignbulksendGetFormsDataV1) | **Get** /1/object/ezsignbulksend/{pkiEzsignbulksendID}/getFormsData | Retrieve an existing Ezsignbulksend&#39;s forms data
[**EzsignbulksendGetListV1**](ObjectEzsignbulksendAPI.md#EzsignbulksendGetListV1) | **Get** /1/object/ezsignbulksend/getList | Retrieve Ezsignbulksend list
[**EzsignbulksendGetObjectV2**](ObjectEzsignbulksendAPI.md#EzsignbulksendGetObjectV2) | **Get** /2/object/ezsignbulksend/{pkiEzsignbulksendID} | Retrieve an existing Ezsignbulksend
[**EzsignbulksendReorderV1**](ObjectEzsignbulksendAPI.md#EzsignbulksendReorderV1) | **Post** /1/object/ezsignbulksend/{pkiEzsignbulksendID}/reorder | Reorder Ezsignbulksenddocumentmappings in the Ezsignbulksend



## EzsignbulksendCreateEzsignbulksendtransmissionV2

> EzsignbulksendCreateEzsignbulksendtransmissionV2Response EzsignbulksendCreateEzsignbulksendtransmissionV2(ctx, pkiEzsignbulksendID).EzsignbulksendCreateEzsignbulksendtransmissionV2Request(ezsignbulksendCreateEzsignbulksendtransmissionV2Request).Execute()

Create a new Ezsignbulksendtransmission in the Ezsignbulksend

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
	pkiEzsignbulksendID := int32(56) // int32 | 
	ezsignbulksendCreateEzsignbulksendtransmissionV2Request := *openapiclient.NewEzsignbulksendCreateEzsignbulksendtransmissionV2Request(int32(2), "Test eZsign Bulk Send Transmission #1", "2020-12-31 23:59:59", int32(30), int32(30), "Hi John,

This is the document I need you to review.

Could you sign it before Monday please.

Best Regards.

Mary", string(123)) // EzsignbulksendCreateEzsignbulksendtransmissionV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignbulksendAPI.EzsignbulksendCreateEzsignbulksendtransmissionV2(context.Background(), pkiEzsignbulksendID).EzsignbulksendCreateEzsignbulksendtransmissionV2Request(ezsignbulksendCreateEzsignbulksendtransmissionV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendAPI.EzsignbulksendCreateEzsignbulksendtransmissionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignbulksendCreateEzsignbulksendtransmissionV2`: EzsignbulksendCreateEzsignbulksendtransmissionV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendAPI.EzsignbulksendCreateEzsignbulksendtransmissionV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendCreateEzsignbulksendtransmissionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignbulksendCreateEzsignbulksendtransmissionV2Request** | [**EzsignbulksendCreateEzsignbulksendtransmissionV2Request**](EzsignbulksendCreateEzsignbulksendtransmissionV2Request.md) |  | 

### Return type

[**EzsignbulksendCreateEzsignbulksendtransmissionV2Response**](EzsignbulksendCreateEzsignbulksendtransmissionV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendCreateObjectV1

> EzsignbulksendCreateObjectV1Response EzsignbulksendCreateObjectV1(ctx).EzsignbulksendCreateObjectV1Request(ezsignbulksendCreateObjectV1Request).Execute()

Create a new Ezsignbulksend



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
	ezsignbulksendCreateObjectV1Request := *openapiclient.NewEzsignbulksendCreateObjectV1Request([]openapiclient.EzsignbulksendRequestCompound{*openapiclient.NewEzsignbulksendRequestCompound(int32(5), int32(2), "Test eZsign Bulk Send", "This is a note", false, true)}) // EzsignbulksendCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignbulksendAPI.EzsignbulksendCreateObjectV1(context.Background()).EzsignbulksendCreateObjectV1Request(ezsignbulksendCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendAPI.EzsignbulksendCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignbulksendCreateObjectV1`: EzsignbulksendCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendAPI.EzsignbulksendCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignbulksendCreateObjectV1Request** | [**EzsignbulksendCreateObjectV1Request**](EzsignbulksendCreateObjectV1Request.md) |  | 

### Return type

[**EzsignbulksendCreateObjectV1Response**](EzsignbulksendCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendDeleteObjectV1

> EzsignbulksendDeleteObjectV1Response EzsignbulksendDeleteObjectV1(ctx, pkiEzsignbulksendID).Execute()

Delete an existing Ezsignbulksend



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
	pkiEzsignbulksendID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignbulksendAPI.EzsignbulksendDeleteObjectV1(context.Background(), pkiEzsignbulksendID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendAPI.EzsignbulksendDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignbulksendDeleteObjectV1`: EzsignbulksendDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendAPI.EzsignbulksendDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignbulksendDeleteObjectV1Response**](EzsignbulksendDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendEditObjectV1

> EzsignbulksendEditObjectV1Response EzsignbulksendEditObjectV1(ctx, pkiEzsignbulksendID).EzsignbulksendEditObjectV1Request(ezsignbulksendEditObjectV1Request).Execute()

Edit an existing Ezsignbulksend



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
	pkiEzsignbulksendID := int32(56) // int32 | 
	ezsignbulksendEditObjectV1Request := *openapiclient.NewEzsignbulksendEditObjectV1Request(*openapiclient.NewEzsignbulksendRequestCompound(int32(5), int32(2), "Test eZsign Bulk Send", "This is a note", false, true)) // EzsignbulksendEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignbulksendAPI.EzsignbulksendEditObjectV1(context.Background(), pkiEzsignbulksendID).EzsignbulksendEditObjectV1Request(ezsignbulksendEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendAPI.EzsignbulksendEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignbulksendEditObjectV1`: EzsignbulksendEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendAPI.EzsignbulksendEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignbulksendEditObjectV1Request** | [**EzsignbulksendEditObjectV1Request**](EzsignbulksendEditObjectV1Request.md) |  | 

### Return type

[**EzsignbulksendEditObjectV1Response**](EzsignbulksendEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendGetCsvTemplateV1

> string EzsignbulksendGetCsvTemplateV1(ctx, pkiEzsignbulksendID).ECsvSeparator(eCsvSeparator).Execute()

Retrieve an existing Ezsignbulksend's empty Csv template



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
	pkiEzsignbulksendID := int32(56) // int32 | 
	eCsvSeparator := "eCsvSeparator_example" // string | Separator that will be used to separate fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignbulksendAPI.EzsignbulksendGetCsvTemplateV1(context.Background(), pkiEzsignbulksendID).ECsvSeparator(eCsvSeparator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendAPI.EzsignbulksendGetCsvTemplateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignbulksendGetCsvTemplateV1`: string
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendAPI.EzsignbulksendGetCsvTemplateV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendGetCsvTemplateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eCsvSeparator** | **string** | Separator that will be used to separate fields | 

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendGetEzsignbulksendtransmissionsV1

> EzsignbulksendGetEzsignbulksendtransmissionsV1Response EzsignbulksendGetEzsignbulksendtransmissionsV1(ctx, pkiEzsignbulksendID).Execute()

Retrieve an existing Ezsignbulksend's Ezsignbulksendtransmissions



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
	pkiEzsignbulksendID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignbulksendAPI.EzsignbulksendGetEzsignbulksendtransmissionsV1(context.Background(), pkiEzsignbulksendID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendAPI.EzsignbulksendGetEzsignbulksendtransmissionsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignbulksendGetEzsignbulksendtransmissionsV1`: EzsignbulksendGetEzsignbulksendtransmissionsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendAPI.EzsignbulksendGetEzsignbulksendtransmissionsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendGetEzsignbulksendtransmissionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignbulksendGetEzsignbulksendtransmissionsV1Response**](EzsignbulksendGetEzsignbulksendtransmissionsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendGetEzsignsignaturesAutomaticV1

> EzsignbulksendGetEzsignsignaturesAutomaticV1Response EzsignbulksendGetEzsignsignaturesAutomaticV1(ctx, pkiEzsignbulksendID).Execute()

Retrieve an existing Ezsignbulksend's automatic Ezsignsignatures



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
	pkiEzsignbulksendID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignbulksendAPI.EzsignbulksendGetEzsignsignaturesAutomaticV1(context.Background(), pkiEzsignbulksendID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendAPI.EzsignbulksendGetEzsignsignaturesAutomaticV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignbulksendGetEzsignsignaturesAutomaticV1`: EzsignbulksendGetEzsignsignaturesAutomaticV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendAPI.EzsignbulksendGetEzsignsignaturesAutomaticV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendGetEzsignsignaturesAutomaticV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignbulksendGetEzsignsignaturesAutomaticV1Response**](EzsignbulksendGetEzsignsignaturesAutomaticV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendGetFormsDataV1

> EzsignbulksendGetFormsDataV1Response EzsignbulksendGetFormsDataV1(ctx, pkiEzsignbulksendID).Execute()

Retrieve an existing Ezsignbulksend's forms data



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
	pkiEzsignbulksendID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignbulksendAPI.EzsignbulksendGetFormsDataV1(context.Background(), pkiEzsignbulksendID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendAPI.EzsignbulksendGetFormsDataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignbulksendGetFormsDataV1`: EzsignbulksendGetFormsDataV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendAPI.EzsignbulksendGetFormsDataV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendGetFormsDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignbulksendGetFormsDataV1Response**](EzsignbulksendGetFormsDataV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendGetListV1

> EzsignbulksendGetListV1Response EzsignbulksendGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Ezsignbulksend list



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
	resp, r, err := apiClient.ObjectEzsignbulksendAPI.EzsignbulksendGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendAPI.EzsignbulksendGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignbulksendGetListV1`: EzsignbulksendGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendAPI.EzsignbulksendGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**EzsignbulksendGetListV1Response**](EzsignbulksendGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendGetObjectV2

> EzsignbulksendGetObjectV2Response EzsignbulksendGetObjectV2(ctx, pkiEzsignbulksendID).Execute()

Retrieve an existing Ezsignbulksend



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
	pkiEzsignbulksendID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignbulksendAPI.EzsignbulksendGetObjectV2(context.Background(), pkiEzsignbulksendID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendAPI.EzsignbulksendGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignbulksendGetObjectV2`: EzsignbulksendGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendAPI.EzsignbulksendGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignbulksendGetObjectV2Response**](EzsignbulksendGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendReorderV1

> EzsignbulksendReorderV1Response EzsignbulksendReorderV1(ctx, pkiEzsignbulksendID).EzsignbulksendReorderV1Request(ezsignbulksendReorderV1Request).Execute()

Reorder Ezsignbulksenddocumentmappings in the Ezsignbulksend

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
	pkiEzsignbulksendID := int32(56) // int32 | 
	ezsignbulksendReorderV1Request := *openapiclient.NewEzsignbulksendReorderV1Request([]int32{int32(48)}) // EzsignbulksendReorderV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignbulksendAPI.EzsignbulksendReorderV1(context.Background(), pkiEzsignbulksendID).EzsignbulksendReorderV1Request(ezsignbulksendReorderV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendAPI.EzsignbulksendReorderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignbulksendReorderV1`: EzsignbulksendReorderV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendAPI.EzsignbulksendReorderV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendReorderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignbulksendReorderV1Request** | [**EzsignbulksendReorderV1Request**](EzsignbulksendReorderV1Request.md) |  | 

### Return type

[**EzsignbulksendReorderV1Response**](EzsignbulksendReorderV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

