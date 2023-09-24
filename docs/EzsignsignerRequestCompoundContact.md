# EzsignsignerRequestCompoundContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SContactFirstname** | **string** | The First name of the contact | 
**SContactLastname** | **string** | The Last name of the contact | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**SPhoneE164** | Pointer to **string** | A phone number in E.164 Format | [optional] 
**SPhoneExtension** | Pointer to **string** | The extension of the phone number.  The extension is the \&quot;123\&quot; section in this sample phone number: (514) 990-1516 x123.  It can also be used with international phone numbers | [optional] 
**SPhoneE164Cell** | Pointer to **string** | A phone number in E.164 Format | [optional] 
**SPhoneNumber** | Pointer to **string** |  | [optional] 
**SPhoneNumberCell** | Pointer to **string** |  | [optional] 

## Methods

### NewEzsignsignerRequestCompoundContact

`func NewEzsignsignerRequestCompoundContact(sContactFirstname string, sContactLastname string, fkiLanguageID int32, ) *EzsignsignerRequestCompoundContact`

NewEzsignsignerRequestCompoundContact instantiates a new EzsignsignerRequestCompoundContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignerRequestCompoundContactWithDefaults

`func NewEzsignsignerRequestCompoundContactWithDefaults() *EzsignsignerRequestCompoundContact`

NewEzsignsignerRequestCompoundContactWithDefaults instantiates a new EzsignsignerRequestCompoundContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSContactFirstname

`func (o *EzsignsignerRequestCompoundContact) GetSContactFirstname() string`

GetSContactFirstname returns the SContactFirstname field if non-nil, zero value otherwise.

### GetSContactFirstnameOk

`func (o *EzsignsignerRequestCompoundContact) GetSContactFirstnameOk() (*string, bool)`

GetSContactFirstnameOk returns a tuple with the SContactFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactFirstname

`func (o *EzsignsignerRequestCompoundContact) SetSContactFirstname(v string)`

SetSContactFirstname sets SContactFirstname field to given value.


### GetSContactLastname

`func (o *EzsignsignerRequestCompoundContact) GetSContactLastname() string`

GetSContactLastname returns the SContactLastname field if non-nil, zero value otherwise.

### GetSContactLastnameOk

`func (o *EzsignsignerRequestCompoundContact) GetSContactLastnameOk() (*string, bool)`

GetSContactLastnameOk returns a tuple with the SContactLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactLastname

`func (o *EzsignsignerRequestCompoundContact) SetSContactLastname(v string)`

SetSContactLastname sets SContactLastname field to given value.


### GetFkiLanguageID

`func (o *EzsignsignerRequestCompoundContact) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsignsignerRequestCompoundContact) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsignsignerRequestCompoundContact) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSEmailAddress

`func (o *EzsignsignerRequestCompoundContact) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *EzsignsignerRequestCompoundContact) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *EzsignsignerRequestCompoundContact) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *EzsignsignerRequestCompoundContact) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetSPhoneE164

`func (o *EzsignsignerRequestCompoundContact) GetSPhoneE164() string`

GetSPhoneE164 returns the SPhoneE164 field if non-nil, zero value otherwise.

### GetSPhoneE164Ok

`func (o *EzsignsignerRequestCompoundContact) GetSPhoneE164Ok() (*string, bool)`

GetSPhoneE164Ok returns a tuple with the SPhoneE164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneE164

`func (o *EzsignsignerRequestCompoundContact) SetSPhoneE164(v string)`

SetSPhoneE164 sets SPhoneE164 field to given value.

### HasSPhoneE164

`func (o *EzsignsignerRequestCompoundContact) HasSPhoneE164() bool`

HasSPhoneE164 returns a boolean if a field has been set.

### GetSPhoneExtension

`func (o *EzsignsignerRequestCompoundContact) GetSPhoneExtension() string`

GetSPhoneExtension returns the SPhoneExtension field if non-nil, zero value otherwise.

### GetSPhoneExtensionOk

`func (o *EzsignsignerRequestCompoundContact) GetSPhoneExtensionOk() (*string, bool)`

GetSPhoneExtensionOk returns a tuple with the SPhoneExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneExtension

`func (o *EzsignsignerRequestCompoundContact) SetSPhoneExtension(v string)`

SetSPhoneExtension sets SPhoneExtension field to given value.

### HasSPhoneExtension

`func (o *EzsignsignerRequestCompoundContact) HasSPhoneExtension() bool`

HasSPhoneExtension returns a boolean if a field has been set.

### GetSPhoneE164Cell

`func (o *EzsignsignerRequestCompoundContact) GetSPhoneE164Cell() string`

GetSPhoneE164Cell returns the SPhoneE164Cell field if non-nil, zero value otherwise.

### GetSPhoneE164CellOk

`func (o *EzsignsignerRequestCompoundContact) GetSPhoneE164CellOk() (*string, bool)`

GetSPhoneE164CellOk returns a tuple with the SPhoneE164Cell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneE164Cell

`func (o *EzsignsignerRequestCompoundContact) SetSPhoneE164Cell(v string)`

SetSPhoneE164Cell sets SPhoneE164Cell field to given value.

### HasSPhoneE164Cell

`func (o *EzsignsignerRequestCompoundContact) HasSPhoneE164Cell() bool`

HasSPhoneE164Cell returns a boolean if a field has been set.

### GetSPhoneNumber

`func (o *EzsignsignerRequestCompoundContact) GetSPhoneNumber() string`

GetSPhoneNumber returns the SPhoneNumber field if non-nil, zero value otherwise.

### GetSPhoneNumberOk

`func (o *EzsignsignerRequestCompoundContact) GetSPhoneNumberOk() (*string, bool)`

GetSPhoneNumberOk returns a tuple with the SPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneNumber

`func (o *EzsignsignerRequestCompoundContact) SetSPhoneNumber(v string)`

SetSPhoneNumber sets SPhoneNumber field to given value.

### HasSPhoneNumber

`func (o *EzsignsignerRequestCompoundContact) HasSPhoneNumber() bool`

HasSPhoneNumber returns a boolean if a field has been set.

### GetSPhoneNumberCell

`func (o *EzsignsignerRequestCompoundContact) GetSPhoneNumberCell() string`

GetSPhoneNumberCell returns the SPhoneNumberCell field if non-nil, zero value otherwise.

### GetSPhoneNumberCellOk

`func (o *EzsignsignerRequestCompoundContact) GetSPhoneNumberCellOk() (*string, bool)`

GetSPhoneNumberCellOk returns a tuple with the SPhoneNumberCell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneNumberCell

`func (o *EzsignsignerRequestCompoundContact) SetSPhoneNumberCell(v string)`

SetSPhoneNumberCell sets SPhoneNumberCell field to given value.

### HasSPhoneNumberCell

`func (o *EzsignsignerRequestCompoundContact) HasSPhoneNumberCell() bool`

HasSPhoneNumberCell returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


