# EzsigntemplateannotationRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateannotationID** | Pointer to **int32** | The unique ID of the Ezsigntemplateannotation | [optional] 
**FkiEzsigntemplatedocumentID** | **int32** | The unique ID of the Ezsigntemplatedocument | 
**EEzsigntemplateannotationHorizontalalignment** | [**FieldEEzsigntemplateannotationHorizontalalignment**](FieldEEzsigntemplateannotationHorizontalalignment.md) |  | 
**EEzsigntemplateannotationVerticalalignment** | [**FieldEEzsigntemplateannotationVerticalalignment**](FieldEEzsigntemplateannotationVerticalalignment.md) |  | 
**EEzsigntemplateannotationType** | [**FieldEEzsigntemplateannotationType**](FieldEEzsigntemplateannotationType.md) |  | 
**IEzsigntemplateannotationX** | **int32** | The x of the Ezsigntemplateannotation | 
**IEzsigntemplateannotationY** | **int32** | The y of the Ezsigntemplateannotation | 
**IEzsigntemplateannotationWidth** | **int32** | The width of the Ezsigntemplateannotation | 
**IEzsigntemplateannotationHeight** | **int32** | The height of the Ezsigntemplateannotation | 
**IEzsigntemplatedocumentpagePagenumber** | **int32** | The page number in the Ezsigntemplatedocument | 
**SEzsigntemplateannotationDescription** | **string** | The description of the Ezsigntemplateannotation | 
**SEzsigntemplateannotationDefaulttext** | **string** | The defaulttext of the Ezsigntemplateannotation | 
**SEzsigntemplateannotationDropdownvalues** | **string** | The ndropdownvalues of the Ezsigntemplateannotation | 
**ObjTextstylestatic** | Pointer to [**TextstylestaticRequestCompound**](TextstylestaticRequestCompound.md) |  | [optional] 

## Methods

### NewEzsigntemplateannotationRequestCompound

`func NewEzsigntemplateannotationRequestCompound(fkiEzsigntemplatedocumentID int32, eEzsigntemplateannotationHorizontalalignment FieldEEzsigntemplateannotationHorizontalalignment, eEzsigntemplateannotationVerticalalignment FieldEEzsigntemplateannotationVerticalalignment, eEzsigntemplateannotationType FieldEEzsigntemplateannotationType, iEzsigntemplateannotationX int32, iEzsigntemplateannotationY int32, iEzsigntemplateannotationWidth int32, iEzsigntemplateannotationHeight int32, iEzsigntemplatedocumentpagePagenumber int32, sEzsigntemplateannotationDescription string, sEzsigntemplateannotationDefaulttext string, sEzsigntemplateannotationDropdownvalues string, ) *EzsigntemplateannotationRequestCompound`

NewEzsigntemplateannotationRequestCompound instantiates a new EzsigntemplateannotationRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateannotationRequestCompoundWithDefaults

`func NewEzsigntemplateannotationRequestCompoundWithDefaults() *EzsigntemplateannotationRequestCompound`

NewEzsigntemplateannotationRequestCompoundWithDefaults instantiates a new EzsigntemplateannotationRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateannotationID

`func (o *EzsigntemplateannotationRequestCompound) GetPkiEzsigntemplateannotationID() int32`

GetPkiEzsigntemplateannotationID returns the PkiEzsigntemplateannotationID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateannotationIDOk

`func (o *EzsigntemplateannotationRequestCompound) GetPkiEzsigntemplateannotationIDOk() (*int32, bool)`

GetPkiEzsigntemplateannotationIDOk returns a tuple with the PkiEzsigntemplateannotationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateannotationID

`func (o *EzsigntemplateannotationRequestCompound) SetPkiEzsigntemplateannotationID(v int32)`

SetPkiEzsigntemplateannotationID sets PkiEzsigntemplateannotationID field to given value.

### HasPkiEzsigntemplateannotationID

`func (o *EzsigntemplateannotationRequestCompound) HasPkiEzsigntemplateannotationID() bool`

HasPkiEzsigntemplateannotationID returns a boolean if a field has been set.

### GetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateannotationRequestCompound) GetFkiEzsigntemplatedocumentID() int32`

GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatedocumentIDOk

`func (o *EzsigntemplateannotationRequestCompound) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool)`

GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatedocumentID

`func (o *EzsigntemplateannotationRequestCompound) SetFkiEzsigntemplatedocumentID(v int32)`

SetFkiEzsigntemplatedocumentID sets FkiEzsigntemplatedocumentID field to given value.


### GetEEzsigntemplateannotationHorizontalalignment

`func (o *EzsigntemplateannotationRequestCompound) GetEEzsigntemplateannotationHorizontalalignment() FieldEEzsigntemplateannotationHorizontalalignment`

GetEEzsigntemplateannotationHorizontalalignment returns the EEzsigntemplateannotationHorizontalalignment field if non-nil, zero value otherwise.

### GetEEzsigntemplateannotationHorizontalalignmentOk

`func (o *EzsigntemplateannotationRequestCompound) GetEEzsigntemplateannotationHorizontalalignmentOk() (*FieldEEzsigntemplateannotationHorizontalalignment, bool)`

GetEEzsigntemplateannotationHorizontalalignmentOk returns a tuple with the EEzsigntemplateannotationHorizontalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateannotationHorizontalalignment

`func (o *EzsigntemplateannotationRequestCompound) SetEEzsigntemplateannotationHorizontalalignment(v FieldEEzsigntemplateannotationHorizontalalignment)`

SetEEzsigntemplateannotationHorizontalalignment sets EEzsigntemplateannotationHorizontalalignment field to given value.


### GetEEzsigntemplateannotationVerticalalignment

`func (o *EzsigntemplateannotationRequestCompound) GetEEzsigntemplateannotationVerticalalignment() FieldEEzsigntemplateannotationVerticalalignment`

GetEEzsigntemplateannotationVerticalalignment returns the EEzsigntemplateannotationVerticalalignment field if non-nil, zero value otherwise.

### GetEEzsigntemplateannotationVerticalalignmentOk

`func (o *EzsigntemplateannotationRequestCompound) GetEEzsigntemplateannotationVerticalalignmentOk() (*FieldEEzsigntemplateannotationVerticalalignment, bool)`

GetEEzsigntemplateannotationVerticalalignmentOk returns a tuple with the EEzsigntemplateannotationVerticalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateannotationVerticalalignment

`func (o *EzsigntemplateannotationRequestCompound) SetEEzsigntemplateannotationVerticalalignment(v FieldEEzsigntemplateannotationVerticalalignment)`

SetEEzsigntemplateannotationVerticalalignment sets EEzsigntemplateannotationVerticalalignment field to given value.


### GetEEzsigntemplateannotationType

`func (o *EzsigntemplateannotationRequestCompound) GetEEzsigntemplateannotationType() FieldEEzsigntemplateannotationType`

GetEEzsigntemplateannotationType returns the EEzsigntemplateannotationType field if non-nil, zero value otherwise.

### GetEEzsigntemplateannotationTypeOk

`func (o *EzsigntemplateannotationRequestCompound) GetEEzsigntemplateannotationTypeOk() (*FieldEEzsigntemplateannotationType, bool)`

GetEEzsigntemplateannotationTypeOk returns a tuple with the EEzsigntemplateannotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateannotationType

`func (o *EzsigntemplateannotationRequestCompound) SetEEzsigntemplateannotationType(v FieldEEzsigntemplateannotationType)`

SetEEzsigntemplateannotationType sets EEzsigntemplateannotationType field to given value.


### GetIEzsigntemplateannotationX

`func (o *EzsigntemplateannotationRequestCompound) GetIEzsigntemplateannotationX() int32`

GetIEzsigntemplateannotationX returns the IEzsigntemplateannotationX field if non-nil, zero value otherwise.

### GetIEzsigntemplateannotationXOk

`func (o *EzsigntemplateannotationRequestCompound) GetIEzsigntemplateannotationXOk() (*int32, bool)`

GetIEzsigntemplateannotationXOk returns a tuple with the IEzsigntemplateannotationX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateannotationX

`func (o *EzsigntemplateannotationRequestCompound) SetIEzsigntemplateannotationX(v int32)`

SetIEzsigntemplateannotationX sets IEzsigntemplateannotationX field to given value.


### GetIEzsigntemplateannotationY

`func (o *EzsigntemplateannotationRequestCompound) GetIEzsigntemplateannotationY() int32`

GetIEzsigntemplateannotationY returns the IEzsigntemplateannotationY field if non-nil, zero value otherwise.

### GetIEzsigntemplateannotationYOk

`func (o *EzsigntemplateannotationRequestCompound) GetIEzsigntemplateannotationYOk() (*int32, bool)`

GetIEzsigntemplateannotationYOk returns a tuple with the IEzsigntemplateannotationY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateannotationY

`func (o *EzsigntemplateannotationRequestCompound) SetIEzsigntemplateannotationY(v int32)`

SetIEzsigntemplateannotationY sets IEzsigntemplateannotationY field to given value.


### GetIEzsigntemplateannotationWidth

`func (o *EzsigntemplateannotationRequestCompound) GetIEzsigntemplateannotationWidth() int32`

GetIEzsigntemplateannotationWidth returns the IEzsigntemplateannotationWidth field if non-nil, zero value otherwise.

### GetIEzsigntemplateannotationWidthOk

`func (o *EzsigntemplateannotationRequestCompound) GetIEzsigntemplateannotationWidthOk() (*int32, bool)`

GetIEzsigntemplateannotationWidthOk returns a tuple with the IEzsigntemplateannotationWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateannotationWidth

`func (o *EzsigntemplateannotationRequestCompound) SetIEzsigntemplateannotationWidth(v int32)`

SetIEzsigntemplateannotationWidth sets IEzsigntemplateannotationWidth field to given value.


### GetIEzsigntemplateannotationHeight

`func (o *EzsigntemplateannotationRequestCompound) GetIEzsigntemplateannotationHeight() int32`

GetIEzsigntemplateannotationHeight returns the IEzsigntemplateannotationHeight field if non-nil, zero value otherwise.

### GetIEzsigntemplateannotationHeightOk

`func (o *EzsigntemplateannotationRequestCompound) GetIEzsigntemplateannotationHeightOk() (*int32, bool)`

GetIEzsigntemplateannotationHeightOk returns a tuple with the IEzsigntemplateannotationHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateannotationHeight

`func (o *EzsigntemplateannotationRequestCompound) SetIEzsigntemplateannotationHeight(v int32)`

SetIEzsigntemplateannotationHeight sets IEzsigntemplateannotationHeight field to given value.


### GetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplateannotationRequestCompound) GetIEzsigntemplatedocumentpagePagenumber() int32`

GetIEzsigntemplatedocumentpagePagenumber returns the IEzsigntemplatedocumentpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagePagenumberOk

`func (o *EzsigntemplateannotationRequestCompound) GetIEzsigntemplatedocumentpagePagenumberOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagePagenumberOk returns a tuple with the IEzsigntemplatedocumentpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplateannotationRequestCompound) SetIEzsigntemplatedocumentpagePagenumber(v int32)`

SetIEzsigntemplatedocumentpagePagenumber sets IEzsigntemplatedocumentpagePagenumber field to given value.


### GetSEzsigntemplateannotationDescription

`func (o *EzsigntemplateannotationRequestCompound) GetSEzsigntemplateannotationDescription() string`

GetSEzsigntemplateannotationDescription returns the SEzsigntemplateannotationDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplateannotationDescriptionOk

`func (o *EzsigntemplateannotationRequestCompound) GetSEzsigntemplateannotationDescriptionOk() (*string, bool)`

GetSEzsigntemplateannotationDescriptionOk returns a tuple with the SEzsigntemplateannotationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateannotationDescription

`func (o *EzsigntemplateannotationRequestCompound) SetSEzsigntemplateannotationDescription(v string)`

SetSEzsigntemplateannotationDescription sets SEzsigntemplateannotationDescription field to given value.


### GetSEzsigntemplateannotationDefaulttext

`func (o *EzsigntemplateannotationRequestCompound) GetSEzsigntemplateannotationDefaulttext() string`

GetSEzsigntemplateannotationDefaulttext returns the SEzsigntemplateannotationDefaulttext field if non-nil, zero value otherwise.

### GetSEzsigntemplateannotationDefaulttextOk

`func (o *EzsigntemplateannotationRequestCompound) GetSEzsigntemplateannotationDefaulttextOk() (*string, bool)`

GetSEzsigntemplateannotationDefaulttextOk returns a tuple with the SEzsigntemplateannotationDefaulttext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateannotationDefaulttext

`func (o *EzsigntemplateannotationRequestCompound) SetSEzsigntemplateannotationDefaulttext(v string)`

SetSEzsigntemplateannotationDefaulttext sets SEzsigntemplateannotationDefaulttext field to given value.


### GetSEzsigntemplateannotationDropdownvalues

`func (o *EzsigntemplateannotationRequestCompound) GetSEzsigntemplateannotationDropdownvalues() string`

GetSEzsigntemplateannotationDropdownvalues returns the SEzsigntemplateannotationDropdownvalues field if non-nil, zero value otherwise.

### GetSEzsigntemplateannotationDropdownvaluesOk

`func (o *EzsigntemplateannotationRequestCompound) GetSEzsigntemplateannotationDropdownvaluesOk() (*string, bool)`

GetSEzsigntemplateannotationDropdownvaluesOk returns a tuple with the SEzsigntemplateannotationDropdownvalues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateannotationDropdownvalues

`func (o *EzsigntemplateannotationRequestCompound) SetSEzsigntemplateannotationDropdownvalues(v string)`

SetSEzsigntemplateannotationDropdownvalues sets SEzsigntemplateannotationDropdownvalues field to given value.


### GetObjTextstylestatic

`func (o *EzsigntemplateannotationRequestCompound) GetObjTextstylestatic() TextstylestaticRequestCompound`

GetObjTextstylestatic returns the ObjTextstylestatic field if non-nil, zero value otherwise.

### GetObjTextstylestaticOk

`func (o *EzsigntemplateannotationRequestCompound) GetObjTextstylestaticOk() (*TextstylestaticRequestCompound, bool)`

GetObjTextstylestaticOk returns a tuple with the ObjTextstylestatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjTextstylestatic

`func (o *EzsigntemplateannotationRequestCompound) SetObjTextstylestatic(v TextstylestaticRequestCompound)`

SetObjTextstylestatic sets ObjTextstylestatic field to given value.

### HasObjTextstylestatic

`func (o *EzsigntemplateannotationRequestCompound) HasObjTextstylestatic() bool`

HasObjTextstylestatic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


