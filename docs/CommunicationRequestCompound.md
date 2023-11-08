# CommunicationRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCommunicationID** | Pointer to **int32** | The unique ID of the Communication. | [optional] 
**ECommunicationImportance** | Pointer to [**FieldECommunicationImportance**](FieldECommunicationImportance.md) |  | [optional] 
**ECommunicationType** | [**FieldECommunicationType**](FieldECommunicationType.md) |  | 
**ObjCommunicationsender** | Pointer to [**CustomCommunicationsenderRequest**](CustomCommunicationsenderRequest.md) |  | [optional] 
**SCommunicationSubject** | Pointer to **string** | The subject of the Communication | [optional] 
**TCommunicationBody** | **string** | The Body of the Communication | 
**BCommunicationPrivate** | **bool** | Whether the Communication is private or not | 
**ECommunicationAttachmenttype** | Pointer to **string** | How the attachment should be included in the email.   Only used if eCommunicationType is **Email** | [optional] 
**ICommunicationAttachmentlinkexpiration** | Pointer to **int32** | The number of days before the attachment link expired.   Only used if eCommunicationType is **Email** and eCommunicationattachmentType is **Link** | [optional] 
**BCommunicationReadreceipt** | Pointer to **bool** | Whether we ask for a read receipt or not. | [optional] 
**AObjCommunicationattachment** | [**[]CustomCommunicationattachmentRequest**](CustomCommunicationattachmentRequest.md) |  | 
**AObjCommunicationrecipient** | [**[]CommunicationrecipientRequestCompound**](CommunicationrecipientRequestCompound.md) |  | 
**AObjCommunicationreference** | [**[]CommunicationreferenceRequestCompound**](CommunicationreferenceRequestCompound.md) |  | 
**AObjCommunicationexternalrecipient** | [**[]CommunicationexternalrecipientRequestCompound**](CommunicationexternalrecipientRequestCompound.md) |  | 

## Methods

### NewCommunicationRequestCompound

`func NewCommunicationRequestCompound(eCommunicationType FieldECommunicationType, tCommunicationBody string, bCommunicationPrivate bool, aObjCommunicationattachment []CustomCommunicationattachmentRequest, aObjCommunicationrecipient []CommunicationrecipientRequestCompound, aObjCommunicationreference []CommunicationreferenceRequestCompound, aObjCommunicationexternalrecipient []CommunicationexternalrecipientRequestCompound, ) *CommunicationRequestCompound`

NewCommunicationRequestCompound instantiates a new CommunicationRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationRequestCompoundWithDefaults

`func NewCommunicationRequestCompoundWithDefaults() *CommunicationRequestCompound`

NewCommunicationRequestCompoundWithDefaults instantiates a new CommunicationRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCommunicationID

`func (o *CommunicationRequestCompound) GetPkiCommunicationID() int32`

GetPkiCommunicationID returns the PkiCommunicationID field if non-nil, zero value otherwise.

### GetPkiCommunicationIDOk

`func (o *CommunicationRequestCompound) GetPkiCommunicationIDOk() (*int32, bool)`

GetPkiCommunicationIDOk returns a tuple with the PkiCommunicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCommunicationID

`func (o *CommunicationRequestCompound) SetPkiCommunicationID(v int32)`

SetPkiCommunicationID sets PkiCommunicationID field to given value.

### HasPkiCommunicationID

`func (o *CommunicationRequestCompound) HasPkiCommunicationID() bool`

HasPkiCommunicationID returns a boolean if a field has been set.

### GetECommunicationImportance

`func (o *CommunicationRequestCompound) GetECommunicationImportance() FieldECommunicationImportance`

GetECommunicationImportance returns the ECommunicationImportance field if non-nil, zero value otherwise.

### GetECommunicationImportanceOk

`func (o *CommunicationRequestCompound) GetECommunicationImportanceOk() (*FieldECommunicationImportance, bool)`

GetECommunicationImportanceOk returns a tuple with the ECommunicationImportance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationImportance

`func (o *CommunicationRequestCompound) SetECommunicationImportance(v FieldECommunicationImportance)`

SetECommunicationImportance sets ECommunicationImportance field to given value.

### HasECommunicationImportance

`func (o *CommunicationRequestCompound) HasECommunicationImportance() bool`

HasECommunicationImportance returns a boolean if a field has been set.

### GetECommunicationType

`func (o *CommunicationRequestCompound) GetECommunicationType() FieldECommunicationType`

GetECommunicationType returns the ECommunicationType field if non-nil, zero value otherwise.

### GetECommunicationTypeOk

`func (o *CommunicationRequestCompound) GetECommunicationTypeOk() (*FieldECommunicationType, bool)`

GetECommunicationTypeOk returns a tuple with the ECommunicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationType

`func (o *CommunicationRequestCompound) SetECommunicationType(v FieldECommunicationType)`

SetECommunicationType sets ECommunicationType field to given value.


### GetObjCommunicationsender

`func (o *CommunicationRequestCompound) GetObjCommunicationsender() CustomCommunicationsenderRequest`

GetObjCommunicationsender returns the ObjCommunicationsender field if non-nil, zero value otherwise.

### GetObjCommunicationsenderOk

`func (o *CommunicationRequestCompound) GetObjCommunicationsenderOk() (*CustomCommunicationsenderRequest, bool)`

GetObjCommunicationsenderOk returns a tuple with the ObjCommunicationsender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCommunicationsender

`func (o *CommunicationRequestCompound) SetObjCommunicationsender(v CustomCommunicationsenderRequest)`

SetObjCommunicationsender sets ObjCommunicationsender field to given value.

### HasObjCommunicationsender

`func (o *CommunicationRequestCompound) HasObjCommunicationsender() bool`

HasObjCommunicationsender returns a boolean if a field has been set.

### GetSCommunicationSubject

`func (o *CommunicationRequestCompound) GetSCommunicationSubject() string`

GetSCommunicationSubject returns the SCommunicationSubject field if non-nil, zero value otherwise.

### GetSCommunicationSubjectOk

`func (o *CommunicationRequestCompound) GetSCommunicationSubjectOk() (*string, bool)`

GetSCommunicationSubjectOk returns a tuple with the SCommunicationSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCommunicationSubject

`func (o *CommunicationRequestCompound) SetSCommunicationSubject(v string)`

SetSCommunicationSubject sets SCommunicationSubject field to given value.

### HasSCommunicationSubject

`func (o *CommunicationRequestCompound) HasSCommunicationSubject() bool`

HasSCommunicationSubject returns a boolean if a field has been set.

### GetTCommunicationBody

`func (o *CommunicationRequestCompound) GetTCommunicationBody() string`

GetTCommunicationBody returns the TCommunicationBody field if non-nil, zero value otherwise.

### GetTCommunicationBodyOk

`func (o *CommunicationRequestCompound) GetTCommunicationBodyOk() (*string, bool)`

GetTCommunicationBodyOk returns a tuple with the TCommunicationBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTCommunicationBody

`func (o *CommunicationRequestCompound) SetTCommunicationBody(v string)`

SetTCommunicationBody sets TCommunicationBody field to given value.


### GetBCommunicationPrivate

`func (o *CommunicationRequestCompound) GetBCommunicationPrivate() bool`

GetBCommunicationPrivate returns the BCommunicationPrivate field if non-nil, zero value otherwise.

### GetBCommunicationPrivateOk

`func (o *CommunicationRequestCompound) GetBCommunicationPrivateOk() (*bool, bool)`

GetBCommunicationPrivateOk returns a tuple with the BCommunicationPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCommunicationPrivate

`func (o *CommunicationRequestCompound) SetBCommunicationPrivate(v bool)`

SetBCommunicationPrivate sets BCommunicationPrivate field to given value.


### GetECommunicationAttachmenttype

`func (o *CommunicationRequestCompound) GetECommunicationAttachmenttype() string`

GetECommunicationAttachmenttype returns the ECommunicationAttachmenttype field if non-nil, zero value otherwise.

### GetECommunicationAttachmenttypeOk

`func (o *CommunicationRequestCompound) GetECommunicationAttachmenttypeOk() (*string, bool)`

GetECommunicationAttachmenttypeOk returns a tuple with the ECommunicationAttachmenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationAttachmenttype

`func (o *CommunicationRequestCompound) SetECommunicationAttachmenttype(v string)`

SetECommunicationAttachmenttype sets ECommunicationAttachmenttype field to given value.

### HasECommunicationAttachmenttype

`func (o *CommunicationRequestCompound) HasECommunicationAttachmenttype() bool`

HasECommunicationAttachmenttype returns a boolean if a field has been set.

### GetICommunicationAttachmentlinkexpiration

`func (o *CommunicationRequestCompound) GetICommunicationAttachmentlinkexpiration() int32`

GetICommunicationAttachmentlinkexpiration returns the ICommunicationAttachmentlinkexpiration field if non-nil, zero value otherwise.

### GetICommunicationAttachmentlinkexpirationOk

`func (o *CommunicationRequestCompound) GetICommunicationAttachmentlinkexpirationOk() (*int32, bool)`

GetICommunicationAttachmentlinkexpirationOk returns a tuple with the ICommunicationAttachmentlinkexpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetICommunicationAttachmentlinkexpiration

`func (o *CommunicationRequestCompound) SetICommunicationAttachmentlinkexpiration(v int32)`

SetICommunicationAttachmentlinkexpiration sets ICommunicationAttachmentlinkexpiration field to given value.

### HasICommunicationAttachmentlinkexpiration

`func (o *CommunicationRequestCompound) HasICommunicationAttachmentlinkexpiration() bool`

HasICommunicationAttachmentlinkexpiration returns a boolean if a field has been set.

### GetBCommunicationReadreceipt

`func (o *CommunicationRequestCompound) GetBCommunicationReadreceipt() bool`

GetBCommunicationReadreceipt returns the BCommunicationReadreceipt field if non-nil, zero value otherwise.

### GetBCommunicationReadreceiptOk

`func (o *CommunicationRequestCompound) GetBCommunicationReadreceiptOk() (*bool, bool)`

GetBCommunicationReadreceiptOk returns a tuple with the BCommunicationReadreceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCommunicationReadreceipt

`func (o *CommunicationRequestCompound) SetBCommunicationReadreceipt(v bool)`

SetBCommunicationReadreceipt sets BCommunicationReadreceipt field to given value.

### HasBCommunicationReadreceipt

`func (o *CommunicationRequestCompound) HasBCommunicationReadreceipt() bool`

HasBCommunicationReadreceipt returns a boolean if a field has been set.

### GetAObjCommunicationattachment

`func (o *CommunicationRequestCompound) GetAObjCommunicationattachment() []CustomCommunicationattachmentRequest`

GetAObjCommunicationattachment returns the AObjCommunicationattachment field if non-nil, zero value otherwise.

### GetAObjCommunicationattachmentOk

`func (o *CommunicationRequestCompound) GetAObjCommunicationattachmentOk() (*[]CustomCommunicationattachmentRequest, bool)`

GetAObjCommunicationattachmentOk returns a tuple with the AObjCommunicationattachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjCommunicationattachment

`func (o *CommunicationRequestCompound) SetAObjCommunicationattachment(v []CustomCommunicationattachmentRequest)`

SetAObjCommunicationattachment sets AObjCommunicationattachment field to given value.


### GetAObjCommunicationrecipient

`func (o *CommunicationRequestCompound) GetAObjCommunicationrecipient() []CommunicationrecipientRequestCompound`

GetAObjCommunicationrecipient returns the AObjCommunicationrecipient field if non-nil, zero value otherwise.

### GetAObjCommunicationrecipientOk

`func (o *CommunicationRequestCompound) GetAObjCommunicationrecipientOk() (*[]CommunicationrecipientRequestCompound, bool)`

GetAObjCommunicationrecipientOk returns a tuple with the AObjCommunicationrecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjCommunicationrecipient

`func (o *CommunicationRequestCompound) SetAObjCommunicationrecipient(v []CommunicationrecipientRequestCompound)`

SetAObjCommunicationrecipient sets AObjCommunicationrecipient field to given value.


### GetAObjCommunicationreference

`func (o *CommunicationRequestCompound) GetAObjCommunicationreference() []CommunicationreferenceRequestCompound`

GetAObjCommunicationreference returns the AObjCommunicationreference field if non-nil, zero value otherwise.

### GetAObjCommunicationreferenceOk

`func (o *CommunicationRequestCompound) GetAObjCommunicationreferenceOk() (*[]CommunicationreferenceRequestCompound, bool)`

GetAObjCommunicationreferenceOk returns a tuple with the AObjCommunicationreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjCommunicationreference

`func (o *CommunicationRequestCompound) SetAObjCommunicationreference(v []CommunicationreferenceRequestCompound)`

SetAObjCommunicationreference sets AObjCommunicationreference field to given value.


### GetAObjCommunicationexternalrecipient

`func (o *CommunicationRequestCompound) GetAObjCommunicationexternalrecipient() []CommunicationexternalrecipientRequestCompound`

GetAObjCommunicationexternalrecipient returns the AObjCommunicationexternalrecipient field if non-nil, zero value otherwise.

### GetAObjCommunicationexternalrecipientOk

`func (o *CommunicationRequestCompound) GetAObjCommunicationexternalrecipientOk() (*[]CommunicationexternalrecipientRequestCompound, bool)`

GetAObjCommunicationexternalrecipientOk returns a tuple with the AObjCommunicationexternalrecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjCommunicationexternalrecipient

`func (o *CommunicationRequestCompound) SetAObjCommunicationexternalrecipient(v []CommunicationexternalrecipientRequestCompound)`

SetAObjCommunicationexternalrecipient sets AObjCommunicationexternalrecipient field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


