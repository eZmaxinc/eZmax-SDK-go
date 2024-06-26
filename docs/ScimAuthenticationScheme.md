# ScimAuthenticationScheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A description of the authentication scheme. | 
**Name** | **string** | The common authentication scheme name | 
**Type** | **string** | The authentication scheme. | 

## Methods

### NewScimAuthenticationScheme

`func NewScimAuthenticationScheme(description string, name string, type_ string, ) *ScimAuthenticationScheme`

NewScimAuthenticationScheme instantiates a new ScimAuthenticationScheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimAuthenticationSchemeWithDefaults

`func NewScimAuthenticationSchemeWithDefaults() *ScimAuthenticationScheme`

NewScimAuthenticationSchemeWithDefaults instantiates a new ScimAuthenticationScheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ScimAuthenticationScheme) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScimAuthenticationScheme) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScimAuthenticationScheme) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *ScimAuthenticationScheme) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimAuthenticationScheme) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScimAuthenticationScheme) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ScimAuthenticationScheme) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScimAuthenticationScheme) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScimAuthenticationScheme) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


