# CommunicationRequest

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

## Methods

### NewCommunicationRequest

`func NewCommunicationRequest(eCommunicationType FieldECommunicationType, tCommunicationBody string, bCommunicationPrivate bool, ) *CommunicationRequest`

NewCommunicationRequest instantiates a new CommunicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationRequestWithDefaults

`func NewCommunicationRequestWithDefaults() *CommunicationRequest`

NewCommunicationRequestWithDefaults instantiates a new CommunicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCommunicationID

`func (o *CommunicationRequest) GetPkiCommunicationID() int32`

GetPkiCommunicationID returns the PkiCommunicationID field if non-nil, zero value otherwise.

### GetPkiCommunicationIDOk

`func (o *CommunicationRequest) GetPkiCommunicationIDOk() (*int32, bool)`

GetPkiCommunicationIDOk returns a tuple with the PkiCommunicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCommunicationID

`func (o *CommunicationRequest) SetPkiCommunicationID(v int32)`

SetPkiCommunicationID sets PkiCommunicationID field to given value.

### HasPkiCommunicationID

`func (o *CommunicationRequest) HasPkiCommunicationID() bool`

HasPkiCommunicationID returns a boolean if a field has been set.

### GetECommunicationImportance

`func (o *CommunicationRequest) GetECommunicationImportance() FieldECommunicationImportance`

GetECommunicationImportance returns the ECommunicationImportance field if non-nil, zero value otherwise.

### GetECommunicationImportanceOk

`func (o *CommunicationRequest) GetECommunicationImportanceOk() (*FieldECommunicationImportance, bool)`

GetECommunicationImportanceOk returns a tuple with the ECommunicationImportance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationImportance

`func (o *CommunicationRequest) SetECommunicationImportance(v FieldECommunicationImportance)`

SetECommunicationImportance sets ECommunicationImportance field to given value.

### HasECommunicationImportance

`func (o *CommunicationRequest) HasECommunicationImportance() bool`

HasECommunicationImportance returns a boolean if a field has been set.

### GetECommunicationType

`func (o *CommunicationRequest) GetECommunicationType() FieldECommunicationType`

GetECommunicationType returns the ECommunicationType field if non-nil, zero value otherwise.

### GetECommunicationTypeOk

`func (o *CommunicationRequest) GetECommunicationTypeOk() (*FieldECommunicationType, bool)`

GetECommunicationTypeOk returns a tuple with the ECommunicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationType

`func (o *CommunicationRequest) SetECommunicationType(v FieldECommunicationType)`

SetECommunicationType sets ECommunicationType field to given value.


### GetObjCommunicationsender

`func (o *CommunicationRequest) GetObjCommunicationsender() CustomCommunicationsenderRequest`

GetObjCommunicationsender returns the ObjCommunicationsender field if non-nil, zero value otherwise.

### GetObjCommunicationsenderOk

`func (o *CommunicationRequest) GetObjCommunicationsenderOk() (*CustomCommunicationsenderRequest, bool)`

GetObjCommunicationsenderOk returns a tuple with the ObjCommunicationsender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCommunicationsender

`func (o *CommunicationRequest) SetObjCommunicationsender(v CustomCommunicationsenderRequest)`

SetObjCommunicationsender sets ObjCommunicationsender field to given value.

### HasObjCommunicationsender

`func (o *CommunicationRequest) HasObjCommunicationsender() bool`

HasObjCommunicationsender returns a boolean if a field has been set.

### GetSCommunicationSubject

`func (o *CommunicationRequest) GetSCommunicationSubject() string`

GetSCommunicationSubject returns the SCommunicationSubject field if non-nil, zero value otherwise.

### GetSCommunicationSubjectOk

`func (o *CommunicationRequest) GetSCommunicationSubjectOk() (*string, bool)`

GetSCommunicationSubjectOk returns a tuple with the SCommunicationSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCommunicationSubject

`func (o *CommunicationRequest) SetSCommunicationSubject(v string)`

SetSCommunicationSubject sets SCommunicationSubject field to given value.

### HasSCommunicationSubject

`func (o *CommunicationRequest) HasSCommunicationSubject() bool`

HasSCommunicationSubject returns a boolean if a field has been set.

### GetTCommunicationBody

`func (o *CommunicationRequest) GetTCommunicationBody() string`

GetTCommunicationBody returns the TCommunicationBody field if non-nil, zero value otherwise.

### GetTCommunicationBodyOk

`func (o *CommunicationRequest) GetTCommunicationBodyOk() (*string, bool)`

GetTCommunicationBodyOk returns a tuple with the TCommunicationBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTCommunicationBody

`func (o *CommunicationRequest) SetTCommunicationBody(v string)`

SetTCommunicationBody sets TCommunicationBody field to given value.


### GetBCommunicationPrivate

`func (o *CommunicationRequest) GetBCommunicationPrivate() bool`

GetBCommunicationPrivate returns the BCommunicationPrivate field if non-nil, zero value otherwise.

### GetBCommunicationPrivateOk

`func (o *CommunicationRequest) GetBCommunicationPrivateOk() (*bool, bool)`

GetBCommunicationPrivateOk returns a tuple with the BCommunicationPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCommunicationPrivate

`func (o *CommunicationRequest) SetBCommunicationPrivate(v bool)`

SetBCommunicationPrivate sets BCommunicationPrivate field to given value.


### GetECommunicationAttachmenttype

`func (o *CommunicationRequest) GetECommunicationAttachmenttype() string`

GetECommunicationAttachmenttype returns the ECommunicationAttachmenttype field if non-nil, zero value otherwise.

### GetECommunicationAttachmenttypeOk

`func (o *CommunicationRequest) GetECommunicationAttachmenttypeOk() (*string, bool)`

GetECommunicationAttachmenttypeOk returns a tuple with the ECommunicationAttachmenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationAttachmenttype

`func (o *CommunicationRequest) SetECommunicationAttachmenttype(v string)`

SetECommunicationAttachmenttype sets ECommunicationAttachmenttype field to given value.

### HasECommunicationAttachmenttype

`func (o *CommunicationRequest) HasECommunicationAttachmenttype() bool`

HasECommunicationAttachmenttype returns a boolean if a field has been set.

### GetICommunicationAttachmentlinkexpiration

`func (o *CommunicationRequest) GetICommunicationAttachmentlinkexpiration() int32`

GetICommunicationAttachmentlinkexpiration returns the ICommunicationAttachmentlinkexpiration field if non-nil, zero value otherwise.

### GetICommunicationAttachmentlinkexpirationOk

`func (o *CommunicationRequest) GetICommunicationAttachmentlinkexpirationOk() (*int32, bool)`

GetICommunicationAttachmentlinkexpirationOk returns a tuple with the ICommunicationAttachmentlinkexpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetICommunicationAttachmentlinkexpiration

`func (o *CommunicationRequest) SetICommunicationAttachmentlinkexpiration(v int32)`

SetICommunicationAttachmentlinkexpiration sets ICommunicationAttachmentlinkexpiration field to given value.

### HasICommunicationAttachmentlinkexpiration

`func (o *CommunicationRequest) HasICommunicationAttachmentlinkexpiration() bool`

HasICommunicationAttachmentlinkexpiration returns a boolean if a field has been set.

### GetBCommunicationReadreceipt

`func (o *CommunicationRequest) GetBCommunicationReadreceipt() bool`

GetBCommunicationReadreceipt returns the BCommunicationReadreceipt field if non-nil, zero value otherwise.

### GetBCommunicationReadreceiptOk

`func (o *CommunicationRequest) GetBCommunicationReadreceiptOk() (*bool, bool)`

GetBCommunicationReadreceiptOk returns a tuple with the BCommunicationReadreceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCommunicationReadreceipt

`func (o *CommunicationRequest) SetBCommunicationReadreceipt(v bool)`

SetBCommunicationReadreceipt sets BCommunicationReadreceipt field to given value.

### HasBCommunicationReadreceipt

`func (o *CommunicationRequest) HasBCommunicationReadreceipt() bool`

HasBCommunicationReadreceipt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


