# EzsignsignatureResponseCompoundV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BEzsignsignatureCustomdate** | Pointer to **bool** | Whether the Ezsignsignature has a custom date format or not. (Only possible when eEzsignsignatureType is **Name** or **Handwritten**) | [optional] 
**AObjEzsignsignaturecustomdate** | Pointer to [**[]EzsignsignaturecustomdateResponseV2**](EzsignsignaturecustomdateResponseV2.md) | An array of custom date blocks that will be filled at the time of signature.  Can only be used if bEzsignsignatureCustomdate is true.  Use an empty array if you don&#39;t want to have a date at all. | [optional] 
**ObjCreditcardtransaction** | Pointer to [**CustomCreditcardtransactionResponse**](CustomCreditcardtransactionResponse.md) |  | [optional] 
**AObjEzsignelementdependency** | Pointer to [**[]EzsignelementdependencyResponse**](EzsignelementdependencyResponse.md) |  | [optional] 

## Methods

### NewEzsignsignatureResponseCompoundV3

`func NewEzsignsignatureResponseCompoundV3() *EzsignsignatureResponseCompoundV3`

NewEzsignsignatureResponseCompoundV3 instantiates a new EzsignsignatureResponseCompoundV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignatureResponseCompoundV3WithDefaults

`func NewEzsignsignatureResponseCompoundV3WithDefaults() *EzsignsignatureResponseCompoundV3`

NewEzsignsignatureResponseCompoundV3WithDefaults instantiates a new EzsignsignatureResponseCompoundV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBEzsignsignatureCustomdate

`func (o *EzsignsignatureResponseCompoundV3) GetBEzsignsignatureCustomdate() bool`

GetBEzsignsignatureCustomdate returns the BEzsignsignatureCustomdate field if non-nil, zero value otherwise.

### GetBEzsignsignatureCustomdateOk

`func (o *EzsignsignatureResponseCompoundV3) GetBEzsignsignatureCustomdateOk() (*bool, bool)`

GetBEzsignsignatureCustomdateOk returns a tuple with the BEzsignsignatureCustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureCustomdate

`func (o *EzsignsignatureResponseCompoundV3) SetBEzsignsignatureCustomdate(v bool)`

SetBEzsignsignatureCustomdate sets BEzsignsignatureCustomdate field to given value.

### HasBEzsignsignatureCustomdate

`func (o *EzsignsignatureResponseCompoundV3) HasBEzsignsignatureCustomdate() bool`

HasBEzsignsignatureCustomdate returns a boolean if a field has been set.

### GetAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureResponseCompoundV3) GetAObjEzsignsignaturecustomdate() []EzsignsignaturecustomdateResponseCompoundV2`

GetAObjEzsignsignaturecustomdate returns the AObjEzsignsignaturecustomdate field if non-nil, zero value otherwise.

### GetAObjEzsignsignaturecustomdateOk

`func (o *EzsignsignatureResponseCompoundV3) GetAObjEzsignsignaturecustomdateOk() (*[]EzsignsignaturecustomdateResponseCompoundV2, bool)`

GetAObjEzsignsignaturecustomdateOk returns a tuple with the AObjEzsignsignaturecustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureResponseCompoundV3) SetAObjEzsignsignaturecustomdate(v []EzsignsignaturecustomdateResponseCompoundV2)`

SetAObjEzsignsignaturecustomdate sets AObjEzsignsignaturecustomdate field to given value.

### HasAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureResponseCompoundV3) HasAObjEzsignsignaturecustomdate() bool`

HasAObjEzsignsignaturecustomdate returns a boolean if a field has been set.

### GetObjCreditcardtransaction

`func (o *EzsignsignatureResponseCompoundV3) GetObjCreditcardtransaction() CustomCreditcardtransactionResponse`

GetObjCreditcardtransaction returns the ObjCreditcardtransaction field if non-nil, zero value otherwise.

### GetObjCreditcardtransactionOk

`func (o *EzsignsignatureResponseCompoundV3) GetObjCreditcardtransactionOk() (*CustomCreditcardtransactionResponse, bool)`

GetObjCreditcardtransactionOk returns a tuple with the ObjCreditcardtransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCreditcardtransaction

`func (o *EzsignsignatureResponseCompoundV3) SetObjCreditcardtransaction(v CustomCreditcardtransactionResponse)`

SetObjCreditcardtransaction sets ObjCreditcardtransaction field to given value.

### HasObjCreditcardtransaction

`func (o *EzsignsignatureResponseCompoundV3) HasObjCreditcardtransaction() bool`

HasObjCreditcardtransaction returns a boolean if a field has been set.

### GetAObjEzsignelementdependency

`func (o *EzsignsignatureResponseCompoundV3) GetAObjEzsignelementdependency() []EzsignelementdependencyResponseCompound`

GetAObjEzsignelementdependency returns the AObjEzsignelementdependency field if non-nil, zero value otherwise.

### GetAObjEzsignelementdependencyOk

`func (o *EzsignsignatureResponseCompoundV3) GetAObjEzsignelementdependencyOk() (*[]EzsignelementdependencyResponseCompound, bool)`

GetAObjEzsignelementdependencyOk returns a tuple with the AObjEzsignelementdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignelementdependency

`func (o *EzsignsignatureResponseCompoundV3) SetAObjEzsignelementdependency(v []EzsignelementdependencyResponseCompound)`

SetAObjEzsignelementdependency sets AObjEzsignelementdependency field to given value.

### HasAObjEzsignelementdependency

`func (o *EzsignsignatureResponseCompoundV3) HasAObjEzsignelementdependency() bool`

HasAObjEzsignelementdependency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


