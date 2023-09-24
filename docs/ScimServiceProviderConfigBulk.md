# ScimServiceProviderConfigBulk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supported** | **bool** | A Boolean value specifying whether or not the operation is supported. | 
**MaxOperations** | **int32** | An integer value specifying the maximum number of operations. | 
**MaxPayloadSize** | **int32** | An integer value specifying the maximum payload size in bytes. | 

## Methods

### NewScimServiceProviderConfigBulk

`func NewScimServiceProviderConfigBulk(supported bool, maxOperations int32, maxPayloadSize int32, ) *ScimServiceProviderConfigBulk`

NewScimServiceProviderConfigBulk instantiates a new ScimServiceProviderConfigBulk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimServiceProviderConfigBulkWithDefaults

`func NewScimServiceProviderConfigBulkWithDefaults() *ScimServiceProviderConfigBulk`

NewScimServiceProviderConfigBulkWithDefaults instantiates a new ScimServiceProviderConfigBulk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupported

`func (o *ScimServiceProviderConfigBulk) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *ScimServiceProviderConfigBulk) GetSupportedOk() (*bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupported

`func (o *ScimServiceProviderConfigBulk) SetSupported(v bool)`

SetSupported sets Supported field to given value.


### GetMaxOperations

`func (o *ScimServiceProviderConfigBulk) GetMaxOperations() int32`

GetMaxOperations returns the MaxOperations field if non-nil, zero value otherwise.

### GetMaxOperationsOk

`func (o *ScimServiceProviderConfigBulk) GetMaxOperationsOk() (*int32, bool)`

GetMaxOperationsOk returns a tuple with the MaxOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOperations

`func (o *ScimServiceProviderConfigBulk) SetMaxOperations(v int32)`

SetMaxOperations sets MaxOperations field to given value.


### GetMaxPayloadSize

`func (o *ScimServiceProviderConfigBulk) GetMaxPayloadSize() int32`

GetMaxPayloadSize returns the MaxPayloadSize field if non-nil, zero value otherwise.

### GetMaxPayloadSizeOk

`func (o *ScimServiceProviderConfigBulk) GetMaxPayloadSizeOk() (*int32, bool)`

GetMaxPayloadSizeOk returns a tuple with the MaxPayloadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayloadSize

`func (o *ScimServiceProviderConfigBulk) SetMaxPayloadSize(v int32)`

SetMaxPayloadSize sets MaxPayloadSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


