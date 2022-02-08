# ApikeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiApikeyID** | Pointer to **int32** | The unique ID of the Apikey | [optional] 
**FkiUserID** | **int32** | The unique ID of the User | 
**ObjApikeyDescription** | [**MultilingualApikeyDescription**](MultilingualApikeyDescription.md) |  | 

## Methods

### NewApikeyRequest

`func NewApikeyRequest(fkiUserID int32, objApikeyDescription MultilingualApikeyDescription, ) *ApikeyRequest`

NewApikeyRequest instantiates a new ApikeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApikeyRequestWithDefaults

`func NewApikeyRequestWithDefaults() *ApikeyRequest`

NewApikeyRequestWithDefaults instantiates a new ApikeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiApikeyID

`func (o *ApikeyRequest) GetPkiApikeyID() int32`

GetPkiApikeyID returns the PkiApikeyID field if non-nil, zero value otherwise.

### GetPkiApikeyIDOk

`func (o *ApikeyRequest) GetPkiApikeyIDOk() (*int32, bool)`

GetPkiApikeyIDOk returns a tuple with the PkiApikeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiApikeyID

`func (o *ApikeyRequest) SetPkiApikeyID(v int32)`

SetPkiApikeyID sets PkiApikeyID field to given value.

### HasPkiApikeyID

`func (o *ApikeyRequest) HasPkiApikeyID() bool`

HasPkiApikeyID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *ApikeyRequest) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *ApikeyRequest) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *ApikeyRequest) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetObjApikeyDescription

`func (o *ApikeyRequest) GetObjApikeyDescription() MultilingualApikeyDescription`

GetObjApikeyDescription returns the ObjApikeyDescription field if non-nil, zero value otherwise.

### GetObjApikeyDescriptionOk

`func (o *ApikeyRequest) GetObjApikeyDescriptionOk() (*MultilingualApikeyDescription, bool)`

GetObjApikeyDescriptionOk returns a tuple with the ObjApikeyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjApikeyDescription

`func (o *ApikeyRequest) SetObjApikeyDescription(v MultilingualApikeyDescription)`

SetObjApikeyDescription sets ObjApikeyDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


