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

// checks if the EzsigntemplatepublicResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepublicResponseCompound{}

// EzsigntemplatepublicResponseCompound A Ezsigntemplatepublic Object
type EzsigntemplatepublicResponseCompound struct {
	// The unique ID of the Ezsigntemplatepublic
	PkiEzsigntemplatepublicID int32 `json:"pkiEzsigntemplatepublicID"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID int32 `json:"fkiEzsignfoldertypeID"`
	// The name of the Ezsignfoldertype in the language of the requester
	SEzsignfoldertypeNameX string `json:"sEzsignfoldertypeNameX"`
	// The unique ID of the Userlogintype  Valid values:  |Value|Description|Detail| |-|-|-| |1|**Email Only**|The Ezsignsigner will receive a secure link by email| |2|**Email and phone or SMS**|The Ezsignsigner will receive a secure link by email and will need to authenticate using SMS or Phone call. **Additional fee applies**| |3|**Email and secret question**|The Ezsignsigner will receive a secure link by email and will need to authenticate using a predefined question and answer| |4|**In person only**|The Ezsignsigner will only be able to sign \"In-Person\" and there won't be any authentication. No email will be sent for invitation to sign. Make sure you evaluate the risk of signature denial and at minimum, we recommend you use a handwritten signature type| |5|**In person with phone or SMS**|The Ezsignsigner will only be able to sign \"In-Person\" and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**| |6|**Embedded**|The Ezsignsigner will only be able to sign in the embedded solution. No email will be sent for invitation to sign. **Additional fee applies**|   |7|**Embedded with phone or SMS**|The Ezsignsigner will only be able to sign in the embedded solution and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**|   |8|**No validation**|The Ezsignsigner will not receive an email and won't have to validate his connection using 2 factor. **Additional fee applies**|      |9|**Sms only**|The Ezsignsigner will not receive an email but will will need to authenticate using SMS. **Additional fee applies**|     
	FkiUserlogintypeID int32 `json:"fkiUserlogintypeID"`
	// The description of the Userlogintype in the language of the requester
	SUserlogintypeDescriptionX string `json:"sUserlogintypeDescriptionX"`
	// The unique ID of the Ezsigntemplate
	FkiEzsigntemplateID *int32 `json:"fkiEzsigntemplateID,omitempty"`
	// The unique ID of the Ezsigntemplatepackage
	FkiEzsigntemplatepackageID *int32 `json:"fkiEzsigntemplatepackageID,omitempty"`
	// The description of the Ezsigntemplatepublic
	SEzsigntemplatepublicDescription string `json:"sEzsigntemplatepublicDescription" validate:"regexp=^.{0,80}$"`
	// The referenceid of the Ezsigntemplatepublic
	SEzsigntemplatepublicReferenceid string `json:"sEzsigntemplatepublicReferenceid" validate:"regexp=^.{0,36}$"`
	// Whether the ezsigntemplatepublic is active or not
	BEzsigntemplatepublicIsactive bool `json:"bEzsigntemplatepublicIsactive"`
	// The note of the Ezsigntemplatepublic
	TEzsigntemplatepublicNote string `json:"tEzsigntemplatepublicNote" validate:"regexp=^.{0,65535}$"`
	EEzsigntemplatepublicLimittype FieldEEzsigntemplatepublicLimittype `json:"eEzsigntemplatepublicLimittype"`
	// The limit of the Ezsigntemplatepublic
	IEzsigntemplatepublicLimit int32 `json:"iEzsigntemplatepublicLimit"`
	// The limitexceeded of the Ezsigntemplatepublic
	IEzsigntemplatepublicLimitexceeded int32 `json:"iEzsigntemplatepublicLimitexceeded"`
	// The limitexceededsince of the Ezsigntemplatepublic
	DtEzsigntemplatepublicLimitexceededsince string `json:"dtEzsigntemplatepublicLimitexceededsince" validate:"regexp=^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) ([01]?[0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$"`
	// The url of the Ezsigntemplatepublic  You can add these value as query parameters to prefill the corresponding role  |Parameter|Description| |-|-| |sEzsigntemplatesignerDescription|The role to fill| |sContactFirstname|The contact firstname| |sContactLastname|The contact lastname| |sEmailAddress|The contact email| |sPhoneE164|The contact phone number| |sPhoneE164Cell|The contact cell phone number|
	SEzsigntemplatepublicUrl string `json:"sEzsigntemplatepublicUrl" validate:"regexp=^(https|http):\\/\\/[^\\\\s\\/$.?#].[^\\\\s]*$"`
	// The Ezsigntemplate/Ezsigntemplatepackage description
	SEzsigntemplatepublicEzsigntemplatedescription string `json:"sEzsigntemplatepublicEzsigntemplatedescription" validate:"regexp=^.{1,80}$"`
	ObjAudit *CommonAudit `json:"objAudit,omitempty"`
	AObjEzsignfolderezsigntemplatepublic []CustomEzsignfolderezsigntemplatepublicResponse `json:"a_objEzsignfolderezsigntemplatepublic"`
}

type _EzsigntemplatepublicResponseCompound EzsigntemplatepublicResponseCompound

// NewEzsigntemplatepublicResponseCompound instantiates a new EzsigntemplatepublicResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepublicResponseCompound(pkiEzsigntemplatepublicID int32, fkiEzsignfoldertypeID int32, sEzsignfoldertypeNameX string, fkiUserlogintypeID int32, sUserlogintypeDescriptionX string, sEzsigntemplatepublicDescription string, sEzsigntemplatepublicReferenceid string, bEzsigntemplatepublicIsactive bool, tEzsigntemplatepublicNote string, eEzsigntemplatepublicLimittype FieldEEzsigntemplatepublicLimittype, iEzsigntemplatepublicLimit int32, iEzsigntemplatepublicLimitexceeded int32, dtEzsigntemplatepublicLimitexceededsince string, sEzsigntemplatepublicUrl string, sEzsigntemplatepublicEzsigntemplatedescription string, aObjEzsignfolderezsigntemplatepublic []CustomEzsignfolderezsigntemplatepublicResponse) *EzsigntemplatepublicResponseCompound {
	this := EzsigntemplatepublicResponseCompound{}
	this.PkiEzsigntemplatepublicID = pkiEzsigntemplatepublicID
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.SEzsignfoldertypeNameX = sEzsignfoldertypeNameX
	this.FkiUserlogintypeID = fkiUserlogintypeID
	this.SUserlogintypeDescriptionX = sUserlogintypeDescriptionX
	this.SEzsigntemplatepublicDescription = sEzsigntemplatepublicDescription
	this.SEzsigntemplatepublicReferenceid = sEzsigntemplatepublicReferenceid
	this.BEzsigntemplatepublicIsactive = bEzsigntemplatepublicIsactive
	this.TEzsigntemplatepublicNote = tEzsigntemplatepublicNote
	this.EEzsigntemplatepublicLimittype = eEzsigntemplatepublicLimittype
	this.IEzsigntemplatepublicLimit = iEzsigntemplatepublicLimit
	this.IEzsigntemplatepublicLimitexceeded = iEzsigntemplatepublicLimitexceeded
	this.DtEzsigntemplatepublicLimitexceededsince = dtEzsigntemplatepublicLimitexceededsince
	this.SEzsigntemplatepublicUrl = sEzsigntemplatepublicUrl
	this.SEzsigntemplatepublicEzsigntemplatedescription = sEzsigntemplatepublicEzsigntemplatedescription
	this.AObjEzsignfolderezsigntemplatepublic = aObjEzsignfolderezsigntemplatepublic
	return &this
}

// NewEzsigntemplatepublicResponseCompoundWithDefaults instantiates a new EzsigntemplatepublicResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepublicResponseCompoundWithDefaults() *EzsigntemplatepublicResponseCompound {
	this := EzsigntemplatepublicResponseCompound{}
	return &this
}

// GetPkiEzsigntemplatepublicID returns the PkiEzsigntemplatepublicID field value
func (o *EzsigntemplatepublicResponseCompound) GetPkiEzsigntemplatepublicID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplatepublicID
}

// GetPkiEzsigntemplatepublicIDOk returns a tuple with the PkiEzsigntemplatepublicID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetPkiEzsigntemplatepublicIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplatepublicID, true
}

// SetPkiEzsigntemplatepublicID sets field value
func (o *EzsigntemplatepublicResponseCompound) SetPkiEzsigntemplatepublicID(v int32) {
	o.PkiEzsigntemplatepublicID = v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value
func (o *EzsigntemplatepublicResponseCompound) GetFkiEzsignfoldertypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfoldertypeID, true
}

// SetFkiEzsignfoldertypeID sets field value
func (o *EzsigntemplatepublicResponseCompound) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = v
}

// GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field value
func (o *EzsigntemplatepublicResponseCompound) GetSEzsignfoldertypeNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfoldertypeNameX
}

// GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetSEzsignfoldertypeNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignfoldertypeNameX, true
}

// SetSEzsignfoldertypeNameX sets field value
func (o *EzsigntemplatepublicResponseCompound) SetSEzsignfoldertypeNameX(v string) {
	o.SEzsignfoldertypeNameX = v
}

// GetFkiUserlogintypeID returns the FkiUserlogintypeID field value
func (o *EzsigntemplatepublicResponseCompound) GetFkiUserlogintypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUserlogintypeID
}

// GetFkiUserlogintypeIDOk returns a tuple with the FkiUserlogintypeID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetFkiUserlogintypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiUserlogintypeID, true
}

// SetFkiUserlogintypeID sets field value
func (o *EzsigntemplatepublicResponseCompound) SetFkiUserlogintypeID(v int32) {
	o.FkiUserlogintypeID = v
}

// GetSUserlogintypeDescriptionX returns the SUserlogintypeDescriptionX field value
func (o *EzsigntemplatepublicResponseCompound) GetSUserlogintypeDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserlogintypeDescriptionX
}

// GetSUserlogintypeDescriptionXOk returns a tuple with the SUserlogintypeDescriptionX field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetSUserlogintypeDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserlogintypeDescriptionX, true
}

// SetSUserlogintypeDescriptionX sets field value
func (o *EzsigntemplatepublicResponseCompound) SetSUserlogintypeDescriptionX(v string) {
	o.SUserlogintypeDescriptionX = v
}

// GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field value if set, zero value otherwise.
func (o *EzsigntemplatepublicResponseCompound) GetFkiEzsigntemplateID() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplateID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplateID
}

// GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetFkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplateID) {
		return nil, false
	}
	return o.FkiEzsigntemplateID, true
}

// HasFkiEzsigntemplateID returns a boolean if a field has been set.
func (o *EzsigntemplatepublicResponseCompound) HasFkiEzsigntemplateID() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplateID) {
		return true
	}

	return false
}

// SetFkiEzsigntemplateID gets a reference to the given int32 and assigns it to the FkiEzsigntemplateID field.
func (o *EzsigntemplatepublicResponseCompound) SetFkiEzsigntemplateID(v int32) {
	o.FkiEzsigntemplateID = &v
}

// GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field value if set, zero value otherwise.
func (o *EzsigntemplatepublicResponseCompound) GetFkiEzsigntemplatepackageID() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplatepackageID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplatepackageID
}

// GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetFkiEzsigntemplatepackageIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplatepackageID) {
		return nil, false
	}
	return o.FkiEzsigntemplatepackageID, true
}

// HasFkiEzsigntemplatepackageID returns a boolean if a field has been set.
func (o *EzsigntemplatepublicResponseCompound) HasFkiEzsigntemplatepackageID() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplatepackageID) {
		return true
	}

	return false
}

// SetFkiEzsigntemplatepackageID gets a reference to the given int32 and assigns it to the FkiEzsigntemplatepackageID field.
func (o *EzsigntemplatepublicResponseCompound) SetFkiEzsigntemplatepackageID(v int32) {
	o.FkiEzsigntemplatepackageID = &v
}

// GetSEzsigntemplatepublicDescription returns the SEzsigntemplatepublicDescription field value
func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatepublicDescription
}

// GetSEzsigntemplatepublicDescriptionOk returns a tuple with the SEzsigntemplatepublicDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatepublicDescription, true
}

// SetSEzsigntemplatepublicDescription sets field value
func (o *EzsigntemplatepublicResponseCompound) SetSEzsigntemplatepublicDescription(v string) {
	o.SEzsigntemplatepublicDescription = v
}

// GetSEzsigntemplatepublicReferenceid returns the SEzsigntemplatepublicReferenceid field value
func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicReferenceid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatepublicReferenceid
}

// GetSEzsigntemplatepublicReferenceidOk returns a tuple with the SEzsigntemplatepublicReferenceid field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicReferenceidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatepublicReferenceid, true
}

// SetSEzsigntemplatepublicReferenceid sets field value
func (o *EzsigntemplatepublicResponseCompound) SetSEzsigntemplatepublicReferenceid(v string) {
	o.SEzsigntemplatepublicReferenceid = v
}

// GetBEzsigntemplatepublicIsactive returns the BEzsigntemplatepublicIsactive field value
func (o *EzsigntemplatepublicResponseCompound) GetBEzsigntemplatepublicIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplatepublicIsactive
}

// GetBEzsigntemplatepublicIsactiveOk returns a tuple with the BEzsigntemplatepublicIsactive field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetBEzsigntemplatepublicIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplatepublicIsactive, true
}

// SetBEzsigntemplatepublicIsactive sets field value
func (o *EzsigntemplatepublicResponseCompound) SetBEzsigntemplatepublicIsactive(v bool) {
	o.BEzsigntemplatepublicIsactive = v
}

// GetTEzsigntemplatepublicNote returns the TEzsigntemplatepublicNote field value
func (o *EzsigntemplatepublicResponseCompound) GetTEzsigntemplatepublicNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TEzsigntemplatepublicNote
}

// GetTEzsigntemplatepublicNoteOk returns a tuple with the TEzsigntemplatepublicNote field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetTEzsigntemplatepublicNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TEzsigntemplatepublicNote, true
}

// SetTEzsigntemplatepublicNote sets field value
func (o *EzsigntemplatepublicResponseCompound) SetTEzsigntemplatepublicNote(v string) {
	o.TEzsigntemplatepublicNote = v
}

// GetEEzsigntemplatepublicLimittype returns the EEzsigntemplatepublicLimittype field value
func (o *EzsigntemplatepublicResponseCompound) GetEEzsigntemplatepublicLimittype() FieldEEzsigntemplatepublicLimittype {
	if o == nil {
		var ret FieldEEzsigntemplatepublicLimittype
		return ret
	}

	return o.EEzsigntemplatepublicLimittype
}

// GetEEzsigntemplatepublicLimittypeOk returns a tuple with the EEzsigntemplatepublicLimittype field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetEEzsigntemplatepublicLimittypeOk() (*FieldEEzsigntemplatepublicLimittype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigntemplatepublicLimittype, true
}

// SetEEzsigntemplatepublicLimittype sets field value
func (o *EzsigntemplatepublicResponseCompound) SetEEzsigntemplatepublicLimittype(v FieldEEzsigntemplatepublicLimittype) {
	o.EEzsigntemplatepublicLimittype = v
}

// GetIEzsigntemplatepublicLimit returns the IEzsigntemplatepublicLimit field value
func (o *EzsigntemplatepublicResponseCompound) GetIEzsigntemplatepublicLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatepublicLimit
}

// GetIEzsigntemplatepublicLimitOk returns a tuple with the IEzsigntemplatepublicLimit field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetIEzsigntemplatepublicLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatepublicLimit, true
}

// SetIEzsigntemplatepublicLimit sets field value
func (o *EzsigntemplatepublicResponseCompound) SetIEzsigntemplatepublicLimit(v int32) {
	o.IEzsigntemplatepublicLimit = v
}

// GetIEzsigntemplatepublicLimitexceeded returns the IEzsigntemplatepublicLimitexceeded field value
func (o *EzsigntemplatepublicResponseCompound) GetIEzsigntemplatepublicLimitexceeded() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatepublicLimitexceeded
}

// GetIEzsigntemplatepublicLimitexceededOk returns a tuple with the IEzsigntemplatepublicLimitexceeded field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetIEzsigntemplatepublicLimitexceededOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatepublicLimitexceeded, true
}

// SetIEzsigntemplatepublicLimitexceeded sets field value
func (o *EzsigntemplatepublicResponseCompound) SetIEzsigntemplatepublicLimitexceeded(v int32) {
	o.IEzsigntemplatepublicLimitexceeded = v
}

// GetDtEzsigntemplatepublicLimitexceededsince returns the DtEzsigntemplatepublicLimitexceededsince field value
func (o *EzsigntemplatepublicResponseCompound) GetDtEzsigntemplatepublicLimitexceededsince() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtEzsigntemplatepublicLimitexceededsince
}

// GetDtEzsigntemplatepublicLimitexceededsinceOk returns a tuple with the DtEzsigntemplatepublicLimitexceededsince field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetDtEzsigntemplatepublicLimitexceededsinceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtEzsigntemplatepublicLimitexceededsince, true
}

// SetDtEzsigntemplatepublicLimitexceededsince sets field value
func (o *EzsigntemplatepublicResponseCompound) SetDtEzsigntemplatepublicLimitexceededsince(v string) {
	o.DtEzsigntemplatepublicLimitexceededsince = v
}

// GetSEzsigntemplatepublicUrl returns the SEzsigntemplatepublicUrl field value
func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatepublicUrl
}

// GetSEzsigntemplatepublicUrlOk returns a tuple with the SEzsigntemplatepublicUrl field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatepublicUrl, true
}

// SetSEzsigntemplatepublicUrl sets field value
func (o *EzsigntemplatepublicResponseCompound) SetSEzsigntemplatepublicUrl(v string) {
	o.SEzsigntemplatepublicUrl = v
}

// GetSEzsigntemplatepublicEzsigntemplatedescription returns the SEzsigntemplatepublicEzsigntemplatedescription field value
func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicEzsigntemplatedescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatepublicEzsigntemplatedescription
}

// GetSEzsigntemplatepublicEzsigntemplatedescriptionOk returns a tuple with the SEzsigntemplatepublicEzsigntemplatedescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetSEzsigntemplatepublicEzsigntemplatedescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatepublicEzsigntemplatedescription, true
}

// SetSEzsigntemplatepublicEzsigntemplatedescription sets field value
func (o *EzsigntemplatepublicResponseCompound) SetSEzsigntemplatepublicEzsigntemplatedescription(v string) {
	o.SEzsigntemplatepublicEzsigntemplatedescription = v
}

// GetObjAudit returns the ObjAudit field value if set, zero value otherwise.
func (o *EzsigntemplatepublicResponseCompound) GetObjAudit() CommonAudit {
	if o == nil || IsNil(o.ObjAudit) {
		var ret CommonAudit
		return ret
	}
	return *o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil || IsNil(o.ObjAudit) {
		return nil, false
	}
	return o.ObjAudit, true
}

// HasObjAudit returns a boolean if a field has been set.
func (o *EzsigntemplatepublicResponseCompound) HasObjAudit() bool {
	if o != nil && !IsNil(o.ObjAudit) {
		return true
	}

	return false
}

// SetObjAudit gets a reference to the given CommonAudit and assigns it to the ObjAudit field.
func (o *EzsigntemplatepublicResponseCompound) SetObjAudit(v CommonAudit) {
	o.ObjAudit = &v
}

// GetAObjEzsignfolderezsigntemplatepublic returns the AObjEzsignfolderezsigntemplatepublic field value
func (o *EzsigntemplatepublicResponseCompound) GetAObjEzsignfolderezsigntemplatepublic() []CustomEzsignfolderezsigntemplatepublicResponse {
	if o == nil {
		var ret []CustomEzsignfolderezsigntemplatepublicResponse
		return ret
	}

	return o.AObjEzsignfolderezsigntemplatepublic
}

// GetAObjEzsignfolderezsigntemplatepublicOk returns a tuple with the AObjEzsignfolderezsigntemplatepublic field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepublicResponseCompound) GetAObjEzsignfolderezsigntemplatepublicOk() ([]CustomEzsignfolderezsigntemplatepublicResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignfolderezsigntemplatepublic, true
}

// SetAObjEzsignfolderezsigntemplatepublic sets field value
func (o *EzsigntemplatepublicResponseCompound) SetAObjEzsignfolderezsigntemplatepublic(v []CustomEzsignfolderezsigntemplatepublicResponse) {
	o.AObjEzsignfolderezsigntemplatepublic = v
}

func (o EzsigntemplatepublicResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepublicResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplatepublicID"] = o.PkiEzsigntemplatepublicID
	toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	toSerialize["sEzsignfoldertypeNameX"] = o.SEzsignfoldertypeNameX
	toSerialize["fkiUserlogintypeID"] = o.FkiUserlogintypeID
	toSerialize["sUserlogintypeDescriptionX"] = o.SUserlogintypeDescriptionX
	if !IsNil(o.FkiEzsigntemplateID) {
		toSerialize["fkiEzsigntemplateID"] = o.FkiEzsigntemplateID
	}
	if !IsNil(o.FkiEzsigntemplatepackageID) {
		toSerialize["fkiEzsigntemplatepackageID"] = o.FkiEzsigntemplatepackageID
	}
	toSerialize["sEzsigntemplatepublicDescription"] = o.SEzsigntemplatepublicDescription
	toSerialize["sEzsigntemplatepublicReferenceid"] = o.SEzsigntemplatepublicReferenceid
	toSerialize["bEzsigntemplatepublicIsactive"] = o.BEzsigntemplatepublicIsactive
	toSerialize["tEzsigntemplatepublicNote"] = o.TEzsigntemplatepublicNote
	toSerialize["eEzsigntemplatepublicLimittype"] = o.EEzsigntemplatepublicLimittype
	toSerialize["iEzsigntemplatepublicLimit"] = o.IEzsigntemplatepublicLimit
	toSerialize["iEzsigntemplatepublicLimitexceeded"] = o.IEzsigntemplatepublicLimitexceeded
	toSerialize["dtEzsigntemplatepublicLimitexceededsince"] = o.DtEzsigntemplatepublicLimitexceededsince
	toSerialize["sEzsigntemplatepublicUrl"] = o.SEzsigntemplatepublicUrl
	toSerialize["sEzsigntemplatepublicEzsigntemplatedescription"] = o.SEzsigntemplatepublicEzsigntemplatedescription
	if !IsNil(o.ObjAudit) {
		toSerialize["objAudit"] = o.ObjAudit
	}
	toSerialize["a_objEzsignfolderezsigntemplatepublic"] = o.AObjEzsignfolderezsigntemplatepublic
	return toSerialize, nil
}

func (o *EzsigntemplatepublicResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplatepublicID",
		"fkiEzsignfoldertypeID",
		"sEzsignfoldertypeNameX",
		"fkiUserlogintypeID",
		"sUserlogintypeDescriptionX",
		"sEzsigntemplatepublicDescription",
		"sEzsigntemplatepublicReferenceid",
		"bEzsigntemplatepublicIsactive",
		"tEzsigntemplatepublicNote",
		"eEzsigntemplatepublicLimittype",
		"iEzsigntemplatepublicLimit",
		"iEzsigntemplatepublicLimitexceeded",
		"dtEzsigntemplatepublicLimitexceededsince",
		"sEzsigntemplatepublicUrl",
		"sEzsigntemplatepublicEzsigntemplatedescription",
		"a_objEzsignfolderezsigntemplatepublic",
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

	varEzsigntemplatepublicResponseCompound := _EzsigntemplatepublicResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepublicResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepublicResponseCompound(varEzsigntemplatepublicResponseCompound)

	return err
}

type NullableEzsigntemplatepublicResponseCompound struct {
	value *EzsigntemplatepublicResponseCompound
	isSet bool
}

func (v NullableEzsigntemplatepublicResponseCompound) Get() *EzsigntemplatepublicResponseCompound {
	return v.value
}

func (v *NullableEzsigntemplatepublicResponseCompound) Set(val *EzsigntemplatepublicResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepublicResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepublicResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepublicResponseCompound(val *EzsigntemplatepublicResponseCompound) *NullableEzsigntemplatepublicResponseCompound {
	return &NullableEzsigntemplatepublicResponseCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplatepublicResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepublicResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


