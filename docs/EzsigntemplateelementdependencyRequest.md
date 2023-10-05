# EzsigntemplateelementdependencyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateelementdependencyID** | Pointer to **int32** | The unique ID of the Ezsigntemplateelementdependency | [optional] 
**FkiEzsigntemplateformfieldIDValidation** | Pointer to **int32** | The unique ID of the Ezsigntemplateformfield | [optional] 
**FkiEzsigntemplateformfieldgroupIDValidation** | Pointer to **int32** | The unique ID of the Ezsigntemplateformfieldgroup | [optional] 
**SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel** | Pointer to **string** | The Label for the Ezsigntemplateformfieldgroup | [optional] 
**SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel** | Pointer to **string** | The Label for the Ezsigntemplateformfield | [optional] 
**EEzsigntemplateelementdependencyValidation** | [**FieldEEzsigntemplateelementdependencyValidation**](FieldEEzsigntemplateelementdependencyValidation.md) |  | 
**BEzsigntemplateelementdependencySelected** | Pointer to **bool** | Whether if it&#39;s selected or not when using eEzsigntemplateelementdependencyValidation &#x3D; Selected | [optional] 
**EEzsigntemplateelementdependencyOperator** | Pointer to [**FieldEEzsigntemplateelementdependencyOperator**](FieldEEzsigntemplateelementdependencyOperator.md) |  | [optional] 
**SEzsigntemplateelementdependencyValue** | Pointer to **string** | The value of the Ezsignelementdependency | [optional] 

## Methods

### NewEzsigntemplateelementdependencyRequest

`func NewEzsigntemplateelementdependencyRequest(eEzsigntemplateelementdependencyValidation FieldEEzsigntemplateelementdependencyValidation, ) *EzsigntemplateelementdependencyRequest`

NewEzsigntemplateelementdependencyRequest instantiates a new EzsigntemplateelementdependencyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateelementdependencyRequestWithDefaults

`func NewEzsigntemplateelementdependencyRequestWithDefaults() *EzsigntemplateelementdependencyRequest`

NewEzsigntemplateelementdependencyRequestWithDefaults instantiates a new EzsigntemplateelementdependencyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateelementdependencyID

`func (o *EzsigntemplateelementdependencyRequest) GetPkiEzsigntemplateelementdependencyID() int32`

GetPkiEzsigntemplateelementdependencyID returns the PkiEzsigntemplateelementdependencyID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateelementdependencyIDOk

`func (o *EzsigntemplateelementdependencyRequest) GetPkiEzsigntemplateelementdependencyIDOk() (*int32, bool)`

GetPkiEzsigntemplateelementdependencyIDOk returns a tuple with the PkiEzsigntemplateelementdependencyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateelementdependencyID

`func (o *EzsigntemplateelementdependencyRequest) SetPkiEzsigntemplateelementdependencyID(v int32)`

SetPkiEzsigntemplateelementdependencyID sets PkiEzsigntemplateelementdependencyID field to given value.

### HasPkiEzsigntemplateelementdependencyID

`func (o *EzsigntemplateelementdependencyRequest) HasPkiEzsigntemplateelementdependencyID() bool`

HasPkiEzsigntemplateelementdependencyID returns a boolean if a field has been set.

### GetFkiEzsigntemplateformfieldIDValidation

`func (o *EzsigntemplateelementdependencyRequest) GetFkiEzsigntemplateformfieldIDValidation() int32`

GetFkiEzsigntemplateformfieldIDValidation returns the FkiEzsigntemplateformfieldIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateformfieldIDValidationOk

`func (o *EzsigntemplateelementdependencyRequest) GetFkiEzsigntemplateformfieldIDValidationOk() (*int32, bool)`

GetFkiEzsigntemplateformfieldIDValidationOk returns a tuple with the FkiEzsigntemplateformfieldIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateformfieldIDValidation

`func (o *EzsigntemplateelementdependencyRequest) SetFkiEzsigntemplateformfieldIDValidation(v int32)`

SetFkiEzsigntemplateformfieldIDValidation sets FkiEzsigntemplateformfieldIDValidation field to given value.

### HasFkiEzsigntemplateformfieldIDValidation

`func (o *EzsigntemplateelementdependencyRequest) HasFkiEzsigntemplateformfieldIDValidation() bool`

HasFkiEzsigntemplateformfieldIDValidation returns a boolean if a field has been set.

### GetFkiEzsigntemplateformfieldgroupIDValidation

`func (o *EzsigntemplateelementdependencyRequest) GetFkiEzsigntemplateformfieldgroupIDValidation() int32`

GetFkiEzsigntemplateformfieldgroupIDValidation returns the FkiEzsigntemplateformfieldgroupIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateformfieldgroupIDValidationOk

`func (o *EzsigntemplateelementdependencyRequest) GetFkiEzsigntemplateformfieldgroupIDValidationOk() (*int32, bool)`

GetFkiEzsigntemplateformfieldgroupIDValidationOk returns a tuple with the FkiEzsigntemplateformfieldgroupIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateformfieldgroupIDValidation

`func (o *EzsigntemplateelementdependencyRequest) SetFkiEzsigntemplateformfieldgroupIDValidation(v int32)`

SetFkiEzsigntemplateformfieldgroupIDValidation sets FkiEzsigntemplateformfieldgroupIDValidation field to given value.

### HasFkiEzsigntemplateformfieldgroupIDValidation

`func (o *EzsigntemplateelementdependencyRequest) HasFkiEzsigntemplateformfieldgroupIDValidation() bool`

HasFkiEzsigntemplateformfieldgroupIDValidation returns a boolean if a field has been set.

### GetSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel

`func (o *EzsigntemplateelementdependencyRequest) GetSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel() string`

GetSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel returns the SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel field if non-nil, zero value otherwise.

### GetSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabelOk

`func (o *EzsigntemplateelementdependencyRequest) GetSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabelOk() (*string, bool)`

GetSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabelOk returns a tuple with the SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel

`func (o *EzsigntemplateelementdependencyRequest) SetSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel(v string)`

SetSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel sets SEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel field to given value.

### HasSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel

`func (o *EzsigntemplateelementdependencyRequest) HasSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel() bool`

HasSEzsigntemplateelementdependencyEzsigntemplateformfieldgrouplabel returns a boolean if a field has been set.

### GetSEzsigntemplateelementdependencyEzsigntemplateformfieldlabel

`func (o *EzsigntemplateelementdependencyRequest) GetSEzsigntemplateelementdependencyEzsigntemplateformfieldlabel() string`

GetSEzsigntemplateelementdependencyEzsigntemplateformfieldlabel returns the SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel field if non-nil, zero value otherwise.

### GetSEzsigntemplateelementdependencyEzsigntemplateformfieldlabelOk

`func (o *EzsigntemplateelementdependencyRequest) GetSEzsigntemplateelementdependencyEzsigntemplateformfieldlabelOk() (*string, bool)`

GetSEzsigntemplateelementdependencyEzsigntemplateformfieldlabelOk returns a tuple with the SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateelementdependencyEzsigntemplateformfieldlabel

`func (o *EzsigntemplateelementdependencyRequest) SetSEzsigntemplateelementdependencyEzsigntemplateformfieldlabel(v string)`

SetSEzsigntemplateelementdependencyEzsigntemplateformfieldlabel sets SEzsigntemplateelementdependencyEzsigntemplateformfieldlabel field to given value.

### HasSEzsigntemplateelementdependencyEzsigntemplateformfieldlabel

`func (o *EzsigntemplateelementdependencyRequest) HasSEzsigntemplateelementdependencyEzsigntemplateformfieldlabel() bool`

HasSEzsigntemplateelementdependencyEzsigntemplateformfieldlabel returns a boolean if a field has been set.

### GetEEzsigntemplateelementdependencyValidation

`func (o *EzsigntemplateelementdependencyRequest) GetEEzsigntemplateelementdependencyValidation() FieldEEzsigntemplateelementdependencyValidation`

GetEEzsigntemplateelementdependencyValidation returns the EEzsigntemplateelementdependencyValidation field if non-nil, zero value otherwise.

### GetEEzsigntemplateelementdependencyValidationOk

`func (o *EzsigntemplateelementdependencyRequest) GetEEzsigntemplateelementdependencyValidationOk() (*FieldEEzsigntemplateelementdependencyValidation, bool)`

GetEEzsigntemplateelementdependencyValidationOk returns a tuple with the EEzsigntemplateelementdependencyValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateelementdependencyValidation

`func (o *EzsigntemplateelementdependencyRequest) SetEEzsigntemplateelementdependencyValidation(v FieldEEzsigntemplateelementdependencyValidation)`

SetEEzsigntemplateelementdependencyValidation sets EEzsigntemplateelementdependencyValidation field to given value.


### GetBEzsigntemplateelementdependencySelected

`func (o *EzsigntemplateelementdependencyRequest) GetBEzsigntemplateelementdependencySelected() bool`

GetBEzsigntemplateelementdependencySelected returns the BEzsigntemplateelementdependencySelected field if non-nil, zero value otherwise.

### GetBEzsigntemplateelementdependencySelectedOk

`func (o *EzsigntemplateelementdependencyRequest) GetBEzsigntemplateelementdependencySelectedOk() (*bool, bool)`

GetBEzsigntemplateelementdependencySelectedOk returns a tuple with the BEzsigntemplateelementdependencySelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateelementdependencySelected

`func (o *EzsigntemplateelementdependencyRequest) SetBEzsigntemplateelementdependencySelected(v bool)`

SetBEzsigntemplateelementdependencySelected sets BEzsigntemplateelementdependencySelected field to given value.

### HasBEzsigntemplateelementdependencySelected

`func (o *EzsigntemplateelementdependencyRequest) HasBEzsigntemplateelementdependencySelected() bool`

HasBEzsigntemplateelementdependencySelected returns a boolean if a field has been set.

### GetEEzsigntemplateelementdependencyOperator

`func (o *EzsigntemplateelementdependencyRequest) GetEEzsigntemplateelementdependencyOperator() FieldEEzsigntemplateelementdependencyOperator`

GetEEzsigntemplateelementdependencyOperator returns the EEzsigntemplateelementdependencyOperator field if non-nil, zero value otherwise.

### GetEEzsigntemplateelementdependencyOperatorOk

`func (o *EzsigntemplateelementdependencyRequest) GetEEzsigntemplateelementdependencyOperatorOk() (*FieldEEzsigntemplateelementdependencyOperator, bool)`

GetEEzsigntemplateelementdependencyOperatorOk returns a tuple with the EEzsigntemplateelementdependencyOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateelementdependencyOperator

`func (o *EzsigntemplateelementdependencyRequest) SetEEzsigntemplateelementdependencyOperator(v FieldEEzsigntemplateelementdependencyOperator)`

SetEEzsigntemplateelementdependencyOperator sets EEzsigntemplateelementdependencyOperator field to given value.

### HasEEzsigntemplateelementdependencyOperator

`func (o *EzsigntemplateelementdependencyRequest) HasEEzsigntemplateelementdependencyOperator() bool`

HasEEzsigntemplateelementdependencyOperator returns a boolean if a field has been set.

### GetSEzsigntemplateelementdependencyValue

`func (o *EzsigntemplateelementdependencyRequest) GetSEzsigntemplateelementdependencyValue() string`

GetSEzsigntemplateelementdependencyValue returns the SEzsigntemplateelementdependencyValue field if non-nil, zero value otherwise.

### GetSEzsigntemplateelementdependencyValueOk

`func (o *EzsigntemplateelementdependencyRequest) GetSEzsigntemplateelementdependencyValueOk() (*string, bool)`

GetSEzsigntemplateelementdependencyValueOk returns a tuple with the SEzsigntemplateelementdependencyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateelementdependencyValue

`func (o *EzsigntemplateelementdependencyRequest) SetSEzsigntemplateelementdependencyValue(v string)`

SetSEzsigntemplateelementdependencyValue sets SEzsigntemplateelementdependencyValue field to given value.

### HasSEzsigntemplateelementdependencyValue

`func (o *EzsigntemplateelementdependencyRequest) HasSEzsigntemplateelementdependencyValue() bool`

HasSEzsigntemplateelementdependencyValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


