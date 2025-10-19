# EzsigntemplatesignaturepaymentdetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatesignaturepaymentdetailID** | **int32** | The unique ID of the Ezsignsignaturepaymentdetail | 
**FkiGlaccountcontainerID** | Pointer to **int32** | The unique ID of the Glaccountcontainer | [optional] 
**TEzsigntemplatesignaturepaymentdetailDescription** | **string** | A description for the Ezsignsignaturepaymentdetail. | 
**DEzsigntemplatesignaturepaymentdetailAmount** | **string** | The amount of the for the Ezsignsignaturepaymentdetail | 
**EEzsigntemplatesignaturepaymentdetailTaxable** | [**FieldEEzsigntemplatesignaturepaymentdetailTaxable**](FieldEEzsigntemplatesignaturepaymentdetailTaxable.md) |  | 

## Methods

### NewEzsigntemplatesignaturepaymentdetailResponse

`func NewEzsigntemplatesignaturepaymentdetailResponse(pkiEzsigntemplatesignaturepaymentdetailID int32, tEzsigntemplatesignaturepaymentdetailDescription string, dEzsigntemplatesignaturepaymentdetailAmount string, eEzsigntemplatesignaturepaymentdetailTaxable FieldEEzsigntemplatesignaturepaymentdetailTaxable, ) *EzsigntemplatesignaturepaymentdetailResponse`

NewEzsigntemplatesignaturepaymentdetailResponse instantiates a new EzsigntemplatesignaturepaymentdetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignaturepaymentdetailResponseWithDefaults

`func NewEzsigntemplatesignaturepaymentdetailResponseWithDefaults() *EzsigntemplatesignaturepaymentdetailResponse`

NewEzsigntemplatesignaturepaymentdetailResponseWithDefaults instantiates a new EzsigntemplatesignaturepaymentdetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatesignaturepaymentdetailID

`func (o *EzsigntemplatesignaturepaymentdetailResponse) GetPkiEzsigntemplatesignaturepaymentdetailID() int32`

GetPkiEzsigntemplatesignaturepaymentdetailID returns the PkiEzsigntemplatesignaturepaymentdetailID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatesignaturepaymentdetailIDOk

`func (o *EzsigntemplatesignaturepaymentdetailResponse) GetPkiEzsigntemplatesignaturepaymentdetailIDOk() (*int32, bool)`

GetPkiEzsigntemplatesignaturepaymentdetailIDOk returns a tuple with the PkiEzsigntemplatesignaturepaymentdetailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatesignaturepaymentdetailID

`func (o *EzsigntemplatesignaturepaymentdetailResponse) SetPkiEzsigntemplatesignaturepaymentdetailID(v int32)`

SetPkiEzsigntemplatesignaturepaymentdetailID sets PkiEzsigntemplatesignaturepaymentdetailID field to given value.


### GetFkiGlaccountcontainerID

`func (o *EzsigntemplatesignaturepaymentdetailResponse) GetFkiGlaccountcontainerID() int32`

GetFkiGlaccountcontainerID returns the FkiGlaccountcontainerID field if non-nil, zero value otherwise.

### GetFkiGlaccountcontainerIDOk

`func (o *EzsigntemplatesignaturepaymentdetailResponse) GetFkiGlaccountcontainerIDOk() (*int32, bool)`

GetFkiGlaccountcontainerIDOk returns a tuple with the FkiGlaccountcontainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountcontainerID

`func (o *EzsigntemplatesignaturepaymentdetailResponse) SetFkiGlaccountcontainerID(v int32)`

SetFkiGlaccountcontainerID sets FkiGlaccountcontainerID field to given value.

### HasFkiGlaccountcontainerID

`func (o *EzsigntemplatesignaturepaymentdetailResponse) HasFkiGlaccountcontainerID() bool`

HasFkiGlaccountcontainerID returns a boolean if a field has been set.

### GetTEzsigntemplatesignaturepaymentdetailDescription

`func (o *EzsigntemplatesignaturepaymentdetailResponse) GetTEzsigntemplatesignaturepaymentdetailDescription() string`

GetTEzsigntemplatesignaturepaymentdetailDescription returns the TEzsigntemplatesignaturepaymentdetailDescription field if non-nil, zero value otherwise.

### GetTEzsigntemplatesignaturepaymentdetailDescriptionOk

`func (o *EzsigntemplatesignaturepaymentdetailResponse) GetTEzsigntemplatesignaturepaymentdetailDescriptionOk() (*string, bool)`

GetTEzsigntemplatesignaturepaymentdetailDescriptionOk returns a tuple with the TEzsigntemplatesignaturepaymentdetailDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplatesignaturepaymentdetailDescription

`func (o *EzsigntemplatesignaturepaymentdetailResponse) SetTEzsigntemplatesignaturepaymentdetailDescription(v string)`

SetTEzsigntemplatesignaturepaymentdetailDescription sets TEzsigntemplatesignaturepaymentdetailDescription field to given value.


### GetDEzsigntemplatesignaturepaymentdetailAmount

`func (o *EzsigntemplatesignaturepaymentdetailResponse) GetDEzsigntemplatesignaturepaymentdetailAmount() string`

GetDEzsigntemplatesignaturepaymentdetailAmount returns the DEzsigntemplatesignaturepaymentdetailAmount field if non-nil, zero value otherwise.

### GetDEzsigntemplatesignaturepaymentdetailAmountOk

`func (o *EzsigntemplatesignaturepaymentdetailResponse) GetDEzsigntemplatesignaturepaymentdetailAmountOk() (*string, bool)`

GetDEzsigntemplatesignaturepaymentdetailAmountOk returns a tuple with the DEzsigntemplatesignaturepaymentdetailAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzsigntemplatesignaturepaymentdetailAmount

`func (o *EzsigntemplatesignaturepaymentdetailResponse) SetDEzsigntemplatesignaturepaymentdetailAmount(v string)`

SetDEzsigntemplatesignaturepaymentdetailAmount sets DEzsigntemplatesignaturepaymentdetailAmount field to given value.


### GetEEzsigntemplatesignaturepaymentdetailTaxable

`func (o *EzsigntemplatesignaturepaymentdetailResponse) GetEEzsigntemplatesignaturepaymentdetailTaxable() FieldEEzsigntemplatesignaturepaymentdetailTaxable`

GetEEzsigntemplatesignaturepaymentdetailTaxable returns the EEzsigntemplatesignaturepaymentdetailTaxable field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignaturepaymentdetailTaxableOk

`func (o *EzsigntemplatesignaturepaymentdetailResponse) GetEEzsigntemplatesignaturepaymentdetailTaxableOk() (*FieldEEzsigntemplatesignaturepaymentdetailTaxable, bool)`

GetEEzsigntemplatesignaturepaymentdetailTaxableOk returns a tuple with the EEzsigntemplatesignaturepaymentdetailTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignaturepaymentdetailTaxable

`func (o *EzsigntemplatesignaturepaymentdetailResponse) SetEEzsigntemplatesignaturepaymentdetailTaxable(v FieldEEzsigntemplatesignaturepaymentdetailTaxable)`

SetEEzsigntemplatesignaturepaymentdetailTaxable sets EEzsigntemplatesignaturepaymentdetailTaxable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


