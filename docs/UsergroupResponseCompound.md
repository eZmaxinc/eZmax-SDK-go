# UsergroupResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUsergroupID** | **int32** | The unique ID of the Usergroup | 
**ObjUsergroupName** | [**MultilingualUsergroupName**](MultilingualUsergroupName.md) |  | 
**SUsergroupNameX** | Pointer to **string** | The Name of the Usergroup in the language of the requester | [optional] 
**ObjEmail** | Pointer to [**EmailRequest**](EmailRequest.md) |  | [optional] 

## Methods

### NewUsergroupResponseCompound

`func NewUsergroupResponseCompound(pkiUsergroupID int32, objUsergroupName MultilingualUsergroupName, ) *UsergroupResponseCompound`

NewUsergroupResponseCompound instantiates a new UsergroupResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsergroupResponseCompoundWithDefaults

`func NewUsergroupResponseCompoundWithDefaults() *UsergroupResponseCompound`

NewUsergroupResponseCompoundWithDefaults instantiates a new UsergroupResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUsergroupID

`func (o *UsergroupResponseCompound) GetPkiUsergroupID() int32`

GetPkiUsergroupID returns the PkiUsergroupID field if non-nil, zero value otherwise.

### GetPkiUsergroupIDOk

`func (o *UsergroupResponseCompound) GetPkiUsergroupIDOk() (*int32, bool)`

GetPkiUsergroupIDOk returns a tuple with the PkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUsergroupID

`func (o *UsergroupResponseCompound) SetPkiUsergroupID(v int32)`

SetPkiUsergroupID sets PkiUsergroupID field to given value.


### GetObjUsergroupName

`func (o *UsergroupResponseCompound) GetObjUsergroupName() MultilingualUsergroupName`

GetObjUsergroupName returns the ObjUsergroupName field if non-nil, zero value otherwise.

### GetObjUsergroupNameOk

`func (o *UsergroupResponseCompound) GetObjUsergroupNameOk() (*MultilingualUsergroupName, bool)`

GetObjUsergroupNameOk returns a tuple with the ObjUsergroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUsergroupName

`func (o *UsergroupResponseCompound) SetObjUsergroupName(v MultilingualUsergroupName)`

SetObjUsergroupName sets ObjUsergroupName field to given value.


### GetSUsergroupNameX

`func (o *UsergroupResponseCompound) GetSUsergroupNameX() string`

GetSUsergroupNameX returns the SUsergroupNameX field if non-nil, zero value otherwise.

### GetSUsergroupNameXOk

`func (o *UsergroupResponseCompound) GetSUsergroupNameXOk() (*string, bool)`

GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUsergroupNameX

`func (o *UsergroupResponseCompound) SetSUsergroupNameX(v string)`

SetSUsergroupNameX sets SUsergroupNameX field to given value.

### HasSUsergroupNameX

`func (o *UsergroupResponseCompound) HasSUsergroupNameX() bool`

HasSUsergroupNameX returns a boolean if a field has been set.

### GetObjEmail

`func (o *UsergroupResponseCompound) GetObjEmail() EmailRequest`

GetObjEmail returns the ObjEmail field if non-nil, zero value otherwise.

### GetObjEmailOk

`func (o *UsergroupResponseCompound) GetObjEmailOk() (*EmailRequest, bool)`

GetObjEmailOk returns a tuple with the ObjEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmail

`func (o *UsergroupResponseCompound) SetObjEmail(v EmailRequest)`

SetObjEmail sets ObjEmail field to given value.

### HasObjEmail

`func (o *UsergroupResponseCompound) HasObjEmail() bool`

HasObjEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


