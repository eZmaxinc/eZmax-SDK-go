# EzsignsignatureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignsignatureID** | **int32** | The unique ID of the Ezsignsignature | 
**FkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 
**FkiEzsignfoldersignerassociationID** | **int32** | The unique ID of the Ezsignfoldersignerassociation | 
**FkiEzsignsigningreasonID** | Pointer to **int32** | The unique ID of the Ezsignsigningreason | [optional] 
**FkiFontID** | Pointer to **int32** | The unique ID of the Font | [optional] 
**SEzsignsigningreasonDescriptionX** | Pointer to **string** | The description of the Ezsignsigningreason in the language of the requester | [optional] 
**IEzsignpagePagenumber** | **int32** | The page number in the Ezsigndocument | 
**IEzsignsignatureX** | **int32** | The X coordinate (Horizontal) where to put the Ezsignsignature on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignsignature 2 inches from the left border of the page, you would use \&quot;200\&quot; for the X coordinate. | 
**IEzsignsignatureY** | **int32** | The Y coordinate (Vertical) where to put the Ezsignsignature on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignsignature 3 inches from the top border of the page, you would use \&quot;300\&quot; for the Y coordinate. | 
**IEzsignsignatureHeight** | Pointer to **int32** | The height of the Ezsignsignature.  Size is calculated at 100dpi (dot per inch). So for example, if you want the Ezsignsignature to have an height of 2 inches, you would use \&quot;200\&quot; for the iEzsignsignatureHeight. | [optional] 
**IEzsignsignatureWidth** | Pointer to **int32** | The width of the Ezsignsignature.  Size is calculated at 100dpi (dot per inch). So for example, if you want the Ezsignsignature to have a width of 2 inches, you would use \&quot;200\&quot; for the iEzsignsignatureWidth. | [optional] 
**IEzsignsignatureStep** | **int32** | The step when the Ezsignsigner will be invited to sign | 
**IEzsignsignatureStepadjusted** | Pointer to **int32** | The step when the Ezsignsigner will be invited to sign | [optional] 
**EEzsignsignatureType** | [**FieldEEzsignsignatureType**](FieldEEzsignsignatureType.md) |  | 
**TEzsignsignatureTooltip** | Pointer to **string** | A tooltip that will be presented to Ezsignsigner about the Ezsignsignature | [optional] 
**EEzsignsignatureTooltipposition** | Pointer to [**FieldEEzsignsignatureTooltipposition**](FieldEEzsignsignatureTooltipposition.md) |  | [optional] 
**EEzsignsignatureFont** | Pointer to [**FieldEEzsignsignatureFont**](FieldEEzsignsignatureFont.md) |  | [optional] 
**IEzsignsignatureValidationstep** | Pointer to **int32** | The step when the Ezsignsigner will be invited to validate the Ezsignsignature of eEzsignsignatureType Attachments | [optional] 
**SEzsignsignatureAttachmentdescription** | Pointer to **string** | The description attached to the attachment name added in Ezsignsignature of eEzsignsignatureType Attachments | [optional] 
**EEzsignsignatureAttachmentnamesource** | Pointer to [**FieldEEzsignsignatureAttachmentnamesource**](FieldEEzsignsignatureAttachmentnamesource.md) |  | [optional] 
**EEzsignsignatureConsultationtrigger** | Pointer to [**FieldEEzsignsignatureConsultationtrigger**](FieldEEzsignsignatureConsultationtrigger.md) |  | [optional] 
**BEzsignsignatureHandwritten** | Pointer to **bool** | Whether the Ezsignsignature must be handwritten or not when eEzsignsignatureType &#x3D; Signature. | [optional] 
**BEzsignsignatureReason** | Pointer to **bool** | Whether the Ezsignsignature must include a reason or not when eEzsignsignatureType &#x3D; Signature. | [optional] 
**BEzsignsignatureRequired** | Pointer to **bool** | Whether the Ezsignsignature is required or not. This field is relevant only with Ezsignsignature with eEzsignsignatureType &#x3D; Attachments, Text or Textarea. | [optional] 
**FkiEzsignfoldersignerassociationIDValidation** | Pointer to **int32** | The unique ID of the Ezsignfoldersignerassociation | [optional] 
**DtEzsignsignatureDate** | Pointer to **string** | The date the Ezsignsignature was signed | [optional] 
**IEzsignsignatureattachmentCount** | Pointer to **int32** | The count of Ezsignsignatureattachment | [optional] 
**SEzsignsignatureDescription** | Pointer to **string** | The value entered while signing Ezsignsignature of eEzsignsignatureType **City**, **FieldText** and **FieldTextarea** | [optional] 
**IEzsignsignatureMaxlength** | Pointer to **int32** | The maximum length for the value in the Ezsignsignature  This can only be set if eEzsignsignatureType is **FieldText** or **FieldTextarea** | [optional] 
**EEzsignsignatureTextvalidation** | Pointer to [**EnumTextvalidation**](EnumTextvalidation.md) |  | [optional] 
**SEzsignsignatureTextvalidationcustommessage** | Pointer to **string** | Description of validation rule. Show by signatory. | [optional] 
**EEzsignsignatureDependencyrequirement** | Pointer to [**FieldEEzsignsignatureDependencyrequirement**](FieldEEzsignsignatureDependencyrequirement.md) |  | [optional] 
**SEzsignsignatureDefaultvalue** | Pointer to **string** | The default value for the Ezsignsignature  You can use the codes below and they will be replaced at signature time.    | Code | Description | Example | | ------------------------- | ------------ | ------------ | | {sUserFirstname} | The first name of the contact | John | | {sUserLastname} | The last name of the contact | Doe | | {sUserJobtitle} | The job title | Sales Representative | | {sCompany} | Company name | eZmax Solutions Inc. | | {sEmailAddress} | The email address | email@example.com | | {sPhoneE164} | A phone number in E.164 Format | +15149901516 | | {sPhoneE164Cell} | A phone number in E.164 Format | +15149901516 | | [optional] 
**SEzsignsignatureRegexp** | Pointer to **string** | A regular expression to indicate what values are acceptable for the Ezsignsignature.  This can only be set if eEzsignsignatureType is **FieldText** or **FieldTextarea** and eEzsignsignatureTextvalidation is **Custom** | [optional] 
**ObjContactName** | [**CustomContactNameResponse**](CustomContactNameResponse.md) |  | 
**ObjContactNameDelegation** | Pointer to [**CustomContactNameResponse**](CustomContactNameResponse.md) |  | [optional] 
**ObjSignature** | Pointer to [**SignatureResponseCompound**](SignatureResponseCompound.md) |  | [optional] 

## Methods

### NewEzsignsignatureResponse

`func NewEzsignsignatureResponse(pkiEzsignsignatureID int32, fkiEzsigndocumentID int32, fkiEzsignfoldersignerassociationID int32, iEzsignpagePagenumber int32, iEzsignsignatureX int32, iEzsignsignatureY int32, iEzsignsignatureStep int32, eEzsignsignatureType FieldEEzsignsignatureType, objContactName CustomContactNameResponse, ) *EzsignsignatureResponse`

NewEzsignsignatureResponse instantiates a new EzsignsignatureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignatureResponseWithDefaults

`func NewEzsignsignatureResponseWithDefaults() *EzsignsignatureResponse`

NewEzsignsignatureResponseWithDefaults instantiates a new EzsignsignatureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignsignatureID

`func (o *EzsignsignatureResponse) GetPkiEzsignsignatureID() int32`

GetPkiEzsignsignatureID returns the PkiEzsignsignatureID field if non-nil, zero value otherwise.

### GetPkiEzsignsignatureIDOk

`func (o *EzsignsignatureResponse) GetPkiEzsignsignatureIDOk() (*int32, bool)`

GetPkiEzsignsignatureIDOk returns a tuple with the PkiEzsignsignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignsignatureID

`func (o *EzsignsignatureResponse) SetPkiEzsignsignatureID(v int32)`

SetPkiEzsignsignatureID sets PkiEzsignsignatureID field to given value.


### GetFkiEzsigndocumentID

`func (o *EzsignsignatureResponse) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *EzsignsignatureResponse) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *EzsignsignatureResponse) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.


### GetFkiEzsignfoldersignerassociationID

`func (o *EzsignsignatureResponse) GetFkiEzsignfoldersignerassociationID() int32`

GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDOk

`func (o *EzsignsignatureResponse) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationID

`func (o *EzsignsignatureResponse) SetFkiEzsignfoldersignerassociationID(v int32)`

SetFkiEzsignfoldersignerassociationID sets FkiEzsignfoldersignerassociationID field to given value.


### GetFkiEzsignsigningreasonID

`func (o *EzsignsignatureResponse) GetFkiEzsignsigningreasonID() int32`

GetFkiEzsignsigningreasonID returns the FkiEzsignsigningreasonID field if non-nil, zero value otherwise.

### GetFkiEzsignsigningreasonIDOk

`func (o *EzsignsignatureResponse) GetFkiEzsignsigningreasonIDOk() (*int32, bool)`

GetFkiEzsignsigningreasonIDOk returns a tuple with the FkiEzsignsigningreasonID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsigningreasonID

`func (o *EzsignsignatureResponse) SetFkiEzsignsigningreasonID(v int32)`

SetFkiEzsignsigningreasonID sets FkiEzsignsigningreasonID field to given value.

### HasFkiEzsignsigningreasonID

`func (o *EzsignsignatureResponse) HasFkiEzsignsigningreasonID() bool`

HasFkiEzsignsigningreasonID returns a boolean if a field has been set.

### GetFkiFontID

`func (o *EzsignsignatureResponse) GetFkiFontID() int32`

GetFkiFontID returns the FkiFontID field if non-nil, zero value otherwise.

### GetFkiFontIDOk

`func (o *EzsignsignatureResponse) GetFkiFontIDOk() (*int32, bool)`

GetFkiFontIDOk returns a tuple with the FkiFontID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontID

`func (o *EzsignsignatureResponse) SetFkiFontID(v int32)`

SetFkiFontID sets FkiFontID field to given value.

### HasFkiFontID

`func (o *EzsignsignatureResponse) HasFkiFontID() bool`

HasFkiFontID returns a boolean if a field has been set.

### GetSEzsignsigningreasonDescriptionX

`func (o *EzsignsignatureResponse) GetSEzsignsigningreasonDescriptionX() string`

GetSEzsignsigningreasonDescriptionX returns the SEzsignsigningreasonDescriptionX field if non-nil, zero value otherwise.

### GetSEzsignsigningreasonDescriptionXOk

`func (o *EzsignsignatureResponse) GetSEzsignsigningreasonDescriptionXOk() (*string, bool)`

GetSEzsignsigningreasonDescriptionXOk returns a tuple with the SEzsignsigningreasonDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsigningreasonDescriptionX

`func (o *EzsignsignatureResponse) SetSEzsignsigningreasonDescriptionX(v string)`

SetSEzsignsigningreasonDescriptionX sets SEzsignsigningreasonDescriptionX field to given value.

### HasSEzsignsigningreasonDescriptionX

`func (o *EzsignsignatureResponse) HasSEzsignsigningreasonDescriptionX() bool`

HasSEzsignsigningreasonDescriptionX returns a boolean if a field has been set.

### GetIEzsignpagePagenumber

`func (o *EzsignsignatureResponse) GetIEzsignpagePagenumber() int32`

GetIEzsignpagePagenumber returns the IEzsignpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsignpagePagenumberOk

`func (o *EzsignsignatureResponse) GetIEzsignpagePagenumberOk() (*int32, bool)`

GetIEzsignpagePagenumberOk returns a tuple with the IEzsignpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignpagePagenumber

`func (o *EzsignsignatureResponse) SetIEzsignpagePagenumber(v int32)`

SetIEzsignpagePagenumber sets IEzsignpagePagenumber field to given value.


### GetIEzsignsignatureX

`func (o *EzsignsignatureResponse) GetIEzsignsignatureX() int32`

GetIEzsignsignatureX returns the IEzsignsignatureX field if non-nil, zero value otherwise.

### GetIEzsignsignatureXOk

`func (o *EzsignsignatureResponse) GetIEzsignsignatureXOk() (*int32, bool)`

GetIEzsignsignatureXOk returns a tuple with the IEzsignsignatureX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureX

`func (o *EzsignsignatureResponse) SetIEzsignsignatureX(v int32)`

SetIEzsignsignatureX sets IEzsignsignatureX field to given value.


### GetIEzsignsignatureY

`func (o *EzsignsignatureResponse) GetIEzsignsignatureY() int32`

GetIEzsignsignatureY returns the IEzsignsignatureY field if non-nil, zero value otherwise.

### GetIEzsignsignatureYOk

`func (o *EzsignsignatureResponse) GetIEzsignsignatureYOk() (*int32, bool)`

GetIEzsignsignatureYOk returns a tuple with the IEzsignsignatureY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureY

`func (o *EzsignsignatureResponse) SetIEzsignsignatureY(v int32)`

SetIEzsignsignatureY sets IEzsignsignatureY field to given value.


### GetIEzsignsignatureHeight

`func (o *EzsignsignatureResponse) GetIEzsignsignatureHeight() int32`

GetIEzsignsignatureHeight returns the IEzsignsignatureHeight field if non-nil, zero value otherwise.

### GetIEzsignsignatureHeightOk

`func (o *EzsignsignatureResponse) GetIEzsignsignatureHeightOk() (*int32, bool)`

GetIEzsignsignatureHeightOk returns a tuple with the IEzsignsignatureHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureHeight

`func (o *EzsignsignatureResponse) SetIEzsignsignatureHeight(v int32)`

SetIEzsignsignatureHeight sets IEzsignsignatureHeight field to given value.

### HasIEzsignsignatureHeight

`func (o *EzsignsignatureResponse) HasIEzsignsignatureHeight() bool`

HasIEzsignsignatureHeight returns a boolean if a field has been set.

### GetIEzsignsignatureWidth

`func (o *EzsignsignatureResponse) GetIEzsignsignatureWidth() int32`

GetIEzsignsignatureWidth returns the IEzsignsignatureWidth field if non-nil, zero value otherwise.

### GetIEzsignsignatureWidthOk

`func (o *EzsignsignatureResponse) GetIEzsignsignatureWidthOk() (*int32, bool)`

GetIEzsignsignatureWidthOk returns a tuple with the IEzsignsignatureWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureWidth

`func (o *EzsignsignatureResponse) SetIEzsignsignatureWidth(v int32)`

SetIEzsignsignatureWidth sets IEzsignsignatureWidth field to given value.

### HasIEzsignsignatureWidth

`func (o *EzsignsignatureResponse) HasIEzsignsignatureWidth() bool`

HasIEzsignsignatureWidth returns a boolean if a field has been set.

### GetIEzsignsignatureStep

`func (o *EzsignsignatureResponse) GetIEzsignsignatureStep() int32`

GetIEzsignsignatureStep returns the IEzsignsignatureStep field if non-nil, zero value otherwise.

### GetIEzsignsignatureStepOk

`func (o *EzsignsignatureResponse) GetIEzsignsignatureStepOk() (*int32, bool)`

GetIEzsignsignatureStepOk returns a tuple with the IEzsignsignatureStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureStep

`func (o *EzsignsignatureResponse) SetIEzsignsignatureStep(v int32)`

SetIEzsignsignatureStep sets IEzsignsignatureStep field to given value.


### GetIEzsignsignatureStepadjusted

`func (o *EzsignsignatureResponse) GetIEzsignsignatureStepadjusted() int32`

GetIEzsignsignatureStepadjusted returns the IEzsignsignatureStepadjusted field if non-nil, zero value otherwise.

### GetIEzsignsignatureStepadjustedOk

`func (o *EzsignsignatureResponse) GetIEzsignsignatureStepadjustedOk() (*int32, bool)`

GetIEzsignsignatureStepadjustedOk returns a tuple with the IEzsignsignatureStepadjusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureStepadjusted

`func (o *EzsignsignatureResponse) SetIEzsignsignatureStepadjusted(v int32)`

SetIEzsignsignatureStepadjusted sets IEzsignsignatureStepadjusted field to given value.

### HasIEzsignsignatureStepadjusted

`func (o *EzsignsignatureResponse) HasIEzsignsignatureStepadjusted() bool`

HasIEzsignsignatureStepadjusted returns a boolean if a field has been set.

### GetEEzsignsignatureType

`func (o *EzsignsignatureResponse) GetEEzsignsignatureType() FieldEEzsignsignatureType`

GetEEzsignsignatureType returns the EEzsignsignatureType field if non-nil, zero value otherwise.

### GetEEzsignsignatureTypeOk

`func (o *EzsignsignatureResponse) GetEEzsignsignatureTypeOk() (*FieldEEzsignsignatureType, bool)`

GetEEzsignsignatureTypeOk returns a tuple with the EEzsignsignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureType

`func (o *EzsignsignatureResponse) SetEEzsignsignatureType(v FieldEEzsignsignatureType)`

SetEEzsignsignatureType sets EEzsignsignatureType field to given value.


### GetTEzsignsignatureTooltip

`func (o *EzsignsignatureResponse) GetTEzsignsignatureTooltip() string`

GetTEzsignsignatureTooltip returns the TEzsignsignatureTooltip field if non-nil, zero value otherwise.

### GetTEzsignsignatureTooltipOk

`func (o *EzsignsignatureResponse) GetTEzsignsignatureTooltipOk() (*string, bool)`

GetTEzsignsignatureTooltipOk returns a tuple with the TEzsignsignatureTooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignsignatureTooltip

`func (o *EzsignsignatureResponse) SetTEzsignsignatureTooltip(v string)`

SetTEzsignsignatureTooltip sets TEzsignsignatureTooltip field to given value.

### HasTEzsignsignatureTooltip

`func (o *EzsignsignatureResponse) HasTEzsignsignatureTooltip() bool`

HasTEzsignsignatureTooltip returns a boolean if a field has been set.

### GetEEzsignsignatureTooltipposition

`func (o *EzsignsignatureResponse) GetEEzsignsignatureTooltipposition() FieldEEzsignsignatureTooltipposition`

GetEEzsignsignatureTooltipposition returns the EEzsignsignatureTooltipposition field if non-nil, zero value otherwise.

### GetEEzsignsignatureTooltippositionOk

`func (o *EzsignsignatureResponse) GetEEzsignsignatureTooltippositionOk() (*FieldEEzsignsignatureTooltipposition, bool)`

GetEEzsignsignatureTooltippositionOk returns a tuple with the EEzsignsignatureTooltipposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureTooltipposition

`func (o *EzsignsignatureResponse) SetEEzsignsignatureTooltipposition(v FieldEEzsignsignatureTooltipposition)`

SetEEzsignsignatureTooltipposition sets EEzsignsignatureTooltipposition field to given value.

### HasEEzsignsignatureTooltipposition

`func (o *EzsignsignatureResponse) HasEEzsignsignatureTooltipposition() bool`

HasEEzsignsignatureTooltipposition returns a boolean if a field has been set.

### GetEEzsignsignatureFont

`func (o *EzsignsignatureResponse) GetEEzsignsignatureFont() FieldEEzsignsignatureFont`

GetEEzsignsignatureFont returns the EEzsignsignatureFont field if non-nil, zero value otherwise.

### GetEEzsignsignatureFontOk

`func (o *EzsignsignatureResponse) GetEEzsignsignatureFontOk() (*FieldEEzsignsignatureFont, bool)`

GetEEzsignsignatureFontOk returns a tuple with the EEzsignsignatureFont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureFont

`func (o *EzsignsignatureResponse) SetEEzsignsignatureFont(v FieldEEzsignsignatureFont)`

SetEEzsignsignatureFont sets EEzsignsignatureFont field to given value.

### HasEEzsignsignatureFont

`func (o *EzsignsignatureResponse) HasEEzsignsignatureFont() bool`

HasEEzsignsignatureFont returns a boolean if a field has been set.

### GetIEzsignsignatureValidationstep

`func (o *EzsignsignatureResponse) GetIEzsignsignatureValidationstep() int32`

GetIEzsignsignatureValidationstep returns the IEzsignsignatureValidationstep field if non-nil, zero value otherwise.

### GetIEzsignsignatureValidationstepOk

`func (o *EzsignsignatureResponse) GetIEzsignsignatureValidationstepOk() (*int32, bool)`

GetIEzsignsignatureValidationstepOk returns a tuple with the IEzsignsignatureValidationstep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureValidationstep

`func (o *EzsignsignatureResponse) SetIEzsignsignatureValidationstep(v int32)`

SetIEzsignsignatureValidationstep sets IEzsignsignatureValidationstep field to given value.

### HasIEzsignsignatureValidationstep

`func (o *EzsignsignatureResponse) HasIEzsignsignatureValidationstep() bool`

HasIEzsignsignatureValidationstep returns a boolean if a field has been set.

### GetSEzsignsignatureAttachmentdescription

`func (o *EzsignsignatureResponse) GetSEzsignsignatureAttachmentdescription() string`

GetSEzsignsignatureAttachmentdescription returns the SEzsignsignatureAttachmentdescription field if non-nil, zero value otherwise.

### GetSEzsignsignatureAttachmentdescriptionOk

`func (o *EzsignsignatureResponse) GetSEzsignsignatureAttachmentdescriptionOk() (*string, bool)`

GetSEzsignsignatureAttachmentdescriptionOk returns a tuple with the SEzsignsignatureAttachmentdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureAttachmentdescription

`func (o *EzsignsignatureResponse) SetSEzsignsignatureAttachmentdescription(v string)`

SetSEzsignsignatureAttachmentdescription sets SEzsignsignatureAttachmentdescription field to given value.

### HasSEzsignsignatureAttachmentdescription

`func (o *EzsignsignatureResponse) HasSEzsignsignatureAttachmentdescription() bool`

HasSEzsignsignatureAttachmentdescription returns a boolean if a field has been set.

### GetEEzsignsignatureAttachmentnamesource

`func (o *EzsignsignatureResponse) GetEEzsignsignatureAttachmentnamesource() FieldEEzsignsignatureAttachmentnamesource`

GetEEzsignsignatureAttachmentnamesource returns the EEzsignsignatureAttachmentnamesource field if non-nil, zero value otherwise.

### GetEEzsignsignatureAttachmentnamesourceOk

`func (o *EzsignsignatureResponse) GetEEzsignsignatureAttachmentnamesourceOk() (*FieldEEzsignsignatureAttachmentnamesource, bool)`

GetEEzsignsignatureAttachmentnamesourceOk returns a tuple with the EEzsignsignatureAttachmentnamesource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureAttachmentnamesource

`func (o *EzsignsignatureResponse) SetEEzsignsignatureAttachmentnamesource(v FieldEEzsignsignatureAttachmentnamesource)`

SetEEzsignsignatureAttachmentnamesource sets EEzsignsignatureAttachmentnamesource field to given value.

### HasEEzsignsignatureAttachmentnamesource

`func (o *EzsignsignatureResponse) HasEEzsignsignatureAttachmentnamesource() bool`

HasEEzsignsignatureAttachmentnamesource returns a boolean if a field has been set.

### GetEEzsignsignatureConsultationtrigger

`func (o *EzsignsignatureResponse) GetEEzsignsignatureConsultationtrigger() FieldEEzsignsignatureConsultationtrigger`

GetEEzsignsignatureConsultationtrigger returns the EEzsignsignatureConsultationtrigger field if non-nil, zero value otherwise.

### GetEEzsignsignatureConsultationtriggerOk

`func (o *EzsignsignatureResponse) GetEEzsignsignatureConsultationtriggerOk() (*FieldEEzsignsignatureConsultationtrigger, bool)`

GetEEzsignsignatureConsultationtriggerOk returns a tuple with the EEzsignsignatureConsultationtrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureConsultationtrigger

`func (o *EzsignsignatureResponse) SetEEzsignsignatureConsultationtrigger(v FieldEEzsignsignatureConsultationtrigger)`

SetEEzsignsignatureConsultationtrigger sets EEzsignsignatureConsultationtrigger field to given value.

### HasEEzsignsignatureConsultationtrigger

`func (o *EzsignsignatureResponse) HasEEzsignsignatureConsultationtrigger() bool`

HasEEzsignsignatureConsultationtrigger returns a boolean if a field has been set.

### GetBEzsignsignatureHandwritten

`func (o *EzsignsignatureResponse) GetBEzsignsignatureHandwritten() bool`

GetBEzsignsignatureHandwritten returns the BEzsignsignatureHandwritten field if non-nil, zero value otherwise.

### GetBEzsignsignatureHandwrittenOk

`func (o *EzsignsignatureResponse) GetBEzsignsignatureHandwrittenOk() (*bool, bool)`

GetBEzsignsignatureHandwrittenOk returns a tuple with the BEzsignsignatureHandwritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureHandwritten

`func (o *EzsignsignatureResponse) SetBEzsignsignatureHandwritten(v bool)`

SetBEzsignsignatureHandwritten sets BEzsignsignatureHandwritten field to given value.

### HasBEzsignsignatureHandwritten

`func (o *EzsignsignatureResponse) HasBEzsignsignatureHandwritten() bool`

HasBEzsignsignatureHandwritten returns a boolean if a field has been set.

### GetBEzsignsignatureReason

`func (o *EzsignsignatureResponse) GetBEzsignsignatureReason() bool`

GetBEzsignsignatureReason returns the BEzsignsignatureReason field if non-nil, zero value otherwise.

### GetBEzsignsignatureReasonOk

`func (o *EzsignsignatureResponse) GetBEzsignsignatureReasonOk() (*bool, bool)`

GetBEzsignsignatureReasonOk returns a tuple with the BEzsignsignatureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureReason

`func (o *EzsignsignatureResponse) SetBEzsignsignatureReason(v bool)`

SetBEzsignsignatureReason sets BEzsignsignatureReason field to given value.

### HasBEzsignsignatureReason

`func (o *EzsignsignatureResponse) HasBEzsignsignatureReason() bool`

HasBEzsignsignatureReason returns a boolean if a field has been set.

### GetBEzsignsignatureRequired

`func (o *EzsignsignatureResponse) GetBEzsignsignatureRequired() bool`

GetBEzsignsignatureRequired returns the BEzsignsignatureRequired field if non-nil, zero value otherwise.

### GetBEzsignsignatureRequiredOk

`func (o *EzsignsignatureResponse) GetBEzsignsignatureRequiredOk() (*bool, bool)`

GetBEzsignsignatureRequiredOk returns a tuple with the BEzsignsignatureRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignsignatureRequired

`func (o *EzsignsignatureResponse) SetBEzsignsignatureRequired(v bool)`

SetBEzsignsignatureRequired sets BEzsignsignatureRequired field to given value.

### HasBEzsignsignatureRequired

`func (o *EzsignsignatureResponse) HasBEzsignsignatureRequired() bool`

HasBEzsignsignatureRequired returns a boolean if a field has been set.

### GetFkiEzsignfoldersignerassociationIDValidation

`func (o *EzsignsignatureResponse) GetFkiEzsignfoldersignerassociationIDValidation() int32`

GetFkiEzsignfoldersignerassociationIDValidation returns the FkiEzsignfoldersignerassociationIDValidation field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDValidationOk

`func (o *EzsignsignatureResponse) GetFkiEzsignfoldersignerassociationIDValidationOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDValidationOk returns a tuple with the FkiEzsignfoldersignerassociationIDValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationIDValidation

`func (o *EzsignsignatureResponse) SetFkiEzsignfoldersignerassociationIDValidation(v int32)`

SetFkiEzsignfoldersignerassociationIDValidation sets FkiEzsignfoldersignerassociationIDValidation field to given value.

### HasFkiEzsignfoldersignerassociationIDValidation

`func (o *EzsignsignatureResponse) HasFkiEzsignfoldersignerassociationIDValidation() bool`

HasFkiEzsignfoldersignerassociationIDValidation returns a boolean if a field has been set.

### GetDtEzsignsignatureDate

`func (o *EzsignsignatureResponse) GetDtEzsignsignatureDate() string`

GetDtEzsignsignatureDate returns the DtEzsignsignatureDate field if non-nil, zero value otherwise.

### GetDtEzsignsignatureDateOk

`func (o *EzsignsignatureResponse) GetDtEzsignsignatureDateOk() (*string, bool)`

GetDtEzsignsignatureDateOk returns a tuple with the DtEzsignsignatureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignsignatureDate

`func (o *EzsignsignatureResponse) SetDtEzsignsignatureDate(v string)`

SetDtEzsignsignatureDate sets DtEzsignsignatureDate field to given value.

### HasDtEzsignsignatureDate

`func (o *EzsignsignatureResponse) HasDtEzsignsignatureDate() bool`

HasDtEzsignsignatureDate returns a boolean if a field has been set.

### GetIEzsignsignatureattachmentCount

`func (o *EzsignsignatureResponse) GetIEzsignsignatureattachmentCount() int32`

GetIEzsignsignatureattachmentCount returns the IEzsignsignatureattachmentCount field if non-nil, zero value otherwise.

### GetIEzsignsignatureattachmentCountOk

`func (o *EzsignsignatureResponse) GetIEzsignsignatureattachmentCountOk() (*int32, bool)`

GetIEzsignsignatureattachmentCountOk returns a tuple with the IEzsignsignatureattachmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureattachmentCount

`func (o *EzsignsignatureResponse) SetIEzsignsignatureattachmentCount(v int32)`

SetIEzsignsignatureattachmentCount sets IEzsignsignatureattachmentCount field to given value.

### HasIEzsignsignatureattachmentCount

`func (o *EzsignsignatureResponse) HasIEzsignsignatureattachmentCount() bool`

HasIEzsignsignatureattachmentCount returns a boolean if a field has been set.

### GetSEzsignsignatureDescription

`func (o *EzsignsignatureResponse) GetSEzsignsignatureDescription() string`

GetSEzsignsignatureDescription returns the SEzsignsignatureDescription field if non-nil, zero value otherwise.

### GetSEzsignsignatureDescriptionOk

`func (o *EzsignsignatureResponse) GetSEzsignsignatureDescriptionOk() (*string, bool)`

GetSEzsignsignatureDescriptionOk returns a tuple with the SEzsignsignatureDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureDescription

`func (o *EzsignsignatureResponse) SetSEzsignsignatureDescription(v string)`

SetSEzsignsignatureDescription sets SEzsignsignatureDescription field to given value.

### HasSEzsignsignatureDescription

`func (o *EzsignsignatureResponse) HasSEzsignsignatureDescription() bool`

HasSEzsignsignatureDescription returns a boolean if a field has been set.

### GetIEzsignsignatureMaxlength

`func (o *EzsignsignatureResponse) GetIEzsignsignatureMaxlength() int32`

GetIEzsignsignatureMaxlength returns the IEzsignsignatureMaxlength field if non-nil, zero value otherwise.

### GetIEzsignsignatureMaxlengthOk

`func (o *EzsignsignatureResponse) GetIEzsignsignatureMaxlengthOk() (*int32, bool)`

GetIEzsignsignatureMaxlengthOk returns a tuple with the IEzsignsignatureMaxlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureMaxlength

`func (o *EzsignsignatureResponse) SetIEzsignsignatureMaxlength(v int32)`

SetIEzsignsignatureMaxlength sets IEzsignsignatureMaxlength field to given value.

### HasIEzsignsignatureMaxlength

`func (o *EzsignsignatureResponse) HasIEzsignsignatureMaxlength() bool`

HasIEzsignsignatureMaxlength returns a boolean if a field has been set.

### GetEEzsignsignatureTextvalidation

`func (o *EzsignsignatureResponse) GetEEzsignsignatureTextvalidation() EnumTextvalidation`

GetEEzsignsignatureTextvalidation returns the EEzsignsignatureTextvalidation field if non-nil, zero value otherwise.

### GetEEzsignsignatureTextvalidationOk

`func (o *EzsignsignatureResponse) GetEEzsignsignatureTextvalidationOk() (*EnumTextvalidation, bool)`

GetEEzsignsignatureTextvalidationOk returns a tuple with the EEzsignsignatureTextvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureTextvalidation

`func (o *EzsignsignatureResponse) SetEEzsignsignatureTextvalidation(v EnumTextvalidation)`

SetEEzsignsignatureTextvalidation sets EEzsignsignatureTextvalidation field to given value.

### HasEEzsignsignatureTextvalidation

`func (o *EzsignsignatureResponse) HasEEzsignsignatureTextvalidation() bool`

HasEEzsignsignatureTextvalidation returns a boolean if a field has been set.

### GetSEzsignsignatureTextvalidationcustommessage

`func (o *EzsignsignatureResponse) GetSEzsignsignatureTextvalidationcustommessage() string`

GetSEzsignsignatureTextvalidationcustommessage returns the SEzsignsignatureTextvalidationcustommessage field if non-nil, zero value otherwise.

### GetSEzsignsignatureTextvalidationcustommessageOk

`func (o *EzsignsignatureResponse) GetSEzsignsignatureTextvalidationcustommessageOk() (*string, bool)`

GetSEzsignsignatureTextvalidationcustommessageOk returns a tuple with the SEzsignsignatureTextvalidationcustommessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureTextvalidationcustommessage

`func (o *EzsignsignatureResponse) SetSEzsignsignatureTextvalidationcustommessage(v string)`

SetSEzsignsignatureTextvalidationcustommessage sets SEzsignsignatureTextvalidationcustommessage field to given value.

### HasSEzsignsignatureTextvalidationcustommessage

`func (o *EzsignsignatureResponse) HasSEzsignsignatureTextvalidationcustommessage() bool`

HasSEzsignsignatureTextvalidationcustommessage returns a boolean if a field has been set.

### GetEEzsignsignatureDependencyrequirement

`func (o *EzsignsignatureResponse) GetEEzsignsignatureDependencyrequirement() FieldEEzsignsignatureDependencyrequirement`

GetEEzsignsignatureDependencyrequirement returns the EEzsignsignatureDependencyrequirement field if non-nil, zero value otherwise.

### GetEEzsignsignatureDependencyrequirementOk

`func (o *EzsignsignatureResponse) GetEEzsignsignatureDependencyrequirementOk() (*FieldEEzsignsignatureDependencyrequirement, bool)`

GetEEzsignsignatureDependencyrequirementOk returns a tuple with the EEzsignsignatureDependencyrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureDependencyrequirement

`func (o *EzsignsignatureResponse) SetEEzsignsignatureDependencyrequirement(v FieldEEzsignsignatureDependencyrequirement)`

SetEEzsignsignatureDependencyrequirement sets EEzsignsignatureDependencyrequirement field to given value.

### HasEEzsignsignatureDependencyrequirement

`func (o *EzsignsignatureResponse) HasEEzsignsignatureDependencyrequirement() bool`

HasEEzsignsignatureDependencyrequirement returns a boolean if a field has been set.

### GetSEzsignsignatureDefaultvalue

`func (o *EzsignsignatureResponse) GetSEzsignsignatureDefaultvalue() string`

GetSEzsignsignatureDefaultvalue returns the SEzsignsignatureDefaultvalue field if non-nil, zero value otherwise.

### GetSEzsignsignatureDefaultvalueOk

`func (o *EzsignsignatureResponse) GetSEzsignsignatureDefaultvalueOk() (*string, bool)`

GetSEzsignsignatureDefaultvalueOk returns a tuple with the SEzsignsignatureDefaultvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureDefaultvalue

`func (o *EzsignsignatureResponse) SetSEzsignsignatureDefaultvalue(v string)`

SetSEzsignsignatureDefaultvalue sets SEzsignsignatureDefaultvalue field to given value.

### HasSEzsignsignatureDefaultvalue

`func (o *EzsignsignatureResponse) HasSEzsignsignatureDefaultvalue() bool`

HasSEzsignsignatureDefaultvalue returns a boolean if a field has been set.

### GetSEzsignsignatureRegexp

`func (o *EzsignsignatureResponse) GetSEzsignsignatureRegexp() string`

GetSEzsignsignatureRegexp returns the SEzsignsignatureRegexp field if non-nil, zero value otherwise.

### GetSEzsignsignatureRegexpOk

`func (o *EzsignsignatureResponse) GetSEzsignsignatureRegexpOk() (*string, bool)`

GetSEzsignsignatureRegexpOk returns a tuple with the SEzsignsignatureRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignsignatureRegexp

`func (o *EzsignsignatureResponse) SetSEzsignsignatureRegexp(v string)`

SetSEzsignsignatureRegexp sets SEzsignsignatureRegexp field to given value.

### HasSEzsignsignatureRegexp

`func (o *EzsignsignatureResponse) HasSEzsignsignatureRegexp() bool`

HasSEzsignsignatureRegexp returns a boolean if a field has been set.

### GetObjContactName

`func (o *EzsignsignatureResponse) GetObjContactName() CustomContactNameResponse`

GetObjContactName returns the ObjContactName field if non-nil, zero value otherwise.

### GetObjContactNameOk

`func (o *EzsignsignatureResponse) GetObjContactNameOk() (*CustomContactNameResponse, bool)`

GetObjContactNameOk returns a tuple with the ObjContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContactName

`func (o *EzsignsignatureResponse) SetObjContactName(v CustomContactNameResponse)`

SetObjContactName sets ObjContactName field to given value.


### GetObjContactNameDelegation

`func (o *EzsignsignatureResponse) GetObjContactNameDelegation() CustomContactNameResponse`

GetObjContactNameDelegation returns the ObjContactNameDelegation field if non-nil, zero value otherwise.

### GetObjContactNameDelegationOk

`func (o *EzsignsignatureResponse) GetObjContactNameDelegationOk() (*CustomContactNameResponse, bool)`

GetObjContactNameDelegationOk returns a tuple with the ObjContactNameDelegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContactNameDelegation

`func (o *EzsignsignatureResponse) SetObjContactNameDelegation(v CustomContactNameResponse)`

SetObjContactNameDelegation sets ObjContactNameDelegation field to given value.

### HasObjContactNameDelegation

`func (o *EzsignsignatureResponse) HasObjContactNameDelegation() bool`

HasObjContactNameDelegation returns a boolean if a field has been set.

### GetObjSignature

`func (o *EzsignsignatureResponse) GetObjSignature() SignatureResponseCompound`

GetObjSignature returns the ObjSignature field if non-nil, zero value otherwise.

### GetObjSignatureOk

`func (o *EzsignsignatureResponse) GetObjSignatureOk() (*SignatureResponseCompound, bool)`

GetObjSignatureOk returns a tuple with the ObjSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjSignature

`func (o *EzsignsignatureResponse) SetObjSignature(v SignatureResponseCompound)`

SetObjSignature sets ObjSignature field to given value.

### HasObjSignature

`func (o *EzsignsignatureResponse) HasObjSignature() bool`

HasObjSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


