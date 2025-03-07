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

// checks if the CommunicationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunicationRequest{}

// CommunicationRequest Request for POST /1/object/communication
type CommunicationRequest struct {
	// The unique ID of the Communication.
	PkiCommunicationID *int32 `json:"pkiCommunicationID,omitempty"`
	ECommunicationImportance *FieldECommunicationImportance `json:"eCommunicationImportance,omitempty"`
	ECommunicationType FieldECommunicationType `json:"eCommunicationType"`
	ObjCommunicationsender *CustomCommunicationsenderRequest `json:"objCommunicationsender,omitempty"`
	// The subject of the Communication
	SCommunicationSubject *string `json:"sCommunicationSubject,omitempty" validate:"regexp=^.{0,200}$"`
	// The Body of the Communication
	TCommunicationBody string `json:"tCommunicationBody"`
	// Whether the Communication is private or not
	BCommunicationPrivate bool `json:"bCommunicationPrivate"`
	// How the attachment should be included in the email.   Only used if eCommunicationType is **Email**
	ECommunicationAttachmenttype *string `json:"eCommunicationAttachmenttype,omitempty"`
	// The number of days before the attachment link expired.   Only used if eCommunicationType is **Email** and eCommunicationattachmentType is **Link**
	ICommunicationAttachmentlinkexpiration *int32 `json:"iCommunicationAttachmentlinkexpiration,omitempty"`
	// Whether we ask for a read receipt or not.
	BCommunicationReadreceipt *bool `json:"bCommunicationReadreceipt,omitempty"`
}

type _CommunicationRequest CommunicationRequest

// NewCommunicationRequest instantiates a new CommunicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunicationRequest(eCommunicationType FieldECommunicationType, tCommunicationBody string, bCommunicationPrivate bool) *CommunicationRequest {
	this := CommunicationRequest{}
	this.ECommunicationType = eCommunicationType
	this.TCommunicationBody = tCommunicationBody
	this.BCommunicationPrivate = bCommunicationPrivate
	return &this
}

// NewCommunicationRequestWithDefaults instantiates a new CommunicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunicationRequestWithDefaults() *CommunicationRequest {
	this := CommunicationRequest{}
	return &this
}

// GetPkiCommunicationID returns the PkiCommunicationID field value if set, zero value otherwise.
func (o *CommunicationRequest) GetPkiCommunicationID() int32 {
	if o == nil || IsNil(o.PkiCommunicationID) {
		var ret int32
		return ret
	}
	return *o.PkiCommunicationID
}

// GetPkiCommunicationIDOk returns a tuple with the PkiCommunicationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRequest) GetPkiCommunicationIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiCommunicationID) {
		return nil, false
	}
	return o.PkiCommunicationID, true
}

// HasPkiCommunicationID returns a boolean if a field has been set.
func (o *CommunicationRequest) HasPkiCommunicationID() bool {
	if o != nil && !IsNil(o.PkiCommunicationID) {
		return true
	}

	return false
}

// SetPkiCommunicationID gets a reference to the given int32 and assigns it to the PkiCommunicationID field.
func (o *CommunicationRequest) SetPkiCommunicationID(v int32) {
	o.PkiCommunicationID = &v
}

// GetECommunicationImportance returns the ECommunicationImportance field value if set, zero value otherwise.
func (o *CommunicationRequest) GetECommunicationImportance() FieldECommunicationImportance {
	if o == nil || IsNil(o.ECommunicationImportance) {
		var ret FieldECommunicationImportance
		return ret
	}
	return *o.ECommunicationImportance
}

// GetECommunicationImportanceOk returns a tuple with the ECommunicationImportance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRequest) GetECommunicationImportanceOk() (*FieldECommunicationImportance, bool) {
	if o == nil || IsNil(o.ECommunicationImportance) {
		return nil, false
	}
	return o.ECommunicationImportance, true
}

// HasECommunicationImportance returns a boolean if a field has been set.
func (o *CommunicationRequest) HasECommunicationImportance() bool {
	if o != nil && !IsNil(o.ECommunicationImportance) {
		return true
	}

	return false
}

// SetECommunicationImportance gets a reference to the given FieldECommunicationImportance and assigns it to the ECommunicationImportance field.
func (o *CommunicationRequest) SetECommunicationImportance(v FieldECommunicationImportance) {
	o.ECommunicationImportance = &v
}

// GetECommunicationType returns the ECommunicationType field value
func (o *CommunicationRequest) GetECommunicationType() FieldECommunicationType {
	if o == nil {
		var ret FieldECommunicationType
		return ret
	}

	return o.ECommunicationType
}

// GetECommunicationTypeOk returns a tuple with the ECommunicationType field value
// and a boolean to check if the value has been set.
func (o *CommunicationRequest) GetECommunicationTypeOk() (*FieldECommunicationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ECommunicationType, true
}

// SetECommunicationType sets field value
func (o *CommunicationRequest) SetECommunicationType(v FieldECommunicationType) {
	o.ECommunicationType = v
}

// GetObjCommunicationsender returns the ObjCommunicationsender field value if set, zero value otherwise.
func (o *CommunicationRequest) GetObjCommunicationsender() CustomCommunicationsenderRequest {
	if o == nil || IsNil(o.ObjCommunicationsender) {
		var ret CustomCommunicationsenderRequest
		return ret
	}
	return *o.ObjCommunicationsender
}

// GetObjCommunicationsenderOk returns a tuple with the ObjCommunicationsender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRequest) GetObjCommunicationsenderOk() (*CustomCommunicationsenderRequest, bool) {
	if o == nil || IsNil(o.ObjCommunicationsender) {
		return nil, false
	}
	return o.ObjCommunicationsender, true
}

// HasObjCommunicationsender returns a boolean if a field has been set.
func (o *CommunicationRequest) HasObjCommunicationsender() bool {
	if o != nil && !IsNil(o.ObjCommunicationsender) {
		return true
	}

	return false
}

// SetObjCommunicationsender gets a reference to the given CustomCommunicationsenderRequest and assigns it to the ObjCommunicationsender field.
func (o *CommunicationRequest) SetObjCommunicationsender(v CustomCommunicationsenderRequest) {
	o.ObjCommunicationsender = &v
}

// GetSCommunicationSubject returns the SCommunicationSubject field value if set, zero value otherwise.
func (o *CommunicationRequest) GetSCommunicationSubject() string {
	if o == nil || IsNil(o.SCommunicationSubject) {
		var ret string
		return ret
	}
	return *o.SCommunicationSubject
}

// GetSCommunicationSubjectOk returns a tuple with the SCommunicationSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRequest) GetSCommunicationSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.SCommunicationSubject) {
		return nil, false
	}
	return o.SCommunicationSubject, true
}

// HasSCommunicationSubject returns a boolean if a field has been set.
func (o *CommunicationRequest) HasSCommunicationSubject() bool {
	if o != nil && !IsNil(o.SCommunicationSubject) {
		return true
	}

	return false
}

// SetSCommunicationSubject gets a reference to the given string and assigns it to the SCommunicationSubject field.
func (o *CommunicationRequest) SetSCommunicationSubject(v string) {
	o.SCommunicationSubject = &v
}

// GetTCommunicationBody returns the TCommunicationBody field value
func (o *CommunicationRequest) GetTCommunicationBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TCommunicationBody
}

// GetTCommunicationBodyOk returns a tuple with the TCommunicationBody field value
// and a boolean to check if the value has been set.
func (o *CommunicationRequest) GetTCommunicationBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TCommunicationBody, true
}

// SetTCommunicationBody sets field value
func (o *CommunicationRequest) SetTCommunicationBody(v string) {
	o.TCommunicationBody = v
}

// GetBCommunicationPrivate returns the BCommunicationPrivate field value
func (o *CommunicationRequest) GetBCommunicationPrivate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCommunicationPrivate
}

// GetBCommunicationPrivateOk returns a tuple with the BCommunicationPrivate field value
// and a boolean to check if the value has been set.
func (o *CommunicationRequest) GetBCommunicationPrivateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCommunicationPrivate, true
}

// SetBCommunicationPrivate sets field value
func (o *CommunicationRequest) SetBCommunicationPrivate(v bool) {
	o.BCommunicationPrivate = v
}

// GetECommunicationAttachmenttype returns the ECommunicationAttachmenttype field value if set, zero value otherwise.
func (o *CommunicationRequest) GetECommunicationAttachmenttype() string {
	if o == nil || IsNil(o.ECommunicationAttachmenttype) {
		var ret string
		return ret
	}
	return *o.ECommunicationAttachmenttype
}

// GetECommunicationAttachmenttypeOk returns a tuple with the ECommunicationAttachmenttype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRequest) GetECommunicationAttachmenttypeOk() (*string, bool) {
	if o == nil || IsNil(o.ECommunicationAttachmenttype) {
		return nil, false
	}
	return o.ECommunicationAttachmenttype, true
}

// HasECommunicationAttachmenttype returns a boolean if a field has been set.
func (o *CommunicationRequest) HasECommunicationAttachmenttype() bool {
	if o != nil && !IsNil(o.ECommunicationAttachmenttype) {
		return true
	}

	return false
}

// SetECommunicationAttachmenttype gets a reference to the given string and assigns it to the ECommunicationAttachmenttype field.
func (o *CommunicationRequest) SetECommunicationAttachmenttype(v string) {
	o.ECommunicationAttachmenttype = &v
}

// GetICommunicationAttachmentlinkexpiration returns the ICommunicationAttachmentlinkexpiration field value if set, zero value otherwise.
func (o *CommunicationRequest) GetICommunicationAttachmentlinkexpiration() int32 {
	if o == nil || IsNil(o.ICommunicationAttachmentlinkexpiration) {
		var ret int32
		return ret
	}
	return *o.ICommunicationAttachmentlinkexpiration
}

// GetICommunicationAttachmentlinkexpirationOk returns a tuple with the ICommunicationAttachmentlinkexpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRequest) GetICommunicationAttachmentlinkexpirationOk() (*int32, bool) {
	if o == nil || IsNil(o.ICommunicationAttachmentlinkexpiration) {
		return nil, false
	}
	return o.ICommunicationAttachmentlinkexpiration, true
}

// HasICommunicationAttachmentlinkexpiration returns a boolean if a field has been set.
func (o *CommunicationRequest) HasICommunicationAttachmentlinkexpiration() bool {
	if o != nil && !IsNil(o.ICommunicationAttachmentlinkexpiration) {
		return true
	}

	return false
}

// SetICommunicationAttachmentlinkexpiration gets a reference to the given int32 and assigns it to the ICommunicationAttachmentlinkexpiration field.
func (o *CommunicationRequest) SetICommunicationAttachmentlinkexpiration(v int32) {
	o.ICommunicationAttachmentlinkexpiration = &v
}

// GetBCommunicationReadreceipt returns the BCommunicationReadreceipt field value if set, zero value otherwise.
func (o *CommunicationRequest) GetBCommunicationReadreceipt() bool {
	if o == nil || IsNil(o.BCommunicationReadreceipt) {
		var ret bool
		return ret
	}
	return *o.BCommunicationReadreceipt
}

// GetBCommunicationReadreceiptOk returns a tuple with the BCommunicationReadreceipt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRequest) GetBCommunicationReadreceiptOk() (*bool, bool) {
	if o == nil || IsNil(o.BCommunicationReadreceipt) {
		return nil, false
	}
	return o.BCommunicationReadreceipt, true
}

// HasBCommunicationReadreceipt returns a boolean if a field has been set.
func (o *CommunicationRequest) HasBCommunicationReadreceipt() bool {
	if o != nil && !IsNil(o.BCommunicationReadreceipt) {
		return true
	}

	return false
}

// SetBCommunicationReadreceipt gets a reference to the given bool and assigns it to the BCommunicationReadreceipt field.
func (o *CommunicationRequest) SetBCommunicationReadreceipt(v bool) {
	o.BCommunicationReadreceipt = &v
}

func (o CommunicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunicationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiCommunicationID) {
		toSerialize["pkiCommunicationID"] = o.PkiCommunicationID
	}
	if !IsNil(o.ECommunicationImportance) {
		toSerialize["eCommunicationImportance"] = o.ECommunicationImportance
	}
	toSerialize["eCommunicationType"] = o.ECommunicationType
	if !IsNil(o.ObjCommunicationsender) {
		toSerialize["objCommunicationsender"] = o.ObjCommunicationsender
	}
	if !IsNil(o.SCommunicationSubject) {
		toSerialize["sCommunicationSubject"] = o.SCommunicationSubject
	}
	toSerialize["tCommunicationBody"] = o.TCommunicationBody
	toSerialize["bCommunicationPrivate"] = o.BCommunicationPrivate
	if !IsNil(o.ECommunicationAttachmenttype) {
		toSerialize["eCommunicationAttachmenttype"] = o.ECommunicationAttachmenttype
	}
	if !IsNil(o.ICommunicationAttachmentlinkexpiration) {
		toSerialize["iCommunicationAttachmentlinkexpiration"] = o.ICommunicationAttachmentlinkexpiration
	}
	if !IsNil(o.BCommunicationReadreceipt) {
		toSerialize["bCommunicationReadreceipt"] = o.BCommunicationReadreceipt
	}
	return toSerialize, nil
}

func (o *CommunicationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eCommunicationType",
		"tCommunicationBody",
		"bCommunicationPrivate",
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

	varCommunicationRequest := _CommunicationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommunicationRequest)

	if err != nil {
		return err
	}

	*o = CommunicationRequest(varCommunicationRequest)

	return err
}

type NullableCommunicationRequest struct {
	value *CommunicationRequest
	isSet bool
}

func (v NullableCommunicationRequest) Get() *CommunicationRequest {
	return v.value
}

func (v *NullableCommunicationRequest) Set(val *CommunicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationRequest(val *CommunicationRequest) *NullableCommunicationRequest {
	return &NullableCommunicationRequest{value: val, isSet: true}
}

func (v NullableCommunicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


