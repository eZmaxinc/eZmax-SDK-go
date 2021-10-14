# CustomAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SCategory** | **string** | The Category for the dropdown or an empty string if not categorized | 
**SLabel** | **string** | The Description of the element | 
**MValue** | [**OneOfintegerstring**](oneOf&lt;integer,string&gt;.md) | The Unique ID of the element | 

## Methods

### NewCustomAutocompleteElementResponse

`func NewCustomAutocompleteElementResponse(sCategory string, sLabel string, mValue OneOfintegerstring, ) *CustomAutocompleteElementResponse`

NewCustomAutocompleteElementResponse instantiates a new CustomAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAutocompleteElementResponseWithDefaults

`func NewCustomAutocompleteElementResponseWithDefaults() *CustomAutocompleteElementResponse`

NewCustomAutocompleteElementResponseWithDefaults instantiates a new CustomAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSCategory

`func (o *CustomAutocompleteElementResponse) GetSCategory() string`

GetSCategory returns the SCategory field if non-nil, zero value otherwise.

### GetSCategoryOk

`func (o *CustomAutocompleteElementResponse) GetSCategoryOk() (*string, bool)`

GetSCategoryOk returns a tuple with the SCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCategory

`func (o *CustomAutocompleteElementResponse) SetSCategory(v string)`

SetSCategory sets SCategory field to given value.


### GetSLabel

`func (o *CustomAutocompleteElementResponse) GetSLabel() string`

GetSLabel returns the SLabel field if non-nil, zero value otherwise.

### GetSLabelOk

`func (o *CustomAutocompleteElementResponse) GetSLabelOk() (*string, bool)`

GetSLabelOk returns a tuple with the SLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLabel

`func (o *CustomAutocompleteElementResponse) SetSLabel(v string)`

SetSLabel sets SLabel field to given value.


### GetMValue

`func (o *CustomAutocompleteElementResponse) GetMValue() OneOfintegerstring`

GetMValue returns the MValue field if non-nil, zero value otherwise.

### GetMValueOk

`func (o *CustomAutocompleteElementResponse) GetMValueOk() (*OneOfintegerstring, bool)`

GetMValueOk returns a tuple with the MValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMValue

`func (o *CustomAutocompleteElementResponse) SetMValue(v OneOfintegerstring)`

SetMValue sets MValue field to given value.


### SetMValueNil

`func (o *CustomAutocompleteElementResponse) SetMValueNil(b bool)`

 SetMValueNil sets the value for MValue to be an explicit nil

### UnsetMValue
`func (o *CustomAutocompleteElementResponse) UnsetMValue()`

UnsetMValue ensures that no value is present for MValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


