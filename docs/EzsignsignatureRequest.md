# EzsignsignatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignsignatureID** | Pointer to **int32** | The unique ID of the Ezsignsignature | [optional] 
**FkiEzsignfoldersignerassociationID** | **int32** | The unique ID of the Ezsignfoldersignerassociation | 
**FkiPaymentgatewayID** | Pointer to **int32** | The unique ID of the Paymentgateway | [optional] 
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

## Methods

### NewEzsignsignatureRequest

`func NewEzsignsignatureRequest(fkiEzsignfoldersignerassociationID int32, iEzsignpagePagenumber int32, iEzsignsignatureX int32, iEzsignsignatureY int32, iEzsignsignatureStep int32, eEzsignsignatureType FieldEEzsignsignatureType, fkiEzsigndocumentID int32, ) *EzsignsignatureRequest`

NewEzsignsignatureRequest instantiates a new EzsignsignatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignatureRequestWithDefaults

`func NewEzsignsignatureRequestWithDefaults() *EzsignsignatureRequest`

NewEzsignsignatureRequestWithDefaults instantiates a new EzsignsignatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignsignatureID

`func (o *EzsignsignatureRequest) GetPkiEzsignsignatureID() int32`

GetPkiEzsignsignatureID returns the PkiEzsignsignatureID field if non-nil, zero value otherwise.

### GetPkiEzsignsignatureIDOk

`func (o *EzsignsignatureRequest) GetPkiEzsignsignatureIDOk() (*int32, bool)`

GetPkiEzsignsignatureIDOk returns a tuple with the PkiEzsignsignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignsignatureID

`func (o *EzsignsignatureRequest) SetPkiEzsignsignatureID(v int32)`

SetPkiEzsignsignatureID sets PkiEzsignsignatureID field to given value.

### HasPkiEzsignsignatureID

`func (o *EzsignsignatureRequest) HasPkiEzsignsignatureID() bool`

HasPkiEzsignsignatureID returns a boolean if a field has been set.

### GetFkiEzsignfoldersignerassociationID

`func (o *EzsignsignatureRequest) GetFkiEzsignfoldersignerassociationID() int32`

GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDOk

`func (o *EzsignsignatureRequest) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationID

`func (o *EzsignsignatureRequest) SetFkiEzsignfoldersignerassociationID(v int32)`

SetFkiEzsignfoldersignerassociationID sets FkiEzsignfoldersignerassociationID field to given value.


### GetFkiPaymentgatewayID

`func (o *EzsignsignatureRequest) GetFkiPaymentgatewayID() int32`

GetFkiPaymentgatewayID returns the FkiPaymentgatewayID field if non-nil, zero value otherwise.

### GetFkiPaymentgatewayIDOk

`func (o *EzsignsignatureRequest) GetFkiPaymentgatewayIDOk() (*int32, bool)`

GetFkiPaymentgatewayIDOk returns a tuple with the FkiPaymentgatewayID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPaymentgatewayID

`func (o *EzsignsignatureRequest) SetFkiPaymentgatewayID(v int32)`

SetFkiPaymentgatewayID sets FkiPaymentgatewayID field to given value.

### HasFkiPaymentgatewayID

`func (o *EzsignsignatureRequest) HasFkiPaymentgatewayID() bool`

HasFkiPaymentgatewayID returns a boolean if a field has been set.

### GetIEzsignpagePagenumber

`func (o *EzsignsignatureRequest) GetIEzsignpagePagenumber() int32`

GetIEzsignpagePagenumber returns the IEzsignpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsignpagePagenumberOk

`func (o *EzsignsignatureRequest) GetIEzsignpagePagenumberOk() (*int32, bool)`

GetIEzsignpagePagenumberOk returns a tuple with the IEzsignpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignpagePagenumber

`func (o *EzsignsignatureRequest) SetIEzsignpagePagenumber(v int32)`

SetIEzsignpagePagenumber sets IEzsignpagePagenumber field to given value.


### GetIEzsignsignatureX

`func (o *EzsignsignatureRequest) GetIEzsignsignatureX() int32`

GetIEzsignsignatureX returns the IEzsignsignatureX field if non-nil, zero value otherwise.

### GetIEzsignsignatureXOk

`func (o *EzsignsignatureRequest) GetIEzsignsignatureXOk() (*int32, bool)`

GetIEzsignsignatureXOk returns a tuple with the IEzsignsignatureX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureX

`func (o *EzsignsignatureRequest) SetIEzsignsignatureX(v int32)`

SetIEzsignsignatureX sets IEzsignsignatureX field to given value.


### GetIEzsignsignatureY

`func (o *EzsignsignatureRequest) GetIEzsignsignatureY() int32`

GetIEzsignsignatureY returns the IEzsignsignatureY field if non-nil, zero value otherwise.

### GetIEzsignsignatureYOk

`func (o *EzsignsignatureRequest) GetIEzsignsignatureYOk() (*int32, bool)`

GetIEzsignsignatureYOk returns a tuple with the IEzsignsignatureY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureY

`func (o *EzsignsignatureRequest) SetIEzsignsignatureY(v int32)`

SetIEzsignsignatureY sets IEzsignsignatureY field to given value.


### GetIEzsignsignatureWidth

`func (o *EzsignsignatureRequest) GetIEzsignsignatureWidth() int32`

GetIEzsignsignatureWidth returns the IEzsignsignatureWidth field if non-nil, zero value otherwise.

### GetIEzsignsignatureWidthOk

`func (o *EzsignsignatureRequest) GetIEzsignsignatureWidthOk() (*int32, bool)`

GetIEzsignsignatureWidthOk returns a tuple with the IEzsignsignatureWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureWidth

`func (o *EzsignsignatureRequest) SetIEzsignsignatureWidth(v int32)`

SetIEzsignsignatureWidth sets IEzsignsignatureWidth field to given value.

### HasIEzsignsignatureWidth

`func (o *EzsignsignatureRequest) HasIEzsignsignatureWidth() bool`

HasIEzsignsignatureWidth returns a boolean if a field has been set.

### GetIEzsignsignatureHeight

`func (o *EzsignsignatureRequest) GetIEzsignsignatureHeight() int32`

GetIEzsignsignatureHeight returns the IEzsignsignatureHeight field if non-nil, zero value otherwise.

### GetIEzsignsignatureHeightOk

`func (o *EzsignsignatureRequest) GetIEzsignsignatureHeightOk() (*int32, bool)`

GetIEzsignsignatureHeightOk returns a tuple with the IEzsignsignatureHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureHeight

`func (o *EzsignsignatureRequest) SetIEzsignsignatureHeight(v int32)`

SetIEzsignsignatureHeight sets IEzsignsignatureHeight field to given value.

### HasIEzsignsignatureHeight

`func (o *EzsignsignatureRequest) HasIEzsignsignatureHeight() bool`

HasIEzsignsignatureHeight returns a boolean if a field has been set.

### GetIEzsignsignatureStep

`func (o *EzsignsignatureRequest) GetIEzsignsignatureStep() int32`

GetIEzsignsignatureStep returns the IEzsignsignatureStep field if non-nil, zero value otherwise.

### GetIEzsignsignatureStepOk

`func (o *EzsignsignatureRequest) GetIEzsignsignatureStepOk() (*int32, bool)`

GetIEzsignsignatureStepOk returns a tuple with the IEzsignsignatureStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureStep

`func (o *EzsignsignatureRequest) SetIEzsignsignatureStep(v int32)`

SetIEzsignsignatureStep sets IEzsignsignatureStep field to given value.


### GetEEzsignsignatureType

`func (o *EzsignsignatureRequest) GetEEzsignsignatureType() FieldEEzsignsignatureType`

GetEEzsignsignatureType returns the EEzsignsignatureType field if non-nil, zero value otherwise.

### GetEEzsignsignatureTypeOk

`func (o *EzsignsignatureRequest) GetEEzsignsignatureTypeOk() (*FieldEEzsignsignatureType, bool)`

GetEEzsignsignatureTypeOk returns a tuple with the EEzsignsignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureType

`func (o *EzsignsignatureRequest) SetEEzsignsignatureType(v FieldEEzsignsignatureType)`

SetEEzsignsignatureType sets EEzsignsignatureType field to given value.


### GetFkiEzsigndocumentID

`func (o *EzsignsignatureRequest) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *EzsignsignatureRequest) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *EzsignsignatureRequest) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.


### GetTEzsignsignatureTooltip

`func (o *EzsignsignatureRequest) GetTEzsignsignatureTooltip() string`

GetTEzsignsignatureTooltip returns the TEzsignsignatureTooltip field if non-nil, zero value otherwise.

### GetTEzsignsignatureTooltipOk

`func (o *EzsignsignatureRequest) GetTEzsignsignatureTooltipOk() (*string, bool)`

GetTEzsignsignatureTooltipOk returns a tuple with the TEzsignsignatureTooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignsignatureTooltip

`func (o *EzsignsignatureRequest) SetTEzsignsignatureTooltip(v string)`

SetTEzsignsignatureTooltip sets TEzsignsignatureTooltip field to given value.

### HasTEzsignsignatureTooltip

`func (o *EzsignsignatureRequest) HasTEzsignsignatureTooltip() bool`

HasTEzsignsignatureTooltip returns a boolean if a field has been set.

### GetEEzsignsignatureTooltipposition

`func (o *EzsignsignatureRequest) GetEEzsignsignatureTooltipposition() FieldEEzsignsignatureTooltipposition`

GetEEzsignsignatureTooltipposition returns the EEzsignsignatureTooltipposition field if non-nil, zero value otherwise.

### GetEEzsignsignatureTooltippositionOk

`func (o *EzsignsignatureRequest) GetEEzsignsignatureTooltippositionOk() (*FieldEEzsignsignatureTooltipposition, bool)`

GetEEzsignsignatureTooltippositionOk returns a tuple with the EEzsignsignatureTooltipposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureTooltipposition

`func (o *EzsignsignatureRequest) SetEEzsignsignatureTooltipposition(v FieldEEzsignsignatureTooltipposition)`

SetEEzsignsignatureTooltipposition sets EEzsignsignatureTooltipposition field to given value.

### HasEEzsignsignatureTooltipposition

`func (o *EzsignsignatureRequest) HasEEzsignsignatureTooltipposition() bool`

HasEEzsignsignatureTooltipposition returns a boolean if a field has been set.

### GetEEzsignsignatureFont

`func (o *EzsignsignatureRequest) GetEEzsignsignatureFont() FieldEEzsignsignatureFont`

GetEEzsignsignatureFont returns the EEzsignsignatureFont field if non-nil, zero value otherwise.

### GetEEzsignsignatureFontOk

`func (o *EzsignsignatureRequest) GetEEzsignsignatureFontOk() (*FieldEEzsignsignatureFont, bool)`

GetEEzsignsignatureFontOk returns a tuple with the EEzsignsignatureFont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureFont

`func (o *EzsignsignatureRequest) SetEEzsignsignatureFont(v FieldEEzsignsignatureFont)`

SetEEzsignsignatureFont sets EEzsignsignatureFont field to given value.

### HasEEzsignsignatureFont

`func (o *EzsignsignatureRequest) HasEEzsignsignatureFont() bool`

HasEEzsignsignatureFont returns a boolean if a field has been set.

### GetFkiEzsignfoldersignerassociationIDValidation

`func (o *EzsignsignatureRequest) GetFkiEzsignfoldersignerassociationIDValidation() int32`

GetFkiEzsignfoldersignerassociationIDValidation returns the FkiEzsignfoldersignerassociationIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDValidationOk

`func (o *EzsignsignatureRequest) GetFkiEzsignfoldersignerassociationIDValidationOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDValidationOk returns a tuple with the FkiEzsignfoldersignerassociationIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationIDValidation

`func (o *EzsignsignatureRequest) SetFkiEzsignfoldersignerassociationIDValidation(v int32)`

SetFkiEzsignfoldersignerassociationIDValidation sets FkiEzsignfoldersignerassociationIDValidation field to given value.

### HasFkiEzsignfoldersignerassociationIDValidation

`func (o *EzsignsignatureRequest) HasFkiEzsignfoldersignerassociationIDValidation() bool`

HasFkiEzsignfoldersignerassociationIDValidation returns a boolean if a field has been set.

### GetBEzsignsignatureHandwritten

`func (o *EzsignsignatureRequest) GetBEzsignsignatureHandwritten() bool`

GetBEzsignsignatureHandwritten returns the BEzsignsignatureHandwritten field if non-nil, zero value otherwise.

### GetBEzsignsignatureHandwrittenOk

`func (o *EzsignsignatureRequest) GetBEzsignsignatureHandwrittenOk() (*bool, bool)`

GetBEzsignsignatureHandwrittenOk returns a tuple with the BEzsignsignatureHandwritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureHandwritten

`func (o *EzsignsignatureRequest) SetBEzsignsignatureHandwritten(v bool)`

SetBEzsignsignatureHandwritten sets BEzsignsignatureHandwritten field to given value.

### HasBEzsignsignatureHandwritten

`func (o *EzsignsignatureRequest) HasBEzsignsignatureHandwritten() bool`

HasBEzsignsignatureHandwritten returns a boolean if a field has been set.

### GetBEzsignsignatureReason

`func (o *EzsignsignatureRequest) GetBEzsignsignatureReason() bool`

GetBEzsignsignatureReason returns the BEzsignsignatureReason field if non-nil, zero value otherwise.

### GetBEzsignsignatureReasonOk

`func (o *EzsignsignatureRequest) GetBEzsignsignatureReasonOk() (*bool, bool)`

GetBEzsignsignatureReasonOk returns a tuple with the BEzsignsignatureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureReason

`func (o *EzsignsignatureRequest) SetBEzsignsignatureReason(v bool)`

SetBEzsignsignatureReason sets BEzsignsignatureReason field to given value.

### HasBEzsignsignatureReason

`func (o *EzsignsignatureRequest) HasBEzsignsignatureReason() bool`

HasBEzsignsignatureReason returns a boolean if a field has been set.

### GetBEzsignsignatureRequired

`func (o *EzsignsignatureRequest) GetBEzsignsignatureRequired() bool`

GetBEzsignsignatureRequired returns the BEzsignsignatureRequired field if non-nil, zero value otherwise.

### GetBEzsignsignatureRequiredOk

`func (o *EzsignsignatureRequest) GetBEzsignsignatureRequiredOk() (*bool, bool)`

GetBEzsignsignatureRequiredOk returns a tuple with the BEzsignsignatureRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureRequired

`func (o *EzsignsignatureRequest) SetBEzsignsignatureRequired(v bool)`

SetBEzsignsignatureRequired sets BEzsignsignatureRequired field to given value.

### HasBEzsignsignatureRequired

`func (o *EzsignsignatureRequest) HasBEzsignsignatureRequired() bool`

HasBEzsignsignatureRequired returns a boolean if a field has been set.

### GetEEzsignsignatureAttachmentnamesource

`func (o *EzsignsignatureRequest) GetEEzsignsignatureAttachmentnamesource() FieldEEzsignsignatureAttachmentnamesource`

GetEEzsignsignatureAttachmentnamesource returns the EEzsignsignatureAttachmentnamesource field if non-nil, zero value otherwise.

### GetEEzsignsignatureAttachmentnamesourceOk

`func (o *EzsignsignatureRequest) GetEEzsignsignatureAttachmentnamesourceOk() (*FieldEEzsignsignatureAttachmentnamesource, bool)`

GetEEzsignsignatureAttachmentnamesourceOk returns a tuple with the EEzsignsignatureAttachmentnamesource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureAttachmentnamesource

`func (o *EzsignsignatureRequest) SetEEzsignsignatureAttachmentnamesource(v FieldEEzsignsignatureAttachmentnamesource)`

SetEEzsignsignatureAttachmentnamesource sets EEzsignsignatureAttachmentnamesource field to given value.

### HasEEzsignsignatureAttachmentnamesource

`func (o *EzsignsignatureRequest) HasEEzsignsignatureAttachmentnamesource() bool`

HasEEzsignsignatureAttachmentnamesource returns a boolean if a field has been set.

### GetSEzsignsignatureAttachmentdescription

`func (o *EzsignsignatureRequest) GetSEzsignsignatureAttachmentdescription() string`

GetSEzsignsignatureAttachmentdescription returns the SEzsignsignatureAttachmentdescription field if non-nil, zero value otherwise.

### GetSEzsignsignatureAttachmentdescriptionOk

`func (o *EzsignsignatureRequest) GetSEzsignsignatureAttachmentdescriptionOk() (*string, bool)`

GetSEzsignsignatureAttachmentdescriptionOk returns a tuple with the SEzsignsignatureAttachmentdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureAttachmentdescription

`func (o *EzsignsignatureRequest) SetSEzsignsignatureAttachmentdescription(v string)`

SetSEzsignsignatureAttachmentdescription sets SEzsignsignatureAttachmentdescription field to given value.

### HasSEzsignsignatureAttachmentdescription

`func (o *EzsignsignatureRequest) HasSEzsignsignatureAttachmentdescription() bool`

HasSEzsignsignatureAttachmentdescription returns a boolean if a field has been set.

### GetEEzsignsignatureConsultationtrigger

`func (o *EzsignsignatureRequest) GetEEzsignsignatureConsultationtrigger() FieldEEzsignsignatureConsultationtrigger`

GetEEzsignsignatureConsultationtrigger returns the EEzsignsignatureConsultationtrigger field if non-nil, zero value otherwise.

### GetEEzsignsignatureConsultationtriggerOk

`func (o *EzsignsignatureRequest) GetEEzsignsignatureConsultationtriggerOk() (*FieldEEzsignsignatureConsultationtrigger, bool)`

GetEEzsignsignatureConsultationtriggerOk returns a tuple with the EEzsignsignatureConsultationtrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureConsultationtrigger

`func (o *EzsignsignatureRequest) SetEEzsignsignatureConsultationtrigger(v FieldEEzsignsignatureConsultationtrigger)`

SetEEzsignsignatureConsultationtrigger sets EEzsignsignatureConsultationtrigger field to given value.

### HasEEzsignsignatureConsultationtrigger

`func (o *EzsignsignatureRequest) HasEEzsignsignatureConsultationtrigger() bool`

HasEEzsignsignatureConsultationtrigger returns a boolean if a field has been set.

### GetIEzsignsignatureValidationstep

`func (o *EzsignsignatureRequest) GetIEzsignsignatureValidationstep() int32`

GetIEzsignsignatureValidationstep returns the IEzsignsignatureValidationstep field if non-nil, zero value otherwise.

### GetIEzsignsignatureValidationstepOk

`func (o *EzsignsignatureRequest) GetIEzsignsignatureValidationstepOk() (*int32, bool)`

GetIEzsignsignatureValidationstepOk returns a tuple with the IEzsignsignatureValidationstep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureValidationstep

`func (o *EzsignsignatureRequest) SetIEzsignsignatureValidationstep(v int32)`

SetIEzsignsignatureValidationstep sets IEzsignsignatureValidationstep field to given value.

### HasIEzsignsignatureValidationstep

`func (o *EzsignsignatureRequest) HasIEzsignsignatureValidationstep() bool`

HasIEzsignsignatureValidationstep returns a boolean if a field has been set.

### GetIEzsignsignatureMaxlength

`func (o *EzsignsignatureRequest) GetIEzsignsignatureMaxlength() int32`

GetIEzsignsignatureMaxlength returns the IEzsignsignatureMaxlength field if non-nil, zero value otherwise.

### GetIEzsignsignatureMaxlengthOk

`func (o *EzsignsignatureRequest) GetIEzsignsignatureMaxlengthOk() (*int32, bool)`

GetIEzsignsignatureMaxlengthOk returns a tuple with the IEzsignsignatureMaxlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureMaxlength

`func (o *EzsignsignatureRequest) SetIEzsignsignatureMaxlength(v int32)`

SetIEzsignsignatureMaxlength sets IEzsignsignatureMaxlength field to given value.

### HasIEzsignsignatureMaxlength

`func (o *EzsignsignatureRequest) HasIEzsignsignatureMaxlength() bool`

HasIEzsignsignatureMaxlength returns a boolean if a field has been set.

### GetSEzsignsignatureDefaultvalue

`func (o *EzsignsignatureRequest) GetSEzsignsignatureDefaultvalue() string`

GetSEzsignsignatureDefaultvalue returns the SEzsignsignatureDefaultvalue field if non-nil, zero value otherwise.

### GetSEzsignsignatureDefaultvalueOk

`func (o *EzsignsignatureRequest) GetSEzsignsignatureDefaultvalueOk() (*string, bool)`

GetSEzsignsignatureDefaultvalueOk returns a tuple with the SEzsignsignatureDefaultvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureDefaultvalue

`func (o *EzsignsignatureRequest) SetSEzsignsignatureDefaultvalue(v string)`

SetSEzsignsignatureDefaultvalue sets SEzsignsignatureDefaultvalue field to given value.

### HasSEzsignsignatureDefaultvalue

`func (o *EzsignsignatureRequest) HasSEzsignsignatureDefaultvalue() bool`

HasSEzsignsignatureDefaultvalue returns a boolean if a field has been set.

### GetEEzsignsignatureTextvalidation

`func (o *EzsignsignatureRequest) GetEEzsignsignatureTextvalidation() EnumTextvalidation`

GetEEzsignsignatureTextvalidation returns the EEzsignsignatureTextvalidation field if non-nil, zero value otherwise.

### GetEEzsignsignatureTextvalidationOk

`func (o *EzsignsignatureRequest) GetEEzsignsignatureTextvalidationOk() (*EnumTextvalidation, bool)`

GetEEzsignsignatureTextvalidationOk returns a tuple with the EEzsignsignatureTextvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureTextvalidation

`func (o *EzsignsignatureRequest) SetEEzsignsignatureTextvalidation(v EnumTextvalidation)`

SetEEzsignsignatureTextvalidation sets EEzsignsignatureTextvalidation field to given value.

### HasEEzsignsignatureTextvalidation

`func (o *EzsignsignatureRequest) HasEEzsignsignatureTextvalidation() bool`

HasEEzsignsignatureTextvalidation returns a boolean if a field has been set.

### GetSEzsignsignatureTextvalidationcustommessage

`func (o *EzsignsignatureRequest) GetSEzsignsignatureTextvalidationcustommessage() string`

GetSEzsignsignatureTextvalidationcustommessage returns the SEzsignsignatureTextvalidationcustommessage field if non-nil, zero value otherwise.

### GetSEzsignsignatureTextvalidationcustommessageOk

`func (o *EzsignsignatureRequest) GetSEzsignsignatureTextvalidationcustommessageOk() (*string, bool)`

GetSEzsignsignatureTextvalidationcustommessageOk returns a tuple with the SEzsignsignatureTextvalidationcustommessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureTextvalidationcustommessage

`func (o *EzsignsignatureRequest) SetSEzsignsignatureTextvalidationcustommessage(v string)`

SetSEzsignsignatureTextvalidationcustommessage sets SEzsignsignatureTextvalidationcustommessage field to given value.

### HasSEzsignsignatureTextvalidationcustommessage

`func (o *EzsignsignatureRequest) HasSEzsignsignatureTextvalidationcustommessage() bool`

HasSEzsignsignatureTextvalidationcustommessage returns a boolean if a field has been set.

### GetSEzsignsignatureRegexp

`func (o *EzsignsignatureRequest) GetSEzsignsignatureRegexp() string`

GetSEzsignsignatureRegexp returns the SEzsignsignatureRegexp field if non-nil, zero value otherwise.

### GetSEzsignsignatureRegexpOk

`func (o *EzsignsignatureRequest) GetSEzsignsignatureRegexpOk() (*string, bool)`

GetSEzsignsignatureRegexpOk returns a tuple with the SEzsignsignatureRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureRegexp

`func (o *EzsignsignatureRequest) SetSEzsignsignatureRegexp(v string)`

SetSEzsignsignatureRegexp sets SEzsignsignatureRegexp field to given value.

### HasSEzsignsignatureRegexp

`func (o *EzsignsignatureRequest) HasSEzsignsignatureRegexp() bool`

HasSEzsignsignatureRegexp returns a boolean if a field has been set.

### GetEEzsignsignatureDependencyrequirement

`func (o *EzsignsignatureRequest) GetEEzsignsignatureDependencyrequirement() FieldEEzsignsignatureDependencyrequirement`

GetEEzsignsignatureDependencyrequirement returns the EEzsignsignatureDependencyrequirement field if non-nil, zero value otherwise.

### GetEEzsignsignatureDependencyrequirementOk

`func (o *EzsignsignatureRequest) GetEEzsignsignatureDependencyrequirementOk() (*FieldEEzsignsignatureDependencyrequirement, bool)`

GetEEzsignsignatureDependencyrequirementOk returns a tuple with the EEzsignsignatureDependencyrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureDependencyrequirement

`func (o *EzsignsignatureRequest) SetEEzsignsignatureDependencyrequirement(v FieldEEzsignsignatureDependencyrequirement)`

SetEEzsignsignatureDependencyrequirement sets EEzsignsignatureDependencyrequirement field to given value.

### HasEEzsignsignatureDependencyrequirement

`func (o *EzsignsignatureRequest) HasEEzsignsignatureDependencyrequirement() bool`

HasEEzsignsignatureDependencyrequirement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


