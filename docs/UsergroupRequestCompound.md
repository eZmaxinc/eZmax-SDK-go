# UsergroupRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUsergroupID** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**ObjEmail** | Pointer to [**EmailRequest**](EmailRequest.md) |  | [optional] 
**ObjUsergroupName** | [**MultilingualUsergroupName**](MultilingualUsergroupName.md) |  | 

## Methods

### NewUsergroupRequestCompound

`func NewUsergroupRequestCompound(objUsergroupName MultilingualUsergroupName, ) *UsergroupRequestCompound`

NewUsergroupRequestCompound instantiates a new UsergroupRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsergroupRequestCompoundWithDefaults

`func NewUsergroupRequestCompoundWithDefaults() *UsergroupRequestCompound`

NewUsergroupRequestCompoundWithDefaults instantiates a new UsergroupRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUsergroupID

`func (o *UsergroupRequestCompound) GetPkiUsergroupID() int32`

GetPkiUsergroupID returns the PkiUsergroupID field if non-nil, zero value otherwise.

### GetPkiUsergroupIDOk

`func (o *UsergroupRequestCompound) GetPkiUsergroupIDOk() (*int32, bool)`

GetPkiUsergroupIDOk returns a tuple with the PkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUsergroupID

`func (o *UsergroupRequestCompound) SetPkiUsergroupID(v int32)`

SetPkiUsergroupID sets PkiUsergroupID field to given value.

### HasPkiUsergroupID

`func (o *UsergroupRequestCompound) HasPkiUsergroupID() bool`

HasPkiUsergroupID returns a boolean if a field has been set.

### GetObjEmail

`func (o *UsergroupRequestCompound) GetObjEmail() EmailRequest`

GetObjEmail returns the ObjEmail field if non-nil, zero value otherwise.

### GetObjEmailOk

`func (o *UsergroupRequestCompound) GetObjEmailOk() (*EmailRequest, bool)`

GetObjEmailOk returns a tuple with the ObjEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmail

`func (o *UsergroupRequestCompound) SetObjEmail(v EmailRequest)`

SetObjEmail sets ObjEmail field to given value.

### HasObjEmail

`func (o *UsergroupRequestCompound) HasObjEmail() bool`

HasObjEmail returns a boolean if a field has been set.

### GetObjUsergroupName

`func (o *UsergroupRequestCompound) GetObjUsergroupName() MultilingualUsergroupName`

GetObjUsergroupName returns the ObjUsergroupName field if non-nil, zero value otherwise.

### GetObjUsergroupNameOk

`func (o *UsergroupRequestCompound) GetObjUsergroupNameOk() (*MultilingualUsergroupName, bool)`

GetObjUsergroupNameOk returns a tuple with the ObjUsergroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUsergroupName

`func (o *UsergroupRequestCompound) SetObjUsergroupName(v MultilingualUsergroupName)`

SetObjUsergroupName sets ObjUsergroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


