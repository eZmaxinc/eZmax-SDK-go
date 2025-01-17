# EzsigntemplatesignatureResponseCompoundV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BEzsigntemplatesignatureCustomdate** | Pointer to **bool** | Whether the Ezsigntemplatesignature has a custom date format or not. (Only possible when eEzsigntemplatesignatureType is **Name** or **Handwritten**) | [optional] 
**AObjEzsigntemplatesignaturecustomdate** | Pointer to [**[]EzsigntemplatesignaturecustomdateResponseV2**](EzsigntemplatesignaturecustomdateResponseV2.md) | An array of custom date blocks that will be filled at the time of signature.  Can only be used if bEzsigntemplatesignatureCustomdate is true.  Use an empty array if you don&#39;t want to have a date at all. | [optional] 
**AObjEzsigntemplateelementdependency** | Pointer to [**[]EzsigntemplateelementdependencyResponse**](EzsigntemplateelementdependencyResponse.md) |  | [optional] 

## Methods

### NewEzsigntemplatesignatureResponseCompoundV3

`func NewEzsigntemplatesignatureResponseCompoundV3() *EzsigntemplatesignatureResponseCompoundV3`

NewEzsigntemplatesignatureResponseCompoundV3 instantiates a new EzsigntemplatesignatureResponseCompoundV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignatureResponseCompoundV3WithDefaults

`func NewEzsigntemplatesignatureResponseCompoundV3WithDefaults() *EzsigntemplatesignatureResponseCompoundV3`

NewEzsigntemplatesignatureResponseCompoundV3WithDefaults instantiates a new EzsigntemplatesignatureResponseCompoundV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetBEzsigntemplatesignatureCustomdate() bool`

GetBEzsigntemplatesignatureCustomdate returns the BEzsigntemplatesignatureCustomdate field if non-nil, zero value otherwise.

### GetBEzsigntemplatesignatureCustomdateOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetBEzsigntemplatesignatureCustomdateOk() (*bool, bool)`

GetBEzsigntemplatesignatureCustomdateOk returns a tuple with the BEzsigntemplatesignatureCustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetBEzsigntemplatesignatureCustomdate(v bool)`

SetBEzsigntemplatesignatureCustomdate sets BEzsigntemplatesignatureCustomdate field to given value.

### HasBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasBEzsigntemplatesignatureCustomdate() bool`

HasBEzsigntemplatesignatureCustomdate returns a boolean if a field has been set.

### GetAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetAObjEzsigntemplatesignaturecustomdate() []EzsigntemplatesignaturecustomdateResponseCompoundV2`

GetAObjEzsigntemplatesignaturecustomdate returns the AObjEzsigntemplatesignaturecustomdate field if non-nil, zero value otherwise.

### GetAObjEzsigntemplatesignaturecustomdateOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetAObjEzsigntemplatesignaturecustomdateOk() (*[]EzsigntemplatesignaturecustomdateResponseCompoundV2, bool)`

GetAObjEzsigntemplatesignaturecustomdateOk returns a tuple with the AObjEzsigntemplatesignaturecustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetAObjEzsigntemplatesignaturecustomdate(v []EzsigntemplatesignaturecustomdateResponseCompoundV2)`

SetAObjEzsigntemplatesignaturecustomdate sets AObjEzsigntemplatesignaturecustomdate field to given value.

### HasAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasAObjEzsigntemplatesignaturecustomdate() bool`

HasAObjEzsigntemplatesignaturecustomdate returns a boolean if a field has been set.

### GetAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetAObjEzsigntemplateelementdependency() []EzsigntemplateelementdependencyResponseCompound`

GetAObjEzsigntemplateelementdependency returns the AObjEzsigntemplateelementdependency field if non-nil, zero value otherwise.

### GetAObjEzsigntemplateelementdependencyOk

`func (o *EzsigntemplatesignatureResponseCompoundV3) GetAObjEzsigntemplateelementdependencyOk() (*[]EzsigntemplateelementdependencyResponseCompound, bool)`

GetAObjEzsigntemplateelementdependencyOk returns a tuple with the AObjEzsigntemplateelementdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureResponseCompoundV3) SetAObjEzsigntemplateelementdependency(v []EzsigntemplateelementdependencyResponseCompound)`

SetAObjEzsigntemplateelementdependency sets AObjEzsigntemplateelementdependency field to given value.

### HasAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureResponseCompoundV3) HasAObjEzsigntemplateelementdependency() bool`

HasAObjEzsigntemplateelementdependency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


