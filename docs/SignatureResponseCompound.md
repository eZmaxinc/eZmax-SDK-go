# SignatureResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSignatureID** | **int32** | The unique ID of the Signature | 
**FkiFontID** | Pointer to **int32** | The unique ID of the Font | [optional] 
**SSignatureUrl** | Pointer to **string** | The URL of the SVG file for the Signature | [optional] 
**SSignatureUrlinitials** | Pointer to **string** | The URL of the SVG file for the Initials | [optional] 

## Methods

### NewSignatureResponseCompound

`func NewSignatureResponseCompound(pkiSignatureID int32, ) *SignatureResponseCompound`

NewSignatureResponseCompound instantiates a new SignatureResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureResponseCompoundWithDefaults

`func NewSignatureResponseCompoundWithDefaults() *SignatureResponseCompound`

NewSignatureResponseCompoundWithDefaults instantiates a new SignatureResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSignatureID

`func (o *SignatureResponseCompound) GetPkiSignatureID() int32`

GetPkiSignatureID returns the PkiSignatureID field if non-nil, zero value otherwise.

### GetPkiSignatureIDOk

`func (o *SignatureResponseCompound) GetPkiSignatureIDOk() (*int32, bool)`

GetPkiSignatureIDOk returns a tuple with the PkiSignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSignatureID

`func (o *SignatureResponseCompound) SetPkiSignatureID(v int32)`

SetPkiSignatureID sets PkiSignatureID field to given value.


### GetFkiFontID

`func (o *SignatureResponseCompound) GetFkiFontID() int32`

GetFkiFontID returns the FkiFontID field if non-nil, zero value otherwise.

### GetFkiFontIDOk

`func (o *SignatureResponseCompound) GetFkiFontIDOk() (*int32, bool)`

GetFkiFontIDOk returns a tuple with the FkiFontID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontID

`func (o *SignatureResponseCompound) SetFkiFontID(v int32)`

SetFkiFontID sets FkiFontID field to given value.

### HasFkiFontID

`func (o *SignatureResponseCompound) HasFkiFontID() bool`

HasFkiFontID returns a boolean if a field has been set.

### GetSSignatureUrl

`func (o *SignatureResponseCompound) GetSSignatureUrl() string`

GetSSignatureUrl returns the SSignatureUrl field if non-nil, zero value otherwise.

### GetSSignatureUrlOk

`func (o *SignatureResponseCompound) GetSSignatureUrlOk() (*string, bool)`

GetSSignatureUrlOk returns a tuple with the SSignatureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSignatureUrl

`func (o *SignatureResponseCompound) SetSSignatureUrl(v string)`

SetSSignatureUrl sets SSignatureUrl field to given value.

### HasSSignatureUrl

`func (o *SignatureResponseCompound) HasSSignatureUrl() bool`

HasSSignatureUrl returns a boolean if a field has been set.

### GetSSignatureUrlinitials

`func (o *SignatureResponseCompound) GetSSignatureUrlinitials() string`

GetSSignatureUrlinitials returns the SSignatureUrlinitials field if non-nil, zero value otherwise.

### GetSSignatureUrlinitialsOk

`func (o *SignatureResponseCompound) GetSSignatureUrlinitialsOk() (*string, bool)`

GetSSignatureUrlinitialsOk returns a tuple with the SSignatureUrlinitials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSignatureUrlinitials

`func (o *SignatureResponseCompound) SetSSignatureUrlinitials(v string)`

SetSSignatureUrlinitials sets SSignatureUrlinitials field to given value.

### HasSSignatureUrlinitials

`func (o *SignatureResponseCompound) HasSSignatureUrlinitials() bool`

HasSSignatureUrlinitials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


