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
)

// checks if the EzsignfoldersignerassociationCreateEmbeddedUrlV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfoldersignerassociationCreateEmbeddedUrlV1Request{}

// EzsignfoldersignerassociationCreateEmbeddedUrlV1Request Request for POST /1/object/ezsignfoldersignerassociation/createEmbeddedUrl
type EzsignfoldersignerassociationCreateEmbeddedUrlV1Request struct {
	// The return Url to redirect after the signing is completed  **Warning** Due to the potential for Ezsignsigners to cancel redirection, close their browser post-signing, or spoof the landing URL, it's advisable not to solely depend on the sReturnUrl for accurate status within your integration.  Once the Ezsignsigner finishes, they are directed back to your application. Your application can retain transaction state details by either storing data in a cookie or incorporating query parameters in the sReturnUrl. For example: https://www.example.com/sReturnUrl?sSessionID=ABC123  The actual url that will be called will have an extra url parameter appended to give details about the process. The possible values are listed in the table below. For example: https://www.example.com/sReturnUrl?sSessionID=ABC123&eEzsignEvent=CompletedEzsignfolder   |**Query parameters appended**| |---| |eEzsignEvent|   |**eEzsignEvent**|**Description**| |---|---| |SessionTimeout|The session timed out| |SessionLogout|The Ezsignsigner signed out| |DeclinedTermOfUse|The Ezsignsigner refused the terms| |DeclinedSign|The Ezsignsigner refused to sign| |Reassigned|The Ezsignsigner reassigned his signatures to someone else| |CompletedStep|The Ezsignsigner completed his step. There is other signatures to complete the Ezsigndocument| |CompletedEzsignfolder|The Ezsignfolder is completed. Everyone signed their signatures|
	SReturnUrl *string `json:"sReturnUrl,omitempty" validate:"regexp=^(https|http):\\/\\/[^\\\\s\\/$.?#].[^\\\\s]*$"`
	// Domain protection for the iFrame
	SIframedomain *string `json:"sIframedomain,omitempty"`
	// Whether the url would be in an iFrame or not
	BIsIframe *bool `json:"bIsIframe,omitempty"`
}

// NewEzsignfoldersignerassociationCreateEmbeddedUrlV1Request instantiates a new EzsignfoldersignerassociationCreateEmbeddedUrlV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationCreateEmbeddedUrlV1Request() *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request {
	this := EzsignfoldersignerassociationCreateEmbeddedUrlV1Request{}
	return &this
}

// NewEzsignfoldersignerassociationCreateEmbeddedUrlV1RequestWithDefaults instantiates a new EzsignfoldersignerassociationCreateEmbeddedUrlV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationCreateEmbeddedUrlV1RequestWithDefaults() *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request {
	this := EzsignfoldersignerassociationCreateEmbeddedUrlV1Request{}
	return &this
}

// GetSReturnUrl returns the SReturnUrl field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) GetSReturnUrl() string {
	if o == nil || IsNil(o.SReturnUrl) {
		var ret string
		return ret
	}
	return *o.SReturnUrl
}

// GetSReturnUrlOk returns a tuple with the SReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) GetSReturnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SReturnUrl) {
		return nil, false
	}
	return o.SReturnUrl, true
}

// HasSReturnUrl returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) HasSReturnUrl() bool {
	if o != nil && !IsNil(o.SReturnUrl) {
		return true
	}

	return false
}

// SetSReturnUrl gets a reference to the given string and assigns it to the SReturnUrl field.
func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) SetSReturnUrl(v string) {
	o.SReturnUrl = &v
}

// GetSIframedomain returns the SIframedomain field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) GetSIframedomain() string {
	if o == nil || IsNil(o.SIframedomain) {
		var ret string
		return ret
	}
	return *o.SIframedomain
}

// GetSIframedomainOk returns a tuple with the SIframedomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) GetSIframedomainOk() (*string, bool) {
	if o == nil || IsNil(o.SIframedomain) {
		return nil, false
	}
	return o.SIframedomain, true
}

// HasSIframedomain returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) HasSIframedomain() bool {
	if o != nil && !IsNil(o.SIframedomain) {
		return true
	}

	return false
}

// SetSIframedomain gets a reference to the given string and assigns it to the SIframedomain field.
func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) SetSIframedomain(v string) {
	o.SIframedomain = &v
}

// GetBIsIframe returns the BIsIframe field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) GetBIsIframe() bool {
	if o == nil || IsNil(o.BIsIframe) {
		var ret bool
		return ret
	}
	return *o.BIsIframe
}

// GetBIsIframeOk returns a tuple with the BIsIframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) GetBIsIframeOk() (*bool, bool) {
	if o == nil || IsNil(o.BIsIframe) {
		return nil, false
	}
	return o.BIsIframe, true
}

// HasBIsIframe returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) HasBIsIframe() bool {
	if o != nil && !IsNil(o.BIsIframe) {
		return true
	}

	return false
}

// SetBIsIframe gets a reference to the given bool and assigns it to the BIsIframe field.
func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) SetBIsIframe(v bool) {
	o.BIsIframe = &v
}

func (o EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SReturnUrl) {
		toSerialize["sReturnUrl"] = o.SReturnUrl
	}
	if !IsNil(o.SIframedomain) {
		toSerialize["sIframedomain"] = o.SIframedomain
	}
	if !IsNil(o.BIsIframe) {
		toSerialize["bIsIframe"] = o.BIsIframe
	}
	return toSerialize, nil
}

type NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1Request struct {
	value *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request
	isSet bool
}

func (v NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1Request) Get() *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1Request) Set(val *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationCreateEmbeddedUrlV1Request(val *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) *NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1Request {
	return &NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1Request{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


