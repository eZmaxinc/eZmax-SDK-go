# EzsigntemplatesignatureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatesignatureID** | **int32** | The unique ID of the Ezsigntemplatesignature | 
**FkiEzsigntemplatedocumentID** | **int32** | The unique ID of the Ezsigntemplatedocument | 
**FkiEzsigntemplatesignerID** | **int32** | The unique ID of the Ezsigntemplatesigner | 
**FkiEzsigntemplatesignerIDValidation** | Pointer to **int32** | The unique ID of the Ezsigntemplatesigner | [optional] 
**IEzsigntemplatedocumentpagePagenumber** | **int32** | The page number in the Ezsigntemplatedocument | 
**IEzsigntemplatesignatureX** | **int32** | The X coordinate (Horizontal) where to put the Ezsigntemplatesignature on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignature 2 inches from the left border of the page, you would use \&quot;200\&quot; for the X coordinate. | 
**IEzsigntemplatesignatureY** | **int32** | The Y coordinate (Vertical) where to put the Ezsigntemplatesignature on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignature 3 inches from the top border of the page, you would use \&quot;300\&quot; for the Y coordinate. | 
**IEzsigntemplatesignatureWidth** | Pointer to **int32** | The width of the Ezsigntemplatesignature.  Size is calculated at 100dpi (dot per inch). So for example, if you want the Ezsigntemplatesignature to have a width of 2 inches, you would use \&quot;200\&quot; for the iEzsigntemplatesignatureWidth. | [optional] 
**IEzsigntemplatesignatureHeight** | Pointer to **int32** | The height of the Ezsigntemplatesignature.  Size is calculated at 100dpi (dot per inch). So for example, if you want the Ezsigntemplatesignature to have an height of 2 inches, you would use \&quot;200\&quot; for the iEzsigntemplatesignatureHeight. | [optional] 
**IEzsigntemplatesignatureStep** | **int32** | The step when the Ezsigntemplatesigner will be invited to sign | 
**EEzsigntemplatesignatureType** | [**FieldEEzsigntemplatesignatureType**](FieldEEzsigntemplatesignatureType.md) |  | 
**TEzsigntemplatesignatureTooltip** | Pointer to **string** | A tooltip that will be presented to Ezsigntemplatesigner about the Ezsigntemplatesignature | [optional] 
**EEzsigntemplatesignatureTooltipposition** | Pointer to [**FieldEEzsigntemplatesignatureTooltipposition**](FieldEEzsigntemplatesignatureTooltipposition.md) |  | [optional] 
**EEzsigntemplatesignatureFont** | Pointer to [**FieldEEzsigntemplatesignatureFont**](FieldEEzsigntemplatesignatureFont.md) |  | [optional] 
**IEzsigntemplatesignatureValidationstep** | Pointer to **int32** | The step when the Ezsigntemplatesigner will be invited to validate the Ezsigntemplatesignature of eEzsigntemplatesignatureType Attachments | [optional] 
**SEzsigntemplatesignatureAttachmentdescription** | Pointer to **string** | The description attached to the attachment name added in Ezsigntemplatesignature of eEzsigntemplatesignatureType Attachments | [optional] 
**EEzsigntemplatesignatureAttachmentnamesource** | Pointer to [**FieldEEzsigntemplatesignatureAttachmentnamesource**](FieldEEzsigntemplatesignatureAttachmentnamesource.md) |  | [optional] 
**BEzsigntemplatesignatureRequired** | Pointer to **bool** | Whether the Ezsigntemplatesignature is required or not. This field is relevant only with Ezsigntemplatesignature with eEzsigntemplatesignatureType &#x3D; Attachments. | [optional] 
**IEzsigntemplatesignatureMaxlength** | Pointer to **int32** | The maximum length for the value in the Ezsigntemplatesignature  This can only be set if eEzsigntemplatesignatureType is **FieldText** or **FieldTextarea** | [optional] 
**SEzsigntemplatesignatureRegexp** | Pointer to **string** | A regular expression to indicate what values are acceptable for the Ezsigntemplatesignature.  This can only be set if eEzsigntemplatesignatureType is **Text** or **Textarea** | [optional] 
**EEzsigntemplatesignatureTextvalidation** | Pointer to [**EnumTextvalidation**](EnumTextvalidation.md) |  | [optional] 
**EEzsigntemplatesignatureDependencyrequirement** | Pointer to [**FieldEEzsigntemplatesignatureDependencyrequirement**](FieldEEzsigntemplatesignatureDependencyrequirement.md) |  | [optional] 

## Methods

### NewEzsigntemplatesignatureResponse

`func NewEzsigntemplatesignatureResponse(pkiEzsigntemplatesignatureID int32, fkiEzsigntemplatedocumentID int32, fkiEzsigntemplatesignerID int32, iEzsigntemplatedocumentpagePagenumber int32, iEzsigntemplatesignatureX int32, iEzsigntemplatesignatureY int32, iEzsigntemplatesignatureStep int32, eEzsigntemplatesignatureType FieldEEzsigntemplatesignatureType, ) *EzsigntemplatesignatureResponse`

NewEzsigntemplatesignatureResponse instantiates a new EzsigntemplatesignatureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatesignatureResponseWithDefaults

`func NewEzsigntemplatesignatureResponseWithDefaults() *EzsigntemplatesignatureResponse`

NewEzsigntemplatesignatureResponseWithDefaults instantiates a new EzsigntemplatesignatureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatesignatureID

`func (o *EzsigntemplatesignatureResponse) GetPkiEzsigntemplatesignatureID() int32`

GetPkiEzsigntemplatesignatureID returns the PkiEzsigntemplatesignatureID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatesignatureIDOk

`func (o *EzsigntemplatesignatureResponse) GetPkiEzsigntemplatesignatureIDOk() (*int32, bool)`

GetPkiEzsigntemplatesignatureIDOk returns a tuple with the PkiEzsigntemplatesignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatesignatureID

`func (o *EzsigntemplatesignatureResponse) SetPkiEzsigntemplatesignatureID(v int32)`

SetPkiEzsigntemplatesignatureID sets PkiEzsigntemplatesignatureID field to given value.


### GetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplatesignatureResponse) GetFkiEzsigntemplatedocumentID() int32`

GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatedocumentIDOk

`func (o *EzsigntemplatesignatureResponse) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool)`

GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplatesignatureResponse) SetFkiEzsigntemplatedocumentID(v int32)`

SetFkiEzsigntemplatedocumentID sets FkiEzsigntemplatedocumentID field to given value.


### GetFkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignatureResponse) GetFkiEzsigntemplatesignerID() int32`

GetFkiEzsigntemplatesignerID returns the FkiEzsigntemplatesignerID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatesignerIDOk

`func (o *EzsigntemplatesignatureResponse) GetFkiEzsigntemplatesignerIDOk() (*int32, bool)`

GetFkiEzsigntemplatesignerIDOk returns a tuple with the FkiEzsigntemplatesignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatesignerID

`func (o *EzsigntemplatesignatureResponse) SetFkiEzsigntemplatesignerID(v int32)`

SetFkiEzsigntemplatesignerID sets FkiEzsigntemplatesignerID field to given value.


### GetFkiEzsigntemplatesignerIDValidation

`func (o *EzsigntemplatesignatureResponse) GetFkiEzsigntemplatesignerIDValidation() int32`

GetFkiEzsigntemplatesignerIDValidation returns the FkiEzsigntemplatesignerIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatesignerIDValidationOk

`func (o *EzsigntemplatesignatureResponse) GetFkiEzsigntemplatesignerIDValidationOk() (*int32, bool)`

GetFkiEzsigntemplatesignerIDValidationOk returns a tuple with the FkiEzsigntemplatesignerIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatesignerIDValidation

`func (o *EzsigntemplatesignatureResponse) SetFkiEzsigntemplatesignerIDValidation(v int32)`

SetFkiEzsigntemplatesignerIDValidation sets FkiEzsigntemplatesignerIDValidation field to given value.

### HasFkiEzsigntemplatesignerIDValidation

`func (o *EzsigntemplatesignatureResponse) HasFkiEzsigntemplatesignerIDValidation() bool`

HasFkiEzsigntemplatesignerIDValidation returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatedocumentpagePagenumber() int32`

GetIEzsigntemplatedocumentpagePagenumber returns the IEzsigntemplatedocumentpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagePagenumberOk

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatedocumentpagePagenumberOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagePagenumberOk returns a tuple with the IEzsigntemplatedocumentpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatedocumentpagePagenumber(v int32)`

SetIEzsigntemplatedocumentpagePagenumber sets IEzsigntemplatedocumentpagePagenumber field to given value.


### GetIEzsigntemplatesignatureX

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureX() int32`

GetIEzsigntemplatesignatureX returns the IEzsigntemplatesignatureX field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureXOk

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureXOk() (*int32, bool)`

GetIEzsigntemplatesignatureXOk returns a tuple with the IEzsigntemplatesignatureX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureX

`func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatesignatureX(v int32)`

SetIEzsigntemplatesignatureX sets IEzsigntemplatesignatureX field to given value.


### GetIEzsigntemplatesignatureY

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureY() int32`

GetIEzsigntemplatesignatureY returns the IEzsigntemplatesignatureY field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureYOk

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureYOk() (*int32, bool)`

GetIEzsigntemplatesignatureYOk returns a tuple with the IEzsigntemplatesignatureY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureY

`func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatesignatureY(v int32)`

SetIEzsigntemplatesignatureY sets IEzsigntemplatesignatureY field to given value.


### GetIEzsigntemplatesignatureWidth

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureWidth() int32`

GetIEzsigntemplatesignatureWidth returns the IEzsigntemplatesignatureWidth field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureWidthOk

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureWidthOk() (*int32, bool)`

GetIEzsigntemplatesignatureWidthOk returns a tuple with the IEzsigntemplatesignatureWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureWidth

`func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatesignatureWidth(v int32)`

SetIEzsigntemplatesignatureWidth sets IEzsigntemplatesignatureWidth field to given value.

### HasIEzsigntemplatesignatureWidth

`func (o *EzsigntemplatesignatureResponse) HasIEzsigntemplatesignatureWidth() bool`

HasIEzsigntemplatesignatureWidth returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureHeight

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureHeight() int32`

GetIEzsigntemplatesignatureHeight returns the IEzsigntemplatesignatureHeight field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureHeightOk

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureHeightOk() (*int32, bool)`

GetIEzsigntemplatesignatureHeightOk returns a tuple with the IEzsigntemplatesignatureHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureHeight

`func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatesignatureHeight(v int32)`

SetIEzsigntemplatesignatureHeight sets IEzsigntemplatesignatureHeight field to given value.

### HasIEzsigntemplatesignatureHeight

`func (o *EzsigntemplatesignatureResponse) HasIEzsigntemplatesignatureHeight() bool`

HasIEzsigntemplatesignatureHeight returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureStep

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureStep() int32`

GetIEzsigntemplatesignatureStep returns the IEzsigntemplatesignatureStep field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureStepOk

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureStepOk() (*int32, bool)`

GetIEzsigntemplatesignatureStepOk returns a tuple with the IEzsigntemplatesignatureStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureStep

`func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatesignatureStep(v int32)`

SetIEzsigntemplatesignatureStep sets IEzsigntemplatesignatureStep field to given value.


### GetEEzsigntemplatesignatureType

`func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureType() FieldEEzsigntemplatesignatureType`

GetEEzsigntemplatesignatureType returns the EEzsigntemplatesignatureType field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureTypeOk

`func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureTypeOk() (*FieldEEzsigntemplatesignatureType, bool)`

GetEEzsigntemplatesignatureTypeOk returns a tuple with the EEzsigntemplatesignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureType

`func (o *EzsigntemplatesignatureResponse) SetEEzsigntemplatesignatureType(v FieldEEzsigntemplatesignatureType)`

SetEEzsigntemplatesignatureType sets EEzsigntemplatesignatureType field to given value.


### GetTEzsigntemplatesignatureTooltip

`func (o *EzsigntemplatesignatureResponse) GetTEzsigntemplatesignatureTooltip() string`

GetTEzsigntemplatesignatureTooltip returns the TEzsigntemplatesignatureTooltip field if non-nil, zero value otherwise.

### GetTEzsigntemplatesignatureTooltipOk

`func (o *EzsigntemplatesignatureResponse) GetTEzsigntemplatesignatureTooltipOk() (*string, bool)`

GetTEzsigntemplatesignatureTooltipOk returns a tuple with the TEzsigntemplatesignatureTooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplatesignatureTooltip

`func (o *EzsigntemplatesignatureResponse) SetTEzsigntemplatesignatureTooltip(v string)`

SetTEzsigntemplatesignatureTooltip sets TEzsigntemplatesignatureTooltip field to given value.

### HasTEzsigntemplatesignatureTooltip

`func (o *EzsigntemplatesignatureResponse) HasTEzsigntemplatesignatureTooltip() bool`

HasTEzsigntemplatesignatureTooltip returns a boolean if a field has been set.

### GetEEzsigntemplatesignatureTooltipposition

`func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureTooltipposition() FieldEEzsigntemplatesignatureTooltipposition`

GetEEzsigntemplatesignatureTooltipposition returns the EEzsigntemplatesignatureTooltipposition field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureTooltippositionOk

`func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureTooltippositionOk() (*FieldEEzsigntemplatesignatureTooltipposition, bool)`

GetEEzsigntemplatesignatureTooltippositionOk returns a tuple with the EEzsigntemplatesignatureTooltipposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureTooltipposition

`func (o *EzsigntemplatesignatureResponse) SetEEzsigntemplatesignatureTooltipposition(v FieldEEzsigntemplatesignatureTooltipposition)`

SetEEzsigntemplatesignatureTooltipposition sets EEzsigntemplatesignatureTooltipposition field to given value.

### HasEEzsigntemplatesignatureTooltipposition

`func (o *EzsigntemplatesignatureResponse) HasEEzsigntemplatesignatureTooltipposition() bool`

HasEEzsigntemplatesignatureTooltipposition returns a boolean if a field has been set.

### GetEEzsigntemplatesignatureFont

`func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureFont() FieldEEzsigntemplatesignatureFont`

GetEEzsigntemplatesignatureFont returns the EEzsigntemplatesignatureFont field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureFontOk

`func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureFontOk() (*FieldEEzsigntemplatesignatureFont, bool)`

GetEEzsigntemplatesignatureFontOk returns a tuple with the EEzsigntemplatesignatureFont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureFont

`func (o *EzsigntemplatesignatureResponse) SetEEzsigntemplatesignatureFont(v FieldEEzsigntemplatesignatureFont)`

SetEEzsigntemplatesignatureFont sets EEzsigntemplatesignatureFont field to given value.

### HasEEzsigntemplatesignatureFont

`func (o *EzsigntemplatesignatureResponse) HasEEzsigntemplatesignatureFont() bool`

HasEEzsigntemplatesignatureFont returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureValidationstep

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureValidationstep() int32`

GetIEzsigntemplatesignatureValidationstep returns the IEzsigntemplatesignatureValidationstep field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureValidationstepOk

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureValidationstepOk() (*int32, bool)`

GetIEzsigntemplatesignatureValidationstepOk returns a tuple with the IEzsigntemplatesignatureValidationstep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureValidationstep

`func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatesignatureValidationstep(v int32)`

SetIEzsigntemplatesignatureValidationstep sets IEzsigntemplatesignatureValidationstep field to given value.

### HasIEzsigntemplatesignatureValidationstep

`func (o *EzsigntemplatesignatureResponse) HasIEzsigntemplatesignatureValidationstep() bool`

HasIEzsigntemplatesignatureValidationstep returns a boolean if a field has been set.

### GetSEzsigntemplatesignatureAttachmentdescription

`func (o *EzsigntemplatesignatureResponse) GetSEzsigntemplatesignatureAttachmentdescription() string`

GetSEzsigntemplatesignatureAttachmentdescription returns the SEzsigntemplatesignatureAttachmentdescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignatureAttachmentdescriptionOk

`func (o *EzsigntemplatesignatureResponse) GetSEzsigntemplatesignatureAttachmentdescriptionOk() (*string, bool)`

GetSEzsigntemplatesignatureAttachmentdescriptionOk returns a tuple with the SEzsigntemplatesignatureAttachmentdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignatureAttachmentdescription

`func (o *EzsigntemplatesignatureResponse) SetSEzsigntemplatesignatureAttachmentdescription(v string)`

SetSEzsigntemplatesignatureAttachmentdescription sets SEzsigntemplatesignatureAttachmentdescription field to given value.

### HasSEzsigntemplatesignatureAttachmentdescription

`func (o *EzsigntemplatesignatureResponse) HasSEzsigntemplatesignatureAttachmentdescription() bool`

HasSEzsigntemplatesignatureAttachmentdescription returns a boolean if a field has been set.

### GetEEzsigntemplatesignatureAttachmentnamesource

`func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureAttachmentnamesource() FieldEEzsigntemplatesignatureAttachmentnamesource`

GetEEzsigntemplatesignatureAttachmentnamesource returns the EEzsigntemplatesignatureAttachmentnamesource field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureAttachmentnamesourceOk

`func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureAttachmentnamesourceOk() (*FieldEEzsigntemplatesignatureAttachmentnamesource, bool)`

GetEEzsigntemplatesignatureAttachmentnamesourceOk returns a tuple with the EEzsigntemplatesignatureAttachmentnamesource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureAttachmentnamesource

`func (o *EzsigntemplatesignatureResponse) SetEEzsigntemplatesignatureAttachmentnamesource(v FieldEEzsigntemplatesignatureAttachmentnamesource)`

SetEEzsigntemplatesignatureAttachmentnamesource sets EEzsigntemplatesignatureAttachmentnamesource field to given value.

### HasEEzsigntemplatesignatureAttachmentnamesource

`func (o *EzsigntemplatesignatureResponse) HasEEzsigntemplatesignatureAttachmentnamesource() bool`

HasEEzsigntemplatesignatureAttachmentnamesource returns a boolean if a field has been set.

### GetBEzsigntemplatesignatureRequired

`func (o *EzsigntemplatesignatureResponse) GetBEzsigntemplatesignatureRequired() bool`

GetBEzsigntemplatesignatureRequired returns the BEzsigntemplatesignatureRequired field if non-nil, zero value otherwise.

### GetBEzsigntemplatesignatureRequiredOk

`func (o *EzsigntemplatesignatureResponse) GetBEzsigntemplatesignatureRequiredOk() (*bool, bool)`

GetBEzsigntemplatesignatureRequiredOk returns a tuple with the BEzsigntemplatesignatureRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatesignatureRequired

`func (o *EzsigntemplatesignatureResponse) SetBEzsigntemplatesignatureRequired(v bool)`

SetBEzsigntemplatesignatureRequired sets BEzsigntemplatesignatureRequired field to given value.

### HasBEzsigntemplatesignatureRequired

`func (o *EzsigntemplatesignatureResponse) HasBEzsigntemplatesignatureRequired() bool`

HasBEzsigntemplatesignatureRequired returns a boolean if a field has been set.

### GetIEzsigntemplatesignatureMaxlength

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureMaxlength() int32`

GetIEzsigntemplatesignatureMaxlength returns the IEzsigntemplatesignatureMaxlength field if non-nil, zero value otherwise.

### GetIEzsigntemplatesignatureMaxlengthOk

`func (o *EzsigntemplatesignatureResponse) GetIEzsigntemplatesignatureMaxlengthOk() (*int32, bool)`

GetIEzsigntemplatesignatureMaxlengthOk returns a tuple with the IEzsigntemplatesignatureMaxlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatesignatureMaxlength

`func (o *EzsigntemplatesignatureResponse) SetIEzsigntemplatesignatureMaxlength(v int32)`

SetIEzsigntemplatesignatureMaxlength sets IEzsigntemplatesignatureMaxlength field to given value.

### HasIEzsigntemplatesignatureMaxlength

`func (o *EzsigntemplatesignatureResponse) HasIEzsigntemplatesignatureMaxlength() bool`

HasIEzsigntemplatesignatureMaxlength returns a boolean if a field has been set.

### GetSEzsigntemplatesignatureRegexp

`func (o *EzsigntemplatesignatureResponse) GetSEzsigntemplatesignatureRegexp() string`

GetSEzsigntemplatesignatureRegexp returns the SEzsigntemplatesignatureRegexp field if non-nil, zero value otherwise.

### GetSEzsigntemplatesignatureRegexpOk

`func (o *EzsigntemplatesignatureResponse) GetSEzsigntemplatesignatureRegexpOk() (*string, bool)`

GetSEzsigntemplatesignatureRegexpOk returns a tuple with the SEzsigntemplatesignatureRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatesignatureRegexp

`func (o *EzsigntemplatesignatureResponse) SetSEzsigntemplatesignatureRegexp(v string)`

SetSEzsigntemplatesignatureRegexp sets SEzsigntemplatesignatureRegexp field to given value.

### HasSEzsigntemplatesignatureRegexp

`func (o *EzsigntemplatesignatureResponse) HasSEzsigntemplatesignatureRegexp() bool`

HasSEzsigntemplatesignatureRegexp returns a boolean if a field has been set.

### GetEEzsigntemplatesignatureTextvalidation

`func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureTextvalidation() EnumTextvalidation`

GetEEzsigntemplatesignatureTextvalidation returns the EEzsigntemplatesignatureTextvalidation field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureTextvalidationOk

`func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureTextvalidationOk() (*EnumTextvalidation, bool)`

GetEEzsigntemplatesignatureTextvalidationOk returns a tuple with the EEzsigntemplatesignatureTextvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureTextvalidation

`func (o *EzsigntemplatesignatureResponse) SetEEzsigntemplatesignatureTextvalidation(v EnumTextvalidation)`

SetEEzsigntemplatesignatureTextvalidation sets EEzsigntemplatesignatureTextvalidation field to given value.

### HasEEzsigntemplatesignatureTextvalidation

`func (o *EzsigntemplatesignatureResponse) HasEEzsigntemplatesignatureTextvalidation() bool`

HasEEzsigntemplatesignatureTextvalidation returns a boolean if a field has been set.

### GetEEzsigntemplatesignatureDependencyrequirement

`func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureDependencyrequirement() FieldEEzsigntemplatesignatureDependencyrequirement`

GetEEzsigntemplatesignatureDependencyrequirement returns the EEzsigntemplatesignatureDependencyrequirement field if non-nil, zero value otherwise.

### GetEEzsigntemplatesignatureDependencyrequirementOk

`func (o *EzsigntemplatesignatureResponse) GetEEzsigntemplatesignatureDependencyrequirementOk() (*FieldEEzsigntemplatesignatureDependencyrequirement, bool)`

GetEEzsigntemplatesignatureDependencyrequirementOk returns a tuple with the EEzsigntemplatesignatureDependencyrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatesignatureDependencyrequirement

`func (o *EzsigntemplatesignatureResponse) SetEEzsigntemplatesignatureDependencyrequirement(v FieldEEzsigntemplatesignatureDependencyrequirement)`

SetEEzsigntemplatesignatureDependencyrequirement sets EEzsigntemplatesignatureDependencyrequirement field to given value.

### HasEEzsigntemplatesignatureDependencyrequirement

`func (o *EzsigntemplatesignatureResponse) HasEEzsigntemplatesignatureDependencyrequirement() bool`

HasEEzsigntemplatesignatureDependencyrequirement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


