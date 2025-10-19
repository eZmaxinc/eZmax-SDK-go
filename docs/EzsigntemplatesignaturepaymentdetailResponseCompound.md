# EzsigntemplatesignaturepaymentdetailResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatesignaturepaymentdetailID** | **int32** | The unique ID of the Ezsignsignaturepaymentdetail | 
**FkiGlaccountcontainerID** | Pointer to **int32** | The unique ID of the Glaccountcontainer | [optional] 
**TEzsigntemplatesignaturepaymentdetailDescription** | **string** | A description for the Ezsignsignaturepaymentdetail. | 
**DEzsigntemplatesignaturepaymentdetailAmount** | **string** | The amount of the for the Ezsignsignaturepaymentdetail | 
**EEzsigntemplatesignaturepaymentdetailTaxable** | [**FieldEEzsigntemplatesignaturepaymentdetailTaxable**](FieldEEzsigntemplatesignaturepaymentdetailTaxable.md) |  | 

## Methods

### NewEzsigntemplatesignaturepaymentdetailResponseCompound

`func NewEzsigntemplatesignaturepaymentdetailResponseCompound(pkiEzsigntemplatesignaturepaymentdetailID int32, tEzsigntemplatesignaturepaymentdetailDescription string, dEzsigntemplatesignaturepaymentdetailAmount string, eEzsigntemplatesignaturepaymentdetailTaxable FieldEEzsigntemplatesignaturepaymentdetailTaxable, ) *EzsigntemplatesignaturepaymentdetailResponseCompound`

NewEzsigntemplatesignaturepaymentdetailResponseCompound instantiates a new EzsigntemplatesignaturepaymentdetailResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignaturepaymentdetailResponseCompoundWithDefaults

`func NewEzsigntemplatesignaturepaymentdetailResponseCompoundWithDefaults() *EzsigntemplatesignaturepaymentdetailResponseCompound`

NewEzsigntemplatesignaturepaymentdetailResponseCompoundWithDefaults instantiates a new EzsigntemplatesignaturepaymentdetailResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatesignaturepaymentdetailID

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) GetPkiEzsigntemplatesignaturepaymentdetailID() int32`

GetPkiEzsigntemplatesignaturepaymentdetailID returns the PkiEzsigntemplatesignaturepaymentdetailID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatesignaturepaymentdetailIDOk

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) GetPkiEzsigntemplatesignaturepaymentdetailIDOk() (*int32, bool)`

GetPkiEzsigntemplatesignaturepaymentdetailIDOk returns a tuple with the PkiEzsigntemplatesignaturepaymentdetailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatesignaturepaymentdetailID

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) SetPkiEzsigntemplatesignaturepaymentdetailID(v int32)`

SetPkiEzsigntemplatesignaturepaymentdetailID sets PkiEzsigntemplatesignaturepaymentdetailID field to given value.


### GetFkiGlaccountcontainerID

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) GetFkiGlaccountcontainerID() int32`

GetFkiGlaccountcontainerID returns the FkiGlaccountcontainerID field if non-nil, zero value otherwise.

### GetFkiGlaccountcontainerIDOk

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) GetFkiGlaccountcontainerIDOk() (*int32, bool)`

GetFkiGlaccountcontainerIDOk returns a tuple with the FkiGlaccountcontainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountcontainerID

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) SetFkiGlaccountcontainerID(v int32)`

SetFkiGlaccountcontainerID sets FkiGlaccountcontainerID field to given value.

### HasFkiGlaccountcontainerID

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) HasFkiGlaccountcontainerID() bool`

HasFkiGlaccountcontainerID returns a boolean if a field has been set.

### GetTEzsigntemplatesignaturepaymentdetailDescription

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) GetTEzsigntemplatesignaturepaymentdetailDescription() string`

GetTEzsigntemplatesignaturepaymentdetailDescription returns the TEzsigntemplatesignaturepaymentdetailDescription field if non-nil, zero value otherwise.

### GetTEzsigntemplatesignaturepaymentdetailDescriptionOk

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) GetTEzsigntemplatesignaturepaymentdetailDescriptionOk() (*string, bool)`

GetTEzsigntemplatesignaturepaymentdetailDescriptionOk returns a tuple with the TEzsigntemplatesignaturepaymentdetailDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplatesignaturepaymentdetailDescription

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) SetTEzsigntemplatesignaturepaymentdetailDescription(v string)`

SetTEzsigntemplatesignaturepaymentdetailDescription sets TEzsigntemplatesignaturepaymentdetailDescription field to given value.


### GetDEzsigntemplatesignaturepaymentdetailAmount

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) GetDEzsigntemplatesignaturepaymentdetailAmount() string`

GetDEzsigntemplatesignaturepaymentdetailAmount returns the DEzsigntemplatesignaturepaymentdetailAmount field if non-nil, zero value otherwise.

### GetDEzsigntemplatesignaturepaymentdetailAmountOk

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) GetDEzsigntemplatesignaturepaymentdetailAmountOk() (*string, bool)`

GetDEzsigntemplatesignaturepaymentdetailAmountOk returns a tuple with the DEzsigntemplatesignaturepaymentdetailAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzsigntemplatesignaturepaymentdetailAmount

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) SetDEzsigntemplatesignaturepaymentdetailAmount(v string)`

SetDEzsigntemplatesignaturepaymentdetailAmount sets DEzsigntemplatesignaturepaymentdetailAmount field to given value.


### GetEEzsigntemplatesignaturepaymentdetailTaxable

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) GetEEzsigntemplatesignaturepaymentdetailTaxable() FieldEEzsigntemplatesignaturepaymentdetailTaxable`

GetEEzsigntemplatesignaturepaymentdetailTaxable returns the EEzsigntemplatesignaturepaymentdetailTaxable field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignaturepaymentdetailTaxableOk

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) GetEEzsigntemplatesignaturepaymentdetailTaxableOk() (*FieldEEzsigntemplatesignaturepaymentdetailTaxable, bool)`

GetEEzsigntemplatesignaturepaymentdetailTaxableOk returns a tuple with the EEzsigntemplatesignaturepaymentdetailTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignaturepaymentdetailTaxable

`func (o *EzsigntemplatesignaturepaymentdetailResponseCompound) SetEEzsigntemplatesignaturepaymentdetailTaxable(v FieldEEzsigntemplatesignaturepaymentdetailTaxable)`

SetEEzsigntemplatesignaturepaymentdetailTaxable sets EEzsigntemplatesignaturepaymentdetailTaxable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


