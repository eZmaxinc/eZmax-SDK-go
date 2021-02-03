# ContactRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjContactinformations** | [**ContactinformationsRequestCompound**](contactinformations-RequestCompound.md) |  | 
**FkiContacttitleID** | **int32** | The unique ID of the Contacttitle.  Valid values:  |Value|Description| |-|-| |1|Ms.| |2|Mr.| |4|(Blank)| |5|Me (For Notaries)| | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SContactFirstname** | **string** | The First name of the contact | 
**SContactLastname** | **string** | The Last name of the contact | 
**SContactCompany** | **string** | The Company name of the contact | 
**DtContactBirthdate** | Pointer to **string** | The Birth Date of the contact | [optional] 

## Methods

### NewContactRequestCompound

`func NewContactRequestCompound(objContactinformations ContactinformationsRequestCompound, fkiContacttitleID int32, fkiLanguageID int32, sContactFirstname string, sContactLastname string, sContactCompany string, ) *ContactRequestCompound`

NewContactRequestCompound instantiates a new ContactRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactRequestCompoundWithDefaults

`func NewContactRequestCompoundWithDefaults() *ContactRequestCompound`

NewContactRequestCompoundWithDefaults instantiates a new ContactRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjContactinformations

`func (o *ContactRequestCompound) GetObjContactinformations() ContactinformationsRequestCompound`

GetObjContactinformations returns the ObjContactinformations field if non-nil, zero value otherwise.

### GetObjContactinformationsOk

`func (o *ContactRequestCompound) GetObjContactinformationsOk() (*ContactinformationsRequestCompound, bool)`

GetObjContactinformationsOk returns a tuple with the ObjContactinformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContactinformations

`func (o *ContactRequestCompound) SetObjContactinformations(v ContactinformationsRequestCompound)`

SetObjContactinformations sets ObjContactinformations field to given value.


### GetFkiContacttitleID

`func (o *ContactRequestCompound) GetFkiContacttitleID() int32`

GetFkiContacttitleID returns the FkiContacttitleID field if non-nil, zero value otherwise.

### GetFkiContacttitleIDOk

`func (o *ContactRequestCompound) GetFkiContacttitleIDOk() (*int32, bool)`

GetFkiContacttitleIDOk returns a tuple with the FkiContacttitleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiContacttitleID

`func (o *ContactRequestCompound) SetFkiContacttitleID(v int32)`

SetFkiContacttitleID sets FkiContacttitleID field to given value.


### GetFkiLanguageID

`func (o *ContactRequestCompound) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *ContactRequestCompound) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *ContactRequestCompound) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSContactFirstname

`func (o *ContactRequestCompound) GetSContactFirstname() string`

GetSContactFirstname returns the SContactFirstname field if non-nil, zero value otherwise.

### GetSContactFirstnameOk

`func (o *ContactRequestCompound) GetSContactFirstnameOk() (*string, bool)`

GetSContactFirstnameOk returns a tuple with the SContactFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactFirstname

`func (o *ContactRequestCompound) SetSContactFirstname(v string)`

SetSContactFirstname sets SContactFirstname field to given value.


### GetSContactLastname

`func (o *ContactRequestCompound) GetSContactLastname() string`

GetSContactLastname returns the SContactLastname field if non-nil, zero value otherwise.

### GetSContactLastnameOk

`func (o *ContactRequestCompound) GetSContactLastnameOk() (*string, bool)`

GetSContactLastnameOk returns a tuple with the SContactLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactLastname

`func (o *ContactRequestCompound) SetSContactLastname(v string)`

SetSContactLastname sets SContactLastname field to given value.


### GetSContactCompany

`func (o *ContactRequestCompound) GetSContactCompany() string`

GetSContactCompany returns the SContactCompany field if non-nil, zero value otherwise.

### GetSContactCompanyOk

`func (o *ContactRequestCompound) GetSContactCompanyOk() (*string, bool)`

GetSContactCompanyOk returns a tuple with the SContactCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactCompany

`func (o *ContactRequestCompound) SetSContactCompany(v string)`

SetSContactCompany sets SContactCompany field to given value.


### GetDtContactBirthdate

`func (o *ContactRequestCompound) GetDtContactBirthdate() string`

GetDtContactBirthdate returns the DtContactBirthdate field if non-nil, zero value otherwise.

### GetDtContactBirthdateOk

`func (o *ContactRequestCompound) GetDtContactBirthdateOk() (*string, bool)`

GetDtContactBirthdateOk returns a tuple with the DtContactBirthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtContactBirthdate

`func (o *ContactRequestCompound) SetDtContactBirthdate(v string)`

SetDtContactBirthdate sets DtContactBirthdate field to given value.

### HasDtContactBirthdate

`func (o *ContactRequestCompound) HasDtContactBirthdate() bool`

HasDtContactBirthdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


