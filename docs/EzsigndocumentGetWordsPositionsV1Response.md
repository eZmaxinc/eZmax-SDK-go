# EzsigndocumentGetWordsPositionsV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MPayload** | [**[]CustomWordPositionWordResponse**](CustomWordPositionWordResponse.md) | Payload for POST /1/object/ezsigndocument/{pkiEzsigndocumentID}/getWordsPositions | 

## Methods

### NewEzsigndocumentGetWordsPositionsV1Response

`func NewEzsigndocumentGetWordsPositionsV1Response(mPayload []CustomWordPositionWordResponse, ) *EzsigndocumentGetWordsPositionsV1Response`

NewEzsigndocumentGetWordsPositionsV1Response instantiates a new EzsigndocumentGetWordsPositionsV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentGetWordsPositionsV1ResponseWithDefaults

`func NewEzsigndocumentGetWordsPositionsV1ResponseWithDefaults() *EzsigndocumentGetWordsPositionsV1Response`

NewEzsigndocumentGetWordsPositionsV1ResponseWithDefaults instantiates a new EzsigndocumentGetWordsPositionsV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMPayload

`func (o *EzsigndocumentGetWordsPositionsV1Response) GetMPayload() []CustomWordPositionWordResponse`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *EzsigndocumentGetWordsPositionsV1Response) GetMPayloadOk() (*[]CustomWordPositionWordResponse, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *EzsigndocumentGetWordsPositionsV1Response) SetMPayload(v []CustomWordPositionWordResponse)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


