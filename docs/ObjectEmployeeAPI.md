# eZmaxAPI\ObjectEmployeeAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmployeeBatchDownloadV1**](ObjectEmployeeAPI.md#EmployeeBatchDownloadV1) | **Post** /1/object/employee/{pkiEmployeeID}/batchDownload | Download multiples attachments from a Employee
[**EmployeeGetAttachmentsV1**](ObjectEmployeeAPI.md#EmployeeGetAttachmentsV1) | **Get** /1/object/employee/{pkiEmployeeID}/getAttachments | Retrieve Employee&#39;s attachments
[**EmployeeGetListV1**](ObjectEmployeeAPI.md#EmployeeGetListV1) | **Get** /1/object/employee/getList | Retrieve Employee list
[**EmployeeImportIntoEDMV1**](ObjectEmployeeAPI.md#EmployeeImportIntoEDMV1) | **Post** /1/object/employee/{pkiEmployeeID}/importIntoEDM | Import attachments into the Employee



## EmployeeBatchDownloadV1

> *os.File EmployeeBatchDownloadV1(ctx, pkiEmployeeID).EmployeeBatchDownloadV1Request(employeeBatchDownloadV1Request).Execute()

Download multiples attachments from a Employee

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
	pkiEmployeeID := int32(56) // int32 | 
	employeeBatchDownloadV1Request := *openapiclient.NewEmployeeBatchDownloadV1Request([]int32{int32(1)}) // EmployeeBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEmployeeAPI.EmployeeBatchDownloadV1(context.Background(), pkiEmployeeID).EmployeeBatchDownloadV1Request(employeeBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEmployeeAPI.EmployeeBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmployeeBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectEmployeeAPI.EmployeeBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEmployeeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmployeeBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **employeeBatchDownloadV1Request** | [**EmployeeBatchDownloadV1Request**](EmployeeBatchDownloadV1Request.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/zip, text/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmployeeGetAttachmentsV1

> EmployeeGetAttachmentsV1Response EmployeeGetAttachmentsV1(ctx, pkiEmployeeID).Execute()

Retrieve Employee's attachments

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
	pkiEmployeeID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEmployeeAPI.EmployeeGetAttachmentsV1(context.Background(), pkiEmployeeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEmployeeAPI.EmployeeGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmployeeGetAttachmentsV1`: EmployeeGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEmployeeAPI.EmployeeGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEmployeeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmployeeGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmployeeGetAttachmentsV1Response**](EmployeeGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmployeeGetListV1

> EmployeeGetListV1Response EmployeeGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Employee list



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
	resp, r, err := apiClient.ObjectEmployeeAPI.EmployeeGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEmployeeAPI.EmployeeGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmployeeGetListV1`: EmployeeGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEmployeeAPI.EmployeeGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmployeeGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**EmployeeGetListV1Response**](EmployeeGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmployeeImportIntoEDMV1

> EmployeeImportIntoEDMV1Response EmployeeImportIntoEDMV1(ctx, pkiEmployeeID).EmployeeImportIntoEDMV1Request(employeeImportIntoEDMV1Request).Execute()

Import attachments into the Employee



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
	pkiEmployeeID := int32(56) // int32 | 
	employeeImportIntoEDMV1Request := *openapiclient.NewEmployeeImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // EmployeeImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEmployeeAPI.EmployeeImportIntoEDMV1(context.Background(), pkiEmployeeID).EmployeeImportIntoEDMV1Request(employeeImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEmployeeAPI.EmployeeImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmployeeImportIntoEDMV1`: EmployeeImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEmployeeAPI.EmployeeImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEmployeeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmployeeImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **employeeImportIntoEDMV1Request** | [**EmployeeImportIntoEDMV1Request**](EmployeeImportIntoEDMV1Request.md) |  | 

### Return type

[**EmployeeImportIntoEDMV1Response**](EmployeeImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

