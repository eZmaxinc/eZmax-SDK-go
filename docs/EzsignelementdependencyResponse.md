# EzsignelementdependencyResponse

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

### NewEzsignelementdependencyResponse

`func NewEzsignelementdependencyResponse(pkiEzsignelementdependencyID int32, eEzsignelementdependencyValidation FieldEEzsignelementdependencyValidation, ) *EzsignelementdependencyResponse`

NewEzsignelementdependencyResponse instantiates a new EzsignelementdependencyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignelementdependencyResponseWithDefaults

`func NewEzsignelementdependencyResponseWithDefaults() *EzsignelementdependencyResponse`

NewEzsignelementdependencyResponseWithDefaults instantiates a new EzsignelementdependencyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignelementdependencyID

`func (o *EzsignelementdependencyResponse) GetPkiEzsignelementdependencyID() int32`

GetPkiEzsignelementdependencyID returns the PkiEzsignelementdependencyID field if non-nil, zero value otherwise.

### GetPkiEzsignelementdependencyIDOk

`func (o *EzsignelementdependencyResponse) GetPkiEzsignelementdependencyIDOk() (*int32, bool)`

GetPkiEzsignelementdependencyIDOk returns a tuple with the PkiEzsignelementdependencyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignelementdependencyID

`func (o *EzsignelementdependencyResponse) SetPkiEzsignelementdependencyID(v int32)`

SetPkiEzsignelementdependencyID sets PkiEzsignelementdependencyID field to given value.


### GetFkiEzsignformfieldID

`func (o *EzsignelementdependencyResponse) GetFkiEzsignformfieldID() int32`

GetFkiEzsignformfieldID returns the FkiEzsignformfieldID field if non-nil, zero value otherwise.

### GetFkiEzsignformfieldIDOk

`func (o *EzsignelementdependencyResponse) GetFkiEzsignformfieldIDOk() (*int32, bool)`

GetFkiEzsignformfieldIDOk returns a tuple with the FkiEzsignformfieldID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignformfieldID

`func (o *EzsignelementdependencyResponse) SetFkiEzsignformfieldID(v int32)`

SetFkiEzsignformfieldID sets FkiEzsignformfieldID field to given value.

### HasFkiEzsignformfieldID

`func (o *EzsignelementdependencyResponse) HasFkiEzsignformfieldID() bool`

HasFkiEzsignformfieldID returns a boolean if a field has been set.

### GetFkiEzsignsignatureID

`func (o *EzsignelementdependencyResponse) GetFkiEzsignsignatureID() int32`

GetFkiEzsignsignatureID returns the FkiEzsignsignatureID field if non-nil, zero value otherwise.

### GetFkiEzsignsignatureIDOk

`func (o *EzsignelementdependencyResponse) GetFkiEzsignsignatureIDOk() (*int32, bool)`

GetFkiEzsignsignatureIDOk returns a tuple with the FkiEzsignsignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsignatureID

`func (o *EzsignelementdependencyResponse) SetFkiEzsignsignatureID(v int32)`

SetFkiEzsignsignatureID sets FkiEzsignsignatureID field to given value.

### HasFkiEzsignsignatureID

`func (o *EzsignelementdependencyResponse) HasFkiEzsignsignatureID() bool`

HasFkiEzsignsignatureID returns a boolean if a field has been set.

### GetFkiEzsignformfieldIDValidation

`func (o *EzsignelementdependencyResponse) GetFkiEzsignformfieldIDValidation() int32`

GetFkiEzsignformfieldIDValidation returns the FkiEzsignformfieldIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsignformfieldIDValidationOk

`func (o *EzsignelementdependencyResponse) GetFkiEzsignformfieldIDValidationOk() (*int32, bool)`

GetFkiEzsignformfieldIDValidationOk returns a tuple with the FkiEzsignformfieldIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignformfieldIDValidation

`func (o *EzsignelementdependencyResponse) SetFkiEzsignformfieldIDValidation(v int32)`

SetFkiEzsignformfieldIDValidation sets FkiEzsignformfieldIDValidation field to given value.

### HasFkiEzsignformfieldIDValidation

`func (o *EzsignelementdependencyResponse) HasFkiEzsignformfieldIDValidation() bool`

HasFkiEzsignformfieldIDValidation returns a boolean if a field has been set.

### GetFkiEzsignformfieldgroupIDValidation

`func (o *EzsignelementdependencyResponse) GetFkiEzsignformfieldgroupIDValidation() int32`

GetFkiEzsignformfieldgroupIDValidation returns the FkiEzsignformfieldgroupIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsignformfieldgroupIDValidationOk

`func (o *EzsignelementdependencyResponse) GetFkiEzsignformfieldgroupIDValidationOk() (*int32, bool)`

GetFkiEzsignformfieldgroupIDValidationOk returns a tuple with the FkiEzsignformfieldgroupIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignformfieldgroupIDValidation

`func (o *EzsignelementdependencyResponse) SetFkiEzsignformfieldgroupIDValidation(v int32)`

SetFkiEzsignformfieldgroupIDValidation sets FkiEzsignformfieldgroupIDValidation field to given value.

### HasFkiEzsignformfieldgroupIDValidation

`func (o *EzsignelementdependencyResponse) HasFkiEzsignformfieldgroupIDValidation() bool`

HasFkiEzsignformfieldgroupIDValidation returns a boolean if a field has been set.

### GetEEzsignelementdependencyValidation

`func (o *EzsignelementdependencyResponse) GetEEzsignelementdependencyValidation() FieldEEzsignelementdependencyValidation`

GetEEzsignelementdependencyValidation returns the EEzsignelementdependencyValidation field if non-nil, zero value otherwise.

### GetEEzsignelementdependencyValidationOk

`func (o *EzsignelementdependencyResponse) GetEEzsignelementdependencyValidationOk() (*FieldEEzsignelementdependencyValidation, bool)`

GetEEzsignelementdependencyValidationOk returns a tuple with the EEzsignelementdependencyValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignelementdependencyValidation

`func (o *EzsignelementdependencyResponse) SetEEzsignelementdependencyValidation(v FieldEEzsignelementdependencyValidation)`

SetEEzsignelementdependencyValidation sets EEzsignelementdependencyValidation field to given value.


### GetBEzsignelementdependencySelected

`func (o *EzsignelementdependencyResponse) GetBEzsignelementdependencySelected() bool`

GetBEzsignelementdependencySelected returns the BEzsignelementdependencySelected field if non-nil, zero value otherwise.

### GetBEzsignelementdependencySelectedOk

`func (o *EzsignelementdependencyResponse) GetBEzsignelementdependencySelectedOk() (*bool, bool)`

GetBEzsignelementdependencySelectedOk returns a tuple with the BEzsignelementdependencySelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignelementdependencySelected

`func (o *EzsignelementdependencyResponse) SetBEzsignelementdependencySelected(v bool)`

SetBEzsignelementdependencySelected sets BEzsignelementdependencySelected field to given value.

### HasBEzsignelementdependencySelected

`func (o *EzsignelementdependencyResponse) HasBEzsignelementdependencySelected() bool`

HasBEzsignelementdependencySelected returns a boolean if a field has been set.

### GetEEzsignelementdependencyOperator

`func (o *EzsignelementdependencyResponse) GetEEzsignelementdependencyOperator() FieldEEzsignelementdependencyOperator`

GetEEzsignelementdependencyOperator returns the EEzsignelementdependencyOperator field if non-nil, zero value otherwise.

### GetEEzsignelementdependencyOperatorOk

`func (o *EzsignelementdependencyResponse) GetEEzsignelementdependencyOperatorOk() (*FieldEEzsignelementdependencyOperator, bool)`

GetEEzsignelementdependencyOperatorOk returns a tuple with the EEzsignelementdependencyOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignelementdependencyOperator

`func (o *EzsignelementdependencyResponse) SetEEzsignelementdependencyOperator(v FieldEEzsignelementdependencyOperator)`

SetEEzsignelementdependencyOperator sets EEzsignelementdependencyOperator field to given value.

### HasEEzsignelementdependencyOperator

`func (o *EzsignelementdependencyResponse) HasEEzsignelementdependencyOperator() bool`

HasEEzsignelementdependencyOperator returns a boolean if a field has been set.

### GetSEzsignelementdependencyValue

`func (o *EzsignelementdependencyResponse) GetSEzsignelementdependencyValue() string`

GetSEzsignelementdependencyValue returns the SEzsignelementdependencyValue field if non-nil, zero value otherwise.

### GetSEzsignelementdependencyValueOk

`func (o *EzsignelementdependencyResponse) GetSEzsignelementdependencyValueOk() (*string, bool)`

GetSEzsignelementdependencyValueOk returns a tuple with the SEzsignelementdependencyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignelementdependencyValue

`func (o *EzsignelementdependencyResponse) SetSEzsignelementdependencyValue(v string)`

SetSEzsignelementdependencyValue sets SEzsignelementdependencyValue field to given value.

### HasSEzsignelementdependencyValue

`func (o *EzsignelementdependencyResponse) HasSEzsignelementdependencyValue() bool`

HasSEzsignelementdependencyValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


