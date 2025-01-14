# EzdoctemplatedocumentRequestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EEzdoctemplatedocumentFormat** | Pointer to **string** | Indicates the format of the template.  This field is Required when sEzdoctemplatedocumentBase64 is set. | [optional] 
**SEzdoctemplatedocumentFields** | Pointer to **string** | List of field in Ezdoctemplatedocument | [optional] 
**SEzdoctemplatedocumentBase64** | Pointer to **string** | The Base64 encoded binary content of the document.  This field is Required when eEzdoctemplatedocumentFormat is set. | [optional] 

## Methods

### NewEzdoctemplatedocumentRequestPatch

`func NewEzdoctemplatedocumentRequestPatch() *EzdoctemplatedocumentRequestPatch`

NewEzdoctemplatedocumentRequestPatch instantiates a new EzdoctemplatedocumentRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzdoctemplatedocumentRequestPatchWithDefaults

`func NewEzdoctemplatedocumentRequestPatchWithDefaults() *EzdoctemplatedocumentRequestPatch`

NewEzdoctemplatedocumentRequestPatchWithDefaults instantiates a new EzdoctemplatedocumentRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEEzdoctemplatedocumentFormat

`func (o *EzdoctemplatedocumentRequestPatch) GetEEzdoctemplatedocumentFormat() string`

GetEEzdoctemplatedocumentFormat returns the EEzdoctemplatedocumentFormat field if non-nil, zero value otherwise.

### GetEEzdoctemplatedocumentFormatOk

`func (o *EzdoctemplatedocumentRequestPatch) GetEEzdoctemplatedocumentFormatOk() (*string, bool)`

GetEEzdoctemplatedocumentFormatOk returns a tuple with the EEzdoctemplatedocumentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzdoctemplatedocumentFormat

`func (o *EzdoctemplatedocumentRequestPatch) SetEEzdoctemplatedocumentFormat(v string)`

SetEEzdoctemplatedocumentFormat sets EEzdoctemplatedocumentFormat field to given value.

### HasEEzdoctemplatedocumentFormat

`func (o *EzdoctemplatedocumentRequestPatch) HasEEzdoctemplatedocumentFormat() bool`

HasEEzdoctemplatedocumentFormat returns a boolean if a field has been set.

### GetSEzdoctemplatedocumentFields

`func (o *EzdoctemplatedocumentRequestPatch) GetSEzdoctemplatedocumentFields() string`

GetSEzdoctemplatedocumentFields returns the SEzdoctemplatedocumentFields field if non-nil, zero value otherwise.

### GetSEzdoctemplatedocumentFieldsOk

`func (o *EzdoctemplatedocumentRequestPatch) GetSEzdoctemplatedocumentFieldsOk() (*string, bool)`

GetSEzdoctemplatedocumentFieldsOk returns a tuple with the SEzdoctemplatedocumentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzdoctemplatedocumentFields

`func (o *EzdoctemplatedocumentRequestPatch) SetSEzdoctemplatedocumentFields(v string)`

SetSEzdoctemplatedocumentFields sets SEzdoctemplatedocumentFields field to given value.

### HasSEzdoctemplatedocumentFields

`func (o *EzdoctemplatedocumentRequestPatch) HasSEzdoctemplatedocumentFields() bool`

HasSEzdoctemplatedocumentFields returns a boolean if a field has been set.

### GetSEzdoctemplatedocumentBase64

`func (o *EzdoctemplatedocumentRequestPatch) GetSEzdoctemplatedocumentBase64() string`

GetSEzdoctemplatedocumentBase64 returns the SEzdoctemplatedocumentBase64 field if non-nil, zero value otherwise.

### GetSEzdoctemplatedocumentBase64Ok

`func (o *EzdoctemplatedocumentRequestPatch) GetSEzdoctemplatedocumentBase64Ok() (*string, bool)`

GetSEzdoctemplatedocumentBase64Ok returns a tuple with the SEzdoctemplatedocumentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzdoctemplatedocumentBase64

`func (o *EzdoctemplatedocumentRequestPatch) SetSEzdoctemplatedocumentBase64(v string)`

SetSEzdoctemplatedocumentBase64 sets SEzdoctemplatedocumentBase64 field to given value.

### HasSEzdoctemplatedocumentBase64

`func (o *EzdoctemplatedocumentRequestPatch) HasSEzdoctemplatedocumentBase64() bool`

HasSEzdoctemplatedocumentBase64 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


