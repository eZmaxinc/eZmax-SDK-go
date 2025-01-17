# DiscussionmembershipResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDiscussionmembershipID** | **int32** | The unique ID of the Discussionmembership | 
**FkiDiscussionID** | **int32** | The unique ID of the Discussion | 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiUsergroupID** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**FkiModulesectionID** | Pointer to **int32** | The unique ID of the Modulesection | [optional] 
**SDiscussionmembershipDescription** | **string** | The Description containing the detail of who the Discussionmembership refers to | 
**DtDiscussionmembershipJoined** | **string** | The joined date of the Discussionmembership | 

## Methods

### NewDiscussionmembershipResponseCompound

`func NewDiscussionmembershipResponseCompound(pkiDiscussionmembershipID int32, fkiDiscussionID int32, sDiscussionmembershipDescription string, dtDiscussionmembershipJoined string, ) *DiscussionmembershipResponseCompound`

NewDiscussionmembershipResponseCompound instantiates a new DiscussionmembershipResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscussionmembershipResponseCompoundWithDefaults

`func NewDiscussionmembershipResponseCompoundWithDefaults() *DiscussionmembershipResponseCompound`

NewDiscussionmembershipResponseCompoundWithDefaults instantiates a new DiscussionmembershipResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDiscussionmembershipID

`func (o *DiscussionmembershipResponseCompound) GetPkiDiscussionmembershipID() int32`

GetPkiDiscussionmembershipID returns the PkiDiscussionmembershipID field if non-nil, zero value otherwise.

### GetPkiDiscussionmembershipIDOk

`func (o *DiscussionmembershipResponseCompound) GetPkiDiscussionmembershipIDOk() (*int32, bool)`

GetPkiDiscussionmembershipIDOk returns a tuple with the PkiDiscussionmembershipID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDiscussionmembershipID

`func (o *DiscussionmembershipResponseCompound) SetPkiDiscussionmembershipID(v int32)`

SetPkiDiscussionmembershipID sets PkiDiscussionmembershipID field to given value.


### GetFkiDiscussionID

`func (o *DiscussionmembershipResponseCompound) GetFkiDiscussionID() int32`

GetFkiDiscussionID returns the FkiDiscussionID field if non-nil, zero value otherwise.

### GetFkiDiscussionIDOk

`func (o *DiscussionmembershipResponseCompound) GetFkiDiscussionIDOk() (*int32, bool)`

GetFkiDiscussionIDOk returns a tuple with the FkiDiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDiscussionID

`func (o *DiscussionmembershipResponseCompound) SetFkiDiscussionID(v int32)`

SetFkiDiscussionID sets FkiDiscussionID field to given value.


### GetFkiUserID

`func (o *DiscussionmembershipResponseCompound) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *DiscussionmembershipResponseCompound) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *DiscussionmembershipResponseCompound) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *DiscussionmembershipResponseCompound) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *DiscussionmembershipResponseCompound) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *DiscussionmembershipResponseCompound) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *DiscussionmembershipResponseCompound) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.

### HasFkiUsergroupID

`func (o *DiscussionmembershipResponseCompound) HasFkiUsergroupID() bool`

HasFkiUsergroupID returns a boolean if a field has been set.

### GetFkiModulesectionID

`func (o *DiscussionmembershipResponseCompound) GetFkiModulesectionID() int32`

GetFkiModulesectionID returns the FkiModulesectionID field if non-nil, zero value otherwise.

### GetFkiModulesectionIDOk

`func (o *DiscussionmembershipResponseCompound) GetFkiModulesectionIDOk() (*int32, bool)`

GetFkiModulesectionIDOk returns a tuple with the FkiModulesectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModulesectionID

`func (o *DiscussionmembershipResponseCompound) SetFkiModulesectionID(v int32)`

SetFkiModulesectionID sets FkiModulesectionID field to given value.

### HasFkiModulesectionID

`func (o *DiscussionmembershipResponseCompound) HasFkiModulesectionID() bool`

HasFkiModulesectionID returns a boolean if a field has been set.

### GetSDiscussionmembershipDescription

`func (o *DiscussionmembershipResponseCompound) GetSDiscussionmembershipDescription() string`

GetSDiscussionmembershipDescription returns the SDiscussionmembershipDescription field if non-nil, zero value otherwise.

### GetSDiscussionmembershipDescriptionOk

`func (o *DiscussionmembershipResponseCompound) GetSDiscussionmembershipDescriptionOk() (*string, bool)`

GetSDiscussionmembershipDescriptionOk returns a tuple with the SDiscussionmembershipDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDiscussionmembershipDescription

`func (o *DiscussionmembershipResponseCompound) SetSDiscussionmembershipDescription(v string)`

SetSDiscussionmembershipDescription sets SDiscussionmembershipDescription field to given value.


### GetDtDiscussionmembershipJoined

`func (o *DiscussionmembershipResponseCompound) GetDtDiscussionmembershipJoined() string`

GetDtDiscussionmembershipJoined returns the DtDiscussionmembershipJoined field if non-nil, zero value otherwise.

### GetDtDiscussionmembershipJoinedOk

`func (o *DiscussionmembershipResponseCompound) GetDtDiscussionmembershipJoinedOk() (*string, bool)`

GetDtDiscussionmembershipJoinedOk returns a tuple with the DtDiscussionmembershipJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtDiscussionmembershipJoined

`func (o *DiscussionmembershipResponseCompound) SetDtDiscussionmembershipJoined(v string)`

SetDtDiscussionmembershipJoined sets DtDiscussionmembershipJoined field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


