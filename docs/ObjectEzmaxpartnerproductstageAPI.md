# eZmaxAPI\ObjectEzmaxpartnerproductstageAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzmaxpartnerproductstageGetEzmaxpartnerproductstagepermissionsV1**](ObjectEzmaxpartnerproductstageAPI.md#EzmaxpartnerproductstageGetEzmaxpartnerproductstagepermissionsV1) | **Get** /1/object/ezmaxpartnerproductstage/{pkiEzmaxpartnerproductstageID}/getEzmaxpartnerproductstagepermissions | Retrieve an existing Ezmaxpartnerproductstage&#39;s Ezmaxpartnerproductstagepermissions
[**EzmaxpartnerproductstageGetEzmaxpartnerproductstagewebhooksV1**](ObjectEzmaxpartnerproductstageAPI.md#EzmaxpartnerproductstageGetEzmaxpartnerproductstagewebhooksV1) | **Get** /1/object/ezmaxpartnerproductstage/{pkiEzmaxpartnerproductstageID}/getEzmaxpartnerproductstagewebhooks | Retrieve an existing Ezmaxpartnerproductstage&#39;s Ezmaxpartnerproductstagewebhooks
[**EzmaxpartnerproductstageGetListV1**](ObjectEzmaxpartnerproductstageAPI.md#EzmaxpartnerproductstageGetListV1) | **Get** /1/object/ezmaxpartnerproductstage/getList | Retrieve Ezmaxpartnerproductstage list
[**EzmaxpartnerproductstageGetObjectV2**](ObjectEzmaxpartnerproductstageAPI.md#EzmaxpartnerproductstageGetObjectV2) | **Get** /2/object/ezmaxpartnerproductstage/{pkiEzmaxpartnerproductstageID} | Retrieve an existing Ezmaxpartnerproductstage
[**EzmaxpartnerproductstageSubscribeV1**](ObjectEzmaxpartnerproductstageAPI.md#EzmaxpartnerproductstageSubscribeV1) | **Post** /1/object/ezmaxpartnerproductstage/{pkiEzmaxpartnerproductstageID}/subscribe | Subscribe to the Ezmaxpartnerproductstage
[**EzmaxpartnerproductstageUnsubscribeV1**](ObjectEzmaxpartnerproductstageAPI.md#EzmaxpartnerproductstageUnsubscribeV1) | **Post** /1/object/ezmaxpartnerproductstage/{pkiEzmaxpartnerproductstageID}/unsubscribe | Unsubscribe to the Ezmaxpartnerproductstage



## EzmaxpartnerproductstageGetEzmaxpartnerproductstagepermissionsV1

> EzmaxpartnerproductstageGetEzmaxpartnerproductstagepermissionsV1Response EzmaxpartnerproductstageGetEzmaxpartnerproductstagepermissionsV1(ctx, pkiEzmaxpartnerproductstageID).Execute()

Retrieve an existing Ezmaxpartnerproductstage's Ezmaxpartnerproductstagepermissions

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
	pkiEzmaxpartnerproductstageID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageGetEzmaxpartnerproductstagepermissionsV1(context.Background(), pkiEzmaxpartnerproductstageID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageGetEzmaxpartnerproductstagepermissionsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxpartnerproductstageGetEzmaxpartnerproductstagepermissionsV1`: EzmaxpartnerproductstageGetEzmaxpartnerproductstagepermissionsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageGetEzmaxpartnerproductstagepermissionsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzmaxpartnerproductstageID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxpartnerproductstageGetEzmaxpartnerproductstagepermissionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzmaxpartnerproductstageGetEzmaxpartnerproductstagepermissionsV1Response**](EzmaxpartnerproductstageGetEzmaxpartnerproductstagepermissionsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzmaxpartnerproductstageGetEzmaxpartnerproductstagewebhooksV1

> EzmaxpartnerproductstageGetEzmaxpartnerproductstagewebhooksV1Response EzmaxpartnerproductstageGetEzmaxpartnerproductstagewebhooksV1(ctx, pkiEzmaxpartnerproductstageID).Execute()

Retrieve an existing Ezmaxpartnerproductstage's Ezmaxpartnerproductstagewebhooks

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
	pkiEzmaxpartnerproductstageID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageGetEzmaxpartnerproductstagewebhooksV1(context.Background(), pkiEzmaxpartnerproductstageID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageGetEzmaxpartnerproductstagewebhooksV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxpartnerproductstageGetEzmaxpartnerproductstagewebhooksV1`: EzmaxpartnerproductstageGetEzmaxpartnerproductstagewebhooksV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageGetEzmaxpartnerproductstagewebhooksV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzmaxpartnerproductstageID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxpartnerproductstageGetEzmaxpartnerproductstagewebhooksV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzmaxpartnerproductstageGetEzmaxpartnerproductstagewebhooksV1Response**](EzmaxpartnerproductstageGetEzmaxpartnerproductstagewebhooksV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzmaxpartnerproductstageGetListV1

> EzmaxpartnerproductstageGetListV1Response EzmaxpartnerproductstageGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Ezmaxpartnerproductstage list



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
	resp, r, err := apiClient.ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxpartnerproductstageGetListV1`: EzmaxpartnerproductstageGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxpartnerproductstageGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**EzmaxpartnerproductstageGetListV1Response**](EzmaxpartnerproductstageGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzmaxpartnerproductstageGetObjectV2

> EzmaxpartnerproductstageGetObjectV2Response EzmaxpartnerproductstageGetObjectV2(ctx, pkiEzmaxpartnerproductstageID).Execute()

Retrieve an existing Ezmaxpartnerproductstage



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
	pkiEzmaxpartnerproductstageID := int32(56) // int32 | The unique ID of the Ezmaxpartnerproductstage

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageGetObjectV2(context.Background(), pkiEzmaxpartnerproductstageID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxpartnerproductstageGetObjectV2`: EzmaxpartnerproductstageGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzmaxpartnerproductstageID** | **int32** | The unique ID of the Ezmaxpartnerproductstage | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxpartnerproductstageGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzmaxpartnerproductstageGetObjectV2Response**](EzmaxpartnerproductstageGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzmaxpartnerproductstageSubscribeV1

> EzmaxpartnerproductstageSubscribeV1Response EzmaxpartnerproductstageSubscribeV1(ctx, pkiEzmaxpartnerproductstageID).Body(body).Execute()

Subscribe to the Ezmaxpartnerproductstage

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
	pkiEzmaxpartnerproductstageID := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageSubscribeV1(context.Background(), pkiEzmaxpartnerproductstageID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageSubscribeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxpartnerproductstageSubscribeV1`: EzmaxpartnerproductstageSubscribeV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageSubscribeV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzmaxpartnerproductstageID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxpartnerproductstageSubscribeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**EzmaxpartnerproductstageSubscribeV1Response**](EzmaxpartnerproductstageSubscribeV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzmaxpartnerproductstageUnsubscribeV1

> EzmaxpartnerproductstageUnsubscribeV1Response EzmaxpartnerproductstageUnsubscribeV1(ctx, pkiEzmaxpartnerproductstageID).Body(body).Execute()

Unsubscribe to the Ezmaxpartnerproductstage

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
	pkiEzmaxpartnerproductstageID := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageUnsubscribeV1(context.Background(), pkiEzmaxpartnerproductstageID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageUnsubscribeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzmaxpartnerproductstageUnsubscribeV1`: EzmaxpartnerproductstageUnsubscribeV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzmaxpartnerproductstageAPI.EzmaxpartnerproductstageUnsubscribeV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzmaxpartnerproductstageID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzmaxpartnerproductstageUnsubscribeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**EzmaxpartnerproductstageUnsubscribeV1Response**](EzmaxpartnerproductstageUnsubscribeV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

