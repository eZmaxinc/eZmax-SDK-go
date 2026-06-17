# UserImpersonateV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IExpirationMinutes** | **int32** | The number of minute before key is no longer active | 

## Methods

### NewUserImpersonateV1Request

`func NewUserImpersonateV1Request(iExpirationMinutes int32, ) *UserImpersonateV1Request`

NewUserImpersonateV1Request instantiates a new UserImpersonateV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserImpersonateV1RequestWithDefaults

`func NewUserImpersonateV1RequestWithDefaults() *UserImpersonateV1Request`

NewUserImpersonateV1RequestWithDefaults instantiates a new UserImpersonateV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIExpirationMinutes

`func (o *UserImpersonateV1Request) GetIExpirationMinutes() int32`

GetIExpirationMinutes returns the IExpirationMinutes field if non-nil, zero value otherwise.

### GetIExpirationMinutesOk

`func (o *UserImpersonateV1Request) GetIExpirationMinutesOk() (*int32, bool)`

GetIExpirationMinutesOk returns a tuple with the IExpirationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIExpirationMinutes

`func (o *UserImpersonateV1Request) SetIExpirationMinutes(v int32)`

SetIExpirationMinutes sets IExpirationMinutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


