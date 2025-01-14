# EzsigntemplateformfieldResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateformfieldID** | **int32** | The unique ID of the Ezsigntemplateformfield | 
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
**ObjTextstylestatic** | Pointer to [**TextstylestaticResponseCompound**](TextstylestaticResponseCompound.md) |  | [optional] 

## Methods

### NewEzsigntemplateformfieldResponse

`func NewEzsigntemplateformfieldResponse(pkiEzsigntemplateformfieldID int32, iEzsigntemplatedocumentpagePagenumber int32, sEzsigntemplateformfieldLabel string, iEzsigntemplateformfieldWidth int32, iEzsigntemplateformfieldHeight int32, ) *EzsigntemplateformfieldResponse`

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


### GetEEzsigntemplateformfieldPositioning

`func (o *EzsigntemplateformfieldResponse) GetEEzsigntemplateformfieldPositioning() FieldEEzsigntemplateformfieldPositioning`

GetEEzsigntemplateformfieldPositioning returns the EEzsigntemplateformfieldPositioning field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldPositioningOk

`func (o *EzsigntemplateformfieldResponse) GetEEzsigntemplateformfieldPositioningOk() (*FieldEEzsigntemplateformfieldPositioning, bool)`

GetEEzsigntemplateformfieldPositioningOk returns a tuple with the EEzsigntemplateformfieldPositioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldPositioning

`func (o *EzsigntemplateformfieldResponse) SetEEzsigntemplateformfieldPositioning(v FieldEEzsigntemplateformfieldPositioning)`

SetEEzsigntemplateformfieldPositioning sets EEzsigntemplateformfieldPositioning field to given value.

### HasEEzsigntemplateformfieldPositioning

`func (o *EzsigntemplateformfieldResponse) HasEEzsigntemplateformfieldPositioning() bool`

HasEEzsigntemplateformfieldPositioning returns a boolean if a field has been set.

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

### HasIEzsigntemplateformfieldX

`func (o *EzsigntemplateformfieldResponse) HasIEzsigntemplateformfieldX() bool`

HasIEzsigntemplateformfieldX returns a boolean if a field has been set.

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

### HasIEzsigntemplateformfieldY

`func (o *EzsigntemplateformfieldResponse) HasIEzsigntemplateformfieldY() bool`

HasIEzsigntemplateformfieldY returns a boolean if a field has been set.

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


### GetBEzsigntemplateformfieldAutocomplete

`func (o *EzsigntemplateformfieldResponse) GetBEzsigntemplateformfieldAutocomplete() bool`

GetBEzsigntemplateformfieldAutocomplete returns the BEzsigntemplateformfieldAutocomplete field if non-nil, zero value otherwise.

### GetBEzsigntemplateformfieldAutocompleteOk

`func (o *EzsigntemplateformfieldResponse) GetBEzsigntemplateformfieldAutocompleteOk() (*bool, bool)`

GetBEzsigntemplateformfieldAutocompleteOk returns a tuple with the BEzsigntemplateformfieldAutocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplateformfieldAutocomplete

`func (o *EzsigntemplateformfieldResponse) SetBEzsigntemplateformfieldAutocomplete(v bool)`

SetBEzsigntemplateformfieldAutocomplete sets BEzsigntemplateformfieldAutocomplete field to given value.

### HasBEzsigntemplateformfieldAutocomplete

`func (o *EzsigntemplateformfieldResponse) HasBEzsigntemplateformfieldAutocomplete() bool`

HasBEzsigntemplateformfieldAutocomplete returns a boolean if a field has been set.

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

### GetEEzsigntemplateformfieldDependencyrequirement

`func (o *EzsigntemplateformfieldResponse) GetEEzsigntemplateformfieldDependencyrequirement() FieldEEzsigntemplateformfieldDependencyrequirement`

GetEEzsigntemplateformfieldDependencyrequirement returns the EEzsigntemplateformfieldDependencyrequirement field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldDependencyrequirementOk

`func (o *EzsigntemplateformfieldResponse) GetEEzsigntemplateformfieldDependencyrequirementOk() (*FieldEEzsigntemplateformfieldDependencyrequirement, bool)`

GetEEzsigntemplateformfieldDependencyrequirementOk returns a tuple with the EEzsigntemplateformfieldDependencyrequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldDependencyrequirement

`func (o *EzsigntemplateformfieldResponse) SetEEzsigntemplateformfieldDependencyrequirement(v FieldEEzsigntemplateformfieldDependencyrequirement)`

SetEEzsigntemplateformfieldDependencyrequirement sets EEzsigntemplateformfieldDependencyrequirement field to given value.

### HasEEzsigntemplateformfieldDependencyrequirement

`func (o *EzsigntemplateformfieldResponse) HasEEzsigntemplateformfieldDependencyrequirement() bool`

HasEEzsigntemplateformfieldDependencyrequirement returns a boolean if a field has been set.

### GetSEzsigntemplateformfieldPositioningpattern

`func (o *EzsigntemplateformfieldResponse) GetSEzsigntemplateformfieldPositioningpattern() string`

GetSEzsigntemplateformfieldPositioningpattern returns the SEzsigntemplateformfieldPositioningpattern field if non-nil, zero value otherwise.

### GetSEzsigntemplateformfieldPositioningpatternOk

`func (o *EzsigntemplateformfieldResponse) GetSEzsigntemplateformfieldPositioningpatternOk() (*string, bool)`

GetSEzsigntemplateformfieldPositioningpatternOk returns a tuple with the SEzsigntemplateformfieldPositioningpattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateformfieldPositioningpattern

`func (o *EzsigntemplateformfieldResponse) SetSEzsigntemplateformfieldPositioningpattern(v string)`

SetSEzsigntemplateformfieldPositioningpattern sets SEzsigntemplateformfieldPositioningpattern field to given value.

### HasSEzsigntemplateformfieldPositioningpattern

`func (o *EzsigntemplateformfieldResponse) HasSEzsigntemplateformfieldPositioningpattern() bool`

HasSEzsigntemplateformfieldPositioningpattern returns a boolean if a field has been set.

### GetIEzsigntemplateformfieldPositioningoffsetx

`func (o *EzsigntemplateformfieldResponse) GetIEzsigntemplateformfieldPositioningoffsetx() int32`

GetIEzsigntemplateformfieldPositioningoffsetx returns the IEzsigntemplateformfieldPositioningoffsetx field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldPositioningoffsetxOk

`func (o *EzsigntemplateformfieldResponse) GetIEzsigntemplateformfieldPositioningoffsetxOk() (*int32, bool)`

GetIEzsigntemplateformfieldPositioningoffsetxOk returns a tuple with the IEzsigntemplateformfieldPositioningoffsetx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldPositioningoffsetx

`func (o *EzsigntemplateformfieldResponse) SetIEzsigntemplateformfieldPositioningoffsetx(v int32)`

SetIEzsigntemplateformfieldPositioningoffsetx sets IEzsigntemplateformfieldPositioningoffsetx field to given value.

### HasIEzsigntemplateformfieldPositioningoffsetx

`func (o *EzsigntemplateformfieldResponse) HasIEzsigntemplateformfieldPositioningoffsetx() bool`

HasIEzsigntemplateformfieldPositioningoffsetx returns a boolean if a field has been set.

### GetIEzsigntemplateformfieldPositioningoffsety

`func (o *EzsigntemplateformfieldResponse) GetIEzsigntemplateformfieldPositioningoffsety() int32`

GetIEzsigntemplateformfieldPositioningoffsety returns the IEzsigntemplateformfieldPositioningoffsety field if non-nil, zero value otherwise.

### GetIEzsigntemplateformfieldPositioningoffsetyOk

`func (o *EzsigntemplateformfieldResponse) GetIEzsigntemplateformfieldPositioningoffsetyOk() (*int32, bool)`

GetIEzsigntemplateformfieldPositioningoffsetyOk returns a tuple with the IEzsigntemplateformfieldPositioningoffsety field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateformfieldPositioningoffsety

`func (o *EzsigntemplateformfieldResponse) SetIEzsigntemplateformfieldPositioningoffsety(v int32)`

SetIEzsigntemplateformfieldPositioningoffsety sets IEzsigntemplateformfieldPositioningoffsety field to given value.

### HasIEzsigntemplateformfieldPositioningoffsety

`func (o *EzsigntemplateformfieldResponse) HasIEzsigntemplateformfieldPositioningoffsety() bool`

HasIEzsigntemplateformfieldPositioningoffsety returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldPositioningoccurence

`func (o *EzsigntemplateformfieldResponse) GetEEzsigntemplateformfieldPositioningoccurence() FieldEEzsigntemplateformfieldPositioningoccurence`

GetEEzsigntemplateformfieldPositioningoccurence returns the EEzsigntemplateformfieldPositioningoccurence field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldPositioningoccurenceOk

`func (o *EzsigntemplateformfieldResponse) GetEEzsigntemplateformfieldPositioningoccurenceOk() (*FieldEEzsigntemplateformfieldPositioningoccurence, bool)`

GetEEzsigntemplateformfieldPositioningoccurenceOk returns a tuple with the EEzsigntemplateformfieldPositioningoccurence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldPositioningoccurence

`func (o *EzsigntemplateformfieldResponse) SetEEzsigntemplateformfieldPositioningoccurence(v FieldEEzsigntemplateformfieldPositioningoccurence)`

SetEEzsigntemplateformfieldPositioningoccurence sets EEzsigntemplateformfieldPositioningoccurence field to given value.

### HasEEzsigntemplateformfieldPositioningoccurence

`func (o *EzsigntemplateformfieldResponse) HasEEzsigntemplateformfieldPositioningoccurence() bool`

HasEEzsigntemplateformfieldPositioningoccurence returns a boolean if a field has been set.

### GetEEzsigntemplateformfieldHorizontalalignment

`func (o *EzsigntemplateformfieldResponse) GetEEzsigntemplateformfieldHorizontalalignment() EnumHorizontalalignment`

GetEEzsigntemplateformfieldHorizontalalignment returns the EEzsigntemplateformfieldHorizontalalignment field if non-nil, zero value otherwise.

### GetEEzsigntemplateformfieldHorizontalalignmentOk

`func (o *EzsigntemplateformfieldResponse) GetEEzsigntemplateformfieldHorizontalalignmentOk() (*EnumHorizontalalignment, bool)`

GetEEzsigntemplateformfieldHorizontalalignmentOk returns a tuple with the EEzsigntemplateformfieldHorizontalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateformfieldHorizontalalignment

`func (o *EzsigntemplateformfieldResponse) SetEEzsigntemplateformfieldHorizontalalignment(v EnumHorizontalalignment)`

SetEEzsigntemplateformfieldHorizontalalignment sets EEzsigntemplateformfieldHorizontalalignment field to given value.

### HasEEzsigntemplateformfieldHorizontalalignment

`func (o *EzsigntemplateformfieldResponse) HasEEzsigntemplateformfieldHorizontalalignment() bool`

HasEEzsigntemplateformfieldHorizontalalignment returns a boolean if a field has been set.

### GetObjTextstylestatic

`func (o *EzsigntemplateformfieldResponse) GetObjTextstylestatic() TextstylestaticResponseCompound`

GetObjTextstylestatic returns the ObjTextstylestatic field if non-nil, zero value otherwise.

### GetObjTextstylestaticOk

`func (o *EzsigntemplateformfieldResponse) GetObjTextstylestaticOk() (*TextstylestaticResponseCompound, bool)`

GetObjTextstylestaticOk returns a tuple with the ObjTextstylestatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjTextstylestatic

`func (o *EzsigntemplateformfieldResponse) SetObjTextstylestatic(v TextstylestaticResponseCompound)`

SetObjTextstylestatic sets ObjTextstylestatic field to given value.

### HasObjTextstylestatic

`func (o *EzsigntemplateformfieldResponse) HasObjTextstylestatic() bool`

HasObjTextstylestatic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


