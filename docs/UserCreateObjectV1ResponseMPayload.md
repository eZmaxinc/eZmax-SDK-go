# UserCreateObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiUserID** | **[]int32** | An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 

## Methods

### NewUserCreateObjectV1ResponseMPayload

`func NewUserCreateObjectV1ResponseMPayload(aPkiUserID []int32, ) *UserCreateObjectV1ResponseMPayload`

NewUserCreateObjectV1ResponseMPayload instantiates a new UserCreateObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateObjectV1ResponseMPayloadWithDefaults

`func NewUserCreateObjectV1ResponseMPayloadWithDefaults() *UserCreateObjectV1ResponseMPayload`

NewUserCreateObjectV1ResponseMPayloadWithDefaults instantiates a new UserCreateObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiUserID

`func (o *UserCreateObjectV1ResponseMPayload) GetAPkiUserID() []int32`

GetAPkiUserID returns the APkiUserID field if non-nil, zero value otherwise.

### GetAPkiUserIDOk

`func (o *UserCreateObjectV1ResponseMPayload) GetAPkiUserIDOk() (*[]int32, bool)`

GetAPkiUserIDOk returns a tuple with the APkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiUserID

`func (o *UserCreateObjectV1ResponseMPayload) SetAPkiUserID(v []int32)`

SetAPkiUserID sets APkiUserID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


