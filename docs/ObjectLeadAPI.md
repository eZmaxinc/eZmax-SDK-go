# eZmaxAPI\ObjectLeadAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadBatchDownloadV1**](ObjectLeadAPI.md#LeadBatchDownloadV1) | **Post** /1/object/lead/{pkiLeadID}/batchDownload | Download multiples attachments from a Lead
[**LeadGetAttachmentsV1**](ObjectLeadAPI.md#LeadGetAttachmentsV1) | **Get** /1/object/lead/{pkiLeadID}/getAttachments | Retrieve Lead&#39;s attachments
[**LeadGetCommunicationCountV1**](ObjectLeadAPI.md#LeadGetCommunicationCountV1) | **Get** /1/object/lead/{pkiLeadID}/getCommunicationCount | Retrieve Communication count
[**LeadGetCommunicationListV1**](ObjectLeadAPI.md#LeadGetCommunicationListV1) | **Get** /1/object/lead/{pkiLeadID}/getCommunicationList | Retrieve Communication list
[**LeadGetCommunicationrecipientsV1**](ObjectLeadAPI.md#LeadGetCommunicationrecipientsV1) | **Get** /1/object/lead/{pkiLeadID}/getCommunicationrecipients | Retrieve Communication recipients
[**LeadGetCommunicationsendersV1**](ObjectLeadAPI.md#LeadGetCommunicationsendersV1) | **Get** /1/object/lead/{pkiLeadID}/getCommunicationsenders | Retrieve Communication senders
[**LeadGetListV1**](ObjectLeadAPI.md#LeadGetListV1) | **Get** /1/object/lead/getList | Retrieve Lead list
[**LeadImportIntoEDMV1**](ObjectLeadAPI.md#LeadImportIntoEDMV1) | **Post** /1/object/lead/{pkiLeadID}/importIntoEDM | Import attachments into the Lead



## LeadBatchDownloadV1

> *os.File LeadBatchDownloadV1(ctx, pkiLeadID).LeadBatchDownloadV1Request(leadBatchDownloadV1Request).Execute()

Download multiples attachments from a Lead

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
	pkiLeadID := int32(56) // int32 | 
	leadBatchDownloadV1Request := *openapiclient.NewLeadBatchDownloadV1Request([]int32{int32(1)}) // LeadBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectLeadAPI.LeadBatchDownloadV1(context.Background(), pkiLeadID).LeadBatchDownloadV1Request(leadBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectLeadAPI.LeadBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LeadBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectLeadAPI.LeadBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiLeadID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeadBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **leadBatchDownloadV1Request** | [**LeadBatchDownloadV1Request**](LeadBatchDownloadV1Request.md) |  | 

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


## LeadGetAttachmentsV1

> LeadGetAttachmentsV1Response LeadGetAttachmentsV1(ctx, pkiLeadID).Execute()

Retrieve Lead's attachments

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
	pkiLeadID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectLeadAPI.LeadGetAttachmentsV1(context.Background(), pkiLeadID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectLeadAPI.LeadGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LeadGetAttachmentsV1`: LeadGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectLeadAPI.LeadGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiLeadID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeadGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LeadGetAttachmentsV1Response**](LeadGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeadGetCommunicationCountV1

> LeadGetCommunicationCountV1Response LeadGetCommunicationCountV1(ctx, pkiLeadID).Execute()

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
	pkiLeadID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectLeadAPI.LeadGetCommunicationCountV1(context.Background(), pkiLeadID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectLeadAPI.LeadGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LeadGetCommunicationCountV1`: LeadGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectLeadAPI.LeadGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiLeadID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeadGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LeadGetCommunicationCountV1Response**](LeadGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeadGetCommunicationListV1

> LeadGetCommunicationListV1Response LeadGetCommunicationListV1(ctx, pkiLeadID).Execute()

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
	pkiLeadID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectLeadAPI.LeadGetCommunicationListV1(context.Background(), pkiLeadID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectLeadAPI.LeadGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LeadGetCommunicationListV1`: LeadGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectLeadAPI.LeadGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiLeadID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeadGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LeadGetCommunicationListV1Response**](LeadGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeadGetCommunicationrecipientsV1

> LeadGetCommunicationrecipientsV1Response LeadGetCommunicationrecipientsV1(ctx, pkiLeadID).Execute()

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
	pkiLeadID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectLeadAPI.LeadGetCommunicationrecipientsV1(context.Background(), pkiLeadID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectLeadAPI.LeadGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LeadGetCommunicationrecipientsV1`: LeadGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectLeadAPI.LeadGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiLeadID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeadGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LeadGetCommunicationrecipientsV1Response**](LeadGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeadGetCommunicationsendersV1

> LeadGetCommunicationsendersV1Response LeadGetCommunicationsendersV1(ctx, pkiLeadID).Execute()

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
	pkiLeadID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectLeadAPI.LeadGetCommunicationsendersV1(context.Background(), pkiLeadID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectLeadAPI.LeadGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LeadGetCommunicationsendersV1`: LeadGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectLeadAPI.LeadGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiLeadID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeadGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LeadGetCommunicationsendersV1Response**](LeadGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeadGetListV1

> LeadGetListV1Response LeadGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Lead list



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
	resp, r, err := apiClient.ObjectLeadAPI.LeadGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectLeadAPI.LeadGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LeadGetListV1`: LeadGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectLeadAPI.LeadGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLeadGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**LeadGetListV1Response**](LeadGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeadImportIntoEDMV1

> LeadImportIntoEDMV1Response LeadImportIntoEDMV1(ctx, pkiLeadID).LeadImportIntoEDMV1Request(leadImportIntoEDMV1Request).Execute()

Import attachments into the Lead



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
	pkiLeadID := int32(56) // int32 | 
	leadImportIntoEDMV1Request := *openapiclient.NewLeadImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // LeadImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectLeadAPI.LeadImportIntoEDMV1(context.Background(), pkiLeadID).LeadImportIntoEDMV1Request(leadImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectLeadAPI.LeadImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LeadImportIntoEDMV1`: LeadImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectLeadAPI.LeadImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiLeadID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeadImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **leadImportIntoEDMV1Request** | [**LeadImportIntoEDMV1Request**](LeadImportIntoEDMV1Request.md) |  | 

### Return type

[**LeadImportIntoEDMV1Response**](LeadImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

