# EzsigntemplateformfieldgroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateformfieldgroupID** | Pointer to **int32** | The unique ID of the Ezsigntemplateformfieldgroup | [optional] 
**FkiEzsigntemplatedocumentID** | **int32** | The unique ID of the Ezsigntemplatedocument | 
**EEzsigntemplateformfieldgroupType** | [**FieldEEzsigntemplateformfieldgroupType**](FieldEEzsigntemplateformfieldgroupType.md) |  | 
**EEzsigntemplateformfieldgroupSignerrequirement** | [**FieldEEzsigntemplateformfieldgroupSignerrequirement**](FieldEEzsigntemplateformfieldgroupSignerrequirement.md) |  | 
**SEzsigntemplateformfieldgroupLabel** | **string** | The Label for the Ezsigntemplateformfieldgroup | 
**IEzsigntemplateformfieldgroupStep** | **int32** | The step when the Ezsigntemplatesigner will be invited to fill the form fields | 
**SEzsigntemplateformfieldgroupDefaultvalue** | **string** | The default value for the Ezsigntemplateformfieldgroup | 
**IEzsigntemplateformfieldgroupFilledmin** | **int32** | The minimum number of Ezsigntemplateformfield that must be filled in the Ezsigntemplateformfieldgroup | 
**IEzsigntemplateformfieldgroupFilledmax** | **int32** | The maximum number of Ezsigntemplateformfield that must be filled in the Ezsigntemplateformfieldgroup | 
**BEzsigntemplateformfieldgroupReadonly** | **bool** | Whether the Ezsigntemplateformfieldgroup is read only or not. | 
**IEzsigntemplateformfieldgroupMaxlength** | Pointer to **int32** | The maximum length for the value in the Ezsigntemplateformfieldgroup  This can only be set if eEzsigntemplateformfieldgroupType is **Text** or **Textarea** | [optional] 
**BEzsigntemplateformfieldgroupEncrypted** | Pointer to **bool** | Whether the Ezsigntemplateformfieldgroup is encrypted in the database or not. Encrypted values are not displayed on the Ezsigndocument. This can only be set if eEzsigntemplateformfieldgroupType is **Text** or **Textarea** | [optional] 
**SEzsigntemplateformfieldgroupRegexp** | Pointer to **string** | A regular expression to indicate what values are acceptable for the Ezsigntemplateformfieldgroup.  This can only be set if eEzsigntemplateformfieldgroupType is **Text** or **Textarea** | [optional] 
**EEzsigntemplateformfieldgroupTextvalidation** | Pointer to [**EnumTextvalidation**](EnumTextvalidation.md) |  | [optional] 
**TEzsigntemplateformfieldgroupTooltip** | Pointer to **string** | A tooltip that will be presented to Ezsigntemplatesigner about the Ezsigntemplateformfieldgroup | [optional] 
**EEzsigntemplateformfieldgroupTooltipposition** | Pointer to [**FieldEEzsigntemplateformfieldgroupTooltipposition**](FieldEEzsigntemplateformfieldgroupTooltipposition.md) |  | [optional] 

## Methods

### NewEzsigntemplateformfieldgroupRequest

`func NewEzsigntemplateformfieldgroupRequest(fkiEzsigntemplatedocumentID int32, eEzsigntemplateformfieldgroupType FieldEEzsigntemplateformfieldgroupType, eEzsigntemplateformfieldgroupSignerrequirement FieldEEzsigntemplateformfieldgroupSignerrequirement, sEzsigntemplateformfieldgroupLabel string, iEzsigntemplateformfieldgroupStep int32, sEzsigntemplateformfieldgroupDefaultvalue string, iEzsigntemplateformfieldgroupFilledmin int32, iEzsigntemplateformfieldgroupFilledmax int32, bEzsigntemplateformfieldgroupReadonly bool, ) *EzsigntemplateformfieldgroupRequest`

NewEzsigntemplateformfieldgroupRequest instantiates a new EzsigntemplateformfieldgroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateformfieldgroupRequestWithDefaults

`func NewEzsigntemplateformfieldgroupRequestWithDefaults() *EzsigntemplateformfieldgroupRequest`

NewEzsigntemplateformfieldgroupRequestWithDefaults instantiates a new EzsigntemplateformfieldgroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateformfieldgroupID

`func (o *EzsigntemplateformfieldgroupRequest) GetPkiEzsigntemplateformfieldgroupID() int32`

GetPkiEzsigntemplateformfieldgroupID returns the PkiEzsigntemplateformfieldgroupID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateformfieldgroupIDOk

`func (o *EzsigntemplateformfieldgroupRequest) GetPkiEzsigntemplateformfieldgroupIDOk() (*int32, bool)`

GetPkiEzsigntemplateformfieldgroupIDOk returns a tuple with the PkiEzsigntemplateformfieldgroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateformfieldgroupID

`func (o *EzsigntemplateformfieldgroupRequest) SetPkiEzsigntemplateformfieldgroupID(v int32)`

SetPkiEzsigntemplateformfieldgroupID sets PkiEzsigntemplateformfieldgroupID field to given value.

### HasPkiEzsigntemplateformfieldgroupID

`func (o *EzsigntemplateformfieldgroupRequest) HasPkiEzsigntemplateformfieldgroupID() bool`

HasPkiEzsigntemplateformfieldgroupID returns a boolean if a field has been set.

### GetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateformfieldgroupRequest) GetFkiEzsigntemplatedocumentID() int32`

GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatedocumentIDOk

`func (o *EzsigntemplateformfieldgroupRequest) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool)`

GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateformfieldgroupRequest) SetFkiEzsigntemplatedocumentID(v int32)`

SetFkiEzsigntemplatedocumentID sets FkiEzsigntemplatedocumentID field to given value.


### GetEEzsigntemplateformfieldgroupType

`func (o *EzsigntemplateformfieldgroupRequest) GetEEzsigntemplateformfieldgroupType() FieldEEzsigntemplateformfieldgroupType`

GetEEzsigntemplateformfieldgroupType returns the EEzsigntemplateformfieldgroupType field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupTypeOk

`func (o *EzsigntemplateformfieldgroupRequest) GetEEzsigntemplateformfieldgroupTypeOk() (*FieldEEzsigntemplateformfieldgroupType, bool)`

GetEEzsigntemplateformfieldgroupTypeOk returns a tuple with the EEzsigntemplateformfieldgroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupType

`func (o *EzsigntemplateformfieldgroupRequest) SetEEzsigntemplateformfieldgroupType(v FieldEEzsigntemplateformfieldgroupType)`

SetEEzsigntemplateformfieldgroupType sets EEzsigntemplateformfieldgroupType field to given value.


### GetEEzsigntemplateformfieldgroupSignerrequirement

`func (o *EzsigntemplateformfieldgroupRequest) GetEEzsigntemplateformfieldgroupSignerrequirement() FieldEEzsigntemplateformfieldgroupSignerrequirement`

GetEEzsigntemplateformfieldgroupSignerrequirement returns the EEzsigntemplateformfieldgroupSignerrequirement field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupSignerrequirementOk

`func (o *EzsigntemplateformfieldgroupRequest) GetEEzsigntemplateformfieldgroupSignerrequirementOk() (*FieldEEzsigntemplateformfieldgroupSignerrequirement, bool)`

GetEEzsigntemplateformfieldgroupSignerrequirementOk returns a tuple with the EEzsigntemplateformfieldgroupSignerrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupSignerrequirement

`func (o *EzsigntemplateformfieldgroupRequest) SetEEzsigntemplateformfieldgroupSignerrequirement(v FieldEEzsigntemplateformfieldgroupSignerrequirement)`

SetEEzsigntemplateformfieldgroupSignerrequirement sets EEzsigntemplateformfieldgroupSignerrequirement field to given value.


### GetSEzsigntemplateformfieldgroupLabel

`func (o *EzsigntemplateformfieldgroupRequest) GetSEzsigntemplateformfieldgroupLabel() string`

GetSEzsigntemplateformfieldgroupLabel returns the SEzsigntemplateformfieldgroupLabel field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldgroupLabelOk

`func (o *EzsigntemplateformfieldgroupRequest) GetSEzsigntemplateformfieldgroupLabelOk() (*string, bool)`

GetSEzsigntemplateformfieldgroupLabelOk returns a tuple with the SEzsigntemplateformfieldgroupLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldgroupLabel

`func (o *EzsigntemplateformfieldgroupRequest) SetSEzsigntemplateformfieldgroupLabel(v string)`

SetSEzsigntemplateformfieldgroupLabel sets SEzsigntemplateformfieldgroupLabel field to given value.


### GetIEzsigntemplateformfieldgroupStep

`func (o *EzsigntemplateformfieldgroupRequest) GetIEzsigntemplateformfieldgroupStep() int32`

GetIEzsigntemplateformfieldgroupStep returns the IEzsigntemplateformfieldgroupStep field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupStepOk

`func (o *EzsigntemplateformfieldgroupRequest) GetIEzsigntemplateformfieldgroupStepOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupStepOk returns a tuple with the IEzsigntemplateformfieldgroupStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupStep

`func (o *EzsigntemplateformfieldgroupRequest) SetIEzsigntemplateformfieldgroupStep(v int32)`

SetIEzsigntemplateformfieldgroupStep sets IEzsigntemplateformfieldgroupStep field to given value.


### GetSEzsigntemplateformfieldgroupDefaultvalue

`func (o *EzsigntemplateformfieldgroupRequest) GetSEzsigntemplateformfieldgroupDefaultvalue() string`

GetSEzsigntemplateformfieldgroupDefaultvalue returns the SEzsigntemplateformfieldgroupDefaultvalue field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldgroupDefaultvalueOk

`func (o *EzsigntemplateformfieldgroupRequest) GetSEzsigntemplateformfieldgroupDefaultvalueOk() (*string, bool)`

GetSEzsigntemplateformfieldgroupDefaultvalueOk returns a tuple with the SEzsigntemplateformfieldgroupDefaultvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldgroupDefaultvalue

`func (o *EzsigntemplateformfieldgroupRequest) SetSEzsigntemplateformfieldgroupDefaultvalue(v string)`

SetSEzsigntemplateformfieldgroupDefaultvalue sets SEzsigntemplateformfieldgroupDefaultvalue field to given value.


### GetIEzsigntemplateformfieldgroupFilledmin

`func (o *EzsigntemplateformfieldgroupRequest) GetIEzsigntemplateformfieldgroupFilledmin() int32`

GetIEzsigntemplateformfieldgroupFilledmin returns the IEzsigntemplateformfieldgroupFilledmin field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupFilledminOk

`func (o *EzsigntemplateformfieldgroupRequest) GetIEzsigntemplateformfieldgroupFilledminOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupFilledminOk returns a tuple with the IEzsigntemplateformfieldgroupFilledmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupFilledmin

`func (o *EzsigntemplateformfieldgroupRequest) SetIEzsigntemplateformfieldgroupFilledmin(v int32)`

SetIEzsigntemplateformfieldgroupFilledmin sets IEzsigntemplateformfieldgroupFilledmin field to given value.


### GetIEzsigntemplateformfieldgroupFilledmax

`func (o *EzsigntemplateformfieldgroupRequest) GetIEzsigntemplateformfieldgroupFilledmax() int32`

GetIEzsigntemplateformfieldgroupFilledmax returns the IEzsigntemplateformfieldgroupFilledmax field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupFilledmaxOk

`func (o *EzsigntemplateformfieldgroupRequest) GetIEzsigntemplateformfieldgroupFilledmaxOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupFilledmaxOk returns a tuple with the IEzsigntemplateformfieldgroupFilledmax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupFilledmax

`func (o *EzsigntemplateformfieldgroupRequest) SetIEzsigntemplateformfieldgroupFilledmax(v int32)`

SetIEzsigntemplateformfieldgroupFilledmax sets IEzsigntemplateformfieldgroupFilledmax field to given value.


### GetBEzsigntemplateformfieldgroupReadonly

`func (o *EzsigntemplateformfieldgroupRequest) GetBEzsigntemplateformfieldgroupReadonly() bool`

GetBEzsigntemplateformfieldgroupReadonly returns the BEzsigntemplateformfieldgroupReadonly field if non-nil, zero value otherwise.

### GetBEzsigntemplateformfieldgroupReadonlyOk

`func (o *EzsigntemplateformfieldgroupRequest) GetBEzsigntemplateformfieldgroupReadonlyOk() (*bool, bool)`

GetBEzsigntemplateformfieldgroupReadonlyOk returns a tuple with the BEzsigntemplateformfieldgroupReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateformfieldgroupReadonly

`func (o *EzsigntemplateformfieldgroupRequest) SetBEzsigntemplateformfieldgroupReadonly(v bool)`

SetBEzsigntemplateformfieldgroupReadonly sets BEzsigntemplateformfieldgroupReadonly field to given value.


### GetIEzsigntemplateformfieldgroupMaxlength

`func (o *EzsigntemplateformfieldgroupRequest) GetIEzsigntemplateformfieldgroupMaxlength() int32`

GetIEzsigntemplateformfieldgroupMaxlength returns the IEzsigntemplateformfieldgroupMaxlength field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupMaxlengthOk

`func (o *EzsigntemplateformfieldgroupRequest) GetIEzsigntemplateformfieldgroupMaxlengthOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupMaxlengthOk returns a tuple with the IEzsigntemplateformfieldgroupMaxlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupMaxlength

`func (o *EzsigntemplateformfieldgroupRequest) SetIEzsigntemplateformfieldgroupMaxlength(v int32)`

SetIEzsigntemplateformfieldgroupMaxlength sets IEzsigntemplateformfieldgroupMaxlength field to given value.

### HasIEzsigntemplateformfieldgroupMaxlength

`func (o *EzsigntemplateformfieldgroupRequest) HasIEzsigntemplateformfieldgroupMaxlength() bool`

HasIEzsigntemplateformfieldgroupMaxlength returns a boolean if a field has been set.

### GetBEzsigntemplateformfieldgroupEncrypted

`func (o *EzsigntemplateformfieldgroupRequest) GetBEzsigntemplateformfieldgroupEncrypted() bool`

GetBEzsigntemplateformfieldgroupEncrypted returns the BEzsigntemplateformfieldgroupEncrypted field if non-nil, zero value otherwise.

### GetBEzsigntemplateformfieldgroupEncryptedOk

`func (o *EzsigntemplateformfieldgroupRequest) GetBEzsigntemplateformfieldgroupEncryptedOk() (*bool, bool)`

GetBEzsigntemplateformfieldgroupEncryptedOk returns a tuple with the BEzsigntemplateformfieldgroupEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateformfieldgroupEncrypted

`func (o *EzsigntemplateformfieldgroupRequest) SetBEzsigntemplateformfieldgroupEncrypted(v bool)`

SetBEzsigntemplateformfieldgroupEncrypted sets BEzsigntemplateformfieldgroupEncrypted field to given value.

### HasBEzsigntemplateformfieldgroupEncrypted

`func (o *EzsigntemplateformfieldgroupRequest) HasBEzsigntemplateformfieldgroupEncrypted() bool`

HasBEzsigntemplateformfieldgroupEncrypted returns a boolean if a field has been set.

### GetSEzsigntemplateformfieldgroupRegexp

`func (o *EzsigntemplateformfieldgroupRequest) GetSEzsigntemplateformfieldgroupRegexp() string`

GetSEzsigntemplateformfieldgroupRegexp returns the SEzsigntemplateformfieldgroupRegexp field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldgroupRegexpOk

`func (o *EzsigntemplateformfieldgroupRequest) GetSEzsigntemplateformfieldgroupRegexpOk() (*string, bool)`

GetSEzsigntemplateformfieldgroupRegexpOk returns a tuple with the SEzsigntemplateformfieldgroupRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldgroupRegexp

`func (o *EzsigntemplateformfieldgroupRequest) SetSEzsigntemplateformfieldgroupRegexp(v string)`

SetSEzsigntemplateformfieldgroupRegexp sets SEzsigntemplateformfieldgroupRegexp field to given value.

### HasSEzsigntemplateformfieldgroupRegexp

`func (o *EzsigntemplateformfieldgroupRequest) HasSEzsigntemplateformfieldgroupRegexp() bool`

HasSEzsigntemplateformfieldgroupRegexp returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldgroupTextvalidation

`func (o *EzsigntemplateformfieldgroupRequest) GetEEzsigntemplateformfieldgroupTextvalidation() EnumTextvalidation`

GetEEzsigntemplateformfieldgroupTextvalidation returns the EEzsigntemplateformfieldgroupTextvalidation field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupTextvalidationOk

`func (o *EzsigntemplateformfieldgroupRequest) GetEEzsigntemplateformfieldgroupTextvalidationOk() (*EnumTextvalidation, bool)`

GetEEzsigntemplateformfieldgroupTextvalidationOk returns a tuple with the EEzsigntemplateformfieldgroupTextvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupTextvalidation

`func (o *EzsigntemplateformfieldgroupRequest) SetEEzsigntemplateformfieldgroupTextvalidation(v EnumTextvalidation)`

SetEEzsigntemplateformfieldgroupTextvalidation sets EEzsigntemplateformfieldgroupTextvalidation field to given value.

### HasEEzsigntemplateformfieldgroupTextvalidation

`func (o *EzsigntemplateformfieldgroupRequest) HasEEzsigntemplateformfieldgroupTextvalidation() bool`

HasEEzsigntemplateformfieldgroupTextvalidation returns a boolean if a field has been set.

### GetTEzsigntemplateformfieldgroupTooltip

`func (o *EzsigntemplateformfieldgroupRequest) GetTEzsigntemplateformfieldgroupTooltip() string`

GetTEzsigntemplateformfieldgroupTooltip returns the TEzsigntemplateformfieldgroupTooltip field if non-nil, zero value otherwise.

### GetTEzsigntemplateformfieldgroupTooltipOk

`func (o *EzsigntemplateformfieldgroupRequest) GetTEzsigntemplateformfieldgroupTooltipOk() (*string, bool)`

GetTEzsigntemplateformfieldgroupTooltipOk returns a tuple with the TEzsigntemplateformfieldgroupTooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplateformfieldgroupTooltip

`func (o *EzsigntemplateformfieldgroupRequest) SetTEzsigntemplateformfieldgroupTooltip(v string)`

SetTEzsigntemplateformfieldgroupTooltip sets TEzsigntemplateformfieldgroupTooltip field to given value.

### HasTEzsigntemplateformfieldgroupTooltip

`func (o *EzsigntemplateformfieldgroupRequest) HasTEzsigntemplateformfieldgroupTooltip() bool`

HasTEzsigntemplateformfieldgroupTooltip returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldgroupTooltipposition

`func (o *EzsigntemplateformfieldgroupRequest) GetEEzsigntemplateformfieldgroupTooltipposition() FieldEEzsigntemplateformfieldgroupTooltipposition`

GetEEzsigntemplateformfieldgroupTooltipposition returns the EEzsigntemplateformfieldgroupTooltipposition field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupTooltippositionOk

`func (o *EzsigntemplateformfieldgroupRequest) GetEEzsigntemplateformfieldgroupTooltippositionOk() (*FieldEEzsigntemplateformfieldgroupTooltipposition, bool)`

GetEEzsigntemplateformfieldgroupTooltippositionOk returns a tuple with the EEzsigntemplateformfieldgroupTooltipposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupTooltipposition

`func (o *EzsigntemplateformfieldgroupRequest) SetEEzsigntemplateformfieldgroupTooltipposition(v FieldEEzsigntemplateformfieldgroupTooltipposition)`

SetEEzsigntemplateformfieldgroupTooltipposition sets EEzsigntemplateformfieldgroupTooltipposition field to given value.

### HasEEzsigntemplateformfieldgroupTooltipposition

`func (o *EzsigntemplateformfieldgroupRequest) HasEEzsigntemplateformfieldgroupTooltipposition() bool`

HasEEzsigntemplateformfieldgroupTooltipposition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


