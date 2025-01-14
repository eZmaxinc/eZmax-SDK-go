# ContactResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiContactID** | **int32** | The unique ID of the Contact | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**FkiContacttitleID** | **int32** | The unique ID of the Contacttitle.  Valid values:  |Value|Description| |-|-| |1|Ms.| |2|Mr.| |4|(Blank)| |5|Me (For Notaries)| | 
**FkiContactinformationsID** | **int32** | The unique ID of the Contactinformations | 
**DtContactBirthdate** | Pointer to **string** | The Birth Date of the contact | [optional] 
**EContactType** | [**FieldEContactType**](FieldEContactType.md) |  | 
**SContactFirstname** | **string** | The First name of the contact | 
**SContactLastname** | **string** | The Last name of the contact | 
**SContactCompany** | Pointer to **string** | The Company name of the contact | [optional] 
**SContactOccupation** | Pointer to **string** | The occupation of the Contact | [optional] 
**TContactNote** | Pointer to **string** | The note of the Contact | [optional] 
**BContactIsactive** | **bool** | Whether the contact is active or not | 
**ObjContactinformations** | [**ContactinformationsResponseCompound**](ContactinformationsResponseCompound.md) |  | 

## Methods

### NewContactResponseCompound

`func NewContactResponseCompound(pkiContactID int32, fkiLanguageID int32, fkiContacttitleID int32, fkiContactinformationsID int32, eContactType FieldEContactType, sContactFirstname string, sContactLastname string, bContactIsactive bool, objContactinformations ContactinformationsResponseCompound, ) *ContactResponseCompound`

NewContactResponseCompound instantiates a new ContactResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactResponseCompoundWithDefaults

`func NewContactResponseCompoundWithDefaults() *ContactResponseCompound`

NewContactResponseCompoundWithDefaults instantiates a new ContactResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiContactID

`func (o *ContactResponseCompound) GetPkiContactID() int32`

GetPkiContactID returns the PkiContactID field if non-nil, zero value otherwise.

### GetPkiContactIDOk

`func (o *ContactResponseCompound) GetPkiContactIDOk() (*int32, bool)`

GetPkiContactIDOk returns a tuple with the PkiContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiContactID

`func (o *ContactResponseCompound) SetPkiContactID(v int32)`

SetPkiContactID sets PkiContactID field to given value.


### GetFkiLanguageID

`func (o *ContactResponseCompound) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *ContactResponseCompound) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *ContactResponseCompound) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetFkiContacttitleID

`func (o *ContactResponseCompound) GetFkiContacttitleID() int32`

GetFkiContacttitleID returns the FkiContacttitleID field if non-nil, zero value otherwise.

### GetFkiContacttitleIDOk

`func (o *ContactResponseCompound) GetFkiContacttitleIDOk() (*int32, bool)`

GetFkiContacttitleIDOk returns a tuple with the FkiContacttitleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiContacttitleID

`func (o *ContactResponseCompound) SetFkiContacttitleID(v int32)`

SetFkiContacttitleID sets FkiContacttitleID field to given value.


### GetFkiContactinformationsID

`func (o *ContactResponseCompound) GetFkiContactinformationsID() int32`

GetFkiContactinformationsID returns the FkiContactinformationsID field if non-nil, zero value otherwise.

### GetFkiContactinformationsIDOk

`func (o *ContactResponseCompound) GetFkiContactinformationsIDOk() (*int32, bool)`

GetFkiContactinformationsIDOk returns a tuple with the FkiContactinformationsID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiContactinformationsID

`func (o *ContactResponseCompound) SetFkiContactinformationsID(v int32)`

SetFkiContactinformationsID sets FkiContactinformationsID field to given value.


### GetDtContactBirthdate

`func (o *ContactResponseCompound) GetDtContactBirthdate() string`

GetDtContactBirthdate returns the DtContactBirthdate field if non-nil, zero value otherwise.

### GetDtContactBirthdateOk

`func (o *ContactResponseCompound) GetDtContactBirthdateOk() (*string, bool)`

GetDtContactBirthdateOk returns a tuple with the DtContactBirthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtContactBirthdate

`func (o *ContactResponseCompound) SetDtContactBirthdate(v string)`

SetDtContactBirthdate sets DtContactBirthdate field to given value.

### HasDtContactBirthdate

`func (o *ContactResponseCompound) HasDtContactBirthdate() bool`

HasDtContactBirthdate returns a boolean if a field has been set.

### GetEContactType

`func (o *ContactResponseCompound) GetEContactType() FieldEContactType`

GetEContactType returns the EContactType field if non-nil, zero value otherwise.

### GetEContactTypeOk

`func (o *ContactResponseCompound) GetEContactTypeOk() (*FieldEContactType, bool)`

GetEContactTypeOk returns a tuple with the EContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEContactType

`func (o *ContactResponseCompound) SetEContactType(v FieldEContactType)`

SetEContactType sets EContactType field to given value.


### GetSContactFirstname

`func (o *ContactResponseCompound) GetSContactFirstname() string`

GetSContactFirstname returns the SContactFirstname field if non-nil, zero value otherwise.

### GetSContactFirstnameOk

`func (o *ContactResponseCompound) GetSContactFirstnameOk() (*string, bool)`

GetSContactFirstnameOk returns a tuple with the SContactFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactFirstname

`func (o *ContactResponseCompound) SetSContactFirstname(v string)`

SetSContactFirstname sets SContactFirstname field to given value.


### GetSContactLastname

`func (o *ContactResponseCompound) GetSContactLastname() string`

GetSContactLastname returns the SContactLastname field if non-nil, zero value otherwise.

### GetSContactLastnameOk

`func (o *ContactResponseCompound) GetSContactLastnameOk() (*string, bool)`

GetSContactLastnameOk returns a tuple with the SContactLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactLastname

`func (o *ContactResponseCompound) SetSContactLastname(v string)`

SetSContactLastname sets SContactLastname field to given value.


### GetSContactCompany

`func (o *ContactResponseCompound) GetSContactCompany() string`

GetSContactCompany returns the SContactCompany field if non-nil, zero value otherwise.

### GetSContactCompanyOk

`func (o *ContactResponseCompound) GetSContactCompanyOk() (*string, bool)`

GetSContactCompanyOk returns a tuple with the SContactCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactCompany

`func (o *ContactResponseCompound) SetSContactCompany(v string)`

SetSContactCompany sets SContactCompany field to given value.

### HasSContactCompany

`func (o *ContactResponseCompound) HasSContactCompany() bool`

HasSContactCompany returns a boolean if a field has been set.

### GetSContactOccupation

`func (o *ContactResponseCompound) GetSContactOccupation() string`

GetSContactOccupation returns the SContactOccupation field if non-nil, zero value otherwise.

### GetSContactOccupationOk

`func (o *ContactResponseCompound) GetSContactOccupationOk() (*string, bool)`

GetSContactOccupationOk returns a tuple with the SContactOccupation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactOccupation

`func (o *ContactResponseCompound) SetSContactOccupation(v string)`

SetSContactOccupation sets SContactOccupation field to given value.

### HasSContactOccupation

`func (o *ContactResponseCompound) HasSContactOccupation() bool`

HasSContactOccupation returns a boolean if a field has been set.

### GetTContactNote

`func (o *ContactResponseCompound) GetTContactNote() string`

GetTContactNote returns the TContactNote field if non-nil, zero value otherwise.

### GetTContactNoteOk

`func (o *ContactResponseCompound) GetTContactNoteOk() (*string, bool)`

GetTContactNoteOk returns a tuple with the TContactNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTContactNote

`func (o *ContactResponseCompound) SetTContactNote(v string)`

SetTContactNote sets TContactNote field to given value.

### HasTContactNote

`func (o *ContactResponseCompound) HasTContactNote() bool`

HasTContactNote returns a boolean if a field has been set.

### GetBContactIsactive

`func (o *ContactResponseCompound) GetBContactIsactive() bool`

GetBContactIsactive returns the BContactIsactive field if non-nil, zero value otherwise.

### GetBContactIsactiveOk

`func (o *ContactResponseCompound) GetBContactIsactiveOk() (*bool, bool)`

GetBContactIsactiveOk returns a tuple with the BContactIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBContactIsactive

`func (o *ContactResponseCompound) SetBContactIsactive(v bool)`

SetBContactIsactive sets BContactIsactive field to given value.


### GetObjContactinformations

`func (o *ContactResponseCompound) GetObjContactinformations() ContactinformationsResponseCompound`

GetObjContactinformations returns the ObjContactinformations field if non-nil, zero value otherwise.

### GetObjContactinformationsOk

`func (o *ContactResponseCompound) GetObjContactinformationsOk() (*ContactinformationsResponseCompound, bool)`

GetObjContactinformationsOk returns a tuple with the ObjContactinformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContactinformations

`func (o *ContactResponseCompound) SetObjContactinformations(v ContactinformationsResponseCompound)`

SetObjContactinformations sets ObjContactinformations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


