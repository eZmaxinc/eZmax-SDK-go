# EzsigntemplatesignatureRequestCompoundV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BEzsigntemplatesignatureCustomdate** | Pointer to **bool** | Whether the Ezsigntemplatesignature has a custom date format or not. (Only possible when eEzsigntemplatesignatureType is **Name** or **Handwritten**) | [optional] 
**AObjEzsigntemplatesignaturecustomdate** | Pointer to [**[]EzsigntemplatesignaturecustomdateRequestV2**](EzsigntemplatesignaturecustomdateRequestV2.md) | An array of custom date blocks that will be filled at the time of signature.  Can only be used if bEzsigntemplatesignatureCustomdate is true.  Use an empty array if you don&#39;t want to have a date at all. | [optional] 
**AObjEzsigntemplateelementdependency** | Pointer to [**[]EzsigntemplateelementdependencyRequest**](EzsigntemplateelementdependencyRequest.md) |  | [optional] 

## Methods

### NewEzsigntemplatesignatureRequestCompoundV2

`func NewEzsigntemplatesignatureRequestCompoundV2() *EzsigntemplatesignatureRequestCompoundV2`

NewEzsigntemplatesignatureRequestCompoundV2 instantiates a new EzsigntemplatesignatureRequestCompoundV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignatureRequestCompoundV2WithDefaults

`func NewEzsigntemplatesignatureRequestCompoundV2WithDefaults() *EzsigntemplatesignatureRequestCompoundV2`

NewEzsigntemplatesignatureRequestCompoundV2WithDefaults instantiates a new EzsigntemplatesignatureRequestCompoundV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetBEzsigntemplatesignatureCustomdate() bool`

GetBEzsigntemplatesignatureCustomdate returns the BEzsigntemplatesignatureCustomdate field if non-nil, zero value otherwise.

### GetBEzsigntemplatesignatureCustomdateOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetBEzsigntemplatesignatureCustomdateOk() (*bool, bool)`

GetBEzsigntemplatesignatureCustomdateOk returns a tuple with the BEzsigntemplatesignatureCustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetBEzsigntemplatesignatureCustomdate(v bool)`

SetBEzsigntemplatesignatureCustomdate sets BEzsigntemplatesignatureCustomdate field to given value.

### HasBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasBEzsigntemplatesignatureCustomdate() bool`

HasBEzsigntemplatesignatureCustomdate returns a boolean if a field has been set.

### GetAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetAObjEzsigntemplatesignaturecustomdate() []EzsigntemplatesignaturecustomdateRequestCompoundV2`

GetAObjEzsigntemplatesignaturecustomdate returns the AObjEzsigntemplatesignaturecustomdate field if non-nil, zero value otherwise.

### GetAObjEzsigntemplatesignaturecustomdateOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetAObjEzsigntemplatesignaturecustomdateOk() (*[]EzsigntemplatesignaturecustomdateRequestCompoundV2, bool)`

GetAObjEzsigntemplatesignaturecustomdateOk returns a tuple with the AObjEzsigntemplatesignaturecustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetAObjEzsigntemplatesignaturecustomdate(v []EzsigntemplatesignaturecustomdateRequestCompoundV2)`

SetAObjEzsigntemplatesignaturecustomdate sets AObjEzsigntemplatesignaturecustomdate field to given value.

### HasAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasAObjEzsigntemplatesignaturecustomdate() bool`

HasAObjEzsigntemplatesignaturecustomdate returns a boolean if a field has been set.

### GetAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetAObjEzsigntemplateelementdependency() []EzsigntemplateelementdependencyRequestCompound`

GetAObjEzsigntemplateelementdependency returns the AObjEzsigntemplateelementdependency field if non-nil, zero value otherwise.

### GetAObjEzsigntemplateelementdependencyOk

`func (o *EzsigntemplatesignatureRequestCompoundV2) GetAObjEzsigntemplateelementdependencyOk() (*[]EzsigntemplateelementdependencyRequestCompound, bool)`

GetAObjEzsigntemplateelementdependencyOk returns a tuple with the AObjEzsigntemplateelementdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureRequestCompoundV2) SetAObjEzsigntemplateelementdependency(v []EzsigntemplateelementdependencyRequestCompound)`

SetAObjEzsigntemplateelementdependency sets AObjEzsigntemplateelementdependency field to given value.

### HasAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureRequestCompoundV2) HasAObjEzsigntemplateelementdependency() bool`

HasAObjEzsigntemplateelementdependency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


