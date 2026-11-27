# eZmaxAPI\ObjectAgentAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentBatchDownloadV1**](ObjectAgentAPI.md#AgentBatchDownloadV1) | **Post** /1/object/agent/{pkiAgentID}/batchDownload | Download multiples attachments from a Agent
[**AgentGetAttachmentsV1**](ObjectAgentAPI.md#AgentGetAttachmentsV1) | **Get** /1/object/agent/{pkiAgentID}/getAttachments | Retrieve Agent&#39;s attachments
[**AgentGetAutocompleteV2**](ObjectAgentAPI.md#AgentGetAutocompleteV2) | **Get** /2/object/agent/getAutocomplete/{sSelector} | Retrieve Agents and IDs
[**AgentGetCommunicationCountV1**](ObjectAgentAPI.md#AgentGetCommunicationCountV1) | **Get** /1/object/agent/{pkiAgentID}/getCommunicationCount | Retrieve Communication count
[**AgentGetCommunicationListV1**](ObjectAgentAPI.md#AgentGetCommunicationListV1) | **Get** /1/object/agent/{pkiAgentID}/getCommunicationList | Retrieve Communication list
[**AgentGetCommunicationrecipientsV1**](ObjectAgentAPI.md#AgentGetCommunicationrecipientsV1) | **Get** /1/object/agent/{pkiAgentID}/getCommunicationrecipients | Retrieve Communication recipients
[**AgentGetCommunicationsendersV1**](ObjectAgentAPI.md#AgentGetCommunicationsendersV1) | **Get** /1/object/agent/{pkiAgentID}/getCommunicationsenders | Retrieve Communication senders
[**AgentGetListV1**](ObjectAgentAPI.md#AgentGetListV1) | **Get** /1/object/agent/getList | Retrieve Agent list
[**AgentImportIntoEDMV1**](ObjectAgentAPI.md#AgentImportIntoEDMV1) | **Post** /1/object/agent/{pkiAgentID}/importIntoEDM | Import attachments into the Agent



## AgentBatchDownloadV1

> *os.File AgentBatchDownloadV1(ctx, pkiAgentID).AgentBatchDownloadV1Request(agentBatchDownloadV1Request).Execute()

Download multiples attachments from a Agent

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
	pkiAgentID := int32(56) // int32 | 
	agentBatchDownloadV1Request := *openapiclient.NewAgentBatchDownloadV1Request([]int32{int32(1)}) // AgentBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAgentAPI.AgentBatchDownloadV1(context.Background(), pkiAgentID).AgentBatchDownloadV1Request(agentBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAgentAPI.AgentBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectAgentAPI.AgentBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAgentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentBatchDownloadV1Request** | [**AgentBatchDownloadV1Request**](AgentBatchDownloadV1Request.md) |  | 

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


## AgentGetAttachmentsV1

> AgentGetAttachmentsV1Response AgentGetAttachmentsV1(ctx, pkiAgentID).Execute()

Retrieve Agent's attachments

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
	pkiAgentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAgentAPI.AgentGetAttachmentsV1(context.Background(), pkiAgentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAgentAPI.AgentGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentGetAttachmentsV1`: AgentGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAgentAPI.AgentGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAgentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentGetAttachmentsV1Response**](AgentGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGetAutocompleteV2

> AgentGetAutocompleteV2Response AgentGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Agents and IDs



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
	sSelector := "sSelector_example" // string | The type of Agents to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAgentAPI.AgentGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAgentAPI.AgentGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentGetAutocompleteV2`: AgentGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAgentAPI.AgentGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Agents to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**AgentGetAutocompleteV2Response**](AgentGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGetCommunicationCountV1

> AgentGetCommunicationCountV1Response AgentGetCommunicationCountV1(ctx, pkiAgentID).Execute()

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
	pkiAgentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAgentAPI.AgentGetCommunicationCountV1(context.Background(), pkiAgentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAgentAPI.AgentGetCommunicationCountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentGetCommunicationCountV1`: AgentGetCommunicationCountV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAgentAPI.AgentGetCommunicationCountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAgentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentGetCommunicationCountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentGetCommunicationCountV1Response**](AgentGetCommunicationCountV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGetCommunicationListV1

> AgentGetCommunicationListV1Response AgentGetCommunicationListV1(ctx, pkiAgentID).Execute()

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
	pkiAgentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAgentAPI.AgentGetCommunicationListV1(context.Background(), pkiAgentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAgentAPI.AgentGetCommunicationListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentGetCommunicationListV1`: AgentGetCommunicationListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAgentAPI.AgentGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAgentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentGetCommunicationListV1Response**](AgentGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGetCommunicationrecipientsV1

> AgentGetCommunicationrecipientsV1Response AgentGetCommunicationrecipientsV1(ctx, pkiAgentID).Execute()

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
	pkiAgentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAgentAPI.AgentGetCommunicationrecipientsV1(context.Background(), pkiAgentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAgentAPI.AgentGetCommunicationrecipientsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentGetCommunicationrecipientsV1`: AgentGetCommunicationrecipientsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAgentAPI.AgentGetCommunicationrecipientsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAgentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentGetCommunicationrecipientsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentGetCommunicationrecipientsV1Response**](AgentGetCommunicationrecipientsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGetCommunicationsendersV1

> AgentGetCommunicationsendersV1Response AgentGetCommunicationsendersV1(ctx, pkiAgentID).Execute()

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
	pkiAgentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAgentAPI.AgentGetCommunicationsendersV1(context.Background(), pkiAgentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAgentAPI.AgentGetCommunicationsendersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentGetCommunicationsendersV1`: AgentGetCommunicationsendersV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAgentAPI.AgentGetCommunicationsendersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAgentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentGetCommunicationsendersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentGetCommunicationsendersV1Response**](AgentGetCommunicationsendersV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGetListV1

> AgentGetListV1Response AgentGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Agent list



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
	resp, r, err := apiClient.ObjectAgentAPI.AgentGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAgentAPI.AgentGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentGetListV1`: AgentGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAgentAPI.AgentGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**AgentGetListV1Response**](AgentGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentImportIntoEDMV1

> AgentImportIntoEDMV1Response AgentImportIntoEDMV1(ctx, pkiAgentID).AgentImportIntoEDMV1Request(agentImportIntoEDMV1Request).Execute()

Import attachments into the Agent



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
	pkiAgentID := int32(56) // int32 | 
	agentImportIntoEDMV1Request := *openapiclient.NewAgentImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // AgentImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAgentAPI.AgentImportIntoEDMV1(context.Background(), pkiAgentID).AgentImportIntoEDMV1Request(agentImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAgentAPI.AgentImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentImportIntoEDMV1`: AgentImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAgentAPI.AgentImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAgentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentImportIntoEDMV1Request** | [**AgentImportIntoEDMV1Request**](AgentImportIntoEDMV1Request.md) |  | 

### Return type

[**AgentImportIntoEDMV1Response**](AgentImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

