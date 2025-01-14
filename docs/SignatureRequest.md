# SignatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSignatureID** | Pointer to **int32** | The unique ID of the Signature | [optional] 
**FkiFontID** | **int32** | The unique ID of the Font | 
**ESignaturePreference** | [**FieldESignaturePreference**](FieldESignaturePreference.md) |  | 
**TSignatureSvg** | Pointer to **string** | The svg of the Signature | [optional] 
**TSignatureSvginitials** | Pointer to **string** | The svg of the Initials | [optional] 

## Methods

### NewSignatureRequest

`func NewSignatureRequest(fkiFontID int32, eSignaturePreference FieldESignaturePreference, ) *SignatureRequest`

NewSignatureRequest instantiates a new SignatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureRequestWithDefaults

`func NewSignatureRequestWithDefaults() *SignatureRequest`

NewSignatureRequestWithDefaults instantiates a new SignatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSignatureID

`func (o *SignatureRequest) GetPkiSignatureID() int32`

GetPkiSignatureID returns the PkiSignatureID field if non-nil, zero value otherwise.

### GetPkiSignatureIDOk

`func (o *SignatureRequest) GetPkiSignatureIDOk() (*int32, bool)`

GetPkiSignatureIDOk returns a tuple with the PkiSignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSignatureID

`func (o *SignatureRequest) SetPkiSignatureID(v int32)`

SetPkiSignatureID sets PkiSignatureID field to given value.

### HasPkiSignatureID

`func (o *SignatureRequest) HasPkiSignatureID() bool`

HasPkiSignatureID returns a boolean if a field has been set.

### GetFkiFontID

`func (o *SignatureRequest) GetFkiFontID() int32`

GetFkiFontID returns the FkiFontID field if non-nil, zero value otherwise.

### GetFkiFontIDOk

`func (o *SignatureRequest) GetFkiFontIDOk() (*int32, bool)`

GetFkiFontIDOk returns a tuple with the FkiFontID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontID

`func (o *SignatureRequest) SetFkiFontID(v int32)`

SetFkiFontID sets FkiFontID field to given value.


### GetESignaturePreference

`func (o *SignatureRequest) GetESignaturePreference() FieldESignaturePreference`

GetESignaturePreference returns the ESignaturePreference field if non-nil, zero value otherwise.

### GetESignaturePreferenceOk

`func (o *SignatureRequest) GetESignaturePreferenceOk() (*FieldESignaturePreference, bool)`

GetESignaturePreferenceOk returns a tuple with the ESignaturePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESignaturePreference

`func (o *SignatureRequest) SetESignaturePreference(v FieldESignaturePreference)`

SetESignaturePreference sets ESignaturePreference field to given value.


### GetTSignatureSvg

`func (o *SignatureRequest) GetTSignatureSvg() string`

GetTSignatureSvg returns the TSignatureSvg field if non-nil, zero value otherwise.

### GetTSignatureSvgOk

`func (o *SignatureRequest) GetTSignatureSvgOk() (*string, bool)`

GetTSignatureSvgOk returns a tuple with the TSignatureSvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTSignatureSvg

`func (o *SignatureRequest) SetTSignatureSvg(v string)`

SetTSignatureSvg sets TSignatureSvg field to given value.

### HasTSignatureSvg

`func (o *SignatureRequest) HasTSignatureSvg() bool`

HasTSignatureSvg returns a boolean if a field has been set.

### GetTSignatureSvginitials

`func (o *SignatureRequest) GetTSignatureSvginitials() string`

GetTSignatureSvginitials returns the TSignatureSvginitials field if non-nil, zero value otherwise.

### GetTSignatureSvginitialsOk

`func (o *SignatureRequest) GetTSignatureSvginitialsOk() (*string, bool)`

GetTSignatureSvginitialsOk returns a tuple with the TSignatureSvginitials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTSignatureSvginitials

`func (o *SignatureRequest) SetTSignatureSvginitials(v string)`

SetTSignatureSvginitials sets TSignatureSvginitials field to given value.

### HasTSignatureSvginitials

`func (o *SignatureRequest) HasTSignatureSvginitials() bool`

HasTSignatureSvginitials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


