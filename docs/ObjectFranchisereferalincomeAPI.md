# eZmaxAPI\ObjectFranchisereferalincomeAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FranchisereferalincomeCreateObjectV2**](ObjectFranchisereferalincomeAPI.md#FranchisereferalincomeCreateObjectV2) | **Post** /2/object/franchisereferalincome | Create a new Franchisereferalincome



## FranchisereferalincomeCreateObjectV2

> FranchisereferalincomeCreateObjectV2Response FranchisereferalincomeCreateObjectV2(ctx).FranchisereferalincomeCreateObjectV2Request(franchisereferalincomeCreateObjectV2Request).Execute()

Create a new Franchisereferalincome



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
	franchisereferalincomeCreateObjectV2Request := *openapiclient.NewFranchisereferalincomeCreateObjectV2Request([]openapiclient.FranchisereferalincomeRequestCompound{*openapiclient.NewFranchisereferalincomeRequestCompound(int32(61), int32(51), int32(21), "500275.62", "275.00", "385.00", "800.00", "2020-12-31", "This is a comment", int32(50), "SFranchisereferalincomeRemoteid_example", []openapiclient.ContactRequestCompound{*openapiclient.NewContactRequestCompound(int32(2), int32(2), "John", "Doe", "eZmax Solutions Inc.", *openapiclient.NewContactinformationsRequestCompound(int32(123), int32(123), int32(123), int32(123), []openapiclient.AddressRequestCompound{*openapiclient.NewAddressRequest(int32(1), "2540", "Daniel-Johnson Blvd.", "Laval", int32(11), int32(1), "H7T2S3")}, []openapiclient.PhoneRequestCompound{*openapiclient.NewPhoneRequest(int32(1))}, []openapiclient.EmailRequestCompound{*openapiclient.NewEmailRequest(int32(1), "email@example.com")}, []openapiclient.WebsiteRequestCompound{*openapiclient.NewWebsiteRequest(int32(1), "https://www.example.com")}))})}) // FranchisereferalincomeCreateObjectV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectFranchisereferalincomeAPI.FranchisereferalincomeCreateObjectV2(context.Background()).FranchisereferalincomeCreateObjectV2Request(franchisereferalincomeCreateObjectV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectFranchisereferalincomeAPI.FranchisereferalincomeCreateObjectV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FranchisereferalincomeCreateObjectV2`: FranchisereferalincomeCreateObjectV2Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectFranchisereferalincomeAPI.FranchisereferalincomeCreateObjectV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFranchisereferalincomeCreateObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **franchisereferalincomeCreateObjectV2Request** | [**FranchisereferalincomeCreateObjectV2Request**](FranchisereferalincomeCreateObjectV2Request.md) |  | 

### Return type

[**FranchisereferalincomeCreateObjectV2Response**](FranchisereferalincomeCreateObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

