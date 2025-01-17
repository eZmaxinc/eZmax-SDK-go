# EzsigndocumentGetObjectV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MPayload** | [**EzsigndocumentResponseCompound**](EzsigndocumentResponseCompound.md) | Payload for GET /1/object/ezsigndocument/{pkiEzsigndocumentID} | 

## Methods

### NewEzsigndocumentGetObjectV1Response

`func NewEzsigndocumentGetObjectV1Response(mPayload EzsigndocumentResponseCompound, ) *EzsigndocumentGetObjectV1Response`

NewEzsigndocumentGetObjectV1Response instantiates a new EzsigndocumentGetObjectV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentGetObjectV1ResponseWithDefaults

`func NewEzsigndocumentGetObjectV1ResponseWithDefaults() *EzsigndocumentGetObjectV1Response`

NewEzsigndocumentGetObjectV1ResponseWithDefaults instantiates a new EzsigndocumentGetObjectV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMPayload

`func (o *EzsigndocumentGetObjectV1Response) GetMPayload() EzsigndocumentResponseCompound`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *EzsigndocumentGetObjectV1Response) GetMPayloadOk() (*EzsigndocumentResponseCompound, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *EzsigndocumentGetObjectV1Response) SetMPayload(v EzsigndocumentResponseCompound)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


