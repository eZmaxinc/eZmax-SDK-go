# eZmaxAPI\ObjectEzsignuserAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignuserEditObjectV1**](ObjectEzsignuserAPI.md#EzsignuserEditObjectV1) | **Put** /1/object/ezsignuser/{pkiEzsignuserID} | Edit an existing Ezsignuser
[**EzsignuserGetObjectV2**](ObjectEzsignuserAPI.md#EzsignuserGetObjectV2) | **Get** /2/object/ezsignuser/{pkiEzsignuserID} | Retrieve an existing Ezsignuser



## EzsignuserEditObjectV1

> CommonResponse EzsignuserEditObjectV1(ctx, pkiEzsignuserID).EzsignuserEditObjectV1Request(ezsignuserEditObjectV1Request).Execute()

Edit an existing Ezsignuser



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
	pkiEzsignuserID := int32(56) // int32 | The unique ID of the Ezsignuser
	ezsignuserEditObjectV1Request := *openapiclient.NewEzsignuserEditObjectV1Request(*openapiclient.NewEzsignuserRequestCompound(int32(21), *openapiclient.NewContactRequestCompoundV2(int32(2), int32(2), openapiclient.Field-eContactType("Agent"), "John", "Doe", *openapiclient.NewContactinformationsRequestCompoundV2(openapiclient.Field-eContactinformationsType("BankAccount"), int32(123), int32(123), int32(123), int32(123), []openapiclient.AddressRequestCompound{*openapiclient.NewAddressRequest(int32(1), "2540", "Daniel-Johnson Blvd.", "Laval", int32(11), int32(1), "H7T2S3")}, []openapiclient.PhoneRequestCompound{*openapiclient.NewPhoneRequest(int32(1))}, []openapiclient.EmailRequestCompound{*openapiclient.NewEmailRequest(int32(1), "email@example.com")}, []openapiclient.WebsiteRequestCompound{*openapiclient.NewWebsiteRequest(int32(1), "https://www.example.com")})))) // EzsignuserEditObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignuserAPI.EzsignuserEditObjectV1(context.Background(), pkiEzsignuserID).EzsignuserEditObjectV1Request(ezsignuserEditObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignuserAPI.EzsignuserEditObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignuserEditObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignuserAPI.EzsignuserEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignuserID** | **int32** | The unique ID of the Ezsignuser | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignuserEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignuserEditObjectV1Request** | [**EzsignuserEditObjectV1Request**](EzsignuserEditObjectV1Request.md) |  | 

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


## EzsignuserGetObjectV2

> EzsignuserGetObjectV2Response EzsignuserGetObjectV2(ctx, pkiEzsignuserID).Execute()

Retrieve an existing Ezsignuser



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
	pkiEzsignuserID := int32(56) // int32 | The unique ID of the Ezsignuser

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectEzsignuserAPI.EzsignuserGetObjectV2(context.Background(), pkiEzsignuserID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignuserAPI.EzsignuserGetObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EzsignuserGetObjectV2`: EzsignuserGetObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignuserAPI.EzsignuserGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignuserID** | **int32** | The unique ID of the Ezsignuser | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignuserGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignuserGetObjectV2Response**](EzsignuserGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

