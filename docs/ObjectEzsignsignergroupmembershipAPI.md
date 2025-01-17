# eZmaxAPI\ObjectEzsignsignergroupmembershipAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignsignergroupmembershipCreateObjectV1**](ObjectEzsignsignergroupmembershipAPI.md#EzsignsignergroupmembershipCreateObjectV1) | **Post** /1/object/ezsignsignergroupmembership | Create a new Ezsignsignergroupmembership
[**EzsignsignergroupmembershipDeleteObjectV1**](ObjectEzsignsignergroupmembershipAPI.md#EzsignsignergroupmembershipDeleteObjectV1) | **Delete** /1/object/ezsignsignergroupmembership/{pkiEzsignsignergroupmembershipID} | Delete an existing Ezsignsignergroupmembership
[**EzsignsignergroupmembershipGetObjectV2**](ObjectEzsignsignergroupmembershipAPI.md#EzsignsignergroupmembershipGetObjectV2) | **Get** /2/object/ezsignsignergroupmembership/{pkiEzsignsignergroupmembershipID} | Retrieve an existing Ezsignsignergroupmembership



## EzsignsignergroupmembershipCreateObjectV1

> EzsignsignergroupmembershipCreateObjectV1Response EzsignsignergroupmembershipCreateObjectV1(ctx).EzsignsignergroupmembershipCreateObjectV1Request(ezsignsignergroupmembershipCreateObjectV1Request).Execute()

Create a new Ezsignsignergroupmembership



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
	ezsignsignergroupmembershipCreateObjectV1Request := *openapiclient.NewEzsignsignergroupmembershipCreateObjectV1Request([]openapiclient.EzsignsignergroupmembershipRequestCompound{*openapiclient.NewEzsignsignergroupmembershipRequestCompound(int32(27))}) // EzsignsignergroupmembershipCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsignergroupmembershipAPI.EzsignsignergroupmembershipCreateObjectV1(context.Background()).EzsignsignergroupmembershipCreateObjectV1Request(ezsignsignergroupmembershipCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignergroupmembershipAPI.EzsignsignergroupmembershipCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsignergroupmembershipCreateObjectV1`: EzsignsignergroupmembershipCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignergroupmembershipAPI.EzsignsignergroupmembershipCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignergroupmembershipCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignsignergroupmembershipCreateObjectV1Request** | [**EzsignsignergroupmembershipCreateObjectV1Request**](EzsignsignergroupmembershipCreateObjectV1Request.md) |  | 

### Return type

[**EzsignsignergroupmembershipCreateObjectV1Response**](EzsignsignergroupmembershipCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignergroupmembershipDeleteObjectV1

> EzsignsignergroupmembershipDeleteObjectV1Response EzsignsignergroupmembershipDeleteObjectV1(ctx, pkiEzsignsignergroupmembershipID).Execute()

Delete an existing Ezsignsignergroupmembership



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
	pkiEzsignsignergroupmembershipID := int32(56) // int32 | The unique ID of the Ezsignsignergroupmembership

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsignergroupmembershipAPI.EzsignsignergroupmembershipDeleteObjectV1(context.Background(), pkiEzsignsignergroupmembershipID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignergroupmembershipAPI.EzsignsignergroupmembershipDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsignergroupmembershipDeleteObjectV1`: EzsignsignergroupmembershipDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignergroupmembershipAPI.EzsignsignergroupmembershipDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignergroupmembershipID** | **int32** | The unique ID of the Ezsignsignergroupmembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignergroupmembershipDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignsignergroupmembershipDeleteObjectV1Response**](EzsignsignergroupmembershipDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignergroupmembershipGetObjectV2

> EzsignsignergroupmembershipGetObjectV2Response EzsignsignergroupmembershipGetObjectV2(ctx, pkiEzsignsignergroupmembershipID).Execute()

Retrieve an existing Ezsignsignergroupmembership



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
	pkiEzsignsignergroupmembershipID := int32(56) // int32 | The unique ID of the Ezsignsignergroupmembership

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignsignergroupmembershipAPI.EzsignsignergroupmembershipGetObjectV2(context.Background(), pkiEzsignsignergroupmembershipID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignergroupmembershipAPI.EzsignsignergroupmembershipGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignsignergroupmembershipGetObjectV2`: EzsignsignergroupmembershipGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignergroupmembershipAPI.EzsignsignergroupmembershipGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignergroupmembershipID** | **int32** | The unique ID of the Ezsignsignergroupmembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignergroupmembershipGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignsignergroupmembershipGetObjectV2Response**](EzsignsignergroupmembershipGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

