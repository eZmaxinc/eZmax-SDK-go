# eZmaxAPI\ObjectAdjustmentAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdjustmentBatchDownloadV1**](ObjectAdjustmentAPI.md#AdjustmentBatchDownloadV1) | **Post** /1/object/adjustment/{pkiAdjustmentID}/batchDownload | Download multiples attachments from an Adjustment
[**AdjustmentGetAttachmentsV1**](ObjectAdjustmentAPI.md#AdjustmentGetAttachmentsV1) | **Get** /1/object/adjustment/{pkiAdjustmentID}/getAttachments | Retrieve Adjustment&#39;s attachments
[**AdjustmentGetCommunicationCountV1**](ObjectAdjustmentAPI.md#AdjustmentGetCommunicationCountV1) | **Get** /1/object/adjustment/{pkiAdjustmentID}/getCommunicationCount | Retrieve Communication count
[**AdjustmentGetCommunicationListV1**](ObjectAdjustmentAPI.md#AdjustmentGetCommunicationListV1) | **Get** /1/object/adjustment/{pkiAdjustmentID}/getCommunicationList | Retrieve Communication list
[**AdjustmentGetCommunicationrecipientsV1**](ObjectAdjustmentAPI.md#AdjustmentGetCommunicationrecipientsV1) | **Get** /1/object/adjustment/{pkiAdjustmentID}/getCommunicationrecipients | Retrieve Communication recipients
[**AdjustmentGetCommunicationsendersV1**](ObjectAdjustmentAPI.md#AdjustmentGetCommunicationsendersV1) | **Get** /1/object/adjustment/{pkiAdjustmentID}/getCommunicationsenders | Retrieve Communication senders
[**AdjustmentImportIntoEDMV1**](ObjectAdjustmentAPI.md#AdjustmentImportIntoEDMV1) | **Post** /1/object/adjustment/{pkiAdjustmentID}/importIntoEDM | Import attachments into the Adjustment



## AdjustmentBatchDownloadV1

> *os.File AdjustmentBatchDownloadV1(ctx, pkiAdjustmentID).AdjustmentBatchDownloadV1Request(adjustmentBatchDownloadV1Request).Execute()

Download multiples attachments from an Adjustment

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
	pkiAdjustmentID := int32(56) // int32 | 
	adjustmentBatchDownloadV1Request := *openapiclient.NewAdjustmentBatchDownloadV1Request([]int32{int32(1)}) // AdjustmentBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAdjustmentAPI.AdjustmentBatchDownloadV1(context.Background(), pkiAdjustmentID).AdjustmentBatchDownloadV1Request(adjustmentBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAdjustmentAPI.AdjustmentBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdjustmentBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectAdjustmentAPI.AdjustmentBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAdjustmentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdjustmentBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adjustmentBatchDownloadV1Request** | [**AdjustmentBatchDownloadV1Request**](AdjustmentBatchDownloadV1Request.md) |  | 

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


## AdjustmentGetAttachmentsV1

> AdjustmentGetAttachmentsV1Response AdjustmentGetAttachmentsV1(ctx, pkiAdjustmentID).Execute()

Retrieve Adjustment's attachments

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
	pkiAdjustmentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAdjustmentAPI.AdjustmentGetAttachmentsV1(context.Background(), pkiAdjustmentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAdjustmentAPI.AdjustmentGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdjustmentGetAttachmentsV1`: AdjustmentGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAdjustmentAPI.AdjustmentGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAdjustmentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdjustmentGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdjustmentGetAttachmentsV1Response**](AdjustmentGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdjustmentGetCommunicationCountV1

> AdjustmentGetCommunicationCountV1Response AdjustmentGetCommunicationCountV1(ctx, pkiAdjustmentID).Execute()

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
	pkiAdjustmentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAdjustmentAPI.AdjustmentGetCommunicationCountV1(context.Background(), pkiAdjustmentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAdjustmentAPI.AdjustmentGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdjustmentGetCommunicationCountV1`: AdjustmentGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAdjustmentAPI.AdjustmentGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAdjustmentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdjustmentGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdjustmentGetCommunicationCountV1Response**](AdjustmentGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdjustmentGetCommunicationListV1

> AdjustmentGetCommunicationListV1Response AdjustmentGetCommunicationListV1(ctx, pkiAdjustmentID).Execute()

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
	pkiAdjustmentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAdjustmentAPI.AdjustmentGetCommunicationListV1(context.Background(), pkiAdjustmentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAdjustmentAPI.AdjustmentGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdjustmentGetCommunicationListV1`: AdjustmentGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAdjustmentAPI.AdjustmentGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAdjustmentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdjustmentGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdjustmentGetCommunicationListV1Response**](AdjustmentGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdjustmentGetCommunicationrecipientsV1

> AdjustmentGetCommunicationrecipientsV1Response AdjustmentGetCommunicationrecipientsV1(ctx, pkiAdjustmentID).Execute()

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
	pkiAdjustmentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAdjustmentAPI.AdjustmentGetCommunicationrecipientsV1(context.Background(), pkiAdjustmentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAdjustmentAPI.AdjustmentGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdjustmentGetCommunicationrecipientsV1`: AdjustmentGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAdjustmentAPI.AdjustmentGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAdjustmentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdjustmentGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdjustmentGetCommunicationrecipientsV1Response**](AdjustmentGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdjustmentGetCommunicationsendersV1

> AdjustmentGetCommunicationsendersV1Response AdjustmentGetCommunicationsendersV1(ctx, pkiAdjustmentID).Execute()

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
	pkiAdjustmentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAdjustmentAPI.AdjustmentGetCommunicationsendersV1(context.Background(), pkiAdjustmentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAdjustmentAPI.AdjustmentGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdjustmentGetCommunicationsendersV1`: AdjustmentGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAdjustmentAPI.AdjustmentGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAdjustmentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdjustmentGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdjustmentGetCommunicationsendersV1Response**](AdjustmentGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdjustmentImportIntoEDMV1

> AdjustmentImportIntoEDMV1Response AdjustmentImportIntoEDMV1(ctx, pkiAdjustmentID).AdjustmentImportIntoEDMV1Request(adjustmentImportIntoEDMV1Request).Execute()

Import attachments into the Adjustment

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
	pkiAdjustmentID := int32(56) // int32 | 
	adjustmentImportIntoEDMV1Request := *openapiclient.NewAdjustmentImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // AdjustmentImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAdjustmentAPI.AdjustmentImportIntoEDMV1(context.Background(), pkiAdjustmentID).AdjustmentImportIntoEDMV1Request(adjustmentImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAdjustmentAPI.AdjustmentImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdjustmentImportIntoEDMV1`: AdjustmentImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAdjustmentAPI.AdjustmentImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAdjustmentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdjustmentImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adjustmentImportIntoEDMV1Request** | [**AdjustmentImportIntoEDMV1Request**](AdjustmentImportIntoEDMV1Request.md) |  | 

### Return type

[**AdjustmentImportIntoEDMV1Response**](AdjustmentImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

