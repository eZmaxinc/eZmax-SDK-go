# EzsigntemplatesignerCreateObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiEzsigntemplatesignerID** | **[]int32** | An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 
**BEzsigntemplatepackageNeedvalidation** | **bool** | Whether the Ezsignbulksend was automatically modified and needs a manual validation | 
**BEzsignbulksendNeedvalidation** | **bool** | Whether the Ezsigntemplatepackage was automatically modified and needs a manual validation | 

## Methods

### NewEzsigntemplatesignerCreateObjectV1ResponseMPayload

`func NewEzsigntemplatesignerCreateObjectV1ResponseMPayload(aPkiEzsigntemplatesignerID []int32, bEzsigntemplatepackageNeedvalidation bool, bEzsignbulksendNeedvalidation bool, ) *EzsigntemplatesignerCreateObjectV1ResponseMPayload`

NewEzsigntemplatesignerCreateObjectV1ResponseMPayload instantiates a new EzsigntemplatesignerCreateObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignerCreateObjectV1ResponseMPayloadWithDefaults

`func NewEzsigntemplatesignerCreateObjectV1ResponseMPayloadWithDefaults() *EzsigntemplatesignerCreateObjectV1ResponseMPayload`

NewEzsigntemplatesignerCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatesignerCreateObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) GetAPkiEzsigntemplatesignerID() []int32`

GetAPkiEzsigntemplatesignerID returns the APkiEzsigntemplatesignerID field if non-nil, zero value otherwise.

### GetAPkiEzsigntemplatesignerIDOk

`func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) GetAPkiEzsigntemplatesignerIDOk() (*[]int32, bool)`

GetAPkiEzsigntemplatesignerIDOk returns a tuple with the APkiEzsigntemplatesignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) SetAPkiEzsigntemplatesignerID(v []int32)`

SetAPkiEzsigntemplatesignerID sets APkiEzsigntemplatesignerID field to given value.


### GetBEzsigntemplatepackageNeedvalidation

`func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) GetBEzsigntemplatepackageNeedvalidation() bool`

GetBEzsigntemplatepackageNeedvalidation returns the BEzsigntemplatepackageNeedvalidation field if non-nil, zero value otherwise.

### GetBEzsigntemplatepackageNeedvalidationOk

`func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) GetBEzsigntemplatepackageNeedvalidationOk() (*bool, bool)`

GetBEzsigntemplatepackageNeedvalidationOk returns a tuple with the BEzsigntemplatepackageNeedvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepackageNeedvalidation

`func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) SetBEzsigntemplatepackageNeedvalidation(v bool)`

SetBEzsigntemplatepackageNeedvalidation sets BEzsigntemplatepackageNeedvalidation field to given value.


### GetBEzsignbulksendNeedvalidation

`func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) GetBEzsignbulksendNeedvalidation() bool`

GetBEzsignbulksendNeedvalidation returns the BEzsignbulksendNeedvalidation field if non-nil, zero value otherwise.

### GetBEzsignbulksendNeedvalidationOk

`func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) GetBEzsignbulksendNeedvalidationOk() (*bool, bool)`

GetBEzsignbulksendNeedvalidationOk returns a tuple with the BEzsignbulksendNeedvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignbulksendNeedvalidation

`func (o *EzsigntemplatesignerCreateObjectV1ResponseMPayload) SetBEzsignbulksendNeedvalidation(v bool)`

SetBEzsignbulksendNeedvalidation sets BEzsignbulksendNeedvalidation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


