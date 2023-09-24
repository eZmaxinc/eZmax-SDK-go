# CorsCreateObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiCorsID** | **[]int32** | An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 

## Methods

### NewCorsCreateObjectV1ResponseMPayload

`func NewCorsCreateObjectV1ResponseMPayload(aPkiCorsID []int32, ) *CorsCreateObjectV1ResponseMPayload`

NewCorsCreateObjectV1ResponseMPayload instantiates a new CorsCreateObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorsCreateObjectV1ResponseMPayloadWithDefaults

`func NewCorsCreateObjectV1ResponseMPayloadWithDefaults() *CorsCreateObjectV1ResponseMPayload`

NewCorsCreateObjectV1ResponseMPayloadWithDefaults instantiates a new CorsCreateObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiCorsID

`func (o *CorsCreateObjectV1ResponseMPayload) GetAPkiCorsID() []int32`

GetAPkiCorsID returns the APkiCorsID field if non-nil, zero value otherwise.

### GetAPkiCorsIDOk

`func (o *CorsCreateObjectV1ResponseMPayload) GetAPkiCorsIDOk() (*[]int32, bool)`

GetAPkiCorsIDOk returns a tuple with the APkiCorsID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiCorsID

`func (o *CorsCreateObjectV1ResponseMPayload) SetAPkiCorsID(v []int32)`

SetAPkiCorsID sets APkiCorsID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


