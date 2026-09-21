# eZmaxAPI\ObjectNotaryAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotaryBatchDownloadV1**](ObjectNotaryAPI.md#NotaryBatchDownloadV1) | **Post** /1/object/notary/{pkiNotaryID}/batchDownload | Download multiples attachments from a Notary
[**NotaryGetAttachmentsV1**](ObjectNotaryAPI.md#NotaryGetAttachmentsV1) | **Get** /1/object/notary/{pkiNotaryID}/getAttachments | Retrieve Notary&#39;s attachments
[**NotaryGetCommunicationCountV1**](ObjectNotaryAPI.md#NotaryGetCommunicationCountV1) | **Get** /1/object/notary/{pkiNotaryID}/getCommunicationCount | Retrieve Communication count
[**NotaryGetCommunicationListV1**](ObjectNotaryAPI.md#NotaryGetCommunicationListV1) | **Get** /1/object/notary/{pkiNotaryID}/getCommunicationList | Retrieve Communication list
[**NotaryGetCommunicationrecipientsV1**](ObjectNotaryAPI.md#NotaryGetCommunicationrecipientsV1) | **Get** /1/object/notary/{pkiNotaryID}/getCommunicationrecipients | Retrieve Communication recipients
[**NotaryGetCommunicationsendersV1**](ObjectNotaryAPI.md#NotaryGetCommunicationsendersV1) | **Get** /1/object/notary/{pkiNotaryID}/getCommunicationsenders | Retrieve Communication senders
[**NotaryGetListV1**](ObjectNotaryAPI.md#NotaryGetListV1) | **Get** /1/object/notary/getList | Retrieve Notary list
[**NotaryImportIntoEDMV1**](ObjectNotaryAPI.md#NotaryImportIntoEDMV1) | **Post** /1/object/notary/{pkiNotaryID}/importIntoEDM | Import attachments into the Notary



## NotaryBatchDownloadV1

> *os.File NotaryBatchDownloadV1(ctx, pkiNotaryID).NotaryBatchDownloadV1Request(notaryBatchDownloadV1Request).Execute()

Download multiples attachments from a Notary

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
	pkiNotaryID := int32(56) // int32 | 
	notaryBatchDownloadV1Request := *openapiclient.NewNotaryBatchDownloadV1Request([]int32{int32(1)}) // NotaryBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectNotaryAPI.NotaryBatchDownloadV1(context.Background(), pkiNotaryID).NotaryBatchDownloadV1Request(notaryBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectNotaryAPI.NotaryBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotaryBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectNotaryAPI.NotaryBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiNotaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotaryBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notaryBatchDownloadV1Request** | [**NotaryBatchDownloadV1Request**](NotaryBatchDownloadV1Request.md) |  | 

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


## NotaryGetAttachmentsV1

> NotaryGetAttachmentsV1Response NotaryGetAttachmentsV1(ctx, pkiNotaryID).Execute()

Retrieve Notary's attachments

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
	pkiNotaryID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectNotaryAPI.NotaryGetAttachmentsV1(context.Background(), pkiNotaryID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectNotaryAPI.NotaryGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotaryGetAttachmentsV1`: NotaryGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectNotaryAPI.NotaryGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiNotaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotaryGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotaryGetAttachmentsV1Response**](NotaryGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotaryGetCommunicationCountV1

> NotaryGetCommunicationCountV1Response NotaryGetCommunicationCountV1(ctx, pkiNotaryID).Execute()

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
	pkiNotaryID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectNotaryAPI.NotaryGetCommunicationCountV1(context.Background(), pkiNotaryID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectNotaryAPI.NotaryGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotaryGetCommunicationCountV1`: NotaryGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectNotaryAPI.NotaryGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiNotaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotaryGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotaryGetCommunicationCountV1Response**](NotaryGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotaryGetCommunicationListV1

> NotaryGetCommunicationListV1Response NotaryGetCommunicationListV1(ctx, pkiNotaryID).Execute()

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
	pkiNotaryID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectNotaryAPI.NotaryGetCommunicationListV1(context.Background(), pkiNotaryID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectNotaryAPI.NotaryGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotaryGetCommunicationListV1`: NotaryGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectNotaryAPI.NotaryGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiNotaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotaryGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotaryGetCommunicationListV1Response**](NotaryGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotaryGetCommunicationrecipientsV1

> NotaryGetCommunicationrecipientsV1Response NotaryGetCommunicationrecipientsV1(ctx, pkiNotaryID).Execute()

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
	pkiNotaryID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectNotaryAPI.NotaryGetCommunicationrecipientsV1(context.Background(), pkiNotaryID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectNotaryAPI.NotaryGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotaryGetCommunicationrecipientsV1`: NotaryGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectNotaryAPI.NotaryGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiNotaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotaryGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotaryGetCommunicationrecipientsV1Response**](NotaryGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotaryGetCommunicationsendersV1

> NotaryGetCommunicationsendersV1Response NotaryGetCommunicationsendersV1(ctx, pkiNotaryID).Execute()

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
	pkiNotaryID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectNotaryAPI.NotaryGetCommunicationsendersV1(context.Background(), pkiNotaryID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectNotaryAPI.NotaryGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotaryGetCommunicationsendersV1`: NotaryGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectNotaryAPI.NotaryGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiNotaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotaryGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotaryGetCommunicationsendersV1Response**](NotaryGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotaryGetListV1

> NotaryGetListV1Response NotaryGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Notary list



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
	resp, r, err := apiClient.ObjectNotaryAPI.NotaryGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectNotaryAPI.NotaryGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotaryGetListV1`: NotaryGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectNotaryAPI.NotaryGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotaryGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**NotaryGetListV1Response**](NotaryGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotaryImportIntoEDMV1

> NotaryImportIntoEDMV1Response NotaryImportIntoEDMV1(ctx, pkiNotaryID).NotaryImportIntoEDMV1Request(notaryImportIntoEDMV1Request).Execute()

Import attachments into the Notary

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
	pkiNotaryID := int32(56) // int32 | 
	notaryImportIntoEDMV1Request := *openapiclient.NewNotaryImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // NotaryImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectNotaryAPI.NotaryImportIntoEDMV1(context.Background(), pkiNotaryID).NotaryImportIntoEDMV1Request(notaryImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectNotaryAPI.NotaryImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotaryImportIntoEDMV1`: NotaryImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectNotaryAPI.NotaryImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiNotaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotaryImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notaryImportIntoEDMV1Request** | [**NotaryImportIntoEDMV1Request**](NotaryImportIntoEDMV1Request.md) |  | 

### Return type

[**NotaryImportIntoEDMV1Response**](NotaryImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

