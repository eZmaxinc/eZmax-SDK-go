# CommonAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiUserIDCreated** | **int32** | The unique ID of the User | 
**FkiUserIDModified** | **int32** | The unique ID of the User | 
**FkiApikeyIDCreated** | Pointer to **int32** | The unique ID of the Apikey | [optional] 
**FkiApikeyIDModified** | Pointer to **int32** | The unique ID of the Apikey | [optional] 
**DtCreatedDate** | **string** | Represent a Date Time. The timezone is the one configured in the User&#39;s profile. | 
**DtModifiedDate** | **string** | Represent a Date Time. The timezone is the one configured in the User&#39;s profile. | 

## Methods

### NewCommonAudit

`func NewCommonAudit(fkiUserIDCreated int32, fkiUserIDModified int32, dtCreatedDate string, dtModifiedDate string, ) *CommonAudit`

NewCommonAudit instantiates a new CommonAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonAuditWithDefaults

`func NewCommonAuditWithDefaults() *CommonAudit`

NewCommonAuditWithDefaults instantiates a new CommonAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiUserIDCreated

`func (o *CommonAudit) GetFkiUserIDCreated() int32`

GetFkiUserIDCreated returns the FkiUserIDCreated field if non-nil, zero value otherwise.

### GetFkiUserIDCreatedOk

`func (o *CommonAudit) GetFkiUserIDCreatedOk() (*int32, bool)`

GetFkiUserIDCreatedOk returns a tuple with the FkiUserIDCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDCreated

`func (o *CommonAudit) SetFkiUserIDCreated(v int32)`

SetFkiUserIDCreated sets FkiUserIDCreated field to given value.


### GetFkiUserIDModified

`func (o *CommonAudit) GetFkiUserIDModified() int32`

GetFkiUserIDModified returns the FkiUserIDModified field if non-nil, zero value otherwise.

### GetFkiUserIDModifiedOk

`func (o *CommonAudit) GetFkiUserIDModifiedOk() (*int32, bool)`

GetFkiUserIDModifiedOk returns a tuple with the FkiUserIDModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDModified

`func (o *CommonAudit) SetFkiUserIDModified(v int32)`

SetFkiUserIDModified sets FkiUserIDModified field to given value.


### GetFkiApikeyIDCreated

`func (o *CommonAudit) GetFkiApikeyIDCreated() int32`

GetFkiApikeyIDCreated returns the FkiApikeyIDCreated field if non-nil, zero value otherwise.

### GetFkiApikeyIDCreatedOk

`func (o *CommonAudit) GetFkiApikeyIDCreatedOk() (*int32, bool)`

GetFkiApikeyIDCreatedOk returns a tuple with the FkiApikeyIDCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiApikeyIDCreated

`func (o *CommonAudit) SetFkiApikeyIDCreated(v int32)`

SetFkiApikeyIDCreated sets FkiApikeyIDCreated field to given value.

### HasFkiApikeyIDCreated

`func (o *CommonAudit) HasFkiApikeyIDCreated() bool`

HasFkiApikeyIDCreated returns a boolean if a field has been set.

### GetFkiApikeyIDModified

`func (o *CommonAudit) GetFkiApikeyIDModified() int32`

GetFkiApikeyIDModified returns the FkiApikeyIDModified field if non-nil, zero value otherwise.

### GetFkiApikeyIDModifiedOk

`func (o *CommonAudit) GetFkiApikeyIDModifiedOk() (*int32, bool)`

GetFkiApikeyIDModifiedOk returns a tuple with the FkiApikeyIDModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiApikeyIDModified

`func (o *CommonAudit) SetFkiApikeyIDModified(v int32)`

SetFkiApikeyIDModified sets FkiApikeyIDModified field to given value.

### HasFkiApikeyIDModified

`func (o *CommonAudit) HasFkiApikeyIDModified() bool`

HasFkiApikeyIDModified returns a boolean if a field has been set.

### GetDtCreatedDate

`func (o *CommonAudit) GetDtCreatedDate() string`

GetDtCreatedDate returns the DtCreatedDate field if non-nil, zero value otherwise.

### GetDtCreatedDateOk

`func (o *CommonAudit) GetDtCreatedDateOk() (*string, bool)`

GetDtCreatedDateOk returns a tuple with the DtCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtCreatedDate

`func (o *CommonAudit) SetDtCreatedDate(v string)`

SetDtCreatedDate sets DtCreatedDate field to given value.


### GetDtModifiedDate

`func (o *CommonAudit) GetDtModifiedDate() string`

GetDtModifiedDate returns the DtModifiedDate field if non-nil, zero value otherwise.

### GetDtModifiedDateOk

`func (o *CommonAudit) GetDtModifiedDateOk() (*string, bool)`

GetDtModifiedDateOk returns a tuple with the DtModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtModifiedDate

`func (o *CommonAudit) SetDtModifiedDate(v string)`

SetDtModifiedDate sets DtModifiedDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


