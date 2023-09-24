# CommunicationResponse

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

## Methods

### NewCommunicationResponse

`func NewCommunicationResponse(pkiCommunicationID int32, eCommunicationImportance FieldECommunicationImportance, eCommunicationType FieldECommunicationType, sCommunicationSubject string, eCommunicationDirection ComputedECommunicationDirection, iCommunicationrecipientCount int32, bCommunicationPrivate bool, objAudit CommonAudit, ) *CommunicationResponse`

NewCommunicationResponse instantiates a new CommunicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationResponseWithDefaults

`func NewCommunicationResponseWithDefaults() *CommunicationResponse`

NewCommunicationResponseWithDefaults instantiates a new CommunicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCommunicationID

`func (o *CommunicationResponse) GetPkiCommunicationID() int32`

GetPkiCommunicationID returns the PkiCommunicationID field if non-nil, zero value otherwise.

### GetPkiCommunicationIDOk

`func (o *CommunicationResponse) GetPkiCommunicationIDOk() (*int32, bool)`

GetPkiCommunicationIDOk returns a tuple with the PkiCommunicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCommunicationID

`func (o *CommunicationResponse) SetPkiCommunicationID(v int32)`

SetPkiCommunicationID sets PkiCommunicationID field to given value.


### GetECommunicationImportance

`func (o *CommunicationResponse) GetECommunicationImportance() FieldECommunicationImportance`

GetECommunicationImportance returns the ECommunicationImportance field if non-nil, zero value otherwise.

### GetECommunicationImportanceOk

`func (o *CommunicationResponse) GetECommunicationImportanceOk() (*FieldECommunicationImportance, bool)`

GetECommunicationImportanceOk returns a tuple with the ECommunicationImportance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationImportance

`func (o *CommunicationResponse) SetECommunicationImportance(v FieldECommunicationImportance)`

SetECommunicationImportance sets ECommunicationImportance field to given value.


### GetECommunicationType

`func (o *CommunicationResponse) GetECommunicationType() FieldECommunicationType`

GetECommunicationType returns the ECommunicationType field if non-nil, zero value otherwise.

### GetECommunicationTypeOk

`func (o *CommunicationResponse) GetECommunicationTypeOk() (*FieldECommunicationType, bool)`

GetECommunicationTypeOk returns a tuple with the ECommunicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationType

`func (o *CommunicationResponse) SetECommunicationType(v FieldECommunicationType)`

SetECommunicationType sets ECommunicationType field to given value.


### GetSCommunicationSubject

`func (o *CommunicationResponse) GetSCommunicationSubject() string`

GetSCommunicationSubject returns the SCommunicationSubject field if non-nil, zero value otherwise.

### GetSCommunicationSubjectOk

`func (o *CommunicationResponse) GetSCommunicationSubjectOk() (*string, bool)`

GetSCommunicationSubjectOk returns a tuple with the SCommunicationSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCommunicationSubject

`func (o *CommunicationResponse) SetSCommunicationSubject(v string)`

SetSCommunicationSubject sets SCommunicationSubject field to given value.


### GetSCommunicationBodyurl

`func (o *CommunicationResponse) GetSCommunicationBodyurl() string`

GetSCommunicationBodyurl returns the SCommunicationBodyurl field if non-nil, zero value otherwise.

### GetSCommunicationBodyurlOk

`func (o *CommunicationResponse) GetSCommunicationBodyurlOk() (*string, bool)`

GetSCommunicationBodyurlOk returns a tuple with the SCommunicationBodyurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCommunicationBodyurl

`func (o *CommunicationResponse) SetSCommunicationBodyurl(v string)`

SetSCommunicationBodyurl sets SCommunicationBodyurl field to given value.

### HasSCommunicationBodyurl

`func (o *CommunicationResponse) HasSCommunicationBodyurl() bool`

HasSCommunicationBodyurl returns a boolean if a field has been set.

### GetECommunicationDirection

`func (o *CommunicationResponse) GetECommunicationDirection() ComputedECommunicationDirection`

GetECommunicationDirection returns the ECommunicationDirection field if non-nil, zero value otherwise.

### GetECommunicationDirectionOk

`func (o *CommunicationResponse) GetECommunicationDirectionOk() (*ComputedECommunicationDirection, bool)`

GetECommunicationDirectionOk returns a tuple with the ECommunicationDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationDirection

`func (o *CommunicationResponse) SetECommunicationDirection(v ComputedECommunicationDirection)`

SetECommunicationDirection sets ECommunicationDirection field to given value.


### GetICommunicationrecipientCount

`func (o *CommunicationResponse) GetICommunicationrecipientCount() int32`

GetICommunicationrecipientCount returns the ICommunicationrecipientCount field if non-nil, zero value otherwise.

### GetICommunicationrecipientCountOk

`func (o *CommunicationResponse) GetICommunicationrecipientCountOk() (*int32, bool)`

GetICommunicationrecipientCountOk returns a tuple with the ICommunicationrecipientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetICommunicationrecipientCount

`func (o *CommunicationResponse) SetICommunicationrecipientCount(v int32)`

SetICommunicationrecipientCount sets ICommunicationrecipientCount field to given value.


### GetBCommunicationPrivate

`func (o *CommunicationResponse) GetBCommunicationPrivate() bool`

GetBCommunicationPrivate returns the BCommunicationPrivate field if non-nil, zero value otherwise.

### GetBCommunicationPrivateOk

`func (o *CommunicationResponse) GetBCommunicationPrivateOk() (*bool, bool)`

GetBCommunicationPrivateOk returns a tuple with the BCommunicationPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCommunicationPrivate

`func (o *CommunicationResponse) SetBCommunicationPrivate(v bool)`

SetBCommunicationPrivate sets BCommunicationPrivate field to given value.


### GetObjDescriptionstaticSender

`func (o *CommunicationResponse) GetObjDescriptionstaticSender() DescriptionstaticResponse`

GetObjDescriptionstaticSender returns the ObjDescriptionstaticSender field if non-nil, zero value otherwise.

### GetObjDescriptionstaticSenderOk

`func (o *CommunicationResponse) GetObjDescriptionstaticSenderOk() (*DescriptionstaticResponse, bool)`

GetObjDescriptionstaticSenderOk returns a tuple with the ObjDescriptionstaticSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDescriptionstaticSender

`func (o *CommunicationResponse) SetObjDescriptionstaticSender(v DescriptionstaticResponse)`

SetObjDescriptionstaticSender sets ObjDescriptionstaticSender field to given value.

### HasObjDescriptionstaticSender

`func (o *CommunicationResponse) HasObjDescriptionstaticSender() bool`

HasObjDescriptionstaticSender returns a boolean if a field has been set.

### GetObjEmailstaticSender

`func (o *CommunicationResponse) GetObjEmailstaticSender() EmailstaticResponse`

GetObjEmailstaticSender returns the ObjEmailstaticSender field if non-nil, zero value otherwise.

### GetObjEmailstaticSenderOk

`func (o *CommunicationResponse) GetObjEmailstaticSenderOk() (*EmailstaticResponse, bool)`

GetObjEmailstaticSenderOk returns a tuple with the ObjEmailstaticSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmailstaticSender

`func (o *CommunicationResponse) SetObjEmailstaticSender(v EmailstaticResponse)`

SetObjEmailstaticSender sets ObjEmailstaticSender field to given value.

### HasObjEmailstaticSender

`func (o *CommunicationResponse) HasObjEmailstaticSender() bool`

HasObjEmailstaticSender returns a boolean if a field has been set.

### GetObjPhonestaticSender

`func (o *CommunicationResponse) GetObjPhonestaticSender() PhonestaticResponse`

GetObjPhonestaticSender returns the ObjPhonestaticSender field if non-nil, zero value otherwise.

### GetObjPhonestaticSenderOk

`func (o *CommunicationResponse) GetObjPhonestaticSenderOk() (*PhonestaticResponse, bool)`

GetObjPhonestaticSenderOk returns a tuple with the ObjPhonestaticSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPhonestaticSender

`func (o *CommunicationResponse) SetObjPhonestaticSender(v PhonestaticResponse)`

SetObjPhonestaticSender sets ObjPhonestaticSender field to given value.

### HasObjPhonestaticSender

`func (o *CommunicationResponse) HasObjPhonestaticSender() bool`

HasObjPhonestaticSender returns a boolean if a field has been set.

### GetObjAudit

`func (o *CommunicationResponse) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *CommunicationResponse) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *CommunicationResponse) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


