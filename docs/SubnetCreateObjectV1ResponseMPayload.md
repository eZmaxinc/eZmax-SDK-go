# SubnetCreateObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiSubnetID** | **[]int32** | An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 

## Methods

### NewSubnetCreateObjectV1ResponseMPayload

`func NewSubnetCreateObjectV1ResponseMPayload(aPkiSubnetID []int32, ) *SubnetCreateObjectV1ResponseMPayload`

NewSubnetCreateObjectV1ResponseMPayload instantiates a new SubnetCreateObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetCreateObjectV1ResponseMPayloadWithDefaults

`func NewSubnetCreateObjectV1ResponseMPayloadWithDefaults() *SubnetCreateObjectV1ResponseMPayload`

NewSubnetCreateObjectV1ResponseMPayloadWithDefaults instantiates a new SubnetCreateObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiSubnetID

`func (o *SubnetCreateObjectV1ResponseMPayload) GetAPkiSubnetID() []int32`

GetAPkiSubnetID returns the APkiSubnetID field if non-nil, zero value otherwise.

### GetAPkiSubnetIDOk

`func (o *SubnetCreateObjectV1ResponseMPayload) GetAPkiSubnetIDOk() (*[]int32, bool)`

GetAPkiSubnetIDOk returns a tuple with the APkiSubnetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiSubnetID

`func (o *SubnetCreateObjectV1ResponseMPayload) SetAPkiSubnetID(v []int32)`

SetAPkiSubnetID sets APkiSubnetID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


