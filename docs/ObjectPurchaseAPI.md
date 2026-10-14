# eZmaxAPI\ObjectPurchaseAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PurchaseBatchDownloadV1**](ObjectPurchaseAPI.md#PurchaseBatchDownloadV1) | **Post** /1/object/purchase/{pkiPurchaseID}/batchDownload | Download multiples attachments from a Purchase
[**PurchaseGetAttachmentsV1**](ObjectPurchaseAPI.md#PurchaseGetAttachmentsV1) | **Get** /1/object/purchase/{pkiPurchaseID}/getAttachments | Retrieve Purchase&#39;s attachments
[**PurchaseGetCommunicationCountV1**](ObjectPurchaseAPI.md#PurchaseGetCommunicationCountV1) | **Get** /1/object/purchase/{pkiPurchaseID}/getCommunicationCount | Retrieve Communication count
[**PurchaseGetCommunicationListV1**](ObjectPurchaseAPI.md#PurchaseGetCommunicationListV1) | **Get** /1/object/purchase/{pkiPurchaseID}/getCommunicationList | Retrieve Communication list
[**PurchaseGetCommunicationrecipientsV1**](ObjectPurchaseAPI.md#PurchaseGetCommunicationrecipientsV1) | **Get** /1/object/purchase/{pkiPurchaseID}/getCommunicationrecipients | Retrieve Communication recipients
[**PurchaseGetCommunicationsendersV1**](ObjectPurchaseAPI.md#PurchaseGetCommunicationsendersV1) | **Get** /1/object/purchase/{pkiPurchaseID}/getCommunicationsenders | Retrieve Communication senders
[**PurchaseGetListV1**](ObjectPurchaseAPI.md#PurchaseGetListV1) | **Get** /1/object/purchase/getList | Retrieve Purchase list
[**PurchaseImportIntoEDMV1**](ObjectPurchaseAPI.md#PurchaseImportIntoEDMV1) | **Post** /1/object/purchase/{pkiPurchaseID}/importIntoEDM | Import attachments into the Purchase



## PurchaseBatchDownloadV1

> *os.File PurchaseBatchDownloadV1(ctx, pkiPurchaseID).PurchaseBatchDownloadV1Request(purchaseBatchDownloadV1Request).Execute()

Download multiples attachments from a Purchase

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
	pkiPurchaseID := int32(56) // int32 | 
	purchaseBatchDownloadV1Request := *openapiclient.NewPurchaseBatchDownloadV1Request([]int32{int32(1)}) // PurchaseBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPurchaseAPI.PurchaseBatchDownloadV1(context.Background(), pkiPurchaseID).PurchaseBatchDownloadV1Request(purchaseBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPurchaseAPI.PurchaseBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectPurchaseAPI.PurchaseBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purchaseBatchDownloadV1Request** | [**PurchaseBatchDownloadV1Request**](PurchaseBatchDownloadV1Request.md) |  | 

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


## PurchaseGetAttachmentsV1

> PurchaseGetAttachmentsV1Response PurchaseGetAttachmentsV1(ctx, pkiPurchaseID).Execute()

Retrieve Purchase's attachments

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
	pkiPurchaseID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPurchaseAPI.PurchaseGetAttachmentsV1(context.Background(), pkiPurchaseID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPurchaseAPI.PurchaseGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseGetAttachmentsV1`: PurchaseGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPurchaseAPI.PurchaseGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PurchaseGetAttachmentsV1Response**](PurchaseGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchaseGetCommunicationCountV1

> PurchaseGetCommunicationCountV1Response PurchaseGetCommunicationCountV1(ctx, pkiPurchaseID).Execute()

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
	pkiPurchaseID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPurchaseAPI.PurchaseGetCommunicationCountV1(context.Background(), pkiPurchaseID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPurchaseAPI.PurchaseGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseGetCommunicationCountV1`: PurchaseGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPurchaseAPI.PurchaseGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PurchaseGetCommunicationCountV1Response**](PurchaseGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchaseGetCommunicationListV1

> PurchaseGetCommunicationListV1Response PurchaseGetCommunicationListV1(ctx, pkiPurchaseID).Execute()

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
	pkiPurchaseID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPurchaseAPI.PurchaseGetCommunicationListV1(context.Background(), pkiPurchaseID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPurchaseAPI.PurchaseGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseGetCommunicationListV1`: PurchaseGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPurchaseAPI.PurchaseGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PurchaseGetCommunicationListV1Response**](PurchaseGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchaseGetCommunicationrecipientsV1

> PurchaseGetCommunicationrecipientsV1Response PurchaseGetCommunicationrecipientsV1(ctx, pkiPurchaseID).Execute()

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
	pkiPurchaseID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPurchaseAPI.PurchaseGetCommunicationrecipientsV1(context.Background(), pkiPurchaseID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPurchaseAPI.PurchaseGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseGetCommunicationrecipientsV1`: PurchaseGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPurchaseAPI.PurchaseGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PurchaseGetCommunicationrecipientsV1Response**](PurchaseGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchaseGetCommunicationsendersV1

> PurchaseGetCommunicationsendersV1Response PurchaseGetCommunicationsendersV1(ctx, pkiPurchaseID).Execute()

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
	pkiPurchaseID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPurchaseAPI.PurchaseGetCommunicationsendersV1(context.Background(), pkiPurchaseID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPurchaseAPI.PurchaseGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseGetCommunicationsendersV1`: PurchaseGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPurchaseAPI.PurchaseGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PurchaseGetCommunicationsendersV1Response**](PurchaseGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchaseGetListV1

> PurchaseGetListV1Response PurchaseGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Purchase list



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
	resp, r, err := apiClient.ObjectPurchaseAPI.PurchaseGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPurchaseAPI.PurchaseGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseGetListV1`: PurchaseGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPurchaseAPI.PurchaseGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**PurchaseGetListV1Response**](PurchaseGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchaseImportIntoEDMV1

> PurchaseImportIntoEDMV1Response PurchaseImportIntoEDMV1(ctx, pkiPurchaseID).PurchaseImportIntoEDMV1Request(purchaseImportIntoEDMV1Request).Execute()

Import attachments into the Purchase

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
	pkiPurchaseID := int32(56) // int32 | 
	purchaseImportIntoEDMV1Request := *openapiclient.NewPurchaseImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // PurchaseImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPurchaseAPI.PurchaseImportIntoEDMV1(context.Background(), pkiPurchaseID).PurchaseImportIntoEDMV1Request(purchaseImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPurchaseAPI.PurchaseImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseImportIntoEDMV1`: PurchaseImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPurchaseAPI.PurchaseImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purchaseImportIntoEDMV1Request** | [**PurchaseImportIntoEDMV1Request**](PurchaseImportIntoEDMV1Request.md) |  | 

### Return type

[**PurchaseImportIntoEDMV1Response**](PurchaseImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

