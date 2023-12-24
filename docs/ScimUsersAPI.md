# eZmaxAPI\ScimUsersAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateObjectScimV2**](ScimUsersAPI.md#UsersCreateObjectScimV2) | **Post** /2/scim/Users | Create a new User
[**UsersDeleteObjectScimV2**](ScimUsersAPI.md#UsersDeleteObjectScimV2) | **Delete** /2/scim/Users/{userId} | Delete an existing User
[**UsersEditObjectScimV2**](ScimUsersAPI.md#UsersEditObjectScimV2) | **Put** /2/scim/Users/{userId} | Edit an existing User
[**UsersGetListScimV2**](ScimUsersAPI.md#UsersGetListScimV2) | **Get** /2/scim/Users | Retrieve User list
[**UsersGetObjectScimV2**](ScimUsersAPI.md#UsersGetObjectScimV2) | **Get** /2/scim/Users/{userId} | Retrieve an existing User



## UsersCreateObjectScimV2

> ScimUser UsersCreateObjectScimV2(ctx).ScimUser(scimUser).Execute()

Create a new User

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
	scimUser := *openapiclient.NewScimUser("UserName_example") // ScimUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScimUsersAPI.UsersCreateObjectScimV2(context.Background()).ScimUser(scimUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScimUsersAPI.UsersCreateObjectScimV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersCreateObjectScimV2`: ScimUser
	fmt.Fprintf(os.Stdout, "Response from `ScimUsersAPI.UsersCreateObjectScimV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateObjectScimV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scimUser** | [**ScimUser**](ScimUser.md) |  | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersDeleteObjectScimV2

> UsersDeleteObjectScimV2(ctx, userId).Execute()

Delete an existing User

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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScimUsersAPI.UsersDeleteObjectScimV2(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScimUsersAPI.UsersDeleteObjectScimV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteObjectScimV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersEditObjectScimV2

> ScimUser UsersEditObjectScimV2(ctx, userId).ScimUser(scimUser).Execute()

Edit an existing User

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
	userId := "userId_example" // string | 
	scimUser := *openapiclient.NewScimUser("UserName_example") // ScimUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScimUsersAPI.UsersEditObjectScimV2(context.Background(), userId).ScimUser(scimUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScimUsersAPI.UsersEditObjectScimV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersEditObjectScimV2`: ScimUser
	fmt.Fprintf(os.Stdout, "Response from `ScimUsersAPI.UsersEditObjectScimV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersEditObjectScimV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scimUser** | [**ScimUser**](ScimUser.md) |  | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetListScimV2

> ScimUserList UsersGetListScimV2(ctx).Filter(filter).Execute()

Retrieve User list

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
	filter := "filter_example" // string | Filter expression for searching users (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScimUsersAPI.UsersGetListScimV2(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScimUsersAPI.UsersGetListScimV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGetListScimV2`: ScimUserList
	fmt.Fprintf(os.Stdout, "Response from `ScimUsersAPI.UsersGetListScimV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetListScimV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter expression for searching users | 

### Return type

[**ScimUserList**](ScimUserList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetObjectScimV2

> ScimUser UsersGetObjectScimV2(ctx, userId).Execute()

Retrieve an existing User

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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScimUsersAPI.UsersGetObjectScimV2(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScimUsersAPI.UsersGetObjectScimV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGetObjectScimV2`: ScimUser
	fmt.Fprintf(os.Stdout, "Response from `ScimUsersAPI.UsersGetObjectScimV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetObjectScimV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

