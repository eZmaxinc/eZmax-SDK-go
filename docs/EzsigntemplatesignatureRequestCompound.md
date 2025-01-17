# EzsigntemplatesignatureRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BEzsigntemplatesignatureCustomdate** | Pointer to **bool** | Whether the Ezsigntemplatesignature has a custom date format or not. (Only possible when eEzsigntemplatesignatureType is **Name** or **Handwritten**) | [optional] 
**AObjEzsigntemplatesignaturecustomdate** | Pointer to [**[]EzsigntemplatesignaturecustomdateRequest**](EzsigntemplatesignaturecustomdateRequest.md) | An array of custom date blocks that will be filled at the time of signature.  Can only be used if bEzsigntemplatesignatureCustomdate is true.  Use an empty array if you don&#39;t want to have a date at all. | [optional] 
**AObjEzsigntemplateelementdependency** | Pointer to [**[]EzsigntemplateelementdependencyRequest**](EzsigntemplateelementdependencyRequest.md) |  | [optional] 

## Methods

### NewEzsigntemplatesignatureRequestCompound

`func NewEzsigntemplatesignatureRequestCompound() *EzsigntemplatesignatureRequestCompound`

NewEzsigntemplatesignatureRequestCompound instantiates a new EzsigntemplatesignatureRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignatureRequestCompoundWithDefaults

`func NewEzsigntemplatesignatureRequestCompoundWithDefaults() *EzsigntemplatesignatureRequestCompound`

NewEzsigntemplatesignatureRequestCompoundWithDefaults instantiates a new EzsigntemplatesignatureRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureRequestCompound) GetBEzsigntemplatesignatureCustomdate() bool`

GetBEzsigntemplatesignatureCustomdate returns the BEzsigntemplatesignatureCustomdate field if non-nil, zero value otherwise.

### GetBEzsigntemplatesignatureCustomdateOk

`func (o *EzsigntemplatesignatureRequestCompound) GetBEzsigntemplatesignatureCustomdateOk() (*bool, bool)`

GetBEzsigntemplatesignatureCustomdateOk returns a tuple with the BEzsigntemplatesignatureCustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureRequestCompound) SetBEzsigntemplatesignatureCustomdate(v bool)`

SetBEzsigntemplatesignatureCustomdate sets BEzsigntemplatesignatureCustomdate field to given value.

### HasBEzsigntemplatesignatureCustomdate

`func (o *EzsigntemplatesignatureRequestCompound) HasBEzsigntemplatesignatureCustomdate() bool`

HasBEzsigntemplatesignatureCustomdate returns a boolean if a field has been set.

### GetAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureRequestCompound) GetAObjEzsigntemplatesignaturecustomdate() []EzsigntemplatesignaturecustomdateRequestCompound`

GetAObjEzsigntemplatesignaturecustomdate returns the AObjEzsigntemplatesignaturecustomdate field if non-nil, zero value otherwise.

### GetAObjEzsigntemplatesignaturecustomdateOk

`func (o *EzsigntemplatesignatureRequestCompound) GetAObjEzsigntemplatesignaturecustomdateOk() (*[]EzsigntemplatesignaturecustomdateRequestCompound, bool)`

GetAObjEzsigntemplatesignaturecustomdateOk returns a tuple with the AObjEzsigntemplatesignaturecustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureRequestCompound) SetAObjEzsigntemplatesignaturecustomdate(v []EzsigntemplatesignaturecustomdateRequestCompound)`

SetAObjEzsigntemplatesignaturecustomdate sets AObjEzsigntemplatesignaturecustomdate field to given value.

### HasAObjEzsigntemplatesignaturecustomdate

`func (o *EzsigntemplatesignatureRequestCompound) HasAObjEzsigntemplatesignaturecustomdate() bool`

HasAObjEzsigntemplatesignaturecustomdate returns a boolean if a field has been set.

### GetAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureRequestCompound) GetAObjEzsigntemplateelementdependency() []EzsigntemplateelementdependencyRequestCompound`

GetAObjEzsigntemplateelementdependency returns the AObjEzsigntemplateelementdependency field if non-nil, zero value otherwise.

### GetAObjEzsigntemplateelementdependencyOk

`func (o *EzsigntemplatesignatureRequestCompound) GetAObjEzsigntemplateelementdependencyOk() (*[]EzsigntemplateelementdependencyRequestCompound, bool)`

GetAObjEzsigntemplateelementdependencyOk returns a tuple with the AObjEzsigntemplateelementdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureRequestCompound) SetAObjEzsigntemplateelementdependency(v []EzsigntemplateelementdependencyRequestCompound)`

SetAObjEzsigntemplateelementdependency sets AObjEzsigntemplateelementdependency field to given value.

### HasAObjEzsigntemplateelementdependency

`func (o *EzsigntemplatesignatureRequestCompound) HasAObjEzsigntemplateelementdependency() bool`

HasAObjEzsigntemplateelementdependency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


