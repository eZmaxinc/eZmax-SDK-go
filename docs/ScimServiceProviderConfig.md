# ScimServiceProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationSchemes** | [**[]ScimAuthenticationScheme**](ScimAuthenticationScheme.md) | A multi-valued complex type that specifies supported authentication scheme properties. | 
**Bulk** | [**ScimServiceProviderConfigBulk**](ScimServiceProviderConfigBulk.md) |  | 
**ChangePassword** | [**ScimServiceProviderConfigChangePassword**](ScimServiceProviderConfigChangePassword.md) |  | 
**DocumentationUri** | **string** | An HTTP-addressable URL pointing to the service provider&#39;s human-consumable help documentation | 
**Etag** | [**ScimServiceProviderConfigEtag**](ScimServiceProviderConfigEtag.md) |  | 
**Filter** | [**ScimServiceProviderConfigFilter**](ScimServiceProviderConfigFilter.md) |  | 
**Patch** | [**ScimServiceProviderConfigPatch**](ScimServiceProviderConfigPatch.md) |  | 
**Sort** | [**ScimServiceProviderConfigSort**](ScimServiceProviderConfigSort.md) |  | 

## Methods

### NewScimServiceProviderConfig

`func NewScimServiceProviderConfig(authenticationSchemes []ScimAuthenticationScheme, bulk ScimServiceProviderConfigBulk, changePassword ScimServiceProviderConfigChangePassword, documentationUri string, etag ScimServiceProviderConfigEtag, filter ScimServiceProviderConfigFilter, patch ScimServiceProviderConfigPatch, sort ScimServiceProviderConfigSort, ) *ScimServiceProviderConfig`

NewScimServiceProviderConfig instantiates a new ScimServiceProviderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimServiceProviderConfigWithDefaults

`func NewScimServiceProviderConfigWithDefaults() *ScimServiceProviderConfig`

NewScimServiceProviderConfigWithDefaults instantiates a new ScimServiceProviderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationSchemes

`func (o *ScimServiceProviderConfig) GetAuthenticationSchemes() []ScimAuthenticationScheme`

GetAuthenticationSchemes returns the AuthenticationSchemes field if non-nil, zero value otherwise.

### GetAuthenticationSchemesOk

`func (o *ScimServiceProviderConfig) GetAuthenticationSchemesOk() (*[]ScimAuthenticationScheme, bool)`

GetAuthenticationSchemesOk returns a tuple with the AuthenticationSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationSchemes

`func (o *ScimServiceProviderConfig) SetAuthenticationSchemes(v []ScimAuthenticationScheme)`

SetAuthenticationSchemes sets AuthenticationSchemes field to given value.


### GetBulk

`func (o *ScimServiceProviderConfig) GetBulk() ScimServiceProviderConfigBulk`

GetBulk returns the Bulk field if non-nil, zero value otherwise.

### GetBulkOk

`func (o *ScimServiceProviderConfig) GetBulkOk() (*ScimServiceProviderConfigBulk, bool)`

GetBulkOk returns a tuple with the Bulk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulk

`func (o *ScimServiceProviderConfig) SetBulk(v ScimServiceProviderConfigBulk)`

SetBulk sets Bulk field to given value.


### GetChangePassword

`func (o *ScimServiceProviderConfig) GetChangePassword() ScimServiceProviderConfigChangePassword`

GetChangePassword returns the ChangePassword field if non-nil, zero value otherwise.

### GetChangePasswordOk

`func (o *ScimServiceProviderConfig) GetChangePasswordOk() (*ScimServiceProviderConfigChangePassword, bool)`

GetChangePasswordOk returns a tuple with the ChangePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassword

`func (o *ScimServiceProviderConfig) SetChangePassword(v ScimServiceProviderConfigChangePassword)`

SetChangePassword sets ChangePassword field to given value.


### GetDocumentationUri

`func (o *ScimServiceProviderConfig) GetDocumentationUri() string`

GetDocumentationUri returns the DocumentationUri field if non-nil, zero value otherwise.

### GetDocumentationUriOk

`func (o *ScimServiceProviderConfig) GetDocumentationUriOk() (*string, bool)`

GetDocumentationUriOk returns a tuple with the DocumentationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUri

`func (o *ScimServiceProviderConfig) SetDocumentationUri(v string)`

SetDocumentationUri sets DocumentationUri field to given value.


### GetEtag

`func (o *ScimServiceProviderConfig) GetEtag() ScimServiceProviderConfigEtag`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *ScimServiceProviderConfig) GetEtagOk() (*ScimServiceProviderConfigEtag, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *ScimServiceProviderConfig) SetEtag(v ScimServiceProviderConfigEtag)`

SetEtag sets Etag field to given value.


### GetFilter

`func (o *ScimServiceProviderConfig) GetFilter() ScimServiceProviderConfigFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ScimServiceProviderConfig) GetFilterOk() (*ScimServiceProviderConfigFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ScimServiceProviderConfig) SetFilter(v ScimServiceProviderConfigFilter)`

SetFilter sets Filter field to given value.


### GetPatch

`func (o *ScimServiceProviderConfig) GetPatch() ScimServiceProviderConfigPatch`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *ScimServiceProviderConfig) GetPatchOk() (*ScimServiceProviderConfigPatch, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *ScimServiceProviderConfig) SetPatch(v ScimServiceProviderConfigPatch)`

SetPatch sets Patch field to given value.


### GetSort

`func (o *ScimServiceProviderConfig) GetSort() ScimServiceProviderConfigSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ScimServiceProviderConfig) GetSortOk() (*ScimServiceProviderConfigSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ScimServiceProviderConfig) SetSort(v ScimServiceProviderConfigSort)`

SetSort sets Sort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


