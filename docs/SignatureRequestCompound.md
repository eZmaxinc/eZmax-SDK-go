# SignatureRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSignatureID** | Pointer to **int32** | The unique ID of the Signature | [optional] 
**TSignatureSvg** | **string** | The svg of the Signature | 

## Methods

### NewSignatureRequestCompound

`func NewSignatureRequestCompound(tSignatureSvg string, ) *SignatureRequestCompound`

NewSignatureRequestCompound instantiates a new SignatureRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureRequestCompoundWithDefaults

`func NewSignatureRequestCompoundWithDefaults() *SignatureRequestCompound`

NewSignatureRequestCompoundWithDefaults instantiates a new SignatureRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSignatureID

`func (o *SignatureRequestCompound) GetPkiSignatureID() int32`

GetPkiSignatureID returns the PkiSignatureID field if non-nil, zero value otherwise.

### GetPkiSignatureIDOk

`func (o *SignatureRequestCompound) GetPkiSignatureIDOk() (*int32, bool)`

GetPkiSignatureIDOk returns a tuple with the PkiSignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSignatureID

`func (o *SignatureRequestCompound) SetPkiSignatureID(v int32)`

SetPkiSignatureID sets PkiSignatureID field to given value.

### HasPkiSignatureID

`func (o *SignatureRequestCompound) HasPkiSignatureID() bool`

HasPkiSignatureID returns a boolean if a field has been set.

### GetTSignatureSvg

`func (o *SignatureRequestCompound) GetTSignatureSvg() string`

GetTSignatureSvg returns the TSignatureSvg field if non-nil, zero value otherwise.

### GetTSignatureSvgOk

`func (o *SignatureRequestCompound) GetTSignatureSvgOk() (*string, bool)`

GetTSignatureSvgOk returns a tuple with the TSignatureSvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTSignatureSvg

`func (o *SignatureRequestCompound) SetTSignatureSvg(v string)`

SetTSignatureSvg sets TSignatureSvg field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


