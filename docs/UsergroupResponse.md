# UsergroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUsergroupID** | **int32** | The unique ID of the Usergroup | 
**ObjUsergroupName** | [**MultilingualUsergroupName**](MultilingualUsergroupName.md) |  | 
**SUsergroupNameX** | Pointer to **string** | The Name of the Usergroup in the language of the requester | [optional] 
**ObjEmail** | Pointer to [**EmailRequest**](EmailRequest.md) |  | [optional] 

## Methods

### NewUsergroupResponse

`func NewUsergroupResponse(pkiUsergroupID int32, objUsergroupName MultilingualUsergroupName, ) *UsergroupResponse`

NewUsergroupResponse instantiates a new UsergroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsergroupResponseWithDefaults

`func NewUsergroupResponseWithDefaults() *UsergroupResponse`

NewUsergroupResponseWithDefaults instantiates a new UsergroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUsergroupID

`func (o *UsergroupResponse) GetPkiUsergroupID() int32`

GetPkiUsergroupID returns the PkiUsergroupID field if non-nil, zero value otherwise.

### GetPkiUsergroupIDOk

`func (o *UsergroupResponse) GetPkiUsergroupIDOk() (*int32, bool)`

GetPkiUsergroupIDOk returns a tuple with the PkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUsergroupID

`func (o *UsergroupResponse) SetPkiUsergroupID(v int32)`

SetPkiUsergroupID sets PkiUsergroupID field to given value.


### GetObjUsergroupName

`func (o *UsergroupResponse) GetObjUsergroupName() MultilingualUsergroupName`

GetObjUsergroupName returns the ObjUsergroupName field if non-nil, zero value otherwise.

### GetObjUsergroupNameOk

`func (o *UsergroupResponse) GetObjUsergroupNameOk() (*MultilingualUsergroupName, bool)`

GetObjUsergroupNameOk returns a tuple with the ObjUsergroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUsergroupName

`func (o *UsergroupResponse) SetObjUsergroupName(v MultilingualUsergroupName)`

SetObjUsergroupName sets ObjUsergroupName field to given value.


### GetSUsergroupNameX

`func (o *UsergroupResponse) GetSUsergroupNameX() string`

GetSUsergroupNameX returns the SUsergroupNameX field if non-nil, zero value otherwise.

### GetSUsergroupNameXOk

`func (o *UsergroupResponse) GetSUsergroupNameXOk() (*string, bool)`

GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUsergroupNameX

`func (o *UsergroupResponse) SetSUsergroupNameX(v string)`

SetSUsergroupNameX sets SUsergroupNameX field to given value.

### HasSUsergroupNameX

`func (o *UsergroupResponse) HasSUsergroupNameX() bool`

HasSUsergroupNameX returns a boolean if a field has been set.

### GetObjEmail

`func (o *UsergroupResponse) GetObjEmail() EmailRequest`

GetObjEmail returns the ObjEmail field if non-nil, zero value otherwise.

### GetObjEmailOk

`func (o *UsergroupResponse) GetObjEmailOk() (*EmailRequest, bool)`

GetObjEmailOk returns a tuple with the ObjEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmail

`func (o *UsergroupResponse) SetObjEmail(v EmailRequest)`

SetObjEmail sets ObjEmail field to given value.

### HasObjEmail

`func (o *UsergroupResponse) HasObjEmail() bool`

HasObjEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


