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

// checks if the EzdoctemplatedocumentListElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzdoctemplatedocumentListElement{}

// EzdoctemplatedocumentListElement A Ezdoctemplatedocument List Element
type EzdoctemplatedocumentListElement struct {
	// The unique ID of the Ezdoctemplatedocument
	PkiEzdoctemplatedocumentID int32 `json:"pkiEzdoctemplatedocumentID"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID *int32 `json:"fkiEzsignfoldertypeID,omitempty"`
	// The unique ID of the Ezdoctemplatetype
	FkiEzdoctemplatetypeID int32 `json:"fkiEzdoctemplatetypeID"`
	// The unique ID of the Ezdoctemplatefieldtypecategory
	FkiEzdoctemplatefieldtypecategoryID int32 `json:"fkiEzdoctemplatefieldtypecategoryID"`
	// The name of the Ezsignfoldertype in the language of the requester
	SEzsignfoldertypeNameX *string `json:"sEzsignfoldertypeNameX,omitempty"`
	// The description of the Ezdoctemplatetype in the language of the requester
	SEzdoctemplatetypeDescriptionX *string `json:"sEzdoctemplatetypeDescriptionX,omitempty" validate:"regexp=^.{0,50}$"`
	// The description of the Ezdoctemplatefieldtypecategory in the language of the requester
	SEzdoctemplatefieldtypecategoryDescriptionX *string `json:"sEzdoctemplatefieldtypecategoryDescriptionX,omitempty" validate:"regexp=^.{0,55}$"`
	EEzdoctemplatedocumentPrivacylevel *FieldEEzdoctemplatedocumentPrivacylevel `json:"eEzdoctemplatedocumentPrivacylevel,omitempty"`
	// Whether the ezdoctemplatedocument is active or not
	BEzdoctemplatedocumentIsactive bool `json:"bEzdoctemplatedocumentIsactive"`
	// The name of the Ezdoctemplatedocument in the language of the requester
	SEzdoctemplatedocumentNameX string `json:"sEzdoctemplatedocumentNameX" validate:"regexp=^.{0,50}$"`
}

type _EzdoctemplatedocumentListElement EzdoctemplatedocumentListElement

// NewEzdoctemplatedocumentListElement instantiates a new EzdoctemplatedocumentListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzdoctemplatedocumentListElement(pkiEzdoctemplatedocumentID int32, fkiLanguageID int32, fkiEzdoctemplatetypeID int32, fkiEzdoctemplatefieldtypecategoryID int32, bEzdoctemplatedocumentIsactive bool, sEzdoctemplatedocumentNameX string) *EzdoctemplatedocumentListElement {
	this := EzdoctemplatedocumentListElement{}
	this.PkiEzdoctemplatedocumentID = pkiEzdoctemplatedocumentID
	this.FkiLanguageID = fkiLanguageID
	this.FkiEzdoctemplatetypeID = fkiEzdoctemplatetypeID
	this.FkiEzdoctemplatefieldtypecategoryID = fkiEzdoctemplatefieldtypecategoryID
	this.BEzdoctemplatedocumentIsactive = bEzdoctemplatedocumentIsactive
	this.SEzdoctemplatedocumentNameX = sEzdoctemplatedocumentNameX
	return &this
}

// NewEzdoctemplatedocumentListElementWithDefaults instantiates a new EzdoctemplatedocumentListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzdoctemplatedocumentListElementWithDefaults() *EzdoctemplatedocumentListElement {
	this := EzdoctemplatedocumentListElement{}
	return &this
}

// GetPkiEzdoctemplatedocumentID returns the PkiEzdoctemplatedocumentID field value
func (o *EzdoctemplatedocumentListElement) GetPkiEzdoctemplatedocumentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzdoctemplatedocumentID
}

// GetPkiEzdoctemplatedocumentIDOk returns a tuple with the PkiEzdoctemplatedocumentID field value
// and a boolean to check if the value has been set.
func (o *EzdoctemplatedocumentListElement) GetPkiEzdoctemplatedocumentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzdoctemplatedocumentID, true
}

// SetPkiEzdoctemplatedocumentID sets field value
func (o *EzdoctemplatedocumentListElement) SetPkiEzdoctemplatedocumentID(v int32) {
	o.PkiEzdoctemplatedocumentID = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzdoctemplatedocumentListElement) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzdoctemplatedocumentListElement) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzdoctemplatedocumentListElement) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value if set, zero value otherwise.
func (o *EzdoctemplatedocumentListElement) GetFkiEzsignfoldertypeID() int32 {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzdoctemplatedocumentListElement) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		return nil, false
	}
	return o.FkiEzsignfoldertypeID, true
}

// HasFkiEzsignfoldertypeID returns a boolean if a field has been set.
func (o *EzdoctemplatedocumentListElement) HasFkiEzsignfoldertypeID() bool {
	if o != nil && !IsNil(o.FkiEzsignfoldertypeID) {
		return true
	}

	return false
}

// SetFkiEzsignfoldertypeID gets a reference to the given int32 and assigns it to the FkiEzsignfoldertypeID field.
func (o *EzdoctemplatedocumentListElement) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = &v
}

// GetFkiEzdoctemplatetypeID returns the FkiEzdoctemplatetypeID field value
func (o *EzdoctemplatedocumentListElement) GetFkiEzdoctemplatetypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzdoctemplatetypeID
}

// GetFkiEzdoctemplatetypeIDOk returns a tuple with the FkiEzdoctemplatetypeID field value
// and a boolean to check if the value has been set.
func (o *EzdoctemplatedocumentListElement) GetFkiEzdoctemplatetypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzdoctemplatetypeID, true
}

// SetFkiEzdoctemplatetypeID sets field value
func (o *EzdoctemplatedocumentListElement) SetFkiEzdoctemplatetypeID(v int32) {
	o.FkiEzdoctemplatetypeID = v
}

// GetFkiEzdoctemplatefieldtypecategoryID returns the FkiEzdoctemplatefieldtypecategoryID field value
func (o *EzdoctemplatedocumentListElement) GetFkiEzdoctemplatefieldtypecategoryID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzdoctemplatefieldtypecategoryID
}

// GetFkiEzdoctemplatefieldtypecategoryIDOk returns a tuple with the FkiEzdoctemplatefieldtypecategoryID field value
// and a boolean to check if the value has been set.
func (o *EzdoctemplatedocumentListElement) GetFkiEzdoctemplatefieldtypecategoryIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzdoctemplatefieldtypecategoryID, true
}

// SetFkiEzdoctemplatefieldtypecategoryID sets field value
func (o *EzdoctemplatedocumentListElement) SetFkiEzdoctemplatefieldtypecategoryID(v int32) {
	o.FkiEzdoctemplatefieldtypecategoryID = v
}

// GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field value if set, zero value otherwise.
func (o *EzdoctemplatedocumentListElement) GetSEzsignfoldertypeNameX() string {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		var ret string
		return ret
	}
	return *o.SEzsignfoldertypeNameX
}

// GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzdoctemplatedocumentListElement) GetSEzsignfoldertypeNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		return nil, false
	}
	return o.SEzsignfoldertypeNameX, true
}

// HasSEzsignfoldertypeNameX returns a boolean if a field has been set.
func (o *EzdoctemplatedocumentListElement) HasSEzsignfoldertypeNameX() bool {
	if o != nil && !IsNil(o.SEzsignfoldertypeNameX) {
		return true
	}

	return false
}

// SetSEzsignfoldertypeNameX gets a reference to the given string and assigns it to the SEzsignfoldertypeNameX field.
func (o *EzdoctemplatedocumentListElement) SetSEzsignfoldertypeNameX(v string) {
	o.SEzsignfoldertypeNameX = &v
}

// GetSEzdoctemplatetypeDescriptionX returns the SEzdoctemplatetypeDescriptionX field value if set, zero value otherwise.
func (o *EzdoctemplatedocumentListElement) GetSEzdoctemplatetypeDescriptionX() string {
	if o == nil || IsNil(o.SEzdoctemplatetypeDescriptionX) {
		var ret string
		return ret
	}
	return *o.SEzdoctemplatetypeDescriptionX
}

// GetSEzdoctemplatetypeDescriptionXOk returns a tuple with the SEzdoctemplatetypeDescriptionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzdoctemplatedocumentListElement) GetSEzdoctemplatetypeDescriptionXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzdoctemplatetypeDescriptionX) {
		return nil, false
	}
	return o.SEzdoctemplatetypeDescriptionX, true
}

// HasSEzdoctemplatetypeDescriptionX returns a boolean if a field has been set.
func (o *EzdoctemplatedocumentListElement) HasSEzdoctemplatetypeDescriptionX() bool {
	if o != nil && !IsNil(o.SEzdoctemplatetypeDescriptionX) {
		return true
	}

	return false
}

// SetSEzdoctemplatetypeDescriptionX gets a reference to the given string and assigns it to the SEzdoctemplatetypeDescriptionX field.
func (o *EzdoctemplatedocumentListElement) SetSEzdoctemplatetypeDescriptionX(v string) {
	o.SEzdoctemplatetypeDescriptionX = &v
}

// GetSEzdoctemplatefieldtypecategoryDescriptionX returns the SEzdoctemplatefieldtypecategoryDescriptionX field value if set, zero value otherwise.
func (o *EzdoctemplatedocumentListElement) GetSEzdoctemplatefieldtypecategoryDescriptionX() string {
	if o == nil || IsNil(o.SEzdoctemplatefieldtypecategoryDescriptionX) {
		var ret string
		return ret
	}
	return *o.SEzdoctemplatefieldtypecategoryDescriptionX
}

// GetSEzdoctemplatefieldtypecategoryDescriptionXOk returns a tuple with the SEzdoctemplatefieldtypecategoryDescriptionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzdoctemplatedocumentListElement) GetSEzdoctemplatefieldtypecategoryDescriptionXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzdoctemplatefieldtypecategoryDescriptionX) {
		return nil, false
	}
	return o.SEzdoctemplatefieldtypecategoryDescriptionX, true
}

// HasSEzdoctemplatefieldtypecategoryDescriptionX returns a boolean if a field has been set.
func (o *EzdoctemplatedocumentListElement) HasSEzdoctemplatefieldtypecategoryDescriptionX() bool {
	if o != nil && !IsNil(o.SEzdoctemplatefieldtypecategoryDescriptionX) {
		return true
	}

	return false
}

// SetSEzdoctemplatefieldtypecategoryDescriptionX gets a reference to the given string and assigns it to the SEzdoctemplatefieldtypecategoryDescriptionX field.
func (o *EzdoctemplatedocumentListElement) SetSEzdoctemplatefieldtypecategoryDescriptionX(v string) {
	o.SEzdoctemplatefieldtypecategoryDescriptionX = &v
}

// GetEEzdoctemplatedocumentPrivacylevel returns the EEzdoctemplatedocumentPrivacylevel field value if set, zero value otherwise.
func (o *EzdoctemplatedocumentListElement) GetEEzdoctemplatedocumentPrivacylevel() FieldEEzdoctemplatedocumentPrivacylevel {
	if o == nil || IsNil(o.EEzdoctemplatedocumentPrivacylevel) {
		var ret FieldEEzdoctemplatedocumentPrivacylevel
		return ret
	}
	return *o.EEzdoctemplatedocumentPrivacylevel
}

// GetEEzdoctemplatedocumentPrivacylevelOk returns a tuple with the EEzdoctemplatedocumentPrivacylevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzdoctemplatedocumentListElement) GetEEzdoctemplatedocumentPrivacylevelOk() (*FieldEEzdoctemplatedocumentPrivacylevel, bool) {
	if o == nil || IsNil(o.EEzdoctemplatedocumentPrivacylevel) {
		return nil, false
	}
	return o.EEzdoctemplatedocumentPrivacylevel, true
}

// HasEEzdoctemplatedocumentPrivacylevel returns a boolean if a field has been set.
func (o *EzdoctemplatedocumentListElement) HasEEzdoctemplatedocumentPrivacylevel() bool {
	if o != nil && !IsNil(o.EEzdoctemplatedocumentPrivacylevel) {
		return true
	}

	return false
}

// SetEEzdoctemplatedocumentPrivacylevel gets a reference to the given FieldEEzdoctemplatedocumentPrivacylevel and assigns it to the EEzdoctemplatedocumentPrivacylevel field.
func (o *EzdoctemplatedocumentListElement) SetEEzdoctemplatedocumentPrivacylevel(v FieldEEzdoctemplatedocumentPrivacylevel) {
	o.EEzdoctemplatedocumentPrivacylevel = &v
}

// GetBEzdoctemplatedocumentIsactive returns the BEzdoctemplatedocumentIsactive field value
func (o *EzdoctemplatedocumentListElement) GetBEzdoctemplatedocumentIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzdoctemplatedocumentIsactive
}

// GetBEzdoctemplatedocumentIsactiveOk returns a tuple with the BEzdoctemplatedocumentIsactive field value
// and a boolean to check if the value has been set.
func (o *EzdoctemplatedocumentListElement) GetBEzdoctemplatedocumentIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzdoctemplatedocumentIsactive, true
}

// SetBEzdoctemplatedocumentIsactive sets field value
func (o *EzdoctemplatedocumentListElement) SetBEzdoctemplatedocumentIsactive(v bool) {
	o.BEzdoctemplatedocumentIsactive = v
}

// GetSEzdoctemplatedocumentNameX returns the SEzdoctemplatedocumentNameX field value
func (o *EzdoctemplatedocumentListElement) GetSEzdoctemplatedocumentNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzdoctemplatedocumentNameX
}

// GetSEzdoctemplatedocumentNameXOk returns a tuple with the SEzdoctemplatedocumentNameX field value
// and a boolean to check if the value has been set.
func (o *EzdoctemplatedocumentListElement) GetSEzdoctemplatedocumentNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzdoctemplatedocumentNameX, true
}

// SetSEzdoctemplatedocumentNameX sets field value
func (o *EzdoctemplatedocumentListElement) SetSEzdoctemplatedocumentNameX(v string) {
	o.SEzdoctemplatedocumentNameX = v
}

func (o EzdoctemplatedocumentListElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzdoctemplatedocumentListElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzdoctemplatedocumentID"] = o.PkiEzdoctemplatedocumentID
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	if !IsNil(o.FkiEzsignfoldertypeID) {
		toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	}
	toSerialize["fkiEzdoctemplatetypeID"] = o.FkiEzdoctemplatetypeID
	toSerialize["fkiEzdoctemplatefieldtypecategoryID"] = o.FkiEzdoctemplatefieldtypecategoryID
	if !IsNil(o.SEzsignfoldertypeNameX) {
		toSerialize["sEzsignfoldertypeNameX"] = o.SEzsignfoldertypeNameX
	}
	if !IsNil(o.SEzdoctemplatetypeDescriptionX) {
		toSerialize["sEzdoctemplatetypeDescriptionX"] = o.SEzdoctemplatetypeDescriptionX
	}
	if !IsNil(o.SEzdoctemplatefieldtypecategoryDescriptionX) {
		toSerialize["sEzdoctemplatefieldtypecategoryDescriptionX"] = o.SEzdoctemplatefieldtypecategoryDescriptionX
	}
	if !IsNil(o.EEzdoctemplatedocumentPrivacylevel) {
		toSerialize["eEzdoctemplatedocumentPrivacylevel"] = o.EEzdoctemplatedocumentPrivacylevel
	}
	toSerialize["bEzdoctemplatedocumentIsactive"] = o.BEzdoctemplatedocumentIsactive
	toSerialize["sEzdoctemplatedocumentNameX"] = o.SEzdoctemplatedocumentNameX
	return toSerialize, nil
}

func (o *EzdoctemplatedocumentListElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzdoctemplatedocumentID",
		"fkiLanguageID",
		"fkiEzdoctemplatetypeID",
		"fkiEzdoctemplatefieldtypecategoryID",
		"bEzdoctemplatedocumentIsactive",
		"sEzdoctemplatedocumentNameX",
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

	varEzdoctemplatedocumentListElement := _EzdoctemplatedocumentListElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzdoctemplatedocumentListElement)

	if err != nil {
		return err
	}

	*o = EzdoctemplatedocumentListElement(varEzdoctemplatedocumentListElement)

	return err
}

type NullableEzdoctemplatedocumentListElement struct {
	value *EzdoctemplatedocumentListElement
	isSet bool
}

func (v NullableEzdoctemplatedocumentListElement) Get() *EzdoctemplatedocumentListElement {
	return v.value
}

func (v *NullableEzdoctemplatedocumentListElement) Set(val *EzdoctemplatedocumentListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableEzdoctemplatedocumentListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableEzdoctemplatedocumentListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzdoctemplatedocumentListElement(val *EzdoctemplatedocumentListElement) *NullableEzdoctemplatedocumentListElement {
	return &NullableEzdoctemplatedocumentListElement{value: val, isSet: true}
}

func (v NullableEzdoctemplatedocumentListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzdoctemplatedocumentListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


