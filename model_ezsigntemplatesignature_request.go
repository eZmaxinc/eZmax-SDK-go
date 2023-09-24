/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// checks if the EzsigntemplatesignatureRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatesignatureRequest{}

// EzsigntemplatesignatureRequest A Ezsigntemplatesignature Object
type EzsigntemplatesignatureRequest struct {
	// The unique ID of the Ezsigntemplatesignature
	PkiEzsigntemplatesignatureID *int32 `json:"pkiEzsigntemplatesignatureID,omitempty"`
	// The unique ID of the Ezsigntemplatedocument
	FkiEzsigntemplatedocumentID int32 `json:"fkiEzsigntemplatedocumentID"`
	// The unique ID of the Ezsigntemplatesigner
	FkiEzsigntemplatesignerID int32 `json:"fkiEzsigntemplatesignerID"`
	// The unique ID of the Ezsigntemplatesigner
	FkiEzsigntemplatesignerIDValidation *int32 `json:"fkiEzsigntemplatesignerIDValidation,omitempty"`
	// The page number in the Ezsigntemplatedocument
	IEzsigntemplatedocumentpagePagenumber int32 `json:"iEzsigntemplatedocumentpagePagenumber"`
	// The X coordinate (Horizontal) where to put the Ezsigntemplatesignature on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignature 2 inches from the left border of the page, you would use \"200\" for the X coordinate.
	IEzsigntemplatesignatureX int32 `json:"iEzsigntemplatesignatureX"`
	// The Y coordinate (Vertical) where to put the Ezsigntemplatesignature on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignature 3 inches from the top border of the page, you would use \"300\" for the Y coordinate.
	IEzsigntemplatesignatureY int32 `json:"iEzsigntemplatesignatureY"`
	// The width of the Ezsigntemplatesignature.  Size is calculated at 100dpi (dot per inch). So for example, if you want the Ezsigntemplatesignature to have a width of 2 inches, you would use \"200\" for the iEzsigntemplatesignatureWidth.
	IEzsigntemplatesignatureWidth *int32 `json:"iEzsigntemplatesignatureWidth,omitempty"`
	// The height of the Ezsigntemplatesignature.  Size is calculated at 100dpi (dot per inch). So for example, if you want the Ezsigntemplatesignature to have an height of 2 inches, you would use \"200\" for the iEzsigntemplatesignatureHeight.
	IEzsigntemplatesignatureHeight *int32 `json:"iEzsigntemplatesignatureHeight,omitempty"`
	// The step when the Ezsigntemplatesigner will be invited to sign
	IEzsigntemplatesignatureStep int32 `json:"iEzsigntemplatesignatureStep"`
	EEzsigntemplatesignatureType FieldEEzsigntemplatesignatureType `json:"eEzsigntemplatesignatureType"`
	// A tooltip that will be presented to Ezsigntemplatesigner about the Ezsigntemplatesignature
	TEzsigntemplatesignatureTooltip *string `json:"tEzsigntemplatesignatureTooltip,omitempty"`
	EEzsigntemplatesignatureTooltipposition *FieldEEzsigntemplatesignatureTooltipposition `json:"eEzsigntemplatesignatureTooltipposition,omitempty"`
	EEzsigntemplatesignatureFont *FieldEEzsigntemplatesignatureFont `json:"eEzsigntemplatesignatureFont,omitempty"`
	// Whether the Ezsigntemplatesignature is required or not. This field is relevant only with Ezsigntemplatesignature with eEzsigntemplatesignatureType = Attachments.
	BEzsigntemplatesignatureRequired *bool `json:"bEzsigntemplatesignatureRequired,omitempty"`
	EEzsigntemplatesignatureAttachmentnamesource *FieldEEzsigntemplatesignatureAttachmentnamesource `json:"eEzsigntemplatesignatureAttachmentnamesource,omitempty"`
	// The description attached to the attachment name added in Ezsigntemplatesignature of eEzsigntemplatesignatureType Attachments
	SEzsigntemplatesignatureAttachmentdescription *string `json:"sEzsigntemplatesignatureAttachmentdescription,omitempty"`
	// The step when the Ezsigntemplatesigner will be invited to validate the Ezsigntemplatesignature of eEzsigntemplatesignatureType Attachments
	IEzsigntemplatesignatureValidationstep *int32 `json:"iEzsigntemplatesignatureValidationstep,omitempty"`
	// The maximum length for the value in the Ezsigntemplatesignature  This can only be set if eEzsigntemplatesignatureType is **FieldText** or **FieldTextarea**
	IEzsigntemplatesignatureMaxlength *int32 `json:"iEzsigntemplatesignatureMaxlength,omitempty"`
	// A regular expression to indicate what values are acceptable for the Ezsigntemplatesignature.  This can only be set if eEzsigntemplatesignatureType is **Text** or **Textarea**
	SEzsigntemplatesignatureRegexp *string `json:"sEzsigntemplatesignatureRegexp,omitempty"`
	EEzsigntemplatesignatureTextvalidation *EnumTextvalidation `json:"eEzsigntemplatesignatureTextvalidation,omitempty"`
}

// NewEzsigntemplatesignatureRequest instantiates a new EzsigntemplatesignatureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatesignatureRequest(fkiEzsigntemplatedocumentID int32, fkiEzsigntemplatesignerID int32, iEzsigntemplatedocumentpagePagenumber int32, iEzsigntemplatesignatureX int32, iEzsigntemplatesignatureY int32, iEzsigntemplatesignatureStep int32, eEzsigntemplatesignatureType FieldEEzsigntemplatesignatureType) *EzsigntemplatesignatureRequest {
	this := EzsigntemplatesignatureRequest{}
	this.FkiEzsigntemplatedocumentID = fkiEzsigntemplatedocumentID
	this.FkiEzsigntemplatesignerID = fkiEzsigntemplatesignerID
	this.IEzsigntemplatedocumentpagePagenumber = iEzsigntemplatedocumentpagePagenumber
	this.IEzsigntemplatesignatureX = iEzsigntemplatesignatureX
	this.IEzsigntemplatesignatureY = iEzsigntemplatesignatureY
	this.IEzsigntemplatesignatureStep = iEzsigntemplatesignatureStep
	this.EEzsigntemplatesignatureType = eEzsigntemplatesignatureType
	return &this
}

// NewEzsigntemplatesignatureRequestWithDefaults instantiates a new EzsigntemplatesignatureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatesignatureRequestWithDefaults() *EzsigntemplatesignatureRequest {
	this := EzsigntemplatesignatureRequest{}
	return &this
}

// GetPkiEzsigntemplatesignatureID returns the PkiEzsigntemplatesignatureID field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureRequest) GetPkiEzsigntemplatesignatureID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplatesignatureID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplatesignatureID
}

// GetPkiEzsigntemplatesignatureIDOk returns a tuple with the PkiEzsigntemplatesignatureID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetPkiEzsigntemplatesignatureIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplatesignatureID) {
		return nil, false
	}
	return o.PkiEzsigntemplatesignatureID, true
}

// HasPkiEzsigntemplatesignatureID returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureRequest) HasPkiEzsigntemplatesignatureID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplatesignatureID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplatesignatureID gets a reference to the given int32 and assigns it to the PkiEzsigntemplatesignatureID field.
func (o *EzsigntemplatesignatureRequest) SetPkiEzsigntemplatesignatureID(v int32) {
	o.PkiEzsigntemplatesignatureID = &v
}

// GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field value
func (o *EzsigntemplatesignatureRequest) GetFkiEzsigntemplatedocumentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatedocumentID
}

// GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatedocumentID, true
}

// SetFkiEzsigntemplatedocumentID sets field value
func (o *EzsigntemplatesignatureRequest) SetFkiEzsigntemplatedocumentID(v int32) {
	o.FkiEzsigntemplatedocumentID = v
}

// GetFkiEzsigntemplatesignerID returns the FkiEzsigntemplatesignerID field value
func (o *EzsigntemplatesignatureRequest) GetFkiEzsigntemplatesignerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatesignerID
}

// GetFkiEzsigntemplatesignerIDOk returns a tuple with the FkiEzsigntemplatesignerID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetFkiEzsigntemplatesignerIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatesignerID, true
}

// SetFkiEzsigntemplatesignerID sets field value
func (o *EzsigntemplatesignatureRequest) SetFkiEzsigntemplatesignerID(v int32) {
	o.FkiEzsigntemplatesignerID = v
}

// GetFkiEzsigntemplatesignerIDValidation returns the FkiEzsigntemplatesignerIDValidation field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureRequest) GetFkiEzsigntemplatesignerIDValidation() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplatesignerIDValidation) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplatesignerIDValidation
}

// GetFkiEzsigntemplatesignerIDValidationOk returns a tuple with the FkiEzsigntemplatesignerIDValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetFkiEzsigntemplatesignerIDValidationOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplatesignerIDValidation) {
		return nil, false
	}
	return o.FkiEzsigntemplatesignerIDValidation, true
}

// HasFkiEzsigntemplatesignerIDValidation returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureRequest) HasFkiEzsigntemplatesignerIDValidation() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplatesignerIDValidation) {
		return true
	}

	return false
}

// SetFkiEzsigntemplatesignerIDValidation gets a reference to the given int32 and assigns it to the FkiEzsigntemplatesignerIDValidation field.
func (o *EzsigntemplatesignatureRequest) SetFkiEzsigntemplatesignerIDValidation(v int32) {
	o.FkiEzsigntemplatesignerIDValidation = &v
}

// GetIEzsigntemplatedocumentpagePagenumber returns the IEzsigntemplatedocumentpagePagenumber field value
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatedocumentpagePagenumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatedocumentpagePagenumber
}

// GetIEzsigntemplatedocumentpagePagenumberOk returns a tuple with the IEzsigntemplatedocumentpagePagenumber field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatedocumentpagePagenumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatedocumentpagePagenumber, true
}

// SetIEzsigntemplatedocumentpagePagenumber sets field value
func (o *EzsigntemplatesignatureRequest) SetIEzsigntemplatedocumentpagePagenumber(v int32) {
	o.IEzsigntemplatedocumentpagePagenumber = v
}

// GetIEzsigntemplatesignatureX returns the IEzsigntemplatesignatureX field value
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatesignatureX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatesignatureX
}

// GetIEzsigntemplatesignatureXOk returns a tuple with the IEzsigntemplatesignatureX field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatesignatureXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatesignatureX, true
}

// SetIEzsigntemplatesignatureX sets field value
func (o *EzsigntemplatesignatureRequest) SetIEzsigntemplatesignatureX(v int32) {
	o.IEzsigntemplatesignatureX = v
}

// GetIEzsigntemplatesignatureY returns the IEzsigntemplatesignatureY field value
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatesignatureY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatesignatureY
}

// GetIEzsigntemplatesignatureYOk returns a tuple with the IEzsigntemplatesignatureY field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatesignatureYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatesignatureY, true
}

// SetIEzsigntemplatesignatureY sets field value
func (o *EzsigntemplatesignatureRequest) SetIEzsigntemplatesignatureY(v int32) {
	o.IEzsigntemplatesignatureY = v
}

// GetIEzsigntemplatesignatureWidth returns the IEzsigntemplatesignatureWidth field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatesignatureWidth() int32 {
	if o == nil || IsNil(o.IEzsigntemplatesignatureWidth) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatesignatureWidth
}

// GetIEzsigntemplatesignatureWidthOk returns a tuple with the IEzsigntemplatesignatureWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatesignatureWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatesignatureWidth) {
		return nil, false
	}
	return o.IEzsigntemplatesignatureWidth, true
}

// HasIEzsigntemplatesignatureWidth returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureRequest) HasIEzsigntemplatesignatureWidth() bool {
	if o != nil && !IsNil(o.IEzsigntemplatesignatureWidth) {
		return true
	}

	return false
}

// SetIEzsigntemplatesignatureWidth gets a reference to the given int32 and assigns it to the IEzsigntemplatesignatureWidth field.
func (o *EzsigntemplatesignatureRequest) SetIEzsigntemplatesignatureWidth(v int32) {
	o.IEzsigntemplatesignatureWidth = &v
}

// GetIEzsigntemplatesignatureHeight returns the IEzsigntemplatesignatureHeight field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatesignatureHeight() int32 {
	if o == nil || IsNil(o.IEzsigntemplatesignatureHeight) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatesignatureHeight
}

// GetIEzsigntemplatesignatureHeightOk returns a tuple with the IEzsigntemplatesignatureHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatesignatureHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatesignatureHeight) {
		return nil, false
	}
	return o.IEzsigntemplatesignatureHeight, true
}

// HasIEzsigntemplatesignatureHeight returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureRequest) HasIEzsigntemplatesignatureHeight() bool {
	if o != nil && !IsNil(o.IEzsigntemplatesignatureHeight) {
		return true
	}

	return false
}

// SetIEzsigntemplatesignatureHeight gets a reference to the given int32 and assigns it to the IEzsigntemplatesignatureHeight field.
func (o *EzsigntemplatesignatureRequest) SetIEzsigntemplatesignatureHeight(v int32) {
	o.IEzsigntemplatesignatureHeight = &v
}

// GetIEzsigntemplatesignatureStep returns the IEzsigntemplatesignatureStep field value
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatesignatureStep() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatesignatureStep
}

// GetIEzsigntemplatesignatureStepOk returns a tuple with the IEzsigntemplatesignatureStep field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatesignatureStepOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatesignatureStep, true
}

// SetIEzsigntemplatesignatureStep sets field value
func (o *EzsigntemplatesignatureRequest) SetIEzsigntemplatesignatureStep(v int32) {
	o.IEzsigntemplatesignatureStep = v
}

// GetEEzsigntemplatesignatureType returns the EEzsigntemplatesignatureType field value
func (o *EzsigntemplatesignatureRequest) GetEEzsigntemplatesignatureType() FieldEEzsigntemplatesignatureType {
	if o == nil {
		var ret FieldEEzsigntemplatesignatureType
		return ret
	}

	return o.EEzsigntemplatesignatureType
}

// GetEEzsigntemplatesignatureTypeOk returns a tuple with the EEzsigntemplatesignatureType field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetEEzsigntemplatesignatureTypeOk() (*FieldEEzsigntemplatesignatureType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigntemplatesignatureType, true
}

// SetEEzsigntemplatesignatureType sets field value
func (o *EzsigntemplatesignatureRequest) SetEEzsigntemplatesignatureType(v FieldEEzsigntemplatesignatureType) {
	o.EEzsigntemplatesignatureType = v
}

// GetTEzsigntemplatesignatureTooltip returns the TEzsigntemplatesignatureTooltip field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureRequest) GetTEzsigntemplatesignatureTooltip() string {
	if o == nil || IsNil(o.TEzsigntemplatesignatureTooltip) {
		var ret string
		return ret
	}
	return *o.TEzsigntemplatesignatureTooltip
}

// GetTEzsigntemplatesignatureTooltipOk returns a tuple with the TEzsigntemplatesignatureTooltip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetTEzsigntemplatesignatureTooltipOk() (*string, bool) {
	if o == nil || IsNil(o.TEzsigntemplatesignatureTooltip) {
		return nil, false
	}
	return o.TEzsigntemplatesignatureTooltip, true
}

// HasTEzsigntemplatesignatureTooltip returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureRequest) HasTEzsigntemplatesignatureTooltip() bool {
	if o != nil && !IsNil(o.TEzsigntemplatesignatureTooltip) {
		return true
	}

	return false
}

// SetTEzsigntemplatesignatureTooltip gets a reference to the given string and assigns it to the TEzsigntemplatesignatureTooltip field.
func (o *EzsigntemplatesignatureRequest) SetTEzsigntemplatesignatureTooltip(v string) {
	o.TEzsigntemplatesignatureTooltip = &v
}

// GetEEzsigntemplatesignatureTooltipposition returns the EEzsigntemplatesignatureTooltipposition field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureRequest) GetEEzsigntemplatesignatureTooltipposition() FieldEEzsigntemplatesignatureTooltipposition {
	if o == nil || IsNil(o.EEzsigntemplatesignatureTooltipposition) {
		var ret FieldEEzsigntemplatesignatureTooltipposition
		return ret
	}
	return *o.EEzsigntemplatesignatureTooltipposition
}

// GetEEzsigntemplatesignatureTooltippositionOk returns a tuple with the EEzsigntemplatesignatureTooltipposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetEEzsigntemplatesignatureTooltippositionOk() (*FieldEEzsigntemplatesignatureTooltipposition, bool) {
	if o == nil || IsNil(o.EEzsigntemplatesignatureTooltipposition) {
		return nil, false
	}
	return o.EEzsigntemplatesignatureTooltipposition, true
}

// HasEEzsigntemplatesignatureTooltipposition returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureRequest) HasEEzsigntemplatesignatureTooltipposition() bool {
	if o != nil && !IsNil(o.EEzsigntemplatesignatureTooltipposition) {
		return true
	}

	return false
}

// SetEEzsigntemplatesignatureTooltipposition gets a reference to the given FieldEEzsigntemplatesignatureTooltipposition and assigns it to the EEzsigntemplatesignatureTooltipposition field.
func (o *EzsigntemplatesignatureRequest) SetEEzsigntemplatesignatureTooltipposition(v FieldEEzsigntemplatesignatureTooltipposition) {
	o.EEzsigntemplatesignatureTooltipposition = &v
}

// GetEEzsigntemplatesignatureFont returns the EEzsigntemplatesignatureFont field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureRequest) GetEEzsigntemplatesignatureFont() FieldEEzsigntemplatesignatureFont {
	if o == nil || IsNil(o.EEzsigntemplatesignatureFont) {
		var ret FieldEEzsigntemplatesignatureFont
		return ret
	}
	return *o.EEzsigntemplatesignatureFont
}

// GetEEzsigntemplatesignatureFontOk returns a tuple with the EEzsigntemplatesignatureFont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetEEzsigntemplatesignatureFontOk() (*FieldEEzsigntemplatesignatureFont, bool) {
	if o == nil || IsNil(o.EEzsigntemplatesignatureFont) {
		return nil, false
	}
	return o.EEzsigntemplatesignatureFont, true
}

// HasEEzsigntemplatesignatureFont returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureRequest) HasEEzsigntemplatesignatureFont() bool {
	if o != nil && !IsNil(o.EEzsigntemplatesignatureFont) {
		return true
	}

	return false
}

// SetEEzsigntemplatesignatureFont gets a reference to the given FieldEEzsigntemplatesignatureFont and assigns it to the EEzsigntemplatesignatureFont field.
func (o *EzsigntemplatesignatureRequest) SetEEzsigntemplatesignatureFont(v FieldEEzsigntemplatesignatureFont) {
	o.EEzsigntemplatesignatureFont = &v
}

// GetBEzsigntemplatesignatureRequired returns the BEzsigntemplatesignatureRequired field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureRequest) GetBEzsigntemplatesignatureRequired() bool {
	if o == nil || IsNil(o.BEzsigntemplatesignatureRequired) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplatesignatureRequired
}

// GetBEzsigntemplatesignatureRequiredOk returns a tuple with the BEzsigntemplatesignatureRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetBEzsigntemplatesignatureRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplatesignatureRequired) {
		return nil, false
	}
	return o.BEzsigntemplatesignatureRequired, true
}

// HasBEzsigntemplatesignatureRequired returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureRequest) HasBEzsigntemplatesignatureRequired() bool {
	if o != nil && !IsNil(o.BEzsigntemplatesignatureRequired) {
		return true
	}

	return false
}

// SetBEzsigntemplatesignatureRequired gets a reference to the given bool and assigns it to the BEzsigntemplatesignatureRequired field.
func (o *EzsigntemplatesignatureRequest) SetBEzsigntemplatesignatureRequired(v bool) {
	o.BEzsigntemplatesignatureRequired = &v
}

// GetEEzsigntemplatesignatureAttachmentnamesource returns the EEzsigntemplatesignatureAttachmentnamesource field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureRequest) GetEEzsigntemplatesignatureAttachmentnamesource() FieldEEzsigntemplatesignatureAttachmentnamesource {
	if o == nil || IsNil(o.EEzsigntemplatesignatureAttachmentnamesource) {
		var ret FieldEEzsigntemplatesignatureAttachmentnamesource
		return ret
	}
	return *o.EEzsigntemplatesignatureAttachmentnamesource
}

// GetEEzsigntemplatesignatureAttachmentnamesourceOk returns a tuple with the EEzsigntemplatesignatureAttachmentnamesource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetEEzsigntemplatesignatureAttachmentnamesourceOk() (*FieldEEzsigntemplatesignatureAttachmentnamesource, bool) {
	if o == nil || IsNil(o.EEzsigntemplatesignatureAttachmentnamesource) {
		return nil, false
	}
	return o.EEzsigntemplatesignatureAttachmentnamesource, true
}

// HasEEzsigntemplatesignatureAttachmentnamesource returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureRequest) HasEEzsigntemplatesignatureAttachmentnamesource() bool {
	if o != nil && !IsNil(o.EEzsigntemplatesignatureAttachmentnamesource) {
		return true
	}

	return false
}

// SetEEzsigntemplatesignatureAttachmentnamesource gets a reference to the given FieldEEzsigntemplatesignatureAttachmentnamesource and assigns it to the EEzsigntemplatesignatureAttachmentnamesource field.
func (o *EzsigntemplatesignatureRequest) SetEEzsigntemplatesignatureAttachmentnamesource(v FieldEEzsigntemplatesignatureAttachmentnamesource) {
	o.EEzsigntemplatesignatureAttachmentnamesource = &v
}

// GetSEzsigntemplatesignatureAttachmentdescription returns the SEzsigntemplatesignatureAttachmentdescription field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureRequest) GetSEzsigntemplatesignatureAttachmentdescription() string {
	if o == nil || IsNil(o.SEzsigntemplatesignatureAttachmentdescription) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplatesignatureAttachmentdescription
}

// GetSEzsigntemplatesignatureAttachmentdescriptionOk returns a tuple with the SEzsigntemplatesignatureAttachmentdescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetSEzsigntemplatesignatureAttachmentdescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplatesignatureAttachmentdescription) {
		return nil, false
	}
	return o.SEzsigntemplatesignatureAttachmentdescription, true
}

// HasSEzsigntemplatesignatureAttachmentdescription returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureRequest) HasSEzsigntemplatesignatureAttachmentdescription() bool {
	if o != nil && !IsNil(o.SEzsigntemplatesignatureAttachmentdescription) {
		return true
	}

	return false
}

// SetSEzsigntemplatesignatureAttachmentdescription gets a reference to the given string and assigns it to the SEzsigntemplatesignatureAttachmentdescription field.
func (o *EzsigntemplatesignatureRequest) SetSEzsigntemplatesignatureAttachmentdescription(v string) {
	o.SEzsigntemplatesignatureAttachmentdescription = &v
}

// GetIEzsigntemplatesignatureValidationstep returns the IEzsigntemplatesignatureValidationstep field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatesignatureValidationstep() int32 {
	if o == nil || IsNil(o.IEzsigntemplatesignatureValidationstep) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatesignatureValidationstep
}

// GetIEzsigntemplatesignatureValidationstepOk returns a tuple with the IEzsigntemplatesignatureValidationstep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatesignatureValidationstepOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatesignatureValidationstep) {
		return nil, false
	}
	return o.IEzsigntemplatesignatureValidationstep, true
}

// HasIEzsigntemplatesignatureValidationstep returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureRequest) HasIEzsigntemplatesignatureValidationstep() bool {
	if o != nil && !IsNil(o.IEzsigntemplatesignatureValidationstep) {
		return true
	}

	return false
}

// SetIEzsigntemplatesignatureValidationstep gets a reference to the given int32 and assigns it to the IEzsigntemplatesignatureValidationstep field.
func (o *EzsigntemplatesignatureRequest) SetIEzsigntemplatesignatureValidationstep(v int32) {
	o.IEzsigntemplatesignatureValidationstep = &v
}

// GetIEzsigntemplatesignatureMaxlength returns the IEzsigntemplatesignatureMaxlength field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatesignatureMaxlength() int32 {
	if o == nil || IsNil(o.IEzsigntemplatesignatureMaxlength) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatesignatureMaxlength
}

// GetIEzsigntemplatesignatureMaxlengthOk returns a tuple with the IEzsigntemplatesignatureMaxlength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetIEzsigntemplatesignatureMaxlengthOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatesignatureMaxlength) {
		return nil, false
	}
	return o.IEzsigntemplatesignatureMaxlength, true
}

// HasIEzsigntemplatesignatureMaxlength returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureRequest) HasIEzsigntemplatesignatureMaxlength() bool {
	if o != nil && !IsNil(o.IEzsigntemplatesignatureMaxlength) {
		return true
	}

	return false
}

// SetIEzsigntemplatesignatureMaxlength gets a reference to the given int32 and assigns it to the IEzsigntemplatesignatureMaxlength field.
func (o *EzsigntemplatesignatureRequest) SetIEzsigntemplatesignatureMaxlength(v int32) {
	o.IEzsigntemplatesignatureMaxlength = &v
}

// GetSEzsigntemplatesignatureRegexp returns the SEzsigntemplatesignatureRegexp field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureRequest) GetSEzsigntemplatesignatureRegexp() string {
	if o == nil || IsNil(o.SEzsigntemplatesignatureRegexp) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplatesignatureRegexp
}

// GetSEzsigntemplatesignatureRegexpOk returns a tuple with the SEzsigntemplatesignatureRegexp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetSEzsigntemplatesignatureRegexpOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplatesignatureRegexp) {
		return nil, false
	}
	return o.SEzsigntemplatesignatureRegexp, true
}

// HasSEzsigntemplatesignatureRegexp returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureRequest) HasSEzsigntemplatesignatureRegexp() bool {
	if o != nil && !IsNil(o.SEzsigntemplatesignatureRegexp) {
		return true
	}

	return false
}

// SetSEzsigntemplatesignatureRegexp gets a reference to the given string and assigns it to the SEzsigntemplatesignatureRegexp field.
func (o *EzsigntemplatesignatureRequest) SetSEzsigntemplatesignatureRegexp(v string) {
	o.SEzsigntemplatesignatureRegexp = &v
}

// GetEEzsigntemplatesignatureTextvalidation returns the EEzsigntemplatesignatureTextvalidation field value if set, zero value otherwise.
func (o *EzsigntemplatesignatureRequest) GetEEzsigntemplatesignatureTextvalidation() EnumTextvalidation {
	if o == nil || IsNil(o.EEzsigntemplatesignatureTextvalidation) {
		var ret EnumTextvalidation
		return ret
	}
	return *o.EEzsigntemplatesignatureTextvalidation
}

// GetEEzsigntemplatesignatureTextvalidationOk returns a tuple with the EEzsigntemplatesignatureTextvalidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignatureRequest) GetEEzsigntemplatesignatureTextvalidationOk() (*EnumTextvalidation, bool) {
	if o == nil || IsNil(o.EEzsigntemplatesignatureTextvalidation) {
		return nil, false
	}
	return o.EEzsigntemplatesignatureTextvalidation, true
}

// HasEEzsigntemplatesignatureTextvalidation returns a boolean if a field has been set.
func (o *EzsigntemplatesignatureRequest) HasEEzsigntemplatesignatureTextvalidation() bool {
	if o != nil && !IsNil(o.EEzsigntemplatesignatureTextvalidation) {
		return true
	}

	return false
}

// SetEEzsigntemplatesignatureTextvalidation gets a reference to the given EnumTextvalidation and assigns it to the EEzsigntemplatesignatureTextvalidation field.
func (o *EzsigntemplatesignatureRequest) SetEEzsigntemplatesignatureTextvalidation(v EnumTextvalidation) {
	o.EEzsigntemplatesignatureTextvalidation = &v
}

func (o EzsigntemplatesignatureRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatesignatureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplatesignatureID) {
		toSerialize["pkiEzsigntemplatesignatureID"] = o.PkiEzsigntemplatesignatureID
	}
	toSerialize["fkiEzsigntemplatedocumentID"] = o.FkiEzsigntemplatedocumentID
	toSerialize["fkiEzsigntemplatesignerID"] = o.FkiEzsigntemplatesignerID
	if !IsNil(o.FkiEzsigntemplatesignerIDValidation) {
		toSerialize["fkiEzsigntemplatesignerIDValidation"] = o.FkiEzsigntemplatesignerIDValidation
	}
	toSerialize["iEzsigntemplatedocumentpagePagenumber"] = o.IEzsigntemplatedocumentpagePagenumber
	toSerialize["iEzsigntemplatesignatureX"] = o.IEzsigntemplatesignatureX
	toSerialize["iEzsigntemplatesignatureY"] = o.IEzsigntemplatesignatureY
	if !IsNil(o.IEzsigntemplatesignatureWidth) {
		toSerialize["iEzsigntemplatesignatureWidth"] = o.IEzsigntemplatesignatureWidth
	}
	if !IsNil(o.IEzsigntemplatesignatureHeight) {
		toSerialize["iEzsigntemplatesignatureHeight"] = o.IEzsigntemplatesignatureHeight
	}
	toSerialize["iEzsigntemplatesignatureStep"] = o.IEzsigntemplatesignatureStep
	toSerialize["eEzsigntemplatesignatureType"] = o.EEzsigntemplatesignatureType
	if !IsNil(o.TEzsigntemplatesignatureTooltip) {
		toSerialize["tEzsigntemplatesignatureTooltip"] = o.TEzsigntemplatesignatureTooltip
	}
	if !IsNil(o.EEzsigntemplatesignatureTooltipposition) {
		toSerialize["eEzsigntemplatesignatureTooltipposition"] = o.EEzsigntemplatesignatureTooltipposition
	}
	if !IsNil(o.EEzsigntemplatesignatureFont) {
		toSerialize["eEzsigntemplatesignatureFont"] = o.EEzsigntemplatesignatureFont
	}
	if !IsNil(o.BEzsigntemplatesignatureRequired) {
		toSerialize["bEzsigntemplatesignatureRequired"] = o.BEzsigntemplatesignatureRequired
	}
	if !IsNil(o.EEzsigntemplatesignatureAttachmentnamesource) {
		toSerialize["eEzsigntemplatesignatureAttachmentnamesource"] = o.EEzsigntemplatesignatureAttachmentnamesource
	}
	if !IsNil(o.SEzsigntemplatesignatureAttachmentdescription) {
		toSerialize["sEzsigntemplatesignatureAttachmentdescription"] = o.SEzsigntemplatesignatureAttachmentdescription
	}
	if !IsNil(o.IEzsigntemplatesignatureValidationstep) {
		toSerialize["iEzsigntemplatesignatureValidationstep"] = o.IEzsigntemplatesignatureValidationstep
	}
	if !IsNil(o.IEzsigntemplatesignatureMaxlength) {
		toSerialize["iEzsigntemplatesignatureMaxlength"] = o.IEzsigntemplatesignatureMaxlength
	}
	if !IsNil(o.SEzsigntemplatesignatureRegexp) {
		toSerialize["sEzsigntemplatesignatureRegexp"] = o.SEzsigntemplatesignatureRegexp
	}
	if !IsNil(o.EEzsigntemplatesignatureTextvalidation) {
		toSerialize["eEzsigntemplatesignatureTextvalidation"] = o.EEzsigntemplatesignatureTextvalidation
	}
	return toSerialize, nil
}

type NullableEzsigntemplatesignatureRequest struct {
	value *EzsigntemplatesignatureRequest
	isSet bool
}

func (v NullableEzsigntemplatesignatureRequest) Get() *EzsigntemplatesignatureRequest {
	return v.value
}

func (v *NullableEzsigntemplatesignatureRequest) Set(val *EzsigntemplatesignatureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatesignatureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatesignatureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatesignatureRequest(val *EzsigntemplatesignatureRequest) *NullableEzsigntemplatesignatureRequest {
	return &NullableEzsigntemplatesignatureRequest{value: val, isSet: true}
}

func (v NullableEzsigntemplatesignatureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatesignatureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


