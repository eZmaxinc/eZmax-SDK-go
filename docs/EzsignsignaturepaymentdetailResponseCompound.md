# EzsignsignaturepaymentdetailResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignsignaturepaymentdetailID** | **int32** | The unique ID of the Ezsignsignaturepaymentdetail | 
**FkiGlaccountcontainerID** | Pointer to **int32** | The unique ID of the Glaccountcontainer | [optional] 
**TEzsignsignaturepaymentdetailDescription** | **string** | A description for the Ezsignsignaturepaymentdetail. | 
**DEzsignsignaturepaymentdetailAmount** | **string** | The amount of the for the Ezsignsignaturepaymentdetail | 
**EEzsignsignaturepaymentdetailTaxable** | [**FieldEEzsignsignaturepaymentdetailTaxable**](FieldEEzsignsignaturepaymentdetailTaxable.md) |  | 

## Methods

### NewEzsignsignaturepaymentdetailResponseCompound

`func NewEzsignsignaturepaymentdetailResponseCompound(pkiEzsignsignaturepaymentdetailID int32, tEzsignsignaturepaymentdetailDescription string, dEzsignsignaturepaymentdetailAmount string, eEzsignsignaturepaymentdetailTaxable FieldEEzsignsignaturepaymentdetailTaxable, ) *EzsignsignaturepaymentdetailResponseCompound`

NewEzsignsignaturepaymentdetailResponseCompound instantiates a new EzsignsignaturepaymentdetailResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignaturepaymentdetailResponseCompoundWithDefaults

`func NewEzsignsignaturepaymentdetailResponseCompoundWithDefaults() *EzsignsignaturepaymentdetailResponseCompound`

NewEzsignsignaturepaymentdetailResponseCompoundWithDefaults instantiates a new EzsignsignaturepaymentdetailResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignsignaturepaymentdetailID

`func (o *EzsignsignaturepaymentdetailResponseCompound) GetPkiEzsignsignaturepaymentdetailID() int32`

GetPkiEzsignsignaturepaymentdetailID returns the PkiEzsignsignaturepaymentdetailID field if non-nil, zero value otherwise.

### GetPkiEzsignsignaturepaymentdetailIDOk

`func (o *EzsignsignaturepaymentdetailResponseCompound) GetPkiEzsignsignaturepaymentdetailIDOk() (*int32, bool)`

GetPkiEzsignsignaturepaymentdetailIDOk returns a tuple with the PkiEzsignsignaturepaymentdetailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignsignaturepaymentdetailID

`func (o *EzsignsignaturepaymentdetailResponseCompound) SetPkiEzsignsignaturepaymentdetailID(v int32)`

SetPkiEzsignsignaturepaymentdetailID sets PkiEzsignsignaturepaymentdetailID field to given value.


### GetFkiGlaccountcontainerID

`func (o *EzsignsignaturepaymentdetailResponseCompound) GetFkiGlaccountcontainerID() int32`

GetFkiGlaccountcontainerID returns the FkiGlaccountcontainerID field if non-nil, zero value otherwise.

### GetFkiGlaccountcontainerIDOk

`func (o *EzsignsignaturepaymentdetailResponseCompound) GetFkiGlaccountcontainerIDOk() (*int32, bool)`

GetFkiGlaccountcontainerIDOk returns a tuple with the FkiGlaccountcontainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountcontainerID

`func (o *EzsignsignaturepaymentdetailResponseCompound) SetFkiGlaccountcontainerID(v int32)`

SetFkiGlaccountcontainerID sets FkiGlaccountcontainerID field to given value.

### HasFkiGlaccountcontainerID

`func (o *EzsignsignaturepaymentdetailResponseCompound) HasFkiGlaccountcontainerID() bool`

HasFkiGlaccountcontainerID returns a boolean if a field has been set.

### GetTEzsignsignaturepaymentdetailDescription

`func (o *EzsignsignaturepaymentdetailResponseCompound) GetTEzsignsignaturepaymentdetailDescription() string`

GetTEzsignsignaturepaymentdetailDescription returns the TEzsignsignaturepaymentdetailDescription field if non-nil, zero value otherwise.

### GetTEzsignsignaturepaymentdetailDescriptionOk

`func (o *EzsignsignaturepaymentdetailResponseCompound) GetTEzsignsignaturepaymentdetailDescriptionOk() (*string, bool)`

GetTEzsignsignaturepaymentdetailDescriptionOk returns a tuple with the TEzsignsignaturepaymentdetailDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignsignaturepaymentdetailDescription

`func (o *EzsignsignaturepaymentdetailResponseCompound) SetTEzsignsignaturepaymentdetailDescription(v string)`

SetTEzsignsignaturepaymentdetailDescription sets TEzsignsignaturepaymentdetailDescription field to given value.


### GetDEzsignsignaturepaymentdetailAmount

`func (o *EzsignsignaturepaymentdetailResponseCompound) GetDEzsignsignaturepaymentdetailAmount() string`

GetDEzsignsignaturepaymentdetailAmount returns the DEzsignsignaturepaymentdetailAmount field if non-nil, zero value otherwise.

### GetDEzsignsignaturepaymentdetailAmountOk

`func (o *EzsignsignaturepaymentdetailResponseCompound) GetDEzsignsignaturepaymentdetailAmountOk() (*string, bool)`

GetDEzsignsignaturepaymentdetailAmountOk returns a tuple with the DEzsignsignaturepaymentdetailAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzsignsignaturepaymentdetailAmount

`func (o *EzsignsignaturepaymentdetailResponseCompound) SetDEzsignsignaturepaymentdetailAmount(v string)`

SetDEzsignsignaturepaymentdetailAmount sets DEzsignsignaturepaymentdetailAmount field to given value.


### GetEEzsignsignaturepaymentdetailTaxable

`func (o *EzsignsignaturepaymentdetailResponseCompound) GetEEzsignsignaturepaymentdetailTaxable() FieldEEzsignsignaturepaymentdetailTaxable`

GetEEzsignsignaturepaymentdetailTaxable returns the EEzsignsignaturepaymentdetailTaxable field if non-nil, zero value otherwise.

### GetEEzsignsignaturepaymentdetailTaxableOk

`func (o *EzsignsignaturepaymentdetailResponseCompound) GetEEzsignsignaturepaymentdetailTaxableOk() (*FieldEEzsignsignaturepaymentdetailTaxable, bool)`

GetEEzsignsignaturepaymentdetailTaxableOk returns a tuple with the EEzsignsignaturepaymentdetailTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignaturepaymentdetailTaxable

`func (o *EzsignsignaturepaymentdetailResponseCompound) SetEEzsignsignaturepaymentdetailTaxable(v FieldEEzsignsignaturepaymentdetailTaxable)`

SetEEzsignsignaturepaymentdetailTaxable sets EEzsignsignaturepaymentdetailTaxable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


