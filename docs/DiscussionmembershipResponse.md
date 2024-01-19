# DiscussionmembershipResponse

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

### NewDiscussionmembershipResponse

`func NewDiscussionmembershipResponse(pkiDiscussionmembershipID int32, fkiDiscussionID int32, sDiscussionmembershipDescription string, dtDiscussionmembershipJoined string, ) *DiscussionmembershipResponse`

NewDiscussionmembershipResponse instantiates a new DiscussionmembershipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscussionmembershipResponseWithDefaults

`func NewDiscussionmembershipResponseWithDefaults() *DiscussionmembershipResponse`

NewDiscussionmembershipResponseWithDefaults instantiates a new DiscussionmembershipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDiscussionmembershipID

`func (o *DiscussionmembershipResponse) GetPkiDiscussionmembershipID() int32`

GetPkiDiscussionmembershipID returns the PkiDiscussionmembershipID field if non-nil, zero value otherwise.

### GetPkiDiscussionmembershipIDOk

`func (o *DiscussionmembershipResponse) GetPkiDiscussionmembershipIDOk() (*int32, bool)`

GetPkiDiscussionmembershipIDOk returns a tuple with the PkiDiscussionmembershipID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDiscussionmembershipID

`func (o *DiscussionmembershipResponse) SetPkiDiscussionmembershipID(v int32)`

SetPkiDiscussionmembershipID sets PkiDiscussionmembershipID field to given value.


### GetFkiDiscussionID

`func (o *DiscussionmembershipResponse) GetFkiDiscussionID() int32`

GetFkiDiscussionID returns the FkiDiscussionID field if non-nil, zero value otherwise.

### GetFkiDiscussionIDOk

`func (o *DiscussionmembershipResponse) GetFkiDiscussionIDOk() (*int32, bool)`

GetFkiDiscussionIDOk returns a tuple with the FkiDiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDiscussionID

`func (o *DiscussionmembershipResponse) SetFkiDiscussionID(v int32)`

SetFkiDiscussionID sets FkiDiscussionID field to given value.


### GetFkiUserID

`func (o *DiscussionmembershipResponse) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *DiscussionmembershipResponse) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *DiscussionmembershipResponse) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *DiscussionmembershipResponse) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *DiscussionmembershipResponse) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *DiscussionmembershipResponse) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *DiscussionmembershipResponse) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.

### HasFkiUsergroupID

`func (o *DiscussionmembershipResponse) HasFkiUsergroupID() bool`

HasFkiUsergroupID returns a boolean if a field has been set.

### GetFkiModulesectionID

`func (o *DiscussionmembershipResponse) GetFkiModulesectionID() int32`

GetFkiModulesectionID returns the FkiModulesectionID field if non-nil, zero value otherwise.

### GetFkiModulesectionIDOk

`func (o *DiscussionmembershipResponse) GetFkiModulesectionIDOk() (*int32, bool)`

GetFkiModulesectionIDOk returns a tuple with the FkiModulesectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModulesectionID

`func (o *DiscussionmembershipResponse) SetFkiModulesectionID(v int32)`

SetFkiModulesectionID sets FkiModulesectionID field to given value.

### HasFkiModulesectionID

`func (o *DiscussionmembershipResponse) HasFkiModulesectionID() bool`

HasFkiModulesectionID returns a boolean if a field has been set.

### GetSDiscussionmembershipDescription

`func (o *DiscussionmembershipResponse) GetSDiscussionmembershipDescription() string`

GetSDiscussionmembershipDescription returns the SDiscussionmembershipDescription field if non-nil, zero value otherwise.

### GetSDiscussionmembershipDescriptionOk

`func (o *DiscussionmembershipResponse) GetSDiscussionmembershipDescriptionOk() (*string, bool)`

GetSDiscussionmembershipDescriptionOk returns a tuple with the SDiscussionmembershipDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDiscussionmembershipDescription

`func (o *DiscussionmembershipResponse) SetSDiscussionmembershipDescription(v string)`

SetSDiscussionmembershipDescription sets SDiscussionmembershipDescription field to given value.


### GetDtDiscussionmembershipJoined

`func (o *DiscussionmembershipResponse) GetDtDiscussionmembershipJoined() string`

GetDtDiscussionmembershipJoined returns the DtDiscussionmembershipJoined field if non-nil, zero value otherwise.

### GetDtDiscussionmembershipJoinedOk

`func (o *DiscussionmembershipResponse) GetDtDiscussionmembershipJoinedOk() (*string, bool)`

GetDtDiscussionmembershipJoinedOk returns a tuple with the DtDiscussionmembershipJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtDiscussionmembershipJoined

`func (o *DiscussionmembershipResponse) SetDtDiscussionmembershipJoined(v string)`

SetDtDiscussionmembershipJoined sets DtDiscussionmembershipJoined field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


