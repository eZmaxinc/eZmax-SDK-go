# eZmaxAPI\ObjectDiscussionmembershipAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscussionmembershipCreateObjectV1**](ObjectDiscussionmembershipAPI.md#DiscussionmembershipCreateObjectV1) | **Post** /1/object/discussionmembership | Create a new Discussionmembership
[**DiscussionmembershipDeleteObjectV1**](ObjectDiscussionmembershipAPI.md#DiscussionmembershipDeleteObjectV1) | **Delete** /1/object/discussionmembership/{pkiDiscussionmembershipID} | Delete an existing Discussionmembership



## DiscussionmembershipCreateObjectV1

> DiscussionmembershipCreateObjectV1Response DiscussionmembershipCreateObjectV1(ctx).DiscussionmembershipCreateObjectV1Request(discussionmembershipCreateObjectV1Request).Execute()

Create a new Discussionmembership



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
	discussionmembershipCreateObjectV1Request := *openapiclient.NewDiscussionmembershipCreateObjectV1Request([]openapiclient.DiscussionmembershipRequestCompound{*openapiclient.NewDiscussionmembershipRequestCompound(int32(125), "2020-12-31 23:59:59")}) // DiscussionmembershipCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDiscussionmembershipAPI.DiscussionmembershipCreateObjectV1(context.Background()).DiscussionmembershipCreateObjectV1Request(discussionmembershipCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDiscussionmembershipAPI.DiscussionmembershipCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscussionmembershipCreateObjectV1`: DiscussionmembershipCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDiscussionmembershipAPI.DiscussionmembershipCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscussionmembershipCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discussionmembershipCreateObjectV1Request** | [**DiscussionmembershipCreateObjectV1Request**](DiscussionmembershipCreateObjectV1Request.md) |  | 

### Return type

[**DiscussionmembershipCreateObjectV1Response**](DiscussionmembershipCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscussionmembershipDeleteObjectV1

> CommonResponse DiscussionmembershipDeleteObjectV1(ctx, pkiDiscussionmembershipID).Execute()

Delete an existing Discussionmembership



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
	pkiDiscussionmembershipID := int32(56) // int32 | The unique ID of the Discussionmembership

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDiscussionmembershipAPI.DiscussionmembershipDeleteObjectV1(context.Background(), pkiDiscussionmembershipID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDiscussionmembershipAPI.DiscussionmembershipDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscussionmembershipDeleteObjectV1`: CommonResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectDiscussionmembershipAPI.DiscussionmembershipDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDiscussionmembershipID** | **int32** | The unique ID of the Discussionmembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscussionmembershipDeleteObjectV1Request struct via the builder pattern


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

