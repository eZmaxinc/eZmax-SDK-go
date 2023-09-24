# SignatureCreateObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiSignatureID** | **[]int32** | An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 

## Methods

### NewSignatureCreateObjectV1ResponseMPayload

`func NewSignatureCreateObjectV1ResponseMPayload(aPkiSignatureID []int32, ) *SignatureCreateObjectV1ResponseMPayload`

NewSignatureCreateObjectV1ResponseMPayload instantiates a new SignatureCreateObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureCreateObjectV1ResponseMPayloadWithDefaults

`func NewSignatureCreateObjectV1ResponseMPayloadWithDefaults() *SignatureCreateObjectV1ResponseMPayload`

NewSignatureCreateObjectV1ResponseMPayloadWithDefaults instantiates a new SignatureCreateObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiSignatureID

`func (o *SignatureCreateObjectV1ResponseMPayload) GetAPkiSignatureID() []int32`

GetAPkiSignatureID returns the APkiSignatureID field if non-nil, zero value otherwise.

### GetAPkiSignatureIDOk

`func (o *SignatureCreateObjectV1ResponseMPayload) GetAPkiSignatureIDOk() (*[]int32, bool)`

GetAPkiSignatureIDOk returns a tuple with the APkiSignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiSignatureID

`func (o *SignatureCreateObjectV1ResponseMPayload) SetAPkiSignatureID(v []int32)`

SetAPkiSignatureID sets APkiSignatureID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


