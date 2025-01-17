# ContactRequestCompoundV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiContacttitleID** | **int32** | The unique ID of the Contacttitle.  Valid values:  |Value|Description| |-|-| |1|Ms.| |2|Mr.| |4|(Blank)| |5|Me (For Notaries)| | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**EContactType** | [**FieldEContactType**](FieldEContactType.md) |  | 
**SContactFirstname** | **string** | The First name of the contact | 
**SContactLastname** | **string** | The Last name of the contact | 
**SContactCompany** | Pointer to **string** | The Company name of the contact | [optional] 
**DtContactBirthdate** | Pointer to **string** | The Birth Date of the contact | [optional] 
**SContactOccupation** | Pointer to **string** | The occupation of the Contact | [optional] 
**TContactNote** | Pointer to **string** | The note of the Contact | [optional] 
**BContactIsactive** | Pointer to **bool** | Whether the contact is active or not | [optional] 
**ObjContactinformations** | [**ContactinformationsRequestCompoundV2**](ContactinformationsRequestCompoundV2.md) |  | 

## Methods

### NewContactRequestCompoundV2

`func NewContactRequestCompoundV2(fkiContacttitleID int32, fkiLanguageID int32, eContactType FieldEContactType, sContactFirstname string, sContactLastname string, objContactinformations ContactinformationsRequestCompoundV2, ) *ContactRequestCompoundV2`

NewContactRequestCompoundV2 instantiates a new ContactRequestCompoundV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactRequestCompoundV2WithDefaults

`func NewContactRequestCompoundV2WithDefaults() *ContactRequestCompoundV2`

NewContactRequestCompoundV2WithDefaults instantiates a new ContactRequestCompoundV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiContacttitleID

`func (o *ContactRequestCompoundV2) GetFkiContacttitleID() int32`

GetFkiContacttitleID returns the FkiContacttitleID field if non-nil, zero value otherwise.

### GetFkiContacttitleIDOk

`func (o *ContactRequestCompoundV2) GetFkiContacttitleIDOk() (*int32, bool)`

GetFkiContacttitleIDOk returns a tuple with the FkiContacttitleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiContacttitleID

`func (o *ContactRequestCompoundV2) SetFkiContacttitleID(v int32)`

SetFkiContacttitleID sets FkiContacttitleID field to given value.


### GetFkiLanguageID

`func (o *ContactRequestCompoundV2) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *ContactRequestCompoundV2) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *ContactRequestCompoundV2) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetEContactType

`func (o *ContactRequestCompoundV2) GetEContactType() FieldEContactType`

GetEContactType returns the EContactType field if non-nil, zero value otherwise.

### GetEContactTypeOk

`func (o *ContactRequestCompoundV2) GetEContactTypeOk() (*FieldEContactType, bool)`

GetEContactTypeOk returns a tuple with the EContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEContactType

`func (o *ContactRequestCompoundV2) SetEContactType(v FieldEContactType)`

SetEContactType sets EContactType field to given value.


### GetSContactFirstname

`func (o *ContactRequestCompoundV2) GetSContactFirstname() string`

GetSContactFirstname returns the SContactFirstname field if non-nil, zero value otherwise.

### GetSContactFirstnameOk

`func (o *ContactRequestCompoundV2) GetSContactFirstnameOk() (*string, bool)`

GetSContactFirstnameOk returns a tuple with the SContactFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactFirstname

`func (o *ContactRequestCompoundV2) SetSContactFirstname(v string)`

SetSContactFirstname sets SContactFirstname field to given value.


### GetSContactLastname

`func (o *ContactRequestCompoundV2) GetSContactLastname() string`

GetSContactLastname returns the SContactLastname field if non-nil, zero value otherwise.

### GetSContactLastnameOk

`func (o *ContactRequestCompoundV2) GetSContactLastnameOk() (*string, bool)`

GetSContactLastnameOk returns a tuple with the SContactLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactLastname

`func (o *ContactRequestCompoundV2) SetSContactLastname(v string)`

SetSContactLastname sets SContactLastname field to given value.


### GetSContactCompany

`func (o *ContactRequestCompoundV2) GetSContactCompany() string`

GetSContactCompany returns the SContactCompany field if non-nil, zero value otherwise.

### GetSContactCompanyOk

`func (o *ContactRequestCompoundV2) GetSContactCompanyOk() (*string, bool)`

GetSContactCompanyOk returns a tuple with the SContactCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactCompany

`func (o *ContactRequestCompoundV2) SetSContactCompany(v string)`

SetSContactCompany sets SContactCompany field to given value.

### HasSContactCompany

`func (o *ContactRequestCompoundV2) HasSContactCompany() bool`

HasSContactCompany returns a boolean if a field has been set.

### GetDtContactBirthdate

`func (o *ContactRequestCompoundV2) GetDtContactBirthdate() string`

GetDtContactBirthdate returns the DtContactBirthdate field if non-nil, zero value otherwise.

### GetDtContactBirthdateOk

`func (o *ContactRequestCompoundV2) GetDtContactBirthdateOk() (*string, bool)`

GetDtContactBirthdateOk returns a tuple with the DtContactBirthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtContactBirthdate

`func (o *ContactRequestCompoundV2) SetDtContactBirthdate(v string)`

SetDtContactBirthdate sets DtContactBirthdate field to given value.

### HasDtContactBirthdate

`func (o *ContactRequestCompoundV2) HasDtContactBirthdate() bool`

HasDtContactBirthdate returns a boolean if a field has been set.

### GetSContactOccupation

`func (o *ContactRequestCompoundV2) GetSContactOccupation() string`

GetSContactOccupation returns the SContactOccupation field if non-nil, zero value otherwise.

### GetSContactOccupationOk

`func (o *ContactRequestCompoundV2) GetSContactOccupationOk() (*string, bool)`

GetSContactOccupationOk returns a tuple with the SContactOccupation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactOccupation

`func (o *ContactRequestCompoundV2) SetSContactOccupation(v string)`

SetSContactOccupation sets SContactOccupation field to given value.

### HasSContactOccupation

`func (o *ContactRequestCompoundV2) HasSContactOccupation() bool`

HasSContactOccupation returns a boolean if a field has been set.

### GetTContactNote

`func (o *ContactRequestCompoundV2) GetTContactNote() string`

GetTContactNote returns the TContactNote field if non-nil, zero value otherwise.

### GetTContactNoteOk

`func (o *ContactRequestCompoundV2) GetTContactNoteOk() (*string, bool)`

GetTContactNoteOk returns a tuple with the TContactNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTContactNote

`func (o *ContactRequestCompoundV2) SetTContactNote(v string)`

SetTContactNote sets TContactNote field to given value.

### HasTContactNote

`func (o *ContactRequestCompoundV2) HasTContactNote() bool`

HasTContactNote returns a boolean if a field has been set.

### GetBContactIsactive

`func (o *ContactRequestCompoundV2) GetBContactIsactive() bool`

GetBContactIsactive returns the BContactIsactive field if non-nil, zero value otherwise.

### GetBContactIsactiveOk

`func (o *ContactRequestCompoundV2) GetBContactIsactiveOk() (*bool, bool)`

GetBContactIsactiveOk returns a tuple with the BContactIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBContactIsactive

`func (o *ContactRequestCompoundV2) SetBContactIsactive(v bool)`

SetBContactIsactive sets BContactIsactive field to given value.

### HasBContactIsactive

`func (o *ContactRequestCompoundV2) HasBContactIsactive() bool`

HasBContactIsactive returns a boolean if a field has been set.

### GetObjContactinformations

`func (o *ContactRequestCompoundV2) GetObjContactinformations() ContactinformationsRequestCompoundV2`

GetObjContactinformations returns the ObjContactinformations field if non-nil, zero value otherwise.

### GetObjContactinformationsOk

`func (o *ContactRequestCompoundV2) GetObjContactinformationsOk() (*ContactinformationsRequestCompoundV2, bool)`

GetObjContactinformationsOk returns a tuple with the ObjContactinformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContactinformations

`func (o *ContactRequestCompoundV2) SetObjContactinformations(v ContactinformationsRequestCompoundV2)`

SetObjContactinformations sets ObjContactinformations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


