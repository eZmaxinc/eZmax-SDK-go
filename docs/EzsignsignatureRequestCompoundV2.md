# EzsignsignatureRequestCompoundV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BEzsignsignatureCustomdate** | Pointer to **bool** | Whether the Ezsignsignature has a custom date format or not. (Only possible when eEzsignsignatureType is **Name** or **Handwritten**) | [optional] 
**AObjEzsignsignaturecustomdate** | Pointer to [**[]EzsignsignaturecustomdateRequestV2**](EzsignsignaturecustomdateRequestV2.md) | An array of custom date blocks that will be filled at the time of signature.  Can only be used if bEzsignsignatureCustomdate is true.  Use an empty array if you don&#39;t want to have a date at all. | [optional] 
**AObjEzsignelementdependency** | Pointer to [**[]EzsignelementdependencyRequest**](EzsignelementdependencyRequest.md) |  | [optional] 

## Methods

### NewEzsignsignatureRequestCompoundV2

`func NewEzsignsignatureRequestCompoundV2() *EzsignsignatureRequestCompoundV2`

NewEzsignsignatureRequestCompoundV2 instantiates a new EzsignsignatureRequestCompoundV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignatureRequestCompoundV2WithDefaults

`func NewEzsignsignatureRequestCompoundV2WithDefaults() *EzsignsignatureRequestCompoundV2`

NewEzsignsignatureRequestCompoundV2WithDefaults instantiates a new EzsignsignatureRequestCompoundV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBEzsignsignatureCustomdate

`func (o *EzsignsignatureRequestCompoundV2) GetBEzsignsignatureCustomdate() bool`

GetBEzsignsignatureCustomdate returns the BEzsignsignatureCustomdate field if non-nil, zero value otherwise.

### GetBEzsignsignatureCustomdateOk

`func (o *EzsignsignatureRequestCompoundV2) GetBEzsignsignatureCustomdateOk() (*bool, bool)`

GetBEzsignsignatureCustomdateOk returns a tuple with the BEzsignsignatureCustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureCustomdate

`func (o *EzsignsignatureRequestCompoundV2) SetBEzsignsignatureCustomdate(v bool)`

SetBEzsignsignatureCustomdate sets BEzsignsignatureCustomdate field to given value.

### HasBEzsignsignatureCustomdate

`func (o *EzsignsignatureRequestCompoundV2) HasBEzsignsignatureCustomdate() bool`

HasBEzsignsignatureCustomdate returns a boolean if a field has been set.

### GetAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureRequestCompoundV2) GetAObjEzsignsignaturecustomdate() []EzsignsignaturecustomdateRequestCompoundV2`

GetAObjEzsignsignaturecustomdate returns the AObjEzsignsignaturecustomdate field if non-nil, zero value otherwise.

### GetAObjEzsignsignaturecustomdateOk

`func (o *EzsignsignatureRequestCompoundV2) GetAObjEzsignsignaturecustomdateOk() (*[]EzsignsignaturecustomdateRequestCompoundV2, bool)`

GetAObjEzsignsignaturecustomdateOk returns a tuple with the AObjEzsignsignaturecustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureRequestCompoundV2) SetAObjEzsignsignaturecustomdate(v []EzsignsignaturecustomdateRequestCompoundV2)`

SetAObjEzsignsignaturecustomdate sets AObjEzsignsignaturecustomdate field to given value.

### HasAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureRequestCompoundV2) HasAObjEzsignsignaturecustomdate() bool`

HasAObjEzsignsignaturecustomdate returns a boolean if a field has been set.

### GetAObjEzsignelementdependency

`func (o *EzsignsignatureRequestCompoundV2) GetAObjEzsignelementdependency() []EzsignelementdependencyRequestCompound`

GetAObjEzsignelementdependency returns the AObjEzsignelementdependency field if non-nil, zero value otherwise.

### GetAObjEzsignelementdependencyOk

`func (o *EzsignsignatureRequestCompoundV2) GetAObjEzsignelementdependencyOk() (*[]EzsignelementdependencyRequestCompound, bool)`

GetAObjEzsignelementdependencyOk returns a tuple with the AObjEzsignelementdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignelementdependency

`func (o *EzsignsignatureRequestCompoundV2) SetAObjEzsignelementdependency(v []EzsignelementdependencyRequestCompound)`

SetAObjEzsignelementdependency sets AObjEzsignelementdependency field to given value.

### HasAObjEzsignelementdependency

`func (o *EzsignsignatureRequestCompoundV2) HasAObjEzsignelementdependency() bool`

HasAObjEzsignelementdependency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


