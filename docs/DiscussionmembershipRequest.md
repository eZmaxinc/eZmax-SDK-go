# DiscussionmembershipRequest

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

### NewDiscussionmembershipRequest

`func NewDiscussionmembershipRequest(fkiDiscussionID int32, dtDiscussionmembershipJoined string, ) *DiscussionmembershipRequest`

NewDiscussionmembershipRequest instantiates a new DiscussionmembershipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscussionmembershipRequestWithDefaults

`func NewDiscussionmembershipRequestWithDefaults() *DiscussionmembershipRequest`

NewDiscussionmembershipRequestWithDefaults instantiates a new DiscussionmembershipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDiscussionmembershipID

`func (o *DiscussionmembershipRequest) GetPkiDiscussionmembershipID() int32`

GetPkiDiscussionmembershipID returns the PkiDiscussionmembershipID field if non-nil, zero value otherwise.

### GetPkiDiscussionmembershipIDOk

`func (o *DiscussionmembershipRequest) GetPkiDiscussionmembershipIDOk() (*int32, bool)`

GetPkiDiscussionmembershipIDOk returns a tuple with the PkiDiscussionmembershipID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDiscussionmembershipID

`func (o *DiscussionmembershipRequest) SetPkiDiscussionmembershipID(v int32)`

SetPkiDiscussionmembershipID sets PkiDiscussionmembershipID field to given value.

### HasPkiDiscussionmembershipID

`func (o *DiscussionmembershipRequest) HasPkiDiscussionmembershipID() bool`

HasPkiDiscussionmembershipID returns a boolean if a field has been set.

### GetFkiDiscussionID

`func (o *DiscussionmembershipRequest) GetFkiDiscussionID() int32`

GetFkiDiscussionID returns the FkiDiscussionID field if non-nil, zero value otherwise.

### GetFkiDiscussionIDOk

`func (o *DiscussionmembershipRequest) GetFkiDiscussionIDOk() (*int32, bool)`

GetFkiDiscussionIDOk returns a tuple with the FkiDiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDiscussionID

`func (o *DiscussionmembershipRequest) SetFkiDiscussionID(v int32)`

SetFkiDiscussionID sets FkiDiscussionID field to given value.


### GetFkiUserID

`func (o *DiscussionmembershipRequest) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *DiscussionmembershipRequest) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *DiscussionmembershipRequest) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *DiscussionmembershipRequest) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *DiscussionmembershipRequest) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *DiscussionmembershipRequest) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *DiscussionmembershipRequest) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.

### HasFkiUsergroupID

`func (o *DiscussionmembershipRequest) HasFkiUsergroupID() bool`

HasFkiUsergroupID returns a boolean if a field has been set.

### GetFkiModulesectionID

`func (o *DiscussionmembershipRequest) GetFkiModulesectionID() int32`

GetFkiModulesectionID returns the FkiModulesectionID field if non-nil, zero value otherwise.

### GetFkiModulesectionIDOk

`func (o *DiscussionmembershipRequest) GetFkiModulesectionIDOk() (*int32, bool)`

GetFkiModulesectionIDOk returns a tuple with the FkiModulesectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModulesectionID

`func (o *DiscussionmembershipRequest) SetFkiModulesectionID(v int32)`

SetFkiModulesectionID sets FkiModulesectionID field to given value.

### HasFkiModulesectionID

`func (o *DiscussionmembershipRequest) HasFkiModulesectionID() bool`

HasFkiModulesectionID returns a boolean if a field has been set.

### GetDtDiscussionmembershipJoined

`func (o *DiscussionmembershipRequest) GetDtDiscussionmembershipJoined() string`

GetDtDiscussionmembershipJoined returns the DtDiscussionmembershipJoined field if non-nil, zero value otherwise.

### GetDtDiscussionmembershipJoinedOk

`func (o *DiscussionmembershipRequest) GetDtDiscussionmembershipJoinedOk() (*string, bool)`

GetDtDiscussionmembershipJoinedOk returns a tuple with the DtDiscussionmembershipJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtDiscussionmembershipJoined

`func (o *DiscussionmembershipRequest) SetDtDiscussionmembershipJoined(v string)`

SetDtDiscussionmembershipJoined sets DtDiscussionmembershipJoined field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


