# eZmaxAPI\ObjectSalaryAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SalaryBatchDownloadV1**](ObjectSalaryAPI.md#SalaryBatchDownloadV1) | **Post** /1/object/salary/{pkiSalaryID}/batchDownload | Download multiples attachments from a Reconciliation
[**SalaryGetAttachmentsV1**](ObjectSalaryAPI.md#SalaryGetAttachmentsV1) | **Get** /1/object/salary/{pkiSalaryID}/getAttachments | Retrieve Salary&#39;s attachments
[**SalaryGetCommunicationCountV1**](ObjectSalaryAPI.md#SalaryGetCommunicationCountV1) | **Get** /1/object/salary/{pkiSalaryID}/getCommunicationCount | Retrieve Communication count
[**SalaryGetCommunicationListV1**](ObjectSalaryAPI.md#SalaryGetCommunicationListV1) | **Get** /1/object/salary/{pkiSalaryID}/getCommunicationList | Retrieve Communication list
[**SalaryGetCommunicationrecipientsV1**](ObjectSalaryAPI.md#SalaryGetCommunicationrecipientsV1) | **Get** /1/object/salary/{pkiSalaryID}/getCommunicationrecipients | Retrieve Communication recipients
[**SalaryGetCommunicationsendersV1**](ObjectSalaryAPI.md#SalaryGetCommunicationsendersV1) | **Get** /1/object/salary/{pkiSalaryID}/getCommunicationsenders | Retrieve Communication senders
[**SalaryGetListV1**](ObjectSalaryAPI.md#SalaryGetListV1) | **Get** /1/object/salary/getList | Retrieve Salary list
[**SalaryImportIntoEDMV1**](ObjectSalaryAPI.md#SalaryImportIntoEDMV1) | **Post** /1/object/salary/{pkiSalaryID}/importIntoEDM | Import attachments into the Salary



## SalaryBatchDownloadV1

> *os.File SalaryBatchDownloadV1(ctx, pkiSalaryID).SalaryBatchDownloadV1Request(salaryBatchDownloadV1Request).Execute()

Download multiples attachments from a Reconciliation

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
	pkiSalaryID := int32(56) // int32 | 
	salaryBatchDownloadV1Request := *openapiclient.NewSalaryBatchDownloadV1Request([]int32{int32(1)}) // SalaryBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSalaryAPI.SalaryBatchDownloadV1(context.Background(), pkiSalaryID).SalaryBatchDownloadV1Request(salaryBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSalaryAPI.SalaryBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalaryBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectSalaryAPI.SalaryBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSalaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSalaryBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **salaryBatchDownloadV1Request** | [**SalaryBatchDownloadV1Request**](SalaryBatchDownloadV1Request.md) |  | 

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


## SalaryGetAttachmentsV1

> SalaryGetAttachmentsV1Response SalaryGetAttachmentsV1(ctx, pkiSalaryID).Execute()

Retrieve Salary's attachments

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
	pkiSalaryID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSalaryAPI.SalaryGetAttachmentsV1(context.Background(), pkiSalaryID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSalaryAPI.SalaryGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalaryGetAttachmentsV1`: SalaryGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSalaryAPI.SalaryGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSalaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSalaryGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalaryGetAttachmentsV1Response**](SalaryGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SalaryGetCommunicationCountV1

> SalaryGetCommunicationCountV1Response SalaryGetCommunicationCountV1(ctx, pkiSalaryID).Execute()

Retrieve Communication count

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
	pkiSalaryID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSalaryAPI.SalaryGetCommunicationCountV1(context.Background(), pkiSalaryID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSalaryAPI.SalaryGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalaryGetCommunicationCountV1`: SalaryGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSalaryAPI.SalaryGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSalaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSalaryGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalaryGetCommunicationCountV1Response**](SalaryGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SalaryGetCommunicationListV1

> SalaryGetCommunicationListV1Response SalaryGetCommunicationListV1(ctx, pkiSalaryID).Execute()

Retrieve Communication list

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
	pkiSalaryID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSalaryAPI.SalaryGetCommunicationListV1(context.Background(), pkiSalaryID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSalaryAPI.SalaryGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalaryGetCommunicationListV1`: SalaryGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSalaryAPI.SalaryGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSalaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSalaryGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalaryGetCommunicationListV1Response**](SalaryGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SalaryGetCommunicationrecipientsV1

> SalaryGetCommunicationrecipientsV1Response SalaryGetCommunicationrecipientsV1(ctx, pkiSalaryID).Execute()

Retrieve Communication recipients

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
	pkiSalaryID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSalaryAPI.SalaryGetCommunicationrecipientsV1(context.Background(), pkiSalaryID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSalaryAPI.SalaryGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalaryGetCommunicationrecipientsV1`: SalaryGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSalaryAPI.SalaryGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSalaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSalaryGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalaryGetCommunicationrecipientsV1Response**](SalaryGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SalaryGetCommunicationsendersV1

> SalaryGetCommunicationsendersV1Response SalaryGetCommunicationsendersV1(ctx, pkiSalaryID).Execute()

Retrieve Communication senders

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
	pkiSalaryID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSalaryAPI.SalaryGetCommunicationsendersV1(context.Background(), pkiSalaryID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSalaryAPI.SalaryGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalaryGetCommunicationsendersV1`: SalaryGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSalaryAPI.SalaryGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSalaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSalaryGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalaryGetCommunicationsendersV1Response**](SalaryGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SalaryGetListV1

> SalaryGetListV1Response SalaryGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Salary list



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
	resp, r, err := apiClient.ObjectSalaryAPI.SalaryGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSalaryAPI.SalaryGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalaryGetListV1`: SalaryGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSalaryAPI.SalaryGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSalaryGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**SalaryGetListV1Response**](SalaryGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SalaryImportIntoEDMV1

> SalaryImportIntoEDMV1Response SalaryImportIntoEDMV1(ctx, pkiSalaryID).SalaryImportIntoEDMV1Request(salaryImportIntoEDMV1Request).Execute()

Import attachments into the Salary

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
	pkiSalaryID := int32(56) // int32 | 
	salaryImportIntoEDMV1Request := *openapiclient.NewSalaryImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // SalaryImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSalaryAPI.SalaryImportIntoEDMV1(context.Background(), pkiSalaryID).SalaryImportIntoEDMV1Request(salaryImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSalaryAPI.SalaryImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalaryImportIntoEDMV1`: SalaryImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSalaryAPI.SalaryImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSalaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSalaryImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **salaryImportIntoEDMV1Request** | [**SalaryImportIntoEDMV1Request**](SalaryImportIntoEDMV1Request.md) |  | 

### Return type

[**SalaryImportIntoEDMV1Response**](SalaryImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

