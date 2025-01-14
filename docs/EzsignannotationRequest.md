# EzsignannotationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignannotationID** | Pointer to **int32** | The unique ID of the Ezsignannotation | [optional] 
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

## Methods

### NewEzsignannotationRequest

`func NewEzsignannotationRequest(fkiEzsigndocumentID int32, eEzsignannotationType FieldEEzsignannotationType, iEzsignannotationX int32, iEzsignannotationY int32, iEzsignpagePagenumber int32, ) *EzsignannotationRequest`

NewEzsignannotationRequest instantiates a new EzsignannotationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignannotationRequestWithDefaults

`func NewEzsignannotationRequestWithDefaults() *EzsignannotationRequest`

NewEzsignannotationRequestWithDefaults instantiates a new EzsignannotationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignannotationID

`func (o *EzsignannotationRequest) GetPkiEzsignannotationID() int32`

GetPkiEzsignannotationID returns the PkiEzsignannotationID field if non-nil, zero value otherwise.

### GetPkiEzsignannotationIDOk

`func (o *EzsignannotationRequest) GetPkiEzsignannotationIDOk() (*int32, bool)`

GetPkiEzsignannotationIDOk returns a tuple with the PkiEzsignannotationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignannotationID

`func (o *EzsignannotationRequest) SetPkiEzsignannotationID(v int32)`

SetPkiEzsignannotationID sets PkiEzsignannotationID field to given value.

### HasPkiEzsignannotationID

`func (o *EzsignannotationRequest) HasPkiEzsignannotationID() bool`

HasPkiEzsignannotationID returns a boolean if a field has been set.

### GetFkiEzsigndocumentID

`func (o *EzsignannotationRequest) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *EzsignannotationRequest) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *EzsignannotationRequest) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.


### GetEEzsignannotationHorizontalalignment

`func (o *EzsignannotationRequest) GetEEzsignannotationHorizontalalignment() EnumHorizontalalignment`

GetEEzsignannotationHorizontalalignment returns the EEzsignannotationHorizontalalignment field if non-nil, zero value otherwise.

### GetEEzsignannotationHorizontalalignmentOk

`func (o *EzsignannotationRequest) GetEEzsignannotationHorizontalalignmentOk() (*EnumHorizontalalignment, bool)`

GetEEzsignannotationHorizontalalignmentOk returns a tuple with the EEzsignannotationHorizontalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignannotationHorizontalalignment

`func (o *EzsignannotationRequest) SetEEzsignannotationHorizontalalignment(v EnumHorizontalalignment)`

SetEEzsignannotationHorizontalalignment sets EEzsignannotationHorizontalalignment field to given value.

### HasEEzsignannotationHorizontalalignment

`func (o *EzsignannotationRequest) HasEEzsignannotationHorizontalalignment() bool`

HasEEzsignannotationHorizontalalignment returns a boolean if a field has been set.

### GetEEzsignannotationVerticalalignment

`func (o *EzsignannotationRequest) GetEEzsignannotationVerticalalignment() EnumVerticalalignment`

GetEEzsignannotationVerticalalignment returns the EEzsignannotationVerticalalignment field if non-nil, zero value otherwise.

### GetEEzsignannotationVerticalalignmentOk

`func (o *EzsignannotationRequest) GetEEzsignannotationVerticalalignmentOk() (*EnumVerticalalignment, bool)`

GetEEzsignannotationVerticalalignmentOk returns a tuple with the EEzsignannotationVerticalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignannotationVerticalalignment

`func (o *EzsignannotationRequest) SetEEzsignannotationVerticalalignment(v EnumVerticalalignment)`

SetEEzsignannotationVerticalalignment sets EEzsignannotationVerticalalignment field to given value.

### HasEEzsignannotationVerticalalignment

`func (o *EzsignannotationRequest) HasEEzsignannotationVerticalalignment() bool`

HasEEzsignannotationVerticalalignment returns a boolean if a field has been set.

### GetEEzsignannotationType

`func (o *EzsignannotationRequest) GetEEzsignannotationType() FieldEEzsignannotationType`

GetEEzsignannotationType returns the EEzsignannotationType field if non-nil, zero value otherwise.

### GetEEzsignannotationTypeOk

`func (o *EzsignannotationRequest) GetEEzsignannotationTypeOk() (*FieldEEzsignannotationType, bool)`

GetEEzsignannotationTypeOk returns a tuple with the EEzsignannotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignannotationType

`func (o *EzsignannotationRequest) SetEEzsignannotationType(v FieldEEzsignannotationType)`

SetEEzsignannotationType sets EEzsignannotationType field to given value.


### GetIEzsignannotationX

`func (o *EzsignannotationRequest) GetIEzsignannotationX() int32`

GetIEzsignannotationX returns the IEzsignannotationX field if non-nil, zero value otherwise.

### GetIEzsignannotationXOk

`func (o *EzsignannotationRequest) GetIEzsignannotationXOk() (*int32, bool)`

GetIEzsignannotationXOk returns a tuple with the IEzsignannotationX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignannotationX

`func (o *EzsignannotationRequest) SetIEzsignannotationX(v int32)`

SetIEzsignannotationX sets IEzsignannotationX field to given value.


### GetIEzsignannotationY

`func (o *EzsignannotationRequest) GetIEzsignannotationY() int32`

GetIEzsignannotationY returns the IEzsignannotationY field if non-nil, zero value otherwise.

### GetIEzsignannotationYOk

`func (o *EzsignannotationRequest) GetIEzsignannotationYOk() (*int32, bool)`

GetIEzsignannotationYOk returns a tuple with the IEzsignannotationY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignannotationY

`func (o *EzsignannotationRequest) SetIEzsignannotationY(v int32)`

SetIEzsignannotationY sets IEzsignannotationY field to given value.


### GetIEzsignannotationWidth

`func (o *EzsignannotationRequest) GetIEzsignannotationWidth() int32`

GetIEzsignannotationWidth returns the IEzsignannotationWidth field if non-nil, zero value otherwise.

### GetIEzsignannotationWidthOk

`func (o *EzsignannotationRequest) GetIEzsignannotationWidthOk() (*int32, bool)`

GetIEzsignannotationWidthOk returns a tuple with the IEzsignannotationWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignannotationWidth

`func (o *EzsignannotationRequest) SetIEzsignannotationWidth(v int32)`

SetIEzsignannotationWidth sets IEzsignannotationWidth field to given value.

### HasIEzsignannotationWidth

`func (o *EzsignannotationRequest) HasIEzsignannotationWidth() bool`

HasIEzsignannotationWidth returns a boolean if a field has been set.

### GetIEzsignannotationHeight

`func (o *EzsignannotationRequest) GetIEzsignannotationHeight() int32`

GetIEzsignannotationHeight returns the IEzsignannotationHeight field if non-nil, zero value otherwise.

### GetIEzsignannotationHeightOk

`func (o *EzsignannotationRequest) GetIEzsignannotationHeightOk() (*int32, bool)`

GetIEzsignannotationHeightOk returns a tuple with the IEzsignannotationHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignannotationHeight

`func (o *EzsignannotationRequest) SetIEzsignannotationHeight(v int32)`

SetIEzsignannotationHeight sets IEzsignannotationHeight field to given value.

### HasIEzsignannotationHeight

`func (o *EzsignannotationRequest) HasIEzsignannotationHeight() bool`

HasIEzsignannotationHeight returns a boolean if a field has been set.

### GetSEzsignannotationText

`func (o *EzsignannotationRequest) GetSEzsignannotationText() string`

GetSEzsignannotationText returns the SEzsignannotationText field if non-nil, zero value otherwise.

### GetSEzsignannotationTextOk

`func (o *EzsignannotationRequest) GetSEzsignannotationTextOk() (*string, bool)`

GetSEzsignannotationTextOk returns a tuple with the SEzsignannotationText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignannotationText

`func (o *EzsignannotationRequest) SetSEzsignannotationText(v string)`

SetSEzsignannotationText sets SEzsignannotationText field to given value.

### HasSEzsignannotationText

`func (o *EzsignannotationRequest) HasSEzsignannotationText() bool`

HasSEzsignannotationText returns a boolean if a field has been set.

### GetIEzsignpagePagenumber

`func (o *EzsignannotationRequest) GetIEzsignpagePagenumber() int32`

GetIEzsignpagePagenumber returns the IEzsignpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsignpagePagenumberOk

`func (o *EzsignannotationRequest) GetIEzsignpagePagenumberOk() (*int32, bool)`

GetIEzsignpagePagenumberOk returns a tuple with the IEzsignpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignpagePagenumber

`func (o *EzsignannotationRequest) SetIEzsignpagePagenumber(v int32)`

SetIEzsignpagePagenumber sets IEzsignpagePagenumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


