# CustomCommunicationListElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCommunicationID** | **int32** | The unique ID of the Communication. | 
**DtCreatedDate** | **string** | The date and time at which the object was created | 
**ECommunicationDirection** | [**ComputedECommunicationDirection**](ComputedECommunicationDirection.md) |  | 
**ECommunicationImportance** | [**FieldECommunicationImportance**](FieldECommunicationImportance.md) |  | 
**ECommunicationType** | [**FieldECommunicationType**](FieldECommunicationType.md) |  | 
**ICommunicationrecipientCount** | **int32** | The count of Communicationrecipient | 
**SCommunicationSubject** | **string** | The subject of the Communication | 
**SCommunicationSender** | **string** | The sender name of the Communication | 
**SCommunicationRecipient** | **string** | The recipients&#39; name of the Communication | 

## Methods

### NewCustomCommunicationListElementResponse

`func NewCustomCommunicationListElementResponse(pkiCommunicationID int32, dtCreatedDate string, eCommunicationDirection ComputedECommunicationDirection, eCommunicationImportance FieldECommunicationImportance, eCommunicationType FieldECommunicationType, iCommunicationrecipientCount int32, sCommunicationSubject string, sCommunicationSender string, sCommunicationRecipient string, ) *CustomCommunicationListElementResponse`

NewCustomCommunicationListElementResponse instantiates a new CustomCommunicationListElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomCommunicationListElementResponseWithDefaults

`func NewCustomCommunicationListElementResponseWithDefaults() *CustomCommunicationListElementResponse`

NewCustomCommunicationListElementResponseWithDefaults instantiates a new CustomCommunicationListElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCommunicationID

`func (o *CustomCommunicationListElementResponse) GetPkiCommunicationID() int32`

GetPkiCommunicationID returns the PkiCommunicationID field if non-nil, zero value otherwise.

### GetPkiCommunicationIDOk

`func (o *CustomCommunicationListElementResponse) GetPkiCommunicationIDOk() (*int32, bool)`

GetPkiCommunicationIDOk returns a tuple with the PkiCommunicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCommunicationID

`func (o *CustomCommunicationListElementResponse) SetPkiCommunicationID(v int32)`

SetPkiCommunicationID sets PkiCommunicationID field to given value.


### GetDtCreatedDate

`func (o *CustomCommunicationListElementResponse) GetDtCreatedDate() string`

GetDtCreatedDate returns the DtCreatedDate field if non-nil, zero value otherwise.

### GetDtCreatedDateOk

`func (o *CustomCommunicationListElementResponse) GetDtCreatedDateOk() (*string, bool)`

GetDtCreatedDateOk returns a tuple with the DtCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtCreatedDate

`func (o *CustomCommunicationListElementResponse) SetDtCreatedDate(v string)`

SetDtCreatedDate sets DtCreatedDate field to given value.


### GetECommunicationDirection

`func (o *CustomCommunicationListElementResponse) GetECommunicationDirection() ComputedECommunicationDirection`

GetECommunicationDirection returns the ECommunicationDirection field if non-nil, zero value otherwise.

### GetECommunicationDirectionOk

`func (o *CustomCommunicationListElementResponse) GetECommunicationDirectionOk() (*ComputedECommunicationDirection, bool)`

GetECommunicationDirectionOk returns a tuple with the ECommunicationDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationDirection

`func (o *CustomCommunicationListElementResponse) SetECommunicationDirection(v ComputedECommunicationDirection)`

SetECommunicationDirection sets ECommunicationDirection field to given value.


### GetECommunicationImportance

`func (o *CustomCommunicationListElementResponse) GetECommunicationImportance() FieldECommunicationImportance`

GetECommunicationImportance returns the ECommunicationImportance field if non-nil, zero value otherwise.

### GetECommunicationImportanceOk

`func (o *CustomCommunicationListElementResponse) GetECommunicationImportanceOk() (*FieldECommunicationImportance, bool)`

GetECommunicationImportanceOk returns a tuple with the ECommunicationImportance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationImportance

`func (o *CustomCommunicationListElementResponse) SetECommunicationImportance(v FieldECommunicationImportance)`

SetECommunicationImportance sets ECommunicationImportance field to given value.


### GetECommunicationType

`func (o *CustomCommunicationListElementResponse) GetECommunicationType() FieldECommunicationType`

GetECommunicationType returns the ECommunicationType field if non-nil, zero value otherwise.

### GetECommunicationTypeOk

`func (o *CustomCommunicationListElementResponse) GetECommunicationTypeOk() (*FieldECommunicationType, bool)`

GetECommunicationTypeOk returns a tuple with the ECommunicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationType

`func (o *CustomCommunicationListElementResponse) SetECommunicationType(v FieldECommunicationType)`

SetECommunicationType sets ECommunicationType field to given value.


### GetICommunicationrecipientCount

`func (o *CustomCommunicationListElementResponse) GetICommunicationrecipientCount() int32`

GetICommunicationrecipientCount returns the ICommunicationrecipientCount field if non-nil, zero value otherwise.

### GetICommunicationrecipientCountOk

`func (o *CustomCommunicationListElementResponse) GetICommunicationrecipientCountOk() (*int32, bool)`

GetICommunicationrecipientCountOk returns a tuple with the ICommunicationrecipientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetICommunicationrecipientCount

`func (o *CustomCommunicationListElementResponse) SetICommunicationrecipientCount(v int32)`

SetICommunicationrecipientCount sets ICommunicationrecipientCount field to given value.


### GetSCommunicationSubject

`func (o *CustomCommunicationListElementResponse) GetSCommunicationSubject() string`

GetSCommunicationSubject returns the SCommunicationSubject field if non-nil, zero value otherwise.

### GetSCommunicationSubjectOk

`func (o *CustomCommunicationListElementResponse) GetSCommunicationSubjectOk() (*string, bool)`

GetSCommunicationSubjectOk returns a tuple with the SCommunicationSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCommunicationSubject

`func (o *CustomCommunicationListElementResponse) SetSCommunicationSubject(v string)`

SetSCommunicationSubject sets SCommunicationSubject field to given value.


### GetSCommunicationSender

`func (o *CustomCommunicationListElementResponse) GetSCommunicationSender() string`

GetSCommunicationSender returns the SCommunicationSender field if non-nil, zero value otherwise.

### GetSCommunicationSenderOk

`func (o *CustomCommunicationListElementResponse) GetSCommunicationSenderOk() (*string, bool)`

GetSCommunicationSenderOk returns a tuple with the SCommunicationSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCommunicationSender

`func (o *CustomCommunicationListElementResponse) SetSCommunicationSender(v string)`

SetSCommunicationSender sets SCommunicationSender field to given value.


### GetSCommunicationRecipient

`func (o *CustomCommunicationListElementResponse) GetSCommunicationRecipient() string`

GetSCommunicationRecipient returns the SCommunicationRecipient field if non-nil, zero value otherwise.

### GetSCommunicationRecipientOk

`func (o *CustomCommunicationListElementResponse) GetSCommunicationRecipientOk() (*string, bool)`

GetSCommunicationRecipientOk returns a tuple with the SCommunicationRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCommunicationRecipient

`func (o *CustomCommunicationListElementResponse) SetSCommunicationRecipient(v string)`

SetSCommunicationRecipient sets SCommunicationRecipient field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


