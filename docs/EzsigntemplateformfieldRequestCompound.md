# EzsigntemplateformfieldRequestCompound

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
**AObjEzsigntemplateelementdependency** | Pointer to [**[]EzsigntemplateelementdependencyRequestCompound**](EzsigntemplateelementdependencyRequestCompound.md) |  | [optional] 

## Methods

### NewEzsigntemplateformfieldRequestCompound

`func NewEzsigntemplateformfieldRequestCompound(iEzsigntemplatedocumentpagePagenumber int32, sEzsigntemplateformfieldLabel string, iEzsigntemplateformfieldWidth int32, iEzsigntemplateformfieldHeight int32, ) *EzsigntemplateformfieldRequestCompound`

NewEzsigntemplateformfieldRequestCompound instantiates a new EzsigntemplateformfieldRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateformfieldRequestCompoundWithDefaults

`func NewEzsigntemplateformfieldRequestCompoundWithDefaults() *EzsigntemplateformfieldRequestCompound`

NewEzsigntemplateformfieldRequestCompoundWithDefaults instantiates a new EzsigntemplateformfieldRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateformfieldID

`func (o *EzsigntemplateformfieldRequestCompound) GetPkiEzsigntemplateformfieldID() int32`

GetPkiEzsigntemplateformfieldID returns the PkiEzsigntemplateformfieldID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateformfieldIDOk

`func (o *EzsigntemplateformfieldRequestCompound) GetPkiEzsigntemplateformfieldIDOk() (*int32, bool)`

GetPkiEzsigntemplateformfieldIDOk returns a tuple with the PkiEzsigntemplateformfieldID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateformfieldID

`func (o *EzsigntemplateformfieldRequestCompound) SetPkiEzsigntemplateformfieldID(v int32)`

SetPkiEzsigntemplateformfieldID sets PkiEzsigntemplateformfieldID field to given value.

### HasPkiEzsigntemplateformfieldID

`func (o *EzsigntemplateformfieldRequestCompound) HasPkiEzsigntemplateformfieldID() bool`

HasPkiEzsigntemplateformfieldID returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldPositioning

`func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldPositioning() FieldEEzsigntemplateformfieldPositioning`

GetEEzsigntemplateformfieldPositioning returns the EEzsigntemplateformfieldPositioning field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldPositioningOk

`func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldPositioningOk() (*FieldEEzsigntemplateformfieldPositioning, bool)`

GetEEzsigntemplateformfieldPositioningOk returns a tuple with the EEzsigntemplateformfieldPositioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldPositioning

`func (o *EzsigntemplateformfieldRequestCompound) SetEEzsigntemplateformfieldPositioning(v FieldEEzsigntemplateformfieldPositioning)`

SetEEzsigntemplateformfieldPositioning sets EEzsigntemplateformfieldPositioning field to given value.

### HasEEzsigntemplateformfieldPositioning

`func (o *EzsigntemplateformfieldRequestCompound) HasEEzsigntemplateformfieldPositioning() bool`

HasEEzsigntemplateformfieldPositioning returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplatedocumentpagePagenumber() int32`

GetIEzsigntemplatedocumentpagePagenumber returns the IEzsigntemplatedocumentpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagePagenumberOk

`func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplatedocumentpagePagenumberOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagePagenumberOk returns a tuple with the IEzsigntemplatedocumentpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplatedocumentpagePagenumber(v int32)`

SetIEzsigntemplatedocumentpagePagenumber sets IEzsigntemplatedocumentpagePagenumber field to given value.


### GetSEzsigntemplateformfieldLabel

`func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldLabel() string`

GetSEzsigntemplateformfieldLabel returns the SEzsigntemplateformfieldLabel field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldLabelOk

`func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldLabelOk() (*string, bool)`

GetSEzsigntemplateformfieldLabelOk returns a tuple with the SEzsigntemplateformfieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldLabel

`func (o *EzsigntemplateformfieldRequestCompound) SetSEzsigntemplateformfieldLabel(v string)`

SetSEzsigntemplateformfieldLabel sets SEzsigntemplateformfieldLabel field to given value.


### GetSEzsigntemplateformfieldValue

`func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldValue() string`

GetSEzsigntemplateformfieldValue returns the SEzsigntemplateformfieldValue field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldValueOk

`func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldValueOk() (*string, bool)`

GetSEzsigntemplateformfieldValueOk returns a tuple with the SEzsigntemplateformfieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldValue

`func (o *EzsigntemplateformfieldRequestCompound) SetSEzsigntemplateformfieldValue(v string)`

SetSEzsigntemplateformfieldValue sets SEzsigntemplateformfieldValue field to given value.

### HasSEzsigntemplateformfieldValue

`func (o *EzsigntemplateformfieldRequestCompound) HasSEzsigntemplateformfieldValue() bool`

HasSEzsigntemplateformfieldValue returns a boolean if a field has been set.

### GetIEzsigntemplateformfieldX

`func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldX() int32`

GetIEzsigntemplateformfieldX returns the IEzsigntemplateformfieldX field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldXOk

`func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldXOk() (*int32, bool)`

GetIEzsigntemplateformfieldXOk returns a tuple with the IEzsigntemplateformfieldX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldX

`func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldX(v int32)`

SetIEzsigntemplateformfieldX sets IEzsigntemplateformfieldX field to given value.

### HasIEzsigntemplateformfieldX

`func (o *EzsigntemplateformfieldRequestCompound) HasIEzsigntemplateformfieldX() bool`

HasIEzsigntemplateformfieldX returns a boolean if a field has been set.

### GetIEzsigntemplateformfieldY

`func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldY() int32`

GetIEzsigntemplateformfieldY returns the IEzsigntemplateformfieldY field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldYOk

`func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldYOk() (*int32, bool)`

GetIEzsigntemplateformfieldYOk returns a tuple with the IEzsigntemplateformfieldY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldY

`func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldY(v int32)`

SetIEzsigntemplateformfieldY sets IEzsigntemplateformfieldY field to given value.

### HasIEzsigntemplateformfieldY

`func (o *EzsigntemplateformfieldRequestCompound) HasIEzsigntemplateformfieldY() bool`

HasIEzsigntemplateformfieldY returns a boolean if a field has been set.

### GetIEzsigntemplateformfieldWidth

`func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldWidth() int32`

GetIEzsigntemplateformfieldWidth returns the IEzsigntemplateformfieldWidth field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldWidthOk

`func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldWidthOk() (*int32, bool)`

GetIEzsigntemplateformfieldWidthOk returns a tuple with the IEzsigntemplateformfieldWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldWidth

`func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldWidth(v int32)`

SetIEzsigntemplateformfieldWidth sets IEzsigntemplateformfieldWidth field to given value.


### GetIEzsigntemplateformfieldHeight

`func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldHeight() int32`

GetIEzsigntemplateformfieldHeight returns the IEzsigntemplateformfieldHeight field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldHeightOk

`func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldHeightOk() (*int32, bool)`

GetIEzsigntemplateformfieldHeightOk returns a tuple with the IEzsigntemplateformfieldHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldHeight

`func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldHeight(v int32)`

SetIEzsigntemplateformfieldHeight sets IEzsigntemplateformfieldHeight field to given value.


### GetBEzsigntemplateformfieldAutocomplete

`func (o *EzsigntemplateformfieldRequestCompound) GetBEzsigntemplateformfieldAutocomplete() bool`

GetBEzsigntemplateformfieldAutocomplete returns the BEzsigntemplateformfieldAutocomplete field if non-nil, zero value otherwise.

### GetBEzsigntemplateformfieldAutocompleteOk

`func (o *EzsigntemplateformfieldRequestCompound) GetBEzsigntemplateformfieldAutocompleteOk() (*bool, bool)`

GetBEzsigntemplateformfieldAutocompleteOk returns a tuple with the BEzsigntemplateformfieldAutocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateformfieldAutocomplete

`func (o *EzsigntemplateformfieldRequestCompound) SetBEzsigntemplateformfieldAutocomplete(v bool)`

SetBEzsigntemplateformfieldAutocomplete sets BEzsigntemplateformfieldAutocomplete field to given value.

### HasBEzsigntemplateformfieldAutocomplete

`func (o *EzsigntemplateformfieldRequestCompound) HasBEzsigntemplateformfieldAutocomplete() bool`

HasBEzsigntemplateformfieldAutocomplete returns a boolean if a field has been set.

### GetBEzsigntemplateformfieldSelected

`func (o *EzsigntemplateformfieldRequestCompound) GetBEzsigntemplateformfieldSelected() bool`

GetBEzsigntemplateformfieldSelected returns the BEzsigntemplateformfieldSelected field if non-nil, zero value otherwise.

### GetBEzsigntemplateformfieldSelectedOk

`func (o *EzsigntemplateformfieldRequestCompound) GetBEzsigntemplateformfieldSelectedOk() (*bool, bool)`

GetBEzsigntemplateformfieldSelectedOk returns a tuple with the BEzsigntemplateformfieldSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateformfieldSelected

`func (o *EzsigntemplateformfieldRequestCompound) SetBEzsigntemplateformfieldSelected(v bool)`

SetBEzsigntemplateformfieldSelected sets BEzsigntemplateformfieldSelected field to given value.

### HasBEzsigntemplateformfieldSelected

`func (o *EzsigntemplateformfieldRequestCompound) HasBEzsigntemplateformfieldSelected() bool`

HasBEzsigntemplateformfieldSelected returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldDependencyrequirement

`func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldDependencyrequirement() FieldEEzsigntemplateformfieldDependencyrequirement`

GetEEzsigntemplateformfieldDependencyrequirement returns the EEzsigntemplateformfieldDependencyrequirement field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldDependencyrequirementOk

`func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldDependencyrequirementOk() (*FieldEEzsigntemplateformfieldDependencyrequirement, bool)`

GetEEzsigntemplateformfieldDependencyrequirementOk returns a tuple with the EEzsigntemplateformfieldDependencyrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldDependencyrequirement

`func (o *EzsigntemplateformfieldRequestCompound) SetEEzsigntemplateformfieldDependencyrequirement(v FieldEEzsigntemplateformfieldDependencyrequirement)`

SetEEzsigntemplateformfieldDependencyrequirement sets EEzsigntemplateformfieldDependencyrequirement field to given value.

### HasEEzsigntemplateformfieldDependencyrequirement

`func (o *EzsigntemplateformfieldRequestCompound) HasEEzsigntemplateformfieldDependencyrequirement() bool`

HasEEzsigntemplateformfieldDependencyrequirement returns a boolean if a field has been set.

### GetSEzsigntemplateformfieldPositioningpattern

`func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldPositioningpattern() string`

GetSEzsigntemplateformfieldPositioningpattern returns the SEzsigntemplateformfieldPositioningpattern field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldPositioningpatternOk

`func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldPositioningpatternOk() (*string, bool)`

GetSEzsigntemplateformfieldPositioningpatternOk returns a tuple with the SEzsigntemplateformfieldPositioningpattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldPositioningpattern

`func (o *EzsigntemplateformfieldRequestCompound) SetSEzsigntemplateformfieldPositioningpattern(v string)`

SetSEzsigntemplateformfieldPositioningpattern sets SEzsigntemplateformfieldPositioningpattern field to given value.

### HasSEzsigntemplateformfieldPositioningpattern

`func (o *EzsigntemplateformfieldRequestCompound) HasSEzsigntemplateformfieldPositioningpattern() bool`

HasSEzsigntemplateformfieldPositioningpattern returns a boolean if a field has been set.

### GetIEzsigntemplateformfieldPositioningoffsetx

`func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldPositioningoffsetx() int32`

GetIEzsigntemplateformfieldPositioningoffsetx returns the IEzsigntemplateformfieldPositioningoffsetx field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldPositioningoffsetxOk

`func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldPositioningoffsetxOk() (*int32, bool)`

GetIEzsigntemplateformfieldPositioningoffsetxOk returns a tuple with the IEzsigntemplateformfieldPositioningoffsetx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldPositioningoffsetx

`func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldPositioningoffsetx(v int32)`

SetIEzsigntemplateformfieldPositioningoffsetx sets IEzsigntemplateformfieldPositioningoffsetx field to given value.

### HasIEzsigntemplateformfieldPositioningoffsetx

`func (o *EzsigntemplateformfieldRequestCompound) HasIEzsigntemplateformfieldPositioningoffsetx() bool`

HasIEzsigntemplateformfieldPositioningoffsetx returns a boolean if a field has been set.

### GetIEzsigntemplateformfieldPositioningoffsety

`func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldPositioningoffsety() int32`

GetIEzsigntemplateformfieldPositioningoffsety returns the IEzsigntemplateformfieldPositioningoffsety field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldPositioningoffsetyOk

`func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldPositioningoffsetyOk() (*int32, bool)`

GetIEzsigntemplateformfieldPositioningoffsetyOk returns a tuple with the IEzsigntemplateformfieldPositioningoffsety field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldPositioningoffsety

`func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldPositioningoffsety(v int32)`

SetIEzsigntemplateformfieldPositioningoffsety sets IEzsigntemplateformfieldPositioningoffsety field to given value.

### HasIEzsigntemplateformfieldPositioningoffsety

`func (o *EzsigntemplateformfieldRequestCompound) HasIEzsigntemplateformfieldPositioningoffsety() bool`

HasIEzsigntemplateformfieldPositioningoffsety returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldPositioningoccurence

`func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldPositioningoccurence() FieldEEzsigntemplateformfieldPositioningoccurence`

GetEEzsigntemplateformfieldPositioningoccurence returns the EEzsigntemplateformfieldPositioningoccurence field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldPositioningoccurenceOk

`func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldPositioningoccurenceOk() (*FieldEEzsigntemplateformfieldPositioningoccurence, bool)`

GetEEzsigntemplateformfieldPositioningoccurenceOk returns a tuple with the EEzsigntemplateformfieldPositioningoccurence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldPositioningoccurence

`func (o *EzsigntemplateformfieldRequestCompound) SetEEzsigntemplateformfieldPositioningoccurence(v FieldEEzsigntemplateformfieldPositioningoccurence)`

SetEEzsigntemplateformfieldPositioningoccurence sets EEzsigntemplateformfieldPositioningoccurence field to given value.

### HasEEzsigntemplateformfieldPositioningoccurence

`func (o *EzsigntemplateformfieldRequestCompound) HasEEzsigntemplateformfieldPositioningoccurence() bool`

HasEEzsigntemplateformfieldPositioningoccurence returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldHorizontalalignment

`func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldHorizontalalignment() EnumHorizontalalignment`

GetEEzsigntemplateformfieldHorizontalalignment returns the EEzsigntemplateformfieldHorizontalalignment field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldHorizontalalignmentOk

`func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldHorizontalalignmentOk() (*EnumHorizontalalignment, bool)`

GetEEzsigntemplateformfieldHorizontalalignmentOk returns a tuple with the EEzsigntemplateformfieldHorizontalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldHorizontalalignment

`func (o *EzsigntemplateformfieldRequestCompound) SetEEzsigntemplateformfieldHorizontalalignment(v EnumHorizontalalignment)`

SetEEzsigntemplateformfieldHorizontalalignment sets EEzsigntemplateformfieldHorizontalalignment field to given value.

### HasEEzsigntemplateformfieldHorizontalalignment

`func (o *EzsigntemplateformfieldRequestCompound) HasEEzsigntemplateformfieldHorizontalalignment() bool`

HasEEzsigntemplateformfieldHorizontalalignment returns a boolean if a field has been set.

### GetObjTextstylestatic

`func (o *EzsigntemplateformfieldRequestCompound) GetObjTextstylestatic() TextstylestaticRequestCompound`

GetObjTextstylestatic returns the ObjTextstylestatic field if non-nil, zero value otherwise.

### GetObjTextstylestaticOk

`func (o *EzsigntemplateformfieldRequestCompound) GetObjTextstylestaticOk() (*TextstylestaticRequestCompound, bool)`

GetObjTextstylestaticOk returns a tuple with the ObjTextstylestatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjTextstylestatic

`func (o *EzsigntemplateformfieldRequestCompound) SetObjTextstylestatic(v TextstylestaticRequestCompound)`

SetObjTextstylestatic sets ObjTextstylestatic field to given value.

### HasObjTextstylestatic

`func (o *EzsigntemplateformfieldRequestCompound) HasObjTextstylestatic() bool`

HasObjTextstylestatic returns a boolean if a field has been set.

### GetAObjEzsigntemplateelementdependency

`func (o *EzsigntemplateformfieldRequestCompound) GetAObjEzsigntemplateelementdependency() []EzsigntemplateelementdependencyRequestCompound`

GetAObjEzsigntemplateelementdependency returns the AObjEzsigntemplateelementdependency field if non-nil, zero value otherwise.

### GetAObjEzsigntemplateelementdependencyOk

`func (o *EzsigntemplateformfieldRequestCompound) GetAObjEzsigntemplateelementdependencyOk() (*[]EzsigntemplateelementdependencyRequestCompound, bool)`

GetAObjEzsigntemplateelementdependencyOk returns a tuple with the AObjEzsigntemplateelementdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplateelementdependency

`func (o *EzsigntemplateformfieldRequestCompound) SetAObjEzsigntemplateelementdependency(v []EzsigntemplateelementdependencyRequestCompound)`

SetAObjEzsigntemplateelementdependency sets AObjEzsigntemplateelementdependency field to given value.

### HasAObjEzsigntemplateelementdependency

`func (o *EzsigntemplateformfieldRequestCompound) HasAObjEzsigntemplateelementdependency() bool`

HasAObjEzsigntemplateelementdependency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


