# eZmaxAPI\ObjectExternalbrokerAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExternalbrokerBatchDownloadV1**](ObjectExternalbrokerAPI.md#ExternalbrokerBatchDownloadV1) | **Post** /1/object/externalbroker/{pkiExternalbrokerID}/batchDownload | Download multiples attachments from an Externalbroker
[**ExternalbrokerGetAttachmentsV1**](ObjectExternalbrokerAPI.md#ExternalbrokerGetAttachmentsV1) | **Get** /1/object/externalbroker/{pkiExternalbrokerID}/getAttachments | Retrieve Externalbroker&#39;s attachments
[**ExternalbrokerGetCommunicationCountV1**](ObjectExternalbrokerAPI.md#ExternalbrokerGetCommunicationCountV1) | **Get** /1/object/externalbroker/{pkiExternalbrokerID}/getCommunicationCount | Retrieve Communication count
[**ExternalbrokerGetCommunicationListV1**](ObjectExternalbrokerAPI.md#ExternalbrokerGetCommunicationListV1) | **Get** /1/object/externalbroker/{pkiExternalbrokerID}/getCommunicationList | Retrieve Communication list
[**ExternalbrokerGetCommunicationrecipientsV1**](ObjectExternalbrokerAPI.md#ExternalbrokerGetCommunicationrecipientsV1) | **Get** /1/object/externalbroker/{pkiExternalbrokerID}/getCommunicationrecipients | Retrieve Communication recipients
[**ExternalbrokerGetCommunicationsendersV1**](ObjectExternalbrokerAPI.md#ExternalbrokerGetCommunicationsendersV1) | **Get** /1/object/externalbroker/{pkiExternalbrokerID}/getCommunicationsenders | Retrieve Communication senders
[**ExternalbrokerImportIntoEDMV1**](ObjectExternalbrokerAPI.md#ExternalbrokerImportIntoEDMV1) | **Post** /1/object/externalbroker/{pkiExternalbrokerID}/importIntoEDM | Import attachments into the Externalbroker



## ExternalbrokerBatchDownloadV1

> *os.File ExternalbrokerBatchDownloadV1(ctx, pkiExternalbrokerID).ExternalbrokerBatchDownloadV1Request(externalbrokerBatchDownloadV1Request).Execute()

Download multiples attachments from an Externalbroker

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
	pkiExternalbrokerID := int32(56) // int32 | 
	externalbrokerBatchDownloadV1Request := *openapiclient.NewExternalbrokerBatchDownloadV1Request([]int32{int32(1)}) // ExternalbrokerBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectExternalbrokerAPI.ExternalbrokerBatchDownloadV1(context.Background(), pkiExternalbrokerID).ExternalbrokerBatchDownloadV1Request(externalbrokerBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectExternalbrokerAPI.ExternalbrokerBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalbrokerBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectExternalbrokerAPI.ExternalbrokerBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiExternalbrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalbrokerBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalbrokerBatchDownloadV1Request** | [**ExternalbrokerBatchDownloadV1Request**](ExternalbrokerBatchDownloadV1Request.md) |  | 

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


## ExternalbrokerGetAttachmentsV1

> ExternalbrokerGetAttachmentsV1Response ExternalbrokerGetAttachmentsV1(ctx, pkiExternalbrokerID).Execute()

Retrieve Externalbroker's attachments

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
	pkiExternalbrokerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectExternalbrokerAPI.ExternalbrokerGetAttachmentsV1(context.Background(), pkiExternalbrokerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectExternalbrokerAPI.ExternalbrokerGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalbrokerGetAttachmentsV1`: ExternalbrokerGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectExternalbrokerAPI.ExternalbrokerGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiExternalbrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalbrokerGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalbrokerGetAttachmentsV1Response**](ExternalbrokerGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalbrokerGetCommunicationCountV1

> ExternalbrokerGetCommunicationCountV1Response ExternalbrokerGetCommunicationCountV1(ctx, pkiExternalbrokerID).Execute()

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
	pkiExternalbrokerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectExternalbrokerAPI.ExternalbrokerGetCommunicationCountV1(context.Background(), pkiExternalbrokerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectExternalbrokerAPI.ExternalbrokerGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalbrokerGetCommunicationCountV1`: ExternalbrokerGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectExternalbrokerAPI.ExternalbrokerGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiExternalbrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalbrokerGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalbrokerGetCommunicationCountV1Response**](ExternalbrokerGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalbrokerGetCommunicationListV1

> ExternalbrokerGetCommunicationListV1Response ExternalbrokerGetCommunicationListV1(ctx, pkiExternalbrokerID).Execute()

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
	pkiExternalbrokerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectExternalbrokerAPI.ExternalbrokerGetCommunicationListV1(context.Background(), pkiExternalbrokerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectExternalbrokerAPI.ExternalbrokerGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalbrokerGetCommunicationListV1`: ExternalbrokerGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectExternalbrokerAPI.ExternalbrokerGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiExternalbrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalbrokerGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalbrokerGetCommunicationListV1Response**](ExternalbrokerGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalbrokerGetCommunicationrecipientsV1

> ExternalbrokerGetCommunicationrecipientsV1Response ExternalbrokerGetCommunicationrecipientsV1(ctx, pkiExternalbrokerID).Execute()

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
	pkiExternalbrokerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectExternalbrokerAPI.ExternalbrokerGetCommunicationrecipientsV1(context.Background(), pkiExternalbrokerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectExternalbrokerAPI.ExternalbrokerGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalbrokerGetCommunicationrecipientsV1`: ExternalbrokerGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectExternalbrokerAPI.ExternalbrokerGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiExternalbrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalbrokerGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalbrokerGetCommunicationrecipientsV1Response**](ExternalbrokerGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalbrokerGetCommunicationsendersV1

> ExternalbrokerGetCommunicationsendersV1Response ExternalbrokerGetCommunicationsendersV1(ctx, pkiExternalbrokerID).Execute()

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
	pkiExternalbrokerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectExternalbrokerAPI.ExternalbrokerGetCommunicationsendersV1(context.Background(), pkiExternalbrokerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectExternalbrokerAPI.ExternalbrokerGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalbrokerGetCommunicationsendersV1`: ExternalbrokerGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectExternalbrokerAPI.ExternalbrokerGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiExternalbrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalbrokerGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalbrokerGetCommunicationsendersV1Response**](ExternalbrokerGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalbrokerImportIntoEDMV1

> ExternalbrokerImportIntoEDMV1Response ExternalbrokerImportIntoEDMV1(ctx, pkiExternalbrokerID).ExternalbrokerImportIntoEDMV1Request(externalbrokerImportIntoEDMV1Request).Execute()

Import attachments into the Externalbroker



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
	pkiExternalbrokerID := int32(56) // int32 | 
	externalbrokerImportIntoEDMV1Request := *openapiclient.NewExternalbrokerImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // ExternalbrokerImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectExternalbrokerAPI.ExternalbrokerImportIntoEDMV1(context.Background(), pkiExternalbrokerID).ExternalbrokerImportIntoEDMV1Request(externalbrokerImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectExternalbrokerAPI.ExternalbrokerImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalbrokerImportIntoEDMV1`: ExternalbrokerImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectExternalbrokerAPI.ExternalbrokerImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiExternalbrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalbrokerImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalbrokerImportIntoEDMV1Request** | [**ExternalbrokerImportIntoEDMV1Request**](ExternalbrokerImportIntoEDMV1Request.md) |  | 

### Return type

[**ExternalbrokerImportIntoEDMV1Response**](ExternalbrokerImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

