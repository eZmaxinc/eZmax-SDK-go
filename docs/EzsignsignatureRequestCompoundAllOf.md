# EzsignsignatureRequestCompoundAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BEzsignsignatureCustomdate** | Pointer to **bool** | Whether the Ezsignsignature has a custom date format or not. (Only possible when eEzsignsignatureType is \&quot;Name\&quot; or \&quot;Handwritten\&quot;) | [optional] 
**AObjEzsignsignaturecustomdate** | Pointer to [**[]EzsignsignaturecustomdateRequest**](EzsignsignaturecustomdateRequest.md) | An array of custom date blocks that will be filled at the time of signature.  Can only be used if bEzsignsignatureCustomdate is true.  Use an empty array if you don&#39;t want to have a date at all. | [optional] 

## Methods

### NewEzsignsignatureRequestCompoundAllOf

`func NewEzsignsignatureRequestCompoundAllOf() *EzsignsignatureRequestCompoundAllOf`

NewEzsignsignatureRequestCompoundAllOf instantiates a new EzsignsignatureRequestCompoundAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignatureRequestCompoundAllOfWithDefaults

`func NewEzsignsignatureRequestCompoundAllOfWithDefaults() *EzsignsignatureRequestCompoundAllOf`

NewEzsignsignatureRequestCompoundAllOfWithDefaults instantiates a new EzsignsignatureRequestCompoundAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBEzsignsignatureCustomdate

`func (o *EzsignsignatureRequestCompoundAllOf) GetBEzsignsignatureCustomdate() bool`

GetBEzsignsignatureCustomdate returns the BEzsignsignatureCustomdate field if non-nil, zero value otherwise.

### GetBEzsignsignatureCustomdateOk

`func (o *EzsignsignatureRequestCompoundAllOf) GetBEzsignsignatureCustomdateOk() (*bool, bool)`

GetBEzsignsignatureCustomdateOk returns a tuple with the BEzsignsignatureCustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureCustomdate

`func (o *EzsignsignatureRequestCompoundAllOf) SetBEzsignsignatureCustomdate(v bool)`

SetBEzsignsignatureCustomdate sets BEzsignsignatureCustomdate field to given value.

### HasBEzsignsignatureCustomdate

`func (o *EzsignsignatureRequestCompoundAllOf) HasBEzsignsignatureCustomdate() bool`

HasBEzsignsignatureCustomdate returns a boolean if a field has been set.

### GetAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureRequestCompoundAllOf) GetAObjEzsignsignaturecustomdate() []EzsignsignaturecustomdateRequest`

GetAObjEzsignsignaturecustomdate returns the AObjEzsignsignaturecustomdate field if non-nil, zero value otherwise.

### GetAObjEzsignsignaturecustomdateOk

`func (o *EzsignsignatureRequestCompoundAllOf) GetAObjEzsignsignaturecustomdateOk() (*[]EzsignsignaturecustomdateRequest, bool)`

GetAObjEzsignsignaturecustomdateOk returns a tuple with the AObjEzsignsignaturecustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureRequestCompoundAllOf) SetAObjEzsignsignaturecustomdate(v []EzsignsignaturecustomdateRequest)`

SetAObjEzsignsignaturecustomdate sets AObjEzsignsignaturecustomdate field to given value.

### HasAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureRequestCompoundAllOf) HasAObjEzsignsignaturecustomdate() bool`

HasAObjEzsignsignaturecustomdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


