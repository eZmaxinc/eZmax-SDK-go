# EzsignsignaturepaymentdetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignsignaturepaymentdetailID** | **int32** | The unique ID of the Ezsignsignaturepaymentdetail | 
**FkiGlaccountcontainerID** | Pointer to **int32** | The unique ID of the Glaccountcontainer | [optional] 
**TEzsignsignaturepaymentdetailDescription** | **string** | A description for the Ezsignsignaturepaymentdetail. | 
**DEzsignsignaturepaymentdetailAmount** | **string** | The amount of the for the Ezsignsignaturepaymentdetail | 
**EEzsignsignaturepaymentdetailTaxable** | [**FieldEEzsignsignaturepaymentdetailTaxable**](FieldEEzsignsignaturepaymentdetailTaxable.md) |  | 

## Methods

### NewEzsignsignaturepaymentdetailResponse

`func NewEzsignsignaturepaymentdetailResponse(pkiEzsignsignaturepaymentdetailID int32, tEzsignsignaturepaymentdetailDescription string, dEzsignsignaturepaymentdetailAmount string, eEzsignsignaturepaymentdetailTaxable FieldEEzsignsignaturepaymentdetailTaxable, ) *EzsignsignaturepaymentdetailResponse`

NewEzsignsignaturepaymentdetailResponse instantiates a new EzsignsignaturepaymentdetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignaturepaymentdetailResponseWithDefaults

`func NewEzsignsignaturepaymentdetailResponseWithDefaults() *EzsignsignaturepaymentdetailResponse`

NewEzsignsignaturepaymentdetailResponseWithDefaults instantiates a new EzsignsignaturepaymentdetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignsignaturepaymentdetailID

`func (o *EzsignsignaturepaymentdetailResponse) GetPkiEzsignsignaturepaymentdetailID() int32`

GetPkiEzsignsignaturepaymentdetailID returns the PkiEzsignsignaturepaymentdetailID field if non-nil, zero value otherwise.

### GetPkiEzsignsignaturepaymentdetailIDOk

`func (o *EzsignsignaturepaymentdetailResponse) GetPkiEzsignsignaturepaymentdetailIDOk() (*int32, bool)`

GetPkiEzsignsignaturepaymentdetailIDOk returns a tuple with the PkiEzsignsignaturepaymentdetailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignsignaturepaymentdetailID

`func (o *EzsignsignaturepaymentdetailResponse) SetPkiEzsignsignaturepaymentdetailID(v int32)`

SetPkiEzsignsignaturepaymentdetailID sets PkiEzsignsignaturepaymentdetailID field to given value.


### GetFkiGlaccountcontainerID

`func (o *EzsignsignaturepaymentdetailResponse) GetFkiGlaccountcontainerID() int32`

GetFkiGlaccountcontainerID returns the FkiGlaccountcontainerID field if non-nil, zero value otherwise.

### GetFkiGlaccountcontainerIDOk

`func (o *EzsignsignaturepaymentdetailResponse) GetFkiGlaccountcontainerIDOk() (*int32, bool)`

GetFkiGlaccountcontainerIDOk returns a tuple with the FkiGlaccountcontainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountcontainerID

`func (o *EzsignsignaturepaymentdetailResponse) SetFkiGlaccountcontainerID(v int32)`

SetFkiGlaccountcontainerID sets FkiGlaccountcontainerID field to given value.

### HasFkiGlaccountcontainerID

`func (o *EzsignsignaturepaymentdetailResponse) HasFkiGlaccountcontainerID() bool`

HasFkiGlaccountcontainerID returns a boolean if a field has been set.

### GetTEzsignsignaturepaymentdetailDescription

`func (o *EzsignsignaturepaymentdetailResponse) GetTEzsignsignaturepaymentdetailDescription() string`

GetTEzsignsignaturepaymentdetailDescription returns the TEzsignsignaturepaymentdetailDescription field if non-nil, zero value otherwise.

### GetTEzsignsignaturepaymentdetailDescriptionOk

`func (o *EzsignsignaturepaymentdetailResponse) GetTEzsignsignaturepaymentdetailDescriptionOk() (*string, bool)`

GetTEzsignsignaturepaymentdetailDescriptionOk returns a tuple with the TEzsignsignaturepaymentdetailDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignsignaturepaymentdetailDescription

`func (o *EzsignsignaturepaymentdetailResponse) SetTEzsignsignaturepaymentdetailDescription(v string)`

SetTEzsignsignaturepaymentdetailDescription sets TEzsignsignaturepaymentdetailDescription field to given value.


### GetDEzsignsignaturepaymentdetailAmount

`func (o *EzsignsignaturepaymentdetailResponse) GetDEzsignsignaturepaymentdetailAmount() string`

GetDEzsignsignaturepaymentdetailAmount returns the DEzsignsignaturepaymentdetailAmount field if non-nil, zero value otherwise.

### GetDEzsignsignaturepaymentdetailAmountOk

`func (o *EzsignsignaturepaymentdetailResponse) GetDEzsignsignaturepaymentdetailAmountOk() (*string, bool)`

GetDEzsignsignaturepaymentdetailAmountOk returns a tuple with the DEzsignsignaturepaymentdetailAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzsignsignaturepaymentdetailAmount

`func (o *EzsignsignaturepaymentdetailResponse) SetDEzsignsignaturepaymentdetailAmount(v string)`

SetDEzsignsignaturepaymentdetailAmount sets DEzsignsignaturepaymentdetailAmount field to given value.


### GetEEzsignsignaturepaymentdetailTaxable

`func (o *EzsignsignaturepaymentdetailResponse) GetEEzsignsignaturepaymentdetailTaxable() FieldEEzsignsignaturepaymentdetailTaxable`

GetEEzsignsignaturepaymentdetailTaxable returns the EEzsignsignaturepaymentdetailTaxable field if non-nil, zero value otherwise.

### GetEEzsignsignaturepaymentdetailTaxableOk

`func (o *EzsignsignaturepaymentdetailResponse) GetEEzsignsignaturepaymentdetailTaxableOk() (*FieldEEzsignsignaturepaymentdetailTaxable, bool)`

GetEEzsignsignaturepaymentdetailTaxableOk returns a tuple with the EEzsignsignaturepaymentdetailTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignaturepaymentdetailTaxable

`func (o *EzsignsignaturepaymentdetailResponse) SetEEzsignsignaturepaymentdetailTaxable(v FieldEEzsignsignaturepaymentdetailTaxable)`

SetEEzsignsignaturepaymentdetailTaxable sets EEzsignsignaturepaymentdetailTaxable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


