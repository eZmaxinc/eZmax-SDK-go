# EzsignuserRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignuserID** | Pointer to **int32** | The unique ID of the Ezsignuser | [optional] 
**FkiContactID** | **int32** | The unique ID of the Contact | 
**ObjContact** | [**ContactRequestCompoundV2**](ContactRequestCompoundV2.md) |  | 

## Methods

### NewEzsignuserRequestCompound

`func NewEzsignuserRequestCompound(fkiContactID int32, objContact ContactRequestCompoundV2, ) *EzsignuserRequestCompound`

NewEzsignuserRequestCompound instantiates a new EzsignuserRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignuserRequestCompoundWithDefaults

`func NewEzsignuserRequestCompoundWithDefaults() *EzsignuserRequestCompound`

NewEzsignuserRequestCompoundWithDefaults instantiates a new EzsignuserRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignuserID

`func (o *EzsignuserRequestCompound) GetPkiEzsignuserID() int32`

GetPkiEzsignuserID returns the PkiEzsignuserID field if non-nil, zero value otherwise.

### GetPkiEzsignuserIDOk

`func (o *EzsignuserRequestCompound) GetPkiEzsignuserIDOk() (*int32, bool)`

GetPkiEzsignuserIDOk returns a tuple with the PkiEzsignuserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignuserID

`func (o *EzsignuserRequestCompound) SetPkiEzsignuserID(v int32)`

SetPkiEzsignuserID sets PkiEzsignuserID field to given value.

### HasPkiEzsignuserID

`func (o *EzsignuserRequestCompound) HasPkiEzsignuserID() bool`

HasPkiEzsignuserID returns a boolean if a field has been set.

### GetFkiContactID

`func (o *EzsignuserRequestCompound) GetFkiContactID() int32`

GetFkiContactID returns the FkiContactID field if non-nil, zero value otherwise.

### GetFkiContactIDOk

`func (o *EzsignuserRequestCompound) GetFkiContactIDOk() (*int32, bool)`

GetFkiContactIDOk returns a tuple with the FkiContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiContactID

`func (o *EzsignuserRequestCompound) SetFkiContactID(v int32)`

SetFkiContactID sets FkiContactID field to given value.


### GetObjContact

`func (o *EzsignuserRequestCompound) GetObjContact() ContactRequestCompoundV2`

GetObjContact returns the ObjContact field if non-nil, zero value otherwise.

### GetObjContactOk

`func (o *EzsignuserRequestCompound) GetObjContactOk() (*ContactRequestCompoundV2, bool)`

GetObjContactOk returns a tuple with the ObjContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContact

`func (o *EzsignuserRequestCompound) SetObjContact(v ContactRequestCompoundV2)`

SetObjContact sets ObjContact field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


