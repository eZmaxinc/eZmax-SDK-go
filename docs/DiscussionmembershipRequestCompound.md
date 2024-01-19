# DiscussionmembershipRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDiscussionmembershipID** | Pointer to **int32** | The unique ID of the Discussionmembership | [optional] 
**FkiDiscussionID** | **int32** | The unique ID of the Discussion | 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiUsergroupID** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**FkiModulesectionID** | Pointer to **int32** | The unique ID of the Modulesection | [optional] 
**DtDiscussionmembershipJoined** | **string** | The joined date of the Discussionmembership | 

## Methods

### NewDiscussionmembershipRequestCompound

`func NewDiscussionmembershipRequestCompound(fkiDiscussionID int32, dtDiscussionmembershipJoined string, ) *DiscussionmembershipRequestCompound`

NewDiscussionmembershipRequestCompound instantiates a new DiscussionmembershipRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscussionmembershipRequestCompoundWithDefaults

`func NewDiscussionmembershipRequestCompoundWithDefaults() *DiscussionmembershipRequestCompound`

NewDiscussionmembershipRequestCompoundWithDefaults instantiates a new DiscussionmembershipRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDiscussionmembershipID

`func (o *DiscussionmembershipRequestCompound) GetPkiDiscussionmembershipID() int32`

GetPkiDiscussionmembershipID returns the PkiDiscussionmembershipID field if non-nil, zero value otherwise.

### GetPkiDiscussionmembershipIDOk

`func (o *DiscussionmembershipRequestCompound) GetPkiDiscussionmembershipIDOk() (*int32, bool)`

GetPkiDiscussionmembershipIDOk returns a tuple with the PkiDiscussionmembershipID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDiscussionmembershipID

`func (o *DiscussionmembershipRequestCompound) SetPkiDiscussionmembershipID(v int32)`

SetPkiDiscussionmembershipID sets PkiDiscussionmembershipID field to given value.

### HasPkiDiscussionmembershipID

`func (o *DiscussionmembershipRequestCompound) HasPkiDiscussionmembershipID() bool`

HasPkiDiscussionmembershipID returns a boolean if a field has been set.

### GetFkiDiscussionID

`func (o *DiscussionmembershipRequestCompound) GetFkiDiscussionID() int32`

GetFkiDiscussionID returns the FkiDiscussionID field if non-nil, zero value otherwise.

### GetFkiDiscussionIDOk

`func (o *DiscussionmembershipRequestCompound) GetFkiDiscussionIDOk() (*int32, bool)`

GetFkiDiscussionIDOk returns a tuple with the FkiDiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDiscussionID

`func (o *DiscussionmembershipRequestCompound) SetFkiDiscussionID(v int32)`

SetFkiDiscussionID sets FkiDiscussionID field to given value.


### GetFkiUserID

`func (o *DiscussionmembershipRequestCompound) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *DiscussionmembershipRequestCompound) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *DiscussionmembershipRequestCompound) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *DiscussionmembershipRequestCompound) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *DiscussionmembershipRequestCompound) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *DiscussionmembershipRequestCompound) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *DiscussionmembershipRequestCompound) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.

### HasFkiUsergroupID

`func (o *DiscussionmembershipRequestCompound) HasFkiUsergroupID() bool`

HasFkiUsergroupID returns a boolean if a field has been set.

### GetFkiModulesectionID

`func (o *DiscussionmembershipRequestCompound) GetFkiModulesectionID() int32`

GetFkiModulesectionID returns the FkiModulesectionID field if non-nil, zero value otherwise.

### GetFkiModulesectionIDOk

`func (o *DiscussionmembershipRequestCompound) GetFkiModulesectionIDOk() (*int32, bool)`

GetFkiModulesectionIDOk returns a tuple with the FkiModulesectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModulesectionID

`func (o *DiscussionmembershipRequestCompound) SetFkiModulesectionID(v int32)`

SetFkiModulesectionID sets FkiModulesectionID field to given value.

### HasFkiModulesectionID

`func (o *DiscussionmembershipRequestCompound) HasFkiModulesectionID() bool`

HasFkiModulesectionID returns a boolean if a field has been set.

### GetDtDiscussionmembershipJoined

`func (o *DiscussionmembershipRequestCompound) GetDtDiscussionmembershipJoined() string`

GetDtDiscussionmembershipJoined returns the DtDiscussionmembershipJoined field if non-nil, zero value otherwise.

### GetDtDiscussionmembershipJoinedOk

`func (o *DiscussionmembershipRequestCompound) GetDtDiscussionmembershipJoinedOk() (*string, bool)`

GetDtDiscussionmembershipJoinedOk returns a tuple with the DtDiscussionmembershipJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtDiscussionmembershipJoined

`func (o *DiscussionmembershipRequestCompound) SetDtDiscussionmembershipJoined(v string)`

SetDtDiscussionmembershipJoined sets DtDiscussionmembershipJoined field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


