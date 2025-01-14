# SignatureResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSignatureID** | **int32** | The unique ID of the Signature | 
**FkiFontID** | **int32** | The unique ID of the Font | 
**ESignaturePreference** | [**FieldESignaturePreference**](FieldESignaturePreference.md) |  | 
**BSignatureSvg** | **bool** | Whether the signature has a SVG or not | 
**BSignatureSvginitials** | **bool** | Whether the initials has a SVG or not | 

## Methods

### NewSignatureResponseV3

`func NewSignatureResponseV3(pkiSignatureID int32, fkiFontID int32, eSignaturePreference FieldESignaturePreference, bSignatureSvg bool, bSignatureSvginitials bool, ) *SignatureResponseV3`

NewSignatureResponseV3 instantiates a new SignatureResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureResponseV3WithDefaults

`func NewSignatureResponseV3WithDefaults() *SignatureResponseV3`

NewSignatureResponseV3WithDefaults instantiates a new SignatureResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSignatureID

`func (o *SignatureResponseV3) GetPkiSignatureID() int32`

GetPkiSignatureID returns the PkiSignatureID field if non-nil, zero value otherwise.

### GetPkiSignatureIDOk

`func (o *SignatureResponseV3) GetPkiSignatureIDOk() (*int32, bool)`

GetPkiSignatureIDOk returns a tuple with the PkiSignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSignatureID

`func (o *SignatureResponseV3) SetPkiSignatureID(v int32)`

SetPkiSignatureID sets PkiSignatureID field to given value.


### GetFkiFontID

`func (o *SignatureResponseV3) GetFkiFontID() int32`

GetFkiFontID returns the FkiFontID field if non-nil, zero value otherwise.

### GetFkiFontIDOk

`func (o *SignatureResponseV3) GetFkiFontIDOk() (*int32, bool)`

GetFkiFontIDOk returns a tuple with the FkiFontID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontID

`func (o *SignatureResponseV3) SetFkiFontID(v int32)`

SetFkiFontID sets FkiFontID field to given value.


### GetESignaturePreference

`func (o *SignatureResponseV3) GetESignaturePreference() FieldESignaturePreference`

GetESignaturePreference returns the ESignaturePreference field if non-nil, zero value otherwise.

### GetESignaturePreferenceOk

`func (o *SignatureResponseV3) GetESignaturePreferenceOk() (*FieldESignaturePreference, bool)`

GetESignaturePreferenceOk returns a tuple with the ESignaturePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESignaturePreference

`func (o *SignatureResponseV3) SetESignaturePreference(v FieldESignaturePreference)`

SetESignaturePreference sets ESignaturePreference field to given value.


### GetBSignatureSvg

`func (o *SignatureResponseV3) GetBSignatureSvg() bool`

GetBSignatureSvg returns the BSignatureSvg field if non-nil, zero value otherwise.

### GetBSignatureSvgOk

`func (o *SignatureResponseV3) GetBSignatureSvgOk() (*bool, bool)`

GetBSignatureSvgOk returns a tuple with the BSignatureSvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSignatureSvg

`func (o *SignatureResponseV3) SetBSignatureSvg(v bool)`

SetBSignatureSvg sets BSignatureSvg field to given value.


### GetBSignatureSvginitials

`func (o *SignatureResponseV3) GetBSignatureSvginitials() bool`

GetBSignatureSvginitials returns the BSignatureSvginitials field if non-nil, zero value otherwise.

### GetBSignatureSvginitialsOk

`func (o *SignatureResponseV3) GetBSignatureSvginitialsOk() (*bool, bool)`

GetBSignatureSvginitialsOk returns a tuple with the BSignatureSvginitials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSignatureSvginitials

`func (o *SignatureResponseV3) SetBSignatureSvginitials(v bool)`

SetBSignatureSvginitials sets BSignatureSvginitials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


