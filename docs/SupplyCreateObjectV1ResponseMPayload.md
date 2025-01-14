# SupplyCreateObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiSupplyID** | **[]int32** | An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 

## Methods

### NewSupplyCreateObjectV1ResponseMPayload

`func NewSupplyCreateObjectV1ResponseMPayload(aPkiSupplyID []int32, ) *SupplyCreateObjectV1ResponseMPayload`

NewSupplyCreateObjectV1ResponseMPayload instantiates a new SupplyCreateObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyCreateObjectV1ResponseMPayloadWithDefaults

`func NewSupplyCreateObjectV1ResponseMPayloadWithDefaults() *SupplyCreateObjectV1ResponseMPayload`

NewSupplyCreateObjectV1ResponseMPayloadWithDefaults instantiates a new SupplyCreateObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiSupplyID

`func (o *SupplyCreateObjectV1ResponseMPayload) GetAPkiSupplyID() []int32`

GetAPkiSupplyID returns the APkiSupplyID field if non-nil, zero value otherwise.

### GetAPkiSupplyIDOk

`func (o *SupplyCreateObjectV1ResponseMPayload) GetAPkiSupplyIDOk() (*[]int32, bool)`

GetAPkiSupplyIDOk returns a tuple with the APkiSupplyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiSupplyID

`func (o *SupplyCreateObjectV1ResponseMPayload) SetAPkiSupplyID(v []int32)`

SetAPkiSupplyID sets APkiSupplyID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


