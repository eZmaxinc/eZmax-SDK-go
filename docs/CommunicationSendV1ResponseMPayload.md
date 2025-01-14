# CommunicationSendV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiCommunicationID** | **[]int32** | An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 

## Methods

### NewCommunicationSendV1ResponseMPayload

`func NewCommunicationSendV1ResponseMPayload(aPkiCommunicationID []int32, ) *CommunicationSendV1ResponseMPayload`

NewCommunicationSendV1ResponseMPayload instantiates a new CommunicationSendV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationSendV1ResponseMPayloadWithDefaults

`func NewCommunicationSendV1ResponseMPayloadWithDefaults() *CommunicationSendV1ResponseMPayload`

NewCommunicationSendV1ResponseMPayloadWithDefaults instantiates a new CommunicationSendV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiCommunicationID

`func (o *CommunicationSendV1ResponseMPayload) GetAPkiCommunicationID() []int32`

GetAPkiCommunicationID returns the APkiCommunicationID field if non-nil, zero value otherwise.

### GetAPkiCommunicationIDOk

`func (o *CommunicationSendV1ResponseMPayload) GetAPkiCommunicationIDOk() (*[]int32, bool)`

GetAPkiCommunicationIDOk returns a tuple with the APkiCommunicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiCommunicationID

`func (o *CommunicationSendV1ResponseMPayload) SetAPkiCommunicationID(v []int32)`

SetAPkiCommunicationID sets APkiCommunicationID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


