# EzmaxinvoicingsummaryinternaldetailResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzmaxinvoicingsummaryinternaldetailID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicingsummaryinternaldetail | [optional] 
**FkiEzmaxinvoicingsummaryinternalID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicingsummaryinternal | [optional] 
**FkiEzmaxproductID** | **int32** | The unique ID of the Ezmaxproduct | 
**SEzmaxproductDescriptionX** | **string** | The description of the Ezmaxproduct in the language of the requester | 
**FkiBillingentityexternalID** | **int32** | The unique ID of the Billingentityexternal | 
**SBillingentityexternalDescription** | **string** | The description of the Billingentityexternal | 
**DEzmaxinvoicingsummaryinternaldetailCountreal** | **string** | The count item invoiced for the product | 
**DEzmaxinvoicingsummaryinternaldetailSubtotal** | **string** | The subtotal invoiced for the product | 
**DEzmaxinvoicingsummaryinternaldetailRebate** | **string** | The rebate for the product | 
**DEzmaxinvoicingsummaryinternaldetailTotal** | **string** | The total invoiced for the product | 
**BEzmaxinvoicingsummaryinternaldetailAdjustment** | **bool** | Whether if it&#39;s an adjustment | 
**TEzmaxproductHelpX** | **string** | The help message of the Ezmaxproduct in the language of the requester | 

## Methods

### NewEzmaxinvoicingsummaryinternaldetailResponseCompound

`func NewEzmaxinvoicingsummaryinternaldetailResponseCompound(fkiEzmaxproductID int32, sEzmaxproductDescriptionX string, fkiBillingentityexternalID int32, sBillingentityexternalDescription string, dEzmaxinvoicingsummaryinternaldetailCountreal string, dEzmaxinvoicingsummaryinternaldetailSubtotal string, dEzmaxinvoicingsummaryinternaldetailRebate string, dEzmaxinvoicingsummaryinternaldetailTotal string, bEzmaxinvoicingsummaryinternaldetailAdjustment bool, tEzmaxproductHelpX string, ) *EzmaxinvoicingsummaryinternaldetailResponseCompound`

NewEzmaxinvoicingsummaryinternaldetailResponseCompound instantiates a new EzmaxinvoicingsummaryinternaldetailResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingsummaryinternaldetailResponseCompoundWithDefaults

`func NewEzmaxinvoicingsummaryinternaldetailResponseCompoundWithDefaults() *EzmaxinvoicingsummaryinternaldetailResponseCompound`

NewEzmaxinvoicingsummaryinternaldetailResponseCompoundWithDefaults instantiates a new EzmaxinvoicingsummaryinternaldetailResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicingsummaryinternaldetailID

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetPkiEzmaxinvoicingsummaryinternaldetailID() int32`

GetPkiEzmaxinvoicingsummaryinternaldetailID returns the PkiEzmaxinvoicingsummaryinternaldetailID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicingsummaryinternaldetailIDOk

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetPkiEzmaxinvoicingsummaryinternaldetailIDOk() (*int32, bool)`

GetPkiEzmaxinvoicingsummaryinternaldetailIDOk returns a tuple with the PkiEzmaxinvoicingsummaryinternaldetailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicingsummaryinternaldetailID

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) SetPkiEzmaxinvoicingsummaryinternaldetailID(v int32)`

SetPkiEzmaxinvoicingsummaryinternaldetailID sets PkiEzmaxinvoicingsummaryinternaldetailID field to given value.

### HasPkiEzmaxinvoicingsummaryinternaldetailID

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) HasPkiEzmaxinvoicingsummaryinternaldetailID() bool`

HasPkiEzmaxinvoicingsummaryinternaldetailID returns a boolean if a field has been set.

### GetFkiEzmaxinvoicingsummaryinternalID

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetFkiEzmaxinvoicingsummaryinternalID() int32`

GetFkiEzmaxinvoicingsummaryinternalID returns the FkiEzmaxinvoicingsummaryinternalID field if non-nil, zero value otherwise.

### GetFkiEzmaxinvoicingsummaryinternalIDOk

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetFkiEzmaxinvoicingsummaryinternalIDOk() (*int32, bool)`

GetFkiEzmaxinvoicingsummaryinternalIDOk returns a tuple with the FkiEzmaxinvoicingsummaryinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxinvoicingsummaryinternalID

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) SetFkiEzmaxinvoicingsummaryinternalID(v int32)`

SetFkiEzmaxinvoicingsummaryinternalID sets FkiEzmaxinvoicingsummaryinternalID field to given value.

### HasFkiEzmaxinvoicingsummaryinternalID

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) HasFkiEzmaxinvoicingsummaryinternalID() bool`

HasFkiEzmaxinvoicingsummaryinternalID returns a boolean if a field has been set.

### GetFkiEzmaxproductID

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetFkiEzmaxproductID() int32`

GetFkiEzmaxproductID returns the FkiEzmaxproductID field if non-nil, zero value otherwise.

### GetFkiEzmaxproductIDOk

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetFkiEzmaxproductIDOk() (*int32, bool)`

GetFkiEzmaxproductIDOk returns a tuple with the FkiEzmaxproductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxproductID

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) SetFkiEzmaxproductID(v int32)`

SetFkiEzmaxproductID sets FkiEzmaxproductID field to given value.


### GetSEzmaxproductDescriptionX

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetSEzmaxproductDescriptionX() string`

GetSEzmaxproductDescriptionX returns the SEzmaxproductDescriptionX field if non-nil, zero value otherwise.

### GetSEzmaxproductDescriptionXOk

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetSEzmaxproductDescriptionXOk() (*string, bool)`

GetSEzmaxproductDescriptionXOk returns a tuple with the SEzmaxproductDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxproductDescriptionX

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) SetSEzmaxproductDescriptionX(v string)`

SetSEzmaxproductDescriptionX sets SEzmaxproductDescriptionX field to given value.


### GetFkiBillingentityexternalID

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetFkiBillingentityexternalID() int32`

GetFkiBillingentityexternalID returns the FkiBillingentityexternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityexternalIDOk

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetFkiBillingentityexternalIDOk() (*int32, bool)`

GetFkiBillingentityexternalIDOk returns a tuple with the FkiBillingentityexternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityexternalID

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) SetFkiBillingentityexternalID(v int32)`

SetFkiBillingentityexternalID sets FkiBillingentityexternalID field to given value.


### GetSBillingentityexternalDescription

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetSBillingentityexternalDescription() string`

GetSBillingentityexternalDescription returns the SBillingentityexternalDescription field if non-nil, zero value otherwise.

### GetSBillingentityexternalDescriptionOk

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetSBillingentityexternalDescriptionOk() (*string, bool)`

GetSBillingentityexternalDescriptionOk returns a tuple with the SBillingentityexternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBillingentityexternalDescription

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) SetSBillingentityexternalDescription(v string)`

SetSBillingentityexternalDescription sets SBillingentityexternalDescription field to given value.


### GetDEzmaxinvoicingsummaryinternaldetailCountreal

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetDEzmaxinvoicingsummaryinternaldetailCountreal() string`

GetDEzmaxinvoicingsummaryinternaldetailCountreal returns the DEzmaxinvoicingsummaryinternaldetailCountreal field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryinternaldetailCountrealOk

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetDEzmaxinvoicingsummaryinternaldetailCountrealOk() (*string, bool)`

GetDEzmaxinvoicingsummaryinternaldetailCountrealOk returns a tuple with the DEzmaxinvoicingsummaryinternaldetailCountreal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryinternaldetailCountreal

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) SetDEzmaxinvoicingsummaryinternaldetailCountreal(v string)`

SetDEzmaxinvoicingsummaryinternaldetailCountreal sets DEzmaxinvoicingsummaryinternaldetailCountreal field to given value.


### GetDEzmaxinvoicingsummaryinternaldetailSubtotal

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetDEzmaxinvoicingsummaryinternaldetailSubtotal() string`

GetDEzmaxinvoicingsummaryinternaldetailSubtotal returns the DEzmaxinvoicingsummaryinternaldetailSubtotal field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryinternaldetailSubtotalOk

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetDEzmaxinvoicingsummaryinternaldetailSubtotalOk() (*string, bool)`

GetDEzmaxinvoicingsummaryinternaldetailSubtotalOk returns a tuple with the DEzmaxinvoicingsummaryinternaldetailSubtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryinternaldetailSubtotal

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) SetDEzmaxinvoicingsummaryinternaldetailSubtotal(v string)`

SetDEzmaxinvoicingsummaryinternaldetailSubtotal sets DEzmaxinvoicingsummaryinternaldetailSubtotal field to given value.


### GetDEzmaxinvoicingsummaryinternaldetailRebate

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetDEzmaxinvoicingsummaryinternaldetailRebate() string`

GetDEzmaxinvoicingsummaryinternaldetailRebate returns the DEzmaxinvoicingsummaryinternaldetailRebate field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryinternaldetailRebateOk

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetDEzmaxinvoicingsummaryinternaldetailRebateOk() (*string, bool)`

GetDEzmaxinvoicingsummaryinternaldetailRebateOk returns a tuple with the DEzmaxinvoicingsummaryinternaldetailRebate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryinternaldetailRebate

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) SetDEzmaxinvoicingsummaryinternaldetailRebate(v string)`

SetDEzmaxinvoicingsummaryinternaldetailRebate sets DEzmaxinvoicingsummaryinternaldetailRebate field to given value.


### GetDEzmaxinvoicingsummaryinternaldetailTotal

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetDEzmaxinvoicingsummaryinternaldetailTotal() string`

GetDEzmaxinvoicingsummaryinternaldetailTotal returns the DEzmaxinvoicingsummaryinternaldetailTotal field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryinternaldetailTotalOk

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetDEzmaxinvoicingsummaryinternaldetailTotalOk() (*string, bool)`

GetDEzmaxinvoicingsummaryinternaldetailTotalOk returns a tuple with the DEzmaxinvoicingsummaryinternaldetailTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryinternaldetailTotal

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) SetDEzmaxinvoicingsummaryinternaldetailTotal(v string)`

SetDEzmaxinvoicingsummaryinternaldetailTotal sets DEzmaxinvoicingsummaryinternaldetailTotal field to given value.


### GetBEzmaxinvoicingsummaryinternaldetailAdjustment

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetBEzmaxinvoicingsummaryinternaldetailAdjustment() bool`

GetBEzmaxinvoicingsummaryinternaldetailAdjustment returns the BEzmaxinvoicingsummaryinternaldetailAdjustment field if non-nil, zero value otherwise.

### GetBEzmaxinvoicingsummaryinternaldetailAdjustmentOk

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetBEzmaxinvoicingsummaryinternaldetailAdjustmentOk() (*bool, bool)`

GetBEzmaxinvoicingsummaryinternaldetailAdjustmentOk returns a tuple with the BEzmaxinvoicingsummaryinternaldetailAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicingsummaryinternaldetailAdjustment

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) SetBEzmaxinvoicingsummaryinternaldetailAdjustment(v bool)`

SetBEzmaxinvoicingsummaryinternaldetailAdjustment sets BEzmaxinvoicingsummaryinternaldetailAdjustment field to given value.


### GetTEzmaxproductHelpX

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetTEzmaxproductHelpX() string`

GetTEzmaxproductHelpX returns the TEzmaxproductHelpX field if non-nil, zero value otherwise.

### GetTEzmaxproductHelpXOk

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) GetTEzmaxproductHelpXOk() (*string, bool)`

GetTEzmaxproductHelpXOk returns a tuple with the TEzmaxproductHelpX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzmaxproductHelpX

`func (o *EzmaxinvoicingsummaryinternaldetailResponseCompound) SetTEzmaxproductHelpX(v string)`

SetTEzmaxproductHelpX sets TEzmaxproductHelpX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


