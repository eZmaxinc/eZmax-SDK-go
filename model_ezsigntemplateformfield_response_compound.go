/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EzsigntemplateformfieldResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateformfieldResponseCompound{}

// EzsigntemplateformfieldResponseCompound An Ezsigntemplateformfield Object and children
type EzsigntemplateformfieldResponseCompound struct {
	// The unique ID of the Ezsigntemplateformfield
	PkiEzsigntemplateformfieldID int32 `json:"pkiEzsigntemplateformfieldID"`
	EEzsigntemplateformfieldPositioning *FieldEEzsigntemplateformfieldPositioning `json:"eEzsigntemplateformfieldPositioning,omitempty"`
	// The page number in the Ezsigntemplatedocument
	IEzsigntemplatedocumentpagePagenumber int32 `json:"iEzsigntemplatedocumentpagePagenumber"`
	// The Label for the Ezsigntemplateformfield
	SEzsigntemplateformfieldLabel string `json:"sEzsigntemplateformfieldLabel"`
	// The value for the Ezsigntemplateformfield
	SEzsigntemplateformfieldValue *string `json:"sEzsigntemplateformfieldValue,omitempty"`
	// The X coordinate (Horizontal) where to put the Ezsigntemplateformfield on the Ezsigntemplatepage.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplateformfield 2 inches from the left border of the page, you would use \"200\" for the X coordinate.
	IEzsigntemplateformfieldX *int32 `json:"iEzsigntemplateformfieldX,omitempty"`
	// The Y coordinate (Vertical) where to put the Ezsigntemplateformfield on the Ezsigntemplatepage.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplateformfield 3 inches from the top border of the page, you would use \"300\" for the Y coordinate.
	IEzsigntemplateformfieldY *int32 `json:"iEzsigntemplateformfieldY,omitempty"`
	// The Width of the Ezsigntemplateformfield in pixels calculated at 100 DPI
	IEzsigntemplateformfieldWidth int32 `json:"iEzsigntemplateformfieldWidth"`
	// The Height of the Ezsigntemplateformfield in pixels calculated at 100 DPI 
	IEzsigntemplateformfieldHeight int32 `json:"iEzsigntemplateformfieldHeight"`
	// Whether the Ezsigntemplateformfield allows the use of the autocomplete of the browser.  This can only be set if eEzsigntemplateformfieldgroupType is **Text**
	BEzsigntemplateformfieldAutocomplete *bool `json:"bEzsigntemplateformfieldAutocomplete,omitempty"`
	// Whether the Ezsigntemplateformfield is selected or not by default.  This can only be set if eEzsigntemplateformfieldgroupType is **Checkbox** or **Radio**
	BEzsigntemplateformfieldSelected *bool `json:"bEzsigntemplateformfieldSelected,omitempty"`
	EEzsigntemplateformfieldDependencyrequirement *FieldEEzsigntemplateformfieldDependencyrequirement `json:"eEzsigntemplateformfieldDependencyrequirement,omitempty"`
	// The string pattern to search for the positioning. **This is not a regexp**  This will be required if **eEzsigntemplateformfieldPositioning** is set to **PerCoordinates**
	SEzsigntemplateformfieldPositioningpattern *string `json:"sEzsigntemplateformfieldPositioningpattern,omitempty" validate:"regexp=^.{0,30}$"`
	// The offset X  This will be required if **eEzsigntemplateformfieldPositioning** is set to **PerCoordinates**
	IEzsigntemplateformfieldPositioningoffsetx *int32 `json:"iEzsigntemplateformfieldPositioningoffsetx,omitempty"`
	// The offset Y  This will be required if **eEzsigntemplateformfieldPositioning** is set to **PerCoordinates**
	IEzsigntemplateformfieldPositioningoffsety *int32 `json:"iEzsigntemplateformfieldPositioningoffsety,omitempty"`
	EEzsigntemplateformfieldPositioningoccurence *FieldEEzsigntemplateformfieldPositioningoccurence `json:"eEzsigntemplateformfieldPositioningoccurence,omitempty"`
	EEzsigntemplateformfieldHorizontalalignment *EnumHorizontalalignment `json:"eEzsigntemplateformfieldHorizontalalignment,omitempty"`
	ObjTextstylestatic *TextstylestaticResponseCompound `json:"objTextstylestatic,omitempty"`
	AObjEzsigntemplateelementdependency []EzsigntemplateelementdependencyResponseCompound `json:"a_objEzsigntemplateelementdependency,omitempty"`
}

type _EzsigntemplateformfieldResponseCompound EzsigntemplateformfieldResponseCompound

// NewEzsigntemplateformfieldResponseCompound instantiates a new EzsigntemplateformfieldResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateformfieldResponseCompound(pkiEzsigntemplateformfieldID int32, iEzsigntemplatedocumentpagePagenumber int32, sEzsigntemplateformfieldLabel string, iEzsigntemplateformfieldWidth int32, iEzsigntemplateformfieldHeight int32) *EzsigntemplateformfieldResponseCompound {
	this := EzsigntemplateformfieldResponseCompound{}
	this.PkiEzsigntemplateformfieldID = pkiEzsigntemplateformfieldID
	var eEzsigntemplateformfieldPositioning FieldEEzsigntemplateformfieldPositioning = PER_COORDINATES
	this.EEzsigntemplateformfieldPositioning = &eEzsigntemplateformfieldPositioning
	this.IEzsigntemplatedocumentpagePagenumber = iEzsigntemplatedocumentpagePagenumber
	this.SEzsigntemplateformfieldLabel = sEzsigntemplateformfieldLabel
	this.IEzsigntemplateformfieldWidth = iEzsigntemplateformfieldWidth
	this.IEzsigntemplateformfieldHeight = iEzsigntemplateformfieldHeight
	return &this
}

// NewEzsigntemplateformfieldResponseCompoundWithDefaults instantiates a new EzsigntemplateformfieldResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateformfieldResponseCompoundWithDefaults() *EzsigntemplateformfieldResponseCompound {
	this := EzsigntemplateformfieldResponseCompound{}
	var eEzsigntemplateformfieldPositioning FieldEEzsigntemplateformfieldPositioning = PER_COORDINATES
	this.EEzsigntemplateformfieldPositioning = &eEzsigntemplateformfieldPositioning
	return &this
}

// GetPkiEzsigntemplateformfieldID returns the PkiEzsigntemplateformfieldID field value
func (o *EzsigntemplateformfieldResponseCompound) GetPkiEzsigntemplateformfieldID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplateformfieldID
}

// GetPkiEzsigntemplateformfieldIDOk returns a tuple with the PkiEzsigntemplateformfieldID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetPkiEzsigntemplateformfieldIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplateformfieldID, true
}

// SetPkiEzsigntemplateformfieldID sets field value
func (o *EzsigntemplateformfieldResponseCompound) SetPkiEzsigntemplateformfieldID(v int32) {
	o.PkiEzsigntemplateformfieldID = v
}

// GetEEzsigntemplateformfieldPositioning returns the EEzsigntemplateformfieldPositioning field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldResponseCompound) GetEEzsigntemplateformfieldPositioning() FieldEEzsigntemplateformfieldPositioning {
	if o == nil || IsNil(o.EEzsigntemplateformfieldPositioning) {
		var ret FieldEEzsigntemplateformfieldPositioning
		return ret
	}
	return *o.EEzsigntemplateformfieldPositioning
}

// GetEEzsigntemplateformfieldPositioningOk returns a tuple with the EEzsigntemplateformfieldPositioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetEEzsigntemplateformfieldPositioningOk() (*FieldEEzsigntemplateformfieldPositioning, bool) {
	if o == nil || IsNil(o.EEzsigntemplateformfieldPositioning) {
		return nil, false
	}
	return o.EEzsigntemplateformfieldPositioning, true
}

// HasEEzsigntemplateformfieldPositioning returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldResponseCompound) HasEEzsigntemplateformfieldPositioning() bool {
	if o != nil && !IsNil(o.EEzsigntemplateformfieldPositioning) {
		return true
	}

	return false
}

// SetEEzsigntemplateformfieldPositioning gets a reference to the given FieldEEzsigntemplateformfieldPositioning and assigns it to the EEzsigntemplateformfieldPositioning field.
func (o *EzsigntemplateformfieldResponseCompound) SetEEzsigntemplateformfieldPositioning(v FieldEEzsigntemplateformfieldPositioning) {
	o.EEzsigntemplateformfieldPositioning = &v
}

// GetIEzsigntemplatedocumentpagePagenumber returns the IEzsigntemplatedocumentpagePagenumber field value
func (o *EzsigntemplateformfieldResponseCompound) GetIEzsigntemplatedocumentpagePagenumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatedocumentpagePagenumber
}

// GetIEzsigntemplatedocumentpagePagenumberOk returns a tuple with the IEzsigntemplatedocumentpagePagenumber field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetIEzsigntemplatedocumentpagePagenumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatedocumentpagePagenumber, true
}

// SetIEzsigntemplatedocumentpagePagenumber sets field value
func (o *EzsigntemplateformfieldResponseCompound) SetIEzsigntemplatedocumentpagePagenumber(v int32) {
	o.IEzsigntemplatedocumentpagePagenumber = v
}

// GetSEzsigntemplateformfieldLabel returns the SEzsigntemplateformfieldLabel field value
func (o *EzsigntemplateformfieldResponseCompound) GetSEzsigntemplateformfieldLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateformfieldLabel
}

// GetSEzsigntemplateformfieldLabelOk returns a tuple with the SEzsigntemplateformfieldLabel field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetSEzsigntemplateformfieldLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateformfieldLabel, true
}

// SetSEzsigntemplateformfieldLabel sets field value
func (o *EzsigntemplateformfieldResponseCompound) SetSEzsigntemplateformfieldLabel(v string) {
	o.SEzsigntemplateformfieldLabel = v
}

// GetSEzsigntemplateformfieldValue returns the SEzsigntemplateformfieldValue field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldResponseCompound) GetSEzsigntemplateformfieldValue() string {
	if o == nil || IsNil(o.SEzsigntemplateformfieldValue) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateformfieldValue
}

// GetSEzsigntemplateformfieldValueOk returns a tuple with the SEzsigntemplateformfieldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetSEzsigntemplateformfieldValueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateformfieldValue) {
		return nil, false
	}
	return o.SEzsigntemplateformfieldValue, true
}

// HasSEzsigntemplateformfieldValue returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldResponseCompound) HasSEzsigntemplateformfieldValue() bool {
	if o != nil && !IsNil(o.SEzsigntemplateformfieldValue) {
		return true
	}

	return false
}

// SetSEzsigntemplateformfieldValue gets a reference to the given string and assigns it to the SEzsigntemplateformfieldValue field.
func (o *EzsigntemplateformfieldResponseCompound) SetSEzsigntemplateformfieldValue(v string) {
	o.SEzsigntemplateformfieldValue = &v
}

// GetIEzsigntemplateformfieldX returns the IEzsigntemplateformfieldX field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldResponseCompound) GetIEzsigntemplateformfieldX() int32 {
	if o == nil || IsNil(o.IEzsigntemplateformfieldX) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplateformfieldX
}

// GetIEzsigntemplateformfieldXOk returns a tuple with the IEzsigntemplateformfieldX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetIEzsigntemplateformfieldXOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplateformfieldX) {
		return nil, false
	}
	return o.IEzsigntemplateformfieldX, true
}

// HasIEzsigntemplateformfieldX returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldResponseCompound) HasIEzsigntemplateformfieldX() bool {
	if o != nil && !IsNil(o.IEzsigntemplateformfieldX) {
		return true
	}

	return false
}

// SetIEzsigntemplateformfieldX gets a reference to the given int32 and assigns it to the IEzsigntemplateformfieldX field.
func (o *EzsigntemplateformfieldResponseCompound) SetIEzsigntemplateformfieldX(v int32) {
	o.IEzsigntemplateformfieldX = &v
}

// GetIEzsigntemplateformfieldY returns the IEzsigntemplateformfieldY field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldResponseCompound) GetIEzsigntemplateformfieldY() int32 {
	if o == nil || IsNil(o.IEzsigntemplateformfieldY) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplateformfieldY
}

// GetIEzsigntemplateformfieldYOk returns a tuple with the IEzsigntemplateformfieldY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetIEzsigntemplateformfieldYOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplateformfieldY) {
		return nil, false
	}
	return o.IEzsigntemplateformfieldY, true
}

// HasIEzsigntemplateformfieldY returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldResponseCompound) HasIEzsigntemplateformfieldY() bool {
	if o != nil && !IsNil(o.IEzsigntemplateformfieldY) {
		return true
	}

	return false
}

// SetIEzsigntemplateformfieldY gets a reference to the given int32 and assigns it to the IEzsigntemplateformfieldY field.
func (o *EzsigntemplateformfieldResponseCompound) SetIEzsigntemplateformfieldY(v int32) {
	o.IEzsigntemplateformfieldY = &v
}

// GetIEzsigntemplateformfieldWidth returns the IEzsigntemplateformfieldWidth field value
func (o *EzsigntemplateformfieldResponseCompound) GetIEzsigntemplateformfieldWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplateformfieldWidth
}

// GetIEzsigntemplateformfieldWidthOk returns a tuple with the IEzsigntemplateformfieldWidth field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetIEzsigntemplateformfieldWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplateformfieldWidth, true
}

// SetIEzsigntemplateformfieldWidth sets field value
func (o *EzsigntemplateformfieldResponseCompound) SetIEzsigntemplateformfieldWidth(v int32) {
	o.IEzsigntemplateformfieldWidth = v
}

// GetIEzsigntemplateformfieldHeight returns the IEzsigntemplateformfieldHeight field value
func (o *EzsigntemplateformfieldResponseCompound) GetIEzsigntemplateformfieldHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplateformfieldHeight
}

// GetIEzsigntemplateformfieldHeightOk returns a tuple with the IEzsigntemplateformfieldHeight field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetIEzsigntemplateformfieldHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplateformfieldHeight, true
}

// SetIEzsigntemplateformfieldHeight sets field value
func (o *EzsigntemplateformfieldResponseCompound) SetIEzsigntemplateformfieldHeight(v int32) {
	o.IEzsigntemplateformfieldHeight = v
}

// GetBEzsigntemplateformfieldAutocomplete returns the BEzsigntemplateformfieldAutocomplete field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldResponseCompound) GetBEzsigntemplateformfieldAutocomplete() bool {
	if o == nil || IsNil(o.BEzsigntemplateformfieldAutocomplete) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplateformfieldAutocomplete
}

// GetBEzsigntemplateformfieldAutocompleteOk returns a tuple with the BEzsigntemplateformfieldAutocomplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetBEzsigntemplateformfieldAutocompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplateformfieldAutocomplete) {
		return nil, false
	}
	return o.BEzsigntemplateformfieldAutocomplete, true
}

// HasBEzsigntemplateformfieldAutocomplete returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldResponseCompound) HasBEzsigntemplateformfieldAutocomplete() bool {
	if o != nil && !IsNil(o.BEzsigntemplateformfieldAutocomplete) {
		return true
	}

	return false
}

// SetBEzsigntemplateformfieldAutocomplete gets a reference to the given bool and assigns it to the BEzsigntemplateformfieldAutocomplete field.
func (o *EzsigntemplateformfieldResponseCompound) SetBEzsigntemplateformfieldAutocomplete(v bool) {
	o.BEzsigntemplateformfieldAutocomplete = &v
}

// GetBEzsigntemplateformfieldSelected returns the BEzsigntemplateformfieldSelected field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldResponseCompound) GetBEzsigntemplateformfieldSelected() bool {
	if o == nil || IsNil(o.BEzsigntemplateformfieldSelected) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplateformfieldSelected
}

// GetBEzsigntemplateformfieldSelectedOk returns a tuple with the BEzsigntemplateformfieldSelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetBEzsigntemplateformfieldSelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplateformfieldSelected) {
		return nil, false
	}
	return o.BEzsigntemplateformfieldSelected, true
}

// HasBEzsigntemplateformfieldSelected returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldResponseCompound) HasBEzsigntemplateformfieldSelected() bool {
	if o != nil && !IsNil(o.BEzsigntemplateformfieldSelected) {
		return true
	}

	return false
}

// SetBEzsigntemplateformfieldSelected gets a reference to the given bool and assigns it to the BEzsigntemplateformfieldSelected field.
func (o *EzsigntemplateformfieldResponseCompound) SetBEzsigntemplateformfieldSelected(v bool) {
	o.BEzsigntemplateformfieldSelected = &v
}

// GetEEzsigntemplateformfieldDependencyrequirement returns the EEzsigntemplateformfieldDependencyrequirement field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldResponseCompound) GetEEzsigntemplateformfieldDependencyrequirement() FieldEEzsigntemplateformfieldDependencyrequirement {
	if o == nil || IsNil(o.EEzsigntemplateformfieldDependencyrequirement) {
		var ret FieldEEzsigntemplateformfieldDependencyrequirement
		return ret
	}
	return *o.EEzsigntemplateformfieldDependencyrequirement
}

// GetEEzsigntemplateformfieldDependencyrequirementOk returns a tuple with the EEzsigntemplateformfieldDependencyrequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetEEzsigntemplateformfieldDependencyrequirementOk() (*FieldEEzsigntemplateformfieldDependencyrequirement, bool) {
	if o == nil || IsNil(o.EEzsigntemplateformfieldDependencyrequirement) {
		return nil, false
	}
	return o.EEzsigntemplateformfieldDependencyrequirement, true
}

// HasEEzsigntemplateformfieldDependencyrequirement returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldResponseCompound) HasEEzsigntemplateformfieldDependencyrequirement() bool {
	if o != nil && !IsNil(o.EEzsigntemplateformfieldDependencyrequirement) {
		return true
	}

	return false
}

// SetEEzsigntemplateformfieldDependencyrequirement gets a reference to the given FieldEEzsigntemplateformfieldDependencyrequirement and assigns it to the EEzsigntemplateformfieldDependencyrequirement field.
func (o *EzsigntemplateformfieldResponseCompound) SetEEzsigntemplateformfieldDependencyrequirement(v FieldEEzsigntemplateformfieldDependencyrequirement) {
	o.EEzsigntemplateformfieldDependencyrequirement = &v
}

// GetSEzsigntemplateformfieldPositioningpattern returns the SEzsigntemplateformfieldPositioningpattern field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldResponseCompound) GetSEzsigntemplateformfieldPositioningpattern() string {
	if o == nil || IsNil(o.SEzsigntemplateformfieldPositioningpattern) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateformfieldPositioningpattern
}

// GetSEzsigntemplateformfieldPositioningpatternOk returns a tuple with the SEzsigntemplateformfieldPositioningpattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetSEzsigntemplateformfieldPositioningpatternOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateformfieldPositioningpattern) {
		return nil, false
	}
	return o.SEzsigntemplateformfieldPositioningpattern, true
}

// HasSEzsigntemplateformfieldPositioningpattern returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldResponseCompound) HasSEzsigntemplateformfieldPositioningpattern() bool {
	if o != nil && !IsNil(o.SEzsigntemplateformfieldPositioningpattern) {
		return true
	}

	return false
}

// SetSEzsigntemplateformfieldPositioningpattern gets a reference to the given string and assigns it to the SEzsigntemplateformfieldPositioningpattern field.
func (o *EzsigntemplateformfieldResponseCompound) SetSEzsigntemplateformfieldPositioningpattern(v string) {
	o.SEzsigntemplateformfieldPositioningpattern = &v
}

// GetIEzsigntemplateformfieldPositioningoffsetx returns the IEzsigntemplateformfieldPositioningoffsetx field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldResponseCompound) GetIEzsigntemplateformfieldPositioningoffsetx() int32 {
	if o == nil || IsNil(o.IEzsigntemplateformfieldPositioningoffsetx) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplateformfieldPositioningoffsetx
}

// GetIEzsigntemplateformfieldPositioningoffsetxOk returns a tuple with the IEzsigntemplateformfieldPositioningoffsetx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetIEzsigntemplateformfieldPositioningoffsetxOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplateformfieldPositioningoffsetx) {
		return nil, false
	}
	return o.IEzsigntemplateformfieldPositioningoffsetx, true
}

// HasIEzsigntemplateformfieldPositioningoffsetx returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldResponseCompound) HasIEzsigntemplateformfieldPositioningoffsetx() bool {
	if o != nil && !IsNil(o.IEzsigntemplateformfieldPositioningoffsetx) {
		return true
	}

	return false
}

// SetIEzsigntemplateformfieldPositioningoffsetx gets a reference to the given int32 and assigns it to the IEzsigntemplateformfieldPositioningoffsetx field.
func (o *EzsigntemplateformfieldResponseCompound) SetIEzsigntemplateformfieldPositioningoffsetx(v int32) {
	o.IEzsigntemplateformfieldPositioningoffsetx = &v
}

// GetIEzsigntemplateformfieldPositioningoffsety returns the IEzsigntemplateformfieldPositioningoffsety field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldResponseCompound) GetIEzsigntemplateformfieldPositioningoffsety() int32 {
	if o == nil || IsNil(o.IEzsigntemplateformfieldPositioningoffsety) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplateformfieldPositioningoffsety
}

// GetIEzsigntemplateformfieldPositioningoffsetyOk returns a tuple with the IEzsigntemplateformfieldPositioningoffsety field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetIEzsigntemplateformfieldPositioningoffsetyOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplateformfieldPositioningoffsety) {
		return nil, false
	}
	return o.IEzsigntemplateformfieldPositioningoffsety, true
}

// HasIEzsigntemplateformfieldPositioningoffsety returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldResponseCompound) HasIEzsigntemplateformfieldPositioningoffsety() bool {
	if o != nil && !IsNil(o.IEzsigntemplateformfieldPositioningoffsety) {
		return true
	}

	return false
}

// SetIEzsigntemplateformfieldPositioningoffsety gets a reference to the given int32 and assigns it to the IEzsigntemplateformfieldPositioningoffsety field.
func (o *EzsigntemplateformfieldResponseCompound) SetIEzsigntemplateformfieldPositioningoffsety(v int32) {
	o.IEzsigntemplateformfieldPositioningoffsety = &v
}

// GetEEzsigntemplateformfieldPositioningoccurence returns the EEzsigntemplateformfieldPositioningoccurence field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldResponseCompound) GetEEzsigntemplateformfieldPositioningoccurence() FieldEEzsigntemplateformfieldPositioningoccurence {
	if o == nil || IsNil(o.EEzsigntemplateformfieldPositioningoccurence) {
		var ret FieldEEzsigntemplateformfieldPositioningoccurence
		return ret
	}
	return *o.EEzsigntemplateformfieldPositioningoccurence
}

// GetEEzsigntemplateformfieldPositioningoccurenceOk returns a tuple with the EEzsigntemplateformfieldPositioningoccurence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetEEzsigntemplateformfieldPositioningoccurenceOk() (*FieldEEzsigntemplateformfieldPositioningoccurence, bool) {
	if o == nil || IsNil(o.EEzsigntemplateformfieldPositioningoccurence) {
		return nil, false
	}
	return o.EEzsigntemplateformfieldPositioningoccurence, true
}

// HasEEzsigntemplateformfieldPositioningoccurence returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldResponseCompound) HasEEzsigntemplateformfieldPositioningoccurence() bool {
	if o != nil && !IsNil(o.EEzsigntemplateformfieldPositioningoccurence) {
		return true
	}

	return false
}

// SetEEzsigntemplateformfieldPositioningoccurence gets a reference to the given FieldEEzsigntemplateformfieldPositioningoccurence and assigns it to the EEzsigntemplateformfieldPositioningoccurence field.
func (o *EzsigntemplateformfieldResponseCompound) SetEEzsigntemplateformfieldPositioningoccurence(v FieldEEzsigntemplateformfieldPositioningoccurence) {
	o.EEzsigntemplateformfieldPositioningoccurence = &v
}

// GetEEzsigntemplateformfieldHorizontalalignment returns the EEzsigntemplateformfieldHorizontalalignment field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldResponseCompound) GetEEzsigntemplateformfieldHorizontalalignment() EnumHorizontalalignment {
	if o == nil || IsNil(o.EEzsigntemplateformfieldHorizontalalignment) {
		var ret EnumHorizontalalignment
		return ret
	}
	return *o.EEzsigntemplateformfieldHorizontalalignment
}

// GetEEzsigntemplateformfieldHorizontalalignmentOk returns a tuple with the EEzsigntemplateformfieldHorizontalalignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetEEzsigntemplateformfieldHorizontalalignmentOk() (*EnumHorizontalalignment, bool) {
	if o == nil || IsNil(o.EEzsigntemplateformfieldHorizontalalignment) {
		return nil, false
	}
	return o.EEzsigntemplateformfieldHorizontalalignment, true
}

// HasEEzsigntemplateformfieldHorizontalalignment returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldResponseCompound) HasEEzsigntemplateformfieldHorizontalalignment() bool {
	if o != nil && !IsNil(o.EEzsigntemplateformfieldHorizontalalignment) {
		return true
	}

	return false
}

// SetEEzsigntemplateformfieldHorizontalalignment gets a reference to the given EnumHorizontalalignment and assigns it to the EEzsigntemplateformfieldHorizontalalignment field.
func (o *EzsigntemplateformfieldResponseCompound) SetEEzsigntemplateformfieldHorizontalalignment(v EnumHorizontalalignment) {
	o.EEzsigntemplateformfieldHorizontalalignment = &v
}

// GetObjTextstylestatic returns the ObjTextstylestatic field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldResponseCompound) GetObjTextstylestatic() TextstylestaticResponseCompound {
	if o == nil || IsNil(o.ObjTextstylestatic) {
		var ret TextstylestaticResponseCompound
		return ret
	}
	return *o.ObjTextstylestatic
}

// GetObjTextstylestaticOk returns a tuple with the ObjTextstylestatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetObjTextstylestaticOk() (*TextstylestaticResponseCompound, bool) {
	if o == nil || IsNil(o.ObjTextstylestatic) {
		return nil, false
	}
	return o.ObjTextstylestatic, true
}

// HasObjTextstylestatic returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldResponseCompound) HasObjTextstylestatic() bool {
	if o != nil && !IsNil(o.ObjTextstylestatic) {
		return true
	}

	return false
}

// SetObjTextstylestatic gets a reference to the given TextstylestaticResponseCompound and assigns it to the ObjTextstylestatic field.
func (o *EzsigntemplateformfieldResponseCompound) SetObjTextstylestatic(v TextstylestaticResponseCompound) {
	o.ObjTextstylestatic = &v
}

// GetAObjEzsigntemplateelementdependency returns the AObjEzsigntemplateelementdependency field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldResponseCompound) GetAObjEzsigntemplateelementdependency() []EzsigntemplateelementdependencyResponseCompound {
	if o == nil || IsNil(o.AObjEzsigntemplateelementdependency) {
		var ret []EzsigntemplateelementdependencyResponseCompound
		return ret
	}
	return o.AObjEzsigntemplateelementdependency
}

// GetAObjEzsigntemplateelementdependencyOk returns a tuple with the AObjEzsigntemplateelementdependency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldResponseCompound) GetAObjEzsigntemplateelementdependencyOk() ([]EzsigntemplateelementdependencyResponseCompound, bool) {
	if o == nil || IsNil(o.AObjEzsigntemplateelementdependency) {
		return nil, false
	}
	return o.AObjEzsigntemplateelementdependency, true
}

// HasAObjEzsigntemplateelementdependency returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldResponseCompound) HasAObjEzsigntemplateelementdependency() bool {
	if o != nil && !IsNil(o.AObjEzsigntemplateelementdependency) {
		return true
	}

	return false
}

// SetAObjEzsigntemplateelementdependency gets a reference to the given []EzsigntemplateelementdependencyResponseCompound and assigns it to the AObjEzsigntemplateelementdependency field.
func (o *EzsigntemplateformfieldResponseCompound) SetAObjEzsigntemplateelementdependency(v []EzsigntemplateelementdependencyResponseCompound) {
	o.AObjEzsigntemplateelementdependency = v
}

func (o EzsigntemplateformfieldResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateformfieldResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplateformfieldID"] = o.PkiEzsigntemplateformfieldID
	if !IsNil(o.EEzsigntemplateformfieldPositioning) {
		toSerialize["eEzsigntemplateformfieldPositioning"] = o.EEzsigntemplateformfieldPositioning
	}
	toSerialize["iEzsigntemplatedocumentpagePagenumber"] = o.IEzsigntemplatedocumentpagePagenumber
	toSerialize["sEzsigntemplateformfieldLabel"] = o.SEzsigntemplateformfieldLabel
	if !IsNil(o.SEzsigntemplateformfieldValue) {
		toSerialize["sEzsigntemplateformfieldValue"] = o.SEzsigntemplateformfieldValue
	}
	if !IsNil(o.IEzsigntemplateformfieldX) {
		toSerialize["iEzsigntemplateformfieldX"] = o.IEzsigntemplateformfieldX
	}
	if !IsNil(o.IEzsigntemplateformfieldY) {
		toSerialize["iEzsigntemplateformfieldY"] = o.IEzsigntemplateformfieldY
	}
	toSerialize["iEzsigntemplateformfieldWidth"] = o.IEzsigntemplateformfieldWidth
	toSerialize["iEzsigntemplateformfieldHeight"] = o.IEzsigntemplateformfieldHeight
	if !IsNil(o.BEzsigntemplateformfieldAutocomplete) {
		toSerialize["bEzsigntemplateformfieldAutocomplete"] = o.BEzsigntemplateformfieldAutocomplete
	}
	if !IsNil(o.BEzsigntemplateformfieldSelected) {
		toSerialize["bEzsigntemplateformfieldSelected"] = o.BEzsigntemplateformfieldSelected
	}
	if !IsNil(o.EEzsigntemplateformfieldDependencyrequirement) {
		toSerialize["eEzsigntemplateformfieldDependencyrequirement"] = o.EEzsigntemplateformfieldDependencyrequirement
	}
	if !IsNil(o.SEzsigntemplateformfieldPositioningpattern) {
		toSerialize["sEzsigntemplateformfieldPositioningpattern"] = o.SEzsigntemplateformfieldPositioningpattern
	}
	if !IsNil(o.IEzsigntemplateformfieldPositioningoffsetx) {
		toSerialize["iEzsigntemplateformfieldPositioningoffsetx"] = o.IEzsigntemplateformfieldPositioningoffsetx
	}
	if !IsNil(o.IEzsigntemplateformfieldPositioningoffsety) {
		toSerialize["iEzsigntemplateformfieldPositioningoffsety"] = o.IEzsigntemplateformfieldPositioningoffsety
	}
	if !IsNil(o.EEzsigntemplateformfieldPositioningoccurence) {
		toSerialize["eEzsigntemplateformfieldPositioningoccurence"] = o.EEzsigntemplateformfieldPositioningoccurence
	}
	if !IsNil(o.EEzsigntemplateformfieldHorizontalalignment) {
		toSerialize["eEzsigntemplateformfieldHorizontalalignment"] = o.EEzsigntemplateformfieldHorizontalalignment
	}
	if !IsNil(o.ObjTextstylestatic) {
		toSerialize["objTextstylestatic"] = o.ObjTextstylestatic
	}
	if !IsNil(o.AObjEzsigntemplateelementdependency) {
		toSerialize["a_objEzsigntemplateelementdependency"] = o.AObjEzsigntemplateelementdependency
	}
	return toSerialize, nil
}

func (o *EzsigntemplateformfieldResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplateformfieldID",
		"iEzsigntemplatedocumentpagePagenumber",
		"sEzsigntemplateformfieldLabel",
		"iEzsigntemplateformfieldWidth",
		"iEzsigntemplateformfieldHeight",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEzsigntemplateformfieldResponseCompound := _EzsigntemplateformfieldResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateformfieldResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplateformfieldResponseCompound(varEzsigntemplateformfieldResponseCompound)

	return err
}

type NullableEzsigntemplateformfieldResponseCompound struct {
	value *EzsigntemplateformfieldResponseCompound
	isSet bool
}

func (v NullableEzsigntemplateformfieldResponseCompound) Get() *EzsigntemplateformfieldResponseCompound {
	return v.value
}

func (v *NullableEzsigntemplateformfieldResponseCompound) Set(val *EzsigntemplateformfieldResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateformfieldResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateformfieldResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateformfieldResponseCompound(val *EzsigntemplateformfieldResponseCompound) *NullableEzsigntemplateformfieldResponseCompound {
	return &NullableEzsigntemplateformfieldResponseCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplateformfieldResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateformfieldResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


