# EzsignsignatureResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DtEzsignsignatureDateInFolderTimezone** | Pointer to **string** | The date the Ezsignsignature was signed in folder&#39;s timezone | [optional] 
**BEzsignsignatureCustomdate** | Pointer to **bool** | Whether the Ezsignsignature has a custom date format or not. (Only possible when eEzsignsignatureType is **Name** or **Handwritten**) | [optional] 
**AObjEzsignsignaturecustomdate** | Pointer to [**[]EzsignsignaturecustomdateResponse**](EzsignsignaturecustomdateResponse.md) | An array of custom date blocks that will be filled at the time of signature.  Can only be used if bEzsignsignatureCustomdate is true.  Use an empty array if you don&#39;t want to have a date at all. | [optional] 
**ObjCreditcardtransaction** | Pointer to [**CustomCreditcardtransactionResponse**](CustomCreditcardtransactionResponse.md) |  | [optional] 
**AObjEzsignelementdependency** | Pointer to [**[]EzsignelementdependencyResponse**](EzsignelementdependencyResponse.md) |  | [optional] 
**ObjTimezone** | Pointer to [**CustomTimezoneWithCodeResponse**](CustomTimezoneWithCodeResponse.md) |  | [optional] 

## Methods

### NewEzsignsignatureResponseCompound

`func NewEzsignsignatureResponseCompound() *EzsignsignatureResponseCompound`

NewEzsignsignatureResponseCompound instantiates a new EzsignsignatureResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignatureResponseCompoundWithDefaults

`func NewEzsignsignatureResponseCompoundWithDefaults() *EzsignsignatureResponseCompound`

NewEzsignsignatureResponseCompoundWithDefaults instantiates a new EzsignsignatureResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDtEzsignsignatureDateInFolderTimezone

`func (o *EzsignsignatureResponseCompound) GetDtEzsignsignatureDateInFolderTimezone() string`

GetDtEzsignsignatureDateInFolderTimezone returns the DtEzsignsignatureDateInFolderTimezone field if non-nil, zero value otherwise.

### GetDtEzsignsignatureDateInFolderTimezoneOk

`func (o *EzsignsignatureResponseCompound) GetDtEzsignsignatureDateInFolderTimezoneOk() (*string, bool)`

GetDtEzsignsignatureDateInFolderTimezoneOk returns a tuple with the DtEzsignsignatureDateInFolderTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignsignatureDateInFolderTimezone

`func (o *EzsignsignatureResponseCompound) SetDtEzsignsignatureDateInFolderTimezone(v string)`

SetDtEzsignsignatureDateInFolderTimezone sets DtEzsignsignatureDateInFolderTimezone field to given value.

### HasDtEzsignsignatureDateInFolderTimezone

`func (o *EzsignsignatureResponseCompound) HasDtEzsignsignatureDateInFolderTimezone() bool`

HasDtEzsignsignatureDateInFolderTimezone returns a boolean if a field has been set.

### GetBEzsignsignatureCustomdate

`func (o *EzsignsignatureResponseCompound) GetBEzsignsignatureCustomdate() bool`

GetBEzsignsignatureCustomdate returns the BEzsignsignatureCustomdate field if non-nil, zero value otherwise.

### GetBEzsignsignatureCustomdateOk

`func (o *EzsignsignatureResponseCompound) GetBEzsignsignatureCustomdateOk() (*bool, bool)`

GetBEzsignsignatureCustomdateOk returns a tuple with the BEzsignsignatureCustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureCustomdate

`func (o *EzsignsignatureResponseCompound) SetBEzsignsignatureCustomdate(v bool)`

SetBEzsignsignatureCustomdate sets BEzsignsignatureCustomdate field to given value.

### HasBEzsignsignatureCustomdate

`func (o *EzsignsignatureResponseCompound) HasBEzsignsignatureCustomdate() bool`

HasBEzsignsignatureCustomdate returns a boolean if a field has been set.

### GetAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureResponseCompound) GetAObjEzsignsignaturecustomdate() []EzsignsignaturecustomdateResponseCompound`

GetAObjEzsignsignaturecustomdate returns the AObjEzsignsignaturecustomdate field if non-nil, zero value otherwise.

### GetAObjEzsignsignaturecustomdateOk

`func (o *EzsignsignatureResponseCompound) GetAObjEzsignsignaturecustomdateOk() (*[]EzsignsignaturecustomdateResponseCompound, bool)`

GetAObjEzsignsignaturecustomdateOk returns a tuple with the AObjEzsignsignaturecustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureResponseCompound) SetAObjEzsignsignaturecustomdate(v []EzsignsignaturecustomdateResponseCompound)`

SetAObjEzsignsignaturecustomdate sets AObjEzsignsignaturecustomdate field to given value.

### HasAObjEzsignsignaturecustomdate

`func (o *EzsignsignatureResponseCompound) HasAObjEzsignsignaturecustomdate() bool`

HasAObjEzsignsignaturecustomdate returns a boolean if a field has been set.

### GetObjCreditcardtransaction

`func (o *EzsignsignatureResponseCompound) GetObjCreditcardtransaction() CustomCreditcardtransactionResponse`

GetObjCreditcardtransaction returns the ObjCreditcardtransaction field if non-nil, zero value otherwise.

### GetObjCreditcardtransactionOk

`func (o *EzsignsignatureResponseCompound) GetObjCreditcardtransactionOk() (*CustomCreditcardtransactionResponse, bool)`

GetObjCreditcardtransactionOk returns a tuple with the ObjCreditcardtransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCreditcardtransaction

`func (o *EzsignsignatureResponseCompound) SetObjCreditcardtransaction(v CustomCreditcardtransactionResponse)`

SetObjCreditcardtransaction sets ObjCreditcardtransaction field to given value.

### HasObjCreditcardtransaction

`func (o *EzsignsignatureResponseCompound) HasObjCreditcardtransaction() bool`

HasObjCreditcardtransaction returns a boolean if a field has been set.

### GetAObjEzsignelementdependency

`func (o *EzsignsignatureResponseCompound) GetAObjEzsignelementdependency() []EzsignelementdependencyResponseCompound`

GetAObjEzsignelementdependency returns the AObjEzsignelementdependency field if non-nil, zero value otherwise.

### GetAObjEzsignelementdependencyOk

`func (o *EzsignsignatureResponseCompound) GetAObjEzsignelementdependencyOk() (*[]EzsignelementdependencyResponseCompound, bool)`

GetAObjEzsignelementdependencyOk returns a tuple with the AObjEzsignelementdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignelementdependency

`func (o *EzsignsignatureResponseCompound) SetAObjEzsignelementdependency(v []EzsignelementdependencyResponseCompound)`

SetAObjEzsignelementdependency sets AObjEzsignelementdependency field to given value.

### HasAObjEzsignelementdependency

`func (o *EzsignsignatureResponseCompound) HasAObjEzsignelementdependency() bool`

HasAObjEzsignelementdependency returns a boolean if a field has been set.

### GetObjTimezone

`func (o *EzsignsignatureResponseCompound) GetObjTimezone() CustomTimezoneWithCodeResponse`

GetObjTimezone returns the ObjTimezone field if non-nil, zero value otherwise.

### GetObjTimezoneOk

`func (o *EzsignsignatureResponseCompound) GetObjTimezoneOk() (*CustomTimezoneWithCodeResponse, bool)`

GetObjTimezoneOk returns a tuple with the ObjTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjTimezone

`func (o *EzsignsignatureResponseCompound) SetObjTimezone(v CustomTimezoneWithCodeResponse)`

SetObjTimezone sets ObjTimezone field to given value.

### HasObjTimezone

`func (o *EzsignsignatureResponseCompound) HasObjTimezone() bool`

HasObjTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


