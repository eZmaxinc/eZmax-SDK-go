# EzsignuserResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignuserID** | **int32** | The unique ID of the Ezsignuser | 
**FkiContactID** | **int32** | The unique ID of the Contact | 
**ObjContact** | [**ContactResponseCompound**](ContactResponseCompound.md) |  | 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 

## Methods

### NewEzsignuserResponseCompound

`func NewEzsignuserResponseCompound(pkiEzsignuserID int32, fkiContactID int32, objContact ContactResponseCompound, objAudit CommonAudit, ) *EzsignuserResponseCompound`

NewEzsignuserResponseCompound instantiates a new EzsignuserResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignuserResponseCompoundWithDefaults

`func NewEzsignuserResponseCompoundWithDefaults() *EzsignuserResponseCompound`

NewEzsignuserResponseCompoundWithDefaults instantiates a new EzsignuserResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignuserID

`func (o *EzsignuserResponseCompound) GetPkiEzsignuserID() int32`

GetPkiEzsignuserID returns the PkiEzsignuserID field if non-nil, zero value otherwise.

### GetPkiEzsignuserIDOk

`func (o *EzsignuserResponseCompound) GetPkiEzsignuserIDOk() (*int32, bool)`

GetPkiEzsignuserIDOk returns a tuple with the PkiEzsignuserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignuserID

`func (o *EzsignuserResponseCompound) SetPkiEzsignuserID(v int32)`

SetPkiEzsignuserID sets PkiEzsignuserID field to given value.


### GetFkiContactID

`func (o *EzsignuserResponseCompound) GetFkiContactID() int32`

GetFkiContactID returns the FkiContactID field if non-nil, zero value otherwise.

### GetFkiContactIDOk

`func (o *EzsignuserResponseCompound) GetFkiContactIDOk() (*int32, bool)`

GetFkiContactIDOk returns a tuple with the FkiContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiContactID

`func (o *EzsignuserResponseCompound) SetFkiContactID(v int32)`

SetFkiContactID sets FkiContactID field to given value.


### GetObjContact

`func (o *EzsignuserResponseCompound) GetObjContact() ContactResponseCompound`

GetObjContact returns the ObjContact field if non-nil, zero value otherwise.

### GetObjContactOk

`func (o *EzsignuserResponseCompound) GetObjContactOk() (*ContactResponseCompound, bool)`

GetObjContactOk returns a tuple with the ObjContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContact

`func (o *EzsignuserResponseCompound) SetObjContact(v ContactResponseCompound)`

SetObjContact sets ObjContact field to given value.


### GetObjAudit

`func (o *EzsignuserResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsignuserResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsignuserResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


