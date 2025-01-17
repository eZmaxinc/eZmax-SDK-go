# CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignsignatureID** | Pointer to **int32** | The unique ID of the Ezsignsignature | [optional] 
**FkiEzsignfoldersignerassociationID** | **int32** | The unique ID of the Ezsignfoldersignerassociation | 
**IEzsignpagePagenumber** | **int32** | The page number in the Ezsigndocument | 
**IEzsignsignatureX** | **int32** | The X coordinate (Horizontal) where to put the Ezsignsignature on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignsignature 2 inches from the left border of the page, you would use \&quot;200\&quot; for the X coordinate. | 
**IEzsignsignatureY** | **int32** | The Y coordinate (Vertical) where to put the Ezsignsignature on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignsignature 3 inches from the top border of the page, you would use \&quot;300\&quot; for the Y coordinate. | 
**IEzsignsignatureWidth** | Pointer to **int32** | The width of the Ezsignsignature.  Size is calculated at 100dpi (dot per inch). So for example, if you want the Ezsignsignature to have a width of 2 inches, you would use \&quot;200\&quot; for the iEzsignsignatureWidth. | [optional] 
**IEzsignsignatureHeight** | Pointer to **int32** | The height of the Ezsignsignature.  Size is calculated at 100dpi (dot per inch). So for example, if you want the Ezsignsignature to have an height of 2 inches, you would use \&quot;200\&quot; for the iEzsignsignatureHeight. | [optional] 
**IEzsignsignatureStep** | **int32** | The step when the Ezsignsigner will be invited to sign | 
**EEzsignsignatureType** | [**FieldEEzsignsignatureType**](FieldEEzsignsignatureType.md) |  | 
**FkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 
**TEzsignsignatureTooltip** | Pointer to **string** | A tooltip that will be presented to Ezsignsigner about the Ezsignsignature | [optional] 
**EEzsignsignatureTooltipposition** | Pointer to [**FieldEEzsignsignatureTooltipposition**](FieldEEzsignsignatureTooltipposition.md) |  | [optional] 
**EEzsignsignatureFont** | Pointer to [**FieldEEzsignsignatureFont**](FieldEEzsignsignatureFont.md) |  | [optional] 
**FkiEzsignfoldersignerassociationIDValidation** | Pointer to **int32** | The unique ID of the Ezsignfoldersignerassociation | [optional] 
**BEzsignsignatureHandwritten** | Pointer to **bool** | Whether the Ezsignsignature must be handwritten or not when eEzsignsignatureType &#x3D; Signature. | [optional] 
**BEzsignsignatureReason** | Pointer to **bool** | Whether the Ezsignsignature must include a reason or not when eEzsignsignatureType &#x3D; Signature. | [optional] 
**BEzsignsignatureRequired** | Pointer to **bool** | Whether the Ezsignsignature is required or not. This field is relevant only with Ezsignsignature with eEzsignsignatureType &#x3D; Attachments, Text or Textarea. | [optional] 
**EEzsignsignatureAttachmentnamesource** | Pointer to [**FieldEEzsignsignatureAttachmentnamesource**](FieldEEzsignsignatureAttachmentnamesource.md) |  | [optional] 
**SEzsignsignatureAttachmentdescription** | Pointer to **string** | The description attached to the attachment name added in Ezsignsignature of eEzsignsignatureType Attachments | [optional] 
**EEzsignsignatureConsultationtrigger** | Pointer to [**FieldEEzsignsignatureConsultationtrigger**](FieldEEzsignsignatureConsultationtrigger.md) |  | [optional] 
**IEzsignsignatureValidationstep** | Pointer to **int32** | The step when the Ezsignsigner will be invited to validate the Ezsignsignature of eEzsignsignatureType Attachments | [optional] 
**IEzsignsignatureMaxlength** | Pointer to **int32** | The maximum length for the value in the Ezsignsignature  This can only be set if eEzsignsignatureType is **FieldText** or **FieldTextarea** | [optional] 
**SEzsignsignatureDefaultvalue** | Pointer to **string** | The default value for the Ezsignsignature  You can use the codes below and they will be replaced at signature time.    | Code | Description | Example | | ------------------------- | ------------ | ------------ | | {sUserFirstname} | The first name of the contact | John | | {sUserLastname} | The last name of the contact | Doe | | {sUserJobtitle} | The job title | Sales Representative | | {sCompany} | Company name | eZmax Solutions Inc. | | {sEmailAddress} | The email address | email@example.com | | {sPhoneE164} | A phone number in E.164 Format | +15149901516 | | {sPhoneE164Cell} | A phone number in E.164 Format | +15149901516 | | [optional] 
**EEzsignsignatureTextvalidation** | Pointer to [**EnumTextvalidation**](EnumTextvalidation.md) |  | [optional] 
**SEzsignsignatureTextvalidationcustommessage** | Pointer to **string** | Description of validation rule. Show by signatory. | [optional] 
**SEzsignsignatureRegexp** | Pointer to **string** | A regular expression to indicate what values are acceptable for the Ezsignsignature.  This can only be set if eEzsignsignatureType is **FieldText** or **FieldTextarea** and eEzsignsignatureTextvalidation is **Custom** | [optional] 
**EEzsignsignatureDependencyrequirement** | Pointer to [**FieldEEzsignsignatureDependencyrequirement**](FieldEEzsignsignatureDependencyrequirement.md) |  | [optional] 
**BEzsignsignatureCustomdate** | Pointer to **bool** | Whether the Ezsignsignature has a custom date format or not. (Only possible when eEzsignsignatureType is **Name** or **Handwritten**) | [optional] 
**AObjEzsignsignaturecustomdate** | Pointer to [**[]EzsignsignaturecustomdateRequestCompound**](EzsignsignaturecustomdateRequestCompound.md) | An array of custom date blocks that will be filled at the time of signature.  Can only be used if bEzsignsignatureCustomdate is true.  Use an empty array if you don&#39;t want to have a date at all. | [optional] 
**AObjEzsignelementdependency** | Pointer to [**[]EzsignelementdependencyRequestCompound**](EzsignelementdependencyRequestCompound.md) |  | [optional] 
**ObjCreateezsignelementspositionedbyword** | [**CustomCreateEzsignelementsPositionedByWordRequest**](CustomCreateEzsignelementsPositionedByWordRequest.md) |  | 

## Methods

### NewCustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest

`func NewCustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest(fkiEzsignfoldersignerassociationID int32, iEzsignpagePagenumber int32, iEzsignsignatureX int32, iEzsignsignatureY int32, iEzsignsignatureStep int32, eEzsignsignatureType FieldEEzsignsignatureType, fkiEzsigndocumentID int32, objCreateezsignelementspositionedbyword CustomCreateEzsignelementsPositionedByWordRequest, ) *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest`

NewCustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest instantiates a new CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzsignsignatureCreateEzsignelementsPositionedByWordRequestWithDefaults

`func NewCustomEzsignsignatureCreateEzsignelementsPositionedByWordRequestWithDefaults() *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest`

NewCustomEzsignsignatureCreateEzsignelementsPositionedByWordRequestWithDefaults instantiates a new CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignsignatureID

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetPkiEzsignsignatureID() int32`

GetPkiEzsignsignatureID returns the PkiEzsignsignatureID field if non-nil, zero value otherwise.

### GetPkiEzsignsignatureIDOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetPkiEzsignsignatureIDOk() (*int32, bool)`

GetPkiEzsignsignatureIDOk returns a tuple with the PkiEzsignsignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignsignatureID

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetPkiEzsignsignatureID(v int32)`

SetPkiEzsignsignatureID sets PkiEzsignsignatureID field to given value.

### HasPkiEzsignsignatureID

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasPkiEzsignsignatureID() bool`

HasPkiEzsignsignatureID returns a boolean if a field has been set.

### GetFkiEzsignfoldersignerassociationID

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetFkiEzsignfoldersignerassociationID() int32`

GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationID

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetFkiEzsignfoldersignerassociationID(v int32)`

SetFkiEzsignfoldersignerassociationID sets FkiEzsignfoldersignerassociationID field to given value.


### GetIEzsignpagePagenumber

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignpagePagenumber() int32`

GetIEzsignpagePagenumber returns the IEzsignpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsignpagePagenumberOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignpagePagenumberOk() (*int32, bool)`

GetIEzsignpagePagenumberOk returns a tuple with the IEzsignpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignpagePagenumber

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetIEzsignpagePagenumber(v int32)`

SetIEzsignpagePagenumber sets IEzsignpagePagenumber field to given value.


### GetIEzsignsignatureX

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignsignatureX() int32`

GetIEzsignsignatureX returns the IEzsignsignatureX field if non-nil, zero value otherwise.

### GetIEzsignsignatureXOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignsignatureXOk() (*int32, bool)`

GetIEzsignsignatureXOk returns a tuple with the IEzsignsignatureX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureX

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetIEzsignsignatureX(v int32)`

SetIEzsignsignatureX sets IEzsignsignatureX field to given value.


### GetIEzsignsignatureY

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignsignatureY() int32`

GetIEzsignsignatureY returns the IEzsignsignatureY field if non-nil, zero value otherwise.

### GetIEzsignsignatureYOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignsignatureYOk() (*int32, bool)`

GetIEzsignsignatureYOk returns a tuple with the IEzsignsignatureY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureY

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetIEzsignsignatureY(v int32)`

SetIEzsignsignatureY sets IEzsignsignatureY field to given value.


### GetIEzsignsignatureWidth

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignsignatureWidth() int32`

GetIEzsignsignatureWidth returns the IEzsignsignatureWidth field if non-nil, zero value otherwise.

### GetIEzsignsignatureWidthOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignsignatureWidthOk() (*int32, bool)`

GetIEzsignsignatureWidthOk returns a tuple with the IEzsignsignatureWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureWidth

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetIEzsignsignatureWidth(v int32)`

SetIEzsignsignatureWidth sets IEzsignsignatureWidth field to given value.

### HasIEzsignsignatureWidth

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasIEzsignsignatureWidth() bool`

HasIEzsignsignatureWidth returns a boolean if a field has been set.

### GetIEzsignsignatureHeight

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignsignatureHeight() int32`

GetIEzsignsignatureHeight returns the IEzsignsignatureHeight field if non-nil, zero value otherwise.

### GetIEzsignsignatureHeightOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignsignatureHeightOk() (*int32, bool)`

GetIEzsignsignatureHeightOk returns a tuple with the IEzsignsignatureHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureHeight

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetIEzsignsignatureHeight(v int32)`

SetIEzsignsignatureHeight sets IEzsignsignatureHeight field to given value.

### HasIEzsignsignatureHeight

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasIEzsignsignatureHeight() bool`

HasIEzsignsignatureHeight returns a boolean if a field has been set.

### GetIEzsignsignatureStep

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignsignatureStep() int32`

GetIEzsignsignatureStep returns the IEzsignsignatureStep field if non-nil, zero value otherwise.

### GetIEzsignsignatureStepOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignsignatureStepOk() (*int32, bool)`

GetIEzsignsignatureStepOk returns a tuple with the IEzsignsignatureStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureStep

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetIEzsignsignatureStep(v int32)`

SetIEzsignsignatureStep sets IEzsignsignatureStep field to given value.


### GetEEzsignsignatureType

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetEEzsignsignatureType() FieldEEzsignsignatureType`

GetEEzsignsignatureType returns the EEzsignsignatureType field if non-nil, zero value otherwise.

### GetEEzsignsignatureTypeOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetEEzsignsignatureTypeOk() (*FieldEEzsignsignatureType, bool)`

GetEEzsignsignatureTypeOk returns a tuple with the EEzsignsignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureType

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetEEzsignsignatureType(v FieldEEzsignsignatureType)`

SetEEzsignsignatureType sets EEzsignsignatureType field to given value.


### GetFkiEzsigndocumentID

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.


### GetTEzsignsignatureTooltip

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetTEzsignsignatureTooltip() string`

GetTEzsignsignatureTooltip returns the TEzsignsignatureTooltip field if non-nil, zero value otherwise.

### GetTEzsignsignatureTooltipOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetTEzsignsignatureTooltipOk() (*string, bool)`

GetTEzsignsignatureTooltipOk returns a tuple with the TEzsignsignatureTooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignsignatureTooltip

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetTEzsignsignatureTooltip(v string)`

SetTEzsignsignatureTooltip sets TEzsignsignatureTooltip field to given value.

### HasTEzsignsignatureTooltip

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasTEzsignsignatureTooltip() bool`

HasTEzsignsignatureTooltip returns a boolean if a field has been set.

### GetEEzsignsignatureTooltipposition

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetEEzsignsignatureTooltipposition() FieldEEzsignsignatureTooltipposition`

GetEEzsignsignatureTooltipposition returns the EEzsignsignatureTooltipposition field if non-nil, zero value otherwise.

### GetEEzsignsignatureTooltippositionOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetEEzsignsignatureTooltippositionOk() (*FieldEEzsignsignatureTooltipposition, bool)`

GetEEzsignsignatureTooltippositionOk returns a tuple with the EEzsignsignatureTooltipposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureTooltipposition

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetEEzsignsignatureTooltipposition(v FieldEEzsignsignatureTooltipposition)`

SetEEzsignsignatureTooltipposition sets EEzsignsignatureTooltipposition field to given value.

### HasEEzsignsignatureTooltipposition

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasEEzsignsignatureTooltipposition() bool`

HasEEzsignsignatureTooltipposition returns a boolean if a field has been set.

### GetEEzsignsignatureFont

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetEEzsignsignatureFont() FieldEEzsignsignatureFont`

GetEEzsignsignatureFont returns the EEzsignsignatureFont field if non-nil, zero value otherwise.

### GetEEzsignsignatureFontOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetEEzsignsignatureFontOk() (*FieldEEzsignsignatureFont, bool)`

GetEEzsignsignatureFontOk returns a tuple with the EEzsignsignatureFont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureFont

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetEEzsignsignatureFont(v FieldEEzsignsignatureFont)`

SetEEzsignsignatureFont sets EEzsignsignatureFont field to given value.

### HasEEzsignsignatureFont

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasEEzsignsignatureFont() bool`

HasEEzsignsignatureFont returns a boolean if a field has been set.

### GetFkiEzsignfoldersignerassociationIDValidation

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetFkiEzsignfoldersignerassociationIDValidation() int32`

GetFkiEzsignfoldersignerassociationIDValidation returns the FkiEzsignfoldersignerassociationIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDValidationOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetFkiEzsignfoldersignerassociationIDValidationOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDValidationOk returns a tuple with the FkiEzsignfoldersignerassociationIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationIDValidation

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetFkiEzsignfoldersignerassociationIDValidation(v int32)`

SetFkiEzsignfoldersignerassociationIDValidation sets FkiEzsignfoldersignerassociationIDValidation field to given value.

### HasFkiEzsignfoldersignerassociationIDValidation

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasFkiEzsignfoldersignerassociationIDValidation() bool`

HasFkiEzsignfoldersignerassociationIDValidation returns a boolean if a field has been set.

### GetBEzsignsignatureHandwritten

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetBEzsignsignatureHandwritten() bool`

GetBEzsignsignatureHandwritten returns the BEzsignsignatureHandwritten field if non-nil, zero value otherwise.

### GetBEzsignsignatureHandwrittenOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetBEzsignsignatureHandwrittenOk() (*bool, bool)`

GetBEzsignsignatureHandwrittenOk returns a tuple with the BEzsignsignatureHandwritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureHandwritten

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetBEzsignsignatureHandwritten(v bool)`

SetBEzsignsignatureHandwritten sets BEzsignsignatureHandwritten field to given value.

### HasBEzsignsignatureHandwritten

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasBEzsignsignatureHandwritten() bool`

HasBEzsignsignatureHandwritten returns a boolean if a field has been set.

### GetBEzsignsignatureReason

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetBEzsignsignatureReason() bool`

GetBEzsignsignatureReason returns the BEzsignsignatureReason field if non-nil, zero value otherwise.

### GetBEzsignsignatureReasonOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetBEzsignsignatureReasonOk() (*bool, bool)`

GetBEzsignsignatureReasonOk returns a tuple with the BEzsignsignatureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureReason

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetBEzsignsignatureReason(v bool)`

SetBEzsignsignatureReason sets BEzsignsignatureReason field to given value.

### HasBEzsignsignatureReason

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasBEzsignsignatureReason() bool`

HasBEzsignsignatureReason returns a boolean if a field has been set.

### GetBEzsignsignatureRequired

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetBEzsignsignatureRequired() bool`

GetBEzsignsignatureRequired returns the BEzsignsignatureRequired field if non-nil, zero value otherwise.

### GetBEzsignsignatureRequiredOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetBEzsignsignatureRequiredOk() (*bool, bool)`

GetBEzsignsignatureRequiredOk returns a tuple with the BEzsignsignatureRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureRequired

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetBEzsignsignatureRequired(v bool)`

SetBEzsignsignatureRequired sets BEzsignsignatureRequired field to given value.

### HasBEzsignsignatureRequired

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasBEzsignsignatureRequired() bool`

HasBEzsignsignatureRequired returns a boolean if a field has been set.

### GetEEzsignsignatureAttachmentnamesource

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetEEzsignsignatureAttachmentnamesource() FieldEEzsignsignatureAttachmentnamesource`

GetEEzsignsignatureAttachmentnamesource returns the EEzsignsignatureAttachmentnamesource field if non-nil, zero value otherwise.

### GetEEzsignsignatureAttachmentnamesourceOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetEEzsignsignatureAttachmentnamesourceOk() (*FieldEEzsignsignatureAttachmentnamesource, bool)`

GetEEzsignsignatureAttachmentnamesourceOk returns a tuple with the EEzsignsignatureAttachmentnamesource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureAttachmentnamesource

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetEEzsignsignatureAttachmentnamesource(v FieldEEzsignsignatureAttachmentnamesource)`

SetEEzsignsignatureAttachmentnamesource sets EEzsignsignatureAttachmentnamesource field to given value.

### HasEEzsignsignatureAttachmentnamesource

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasEEzsignsignatureAttachmentnamesource() bool`

HasEEzsignsignatureAttachmentnamesource returns a boolean if a field has been set.

### GetSEzsignsignatureAttachmentdescription

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetSEzsignsignatureAttachmentdescription() string`

GetSEzsignsignatureAttachmentdescription returns the SEzsignsignatureAttachmentdescription field if non-nil, zero value otherwise.

### GetSEzsignsignatureAttachmentdescriptionOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetSEzsignsignatureAttachmentdescriptionOk() (*string, bool)`

GetSEzsignsignatureAttachmentdescriptionOk returns a tuple with the SEzsignsignatureAttachmentdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureAttachmentdescription

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetSEzsignsignatureAttachmentdescription(v string)`

SetSEzsignsignatureAttachmentdescription sets SEzsignsignatureAttachmentdescription field to given value.

### HasSEzsignsignatureAttachmentdescription

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasSEzsignsignatureAttachmentdescription() bool`

HasSEzsignsignatureAttachmentdescription returns a boolean if a field has been set.

### GetEEzsignsignatureConsultationtrigger

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetEEzsignsignatureConsultationtrigger() FieldEEzsignsignatureConsultationtrigger`

GetEEzsignsignatureConsultationtrigger returns the EEzsignsignatureConsultationtrigger field if non-nil, zero value otherwise.

### GetEEzsignsignatureConsultationtriggerOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetEEzsignsignatureConsultationtriggerOk() (*FieldEEzsignsignatureConsultationtrigger, bool)`

GetEEzsignsignatureConsultationtriggerOk returns a tuple with the EEzsignsignatureConsultationtrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureConsultationtrigger

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetEEzsignsignatureConsultationtrigger(v FieldEEzsignsignatureConsultationtrigger)`

SetEEzsignsignatureConsultationtrigger sets EEzsignsignatureConsultationtrigger field to given value.

### HasEEzsignsignatureConsultationtrigger

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasEEzsignsignatureConsultationtrigger() bool`

HasEEzsignsignatureConsultationtrigger returns a boolean if a field has been set.

### GetIEzsignsignatureValidationstep

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignsignatureValidationstep() int32`

GetIEzsignsignatureValidationstep returns the IEzsignsignatureValidationstep field if non-nil, zero value otherwise.

### GetIEzsignsignatureValidationstepOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignsignatureValidationstepOk() (*int32, bool)`

GetIEzsignsignatureValidationstepOk returns a tuple with the IEzsignsignatureValidationstep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureValidationstep

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetIEzsignsignatureValidationstep(v int32)`

SetIEzsignsignatureValidationstep sets IEzsignsignatureValidationstep field to given value.

### HasIEzsignsignatureValidationstep

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasIEzsignsignatureValidationstep() bool`

HasIEzsignsignatureValidationstep returns a boolean if a field has been set.

### GetIEzsignsignatureMaxlength

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignsignatureMaxlength() int32`

GetIEzsignsignatureMaxlength returns the IEzsignsignatureMaxlength field if non-nil, zero value otherwise.

### GetIEzsignsignatureMaxlengthOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetIEzsignsignatureMaxlengthOk() (*int32, bool)`

GetIEzsignsignatureMaxlengthOk returns a tuple with the IEzsignsignatureMaxlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureMaxlength

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetIEzsignsignatureMaxlength(v int32)`

SetIEzsignsignatureMaxlength sets IEzsignsignatureMaxlength field to given value.

### HasIEzsignsignatureMaxlength

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasIEzsignsignatureMaxlength() bool`

HasIEzsignsignatureMaxlength returns a boolean if a field has been set.

### GetSEzsignsignatureDefaultvalue

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetSEzsignsignatureDefaultvalue() string`

GetSEzsignsignatureDefaultvalue returns the SEzsignsignatureDefaultvalue field if non-nil, zero value otherwise.

### GetSEzsignsignatureDefaultvalueOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetSEzsignsignatureDefaultvalueOk() (*string, bool)`

GetSEzsignsignatureDefaultvalueOk returns a tuple with the SEzsignsignatureDefaultvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureDefaultvalue

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetSEzsignsignatureDefaultvalue(v string)`

SetSEzsignsignatureDefaultvalue sets SEzsignsignatureDefaultvalue field to given value.

### HasSEzsignsignatureDefaultvalue

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasSEzsignsignatureDefaultvalue() bool`

HasSEzsignsignatureDefaultvalue returns a boolean if a field has been set.

### GetEEzsignsignatureTextvalidation

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetEEzsignsignatureTextvalidation() EnumTextvalidation`

GetEEzsignsignatureTextvalidation returns the EEzsignsignatureTextvalidation field if non-nil, zero value otherwise.

### GetEEzsignsignatureTextvalidationOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetEEzsignsignatureTextvalidationOk() (*EnumTextvalidation, bool)`

GetEEzsignsignatureTextvalidationOk returns a tuple with the EEzsignsignatureTextvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureTextvalidation

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetEEzsignsignatureTextvalidation(v EnumTextvalidation)`

SetEEzsignsignatureTextvalidation sets EEzsignsignatureTextvalidation field to given value.

### HasEEzsignsignatureTextvalidation

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasEEzsignsignatureTextvalidation() bool`

HasEEzsignsignatureTextvalidation returns a boolean if a field has been set.

### GetSEzsignsignatureTextvalidationcustommessage

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetSEzsignsignatureTextvalidationcustommessage() string`

GetSEzsignsignatureTextvalidationcustommessage returns the SEzsignsignatureTextvalidationcustommessage field if non-nil, zero value otherwise.

### GetSEzsignsignatureTextvalidationcustommessageOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetSEzsignsignatureTextvalidationcustommessageOk() (*string, bool)`

GetSEzsignsignatureTextvalidationcustommessageOk returns a tuple with the SEzsignsignatureTextvalidationcustommessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureTextvalidationcustommessage

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetSEzsignsignatureTextvalidationcustommessage(v string)`

SetSEzsignsignatureTextvalidationcustommessage sets SEzsignsignatureTextvalidationcustommessage field to given value.

### HasSEzsignsignatureTextvalidationcustommessage

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasSEzsignsignatureTextvalidationcustommessage() bool`

HasSEzsignsignatureTextvalidationcustommessage returns a boolean if a field has been set.

### GetSEzsignsignatureRegexp

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetSEzsignsignatureRegexp() string`

GetSEzsignsignatureRegexp returns the SEzsignsignatureRegexp field if non-nil, zero value otherwise.

### GetSEzsignsignatureRegexpOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetSEzsignsignatureRegexpOk() (*string, bool)`

GetSEzsignsignatureRegexpOk returns a tuple with the SEzsignsignatureRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureRegexp

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetSEzsignsignatureRegexp(v string)`

SetSEzsignsignatureRegexp sets SEzsignsignatureRegexp field to given value.

### HasSEzsignsignatureRegexp

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasSEzsignsignatureRegexp() bool`

HasSEzsignsignatureRegexp returns a boolean if a field has been set.

### GetEEzsignsignatureDependencyrequirement

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetEEzsignsignatureDependencyrequirement() FieldEEzsignsignatureDependencyrequirement`

GetEEzsignsignatureDependencyrequirement returns the EEzsignsignatureDependencyrequirement field if non-nil, zero value otherwise.

### GetEEzsignsignatureDependencyrequirementOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetEEzsignsignatureDependencyrequirementOk() (*FieldEEzsignsignatureDependencyrequirement, bool)`

GetEEzsignsignatureDependencyrequirementOk returns a tuple with the EEzsignsignatureDependencyrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureDependencyrequirement

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetEEzsignsignatureDependencyrequirement(v FieldEEzsignsignatureDependencyrequirement)`

SetEEzsignsignatureDependencyrequirement sets EEzsignsignatureDependencyrequirement field to given value.

### HasEEzsignsignatureDependencyrequirement

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasEEzsignsignatureDependencyrequirement() bool`

HasEEzsignsignatureDependencyrequirement returns a boolean if a field has been set.

### GetBEzsignsignatureCustomdate

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetBEzsignsignatureCustomdate() bool`

GetBEzsignsignatureCustomdate returns the BEzsignsignatureCustomdate field if non-nil, zero value otherwise.

### GetBEzsignsignatureCustomdateOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetBEzsignsignatureCustomdateOk() (*bool, bool)`

GetBEzsignsignatureCustomdateOk returns a tuple with the BEzsignsignatureCustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureCustomdate

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetBEzsignsignatureCustomdate(v bool)`

SetBEzsignsignatureCustomdate sets BEzsignsignatureCustomdate field to given value.

### HasBEzsignsignatureCustomdate

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasBEzsignsignatureCustomdate() bool`

HasBEzsignsignatureCustomdate returns a boolean if a field has been set.

### GetAObjEzsignsignaturecustomdate

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetAObjEzsignsignaturecustomdate() []EzsignsignaturecustomdateRequestCompound`

GetAObjEzsignsignaturecustomdate returns the AObjEzsignsignaturecustomdate field if non-nil, zero value otherwise.

### GetAObjEzsignsignaturecustomdateOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetAObjEzsignsignaturecustomdateOk() (*[]EzsignsignaturecustomdateRequestCompound, bool)`

GetAObjEzsignsignaturecustomdateOk returns a tuple with the AObjEzsignsignaturecustomdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignsignaturecustomdate

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetAObjEzsignsignaturecustomdate(v []EzsignsignaturecustomdateRequestCompound)`

SetAObjEzsignsignaturecustomdate sets AObjEzsignsignaturecustomdate field to given value.

### HasAObjEzsignsignaturecustomdate

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasAObjEzsignsignaturecustomdate() bool`

HasAObjEzsignsignaturecustomdate returns a boolean if a field has been set.

### GetAObjEzsignelementdependency

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetAObjEzsignelementdependency() []EzsignelementdependencyRequestCompound`

GetAObjEzsignelementdependency returns the AObjEzsignelementdependency field if non-nil, zero value otherwise.

### GetAObjEzsignelementdependencyOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetAObjEzsignelementdependencyOk() (*[]EzsignelementdependencyRequestCompound, bool)`

GetAObjEzsignelementdependencyOk returns a tuple with the AObjEzsignelementdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignelementdependency

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetAObjEzsignelementdependency(v []EzsignelementdependencyRequestCompound)`

SetAObjEzsignelementdependency sets AObjEzsignelementdependency field to given value.

### HasAObjEzsignelementdependency

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) HasAObjEzsignelementdependency() bool`

HasAObjEzsignelementdependency returns a boolean if a field has been set.

### GetObjCreateezsignelementspositionedbyword

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetObjCreateezsignelementspositionedbyword() CustomCreateEzsignelementsPositionedByWordRequest`

GetObjCreateezsignelementspositionedbyword returns the ObjCreateezsignelementspositionedbyword field if non-nil, zero value otherwise.

### GetObjCreateezsignelementspositionedbywordOk

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) GetObjCreateezsignelementspositionedbywordOk() (*CustomCreateEzsignelementsPositionedByWordRequest, bool)`

GetObjCreateezsignelementspositionedbywordOk returns a tuple with the ObjCreateezsignelementspositionedbyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCreateezsignelementspositionedbyword

`func (o *CustomEzsignsignatureCreateEzsignelementsPositionedByWordRequest) SetObjCreateezsignelementspositionedbyword(v CustomCreateEzsignelementsPositionedByWordRequest)`

SetObjCreateezsignelementspositionedbyword sets ObjCreateezsignelementspositionedbyword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


