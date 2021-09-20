# EzsigndocumentGetWordsPositionsV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ASWords** | [**[]WordPositionResponse**](WordPositionResponse.md) | An array of words with an array of pages and positions X,Y  They are returned with the sames words that was sent in the request. | 

## Methods

### NewEzsigndocumentGetWordsPositionsV1ResponseMPayload

`func NewEzsigndocumentGetWordsPositionsV1ResponseMPayload(aSWords []WordPositionResponse, ) *EzsigndocumentGetWordsPositionsV1ResponseMPayload`

NewEzsigndocumentGetWordsPositionsV1ResponseMPayload instantiates a new EzsigndocumentGetWordsPositionsV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentGetWordsPositionsV1ResponseMPayloadWithDefaults

`func NewEzsigndocumentGetWordsPositionsV1ResponseMPayloadWithDefaults() *EzsigndocumentGetWordsPositionsV1ResponseMPayload`

NewEzsigndocumentGetWordsPositionsV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentGetWordsPositionsV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetASWords

`func (o *EzsigndocumentGetWordsPositionsV1ResponseMPayload) GetASWords() []WordPositionResponse`

GetASWords returns the ASWords field if non-nil, zero value otherwise.

### GetASWordsOk

`func (o *EzsigndocumentGetWordsPositionsV1ResponseMPayload) GetASWordsOk() (*[]WordPositionResponse, bool)`

GetASWordsOk returns a tuple with the ASWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASWords

`func (o *EzsigndocumentGetWordsPositionsV1ResponseMPayload) SetASWords(v []WordPositionResponse)`

SetASWords sets ASWords field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


