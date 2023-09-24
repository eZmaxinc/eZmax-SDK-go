# ApikeyRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiApikeyID** | Pointer to **int32** | The unique ID of the Apikey | [optional] 
**FkiUserID** | **int32** | The unique ID of the User | 
**ObjApikeyDescription** | [**MultilingualApikeyDescription**](MultilingualApikeyDescription.md) |  | 
**BApikeyIsactive** | Pointer to **bool** | Whether the apikey is active or not | [optional] 
**BApikeyIssigned** | Pointer to **bool** | Whether the apikey is signed or not | [optional] 

## Methods

### NewApikeyRequestCompound

`func NewApikeyRequestCompound(fkiUserID int32, objApikeyDescription MultilingualApikeyDescription, ) *ApikeyRequestCompound`

NewApikeyRequestCompound instantiates a new ApikeyRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApikeyRequestCompoundWithDefaults

`func NewApikeyRequestCompoundWithDefaults() *ApikeyRequestCompound`

NewApikeyRequestCompoundWithDefaults instantiates a new ApikeyRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiApikeyID

`func (o *ApikeyRequestCompound) GetPkiApikeyID() int32`

GetPkiApikeyID returns the PkiApikeyID field if non-nil, zero value otherwise.

### GetPkiApikeyIDOk

`func (o *ApikeyRequestCompound) GetPkiApikeyIDOk() (*int32, bool)`

GetPkiApikeyIDOk returns a tuple with the PkiApikeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiApikeyID

`func (o *ApikeyRequestCompound) SetPkiApikeyID(v int32)`

SetPkiApikeyID sets PkiApikeyID field to given value.

### HasPkiApikeyID

`func (o *ApikeyRequestCompound) HasPkiApikeyID() bool`

HasPkiApikeyID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *ApikeyRequestCompound) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *ApikeyRequestCompound) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *ApikeyRequestCompound) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetObjApikeyDescription

`func (o *ApikeyRequestCompound) GetObjApikeyDescription() MultilingualApikeyDescription`

GetObjApikeyDescription returns the ObjApikeyDescription field if non-nil, zero value otherwise.

### GetObjApikeyDescriptionOk

`func (o *ApikeyRequestCompound) GetObjApikeyDescriptionOk() (*MultilingualApikeyDescription, bool)`

GetObjApikeyDescriptionOk returns a tuple with the ObjApikeyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjApikeyDescription

`func (o *ApikeyRequestCompound) SetObjApikeyDescription(v MultilingualApikeyDescription)`

SetObjApikeyDescription sets ObjApikeyDescription field to given value.


### GetBApikeyIsactive

`func (o *ApikeyRequestCompound) GetBApikeyIsactive() bool`

GetBApikeyIsactive returns the BApikeyIsactive field if non-nil, zero value otherwise.

### GetBApikeyIsactiveOk

`func (o *ApikeyRequestCompound) GetBApikeyIsactiveOk() (*bool, bool)`

GetBApikeyIsactiveOk returns a tuple with the BApikeyIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBApikeyIsactive

`func (o *ApikeyRequestCompound) SetBApikeyIsactive(v bool)`

SetBApikeyIsactive sets BApikeyIsactive field to given value.

### HasBApikeyIsactive

`func (o *ApikeyRequestCompound) HasBApikeyIsactive() bool`

HasBApikeyIsactive returns a boolean if a field has been set.

### GetBApikeyIssigned

`func (o *ApikeyRequestCompound) GetBApikeyIssigned() bool`

GetBApikeyIssigned returns the BApikeyIssigned field if non-nil, zero value otherwise.

### GetBApikeyIssignedOk

`func (o *ApikeyRequestCompound) GetBApikeyIssignedOk() (*bool, bool)`

GetBApikeyIssignedOk returns a tuple with the BApikeyIssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBApikeyIssigned

`func (o *ApikeyRequestCompound) SetBApikeyIssigned(v bool)`

SetBApikeyIssigned sets BApikeyIssigned field to given value.

### HasBApikeyIssigned

`func (o *ApikeyRequestCompound) HasBApikeyIssigned() bool`

HasBApikeyIssigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


