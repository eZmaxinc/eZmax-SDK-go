# WebsocketResponseErrorV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EWebsocketMessagetype** | **string** | The Type of message | 
**SWebsocketChannel** | **string** | The Channel on which to route the websocket message | 
**MPayload** | [**WebsocketResponseErrorV1MPayload**](WebsocketResponseErrorV1MPayload.md) |  | 

## Methods

### NewWebsocketResponseErrorV1

`func NewWebsocketResponseErrorV1(eWebsocketMessagetype string, sWebsocketChannel string, mPayload WebsocketResponseErrorV1MPayload, ) *WebsocketResponseErrorV1`

NewWebsocketResponseErrorV1 instantiates a new WebsocketResponseErrorV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsocketResponseErrorV1WithDefaults

`func NewWebsocketResponseErrorV1WithDefaults() *WebsocketResponseErrorV1`

NewWebsocketResponseErrorV1WithDefaults instantiates a new WebsocketResponseErrorV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEWebsocketMessagetype

`func (o *WebsocketResponseErrorV1) GetEWebsocketMessagetype() string`

GetEWebsocketMessagetype returns the EWebsocketMessagetype field if non-nil, zero value otherwise.

### GetEWebsocketMessagetypeOk

`func (o *WebsocketResponseErrorV1) GetEWebsocketMessagetypeOk() (*string, bool)`

GetEWebsocketMessagetypeOk returns a tuple with the EWebsocketMessagetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebsocketMessagetype

`func (o *WebsocketResponseErrorV1) SetEWebsocketMessagetype(v string)`

SetEWebsocketMessagetype sets EWebsocketMessagetype field to given value.


### GetSWebsocketChannel

`func (o *WebsocketResponseErrorV1) GetSWebsocketChannel() string`

GetSWebsocketChannel returns the SWebsocketChannel field if non-nil, zero value otherwise.

### GetSWebsocketChannelOk

`func (o *WebsocketResponseErrorV1) GetSWebsocketChannelOk() (*string, bool)`

GetSWebsocketChannelOk returns a tuple with the SWebsocketChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebsocketChannel

`func (o *WebsocketResponseErrorV1) SetSWebsocketChannel(v string)`

SetSWebsocketChannel sets SWebsocketChannel field to given value.


### GetMPayload

`func (o *WebsocketResponseErrorV1) GetMPayload() WebsocketResponseErrorV1MPayload`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *WebsocketResponseErrorV1) GetMPayloadOk() (*WebsocketResponseErrorV1MPayload, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *WebsocketResponseErrorV1) SetMPayload(v WebsocketResponseErrorV1MPayload)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


