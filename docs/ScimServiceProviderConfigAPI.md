# eZmaxAPI\ScimServiceProviderConfigAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceProviderConfigGetObjectScimV2**](ScimServiceProviderConfigAPI.md#ServiceProviderConfigGetObjectScimV2) | **Get** /2/scim/ServiceProviderConfig | Get Service Provider Configuration



## ServiceProviderConfigGetObjectScimV2

> ScimServiceProviderConfig ServiceProviderConfigGetObjectScimV2(ctx).Execute()

Get Service Provider Configuration

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimServiceProviderConfigAPI.ServiceProviderConfigGetObjectScimV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimServiceProviderConfigAPI.ServiceProviderConfigGetObjectScimV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceProviderConfigGetObjectScimV2`: ScimServiceProviderConfig
    fmt.Fprintf(os.Stdout, "Response from `ScimServiceProviderConfigAPI.ServiceProviderConfigGetObjectScimV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServiceProviderConfigGetObjectScimV2Request struct via the builder pattern


### Return type

[**ScimServiceProviderConfig**](ScimServiceProviderConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

