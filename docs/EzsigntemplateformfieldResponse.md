# EzsigntemplateformfieldResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateformfieldID** | **int32** | The unique ID of the Ezsigntemplateformfield | 
**IEzsigntemplatedocumentpagePagenumber** | **int32** | The page number in the Ezsigntemplatedocument | 
**SEzsigntemplateformfieldLabel** | **string** | The Label for the Ezsigntemplateformfield | 
**SEzsigntemplateformfieldValue** | Pointer to **string** | The value for the Ezsigntemplateformfield | [optional] 
**IEzsigntemplateformfieldX** | **int32** | The X coordinate (Horizontal) where to put the Ezsigntemplateformfield on the Ezsigntemplatepage.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplateformfield 2 inches from the left border of the page, you would use \&quot;200\&quot; for the X coordinate. | 
**IEzsigntemplateformfieldY** | **int32** | The Y coordinate (Vertical) where to put the Ezsigntemplateformfield on the Ezsigntemplatepage.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplateformfield 3 inches from the top border of the page, you would use \&quot;300\&quot; for the Y coordinate. | 
**IEzsigntemplateformfieldWidth** | **int32** | The Width of the Ezsigntemplateformfield in pixels calculated at 100 DPI  The allowed values are varying based on the eEzsigntemplateformfieldgroupType.  | eEzsigntemplateformfieldgroupType | Valid values | | ------------------------- | ------------ | | Checkbox                  | 22           | | Dropdown                  | 22-65535     | | Radio                     | 22           | | Text                      | 22-65535     | | Textarea                  | 22-65535     | | 
**IEzsigntemplateformfieldHeight** | **int32** | The Height of the Ezsigntemplateformfield in pixels calculated at 100 DPI  The allowed values are varying based on the eEzsigntemplateformfieldgroupType.  | eEzsigntemplateformfieldgroupType | Valid values | | ------------------------- | ------------ | | Checkbox                  | 22           | | Dropdown                  | 22           | | Radio                     | 22           | | Text                      | 22           | | Textarea                  | 22-65535     |  | 
**BEzsigntemplateformfieldSelected** | Pointer to **bool** | Whether the Ezsigntemplateformfield is selected or not by default.  This can only be set if eEzsigntemplateformfieldgroupType is **Checkbox** or **Radio** | [optional] 

## Methods

### NewEzsigntemplateformfieldResponse

`func NewEzsigntemplateformfieldResponse(pkiEzsigntemplateformfieldID int32, iEzsigntemplatedocumentpagePagenumber int32, sEzsigntemplateformfieldLabel string, iEzsigntemplateformfieldX int32, iEzsigntemplateformfieldY int32, iEzsigntemplateformfieldWidth int32, iEzsigntemplateformfieldHeight int32, ) *EzsigntemplateformfieldResponse`

NewEzsigntemplateformfieldResponse instantiates a new EzsigntemplateformfieldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateformfieldResponseWithDefaults

`func NewEzsigntemplateformfieldResponseWithDefaults() *EzsigntemplateformfieldResponse`

NewEzsigntemplateformfieldResponseWithDefaults instantiates a new EzsigntemplateformfieldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateformfieldID

`func (o *EzsigntemplateformfieldResponse) GetPkiEzsigntemplateformfieldID() int32`

GetPkiEzsigntemplateformfieldID returns the PkiEzsigntemplateformfieldID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateformfieldIDOk

`func (o *EzsigntemplateformfieldResponse) GetPkiEzsigntemplateformfieldIDOk() (*int32, bool)`

GetPkiEzsigntemplateformfieldIDOk returns a tuple with the PkiEzsigntemplateformfieldID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateformfieldID

`func (o *EzsigntemplateformfieldResponse) SetPkiEzsigntemplateformfieldID(v int32)`

SetPkiEzsigntemplateformfieldID sets PkiEzsigntemplateformfieldID field to given value.


### GetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplateformfieldResponse) GetIEzsigntemplatedocumentpagePagenumber() int32`

GetIEzsigntemplatedocumentpagePagenumber returns the IEzsigntemplatedocumentpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagePagenumberOk

`func (o *EzsigntemplateformfieldResponse) GetIEzsigntemplatedocumentpagePagenumberOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagePagenumberOk returns a tuple with the IEzsigntemplatedocumentpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplateformfieldResponse) SetIEzsigntemplatedocumentpagePagenumber(v int32)`

SetIEzsigntemplatedocumentpagePagenumber sets IEzsigntemplatedocumentpagePagenumber field to given value.


### GetSEzsigntemplateformfieldLabel

`func (o *EzsigntemplateformfieldResponse) GetSEzsigntemplateformfieldLabel() string`

GetSEzsigntemplateformfieldLabel returns the SEzsigntemplateformfieldLabel field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldLabelOk

`func (o *EzsigntemplateformfieldResponse) GetSEzsigntemplateformfieldLabelOk() (*string, bool)`

GetSEzsigntemplateformfieldLabelOk returns a tuple with the SEzsigntemplateformfieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldLabel

`func (o *EzsigntemplateformfieldResponse) SetSEzsigntemplateformfieldLabel(v string)`

SetSEzsigntemplateformfieldLabel sets SEzsigntemplateformfieldLabel field to given value.


### GetSEzsigntemplateformfieldValue

`func (o *EzsigntemplateformfieldResponse) GetSEzsigntemplateformfieldValue() string`

GetSEzsigntemplateformfieldValue returns the SEzsigntemplateformfieldValue field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldValueOk

`func (o *EzsigntemplateformfieldResponse) GetSEzsigntemplateformfieldValueOk() (*string, bool)`

GetSEzsigntemplateformfieldValueOk returns a tuple with the SEzsigntemplateformfieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldValue

`func (o *EzsigntemplateformfieldResponse) SetSEzsigntemplateformfieldValue(v string)`

SetSEzsigntemplateformfieldValue sets SEzsigntemplateformfieldValue field to given value.

### HasSEzsigntemplateformfieldValue

`func (o *EzsigntemplateformfieldResponse) HasSEzsigntemplateformfieldValue() bool`

HasSEzsigntemplateformfieldValue returns a boolean if a field has been set.

### GetIEzsigntemplateformfieldX

`func (o *EzsigntemplateformfieldResponse) GetIEzsigntemplateformfieldX() int32`

GetIEzsigntemplateformfieldX returns the IEzsigntemplateformfieldX field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldXOk

`func (o *EzsigntemplateformfieldResponse) GetIEzsigntemplateformfieldXOk() (*int32, bool)`

GetIEzsigntemplateformfieldXOk returns a tuple with the IEzsigntemplateformfieldX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldX

`func (o *EzsigntemplateformfieldResponse) SetIEzsigntemplateformfieldX(v int32)`

SetIEzsigntemplateformfieldX sets IEzsigntemplateformfieldX field to given value.


### GetIEzsigntemplateformfieldY

`func (o *EzsigntemplateformfieldResponse) GetIEzsigntemplateformfieldY() int32`

GetIEzsigntemplateformfieldY returns the IEzsigntemplateformfieldY field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldYOk

`func (o *EzsigntemplateformfieldResponse) GetIEzsigntemplateformfieldYOk() (*int32, bool)`

GetIEzsigntemplateformfieldYOk returns a tuple with the IEzsigntemplateformfieldY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldY

`func (o *EzsigntemplateformfieldResponse) SetIEzsigntemplateformfieldY(v int32)`

SetIEzsigntemplateformfieldY sets IEzsigntemplateformfieldY field to given value.


### GetIEzsigntemplateformfieldWidth

`func (o *EzsigntemplateformfieldResponse) GetIEzsigntemplateformfieldWidth() int32`

GetIEzsigntemplateformfieldWidth returns the IEzsigntemplateformfieldWidth field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldWidthOk

`func (o *EzsigntemplateformfieldResponse) GetIEzsigntemplateformfieldWidthOk() (*int32, bool)`

GetIEzsigntemplateformfieldWidthOk returns a tuple with the IEzsigntemplateformfieldWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldWidth

`func (o *EzsigntemplateformfieldResponse) SetIEzsigntemplateformfieldWidth(v int32)`

SetIEzsigntemplateformfieldWidth sets IEzsigntemplateformfieldWidth field to given value.


### GetIEzsigntemplateformfieldHeight

`func (o *EzsigntemplateformfieldResponse) GetIEzsigntemplateformfieldHeight() int32`

GetIEzsigntemplateformfieldHeight returns the IEzsigntemplateformfieldHeight field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldHeightOk

`func (o *EzsigntemplateformfieldResponse) GetIEzsigntemplateformfieldHeightOk() (*int32, bool)`

GetIEzsigntemplateformfieldHeightOk returns a tuple with the IEzsigntemplateformfieldHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldHeight

`func (o *EzsigntemplateformfieldResponse) SetIEzsigntemplateformfieldHeight(v int32)`

SetIEzsigntemplateformfieldHeight sets IEzsigntemplateformfieldHeight field to given value.


### GetBEzsigntemplateformfieldSelected

`func (o *EzsigntemplateformfieldResponse) GetBEzsigntemplateformfieldSelected() bool`

GetBEzsigntemplateformfieldSelected returns the BEzsigntemplateformfieldSelected field if non-nil, zero value otherwise.

### GetBEzsigntemplateformfieldSelectedOk

`func (o *EzsigntemplateformfieldResponse) GetBEzsigntemplateformfieldSelectedOk() (*bool, bool)`

GetBEzsigntemplateformfieldSelectedOk returns a tuple with the BEzsigntemplateformfieldSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateformfieldSelected

`func (o *EzsigntemplateformfieldResponse) SetBEzsigntemplateformfieldSelected(v bool)`

SetBEzsigntemplateformfieldSelected sets BEzsigntemplateformfieldSelected field to given value.

### HasBEzsigntemplateformfieldSelected

`func (o *EzsigntemplateformfieldResponse) HasBEzsigntemplateformfieldSelected() bool`

HasBEzsigntemplateformfieldSelected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


