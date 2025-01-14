# eZmaxAPI\ObjectDomainAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainCreateObjectV1**](ObjectDomainAPI.md#DomainCreateObjectV1) | **Post** /1/object/domain | Create a new Domain
[**DomainDeleteObjectV1**](ObjectDomainAPI.md#DomainDeleteObjectV1) | **Delete** /1/object/domain/{pkiDomainID} | Delete an existing Domain
[**DomainGetListV1**](ObjectDomainAPI.md#DomainGetListV1) | **Get** /1/object/domain/getList | Retrieve Domain list
[**DomainGetObjectV2**](ObjectDomainAPI.md#DomainGetObjectV2) | **Get** /2/object/domain/{pkiDomainID} | Retrieve an existing Domain



## DomainCreateObjectV1

> DomainCreateObjectV1Response DomainCreateObjectV1(ctx).DomainCreateObjectV1Request(domainCreateObjectV1Request).Execute()

Create a new Domain



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
	domainCreateObjectV1Request := *openapiclient.NewDomainCreateObjectV1Request([]openapiclient.DomainRequestCompound{*openapiclient.NewDomainRequestCompound("ezsign.ca")}) // DomainCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDomainAPI.DomainCreateObjectV1(context.Background()).DomainCreateObjectV1Request(domainCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDomainAPI.DomainCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainCreateObjectV1`: DomainCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDomainAPI.DomainCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainCreateObjectV1Request** | [**DomainCreateObjectV1Request**](DomainCreateObjectV1Request.md) |  | 

### Return type

[**DomainCreateObjectV1Response**](DomainCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDeleteObjectV1

> DomainDeleteObjectV1Response DomainDeleteObjectV1(ctx, pkiDomainID).Execute()

Delete an existing Domain



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
	pkiDomainID := int32(56) // int32 | The unique ID of the Domain

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDomainAPI.DomainDeleteObjectV1(context.Background(), pkiDomainID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDomainAPI.DomainDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainDeleteObjectV1`: DomainDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDomainAPI.DomainDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDomainID** | **int32** | The unique ID of the Domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainDeleteObjectV1Response**](DomainDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainGetListV1

> DomainGetListV1Response DomainGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Domain list



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
	resp, r, err := apiClient.ObjectDomainAPI.DomainGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDomainAPI.DomainGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainGetListV1`: DomainGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDomainAPI.DomainGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**DomainGetListV1Response**](DomainGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainGetObjectV2

> DomainGetObjectV2Response DomainGetObjectV2(ctx, pkiDomainID).Execute()

Retrieve an existing Domain



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
	pkiDomainID := int32(56) // int32 | The unique ID of the Domain

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDomainAPI.DomainGetObjectV2(context.Background(), pkiDomainID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDomainAPI.DomainGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainGetObjectV2`: DomainGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDomainAPI.DomainGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDomainID** | **int32** | The unique ID of the Domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainGetObjectV2Response**](DomainGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

