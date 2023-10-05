# EzsignelementdependencyResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignelementdependencyID** | **int32** | The unique ID of the Ezsignelementdependency | 
**FkiEzsignformfieldID** | Pointer to **int32** | The unique ID of the Ezsignformfield | [optional] 
**FkiEzsignsignatureID** | Pointer to **int32** | The unique ID of the Ezsignsignature | [optional] 
**FkiEzsignformfieldIDValidation** | Pointer to **int32** | The unique ID of the Ezsignformfield | [optional] 
**FkiEzsignformfieldgroupIDValidation** | Pointer to **int32** | The unique ID of the Ezsignformfieldgroup | [optional] 
**EEzsignelementdependencyValidation** | [**FieldEEzsignelementdependencyValidation**](FieldEEzsignelementdependencyValidation.md) |  | 
**BEzsignelementdependencySelected** | Pointer to **bool** | Whether if it&#39;s selected or not when using eEzsignelementdependencyValidation &#x3D; Selected | [optional] 
**EEzsignelementdependencyOperator** | Pointer to [**FieldEEzsignelementdependencyOperator**](FieldEEzsignelementdependencyOperator.md) |  | [optional] 
**SEzsignelementdependencyValue** | Pointer to **string** | The value of the Ezsignelementdependency | [optional] 

## Methods

### NewEzsignelementdependencyResponseCompound

`func NewEzsignelementdependencyResponseCompound(pkiEzsignelementdependencyID int32, eEzsignelementdependencyValidation FieldEEzsignelementdependencyValidation, ) *EzsignelementdependencyResponseCompound`

NewEzsignelementdependencyResponseCompound instantiates a new EzsignelementdependencyResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignelementdependencyResponseCompoundWithDefaults

`func NewEzsignelementdependencyResponseCompoundWithDefaults() *EzsignelementdependencyResponseCompound`

NewEzsignelementdependencyResponseCompoundWithDefaults instantiates a new EzsignelementdependencyResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignelementdependencyID

`func (o *EzsignelementdependencyResponseCompound) GetPkiEzsignelementdependencyID() int32`

GetPkiEzsignelementdependencyID returns the PkiEzsignelementdependencyID field if non-nil, zero value otherwise.

### GetPkiEzsignelementdependencyIDOk

`func (o *EzsignelementdependencyResponseCompound) GetPkiEzsignelementdependencyIDOk() (*int32, bool)`

GetPkiEzsignelementdependencyIDOk returns a tuple with the PkiEzsignelementdependencyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignelementdependencyID

`func (o *EzsignelementdependencyResponseCompound) SetPkiEzsignelementdependencyID(v int32)`

SetPkiEzsignelementdependencyID sets PkiEzsignelementdependencyID field to given value.


### GetFkiEzsignformfieldID

`func (o *EzsignelementdependencyResponseCompound) GetFkiEzsignformfieldID() int32`

GetFkiEzsignformfieldID returns the FkiEzsignformfieldID field if non-nil, zero value otherwise.

### GetFkiEzsignformfieldIDOk

`func (o *EzsignelementdependencyResponseCompound) GetFkiEzsignformfieldIDOk() (*int32, bool)`

GetFkiEzsignformfieldIDOk returns a tuple with the FkiEzsignformfieldID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignformfieldID

`func (o *EzsignelementdependencyResponseCompound) SetFkiEzsignformfieldID(v int32)`

SetFkiEzsignformfieldID sets FkiEzsignformfieldID field to given value.

### HasFkiEzsignformfieldID

`func (o *EzsignelementdependencyResponseCompound) HasFkiEzsignformfieldID() bool`

HasFkiEzsignformfieldID returns a boolean if a field has been set.

### GetFkiEzsignsignatureID

`func (o *EzsignelementdependencyResponseCompound) GetFkiEzsignsignatureID() int32`

GetFkiEzsignsignatureID returns the FkiEzsignsignatureID field if non-nil, zero value otherwise.

### GetFkiEzsignsignatureIDOk

`func (o *EzsignelementdependencyResponseCompound) GetFkiEzsignsignatureIDOk() (*int32, bool)`

GetFkiEzsignsignatureIDOk returns a tuple with the FkiEzsignsignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsignatureID

`func (o *EzsignelementdependencyResponseCompound) SetFkiEzsignsignatureID(v int32)`

SetFkiEzsignsignatureID sets FkiEzsignsignatureID field to given value.

### HasFkiEzsignsignatureID

`func (o *EzsignelementdependencyResponseCompound) HasFkiEzsignsignatureID() bool`

HasFkiEzsignsignatureID returns a boolean if a field has been set.

### GetFkiEzsignformfieldIDValidation

`func (o *EzsignelementdependencyResponseCompound) GetFkiEzsignformfieldIDValidation() int32`

GetFkiEzsignformfieldIDValidation returns the FkiEzsignformfieldIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsignformfieldIDValidationOk

`func (o *EzsignelementdependencyResponseCompound) GetFkiEzsignformfieldIDValidationOk() (*int32, bool)`

GetFkiEzsignformfieldIDValidationOk returns a tuple with the FkiEzsignformfieldIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignformfieldIDValidation

`func (o *EzsignelementdependencyResponseCompound) SetFkiEzsignformfieldIDValidation(v int32)`

SetFkiEzsignformfieldIDValidation sets FkiEzsignformfieldIDValidation field to given value.

### HasFkiEzsignformfieldIDValidation

`func (o *EzsignelementdependencyResponseCompound) HasFkiEzsignformfieldIDValidation() bool`

HasFkiEzsignformfieldIDValidation returns a boolean if a field has been set.

### GetFkiEzsignformfieldgroupIDValidation

`func (o *EzsignelementdependencyResponseCompound) GetFkiEzsignformfieldgroupIDValidation() int32`

GetFkiEzsignformfieldgroupIDValidation returns the FkiEzsignformfieldgroupIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsignformfieldgroupIDValidationOk

`func (o *EzsignelementdependencyResponseCompound) GetFkiEzsignformfieldgroupIDValidationOk() (*int32, bool)`

GetFkiEzsignformfieldgroupIDValidationOk returns a tuple with the FkiEzsignformfieldgroupIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignformfieldgroupIDValidation

`func (o *EzsignelementdependencyResponseCompound) SetFkiEzsignformfieldgroupIDValidation(v int32)`

SetFkiEzsignformfieldgroupIDValidation sets FkiEzsignformfieldgroupIDValidation field to given value.

### HasFkiEzsignformfieldgroupIDValidation

`func (o *EzsignelementdependencyResponseCompound) HasFkiEzsignformfieldgroupIDValidation() bool`

HasFkiEzsignformfieldgroupIDValidation returns a boolean if a field has been set.

### GetEEzsignelementdependencyValidation

`func (o *EzsignelementdependencyResponseCompound) GetEEzsignelementdependencyValidation() FieldEEzsignelementdependencyValidation`

GetEEzsignelementdependencyValidation returns the EEzsignelementdependencyValidation field if non-nil, zero value otherwise.

### GetEEzsignelementdependencyValidationOk

`func (o *EzsignelementdependencyResponseCompound) GetEEzsignelementdependencyValidationOk() (*FieldEEzsignelementdependencyValidation, bool)`

GetEEzsignelementdependencyValidationOk returns a tuple with the EEzsignelementdependencyValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignelementdependencyValidation

`func (o *EzsignelementdependencyResponseCompound) SetEEzsignelementdependencyValidation(v FieldEEzsignelementdependencyValidation)`

SetEEzsignelementdependencyValidation sets EEzsignelementdependencyValidation field to given value.


### GetBEzsignelementdependencySelected

`func (o *EzsignelementdependencyResponseCompound) GetBEzsignelementdependencySelected() bool`

GetBEzsignelementdependencySelected returns the BEzsignelementdependencySelected field if non-nil, zero value otherwise.

### GetBEzsignelementdependencySelectedOk

`func (o *EzsignelementdependencyResponseCompound) GetBEzsignelementdependencySelectedOk() (*bool, bool)`

GetBEzsignelementdependencySelectedOk returns a tuple with the BEzsignelementdependencySelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignelementdependencySelected

`func (o *EzsignelementdependencyResponseCompound) SetBEzsignelementdependencySelected(v bool)`

SetBEzsignelementdependencySelected sets BEzsignelementdependencySelected field to given value.

### HasBEzsignelementdependencySelected

`func (o *EzsignelementdependencyResponseCompound) HasBEzsignelementdependencySelected() bool`

HasBEzsignelementdependencySelected returns a boolean if a field has been set.

### GetEEzsignelementdependencyOperator

`func (o *EzsignelementdependencyResponseCompound) GetEEzsignelementdependencyOperator() FieldEEzsignelementdependencyOperator`

GetEEzsignelementdependencyOperator returns the EEzsignelementdependencyOperator field if non-nil, zero value otherwise.

### GetEEzsignelementdependencyOperatorOk

`func (o *EzsignelementdependencyResponseCompound) GetEEzsignelementdependencyOperatorOk() (*FieldEEzsignelementdependencyOperator, bool)`

GetEEzsignelementdependencyOperatorOk returns a tuple with the EEzsignelementdependencyOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignelementdependencyOperator

`func (o *EzsignelementdependencyResponseCompound) SetEEzsignelementdependencyOperator(v FieldEEzsignelementdependencyOperator)`

SetEEzsignelementdependencyOperator sets EEzsignelementdependencyOperator field to given value.

### HasEEzsignelementdependencyOperator

`func (o *EzsignelementdependencyResponseCompound) HasEEzsignelementdependencyOperator() bool`

HasEEzsignelementdependencyOperator returns a boolean if a field has been set.

### GetSEzsignelementdependencyValue

`func (o *EzsignelementdependencyResponseCompound) GetSEzsignelementdependencyValue() string`

GetSEzsignelementdependencyValue returns the SEzsignelementdependencyValue field if non-nil, zero value otherwise.

### GetSEzsignelementdependencyValueOk

`func (o *EzsignelementdependencyResponseCompound) GetSEzsignelementdependencyValueOk() (*string, bool)`

GetSEzsignelementdependencyValueOk returns a tuple with the SEzsignelementdependencyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignelementdependencyValue

`func (o *EzsignelementdependencyResponseCompound) SetSEzsignelementdependencyValue(v string)`

SetSEzsignelementdependencyValue sets SEzsignelementdependencyValue field to given value.

### HasSEzsignelementdependencyValue

`func (o *EzsignelementdependencyResponseCompound) HasSEzsignelementdependencyValue() bool`

HasSEzsignelementdependencyValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


