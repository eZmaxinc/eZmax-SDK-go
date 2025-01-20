# eZmaxAPI\ObjectEzsigntemplatepublicAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigntemplatepublicCreateEzsignfolderV1**](ObjectEzsigntemplatepublicAPI.md#EzsigntemplatepublicCreateEzsignfolderV1) | **Post** /1/object/ezsigntemplatepublic/createEzsignfolder | Create an Ezsignfolder
[**EzsigntemplatepublicCreateObjectV1**](ObjectEzsigntemplatepublicAPI.md#EzsigntemplatepublicCreateObjectV1) | **Post** /1/object/ezsigntemplatepublic | Create a new Ezsigntemplatepublic
[**EzsigntemplatepublicEditObjectV1**](ObjectEzsigntemplatepublicAPI.md#EzsigntemplatepublicEditObjectV1) | **Put** /1/object/ezsigntemplatepublic/{pkiEzsigntemplatepublicID} | Edit an existing Ezsigntemplatepublic
[**EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1**](ObjectEzsigntemplatepublicAPI.md#EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1) | **Post** /1/object/ezsigntemplatepublic/getEzsigntemplatepublicDetails | Retrieve the Ezsigntemplatepublic details
[**EzsigntemplatepublicGetFormsDataV1**](ObjectEzsigntemplatepublicAPI.md#EzsigntemplatepublicGetFormsDataV1) | **Get** /1/object/ezsigntemplatepublic/{pkiEzsigntemplatepublicID}/getFormsData | Retrieve an existing Ezsigntemplatepublic&#39;s forms data
[**EzsigntemplatepublicGetListV1**](ObjectEzsigntemplatepublicAPI.md#EzsigntemplatepublicGetListV1) | **Get** /1/object/ezsigntemplatepublic/getList | Retrieve Ezsigntemplatepublic list
[**EzsigntemplatepublicGetObjectV2**](ObjectEzsigntemplatepublicAPI.md#EzsigntemplatepublicGetObjectV2) | **Get** /2/object/ezsigntemplatepublic/{pkiEzsigntemplatepublicID} | Retrieve an existing Ezsigntemplatepublic
[**EzsigntemplatepublicResetLimitExceededCounterV1**](ObjectEzsigntemplatepublicAPI.md#EzsigntemplatepublicResetLimitExceededCounterV1) | **Post** /1/object/ezsigntemplatepublic/{pkiEzsigntemplatepublicID}/resetLimitExceededCounter | Reset the limit exceeded counter
[**EzsigntemplatepublicResetUrlV1**](ObjectEzsigntemplatepublicAPI.md#EzsigntemplatepublicResetUrlV1) | **Post** /1/object/ezsigntemplatepublic/{pkiEzsigntemplatepublicID}/resetUrl | Reset the Ezsigntemplatepublic url



## EzsigntemplatepublicCreateEzsignfolderV1

> EzsigntemplatepublicCreateEzsignfolderV1Response EzsigntemplatepublicCreateEzsignfolderV1(ctx).EzsigntemplatepublicCreateEzsignfolderV1Request(ezsigntemplatepublicCreateEzsignfolderV1Request).Execute()

Create an Ezsignfolder



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
	ezsigntemplatepublicCreateEzsignfolderV1Request := *openapiclient.NewEzsigntemplatepublicCreateEzsignfolderV1Request("demo", "6B29FC40-CA47-1067-B31D-00DD010662DA", []string{"http://www.website.com/avatar.jpg"}, []openapiclient.EzsignsignerRequestCompound{*openapiclient.NewEzsignsignerRequestCompound(int32(1), *openapiclient.NewEzsignsignerRequestCompoundContact("John", "Doe", int32(2)))}) // EzsigntemplatepublicCreateEzsignfolderV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicCreateEzsignfolderV1(context.Background()).EzsigntemplatepublicCreateEzsignfolderV1Request(ezsigntemplatepublicCreateEzsignfolderV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicCreateEzsignfolderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepublicCreateEzsignfolderV1`: EzsigntemplatepublicCreateEzsignfolderV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicCreateEzsignfolderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepublicCreateEzsignfolderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplatepublicCreateEzsignfolderV1Request** | [**EzsigntemplatepublicCreateEzsignfolderV1Request**](EzsigntemplatepublicCreateEzsignfolderV1Request.md) |  | 

### Return type

[**EzsigntemplatepublicCreateEzsignfolderV1Response**](EzsigntemplatepublicCreateEzsignfolderV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepublicCreateObjectV1

> EzsigntemplatepublicCreateObjectV1Response EzsigntemplatepublicCreateObjectV1(ctx).EzsigntemplatepublicCreateObjectV1Request(ezsigntemplatepublicCreateObjectV1Request).Execute()

Create a new Ezsigntemplatepublic



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
	ezsigntemplatepublicCreateObjectV1Request := *openapiclient.NewEzsigntemplatepublicCreateObjectV1Request([]openapiclient.EzsigntemplatepublicRequestCompound{*openapiclient.NewEzsigntemplatepublicRequestCompound(int32(5), int32(2), "Inscription form", true, "This is a note", openapiclient.Field-eEzsigntemplatepublicLimittype("Hour"), int32(10))}) // EzsigntemplatepublicCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicCreateObjectV1(context.Background()).EzsigntemplatepublicCreateObjectV1Request(ezsigntemplatepublicCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepublicCreateObjectV1`: EzsigntemplatepublicCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepublicCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplatepublicCreateObjectV1Request** | [**EzsigntemplatepublicCreateObjectV1Request**](EzsigntemplatepublicCreateObjectV1Request.md) |  | 

### Return type

[**EzsigntemplatepublicCreateObjectV1Response**](EzsigntemplatepublicCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepublicEditObjectV1

> CommonResponse EzsigntemplatepublicEditObjectV1(ctx, pkiEzsigntemplatepublicID).EzsigntemplatepublicEditObjectV1Request(ezsigntemplatepublicEditObjectV1Request).Execute()

Edit an existing Ezsigntemplatepublic



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
	pkiEzsigntemplatepublicID := int32(56) // int32 | The unique ID of the Ezsigntemplatepublic
	ezsigntemplatepublicEditObjectV1Request := *openapiclient.NewEzsigntemplatepublicEditObjectV1Request(*openapiclient.NewEzsigntemplatepublicRequestCompound(int32(5), int32(2), "Inscription form", true, "This is a note", openapiclient.Field-eEzsigntemplatepublicLimittype("Hour"), int32(10))) // EzsigntemplatepublicEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicEditObjectV1(context.Background(), pkiEzsigntemplatepublicID).EzsigntemplatepublicEditObjectV1Request(ezsigntemplatepublicEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepublicEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepublicID** | **int32** | The unique ID of the Ezsigntemplatepublic | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepublicEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplatepublicEditObjectV1Request** | [**EzsigntemplatepublicEditObjectV1Request**](EzsigntemplatepublicEditObjectV1Request.md) |  | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1

> EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1Response EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1(ctx).EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1Request(ezsigntemplatepublicGetEzsigntemplatepublicDetailsV1Request).Execute()

Retrieve the Ezsigntemplatepublic details



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
	ezsigntemplatepublicGetEzsigntemplatepublicDetailsV1Request := *openapiclient.NewEzsigntemplatepublicGetEzsigntemplatepublicDetailsV1Request("demo", "6B29FC40-CA47-1067-B31D-00DD010662DA") // EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1(context.Background()).EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1Request(ezsigntemplatepublicGetEzsigntemplatepublicDetailsV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1`: EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepublicGetEzsigntemplatepublicDetailsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplatepublicGetEzsigntemplatepublicDetailsV1Request** | [**EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1Request**](EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1Request.md) |  | 

### Return type

[**EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1Response**](EzsigntemplatepublicGetEzsigntemplatepublicDetailsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepublicGetFormsDataV1

> EzsigntemplatepublicGetFormsDataV1Response EzsigntemplatepublicGetFormsDataV1(ctx, pkiEzsigntemplatepublicID).Execute()

Retrieve an existing Ezsigntemplatepublic's forms data



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
	pkiEzsigntemplatepublicID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicGetFormsDataV1(context.Background(), pkiEzsigntemplatepublicID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicGetFormsDataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepublicGetFormsDataV1`: EzsigntemplatepublicGetFormsDataV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicGetFormsDataV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepublicID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepublicGetFormsDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatepublicGetFormsDataV1Response**](EzsigntemplatepublicGetFormsDataV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepublicGetListV1

> EzsigntemplatepublicGetListV1Response EzsigntemplatepublicGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Ezsigntemplatepublic list



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
	resp, r, err := apiClient.ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepublicGetListV1`: EzsigntemplatepublicGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepublicGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**EzsigntemplatepublicGetListV1Response**](EzsigntemplatepublicGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepublicGetObjectV2

> EzsigntemplatepublicGetObjectV2Response EzsigntemplatepublicGetObjectV2(ctx, pkiEzsigntemplatepublicID).Execute()

Retrieve an existing Ezsigntemplatepublic



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
	pkiEzsigntemplatepublicID := int32(56) // int32 | The unique ID of the Ezsigntemplatepublic

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicGetObjectV2(context.Background(), pkiEzsigntemplatepublicID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepublicGetObjectV2`: EzsigntemplatepublicGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepublicID** | **int32** | The unique ID of the Ezsigntemplatepublic | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepublicGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatepublicGetObjectV2Response**](EzsigntemplatepublicGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepublicResetLimitExceededCounterV1

> EzsigntemplatepublicResetLimitExceededCounterV1Response EzsigntemplatepublicResetLimitExceededCounterV1(ctx, pkiEzsigntemplatepublicID).Body(body).Execute()

Reset the limit exceeded counter



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
	pkiEzsigntemplatepublicID := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicResetLimitExceededCounterV1(context.Background(), pkiEzsigntemplatepublicID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicResetLimitExceededCounterV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepublicResetLimitExceededCounterV1`: EzsigntemplatepublicResetLimitExceededCounterV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicResetLimitExceededCounterV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepublicID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepublicResetLimitExceededCounterV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**EzsigntemplatepublicResetLimitExceededCounterV1Response**](EzsigntemplatepublicResetLimitExceededCounterV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatepublicResetUrlV1

> EzsigntemplatepublicResetUrlV1Response EzsigntemplatepublicResetUrlV1(ctx, pkiEzsigntemplatepublicID).Body(body).Execute()

Reset the Ezsigntemplatepublic url



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
	pkiEzsigntemplatepublicID := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicResetUrlV1(context.Background(), pkiEzsigntemplatepublicID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicResetUrlV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsigntemplatepublicResetUrlV1`: EzsigntemplatepublicResetUrlV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatepublicAPI.EzsigntemplatepublicResetUrlV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatepublicID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatepublicResetUrlV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**EzsigntemplatepublicResetUrlV1Response**](EzsigntemplatepublicResetUrlV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

