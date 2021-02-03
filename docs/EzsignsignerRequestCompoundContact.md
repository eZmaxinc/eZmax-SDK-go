# EzsignsignerRequestCompoundContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SContactFirstname** | **string** | The first name of the Contact | 
**SContactLastname** | **string** | The last name of the Contact | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SEmailAddress** | Pointer to **string** | The email address of the contact. Must be filled if email authentification was requested | [optional] 
**SPhoneNumber** | Pointer to **string** | The Phone number of the contact. Use format \&quot;5149901516\&quot; for North American Numbers (Without \&quot;1\&quot; for long distance code) you would dial like this: 1-514-990-1516. Use format \&quot;498945233886\&quot; for international numbers (Without \&quot;011\&quot;) you would dial like this: +49 89 452 33 88-6. In this example \&quot;49\&quot; is the country code of Germany. | [optional] 
**SPhoneNumberCell** | Pointer to **string** | The Cell Phone number of the contact. Use format \&quot;5149901516\&quot; for North American Numbers (Without \&quot;1\&quot; for long distance code) you would dial like this: 1-514-990-1516. Use format \&quot;498945233886\&quot; for international numbers (Without \&quot;011\&quot;) you would dial like this: +49 89 452 33 88-6. In this example \&quot;49\&quot; is the country code of Germany. | [optional] 

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


