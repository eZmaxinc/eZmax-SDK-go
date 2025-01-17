# VariableexpenseResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiVariableexpenseID** | **int32** | The unique ID of the Variableexpense | 
**SVariableexpenseCode** | Pointer to **string** | The code of the Variableexpense | [optional] 
**ObjVariableexpenseDescription** | [**MultilingualVariableexpenseDescription**](MultilingualVariableexpenseDescription.md) |  | 
**EVariableexpenseTaxable** | Pointer to [**FieldEVariableexpenseTaxable**](FieldEVariableexpenseTaxable.md) |  | [optional] 
**BVariableexpenseIsactive** | Pointer to **bool** | Whether the variableexpense is active or not | [optional] 

## Methods

### NewVariableexpenseResponseCompound

`func NewVariableexpenseResponseCompound(pkiVariableexpenseID int32, objVariableexpenseDescription MultilingualVariableexpenseDescription, ) *VariableexpenseResponseCompound`

NewVariableexpenseResponseCompound instantiates a new VariableexpenseResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableexpenseResponseCompoundWithDefaults

`func NewVariableexpenseResponseCompoundWithDefaults() *VariableexpenseResponseCompound`

NewVariableexpenseResponseCompoundWithDefaults instantiates a new VariableexpenseResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiVariableexpenseID

`func (o *VariableexpenseResponseCompound) GetPkiVariableexpenseID() int32`

GetPkiVariableexpenseID returns the PkiVariableexpenseID field if non-nil, zero value otherwise.

### GetPkiVariableexpenseIDOk

`func (o *VariableexpenseResponseCompound) GetPkiVariableexpenseIDOk() (*int32, bool)`

GetPkiVariableexpenseIDOk returns a tuple with the PkiVariableexpenseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiVariableexpenseID

`func (o *VariableexpenseResponseCompound) SetPkiVariableexpenseID(v int32)`

SetPkiVariableexpenseID sets PkiVariableexpenseID field to given value.


### GetSVariableexpenseCode

`func (o *VariableexpenseResponseCompound) GetSVariableexpenseCode() string`

GetSVariableexpenseCode returns the SVariableexpenseCode field if non-nil, zero value otherwise.

### GetSVariableexpenseCodeOk

`func (o *VariableexpenseResponseCompound) GetSVariableexpenseCodeOk() (*string, bool)`

GetSVariableexpenseCodeOk returns a tuple with the SVariableexpenseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSVariableexpenseCode

`func (o *VariableexpenseResponseCompound) SetSVariableexpenseCode(v string)`

SetSVariableexpenseCode sets SVariableexpenseCode field to given value.

### HasSVariableexpenseCode

`func (o *VariableexpenseResponseCompound) HasSVariableexpenseCode() bool`

HasSVariableexpenseCode returns a boolean if a field has been set.

### GetObjVariableexpenseDescription

`func (o *VariableexpenseResponseCompound) GetObjVariableexpenseDescription() MultilingualVariableexpenseDescription`

GetObjVariableexpenseDescription returns the ObjVariableexpenseDescription field if non-nil, zero value otherwise.

### GetObjVariableexpenseDescriptionOk

`func (o *VariableexpenseResponseCompound) GetObjVariableexpenseDescriptionOk() (*MultilingualVariableexpenseDescription, bool)`

GetObjVariableexpenseDescriptionOk returns a tuple with the ObjVariableexpenseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjVariableexpenseDescription

`func (o *VariableexpenseResponseCompound) SetObjVariableexpenseDescription(v MultilingualVariableexpenseDescription)`

SetObjVariableexpenseDescription sets ObjVariableexpenseDescription field to given value.


### GetEVariableexpenseTaxable

`func (o *VariableexpenseResponseCompound) GetEVariableexpenseTaxable() FieldEVariableexpenseTaxable`

GetEVariableexpenseTaxable returns the EVariableexpenseTaxable field if non-nil, zero value otherwise.

### GetEVariableexpenseTaxableOk

`func (o *VariableexpenseResponseCompound) GetEVariableexpenseTaxableOk() (*FieldEVariableexpenseTaxable, bool)`

GetEVariableexpenseTaxableOk returns a tuple with the EVariableexpenseTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEVariableexpenseTaxable

`func (o *VariableexpenseResponseCompound) SetEVariableexpenseTaxable(v FieldEVariableexpenseTaxable)`

SetEVariableexpenseTaxable sets EVariableexpenseTaxable field to given value.

### HasEVariableexpenseTaxable

`func (o *VariableexpenseResponseCompound) HasEVariableexpenseTaxable() bool`

HasEVariableexpenseTaxable returns a boolean if a field has been set.

### GetBVariableexpenseIsactive

`func (o *VariableexpenseResponseCompound) GetBVariableexpenseIsactive() bool`

GetBVariableexpenseIsactive returns the BVariableexpenseIsactive field if non-nil, zero value otherwise.

### GetBVariableexpenseIsactiveOk

`func (o *VariableexpenseResponseCompound) GetBVariableexpenseIsactiveOk() (*bool, bool)`

GetBVariableexpenseIsactiveOk returns a tuple with the BVariableexpenseIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBVariableexpenseIsactive

`func (o *VariableexpenseResponseCompound) SetBVariableexpenseIsactive(v bool)`

SetBVariableexpenseIsactive sets BVariableexpenseIsactive field to given value.

### HasBVariableexpenseIsactive

`func (o *VariableexpenseResponseCompound) HasBVariableexpenseIsactive() bool`

HasBVariableexpenseIsactive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


