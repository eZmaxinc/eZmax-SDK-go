# SubnetRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSubnetID** | Pointer to **int32** | The unique ID of the Subnet | [optional] 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiApikeyID** | Pointer to **int32** | The unique ID of the Apikey | [optional] 
**ObjSubnetDescription** | [**MultilingualSubnetDescription**](MultilingualSubnetDescription.md) |  | 
**ISubnetNetwork** | **int64** | The network of the Subnet in integer form. For example 8.8.8.0 would be 134744064 | 
**ISubnetMask** | **int64** | The mask of the Subnet  in integer form. For example 255.255.255.0 would be 4294967040 | 

## Methods

### NewSubnetRequestCompound

`func NewSubnetRequestCompound(objSubnetDescription MultilingualSubnetDescription, iSubnetNetwork int64, iSubnetMask int64, ) *SubnetRequestCompound`

NewSubnetRequestCompound instantiates a new SubnetRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetRequestCompoundWithDefaults

`func NewSubnetRequestCompoundWithDefaults() *SubnetRequestCompound`

NewSubnetRequestCompoundWithDefaults instantiates a new SubnetRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSubnetID

`func (o *SubnetRequestCompound) GetPkiSubnetID() int32`

GetPkiSubnetID returns the PkiSubnetID field if non-nil, zero value otherwise.

### GetPkiSubnetIDOk

`func (o *SubnetRequestCompound) GetPkiSubnetIDOk() (*int32, bool)`

GetPkiSubnetIDOk returns a tuple with the PkiSubnetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSubnetID

`func (o *SubnetRequestCompound) SetPkiSubnetID(v int32)`

SetPkiSubnetID sets PkiSubnetID field to given value.

### HasPkiSubnetID

`func (o *SubnetRequestCompound) HasPkiSubnetID() bool`

HasPkiSubnetID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *SubnetRequestCompound) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *SubnetRequestCompound) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *SubnetRequestCompound) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *SubnetRequestCompound) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiApikeyID

`func (o *SubnetRequestCompound) GetFkiApikeyID() int32`

GetFkiApikeyID returns the FkiApikeyID field if non-nil, zero value otherwise.

### GetFkiApikeyIDOk

`func (o *SubnetRequestCompound) GetFkiApikeyIDOk() (*int32, bool)`

GetFkiApikeyIDOk returns a tuple with the FkiApikeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiApikeyID

`func (o *SubnetRequestCompound) SetFkiApikeyID(v int32)`

SetFkiApikeyID sets FkiApikeyID field to given value.

### HasFkiApikeyID

`func (o *SubnetRequestCompound) HasFkiApikeyID() bool`

HasFkiApikeyID returns a boolean if a field has been set.

### GetObjSubnetDescription

`func (o *SubnetRequestCompound) GetObjSubnetDescription() MultilingualSubnetDescription`

GetObjSubnetDescription returns the ObjSubnetDescription field if non-nil, zero value otherwise.

### GetObjSubnetDescriptionOk

`func (o *SubnetRequestCompound) GetObjSubnetDescriptionOk() (*MultilingualSubnetDescription, bool)`

GetObjSubnetDescriptionOk returns a tuple with the ObjSubnetDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjSubnetDescription

`func (o *SubnetRequestCompound) SetObjSubnetDescription(v MultilingualSubnetDescription)`

SetObjSubnetDescription sets ObjSubnetDescription field to given value.


### GetISubnetNetwork

`func (o *SubnetRequestCompound) GetISubnetNetwork() int64`

GetISubnetNetwork returns the ISubnetNetwork field if non-nil, zero value otherwise.

### GetISubnetNetworkOk

`func (o *SubnetRequestCompound) GetISubnetNetworkOk() (*int64, bool)`

GetISubnetNetworkOk returns a tuple with the ISubnetNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISubnetNetwork

`func (o *SubnetRequestCompound) SetISubnetNetwork(v int64)`

SetISubnetNetwork sets ISubnetNetwork field to given value.


### GetISubnetMask

`func (o *SubnetRequestCompound) GetISubnetMask() int64`

GetISubnetMask returns the ISubnetMask field if non-nil, zero value otherwise.

### GetISubnetMaskOk

`func (o *SubnetRequestCompound) GetISubnetMaskOk() (*int64, bool)`

GetISubnetMaskOk returns a tuple with the ISubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISubnetMask

`func (o *SubnetRequestCompound) SetISubnetMask(v int64)`

SetISubnetMask sets ISubnetMask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


