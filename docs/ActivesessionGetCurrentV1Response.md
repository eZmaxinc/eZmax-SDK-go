# ActivesessionGetCurrentV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MPayload** | [**ActivesessionResponseCompound**](ActivesessionResponseCompound.md) | Payload for GET /1/object/activesession/getCurrent | 

## Methods

### NewActivesessionGetCurrentV1Response

`func NewActivesessionGetCurrentV1Response(mPayload ActivesessionResponseCompound, ) *ActivesessionGetCurrentV1Response`

NewActivesessionGetCurrentV1Response instantiates a new ActivesessionGetCurrentV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivesessionGetCurrentV1ResponseWithDefaults

`func NewActivesessionGetCurrentV1ResponseWithDefaults() *ActivesessionGetCurrentV1Response`

NewActivesessionGetCurrentV1ResponseWithDefaults instantiates a new ActivesessionGetCurrentV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMPayload

`func (o *ActivesessionGetCurrentV1Response) GetMPayload() ActivesessionResponseCompound`

GetMPayload returns the MPayload field if non-nil, zero value otherwise.

### GetMPayloadOk

`func (o *ActivesessionGetCurrentV1Response) GetMPayloadOk() (*ActivesessionResponseCompound, bool)`

GetMPayloadOk returns a tuple with the MPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMPayload

`func (o *ActivesessionGetCurrentV1Response) SetMPayload(v ActivesessionResponseCompound)`

SetMPayload sets MPayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


