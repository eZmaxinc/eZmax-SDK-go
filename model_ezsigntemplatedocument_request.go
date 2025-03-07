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

// checks if the EzsigntemplatedocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatedocumentRequest{}

// EzsigntemplatedocumentRequest A Ezsigntemplatedocument Object
type EzsigntemplatedocumentRequest struct {
	// The unique ID of the Ezsigntemplatedocument
	PkiEzsigntemplatedocumentID *int32 `json:"pkiEzsigntemplatedocumentID,omitempty"`
	// The unique ID of the Ezsigntemplate
	FkiEzsigntemplateID int32 `json:"fkiEzsigntemplateID"`
	// The unique ID of the Ezsigndocument
	FkiEzsigndocumentID *int32 `json:"fkiEzsigndocumentID,omitempty"`
	// The unique ID of the Ezsigntemplatesigner
	FkiEzsigntemplatesignerID *int32 `json:"fkiEzsigntemplatesignerID,omitempty"`
	// The name of the Ezsigntemplatedocument.
	SEzsigntemplatedocumentName string `json:"sEzsigntemplatedocumentName"`
	// Indicates where to look for the document binary content.
	EEzsigntemplatedocumentSource string `json:"eEzsigntemplatedocumentSource"`
	// Indicates the format of the template.
	EEzsigntemplatedocumentFormat *string `json:"eEzsigntemplatedocumentFormat,omitempty"`
	// The Base64 encoded binary content of the document.  This field is Required when eEzsigntemplatedocumentSource = Base64.
	SEzsigntemplatedocumentBase64 *string `json:"sEzsigntemplatedocumentBase64,omitempty"`
	// The url where the document content resides.  This field is Required when eEzsigntemplatedocumentSource = Url.
	SEzsigntemplatedocumentUrl *string `json:"sEzsigntemplatedocumentUrl,omitempty"`
	// Try to repair the document or flatten it if it cannot be used for electronic signature.
	BEzsigntemplatedocumentForcerepair *bool `json:"bEzsigntemplatedocumentForcerepair,omitempty"`
	// If the document contains an existing PDF form this property must be set.  **Keep** leaves the form as-is in the document.  **Convert** removes the form and convert all the existing fields to Ezsigntemplateformfieldgroups and assign them to the specified **fkiEzsigntemplatesignerID**  **Discard** removes the form from the document  **Flatten** prints the form values in the document.
	EEzsigntemplatedocumentForm *string `json:"eEzsigntemplatedocumentForm,omitempty"`
	// If the source template is password protected, the password to open/modify it.
	SEzsigntemplatedocumentPassword *string `json:"sEzsigntemplatedocumentPassword,omitempty"`
}

type _EzsigntemplatedocumentRequest EzsigntemplatedocumentRequest

// NewEzsigntemplatedocumentRequest instantiates a new EzsigntemplatedocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatedocumentRequest(fkiEzsigntemplateID int32, sEzsigntemplatedocumentName string, eEzsigntemplatedocumentSource string) *EzsigntemplatedocumentRequest {
	this := EzsigntemplatedocumentRequest{}
	this.FkiEzsigntemplateID = fkiEzsigntemplateID
	this.SEzsigntemplatedocumentName = sEzsigntemplatedocumentName
	this.EEzsigntemplatedocumentSource = eEzsigntemplatedocumentSource
	var sEzsigntemplatedocumentPassword string = ""
	this.SEzsigntemplatedocumentPassword = &sEzsigntemplatedocumentPassword
	return &this
}

// NewEzsigntemplatedocumentRequestWithDefaults instantiates a new EzsigntemplatedocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatedocumentRequestWithDefaults() *EzsigntemplatedocumentRequest {
	this := EzsigntemplatedocumentRequest{}
	var sEzsigntemplatedocumentPassword string = ""
	this.SEzsigntemplatedocumentPassword = &sEzsigntemplatedocumentPassword
	return &this
}

// GetPkiEzsigntemplatedocumentID returns the PkiEzsigntemplatedocumentID field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentRequest) GetPkiEzsigntemplatedocumentID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplatedocumentID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplatedocumentID
}

// GetPkiEzsigntemplatedocumentIDOk returns a tuple with the PkiEzsigntemplatedocumentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentRequest) GetPkiEzsigntemplatedocumentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplatedocumentID) {
		return nil, false
	}
	return o.PkiEzsigntemplatedocumentID, true
}

// HasPkiEzsigntemplatedocumentID returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentRequest) HasPkiEzsigntemplatedocumentID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplatedocumentID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplatedocumentID gets a reference to the given int32 and assigns it to the PkiEzsigntemplatedocumentID field.
func (o *EzsigntemplatedocumentRequest) SetPkiEzsigntemplatedocumentID(v int32) {
	o.PkiEzsigntemplatedocumentID = &v
}

// GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field value
func (o *EzsigntemplatedocumentRequest) GetFkiEzsigntemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplateID
}

// GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentRequest) GetFkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplateID, true
}

// SetFkiEzsigntemplateID sets field value
func (o *EzsigntemplatedocumentRequest) SetFkiEzsigntemplateID(v int32) {
	o.FkiEzsigntemplateID = v
}

// GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentRequest) GetFkiEzsigndocumentID() int32 {
	if o == nil || IsNil(o.FkiEzsigndocumentID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigndocumentID
}

// GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentRequest) GetFkiEzsigndocumentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigndocumentID) {
		return nil, false
	}
	return o.FkiEzsigndocumentID, true
}

// HasFkiEzsigndocumentID returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentRequest) HasFkiEzsigndocumentID() bool {
	if o != nil && !IsNil(o.FkiEzsigndocumentID) {
		return true
	}

	return false
}

// SetFkiEzsigndocumentID gets a reference to the given int32 and assigns it to the FkiEzsigndocumentID field.
func (o *EzsigntemplatedocumentRequest) SetFkiEzsigndocumentID(v int32) {
	o.FkiEzsigndocumentID = &v
}

// GetFkiEzsigntemplatesignerID returns the FkiEzsigntemplatesignerID field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentRequest) GetFkiEzsigntemplatesignerID() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplatesignerID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplatesignerID
}

// GetFkiEzsigntemplatesignerIDOk returns a tuple with the FkiEzsigntemplatesignerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentRequest) GetFkiEzsigntemplatesignerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplatesignerID) {
		return nil, false
	}
	return o.FkiEzsigntemplatesignerID, true
}

// HasFkiEzsigntemplatesignerID returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentRequest) HasFkiEzsigntemplatesignerID() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplatesignerID) {
		return true
	}

	return false
}

// SetFkiEzsigntemplatesignerID gets a reference to the given int32 and assigns it to the FkiEzsigntemplatesignerID field.
func (o *EzsigntemplatedocumentRequest) SetFkiEzsigntemplatesignerID(v int32) {
	o.FkiEzsigntemplatesignerID = &v
}

// GetSEzsigntemplatedocumentName returns the SEzsigntemplatedocumentName field value
func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatedocumentName
}

// GetSEzsigntemplatedocumentNameOk returns a tuple with the SEzsigntemplatedocumentName field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatedocumentName, true
}

// SetSEzsigntemplatedocumentName sets field value
func (o *EzsigntemplatedocumentRequest) SetSEzsigntemplatedocumentName(v string) {
	o.SEzsigntemplatedocumentName = v
}

// GetEEzsigntemplatedocumentSource returns the EEzsigntemplatedocumentSource field value
func (o *EzsigntemplatedocumentRequest) GetEEzsigntemplatedocumentSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EEzsigntemplatedocumentSource
}

// GetEEzsigntemplatedocumentSourceOk returns a tuple with the EEzsigntemplatedocumentSource field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentRequest) GetEEzsigntemplatedocumentSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigntemplatedocumentSource, true
}

// SetEEzsigntemplatedocumentSource sets field value
func (o *EzsigntemplatedocumentRequest) SetEEzsigntemplatedocumentSource(v string) {
	o.EEzsigntemplatedocumentSource = v
}

// GetEEzsigntemplatedocumentFormat returns the EEzsigntemplatedocumentFormat field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentRequest) GetEEzsigntemplatedocumentFormat() string {
	if o == nil || IsNil(o.EEzsigntemplatedocumentFormat) {
		var ret string
		return ret
	}
	return *o.EEzsigntemplatedocumentFormat
}

// GetEEzsigntemplatedocumentFormatOk returns a tuple with the EEzsigntemplatedocumentFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentRequest) GetEEzsigntemplatedocumentFormatOk() (*string, bool) {
	if o == nil || IsNil(o.EEzsigntemplatedocumentFormat) {
		return nil, false
	}
	return o.EEzsigntemplatedocumentFormat, true
}

// HasEEzsigntemplatedocumentFormat returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentRequest) HasEEzsigntemplatedocumentFormat() bool {
	if o != nil && !IsNil(o.EEzsigntemplatedocumentFormat) {
		return true
	}

	return false
}

// SetEEzsigntemplatedocumentFormat gets a reference to the given string and assigns it to the EEzsigntemplatedocumentFormat field.
func (o *EzsigntemplatedocumentRequest) SetEEzsigntemplatedocumentFormat(v string) {
	o.EEzsigntemplatedocumentFormat = &v
}

// GetSEzsigntemplatedocumentBase64 returns the SEzsigntemplatedocumentBase64 field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentBase64() string {
	if o == nil || IsNil(o.SEzsigntemplatedocumentBase64) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplatedocumentBase64
}

// GetSEzsigntemplatedocumentBase64Ok returns a tuple with the SEzsigntemplatedocumentBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentBase64Ok() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplatedocumentBase64) {
		return nil, false
	}
	return o.SEzsigntemplatedocumentBase64, true
}

// HasSEzsigntemplatedocumentBase64 returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentRequest) HasSEzsigntemplatedocumentBase64() bool {
	if o != nil && !IsNil(o.SEzsigntemplatedocumentBase64) {
		return true
	}

	return false
}

// SetSEzsigntemplatedocumentBase64 gets a reference to the given string and assigns it to the SEzsigntemplatedocumentBase64 field.
func (o *EzsigntemplatedocumentRequest) SetSEzsigntemplatedocumentBase64(v string) {
	o.SEzsigntemplatedocumentBase64 = &v
}

// GetSEzsigntemplatedocumentUrl returns the SEzsigntemplatedocumentUrl field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentUrl() string {
	if o == nil || IsNil(o.SEzsigntemplatedocumentUrl) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplatedocumentUrl
}

// GetSEzsigntemplatedocumentUrlOk returns a tuple with the SEzsigntemplatedocumentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplatedocumentUrl) {
		return nil, false
	}
	return o.SEzsigntemplatedocumentUrl, true
}

// HasSEzsigntemplatedocumentUrl returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentRequest) HasSEzsigntemplatedocumentUrl() bool {
	if o != nil && !IsNil(o.SEzsigntemplatedocumentUrl) {
		return true
	}

	return false
}

// SetSEzsigntemplatedocumentUrl gets a reference to the given string and assigns it to the SEzsigntemplatedocumentUrl field.
func (o *EzsigntemplatedocumentRequest) SetSEzsigntemplatedocumentUrl(v string) {
	o.SEzsigntemplatedocumentUrl = &v
}

// GetBEzsigntemplatedocumentForcerepair returns the BEzsigntemplatedocumentForcerepair field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentRequest) GetBEzsigntemplatedocumentForcerepair() bool {
	if o == nil || IsNil(o.BEzsigntemplatedocumentForcerepair) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplatedocumentForcerepair
}

// GetBEzsigntemplatedocumentForcerepairOk returns a tuple with the BEzsigntemplatedocumentForcerepair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentRequest) GetBEzsigntemplatedocumentForcerepairOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplatedocumentForcerepair) {
		return nil, false
	}
	return o.BEzsigntemplatedocumentForcerepair, true
}

// HasBEzsigntemplatedocumentForcerepair returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentRequest) HasBEzsigntemplatedocumentForcerepair() bool {
	if o != nil && !IsNil(o.BEzsigntemplatedocumentForcerepair) {
		return true
	}

	return false
}

// SetBEzsigntemplatedocumentForcerepair gets a reference to the given bool and assigns it to the BEzsigntemplatedocumentForcerepair field.
func (o *EzsigntemplatedocumentRequest) SetBEzsigntemplatedocumentForcerepair(v bool) {
	o.BEzsigntemplatedocumentForcerepair = &v
}

// GetEEzsigntemplatedocumentForm returns the EEzsigntemplatedocumentForm field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentRequest) GetEEzsigntemplatedocumentForm() string {
	if o == nil || IsNil(o.EEzsigntemplatedocumentForm) {
		var ret string
		return ret
	}
	return *o.EEzsigntemplatedocumentForm
}

// GetEEzsigntemplatedocumentFormOk returns a tuple with the EEzsigntemplatedocumentForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentRequest) GetEEzsigntemplatedocumentFormOk() (*string, bool) {
	if o == nil || IsNil(o.EEzsigntemplatedocumentForm) {
		return nil, false
	}
	return o.EEzsigntemplatedocumentForm, true
}

// HasEEzsigntemplatedocumentForm returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentRequest) HasEEzsigntemplatedocumentForm() bool {
	if o != nil && !IsNil(o.EEzsigntemplatedocumentForm) {
		return true
	}

	return false
}

// SetEEzsigntemplatedocumentForm gets a reference to the given string and assigns it to the EEzsigntemplatedocumentForm field.
func (o *EzsigntemplatedocumentRequest) SetEEzsigntemplatedocumentForm(v string) {
	o.EEzsigntemplatedocumentForm = &v
}

// GetSEzsigntemplatedocumentPassword returns the SEzsigntemplatedocumentPassword field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentPassword() string {
	if o == nil || IsNil(o.SEzsigntemplatedocumentPassword) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplatedocumentPassword
}

// GetSEzsigntemplatedocumentPasswordOk returns a tuple with the SEzsigntemplatedocumentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentRequest) GetSEzsigntemplatedocumentPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplatedocumentPassword) {
		return nil, false
	}
	return o.SEzsigntemplatedocumentPassword, true
}

// HasSEzsigntemplatedocumentPassword returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentRequest) HasSEzsigntemplatedocumentPassword() bool {
	if o != nil && !IsNil(o.SEzsigntemplatedocumentPassword) {
		return true
	}

	return false
}

// SetSEzsigntemplatedocumentPassword gets a reference to the given string and assigns it to the SEzsigntemplatedocumentPassword field.
func (o *EzsigntemplatedocumentRequest) SetSEzsigntemplatedocumentPassword(v string) {
	o.SEzsigntemplatedocumentPassword = &v
}

func (o EzsigntemplatedocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatedocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplatedocumentID) {
		toSerialize["pkiEzsigntemplatedocumentID"] = o.PkiEzsigntemplatedocumentID
	}
	toSerialize["fkiEzsigntemplateID"] = o.FkiEzsigntemplateID
	if !IsNil(o.FkiEzsigndocumentID) {
		toSerialize["fkiEzsigndocumentID"] = o.FkiEzsigndocumentID
	}
	if !IsNil(o.FkiEzsigntemplatesignerID) {
		toSerialize["fkiEzsigntemplatesignerID"] = o.FkiEzsigntemplatesignerID
	}
	toSerialize["sEzsigntemplatedocumentName"] = o.SEzsigntemplatedocumentName
	toSerialize["eEzsigntemplatedocumentSource"] = o.EEzsigntemplatedocumentSource
	if !IsNil(o.EEzsigntemplatedocumentFormat) {
		toSerialize["eEzsigntemplatedocumentFormat"] = o.EEzsigntemplatedocumentFormat
	}
	if !IsNil(o.SEzsigntemplatedocumentBase64) {
		toSerialize["sEzsigntemplatedocumentBase64"] = o.SEzsigntemplatedocumentBase64
	}
	if !IsNil(o.SEzsigntemplatedocumentUrl) {
		toSerialize["sEzsigntemplatedocumentUrl"] = o.SEzsigntemplatedocumentUrl
	}
	if !IsNil(o.BEzsigntemplatedocumentForcerepair) {
		toSerialize["bEzsigntemplatedocumentForcerepair"] = o.BEzsigntemplatedocumentForcerepair
	}
	if !IsNil(o.EEzsigntemplatedocumentForm) {
		toSerialize["eEzsigntemplatedocumentForm"] = o.EEzsigntemplatedocumentForm
	}
	if !IsNil(o.SEzsigntemplatedocumentPassword) {
		toSerialize["sEzsigntemplatedocumentPassword"] = o.SEzsigntemplatedocumentPassword
	}
	return toSerialize, nil
}

func (o *EzsigntemplatedocumentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsigntemplateID",
		"sEzsigntemplatedocumentName",
		"eEzsigntemplatedocumentSource",
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

	varEzsigntemplatedocumentRequest := _EzsigntemplatedocumentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatedocumentRequest)

	if err != nil {
		return err
	}

	*o = EzsigntemplatedocumentRequest(varEzsigntemplatedocumentRequest)

	return err
}

type NullableEzsigntemplatedocumentRequest struct {
	value *EzsigntemplatedocumentRequest
	isSet bool
}

func (v NullableEzsigntemplatedocumentRequest) Get() *EzsigntemplatedocumentRequest {
	return v.value
}

func (v *NullableEzsigntemplatedocumentRequest) Set(val *EzsigntemplatedocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatedocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatedocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatedocumentRequest(val *EzsigntemplatedocumentRequest) *NullableEzsigntemplatedocumentRequest {
	return &NullableEzsigntemplatedocumentRequest{value: val, isSet: true}
}

func (v NullableEzsigntemplatedocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatedocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


