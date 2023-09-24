# ScimGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** | The Name of the Usergroup in the language of the requester | 
**Members** | Pointer to [**[]ScimGroupMember**](ScimGroupMember.md) |  | [optional] 

## Methods

### NewScimGroup

`func NewScimGroup(displayName string, ) *ScimGroup`

NewScimGroup instantiates a new ScimGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimGroupWithDefaults

`func NewScimGroupWithDefaults() *ScimGroup`

NewScimGroupWithDefaults instantiates a new ScimGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScimGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScimGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *ScimGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScimGroup) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ScimGroup) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMembers

`func (o *ScimGroup) GetMembers() []ScimGroupMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ScimGroup) GetMembersOk() (*[]ScimGroupMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ScimGroup) SetMembers(v []ScimGroupMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ScimGroup) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


