# EzsignbulksendtransmissionResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignbulksendtransmissionID** | **int32** | The unique ID of the Ezsignbulksendtransmission | 
**FkiEzsignbulksendID** | **int32** | The unique ID of the Ezsignbulksend | 
**SEzsignbulksendtransmissionDescription** | **string** | The description of the Ezsignbulksendtransmission | 
**IEzsignbulksendtransmissionErrors** | **int32** | The number of errors during the Ezsignbulksendtransmission | 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 
**AObjEzsignfoldertransmission** | [**[]CustomEzsignfoldertransmissionResponse**](CustomEzsignfoldertransmissionResponse.md) |  | 

## Methods

### NewEzsignbulksendtransmissionResponseCompound

`func NewEzsignbulksendtransmissionResponseCompound(pkiEzsignbulksendtransmissionID int32, fkiEzsignbulksendID int32, sEzsignbulksendtransmissionDescription string, iEzsignbulksendtransmissionErrors int32, objAudit CommonAudit, aObjEzsignfoldertransmission []CustomEzsignfoldertransmissionResponse, ) *EzsignbulksendtransmissionResponseCompound`

NewEzsignbulksendtransmissionResponseCompound instantiates a new EzsignbulksendtransmissionResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignbulksendtransmissionResponseCompoundWithDefaults

`func NewEzsignbulksendtransmissionResponseCompoundWithDefaults() *EzsignbulksendtransmissionResponseCompound`

NewEzsignbulksendtransmissionResponseCompoundWithDefaults instantiates a new EzsignbulksendtransmissionResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignbulksendtransmissionID

`func (o *EzsignbulksendtransmissionResponseCompound) GetPkiEzsignbulksendtransmissionID() int32`

GetPkiEzsignbulksendtransmissionID returns the PkiEzsignbulksendtransmissionID field if non-nil, zero value otherwise.

### GetPkiEzsignbulksendtransmissionIDOk

`func (o *EzsignbulksendtransmissionResponseCompound) GetPkiEzsignbulksendtransmissionIDOk() (*int32, bool)`

GetPkiEzsignbulksendtransmissionIDOk returns a tuple with the PkiEzsignbulksendtransmissionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignbulksendtransmissionID

`func (o *EzsignbulksendtransmissionResponseCompound) SetPkiEzsignbulksendtransmissionID(v int32)`

SetPkiEzsignbulksendtransmissionID sets PkiEzsignbulksendtransmissionID field to given value.


### GetFkiEzsignbulksendID

`func (o *EzsignbulksendtransmissionResponseCompound) GetFkiEzsignbulksendID() int32`

GetFkiEzsignbulksendID returns the FkiEzsignbulksendID field if non-nil, zero value otherwise.

### GetFkiEzsignbulksendIDOk

`func (o *EzsignbulksendtransmissionResponseCompound) GetFkiEzsignbulksendIDOk() (*int32, bool)`

GetFkiEzsignbulksendIDOk returns a tuple with the FkiEzsignbulksendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignbulksendID

`func (o *EzsignbulksendtransmissionResponseCompound) SetFkiEzsignbulksendID(v int32)`

SetFkiEzsignbulksendID sets FkiEzsignbulksendID field to given value.


### GetSEzsignbulksendtransmissionDescription

`func (o *EzsignbulksendtransmissionResponseCompound) GetSEzsignbulksendtransmissionDescription() string`

GetSEzsignbulksendtransmissionDescription returns the SEzsignbulksendtransmissionDescription field if non-nil, zero value otherwise.

### GetSEzsignbulksendtransmissionDescriptionOk

`func (o *EzsignbulksendtransmissionResponseCompound) GetSEzsignbulksendtransmissionDescriptionOk() (*string, bool)`

GetSEzsignbulksendtransmissionDescriptionOk returns a tuple with the SEzsignbulksendtransmissionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignbulksendtransmissionDescription

`func (o *EzsignbulksendtransmissionResponseCompound) SetSEzsignbulksendtransmissionDescription(v string)`

SetSEzsignbulksendtransmissionDescription sets SEzsignbulksendtransmissionDescription field to given value.


### GetIEzsignbulksendtransmissionErrors

`func (o *EzsignbulksendtransmissionResponseCompound) GetIEzsignbulksendtransmissionErrors() int32`

GetIEzsignbulksendtransmissionErrors returns the IEzsignbulksendtransmissionErrors field if non-nil, zero value otherwise.

### GetIEzsignbulksendtransmissionErrorsOk

`func (o *EzsignbulksendtransmissionResponseCompound) GetIEzsignbulksendtransmissionErrorsOk() (*int32, bool)`

GetIEzsignbulksendtransmissionErrorsOk returns a tuple with the IEzsignbulksendtransmissionErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignbulksendtransmissionErrors

`func (o *EzsignbulksendtransmissionResponseCompound) SetIEzsignbulksendtransmissionErrors(v int32)`

SetIEzsignbulksendtransmissionErrors sets IEzsignbulksendtransmissionErrors field to given value.


### GetObjAudit

`func (o *EzsignbulksendtransmissionResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsignbulksendtransmissionResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsignbulksendtransmissionResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.


### GetAObjEzsignfoldertransmission

`func (o *EzsignbulksendtransmissionResponseCompound) GetAObjEzsignfoldertransmission() []CustomEzsignfoldertransmissionResponse`

GetAObjEzsignfoldertransmission returns the AObjEzsignfoldertransmission field if non-nil, zero value otherwise.

### GetAObjEzsignfoldertransmissionOk

`func (o *EzsignbulksendtransmissionResponseCompound) GetAObjEzsignfoldertransmissionOk() (*[]CustomEzsignfoldertransmissionResponse, bool)`

GetAObjEzsignfoldertransmissionOk returns a tuple with the AObjEzsignfoldertransmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignfoldertransmission

`func (o *EzsignbulksendtransmissionResponseCompound) SetAObjEzsignfoldertransmission(v []CustomEzsignfoldertransmissionResponse)`

SetAObjEzsignfoldertransmission sets AObjEzsignfoldertransmission field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


