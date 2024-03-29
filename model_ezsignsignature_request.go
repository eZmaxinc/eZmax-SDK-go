/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.3
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignsignatureRequest An Ezsignsignature Object
type EzsignsignatureRequest struct {
	// The unique ID of the Ezsignfoldersignerassociation
	FkiEzsignfoldersignerassociationID int32 `json:"fkiEzsignfoldersignerassociationID"`
	// The page number in the Ezsigndocument
	IEzsignpagePagenumber int32 `json:"iEzsignpagePagenumber"`
	// The X coordinate (Horizontal) where to put the signature block on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the signature block 2 inches from the left border of the page, you would use \"200\" for the X coordinate.
	IEzsignsignatureX int32 `json:"iEzsignsignatureX"`
	// The Y coordinate (Vertical) where to put the signature block on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the signature block 3 inches from the top border of the page, you would use \"300\" for the Y coordinate.
	IEzsignsignatureY int32 `json:"iEzsignsignatureY"`
	// The step when the Ezsignsigner will be invited to sign.
	IEzsignsignatureStep int32 `json:"iEzsignsignatureStep"`
	EEzsignsignatureType FieldEEzsignsignatureType `json:"eEzsignsignatureType"`
	// The unique ID of the Ezsigndocument
	FkiEzsigndocumentID int32 `json:"fkiEzsigndocumentID"`
}

// NewEzsignsignatureRequest instantiates a new EzsignsignatureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignatureRequest(fkiEzsignfoldersignerassociationID int32, iEzsignpagePagenumber int32, iEzsignsignatureX int32, iEzsignsignatureY int32, iEzsignsignatureStep int32, eEzsignsignatureType FieldEEzsignsignatureType, fkiEzsigndocumentID int32) *EzsignsignatureRequest {
	this := EzsignsignatureRequest{}
	this.FkiEzsignfoldersignerassociationID = fkiEzsignfoldersignerassociationID
	this.IEzsignpagePagenumber = iEzsignpagePagenumber
	this.IEzsignsignatureX = iEzsignsignatureX
	this.IEzsignsignatureY = iEzsignsignatureY
	this.IEzsignsignatureStep = iEzsignsignatureStep
	this.EEzsignsignatureType = eEzsignsignatureType
	this.FkiEzsigndocumentID = fkiEzsigndocumentID
	return &this
}

// NewEzsignsignatureRequestWithDefaults instantiates a new EzsignsignatureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignatureRequestWithDefaults() *EzsignsignatureRequest {
	this := EzsignsignatureRequest{}
	return &this
}

// GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field value
func (o *EzsignsignatureRequest) GetFkiEzsignfoldersignerassociationID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldersignerassociationID
}

// GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureRequest) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsignfoldersignerassociationID, true
}

// SetFkiEzsignfoldersignerassociationID sets field value
func (o *EzsignsignatureRequest) SetFkiEzsignfoldersignerassociationID(v int32) {
	o.FkiEzsignfoldersignerassociationID = v
}

// GetIEzsignpagePagenumber returns the IEzsignpagePagenumber field value
func (o *EzsignsignatureRequest) GetIEzsignpagePagenumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignpagePagenumber
}

// GetIEzsignpagePagenumberOk returns a tuple with the IEzsignpagePagenumber field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureRequest) GetIEzsignpagePagenumberOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsignpagePagenumber, true
}

// SetIEzsignpagePagenumber sets field value
func (o *EzsignsignatureRequest) SetIEzsignpagePagenumber(v int32) {
	o.IEzsignpagePagenumber = v
}

// GetIEzsignsignatureX returns the IEzsignsignatureX field value
func (o *EzsignsignatureRequest) GetIEzsignsignatureX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignatureX
}

// GetIEzsignsignatureXOk returns a tuple with the IEzsignsignatureX field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureRequest) GetIEzsignsignatureXOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsignsignatureX, true
}

// SetIEzsignsignatureX sets field value
func (o *EzsignsignatureRequest) SetIEzsignsignatureX(v int32) {
	o.IEzsignsignatureX = v
}

// GetIEzsignsignatureY returns the IEzsignsignatureY field value
func (o *EzsignsignatureRequest) GetIEzsignsignatureY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignatureY
}

// GetIEzsignsignatureYOk returns a tuple with the IEzsignsignatureY field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureRequest) GetIEzsignsignatureYOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsignsignatureY, true
}

// SetIEzsignsignatureY sets field value
func (o *EzsignsignatureRequest) SetIEzsignsignatureY(v int32) {
	o.IEzsignsignatureY = v
}

// GetIEzsignsignatureStep returns the IEzsignsignatureStep field value
func (o *EzsignsignatureRequest) GetIEzsignsignatureStep() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignatureStep
}

// GetIEzsignsignatureStepOk returns a tuple with the IEzsignsignatureStep field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureRequest) GetIEzsignsignatureStepOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsignsignatureStep, true
}

// SetIEzsignsignatureStep sets field value
func (o *EzsignsignatureRequest) SetIEzsignsignatureStep(v int32) {
	o.IEzsignsignatureStep = v
}

// GetEEzsignsignatureType returns the EEzsignsignatureType field value
func (o *EzsignsignatureRequest) GetEEzsignsignatureType() FieldEEzsignsignatureType {
	if o == nil {
		var ret FieldEEzsignsignatureType
		return ret
	}

	return o.EEzsignsignatureType
}

// GetEEzsignsignatureTypeOk returns a tuple with the EEzsignsignatureType field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureRequest) GetEEzsignsignatureTypeOk() (*FieldEEzsignsignatureType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EEzsignsignatureType, true
}

// SetEEzsignsignatureType sets field value
func (o *EzsignsignatureRequest) SetEEzsignsignatureType(v FieldEEzsignsignatureType) {
	o.EEzsignsignatureType = v
}

// GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field value
func (o *EzsignsignatureRequest) GetFkiEzsigndocumentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigndocumentID
}

// GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureRequest) GetFkiEzsigndocumentIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsigndocumentID, true
}

// SetFkiEzsigndocumentID sets field value
func (o *EzsignsignatureRequest) SetFkiEzsigndocumentID(v int32) {
	o.FkiEzsigndocumentID = v
}

func (o EzsignsignatureRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fkiEzsignfoldersignerassociationID"] = o.FkiEzsignfoldersignerassociationID
	}
	if true {
		toSerialize["iEzsignpagePagenumber"] = o.IEzsignpagePagenumber
	}
	if true {
		toSerialize["iEzsignsignatureX"] = o.IEzsignsignatureX
	}
	if true {
		toSerialize["iEzsignsignatureY"] = o.IEzsignsignatureY
	}
	if true {
		toSerialize["iEzsignsignatureStep"] = o.IEzsignsignatureStep
	}
	if true {
		toSerialize["eEzsignsignatureType"] = o.EEzsignsignatureType
	}
	if true {
		toSerialize["fkiEzsigndocumentID"] = o.FkiEzsigndocumentID
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignsignatureRequest struct {
	value *EzsignsignatureRequest
	isSet bool
}

func (v NullableEzsignsignatureRequest) Get() *EzsignsignatureRequest {
	return v.value
}

func (v *NullableEzsignsignatureRequest) Set(val *EzsignsignatureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignatureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignatureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignatureRequest(val *EzsignsignatureRequest) *NullableEzsignsignatureRequest {
	return &NullableEzsignsignatureRequest{value: val, isSet: true}
}

func (v NullableEzsignsignatureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignatureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


