# EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiEzsigntemplatepackagemembershipID** | **[]int32** | An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 
**BEzsigntemplatepackageNeedvalidation** | **bool** | Whether the Ezsignbulksend was automatically modified and needs a manual validation | 
**BEzsignbulksendNeedvalidation** | **bool** | Whether the Ezsigntemplatepackage was automatically modified and needs a manual validation | 

## Methods

### NewEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload

`func NewEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload(aPkiEzsigntemplatepackagemembershipID []int32, bEzsigntemplatepackageNeedvalidation bool, bEzsignbulksendNeedvalidation bool, ) *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload`

NewEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload instantiates a new EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayloadWithDefaults

`func NewEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayloadWithDefaults() *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload`

NewEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiEzsigntemplatepackagemembershipID

`func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) GetAPkiEzsigntemplatepackagemembershipID() []int32`

GetAPkiEzsigntemplatepackagemembershipID returns the APkiEzsigntemplatepackagemembershipID field if non-nil, zero value otherwise.

### GetAPkiEzsigntemplatepackagemembershipIDOk

`func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) GetAPkiEzsigntemplatepackagemembershipIDOk() (*[]int32, bool)`

GetAPkiEzsigntemplatepackagemembershipIDOk returns a tuple with the APkiEzsigntemplatepackagemembershipID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiEzsigntemplatepackagemembershipID

`func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) SetAPkiEzsigntemplatepackagemembershipID(v []int32)`

SetAPkiEzsigntemplatepackagemembershipID sets APkiEzsigntemplatepackagemembershipID field to given value.


### GetBEzsigntemplatepackageNeedvalidation

`func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) GetBEzsigntemplatepackageNeedvalidation() bool`

GetBEzsigntemplatepackageNeedvalidation returns the BEzsigntemplatepackageNeedvalidation field if non-nil, zero value otherwise.

### GetBEzsigntemplatepackageNeedvalidationOk

`func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) GetBEzsigntemplatepackageNeedvalidationOk() (*bool, bool)`

GetBEzsigntemplatepackageNeedvalidationOk returns a tuple with the BEzsigntemplatepackageNeedvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepackageNeedvalidation

`func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) SetBEzsigntemplatepackageNeedvalidation(v bool)`

SetBEzsigntemplatepackageNeedvalidation sets BEzsigntemplatepackageNeedvalidation field to given value.


### GetBEzsignbulksendNeedvalidation

`func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) GetBEzsignbulksendNeedvalidation() bool`

GetBEzsignbulksendNeedvalidation returns the BEzsignbulksendNeedvalidation field if non-nil, zero value otherwise.

### GetBEzsignbulksendNeedvalidationOk

`func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) GetBEzsignbulksendNeedvalidationOk() (*bool, bool)`

GetBEzsignbulksendNeedvalidationOk returns a tuple with the BEzsignbulksendNeedvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignbulksendNeedvalidation

`func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) SetBEzsignbulksendNeedvalidation(v bool)`

SetBEzsignbulksendNeedvalidation sets BEzsignbulksendNeedvalidation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


