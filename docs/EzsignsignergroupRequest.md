# EzsignsignergroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignsignergroupID** | Pointer to **int32** | The unique ID of the Ezsignsignergroup | [optional] 
**FkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 
**ObjEzsignsignergroupDescription** | [**MultilingualEzsignsignergroupDescription**](MultilingualEzsignsignergroupDescription.md) |  | 

## Methods

### NewEzsignsignergroupRequest

`func NewEzsignsignergroupRequest(fkiEzsignfolderID int32, objEzsignsignergroupDescription MultilingualEzsignsignergroupDescription, ) *EzsignsignergroupRequest`

NewEzsignsignergroupRequest instantiates a new EzsignsignergroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignergroupRequestWithDefaults

`func NewEzsignsignergroupRequestWithDefaults() *EzsignsignergroupRequest`

NewEzsignsignergroupRequestWithDefaults instantiates a new EzsignsignergroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignsignergroupID

`func (o *EzsignsignergroupRequest) GetPkiEzsignsignergroupID() int32`

GetPkiEzsignsignergroupID returns the PkiEzsignsignergroupID field if non-nil, zero value otherwise.

### GetPkiEzsignsignergroupIDOk

`func (o *EzsignsignergroupRequest) GetPkiEzsignsignergroupIDOk() (*int32, bool)`

GetPkiEzsignsignergroupIDOk returns a tuple with the PkiEzsignsignergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignsignergroupID

`func (o *EzsignsignergroupRequest) SetPkiEzsignsignergroupID(v int32)`

SetPkiEzsignsignergroupID sets PkiEzsignsignergroupID field to given value.

### HasPkiEzsignsignergroupID

`func (o *EzsignsignergroupRequest) HasPkiEzsignsignergroupID() bool`

HasPkiEzsignsignergroupID returns a boolean if a field has been set.

### GetFkiEzsignfolderID

`func (o *EzsignsignergroupRequest) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *EzsignsignergroupRequest) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *EzsignsignergroupRequest) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.


### GetObjEzsignsignergroupDescription

`func (o *EzsignsignergroupRequest) GetObjEzsignsignergroupDescription() MultilingualEzsignsignergroupDescription`

GetObjEzsignsignergroupDescription returns the ObjEzsignsignergroupDescription field if non-nil, zero value otherwise.

### GetObjEzsignsignergroupDescriptionOk

`func (o *EzsignsignergroupRequest) GetObjEzsignsignergroupDescriptionOk() (*MultilingualEzsignsignergroupDescription, bool)`

GetObjEzsignsignergroupDescriptionOk returns a tuple with the ObjEzsignsignergroupDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignsignergroupDescription

`func (o *EzsignsignergroupRequest) SetObjEzsignsignergroupDescription(v MultilingualEzsignsignergroupDescription)`

SetObjEzsignsignergroupDescription sets ObjEzsignsignergroupDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


