# eZmaxAPI\ObjectFolderAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FolderBatchDownloadV1**](ObjectFolderAPI.md#FolderBatchDownloadV1) | **Post** /1/object/folder/{pkiFolderID}/batchDownload | Download multiples attachments from an Folder
[**FolderGetAttachmentsV1**](ObjectFolderAPI.md#FolderGetAttachmentsV1) | **Get** /1/object/folder/{pkiFolderID}/getAttachments | Retrieve Folder&#39;s attachments
[**FolderGetCommunicationCountV1**](ObjectFolderAPI.md#FolderGetCommunicationCountV1) | **Get** /1/object/folder/{pkiFolderID}/getCommunicationCount | Retrieve Communication count
[**FolderGetCommunicationListV1**](ObjectFolderAPI.md#FolderGetCommunicationListV1) | **Get** /1/object/folder/{pkiFolderID}/getCommunicationList | Retrieve Communication list
[**FolderGetCommunicationrecipientsV1**](ObjectFolderAPI.md#FolderGetCommunicationrecipientsV1) | **Get** /1/object/folder/{pkiFolderID}/getCommunicationrecipients | Retrieve Communication recipients
[**FolderGetCommunicationsendersV1**](ObjectFolderAPI.md#FolderGetCommunicationsendersV1) | **Get** /1/object/folder/{pkiFolderID}/getCommunicationsenders | Retrieve Communication senders
[**FolderImportIntoEDMV1**](ObjectFolderAPI.md#FolderImportIntoEDMV1) | **Post** /1/object/folder/{pkiFolderID}/importIntoEDM | Import attachments into the Folder



## FolderBatchDownloadV1

> *os.File FolderBatchDownloadV1(ctx, pkiFolderID).FolderBatchDownloadV1Request(folderBatchDownloadV1Request).Execute()

Download multiples attachments from an Folder

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
	pkiFolderID := int32(56) // int32 | 
	folderBatchDownloadV1Request := *openapiclient.NewFolderBatchDownloadV1Request([]int32{int32(1)}) // FolderBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectFolderAPI.FolderBatchDownloadV1(context.Background(), pkiFolderID).FolderBatchDownloadV1Request(folderBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectFolderAPI.FolderBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FolderBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectFolderAPI.FolderBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiFolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFolderBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderBatchDownloadV1Request** | [**FolderBatchDownloadV1Request**](FolderBatchDownloadV1Request.md) |  | 

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


## FolderGetAttachmentsV1

> FolderGetAttachmentsV1Response FolderGetAttachmentsV1(ctx, pkiFolderID).Execute()

Retrieve Folder's attachments

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
	pkiFolderID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectFolderAPI.FolderGetAttachmentsV1(context.Background(), pkiFolderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectFolderAPI.FolderGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FolderGetAttachmentsV1`: FolderGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectFolderAPI.FolderGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiFolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFolderGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderGetAttachmentsV1Response**](FolderGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FolderGetCommunicationCountV1

> FolderGetCommunicationCountV1Response FolderGetCommunicationCountV1(ctx, pkiFolderID).Execute()

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
	pkiFolderID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectFolderAPI.FolderGetCommunicationCountV1(context.Background(), pkiFolderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectFolderAPI.FolderGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FolderGetCommunicationCountV1`: FolderGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectFolderAPI.FolderGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiFolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFolderGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderGetCommunicationCountV1Response**](FolderGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FolderGetCommunicationListV1

> FolderGetCommunicationListV1Response FolderGetCommunicationListV1(ctx, pkiFolderID).Execute()

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
	pkiFolderID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectFolderAPI.FolderGetCommunicationListV1(context.Background(), pkiFolderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectFolderAPI.FolderGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FolderGetCommunicationListV1`: FolderGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectFolderAPI.FolderGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiFolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFolderGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderGetCommunicationListV1Response**](FolderGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FolderGetCommunicationrecipientsV1

> FolderGetCommunicationrecipientsV1Response FolderGetCommunicationrecipientsV1(ctx, pkiFolderID).Execute()

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
	pkiFolderID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectFolderAPI.FolderGetCommunicationrecipientsV1(context.Background(), pkiFolderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectFolderAPI.FolderGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FolderGetCommunicationrecipientsV1`: FolderGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectFolderAPI.FolderGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiFolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFolderGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderGetCommunicationrecipientsV1Response**](FolderGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FolderGetCommunicationsendersV1

> FolderGetCommunicationsendersV1Response FolderGetCommunicationsendersV1(ctx, pkiFolderID).Execute()

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
	pkiFolderID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectFolderAPI.FolderGetCommunicationsendersV1(context.Background(), pkiFolderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectFolderAPI.FolderGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FolderGetCommunicationsendersV1`: FolderGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectFolderAPI.FolderGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiFolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFolderGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderGetCommunicationsendersV1Response**](FolderGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FolderImportIntoEDMV1

> FolderImportIntoEDMV1Response FolderImportIntoEDMV1(ctx, pkiFolderID).FolderImportIntoEDMV1Request(folderImportIntoEDMV1Request).Execute()

Import attachments into the Folder

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
	pkiFolderID := int32(56) // int32 | 
	folderImportIntoEDMV1Request := *openapiclient.NewFolderImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // FolderImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectFolderAPI.FolderImportIntoEDMV1(context.Background(), pkiFolderID).FolderImportIntoEDMV1Request(folderImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectFolderAPI.FolderImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FolderImportIntoEDMV1`: FolderImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectFolderAPI.FolderImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiFolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFolderImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderImportIntoEDMV1Request** | [**FolderImportIntoEDMV1Request**](FolderImportIntoEDMV1Request.md) |  | 

### Return type

[**FolderImportIntoEDMV1Response**](FolderImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

