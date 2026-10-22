# eZmaxAPI\ObjectDocumenttypechecklistAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocumenttypechecklistCreateObjectV1**](ObjectDocumenttypechecklistAPI.md#DocumenttypechecklistCreateObjectV1) | **Post** /1/object/documenttypechecklist | Create a new Documenttypechecklist
[**DocumenttypechecklistDeleteObjectV1**](ObjectDocumenttypechecklistAPI.md#DocumenttypechecklistDeleteObjectV1) | **Delete** /1/object/documenttypechecklist/{pkiDocumenttypechecklistID} | Delete an existing Documenttypechecklist
[**DocumenttypechecklistEditObjectV1**](ObjectDocumenttypechecklistAPI.md#DocumenttypechecklistEditObjectV1) | **Put** /1/object/documenttypechecklist/{pkiDocumenttypechecklistID} | Edit an existing Documenttypechecklist
[**DocumenttypechecklistGetListV1**](ObjectDocumenttypechecklistAPI.md#DocumenttypechecklistGetListV1) | **Get** /1/object/documenttypechecklist/getList | Retrieve Documenttypechecklist list
[**DocumenttypechecklistGetObjectV2**](ObjectDocumenttypechecklistAPI.md#DocumenttypechecklistGetObjectV2) | **Get** /2/object/documenttypechecklist/{pkiDocumenttypechecklistID} | Retrieve an existing Documenttypechecklist



## DocumenttypechecklistCreateObjectV1

> DocumenttypechecklistCreateObjectV1Response DocumenttypechecklistCreateObjectV1(ctx).DocumenttypechecklistCreateObjectV1Request(documenttypechecklistCreateObjectV1Request).Execute()

Create a new Documenttypechecklist



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
	documenttypechecklistCreateObjectV1Request := *openapiclient.NewDocumenttypechecklistCreateObjectV1Request([]openapiclient.DocumenttypechecklistRequestCompound{*openapiclient.NewDocumenttypechecklistRequestCompound(int32(172), openapiclient.Field-eDocumenttypechecklistStatus("Not required"), "This document is optional.")}) // DocumenttypechecklistCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDocumenttypechecklistAPI.DocumenttypechecklistCreateObjectV1(context.Background()).DocumenttypechecklistCreateObjectV1Request(documenttypechecklistCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDocumenttypechecklistAPI.DocumenttypechecklistCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumenttypechecklistCreateObjectV1`: DocumenttypechecklistCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDocumenttypechecklistAPI.DocumenttypechecklistCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumenttypechecklistCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documenttypechecklistCreateObjectV1Request** | [**DocumenttypechecklistCreateObjectV1Request**](DocumenttypechecklistCreateObjectV1Request.md) |  | 

### Return type

[**DocumenttypechecklistCreateObjectV1Response**](DocumenttypechecklistCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumenttypechecklistDeleteObjectV1

> DocumenttypechecklistDeleteObjectV1Response DocumenttypechecklistDeleteObjectV1(ctx, pkiDocumenttypechecklistID).Execute()

Delete an existing Documenttypechecklist



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
	pkiDocumenttypechecklistID := int32(56) // int32 | The unique ID of the Documenttypechecklist

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDocumenttypechecklistAPI.DocumenttypechecklistDeleteObjectV1(context.Background(), pkiDocumenttypechecklistID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDocumenttypechecklistAPI.DocumenttypechecklistDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumenttypechecklistDeleteObjectV1`: DocumenttypechecklistDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDocumenttypechecklistAPI.DocumenttypechecklistDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDocumenttypechecklistID** | **int32** | The unique ID of the Documenttypechecklist | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocumenttypechecklistDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DocumenttypechecklistDeleteObjectV1Response**](DocumenttypechecklistDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumenttypechecklistEditObjectV1

> DocumenttypechecklistEditObjectV1Response DocumenttypechecklistEditObjectV1(ctx, pkiDocumenttypechecklistID).DocumenttypechecklistEditObjectV1Request(documenttypechecklistEditObjectV1Request).Execute()

Edit an existing Documenttypechecklist



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
	pkiDocumenttypechecklistID := int32(56) // int32 | The unique ID of the Documenttypechecklist
	documenttypechecklistEditObjectV1Request := *openapiclient.NewDocumenttypechecklistEditObjectV1Request(*openapiclient.NewDocumenttypechecklistRequestCompound(int32(172), openapiclient.Field-eDocumenttypechecklistStatus("Not required"), "This document is optional.")) // DocumenttypechecklistEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDocumenttypechecklistAPI.DocumenttypechecklistEditObjectV1(context.Background(), pkiDocumenttypechecklistID).DocumenttypechecklistEditObjectV1Request(documenttypechecklistEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDocumenttypechecklistAPI.DocumenttypechecklistEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumenttypechecklistEditObjectV1`: DocumenttypechecklistEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDocumenttypechecklistAPI.DocumenttypechecklistEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDocumenttypechecklistID** | **int32** | The unique ID of the Documenttypechecklist | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocumenttypechecklistEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **documenttypechecklistEditObjectV1Request** | [**DocumenttypechecklistEditObjectV1Request**](DocumenttypechecklistEditObjectV1Request.md) |  | 

### Return type

[**DocumenttypechecklistEditObjectV1Response**](DocumenttypechecklistEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumenttypechecklistGetListV1

> DocumenttypechecklistGetListV1Response DocumenttypechecklistGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Documenttypechecklist list



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
	resp, r, err := apiClient.ObjectDocumenttypechecklistAPI.DocumenttypechecklistGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDocumenttypechecklistAPI.DocumenttypechecklistGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumenttypechecklistGetListV1`: DocumenttypechecklistGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDocumenttypechecklistAPI.DocumenttypechecklistGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumenttypechecklistGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**DocumenttypechecklistGetListV1Response**](DocumenttypechecklistGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumenttypechecklistGetObjectV2

> DocumenttypechecklistGetObjectV2Response DocumenttypechecklistGetObjectV2(ctx, pkiDocumenttypechecklistID).Execute()

Retrieve an existing Documenttypechecklist



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
	pkiDocumenttypechecklistID := int32(56) // int32 | The unique ID of the Documenttypechecklist

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDocumenttypechecklistAPI.DocumenttypechecklistGetObjectV2(context.Background(), pkiDocumenttypechecklistID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDocumenttypechecklistAPI.DocumenttypechecklistGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumenttypechecklistGetObjectV2`: DocumenttypechecklistGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDocumenttypechecklistAPI.DocumenttypechecklistGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDocumenttypechecklistID** | **int32** | The unique ID of the Documenttypechecklist | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocumenttypechecklistGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DocumenttypechecklistGetObjectV2Response**](DocumenttypechecklistGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

