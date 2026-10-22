# DocumenttypeCreateObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiDocumenttypeID** | **[]int32** | An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 

## Methods

### NewDocumenttypeCreateObjectV1ResponseMPayload

`func NewDocumenttypeCreateObjectV1ResponseMPayload(aPkiDocumenttypeID []int32, ) *DocumenttypeCreateObjectV1ResponseMPayload`

NewDocumenttypeCreateObjectV1ResponseMPayload instantiates a new DocumenttypeCreateObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumenttypeCreateObjectV1ResponseMPayloadWithDefaults

`func NewDocumenttypeCreateObjectV1ResponseMPayloadWithDefaults() *DocumenttypeCreateObjectV1ResponseMPayload`

NewDocumenttypeCreateObjectV1ResponseMPayloadWithDefaults instantiates a new DocumenttypeCreateObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiDocumenttypeID

`func (o *DocumenttypeCreateObjectV1ResponseMPayload) GetAPkiDocumenttypeID() []int32`

GetAPkiDocumenttypeID returns the APkiDocumenttypeID field if non-nil, zero value otherwise.

### GetAPkiDocumenttypeIDOk

`func (o *DocumenttypeCreateObjectV1ResponseMPayload) GetAPkiDocumenttypeIDOk() (*[]int32, bool)`

GetAPkiDocumenttypeIDOk returns a tuple with the APkiDocumenttypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiDocumenttypeID

`func (o *DocumenttypeCreateObjectV1ResponseMPayload) SetAPkiDocumenttypeID(v []int32)`

SetAPkiDocumenttypeID sets APkiDocumenttypeID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


