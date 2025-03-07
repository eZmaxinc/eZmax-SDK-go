# EzsignsignaturepaymentdetailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignsignaturepaymentdetailID** | Pointer to **int32** | The unique ID of the Ezsignsignaturepaymentdetail | [optional] 
**FkiGlaccountcontainerID** | Pointer to **int32** | The unique ID of the Glaccountcontainer | [optional] 
**TEzsignsignaturepaymentdetailDescription** | **string** | A description for the Ezsignsignaturepaymentdetail. | 
**DEzsignsignaturepaymentdetailAmount** | **string** | The amount of the for the Ezsignsignaturepaymentdetail | 
**EEzsignsignaturepaymentdetailTaxable** | [**FieldEEzsignsignaturepaymentdetailTaxable**](FieldEEzsignsignaturepaymentdetailTaxable.md) |  | 

## Methods

### NewEzsignsignaturepaymentdetailRequest

`func NewEzsignsignaturepaymentdetailRequest(tEzsignsignaturepaymentdetailDescription string, dEzsignsignaturepaymentdetailAmount string, eEzsignsignaturepaymentdetailTaxable FieldEEzsignsignaturepaymentdetailTaxable, ) *EzsignsignaturepaymentdetailRequest`

NewEzsignsignaturepaymentdetailRequest instantiates a new EzsignsignaturepaymentdetailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignaturepaymentdetailRequestWithDefaults

`func NewEzsignsignaturepaymentdetailRequestWithDefaults() *EzsignsignaturepaymentdetailRequest`

NewEzsignsignaturepaymentdetailRequestWithDefaults instantiates a new EzsignsignaturepaymentdetailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignsignaturepaymentdetailID

`func (o *EzsignsignaturepaymentdetailRequest) GetPkiEzsignsignaturepaymentdetailID() int32`

GetPkiEzsignsignaturepaymentdetailID returns the PkiEzsignsignaturepaymentdetailID field if non-nil, zero value otherwise.

### GetPkiEzsignsignaturepaymentdetailIDOk

`func (o *EzsignsignaturepaymentdetailRequest) GetPkiEzsignsignaturepaymentdetailIDOk() (*int32, bool)`

GetPkiEzsignsignaturepaymentdetailIDOk returns a tuple with the PkiEzsignsignaturepaymentdetailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignsignaturepaymentdetailID

`func (o *EzsignsignaturepaymentdetailRequest) SetPkiEzsignsignaturepaymentdetailID(v int32)`

SetPkiEzsignsignaturepaymentdetailID sets PkiEzsignsignaturepaymentdetailID field to given value.

### HasPkiEzsignsignaturepaymentdetailID

`func (o *EzsignsignaturepaymentdetailRequest) HasPkiEzsignsignaturepaymentdetailID() bool`

HasPkiEzsignsignaturepaymentdetailID returns a boolean if a field has been set.

### GetFkiGlaccountcontainerID

`func (o *EzsignsignaturepaymentdetailRequest) GetFkiGlaccountcontainerID() int32`

GetFkiGlaccountcontainerID returns the FkiGlaccountcontainerID field if non-nil, zero value otherwise.

### GetFkiGlaccountcontainerIDOk

`func (o *EzsignsignaturepaymentdetailRequest) GetFkiGlaccountcontainerIDOk() (*int32, bool)`

GetFkiGlaccountcontainerIDOk returns a tuple with the FkiGlaccountcontainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountcontainerID

`func (o *EzsignsignaturepaymentdetailRequest) SetFkiGlaccountcontainerID(v int32)`

SetFkiGlaccountcontainerID sets FkiGlaccountcontainerID field to given value.

### HasFkiGlaccountcontainerID

`func (o *EzsignsignaturepaymentdetailRequest) HasFkiGlaccountcontainerID() bool`

HasFkiGlaccountcontainerID returns a boolean if a field has been set.

### GetTEzsignsignaturepaymentdetailDescription

`func (o *EzsignsignaturepaymentdetailRequest) GetTEzsignsignaturepaymentdetailDescription() string`

GetTEzsignsignaturepaymentdetailDescription returns the TEzsignsignaturepaymentdetailDescription field if non-nil, zero value otherwise.

### GetTEzsignsignaturepaymentdetailDescriptionOk

`func (o *EzsignsignaturepaymentdetailRequest) GetTEzsignsignaturepaymentdetailDescriptionOk() (*string, bool)`

GetTEzsignsignaturepaymentdetailDescriptionOk returns a tuple with the TEzsignsignaturepaymentdetailDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignsignaturepaymentdetailDescription

`func (o *EzsignsignaturepaymentdetailRequest) SetTEzsignsignaturepaymentdetailDescription(v string)`

SetTEzsignsignaturepaymentdetailDescription sets TEzsignsignaturepaymentdetailDescription field to given value.


### GetDEzsignsignaturepaymentdetailAmount

`func (o *EzsignsignaturepaymentdetailRequest) GetDEzsignsignaturepaymentdetailAmount() string`

GetDEzsignsignaturepaymentdetailAmount returns the DEzsignsignaturepaymentdetailAmount field if non-nil, zero value otherwise.

### GetDEzsignsignaturepaymentdetailAmountOk

`func (o *EzsignsignaturepaymentdetailRequest) GetDEzsignsignaturepaymentdetailAmountOk() (*string, bool)`

GetDEzsignsignaturepaymentdetailAmountOk returns a tuple with the DEzsignsignaturepaymentdetailAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzsignsignaturepaymentdetailAmount

`func (o *EzsignsignaturepaymentdetailRequest) SetDEzsignsignaturepaymentdetailAmount(v string)`

SetDEzsignsignaturepaymentdetailAmount sets DEzsignsignaturepaymentdetailAmount field to given value.


### GetEEzsignsignaturepaymentdetailTaxable

`func (o *EzsignsignaturepaymentdetailRequest) GetEEzsignsignaturepaymentdetailTaxable() FieldEEzsignsignaturepaymentdetailTaxable`

GetEEzsignsignaturepaymentdetailTaxable returns the EEzsignsignaturepaymentdetailTaxable field if non-nil, zero value otherwise.

### GetEEzsignsignaturepaymentdetailTaxableOk

`func (o *EzsignsignaturepaymentdetailRequest) GetEEzsignsignaturepaymentdetailTaxableOk() (*FieldEEzsignsignaturepaymentdetailTaxable, bool)`

GetEEzsignsignaturepaymentdetailTaxableOk returns a tuple with the EEzsignsignaturepaymentdetailTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignaturepaymentdetailTaxable

`func (o *EzsignsignaturepaymentdetailRequest) SetEEzsignsignaturepaymentdetailTaxable(v FieldEEzsignsignaturepaymentdetailTaxable)`

SetEEzsignsignaturepaymentdetailTaxable sets EEzsignsignaturepaymentdetailTaxable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


