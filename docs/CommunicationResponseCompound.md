# CommunicationResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCommunicationID** | **int32** | The unique ID of the Communication. | 
**ECommunicationImportance** | [**FieldECommunicationImportance**](FieldECommunicationImportance.md) |  | 
**ECommunicationType** | [**FieldECommunicationType**](FieldECommunicationType.md) |  | 
**SCommunicationSubject** | **string** | The subject of the Communication | 
**SCommunicationBodyurl** | Pointer to **string** | The url of the body used as body in the Communication | [optional] 
**ECommunicationDirection** | [**ComputedECommunicationDirection**](ComputedECommunicationDirection.md) |  | 
**ICommunicationrecipientCount** | **int32** | The count of Communicationrecipient | 
**BCommunicationPrivate** | **bool** | Whether the Communication is private or not | 
**ObjDescriptionstaticSender** | Pointer to [**DescriptionstaticResponse**](DescriptionstaticResponse.md) |  | [optional] 
**ObjEmailstaticSender** | Pointer to [**EmailstaticResponse**](EmailstaticResponse.md) |  | [optional] 
**ObjPhonestaticSender** | Pointer to [**PhonestaticResponse**](PhonestaticResponse.md) |  | [optional] 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 
**AObjCommunicationattachment** | [**[]CommunicationattachmentResponseCompound**](CommunicationattachmentResponseCompound.md) |  | 
**AObjCommunicationrecipient** | [**[]CommunicationrecipientResponseCompound**](CommunicationrecipientResponseCompound.md) |  | 
**AObjCommunicationexternalrecipient** | [**[]CommunicationexternalrecipientResponseCompound**](CommunicationexternalrecipientResponseCompound.md) |  | 

## Methods

### NewCommunicationResponseCompound

`func NewCommunicationResponseCompound(pkiCommunicationID int32, eCommunicationImportance FieldECommunicationImportance, eCommunicationType FieldECommunicationType, sCommunicationSubject string, eCommunicationDirection ComputedECommunicationDirection, iCommunicationrecipientCount int32, bCommunicationPrivate bool, objAudit CommonAudit, aObjCommunicationattachment []CommunicationattachmentResponseCompound, aObjCommunicationrecipient []CommunicationrecipientResponseCompound, aObjCommunicationexternalrecipient []CommunicationexternalrecipientResponseCompound, ) *CommunicationResponseCompound`

NewCommunicationResponseCompound instantiates a new CommunicationResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationResponseCompoundWithDefaults

`func NewCommunicationResponseCompoundWithDefaults() *CommunicationResponseCompound`

NewCommunicationResponseCompoundWithDefaults instantiates a new CommunicationResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCommunicationID

`func (o *CommunicationResponseCompound) GetPkiCommunicationID() int32`

GetPkiCommunicationID returns the PkiCommunicationID field if non-nil, zero value otherwise.

### GetPkiCommunicationIDOk

`func (o *CommunicationResponseCompound) GetPkiCommunicationIDOk() (*int32, bool)`

GetPkiCommunicationIDOk returns a tuple with the PkiCommunicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCommunicationID

`func (o *CommunicationResponseCompound) SetPkiCommunicationID(v int32)`

SetPkiCommunicationID sets PkiCommunicationID field to given value.


### GetECommunicationImportance

`func (o *CommunicationResponseCompound) GetECommunicationImportance() FieldECommunicationImportance`

GetECommunicationImportance returns the ECommunicationImportance field if non-nil, zero value otherwise.

### GetECommunicationImportanceOk

`func (o *CommunicationResponseCompound) GetECommunicationImportanceOk() (*FieldECommunicationImportance, bool)`

GetECommunicationImportanceOk returns a tuple with the ECommunicationImportance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationImportance

`func (o *CommunicationResponseCompound) SetECommunicationImportance(v FieldECommunicationImportance)`

SetECommunicationImportance sets ECommunicationImportance field to given value.


### GetECommunicationType

`func (o *CommunicationResponseCompound) GetECommunicationType() FieldECommunicationType`

GetECommunicationType returns the ECommunicationType field if non-nil, zero value otherwise.

### GetECommunicationTypeOk

`func (o *CommunicationResponseCompound) GetECommunicationTypeOk() (*FieldECommunicationType, bool)`

GetECommunicationTypeOk returns a tuple with the ECommunicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationType

`func (o *CommunicationResponseCompound) SetECommunicationType(v FieldECommunicationType)`

SetECommunicationType sets ECommunicationType field to given value.


### GetSCommunicationSubject

`func (o *CommunicationResponseCompound) GetSCommunicationSubject() string`

GetSCommunicationSubject returns the SCommunicationSubject field if non-nil, zero value otherwise.

### GetSCommunicationSubjectOk

`func (o *CommunicationResponseCompound) GetSCommunicationSubjectOk() (*string, bool)`

GetSCommunicationSubjectOk returns a tuple with the SCommunicationSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCommunicationSubject

`func (o *CommunicationResponseCompound) SetSCommunicationSubject(v string)`

SetSCommunicationSubject sets SCommunicationSubject field to given value.


### GetSCommunicationBodyurl

`func (o *CommunicationResponseCompound) GetSCommunicationBodyurl() string`

GetSCommunicationBodyurl returns the SCommunicationBodyurl field if non-nil, zero value otherwise.

### GetSCommunicationBodyurlOk

`func (o *CommunicationResponseCompound) GetSCommunicationBodyurlOk() (*string, bool)`

GetSCommunicationBodyurlOk returns a tuple with the SCommunicationBodyurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCommunicationBodyurl

`func (o *CommunicationResponseCompound) SetSCommunicationBodyurl(v string)`

SetSCommunicationBodyurl sets SCommunicationBodyurl field to given value.

### HasSCommunicationBodyurl

`func (o *CommunicationResponseCompound) HasSCommunicationBodyurl() bool`

HasSCommunicationBodyurl returns a boolean if a field has been set.

### GetECommunicationDirection

`func (o *CommunicationResponseCompound) GetECommunicationDirection() ComputedECommunicationDirection`

GetECommunicationDirection returns the ECommunicationDirection field if non-nil, zero value otherwise.

### GetECommunicationDirectionOk

`func (o *CommunicationResponseCompound) GetECommunicationDirectionOk() (*ComputedECommunicationDirection, bool)`

GetECommunicationDirectionOk returns a tuple with the ECommunicationDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationDirection

`func (o *CommunicationResponseCompound) SetECommunicationDirection(v ComputedECommunicationDirection)`

SetECommunicationDirection sets ECommunicationDirection field to given value.


### GetICommunicationrecipientCount

`func (o *CommunicationResponseCompound) GetICommunicationrecipientCount() int32`

GetICommunicationrecipientCount returns the ICommunicationrecipientCount field if non-nil, zero value otherwise.

### GetICommunicationrecipientCountOk

`func (o *CommunicationResponseCompound) GetICommunicationrecipientCountOk() (*int32, bool)`

GetICommunicationrecipientCountOk returns a tuple with the ICommunicationrecipientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetICommunicationrecipientCount

`func (o *CommunicationResponseCompound) SetICommunicationrecipientCount(v int32)`

SetICommunicationrecipientCount sets ICommunicationrecipientCount field to given value.


### GetBCommunicationPrivate

`func (o *CommunicationResponseCompound) GetBCommunicationPrivate() bool`

GetBCommunicationPrivate returns the BCommunicationPrivate field if non-nil, zero value otherwise.

### GetBCommunicationPrivateOk

`func (o *CommunicationResponseCompound) GetBCommunicationPrivateOk() (*bool, bool)`

GetBCommunicationPrivateOk returns a tuple with the BCommunicationPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCommunicationPrivate

`func (o *CommunicationResponseCompound) SetBCommunicationPrivate(v bool)`

SetBCommunicationPrivate sets BCommunicationPrivate field to given value.


### GetObjDescriptionstaticSender

`func (o *CommunicationResponseCompound) GetObjDescriptionstaticSender() DescriptionstaticResponse`

GetObjDescriptionstaticSender returns the ObjDescriptionstaticSender field if non-nil, zero value otherwise.

### GetObjDescriptionstaticSenderOk

`func (o *CommunicationResponseCompound) GetObjDescriptionstaticSenderOk() (*DescriptionstaticResponse, bool)`

GetObjDescriptionstaticSenderOk returns a tuple with the ObjDescriptionstaticSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDescriptionstaticSender

`func (o *CommunicationResponseCompound) SetObjDescriptionstaticSender(v DescriptionstaticResponse)`

SetObjDescriptionstaticSender sets ObjDescriptionstaticSender field to given value.

### HasObjDescriptionstaticSender

`func (o *CommunicationResponseCompound) HasObjDescriptionstaticSender() bool`

HasObjDescriptionstaticSender returns a boolean if a field has been set.

### GetObjEmailstaticSender

`func (o *CommunicationResponseCompound) GetObjEmailstaticSender() EmailstaticResponse`

GetObjEmailstaticSender returns the ObjEmailstaticSender field if non-nil, zero value otherwise.

### GetObjEmailstaticSenderOk

`func (o *CommunicationResponseCompound) GetObjEmailstaticSenderOk() (*EmailstaticResponse, bool)`

GetObjEmailstaticSenderOk returns a tuple with the ObjEmailstaticSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmailstaticSender

`func (o *CommunicationResponseCompound) SetObjEmailstaticSender(v EmailstaticResponse)`

SetObjEmailstaticSender sets ObjEmailstaticSender field to given value.

### HasObjEmailstaticSender

`func (o *CommunicationResponseCompound) HasObjEmailstaticSender() bool`

HasObjEmailstaticSender returns a boolean if a field has been set.

### GetObjPhonestaticSender

`func (o *CommunicationResponseCompound) GetObjPhonestaticSender() PhonestaticResponse`

GetObjPhonestaticSender returns the ObjPhonestaticSender field if non-nil, zero value otherwise.

### GetObjPhonestaticSenderOk

`func (o *CommunicationResponseCompound) GetObjPhonestaticSenderOk() (*PhonestaticResponse, bool)`

GetObjPhonestaticSenderOk returns a tuple with the ObjPhonestaticSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPhonestaticSender

`func (o *CommunicationResponseCompound) SetObjPhonestaticSender(v PhonestaticResponse)`

SetObjPhonestaticSender sets ObjPhonestaticSender field to given value.

### HasObjPhonestaticSender

`func (o *CommunicationResponseCompound) HasObjPhonestaticSender() bool`

HasObjPhonestaticSender returns a boolean if a field has been set.

### GetObjAudit

`func (o *CommunicationResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *CommunicationResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *CommunicationResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.


### GetAObjCommunicationattachment

`func (o *CommunicationResponseCompound) GetAObjCommunicationattachment() []CommunicationattachmentResponseCompound`

GetAObjCommunicationattachment returns the AObjCommunicationattachment field if non-nil, zero value otherwise.

### GetAObjCommunicationattachmentOk

`func (o *CommunicationResponseCompound) GetAObjCommunicationattachmentOk() (*[]CommunicationattachmentResponseCompound, bool)`

GetAObjCommunicationattachmentOk returns a tuple with the AObjCommunicationattachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjCommunicationattachment

`func (o *CommunicationResponseCompound) SetAObjCommunicationattachment(v []CommunicationattachmentResponseCompound)`

SetAObjCommunicationattachment sets AObjCommunicationattachment field to given value.


### GetAObjCommunicationrecipient

`func (o *CommunicationResponseCompound) GetAObjCommunicationrecipient() []CommunicationrecipientResponseCompound`

GetAObjCommunicationrecipient returns the AObjCommunicationrecipient field if non-nil, zero value otherwise.

### GetAObjCommunicationrecipientOk

`func (o *CommunicationResponseCompound) GetAObjCommunicationrecipientOk() (*[]CommunicationrecipientResponseCompound, bool)`

GetAObjCommunicationrecipientOk returns a tuple with the AObjCommunicationrecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjCommunicationrecipient

`func (o *CommunicationResponseCompound) SetAObjCommunicationrecipient(v []CommunicationrecipientResponseCompound)`

SetAObjCommunicationrecipient sets AObjCommunicationrecipient field to given value.


### GetAObjCommunicationexternalrecipient

`func (o *CommunicationResponseCompound) GetAObjCommunicationexternalrecipient() []CommunicationexternalrecipientResponseCompound`

GetAObjCommunicationexternalrecipient returns the AObjCommunicationexternalrecipient field if non-nil, zero value otherwise.

### GetAObjCommunicationexternalrecipientOk

`func (o *CommunicationResponseCompound) GetAObjCommunicationexternalrecipientOk() (*[]CommunicationexternalrecipientResponseCompound, bool)`

GetAObjCommunicationexternalrecipientOk returns a tuple with the AObjCommunicationexternalrecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjCommunicationexternalrecipient

`func (o *CommunicationResponseCompound) SetAObjCommunicationexternalrecipient(v []CommunicationexternalrecipientResponseCompound)`

SetAObjCommunicationexternalrecipient sets AObjCommunicationexternalrecipient field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


