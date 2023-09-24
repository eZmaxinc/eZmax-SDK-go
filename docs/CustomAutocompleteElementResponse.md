# CustomAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SCategory** | **string** | The Category for the dropdown or an empty string if not categorized | 
**SLabel** | **string** | The Description of the element | 
**SValue** | **string** | The Unique ID of the element | 
**MValue** | Pointer to **string** | The Unique ID of the element | [optional] 
**BActive** | **bool** | Indicates if the element is active | 

## Methods

### NewCustomAutocompleteElementResponse

`func NewCustomAutocompleteElementResponse(sCategory string, sLabel string, sValue string, bActive bool, ) *CustomAutocompleteElementResponse`

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


### GetSValue

`func (o *CustomAutocompleteElementResponse) GetSValue() string`

GetSValue returns the SValue field if non-nil, zero value otherwise.

### GetSValueOk

`func (o *CustomAutocompleteElementResponse) GetSValueOk() (*string, bool)`

GetSValueOk returns a tuple with the SValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSValue

`func (o *CustomAutocompleteElementResponse) SetSValue(v string)`

SetSValue sets SValue field to given value.


### GetMValue

`func (o *CustomAutocompleteElementResponse) GetMValue() string`

GetMValue returns the MValue field if non-nil, zero value otherwise.

### GetMValueOk

`func (o *CustomAutocompleteElementResponse) GetMValueOk() (*string, bool)`

GetMValueOk returns a tuple with the MValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMValue

`func (o *CustomAutocompleteElementResponse) SetMValue(v string)`

SetMValue sets MValue field to given value.

### HasMValue

`func (o *CustomAutocompleteElementResponse) HasMValue() bool`

HasMValue returns a boolean if a field has been set.

### GetBActive

`func (o *CustomAutocompleteElementResponse) GetBActive() bool`

GetBActive returns the BActive field if non-nil, zero value otherwise.

### GetBActiveOk

`func (o *CustomAutocompleteElementResponse) GetBActiveOk() (*bool, bool)`

GetBActiveOk returns a tuple with the BActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBActive

`func (o *CustomAutocompleteElementResponse) SetBActive(v bool)`

SetBActive sets BActive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


