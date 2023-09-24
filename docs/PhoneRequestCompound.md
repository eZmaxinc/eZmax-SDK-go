# PhoneRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiPhoneID** | Pointer to **int32** | The unique ID of the Phone. | [optional] 
**FkiPhonetypeID** | **int32** | The unique ID of the Phonetype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home| |3|Mobile| |4|Fax| |5|Pager| |6|Toll Free| | 
**EPhoneType** | Pointer to [**FieldEPhoneType**](FieldEPhoneType.md) |  | [optional] 
**SPhoneRegion** | Pointer to **string** | The region of the phone number. (For a North America Number only)  The region is the \&quot;514\&quot; section in this sample phone number: (514) 990-1516 x123 | [optional] 
**SPhoneExchange** | Pointer to **string** | The exchange of the phone number. (For a North America Number only)  The exchange is the \&quot;990\&quot; section in this sample phone number: (514) 990-1516 x123 | [optional] 
**SPhoneNumber** | Pointer to **string** | The number of the phone number. (For a North America Number only)  The number is the \&quot;1516\&quot; section in this sample phone number: (514) 990-1516 x123 | [optional] 
**SPhoneInternational** | Pointer to **string** | The international phone number. | [optional] 
**SPhoneExtension** | Pointer to **string** | The extension of the phone number.  The extension is the \&quot;123\&quot; section in this sample phone number: (514) 990-1516 x123.  It can also be used with international phone numbers | [optional] 
**SPhoneE164** | Pointer to **string** | A phone number in E.164 Format | [optional] 

## Methods

### NewPhoneRequestCompound

`func NewPhoneRequestCompound(fkiPhonetypeID int32, ) *PhoneRequestCompound`

NewPhoneRequestCompound instantiates a new PhoneRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneRequestCompoundWithDefaults

`func NewPhoneRequestCompoundWithDefaults() *PhoneRequestCompound`

NewPhoneRequestCompoundWithDefaults instantiates a new PhoneRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPhoneID

`func (o *PhoneRequestCompound) GetPkiPhoneID() int32`

GetPkiPhoneID returns the PkiPhoneID field if non-nil, zero value otherwise.

### GetPkiPhoneIDOk

`func (o *PhoneRequestCompound) GetPkiPhoneIDOk() (*int32, bool)`

GetPkiPhoneIDOk returns a tuple with the PkiPhoneID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPhoneID

`func (o *PhoneRequestCompound) SetPkiPhoneID(v int32)`

SetPkiPhoneID sets PkiPhoneID field to given value.

### HasPkiPhoneID

`func (o *PhoneRequestCompound) HasPkiPhoneID() bool`

HasPkiPhoneID returns a boolean if a field has been set.

### GetFkiPhonetypeID

`func (o *PhoneRequestCompound) GetFkiPhonetypeID() int32`

GetFkiPhonetypeID returns the FkiPhonetypeID field if non-nil, zero value otherwise.

### GetFkiPhonetypeIDOk

`func (o *PhoneRequestCompound) GetFkiPhonetypeIDOk() (*int32, bool)`

GetFkiPhonetypeIDOk returns a tuple with the FkiPhonetypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPhonetypeID

`func (o *PhoneRequestCompound) SetFkiPhonetypeID(v int32)`

SetFkiPhonetypeID sets FkiPhonetypeID field to given value.


### GetEPhoneType

`func (o *PhoneRequestCompound) GetEPhoneType() FieldEPhoneType`

GetEPhoneType returns the EPhoneType field if non-nil, zero value otherwise.

### GetEPhoneTypeOk

`func (o *PhoneRequestCompound) GetEPhoneTypeOk() (*FieldEPhoneType, bool)`

GetEPhoneTypeOk returns a tuple with the EPhoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPhoneType

`func (o *PhoneRequestCompound) SetEPhoneType(v FieldEPhoneType)`

SetEPhoneType sets EPhoneType field to given value.

### HasEPhoneType

`func (o *PhoneRequestCompound) HasEPhoneType() bool`

HasEPhoneType returns a boolean if a field has been set.

### GetSPhoneRegion

`func (o *PhoneRequestCompound) GetSPhoneRegion() string`

GetSPhoneRegion returns the SPhoneRegion field if non-nil, zero value otherwise.

### GetSPhoneRegionOk

`func (o *PhoneRequestCompound) GetSPhoneRegionOk() (*string, bool)`

GetSPhoneRegionOk returns a tuple with the SPhoneRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneRegion

`func (o *PhoneRequestCompound) SetSPhoneRegion(v string)`

SetSPhoneRegion sets SPhoneRegion field to given value.

### HasSPhoneRegion

`func (o *PhoneRequestCompound) HasSPhoneRegion() bool`

HasSPhoneRegion returns a boolean if a field has been set.

### GetSPhoneExchange

`func (o *PhoneRequestCompound) GetSPhoneExchange() string`

GetSPhoneExchange returns the SPhoneExchange field if non-nil, zero value otherwise.

### GetSPhoneExchangeOk

`func (o *PhoneRequestCompound) GetSPhoneExchangeOk() (*string, bool)`

GetSPhoneExchangeOk returns a tuple with the SPhoneExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneExchange

`func (o *PhoneRequestCompound) SetSPhoneExchange(v string)`

SetSPhoneExchange sets SPhoneExchange field to given value.

### HasSPhoneExchange

`func (o *PhoneRequestCompound) HasSPhoneExchange() bool`

HasSPhoneExchange returns a boolean if a field has been set.

### GetSPhoneNumber

`func (o *PhoneRequestCompound) GetSPhoneNumber() string`

GetSPhoneNumber returns the SPhoneNumber field if non-nil, zero value otherwise.

### GetSPhoneNumberOk

`func (o *PhoneRequestCompound) GetSPhoneNumberOk() (*string, bool)`

GetSPhoneNumberOk returns a tuple with the SPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneNumber

`func (o *PhoneRequestCompound) SetSPhoneNumber(v string)`

SetSPhoneNumber sets SPhoneNumber field to given value.

### HasSPhoneNumber

`func (o *PhoneRequestCompound) HasSPhoneNumber() bool`

HasSPhoneNumber returns a boolean if a field has been set.

### GetSPhoneInternational

`func (o *PhoneRequestCompound) GetSPhoneInternational() string`

GetSPhoneInternational returns the SPhoneInternational field if non-nil, zero value otherwise.

### GetSPhoneInternationalOk

`func (o *PhoneRequestCompound) GetSPhoneInternationalOk() (*string, bool)`

GetSPhoneInternationalOk returns a tuple with the SPhoneInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneInternational

`func (o *PhoneRequestCompound) SetSPhoneInternational(v string)`

SetSPhoneInternational sets SPhoneInternational field to given value.

### HasSPhoneInternational

`func (o *PhoneRequestCompound) HasSPhoneInternational() bool`

HasSPhoneInternational returns a boolean if a field has been set.

### GetSPhoneExtension

`func (o *PhoneRequestCompound) GetSPhoneExtension() string`

GetSPhoneExtension returns the SPhoneExtension field if non-nil, zero value otherwise.

### GetSPhoneExtensionOk

`func (o *PhoneRequestCompound) GetSPhoneExtensionOk() (*string, bool)`

GetSPhoneExtensionOk returns a tuple with the SPhoneExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneExtension

`func (o *PhoneRequestCompound) SetSPhoneExtension(v string)`

SetSPhoneExtension sets SPhoneExtension field to given value.

### HasSPhoneExtension

`func (o *PhoneRequestCompound) HasSPhoneExtension() bool`

HasSPhoneExtension returns a boolean if a field has been set.

### GetSPhoneE164

`func (o *PhoneRequestCompound) GetSPhoneE164() string`

GetSPhoneE164 returns the SPhoneE164 field if non-nil, zero value otherwise.

### GetSPhoneE164Ok

`func (o *PhoneRequestCompound) GetSPhoneE164Ok() (*string, bool)`

GetSPhoneE164Ok returns a tuple with the SPhoneE164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneE164

`func (o *PhoneRequestCompound) SetSPhoneE164(v string)`

SetSPhoneE164 sets SPhoneE164 field to given value.

### HasSPhoneE164

`func (o *PhoneRequestCompound) HasSPhoneE164() bool`

HasSPhoneE164 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


