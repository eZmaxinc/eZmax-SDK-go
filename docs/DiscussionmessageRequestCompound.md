# DiscussionmessageRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDiscussionmessageID** | Pointer to **int32** | The unique ID of the Discussionmessage | [optional] 
**FkiDiscussionID** | **int32** | The unique ID of the Discussion | 
**FkiDiscussionmembershipIDActionrequired** | Pointer to **int32** | The unique ID of the Discussionmembership | [optional] 
**TDiscussionmessageContent** | **string** | The content of the Discussionmessage | 

## Methods

### NewDiscussionmessageRequestCompound

`func NewDiscussionmessageRequestCompound(fkiDiscussionID int32, tDiscussionmessageContent string, ) *DiscussionmessageRequestCompound`

NewDiscussionmessageRequestCompound instantiates a new DiscussionmessageRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscussionmessageRequestCompoundWithDefaults

`func NewDiscussionmessageRequestCompoundWithDefaults() *DiscussionmessageRequestCompound`

NewDiscussionmessageRequestCompoundWithDefaults instantiates a new DiscussionmessageRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDiscussionmessageID

`func (o *DiscussionmessageRequestCompound) GetPkiDiscussionmessageID() int32`

GetPkiDiscussionmessageID returns the PkiDiscussionmessageID field if non-nil, zero value otherwise.

### GetPkiDiscussionmessageIDOk

`func (o *DiscussionmessageRequestCompound) GetPkiDiscussionmessageIDOk() (*int32, bool)`

GetPkiDiscussionmessageIDOk returns a tuple with the PkiDiscussionmessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDiscussionmessageID

`func (o *DiscussionmessageRequestCompound) SetPkiDiscussionmessageID(v int32)`

SetPkiDiscussionmessageID sets PkiDiscussionmessageID field to given value.

### HasPkiDiscussionmessageID

`func (o *DiscussionmessageRequestCompound) HasPkiDiscussionmessageID() bool`

HasPkiDiscussionmessageID returns a boolean if a field has been set.

### GetFkiDiscussionID

`func (o *DiscussionmessageRequestCompound) GetFkiDiscussionID() int32`

GetFkiDiscussionID returns the FkiDiscussionID field if non-nil, zero value otherwise.

### GetFkiDiscussionIDOk

`func (o *DiscussionmessageRequestCompound) GetFkiDiscussionIDOk() (*int32, bool)`

GetFkiDiscussionIDOk returns a tuple with the FkiDiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDiscussionID

`func (o *DiscussionmessageRequestCompound) SetFkiDiscussionID(v int32)`

SetFkiDiscussionID sets FkiDiscussionID field to given value.


### GetFkiDiscussionmembershipIDActionrequired

`func (o *DiscussionmessageRequestCompound) GetFkiDiscussionmembershipIDActionrequired() int32`

GetFkiDiscussionmembershipIDActionrequired returns the FkiDiscussionmembershipIDActionrequired field if non-nil, zero value otherwise.

### GetFkiDiscussionmembershipIDActionrequiredOk

`func (o *DiscussionmessageRequestCompound) GetFkiDiscussionmembershipIDActionrequiredOk() (*int32, bool)`

GetFkiDiscussionmembershipIDActionrequiredOk returns a tuple with the FkiDiscussionmembershipIDActionrequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDiscussionmembershipIDActionrequired

`func (o *DiscussionmessageRequestCompound) SetFkiDiscussionmembershipIDActionrequired(v int32)`

SetFkiDiscussionmembershipIDActionrequired sets FkiDiscussionmembershipIDActionrequired field to given value.

### HasFkiDiscussionmembershipIDActionrequired

`func (o *DiscussionmessageRequestCompound) HasFkiDiscussionmembershipIDActionrequired() bool`

HasFkiDiscussionmembershipIDActionrequired returns a boolean if a field has been set.

### GetTDiscussionmessageContent

`func (o *DiscussionmessageRequestCompound) GetTDiscussionmessageContent() string`

GetTDiscussionmessageContent returns the TDiscussionmessageContent field if non-nil, zero value otherwise.

### GetTDiscussionmessageContentOk

`func (o *DiscussionmessageRequestCompound) GetTDiscussionmessageContentOk() (*string, bool)`

GetTDiscussionmessageContentOk returns a tuple with the TDiscussionmessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTDiscussionmessageContent

`func (o *DiscussionmessageRequestCompound) SetTDiscussionmessageContent(v string)`

SetTDiscussionmessageContent sets TDiscussionmessageContent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


