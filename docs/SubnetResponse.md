# SubnetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSubnetID** | **int32** | The unique ID of the Subnet | 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiApikeyID** | Pointer to **int32** | The unique ID of the Apikey | [optional] 
**ObjSubnetDescription** | [**MultilingualSubnetDescription**](MultilingualSubnetDescription.md) |  | 
**ISubnetNetwork** | **int64** | The network of the Subnet in integer form. For example 8.8.8.0 would be 134744064 | 
**ISubnetMask** | **int64** | The mask of the Subnet  in integer form. For example 255.255.255.0 would be 4294967040 | 

## Methods

### NewSubnetResponse

`func NewSubnetResponse(pkiSubnetID int32, objSubnetDescription MultilingualSubnetDescription, iSubnetNetwork int64, iSubnetMask int64, ) *SubnetResponse`

NewSubnetResponse instantiates a new SubnetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetResponseWithDefaults

`func NewSubnetResponseWithDefaults() *SubnetResponse`

NewSubnetResponseWithDefaults instantiates a new SubnetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSubnetID

`func (o *SubnetResponse) GetPkiSubnetID() int32`

GetPkiSubnetID returns the PkiSubnetID field if non-nil, zero value otherwise.

### GetPkiSubnetIDOk

`func (o *SubnetResponse) GetPkiSubnetIDOk() (*int32, bool)`

GetPkiSubnetIDOk returns a tuple with the PkiSubnetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSubnetID

`func (o *SubnetResponse) SetPkiSubnetID(v int32)`

SetPkiSubnetID sets PkiSubnetID field to given value.


### GetFkiUserID

`func (o *SubnetResponse) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *SubnetResponse) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *SubnetResponse) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *SubnetResponse) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiApikeyID

`func (o *SubnetResponse) GetFkiApikeyID() int32`

GetFkiApikeyID returns the FkiApikeyID field if non-nil, zero value otherwise.

### GetFkiApikeyIDOk

`func (o *SubnetResponse) GetFkiApikeyIDOk() (*int32, bool)`

GetFkiApikeyIDOk returns a tuple with the FkiApikeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiApikeyID

`func (o *SubnetResponse) SetFkiApikeyID(v int32)`

SetFkiApikeyID sets FkiApikeyID field to given value.

### HasFkiApikeyID

`func (o *SubnetResponse) HasFkiApikeyID() bool`

HasFkiApikeyID returns a boolean if a field has been set.

### GetObjSubnetDescription

`func (o *SubnetResponse) GetObjSubnetDescription() MultilingualSubnetDescription`

GetObjSubnetDescription returns the ObjSubnetDescription field if non-nil, zero value otherwise.

### GetObjSubnetDescriptionOk

`func (o *SubnetResponse) GetObjSubnetDescriptionOk() (*MultilingualSubnetDescription, bool)`

GetObjSubnetDescriptionOk returns a tuple with the ObjSubnetDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjSubnetDescription

`func (o *SubnetResponse) SetObjSubnetDescription(v MultilingualSubnetDescription)`

SetObjSubnetDescription sets ObjSubnetDescription field to given value.


### GetISubnetNetwork

`func (o *SubnetResponse) GetISubnetNetwork() int64`

GetISubnetNetwork returns the ISubnetNetwork field if non-nil, zero value otherwise.

### GetISubnetNetworkOk

`func (o *SubnetResponse) GetISubnetNetworkOk() (*int64, bool)`

GetISubnetNetworkOk returns a tuple with the ISubnetNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISubnetNetwork

`func (o *SubnetResponse) SetISubnetNetwork(v int64)`

SetISubnetNetwork sets ISubnetNetwork field to given value.


### GetISubnetMask

`func (o *SubnetResponse) GetISubnetMask() int64`

GetISubnetMask returns the ISubnetMask field if non-nil, zero value otherwise.

### GetISubnetMaskOk

`func (o *SubnetResponse) GetISubnetMaskOk() (*int64, bool)`

GetISubnetMaskOk returns a tuple with the ISubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISubnetMask

`func (o *SubnetResponse) SetISubnetMask(v int64)`

SetISubnetMask sets ISubnetMask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


