# EzsignsignerResponseCompoundContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiContactID** | **int32** | The unique ID of the Contact | 
**SContactFirstname** | **string** | The First name of the contact | 
**SContactLastname** | **string** | The Last name of the contact | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**SPhoneE164** | Pointer to **string** | A phone number in E.164 Format | [optional] 
**SPhoneExtension** | Pointer to **string** | The extension of the phone number.  The extension is the \&quot;123\&quot; section in this sample phone number: (514) 990-1516 x123.  It can also be used with international phone numbers | [optional] 
**SPhoneE164Cell** | Pointer to **string** | A phone number in E.164 Format | [optional] 

## Methods

### NewEzsignsignerResponseCompoundContact

`func NewEzsignsignerResponseCompoundContact(pkiContactID int32, sContactFirstname string, sContactLastname string, fkiLanguageID int32, ) *EzsignsignerResponseCompoundContact`

NewEzsignsignerResponseCompoundContact instantiates a new EzsignsignerResponseCompoundContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignerResponseCompoundContactWithDefaults

`func NewEzsignsignerResponseCompoundContactWithDefaults() *EzsignsignerResponseCompoundContact`

NewEzsignsignerResponseCompoundContactWithDefaults instantiates a new EzsignsignerResponseCompoundContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiContactID

`func (o *EzsignsignerResponseCompoundContact) GetPkiContactID() int32`

GetPkiContactID returns the PkiContactID field if non-nil, zero value otherwise.

### GetPkiContactIDOk

`func (o *EzsignsignerResponseCompoundContact) GetPkiContactIDOk() (*int32, bool)`

GetPkiContactIDOk returns a tuple with the PkiContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiContactID

`func (o *EzsignsignerResponseCompoundContact) SetPkiContactID(v int32)`

SetPkiContactID sets PkiContactID field to given value.


### GetSContactFirstname

`func (o *EzsignsignerResponseCompoundContact) GetSContactFirstname() string`

GetSContactFirstname returns the SContactFirstname field if non-nil, zero value otherwise.

### GetSContactFirstnameOk

`func (o *EzsignsignerResponseCompoundContact) GetSContactFirstnameOk() (*string, bool)`

GetSContactFirstnameOk returns a tuple with the SContactFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactFirstname

`func (o *EzsignsignerResponseCompoundContact) SetSContactFirstname(v string)`

SetSContactFirstname sets SContactFirstname field to given value.


### GetSContactLastname

`func (o *EzsignsignerResponseCompoundContact) GetSContactLastname() string`

GetSContactLastname returns the SContactLastname field if non-nil, zero value otherwise.

### GetSContactLastnameOk

`func (o *EzsignsignerResponseCompoundContact) GetSContactLastnameOk() (*string, bool)`

GetSContactLastnameOk returns a tuple with the SContactLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactLastname

`func (o *EzsignsignerResponseCompoundContact) SetSContactLastname(v string)`

SetSContactLastname sets SContactLastname field to given value.


### GetFkiLanguageID

`func (o *EzsignsignerResponseCompoundContact) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsignsignerResponseCompoundContact) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsignsignerResponseCompoundContact) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSEmailAddress

`func (o *EzsignsignerResponseCompoundContact) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *EzsignsignerResponseCompoundContact) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *EzsignsignerResponseCompoundContact) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *EzsignsignerResponseCompoundContact) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetSPhoneE164

`func (o *EzsignsignerResponseCompoundContact) GetSPhoneE164() string`

GetSPhoneE164 returns the SPhoneE164 field if non-nil, zero value otherwise.

### GetSPhoneE164Ok

`func (o *EzsignsignerResponseCompoundContact) GetSPhoneE164Ok() (*string, bool)`

GetSPhoneE164Ok returns a tuple with the SPhoneE164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneE164

`func (o *EzsignsignerResponseCompoundContact) SetSPhoneE164(v string)`

SetSPhoneE164 sets SPhoneE164 field to given value.

### HasSPhoneE164

`func (o *EzsignsignerResponseCompoundContact) HasSPhoneE164() bool`

HasSPhoneE164 returns a boolean if a field has been set.

### GetSPhoneExtension

`func (o *EzsignsignerResponseCompoundContact) GetSPhoneExtension() string`

GetSPhoneExtension returns the SPhoneExtension field if non-nil, zero value otherwise.

### GetSPhoneExtensionOk

`func (o *EzsignsignerResponseCompoundContact) GetSPhoneExtensionOk() (*string, bool)`

GetSPhoneExtensionOk returns a tuple with the SPhoneExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneExtension

`func (o *EzsignsignerResponseCompoundContact) SetSPhoneExtension(v string)`

SetSPhoneExtension sets SPhoneExtension field to given value.

### HasSPhoneExtension

`func (o *EzsignsignerResponseCompoundContact) HasSPhoneExtension() bool`

HasSPhoneExtension returns a boolean if a field has been set.

### GetSPhoneE164Cell

`func (o *EzsignsignerResponseCompoundContact) GetSPhoneE164Cell() string`

GetSPhoneE164Cell returns the SPhoneE164Cell field if non-nil, zero value otherwise.

### GetSPhoneE164CellOk

`func (o *EzsignsignerResponseCompoundContact) GetSPhoneE164CellOk() (*string, bool)`

GetSPhoneE164CellOk returns a tuple with the SPhoneE164Cell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneE164Cell

`func (o *EzsignsignerResponseCompoundContact) SetSPhoneE164Cell(v string)`

SetSPhoneE164Cell sets SPhoneE164Cell field to given value.

### HasSPhoneE164Cell

`func (o *EzsignsignerResponseCompoundContact) HasSPhoneE164Cell() bool`

HasSPhoneE164Cell returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


