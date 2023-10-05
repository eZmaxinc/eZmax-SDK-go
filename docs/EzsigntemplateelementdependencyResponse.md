# EzsigntemplateelementdependencyResponse

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

### NewEzsigntemplateelementdependencyResponse

`func NewEzsigntemplateelementdependencyResponse(pkiEzsigntemplateelementdependencyID int32, eEzsigntemplateelementdependencyValidation FieldEEzsigntemplateelementdependencyValidation, ) *EzsigntemplateelementdependencyResponse`

NewEzsigntemplateelementdependencyResponse instantiates a new EzsigntemplateelementdependencyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateelementdependencyResponseWithDefaults

`func NewEzsigntemplateelementdependencyResponseWithDefaults() *EzsigntemplateelementdependencyResponse`

NewEzsigntemplateelementdependencyResponseWithDefaults instantiates a new EzsigntemplateelementdependencyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateelementdependencyID

`func (o *EzsigntemplateelementdependencyResponse) GetPkiEzsigntemplateelementdependencyID() int32`

GetPkiEzsigntemplateelementdependencyID returns the PkiEzsigntemplateelementdependencyID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateelementdependencyIDOk

`func (o *EzsigntemplateelementdependencyResponse) GetPkiEzsigntemplateelementdependencyIDOk() (*int32, bool)`

GetPkiEzsigntemplateelementdependencyIDOk returns a tuple with the PkiEzsigntemplateelementdependencyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateelementdependencyID

`func (o *EzsigntemplateelementdependencyResponse) SetPkiEzsigntemplateelementdependencyID(v int32)`

SetPkiEzsigntemplateelementdependencyID sets PkiEzsigntemplateelementdependencyID field to given value.


### GetFkiEzsigntemplateformfieldID

`func (o *EzsigntemplateelementdependencyResponse) GetFkiEzsigntemplateformfieldID() int32`

GetFkiEzsigntemplateformfieldID returns the FkiEzsigntemplateformfieldID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateformfieldIDOk

`func (o *EzsigntemplateelementdependencyResponse) GetFkiEzsigntemplateformfieldIDOk() (*int32, bool)`

GetFkiEzsigntemplateformfieldIDOk returns a tuple with the FkiEzsigntemplateformfieldID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateformfieldID

`func (o *EzsigntemplateelementdependencyResponse) SetFkiEzsigntemplateformfieldID(v int32)`

SetFkiEzsigntemplateformfieldID sets FkiEzsigntemplateformfieldID field to given value.

### HasFkiEzsigntemplateformfieldID

`func (o *EzsigntemplateelementdependencyResponse) HasFkiEzsigntemplateformfieldID() bool`

HasFkiEzsigntemplateformfieldID returns a boolean if a field has been set.

### GetFkiEzsigntemplatesignatureID

`func (o *EzsigntemplateelementdependencyResponse) GetFkiEzsigntemplatesignatureID() int32`

GetFkiEzsigntemplatesignatureID returns the FkiEzsigntemplatesignatureID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatesignatureIDOk

`func (o *EzsigntemplateelementdependencyResponse) GetFkiEzsigntemplatesignatureIDOk() (*int32, bool)`

GetFkiEzsigntemplatesignatureIDOk returns a tuple with the FkiEzsigntemplatesignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatesignatureID

`func (o *EzsigntemplateelementdependencyResponse) SetFkiEzsigntemplatesignatureID(v int32)`

SetFkiEzsigntemplatesignatureID sets FkiEzsigntemplatesignatureID field to given value.

### HasFkiEzsigntemplatesignatureID

`func (o *EzsigntemplateelementdependencyResponse) HasFkiEzsigntemplatesignatureID() bool`

HasFkiEzsigntemplatesignatureID returns a boolean if a field has been set.

### GetFkiEzsigntemplateformfieldIDValidation

`func (o *EzsigntemplateelementdependencyResponse) GetFkiEzsigntemplateformfieldIDValidation() int32`

GetFkiEzsigntemplateformfieldIDValidation returns the FkiEzsigntemplateformfieldIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateformfieldIDValidationOk

`func (o *EzsigntemplateelementdependencyResponse) GetFkiEzsigntemplateformfieldIDValidationOk() (*int32, bool)`

GetFkiEzsigntemplateformfieldIDValidationOk returns a tuple with the FkiEzsigntemplateformfieldIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateformfieldIDValidation

`func (o *EzsigntemplateelementdependencyResponse) SetFkiEzsigntemplateformfieldIDValidation(v int32)`

SetFkiEzsigntemplateformfieldIDValidation sets FkiEzsigntemplateformfieldIDValidation field to given value.

### HasFkiEzsigntemplateformfieldIDValidation

`func (o *EzsigntemplateelementdependencyResponse) HasFkiEzsigntemplateformfieldIDValidation() bool`

HasFkiEzsigntemplateformfieldIDValidation returns a boolean if a field has been set.

### GetFkiEzsigntemplateformfieldgroupIDValidation

`func (o *EzsigntemplateelementdependencyResponse) GetFkiEzsigntemplateformfieldgroupIDValidation() int32`

GetFkiEzsigntemplateformfieldgroupIDValidation returns the FkiEzsigntemplateformfieldgroupIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateformfieldgroupIDValidationOk

`func (o *EzsigntemplateelementdependencyResponse) GetFkiEzsigntemplateformfieldgroupIDValidationOk() (*int32, bool)`

GetFkiEzsigntemplateformfieldgroupIDValidationOk returns a tuple with the FkiEzsigntemplateformfieldgroupIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateformfieldgroupIDValidation

`func (o *EzsigntemplateelementdependencyResponse) SetFkiEzsigntemplateformfieldgroupIDValidation(v int32)`

SetFkiEzsigntemplateformfieldgroupIDValidation sets FkiEzsigntemplateformfieldgroupIDValidation field to given value.

### HasFkiEzsigntemplateformfieldgroupIDValidation

`func (o *EzsigntemplateelementdependencyResponse) HasFkiEzsigntemplateformfieldgroupIDValidation() bool`

HasFkiEzsigntemplateformfieldgroupIDValidation returns a boolean if a field has been set.

### GetEEzsigntemplateelementdependencyValidation

`func (o *EzsigntemplateelementdependencyResponse) GetEEzsigntemplateelementdependencyValidation() FieldEEzsigntemplateelementdependencyValidation`

GetEEzsigntemplateelementdependencyValidation returns the EEzsigntemplateelementdependencyValidation field if non-nil, zero value otherwise.

### GetEEzsigntemplateelementdependencyValidationOk

`func (o *EzsigntemplateelementdependencyResponse) GetEEzsigntemplateelementdependencyValidationOk() (*FieldEEzsigntemplateelementdependencyValidation, bool)`

GetEEzsigntemplateelementdependencyValidationOk returns a tuple with the EEzsigntemplateelementdependencyValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateelementdependencyValidation

`func (o *EzsigntemplateelementdependencyResponse) SetEEzsigntemplateelementdependencyValidation(v FieldEEzsigntemplateelementdependencyValidation)`

SetEEzsigntemplateelementdependencyValidation sets EEzsigntemplateelementdependencyValidation field to given value.


### GetBEzsigntemplateelementdependencySelected

`func (o *EzsigntemplateelementdependencyResponse) GetBEzsigntemplateelementdependencySelected() bool`

GetBEzsigntemplateelementdependencySelected returns the BEzsigntemplateelementdependencySelected field if non-nil, zero value otherwise.

### GetBEzsigntemplateelementdependencySelectedOk

`func (o *EzsigntemplateelementdependencyResponse) GetBEzsigntemplateelementdependencySelectedOk() (*bool, bool)`

GetBEzsigntemplateelementdependencySelectedOk returns a tuple with the BEzsigntemplateelementdependencySelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateelementdependencySelected

`func (o *EzsigntemplateelementdependencyResponse) SetBEzsigntemplateelementdependencySelected(v bool)`

SetBEzsigntemplateelementdependencySelected sets BEzsigntemplateelementdependencySelected field to given value.

### HasBEzsigntemplateelementdependencySelected

`func (o *EzsigntemplateelementdependencyResponse) HasBEzsigntemplateelementdependencySelected() bool`

HasBEzsigntemplateelementdependencySelected returns a boolean if a field has been set.

### GetEEzsigntemplateelementdependencyOperator

`func (o *EzsigntemplateelementdependencyResponse) GetEEzsigntemplateelementdependencyOperator() FieldEEzsigntemplateelementdependencyOperator`

GetEEzsigntemplateelementdependencyOperator returns the EEzsigntemplateelementdependencyOperator field if non-nil, zero value otherwise.

### GetEEzsigntemplateelementdependencyOperatorOk

`func (o *EzsigntemplateelementdependencyResponse) GetEEzsigntemplateelementdependencyOperatorOk() (*FieldEEzsigntemplateelementdependencyOperator, bool)`

GetEEzsigntemplateelementdependencyOperatorOk returns a tuple with the EEzsigntemplateelementdependencyOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateelementdependencyOperator

`func (o *EzsigntemplateelementdependencyResponse) SetEEzsigntemplateelementdependencyOperator(v FieldEEzsigntemplateelementdependencyOperator)`

SetEEzsigntemplateelementdependencyOperator sets EEzsigntemplateelementdependencyOperator field to given value.

### HasEEzsigntemplateelementdependencyOperator

`func (o *EzsigntemplateelementdependencyResponse) HasEEzsigntemplateelementdependencyOperator() bool`

HasEEzsigntemplateelementdependencyOperator returns a boolean if a field has been set.

### GetSEzsigntemplateelementdependencyValue

`func (o *EzsigntemplateelementdependencyResponse) GetSEzsigntemplateelementdependencyValue() string`

GetSEzsigntemplateelementdependencyValue returns the SEzsigntemplateelementdependencyValue field if non-nil, zero value otherwise.

### GetSEzsigntemplateelementdependencyValueOk

`func (o *EzsigntemplateelementdependencyResponse) GetSEzsigntemplateelementdependencyValueOk() (*string, bool)`

GetSEzsigntemplateelementdependencyValueOk returns a tuple with the SEzsigntemplateelementdependencyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateelementdependencyValue

`func (o *EzsigntemplateelementdependencyResponse) SetSEzsigntemplateelementdependencyValue(v string)`

SetSEzsigntemplateelementdependencyValue sets SEzsigntemplateelementdependencyValue field to given value.

### HasSEzsigntemplateelementdependencyValue

`func (o *EzsigntemplateelementdependencyResponse) HasSEzsigntemplateelementdependencyValue() bool`

HasSEzsigntemplateelementdependencyValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


