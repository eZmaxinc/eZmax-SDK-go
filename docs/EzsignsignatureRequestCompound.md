# EzsignsignatureRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BEzsignsignatureCustomdate** | Pointer to **bool** | Whether the Ezsignsignature has a custom date format or not. (Only possible when eEzsignsignatureType is **Name** or **Handwritten**) | [optional] 
**AObjEzsignsignaturecustomdate** | Pointer to [**[]EzsignsignaturecustomdateRequest**](EzsignsignaturecustomdateRequest.md) | An array of custom date blocks that will be filled at the time of signature.  Can only be used if bEzsignsignatureCustomdate is true.  Use an empty array if you don&#39;t want to have a date at all. | [optional] 
**AObjEzsignelementdependency** | Pointer to [**[]EzsignelementdependencyRequest**](EzsignelementdependencyRequest.md) |  | [optional] 

## Methods

### NewEzsignsignatureRequestCompound

`func NewEzsignsignatureRequestCompound() *EzsignsignatureRequestCompound`

NewEzsignsignatureRequestCompound instantiates a new EzsignsignatureRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignatureRequestCompoundWithDefaults

`func NewEzsignsignatureRequestCompoundWithDefaults() *EzsignsignatureRequestCompound`

NewEzsignsignatureRequestCompoundWithDefaults instantiates a new EzsignsignatureRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBEzsignsignatureCustomdate

`func (o *EzsignsignatureRequestCompound) GetBEzsignsignatureCustomdate() bool`

GetBEzsignsignatureCustomdate returns the BEzsignsignatureCustomdate field if non-nil, zero value otherwise.

### GetBEzsignsignatureCustomdateOk

`func (o *EzsignsignatureRequestCompound) GetBEzsignsignatureCustomdateOk() (*bool, bool)`

GetBEzsignsignatureCustomdateOk returns a tuple with the BEzsignsignatureCustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureCustomdate

`func (o *EzsignsignatureRequestCompound) SetBEzsignsignatureCustomdate(v bool)`

SetBEzsignsignatureCustomdate sets BEzsignsignatureCustomdate field to given value.

### HasBEzsignsignatureCustomdate

`func (o *EzsignsignatureRequestCompound) HasBEzsignsignatureCustomdate() bool`

HasBEzsignsignatureCustomdate returns a boolean if a field has been set.

### GetAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureRequestCompound) GetAObjEzsignsignaturecustomdate() []EzsignsignaturecustomdateRequestCompound`

GetAObjEzsignsignaturecustomdate returns the AObjEzsignsignaturecustomdate field if non-nil, zero value otherwise.

### GetAObjEzsignsignaturecustomdateOk

`func (o *EzsignsignatureRequestCompound) GetAObjEzsignsignaturecustomdateOk() (*[]EzsignsignaturecustomdateRequestCompound, bool)`

GetAObjEzsignsignaturecustomdateOk returns a tuple with the AObjEzsignsignaturecustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureRequestCompound) SetAObjEzsignsignaturecustomdate(v []EzsignsignaturecustomdateRequestCompound)`

SetAObjEzsignsignaturecustomdate sets AObjEzsignsignaturecustomdate field to given value.

### HasAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureRequestCompound) HasAObjEzsignsignaturecustomdate() bool`

HasAObjEzsignsignaturecustomdate returns a boolean if a field has been set.

### GetAObjEzsignelementdependency

`func (o *EzsignsignatureRequestCompound) GetAObjEzsignelementdependency() []EzsignelementdependencyRequestCompound`

GetAObjEzsignelementdependency returns the AObjEzsignelementdependency field if non-nil, zero value otherwise.

### GetAObjEzsignelementdependencyOk

`func (o *EzsignsignatureRequestCompound) GetAObjEzsignelementdependencyOk() (*[]EzsignelementdependencyRequestCompound, bool)`

GetAObjEzsignelementdependencyOk returns a tuple with the AObjEzsignelementdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignelementdependency

`func (o *EzsignsignatureRequestCompound) SetAObjEzsignelementdependency(v []EzsignelementdependencyRequestCompound)`

SetAObjEzsignelementdependency sets AObjEzsignelementdependency field to given value.

### HasAObjEzsignelementdependency

`func (o *EzsignsignatureRequestCompound) HasAObjEzsignelementdependency() bool`

HasAObjEzsignelementdependency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


