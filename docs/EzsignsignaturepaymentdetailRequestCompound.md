# EzsignsignaturepaymentdetailRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignsignaturepaymentdetailID** | Pointer to **int32** | The unique ID of the Ezsignsignaturepaymentdetail | [optional] 
**FkiGlaccountcontainerID** | Pointer to **int32** | The unique ID of the Glaccountcontainer | [optional] 
**TEzsignsignaturepaymentdetailDescription** | **string** | A description for the Ezsignsignaturepaymentdetail. | 
**DEzsignsignaturepaymentdetailAmount** | **string** | The amount of the for the Ezsignsignaturepaymentdetail | 
**EEzsignsignaturepaymentdetailTaxable** | [**FieldEEzsignsignaturepaymentdetailTaxable**](FieldEEzsignsignaturepaymentdetailTaxable.md) |  | 

## Methods

### NewEzsignsignaturepaymentdetailRequestCompound

`func NewEzsignsignaturepaymentdetailRequestCompound(tEzsignsignaturepaymentdetailDescription string, dEzsignsignaturepaymentdetailAmount string, eEzsignsignaturepaymentdetailTaxable FieldEEzsignsignaturepaymentdetailTaxable, ) *EzsignsignaturepaymentdetailRequestCompound`

NewEzsignsignaturepaymentdetailRequestCompound instantiates a new EzsignsignaturepaymentdetailRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignaturepaymentdetailRequestCompoundWithDefaults

`func NewEzsignsignaturepaymentdetailRequestCompoundWithDefaults() *EzsignsignaturepaymentdetailRequestCompound`

NewEzsignsignaturepaymentdetailRequestCompoundWithDefaults instantiates a new EzsignsignaturepaymentdetailRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignsignaturepaymentdetailID

`func (o *EzsignsignaturepaymentdetailRequestCompound) GetPkiEzsignsignaturepaymentdetailID() int32`

GetPkiEzsignsignaturepaymentdetailID returns the PkiEzsignsignaturepaymentdetailID field if non-nil, zero value otherwise.

### GetPkiEzsignsignaturepaymentdetailIDOk

`func (o *EzsignsignaturepaymentdetailRequestCompound) GetPkiEzsignsignaturepaymentdetailIDOk() (*int32, bool)`

GetPkiEzsignsignaturepaymentdetailIDOk returns a tuple with the PkiEzsignsignaturepaymentdetailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignsignaturepaymentdetailID

`func (o *EzsignsignaturepaymentdetailRequestCompound) SetPkiEzsignsignaturepaymentdetailID(v int32)`

SetPkiEzsignsignaturepaymentdetailID sets PkiEzsignsignaturepaymentdetailID field to given value.

### HasPkiEzsignsignaturepaymentdetailID

`func (o *EzsignsignaturepaymentdetailRequestCompound) HasPkiEzsignsignaturepaymentdetailID() bool`

HasPkiEzsignsignaturepaymentdetailID returns a boolean if a field has been set.

### GetFkiGlaccountcontainerID

`func (o *EzsignsignaturepaymentdetailRequestCompound) GetFkiGlaccountcontainerID() int32`

GetFkiGlaccountcontainerID returns the FkiGlaccountcontainerID field if non-nil, zero value otherwise.

### GetFkiGlaccountcontainerIDOk

`func (o *EzsignsignaturepaymentdetailRequestCompound) GetFkiGlaccountcontainerIDOk() (*int32, bool)`

GetFkiGlaccountcontainerIDOk returns a tuple with the FkiGlaccountcontainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountcontainerID

`func (o *EzsignsignaturepaymentdetailRequestCompound) SetFkiGlaccountcontainerID(v int32)`

SetFkiGlaccountcontainerID sets FkiGlaccountcontainerID field to given value.

### HasFkiGlaccountcontainerID

`func (o *EzsignsignaturepaymentdetailRequestCompound) HasFkiGlaccountcontainerID() bool`

HasFkiGlaccountcontainerID returns a boolean if a field has been set.

### GetTEzsignsignaturepaymentdetailDescription

`func (o *EzsignsignaturepaymentdetailRequestCompound) GetTEzsignsignaturepaymentdetailDescription() string`

GetTEzsignsignaturepaymentdetailDescription returns the TEzsignsignaturepaymentdetailDescription field if non-nil, zero value otherwise.

### GetTEzsignsignaturepaymentdetailDescriptionOk

`func (o *EzsignsignaturepaymentdetailRequestCompound) GetTEzsignsignaturepaymentdetailDescriptionOk() (*string, bool)`

GetTEzsignsignaturepaymentdetailDescriptionOk returns a tuple with the TEzsignsignaturepaymentdetailDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignsignaturepaymentdetailDescription

`func (o *EzsignsignaturepaymentdetailRequestCompound) SetTEzsignsignaturepaymentdetailDescription(v string)`

SetTEzsignsignaturepaymentdetailDescription sets TEzsignsignaturepaymentdetailDescription field to given value.


### GetDEzsignsignaturepaymentdetailAmount

`func (o *EzsignsignaturepaymentdetailRequestCompound) GetDEzsignsignaturepaymentdetailAmount() string`

GetDEzsignsignaturepaymentdetailAmount returns the DEzsignsignaturepaymentdetailAmount field if non-nil, zero value otherwise.

### GetDEzsignsignaturepaymentdetailAmountOk

`func (o *EzsignsignaturepaymentdetailRequestCompound) GetDEzsignsignaturepaymentdetailAmountOk() (*string, bool)`

GetDEzsignsignaturepaymentdetailAmountOk returns a tuple with the DEzsignsignaturepaymentdetailAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzsignsignaturepaymentdetailAmount

`func (o *EzsignsignaturepaymentdetailRequestCompound) SetDEzsignsignaturepaymentdetailAmount(v string)`

SetDEzsignsignaturepaymentdetailAmount sets DEzsignsignaturepaymentdetailAmount field to given value.


### GetEEzsignsignaturepaymentdetailTaxable

`func (o *EzsignsignaturepaymentdetailRequestCompound) GetEEzsignsignaturepaymentdetailTaxable() FieldEEzsignsignaturepaymentdetailTaxable`

GetEEzsignsignaturepaymentdetailTaxable returns the EEzsignsignaturepaymentdetailTaxable field if non-nil, zero value otherwise.

### GetEEzsignsignaturepaymentdetailTaxableOk

`func (o *EzsignsignaturepaymentdetailRequestCompound) GetEEzsignsignaturepaymentdetailTaxableOk() (*FieldEEzsignsignaturepaymentdetailTaxable, bool)`

GetEEzsignsignaturepaymentdetailTaxableOk returns a tuple with the EEzsignsignaturepaymentdetailTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignaturepaymentdetailTaxable

`func (o *EzsignsignaturepaymentdetailRequestCompound) SetEEzsignsignaturepaymentdetailTaxable(v FieldEEzsignsignaturepaymentdetailTaxable)`

SetEEzsignsignaturepaymentdetailTaxable sets EEzsignsignaturepaymentdetailTaxable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


