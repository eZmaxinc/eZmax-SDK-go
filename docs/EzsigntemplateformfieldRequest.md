# EzsigntemplateformfieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateformfieldID** | Pointer to **int32** | The unique ID of the Ezsigntemplateformfield | [optional] 
**EEzsigntemplateformfieldPositioning** | Pointer to [**FieldEEzsigntemplateformfieldPositioning**](FieldEEzsigntemplateformfieldPositioning.md) |  | [optional] [default to PER_COORDINATES]
**IEzsigntemplatedocumentpagePagenumber** | **int32** | The page number in the Ezsigntemplatedocument | 
**SEzsigntemplateformfieldLabel** | **string** | The Label for the Ezsigntemplateformfield | 
**SEzsigntemplateformfieldValue** | Pointer to **string** | The value for the Ezsigntemplateformfield | [optional] 
**IEzsigntemplateformfieldX** | Pointer to **int32** | The X coordinate (Horizontal) where to put the Ezsigntemplateformfield on the Ezsigntemplatepage.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplateformfield 2 inches from the left border of the page, you would use \&quot;200\&quot; for the X coordinate. | [optional] 
**IEzsigntemplateformfieldY** | Pointer to **int32** | The Y coordinate (Vertical) where to put the Ezsigntemplateformfield on the Ezsigntemplatepage.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplateformfield 3 inches from the top border of the page, you would use \&quot;300\&quot; for the Y coordinate. | [optional] 
**IEzsigntemplateformfieldWidth** | **int32** | The Width of the Ezsigntemplateformfield in pixels calculated at 100 DPI | 
**IEzsigntemplateformfieldHeight** | **int32** | The Height of the Ezsigntemplateformfield in pixels calculated at 100 DPI  | 
**BEzsigntemplateformfieldAutocomplete** | Pointer to **bool** | Whether the Ezsigntemplateformfield allows the use of the autocomplete of the browser.  This can only be set if eEzsigntemplateformfieldgroupType is **Text** | [optional] 
**BEzsigntemplateformfieldSelected** | Pointer to **bool** | Whether the Ezsigntemplateformfield is selected or not by default.  This can only be set if eEzsigntemplateformfieldgroupType is **Checkbox** or **Radio** | [optional] 
**EEzsigntemplateformfieldDependencyrequirement** | Pointer to [**FieldEEzsigntemplateformfieldDependencyrequirement**](FieldEEzsigntemplateformfieldDependencyrequirement.md) |  | [optional] 
**SEzsigntemplateformfieldPositioningpattern** | Pointer to **string** | The string pattern to search for the positioning. **This is not a regexp**  This will be required if **eEzsigntemplateformfieldPositioning** is set to **PerCoordinates** | [optional] 
**IEzsigntemplateformfieldPositioningoffsetx** | Pointer to **int32** | The offset X  This will be required if **eEzsigntemplateformfieldPositioning** is set to **PerCoordinates** | [optional] 
**IEzsigntemplateformfieldPositioningoffsety** | Pointer to **int32** | The offset Y  This will be required if **eEzsigntemplateformfieldPositioning** is set to **PerCoordinates** | [optional] 
**EEzsigntemplateformfieldPositioningoccurence** | Pointer to [**FieldEEzsigntemplateformfieldPositioningoccurence**](FieldEEzsigntemplateformfieldPositioningoccurence.md) |  | [optional] 
**EEzsigntemplateformfieldHorizontalalignment** | Pointer to [**EnumHorizontalalignment**](EnumHorizontalalignment.md) |  | [optional] 
**ObjTextstylestatic** | Pointer to [**TextstylestaticRequestCompound**](TextstylestaticRequestCompound.md) |  | [optional] 

## Methods

### NewEzsigntemplateformfieldRequest

`func NewEzsigntemplateformfieldRequest(iEzsigntemplatedocumentpagePagenumber int32, sEzsigntemplateformfieldLabel string, iEzsigntemplateformfieldWidth int32, iEzsigntemplateformfieldHeight int32, ) *EzsigntemplateformfieldRequest`

NewEzsigntemplateformfieldRequest instantiates a new EzsigntemplateformfieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateformfieldRequestWithDefaults

`func NewEzsigntemplateformfieldRequestWithDefaults() *EzsigntemplateformfieldRequest`

NewEzsigntemplateformfieldRequestWithDefaults instantiates a new EzsigntemplateformfieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateformfieldID

`func (o *EzsigntemplateformfieldRequest) GetPkiEzsigntemplateformfieldID() int32`

GetPkiEzsigntemplateformfieldID returns the PkiEzsigntemplateformfieldID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateformfieldIDOk

`func (o *EzsigntemplateformfieldRequest) GetPkiEzsigntemplateformfieldIDOk() (*int32, bool)`

GetPkiEzsigntemplateformfieldIDOk returns a tuple with the PkiEzsigntemplateformfieldID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateformfieldID

`func (o *EzsigntemplateformfieldRequest) SetPkiEzsigntemplateformfieldID(v int32)`

SetPkiEzsigntemplateformfieldID sets PkiEzsigntemplateformfieldID field to given value.

### HasPkiEzsigntemplateformfieldID

`func (o *EzsigntemplateformfieldRequest) HasPkiEzsigntemplateformfieldID() bool`

HasPkiEzsigntemplateformfieldID returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldPositioning

`func (o *EzsigntemplateformfieldRequest) GetEEzsigntemplateformfieldPositioning() FieldEEzsigntemplateformfieldPositioning`

GetEEzsigntemplateformfieldPositioning returns the EEzsigntemplateformfieldPositioning field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldPositioningOk

`func (o *EzsigntemplateformfieldRequest) GetEEzsigntemplateformfieldPositioningOk() (*FieldEEzsigntemplateformfieldPositioning, bool)`

GetEEzsigntemplateformfieldPositioningOk returns a tuple with the EEzsigntemplateformfieldPositioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldPositioning

`func (o *EzsigntemplateformfieldRequest) SetEEzsigntemplateformfieldPositioning(v FieldEEzsigntemplateformfieldPositioning)`

SetEEzsigntemplateformfieldPositioning sets EEzsigntemplateformfieldPositioning field to given value.

### HasEEzsigntemplateformfieldPositioning

`func (o *EzsigntemplateformfieldRequest) HasEEzsigntemplateformfieldPositioning() bool`

HasEEzsigntemplateformfieldPositioning returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplateformfieldRequest) GetIEzsigntemplatedocumentpagePagenumber() int32`

GetIEzsigntemplatedocumentpagePagenumber returns the IEzsigntemplatedocumentpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagePagenumberOk

`func (o *EzsigntemplateformfieldRequest) GetIEzsigntemplatedocumentpagePagenumberOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagePagenumberOk returns a tuple with the IEzsigntemplatedocumentpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplateformfieldRequest) SetIEzsigntemplatedocumentpagePagenumber(v int32)`

SetIEzsigntemplatedocumentpagePagenumber sets IEzsigntemplatedocumentpagePagenumber field to given value.


### GetSEzsigntemplateformfieldLabel

`func (o *EzsigntemplateformfieldRequest) GetSEzsigntemplateformfieldLabel() string`

GetSEzsigntemplateformfieldLabel returns the SEzsigntemplateformfieldLabel field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldLabelOk

`func (o *EzsigntemplateformfieldRequest) GetSEzsigntemplateformfieldLabelOk() (*string, bool)`

GetSEzsigntemplateformfieldLabelOk returns a tuple with the SEzsigntemplateformfieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldLabel

`func (o *EzsigntemplateformfieldRequest) SetSEzsigntemplateformfieldLabel(v string)`

SetSEzsigntemplateformfieldLabel sets SEzsigntemplateformfieldLabel field to given value.


### GetSEzsigntemplateformfieldValue

`func (o *EzsigntemplateformfieldRequest) GetSEzsigntemplateformfieldValue() string`

GetSEzsigntemplateformfieldValue returns the SEzsigntemplateformfieldValue field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldValueOk

`func (o *EzsigntemplateformfieldRequest) GetSEzsigntemplateformfieldValueOk() (*string, bool)`

GetSEzsigntemplateformfieldValueOk returns a tuple with the SEzsigntemplateformfieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldValue

`func (o *EzsigntemplateformfieldRequest) SetSEzsigntemplateformfieldValue(v string)`

SetSEzsigntemplateformfieldValue sets SEzsigntemplateformfieldValue field to given value.

### HasSEzsigntemplateformfieldValue

`func (o *EzsigntemplateformfieldRequest) HasSEzsigntemplateformfieldValue() bool`

HasSEzsigntemplateformfieldValue returns a boolean if a field has been set.

### GetIEzsigntemplateformfieldX

`func (o *EzsigntemplateformfieldRequest) GetIEzsigntemplateformfieldX() int32`

GetIEzsigntemplateformfieldX returns the IEzsigntemplateformfieldX field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldXOk

`func (o *EzsigntemplateformfieldRequest) GetIEzsigntemplateformfieldXOk() (*int32, bool)`

GetIEzsigntemplateformfieldXOk returns a tuple with the IEzsigntemplateformfieldX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldX

`func (o *EzsigntemplateformfieldRequest) SetIEzsigntemplateformfieldX(v int32)`

SetIEzsigntemplateformfieldX sets IEzsigntemplateformfieldX field to given value.

### HasIEzsigntemplateformfieldX

`func (o *EzsigntemplateformfieldRequest) HasIEzsigntemplateformfieldX() bool`

HasIEzsigntemplateformfieldX returns a boolean if a field has been set.

### GetIEzsigntemplateformfieldY

`func (o *EzsigntemplateformfieldRequest) GetIEzsigntemplateformfieldY() int32`

GetIEzsigntemplateformfieldY returns the IEzsigntemplateformfieldY field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldYOk

`func (o *EzsigntemplateformfieldRequest) GetIEzsigntemplateformfieldYOk() (*int32, bool)`

GetIEzsigntemplateformfieldYOk returns a tuple with the IEzsigntemplateformfieldY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldY

`func (o *EzsigntemplateformfieldRequest) SetIEzsigntemplateformfieldY(v int32)`

SetIEzsigntemplateformfieldY sets IEzsigntemplateformfieldY field to given value.

### HasIEzsigntemplateformfieldY

`func (o *EzsigntemplateformfieldRequest) HasIEzsigntemplateformfieldY() bool`

HasIEzsigntemplateformfieldY returns a boolean if a field has been set.

### GetIEzsigntemplateformfieldWidth

`func (o *EzsigntemplateformfieldRequest) GetIEzsigntemplateformfieldWidth() int32`

GetIEzsigntemplateformfieldWidth returns the IEzsigntemplateformfieldWidth field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldWidthOk

`func (o *EzsigntemplateformfieldRequest) GetIEzsigntemplateformfieldWidthOk() (*int32, bool)`

GetIEzsigntemplateformfieldWidthOk returns a tuple with the IEzsigntemplateformfieldWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldWidth

`func (o *EzsigntemplateformfieldRequest) SetIEzsigntemplateformfieldWidth(v int32)`

SetIEzsigntemplateformfieldWidth sets IEzsigntemplateformfieldWidth field to given value.


### GetIEzsigntemplateformfieldHeight

`func (o *EzsigntemplateformfieldRequest) GetIEzsigntemplateformfieldHeight() int32`

GetIEzsigntemplateformfieldHeight returns the IEzsigntemplateformfieldHeight field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldHeightOk

`func (o *EzsigntemplateformfieldRequest) GetIEzsigntemplateformfieldHeightOk() (*int32, bool)`

GetIEzsigntemplateformfieldHeightOk returns a tuple with the IEzsigntemplateformfieldHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldHeight

`func (o *EzsigntemplateformfieldRequest) SetIEzsigntemplateformfieldHeight(v int32)`

SetIEzsigntemplateformfieldHeight sets IEzsigntemplateformfieldHeight field to given value.


### GetBEzsigntemplateformfieldAutocomplete

`func (o *EzsigntemplateformfieldRequest) GetBEzsigntemplateformfieldAutocomplete() bool`

GetBEzsigntemplateformfieldAutocomplete returns the BEzsigntemplateformfieldAutocomplete field if non-nil, zero value otherwise.

### GetBEzsigntemplateformfieldAutocompleteOk

`func (o *EzsigntemplateformfieldRequest) GetBEzsigntemplateformfieldAutocompleteOk() (*bool, bool)`

GetBEzsigntemplateformfieldAutocompleteOk returns a tuple with the BEzsigntemplateformfieldAutocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateformfieldAutocomplete

`func (o *EzsigntemplateformfieldRequest) SetBEzsigntemplateformfieldAutocomplete(v bool)`

SetBEzsigntemplateformfieldAutocomplete sets BEzsigntemplateformfieldAutocomplete field to given value.

### HasBEzsigntemplateformfieldAutocomplete

`func (o *EzsigntemplateformfieldRequest) HasBEzsigntemplateformfieldAutocomplete() bool`

HasBEzsigntemplateformfieldAutocomplete returns a boolean if a field has been set.

### GetBEzsigntemplateformfieldSelected

`func (o *EzsigntemplateformfieldRequest) GetBEzsigntemplateformfieldSelected() bool`

GetBEzsigntemplateformfieldSelected returns the BEzsigntemplateformfieldSelected field if non-nil, zero value otherwise.

### GetBEzsigntemplateformfieldSelectedOk

`func (o *EzsigntemplateformfieldRequest) GetBEzsigntemplateformfieldSelectedOk() (*bool, bool)`

GetBEzsigntemplateformfieldSelectedOk returns a tuple with the BEzsigntemplateformfieldSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateformfieldSelected

`func (o *EzsigntemplateformfieldRequest) SetBEzsigntemplateformfieldSelected(v bool)`

SetBEzsigntemplateformfieldSelected sets BEzsigntemplateformfieldSelected field to given value.

### HasBEzsigntemplateformfieldSelected

`func (o *EzsigntemplateformfieldRequest) HasBEzsigntemplateformfieldSelected() bool`

HasBEzsigntemplateformfieldSelected returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldDependencyrequirement

`func (o *EzsigntemplateformfieldRequest) GetEEzsigntemplateformfieldDependencyrequirement() FieldEEzsigntemplateformfieldDependencyrequirement`

GetEEzsigntemplateformfieldDependencyrequirement returns the EEzsigntemplateformfieldDependencyrequirement field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldDependencyrequirementOk

`func (o *EzsigntemplateformfieldRequest) GetEEzsigntemplateformfieldDependencyrequirementOk() (*FieldEEzsigntemplateformfieldDependencyrequirement, bool)`

GetEEzsigntemplateformfieldDependencyrequirementOk returns a tuple with the EEzsigntemplateformfieldDependencyrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldDependencyrequirement

`func (o *EzsigntemplateformfieldRequest) SetEEzsigntemplateformfieldDependencyrequirement(v FieldEEzsigntemplateformfieldDependencyrequirement)`

SetEEzsigntemplateformfieldDependencyrequirement sets EEzsigntemplateformfieldDependencyrequirement field to given value.

### HasEEzsigntemplateformfieldDependencyrequirement

`func (o *EzsigntemplateformfieldRequest) HasEEzsigntemplateformfieldDependencyrequirement() bool`

HasEEzsigntemplateformfieldDependencyrequirement returns a boolean if a field has been set.

### GetSEzsigntemplateformfieldPositioningpattern

`func (o *EzsigntemplateformfieldRequest) GetSEzsigntemplateformfieldPositioningpattern() string`

GetSEzsigntemplateformfieldPositioningpattern returns the SEzsigntemplateformfieldPositioningpattern field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldPositioningpatternOk

`func (o *EzsigntemplateformfieldRequest) GetSEzsigntemplateformfieldPositioningpatternOk() (*string, bool)`

GetSEzsigntemplateformfieldPositioningpatternOk returns a tuple with the SEzsigntemplateformfieldPositioningpattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldPositioningpattern

`func (o *EzsigntemplateformfieldRequest) SetSEzsigntemplateformfieldPositioningpattern(v string)`

SetSEzsigntemplateformfieldPositioningpattern sets SEzsigntemplateformfieldPositioningpattern field to given value.

### HasSEzsigntemplateformfieldPositioningpattern

`func (o *EzsigntemplateformfieldRequest) HasSEzsigntemplateformfieldPositioningpattern() bool`

HasSEzsigntemplateformfieldPositioningpattern returns a boolean if a field has been set.

### GetIEzsigntemplateformfieldPositioningoffsetx

`func (o *EzsigntemplateformfieldRequest) GetIEzsigntemplateformfieldPositioningoffsetx() int32`

GetIEzsigntemplateformfieldPositioningoffsetx returns the IEzsigntemplateformfieldPositioningoffsetx field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldPositioningoffsetxOk

`func (o *EzsigntemplateformfieldRequest) GetIEzsigntemplateformfieldPositioningoffsetxOk() (*int32, bool)`

GetIEzsigntemplateformfieldPositioningoffsetxOk returns a tuple with the IEzsigntemplateformfieldPositioningoffsetx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldPositioningoffsetx

`func (o *EzsigntemplateformfieldRequest) SetIEzsigntemplateformfieldPositioningoffsetx(v int32)`

SetIEzsigntemplateformfieldPositioningoffsetx sets IEzsigntemplateformfieldPositioningoffsetx field to given value.

### HasIEzsigntemplateformfieldPositioningoffsetx

`func (o *EzsigntemplateformfieldRequest) HasIEzsigntemplateformfieldPositioningoffsetx() bool`

HasIEzsigntemplateformfieldPositioningoffsetx returns a boolean if a field has been set.

### GetIEzsigntemplateformfieldPositioningoffsety

`func (o *EzsigntemplateformfieldRequest) GetIEzsigntemplateformfieldPositioningoffsety() int32`

GetIEzsigntemplateformfieldPositioningoffsety returns the IEzsigntemplateformfieldPositioningoffsety field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldPositioningoffsetyOk

`func (o *EzsigntemplateformfieldRequest) GetIEzsigntemplateformfieldPositioningoffsetyOk() (*int32, bool)`

GetIEzsigntemplateformfieldPositioningoffsetyOk returns a tuple with the IEzsigntemplateformfieldPositioningoffsety field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldPositioningoffsety

`func (o *EzsigntemplateformfieldRequest) SetIEzsigntemplateformfieldPositioningoffsety(v int32)`

SetIEzsigntemplateformfieldPositioningoffsety sets IEzsigntemplateformfieldPositioningoffsety field to given value.

### HasIEzsigntemplateformfieldPositioningoffsety

`func (o *EzsigntemplateformfieldRequest) HasIEzsigntemplateformfieldPositioningoffsety() bool`

HasIEzsigntemplateformfieldPositioningoffsety returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldPositioningoccurence

`func (o *EzsigntemplateformfieldRequest) GetEEzsigntemplateformfieldPositioningoccurence() FieldEEzsigntemplateformfieldPositioningoccurence`

GetEEzsigntemplateformfieldPositioningoccurence returns the EEzsigntemplateformfieldPositioningoccurence field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldPositioningoccurenceOk

`func (o *EzsigntemplateformfieldRequest) GetEEzsigntemplateformfieldPositioningoccurenceOk() (*FieldEEzsigntemplateformfieldPositioningoccurence, bool)`

GetEEzsigntemplateformfieldPositioningoccurenceOk returns a tuple with the EEzsigntemplateformfieldPositioningoccurence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldPositioningoccurence

`func (o *EzsigntemplateformfieldRequest) SetEEzsigntemplateformfieldPositioningoccurence(v FieldEEzsigntemplateformfieldPositioningoccurence)`

SetEEzsigntemplateformfieldPositioningoccurence sets EEzsigntemplateformfieldPositioningoccurence field to given value.

### HasEEzsigntemplateformfieldPositioningoccurence

`func (o *EzsigntemplateformfieldRequest) HasEEzsigntemplateformfieldPositioningoccurence() bool`

HasEEzsigntemplateformfieldPositioningoccurence returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldHorizontalalignment

`func (o *EzsigntemplateformfieldRequest) GetEEzsigntemplateformfieldHorizontalalignment() EnumHorizontalalignment`

GetEEzsigntemplateformfieldHorizontalalignment returns the EEzsigntemplateformfieldHorizontalalignment field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldHorizontalalignmentOk

`func (o *EzsigntemplateformfieldRequest) GetEEzsigntemplateformfieldHorizontalalignmentOk() (*EnumHorizontalalignment, bool)`

GetEEzsigntemplateformfieldHorizontalalignmentOk returns a tuple with the EEzsigntemplateformfieldHorizontalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldHorizontalalignment

`func (o *EzsigntemplateformfieldRequest) SetEEzsigntemplateformfieldHorizontalalignment(v EnumHorizontalalignment)`

SetEEzsigntemplateformfieldHorizontalalignment sets EEzsigntemplateformfieldHorizontalalignment field to given value.

### HasEEzsigntemplateformfieldHorizontalalignment

`func (o *EzsigntemplateformfieldRequest) HasEEzsigntemplateformfieldHorizontalalignment() bool`

HasEEzsigntemplateformfieldHorizontalalignment returns a boolean if a field has been set.

### GetObjTextstylestatic

`func (o *EzsigntemplateformfieldRequest) GetObjTextstylestatic() TextstylestaticRequestCompound`

GetObjTextstylestatic returns the ObjTextstylestatic field if non-nil, zero value otherwise.

### GetObjTextstylestaticOk

`func (o *EzsigntemplateformfieldRequest) GetObjTextstylestaticOk() (*TextstylestaticRequestCompound, bool)`

GetObjTextstylestaticOk returns a tuple with the ObjTextstylestatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjTextstylestatic

`func (o *EzsigntemplateformfieldRequest) SetObjTextstylestatic(v TextstylestaticRequestCompound)`

SetObjTextstylestatic sets ObjTextstylestatic field to given value.

### HasObjTextstylestatic

`func (o *EzsigntemplateformfieldRequest) HasObjTextstylestatic() bool`

HasObjTextstylestatic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


