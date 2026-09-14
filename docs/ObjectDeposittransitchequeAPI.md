# eZmaxAPI\ObjectDeposittransitchequeAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeposittransitchequeBatchDownloadV1**](ObjectDeposittransitchequeAPI.md#DeposittransitchequeBatchDownloadV1) | **Post** /1/object/deposittransitcheque/{pkiDeposittransitchequeID}/batchDownload | Download multiples attachments from a Deposittransitcheque
[**DeposittransitchequeGetAttachmentsV1**](ObjectDeposittransitchequeAPI.md#DeposittransitchequeGetAttachmentsV1) | **Get** /1/object/deposittransitcheque/{pkiDeposittransitchequeID}/getAttachments | Retrieve Deposittransitcheque&#39;s attachments
[**DeposittransitchequeGetCommunicationCountV1**](ObjectDeposittransitchequeAPI.md#DeposittransitchequeGetCommunicationCountV1) | **Get** /1/object/deposittransitcheque/{pkiDeposittransitchequeID}/getCommunicationCount | Retrieve Communication count
[**DeposittransitchequeGetCommunicationListV1**](ObjectDeposittransitchequeAPI.md#DeposittransitchequeGetCommunicationListV1) | **Get** /1/object/deposittransitcheque/{pkiDeposittransitchequeID}/getCommunicationList | Retrieve Communication list
[**DeposittransitchequeGetCommunicationrecipientsV1**](ObjectDeposittransitchequeAPI.md#DeposittransitchequeGetCommunicationrecipientsV1) | **Get** /1/object/deposittransitcheque/{pkiDeposittransitchequeID}/getCommunicationrecipients | Retrieve Communication recipients
[**DeposittransitchequeGetCommunicationsendersV1**](ObjectDeposittransitchequeAPI.md#DeposittransitchequeGetCommunicationsendersV1) | **Get** /1/object/deposittransitcheque/{pkiDeposittransitchequeID}/getCommunicationsenders | Retrieve Communication senders
[**DeposittransitchequeImportIntoEDMV1**](ObjectDeposittransitchequeAPI.md#DeposittransitchequeImportIntoEDMV1) | **Post** /1/object/deposittransitcheque/{pkiDeposittransitchequeID}/importIntoEDM | Import attachments into the Deposittransitcheque



## DeposittransitchequeBatchDownloadV1

> *os.File DeposittransitchequeBatchDownloadV1(ctx, pkiDeposittransitchequeID).DeposittransitchequeBatchDownloadV1Request(deposittransitchequeBatchDownloadV1Request).Execute()

Download multiples attachments from a Deposittransitcheque

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
	pkiDeposittransitchequeID := int32(56) // int32 | 
	deposittransitchequeBatchDownloadV1Request := *openapiclient.NewDeposittransitchequeBatchDownloadV1Request([]int32{int32(1)}) // DeposittransitchequeBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDeposittransitchequeAPI.DeposittransitchequeBatchDownloadV1(context.Background(), pkiDeposittransitchequeID).DeposittransitchequeBatchDownloadV1Request(deposittransitchequeBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDeposittransitchequeAPI.DeposittransitchequeBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeposittransitchequeBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectDeposittransitchequeAPI.DeposittransitchequeBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDeposittransitchequeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeposittransitchequeBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deposittransitchequeBatchDownloadV1Request** | [**DeposittransitchequeBatchDownloadV1Request**](DeposittransitchequeBatchDownloadV1Request.md) |  | 

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


## DeposittransitchequeGetAttachmentsV1

> DeposittransitchequeGetAttachmentsV1Response DeposittransitchequeGetAttachmentsV1(ctx, pkiDeposittransitchequeID).Execute()

Retrieve Deposittransitcheque's attachments

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
	pkiDeposittransitchequeID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDeposittransitchequeAPI.DeposittransitchequeGetAttachmentsV1(context.Background(), pkiDeposittransitchequeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDeposittransitchequeAPI.DeposittransitchequeGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeposittransitchequeGetAttachmentsV1`: DeposittransitchequeGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDeposittransitchequeAPI.DeposittransitchequeGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDeposittransitchequeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeposittransitchequeGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeposittransitchequeGetAttachmentsV1Response**](DeposittransitchequeGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeposittransitchequeGetCommunicationCountV1

> DeposittransitchequeGetCommunicationCountV1Response DeposittransitchequeGetCommunicationCountV1(ctx, pkiDeposittransitchequeID).Execute()

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
	pkiDeposittransitchequeID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDeposittransitchequeAPI.DeposittransitchequeGetCommunicationCountV1(context.Background(), pkiDeposittransitchequeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDeposittransitchequeAPI.DeposittransitchequeGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeposittransitchequeGetCommunicationCountV1`: DeposittransitchequeGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDeposittransitchequeAPI.DeposittransitchequeGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDeposittransitchequeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeposittransitchequeGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeposittransitchequeGetCommunicationCountV1Response**](DeposittransitchequeGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeposittransitchequeGetCommunicationListV1

> DeposittransitchequeGetCommunicationListV1Response DeposittransitchequeGetCommunicationListV1(ctx, pkiDeposittransitchequeID).Execute()

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
	pkiDeposittransitchequeID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDeposittransitchequeAPI.DeposittransitchequeGetCommunicationListV1(context.Background(), pkiDeposittransitchequeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDeposittransitchequeAPI.DeposittransitchequeGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeposittransitchequeGetCommunicationListV1`: DeposittransitchequeGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDeposittransitchequeAPI.DeposittransitchequeGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDeposittransitchequeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeposittransitchequeGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeposittransitchequeGetCommunicationListV1Response**](DeposittransitchequeGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeposittransitchequeGetCommunicationrecipientsV1

> DeposittransitchequeGetCommunicationrecipientsV1Response DeposittransitchequeGetCommunicationrecipientsV1(ctx, pkiDeposittransitchequeID).Execute()

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
	pkiDeposittransitchequeID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDeposittransitchequeAPI.DeposittransitchequeGetCommunicationrecipientsV1(context.Background(), pkiDeposittransitchequeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDeposittransitchequeAPI.DeposittransitchequeGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeposittransitchequeGetCommunicationrecipientsV1`: DeposittransitchequeGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDeposittransitchequeAPI.DeposittransitchequeGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDeposittransitchequeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeposittransitchequeGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeposittransitchequeGetCommunicationrecipientsV1Response**](DeposittransitchequeGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeposittransitchequeGetCommunicationsendersV1

> DeposittransitchequeGetCommunicationsendersV1Response DeposittransitchequeGetCommunicationsendersV1(ctx, pkiDeposittransitchequeID).Execute()

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
	pkiDeposittransitchequeID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDeposittransitchequeAPI.DeposittransitchequeGetCommunicationsendersV1(context.Background(), pkiDeposittransitchequeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDeposittransitchequeAPI.DeposittransitchequeGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeposittransitchequeGetCommunicationsendersV1`: DeposittransitchequeGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDeposittransitchequeAPI.DeposittransitchequeGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDeposittransitchequeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeposittransitchequeGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeposittransitchequeGetCommunicationsendersV1Response**](DeposittransitchequeGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeposittransitchequeImportIntoEDMV1

> DeposittransitchequeImportIntoEDMV1Response DeposittransitchequeImportIntoEDMV1(ctx, pkiDeposittransitchequeID).DeposittransitchequeImportIntoEDMV1Request(deposittransitchequeImportIntoEDMV1Request).Execute()

Import attachments into the Deposittransitcheque

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
	pkiDeposittransitchequeID := int32(56) // int32 | 
	deposittransitchequeImportIntoEDMV1Request := *openapiclient.NewDeposittransitchequeImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // DeposittransitchequeImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDeposittransitchequeAPI.DeposittransitchequeImportIntoEDMV1(context.Background(), pkiDeposittransitchequeID).DeposittransitchequeImportIntoEDMV1Request(deposittransitchequeImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDeposittransitchequeAPI.DeposittransitchequeImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeposittransitchequeImportIntoEDMV1`: DeposittransitchequeImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDeposittransitchequeAPI.DeposittransitchequeImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDeposittransitchequeID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeposittransitchequeImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deposittransitchequeImportIntoEDMV1Request** | [**DeposittransitchequeImportIntoEDMV1Request**](DeposittransitchequeImportIntoEDMV1Request.md) |  | 

### Return type

[**DeposittransitchequeImportIntoEDMV1Response**](DeposittransitchequeImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

