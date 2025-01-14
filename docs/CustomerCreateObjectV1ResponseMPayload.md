# CustomerCreateObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiCustomerID** | **[]int32** | An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 

## Methods

### NewCustomerCreateObjectV1ResponseMPayload

`func NewCustomerCreateObjectV1ResponseMPayload(aPkiCustomerID []int32, ) *CustomerCreateObjectV1ResponseMPayload`

NewCustomerCreateObjectV1ResponseMPayload instantiates a new CustomerCreateObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerCreateObjectV1ResponseMPayloadWithDefaults

`func NewCustomerCreateObjectV1ResponseMPayloadWithDefaults() *CustomerCreateObjectV1ResponseMPayload`

NewCustomerCreateObjectV1ResponseMPayloadWithDefaults instantiates a new CustomerCreateObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiCustomerID

`func (o *CustomerCreateObjectV1ResponseMPayload) GetAPkiCustomerID() []int32`

GetAPkiCustomerID returns the APkiCustomerID field if non-nil, zero value otherwise.

### GetAPkiCustomerIDOk

`func (o *CustomerCreateObjectV1ResponseMPayload) GetAPkiCustomerIDOk() (*[]int32, bool)`

GetAPkiCustomerIDOk returns a tuple with the APkiCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiCustomerID

`func (o *CustomerCreateObjectV1ResponseMPayload) SetAPkiCustomerID(v []int32)`

SetAPkiCustomerID sets APkiCustomerID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


