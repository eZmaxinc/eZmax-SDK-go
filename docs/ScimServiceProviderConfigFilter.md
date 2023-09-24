# ScimServiceProviderConfigFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supported** | **bool** | A Boolean value specifying whether or not the operation is supported. | 
**MaxResults** | **int32** | An integer value specifying the maximum number of resources returned in a response. | 

## Methods

### NewScimServiceProviderConfigFilter

`func NewScimServiceProviderConfigFilter(supported bool, maxResults int32, ) *ScimServiceProviderConfigFilter`

NewScimServiceProviderConfigFilter instantiates a new ScimServiceProviderConfigFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimServiceProviderConfigFilterWithDefaults

`func NewScimServiceProviderConfigFilterWithDefaults() *ScimServiceProviderConfigFilter`

NewScimServiceProviderConfigFilterWithDefaults instantiates a new ScimServiceProviderConfigFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupported

`func (o *ScimServiceProviderConfigFilter) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *ScimServiceProviderConfigFilter) GetSupportedOk() (*bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupported

`func (o *ScimServiceProviderConfigFilter) SetSupported(v bool)`

SetSupported sets Supported field to given value.


### GetMaxResults

`func (o *ScimServiceProviderConfigFilter) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *ScimServiceProviderConfigFilter) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *ScimServiceProviderConfigFilter) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


