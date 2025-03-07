# CustomPrefillEzsignformValueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SEzsignformfieldgroupLabel** | **string** | The Label for the Ezsignformfieldgroup | 
**SEzsignformfieldLabel** | Pointer to **string** | The Label for the Ezsignformfield | [optional] 
**SEzsignformfieldEnteredvalue** | Pointer to **string** | This is the value enterred for the Ezsignformfield  This can only be set if eEzsignformfieldgroupType is **Dropdown**, **Text** or **Textarea** | [optional] 
**BEzsignformfieldSelected** | Pointer to **bool** | Whether the Ezsignformfield is selected or not by default.  This can only be set if eEzsignformfieldgroupType is **Checkbox** or **Radio** | [optional] 

## Methods

### NewCustomPrefillEzsignformValueRequest

`func NewCustomPrefillEzsignformValueRequest(sEzsignformfieldgroupLabel string, ) *CustomPrefillEzsignformValueRequest`

NewCustomPrefillEzsignformValueRequest instantiates a new CustomPrefillEzsignformValueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPrefillEzsignformValueRequestWithDefaults

`func NewCustomPrefillEzsignformValueRequestWithDefaults() *CustomPrefillEzsignformValueRequest`

NewCustomPrefillEzsignformValueRequestWithDefaults instantiates a new CustomPrefillEzsignformValueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSEzsignformfieldgroupLabel

`func (o *CustomPrefillEzsignformValueRequest) GetSEzsignformfieldgroupLabel() string`

GetSEzsignformfieldgroupLabel returns the SEzsignformfieldgroupLabel field if non-nil, zero value otherwise.

### GetSEzsignformfieldgroupLabelOk

`func (o *CustomPrefillEzsignformValueRequest) GetSEzsignformfieldgroupLabelOk() (*string, bool)`

GetSEzsignformfieldgroupLabelOk returns a tuple with the SEzsignformfieldgroupLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldgroupLabel

`func (o *CustomPrefillEzsignformValueRequest) SetSEzsignformfieldgroupLabel(v string)`

SetSEzsignformfieldgroupLabel sets SEzsignformfieldgroupLabel field to given value.


### GetSEzsignformfieldLabel

`func (o *CustomPrefillEzsignformValueRequest) GetSEzsignformfieldLabel() string`

GetSEzsignformfieldLabel returns the SEzsignformfieldLabel field if non-nil, zero value otherwise.

### GetSEzsignformfieldLabelOk

`func (o *CustomPrefillEzsignformValueRequest) GetSEzsignformfieldLabelOk() (*string, bool)`

GetSEzsignformfieldLabelOk returns a tuple with the SEzsignformfieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldLabel

`func (o *CustomPrefillEzsignformValueRequest) SetSEzsignformfieldLabel(v string)`

SetSEzsignformfieldLabel sets SEzsignformfieldLabel field to given value.

### HasSEzsignformfieldLabel

`func (o *CustomPrefillEzsignformValueRequest) HasSEzsignformfieldLabel() bool`

HasSEzsignformfieldLabel returns a boolean if a field has been set.

### GetSEzsignformfieldEnteredvalue

`func (o *CustomPrefillEzsignformValueRequest) GetSEzsignformfieldEnteredvalue() string`

GetSEzsignformfieldEnteredvalue returns the SEzsignformfieldEnteredvalue field if non-nil, zero value otherwise.

### GetSEzsignformfieldEnteredvalueOk

`func (o *CustomPrefillEzsignformValueRequest) GetSEzsignformfieldEnteredvalueOk() (*string, bool)`

GetSEzsignformfieldEnteredvalueOk returns a tuple with the SEzsignformfieldEnteredvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldEnteredvalue

`func (o *CustomPrefillEzsignformValueRequest) SetSEzsignformfieldEnteredvalue(v string)`

SetSEzsignformfieldEnteredvalue sets SEzsignformfieldEnteredvalue field to given value.

### HasSEzsignformfieldEnteredvalue

`func (o *CustomPrefillEzsignformValueRequest) HasSEzsignformfieldEnteredvalue() bool`

HasSEzsignformfieldEnteredvalue returns a boolean if a field has been set.

### GetBEzsignformfieldSelected

`func (o *CustomPrefillEzsignformValueRequest) GetBEzsignformfieldSelected() bool`

GetBEzsignformfieldSelected returns the BEzsignformfieldSelected field if non-nil, zero value otherwise.

### GetBEzsignformfieldSelectedOk

`func (o *CustomPrefillEzsignformValueRequest) GetBEzsignformfieldSelectedOk() (*bool, bool)`

GetBEzsignformfieldSelectedOk returns a tuple with the BEzsignformfieldSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignformfieldSelected

`func (o *CustomPrefillEzsignformValueRequest) SetBEzsignformfieldSelected(v bool)`

SetBEzsignformfieldSelected sets BEzsignformfieldSelected field to given value.

### HasBEzsignformfieldSelected

`func (o *CustomPrefillEzsignformValueRequest) HasBEzsignformfieldSelected() bool`

HasBEzsignformfieldSelected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


