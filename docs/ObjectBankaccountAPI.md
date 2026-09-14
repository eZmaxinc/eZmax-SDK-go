# eZmaxAPI\ObjectBankaccountAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BankaccountBatchDownloadV1**](ObjectBankaccountAPI.md#BankaccountBatchDownloadV1) | **Post** /1/object/bankaccount/{pkiBankaccountID}/batchDownload | Download multiples attachments from a Bankaccount
[**BankaccountGetAttachmentsV1**](ObjectBankaccountAPI.md#BankaccountGetAttachmentsV1) | **Get** /1/object/bankaccount/{pkiBankaccountID}/getAttachments | Retrieve Bankaccount&#39;s attachments
[**BankaccountGetAutocompleteV2**](ObjectBankaccountAPI.md#BankaccountGetAutocompleteV2) | **Get** /2/object/bankaccount/getAutocomplete/{sSelector} | Retrieve Bankaccounts and IDs
[**BankaccountGetCommunicationCountV1**](ObjectBankaccountAPI.md#BankaccountGetCommunicationCountV1) | **Get** /1/object/bankaccount/{pkiBankaccountID}/getCommunicationCount | Retrieve Communication count
[**BankaccountGetCommunicationListV1**](ObjectBankaccountAPI.md#BankaccountGetCommunicationListV1) | **Get** /1/object/bankaccount/{pkiBankaccountID}/getCommunicationList | Retrieve Communication list
[**BankaccountGetCommunicationrecipientsV1**](ObjectBankaccountAPI.md#BankaccountGetCommunicationrecipientsV1) | **Get** /1/object/bankaccount/{pkiBankaccountID}/getCommunicationrecipients | Retrieve Communication recipients
[**BankaccountGetCommunicationsendersV1**](ObjectBankaccountAPI.md#BankaccountGetCommunicationsendersV1) | **Get** /1/object/bankaccount/{pkiBankaccountID}/getCommunicationsenders | Retrieve Communication senders
[**BankaccountImportIntoEDMV1**](ObjectBankaccountAPI.md#BankaccountImportIntoEDMV1) | **Post** /1/object/bankaccount/{pkiBankaccountID}/importIntoEDM | Import attachments into the Bankaccount



## BankaccountBatchDownloadV1

> *os.File BankaccountBatchDownloadV1(ctx, pkiBankaccountID).BankaccountBatchDownloadV1Request(bankaccountBatchDownloadV1Request).Execute()

Download multiples attachments from a Bankaccount

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
	pkiBankaccountID := int32(56) // int32 | 
	bankaccountBatchDownloadV1Request := *openapiclient.NewBankaccountBatchDownloadV1Request([]int32{int32(1)}) // BankaccountBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBankaccountAPI.BankaccountBatchDownloadV1(context.Background(), pkiBankaccountID).BankaccountBatchDownloadV1Request(bankaccountBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBankaccountAPI.BankaccountBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BankaccountBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectBankaccountAPI.BankaccountBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBankaccountID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBankaccountBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bankaccountBatchDownloadV1Request** | [**BankaccountBatchDownloadV1Request**](BankaccountBatchDownloadV1Request.md) |  | 

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


## BankaccountGetAttachmentsV1

> BankaccountGetAttachmentsV1Response BankaccountGetAttachmentsV1(ctx, pkiBankaccountID).Execute()

Retrieve Bankaccount's attachments

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
	pkiBankaccountID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBankaccountAPI.BankaccountGetAttachmentsV1(context.Background(), pkiBankaccountID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBankaccountAPI.BankaccountGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BankaccountGetAttachmentsV1`: BankaccountGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBankaccountAPI.BankaccountGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBankaccountID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBankaccountGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BankaccountGetAttachmentsV1Response**](BankaccountGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankaccountGetAutocompleteV2

> BankaccountGetAutocompleteV2Response BankaccountGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Bankaccounts and IDs



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
	sSelector := "sSelector_example" // string | The type of Bankaccounts to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBankaccountAPI.BankaccountGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBankaccountAPI.BankaccountGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BankaccountGetAutocompleteV2`: BankaccountGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBankaccountAPI.BankaccountGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Bankaccounts to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiBankaccountGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**BankaccountGetAutocompleteV2Response**](BankaccountGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankaccountGetCommunicationCountV1

> BankaccountGetCommunicationCountV1Response BankaccountGetCommunicationCountV1(ctx, pkiBankaccountID).Execute()

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
	pkiBankaccountID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBankaccountAPI.BankaccountGetCommunicationCountV1(context.Background(), pkiBankaccountID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBankaccountAPI.BankaccountGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BankaccountGetCommunicationCountV1`: BankaccountGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBankaccountAPI.BankaccountGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBankaccountID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBankaccountGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BankaccountGetCommunicationCountV1Response**](BankaccountGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankaccountGetCommunicationListV1

> BankaccountGetCommunicationListV1Response BankaccountGetCommunicationListV1(ctx, pkiBankaccountID).Execute()

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
	pkiBankaccountID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBankaccountAPI.BankaccountGetCommunicationListV1(context.Background(), pkiBankaccountID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBankaccountAPI.BankaccountGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BankaccountGetCommunicationListV1`: BankaccountGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBankaccountAPI.BankaccountGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBankaccountID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBankaccountGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BankaccountGetCommunicationListV1Response**](BankaccountGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankaccountGetCommunicationrecipientsV1

> BankaccountGetCommunicationrecipientsV1Response BankaccountGetCommunicationrecipientsV1(ctx, pkiBankaccountID).Execute()

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
	pkiBankaccountID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBankaccountAPI.BankaccountGetCommunicationrecipientsV1(context.Background(), pkiBankaccountID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBankaccountAPI.BankaccountGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BankaccountGetCommunicationrecipientsV1`: BankaccountGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBankaccountAPI.BankaccountGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBankaccountID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBankaccountGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BankaccountGetCommunicationrecipientsV1Response**](BankaccountGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankaccountGetCommunicationsendersV1

> BankaccountGetCommunicationsendersV1Response BankaccountGetCommunicationsendersV1(ctx, pkiBankaccountID).Execute()

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
	pkiBankaccountID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBankaccountAPI.BankaccountGetCommunicationsendersV1(context.Background(), pkiBankaccountID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBankaccountAPI.BankaccountGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BankaccountGetCommunicationsendersV1`: BankaccountGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBankaccountAPI.BankaccountGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBankaccountID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBankaccountGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BankaccountGetCommunicationsendersV1Response**](BankaccountGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankaccountImportIntoEDMV1

> BankaccountImportIntoEDMV1Response BankaccountImportIntoEDMV1(ctx, pkiBankaccountID).BankaccountImportIntoEDMV1Request(bankaccountImportIntoEDMV1Request).Execute()

Import attachments into the Bankaccount

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
	pkiBankaccountID := int32(56) // int32 | 
	bankaccountImportIntoEDMV1Request := *openapiclient.NewBankaccountImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // BankaccountImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBankaccountAPI.BankaccountImportIntoEDMV1(context.Background(), pkiBankaccountID).BankaccountImportIntoEDMV1Request(bankaccountImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBankaccountAPI.BankaccountImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BankaccountImportIntoEDMV1`: BankaccountImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBankaccountAPI.BankaccountImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBankaccountID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBankaccountImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bankaccountImportIntoEDMV1Request** | [**BankaccountImportIntoEDMV1Request**](BankaccountImportIntoEDMV1Request.md) |  | 

### Return type

[**BankaccountImportIntoEDMV1Response**](BankaccountImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

