# EzsignelementdependencyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignelementdependencyID** | Pointer to **int32** | The unique ID of the Ezsignelementdependency | [optional] 
**FkiEzsignformfieldIDValidation** | Pointer to **int32** | The unique ID of the Ezsignformfield | [optional] 
**FkiEzsignformfieldgroupIDValidation** | Pointer to **int32** | The unique ID of the Ezsignformfieldgroup | [optional] 
**SEzsignelementdependencyEzsignformfieldgrouplabel** | Pointer to **string** | The Label for the Ezsignformfieldgroup | [optional] 
**SEzsignelementdependencyEzsignformfieldlabel** | Pointer to **string** | The Label for the Ezsignformfield | [optional] 
**EEzsignelementdependencyValidation** | [**FieldEEzsignelementdependencyValidation**](FieldEEzsignelementdependencyValidation.md) |  | 
**BEzsignelementdependencySelected** | Pointer to **bool** | Whether if it&#39;s selected or not when using eEzsignelementdependencyValidation &#x3D; Selected | [optional] 
**EEzsignelementdependencyOperator** | Pointer to [**FieldEEzsignelementdependencyOperator**](FieldEEzsignelementdependencyOperator.md) |  | [optional] 
**SEzsignelementdependencyValue** | Pointer to **string** | The value of the Ezsignelementdependency | [optional] 

## Methods

### NewEzsignelementdependencyRequest

`func NewEzsignelementdependencyRequest(eEzsignelementdependencyValidation FieldEEzsignelementdependencyValidation, ) *EzsignelementdependencyRequest`

NewEzsignelementdependencyRequest instantiates a new EzsignelementdependencyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignelementdependencyRequestWithDefaults

`func NewEzsignelementdependencyRequestWithDefaults() *EzsignelementdependencyRequest`

NewEzsignelementdependencyRequestWithDefaults instantiates a new EzsignelementdependencyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignelementdependencyID

`func (o *EzsignelementdependencyRequest) GetPkiEzsignelementdependencyID() int32`

GetPkiEzsignelementdependencyID returns the PkiEzsignelementdependencyID field if non-nil, zero value otherwise.

### GetPkiEzsignelementdependencyIDOk

`func (o *EzsignelementdependencyRequest) GetPkiEzsignelementdependencyIDOk() (*int32, bool)`

GetPkiEzsignelementdependencyIDOk returns a tuple with the PkiEzsignelementdependencyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignelementdependencyID

`func (o *EzsignelementdependencyRequest) SetPkiEzsignelementdependencyID(v int32)`

SetPkiEzsignelementdependencyID sets PkiEzsignelementdependencyID field to given value.

### HasPkiEzsignelementdependencyID

`func (o *EzsignelementdependencyRequest) HasPkiEzsignelementdependencyID() bool`

HasPkiEzsignelementdependencyID returns a boolean if a field has been set.

### GetFkiEzsignformfieldIDValidation

`func (o *EzsignelementdependencyRequest) GetFkiEzsignformfieldIDValidation() int32`

GetFkiEzsignformfieldIDValidation returns the FkiEzsignformfieldIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsignformfieldIDValidationOk

`func (o *EzsignelementdependencyRequest) GetFkiEzsignformfieldIDValidationOk() (*int32, bool)`

GetFkiEzsignformfieldIDValidationOk returns a tuple with the FkiEzsignformfieldIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignformfieldIDValidation

`func (o *EzsignelementdependencyRequest) SetFkiEzsignformfieldIDValidation(v int32)`

SetFkiEzsignformfieldIDValidation sets FkiEzsignformfieldIDValidation field to given value.

### HasFkiEzsignformfieldIDValidation

`func (o *EzsignelementdependencyRequest) HasFkiEzsignformfieldIDValidation() bool`

HasFkiEzsignformfieldIDValidation returns a boolean if a field has been set.

### GetFkiEzsignformfieldgroupIDValidation

`func (o *EzsignelementdependencyRequest) GetFkiEzsignformfieldgroupIDValidation() int32`

GetFkiEzsignformfieldgroupIDValidation returns the FkiEzsignformfieldgroupIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsignformfieldgroupIDValidationOk

`func (o *EzsignelementdependencyRequest) GetFkiEzsignformfieldgroupIDValidationOk() (*int32, bool)`

GetFkiEzsignformfieldgroupIDValidationOk returns a tuple with the FkiEzsignformfieldgroupIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignformfieldgroupIDValidation

`func (o *EzsignelementdependencyRequest) SetFkiEzsignformfieldgroupIDValidation(v int32)`

SetFkiEzsignformfieldgroupIDValidation sets FkiEzsignformfieldgroupIDValidation field to given value.

### HasFkiEzsignformfieldgroupIDValidation

`func (o *EzsignelementdependencyRequest) HasFkiEzsignformfieldgroupIDValidation() bool`

HasFkiEzsignformfieldgroupIDValidation returns a boolean if a field has been set.

### GetSEzsignelementdependencyEzsignformfieldgrouplabel

`func (o *EzsignelementdependencyRequest) GetSEzsignelementdependencyEzsignformfieldgrouplabel() string`

GetSEzsignelementdependencyEzsignformfieldgrouplabel returns the SEzsignelementdependencyEzsignformfieldgrouplabel field if non-nil, zero value otherwise.

### GetSEzsignelementdependencyEzsignformfieldgrouplabelOk

`func (o *EzsignelementdependencyRequest) GetSEzsignelementdependencyEzsignformfieldgrouplabelOk() (*string, bool)`

GetSEzsignelementdependencyEzsignformfieldgrouplabelOk returns a tuple with the SEzsignelementdependencyEzsignformfieldgrouplabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignelementdependencyEzsignformfieldgrouplabel

`func (o *EzsignelementdependencyRequest) SetSEzsignelementdependencyEzsignformfieldgrouplabel(v string)`

SetSEzsignelementdependencyEzsignformfieldgrouplabel sets SEzsignelementdependencyEzsignformfieldgrouplabel field to given value.

### HasSEzsignelementdependencyEzsignformfieldgrouplabel

`func (o *EzsignelementdependencyRequest) HasSEzsignelementdependencyEzsignformfieldgrouplabel() bool`

HasSEzsignelementdependencyEzsignformfieldgrouplabel returns a boolean if a field has been set.

### GetSEzsignelementdependencyEzsignformfieldlabel

`func (o *EzsignelementdependencyRequest) GetSEzsignelementdependencyEzsignformfieldlabel() string`

GetSEzsignelementdependencyEzsignformfieldlabel returns the SEzsignelementdependencyEzsignformfieldlabel field if non-nil, zero value otherwise.

### GetSEzsignelementdependencyEzsignformfieldlabelOk

`func (o *EzsignelementdependencyRequest) GetSEzsignelementdependencyEzsignformfieldlabelOk() (*string, bool)`

GetSEzsignelementdependencyEzsignformfieldlabelOk returns a tuple with the SEzsignelementdependencyEzsignformfieldlabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignelementdependencyEzsignformfieldlabel

`func (o *EzsignelementdependencyRequest) SetSEzsignelementdependencyEzsignformfieldlabel(v string)`

SetSEzsignelementdependencyEzsignformfieldlabel sets SEzsignelementdependencyEzsignformfieldlabel field to given value.

### HasSEzsignelementdependencyEzsignformfieldlabel

`func (o *EzsignelementdependencyRequest) HasSEzsignelementdependencyEzsignformfieldlabel() bool`

HasSEzsignelementdependencyEzsignformfieldlabel returns a boolean if a field has been set.

### GetEEzsignelementdependencyValidation

`func (o *EzsignelementdependencyRequest) GetEEzsignelementdependencyValidation() FieldEEzsignelementdependencyValidation`

GetEEzsignelementdependencyValidation returns the EEzsignelementdependencyValidation field if non-nil, zero value otherwise.

### GetEEzsignelementdependencyValidationOk

`func (o *EzsignelementdependencyRequest) GetEEzsignelementdependencyValidationOk() (*FieldEEzsignelementdependencyValidation, bool)`

GetEEzsignelementdependencyValidationOk returns a tuple with the EEzsignelementdependencyValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignelementdependencyValidation

`func (o *EzsignelementdependencyRequest) SetEEzsignelementdependencyValidation(v FieldEEzsignelementdependencyValidation)`

SetEEzsignelementdependencyValidation sets EEzsignelementdependencyValidation field to given value.


### GetBEzsignelementdependencySelected

`func (o *EzsignelementdependencyRequest) GetBEzsignelementdependencySelected() bool`

GetBEzsignelementdependencySelected returns the BEzsignelementdependencySelected field if non-nil, zero value otherwise.

### GetBEzsignelementdependencySelectedOk

`func (o *EzsignelementdependencyRequest) GetBEzsignelementdependencySelectedOk() (*bool, bool)`

GetBEzsignelementdependencySelectedOk returns a tuple with the BEzsignelementdependencySelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignelementdependencySelected

`func (o *EzsignelementdependencyRequest) SetBEzsignelementdependencySelected(v bool)`

SetBEzsignelementdependencySelected sets BEzsignelementdependencySelected field to given value.

### HasBEzsignelementdependencySelected

`func (o *EzsignelementdependencyRequest) HasBEzsignelementdependencySelected() bool`

HasBEzsignelementdependencySelected returns a boolean if a field has been set.

### GetEEzsignelementdependencyOperator

`func (o *EzsignelementdependencyRequest) GetEEzsignelementdependencyOperator() FieldEEzsignelementdependencyOperator`

GetEEzsignelementdependencyOperator returns the EEzsignelementdependencyOperator field if non-nil, zero value otherwise.

### GetEEzsignelementdependencyOperatorOk

`func (o *EzsignelementdependencyRequest) GetEEzsignelementdependencyOperatorOk() (*FieldEEzsignelementdependencyOperator, bool)`

GetEEzsignelementdependencyOperatorOk returns a tuple with the EEzsignelementdependencyOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignelementdependencyOperator

`func (o *EzsignelementdependencyRequest) SetEEzsignelementdependencyOperator(v FieldEEzsignelementdependencyOperator)`

SetEEzsignelementdependencyOperator sets EEzsignelementdependencyOperator field to given value.

### HasEEzsignelementdependencyOperator

`func (o *EzsignelementdependencyRequest) HasEEzsignelementdependencyOperator() bool`

HasEEzsignelementdependencyOperator returns a boolean if a field has been set.

### GetSEzsignelementdependencyValue

`func (o *EzsignelementdependencyRequest) GetSEzsignelementdependencyValue() string`

GetSEzsignelementdependencyValue returns the SEzsignelementdependencyValue field if non-nil, zero value otherwise.

### GetSEzsignelementdependencyValueOk

`func (o *EzsignelementdependencyRequest) GetSEzsignelementdependencyValueOk() (*string, bool)`

GetSEzsignelementdependencyValueOk returns a tuple with the SEzsignelementdependencyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignelementdependencyValue

`func (o *EzsignelementdependencyRequest) SetSEzsignelementdependencyValue(v string)`

SetSEzsignelementdependencyValue sets SEzsignelementdependencyValue field to given value.

### HasSEzsignelementdependencyValue

`func (o *EzsignelementdependencyRequest) HasSEzsignelementdependencyValue() bool`

HasSEzsignelementdependencyValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


