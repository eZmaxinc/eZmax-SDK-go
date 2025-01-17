# EzsigndocumentResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EEzsigndocumentSteptype** | [**ComputedEEzsigndocumentSteptype**](ComputedEEzsigndocumentSteptype.md) |  | 
**IEzsigndocumentStepformtotal** | **int32** | The total number of steps in the form filling phase | 
**IEzsigndocumentStepformcurrent** | **int32** | The current step in the form filling phase | 
**IEzsigndocumentStepsignaturetotal** | **int32** | The total number of steps in the signature filling phase | 
**IEzsigndocumentStepsignatureCurrent** | **int32** | The current step in the signature phase | 
**AObjEzsignfoldersignerassociationstatus** | [**[]CustomEzsignfoldersignerassociationstatusResponse**](CustomEzsignfoldersignerassociationstatusResponse.md) |  | 
**AObjEzsigndocumentdependency** | Pointer to [**[]EzsigndocumentdependencyResponse**](EzsigndocumentdependencyResponse.md) |  | [optional] 

## Methods

### NewEzsigndocumentResponseCompound

`func NewEzsigndocumentResponseCompound(eEzsigndocumentSteptype ComputedEEzsigndocumentSteptype, iEzsigndocumentStepformtotal int32, iEzsigndocumentStepformcurrent int32, iEzsigndocumentStepsignaturetotal int32, iEzsigndocumentStepsignatureCurrent int32, aObjEzsignfoldersignerassociationstatus []CustomEzsignfoldersignerassociationstatusResponse, ) *EzsigndocumentResponseCompound`

NewEzsigndocumentResponseCompound instantiates a new EzsigndocumentResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentResponseCompoundWithDefaults

`func NewEzsigndocumentResponseCompoundWithDefaults() *EzsigndocumentResponseCompound`

NewEzsigndocumentResponseCompoundWithDefaults instantiates a new EzsigndocumentResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEEzsigndocumentSteptype

`func (o *EzsigndocumentResponseCompound) GetEEzsigndocumentSteptype() ComputedEEzsigndocumentSteptype`

GetEEzsigndocumentSteptype returns the EEzsigndocumentSteptype field if non-nil, zero value otherwise.

### GetEEzsigndocumentSteptypeOk

`func (o *EzsigndocumentResponseCompound) GetEEzsigndocumentSteptypeOk() (*ComputedEEzsigndocumentSteptype, bool)`

GetEEzsigndocumentSteptypeOk returns a tuple with the EEzsigndocumentSteptype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigndocumentSteptype

`func (o *EzsigndocumentResponseCompound) SetEEzsigndocumentSteptype(v ComputedEEzsigndocumentSteptype)`

SetEEzsigndocumentSteptype sets EEzsigndocumentSteptype field to given value.


### GetIEzsigndocumentStepformtotal

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepformtotal() int32`

GetIEzsigndocumentStepformtotal returns the IEzsigndocumentStepformtotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentStepformtotalOk

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepformtotalOk() (*int32, bool)`

GetIEzsigndocumentStepformtotalOk returns a tuple with the IEzsigndocumentStepformtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentStepformtotal

`func (o *EzsigndocumentResponseCompound) SetIEzsigndocumentStepformtotal(v int32)`

SetIEzsigndocumentStepformtotal sets IEzsigndocumentStepformtotal field to given value.


### GetIEzsigndocumentStepformcurrent

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepformcurrent() int32`

GetIEzsigndocumentStepformcurrent returns the IEzsigndocumentStepformcurrent field if non-nil, zero value otherwise.

### GetIEzsigndocumentStepformcurrentOk

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepformcurrentOk() (*int32, bool)`

GetIEzsigndocumentStepformcurrentOk returns a tuple with the IEzsigndocumentStepformcurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentStepformcurrent

`func (o *EzsigndocumentResponseCompound) SetIEzsigndocumentStepformcurrent(v int32)`

SetIEzsigndocumentStepformcurrent sets IEzsigndocumentStepformcurrent field to given value.


### GetIEzsigndocumentStepsignaturetotal

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepsignaturetotal() int32`

GetIEzsigndocumentStepsignaturetotal returns the IEzsigndocumentStepsignaturetotal field if non-nil, zero value otherwise.

### GetIEzsigndocumentStepsignaturetotalOk

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepsignaturetotalOk() (*int32, bool)`

GetIEzsigndocumentStepsignaturetotalOk returns a tuple with the IEzsigndocumentStepsignaturetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentStepsignaturetotal

`func (o *EzsigndocumentResponseCompound) SetIEzsigndocumentStepsignaturetotal(v int32)`

SetIEzsigndocumentStepsignaturetotal sets IEzsigndocumentStepsignaturetotal field to given value.


### GetIEzsigndocumentStepsignatureCurrent

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepsignatureCurrent() int32`

GetIEzsigndocumentStepsignatureCurrent returns the IEzsigndocumentStepsignatureCurrent field if non-nil, zero value otherwise.

### GetIEzsigndocumentStepsignatureCurrentOk

`func (o *EzsigndocumentResponseCompound) GetIEzsigndocumentStepsignatureCurrentOk() (*int32, bool)`

GetIEzsigndocumentStepsignatureCurrentOk returns a tuple with the IEzsigndocumentStepsignatureCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocumentStepsignatureCurrent

`func (o *EzsigndocumentResponseCompound) SetIEzsigndocumentStepsignatureCurrent(v int32)`

SetIEzsigndocumentStepsignatureCurrent sets IEzsigndocumentStepsignatureCurrent field to given value.


### GetAObjEzsignfoldersignerassociationstatus

`func (o *EzsigndocumentResponseCompound) GetAObjEzsignfoldersignerassociationstatus() []CustomEzsignfoldersignerassociationstatusResponse`

GetAObjEzsignfoldersignerassociationstatus returns the AObjEzsignfoldersignerassociationstatus field if non-nil, zero value otherwise.

### GetAObjEzsignfoldersignerassociationstatusOk

`func (o *EzsigndocumentResponseCompound) GetAObjEzsignfoldersignerassociationstatusOk() (*[]CustomEzsignfoldersignerassociationstatusResponse, bool)`

GetAObjEzsignfoldersignerassociationstatusOk returns a tuple with the AObjEzsignfoldersignerassociationstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignfoldersignerassociationstatus

`func (o *EzsigndocumentResponseCompound) SetAObjEzsignfoldersignerassociationstatus(v []CustomEzsignfoldersignerassociationstatusResponse)`

SetAObjEzsignfoldersignerassociationstatus sets AObjEzsignfoldersignerassociationstatus field to given value.


### GetAObjEzsigndocumentdependency

`func (o *EzsigndocumentResponseCompound) GetAObjEzsigndocumentdependency() []EzsigndocumentdependencyResponse`

GetAObjEzsigndocumentdependency returns the AObjEzsigndocumentdependency field if non-nil, zero value otherwise.

### GetAObjEzsigndocumentdependencyOk

`func (o *EzsigndocumentResponseCompound) GetAObjEzsigndocumentdependencyOk() (*[]EzsigndocumentdependencyResponse, bool)`

GetAObjEzsigndocumentdependencyOk returns a tuple with the AObjEzsigndocumentdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigndocumentdependency

`func (o *EzsigndocumentResponseCompound) SetAObjEzsigndocumentdependency(v []EzsigndocumentdependencyResponse)`

SetAObjEzsigndocumentdependency sets AObjEzsigndocumentdependency field to given value.

### HasAObjEzsigndocumentdependency

`func (o *EzsigndocumentResponseCompound) HasAObjEzsigndocumentdependency() bool`

HasAObjEzsigndocumentdependency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


