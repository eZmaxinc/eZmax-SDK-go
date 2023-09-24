# PhonestaticResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiPhonestaticID** | **int32** | The unique ID of the Phone. | 
**SPhonestaticE164** | Pointer to **string** | A phone number in E.164 Format | [optional] 
**SPhonestaticExtension** | Pointer to **string** | The extension of the phone number. | [optional] 

## Methods

### NewPhonestaticResponseCompound

`func NewPhonestaticResponseCompound(pkiPhonestaticID int32, ) *PhonestaticResponseCompound`

NewPhonestaticResponseCompound instantiates a new PhonestaticResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhonestaticResponseCompoundWithDefaults

`func NewPhonestaticResponseCompoundWithDefaults() *PhonestaticResponseCompound`

NewPhonestaticResponseCompoundWithDefaults instantiates a new PhonestaticResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPhonestaticID

`func (o *PhonestaticResponseCompound) GetPkiPhonestaticID() int32`

GetPkiPhonestaticID returns the PkiPhonestaticID field if non-nil, zero value otherwise.

### GetPkiPhonestaticIDOk

`func (o *PhonestaticResponseCompound) GetPkiPhonestaticIDOk() (*int32, bool)`

GetPkiPhonestaticIDOk returns a tuple with the PkiPhonestaticID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPhonestaticID

`func (o *PhonestaticResponseCompound) SetPkiPhonestaticID(v int32)`

SetPkiPhonestaticID sets PkiPhonestaticID field to given value.


### GetSPhonestaticE164

`func (o *PhonestaticResponseCompound) GetSPhonestaticE164() string`

GetSPhonestaticE164 returns the SPhonestaticE164 field if non-nil, zero value otherwise.

### GetSPhonestaticE164Ok

`func (o *PhonestaticResponseCompound) GetSPhonestaticE164Ok() (*string, bool)`

GetSPhonestaticE164Ok returns a tuple with the SPhonestaticE164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhonestaticE164

`func (o *PhonestaticResponseCompound) SetSPhonestaticE164(v string)`

SetSPhonestaticE164 sets SPhonestaticE164 field to given value.

### HasSPhonestaticE164

`func (o *PhonestaticResponseCompound) HasSPhonestaticE164() bool`

HasSPhonestaticE164 returns a boolean if a field has been set.

### GetSPhonestaticExtension

`func (o *PhonestaticResponseCompound) GetSPhonestaticExtension() string`

GetSPhonestaticExtension returns the SPhonestaticExtension field if non-nil, zero value otherwise.

### GetSPhonestaticExtensionOk

`func (o *PhonestaticResponseCompound) GetSPhonestaticExtensionOk() (*string, bool)`

GetSPhonestaticExtensionOk returns a tuple with the SPhonestaticExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhonestaticExtension

`func (o *PhonestaticResponseCompound) SetSPhonestaticExtension(v string)`

SetSPhonestaticExtension sets SPhonestaticExtension field to given value.

### HasSPhonestaticExtension

`func (o *PhonestaticResponseCompound) HasSPhonestaticExtension() bool`

HasSPhonestaticExtension returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


