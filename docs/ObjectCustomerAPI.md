# eZmaxAPI\ObjectCustomerAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerBatchDownloadV1**](ObjectCustomerAPI.md#CustomerBatchDownloadV1) | **Post** /1/object/customer/{pkiCustomerID}/batchDownload | Download multiples attachments from a Customer
[**CustomerGetAttachmentsV1**](ObjectCustomerAPI.md#CustomerGetAttachmentsV1) | **Get** /1/object/customer/{pkiCustomerID}/getAttachments | Retrieve Customer&#39;s attachments
[**CustomerGetAutocompleteV2**](ObjectCustomerAPI.md#CustomerGetAutocompleteV2) | **Get** /2/object/customer/getAutocomplete/{sSelector} | Retrieve Customers and IDs
[**CustomerGetCommunicationCountV1**](ObjectCustomerAPI.md#CustomerGetCommunicationCountV1) | **Get** /1/object/customer/{pkiCustomerID}/getCommunicationCount | Retrieve Communication count
[**CustomerGetCommunicationListV1**](ObjectCustomerAPI.md#CustomerGetCommunicationListV1) | **Get** /1/object/customer/{pkiCustomerID}/getCommunicationList | Retrieve Communication list
[**CustomerGetCommunicationrecipientsV1**](ObjectCustomerAPI.md#CustomerGetCommunicationrecipientsV1) | **Get** /1/object/customer/{pkiCustomerID}/getCommunicationrecipients | Retrieve Communication recipients
[**CustomerGetCommunicationsendersV1**](ObjectCustomerAPI.md#CustomerGetCommunicationsendersV1) | **Get** /1/object/customer/{pkiCustomerID}/getCommunicationsenders | Retrieve Communication senders
[**CustomerGetObjectV2**](ObjectCustomerAPI.md#CustomerGetObjectV2) | **Get** /2/object/customer/{pkiCustomerID} | Retrieve an existing Customer
[**CustomerImportIntoEDMV1**](ObjectCustomerAPI.md#CustomerImportIntoEDMV1) | **Post** /1/object/customer/{pkiCustomerID}/importIntoEDM | Import attachments into the Customer



## CustomerBatchDownloadV1

> *os.File CustomerBatchDownloadV1(ctx, pkiCustomerID).CustomerBatchDownloadV1Request(customerBatchDownloadV1Request).Execute()

Download multiples attachments from a Customer

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
	pkiCustomerID := int32(56) // int32 | 
	customerBatchDownloadV1Request := *openapiclient.NewCustomerBatchDownloadV1Request([]int32{int32(1)}) // CustomerBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCustomerAPI.CustomerBatchDownloadV1(context.Background(), pkiCustomerID).CustomerBatchDownloadV1Request(customerBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCustomerAPI.CustomerBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectCustomerAPI.CustomerBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCustomerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerBatchDownloadV1Request** | [**CustomerBatchDownloadV1Request**](CustomerBatchDownloadV1Request.md) |  | 

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


## CustomerGetAttachmentsV1

> CustomerGetAttachmentsV1Response CustomerGetAttachmentsV1(ctx, pkiCustomerID).Execute()

Retrieve Customer's attachments

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
	pkiCustomerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCustomerAPI.CustomerGetAttachmentsV1(context.Background(), pkiCustomerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCustomerAPI.CustomerGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerGetAttachmentsV1`: CustomerGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCustomerAPI.CustomerGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCustomerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerGetAttachmentsV1Response**](CustomerGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerGetAutocompleteV2

> CustomerGetAutocompleteV2Response CustomerGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Customers and IDs



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
	sSelector := "sSelector_example" // string | The type of Customers to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCustomerAPI.CustomerGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCustomerAPI.CustomerGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerGetAutocompleteV2`: CustomerGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCustomerAPI.CustomerGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Customers to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**CustomerGetAutocompleteV2Response**](CustomerGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerGetCommunicationCountV1

> CustomerGetCommunicationCountV1Response CustomerGetCommunicationCountV1(ctx, pkiCustomerID).Execute()

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
	pkiCustomerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCustomerAPI.CustomerGetCommunicationCountV1(context.Background(), pkiCustomerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCustomerAPI.CustomerGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerGetCommunicationCountV1`: CustomerGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCustomerAPI.CustomerGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCustomerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerGetCommunicationCountV1Response**](CustomerGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerGetCommunicationListV1

> CustomerGetCommunicationListV1Response CustomerGetCommunicationListV1(ctx, pkiCustomerID).Execute()

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
	pkiCustomerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCustomerAPI.CustomerGetCommunicationListV1(context.Background(), pkiCustomerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCustomerAPI.CustomerGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerGetCommunicationListV1`: CustomerGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCustomerAPI.CustomerGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCustomerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerGetCommunicationListV1Response**](CustomerGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerGetCommunicationrecipientsV1

> CustomerGetCommunicationrecipientsV1Response CustomerGetCommunicationrecipientsV1(ctx, pkiCustomerID).Execute()

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
	pkiCustomerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCustomerAPI.CustomerGetCommunicationrecipientsV1(context.Background(), pkiCustomerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCustomerAPI.CustomerGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerGetCommunicationrecipientsV1`: CustomerGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCustomerAPI.CustomerGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCustomerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerGetCommunicationrecipientsV1Response**](CustomerGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerGetCommunicationsendersV1

> CustomerGetCommunicationsendersV1Response CustomerGetCommunicationsendersV1(ctx, pkiCustomerID).Execute()

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
	pkiCustomerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCustomerAPI.CustomerGetCommunicationsendersV1(context.Background(), pkiCustomerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCustomerAPI.CustomerGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerGetCommunicationsendersV1`: CustomerGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCustomerAPI.CustomerGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCustomerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerGetCommunicationsendersV1Response**](CustomerGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerGetObjectV2

> CustomerGetObjectV2Response CustomerGetObjectV2(ctx, pkiCustomerID).Execute()

Retrieve an existing Customer



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
	pkiCustomerID := int32(56) // int32 | The unique ID of the Customer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCustomerAPI.CustomerGetObjectV2(context.Background(), pkiCustomerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCustomerAPI.CustomerGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerGetObjectV2`: CustomerGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCustomerAPI.CustomerGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCustomerID** | **int32** | The unique ID of the Customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerGetObjectV2Response**](CustomerGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerImportIntoEDMV1

> CustomerImportIntoEDMV1Response CustomerImportIntoEDMV1(ctx, pkiCustomerID).CustomerImportIntoEDMV1Request(customerImportIntoEDMV1Request).Execute()

Import attachments into the Customer



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
	pkiCustomerID := int32(56) // int32 | 
	customerImportIntoEDMV1Request := *openapiclient.NewCustomerImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // CustomerImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCustomerAPI.CustomerImportIntoEDMV1(context.Background(), pkiCustomerID).CustomerImportIntoEDMV1Request(customerImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCustomerAPI.CustomerImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerImportIntoEDMV1`: CustomerImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCustomerAPI.CustomerImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCustomerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerImportIntoEDMV1Request** | [**CustomerImportIntoEDMV1Request**](CustomerImportIntoEDMV1Request.md) |  | 

### Return type

[**CustomerImportIntoEDMV1Response**](CustomerImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

