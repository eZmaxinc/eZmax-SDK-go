# DiscussionmessageResponse

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

### NewDiscussionmessageResponse

`func NewDiscussionmessageResponse(pkiDiscussionmessageID int32, fkiDiscussionID int32, eDiscussionmessageStatus FieldEDiscussionmessageStatus, tDiscussionmessageContent string, sDiscussionmessageCreatorname string, objAudit CommonAudit, ) *DiscussionmessageResponse`

NewDiscussionmessageResponse instantiates a new DiscussionmessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscussionmessageResponseWithDefaults

`func NewDiscussionmessageResponseWithDefaults() *DiscussionmessageResponse`

NewDiscussionmessageResponseWithDefaults instantiates a new DiscussionmessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDiscussionmessageID

`func (o *DiscussionmessageResponse) GetPkiDiscussionmessageID() int32`

GetPkiDiscussionmessageID returns the PkiDiscussionmessageID field if non-nil, zero value otherwise.

### GetPkiDiscussionmessageIDOk

`func (o *DiscussionmessageResponse) GetPkiDiscussionmessageIDOk() (*int32, bool)`

GetPkiDiscussionmessageIDOk returns a tuple with the PkiDiscussionmessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDiscussionmessageID

`func (o *DiscussionmessageResponse) SetPkiDiscussionmessageID(v int32)`

SetPkiDiscussionmessageID sets PkiDiscussionmessageID field to given value.


### GetFkiDiscussionID

`func (o *DiscussionmessageResponse) GetFkiDiscussionID() int32`

GetFkiDiscussionID returns the FkiDiscussionID field if non-nil, zero value otherwise.

### GetFkiDiscussionIDOk

`func (o *DiscussionmessageResponse) GetFkiDiscussionIDOk() (*int32, bool)`

GetFkiDiscussionIDOk returns a tuple with the FkiDiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDiscussionID

`func (o *DiscussionmessageResponse) SetFkiDiscussionID(v int32)`

SetFkiDiscussionID sets FkiDiscussionID field to given value.


### GetFkiDiscussionmembershipID

`func (o *DiscussionmessageResponse) GetFkiDiscussionmembershipID() int32`

GetFkiDiscussionmembershipID returns the FkiDiscussionmembershipID field if non-nil, zero value otherwise.

### GetFkiDiscussionmembershipIDOk

`func (o *DiscussionmessageResponse) GetFkiDiscussionmembershipIDOk() (*int32, bool)`

GetFkiDiscussionmembershipIDOk returns a tuple with the FkiDiscussionmembershipID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDiscussionmembershipID

`func (o *DiscussionmessageResponse) SetFkiDiscussionmembershipID(v int32)`

SetFkiDiscussionmembershipID sets FkiDiscussionmembershipID field to given value.

### HasFkiDiscussionmembershipID

`func (o *DiscussionmessageResponse) HasFkiDiscussionmembershipID() bool`

HasFkiDiscussionmembershipID returns a boolean if a field has been set.

### GetFkiDiscussionmembershipIDActionrequired

`func (o *DiscussionmessageResponse) GetFkiDiscussionmembershipIDActionrequired() int32`

GetFkiDiscussionmembershipIDActionrequired returns the FkiDiscussionmembershipIDActionrequired field if non-nil, zero value otherwise.

### GetFkiDiscussionmembershipIDActionrequiredOk

`func (o *DiscussionmessageResponse) GetFkiDiscussionmembershipIDActionrequiredOk() (*int32, bool)`

GetFkiDiscussionmembershipIDActionrequiredOk returns a tuple with the FkiDiscussionmembershipIDActionrequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDiscussionmembershipIDActionrequired

`func (o *DiscussionmessageResponse) SetFkiDiscussionmembershipIDActionrequired(v int32)`

SetFkiDiscussionmembershipIDActionrequired sets FkiDiscussionmembershipIDActionrequired field to given value.

### HasFkiDiscussionmembershipIDActionrequired

`func (o *DiscussionmessageResponse) HasFkiDiscussionmembershipIDActionrequired() bool`

HasFkiDiscussionmembershipIDActionrequired returns a boolean if a field has been set.

### GetEDiscussionmessageStatus

`func (o *DiscussionmessageResponse) GetEDiscussionmessageStatus() FieldEDiscussionmessageStatus`

GetEDiscussionmessageStatus returns the EDiscussionmessageStatus field if non-nil, zero value otherwise.

### GetEDiscussionmessageStatusOk

`func (o *DiscussionmessageResponse) GetEDiscussionmessageStatusOk() (*FieldEDiscussionmessageStatus, bool)`

GetEDiscussionmessageStatusOk returns a tuple with the EDiscussionmessageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDiscussionmessageStatus

`func (o *DiscussionmessageResponse) SetEDiscussionmessageStatus(v FieldEDiscussionmessageStatus)`

SetEDiscussionmessageStatus sets EDiscussionmessageStatus field to given value.


### GetTDiscussionmessageContent

`func (o *DiscussionmessageResponse) GetTDiscussionmessageContent() string`

GetTDiscussionmessageContent returns the TDiscussionmessageContent field if non-nil, zero value otherwise.

### GetTDiscussionmessageContentOk

`func (o *DiscussionmessageResponse) GetTDiscussionmessageContentOk() (*string, bool)`

GetTDiscussionmessageContentOk returns a tuple with the TDiscussionmessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTDiscussionmessageContent

`func (o *DiscussionmessageResponse) SetTDiscussionmessageContent(v string)`

SetTDiscussionmessageContent sets TDiscussionmessageContent field to given value.


### GetSDiscussionmessageCreatorname

`func (o *DiscussionmessageResponse) GetSDiscussionmessageCreatorname() string`

GetSDiscussionmessageCreatorname returns the SDiscussionmessageCreatorname field if non-nil, zero value otherwise.

### GetSDiscussionmessageCreatornameOk

`func (o *DiscussionmessageResponse) GetSDiscussionmessageCreatornameOk() (*string, bool)`

GetSDiscussionmessageCreatornameOk returns a tuple with the SDiscussionmessageCreatorname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDiscussionmessageCreatorname

`func (o *DiscussionmessageResponse) SetSDiscussionmessageCreatorname(v string)`

SetSDiscussionmessageCreatorname sets SDiscussionmessageCreatorname field to given value.


### GetSDiscussionmessageActionrequiredname

`func (o *DiscussionmessageResponse) GetSDiscussionmessageActionrequiredname() string`

GetSDiscussionmessageActionrequiredname returns the SDiscussionmessageActionrequiredname field if non-nil, zero value otherwise.

### GetSDiscussionmessageActionrequirednameOk

`func (o *DiscussionmessageResponse) GetSDiscussionmessageActionrequirednameOk() (*string, bool)`

GetSDiscussionmessageActionrequirednameOk returns a tuple with the SDiscussionmessageActionrequiredname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDiscussionmessageActionrequiredname

`func (o *DiscussionmessageResponse) SetSDiscussionmessageActionrequiredname(v string)`

SetSDiscussionmessageActionrequiredname sets SDiscussionmessageActionrequiredname field to given value.

### HasSDiscussionmessageActionrequiredname

`func (o *DiscussionmessageResponse) HasSDiscussionmessageActionrequiredname() bool`

HasSDiscussionmessageActionrequiredname returns a boolean if a field has been set.

### GetObjAudit

`func (o *DiscussionmessageResponse) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *DiscussionmessageResponse) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *DiscussionmessageResponse) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


