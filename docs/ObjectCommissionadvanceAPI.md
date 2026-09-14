# eZmaxAPI\ObjectCommissionadvanceAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommissionadvanceBatchDownloadV1**](ObjectCommissionadvanceAPI.md#CommissionadvanceBatchDownloadV1) | **Post** /1/object/commissionadvance/{pkiCommissionadvanceID}/batchDownload | Download multiples attachments from a Commission advance
[**CommissionadvanceGetAttachmentsV1**](ObjectCommissionadvanceAPI.md#CommissionadvanceGetAttachmentsV1) | **Get** /1/object/commissionadvance/{pkiCommissionadvanceID}/getAttachments | Retrieve Commissionadvance&#39;s attachments
[**CommissionadvanceGetCommunicationCountV1**](ObjectCommissionadvanceAPI.md#CommissionadvanceGetCommunicationCountV1) | **Get** /1/object/commissionadvance/{pkiCommissionadvanceID}/getCommunicationCount | Retrieve Communication count
[**CommissionadvanceGetCommunicationListV1**](ObjectCommissionadvanceAPI.md#CommissionadvanceGetCommunicationListV1) | **Get** /1/object/commissionadvance/{pkiCommissionadvanceID}/getCommunicationList | Retrieve Communication list
[**CommissionadvanceGetCommunicationrecipientsV1**](ObjectCommissionadvanceAPI.md#CommissionadvanceGetCommunicationrecipientsV1) | **Get** /1/object/commissionadvance/{pkiCommissionadvanceID}/getCommunicationrecipients | Retrieve Communication recipients
[**CommissionadvanceGetCommunicationsendersV1**](ObjectCommissionadvanceAPI.md#CommissionadvanceGetCommunicationsendersV1) | **Get** /1/object/commissionadvance/{pkiCommissionadvanceID}/getCommunicationsenders | Retrieve Communication senders
[**CommissionadvanceImportIntoEDMV1**](ObjectCommissionadvanceAPI.md#CommissionadvanceImportIntoEDMV1) | **Post** /1/object/commissionadvance/{pkiCommissionadvanceID}/importIntoEDM | Import attachments into the Commissionadvance



## CommissionadvanceBatchDownloadV1

> *os.File CommissionadvanceBatchDownloadV1(ctx, pkiCommissionadvanceID).CommissionadvanceBatchDownloadV1Request(commissionadvanceBatchDownloadV1Request).Execute()

Download multiples attachments from a Commission advance

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
	pkiCommissionadvanceID := int32(56) // int32 | 
	commissionadvanceBatchDownloadV1Request := *openapiclient.NewCommissionadvanceBatchDownloadV1Request([]int32{int32(1)}) // CommissionadvanceBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCommissionadvanceAPI.CommissionadvanceBatchDownloadV1(context.Background(), pkiCommissionadvanceID).CommissionadvanceBatchDownloadV1Request(commissionadvanceBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCommissionadvanceAPI.CommissionadvanceBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommissionadvanceBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectCommissionadvanceAPI.CommissionadvanceBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCommissionadvanceID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommissionadvanceBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commissionadvanceBatchDownloadV1Request** | [**CommissionadvanceBatchDownloadV1Request**](CommissionadvanceBatchDownloadV1Request.md) |  | 

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


## CommissionadvanceGetAttachmentsV1

> CommissionadvanceGetAttachmentsV1Response CommissionadvanceGetAttachmentsV1(ctx, pkiCommissionadvanceID).Execute()

Retrieve Commissionadvance's attachments

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
	pkiCommissionadvanceID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCommissionadvanceAPI.CommissionadvanceGetAttachmentsV1(context.Background(), pkiCommissionadvanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCommissionadvanceAPI.CommissionadvanceGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommissionadvanceGetAttachmentsV1`: CommissionadvanceGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCommissionadvanceAPI.CommissionadvanceGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCommissionadvanceID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommissionadvanceGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommissionadvanceGetAttachmentsV1Response**](CommissionadvanceGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommissionadvanceGetCommunicationCountV1

> CommissionadvanceGetCommunicationCountV1Response CommissionadvanceGetCommunicationCountV1(ctx, pkiCommissionadvanceID).Execute()

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
	pkiCommissionadvanceID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCommissionadvanceAPI.CommissionadvanceGetCommunicationCountV1(context.Background(), pkiCommissionadvanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCommissionadvanceAPI.CommissionadvanceGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommissionadvanceGetCommunicationCountV1`: CommissionadvanceGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCommissionadvanceAPI.CommissionadvanceGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCommissionadvanceID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommissionadvanceGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommissionadvanceGetCommunicationCountV1Response**](CommissionadvanceGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommissionadvanceGetCommunicationListV1

> CommissionadvanceGetCommunicationListV1Response CommissionadvanceGetCommunicationListV1(ctx, pkiCommissionadvanceID).Execute()

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
	pkiCommissionadvanceID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCommissionadvanceAPI.CommissionadvanceGetCommunicationListV1(context.Background(), pkiCommissionadvanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCommissionadvanceAPI.CommissionadvanceGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommissionadvanceGetCommunicationListV1`: CommissionadvanceGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCommissionadvanceAPI.CommissionadvanceGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCommissionadvanceID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommissionadvanceGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommissionadvanceGetCommunicationListV1Response**](CommissionadvanceGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommissionadvanceGetCommunicationrecipientsV1

> CommissionadvanceGetCommunicationrecipientsV1Response CommissionadvanceGetCommunicationrecipientsV1(ctx, pkiCommissionadvanceID).Execute()

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
	pkiCommissionadvanceID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCommissionadvanceAPI.CommissionadvanceGetCommunicationrecipientsV1(context.Background(), pkiCommissionadvanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCommissionadvanceAPI.CommissionadvanceGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommissionadvanceGetCommunicationrecipientsV1`: CommissionadvanceGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCommissionadvanceAPI.CommissionadvanceGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCommissionadvanceID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommissionadvanceGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommissionadvanceGetCommunicationrecipientsV1Response**](CommissionadvanceGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommissionadvanceGetCommunicationsendersV1

> CommissionadvanceGetCommunicationsendersV1Response CommissionadvanceGetCommunicationsendersV1(ctx, pkiCommissionadvanceID).Execute()

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
	pkiCommissionadvanceID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCommissionadvanceAPI.CommissionadvanceGetCommunicationsendersV1(context.Background(), pkiCommissionadvanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCommissionadvanceAPI.CommissionadvanceGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommissionadvanceGetCommunicationsendersV1`: CommissionadvanceGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCommissionadvanceAPI.CommissionadvanceGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCommissionadvanceID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommissionadvanceGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommissionadvanceGetCommunicationsendersV1Response**](CommissionadvanceGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommissionadvanceImportIntoEDMV1

> CommissionadvanceImportIntoEDMV1Response CommissionadvanceImportIntoEDMV1(ctx, pkiCommissionadvanceID).CommissionadvanceImportIntoEDMV1Request(commissionadvanceImportIntoEDMV1Request).Execute()

Import attachments into the Commissionadvance

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
	pkiCommissionadvanceID := int32(56) // int32 | 
	commissionadvanceImportIntoEDMV1Request := *openapiclient.NewCommissionadvanceImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // CommissionadvanceImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCommissionadvanceAPI.CommissionadvanceImportIntoEDMV1(context.Background(), pkiCommissionadvanceID).CommissionadvanceImportIntoEDMV1Request(commissionadvanceImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCommissionadvanceAPI.CommissionadvanceImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommissionadvanceImportIntoEDMV1`: CommissionadvanceImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCommissionadvanceAPI.CommissionadvanceImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCommissionadvanceID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommissionadvanceImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commissionadvanceImportIntoEDMV1Request** | [**CommissionadvanceImportIntoEDMV1Request**](CommissionadvanceImportIntoEDMV1Request.md) |  | 

### Return type

[**CommissionadvanceImportIntoEDMV1Response**](CommissionadvanceImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

