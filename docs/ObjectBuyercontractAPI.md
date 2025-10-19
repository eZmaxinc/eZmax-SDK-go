# eZmaxAPI\ObjectBuyercontractAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuyercontractGetCommunicationCountV1**](ObjectBuyercontractAPI.md#BuyercontractGetCommunicationCountV1) | **Get** /1/object/buyercontract/{pkiBuyercontractID}/getCommunicationCount | Retrieve Communication count
[**BuyercontractGetCommunicationListV1**](ObjectBuyercontractAPI.md#BuyercontractGetCommunicationListV1) | **Get** /1/object/buyercontract/{pkiBuyercontractID}/getCommunicationList | Retrieve Communication list
[**BuyercontractGetCommunicationrecipientsV1**](ObjectBuyercontractAPI.md#BuyercontractGetCommunicationrecipientsV1) | **Get** /1/object/buyercontract/{pkiBuyercontractID}/getCommunicationrecipients | Retrieve Buyercontract&#39;s Communicationrecipient
[**BuyercontractGetCommunicationsendersV1**](ObjectBuyercontractAPI.md#BuyercontractGetCommunicationsendersV1) | **Get** /1/object/buyercontract/{pkiBuyercontractID}/getCommunicationsenders | Retrieve Buyercontract&#39;s Communicationsender
[**BuyercontractGetListV1**](ObjectBuyercontractAPI.md#BuyercontractGetListV1) | **Get** /1/object/buyercontract/getList | Retrieve Buyercontract list
[**BuyercontractImportIntoEDMV1**](ObjectBuyercontractAPI.md#BuyercontractImportIntoEDMV1) | **Post** /1/object/buyercontract/{pkiBuyercontractID}/importIntoEDM | Import attachments into the Buyercontract



## BuyercontractGetCommunicationCountV1

> BuyercontractGetCommunicationCountV1Response BuyercontractGetCommunicationCountV1(ctx, pkiBuyercontractID).Execute()

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
	pkiBuyercontractID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBuyercontractAPI.BuyercontractGetCommunicationCountV1(context.Background(), pkiBuyercontractID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBuyercontractAPI.BuyercontractGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuyercontractGetCommunicationCountV1`: BuyercontractGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBuyercontractAPI.BuyercontractGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBuyercontractID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuyercontractGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuyercontractGetCommunicationCountV1Response**](BuyercontractGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuyercontractGetCommunicationListV1

> BuyercontractGetCommunicationListV1Response BuyercontractGetCommunicationListV1(ctx, pkiBuyercontractID).Execute()

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
	pkiBuyercontractID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBuyercontractAPI.BuyercontractGetCommunicationListV1(context.Background(), pkiBuyercontractID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBuyercontractAPI.BuyercontractGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuyercontractGetCommunicationListV1`: BuyercontractGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBuyercontractAPI.BuyercontractGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBuyercontractID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuyercontractGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuyercontractGetCommunicationListV1Response**](BuyercontractGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuyercontractGetCommunicationrecipientsV1

> BuyercontractGetCommunicationrecipientsV1Response BuyercontractGetCommunicationrecipientsV1(ctx, pkiBuyercontractID).Execute()

Retrieve Buyercontract's Communicationrecipient



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
	pkiBuyercontractID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBuyercontractAPI.BuyercontractGetCommunicationrecipientsV1(context.Background(), pkiBuyercontractID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBuyercontractAPI.BuyercontractGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuyercontractGetCommunicationrecipientsV1`: BuyercontractGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBuyercontractAPI.BuyercontractGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBuyercontractID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuyercontractGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuyercontractGetCommunicationrecipientsV1Response**](BuyercontractGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuyercontractGetCommunicationsendersV1

> BuyercontractGetCommunicationsendersV1Response BuyercontractGetCommunicationsendersV1(ctx, pkiBuyercontractID).Execute()

Retrieve Buyercontract's Communicationsender



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
	pkiBuyercontractID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBuyercontractAPI.BuyercontractGetCommunicationsendersV1(context.Background(), pkiBuyercontractID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBuyercontractAPI.BuyercontractGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuyercontractGetCommunicationsendersV1`: BuyercontractGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBuyercontractAPI.BuyercontractGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBuyercontractID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuyercontractGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuyercontractGetCommunicationsendersV1Response**](BuyercontractGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuyercontractGetListV1

> BuyercontractGetListV1Response BuyercontractGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Buyercontract list



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
	resp, r, err := apiClient.ObjectBuyercontractAPI.BuyercontractGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBuyercontractAPI.BuyercontractGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuyercontractGetListV1`: BuyercontractGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBuyercontractAPI.BuyercontractGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuyercontractGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**BuyercontractGetListV1Response**](BuyercontractGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuyercontractImportIntoEDMV1

> BuyercontractImportIntoEDMV1Response BuyercontractImportIntoEDMV1(ctx, pkiBuyercontractID).BuyercontractImportIntoEDMV1Request(buyercontractImportIntoEDMV1Request).Execute()

Import attachments into the Buyercontract



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
	pkiBuyercontractID := int32(56) // int32 | 
	buyercontractImportIntoEDMV1Request := *openapiclient.NewBuyercontractImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // BuyercontractImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectBuyercontractAPI.BuyercontractImportIntoEDMV1(context.Background(), pkiBuyercontractID).BuyercontractImportIntoEDMV1Request(buyercontractImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectBuyercontractAPI.BuyercontractImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuyercontractImportIntoEDMV1`: BuyercontractImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectBuyercontractAPI.BuyercontractImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBuyercontractID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuyercontractImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buyercontractImportIntoEDMV1Request** | [**BuyercontractImportIntoEDMV1Request**](BuyercontractImportIntoEDMV1Request.md) |  | 

### Return type

[**BuyercontractImportIntoEDMV1Response**](BuyercontractImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

