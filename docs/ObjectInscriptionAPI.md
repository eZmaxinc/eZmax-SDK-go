# eZmaxAPI\ObjectInscriptionAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InscriptionGetAttachmentsV1**](ObjectInscriptionAPI.md#InscriptionGetAttachmentsV1) | **Get** /1/object/inscription/{pkiInscriptionID}/getAttachments | Retrieve Inscription&#39;s Attachments
[**InscriptionGetCommunicationCountV1**](ObjectInscriptionAPI.md#InscriptionGetCommunicationCountV1) | **Get** /1/object/inscription/{pkiInscriptionID}/getCommunicationCount | Retrieve Communication count
[**InscriptionGetCommunicationListV1**](ObjectInscriptionAPI.md#InscriptionGetCommunicationListV1) | **Get** /1/object/inscription/{pkiInscriptionID}/getCommunicationList | Retrieve Communication list
[**InscriptionGetCommunicationrecipientsV1**](ObjectInscriptionAPI.md#InscriptionGetCommunicationrecipientsV1) | **Get** /1/object/inscription/{pkiInscriptionID}/getCommunicationrecipients | Retrieve Inscription&#39;s Communicationrecipient
[**InscriptionGetCommunicationsendersV1**](ObjectInscriptionAPI.md#InscriptionGetCommunicationsendersV1) | **Get** /1/object/inscription/{pkiInscriptionID}/getCommunicationsenders | Retrieve Inscription&#39;s Communicationsender
[**InscriptionGetListV1**](ObjectInscriptionAPI.md#InscriptionGetListV1) | **Get** /1/object/inscription/getList | Retrieve Inscription list
[**InscriptionImportIntoEDMV1**](ObjectInscriptionAPI.md#InscriptionImportIntoEDMV1) | **Post** /1/object/inscription/{pkiInscriptionID}/importIntoEDM | Import attachments into the Inscription
[**InscriptionPrepareFilesTransferV1**](ObjectInscriptionAPI.md#InscriptionPrepareFilesTransferV1) | **Post** /1/object/inscription/{pkiInscriptionID}/prepareFilesTransfer | Prepares file transfer into EDM



## InscriptionGetAttachmentsV1

> InscriptionGetAttachmentsV1Response InscriptionGetAttachmentsV1(ctx, pkiInscriptionID).Execute()

Retrieve Inscription's Attachments



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
	pkiInscriptionID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionAPI.InscriptionGetAttachmentsV1(context.Background(), pkiInscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionAPI.InscriptionGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionGetAttachmentsV1`: InscriptionGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionAPI.InscriptionGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptionGetAttachmentsV1Response**](InscriptionGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionGetCommunicationCountV1

> InscriptionGetCommunicationCountV1Response InscriptionGetCommunicationCountV1(ctx, pkiInscriptionID).Execute()

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
	pkiInscriptionID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionAPI.InscriptionGetCommunicationCountV1(context.Background(), pkiInscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionAPI.InscriptionGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionGetCommunicationCountV1`: InscriptionGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionAPI.InscriptionGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptionGetCommunicationCountV1Response**](InscriptionGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionGetCommunicationListV1

> InscriptionGetCommunicationListV1Response InscriptionGetCommunicationListV1(ctx, pkiInscriptionID).Execute()

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
	pkiInscriptionID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionAPI.InscriptionGetCommunicationListV1(context.Background(), pkiInscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionAPI.InscriptionGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionGetCommunicationListV1`: InscriptionGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionAPI.InscriptionGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptionGetCommunicationListV1Response**](InscriptionGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionGetCommunicationrecipientsV1

> InscriptionGetCommunicationrecipientsV1Response InscriptionGetCommunicationrecipientsV1(ctx, pkiInscriptionID).Execute()

Retrieve Inscription's Communicationrecipient



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
	pkiInscriptionID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionAPI.InscriptionGetCommunicationrecipientsV1(context.Background(), pkiInscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionAPI.InscriptionGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionGetCommunicationrecipientsV1`: InscriptionGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionAPI.InscriptionGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptionGetCommunicationrecipientsV1Response**](InscriptionGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionGetCommunicationsendersV1

> InscriptionGetCommunicationsendersV1Response InscriptionGetCommunicationsendersV1(ctx, pkiInscriptionID).Execute()

Retrieve Inscription's Communicationsender



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
	pkiInscriptionID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionAPI.InscriptionGetCommunicationsendersV1(context.Background(), pkiInscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionAPI.InscriptionGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionGetCommunicationsendersV1`: InscriptionGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionAPI.InscriptionGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InscriptionGetCommunicationsendersV1Response**](InscriptionGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionGetListV1

> InscriptionGetListV1Response InscriptionGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Inscription list



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
	resp, r, err := apiClient.ObjectInscriptionAPI.InscriptionGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionAPI.InscriptionGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionGetListV1`: InscriptionGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionAPI.InscriptionGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**InscriptionGetListV1Response**](InscriptionGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionImportIntoEDMV1

> InscriptionImportIntoEDMV1Response InscriptionImportIntoEDMV1(ctx, pkiInscriptionID).InscriptionImportIntoEDMV1Request(inscriptionImportIntoEDMV1Request).Execute()

Import attachments into the Inscription



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
	pkiInscriptionID := int32(56) // int32 | 
	inscriptionImportIntoEDMV1Request := *openapiclient.NewInscriptionImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // InscriptionImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionAPI.InscriptionImportIntoEDMV1(context.Background(), pkiInscriptionID).InscriptionImportIntoEDMV1Request(inscriptionImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionAPI.InscriptionImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionImportIntoEDMV1`: InscriptionImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionAPI.InscriptionImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inscriptionImportIntoEDMV1Request** | [**InscriptionImportIntoEDMV1Request**](InscriptionImportIntoEDMV1Request.md) |  | 

### Return type

[**InscriptionImportIntoEDMV1Response**](InscriptionImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InscriptionPrepareFilesTransferV1

> InscriptionPrepareFilesTransferV1Response InscriptionPrepareFilesTransferV1(ctx, pkiInscriptionID).InscriptionPrepareFilesTransferV1Request(inscriptionPrepareFilesTransferV1Request).Execute()

Prepares file transfer into EDM



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
	pkiInscriptionID := int32(56) // int32 | 
	inscriptionPrepareFilesTransferV1Request := *openapiclient.NewInscriptionPrepareFilesTransferV1Request([]openapiclient.CustomAttachmentPrepareFilesTransferRequest{*openapiclient.NewCustomAttachmentPrepareFilesTransferRequest("Document.pdf", "2cb29026e8a93c930e71598579400db6")}) // InscriptionPrepareFilesTransferV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectInscriptionAPI.InscriptionPrepareFilesTransferV1(context.Background(), pkiInscriptionID).InscriptionPrepareFilesTransferV1Request(inscriptionPrepareFilesTransferV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectInscriptionAPI.InscriptionPrepareFilesTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InscriptionPrepareFilesTransferV1`: InscriptionPrepareFilesTransferV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectInscriptionAPI.InscriptionPrepareFilesTransferV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInscriptionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInscriptionPrepareFilesTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inscriptionPrepareFilesTransferV1Request** | [**InscriptionPrepareFilesTransferV1Request**](InscriptionPrepareFilesTransferV1Request.md) |  | 

### Return type

[**InscriptionPrepareFilesTransferV1Response**](InscriptionPrepareFilesTransferV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

