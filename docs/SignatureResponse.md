# SignatureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSignatureID** | **int32** | The unique ID of the Signature | 
**SSignatureUrl** | **string** | The URL of the SVG file for the Signature | 

## Methods

### NewSignatureResponse

`func NewSignatureResponse(pkiSignatureID int32, sSignatureUrl string, ) *SignatureResponse`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


