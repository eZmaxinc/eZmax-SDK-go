# eZmaxAPI\ObjectUsergroupexternalAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsergroupexternalCreateObjectV1**](ObjectUsergroupexternalAPI.md#UsergroupexternalCreateObjectV1) | **Post** /1/object/usergroupexternal | Create a new Usergroupexternal
[**UsergroupexternalDeleteObjectV1**](ObjectUsergroupexternalAPI.md#UsergroupexternalDeleteObjectV1) | **Delete** /1/object/usergroupexternal/{pkiUsergroupexternalID} | Delete an existing Usergroupexternal
[**UsergroupexternalEditObjectV1**](ObjectUsergroupexternalAPI.md#UsergroupexternalEditObjectV1) | **Put** /1/object/usergroupexternal/{pkiUsergroupexternalID} | Edit an existing Usergroupexternal
[**UsergroupexternalGetAutocompleteV2**](ObjectUsergroupexternalAPI.md#UsergroupexternalGetAutocompleteV2) | **Get** /2/object/usergroupexternal/getAutocomplete/{sSelector} | Retrieve Usergroupexternals and IDs
[**UsergroupexternalGetListV1**](ObjectUsergroupexternalAPI.md#UsergroupexternalGetListV1) | **Get** /1/object/usergroupexternal/getList | Retrieve Usergroupexternal list
[**UsergroupexternalGetObjectV2**](ObjectUsergroupexternalAPI.md#UsergroupexternalGetObjectV2) | **Get** /2/object/usergroupexternal/{pkiUsergroupexternalID} | Retrieve an existing Usergroupexternal
[**UsergroupexternalGetUsergroupexternalmembershipsV1**](ObjectUsergroupexternalAPI.md#UsergroupexternalGetUsergroupexternalmembershipsV1) | **Get** /1/object/usergroupexternal/{pkiUsergroupexternalID}/getUsergroupexternalmemberships | Retrieve an existing Usergroupexternal&#39;s Usergroupexternalmemberships
[**UsergroupexternalGetUsergroupsV1**](ObjectUsergroupexternalAPI.md#UsergroupexternalGetUsergroupsV1) | **Get** /1/object/usergroupexternal/{pkiUsergroupexternalID}/getUsergroups | Get Usergroupexternal&#39;s Usergroups



## UsergroupexternalCreateObjectV1

> UsergroupexternalCreateObjectV1Response UsergroupexternalCreateObjectV1(ctx).UsergroupexternalCreateObjectV1Request(usergroupexternalCreateObjectV1Request).Execute()

Create a new Usergroupexternal



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
	usergroupexternalCreateObjectV1Request := *openapiclient.NewUsergroupexternalCreateObjectV1Request([]openapiclient.UsergroupexternalRequestCompound{*openapiclient.NewUsergroupexternalRequestCompound("Administrators", "5140-1542")}) // UsergroupexternalCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalCreateObjectV1(context.Background()).UsergroupexternalCreateObjectV1Request(usergroupexternalCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupexternalAPI.UsergroupexternalCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupexternalCreateObjectV1`: UsergroupexternalCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupexternalAPI.UsergroupexternalCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupexternalCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usergroupexternalCreateObjectV1Request** | [**UsergroupexternalCreateObjectV1Request**](UsergroupexternalCreateObjectV1Request.md) |  | 

### Return type

[**UsergroupexternalCreateObjectV1Response**](UsergroupexternalCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupexternalDeleteObjectV1

> UsergroupexternalDeleteObjectV1Response UsergroupexternalDeleteObjectV1(ctx, pkiUsergroupexternalID).Execute()

Delete an existing Usergroupexternal



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
	pkiUsergroupexternalID := int32(56) // int32 | The unique ID of the Usergroupexternal

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalDeleteObjectV1(context.Background(), pkiUsergroupexternalID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupexternalAPI.UsergroupexternalDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupexternalDeleteObjectV1`: UsergroupexternalDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupexternalAPI.UsergroupexternalDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupexternalID** | **int32** | The unique ID of the Usergroupexternal | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupexternalDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsergroupexternalDeleteObjectV1Response**](UsergroupexternalDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupexternalEditObjectV1

> UsergroupexternalEditObjectV1Response UsergroupexternalEditObjectV1(ctx, pkiUsergroupexternalID).UsergroupexternalEditObjectV1Request(usergroupexternalEditObjectV1Request).Execute()

Edit an existing Usergroupexternal



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
	pkiUsergroupexternalID := int32(56) // int32 | The unique ID of the Usergroupexternal
	usergroupexternalEditObjectV1Request := *openapiclient.NewUsergroupexternalEditObjectV1Request(*openapiclient.NewUsergroupexternalRequestCompound("Administrators", "5140-1542")) // UsergroupexternalEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalEditObjectV1(context.Background(), pkiUsergroupexternalID).UsergroupexternalEditObjectV1Request(usergroupexternalEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupexternalAPI.UsergroupexternalEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupexternalEditObjectV1`: UsergroupexternalEditObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupexternalAPI.UsergroupexternalEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupexternalID** | **int32** | The unique ID of the Usergroupexternal | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupexternalEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usergroupexternalEditObjectV1Request** | [**UsergroupexternalEditObjectV1Request**](UsergroupexternalEditObjectV1Request.md) |  | 

### Return type

[**UsergroupexternalEditObjectV1Response**](UsergroupexternalEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupexternalGetAutocompleteV2

> UsergroupexternalGetAutocompleteV2Response UsergroupexternalGetAutocompleteV2(ctx, sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()

Retrieve Usergroupexternals and IDs



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
	sSelector := "sSelector_example" // string | The type of Usergroupexternals to return
	eFilterActive := "eFilterActive_example" // string | Specify which results we want to display. (optional) (default to "Active")
	sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalGetAutocompleteV2(context.Background(), sSelector).EFilterActive(eFilterActive).SQuery(sQuery).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupexternalAPI.UsergroupexternalGetAutocompleteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupexternalGetAutocompleteV2`: UsergroupexternalGetAutocompleteV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupexternalAPI.UsergroupexternalGetAutocompleteV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Usergroupexternals to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupexternalGetAutocompleteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eFilterActive** | **string** | Specify which results we want to display. | [default to &quot;Active&quot;]
 **sQuery** | **string** | Allow to filter the returned results | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 

### Return type

[**UsergroupexternalGetAutocompleteV2Response**](UsergroupexternalGetAutocompleteV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupexternalGetListV1

> UsergroupexternalGetListV1Response UsergroupexternalGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Usergroupexternal list



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
	resp, r, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupexternalAPI.UsergroupexternalGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupexternalGetListV1`: UsergroupexternalGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupexternalAPI.UsergroupexternalGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupexternalGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**UsergroupexternalGetListV1Response**](UsergroupexternalGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupexternalGetObjectV2

> UsergroupexternalGetObjectV2Response UsergroupexternalGetObjectV2(ctx, pkiUsergroupexternalID).Execute()

Retrieve an existing Usergroupexternal



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
	pkiUsergroupexternalID := int32(56) // int32 | The unique ID of the Usergroupexternal

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalGetObjectV2(context.Background(), pkiUsergroupexternalID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupexternalAPI.UsergroupexternalGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupexternalGetObjectV2`: UsergroupexternalGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupexternalAPI.UsergroupexternalGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupexternalID** | **int32** | The unique ID of the Usergroupexternal | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupexternalGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsergroupexternalGetObjectV2Response**](UsergroupexternalGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupexternalGetUsergroupexternalmembershipsV1

> UsergroupexternalGetUsergroupexternalmembershipsV1Response UsergroupexternalGetUsergroupexternalmembershipsV1(ctx, pkiUsergroupexternalID).Execute()

Retrieve an existing Usergroupexternal's Usergroupexternalmemberships

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
	pkiUsergroupexternalID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalGetUsergroupexternalmembershipsV1(context.Background(), pkiUsergroupexternalID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupexternalAPI.UsergroupexternalGetUsergroupexternalmembershipsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupexternalGetUsergroupexternalmembershipsV1`: UsergroupexternalGetUsergroupexternalmembershipsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupexternalAPI.UsergroupexternalGetUsergroupexternalmembershipsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupexternalID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupexternalGetUsergroupexternalmembershipsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsergroupexternalGetUsergroupexternalmembershipsV1Response**](UsergroupexternalGetUsergroupexternalmembershipsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupexternalGetUsergroupsV1

> UsergroupexternalGetUsergroupsV1Response UsergroupexternalGetUsergroupsV1(ctx, pkiUsergroupexternalID).Execute()

Get Usergroupexternal's Usergroups

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
	pkiUsergroupexternalID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUsergroupexternalAPI.UsergroupexternalGetUsergroupsV1(context.Background(), pkiUsergroupexternalID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupexternalAPI.UsergroupexternalGetUsergroupsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupexternalGetUsergroupsV1`: UsergroupexternalGetUsergroupsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupexternalAPI.UsergroupexternalGetUsergroupsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupexternalID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupexternalGetUsergroupsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsergroupexternalGetUsergroupsV1Response**](UsergroupexternalGetUsergroupsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

