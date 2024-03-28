# PhoneRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiPhoneID** | Pointer to **int32** | The unique ID of the Phone. | [optional] 
**FkiPhonetypeID** | **int32** | The unique ID of the Phonetype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home| |3|Mobile| |4|Fax| |5|Pager| |6|Toll Free| | 
**SPhoneExtension** | Pointer to **string** | The extension of the phone number.  The extension is the \&quot;123\&quot; section in this sample phone number: (514) 990-1516 x123.  It can also be used with international phone numbers | [optional] 
**SPhoneE164** | Pointer to **string** | A phone number in E.164 Format | [optional] 

## Methods

### NewPhoneRequestV2

`func NewPhoneRequestV2(fkiPhonetypeID int32, ) *PhoneRequestV2`

NewPhoneRequestV2 instantiates a new PhoneRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneRequestV2WithDefaults

`func NewPhoneRequestV2WithDefaults() *PhoneRequestV2`

NewPhoneRequestV2WithDefaults instantiates a new PhoneRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPhoneID

`func (o *PhoneRequestV2) GetPkiPhoneID() int32`

GetPkiPhoneID returns the PkiPhoneID field if non-nil, zero value otherwise.

### GetPkiPhoneIDOk

`func (o *PhoneRequestV2) GetPkiPhoneIDOk() (*int32, bool)`

GetPkiPhoneIDOk returns a tuple with the PkiPhoneID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPhoneID

`func (o *PhoneRequestV2) SetPkiPhoneID(v int32)`

SetPkiPhoneID sets PkiPhoneID field to given value.

### HasPkiPhoneID

`func (o *PhoneRequestV2) HasPkiPhoneID() bool`

HasPkiPhoneID returns a boolean if a field has been set.

### GetFkiPhonetypeID

`func (o *PhoneRequestV2) GetFkiPhonetypeID() int32`

GetFkiPhonetypeID returns the FkiPhonetypeID field if non-nil, zero value otherwise.

### GetFkiPhonetypeIDOk

`func (o *PhoneRequestV2) GetFkiPhonetypeIDOk() (*int32, bool)`

GetFkiPhonetypeIDOk returns a tuple with the FkiPhonetypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPhonetypeID

`func (o *PhoneRequestV2) SetFkiPhonetypeID(v int32)`

SetFkiPhonetypeID sets FkiPhonetypeID field to given value.


### GetSPhoneExtension

`func (o *PhoneRequestV2) GetSPhoneExtension() string`

GetSPhoneExtension returns the SPhoneExtension field if non-nil, zero value otherwise.

### GetSPhoneExtensionOk

`func (o *PhoneRequestV2) GetSPhoneExtensionOk() (*string, bool)`

GetSPhoneExtensionOk returns a tuple with the SPhoneExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneExtension

`func (o *PhoneRequestV2) SetSPhoneExtension(v string)`

SetSPhoneExtension sets SPhoneExtension field to given value.

### HasSPhoneExtension

`func (o *PhoneRequestV2) HasSPhoneExtension() bool`

HasSPhoneExtension returns a boolean if a field has been set.

### GetSPhoneE164

`func (o *PhoneRequestV2) GetSPhoneE164() string`

GetSPhoneE164 returns the SPhoneE164 field if non-nil, zero value otherwise.

### GetSPhoneE164Ok

`func (o *PhoneRequestV2) GetSPhoneE164Ok() (*string, bool)`

GetSPhoneE164Ok returns a tuple with the SPhoneE164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneE164

`func (o *PhoneRequestV2) SetSPhoneE164(v string)`

SetSPhoneE164 sets SPhoneE164 field to given value.

### HasSPhoneE164

`func (o *PhoneRequestV2) HasSPhoneE164() bool`

HasSPhoneE164 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


