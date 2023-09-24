# CommunicationexternalrecipientResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCommunicationexternalrecipientID** | **int32** | The unique ID of the Communicationexternalrecipient | 
**ECommunicationexternalrecipientType** | [**FieldECommunicationexternalrecipientType**](FieldECommunicationexternalrecipientType.md) |  | 
**ObjDescriptionstatic** | [**DescriptionstaticResponseCompound**](DescriptionstaticResponseCompound.md) |  | 
**ObjEmailstatic** | Pointer to [**EmailstaticResponseCompound**](EmailstaticResponseCompound.md) |  | [optional] 
**ObjPhonestatic** | Pointer to [**PhonestaticResponseCompound**](PhonestaticResponseCompound.md) |  | [optional] 

## Methods

### NewCommunicationexternalrecipientResponse

`func NewCommunicationexternalrecipientResponse(pkiCommunicationexternalrecipientID int32, eCommunicationexternalrecipientType FieldECommunicationexternalrecipientType, objDescriptionstatic DescriptionstaticResponseCompound, ) *CommunicationexternalrecipientResponse`

NewCommunicationexternalrecipientResponse instantiates a new CommunicationexternalrecipientResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationexternalrecipientResponseWithDefaults

`func NewCommunicationexternalrecipientResponseWithDefaults() *CommunicationexternalrecipientResponse`

NewCommunicationexternalrecipientResponseWithDefaults instantiates a new CommunicationexternalrecipientResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCommunicationexternalrecipientID

`func (o *CommunicationexternalrecipientResponse) GetPkiCommunicationexternalrecipientID() int32`

GetPkiCommunicationexternalrecipientID returns the PkiCommunicationexternalrecipientID field if non-nil, zero value otherwise.

### GetPkiCommunicationexternalrecipientIDOk

`func (o *CommunicationexternalrecipientResponse) GetPkiCommunicationexternalrecipientIDOk() (*int32, bool)`

GetPkiCommunicationexternalrecipientIDOk returns a tuple with the PkiCommunicationexternalrecipientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCommunicationexternalrecipientID

`func (o *CommunicationexternalrecipientResponse) SetPkiCommunicationexternalrecipientID(v int32)`

SetPkiCommunicationexternalrecipientID sets PkiCommunicationexternalrecipientID field to given value.


### GetECommunicationexternalrecipientType

`func (o *CommunicationexternalrecipientResponse) GetECommunicationexternalrecipientType() FieldECommunicationexternalrecipientType`

GetECommunicationexternalrecipientType returns the ECommunicationexternalrecipientType field if non-nil, zero value otherwise.

### GetECommunicationexternalrecipientTypeOk

`func (o *CommunicationexternalrecipientResponse) GetECommunicationexternalrecipientTypeOk() (*FieldECommunicationexternalrecipientType, bool)`

GetECommunicationexternalrecipientTypeOk returns a tuple with the ECommunicationexternalrecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationexternalrecipientType

`func (o *CommunicationexternalrecipientResponse) SetECommunicationexternalrecipientType(v FieldECommunicationexternalrecipientType)`

SetECommunicationexternalrecipientType sets ECommunicationexternalrecipientType field to given value.


### GetObjDescriptionstatic

`func (o *CommunicationexternalrecipientResponse) GetObjDescriptionstatic() DescriptionstaticResponseCompound`

GetObjDescriptionstatic returns the ObjDescriptionstatic field if non-nil, zero value otherwise.

### GetObjDescriptionstaticOk

`func (o *CommunicationexternalrecipientResponse) GetObjDescriptionstaticOk() (*DescriptionstaticResponseCompound, bool)`

GetObjDescriptionstaticOk returns a tuple with the ObjDescriptionstatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDescriptionstatic

`func (o *CommunicationexternalrecipientResponse) SetObjDescriptionstatic(v DescriptionstaticResponseCompound)`

SetObjDescriptionstatic sets ObjDescriptionstatic field to given value.


### GetObjEmailstatic

`func (o *CommunicationexternalrecipientResponse) GetObjEmailstatic() EmailstaticResponseCompound`

GetObjEmailstatic returns the ObjEmailstatic field if non-nil, zero value otherwise.

### GetObjEmailstaticOk

`func (o *CommunicationexternalrecipientResponse) GetObjEmailstaticOk() (*EmailstaticResponseCompound, bool)`

GetObjEmailstaticOk returns a tuple with the ObjEmailstatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmailstatic

`func (o *CommunicationexternalrecipientResponse) SetObjEmailstatic(v EmailstaticResponseCompound)`

SetObjEmailstatic sets ObjEmailstatic field to given value.

### HasObjEmailstatic

`func (o *CommunicationexternalrecipientResponse) HasObjEmailstatic() bool`

HasObjEmailstatic returns a boolean if a field has been set.

### GetObjPhonestatic

`func (o *CommunicationexternalrecipientResponse) GetObjPhonestatic() PhonestaticResponseCompound`

GetObjPhonestatic returns the ObjPhonestatic field if non-nil, zero value otherwise.

### GetObjPhonestaticOk

`func (o *CommunicationexternalrecipientResponse) GetObjPhonestaticOk() (*PhonestaticResponseCompound, bool)`

GetObjPhonestaticOk returns a tuple with the ObjPhonestatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPhonestatic

`func (o *CommunicationexternalrecipientResponse) SetObjPhonestatic(v PhonestaticResponseCompound)`

SetObjPhonestatic sets ObjPhonestatic field to given value.

### HasObjPhonestatic

`func (o *CommunicationexternalrecipientResponse) HasObjPhonestatic() bool`

HasObjPhonestatic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


