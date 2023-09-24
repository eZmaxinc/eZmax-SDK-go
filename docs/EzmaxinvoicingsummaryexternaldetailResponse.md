# EzmaxinvoicingsummaryexternaldetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzmaxinvoicingsummaryexternaldetailID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicingsummaryexternaldetail | [optional] 
**FkiEzmaxinvoicingsummaryexternalID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicingsummaryexternal | [optional] 
**FkiEzmaxproductID** | **int32** | The unique ID of the Ezmaxproduct | 
**SEzmaxproductDescriptionX** | **string** | The description of the Ezmaxproduct in the language of the requester | 
**DEzmaxinvoicingsummaryexternaldetailCountreal** | **string** | The count item invoiced for the product | 
**DEzmaxinvoicingsummaryexternaldetailSubtotal** | **string** | The subtotal invoiced for the product | 
**DEzmaxinvoicingsummaryexternaldetailRebate** | **string** | The rebate for the product | 
**DEzmaxinvoicingsummaryexternaldetailTotal** | **string** | The total invoiced for the product | 
**BEzmaxinvoicingsummaryexternaldetailAdjustment** | **bool** | Whether it&#39;s an adjustment | 
**TEzmaxproductHelpX** | **string** | The help message of the Ezmaxproduct in the language of the requester | 

## Methods

### NewEzmaxinvoicingsummaryexternaldetailResponse

`func NewEzmaxinvoicingsummaryexternaldetailResponse(fkiEzmaxproductID int32, sEzmaxproductDescriptionX string, dEzmaxinvoicingsummaryexternaldetailCountreal string, dEzmaxinvoicingsummaryexternaldetailSubtotal string, dEzmaxinvoicingsummaryexternaldetailRebate string, dEzmaxinvoicingsummaryexternaldetailTotal string, bEzmaxinvoicingsummaryexternaldetailAdjustment bool, tEzmaxproductHelpX string, ) *EzmaxinvoicingsummaryexternaldetailResponse`

NewEzmaxinvoicingsummaryexternaldetailResponse instantiates a new EzmaxinvoicingsummaryexternaldetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingsummaryexternaldetailResponseWithDefaults

`func NewEzmaxinvoicingsummaryexternaldetailResponseWithDefaults() *EzmaxinvoicingsummaryexternaldetailResponse`

NewEzmaxinvoicingsummaryexternaldetailResponseWithDefaults instantiates a new EzmaxinvoicingsummaryexternaldetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicingsummaryexternaldetailID

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetPkiEzmaxinvoicingsummaryexternaldetailID() int32`

GetPkiEzmaxinvoicingsummaryexternaldetailID returns the PkiEzmaxinvoicingsummaryexternaldetailID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicingsummaryexternaldetailIDOk

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetPkiEzmaxinvoicingsummaryexternaldetailIDOk() (*int32, bool)`

GetPkiEzmaxinvoicingsummaryexternaldetailIDOk returns a tuple with the PkiEzmaxinvoicingsummaryexternaldetailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicingsummaryexternaldetailID

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetPkiEzmaxinvoicingsummaryexternaldetailID(v int32)`

SetPkiEzmaxinvoicingsummaryexternaldetailID sets PkiEzmaxinvoicingsummaryexternaldetailID field to given value.

### HasPkiEzmaxinvoicingsummaryexternaldetailID

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) HasPkiEzmaxinvoicingsummaryexternaldetailID() bool`

HasPkiEzmaxinvoicingsummaryexternaldetailID returns a boolean if a field has been set.

### GetFkiEzmaxinvoicingsummaryexternalID

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetFkiEzmaxinvoicingsummaryexternalID() int32`

GetFkiEzmaxinvoicingsummaryexternalID returns the FkiEzmaxinvoicingsummaryexternalID field if non-nil, zero value otherwise.

### GetFkiEzmaxinvoicingsummaryexternalIDOk

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetFkiEzmaxinvoicingsummaryexternalIDOk() (*int32, bool)`

GetFkiEzmaxinvoicingsummaryexternalIDOk returns a tuple with the FkiEzmaxinvoicingsummaryexternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxinvoicingsummaryexternalID

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetFkiEzmaxinvoicingsummaryexternalID(v int32)`

SetFkiEzmaxinvoicingsummaryexternalID sets FkiEzmaxinvoicingsummaryexternalID field to given value.

### HasFkiEzmaxinvoicingsummaryexternalID

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) HasFkiEzmaxinvoicingsummaryexternalID() bool`

HasFkiEzmaxinvoicingsummaryexternalID returns a boolean if a field has been set.

### GetFkiEzmaxproductID

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetFkiEzmaxproductID() int32`

GetFkiEzmaxproductID returns the FkiEzmaxproductID field if non-nil, zero value otherwise.

### GetFkiEzmaxproductIDOk

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetFkiEzmaxproductIDOk() (*int32, bool)`

GetFkiEzmaxproductIDOk returns a tuple with the FkiEzmaxproductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxproductID

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetFkiEzmaxproductID(v int32)`

SetFkiEzmaxproductID sets FkiEzmaxproductID field to given value.


### GetSEzmaxproductDescriptionX

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetSEzmaxproductDescriptionX() string`

GetSEzmaxproductDescriptionX returns the SEzmaxproductDescriptionX field if non-nil, zero value otherwise.

### GetSEzmaxproductDescriptionXOk

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetSEzmaxproductDescriptionXOk() (*string, bool)`

GetSEzmaxproductDescriptionXOk returns a tuple with the SEzmaxproductDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxproductDescriptionX

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetSEzmaxproductDescriptionX(v string)`

SetSEzmaxproductDescriptionX sets SEzmaxproductDescriptionX field to given value.


### GetDEzmaxinvoicingsummaryexternaldetailCountreal

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailCountreal() string`

GetDEzmaxinvoicingsummaryexternaldetailCountreal returns the DEzmaxinvoicingsummaryexternaldetailCountreal field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryexternaldetailCountrealOk

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailCountrealOk() (*string, bool)`

GetDEzmaxinvoicingsummaryexternaldetailCountrealOk returns a tuple with the DEzmaxinvoicingsummaryexternaldetailCountreal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryexternaldetailCountreal

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetDEzmaxinvoicingsummaryexternaldetailCountreal(v string)`

SetDEzmaxinvoicingsummaryexternaldetailCountreal sets DEzmaxinvoicingsummaryexternaldetailCountreal field to given value.


### GetDEzmaxinvoicingsummaryexternaldetailSubtotal

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailSubtotal() string`

GetDEzmaxinvoicingsummaryexternaldetailSubtotal returns the DEzmaxinvoicingsummaryexternaldetailSubtotal field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryexternaldetailSubtotalOk

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailSubtotalOk() (*string, bool)`

GetDEzmaxinvoicingsummaryexternaldetailSubtotalOk returns a tuple with the DEzmaxinvoicingsummaryexternaldetailSubtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryexternaldetailSubtotal

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetDEzmaxinvoicingsummaryexternaldetailSubtotal(v string)`

SetDEzmaxinvoicingsummaryexternaldetailSubtotal sets DEzmaxinvoicingsummaryexternaldetailSubtotal field to given value.


### GetDEzmaxinvoicingsummaryexternaldetailRebate

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailRebate() string`

GetDEzmaxinvoicingsummaryexternaldetailRebate returns the DEzmaxinvoicingsummaryexternaldetailRebate field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryexternaldetailRebateOk

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailRebateOk() (*string, bool)`

GetDEzmaxinvoicingsummaryexternaldetailRebateOk returns a tuple with the DEzmaxinvoicingsummaryexternaldetailRebate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryexternaldetailRebate

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetDEzmaxinvoicingsummaryexternaldetailRebate(v string)`

SetDEzmaxinvoicingsummaryexternaldetailRebate sets DEzmaxinvoicingsummaryexternaldetailRebate field to given value.


### GetDEzmaxinvoicingsummaryexternaldetailTotal

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailTotal() string`

GetDEzmaxinvoicingsummaryexternaldetailTotal returns the DEzmaxinvoicingsummaryexternaldetailTotal field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryexternaldetailTotalOk

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetDEzmaxinvoicingsummaryexternaldetailTotalOk() (*string, bool)`

GetDEzmaxinvoicingsummaryexternaldetailTotalOk returns a tuple with the DEzmaxinvoicingsummaryexternaldetailTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryexternaldetailTotal

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetDEzmaxinvoicingsummaryexternaldetailTotal(v string)`

SetDEzmaxinvoicingsummaryexternaldetailTotal sets DEzmaxinvoicingsummaryexternaldetailTotal field to given value.


### GetBEzmaxinvoicingsummaryexternaldetailAdjustment

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetBEzmaxinvoicingsummaryexternaldetailAdjustment() bool`

GetBEzmaxinvoicingsummaryexternaldetailAdjustment returns the BEzmaxinvoicingsummaryexternaldetailAdjustment field if non-nil, zero value otherwise.

### GetBEzmaxinvoicingsummaryexternaldetailAdjustmentOk

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetBEzmaxinvoicingsummaryexternaldetailAdjustmentOk() (*bool, bool)`

GetBEzmaxinvoicingsummaryexternaldetailAdjustmentOk returns a tuple with the BEzmaxinvoicingsummaryexternaldetailAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicingsummaryexternaldetailAdjustment

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetBEzmaxinvoicingsummaryexternaldetailAdjustment(v bool)`

SetBEzmaxinvoicingsummaryexternaldetailAdjustment sets BEzmaxinvoicingsummaryexternaldetailAdjustment field to given value.


### GetTEzmaxproductHelpX

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetTEzmaxproductHelpX() string`

GetTEzmaxproductHelpX returns the TEzmaxproductHelpX field if non-nil, zero value otherwise.

### GetTEzmaxproductHelpXOk

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) GetTEzmaxproductHelpXOk() (*string, bool)`

GetTEzmaxproductHelpXOk returns a tuple with the TEzmaxproductHelpX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzmaxproductHelpX

`func (o *EzmaxinvoicingsummaryexternaldetailResponse) SetTEzmaxproductHelpX(v string)`

SetTEzmaxproductHelpX sets TEzmaxproductHelpX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


