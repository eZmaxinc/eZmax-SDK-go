# WebsocketResponseInformationV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EWebsocketMessagetype** | **string** | The Type of message | 
**SWebsocketChannel** | **string** | The Channel on which to route the websocket message | 
**MPayload** | [**WebsocketResponseInformationV1MPayload**](WebsocketResponseInformationV1MPayload.md) |  | 

## Methods

### NewWebsocketResponseInformationV1

`func NewWebsocketResponseInformationV1(eWebsocketMessagetype string, sWebsocketChannel string, mPayload WebsocketResponseInformationV1MPayload, ) *WebsocketResponseInformationV1`

NewWebsocketResponseInformationV1 instantiates a new WebsocketResponseInformationV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsocketResponseInformationV1WithDefaults

`func NewWebsocketResponseInformationV1WithDefaults() *WebsocketResponseInformationV1`

NewWebsocketResponseInformationV1WithDefaults instantiates a new WebsocketResponseInformationV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEWebsocketMessagetype

`func (o *WebsocketResponseInformationV1) GetEWebsocketMessagetype() string`

GetEWebsocketMessagetype returns the EWebsocketMessagetype field if non-nil, zero value otherwise.

### GetEWebsocketMessagetypeOk

`func (o *WebsocketResponseInformationV1) GetEWebsocketMessagetypeOk() (*string, bool)`

GetEWebsocketMessagetypeOk returns a tuple with the EWebsocketMessagetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebsocketMessagetype

`func (o *WebsocketResponseInformationV1) SetEWebsocketMessagetype(v string)`

SetEWebsocketMessagetype sets EWebsocketMessagetype field to given value.


### GetSWebsocketChannel

`func (o *WebsocketResponseInformationV1) GetSWebsocketChannel() string`

GetSWebsocketChannel returns the SWebsocketChannel field if non-nil, zero value otherwise.

### GetSWebsocketChannelOk

`func (o *WebsocketResponseInformationV1) GetSWebsocketChannelOk() (*string, bool)`

GetSWebsocketChannelOk returns a tuple with the SWebsocketChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebsocketChannel

`func (o *WebsocketResponseInformationV1) SetSWebsocketChannel(v string)`

SetSWebsocketChannel sets SWebsocketChannel field to given value.


### GetMPayload

`func (o *WebsocketResponseInformationV1) GetMPayload() WebsocketResponseInformationV1MPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *WebsocketResponseInformationV1) GetMPayloadOk() (*WebsocketResponseInformationV1MPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *WebsocketResponseInformationV1) SetMPayload(v WebsocketResponseInformationV1MPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


