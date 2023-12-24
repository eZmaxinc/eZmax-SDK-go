# PhoneResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiPhoneID** | **int32** | The unique ID of the Phone. | 
**FkiPhonetypeID** | **int32** | The unique ID of the Phonetype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home| |3|Mobile| |4|Fax| |5|Pager| |6|Toll Free| | 
**EPhoneType** | Pointer to [**FieldEPhoneType**](FieldEPhoneType.md) |  | [optional] 
**SPhoneE164** | Pointer to **string** | A phone number in E.164 Format | [optional] 
**SPhoneExtension** | Pointer to **string** | The extension of the phone number.  The extension is the \&quot;123\&quot; section in this sample phone number: (514) 990-1516 x123.  It can also be used with international phone numbers | [optional] 
**BPhoneInternational** | Pointer to **bool** | Indicate the phone number is an international phone number. | [optional] 

## Methods

### NewPhoneResponseCompound

`func NewPhoneResponseCompound(pkiPhoneID int32, fkiPhonetypeID int32, ) *PhoneResponseCompound`

NewPhoneResponseCompound instantiates a new PhoneResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneResponseCompoundWithDefaults

`func NewPhoneResponseCompoundWithDefaults() *PhoneResponseCompound`

NewPhoneResponseCompoundWithDefaults instantiates a new PhoneResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPhoneID

`func (o *PhoneResponseCompound) GetPkiPhoneID() int32`

GetPkiPhoneID returns the PkiPhoneID field if non-nil, zero value otherwise.

### GetPkiPhoneIDOk

`func (o *PhoneResponseCompound) GetPkiPhoneIDOk() (*int32, bool)`

GetPkiPhoneIDOk returns a tuple with the PkiPhoneID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPhoneID

`func (o *PhoneResponseCompound) SetPkiPhoneID(v int32)`

SetPkiPhoneID sets PkiPhoneID field to given value.


### GetFkiPhonetypeID

`func (o *PhoneResponseCompound) GetFkiPhonetypeID() int32`

GetFkiPhonetypeID returns the FkiPhonetypeID field if non-nil, zero value otherwise.

### GetFkiPhonetypeIDOk

`func (o *PhoneResponseCompound) GetFkiPhonetypeIDOk() (*int32, bool)`

GetFkiPhonetypeIDOk returns a tuple with the FkiPhonetypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPhonetypeID

`func (o *PhoneResponseCompound) SetFkiPhonetypeID(v int32)`

SetFkiPhonetypeID sets FkiPhonetypeID field to given value.


### GetEPhoneType

`func (o *PhoneResponseCompound) GetEPhoneType() FieldEPhoneType`

GetEPhoneType returns the EPhoneType field if non-nil, zero value otherwise.

### GetEPhoneTypeOk

`func (o *PhoneResponseCompound) GetEPhoneTypeOk() (*FieldEPhoneType, bool)`

GetEPhoneTypeOk returns a tuple with the EPhoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPhoneType

`func (o *PhoneResponseCompound) SetEPhoneType(v FieldEPhoneType)`

SetEPhoneType sets EPhoneType field to given value.

### HasEPhoneType

`func (o *PhoneResponseCompound) HasEPhoneType() bool`

HasEPhoneType returns a boolean if a field has been set.

### GetSPhoneE164

`func (o *PhoneResponseCompound) GetSPhoneE164() string`

GetSPhoneE164 returns the SPhoneE164 field if non-nil, zero value otherwise.

### GetSPhoneE164Ok

`func (o *PhoneResponseCompound) GetSPhoneE164Ok() (*string, bool)`

GetSPhoneE164Ok returns a tuple with the SPhoneE164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneE164

`func (o *PhoneResponseCompound) SetSPhoneE164(v string)`

SetSPhoneE164 sets SPhoneE164 field to given value.

### HasSPhoneE164

`func (o *PhoneResponseCompound) HasSPhoneE164() bool`

HasSPhoneE164 returns a boolean if a field has been set.

### GetSPhoneExtension

`func (o *PhoneResponseCompound) GetSPhoneExtension() string`

GetSPhoneExtension returns the SPhoneExtension field if non-nil, zero value otherwise.

### GetSPhoneExtensionOk

`func (o *PhoneResponseCompound) GetSPhoneExtensionOk() (*string, bool)`

GetSPhoneExtensionOk returns a tuple with the SPhoneExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneExtension

`func (o *PhoneResponseCompound) SetSPhoneExtension(v string)`

SetSPhoneExtension sets SPhoneExtension field to given value.

### HasSPhoneExtension

`func (o *PhoneResponseCompound) HasSPhoneExtension() bool`

HasSPhoneExtension returns a boolean if a field has been set.

### GetBPhoneInternational

`func (o *PhoneResponseCompound) GetBPhoneInternational() bool`

GetBPhoneInternational returns the BPhoneInternational field if non-nil, zero value otherwise.

### GetBPhoneInternationalOk

`func (o *PhoneResponseCompound) GetBPhoneInternationalOk() (*bool, bool)`

GetBPhoneInternationalOk returns a tuple with the BPhoneInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBPhoneInternational

`func (o *PhoneResponseCompound) SetBPhoneInternational(v bool)`

SetBPhoneInternational sets BPhoneInternational field to given value.

### HasBPhoneInternational

`func (o *PhoneResponseCompound) HasBPhoneInternational() bool`

HasBPhoneInternational returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


