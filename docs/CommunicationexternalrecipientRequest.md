# CommunicationexternalrecipientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCommunicationexternalrecipientID** | Pointer to **int32** | The unique ID of the Communicationexternalrecipient | [optional] 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**SPhoneE164** | Pointer to **string** | A phone number in E.164 Format | [optional] 
**ECommunicationexternalrecipientType** | Pointer to [**FieldECommunicationexternalrecipientType**](FieldECommunicationexternalrecipientType.md) |  | [optional] 
**SCommunicationexternalrecipientName** | **string** | The name of the Communicationexternalrecipient | 

## Methods

### NewCommunicationexternalrecipientRequest

`func NewCommunicationexternalrecipientRequest(sCommunicationexternalrecipientName string, ) *CommunicationexternalrecipientRequest`

NewCommunicationexternalrecipientRequest instantiates a new CommunicationexternalrecipientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationexternalrecipientRequestWithDefaults

`func NewCommunicationexternalrecipientRequestWithDefaults() *CommunicationexternalrecipientRequest`

NewCommunicationexternalrecipientRequestWithDefaults instantiates a new CommunicationexternalrecipientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCommunicationexternalrecipientID

`func (o *CommunicationexternalrecipientRequest) GetPkiCommunicationexternalrecipientID() int32`

GetPkiCommunicationexternalrecipientID returns the PkiCommunicationexternalrecipientID field if non-nil, zero value otherwise.

### GetPkiCommunicationexternalrecipientIDOk

`func (o *CommunicationexternalrecipientRequest) GetPkiCommunicationexternalrecipientIDOk() (*int32, bool)`

GetPkiCommunicationexternalrecipientIDOk returns a tuple with the PkiCommunicationexternalrecipientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCommunicationexternalrecipientID

`func (o *CommunicationexternalrecipientRequest) SetPkiCommunicationexternalrecipientID(v int32)`

SetPkiCommunicationexternalrecipientID sets PkiCommunicationexternalrecipientID field to given value.

### HasPkiCommunicationexternalrecipientID

`func (o *CommunicationexternalrecipientRequest) HasPkiCommunicationexternalrecipientID() bool`

HasPkiCommunicationexternalrecipientID returns a boolean if a field has been set.

### GetSEmailAddress

`func (o *CommunicationexternalrecipientRequest) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *CommunicationexternalrecipientRequest) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *CommunicationexternalrecipientRequest) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *CommunicationexternalrecipientRequest) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetSPhoneE164

`func (o *CommunicationexternalrecipientRequest) GetSPhoneE164() string`

GetSPhoneE164 returns the SPhoneE164 field if non-nil, zero value otherwise.

### GetSPhoneE164Ok

`func (o *CommunicationexternalrecipientRequest) GetSPhoneE164Ok() (*string, bool)`

GetSPhoneE164Ok returns a tuple with the SPhoneE164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneE164

`func (o *CommunicationexternalrecipientRequest) SetSPhoneE164(v string)`

SetSPhoneE164 sets SPhoneE164 field to given value.

### HasSPhoneE164

`func (o *CommunicationexternalrecipientRequest) HasSPhoneE164() bool`

HasSPhoneE164 returns a boolean if a field has been set.

### GetECommunicationexternalrecipientType

`func (o *CommunicationexternalrecipientRequest) GetECommunicationexternalrecipientType() FieldECommunicationexternalrecipientType`

GetECommunicationexternalrecipientType returns the ECommunicationexternalrecipientType field if non-nil, zero value otherwise.

### GetECommunicationexternalrecipientTypeOk

`func (o *CommunicationexternalrecipientRequest) GetECommunicationexternalrecipientTypeOk() (*FieldECommunicationexternalrecipientType, bool)`

GetECommunicationexternalrecipientTypeOk returns a tuple with the ECommunicationexternalrecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationexternalrecipientType

`func (o *CommunicationexternalrecipientRequest) SetECommunicationexternalrecipientType(v FieldECommunicationexternalrecipientType)`

SetECommunicationexternalrecipientType sets ECommunicationexternalrecipientType field to given value.

### HasECommunicationexternalrecipientType

`func (o *CommunicationexternalrecipientRequest) HasECommunicationexternalrecipientType() bool`

HasECommunicationexternalrecipientType returns a boolean if a field has been set.

### GetSCommunicationexternalrecipientName

`func (o *CommunicationexternalrecipientRequest) GetSCommunicationexternalrecipientName() string`

GetSCommunicationexternalrecipientName returns the SCommunicationexternalrecipientName field if non-nil, zero value otherwise.

### GetSCommunicationexternalrecipientNameOk

`func (o *CommunicationexternalrecipientRequest) GetSCommunicationexternalrecipientNameOk() (*string, bool)`

GetSCommunicationexternalrecipientNameOk returns a tuple with the SCommunicationexternalrecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCommunicationexternalrecipientName

`func (o *CommunicationexternalrecipientRequest) SetSCommunicationexternalrecipientName(v string)`

SetSCommunicationexternalrecipientName sets SCommunicationexternalrecipientName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


