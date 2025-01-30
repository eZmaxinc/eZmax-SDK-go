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

// checks if the EzsignbulksenddocumentmappingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksenddocumentmappingResponse{}

// EzsignbulksenddocumentmappingResponse A Ezsignbulksenddocumentmapping Object
type EzsignbulksenddocumentmappingResponse struct {
	// The unique ID of the Ezsignbulksenddocumentmapping.
	PkiEzsignbulksenddocumentmappingID int32 `json:"pkiEzsignbulksenddocumentmappingID"`
	// The unique ID of the Ezsignbulksend
	FkiEzsignbulksendID int32 `json:"fkiEzsignbulksendID"`
	// The unique ID of the Ezsigntemplatepackage
	FkiEzsigntemplatepackageID *int32 `json:"fkiEzsigntemplatepackageID,omitempty"`
	// The unique ID of the Ezsigntemplate
	FkiEzsigntemplateID *int32 `json:"fkiEzsigntemplateID,omitempty"`
	// The order in which the Ezsigntemplate or Ezsigntemplatepackage will be presented to the signatory in the Ezsignfolder.
	IEzsignbulksenddocumentmappingOrder int32 `json:"iEzsignbulksenddocumentmappingOrder"`
}

type _EzsignbulksenddocumentmappingResponse EzsignbulksenddocumentmappingResponse

// NewEzsignbulksenddocumentmappingResponse instantiates a new EzsignbulksenddocumentmappingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksenddocumentmappingResponse(pkiEzsignbulksenddocumentmappingID int32, fkiEzsignbulksendID int32, iEzsignbulksenddocumentmappingOrder int32) *EzsignbulksenddocumentmappingResponse {
	this := EzsignbulksenddocumentmappingResponse{}
	this.PkiEzsignbulksenddocumentmappingID = pkiEzsignbulksenddocumentmappingID
	this.FkiEzsignbulksendID = fkiEzsignbulksendID
	this.IEzsignbulksenddocumentmappingOrder = iEzsignbulksenddocumentmappingOrder
	return &this
}

// NewEzsignbulksenddocumentmappingResponseWithDefaults instantiates a new EzsignbulksenddocumentmappingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksenddocumentmappingResponseWithDefaults() *EzsignbulksenddocumentmappingResponse {
	this := EzsignbulksenddocumentmappingResponse{}
	return &this
}

// GetPkiEzsignbulksenddocumentmappingID returns the PkiEzsignbulksenddocumentmappingID field value
func (o *EzsignbulksenddocumentmappingResponse) GetPkiEzsignbulksenddocumentmappingID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignbulksenddocumentmappingID
}

// GetPkiEzsignbulksenddocumentmappingIDOk returns a tuple with the PkiEzsignbulksenddocumentmappingID field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksenddocumentmappingResponse) GetPkiEzsignbulksenddocumentmappingIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignbulksenddocumentmappingID, true
}

// SetPkiEzsignbulksenddocumentmappingID sets field value
func (o *EzsignbulksenddocumentmappingResponse) SetPkiEzsignbulksenddocumentmappingID(v int32) {
	o.PkiEzsignbulksenddocumentmappingID = v
}

// GetFkiEzsignbulksendID returns the FkiEzsignbulksendID field value
func (o *EzsignbulksenddocumentmappingResponse) GetFkiEzsignbulksendID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignbulksendID
}

// GetFkiEzsignbulksendIDOk returns a tuple with the FkiEzsignbulksendID field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksenddocumentmappingResponse) GetFkiEzsignbulksendIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignbulksendID, true
}

// SetFkiEzsignbulksendID sets field value
func (o *EzsignbulksenddocumentmappingResponse) SetFkiEzsignbulksendID(v int32) {
	o.FkiEzsignbulksendID = v
}

// GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field value if set, zero value otherwise.
func (o *EzsignbulksenddocumentmappingResponse) GetFkiEzsigntemplatepackageID() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplatepackageID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplatepackageID
}

// GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignbulksenddocumentmappingResponse) GetFkiEzsigntemplatepackageIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplatepackageID) {
		return nil, false
	}
	return o.FkiEzsigntemplatepackageID, true
}

// HasFkiEzsigntemplatepackageID returns a boolean if a field has been set.
func (o *EzsignbulksenddocumentmappingResponse) HasFkiEzsigntemplatepackageID() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplatepackageID) {
		return true
	}

	return false
}

// SetFkiEzsigntemplatepackageID gets a reference to the given int32 and assigns it to the FkiEzsigntemplatepackageID field.
func (o *EzsignbulksenddocumentmappingResponse) SetFkiEzsigntemplatepackageID(v int32) {
	o.FkiEzsigntemplatepackageID = &v
}

// GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field value if set, zero value otherwise.
func (o *EzsignbulksenddocumentmappingResponse) GetFkiEzsigntemplateID() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplateID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplateID
}

// GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignbulksenddocumentmappingResponse) GetFkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplateID) {
		return nil, false
	}
	return o.FkiEzsigntemplateID, true
}

// HasFkiEzsigntemplateID returns a boolean if a field has been set.
func (o *EzsignbulksenddocumentmappingResponse) HasFkiEzsigntemplateID() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplateID) {
		return true
	}

	return false
}

// SetFkiEzsigntemplateID gets a reference to the given int32 and assigns it to the FkiEzsigntemplateID field.
func (o *EzsignbulksenddocumentmappingResponse) SetFkiEzsigntemplateID(v int32) {
	o.FkiEzsigntemplateID = &v
}

// GetIEzsignbulksenddocumentmappingOrder returns the IEzsignbulksenddocumentmappingOrder field value
func (o *EzsignbulksenddocumentmappingResponse) GetIEzsignbulksenddocumentmappingOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignbulksenddocumentmappingOrder
}

// GetIEzsignbulksenddocumentmappingOrderOk returns a tuple with the IEzsignbulksenddocumentmappingOrder field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksenddocumentmappingResponse) GetIEzsignbulksenddocumentmappingOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignbulksenddocumentmappingOrder, true
}

// SetIEzsignbulksenddocumentmappingOrder sets field value
func (o *EzsignbulksenddocumentmappingResponse) SetIEzsignbulksenddocumentmappingOrder(v int32) {
	o.IEzsignbulksenddocumentmappingOrder = v
}

func (o EzsignbulksenddocumentmappingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksenddocumentmappingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignbulksenddocumentmappingID"] = o.PkiEzsignbulksenddocumentmappingID
	toSerialize["fkiEzsignbulksendID"] = o.FkiEzsignbulksendID
	if !IsNil(o.FkiEzsigntemplatepackageID) {
		toSerialize["fkiEzsigntemplatepackageID"] = o.FkiEzsigntemplatepackageID
	}
	if !IsNil(o.FkiEzsigntemplateID) {
		toSerialize["fkiEzsigntemplateID"] = o.FkiEzsigntemplateID
	}
	toSerialize["iEzsignbulksenddocumentmappingOrder"] = o.IEzsignbulksenddocumentmappingOrder
	return toSerialize, nil
}

func (o *EzsignbulksenddocumentmappingResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignbulksenddocumentmappingID",
		"fkiEzsignbulksendID",
		"iEzsignbulksenddocumentmappingOrder",
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

	varEzsignbulksenddocumentmappingResponse := _EzsignbulksenddocumentmappingResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignbulksenddocumentmappingResponse)

	if err != nil {
		return err
	}

	*o = EzsignbulksenddocumentmappingResponse(varEzsignbulksenddocumentmappingResponse)

	return err
}

type NullableEzsignbulksenddocumentmappingResponse struct {
	value *EzsignbulksenddocumentmappingResponse
	isSet bool
}

func (v NullableEzsignbulksenddocumentmappingResponse) Get() *EzsignbulksenddocumentmappingResponse {
	return v.value
}

func (v *NullableEzsignbulksenddocumentmappingResponse) Set(val *EzsignbulksenddocumentmappingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksenddocumentmappingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksenddocumentmappingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksenddocumentmappingResponse(val *EzsignbulksenddocumentmappingResponse) *NullableEzsignbulksenddocumentmappingResponse {
	return &NullableEzsignbulksenddocumentmappingResponse{value: val, isSet: true}
}

func (v NullableEzsignbulksenddocumentmappingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksenddocumentmappingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


