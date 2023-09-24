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

// checks if the EzsignsignerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignerResponse{}

// EzsignsignerResponse An Ezsignsigner Object
type EzsignsignerResponse struct {
	// The unique ID of the Ezsignsigner
	PkiEzsignsignerID int32 `json:"pkiEzsignsignerID"`
	// The unique ID of the Taxassignment.  Valid values:  |Value|Description| |-|-| |1|No tax| |2|GST| |3|HST (ON)| |4|HST (NB)| |5|HST (NS)| |6|HST (NL)| |7|HST (PE)| |8|GST + QST (QC)| |9|GST + QST (QC) Non-Recoverable| |10|GST + PST (BC)| |11|GST + PST (SK)| |12|GST + RST (MB)| |13|GST + PST (BC) Non-Recoverable| |14|GST + PST (SK) Non-Recoverable| |15|GST + RST (MB) Non-Recoverable|
	FkiTaxassignmentID int32 `json:"fkiTaxassignmentID"`
	// The unique ID of the Secretquestion.  Valid values:  |Value|Description| |-|-| |1|The name of the hospital in which you were born| |2|The name of your grade school| |3|The last name of your favorite teacher| |4|Your favorite sports team| |5|Your favorite TV show| |6|Your favorite movie| |7|The name of the street on which you grew up| |8|The name of your first employer| |9|Your first car| |10|Your favorite food| |11|The name of your first pet| |12|Favorite musician/band| |13|What instrument you play| |14|Your father's middle name| |15|Your mother's maiden name| |16|Name of your eldest child| |17|Your spouse's middle name| |18|Favorite restaurant| |19|Childhood nickname| |20|Favorite vacation destination| |21|Your boat's name| |22|Date of Birth (YYYY-MM-DD)|
	FkiSecretquestionID *int32 `json:"fkiSecretquestionID,omitempty"`
	// The unique ID of the Userlogintype  Valid values:  |Value|Description|Detail| |-|-|-| |1|**Email Only**|The Ezsignsigner will receive a secure link by email| |2|**Email and phone or SMS**|The Ezsignsigner will receive a secure link by email and will need to authenticate using SMS or Phone call. **Additional fee applies**| |3|**Email and secret question**|The Ezsignsigner will receive a secure link by email and will need to authenticate using a predefined question and answer| |4|**In person only**|The Ezsignsigner will only be able to sign \"In-Person\" and there won't be any authentication. No email will be sent for invitation to sign. Make sure you evaluate the risk of signature denial and at minimum, we recommend you use a handwritten signature type| |5|**In person with phone or SMS**|The Ezsignsigner will only be able to sign \"In-Person\" and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**|
	FkiUserlogintypeID int32 `json:"fkiUserlogintypeID"`
	// The description of the Userlogintype in the language of the requester
	SUserlogintypeDescriptionX string `json:"sUserlogintypeDescriptionX"`
}

// NewEzsignsignerResponse instantiates a new EzsignsignerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignerResponse(pkiEzsignsignerID int32, fkiTaxassignmentID int32, fkiUserlogintypeID int32, sUserlogintypeDescriptionX string) *EzsignsignerResponse {
	this := EzsignsignerResponse{}
	this.PkiEzsignsignerID = pkiEzsignsignerID
	this.FkiTaxassignmentID = fkiTaxassignmentID
	this.FkiUserlogintypeID = fkiUserlogintypeID
	this.SUserlogintypeDescriptionX = sUserlogintypeDescriptionX
	return &this
}

// NewEzsignsignerResponseWithDefaults instantiates a new EzsignsignerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignerResponseWithDefaults() *EzsignsignerResponse {
	this := EzsignsignerResponse{}
	return &this
}

// GetPkiEzsignsignerID returns the PkiEzsignsignerID field value
func (o *EzsignsignerResponse) GetPkiEzsignsignerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignsignerID
}

// GetPkiEzsignsignerIDOk returns a tuple with the PkiEzsignsignerID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignerResponse) GetPkiEzsignsignerIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignsignerID, true
}

// SetPkiEzsignsignerID sets field value
func (o *EzsignsignerResponse) SetPkiEzsignsignerID(v int32) {
	o.PkiEzsignsignerID = v
}

// GetFkiTaxassignmentID returns the FkiTaxassignmentID field value
func (o *EzsignsignerResponse) GetFkiTaxassignmentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiTaxassignmentID
}

// GetFkiTaxassignmentIDOk returns a tuple with the FkiTaxassignmentID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignerResponse) GetFkiTaxassignmentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiTaxassignmentID, true
}

// SetFkiTaxassignmentID sets field value
func (o *EzsignsignerResponse) SetFkiTaxassignmentID(v int32) {
	o.FkiTaxassignmentID = v
}

// GetFkiSecretquestionID returns the FkiSecretquestionID field value if set, zero value otherwise.
func (o *EzsignsignerResponse) GetFkiSecretquestionID() int32 {
	if o == nil || IsNil(o.FkiSecretquestionID) {
		var ret int32
		return ret
	}
	return *o.FkiSecretquestionID
}

// GetFkiSecretquestionIDOk returns a tuple with the FkiSecretquestionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignerResponse) GetFkiSecretquestionIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiSecretquestionID) {
		return nil, false
	}
	return o.FkiSecretquestionID, true
}

// HasFkiSecretquestionID returns a boolean if a field has been set.
func (o *EzsignsignerResponse) HasFkiSecretquestionID() bool {
	if o != nil && !IsNil(o.FkiSecretquestionID) {
		return true
	}

	return false
}

// SetFkiSecretquestionID gets a reference to the given int32 and assigns it to the FkiSecretquestionID field.
func (o *EzsignsignerResponse) SetFkiSecretquestionID(v int32) {
	o.FkiSecretquestionID = &v
}

// GetFkiUserlogintypeID returns the FkiUserlogintypeID field value
func (o *EzsignsignerResponse) GetFkiUserlogintypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUserlogintypeID
}

// GetFkiUserlogintypeIDOk returns a tuple with the FkiUserlogintypeID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignerResponse) GetFkiUserlogintypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiUserlogintypeID, true
}

// SetFkiUserlogintypeID sets field value
func (o *EzsignsignerResponse) SetFkiUserlogintypeID(v int32) {
	o.FkiUserlogintypeID = v
}

// GetSUserlogintypeDescriptionX returns the SUserlogintypeDescriptionX field value
func (o *EzsignsignerResponse) GetSUserlogintypeDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserlogintypeDescriptionX
}

// GetSUserlogintypeDescriptionXOk returns a tuple with the SUserlogintypeDescriptionX field value
// and a boolean to check if the value has been set.
func (o *EzsignsignerResponse) GetSUserlogintypeDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserlogintypeDescriptionX, true
}

// SetSUserlogintypeDescriptionX sets field value
func (o *EzsignsignerResponse) SetSUserlogintypeDescriptionX(v string) {
	o.SUserlogintypeDescriptionX = v
}

func (o EzsignsignerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignsignerID"] = o.PkiEzsignsignerID
	toSerialize["fkiTaxassignmentID"] = o.FkiTaxassignmentID
	if !IsNil(o.FkiSecretquestionID) {
		toSerialize["fkiSecretquestionID"] = o.FkiSecretquestionID
	}
	toSerialize["fkiUserlogintypeID"] = o.FkiUserlogintypeID
	toSerialize["sUserlogintypeDescriptionX"] = o.SUserlogintypeDescriptionX
	return toSerialize, nil
}

type NullableEzsignsignerResponse struct {
	value *EzsignsignerResponse
	isSet bool
}

func (v NullableEzsignsignerResponse) Get() *EzsignsignerResponse {
	return v.value
}

func (v *NullableEzsignsignerResponse) Set(val *EzsignsignerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignerResponse(val *EzsignsignerResponse) *NullableEzsignsignerResponse {
	return &NullableEzsignsignerResponse{value: val, isSet: true}
}

func (v NullableEzsignsignerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


