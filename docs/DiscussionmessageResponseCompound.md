# DiscussionmessageResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDiscussionmessageID** | **int32** | The unique ID of the Discussionmessage | 
**FkiDiscussionID** | **int32** | The unique ID of the Discussion | 
**FkiDiscussionmembershipID** | Pointer to **int32** | The unique ID of the Discussionmembership | [optional] 
**FkiDiscussionmembershipIDActionrequired** | Pointer to **int32** | The unique ID of the Discussionmembership | [optional] 
**EDiscussionmessageStatus** | [**FieldEDiscussionmessageStatus**](FieldEDiscussionmessageStatus.md) |  | 
**TDiscussionmessageContent** | **string** | The content of the Discussionmessage | 
**SDiscussionmessageCreatorname** | **string** | The name the creator of the Discussionmessage. | 
**SDiscussionmessageActionrequiredname** | Pointer to **string** | The name the Actionrequired of the Discussionmessage. | [optional] 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 

## Methods

### NewDiscussionmessageResponseCompound

`func NewDiscussionmessageResponseCompound(pkiDiscussionmessageID int32, fkiDiscussionID int32, eDiscussionmessageStatus FieldEDiscussionmessageStatus, tDiscussionmessageContent string, sDiscussionmessageCreatorname string, objAudit CommonAudit, ) *DiscussionmessageResponseCompound`

NewDiscussionmessageResponseCompound instantiates a new DiscussionmessageResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscussionmessageResponseCompoundWithDefaults

`func NewDiscussionmessageResponseCompoundWithDefaults() *DiscussionmessageResponseCompound`

NewDiscussionmessageResponseCompoundWithDefaults instantiates a new DiscussionmessageResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDiscussionmessageID

`func (o *DiscussionmessageResponseCompound) GetPkiDiscussionmessageID() int32`

GetPkiDiscussionmessageID returns the PkiDiscussionmessageID field if non-nil, zero value otherwise.

### GetPkiDiscussionmessageIDOk

`func (o *DiscussionmessageResponseCompound) GetPkiDiscussionmessageIDOk() (*int32, bool)`

GetPkiDiscussionmessageIDOk returns a tuple with the PkiDiscussionmessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDiscussionmessageID

`func (o *DiscussionmessageResponseCompound) SetPkiDiscussionmessageID(v int32)`

SetPkiDiscussionmessageID sets PkiDiscussionmessageID field to given value.


### GetFkiDiscussionID

`func (o *DiscussionmessageResponseCompound) GetFkiDiscussionID() int32`

GetFkiDiscussionID returns the FkiDiscussionID field if non-nil, zero value otherwise.

### GetFkiDiscussionIDOk

`func (o *DiscussionmessageResponseCompound) GetFkiDiscussionIDOk() (*int32, bool)`

GetFkiDiscussionIDOk returns a tuple with the FkiDiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDiscussionID

`func (o *DiscussionmessageResponseCompound) SetFkiDiscussionID(v int32)`

SetFkiDiscussionID sets FkiDiscussionID field to given value.


### GetFkiDiscussionmembershipID

`func (o *DiscussionmessageResponseCompound) GetFkiDiscussionmembershipID() int32`

GetFkiDiscussionmembershipID returns the FkiDiscussionmembershipID field if non-nil, zero value otherwise.

### GetFkiDiscussionmembershipIDOk

`func (o *DiscussionmessageResponseCompound) GetFkiDiscussionmembershipIDOk() (*int32, bool)`

GetFkiDiscussionmembershipIDOk returns a tuple with the FkiDiscussionmembershipID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDiscussionmembershipID

`func (o *DiscussionmessageResponseCompound) SetFkiDiscussionmembershipID(v int32)`

SetFkiDiscussionmembershipID sets FkiDiscussionmembershipID field to given value.

### HasFkiDiscussionmembershipID

`func (o *DiscussionmessageResponseCompound) HasFkiDiscussionmembershipID() bool`

HasFkiDiscussionmembershipID returns a boolean if a field has been set.

### GetFkiDiscussionmembershipIDActionrequired

`func (o *DiscussionmessageResponseCompound) GetFkiDiscussionmembershipIDActionrequired() int32`

GetFkiDiscussionmembershipIDActionrequired returns the FkiDiscussionmembershipIDActionrequired field if non-nil, zero value otherwise.

### GetFkiDiscussionmembershipIDActionrequiredOk

`func (o *DiscussionmessageResponseCompound) GetFkiDiscussionmembershipIDActionrequiredOk() (*int32, bool)`

GetFkiDiscussionmembershipIDActionrequiredOk returns a tuple with the FkiDiscussionmembershipIDActionrequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDiscussionmembershipIDActionrequired

`func (o *DiscussionmessageResponseCompound) SetFkiDiscussionmembershipIDActionrequired(v int32)`

SetFkiDiscussionmembershipIDActionrequired sets FkiDiscussionmembershipIDActionrequired field to given value.

### HasFkiDiscussionmembershipIDActionrequired

`func (o *DiscussionmessageResponseCompound) HasFkiDiscussionmembershipIDActionrequired() bool`

HasFkiDiscussionmembershipIDActionrequired returns a boolean if a field has been set.

### GetEDiscussionmessageStatus

`func (o *DiscussionmessageResponseCompound) GetEDiscussionmessageStatus() FieldEDiscussionmessageStatus`

GetEDiscussionmessageStatus returns the EDiscussionmessageStatus field if non-nil, zero value otherwise.

### GetEDiscussionmessageStatusOk

`func (o *DiscussionmessageResponseCompound) GetEDiscussionmessageStatusOk() (*FieldEDiscussionmessageStatus, bool)`

GetEDiscussionmessageStatusOk returns a tuple with the EDiscussionmessageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDiscussionmessageStatus

`func (o *DiscussionmessageResponseCompound) SetEDiscussionmessageStatus(v FieldEDiscussionmessageStatus)`

SetEDiscussionmessageStatus sets EDiscussionmessageStatus field to given value.


### GetTDiscussionmessageContent

`func (o *DiscussionmessageResponseCompound) GetTDiscussionmessageContent() string`

GetTDiscussionmessageContent returns the TDiscussionmessageContent field if non-nil, zero value otherwise.

### GetTDiscussionmessageContentOk

`func (o *DiscussionmessageResponseCompound) GetTDiscussionmessageContentOk() (*string, bool)`

GetTDiscussionmessageContentOk returns a tuple with the TDiscussionmessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTDiscussionmessageContent

`func (o *DiscussionmessageResponseCompound) SetTDiscussionmessageContent(v string)`

SetTDiscussionmessageContent sets TDiscussionmessageContent field to given value.


### GetSDiscussionmessageCreatorname

`func (o *DiscussionmessageResponseCompound) GetSDiscussionmessageCreatorname() string`

GetSDiscussionmessageCreatorname returns the SDiscussionmessageCreatorname field if non-nil, zero value otherwise.

### GetSDiscussionmessageCreatornameOk

`func (o *DiscussionmessageResponseCompound) GetSDiscussionmessageCreatornameOk() (*string, bool)`

GetSDiscussionmessageCreatornameOk returns a tuple with the SDiscussionmessageCreatorname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDiscussionmessageCreatorname

`func (o *DiscussionmessageResponseCompound) SetSDiscussionmessageCreatorname(v string)`

SetSDiscussionmessageCreatorname sets SDiscussionmessageCreatorname field to given value.


### GetSDiscussionmessageActionrequiredname

`func (o *DiscussionmessageResponseCompound) GetSDiscussionmessageActionrequiredname() string`

GetSDiscussionmessageActionrequiredname returns the SDiscussionmessageActionrequiredname field if non-nil, zero value otherwise.

### GetSDiscussionmessageActionrequirednameOk

`func (o *DiscussionmessageResponseCompound) GetSDiscussionmessageActionrequirednameOk() (*string, bool)`

GetSDiscussionmessageActionrequirednameOk returns a tuple with the SDiscussionmessageActionrequiredname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDiscussionmessageActionrequiredname

`func (o *DiscussionmessageResponseCompound) SetSDiscussionmessageActionrequiredname(v string)`

SetSDiscussionmessageActionrequiredname sets SDiscussionmessageActionrequiredname field to given value.

### HasSDiscussionmessageActionrequiredname

`func (o *DiscussionmessageResponseCompound) HasSDiscussionmessageActionrequiredname() bool`

HasSDiscussionmessageActionrequiredname returns a boolean if a field has been set.

### GetObjAudit

`func (o *DiscussionmessageResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *DiscussionmessageResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *DiscussionmessageResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


