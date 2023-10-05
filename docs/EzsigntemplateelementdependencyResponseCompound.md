# EzsigntemplateelementdependencyResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateelementdependencyID** | **int32** | The unique ID of the Ezsigntemplateelementdependency | 
**FkiEzsigntemplateformfieldID** | Pointer to **int32** | The unique ID of the Ezsigntemplateformfield | [optional] 
**FkiEzsigntemplatesignatureID** | Pointer to **int32** | The unique ID of the Ezsigntemplatesignature | [optional] 
**FkiEzsigntemplateformfieldIDValidation** | Pointer to **int32** | The unique ID of the Ezsigntemplateformfield | [optional] 
**FkiEzsigntemplateformfieldgroupIDValidation** | Pointer to **int32** | The unique ID of the Ezsigntemplateformfieldgroup | [optional] 
**EEzsigntemplateelementdependencyValidation** | [**FieldEEzsigntemplateelementdependencyValidation**](FieldEEzsigntemplateelementdependencyValidation.md) |  | 
**BEzsigntemplateelementdependencySelected** | Pointer to **bool** | Whether if it&#39;s selected or not when using eEzsigntemplateelementdependencyValidation &#x3D; Selected | [optional] 
**EEzsigntemplateelementdependencyOperator** | Pointer to [**FieldEEzsigntemplateelementdependencyOperator**](FieldEEzsigntemplateelementdependencyOperator.md) |  | [optional] 
**SEzsigntemplateelementdependencyValue** | Pointer to **string** | The value of the Ezsignelementdependency | [optional] 

## Methods

### NewEzsigntemplateelementdependencyResponseCompound

`func NewEzsigntemplateelementdependencyResponseCompound(pkiEzsigntemplateelementdependencyID int32, eEzsigntemplateelementdependencyValidation FieldEEzsigntemplateelementdependencyValidation, ) *EzsigntemplateelementdependencyResponseCompound`

NewEzsigntemplateelementdependencyResponseCompound instantiates a new EzsigntemplateelementdependencyResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateelementdependencyResponseCompoundWithDefaults

`func NewEzsigntemplateelementdependencyResponseCompoundWithDefaults() *EzsigntemplateelementdependencyResponseCompound`

NewEzsigntemplateelementdependencyResponseCompoundWithDefaults instantiates a new EzsigntemplateelementdependencyResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateelementdependencyID

`func (o *EzsigntemplateelementdependencyResponseCompound) GetPkiEzsigntemplateelementdependencyID() int32`

GetPkiEzsigntemplateelementdependencyID returns the PkiEzsigntemplateelementdependencyID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateelementdependencyIDOk

`func (o *EzsigntemplateelementdependencyResponseCompound) GetPkiEzsigntemplateelementdependencyIDOk() (*int32, bool)`

GetPkiEzsigntemplateelementdependencyIDOk returns a tuple with the PkiEzsigntemplateelementdependencyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateelementdependencyID

`func (o *EzsigntemplateelementdependencyResponseCompound) SetPkiEzsigntemplateelementdependencyID(v int32)`

SetPkiEzsigntemplateelementdependencyID sets PkiEzsigntemplateelementdependencyID field to given value.


### GetFkiEzsigntemplateformfieldID

`func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplateformfieldID() int32`

GetFkiEzsigntemplateformfieldID returns the FkiEzsigntemplateformfieldID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateformfieldIDOk

`func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplateformfieldIDOk() (*int32, bool)`

GetFkiEzsigntemplateformfieldIDOk returns a tuple with the FkiEzsigntemplateformfieldID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateformfieldID

`func (o *EzsigntemplateelementdependencyResponseCompound) SetFkiEzsigntemplateformfieldID(v int32)`

SetFkiEzsigntemplateformfieldID sets FkiEzsigntemplateformfieldID field to given value.

### HasFkiEzsigntemplateformfieldID

`func (o *EzsigntemplateelementdependencyResponseCompound) HasFkiEzsigntemplateformfieldID() bool`

HasFkiEzsigntemplateformfieldID returns a boolean if a field has been set.

### GetFkiEzsigntemplatesignatureID

`func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplatesignatureID() int32`

GetFkiEzsigntemplatesignatureID returns the FkiEzsigntemplatesignatureID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatesignatureIDOk

`func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplatesignatureIDOk() (*int32, bool)`

GetFkiEzsigntemplatesignatureIDOk returns a tuple with the FkiEzsigntemplatesignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatesignatureID

`func (o *EzsigntemplateelementdependencyResponseCompound) SetFkiEzsigntemplatesignatureID(v int32)`

SetFkiEzsigntemplatesignatureID sets FkiEzsigntemplatesignatureID field to given value.

### HasFkiEzsigntemplatesignatureID

`func (o *EzsigntemplateelementdependencyResponseCompound) HasFkiEzsigntemplatesignatureID() bool`

HasFkiEzsigntemplatesignatureID returns a boolean if a field has been set.

### GetFkiEzsigntemplateformfieldIDValidation

`func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplateformfieldIDValidation() int32`

GetFkiEzsigntemplateformfieldIDValidation returns the FkiEzsigntemplateformfieldIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateformfieldIDValidationOk

`func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplateformfieldIDValidationOk() (*int32, bool)`

GetFkiEzsigntemplateformfieldIDValidationOk returns a tuple with the FkiEzsigntemplateformfieldIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateformfieldIDValidation

`func (o *EzsigntemplateelementdependencyResponseCompound) SetFkiEzsigntemplateformfieldIDValidation(v int32)`

SetFkiEzsigntemplateformfieldIDValidation sets FkiEzsigntemplateformfieldIDValidation field to given value.

### HasFkiEzsigntemplateformfieldIDValidation

`func (o *EzsigntemplateelementdependencyResponseCompound) HasFkiEzsigntemplateformfieldIDValidation() bool`

HasFkiEzsigntemplateformfieldIDValidation returns a boolean if a field has been set.

### GetFkiEzsigntemplateformfieldgroupIDValidation

`func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplateformfieldgroupIDValidation() int32`

GetFkiEzsigntemplateformfieldgroupIDValidation returns the FkiEzsigntemplateformfieldgroupIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateformfieldgroupIDValidationOk

`func (o *EzsigntemplateelementdependencyResponseCompound) GetFkiEzsigntemplateformfieldgroupIDValidationOk() (*int32, bool)`

GetFkiEzsigntemplateformfieldgroupIDValidationOk returns a tuple with the FkiEzsigntemplateformfieldgroupIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateformfieldgroupIDValidation

`func (o *EzsigntemplateelementdependencyResponseCompound) SetFkiEzsigntemplateformfieldgroupIDValidation(v int32)`

SetFkiEzsigntemplateformfieldgroupIDValidation sets FkiEzsigntemplateformfieldgroupIDValidation field to given value.

### HasFkiEzsigntemplateformfieldgroupIDValidation

`func (o *EzsigntemplateelementdependencyResponseCompound) HasFkiEzsigntemplateformfieldgroupIDValidation() bool`

HasFkiEzsigntemplateformfieldgroupIDValidation returns a boolean if a field has been set.

### GetEEzsigntemplateelementdependencyValidation

`func (o *EzsigntemplateelementdependencyResponseCompound) GetEEzsigntemplateelementdependencyValidation() FieldEEzsigntemplateelementdependencyValidation`

GetEEzsigntemplateelementdependencyValidation returns the EEzsigntemplateelementdependencyValidation field if non-nil, zero value otherwise.

### GetEEzsigntemplateelementdependencyValidationOk

`func (o *EzsigntemplateelementdependencyResponseCompound) GetEEzsigntemplateelementdependencyValidationOk() (*FieldEEzsigntemplateelementdependencyValidation, bool)`

GetEEzsigntemplateelementdependencyValidationOk returns a tuple with the EEzsigntemplateelementdependencyValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateelementdependencyValidation

`func (o *EzsigntemplateelementdependencyResponseCompound) SetEEzsigntemplateelementdependencyValidation(v FieldEEzsigntemplateelementdependencyValidation)`

SetEEzsigntemplateelementdependencyValidation sets EEzsigntemplateelementdependencyValidation field to given value.


### GetBEzsigntemplateelementdependencySelected

`func (o *EzsigntemplateelementdependencyResponseCompound) GetBEzsigntemplateelementdependencySelected() bool`

GetBEzsigntemplateelementdependencySelected returns the BEzsigntemplateelementdependencySelected field if non-nil, zero value otherwise.

### GetBEzsigntemplateelementdependencySelectedOk

`func (o *EzsigntemplateelementdependencyResponseCompound) GetBEzsigntemplateelementdependencySelectedOk() (*bool, bool)`

GetBEzsigntemplateelementdependencySelectedOk returns a tuple with the BEzsigntemplateelementdependencySelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateelementdependencySelected

`func (o *EzsigntemplateelementdependencyResponseCompound) SetBEzsigntemplateelementdependencySelected(v bool)`

SetBEzsigntemplateelementdependencySelected sets BEzsigntemplateelementdependencySelected field to given value.

### HasBEzsigntemplateelementdependencySelected

`func (o *EzsigntemplateelementdependencyResponseCompound) HasBEzsigntemplateelementdependencySelected() bool`

HasBEzsigntemplateelementdependencySelected returns a boolean if a field has been set.

### GetEEzsigntemplateelementdependencyOperator

`func (o *EzsigntemplateelementdependencyResponseCompound) GetEEzsigntemplateelementdependencyOperator() FieldEEzsigntemplateelementdependencyOperator`

GetEEzsigntemplateelementdependencyOperator returns the EEzsigntemplateelementdependencyOperator field if non-nil, zero value otherwise.

### GetEEzsigntemplateelementdependencyOperatorOk

`func (o *EzsigntemplateelementdependencyResponseCompound) GetEEzsigntemplateelementdependencyOperatorOk() (*FieldEEzsigntemplateelementdependencyOperator, bool)`

GetEEzsigntemplateelementdependencyOperatorOk returns a tuple with the EEzsigntemplateelementdependencyOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateelementdependencyOperator

`func (o *EzsigntemplateelementdependencyResponseCompound) SetEEzsigntemplateelementdependencyOperator(v FieldEEzsigntemplateelementdependencyOperator)`

SetEEzsigntemplateelementdependencyOperator sets EEzsigntemplateelementdependencyOperator field to given value.

### HasEEzsigntemplateelementdependencyOperator

`func (o *EzsigntemplateelementdependencyResponseCompound) HasEEzsigntemplateelementdependencyOperator() bool`

HasEEzsigntemplateelementdependencyOperator returns a boolean if a field has been set.

### GetSEzsigntemplateelementdependencyValue

`func (o *EzsigntemplateelementdependencyResponseCompound) GetSEzsigntemplateelementdependencyValue() string`

GetSEzsigntemplateelementdependencyValue returns the SEzsigntemplateelementdependencyValue field if non-nil, zero value otherwise.

### GetSEzsigntemplateelementdependencyValueOk

`func (o *EzsigntemplateelementdependencyResponseCompound) GetSEzsigntemplateelementdependencyValueOk() (*string, bool)`

GetSEzsigntemplateelementdependencyValueOk returns a tuple with the SEzsigntemplateelementdependencyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateelementdependencyValue

`func (o *EzsigntemplateelementdependencyResponseCompound) SetSEzsigntemplateelementdependencyValue(v string)`

SetSEzsigntemplateelementdependencyValue sets SEzsigntemplateelementdependencyValue field to given value.

### HasSEzsigntemplateelementdependencyValue

`func (o *EzsigntemplateelementdependencyResponseCompound) HasSEzsigntemplateelementdependencyValue() bool`

HasSEzsigntemplateelementdependencyValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


