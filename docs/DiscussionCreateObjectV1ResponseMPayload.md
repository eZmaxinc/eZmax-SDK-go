# DiscussionCreateObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiDiscussionID** | **[]int32** | An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 

## Methods

### NewDiscussionCreateObjectV1ResponseMPayload

`func NewDiscussionCreateObjectV1ResponseMPayload(aPkiDiscussionID []int32, ) *DiscussionCreateObjectV1ResponseMPayload`

NewDiscussionCreateObjectV1ResponseMPayload instantiates a new DiscussionCreateObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscussionCreateObjectV1ResponseMPayloadWithDefaults

`func NewDiscussionCreateObjectV1ResponseMPayloadWithDefaults() *DiscussionCreateObjectV1ResponseMPayload`

NewDiscussionCreateObjectV1ResponseMPayloadWithDefaults instantiates a new DiscussionCreateObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiDiscussionID

`func (o *DiscussionCreateObjectV1ResponseMPayload) GetAPkiDiscussionID() []int32`

GetAPkiDiscussionID returns the APkiDiscussionID field if non-nil, zero value otherwise.

### GetAPkiDiscussionIDOk

`func (o *DiscussionCreateObjectV1ResponseMPayload) GetAPkiDiscussionIDOk() (*[]int32, bool)`

GetAPkiDiscussionIDOk returns a tuple with the APkiDiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiDiscussionID

`func (o *DiscussionCreateObjectV1ResponseMPayload) SetAPkiDiscussionID(v []int32)`

SetAPkiDiscussionID sets APkiDiscussionID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


