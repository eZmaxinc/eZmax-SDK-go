# ScimEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The email address. | [optional] 
**Primary** | Pointer to **bool** |  | [optional] 

## Methods

### NewScimEmail

`func NewScimEmail() *ScimEmail`

NewScimEmail instantiates a new ScimEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimEmailWithDefaults

`func NewScimEmailWithDefaults() *ScimEmail`

NewScimEmailWithDefaults instantiates a new ScimEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ScimEmail) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ScimEmail) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ScimEmail) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ScimEmail) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPrimary

`func (o *ScimEmail) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *ScimEmail) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *ScimEmail) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *ScimEmail) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


