# DomainCreateObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiDomainID** | **[]int32** | An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 

## Methods

### NewDomainCreateObjectV1ResponseMPayload

`func NewDomainCreateObjectV1ResponseMPayload(aPkiDomainID []int32, ) *DomainCreateObjectV1ResponseMPayload`

NewDomainCreateObjectV1ResponseMPayload instantiates a new DomainCreateObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainCreateObjectV1ResponseMPayloadWithDefaults

`func NewDomainCreateObjectV1ResponseMPayloadWithDefaults() *DomainCreateObjectV1ResponseMPayload`

NewDomainCreateObjectV1ResponseMPayloadWithDefaults instantiates a new DomainCreateObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiDomainID

`func (o *DomainCreateObjectV1ResponseMPayload) GetAPkiDomainID() []int32`

GetAPkiDomainID returns the APkiDomainID field if non-nil, zero value otherwise.

### GetAPkiDomainIDOk

`func (o *DomainCreateObjectV1ResponseMPayload) GetAPkiDomainIDOk() (*[]int32, bool)`

GetAPkiDomainIDOk returns a tuple with the APkiDomainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiDomainID

`func (o *DomainCreateObjectV1ResponseMPayload) SetAPkiDomainID(v []int32)`

SetAPkiDomainID sets APkiDomainID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


