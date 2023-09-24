# UsergroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUsergroupID** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**ObjUsergroupName** | [**MultilingualUsergroupName**](MultilingualUsergroupName.md) |  | 

## Methods

### NewUsergroupRequest

`func NewUsergroupRequest(objUsergroupName MultilingualUsergroupName, ) *UsergroupRequest`

NewUsergroupRequest instantiates a new UsergroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsergroupRequestWithDefaults

`func NewUsergroupRequestWithDefaults() *UsergroupRequest`

NewUsergroupRequestWithDefaults instantiates a new UsergroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUsergroupID

`func (o *UsergroupRequest) GetPkiUsergroupID() int32`

GetPkiUsergroupID returns the PkiUsergroupID field if non-nil, zero value otherwise.

### GetPkiUsergroupIDOk

`func (o *UsergroupRequest) GetPkiUsergroupIDOk() (*int32, bool)`

GetPkiUsergroupIDOk returns a tuple with the PkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUsergroupID

`func (o *UsergroupRequest) SetPkiUsergroupID(v int32)`

SetPkiUsergroupID sets PkiUsergroupID field to given value.

### HasPkiUsergroupID

`func (o *UsergroupRequest) HasPkiUsergroupID() bool`

HasPkiUsergroupID returns a boolean if a field has been set.

### GetObjUsergroupName

`func (o *UsergroupRequest) GetObjUsergroupName() MultilingualUsergroupName`

GetObjUsergroupName returns the ObjUsergroupName field if non-nil, zero value otherwise.

### GetObjUsergroupNameOk

`func (o *UsergroupRequest) GetObjUsergroupNameOk() (*MultilingualUsergroupName, bool)`

GetObjUsergroupNameOk returns a tuple with the ObjUsergroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUsergroupName

`func (o *UsergroupRequest) SetObjUsergroupName(v MultilingualUsergroupName)`

SetObjUsergroupName sets ObjUsergroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


