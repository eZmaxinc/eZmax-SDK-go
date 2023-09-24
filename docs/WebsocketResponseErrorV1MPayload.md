# WebsocketResponseErrorV1MPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SErrorMessage** | **string** | The message giving details about the error | 
**EErrorCode** | [**FieldEErrorCode**](FieldEErrorCode.md) |  | 

## Methods

### NewWebsocketResponseErrorV1MPayload

`func NewWebsocketResponseErrorV1MPayload(sErrorMessage string, eErrorCode FieldEErrorCode, ) *WebsocketResponseErrorV1MPayload`

NewWebsocketResponseErrorV1MPayload instantiates a new WebsocketResponseErrorV1MPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsocketResponseErrorV1MPayloadWithDefaults

`func NewWebsocketResponseErrorV1MPayloadWithDefaults() *WebsocketResponseErrorV1MPayload`

NewWebsocketResponseErrorV1MPayloadWithDefaults instantiates a new WebsocketResponseErrorV1MPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSErrorMessage

`func (o *WebsocketResponseErrorV1MPayload) GetSErrorMessage() string`

GetSErrorMessage returns the SErrorMessage field if non-nil, zero value otherwise.

### GetSErrorMessageOk

`func (o *WebsocketResponseErrorV1MPayload) GetSErrorMessageOk() (*string, bool)`

GetSErrorMessageOk returns a tuple with the SErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSErrorMessage

`func (o *WebsocketResponseErrorV1MPayload) SetSErrorMessage(v string)`

SetSErrorMessage sets SErrorMessage field to given value.


### GetEErrorCode

`func (o *WebsocketResponseErrorV1MPayload) GetEErrorCode() FieldEErrorCode`

GetEErrorCode returns the EErrorCode field if non-nil, zero value otherwise.

### GetEErrorCodeOk

`func (o *WebsocketResponseErrorV1MPayload) GetEErrorCodeOk() (*FieldEErrorCode, bool)`

GetEErrorCodeOk returns a tuple with the EErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEErrorCode

`func (o *WebsocketResponseErrorV1MPayload) SetEErrorCode(v FieldEErrorCode)`

SetEErrorCode sets EErrorCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


