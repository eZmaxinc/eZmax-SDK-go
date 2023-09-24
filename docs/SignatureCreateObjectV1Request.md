# SignatureCreateObjectV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AObjSignature** | [**[]SignatureRequestCompound**](SignatureRequestCompound.md) |  | 

## Methods

### NewSignatureCreateObjectV1Request

`func NewSignatureCreateObjectV1Request(aObjSignature []SignatureRequestCompound, ) *SignatureCreateObjectV1Request`

NewSignatureCreateObjectV1Request instantiates a new SignatureCreateObjectV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureCreateObjectV1RequestWithDefaults

`func NewSignatureCreateObjectV1RequestWithDefaults() *SignatureCreateObjectV1Request`

NewSignatureCreateObjectV1RequestWithDefaults instantiates a new SignatureCreateObjectV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAObjSignature

`func (o *SignatureCreateObjectV1Request) GetAObjSignature() []SignatureRequestCompound`

GetAObjSignature returns the AObjSignature field if non-nil, zero value otherwise.

### GetAObjSignatureOk

`func (o *SignatureCreateObjectV1Request) GetAObjSignatureOk() (*[]SignatureRequestCompound, bool)`

GetAObjSignatureOk returns a tuple with the AObjSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjSignature

`func (o *SignatureCreateObjectV1Request) SetAObjSignature(v []SignatureRequestCompound)`

SetAObjSignature sets AObjSignature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


