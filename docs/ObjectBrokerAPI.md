# eZmaxAPI\ObjectBrokerAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrokerBatchDownloadV1**](ObjectBrokerAPI.md#BrokerBatchDownloadV1) | **Post** /1/object/broker/{pkiBrokerID}/batchDownload | Download multiples attachments from a Broker
[**BrokerGetAttachmentsV1**](ObjectBrokerAPI.md#BrokerGetAttachmentsV1) | **Get** /1/object/broker/{pkiBrokerID}/getAttachments | Retrieve Broker&#39;s attachments
[**BrokerGetAutocompleteV2**](ObjectBrokerAPI.md#BrokerGetAutocompleteV2) | **Get** /2/object/broker/getAutocomplete/{sSelector} | Retrieve Brokers and IDs
[**BrokerGetCommunicationCountV1**](ObjectBrokerAPI.md#BrokerGetCommunicationCountV1) | **Get** /1/object/broker/{pkiBrokerID}/getCommunicationCount | Retrieve Communication count
[**BrokerGetCommunicationListV1**](ObjectBrokerAPI.md#BrokerGetCommunicationListV1) | **Get** /1/object/broker/{pkiBrokerID}/getCommunicationList | Retrieve Communication list
[**BrokerGetCommunicationrecipientsV1**](ObjectBrokerAPI.md#BrokerGetCommunicationrecipientsV1) | **Get** /1/object/broker/{pkiBrokerID}/getCommunicationrecipients | Retrieve Communication recipients
[**BrokerGetCommunicationsendersV1**](ObjectBrokerAPI.md#BrokerGetCommunicationsendersV1) | **Get** /1/object/broker/{pkiBrokerID}/getCommunicationsenders | Retrieve Communication senders
[**BrokerGetListV1**](ObjectBrokerAPI.md#BrokerGetListV1) | **Get** /1/object/broker/getList | Retrieve Broker list
[**BrokerImportIntoEDMV1**](ObjectBrokerAPI.md#BrokerImportIntoEDMV1) | **Post** /1/object/broker/{pkiBrokerID}/importIntoEDM | Import attachments into the Broker



## BrokerBatchDownloadV1

> *os.File BrokerBatchDownloadV1(ctx, pkiBrokerID).BrokerBatchDownloadV1Request(brokerBatchDownloadV1Request).Execute()

Download multiples attachments from a Broker

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
	pkiBrokerID := int32(56) // int32 | 
	brokerBatchDownloadV1Request := *openapiclient.NewBrokerBatchDownloadV1Request([]int32{int32(1)}) // BrokerBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBrokerAPI.BrokerBatchDownloadV1(context.Background(), pkiBrokerID).BrokerBatchDownloadV1Request(brokerBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBrokerAPI.BrokerBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrokerBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectBrokerAPI.BrokerBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrokerBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **brokerBatchDownloadV1Request** | [**BrokerBatchDownloadV1Request**](BrokerBatchDownloadV1Request.md) |  | 

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


## BrokerGetAttachmentsV1

> BrokerGetAttachmentsV1Response BrokerGetAttachmentsV1(ctx, pkiBrokerID).Execute()

Retrieve Broker's attachments

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
	pkiBrokerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBrokerAPI.BrokerGetAttachmentsV1(context.Background(), pkiBrokerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBrokerAPI.BrokerGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrokerGetAttachmentsV1`: BrokerGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBrokerAPI.BrokerGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrokerGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BrokerGetAttachmentsV1Response**](BrokerGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrokerGetAutocompleteV2

> BrokerGetAutocompleteV2Response BrokerGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Brokers and IDs



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
	sSelector := "sSelector_example" // string | The type of Brokers to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBrokerAPI.BrokerGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBrokerAPI.BrokerGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrokerGetAutocompleteV2`: BrokerGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBrokerAPI.BrokerGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Brokers to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrokerGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**BrokerGetAutocompleteV2Response**](BrokerGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrokerGetCommunicationCountV1

> BrokerGetCommunicationCountV1Response BrokerGetCommunicationCountV1(ctx, pkiBrokerID).Execute()

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
	pkiBrokerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBrokerAPI.BrokerGetCommunicationCountV1(context.Background(), pkiBrokerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBrokerAPI.BrokerGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrokerGetCommunicationCountV1`: BrokerGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBrokerAPI.BrokerGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrokerGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BrokerGetCommunicationCountV1Response**](BrokerGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrokerGetCommunicationListV1

> BrokerGetCommunicationListV1Response BrokerGetCommunicationListV1(ctx, pkiBrokerID).Execute()

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
	pkiBrokerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBrokerAPI.BrokerGetCommunicationListV1(context.Background(), pkiBrokerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBrokerAPI.BrokerGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrokerGetCommunicationListV1`: BrokerGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBrokerAPI.BrokerGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrokerGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BrokerGetCommunicationListV1Response**](BrokerGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrokerGetCommunicationrecipientsV1

> BrokerGetCommunicationrecipientsV1Response BrokerGetCommunicationrecipientsV1(ctx, pkiBrokerID).Execute()

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
	pkiBrokerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBrokerAPI.BrokerGetCommunicationrecipientsV1(context.Background(), pkiBrokerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBrokerAPI.BrokerGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrokerGetCommunicationrecipientsV1`: BrokerGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBrokerAPI.BrokerGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrokerGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BrokerGetCommunicationrecipientsV1Response**](BrokerGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrokerGetCommunicationsendersV1

> BrokerGetCommunicationsendersV1Response BrokerGetCommunicationsendersV1(ctx, pkiBrokerID).Execute()

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
	pkiBrokerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBrokerAPI.BrokerGetCommunicationsendersV1(context.Background(), pkiBrokerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBrokerAPI.BrokerGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrokerGetCommunicationsendersV1`: BrokerGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBrokerAPI.BrokerGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrokerGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BrokerGetCommunicationsendersV1Response**](BrokerGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrokerGetListV1

> BrokerGetListV1Response BrokerGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Broker list



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
	resp, r, err := apiClient.ObjectBrokerAPI.BrokerGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBrokerAPI.BrokerGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrokerGetListV1`: BrokerGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBrokerAPI.BrokerGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBrokerGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**BrokerGetListV1Response**](BrokerGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrokerImportIntoEDMV1

> BrokerImportIntoEDMV1Response BrokerImportIntoEDMV1(ctx, pkiBrokerID).BrokerImportIntoEDMV1Request(brokerImportIntoEDMV1Request).Execute()

Import attachments into the Broker



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
	pkiBrokerID := int32(56) // int32 | 
	brokerImportIntoEDMV1Request := *openapiclient.NewBrokerImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // BrokerImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBrokerAPI.BrokerImportIntoEDMV1(context.Background(), pkiBrokerID).BrokerImportIntoEDMV1Request(brokerImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBrokerAPI.BrokerImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrokerImportIntoEDMV1`: BrokerImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBrokerAPI.BrokerImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrokerImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **brokerImportIntoEDMV1Request** | [**BrokerImportIntoEDMV1Request**](BrokerImportIntoEDMV1Request.md) |  | 

### Return type

[**BrokerImportIntoEDMV1Response**](BrokerImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

