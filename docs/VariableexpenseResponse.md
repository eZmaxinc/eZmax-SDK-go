# VariableexpenseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiVariableexpenseID** | **int32** | The unique ID of the Variableexpense | 
**SVariableexpenseCode** | Pointer to **string** | The code of the Variableexpense | [optional] 
**ObjVariableexpenseDescription** | [**MultilingualVariableexpenseDescription**](MultilingualVariableexpenseDescription.md) |  | 
**EVariableexpenseTaxable** | Pointer to [**FieldEVariableexpenseTaxable**](FieldEVariableexpenseTaxable.md) |  | [optional] 
**BVariableexpenseIsactive** | Pointer to **bool** | Whether the variableexpense is active or not | [optional] 

## Methods

### NewVariableexpenseResponse

`func NewVariableexpenseResponse(pkiVariableexpenseID int32, objVariableexpenseDescription MultilingualVariableexpenseDescription, ) *VariableexpenseResponse`

NewVariableexpenseResponse instantiates a new VariableexpenseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableexpenseResponseWithDefaults

`func NewVariableexpenseResponseWithDefaults() *VariableexpenseResponse`

NewVariableexpenseResponseWithDefaults instantiates a new VariableexpenseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiVariableexpenseID

`func (o *VariableexpenseResponse) GetPkiVariableexpenseID() int32`

GetPkiVariableexpenseID returns the PkiVariableexpenseID field if non-nil, zero value otherwise.

### GetPkiVariableexpenseIDOk

`func (o *VariableexpenseResponse) GetPkiVariableexpenseIDOk() (*int32, bool)`

GetPkiVariableexpenseIDOk returns a tuple with the PkiVariableexpenseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiVariableexpenseID

`func (o *VariableexpenseResponse) SetPkiVariableexpenseID(v int32)`

SetPkiVariableexpenseID sets PkiVariableexpenseID field to given value.


### GetSVariableexpenseCode

`func (o *VariableexpenseResponse) GetSVariableexpenseCode() string`

GetSVariableexpenseCode returns the SVariableexpenseCode field if non-nil, zero value otherwise.

### GetSVariableexpenseCodeOk

`func (o *VariableexpenseResponse) GetSVariableexpenseCodeOk() (*string, bool)`

GetSVariableexpenseCodeOk returns a tuple with the SVariableexpenseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSVariableexpenseCode

`func (o *VariableexpenseResponse) SetSVariableexpenseCode(v string)`

SetSVariableexpenseCode sets SVariableexpenseCode field to given value.

### HasSVariableexpenseCode

`func (o *VariableexpenseResponse) HasSVariableexpenseCode() bool`

HasSVariableexpenseCode returns a boolean if a field has been set.

### GetObjVariableexpenseDescription

`func (o *VariableexpenseResponse) GetObjVariableexpenseDescription() MultilingualVariableexpenseDescription`

GetObjVariableexpenseDescription returns the ObjVariableexpenseDescription field if non-nil, zero value otherwise.

### GetObjVariableexpenseDescriptionOk

`func (o *VariableexpenseResponse) GetObjVariableexpenseDescriptionOk() (*MultilingualVariableexpenseDescription, bool)`

GetObjVariableexpenseDescriptionOk returns a tuple with the ObjVariableexpenseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjVariableexpenseDescription

`func (o *VariableexpenseResponse) SetObjVariableexpenseDescription(v MultilingualVariableexpenseDescription)`

SetObjVariableexpenseDescription sets ObjVariableexpenseDescription field to given value.


### GetEVariableexpenseTaxable

`func (o *VariableexpenseResponse) GetEVariableexpenseTaxable() FieldEVariableexpenseTaxable`

GetEVariableexpenseTaxable returns the EVariableexpenseTaxable field if non-nil, zero value otherwise.

### GetEVariableexpenseTaxableOk

`func (o *VariableexpenseResponse) GetEVariableexpenseTaxableOk() (*FieldEVariableexpenseTaxable, bool)`

GetEVariableexpenseTaxableOk returns a tuple with the EVariableexpenseTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEVariableexpenseTaxable

`func (o *VariableexpenseResponse) SetEVariableexpenseTaxable(v FieldEVariableexpenseTaxable)`

SetEVariableexpenseTaxable sets EVariableexpenseTaxable field to given value.

### HasEVariableexpenseTaxable

`func (o *VariableexpenseResponse) HasEVariableexpenseTaxable() bool`

HasEVariableexpenseTaxable returns a boolean if a field has been set.

### GetBVariableexpenseIsactive

`func (o *VariableexpenseResponse) GetBVariableexpenseIsactive() bool`

GetBVariableexpenseIsactive returns the BVariableexpenseIsactive field if non-nil, zero value otherwise.

### GetBVariableexpenseIsactiveOk

`func (o *VariableexpenseResponse) GetBVariableexpenseIsactiveOk() (*bool, bool)`

GetBVariableexpenseIsactiveOk returns a tuple with the BVariableexpenseIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBVariableexpenseIsactive

`func (o *VariableexpenseResponse) SetBVariableexpenseIsactive(v bool)`

SetBVariableexpenseIsactive sets BVariableexpenseIsactive field to given value.

### HasBVariableexpenseIsactive

`func (o *VariableexpenseResponse) HasBVariableexpenseIsactive() bool`

HasBVariableexpenseIsactive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


