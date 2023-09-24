# UsergroupCreateObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiUsergroupID** | **[]int32** | An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 

## Methods

### NewUsergroupCreateObjectV1ResponseMPayload

`func NewUsergroupCreateObjectV1ResponseMPayload(aPkiUsergroupID []int32, ) *UsergroupCreateObjectV1ResponseMPayload`

NewUsergroupCreateObjectV1ResponseMPayload instantiates a new UsergroupCreateObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsergroupCreateObjectV1ResponseMPayloadWithDefaults

`func NewUsergroupCreateObjectV1ResponseMPayloadWithDefaults() *UsergroupCreateObjectV1ResponseMPayload`

NewUsergroupCreateObjectV1ResponseMPayloadWithDefaults instantiates a new UsergroupCreateObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiUsergroupID

`func (o *UsergroupCreateObjectV1ResponseMPayload) GetAPkiUsergroupID() []int32`

GetAPkiUsergroupID returns the APkiUsergroupID field if non-nil, zero value otherwise.

### GetAPkiUsergroupIDOk

`func (o *UsergroupCreateObjectV1ResponseMPayload) GetAPkiUsergroupIDOk() (*[]int32, bool)`

GetAPkiUsergroupIDOk returns a tuple with the APkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiUsergroupID

`func (o *UsergroupCreateObjectV1ResponseMPayload) SetAPkiUsergroupID(v []int32)`

SetAPkiUsergroupID sets APkiUsergroupID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


