# EzsigntemplatesignatureResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BEzsigntemplatesignatureCustomdate** | Pointer to **bool** | Whether the Ezsigntemplatesignature has a custom date format or not. (Only possible when eEzsigntemplatesignatureType is **Name** or **Handwritten**) | [optional] 
**AObjEzsigntemplatesignaturecustomdate** | Pointer to [**[]EzsigntemplatesignaturecustomdateResponse**](EzsigntemplatesignaturecustomdateResponse.md) | An array of custom date blocks that will be filled at the time of signature.  Can only be used if bEzsigntemplatesignatureCustomdate is true.  Use an empty array if you don&#39;t want to have a date at all. | [optional] 
**AObjEzsigntemplateelementdependency** | Pointer to [**[]EzsigntemplateelementdependencyResponse**](EzsigntemplateelementdependencyResponse.md) |  | [optional] 

## Methods

### NewEzsigntemplatesignatureResponseCompound

`func NewEzsigntemplatesignatureResponseCompound() *EzsigntemplatesignatureResponseCompound`

NewEzsigntemplatesignatureResponseCompound instantiates a new EzsigntemplatesignatureResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignatureResponseCompoundWithDefaults

`func NewEzsigntemplatesignatureResponseCompoundWithDefaults() *EzsigntemplatesignatureResponseCompound`

NewEzsigntemplatesignatureResponseCompoundWithDefaults instantiates a new EzsigntemplatesignatureResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureResponseCompound) GetBEzsigntemplatesignatureCustomdate() bool`

GetBEzsigntemplatesignatureCustomdate returns the BEzsigntemplatesignatureCustomdate field if non-nil, zero value otherwise.

### GetBEzsigntemplatesignatureCustomdateOk

`func (o *EzsigntemplatesignatureResponseCompound) GetBEzsigntemplatesignatureCustomdateOk() (*bool, bool)`

GetBEzsigntemplatesignatureCustomdateOk returns a tuple with the BEzsigntemplatesignatureCustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureResponseCompound) SetBEzsigntemplatesignatureCustomdate(v bool)`

SetBEzsigntemplatesignatureCustomdate sets BEzsigntemplatesignatureCustomdate field to given value.

### HasBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureResponseCompound) HasBEzsigntemplatesignatureCustomdate() bool`

HasBEzsigntemplatesignatureCustomdate returns a boolean if a field has been set.

### GetAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureResponseCompound) GetAObjEzsigntemplatesignaturecustomdate() []EzsigntemplatesignaturecustomdateResponseCompound`

GetAObjEzsigntemplatesignaturecustomdate returns the AObjEzsigntemplatesignaturecustomdate field if non-nil, zero value otherwise.

### GetAObjEzsigntemplatesignaturecustomdateOk

`func (o *EzsigntemplatesignatureResponseCompound) GetAObjEzsigntemplatesignaturecustomdateOk() (*[]EzsigntemplatesignaturecustomdateResponseCompound, bool)`

GetAObjEzsigntemplatesignaturecustomdateOk returns a tuple with the AObjEzsigntemplatesignaturecustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureResponseCompound) SetAObjEzsigntemplatesignaturecustomdate(v []EzsigntemplatesignaturecustomdateResponseCompound)`

SetAObjEzsigntemplatesignaturecustomdate sets AObjEzsigntemplatesignaturecustomdate field to given value.

### HasAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureResponseCompound) HasAObjEzsigntemplatesignaturecustomdate() bool`

HasAObjEzsigntemplatesignaturecustomdate returns a boolean if a field has been set.

### GetAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureResponseCompound) GetAObjEzsigntemplateelementdependency() []EzsigntemplateelementdependencyResponseCompound`

GetAObjEzsigntemplateelementdependency returns the AObjEzsigntemplateelementdependency field if non-nil, zero value otherwise.

### GetAObjEzsigntemplateelementdependencyOk

`func (o *EzsigntemplatesignatureResponseCompound) GetAObjEzsigntemplateelementdependencyOk() (*[]EzsigntemplateelementdependencyResponseCompound, bool)`

GetAObjEzsigntemplateelementdependencyOk returns a tuple with the AObjEzsigntemplateelementdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureResponseCompound) SetAObjEzsigntemplateelementdependency(v []EzsigntemplateelementdependencyResponseCompound)`

SetAObjEzsigntemplateelementdependency sets AObjEzsigntemplateelementdependency field to given value.

### HasAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureResponseCompound) HasAObjEzsigntemplateelementdependency() bool`

HasAObjEzsigntemplateelementdependency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


