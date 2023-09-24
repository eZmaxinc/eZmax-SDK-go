# SignatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSignatureID** | Pointer to **int32** | The unique ID of the Signature | [optional] 
**TSignatureSvg** | **string** | The svg of the Signature | 

## Methods

### NewSignatureRequest

`func NewSignatureRequest(tSignatureSvg string, ) *SignatureRequest`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


