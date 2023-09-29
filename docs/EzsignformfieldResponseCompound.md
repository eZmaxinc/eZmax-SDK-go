# EzsignformfieldResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignformfieldID** | **int32** | The unique ID of the Ezsignformfield | 
**IEzsignpagePagenumber** | **int32** | The page number in the Ezsigndocument | 
**SEzsignformfieldLabel** | **string** | The Label for the Ezsignformfield | 
**SEzsignformfieldValue** | Pointer to **string** | The value for the Ezsignformfield  This can only be set if eEzsignformfieldgroupType is Checkbox or Radio | [optional] 
**IEzsignformfieldX** | **int32** | The X coordinate (Horizontal) where to put the Ezsignformfield on the Ezsignpage.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignformfield 2 inches from the left border of the page, you would use \&quot;200\&quot; for the X coordinate. | 
**IEzsignformfieldY** | **int32** | The Y coordinate (Vertical) where to put the Ezsignformfield on the Ezsignpage.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignformfield 3 inches from the top border of the page, you would use \&quot;300\&quot; for the Y coordinate. | 
**IEzsignformfieldWidth** | **int32** | The Width of the Ezsignformfield in pixels calculated at 100 DPI  The allowed values are varying based on the eEzsignformfieldgroupType.  | eEzsignformfieldgroupType | Valid values | | ------------------------- | ------------ | | Checkbox                  | 22           | | Dropdown                  | 22-65535     | | Radio                     | 22           | | Text                      | 22-65535     | | Textarea                  | 22-65535     | | 
**IEzsignformfieldHeight** | **int32** | The Height of the Ezsignformfield in pixels calculated at 100 DPI  The allowed values are varying based on the eEzsignformfieldgroupType.  | eEzsignformfieldgroupType | Valid values | | ------------------------- | ------------ | | Checkbox                  | 22           | | Dropdown                  | 22           | | Radio                     | 22           | | Text                      | 22           | | Textarea                  | 22-65535     |  | 
**BEzsignformfieldAutocomplete** | Pointer to **bool** | Whether the Ezsignformfield allows the use of the autocomplete of the browser.  This can only be set if eEzsignformfieldgroupType is **Text** | [optional] 
**BEzsignformfieldSelected** | Pointer to **bool** | Whether the Ezsignformfield is selected or not by default.  This can only be set if eEzsignformfieldgroupType is **Checkbox** or **Radio** | [optional] 
**SEzsignformfieldEnteredvalue** | Pointer to **string** | This is the value enterred for the Ezsignformfield  This can only be set if eEzsignformfieldgroupType is **Dropdown**, **Text** or **Textarea** | [optional] 

## Methods

### NewEzsignformfieldResponseCompound

`func NewEzsignformfieldResponseCompound(pkiEzsignformfieldID int32, iEzsignpagePagenumber int32, sEzsignformfieldLabel string, iEzsignformfieldX int32, iEzsignformfieldY int32, iEzsignformfieldWidth int32, iEzsignformfieldHeight int32, ) *EzsignformfieldResponseCompound`

NewEzsignformfieldResponseCompound instantiates a new EzsignformfieldResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignformfieldResponseCompoundWithDefaults

`func NewEzsignformfieldResponseCompoundWithDefaults() *EzsignformfieldResponseCompound`

NewEzsignformfieldResponseCompoundWithDefaults instantiates a new EzsignformfieldResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignformfieldID

`func (o *EzsignformfieldResponseCompound) GetPkiEzsignformfieldID() int32`

GetPkiEzsignformfieldID returns the PkiEzsignformfieldID field if non-nil, zero value otherwise.

### GetPkiEzsignformfieldIDOk

`func (o *EzsignformfieldResponseCompound) GetPkiEzsignformfieldIDOk() (*int32, bool)`

GetPkiEzsignformfieldIDOk returns a tuple with the PkiEzsignformfieldID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignformfieldID

`func (o *EzsignformfieldResponseCompound) SetPkiEzsignformfieldID(v int32)`

SetPkiEzsignformfieldID sets PkiEzsignformfieldID field to given value.


### GetIEzsignpagePagenumber

`func (o *EzsignformfieldResponseCompound) GetIEzsignpagePagenumber() int32`

GetIEzsignpagePagenumber returns the IEzsignpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsignpagePagenumberOk

`func (o *EzsignformfieldResponseCompound) GetIEzsignpagePagenumberOk() (*int32, bool)`

GetIEzsignpagePagenumberOk returns a tuple with the IEzsignpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignpagePagenumber

`func (o *EzsignformfieldResponseCompound) SetIEzsignpagePagenumber(v int32)`

SetIEzsignpagePagenumber sets IEzsignpagePagenumber field to given value.


### GetSEzsignformfieldLabel

`func (o *EzsignformfieldResponseCompound) GetSEzsignformfieldLabel() string`

GetSEzsignformfieldLabel returns the SEzsignformfieldLabel field if non-nil, zero value otherwise.

### GetSEzsignformfieldLabelOk

`func (o *EzsignformfieldResponseCompound) GetSEzsignformfieldLabelOk() (*string, bool)`

GetSEzsignformfieldLabelOk returns a tuple with the SEzsignformfieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldLabel

`func (o *EzsignformfieldResponseCompound) SetSEzsignformfieldLabel(v string)`

SetSEzsignformfieldLabel sets SEzsignformfieldLabel field to given value.


### GetSEzsignformfieldValue

`func (o *EzsignformfieldResponseCompound) GetSEzsignformfieldValue() string`

GetSEzsignformfieldValue returns the SEzsignformfieldValue field if non-nil, zero value otherwise.

### GetSEzsignformfieldValueOk

`func (o *EzsignformfieldResponseCompound) GetSEzsignformfieldValueOk() (*string, bool)`

GetSEzsignformfieldValueOk returns a tuple with the SEzsignformfieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldValue

`func (o *EzsignformfieldResponseCompound) SetSEzsignformfieldValue(v string)`

SetSEzsignformfieldValue sets SEzsignformfieldValue field to given value.

### HasSEzsignformfieldValue

`func (o *EzsignformfieldResponseCompound) HasSEzsignformfieldValue() bool`

HasSEzsignformfieldValue returns a boolean if a field has been set.

### GetIEzsignformfieldX

`func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldX() int32`

GetIEzsignformfieldX returns the IEzsignformfieldX field if non-nil, zero value otherwise.

### GetIEzsignformfieldXOk

`func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldXOk() (*int32, bool)`

GetIEzsignformfieldXOk returns a tuple with the IEzsignformfieldX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldX

`func (o *EzsignformfieldResponseCompound) SetIEzsignformfieldX(v int32)`

SetIEzsignformfieldX sets IEzsignformfieldX field to given value.


### GetIEzsignformfieldY

`func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldY() int32`

GetIEzsignformfieldY returns the IEzsignformfieldY field if non-nil, zero value otherwise.

### GetIEzsignformfieldYOk

`func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldYOk() (*int32, bool)`

GetIEzsignformfieldYOk returns a tuple with the IEzsignformfieldY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldY

`func (o *EzsignformfieldResponseCompound) SetIEzsignformfieldY(v int32)`

SetIEzsignformfieldY sets IEzsignformfieldY field to given value.


### GetIEzsignformfieldWidth

`func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldWidth() int32`

GetIEzsignformfieldWidth returns the IEzsignformfieldWidth field if non-nil, zero value otherwise.

### GetIEzsignformfieldWidthOk

`func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldWidthOk() (*int32, bool)`

GetIEzsignformfieldWidthOk returns a tuple with the IEzsignformfieldWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldWidth

`func (o *EzsignformfieldResponseCompound) SetIEzsignformfieldWidth(v int32)`

SetIEzsignformfieldWidth sets IEzsignformfieldWidth field to given value.


### GetIEzsignformfieldHeight

`func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldHeight() int32`

GetIEzsignformfieldHeight returns the IEzsignformfieldHeight field if non-nil, zero value otherwise.

### GetIEzsignformfieldHeightOk

`func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldHeightOk() (*int32, bool)`

GetIEzsignformfieldHeightOk returns a tuple with the IEzsignformfieldHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignformfieldHeight

`func (o *EzsignformfieldResponseCompound) SetIEzsignformfieldHeight(v int32)`

SetIEzsignformfieldHeight sets IEzsignformfieldHeight field to given value.


### GetBEzsignformfieldAutocomplete

`func (o *EzsignformfieldResponseCompound) GetBEzsignformfieldAutocomplete() bool`

GetBEzsignformfieldAutocomplete returns the BEzsignformfieldAutocomplete field if non-nil, zero value otherwise.

### GetBEzsignformfieldAutocompleteOk

`func (o *EzsignformfieldResponseCompound) GetBEzsignformfieldAutocompleteOk() (*bool, bool)`

GetBEzsignformfieldAutocompleteOk returns a tuple with the BEzsignformfieldAutocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignformfieldAutocomplete

`func (o *EzsignformfieldResponseCompound) SetBEzsignformfieldAutocomplete(v bool)`

SetBEzsignformfieldAutocomplete sets BEzsignformfieldAutocomplete field to given value.

### HasBEzsignformfieldAutocomplete

`func (o *EzsignformfieldResponseCompound) HasBEzsignformfieldAutocomplete() bool`

HasBEzsignformfieldAutocomplete returns a boolean if a field has been set.

### GetBEzsignformfieldSelected

`func (o *EzsignformfieldResponseCompound) GetBEzsignformfieldSelected() bool`

GetBEzsignformfieldSelected returns the BEzsignformfieldSelected field if non-nil, zero value otherwise.

### GetBEzsignformfieldSelectedOk

`func (o *EzsignformfieldResponseCompound) GetBEzsignformfieldSelectedOk() (*bool, bool)`

GetBEzsignformfieldSelectedOk returns a tuple with the BEzsignformfieldSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignformfieldSelected

`func (o *EzsignformfieldResponseCompound) SetBEzsignformfieldSelected(v bool)`

SetBEzsignformfieldSelected sets BEzsignformfieldSelected field to given value.

### HasBEzsignformfieldSelected

`func (o *EzsignformfieldResponseCompound) HasBEzsignformfieldSelected() bool`

HasBEzsignformfieldSelected returns a boolean if a field has been set.

### GetSEzsignformfieldEnteredvalue

`func (o *EzsignformfieldResponseCompound) GetSEzsignformfieldEnteredvalue() string`

GetSEzsignformfieldEnteredvalue returns the SEzsignformfieldEnteredvalue field if non-nil, zero value otherwise.

### GetSEzsignformfieldEnteredvalueOk

`func (o *EzsignformfieldResponseCompound) GetSEzsignformfieldEnteredvalueOk() (*string, bool)`

GetSEzsignformfieldEnteredvalueOk returns a tuple with the SEzsignformfieldEnteredvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldEnteredvalue

`func (o *EzsignformfieldResponseCompound) SetSEzsignformfieldEnteredvalue(v string)`

SetSEzsignformfieldEnteredvalue sets SEzsignformfieldEnteredvalue field to given value.

### HasSEzsignformfieldEnteredvalue

`func (o *EzsignformfieldResponseCompound) HasSEzsignformfieldEnteredvalue() bool`

HasSEzsignformfieldEnteredvalue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


