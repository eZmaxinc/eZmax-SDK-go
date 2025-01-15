# eZmaxAPI\ObjectUsergroupmembershipAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsergroupmembershipCreateObjectV1**](ObjectUsergroupmembershipAPI.md#UsergroupmembershipCreateObjectV1) | **Post** /1/object/usergroupmembership | Create a new Usergroupmembership
[**UsergroupmembershipDeleteObjectV1**](ObjectUsergroupmembershipAPI.md#UsergroupmembershipDeleteObjectV1) | **Delete** /1/object/usergroupmembership/{pkiUsergroupmembershipID} | Delete an existing Usergroupmembership
[**UsergroupmembershipEditObjectV1**](ObjectUsergroupmembershipAPI.md#UsergroupmembershipEditObjectV1) | **Put** /1/object/usergroupmembership/{pkiUsergroupmembershipID} | Edit an existing Usergroupmembership
[**UsergroupmembershipGetObjectV2**](ObjectUsergroupmembershipAPI.md#UsergroupmembershipGetObjectV2) | **Get** /2/object/usergroupmembership/{pkiUsergroupmembershipID} | Retrieve an existing Usergroupmembership



## UsergroupmembershipCreateObjectV1

> UsergroupmembershipCreateObjectV1Response UsergroupmembershipCreateObjectV1(ctx).UsergroupmembershipCreateObjectV1Request(usergroupmembershipCreateObjectV1Request).Execute()

Create a new Usergroupmembership



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
	usergroupmembershipCreateObjectV1Request := *openapiclient.NewUsergroupmembershipCreateObjectV1Request([]openapiclient.UsergroupmembershipRequestCompound{*openapiclient.NewUsergroupmembershipRequestCompound(int32(2))}) // UsergroupmembershipCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUsergroupmembershipAPI.UsergroupmembershipCreateObjectV1(context.Background()).UsergroupmembershipCreateObjectV1Request(usergroupmembershipCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupmembershipAPI.UsergroupmembershipCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupmembershipCreateObjectV1`: UsergroupmembershipCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupmembershipAPI.UsergroupmembershipCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupmembershipCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usergroupmembershipCreateObjectV1Request** | [**UsergroupmembershipCreateObjectV1Request**](UsergroupmembershipCreateObjectV1Request.md) |  | 

### Return type

[**UsergroupmembershipCreateObjectV1Response**](UsergroupmembershipCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupmembershipDeleteObjectV1

> CommonResponse UsergroupmembershipDeleteObjectV1(ctx, pkiUsergroupmembershipID).Execute()

Delete an existing Usergroupmembership



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
	pkiUsergroupmembershipID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUsergroupmembershipAPI.UsergroupmembershipDeleteObjectV1(context.Background(), pkiUsergroupmembershipID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupmembershipAPI.UsergroupmembershipDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupmembershipDeleteObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupmembershipAPI.UsergroupmembershipDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupmembershipID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupmembershipDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsergroupmembershipEditObjectV1

> CommonResponse UsergroupmembershipEditObjectV1(ctx, pkiUsergroupmembershipID).UsergroupmembershipEditObjectV1Request(usergroupmembershipEditObjectV1Request).Execute()

Edit an existing Usergroupmembership



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
	pkiUsergroupmembershipID := int32(56) // int32 | 
	usergroupmembershipEditObjectV1Request := *openapiclient.NewUsergroupmembershipEditObjectV1Request(*openapiclient.NewUsergroupmembershipRequestCompound(int32(2))) // UsergroupmembershipEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUsergroupmembershipAPI.UsergroupmembershipEditObjectV1(context.Background(), pkiUsergroupmembershipID).UsergroupmembershipEditObjectV1Request(usergroupmembershipEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupmembershipAPI.UsergroupmembershipEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupmembershipEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupmembershipAPI.UsergroupmembershipEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupmembershipID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupmembershipEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usergroupmembershipEditObjectV1Request** | [**UsergroupmembershipEditObjectV1Request**](UsergroupmembershipEditObjectV1Request.md) |  | 

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


## UsergroupmembershipGetObjectV2

> UsergroupmembershipGetObjectV2Response UsergroupmembershipGetObjectV2(ctx, pkiUsergroupmembershipID).Execute()

Retrieve an existing Usergroupmembership



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
	pkiUsergroupmembershipID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectUsergroupmembershipAPI.UsergroupmembershipGetObjectV2(context.Background(), pkiUsergroupmembershipID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectUsergroupmembershipAPI.UsergroupmembershipGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsergroupmembershipGetObjectV2`: UsergroupmembershipGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectUsergroupmembershipAPI.UsergroupmembershipGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiUsergroupmembershipID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsergroupmembershipGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsergroupmembershipGetObjectV2Response**](UsergroupmembershipGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

