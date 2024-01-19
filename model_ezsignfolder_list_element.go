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
	"bytes"
	"fmt"
)

// checks if the EzsignfolderListElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderListElement{}

// EzsignfolderListElement An Ezsignfolder List Element
type EzsignfolderListElement struct {
	// The unique ID of the Ezsignfolder
	PkiEzsignfolderID int32 `json:"pkiEzsignfolderID"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID int32 `json:"fkiEzsignfoldertypeID"`
	EEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel `json:"eEzsignfoldertypePrivacylevel"`
	// The name of the Ezsignfoldertype in the language of the requester
	SEzsignfoldertypeNameX string `json:"sEzsignfoldertypeNameX"`
	// The description of the Ezsignfolder
	SEzsignfolderDescription string `json:"sEzsignfolderDescription"`
	EEzsignfolderStep FieldEEzsignfolderStep `json:"eEzsignfolderStep"`
	// The date and time at which the object was created
	DtCreatedDate string `json:"dtCreatedDate"`
	// The date and time at which the Ezsignfolder will be sent in the future.
	DtEzsignfolderDelayedsenddate *string `json:"dtEzsignfolderDelayedsenddate,omitempty"`
	// The date and time at which the Ezsignfolder was sent the last time.
	DtEzsignfolderSentdate *string `json:"dtEzsignfolderSentdate,omitempty"`
	// The maximum date and time at which the Ezsignfolder can be signed.
	DtEzsignfolderDuedate *string `json:"dtEzsignfolderDuedate,omitempty"`
	// The total number of Ezsigndocument in the folder
	IEzsigndocument int32 `json:"iEzsigndocument"`
	// The total number of Ezsigndocument in the folder that were saved in the edm system
	IEzsigndocumentEdm int32 `json:"iEzsigndocumentEdm"`
	// The total number of signature blocks in all Ezsigndocuments in the folder
	IEzsignsignature int32 `json:"iEzsignsignature"`
	// The total number of already signed signature blocks in all Ezsigndocuments in the folder
	IEzsignsignatureSigned int32 `json:"iEzsignsignatureSigned"`
}

type _EzsignfolderListElement EzsignfolderListElement

// NewEzsignfolderListElement instantiates a new EzsignfolderListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderListElement(pkiEzsignfolderID int32, fkiEzsignfoldertypeID int32, eEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel, sEzsignfoldertypeNameX string, sEzsignfolderDescription string, eEzsignfolderStep FieldEEzsignfolderStep, dtCreatedDate string, iEzsigndocument int32, iEzsigndocumentEdm int32, iEzsignsignature int32, iEzsignsignatureSigned int32) *EzsignfolderListElement {
	this := EzsignfolderListElement{}
	this.PkiEzsignfolderID = pkiEzsignfolderID
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.EEzsignfoldertypePrivacylevel = eEzsignfoldertypePrivacylevel
	this.SEzsignfoldertypeNameX = sEzsignfoldertypeNameX
	this.SEzsignfolderDescription = sEzsignfolderDescription
	this.EEzsignfolderStep = eEzsignfolderStep
	this.DtCreatedDate = dtCreatedDate
	this.IEzsigndocument = iEzsigndocument
	this.IEzsigndocumentEdm = iEzsigndocumentEdm
	this.IEzsignsignature = iEzsignsignature
	this.IEzsignsignatureSigned = iEzsignsignatureSigned
	return &this
}

// NewEzsignfolderListElementWithDefaults instantiates a new EzsignfolderListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderListElementWithDefaults() *EzsignfolderListElement {
	this := EzsignfolderListElement{}
	return &this
}

// GetPkiEzsignfolderID returns the PkiEzsignfolderID field value
func (o *EzsignfolderListElement) GetPkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignfolderID
}

// GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetPkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignfolderID, true
}

// SetPkiEzsignfolderID sets field value
func (o *EzsignfolderListElement) SetPkiEzsignfolderID(v int32) {
	o.PkiEzsignfolderID = v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value
func (o *EzsignfolderListElement) GetFkiEzsignfoldertypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfoldertypeID, true
}

// SetFkiEzsignfoldertypeID sets field value
func (o *EzsignfolderListElement) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = v
}

// GetEEzsignfoldertypePrivacylevel returns the EEzsignfoldertypePrivacylevel field value
func (o *EzsignfolderListElement) GetEEzsignfoldertypePrivacylevel() FieldEEzsignfoldertypePrivacylevel {
	if o == nil {
		var ret FieldEEzsignfoldertypePrivacylevel
		return ret
	}

	return o.EEzsignfoldertypePrivacylevel
}

// GetEEzsignfoldertypePrivacylevelOk returns a tuple with the EEzsignfoldertypePrivacylevel field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetEEzsignfoldertypePrivacylevelOk() (*FieldEEzsignfoldertypePrivacylevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsignfoldertypePrivacylevel, true
}

// SetEEzsignfoldertypePrivacylevel sets field value
func (o *EzsignfolderListElement) SetEEzsignfoldertypePrivacylevel(v FieldEEzsignfoldertypePrivacylevel) {
	o.EEzsignfoldertypePrivacylevel = v
}

// GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field value
func (o *EzsignfolderListElement) GetSEzsignfoldertypeNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfoldertypeNameX
}

// GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetSEzsignfoldertypeNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignfoldertypeNameX, true
}

// SetSEzsignfoldertypeNameX sets field value
func (o *EzsignfolderListElement) SetSEzsignfoldertypeNameX(v string) {
	o.SEzsignfoldertypeNameX = v
}

// GetSEzsignfolderDescription returns the SEzsignfolderDescription field value
func (o *EzsignfolderListElement) GetSEzsignfolderDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfolderDescription
}

// GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetSEzsignfolderDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignfolderDescription, true
}

// SetSEzsignfolderDescription sets field value
func (o *EzsignfolderListElement) SetSEzsignfolderDescription(v string) {
	o.SEzsignfolderDescription = v
}

// GetEEzsignfolderStep returns the EEzsignfolderStep field value
func (o *EzsignfolderListElement) GetEEzsignfolderStep() FieldEEzsignfolderStep {
	if o == nil {
		var ret FieldEEzsignfolderStep
		return ret
	}

	return o.EEzsignfolderStep
}

// GetEEzsignfolderStepOk returns a tuple with the EEzsignfolderStep field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetEEzsignfolderStepOk() (*FieldEEzsignfolderStep, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsignfolderStep, true
}

// SetEEzsignfolderStep sets field value
func (o *EzsignfolderListElement) SetEEzsignfolderStep(v FieldEEzsignfolderStep) {
	o.EEzsignfolderStep = v
}

// GetDtCreatedDate returns the DtCreatedDate field value
func (o *EzsignfolderListElement) GetDtCreatedDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtCreatedDate
}

// GetDtCreatedDateOk returns a tuple with the DtCreatedDate field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetDtCreatedDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtCreatedDate, true
}

// SetDtCreatedDate sets field value
func (o *EzsignfolderListElement) SetDtCreatedDate(v string) {
	o.DtCreatedDate = v
}

// GetDtEzsignfolderDelayedsenddate returns the DtEzsignfolderDelayedsenddate field value if set, zero value otherwise.
func (o *EzsignfolderListElement) GetDtEzsignfolderDelayedsenddate() string {
	if o == nil || IsNil(o.DtEzsignfolderDelayedsenddate) {
		var ret string
		return ret
	}
	return *o.DtEzsignfolderDelayedsenddate
}

// GetDtEzsignfolderDelayedsenddateOk returns a tuple with the DtEzsignfolderDelayedsenddate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetDtEzsignfolderDelayedsenddateOk() (*string, bool) {
	if o == nil || IsNil(o.DtEzsignfolderDelayedsenddate) {
		return nil, false
	}
	return o.DtEzsignfolderDelayedsenddate, true
}

// HasDtEzsignfolderDelayedsenddate returns a boolean if a field has been set.
func (o *EzsignfolderListElement) HasDtEzsignfolderDelayedsenddate() bool {
	if o != nil && !IsNil(o.DtEzsignfolderDelayedsenddate) {
		return true
	}

	return false
}

// SetDtEzsignfolderDelayedsenddate gets a reference to the given string and assigns it to the DtEzsignfolderDelayedsenddate field.
func (o *EzsignfolderListElement) SetDtEzsignfolderDelayedsenddate(v string) {
	o.DtEzsignfolderDelayedsenddate = &v
}

// GetDtEzsignfolderSentdate returns the DtEzsignfolderSentdate field value if set, zero value otherwise.
func (o *EzsignfolderListElement) GetDtEzsignfolderSentdate() string {
	if o == nil || IsNil(o.DtEzsignfolderSentdate) {
		var ret string
		return ret
	}
	return *o.DtEzsignfolderSentdate
}

// GetDtEzsignfolderSentdateOk returns a tuple with the DtEzsignfolderSentdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetDtEzsignfolderSentdateOk() (*string, bool) {
	if o == nil || IsNil(o.DtEzsignfolderSentdate) {
		return nil, false
	}
	return o.DtEzsignfolderSentdate, true
}

// HasDtEzsignfolderSentdate returns a boolean if a field has been set.
func (o *EzsignfolderListElement) HasDtEzsignfolderSentdate() bool {
	if o != nil && !IsNil(o.DtEzsignfolderSentdate) {
		return true
	}

	return false
}

// SetDtEzsignfolderSentdate gets a reference to the given string and assigns it to the DtEzsignfolderSentdate field.
func (o *EzsignfolderListElement) SetDtEzsignfolderSentdate(v string) {
	o.DtEzsignfolderSentdate = &v
}

// GetDtEzsignfolderDuedate returns the DtEzsignfolderDuedate field value if set, zero value otherwise.
func (o *EzsignfolderListElement) GetDtEzsignfolderDuedate() string {
	if o == nil || IsNil(o.DtEzsignfolderDuedate) {
		var ret string
		return ret
	}
	return *o.DtEzsignfolderDuedate
}

// GetDtEzsignfolderDuedateOk returns a tuple with the DtEzsignfolderDuedate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetDtEzsignfolderDuedateOk() (*string, bool) {
	if o == nil || IsNil(o.DtEzsignfolderDuedate) {
		return nil, false
	}
	return o.DtEzsignfolderDuedate, true
}

// HasDtEzsignfolderDuedate returns a boolean if a field has been set.
func (o *EzsignfolderListElement) HasDtEzsignfolderDuedate() bool {
	if o != nil && !IsNil(o.DtEzsignfolderDuedate) {
		return true
	}

	return false
}

// SetDtEzsignfolderDuedate gets a reference to the given string and assigns it to the DtEzsignfolderDuedate field.
func (o *EzsignfolderListElement) SetDtEzsignfolderDuedate(v string) {
	o.DtEzsignfolderDuedate = &v
}

// GetIEzsigndocument returns the IEzsigndocument field value
func (o *EzsignfolderListElement) GetIEzsigndocument() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigndocument
}

// GetIEzsigndocumentOk returns a tuple with the IEzsigndocument field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetIEzsigndocumentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigndocument, true
}

// SetIEzsigndocument sets field value
func (o *EzsignfolderListElement) SetIEzsigndocument(v int32) {
	o.IEzsigndocument = v
}

// GetIEzsigndocumentEdm returns the IEzsigndocumentEdm field value
func (o *EzsignfolderListElement) GetIEzsigndocumentEdm() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigndocumentEdm
}

// GetIEzsigndocumentEdmOk returns a tuple with the IEzsigndocumentEdm field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetIEzsigndocumentEdmOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigndocumentEdm, true
}

// SetIEzsigndocumentEdm sets field value
func (o *EzsignfolderListElement) SetIEzsigndocumentEdm(v int32) {
	o.IEzsigndocumentEdm = v
}

// GetIEzsignsignature returns the IEzsignsignature field value
func (o *EzsignfolderListElement) GetIEzsignsignature() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignature
}

// GetIEzsignsignatureOk returns a tuple with the IEzsignsignature field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetIEzsignsignatureOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignsignature, true
}

// SetIEzsignsignature sets field value
func (o *EzsignfolderListElement) SetIEzsignsignature(v int32) {
	o.IEzsignsignature = v
}

// GetIEzsignsignatureSigned returns the IEzsignsignatureSigned field value
func (o *EzsignfolderListElement) GetIEzsignsignatureSigned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignatureSigned
}

// GetIEzsignsignatureSignedOk returns a tuple with the IEzsignsignatureSigned field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderListElement) GetIEzsignsignatureSignedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignsignatureSigned, true
}

// SetIEzsignsignatureSigned sets field value
func (o *EzsignfolderListElement) SetIEzsignsignatureSigned(v int32) {
	o.IEzsignsignatureSigned = v
}

func (o EzsignfolderListElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderListElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignfolderID"] = o.PkiEzsignfolderID
	toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	toSerialize["eEzsignfoldertypePrivacylevel"] = o.EEzsignfoldertypePrivacylevel
	toSerialize["sEzsignfoldertypeNameX"] = o.SEzsignfoldertypeNameX
	toSerialize["sEzsignfolderDescription"] = o.SEzsignfolderDescription
	toSerialize["eEzsignfolderStep"] = o.EEzsignfolderStep
	toSerialize["dtCreatedDate"] = o.DtCreatedDate
	if !IsNil(o.DtEzsignfolderDelayedsenddate) {
		toSerialize["dtEzsignfolderDelayedsenddate"] = o.DtEzsignfolderDelayedsenddate
	}
	if !IsNil(o.DtEzsignfolderSentdate) {
		toSerialize["dtEzsignfolderSentdate"] = o.DtEzsignfolderSentdate
	}
	if !IsNil(o.DtEzsignfolderDuedate) {
		toSerialize["dtEzsignfolderDuedate"] = o.DtEzsignfolderDuedate
	}
	toSerialize["iEzsigndocument"] = o.IEzsigndocument
	toSerialize["iEzsigndocumentEdm"] = o.IEzsigndocumentEdm
	toSerialize["iEzsignsignature"] = o.IEzsignsignature
	toSerialize["iEzsignsignatureSigned"] = o.IEzsignsignatureSigned
	return toSerialize, nil
}

func (o *EzsignfolderListElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignfolderID",
		"fkiEzsignfoldertypeID",
		"eEzsignfoldertypePrivacylevel",
		"sEzsignfoldertypeNameX",
		"sEzsignfolderDescription",
		"eEzsignfolderStep",
		"dtCreatedDate",
		"iEzsigndocument",
		"iEzsigndocumentEdm",
		"iEzsignsignature",
		"iEzsignsignatureSigned",
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

	varEzsignfolderListElement := _EzsignfolderListElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfolderListElement)

	if err != nil {
		return err
	}

	*o = EzsignfolderListElement(varEzsignfolderListElement)

	return err
}

type NullableEzsignfolderListElement struct {
	value *EzsignfolderListElement
	isSet bool
}

func (v NullableEzsignfolderListElement) Get() *EzsignfolderListElement {
	return v.value
}

func (v *NullableEzsignfolderListElement) Set(val *EzsignfolderListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderListElement(val *EzsignfolderListElement) *NullableEzsignfolderListElement {
	return &NullableEzsignfolderListElement{value: val, isSet: true}
}

func (v NullableEzsignfolderListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


