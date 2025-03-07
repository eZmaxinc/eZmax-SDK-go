# eZmaxAPI\ObjectEzsignimportfolderAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignimportfolderDeleteObjectV1**](ObjectEzsignimportfolderAPI.md#EzsignimportfolderDeleteObjectV1) | **Delete** /1/object/ezsignimportfolder/{pkiEzsignimportfolderID} | Delete an existing Ezsignimportfolder
[**EzsignimportfolderGetListV1**](ObjectEzsignimportfolderAPI.md#EzsignimportfolderGetListV1) | **Get** /1/object/ezsignimportfolder/getList | Retrieve Ezsignimportfolder list
[**EzsignimportfolderGetObjectV2**](ObjectEzsignimportfolderAPI.md#EzsignimportfolderGetObjectV2) | **Get** /2/object/ezsignimportfolder/{pkiEzsignimportfolderID} | Retrieve an existing Ezsignimportfolder



## EzsignimportfolderDeleteObjectV1

> EzsignimportfolderDeleteObjectV1Response EzsignimportfolderDeleteObjectV1(ctx, pkiEzsignimportfolderID).Execute()

Delete an existing Ezsignimportfolder



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
	pkiEzsignimportfolderID := int32(56) // int32 | The unique ID of the Ezsignimportfolder

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignimportfolderAPI.EzsignimportfolderDeleteObjectV1(context.Background(), pkiEzsignimportfolderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignimportfolderAPI.EzsignimportfolderDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignimportfolderDeleteObjectV1`: EzsignimportfolderDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignimportfolderAPI.EzsignimportfolderDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignimportfolderID** | **int32** | The unique ID of the Ezsignimportfolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignimportfolderDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignimportfolderDeleteObjectV1Response**](EzsignimportfolderDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignimportfolderGetListV1

> EzsignimportfolderGetListV1Response EzsignimportfolderGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Ezsignimportfolder list



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
	resp, r, err := apiClient.ObjectEzsignimportfolderAPI.EzsignimportfolderGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignimportfolderAPI.EzsignimportfolderGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignimportfolderGetListV1`: EzsignimportfolderGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignimportfolderAPI.EzsignimportfolderGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignimportfolderGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**EzsignimportfolderGetListV1Response**](EzsignimportfolderGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignimportfolderGetObjectV2

> EzsignimportfolderGetObjectV2Response EzsignimportfolderGetObjectV2(ctx, pkiEzsignimportfolderID).Execute()

Retrieve an existing Ezsignimportfolder



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
	pkiEzsignimportfolderID := int32(56) // int32 | The unique ID of the Ezsignimportfolder

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignimportfolderAPI.EzsignimportfolderGetObjectV2(context.Background(), pkiEzsignimportfolderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignimportfolderAPI.EzsignimportfolderGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignimportfolderGetObjectV2`: EzsignimportfolderGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignimportfolderAPI.EzsignimportfolderGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignimportfolderID** | **int32** | The unique ID of the Ezsignimportfolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignimportfolderGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignimportfolderGetObjectV2Response**](EzsignimportfolderGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

