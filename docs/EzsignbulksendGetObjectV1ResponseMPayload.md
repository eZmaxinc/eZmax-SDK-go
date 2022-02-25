# EzsignbulksendGetObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignbulksendID** | **int32** | The unique ID of the Ezsignbulksend | 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SEzsignbulksendDescription** | **string** | The description of the Ezsignbulksend | 
**TEzsignbulksendNote** | **string** | Note about the Ezsignbulksend | 
**BEzsignbulksendIsactive** | **bool** | Whether the Ezsignbulksend is active or not | 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 

## Methods

### NewEzsignbulksendGetObjectV1ResponseMPayload

`func NewEzsignbulksendGetObjectV1ResponseMPayload(pkiEzsignbulksendID int32, fkiEzsignfoldertypeID int32, fkiLanguageID int32, sEzsignbulksendDescription string, tEzsignbulksendNote string, bEzsignbulksendIsactive bool, objAudit CommonAudit, ) *EzsignbulksendGetObjectV1ResponseMPayload`

NewEzsignbulksendGetObjectV1ResponseMPayload instantiates a new EzsignbulksendGetObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignbulksendGetObjectV1ResponseMPayloadWithDefaults

`func NewEzsignbulksendGetObjectV1ResponseMPayloadWithDefaults() *EzsignbulksendGetObjectV1ResponseMPayload`

NewEzsignbulksendGetObjectV1ResponseMPayloadWithDefaults instantiates a new EzsignbulksendGetObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignbulksendID

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) GetPkiEzsignbulksendID() int32`

GetPkiEzsignbulksendID returns the PkiEzsignbulksendID field if non-nil, zero value otherwise.

### GetPkiEzsignbulksendIDOk

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) GetPkiEzsignbulksendIDOk() (*int32, bool)`

GetPkiEzsignbulksendIDOk returns a tuple with the PkiEzsignbulksendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignbulksendID

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) SetPkiEzsignbulksendID(v int32)`

SetPkiEzsignbulksendID sets PkiEzsignbulksendID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetFkiLanguageID

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSEzsignbulksendDescription

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) GetSEzsignbulksendDescription() string`

GetSEzsignbulksendDescription returns the SEzsignbulksendDescription field if non-nil, zero value otherwise.

### GetSEzsignbulksendDescriptionOk

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) GetSEzsignbulksendDescriptionOk() (*string, bool)`

GetSEzsignbulksendDescriptionOk returns a tuple with the SEzsignbulksendDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignbulksendDescription

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) SetSEzsignbulksendDescription(v string)`

SetSEzsignbulksendDescription sets SEzsignbulksendDescription field to given value.


### GetTEzsignbulksendNote

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) GetTEzsignbulksendNote() string`

GetTEzsignbulksendNote returns the TEzsignbulksendNote field if non-nil, zero value otherwise.

### GetTEzsignbulksendNoteOk

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) GetTEzsignbulksendNoteOk() (*string, bool)`

GetTEzsignbulksendNoteOk returns a tuple with the TEzsignbulksendNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignbulksendNote

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) SetTEzsignbulksendNote(v string)`

SetTEzsignbulksendNote sets TEzsignbulksendNote field to given value.


### GetBEzsignbulksendIsactive

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) GetBEzsignbulksendIsactive() bool`

GetBEzsignbulksendIsactive returns the BEzsignbulksendIsactive field if non-nil, zero value otherwise.

### GetBEzsignbulksendIsactiveOk

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) GetBEzsignbulksendIsactiveOk() (*bool, bool)`

GetBEzsignbulksendIsactiveOk returns a tuple with the BEzsignbulksendIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignbulksendIsactive

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) SetBEzsignbulksendIsactive(v bool)`

SetBEzsignbulksendIsactive sets BEzsignbulksendIsactive field to given value.


### GetObjAudit

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *EzsignbulksendGetObjectV1ResponseMPayload) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


