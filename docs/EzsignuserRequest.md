# EzsignuserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignuserID** | Pointer to **int32** | The unique ID of the Ezsignuser | [optional] 
**FkiContactID** | **int32** | The unique ID of the Contact | 
**ObjContact** | [**ContactRequestCompoundV2**](ContactRequestCompoundV2.md) |  | 

## Methods

### NewEzsignuserRequest

`func NewEzsignuserRequest(fkiContactID int32, objContact ContactRequestCompoundV2, ) *EzsignuserRequest`

NewEzsignuserRequest instantiates a new EzsignuserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignuserRequestWithDefaults

`func NewEzsignuserRequestWithDefaults() *EzsignuserRequest`

NewEzsignuserRequestWithDefaults instantiates a new EzsignuserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignuserID

`func (o *EzsignuserRequest) GetPkiEzsignuserID() int32`

GetPkiEzsignuserID returns the PkiEzsignuserID field if non-nil, zero value otherwise.

### GetPkiEzsignuserIDOk

`func (o *EzsignuserRequest) GetPkiEzsignuserIDOk() (*int32, bool)`

GetPkiEzsignuserIDOk returns a tuple with the PkiEzsignuserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignuserID

`func (o *EzsignuserRequest) SetPkiEzsignuserID(v int32)`

SetPkiEzsignuserID sets PkiEzsignuserID field to given value.

### HasPkiEzsignuserID

`func (o *EzsignuserRequest) HasPkiEzsignuserID() bool`

HasPkiEzsignuserID returns a boolean if a field has been set.

### GetFkiContactID

`func (o *EzsignuserRequest) GetFkiContactID() int32`

GetFkiContactID returns the FkiContactID field if non-nil, zero value otherwise.

### GetFkiContactIDOk

`func (o *EzsignuserRequest) GetFkiContactIDOk() (*int32, bool)`

GetFkiContactIDOk returns a tuple with the FkiContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiContactID

`func (o *EzsignuserRequest) SetFkiContactID(v int32)`

SetFkiContactID sets FkiContactID field to given value.


### GetObjContact

`func (o *EzsignuserRequest) GetObjContact() ContactRequestCompoundV2`

GetObjContact returns the ObjContact field if non-nil, zero value otherwise.

### GetObjContactOk

`func (o *EzsignuserRequest) GetObjContactOk() (*ContactRequestCompoundV2, bool)`

GetObjContactOk returns a tuple with the ObjContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContact

`func (o *EzsignuserRequest) SetObjContact(v ContactRequestCompoundV2)`

SetObjContact sets ObjContact field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


