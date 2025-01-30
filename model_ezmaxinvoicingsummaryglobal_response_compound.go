/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EzmaxinvoicingsummaryglobalResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxinvoicingsummaryglobalResponseCompound{}

// EzmaxinvoicingsummaryglobalResponseCompound A Ezmaxinvoicingsummaryglobal Object
type EzmaxinvoicingsummaryglobalResponseCompound struct {
	// The unique ID of the Ezmaxinvoicingsummaryglobal
	PkiEzmaxinvoicingsummaryglobalID *int32 `json:"pkiEzmaxinvoicingsummaryglobalID,omitempty"`
	// The unique ID of the Ezmaxinvoicing
	FkiEzmaxinvoicingID *int32 `json:"fkiEzmaxinvoicingID,omitempty"`
	// The unique ID of the Ezmaxproduct
	FkiEzmaxproductID int32 `json:"fkiEzmaxproductID"`
	// The description of the Ezmaxproduct in the language of the requester
	SEzmaxproductDescriptionX string `json:"sEzmaxproductDescriptionX"`
	// The start date for the Ezmaxinvoicingsummaryglobal
	DtEzmaxinvoicingsummaryglobalStart string `json:"dtEzmaxinvoicingsummaryglobalStart"`
	// The end date for the Ezmaxinvoicingsummaryglobal
	DtEzmaxinvoicingsummaryglobalEnd string `json:"dtEzmaxinvoicingsummaryglobalEnd"`
	// The number of days for the Ezmaxinvoicingsummaryglobal
	IEzmaxinvoicingsummaryglobalDays int32 `json:"iEzmaxinvoicingsummaryglobalDays"`
	// The count item calculated
	DEzmaxinvoicingsummaryglobalCountreal string `json:"dEzmaxinvoicingsummaryglobalCountreal" validate:"regexp=^-{0,1}[\\\\d]{1,6}?\\\\.[\\\\d]{2}$"`
	// The count item billed
	DEzmaxinvoicingsummaryglobalCountbilled string `json:"dEzmaxinvoicingsummaryglobalCountbilled" validate:"regexp=^-{0,1}[\\\\d]{1,6}?\\\\.[\\\\d]{2}$"`
	// The Ezmaxinvoicingsummaryglobal subtotal
	DEzmaxinvoicingsummaryglobalSubtotal string `json:"dEzmaxinvoicingsummaryglobalSubtotal" validate:"regexp=^-{0,1}[\\\\d]{1,9}?\\\\.[\\\\d]{2}$"`
	// The rebate amount for the Ezmaxinvoicingsummaryglobal
	DEzmaxinvoicingsummaryglobalRebateamount string `json:"dEzmaxinvoicingsummaryglobalRebateamount" validate:"regexp=^-{0,1}[\\\\d]{1,9}?\\\\.[\\\\d]{2}$"`
	// The rebate percentage of the Ezmaxinvoicingsummaryglobal
	DEzmaxinvoicingsummaryglobalRebatepercent string `json:"dEzmaxinvoicingsummaryglobalRebatepercent" validate:"regexp=^-{0,1}[\\\\d]{1,3}?\\\\.[\\\\d]{2}$"`
	// The rebate amount total for the Ezmaxinvoicingsummaryglobal
	DEzmaxinvoicingsummaryglobalRebatetotal string `json:"dEzmaxinvoicingsummaryglobalRebatetotal" validate:"regexp=^-{0,1}[\\\\d]{1,9}?\\\\.[\\\\d]{2}$"`
	// The Ezmaxinvoicingsummaryglobal total
	DEzmaxinvoicingsummaryglobalTotal string `json:"dEzmaxinvoicingsummaryglobalTotal" validate:"regexp=^-{0,1}[\\\\d]{1,9}?\\\\.[\\\\d]{2}$"`
	// The amount of commission for the representative
	DEzmaxinvoicingsummaryglobalRepresentative *string `json:"dEzmaxinvoicingsummaryglobalRepresentative,omitempty" validate:"regexp=^-{0,1}[\\\\d]{1,9}?\\\\.[\\\\d]{2}$"`
	// The amount of commission for the partner
	DEzmaxinvoicingsummaryglobalPartner *string `json:"dEzmaxinvoicingsummaryglobalPartner,omitempty" validate:"regexp=^-{0,1}[\\\\d]{1,9}?\\\\.[\\\\d]{2}$"`
	// The net amount of the Ezmaxinvoicingsummaryglobal
	DEzmaxinvoicingsummaryglobalNet *string `json:"dEzmaxinvoicingsummaryglobalNet,omitempty" validate:"regexp=^-{0,1}[\\\\d]{1,9}?\\\\.[\\\\d]{2}$"`
	// Whether it is adjustment for the Ezmaxinvoicingsummaryglobal
	BEzmaxinvoicingsummaryglobalAdjustment bool `json:"bEzmaxinvoicingsummaryglobalAdjustment"`
	// The help message of the Ezmaxproduct in the language of the requester
	TEzmaxproductHelpX string `json:"tEzmaxproductHelpX"`
	AObjEzmaxinvoicingcommission []EzmaxinvoicingcommissionResponseCompound `json:"a_objEzmaxinvoicingcommission,omitempty"`
}

type _EzmaxinvoicingsummaryglobalResponseCompound EzmaxinvoicingsummaryglobalResponseCompound

// NewEzmaxinvoicingsummaryglobalResponseCompound instantiates a new EzmaxinvoicingsummaryglobalResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxinvoicingsummaryglobalResponseCompound(fkiEzmaxproductID int32, sEzmaxproductDescriptionX string, dtEzmaxinvoicingsummaryglobalStart string, dtEzmaxinvoicingsummaryglobalEnd string, iEzmaxinvoicingsummaryglobalDays int32, dEzmaxinvoicingsummaryglobalCountreal string, dEzmaxinvoicingsummaryglobalCountbilled string, dEzmaxinvoicingsummaryglobalSubtotal string, dEzmaxinvoicingsummaryglobalRebateamount string, dEzmaxinvoicingsummaryglobalRebatepercent string, dEzmaxinvoicingsummaryglobalRebatetotal string, dEzmaxinvoicingsummaryglobalTotal string, bEzmaxinvoicingsummaryglobalAdjustment bool, tEzmaxproductHelpX string) *EzmaxinvoicingsummaryglobalResponseCompound {
	this := EzmaxinvoicingsummaryglobalResponseCompound{}
	this.FkiEzmaxproductID = fkiEzmaxproductID
	this.SEzmaxproductDescriptionX = sEzmaxproductDescriptionX
	this.DtEzmaxinvoicingsummaryglobalStart = dtEzmaxinvoicingsummaryglobalStart
	this.DtEzmaxinvoicingsummaryglobalEnd = dtEzmaxinvoicingsummaryglobalEnd
	this.IEzmaxinvoicingsummaryglobalDays = iEzmaxinvoicingsummaryglobalDays
	this.DEzmaxinvoicingsummaryglobalCountreal = dEzmaxinvoicingsummaryglobalCountreal
	this.DEzmaxinvoicingsummaryglobalCountbilled = dEzmaxinvoicingsummaryglobalCountbilled
	this.DEzmaxinvoicingsummaryglobalSubtotal = dEzmaxinvoicingsummaryglobalSubtotal
	this.DEzmaxinvoicingsummaryglobalRebateamount = dEzmaxinvoicingsummaryglobalRebateamount
	this.DEzmaxinvoicingsummaryglobalRebatepercent = dEzmaxinvoicingsummaryglobalRebatepercent
	this.DEzmaxinvoicingsummaryglobalRebatetotal = dEzmaxinvoicingsummaryglobalRebatetotal
	this.DEzmaxinvoicingsummaryglobalTotal = dEzmaxinvoicingsummaryglobalTotal
	this.BEzmaxinvoicingsummaryglobalAdjustment = bEzmaxinvoicingsummaryglobalAdjustment
	this.TEzmaxproductHelpX = tEzmaxproductHelpX
	return &this
}

// NewEzmaxinvoicingsummaryglobalResponseCompoundWithDefaults instantiates a new EzmaxinvoicingsummaryglobalResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxinvoicingsummaryglobalResponseCompoundWithDefaults() *EzmaxinvoicingsummaryglobalResponseCompound {
	this := EzmaxinvoicingsummaryglobalResponseCompound{}
	return &this
}

// GetPkiEzmaxinvoicingsummaryglobalID returns the PkiEzmaxinvoicingsummaryglobalID field value if set, zero value otherwise.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetPkiEzmaxinvoicingsummaryglobalID() int32 {
	if o == nil || IsNil(o.PkiEzmaxinvoicingsummaryglobalID) {
		var ret int32
		return ret
	}
	return *o.PkiEzmaxinvoicingsummaryglobalID
}

// GetPkiEzmaxinvoicingsummaryglobalIDOk returns a tuple with the PkiEzmaxinvoicingsummaryglobalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetPkiEzmaxinvoicingsummaryglobalIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzmaxinvoicingsummaryglobalID) {
		return nil, false
	}
	return o.PkiEzmaxinvoicingsummaryglobalID, true
}

// HasPkiEzmaxinvoicingsummaryglobalID returns a boolean if a field has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) HasPkiEzmaxinvoicingsummaryglobalID() bool {
	if o != nil && !IsNil(o.PkiEzmaxinvoicingsummaryglobalID) {
		return true
	}

	return false
}

// SetPkiEzmaxinvoicingsummaryglobalID gets a reference to the given int32 and assigns it to the PkiEzmaxinvoicingsummaryglobalID field.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetPkiEzmaxinvoicingsummaryglobalID(v int32) {
	o.PkiEzmaxinvoicingsummaryglobalID = &v
}

// GetFkiEzmaxinvoicingID returns the FkiEzmaxinvoicingID field value if set, zero value otherwise.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetFkiEzmaxinvoicingID() int32 {
	if o == nil || IsNil(o.FkiEzmaxinvoicingID) {
		var ret int32
		return ret
	}
	return *o.FkiEzmaxinvoicingID
}

// GetFkiEzmaxinvoicingIDOk returns a tuple with the FkiEzmaxinvoicingID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetFkiEzmaxinvoicingIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzmaxinvoicingID) {
		return nil, false
	}
	return o.FkiEzmaxinvoicingID, true
}

// HasFkiEzmaxinvoicingID returns a boolean if a field has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) HasFkiEzmaxinvoicingID() bool {
	if o != nil && !IsNil(o.FkiEzmaxinvoicingID) {
		return true
	}

	return false
}

// SetFkiEzmaxinvoicingID gets a reference to the given int32 and assigns it to the FkiEzmaxinvoicingID field.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetFkiEzmaxinvoicingID(v int32) {
	o.FkiEzmaxinvoicingID = &v
}

// GetFkiEzmaxproductID returns the FkiEzmaxproductID field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetFkiEzmaxproductID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzmaxproductID
}

// GetFkiEzmaxproductIDOk returns a tuple with the FkiEzmaxproductID field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetFkiEzmaxproductIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzmaxproductID, true
}

// SetFkiEzmaxproductID sets field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetFkiEzmaxproductID(v int32) {
	o.FkiEzmaxproductID = v
}

// GetSEzmaxproductDescriptionX returns the SEzmaxproductDescriptionX field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetSEzmaxproductDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzmaxproductDescriptionX
}

// GetSEzmaxproductDescriptionXOk returns a tuple with the SEzmaxproductDescriptionX field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetSEzmaxproductDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzmaxproductDescriptionX, true
}

// SetSEzmaxproductDescriptionX sets field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetSEzmaxproductDescriptionX(v string) {
	o.SEzmaxproductDescriptionX = v
}

// GetDtEzmaxinvoicingsummaryglobalStart returns the DtEzmaxinvoicingsummaryglobalStart field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDtEzmaxinvoicingsummaryglobalStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtEzmaxinvoicingsummaryglobalStart
}

// GetDtEzmaxinvoicingsummaryglobalStartOk returns a tuple with the DtEzmaxinvoicingsummaryglobalStart field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDtEzmaxinvoicingsummaryglobalStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtEzmaxinvoicingsummaryglobalStart, true
}

// SetDtEzmaxinvoicingsummaryglobalStart sets field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDtEzmaxinvoicingsummaryglobalStart(v string) {
	o.DtEzmaxinvoicingsummaryglobalStart = v
}

// GetDtEzmaxinvoicingsummaryglobalEnd returns the DtEzmaxinvoicingsummaryglobalEnd field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDtEzmaxinvoicingsummaryglobalEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtEzmaxinvoicingsummaryglobalEnd
}

// GetDtEzmaxinvoicingsummaryglobalEndOk returns a tuple with the DtEzmaxinvoicingsummaryglobalEnd field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDtEzmaxinvoicingsummaryglobalEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtEzmaxinvoicingsummaryglobalEnd, true
}

// SetDtEzmaxinvoicingsummaryglobalEnd sets field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDtEzmaxinvoicingsummaryglobalEnd(v string) {
	o.DtEzmaxinvoicingsummaryglobalEnd = v
}

// GetIEzmaxinvoicingsummaryglobalDays returns the IEzmaxinvoicingsummaryglobalDays field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetIEzmaxinvoicingsummaryglobalDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzmaxinvoicingsummaryglobalDays
}

// GetIEzmaxinvoicingsummaryglobalDaysOk returns a tuple with the IEzmaxinvoicingsummaryglobalDays field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetIEzmaxinvoicingsummaryglobalDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzmaxinvoicingsummaryglobalDays, true
}

// SetIEzmaxinvoicingsummaryglobalDays sets field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetIEzmaxinvoicingsummaryglobalDays(v int32) {
	o.IEzmaxinvoicingsummaryglobalDays = v
}

// GetDEzmaxinvoicingsummaryglobalCountreal returns the DEzmaxinvoicingsummaryglobalCountreal field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalCountreal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingsummaryglobalCountreal
}

// GetDEzmaxinvoicingsummaryglobalCountrealOk returns a tuple with the DEzmaxinvoicingsummaryglobalCountreal field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalCountrealOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingsummaryglobalCountreal, true
}

// SetDEzmaxinvoicingsummaryglobalCountreal sets field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalCountreal(v string) {
	o.DEzmaxinvoicingsummaryglobalCountreal = v
}

// GetDEzmaxinvoicingsummaryglobalCountbilled returns the DEzmaxinvoicingsummaryglobalCountbilled field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalCountbilled() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingsummaryglobalCountbilled
}

// GetDEzmaxinvoicingsummaryglobalCountbilledOk returns a tuple with the DEzmaxinvoicingsummaryglobalCountbilled field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalCountbilledOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingsummaryglobalCountbilled, true
}

// SetDEzmaxinvoicingsummaryglobalCountbilled sets field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalCountbilled(v string) {
	o.DEzmaxinvoicingsummaryglobalCountbilled = v
}

// GetDEzmaxinvoicingsummaryglobalSubtotal returns the DEzmaxinvoicingsummaryglobalSubtotal field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalSubtotal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingsummaryglobalSubtotal
}

// GetDEzmaxinvoicingsummaryglobalSubtotalOk returns a tuple with the DEzmaxinvoicingsummaryglobalSubtotal field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalSubtotalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingsummaryglobalSubtotal, true
}

// SetDEzmaxinvoicingsummaryglobalSubtotal sets field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalSubtotal(v string) {
	o.DEzmaxinvoicingsummaryglobalSubtotal = v
}

// GetDEzmaxinvoicingsummaryglobalRebateamount returns the DEzmaxinvoicingsummaryglobalRebateamount field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRebateamount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingsummaryglobalRebateamount
}

// GetDEzmaxinvoicingsummaryglobalRebateamountOk returns a tuple with the DEzmaxinvoicingsummaryglobalRebateamount field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRebateamountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingsummaryglobalRebateamount, true
}

// SetDEzmaxinvoicingsummaryglobalRebateamount sets field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalRebateamount(v string) {
	o.DEzmaxinvoicingsummaryglobalRebateamount = v
}

// GetDEzmaxinvoicingsummaryglobalRebatepercent returns the DEzmaxinvoicingsummaryglobalRebatepercent field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRebatepercent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingsummaryglobalRebatepercent
}

// GetDEzmaxinvoicingsummaryglobalRebatepercentOk returns a tuple with the DEzmaxinvoicingsummaryglobalRebatepercent field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRebatepercentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingsummaryglobalRebatepercent, true
}

// SetDEzmaxinvoicingsummaryglobalRebatepercent sets field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalRebatepercent(v string) {
	o.DEzmaxinvoicingsummaryglobalRebatepercent = v
}

// GetDEzmaxinvoicingsummaryglobalRebatetotal returns the DEzmaxinvoicingsummaryglobalRebatetotal field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRebatetotal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingsummaryglobalRebatetotal
}

// GetDEzmaxinvoicingsummaryglobalRebatetotalOk returns a tuple with the DEzmaxinvoicingsummaryglobalRebatetotal field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRebatetotalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingsummaryglobalRebatetotal, true
}

// SetDEzmaxinvoicingsummaryglobalRebatetotal sets field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalRebatetotal(v string) {
	o.DEzmaxinvoicingsummaryglobalRebatetotal = v
}

// GetDEzmaxinvoicingsummaryglobalTotal returns the DEzmaxinvoicingsummaryglobalTotal field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalTotal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingsummaryglobalTotal
}

// GetDEzmaxinvoicingsummaryglobalTotalOk returns a tuple with the DEzmaxinvoicingsummaryglobalTotal field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalTotalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingsummaryglobalTotal, true
}

// SetDEzmaxinvoicingsummaryglobalTotal sets field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalTotal(v string) {
	o.DEzmaxinvoicingsummaryglobalTotal = v
}

// GetDEzmaxinvoicingsummaryglobalRepresentative returns the DEzmaxinvoicingsummaryglobalRepresentative field value if set, zero value otherwise.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRepresentative() string {
	if o == nil || IsNil(o.DEzmaxinvoicingsummaryglobalRepresentative) {
		var ret string
		return ret
	}
	return *o.DEzmaxinvoicingsummaryglobalRepresentative
}

// GetDEzmaxinvoicingsummaryglobalRepresentativeOk returns a tuple with the DEzmaxinvoicingsummaryglobalRepresentative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRepresentativeOk() (*string, bool) {
	if o == nil || IsNil(o.DEzmaxinvoicingsummaryglobalRepresentative) {
		return nil, false
	}
	return o.DEzmaxinvoicingsummaryglobalRepresentative, true
}

// HasDEzmaxinvoicingsummaryglobalRepresentative returns a boolean if a field has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) HasDEzmaxinvoicingsummaryglobalRepresentative() bool {
	if o != nil && !IsNil(o.DEzmaxinvoicingsummaryglobalRepresentative) {
		return true
	}

	return false
}

// SetDEzmaxinvoicingsummaryglobalRepresentative gets a reference to the given string and assigns it to the DEzmaxinvoicingsummaryglobalRepresentative field.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalRepresentative(v string) {
	o.DEzmaxinvoicingsummaryglobalRepresentative = &v
}

// GetDEzmaxinvoicingsummaryglobalPartner returns the DEzmaxinvoicingsummaryglobalPartner field value if set, zero value otherwise.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalPartner() string {
	if o == nil || IsNil(o.DEzmaxinvoicingsummaryglobalPartner) {
		var ret string
		return ret
	}
	return *o.DEzmaxinvoicingsummaryglobalPartner
}

// GetDEzmaxinvoicingsummaryglobalPartnerOk returns a tuple with the DEzmaxinvoicingsummaryglobalPartner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalPartnerOk() (*string, bool) {
	if o == nil || IsNil(o.DEzmaxinvoicingsummaryglobalPartner) {
		return nil, false
	}
	return o.DEzmaxinvoicingsummaryglobalPartner, true
}

// HasDEzmaxinvoicingsummaryglobalPartner returns a boolean if a field has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) HasDEzmaxinvoicingsummaryglobalPartner() bool {
	if o != nil && !IsNil(o.DEzmaxinvoicingsummaryglobalPartner) {
		return true
	}

	return false
}

// SetDEzmaxinvoicingsummaryglobalPartner gets a reference to the given string and assigns it to the DEzmaxinvoicingsummaryglobalPartner field.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalPartner(v string) {
	o.DEzmaxinvoicingsummaryglobalPartner = &v
}

// GetDEzmaxinvoicingsummaryglobalNet returns the DEzmaxinvoicingsummaryglobalNet field value if set, zero value otherwise.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalNet() string {
	if o == nil || IsNil(o.DEzmaxinvoicingsummaryglobalNet) {
		var ret string
		return ret
	}
	return *o.DEzmaxinvoicingsummaryglobalNet
}

// GetDEzmaxinvoicingsummaryglobalNetOk returns a tuple with the DEzmaxinvoicingsummaryglobalNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalNetOk() (*string, bool) {
	if o == nil || IsNil(o.DEzmaxinvoicingsummaryglobalNet) {
		return nil, false
	}
	return o.DEzmaxinvoicingsummaryglobalNet, true
}

// HasDEzmaxinvoicingsummaryglobalNet returns a boolean if a field has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) HasDEzmaxinvoicingsummaryglobalNet() bool {
	if o != nil && !IsNil(o.DEzmaxinvoicingsummaryglobalNet) {
		return true
	}

	return false
}

// SetDEzmaxinvoicingsummaryglobalNet gets a reference to the given string and assigns it to the DEzmaxinvoicingsummaryglobalNet field.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalNet(v string) {
	o.DEzmaxinvoicingsummaryglobalNet = &v
}

// GetBEzmaxinvoicingsummaryglobalAdjustment returns the BEzmaxinvoicingsummaryglobalAdjustment field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetBEzmaxinvoicingsummaryglobalAdjustment() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzmaxinvoicingsummaryglobalAdjustment
}

// GetBEzmaxinvoicingsummaryglobalAdjustmentOk returns a tuple with the BEzmaxinvoicingsummaryglobalAdjustment field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetBEzmaxinvoicingsummaryglobalAdjustmentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzmaxinvoicingsummaryglobalAdjustment, true
}

// SetBEzmaxinvoicingsummaryglobalAdjustment sets field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetBEzmaxinvoicingsummaryglobalAdjustment(v bool) {
	o.BEzmaxinvoicingsummaryglobalAdjustment = v
}

// GetTEzmaxproductHelpX returns the TEzmaxproductHelpX field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetTEzmaxproductHelpX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TEzmaxproductHelpX
}

// GetTEzmaxproductHelpXOk returns a tuple with the TEzmaxproductHelpX field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetTEzmaxproductHelpXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TEzmaxproductHelpX, true
}

// SetTEzmaxproductHelpX sets field value
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetTEzmaxproductHelpX(v string) {
	o.TEzmaxproductHelpX = v
}

// GetAObjEzmaxinvoicingcommission returns the AObjEzmaxinvoicingcommission field value if set, zero value otherwise.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetAObjEzmaxinvoicingcommission() []EzmaxinvoicingcommissionResponseCompound {
	if o == nil || IsNil(o.AObjEzmaxinvoicingcommission) {
		var ret []EzmaxinvoicingcommissionResponseCompound
		return ret
	}
	return o.AObjEzmaxinvoicingcommission
}

// GetAObjEzmaxinvoicingcommissionOk returns a tuple with the AObjEzmaxinvoicingcommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetAObjEzmaxinvoicingcommissionOk() ([]EzmaxinvoicingcommissionResponseCompound, bool) {
	if o == nil || IsNil(o.AObjEzmaxinvoicingcommission) {
		return nil, false
	}
	return o.AObjEzmaxinvoicingcommission, true
}

// HasAObjEzmaxinvoicingcommission returns a boolean if a field has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) HasAObjEzmaxinvoicingcommission() bool {
	if o != nil && !IsNil(o.AObjEzmaxinvoicingcommission) {
		return true
	}

	return false
}

// SetAObjEzmaxinvoicingcommission gets a reference to the given []EzmaxinvoicingcommissionResponseCompound and assigns it to the AObjEzmaxinvoicingcommission field.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetAObjEzmaxinvoicingcommission(v []EzmaxinvoicingcommissionResponseCompound) {
	o.AObjEzmaxinvoicingcommission = v
}

func (o EzmaxinvoicingsummaryglobalResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxinvoicingsummaryglobalResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzmaxinvoicingsummaryglobalID) {
		toSerialize["pkiEzmaxinvoicingsummaryglobalID"] = o.PkiEzmaxinvoicingsummaryglobalID
	}
	if !IsNil(o.FkiEzmaxinvoicingID) {
		toSerialize["fkiEzmaxinvoicingID"] = o.FkiEzmaxinvoicingID
	}
	toSerialize["fkiEzmaxproductID"] = o.FkiEzmaxproductID
	toSerialize["sEzmaxproductDescriptionX"] = o.SEzmaxproductDescriptionX
	toSerialize["dtEzmaxinvoicingsummaryglobalStart"] = o.DtEzmaxinvoicingsummaryglobalStart
	toSerialize["dtEzmaxinvoicingsummaryglobalEnd"] = o.DtEzmaxinvoicingsummaryglobalEnd
	toSerialize["iEzmaxinvoicingsummaryglobalDays"] = o.IEzmaxinvoicingsummaryglobalDays
	toSerialize["dEzmaxinvoicingsummaryglobalCountreal"] = o.DEzmaxinvoicingsummaryglobalCountreal
	toSerialize["dEzmaxinvoicingsummaryglobalCountbilled"] = o.DEzmaxinvoicingsummaryglobalCountbilled
	toSerialize["dEzmaxinvoicingsummaryglobalSubtotal"] = o.DEzmaxinvoicingsummaryglobalSubtotal
	toSerialize["dEzmaxinvoicingsummaryglobalRebateamount"] = o.DEzmaxinvoicingsummaryglobalRebateamount
	toSerialize["dEzmaxinvoicingsummaryglobalRebatepercent"] = o.DEzmaxinvoicingsummaryglobalRebatepercent
	toSerialize["dEzmaxinvoicingsummaryglobalRebatetotal"] = o.DEzmaxinvoicingsummaryglobalRebatetotal
	toSerialize["dEzmaxinvoicingsummaryglobalTotal"] = o.DEzmaxinvoicingsummaryglobalTotal
	if !IsNil(o.DEzmaxinvoicingsummaryglobalRepresentative) {
		toSerialize["dEzmaxinvoicingsummaryglobalRepresentative"] = o.DEzmaxinvoicingsummaryglobalRepresentative
	}
	if !IsNil(o.DEzmaxinvoicingsummaryglobalPartner) {
		toSerialize["dEzmaxinvoicingsummaryglobalPartner"] = o.DEzmaxinvoicingsummaryglobalPartner
	}
	if !IsNil(o.DEzmaxinvoicingsummaryglobalNet) {
		toSerialize["dEzmaxinvoicingsummaryglobalNet"] = o.DEzmaxinvoicingsummaryglobalNet
	}
	toSerialize["bEzmaxinvoicingsummaryglobalAdjustment"] = o.BEzmaxinvoicingsummaryglobalAdjustment
	toSerialize["tEzmaxproductHelpX"] = o.TEzmaxproductHelpX
	if !IsNil(o.AObjEzmaxinvoicingcommission) {
		toSerialize["a_objEzmaxinvoicingcommission"] = o.AObjEzmaxinvoicingcommission
	}
	return toSerialize, nil
}

func (o *EzmaxinvoicingsummaryglobalResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzmaxproductID",
		"sEzmaxproductDescriptionX",
		"dtEzmaxinvoicingsummaryglobalStart",
		"dtEzmaxinvoicingsummaryglobalEnd",
		"iEzmaxinvoicingsummaryglobalDays",
		"dEzmaxinvoicingsummaryglobalCountreal",
		"dEzmaxinvoicingsummaryglobalCountbilled",
		"dEzmaxinvoicingsummaryglobalSubtotal",
		"dEzmaxinvoicingsummaryglobalRebateamount",
		"dEzmaxinvoicingsummaryglobalRebatepercent",
		"dEzmaxinvoicingsummaryglobalRebatetotal",
		"dEzmaxinvoicingsummaryglobalTotal",
		"bEzmaxinvoicingsummaryglobalAdjustment",
		"tEzmaxproductHelpX",
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

	varEzmaxinvoicingsummaryglobalResponseCompound := _EzmaxinvoicingsummaryglobalResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzmaxinvoicingsummaryglobalResponseCompound)

	if err != nil {
		return err
	}

	*o = EzmaxinvoicingsummaryglobalResponseCompound(varEzmaxinvoicingsummaryglobalResponseCompound)

	return err
}

type NullableEzmaxinvoicingsummaryglobalResponseCompound struct {
	value *EzmaxinvoicingsummaryglobalResponseCompound
	isSet bool
}

func (v NullableEzmaxinvoicingsummaryglobalResponseCompound) Get() *EzmaxinvoicingsummaryglobalResponseCompound {
	return v.value
}

func (v *NullableEzmaxinvoicingsummaryglobalResponseCompound) Set(val *EzmaxinvoicingsummaryglobalResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxinvoicingsummaryglobalResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxinvoicingsummaryglobalResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxinvoicingsummaryglobalResponseCompound(val *EzmaxinvoicingsummaryglobalResponseCompound) *NullableEzmaxinvoicingsummaryglobalResponseCompound {
	return &NullableEzmaxinvoicingsummaryglobalResponseCompound{value: val, isSet: true}
}

func (v NullableEzmaxinvoicingsummaryglobalResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxinvoicingsummaryglobalResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


