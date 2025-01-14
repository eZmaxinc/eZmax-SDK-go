# SignatureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSignatureID** | **int32** | The unique ID of the Signature | 
**FkiFontID** | Pointer to **int32** | The unique ID of the Font | [optional] 
**SSignatureUrl** | Pointer to **string** | The URL of the SVG file for the Signature | [optional] 
**SSignatureUrlinitials** | Pointer to **string** | The URL of the SVG file for the Initials | [optional] 

## Methods

### NewSignatureResponse

`func NewSignatureResponse(pkiSignatureID int32, ) *SignatureResponse`

NewSignatureResponse instantiates a new SignatureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureResponseWithDefaults

`func NewSignatureResponseWithDefaults() *SignatureResponse`

NewSignatureResponseWithDefaults instantiates a new SignatureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSignatureID

`func (o *SignatureResponse) GetPkiSignatureID() int32`

GetPkiSignatureID returns the PkiSignatureID field if non-nil, zero value otherwise.

### GetPkiSignatureIDOk

`func (o *SignatureResponse) GetPkiSignatureIDOk() (*int32, bool)`

GetPkiSignatureIDOk returns a tuple with the PkiSignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSignatureID

`func (o *SignatureResponse) SetPkiSignatureID(v int32)`

SetPkiSignatureID sets PkiSignatureID field to given value.


### GetFkiFontID

`func (o *SignatureResponse) GetFkiFontID() int32`

GetFkiFontID returns the FkiFontID field if non-nil, zero value otherwise.

### GetFkiFontIDOk

`func (o *SignatureResponse) GetFkiFontIDOk() (*int32, bool)`

GetFkiFontIDOk returns a tuple with the FkiFontID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontID

`func (o *SignatureResponse) SetFkiFontID(v int32)`

SetFkiFontID sets FkiFontID field to given value.

### HasFkiFontID

`func (o *SignatureResponse) HasFkiFontID() bool`

HasFkiFontID returns a boolean if a field has been set.

### GetSSignatureUrl

`func (o *SignatureResponse) GetSSignatureUrl() string`

GetSSignatureUrl returns the SSignatureUrl field if non-nil, zero value otherwise.

### GetSSignatureUrlOk

`func (o *SignatureResponse) GetSSignatureUrlOk() (*string, bool)`

GetSSignatureUrlOk returns a tuple with the SSignatureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSignatureUrl

`func (o *SignatureResponse) SetSSignatureUrl(v string)`

SetSSignatureUrl sets SSignatureUrl field to given value.

### HasSSignatureUrl

`func (o *SignatureResponse) HasSSignatureUrl() bool`

HasSSignatureUrl returns a boolean if a field has been set.

### GetSSignatureUrlinitials

`func (o *SignatureResponse) GetSSignatureUrlinitials() string`

GetSSignatureUrlinitials returns the SSignatureUrlinitials field if non-nil, zero value otherwise.

### GetSSignatureUrlinitialsOk

`func (o *SignatureResponse) GetSSignatureUrlinitialsOk() (*string, bool)`

GetSSignatureUrlinitialsOk returns a tuple with the SSignatureUrlinitials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSignatureUrlinitials

`func (o *SignatureResponse) SetSSignatureUrlinitials(v string)`

SetSSignatureUrlinitials sets SSignatureUrlinitials field to given value.

### HasSSignatureUrlinitials

`func (o *SignatureResponse) HasSSignatureUrlinitials() bool`

HasSSignatureUrlinitials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


