# EzsignuserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignuserID** | **int32** | The unique ID of the Ezsignuser | 
**FkiContactID** | **int32** | The unique ID of the Contact | 
**ObjContact** | [**ContactResponseCompound**](ContactResponseCompound.md) |  | 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 

## Methods

### NewEzsignuserResponse

`func NewEzsignuserResponse(pkiEzsignuserID int32, fkiContactID int32, objContact ContactResponseCompound, objAudit CommonAudit, ) *EzsignuserResponse`

NewEzsignuserResponse instantiates a new EzsignuserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignuserResponseWithDefaults

`func NewEzsignuserResponseWithDefaults() *EzsignuserResponse`

NewEzsignuserResponseWithDefaults instantiates a new EzsignuserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignuserID

`func (o *EzsignuserResponse) GetPkiEzsignuserID() int32`

GetPkiEzsignuserID returns the PkiEzsignuserID field if non-nil, zero value otherwise.

### GetPkiEzsignuserIDOk

`func (o *EzsignuserResponse) GetPkiEzsignuserIDOk() (*int32, bool)`

GetPkiEzsignuserIDOk returns a tuple with the PkiEzsignuserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignuserID

`func (o *EzsignuserResponse) SetPkiEzsignuserID(v int32)`

SetPkiEzsignuserID sets PkiEzsignuserID field to given value.


### GetFkiContactID

`func (o *EzsignuserResponse) GetFkiContactID() int32`

GetFkiContactID returns the FkiContactID field if non-nil, zero value otherwise.

### GetFkiContactIDOk

`func (o *EzsignuserResponse) GetFkiContactIDOk() (*int32, bool)`

GetFkiContactIDOk returns a tuple with the FkiContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiContactID

`func (o *EzsignuserResponse) SetFkiContactID(v int32)`

SetFkiContactID sets FkiContactID field to given value.


### GetObjContact

`func (o *EzsignuserResponse) GetObjContact() ContactResponseCompound`

GetObjContact returns the ObjContact field if non-nil, zero value otherwise.

### GetObjContactOk

`func (o *EzsignuserResponse) GetObjContactOk() (*ContactResponseCompound, bool)`

GetObjContactOk returns a tuple with the ObjContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContact

`func (o *EzsignuserResponse) SetObjContact(v ContactResponseCompound)`

SetObjContact sets ObjContact field to given value.


### GetObjAudit

`func (o *EzsignuserResponse) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsignuserResponse) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsignuserResponse) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


