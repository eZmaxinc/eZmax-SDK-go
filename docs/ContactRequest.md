# ContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiContacttitleID** | **int32** | The unique ID of the Contacttitle.  Valid values:  |Value|Description| |-|-| |1|Ms.| |2|Mr.| |4|(Blank)| |5|Me (For Notaries)| | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SContactFirstname** | **string** | The First name of the contact | 
**SContactLastname** | **string** | The Last name of the contact | 
**SContactCompany** | **string** | The Company name of the contact | 
**DtContactBirthdate** | Pointer to **string** | The Birth Date of the contact | [optional] 

## Methods

### NewContactRequest

`func NewContactRequest(fkiContacttitleID int32, fkiLanguageID int32, sContactFirstname string, sContactLastname string, sContactCompany string, ) *ContactRequest`

NewContactRequest instantiates a new ContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactRequestWithDefaults

`func NewContactRequestWithDefaults() *ContactRequest`

NewContactRequestWithDefaults instantiates a new ContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiContacttitleID

`func (o *ContactRequest) GetFkiContacttitleID() int32`

GetFkiContacttitleID returns the FkiContacttitleID field if non-nil, zero value otherwise.

### GetFkiContacttitleIDOk

`func (o *ContactRequest) GetFkiContacttitleIDOk() (*int32, bool)`

GetFkiContacttitleIDOk returns a tuple with the FkiContacttitleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiContacttitleID

`func (o *ContactRequest) SetFkiContacttitleID(v int32)`

SetFkiContacttitleID sets FkiContacttitleID field to given value.


### GetFkiLanguageID

`func (o *ContactRequest) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *ContactRequest) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *ContactRequest) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSContactFirstname

`func (o *ContactRequest) GetSContactFirstname() string`

GetSContactFirstname returns the SContactFirstname field if non-nil, zero value otherwise.

### GetSContactFirstnameOk

`func (o *ContactRequest) GetSContactFirstnameOk() (*string, bool)`

GetSContactFirstnameOk returns a tuple with the SContactFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactFirstname

`func (o *ContactRequest) SetSContactFirstname(v string)`

SetSContactFirstname sets SContactFirstname field to given value.


### GetSContactLastname

`func (o *ContactRequest) GetSContactLastname() string`

GetSContactLastname returns the SContactLastname field if non-nil, zero value otherwise.

### GetSContactLastnameOk

`func (o *ContactRequest) GetSContactLastnameOk() (*string, bool)`

GetSContactLastnameOk returns a tuple with the SContactLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactLastname

`func (o *ContactRequest) SetSContactLastname(v string)`

SetSContactLastname sets SContactLastname field to given value.


### GetSContactCompany

`func (o *ContactRequest) GetSContactCompany() string`

GetSContactCompany returns the SContactCompany field if non-nil, zero value otherwise.

### GetSContactCompanyOk

`func (o *ContactRequest) GetSContactCompanyOk() (*string, bool)`

GetSContactCompanyOk returns a tuple with the SContactCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactCompany

`func (o *ContactRequest) SetSContactCompany(v string)`

SetSContactCompany sets SContactCompany field to given value.


### GetDtContactBirthdate

`func (o *ContactRequest) GetDtContactBirthdate() string`

GetDtContactBirthdate returns the DtContactBirthdate field if non-nil, zero value otherwise.

### GetDtContactBirthdateOk

`func (o *ContactRequest) GetDtContactBirthdateOk() (*string, bool)`

GetDtContactBirthdateOk returns a tuple with the DtContactBirthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtContactBirthdate

`func (o *ContactRequest) SetDtContactBirthdate(v string)`

SetDtContactBirthdate sets DtContactBirthdate field to given value.

### HasDtContactBirthdate

`func (o *ContactRequest) HasDtContactBirthdate() bool`

HasDtContactBirthdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


