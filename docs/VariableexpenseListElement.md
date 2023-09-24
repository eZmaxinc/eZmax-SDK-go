# VariableexpenseListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiVariableexpenseID** | **int32** | The unique ID of the Variableexpense | 
**SVariableexpenseCode** | Pointer to **string** | The code of the Variableexpense | [optional] 
**SVariableexpenseDescriptionX** | Pointer to **string** | The description of the Variableexpense in the language of the requester | [optional] 
**EVariableexpenseTaxable** | Pointer to [**FieldEVariableexpenseTaxable**](FieldEVariableexpenseTaxable.md) |  | [optional] 
**BVariableexpenseIsactive** | Pointer to **bool** | Whether the variableexpense is active or not | [optional] 

## Methods

### NewVariableexpenseListElement

`func NewVariableexpenseListElement(pkiVariableexpenseID int32, ) *VariableexpenseListElement`

NewVariableexpenseListElement instantiates a new VariableexpenseListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableexpenseListElementWithDefaults

`func NewVariableexpenseListElementWithDefaults() *VariableexpenseListElement`

NewVariableexpenseListElementWithDefaults instantiates a new VariableexpenseListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiVariableexpenseID

`func (o *VariableexpenseListElement) GetPkiVariableexpenseID() int32`

GetPkiVariableexpenseID returns the PkiVariableexpenseID field if non-nil, zero value otherwise.

### GetPkiVariableexpenseIDOk

`func (o *VariableexpenseListElement) GetPkiVariableexpenseIDOk() (*int32, bool)`

GetPkiVariableexpenseIDOk returns a tuple with the PkiVariableexpenseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiVariableexpenseID

`func (o *VariableexpenseListElement) SetPkiVariableexpenseID(v int32)`

SetPkiVariableexpenseID sets PkiVariableexpenseID field to given value.


### GetSVariableexpenseCode

`func (o *VariableexpenseListElement) GetSVariableexpenseCode() string`

GetSVariableexpenseCode returns the SVariableexpenseCode field if non-nil, zero value otherwise.

### GetSVariableexpenseCodeOk

`func (o *VariableexpenseListElement) GetSVariableexpenseCodeOk() (*string, bool)`

GetSVariableexpenseCodeOk returns a tuple with the SVariableexpenseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSVariableexpenseCode

`func (o *VariableexpenseListElement) SetSVariableexpenseCode(v string)`

SetSVariableexpenseCode sets SVariableexpenseCode field to given value.

### HasSVariableexpenseCode

`func (o *VariableexpenseListElement) HasSVariableexpenseCode() bool`

HasSVariableexpenseCode returns a boolean if a field has been set.

### GetSVariableexpenseDescriptionX

`func (o *VariableexpenseListElement) GetSVariableexpenseDescriptionX() string`

GetSVariableexpenseDescriptionX returns the SVariableexpenseDescriptionX field if non-nil, zero value otherwise.

### GetSVariableexpenseDescriptionXOk

`func (o *VariableexpenseListElement) GetSVariableexpenseDescriptionXOk() (*string, bool)`

GetSVariableexpenseDescriptionXOk returns a tuple with the SVariableexpenseDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSVariableexpenseDescriptionX

`func (o *VariableexpenseListElement) SetSVariableexpenseDescriptionX(v string)`

SetSVariableexpenseDescriptionX sets SVariableexpenseDescriptionX field to given value.

### HasSVariableexpenseDescriptionX

`func (o *VariableexpenseListElement) HasSVariableexpenseDescriptionX() bool`

HasSVariableexpenseDescriptionX returns a boolean if a field has been set.

### GetEVariableexpenseTaxable

`func (o *VariableexpenseListElement) GetEVariableexpenseTaxable() FieldEVariableexpenseTaxable`

GetEVariableexpenseTaxable returns the EVariableexpenseTaxable field if non-nil, zero value otherwise.

### GetEVariableexpenseTaxableOk

`func (o *VariableexpenseListElement) GetEVariableexpenseTaxableOk() (*FieldEVariableexpenseTaxable, bool)`

GetEVariableexpenseTaxableOk returns a tuple with the EVariableexpenseTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEVariableexpenseTaxable

`func (o *VariableexpenseListElement) SetEVariableexpenseTaxable(v FieldEVariableexpenseTaxable)`

SetEVariableexpenseTaxable sets EVariableexpenseTaxable field to given value.

### HasEVariableexpenseTaxable

`func (o *VariableexpenseListElement) HasEVariableexpenseTaxable() bool`

HasEVariableexpenseTaxable returns a boolean if a field has been set.

### GetBVariableexpenseIsactive

`func (o *VariableexpenseListElement) GetBVariableexpenseIsactive() bool`

GetBVariableexpenseIsactive returns the BVariableexpenseIsactive field if non-nil, zero value otherwise.

### GetBVariableexpenseIsactiveOk

`func (o *VariableexpenseListElement) GetBVariableexpenseIsactiveOk() (*bool, bool)`

GetBVariableexpenseIsactiveOk returns a tuple with the BVariableexpenseIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBVariableexpenseIsactive

`func (o *VariableexpenseListElement) SetBVariableexpenseIsactive(v bool)`

SetBVariableexpenseIsactive sets BVariableexpenseIsactive field to given value.

### HasBVariableexpenseIsactive

`func (o *VariableexpenseListElement) HasBVariableexpenseIsactive() bool`

HasBVariableexpenseIsactive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


