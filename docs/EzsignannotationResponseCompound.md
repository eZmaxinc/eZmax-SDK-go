# EzsignannotationResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignannotationID** | **int32** | The unique ID of the Ezsignannotation | 
**FkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 
**EEzsignannotationHorizontalalignment** | Pointer to [**EnumHorizontalalignment**](EnumHorizontalalignment.md) |  | [optional] 
**EEzsignannotationVerticalalignment** | Pointer to [**EnumVerticalalignment**](EnumVerticalalignment.md) |  | [optional] 
**EEzsignannotationType** | [**FieldEEzsignannotationType**](FieldEEzsignannotationType.md) |  | 
**IEzsignannotationX** | **int32** | The X coordinate (Horizontal) where to put the Ezsignannotation on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignannotation 2 inches from the left border of the page, you would use \&quot;200\&quot; for the X coordinate. | 
**IEzsignannotationY** | **int32** | The Y coordinate (Vertical) where to put the Ezsignannotation on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignannotation 3 inches from the top border of the page, you would use \&quot;300\&quot; for the Y coordinate. | 
**IEzsignannotationWidth** | Pointer to **int32** | The Width of the Ezsignannotation.  Width is calculated at 100dpi (dot per inch). So for example, if you want to have the width of the Ezsignannotation to be 3 inches, you would use \&quot;300\&quot; for the Width. | [optional] 
**IEzsignannotationHeight** | Pointer to **int32** | The Height of the Ezsignannotation.  Height is calculated at 100dpi (dot per inch). So for example, if you want to have the height of the Ezsignannotation to be 2 inches, you would use \&quot;200\&quot; for the Height.  This can only be set if eEzsignannotationType is **StrikethroughBlock** or **Text** | [optional] 
**SEzsignannotationText** | Pointer to **string** | The Text of the Ezsignannotation | [optional] 
**IEzsignpagePagenumber** | **int32** | The page number in the Ezsigndocument | 
**ObjTextstylestatic** | Pointer to [**TextstylestaticResponseCompound**](TextstylestaticResponseCompound.md) |  | [optional] 

## Methods

### NewEzsignannotationResponseCompound

`func NewEzsignannotationResponseCompound(pkiEzsignannotationID int32, fkiEzsigndocumentID int32, eEzsignannotationType FieldEEzsignannotationType, iEzsignannotationX int32, iEzsignannotationY int32, iEzsignpagePagenumber int32, ) *EzsignannotationResponseCompound`

NewEzsignannotationResponseCompound instantiates a new EzsignannotationResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignannotationResponseCompoundWithDefaults

`func NewEzsignannotationResponseCompoundWithDefaults() *EzsignannotationResponseCompound`

NewEzsignannotationResponseCompoundWithDefaults instantiates a new EzsignannotationResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignannotationID

`func (o *EzsignannotationResponseCompound) GetPkiEzsignannotationID() int32`

GetPkiEzsignannotationID returns the PkiEzsignannotationID field if non-nil, zero value otherwise.

### GetPkiEzsignannotationIDOk

`func (o *EzsignannotationResponseCompound) GetPkiEzsignannotationIDOk() (*int32, bool)`

GetPkiEzsignannotationIDOk returns a tuple with the PkiEzsignannotationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignannotationID

`func (o *EzsignannotationResponseCompound) SetPkiEzsignannotationID(v int32)`

SetPkiEzsignannotationID sets PkiEzsignannotationID field to given value.


### GetFkiEzsigndocumentID

`func (o *EzsignannotationResponseCompound) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *EzsignannotationResponseCompound) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *EzsignannotationResponseCompound) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.


### GetEEzsignannotationHorizontalalignment

`func (o *EzsignannotationResponseCompound) GetEEzsignannotationHorizontalalignment() EnumHorizontalalignment`

GetEEzsignannotationHorizontalalignment returns the EEzsignannotationHorizontalalignment field if non-nil, zero value otherwise.

### GetEEzsignannotationHorizontalalignmentOk

`func (o *EzsignannotationResponseCompound) GetEEzsignannotationHorizontalalignmentOk() (*EnumHorizontalalignment, bool)`

GetEEzsignannotationHorizontalalignmentOk returns a tuple with the EEzsignannotationHorizontalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignannotationHorizontalalignment

`func (o *EzsignannotationResponseCompound) SetEEzsignannotationHorizontalalignment(v EnumHorizontalalignment)`

SetEEzsignannotationHorizontalalignment sets EEzsignannotationHorizontalalignment field to given value.

### HasEEzsignannotationHorizontalalignment

`func (o *EzsignannotationResponseCompound) HasEEzsignannotationHorizontalalignment() bool`

HasEEzsignannotationHorizontalalignment returns a boolean if a field has been set.

### GetEEzsignannotationVerticalalignment

`func (o *EzsignannotationResponseCompound) GetEEzsignannotationVerticalalignment() EnumVerticalalignment`

GetEEzsignannotationVerticalalignment returns the EEzsignannotationVerticalalignment field if non-nil, zero value otherwise.

### GetEEzsignannotationVerticalalignmentOk

`func (o *EzsignannotationResponseCompound) GetEEzsignannotationVerticalalignmentOk() (*EnumVerticalalignment, bool)`

GetEEzsignannotationVerticalalignmentOk returns a tuple with the EEzsignannotationVerticalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignannotationVerticalalignment

`func (o *EzsignannotationResponseCompound) SetEEzsignannotationVerticalalignment(v EnumVerticalalignment)`

SetEEzsignannotationVerticalalignment sets EEzsignannotationVerticalalignment field to given value.

### HasEEzsignannotationVerticalalignment

`func (o *EzsignannotationResponseCompound) HasEEzsignannotationVerticalalignment() bool`

HasEEzsignannotationVerticalalignment returns a boolean if a field has been set.

### GetEEzsignannotationType

`func (o *EzsignannotationResponseCompound) GetEEzsignannotationType() FieldEEzsignannotationType`

GetEEzsignannotationType returns the EEzsignannotationType field if non-nil, zero value otherwise.

### GetEEzsignannotationTypeOk

`func (o *EzsignannotationResponseCompound) GetEEzsignannotationTypeOk() (*FieldEEzsignannotationType, bool)`

GetEEzsignannotationTypeOk returns a tuple with the EEzsignannotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignannotationType

`func (o *EzsignannotationResponseCompound) SetEEzsignannotationType(v FieldEEzsignannotationType)`

SetEEzsignannotationType sets EEzsignannotationType field to given value.


### GetIEzsignannotationX

`func (o *EzsignannotationResponseCompound) GetIEzsignannotationX() int32`

GetIEzsignannotationX returns the IEzsignannotationX field if non-nil, zero value otherwise.

### GetIEzsignannotationXOk

`func (o *EzsignannotationResponseCompound) GetIEzsignannotationXOk() (*int32, bool)`

GetIEzsignannotationXOk returns a tuple with the IEzsignannotationX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignannotationX

`func (o *EzsignannotationResponseCompound) SetIEzsignannotationX(v int32)`

SetIEzsignannotationX sets IEzsignannotationX field to given value.


### GetIEzsignannotationY

`func (o *EzsignannotationResponseCompound) GetIEzsignannotationY() int32`

GetIEzsignannotationY returns the IEzsignannotationY field if non-nil, zero value otherwise.

### GetIEzsignannotationYOk

`func (o *EzsignannotationResponseCompound) GetIEzsignannotationYOk() (*int32, bool)`

GetIEzsignannotationYOk returns a tuple with the IEzsignannotationY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignannotationY

`func (o *EzsignannotationResponseCompound) SetIEzsignannotationY(v int32)`

SetIEzsignannotationY sets IEzsignannotationY field to given value.


### GetIEzsignannotationWidth

`func (o *EzsignannotationResponseCompound) GetIEzsignannotationWidth() int32`

GetIEzsignannotationWidth returns the IEzsignannotationWidth field if non-nil, zero value otherwise.

### GetIEzsignannotationWidthOk

`func (o *EzsignannotationResponseCompound) GetIEzsignannotationWidthOk() (*int32, bool)`

GetIEzsignannotationWidthOk returns a tuple with the IEzsignannotationWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignannotationWidth

`func (o *EzsignannotationResponseCompound) SetIEzsignannotationWidth(v int32)`

SetIEzsignannotationWidth sets IEzsignannotationWidth field to given value.

### HasIEzsignannotationWidth

`func (o *EzsignannotationResponseCompound) HasIEzsignannotationWidth() bool`

HasIEzsignannotationWidth returns a boolean if a field has been set.

### GetIEzsignannotationHeight

`func (o *EzsignannotationResponseCompound) GetIEzsignannotationHeight() int32`

GetIEzsignannotationHeight returns the IEzsignannotationHeight field if non-nil, zero value otherwise.

### GetIEzsignannotationHeightOk

`func (o *EzsignannotationResponseCompound) GetIEzsignannotationHeightOk() (*int32, bool)`

GetIEzsignannotationHeightOk returns a tuple with the IEzsignannotationHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignannotationHeight

`func (o *EzsignannotationResponseCompound) SetIEzsignannotationHeight(v int32)`

SetIEzsignannotationHeight sets IEzsignannotationHeight field to given value.

### HasIEzsignannotationHeight

`func (o *EzsignannotationResponseCompound) HasIEzsignannotationHeight() bool`

HasIEzsignannotationHeight returns a boolean if a field has been set.

### GetSEzsignannotationText

`func (o *EzsignannotationResponseCompound) GetSEzsignannotationText() string`

GetSEzsignannotationText returns the SEzsignannotationText field if non-nil, zero value otherwise.

### GetSEzsignannotationTextOk

`func (o *EzsignannotationResponseCompound) GetSEzsignannotationTextOk() (*string, bool)`

GetSEzsignannotationTextOk returns a tuple with the SEzsignannotationText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignannotationText

`func (o *EzsignannotationResponseCompound) SetSEzsignannotationText(v string)`

SetSEzsignannotationText sets SEzsignannotationText field to given value.

### HasSEzsignannotationText

`func (o *EzsignannotationResponseCompound) HasSEzsignannotationText() bool`

HasSEzsignannotationText returns a boolean if a field has been set.

### GetIEzsignpagePagenumber

`func (o *EzsignannotationResponseCompound) GetIEzsignpagePagenumber() int32`

GetIEzsignpagePagenumber returns the IEzsignpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsignpagePagenumberOk

`func (o *EzsignannotationResponseCompound) GetIEzsignpagePagenumberOk() (*int32, bool)`

GetIEzsignpagePagenumberOk returns a tuple with the IEzsignpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignpagePagenumber

`func (o *EzsignannotationResponseCompound) SetIEzsignpagePagenumber(v int32)`

SetIEzsignpagePagenumber sets IEzsignpagePagenumber field to given value.


### GetObjTextstylestatic

`func (o *EzsignannotationResponseCompound) GetObjTextstylestatic() TextstylestaticResponseCompound`

GetObjTextstylestatic returns the ObjTextstylestatic field if non-nil, zero value otherwise.

### GetObjTextstylestaticOk

`func (o *EzsignannotationResponseCompound) GetObjTextstylestaticOk() (*TextstylestaticResponseCompound, bool)`

GetObjTextstylestaticOk returns a tuple with the ObjTextstylestatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjTextstylestatic

`func (o *EzsignannotationResponseCompound) SetObjTextstylestatic(v TextstylestaticResponseCompound)`

SetObjTextstylestatic sets ObjTextstylestatic field to given value.

### HasObjTextstylestatic

`func (o *EzsignannotationResponseCompound) HasObjTextstylestatic() bool`

HasObjTextstylestatic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


