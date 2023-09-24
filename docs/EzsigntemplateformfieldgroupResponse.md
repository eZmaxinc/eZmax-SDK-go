# EzsigntemplateformfieldgroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateformfieldgroupID** | **int32** | The unique ID of the Ezsigntemplateformfieldgroup | 
**FkiEzsigntemplatedocumentID** | **int32** | The unique ID of the Ezsigntemplatedocument | 
**EEzsigntemplateformfieldgroupType** | [**FieldEEzsigntemplateformfieldgroupType**](FieldEEzsigntemplateformfieldgroupType.md) |  | 
**EEzsigntemplateformfieldgroupSignerrequirement** | [**FieldEEzsigntemplateformfieldgroupSignerrequirement**](FieldEEzsigntemplateformfieldgroupSignerrequirement.md) |  | 
**SEzsigntemplateformfieldgroupLabel** | **string** | The Label for the Ezsigntemplateformfieldgroup | 
**IEzsigntemplateformfieldgroupStep** | **int32** | The step when the Ezsigntemplatesigner will be invited to fill the form fields | 
**SEzsigntemplateformfieldgroupDefaultvalue** | Pointer to **string** | The default value for the Ezsigntemplateformfieldgroup | [optional] 
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

### NewEzsigntemplateformfieldgroupResponse

`func NewEzsigntemplateformfieldgroupResponse(pkiEzsigntemplateformfieldgroupID int32, fkiEzsigntemplatedocumentID int32, eEzsigntemplateformfieldgroupType FieldEEzsigntemplateformfieldgroupType, eEzsigntemplateformfieldgroupSignerrequirement FieldEEzsigntemplateformfieldgroupSignerrequirement, sEzsigntemplateformfieldgroupLabel string, iEzsigntemplateformfieldgroupStep int32, iEzsigntemplateformfieldgroupFilledmin int32, iEzsigntemplateformfieldgroupFilledmax int32, bEzsigntemplateformfieldgroupReadonly bool, ) *EzsigntemplateformfieldgroupResponse`

NewEzsigntemplateformfieldgroupResponse instantiates a new EzsigntemplateformfieldgroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateformfieldgroupResponseWithDefaults

`func NewEzsigntemplateformfieldgroupResponseWithDefaults() *EzsigntemplateformfieldgroupResponse`

NewEzsigntemplateformfieldgroupResponseWithDefaults instantiates a new EzsigntemplateformfieldgroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateformfieldgroupID

`func (o *EzsigntemplateformfieldgroupResponse) GetPkiEzsigntemplateformfieldgroupID() int32`

GetPkiEzsigntemplateformfieldgroupID returns the PkiEzsigntemplateformfieldgroupID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateformfieldgroupIDOk

`func (o *EzsigntemplateformfieldgroupResponse) GetPkiEzsigntemplateformfieldgroupIDOk() (*int32, bool)`

GetPkiEzsigntemplateformfieldgroupIDOk returns a tuple with the PkiEzsigntemplateformfieldgroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateformfieldgroupID

`func (o *EzsigntemplateformfieldgroupResponse) SetPkiEzsigntemplateformfieldgroupID(v int32)`

SetPkiEzsigntemplateformfieldgroupID sets PkiEzsigntemplateformfieldgroupID field to given value.


### GetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateformfieldgroupResponse) GetFkiEzsigntemplatedocumentID() int32`

GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatedocumentIDOk

`func (o *EzsigntemplateformfieldgroupResponse) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool)`

GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateformfieldgroupResponse) SetFkiEzsigntemplatedocumentID(v int32)`

SetFkiEzsigntemplatedocumentID sets FkiEzsigntemplatedocumentID field to given value.


### GetEEzsigntemplateformfieldgroupType

`func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupType() FieldEEzsigntemplateformfieldgroupType`

GetEEzsigntemplateformfieldgroupType returns the EEzsigntemplateformfieldgroupType field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupTypeOk

`func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupTypeOk() (*FieldEEzsigntemplateformfieldgroupType, bool)`

GetEEzsigntemplateformfieldgroupTypeOk returns a tuple with the EEzsigntemplateformfieldgroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupType

`func (o *EzsigntemplateformfieldgroupResponse) SetEEzsigntemplateformfieldgroupType(v FieldEEzsigntemplateformfieldgroupType)`

SetEEzsigntemplateformfieldgroupType sets EEzsigntemplateformfieldgroupType field to given value.


### GetEEzsigntemplateformfieldgroupSignerrequirement

`func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupSignerrequirement() FieldEEzsigntemplateformfieldgroupSignerrequirement`

GetEEzsigntemplateformfieldgroupSignerrequirement returns the EEzsigntemplateformfieldgroupSignerrequirement field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupSignerrequirementOk

`func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupSignerrequirementOk() (*FieldEEzsigntemplateformfieldgroupSignerrequirement, bool)`

GetEEzsigntemplateformfieldgroupSignerrequirementOk returns a tuple with the EEzsigntemplateformfieldgroupSignerrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupSignerrequirement

`func (o *EzsigntemplateformfieldgroupResponse) SetEEzsigntemplateformfieldgroupSignerrequirement(v FieldEEzsigntemplateformfieldgroupSignerrequirement)`

SetEEzsigntemplateformfieldgroupSignerrequirement sets EEzsigntemplateformfieldgroupSignerrequirement field to given value.


### GetSEzsigntemplateformfieldgroupLabel

`func (o *EzsigntemplateformfieldgroupResponse) GetSEzsigntemplateformfieldgroupLabel() string`

GetSEzsigntemplateformfieldgroupLabel returns the SEzsigntemplateformfieldgroupLabel field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldgroupLabelOk

`func (o *EzsigntemplateformfieldgroupResponse) GetSEzsigntemplateformfieldgroupLabelOk() (*string, bool)`

GetSEzsigntemplateformfieldgroupLabelOk returns a tuple with the SEzsigntemplateformfieldgroupLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldgroupLabel

`func (o *EzsigntemplateformfieldgroupResponse) SetSEzsigntemplateformfieldgroupLabel(v string)`

SetSEzsigntemplateformfieldgroupLabel sets SEzsigntemplateformfieldgroupLabel field to given value.


### GetIEzsigntemplateformfieldgroupStep

`func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupStep() int32`

GetIEzsigntemplateformfieldgroupStep returns the IEzsigntemplateformfieldgroupStep field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupStepOk

`func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupStepOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupStepOk returns a tuple with the IEzsigntemplateformfieldgroupStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupStep

`func (o *EzsigntemplateformfieldgroupResponse) SetIEzsigntemplateformfieldgroupStep(v int32)`

SetIEzsigntemplateformfieldgroupStep sets IEzsigntemplateformfieldgroupStep field to given value.


### GetSEzsigntemplateformfieldgroupDefaultvalue

`func (o *EzsigntemplateformfieldgroupResponse) GetSEzsigntemplateformfieldgroupDefaultvalue() string`

GetSEzsigntemplateformfieldgroupDefaultvalue returns the SEzsigntemplateformfieldgroupDefaultvalue field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldgroupDefaultvalueOk

`func (o *EzsigntemplateformfieldgroupResponse) GetSEzsigntemplateformfieldgroupDefaultvalueOk() (*string, bool)`

GetSEzsigntemplateformfieldgroupDefaultvalueOk returns a tuple with the SEzsigntemplateformfieldgroupDefaultvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldgroupDefaultvalue

`func (o *EzsigntemplateformfieldgroupResponse) SetSEzsigntemplateformfieldgroupDefaultvalue(v string)`

SetSEzsigntemplateformfieldgroupDefaultvalue sets SEzsigntemplateformfieldgroupDefaultvalue field to given value.

### HasSEzsigntemplateformfieldgroupDefaultvalue

`func (o *EzsigntemplateformfieldgroupResponse) HasSEzsigntemplateformfieldgroupDefaultvalue() bool`

HasSEzsigntemplateformfieldgroupDefaultvalue returns a boolean if a field has been set.

### GetIEzsigntemplateformfieldgroupFilledmin

`func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupFilledmin() int32`

GetIEzsigntemplateformfieldgroupFilledmin returns the IEzsigntemplateformfieldgroupFilledmin field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupFilledminOk

`func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupFilledminOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupFilledminOk returns a tuple with the IEzsigntemplateformfieldgroupFilledmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupFilledmin

`func (o *EzsigntemplateformfieldgroupResponse) SetIEzsigntemplateformfieldgroupFilledmin(v int32)`

SetIEzsigntemplateformfieldgroupFilledmin sets IEzsigntemplateformfieldgroupFilledmin field to given value.


### GetIEzsigntemplateformfieldgroupFilledmax

`func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupFilledmax() int32`

GetIEzsigntemplateformfieldgroupFilledmax returns the IEzsigntemplateformfieldgroupFilledmax field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupFilledmaxOk

`func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupFilledmaxOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupFilledmaxOk returns a tuple with the IEzsigntemplateformfieldgroupFilledmax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupFilledmax

`func (o *EzsigntemplateformfieldgroupResponse) SetIEzsigntemplateformfieldgroupFilledmax(v int32)`

SetIEzsigntemplateformfieldgroupFilledmax sets IEzsigntemplateformfieldgroupFilledmax field to given value.


### GetBEzsigntemplateformfieldgroupReadonly

`func (o *EzsigntemplateformfieldgroupResponse) GetBEzsigntemplateformfieldgroupReadonly() bool`

GetBEzsigntemplateformfieldgroupReadonly returns the BEzsigntemplateformfieldgroupReadonly field if non-nil, zero value otherwise.

### GetBEzsigntemplateformfieldgroupReadonlyOk

`func (o *EzsigntemplateformfieldgroupResponse) GetBEzsigntemplateformfieldgroupReadonlyOk() (*bool, bool)`

GetBEzsigntemplateformfieldgroupReadonlyOk returns a tuple with the BEzsigntemplateformfieldgroupReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateformfieldgroupReadonly

`func (o *EzsigntemplateformfieldgroupResponse) SetBEzsigntemplateformfieldgroupReadonly(v bool)`

SetBEzsigntemplateformfieldgroupReadonly sets BEzsigntemplateformfieldgroupReadonly field to given value.


### GetIEzsigntemplateformfieldgroupMaxlength

`func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupMaxlength() int32`

GetIEzsigntemplateformfieldgroupMaxlength returns the IEzsigntemplateformfieldgroupMaxlength field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldgroupMaxlengthOk

`func (o *EzsigntemplateformfieldgroupResponse) GetIEzsigntemplateformfieldgroupMaxlengthOk() (*int32, bool)`

GetIEzsigntemplateformfieldgroupMaxlengthOk returns a tuple with the IEzsigntemplateformfieldgroupMaxlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldgroupMaxlength

`func (o *EzsigntemplateformfieldgroupResponse) SetIEzsigntemplateformfieldgroupMaxlength(v int32)`

SetIEzsigntemplateformfieldgroupMaxlength sets IEzsigntemplateformfieldgroupMaxlength field to given value.

### HasIEzsigntemplateformfieldgroupMaxlength

`func (o *EzsigntemplateformfieldgroupResponse) HasIEzsigntemplateformfieldgroupMaxlength() bool`

HasIEzsigntemplateformfieldgroupMaxlength returns a boolean if a field has been set.

### GetBEzsigntemplateformfieldgroupEncrypted

`func (o *EzsigntemplateformfieldgroupResponse) GetBEzsigntemplateformfieldgroupEncrypted() bool`

GetBEzsigntemplateformfieldgroupEncrypted returns the BEzsigntemplateformfieldgroupEncrypted field if non-nil, zero value otherwise.

### GetBEzsigntemplateformfieldgroupEncryptedOk

`func (o *EzsigntemplateformfieldgroupResponse) GetBEzsigntemplateformfieldgroupEncryptedOk() (*bool, bool)`

GetBEzsigntemplateformfieldgroupEncryptedOk returns a tuple with the BEzsigntemplateformfieldgroupEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateformfieldgroupEncrypted

`func (o *EzsigntemplateformfieldgroupResponse) SetBEzsigntemplateformfieldgroupEncrypted(v bool)`

SetBEzsigntemplateformfieldgroupEncrypted sets BEzsigntemplateformfieldgroupEncrypted field to given value.

### HasBEzsigntemplateformfieldgroupEncrypted

`func (o *EzsigntemplateformfieldgroupResponse) HasBEzsigntemplateformfieldgroupEncrypted() bool`

HasBEzsigntemplateformfieldgroupEncrypted returns a boolean if a field has been set.

### GetSEzsigntemplateformfieldgroupRegexp

`func (o *EzsigntemplateformfieldgroupResponse) GetSEzsigntemplateformfieldgroupRegexp() string`

GetSEzsigntemplateformfieldgroupRegexp returns the SEzsigntemplateformfieldgroupRegexp field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldgroupRegexpOk

`func (o *EzsigntemplateformfieldgroupResponse) GetSEzsigntemplateformfieldgroupRegexpOk() (*string, bool)`

GetSEzsigntemplateformfieldgroupRegexpOk returns a tuple with the SEzsigntemplateformfieldgroupRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldgroupRegexp

`func (o *EzsigntemplateformfieldgroupResponse) SetSEzsigntemplateformfieldgroupRegexp(v string)`

SetSEzsigntemplateformfieldgroupRegexp sets SEzsigntemplateformfieldgroupRegexp field to given value.

### HasSEzsigntemplateformfieldgroupRegexp

`func (o *EzsigntemplateformfieldgroupResponse) HasSEzsigntemplateformfieldgroupRegexp() bool`

HasSEzsigntemplateformfieldgroupRegexp returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldgroupTextvalidation

`func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupTextvalidation() EnumTextvalidation`

GetEEzsigntemplateformfieldgroupTextvalidation returns the EEzsigntemplateformfieldgroupTextvalidation field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupTextvalidationOk

`func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupTextvalidationOk() (*EnumTextvalidation, bool)`

GetEEzsigntemplateformfieldgroupTextvalidationOk returns a tuple with the EEzsigntemplateformfieldgroupTextvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupTextvalidation

`func (o *EzsigntemplateformfieldgroupResponse) SetEEzsigntemplateformfieldgroupTextvalidation(v EnumTextvalidation)`

SetEEzsigntemplateformfieldgroupTextvalidation sets EEzsigntemplateformfieldgroupTextvalidation field to given value.

### HasEEzsigntemplateformfieldgroupTextvalidation

`func (o *EzsigntemplateformfieldgroupResponse) HasEEzsigntemplateformfieldgroupTextvalidation() bool`

HasEEzsigntemplateformfieldgroupTextvalidation returns a boolean if a field has been set.

### GetTEzsigntemplateformfieldgroupTooltip

`func (o *EzsigntemplateformfieldgroupResponse) GetTEzsigntemplateformfieldgroupTooltip() string`

GetTEzsigntemplateformfieldgroupTooltip returns the TEzsigntemplateformfieldgroupTooltip field if non-nil, zero value otherwise.

### GetTEzsigntemplateformfieldgroupTooltipOk

`func (o *EzsigntemplateformfieldgroupResponse) GetTEzsigntemplateformfieldgroupTooltipOk() (*string, bool)`

GetTEzsigntemplateformfieldgroupTooltipOk returns a tuple with the TEzsigntemplateformfieldgroupTooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplateformfieldgroupTooltip

`func (o *EzsigntemplateformfieldgroupResponse) SetTEzsigntemplateformfieldgroupTooltip(v string)`

SetTEzsigntemplateformfieldgroupTooltip sets TEzsigntemplateformfieldgroupTooltip field to given value.

### HasTEzsigntemplateformfieldgroupTooltip

`func (o *EzsigntemplateformfieldgroupResponse) HasTEzsigntemplateformfieldgroupTooltip() bool`

HasTEzsigntemplateformfieldgroupTooltip returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldgroupTooltipposition

`func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupTooltipposition() FieldEEzsigntemplateformfieldgroupTooltipposition`

GetEEzsigntemplateformfieldgroupTooltipposition returns the EEzsigntemplateformfieldgroupTooltipposition field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldgroupTooltippositionOk

`func (o *EzsigntemplateformfieldgroupResponse) GetEEzsigntemplateformfieldgroupTooltippositionOk() (*FieldEEzsigntemplateformfieldgroupTooltipposition, bool)`

GetEEzsigntemplateformfieldgroupTooltippositionOk returns a tuple with the EEzsigntemplateformfieldgroupTooltipposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldgroupTooltipposition

`func (o *EzsigntemplateformfieldgroupResponse) SetEEzsigntemplateformfieldgroupTooltipposition(v FieldEEzsigntemplateformfieldgroupTooltipposition)`

SetEEzsigntemplateformfieldgroupTooltipposition sets EEzsigntemplateformfieldgroupTooltipposition field to given value.

### HasEEzsigntemplateformfieldgroupTooltipposition

`func (o *EzsigntemplateformfieldgroupResponse) HasEEzsigntemplateformfieldgroupTooltipposition() bool`

HasEEzsigntemplateformfieldgroupTooltipposition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


