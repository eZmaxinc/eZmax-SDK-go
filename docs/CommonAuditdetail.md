# CommonAuditdetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiUserID** | **int32** | The unique ID of the User | 
**FkiApikeyID** | Pointer to **int32** | The unique ID of the Apikey | [optional] 
**SUserLoginname** | **string** | The login name of the User. | 
**SUserLastname** | **string** | The last name of the user | 
**SUserFirstname** | **string** | The first name of the user | 
**SApikeyDescriptionX** | Pointer to **string** | The description of the Apikey in the language of the requester | [optional] 
**DtAuditdetailDate** | **string** | Represent a Date Time. The timezone is the one configured in the User&#39;s profile. | 

## Methods

### NewCommonAuditdetail

`func NewCommonAuditdetail(fkiUserID int32, sUserLoginname string, sUserLastname string, sUserFirstname string, dtAuditdetailDate string, ) *CommonAuditdetail`

NewCommonAuditdetail instantiates a new CommonAuditdetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonAuditdetailWithDefaults

`func NewCommonAuditdetailWithDefaults() *CommonAuditdetail`

NewCommonAuditdetailWithDefaults instantiates a new CommonAuditdetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiUserID

`func (o *CommonAuditdetail) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *CommonAuditdetail) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *CommonAuditdetail) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetFkiApikeyID

`func (o *CommonAuditdetail) GetFkiApikeyID() int32`

GetFkiApikeyID returns the FkiApikeyID field if non-nil, zero value otherwise.

### GetFkiApikeyIDOk

`func (o *CommonAuditdetail) GetFkiApikeyIDOk() (*int32, bool)`

GetFkiApikeyIDOk returns a tuple with the FkiApikeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiApikeyID

`func (o *CommonAuditdetail) SetFkiApikeyID(v int32)`

SetFkiApikeyID sets FkiApikeyID field to given value.

### HasFkiApikeyID

`func (o *CommonAuditdetail) HasFkiApikeyID() bool`

HasFkiApikeyID returns a boolean if a field has been set.

### GetSUserLoginname

`func (o *CommonAuditdetail) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *CommonAuditdetail) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *CommonAuditdetail) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.


### GetSUserLastname

`func (o *CommonAuditdetail) GetSUserLastname() string`

GetSUserLastname returns the SUserLastname field if non-nil, zero value otherwise.

### GetSUserLastnameOk

`func (o *CommonAuditdetail) GetSUserLastnameOk() (*string, bool)`

GetSUserLastnameOk returns a tuple with the SUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastname

`func (o *CommonAuditdetail) SetSUserLastname(v string)`

SetSUserLastname sets SUserLastname field to given value.


### GetSUserFirstname

`func (o *CommonAuditdetail) GetSUserFirstname() string`

GetSUserFirstname returns the SUserFirstname field if non-nil, zero value otherwise.

### GetSUserFirstnameOk

`func (o *CommonAuditdetail) GetSUserFirstnameOk() (*string, bool)`

GetSUserFirstnameOk returns a tuple with the SUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstname

`func (o *CommonAuditdetail) SetSUserFirstname(v string)`

SetSUserFirstname sets SUserFirstname field to given value.


### GetSApikeyDescriptionX

`func (o *CommonAuditdetail) GetSApikeyDescriptionX() string`

GetSApikeyDescriptionX returns the SApikeyDescriptionX field if non-nil, zero value otherwise.

### GetSApikeyDescriptionXOk

`func (o *CommonAuditdetail) GetSApikeyDescriptionXOk() (*string, bool)`

GetSApikeyDescriptionXOk returns a tuple with the SApikeyDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSApikeyDescriptionX

`func (o *CommonAuditdetail) SetSApikeyDescriptionX(v string)`

SetSApikeyDescriptionX sets SApikeyDescriptionX field to given value.

### HasSApikeyDescriptionX

`func (o *CommonAuditdetail) HasSApikeyDescriptionX() bool`

HasSApikeyDescriptionX returns a boolean if a field has been set.

### GetDtAuditdetailDate

`func (o *CommonAuditdetail) GetDtAuditdetailDate() string`

GetDtAuditdetailDate returns the DtAuditdetailDate field if non-nil, zero value otherwise.

### GetDtAuditdetailDateOk

`func (o *CommonAuditdetail) GetDtAuditdetailDateOk() (*string, bool)`

GetDtAuditdetailDateOk returns a tuple with the DtAuditdetailDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtAuditdetailDate

`func (o *CommonAuditdetail) SetDtAuditdetailDate(v string)`

SetDtAuditdetailDate sets DtAuditdetailDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


