# VariableexpenseRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiVariableexpenseID** | Pointer to **int32** | The unique ID of the Variableexpense | [optional] 
**SVariableexpenseCode** | **string** | The code of the Variableexpense | 
**ObjVariableexpenseDescription** | [**MultilingualVariableexpenseDescription**](MultilingualVariableexpenseDescription.md) |  | 
**EVariableexpenseTaxable** | [**FieldEVariableexpenseTaxable**](FieldEVariableexpenseTaxable.md) |  | 
**BVariableexpenseIsactive** | **bool** | Whether the variableexpense is active or not | 

## Methods

### NewVariableexpenseRequestCompound

`func NewVariableexpenseRequestCompound(sVariableexpenseCode string, objVariableexpenseDescription MultilingualVariableexpenseDescription, eVariableexpenseTaxable FieldEVariableexpenseTaxable, bVariableexpenseIsactive bool, ) *VariableexpenseRequestCompound`

NewVariableexpenseRequestCompound instantiates a new VariableexpenseRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableexpenseRequestCompoundWithDefaults

`func NewVariableexpenseRequestCompoundWithDefaults() *VariableexpenseRequestCompound`

NewVariableexpenseRequestCompoundWithDefaults instantiates a new VariableexpenseRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiVariableexpenseID

`func (o *VariableexpenseRequestCompound) GetPkiVariableexpenseID() int32`

GetPkiVariableexpenseID returns the PkiVariableexpenseID field if non-nil, zero value otherwise.

### GetPkiVariableexpenseIDOk

`func (o *VariableexpenseRequestCompound) GetPkiVariableexpenseIDOk() (*int32, bool)`

GetPkiVariableexpenseIDOk returns a tuple with the PkiVariableexpenseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiVariableexpenseID

`func (o *VariableexpenseRequestCompound) SetPkiVariableexpenseID(v int32)`

SetPkiVariableexpenseID sets PkiVariableexpenseID field to given value.

### HasPkiVariableexpenseID

`func (o *VariableexpenseRequestCompound) HasPkiVariableexpenseID() bool`

HasPkiVariableexpenseID returns a boolean if a field has been set.

### GetSVariableexpenseCode

`func (o *VariableexpenseRequestCompound) GetSVariableexpenseCode() string`

GetSVariableexpenseCode returns the SVariableexpenseCode field if non-nil, zero value otherwise.

### GetSVariableexpenseCodeOk

`func (o *VariableexpenseRequestCompound) GetSVariableexpenseCodeOk() (*string, bool)`

GetSVariableexpenseCodeOk returns a tuple with the SVariableexpenseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSVariableexpenseCode

`func (o *VariableexpenseRequestCompound) SetSVariableexpenseCode(v string)`

SetSVariableexpenseCode sets SVariableexpenseCode field to given value.


### GetObjVariableexpenseDescription

`func (o *VariableexpenseRequestCompound) GetObjVariableexpenseDescription() MultilingualVariableexpenseDescription`

GetObjVariableexpenseDescription returns the ObjVariableexpenseDescription field if non-nil, zero value otherwise.

### GetObjVariableexpenseDescriptionOk

`func (o *VariableexpenseRequestCompound) GetObjVariableexpenseDescriptionOk() (*MultilingualVariableexpenseDescription, bool)`

GetObjVariableexpenseDescriptionOk returns a tuple with the ObjVariableexpenseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjVariableexpenseDescription

`func (o *VariableexpenseRequestCompound) SetObjVariableexpenseDescription(v MultilingualVariableexpenseDescription)`

SetObjVariableexpenseDescription sets ObjVariableexpenseDescription field to given value.


### GetEVariableexpenseTaxable

`func (o *VariableexpenseRequestCompound) GetEVariableexpenseTaxable() FieldEVariableexpenseTaxable`

GetEVariableexpenseTaxable returns the EVariableexpenseTaxable field if non-nil, zero value otherwise.

### GetEVariableexpenseTaxableOk

`func (o *VariableexpenseRequestCompound) GetEVariableexpenseTaxableOk() (*FieldEVariableexpenseTaxable, bool)`

GetEVariableexpenseTaxableOk returns a tuple with the EVariableexpenseTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEVariableexpenseTaxable

`func (o *VariableexpenseRequestCompound) SetEVariableexpenseTaxable(v FieldEVariableexpenseTaxable)`

SetEVariableexpenseTaxable sets EVariableexpenseTaxable field to given value.


### GetBVariableexpenseIsactive

`func (o *VariableexpenseRequestCompound) GetBVariableexpenseIsactive() bool`

GetBVariableexpenseIsactive returns the BVariableexpenseIsactive field if non-nil, zero value otherwise.

### GetBVariableexpenseIsactiveOk

`func (o *VariableexpenseRequestCompound) GetBVariableexpenseIsactiveOk() (*bool, bool)`

GetBVariableexpenseIsactiveOk returns a tuple with the BVariableexpenseIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBVariableexpenseIsactive

`func (o *VariableexpenseRequestCompound) SetBVariableexpenseIsactive(v bool)`

SetBVariableexpenseIsactive sets BVariableexpenseIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


