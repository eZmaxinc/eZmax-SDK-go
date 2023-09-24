# PermissionCreateObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiPermissionID** | **[]int32** | An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 

## Methods

### NewPermissionCreateObjectV1ResponseMPayload

`func NewPermissionCreateObjectV1ResponseMPayload(aPkiPermissionID []int32, ) *PermissionCreateObjectV1ResponseMPayload`

NewPermissionCreateObjectV1ResponseMPayload instantiates a new PermissionCreateObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionCreateObjectV1ResponseMPayloadWithDefaults

`func NewPermissionCreateObjectV1ResponseMPayloadWithDefaults() *PermissionCreateObjectV1ResponseMPayload`

NewPermissionCreateObjectV1ResponseMPayloadWithDefaults instantiates a new PermissionCreateObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiPermissionID

`func (o *PermissionCreateObjectV1ResponseMPayload) GetAPkiPermissionID() []int32`

GetAPkiPermissionID returns the APkiPermissionID field if non-nil, zero value otherwise.

### GetAPkiPermissionIDOk

`func (o *PermissionCreateObjectV1ResponseMPayload) GetAPkiPermissionIDOk() (*[]int32, bool)`

GetAPkiPermissionIDOk returns a tuple with the APkiPermissionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiPermissionID

`func (o *PermissionCreateObjectV1ResponseMPayload) SetAPkiPermissionID(v []int32)`

SetAPkiPermissionID sets APkiPermissionID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


