# eZmaxAPI\ObjectInvoiceAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceGetAttachmentsV1**](ObjectInvoiceAPI.md#InvoiceGetAttachmentsV1) | **Get** /1/object/invoice/{pkiInvoiceID}/getAttachments | Retrieve Invoice&#39;s Attachments
[**InvoiceGetCommunicationCountV1**](ObjectInvoiceAPI.md#InvoiceGetCommunicationCountV1) | **Get** /1/object/invoice/{pkiInvoiceID}/getCommunicationCount | Retrieve Communication count
[**InvoiceGetCommunicationListV1**](ObjectInvoiceAPI.md#InvoiceGetCommunicationListV1) | **Get** /1/object/invoice/{pkiInvoiceID}/getCommunicationList | Retrieve Communication list
[**InvoiceGetCommunicationrecipientsV1**](ObjectInvoiceAPI.md#InvoiceGetCommunicationrecipientsV1) | **Get** /1/object/invoice/{pkiInvoiceID}/getCommunicationrecipients | Retrieve Invoice&#39;s Communicationrecipient
[**InvoiceGetCommunicationsendersV1**](ObjectInvoiceAPI.md#InvoiceGetCommunicationsendersV1) | **Get** /1/object/invoice/{pkiInvoiceID}/getCommunicationsenders | Retrieve Invoice&#39;s Communicationsender
[**InvoiceImportIntoEDMV1**](ObjectInvoiceAPI.md#InvoiceImportIntoEDMV1) | **Post** /1/object/invoice/{pkiInvoiceID}/importIntoEDM | Import attachments into the Invoice



## InvoiceGetAttachmentsV1

> InvoiceGetAttachmentsV1Response InvoiceGetAttachmentsV1(ctx, pkiInvoiceID).Execute()

Retrieve Invoice's Attachments



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
	pkiInvoiceID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInvoiceAPI.InvoiceGetAttachmentsV1(context.Background(), pkiInvoiceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInvoiceAPI.InvoiceGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceGetAttachmentsV1`: InvoiceGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInvoiceAPI.InvoiceGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInvoiceID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceGetAttachmentsV1Response**](InvoiceGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceGetCommunicationCountV1

> InvoiceGetCommunicationCountV1Response InvoiceGetCommunicationCountV1(ctx, pkiInvoiceID).Execute()

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
	pkiInvoiceID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInvoiceAPI.InvoiceGetCommunicationCountV1(context.Background(), pkiInvoiceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInvoiceAPI.InvoiceGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceGetCommunicationCountV1`: InvoiceGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInvoiceAPI.InvoiceGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInvoiceID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceGetCommunicationCountV1Response**](InvoiceGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceGetCommunicationListV1

> InvoiceGetCommunicationListV1Response InvoiceGetCommunicationListV1(ctx, pkiInvoiceID).Execute()

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
	pkiInvoiceID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInvoiceAPI.InvoiceGetCommunicationListV1(context.Background(), pkiInvoiceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInvoiceAPI.InvoiceGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceGetCommunicationListV1`: InvoiceGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInvoiceAPI.InvoiceGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInvoiceID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceGetCommunicationListV1Response**](InvoiceGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceGetCommunicationrecipientsV1

> InvoiceGetCommunicationrecipientsV1Response InvoiceGetCommunicationrecipientsV1(ctx, pkiInvoiceID).Execute()

Retrieve Invoice's Communicationrecipient



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
	pkiInvoiceID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInvoiceAPI.InvoiceGetCommunicationrecipientsV1(context.Background(), pkiInvoiceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInvoiceAPI.InvoiceGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceGetCommunicationrecipientsV1`: InvoiceGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInvoiceAPI.InvoiceGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInvoiceID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceGetCommunicationrecipientsV1Response**](InvoiceGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceGetCommunicationsendersV1

> InvoiceGetCommunicationsendersV1Response InvoiceGetCommunicationsendersV1(ctx, pkiInvoiceID).Execute()

Retrieve Invoice's Communicationsender



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
	pkiInvoiceID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInvoiceAPI.InvoiceGetCommunicationsendersV1(context.Background(), pkiInvoiceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInvoiceAPI.InvoiceGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceGetCommunicationsendersV1`: InvoiceGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInvoiceAPI.InvoiceGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInvoiceID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceGetCommunicationsendersV1Response**](InvoiceGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceImportIntoEDMV1

> InvoiceImportIntoEDMV1Response InvoiceImportIntoEDMV1(ctx, pkiInvoiceID).InvoiceImportIntoEDMV1Request(invoiceImportIntoEDMV1Request).Execute()

Import attachments into the Invoice



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
	pkiInvoiceID := int32(56) // int32 | 
	invoiceImportIntoEDMV1Request := *openapiclient.NewInvoiceImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // InvoiceImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInvoiceAPI.InvoiceImportIntoEDMV1(context.Background(), pkiInvoiceID).InvoiceImportIntoEDMV1Request(invoiceImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInvoiceAPI.InvoiceImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceImportIntoEDMV1`: InvoiceImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInvoiceAPI.InvoiceImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInvoiceID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceImportIntoEDMV1Request** | [**InvoiceImportIntoEDMV1Request**](InvoiceImportIntoEDMV1Request.md) |  | 

### Return type

[**InvoiceImportIntoEDMV1Response**](InvoiceImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

