# EzsigntemplatesignaturepaymentdetailRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatesignaturepaymentdetailID** | Pointer to **int32** | The unique ID of the Ezsignsignaturepaymentdetail | [optional] 
**FkiGlaccountcontainerID** | **int32** | The unique ID of the Glaccountcontainer | 
**TEzsigntemplatesignaturepaymentdetailDescription** | **string** | A description for the Ezsignsignaturepaymentdetail. | 
**DEzsigntemplatesignaturepaymentdetailAmount** | **string** | The amount of the for the Ezsignsignaturepaymentdetail | 
**EEzsigntemplatesignaturepaymentdetailTaxable** | [**FieldEEzsigntemplatesignaturepaymentdetailTaxable**](FieldEEzsigntemplatesignaturepaymentdetailTaxable.md) |  | 

## Methods

### NewEzsigntemplatesignaturepaymentdetailRequestCompound

`func NewEzsigntemplatesignaturepaymentdetailRequestCompound(fkiGlaccountcontainerID int32, tEzsigntemplatesignaturepaymentdetailDescription string, dEzsigntemplatesignaturepaymentdetailAmount string, eEzsigntemplatesignaturepaymentdetailTaxable FieldEEzsigntemplatesignaturepaymentdetailTaxable, ) *EzsigntemplatesignaturepaymentdetailRequestCompound`

NewEzsigntemplatesignaturepaymentdetailRequestCompound instantiates a new EzsigntemplatesignaturepaymentdetailRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignaturepaymentdetailRequestCompoundWithDefaults

`func NewEzsigntemplatesignaturepaymentdetailRequestCompoundWithDefaults() *EzsigntemplatesignaturepaymentdetailRequestCompound`

NewEzsigntemplatesignaturepaymentdetailRequestCompoundWithDefaults instantiates a new EzsigntemplatesignaturepaymentdetailRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatesignaturepaymentdetailID

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) GetPkiEzsigntemplatesignaturepaymentdetailID() int32`

GetPkiEzsigntemplatesignaturepaymentdetailID returns the PkiEzsigntemplatesignaturepaymentdetailID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatesignaturepaymentdetailIDOk

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) GetPkiEzsigntemplatesignaturepaymentdetailIDOk() (*int32, bool)`

GetPkiEzsigntemplatesignaturepaymentdetailIDOk returns a tuple with the PkiEzsigntemplatesignaturepaymentdetailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatesignaturepaymentdetailID

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) SetPkiEzsigntemplatesignaturepaymentdetailID(v int32)`

SetPkiEzsigntemplatesignaturepaymentdetailID sets PkiEzsigntemplatesignaturepaymentdetailID field to given value.

### HasPkiEzsigntemplatesignaturepaymentdetailID

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) HasPkiEzsigntemplatesignaturepaymentdetailID() bool`

HasPkiEzsigntemplatesignaturepaymentdetailID returns a boolean if a field has been set.

### GetFkiGlaccountcontainerID

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) GetFkiGlaccountcontainerID() int32`

GetFkiGlaccountcontainerID returns the FkiGlaccountcontainerID field if non-nil, zero value otherwise.

### GetFkiGlaccountcontainerIDOk

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) GetFkiGlaccountcontainerIDOk() (*int32, bool)`

GetFkiGlaccountcontainerIDOk returns a tuple with the FkiGlaccountcontainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountcontainerID

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) SetFkiGlaccountcontainerID(v int32)`

SetFkiGlaccountcontainerID sets FkiGlaccountcontainerID field to given value.


### GetTEzsigntemplatesignaturepaymentdetailDescription

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) GetTEzsigntemplatesignaturepaymentdetailDescription() string`

GetTEzsigntemplatesignaturepaymentdetailDescription returns the TEzsigntemplatesignaturepaymentdetailDescription field if non-nil, zero value otherwise.

### GetTEzsigntemplatesignaturepaymentdetailDescriptionOk

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) GetTEzsigntemplatesignaturepaymentdetailDescriptionOk() (*string, bool)`

GetTEzsigntemplatesignaturepaymentdetailDescriptionOk returns a tuple with the TEzsigntemplatesignaturepaymentdetailDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplatesignaturepaymentdetailDescription

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) SetTEzsigntemplatesignaturepaymentdetailDescription(v string)`

SetTEzsigntemplatesignaturepaymentdetailDescription sets TEzsigntemplatesignaturepaymentdetailDescription field to given value.


### GetDEzsigntemplatesignaturepaymentdetailAmount

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) GetDEzsigntemplatesignaturepaymentdetailAmount() string`

GetDEzsigntemplatesignaturepaymentdetailAmount returns the DEzsigntemplatesignaturepaymentdetailAmount field if non-nil, zero value otherwise.

### GetDEzsigntemplatesignaturepaymentdetailAmountOk

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) GetDEzsigntemplatesignaturepaymentdetailAmountOk() (*string, bool)`

GetDEzsigntemplatesignaturepaymentdetailAmountOk returns a tuple with the DEzsigntemplatesignaturepaymentdetailAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzsigntemplatesignaturepaymentdetailAmount

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) SetDEzsigntemplatesignaturepaymentdetailAmount(v string)`

SetDEzsigntemplatesignaturepaymentdetailAmount sets DEzsigntemplatesignaturepaymentdetailAmount field to given value.


### GetEEzsigntemplatesignaturepaymentdetailTaxable

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) GetEEzsigntemplatesignaturepaymentdetailTaxable() FieldEEzsigntemplatesignaturepaymentdetailTaxable`

GetEEzsigntemplatesignaturepaymentdetailTaxable returns the EEzsigntemplatesignaturepaymentdetailTaxable field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignaturepaymentdetailTaxableOk

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) GetEEzsigntemplatesignaturepaymentdetailTaxableOk() (*FieldEEzsigntemplatesignaturepaymentdetailTaxable, bool)`

GetEEzsigntemplatesignaturepaymentdetailTaxableOk returns a tuple with the EEzsigntemplatesignaturepaymentdetailTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignaturepaymentdetailTaxable

`func (o *EzsigntemplatesignaturepaymentdetailRequestCompound) SetEEzsigntemplatesignaturepaymentdetailTaxable(v FieldEEzsigntemplatesignaturepaymentdetailTaxable)`

SetEEzsigntemplatesignaturepaymentdetailTaxable sets EEzsigntemplatesignaturepaymentdetailTaxable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


