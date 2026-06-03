# EzsigntemplateglobalannotationResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateglobalannotationID** | **int32** | The unique ID of the Ezsigntemplateglobalannotation | 
**FkiTextstylestaticID** | Pointer to **int32** | The unique ID of the Textstylestatic | [optional] 
**ObjTextstylestatic** | Pointer to [**TextstylestaticRequestCompound**](TextstylestaticRequestCompound.md) |  | [optional] 
**EEzsigntemplateglobalannotationHorizontalalignment** | [**FieldEEzsigntemplateglobalannotationHorizontalalignment**](FieldEEzsigntemplateglobalannotationHorizontalalignment.md) |  | 
**EEzsigntemplateglobalannotationVerticalalignment** | [**FieldEEzsigntemplateglobalannotationVerticalalignment**](FieldEEzsigntemplateglobalannotationVerticalalignment.md) |  | 
**EEzsigntemplateglobalannotationType** | [**FieldEEzsigntemplateglobalannotationType**](FieldEEzsigntemplateglobalannotationType.md) |  | 
**IEzsigntemplateglobalannotationX** | **int32** | The x of the Ezsigntemplateglobalannotation | 
**IEzsigntemplateglobalannotationY** | **int32** | The y of the Ezsigntemplateglobalannotation | 
**IEzsigntemplateglobalannotationWidth** | **int32** | The width of the Ezsigntemplateglobalannotation | 
**IEzsigntemplateglobalannotationHeight** | **int32** | The height of the Ezsigntemplateglobalannotation | 
**IEzsigntemplateglobaldocumentpagePagenumber** | **int32** | The page number in the Ezsigntemplateglobaldocument | 
**SEzsigntemplateglobalannotationDescription** | **string** | The description of the Ezsigntemplateglobalannotation | 
**SEzsigntemplateglobalannotationDefaulttext** | **string** | The defaulttext of the Ezsigntemplateglobalannotation | 
**SEzsigntemplateglobalannotationDropdownvalues** | **string** | The dropdownvalues of the Ezsigntemplateglobalannotation | 

## Methods

### NewEzsigntemplateglobalannotationResponseCompound

`func NewEzsigntemplateglobalannotationResponseCompound(pkiEzsigntemplateglobalannotationID int32, eEzsigntemplateglobalannotationHorizontalalignment FieldEEzsigntemplateglobalannotationHorizontalalignment, eEzsigntemplateglobalannotationVerticalalignment FieldEEzsigntemplateglobalannotationVerticalalignment, eEzsigntemplateglobalannotationType FieldEEzsigntemplateglobalannotationType, iEzsigntemplateglobalannotationX int32, iEzsigntemplateglobalannotationY int32, iEzsigntemplateglobalannotationWidth int32, iEzsigntemplateglobalannotationHeight int32, iEzsigntemplateglobaldocumentpagePagenumber int32, sEzsigntemplateglobalannotationDescription string, sEzsigntemplateglobalannotationDefaulttext string, sEzsigntemplateglobalannotationDropdownvalues string, ) *EzsigntemplateglobalannotationResponseCompound`

NewEzsigntemplateglobalannotationResponseCompound instantiates a new EzsigntemplateglobalannotationResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateglobalannotationResponseCompoundWithDefaults

`func NewEzsigntemplateglobalannotationResponseCompoundWithDefaults() *EzsigntemplateglobalannotationResponseCompound`

NewEzsigntemplateglobalannotationResponseCompoundWithDefaults instantiates a new EzsigntemplateglobalannotationResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateglobalannotationID

`func (o *EzsigntemplateglobalannotationResponseCompound) GetPkiEzsigntemplateglobalannotationID() int32`

GetPkiEzsigntemplateglobalannotationID returns the PkiEzsigntemplateglobalannotationID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateglobalannotationIDOk

`func (o *EzsigntemplateglobalannotationResponseCompound) GetPkiEzsigntemplateglobalannotationIDOk() (*int32, bool)`

GetPkiEzsigntemplateglobalannotationIDOk returns a tuple with the PkiEzsigntemplateglobalannotationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateglobalannotationID

`func (o *EzsigntemplateglobalannotationResponseCompound) SetPkiEzsigntemplateglobalannotationID(v int32)`

SetPkiEzsigntemplateglobalannotationID sets PkiEzsigntemplateglobalannotationID field to given value.


### GetFkiTextstylestaticID

`func (o *EzsigntemplateglobalannotationResponseCompound) GetFkiTextstylestaticID() int32`

GetFkiTextstylestaticID returns the FkiTextstylestaticID field if non-nil, zero value otherwise.

### GetFkiTextstylestaticIDOk

`func (o *EzsigntemplateglobalannotationResponseCompound) GetFkiTextstylestaticIDOk() (*int32, bool)`

GetFkiTextstylestaticIDOk returns a tuple with the FkiTextstylestaticID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTextstylestaticID

`func (o *EzsigntemplateglobalannotationResponseCompound) SetFkiTextstylestaticID(v int32)`

SetFkiTextstylestaticID sets FkiTextstylestaticID field to given value.

### HasFkiTextstylestaticID

`func (o *EzsigntemplateglobalannotationResponseCompound) HasFkiTextstylestaticID() bool`

HasFkiTextstylestaticID returns a boolean if a field has been set.

### GetObjTextstylestatic

`func (o *EzsigntemplateglobalannotationResponseCompound) GetObjTextstylestatic() TextstylestaticRequestCompound`

GetObjTextstylestatic returns the ObjTextstylestatic field if non-nil, zero value otherwise.

### GetObjTextstylestaticOk

`func (o *EzsigntemplateglobalannotationResponseCompound) GetObjTextstylestaticOk() (*TextstylestaticRequestCompound, bool)`

GetObjTextstylestaticOk returns a tuple with the ObjTextstylestatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjTextstylestatic

`func (o *EzsigntemplateglobalannotationResponseCompound) SetObjTextstylestatic(v TextstylestaticRequestCompound)`

SetObjTextstylestatic sets ObjTextstylestatic field to given value.

### HasObjTextstylestatic

`func (o *EzsigntemplateglobalannotationResponseCompound) HasObjTextstylestatic() bool`

HasObjTextstylestatic returns a boolean if a field has been set.

### GetEEzsigntemplateglobalannotationHorizontalalignment

`func (o *EzsigntemplateglobalannotationResponseCompound) GetEEzsigntemplateglobalannotationHorizontalalignment() FieldEEzsigntemplateglobalannotationHorizontalalignment`

GetEEzsigntemplateglobalannotationHorizontalalignment returns the EEzsigntemplateglobalannotationHorizontalalignment field if non-nil, zero value otherwise.

### GetEEzsigntemplateglobalannotationHorizontalalignmentOk

`func (o *EzsigntemplateglobalannotationResponseCompound) GetEEzsigntemplateglobalannotationHorizontalalignmentOk() (*FieldEEzsigntemplateglobalannotationHorizontalalignment, bool)`

GetEEzsigntemplateglobalannotationHorizontalalignmentOk returns a tuple with the EEzsigntemplateglobalannotationHorizontalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateglobalannotationHorizontalalignment

`func (o *EzsigntemplateglobalannotationResponseCompound) SetEEzsigntemplateglobalannotationHorizontalalignment(v FieldEEzsigntemplateglobalannotationHorizontalalignment)`

SetEEzsigntemplateglobalannotationHorizontalalignment sets EEzsigntemplateglobalannotationHorizontalalignment field to given value.


### GetEEzsigntemplateglobalannotationVerticalalignment

`func (o *EzsigntemplateglobalannotationResponseCompound) GetEEzsigntemplateglobalannotationVerticalalignment() FieldEEzsigntemplateglobalannotationVerticalalignment`

GetEEzsigntemplateglobalannotationVerticalalignment returns the EEzsigntemplateglobalannotationVerticalalignment field if non-nil, zero value otherwise.

### GetEEzsigntemplateglobalannotationVerticalalignmentOk

`func (o *EzsigntemplateglobalannotationResponseCompound) GetEEzsigntemplateglobalannotationVerticalalignmentOk() (*FieldEEzsigntemplateglobalannotationVerticalalignment, bool)`

GetEEzsigntemplateglobalannotationVerticalalignmentOk returns a tuple with the EEzsigntemplateglobalannotationVerticalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateglobalannotationVerticalalignment

`func (o *EzsigntemplateglobalannotationResponseCompound) SetEEzsigntemplateglobalannotationVerticalalignment(v FieldEEzsigntemplateglobalannotationVerticalalignment)`

SetEEzsigntemplateglobalannotationVerticalalignment sets EEzsigntemplateglobalannotationVerticalalignment field to given value.


### GetEEzsigntemplateglobalannotationType

`func (o *EzsigntemplateglobalannotationResponseCompound) GetEEzsigntemplateglobalannotationType() FieldEEzsigntemplateglobalannotationType`

GetEEzsigntemplateglobalannotationType returns the EEzsigntemplateglobalannotationType field if non-nil, zero value otherwise.

### GetEEzsigntemplateglobalannotationTypeOk

`func (o *EzsigntemplateglobalannotationResponseCompound) GetEEzsigntemplateglobalannotationTypeOk() (*FieldEEzsigntemplateglobalannotationType, bool)`

GetEEzsigntemplateglobalannotationTypeOk returns a tuple with the EEzsigntemplateglobalannotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateglobalannotationType

`func (o *EzsigntemplateglobalannotationResponseCompound) SetEEzsigntemplateglobalannotationType(v FieldEEzsigntemplateglobalannotationType)`

SetEEzsigntemplateglobalannotationType sets EEzsigntemplateglobalannotationType field to given value.


### GetIEzsigntemplateglobalannotationX

`func (o *EzsigntemplateglobalannotationResponseCompound) GetIEzsigntemplateglobalannotationX() int32`

GetIEzsigntemplateglobalannotationX returns the IEzsigntemplateglobalannotationX field if non-nil, zero value otherwise.

### GetIEzsigntemplateglobalannotationXOk

`func (o *EzsigntemplateglobalannotationResponseCompound) GetIEzsigntemplateglobalannotationXOk() (*int32, bool)`

GetIEzsigntemplateglobalannotationXOk returns a tuple with the IEzsigntemplateglobalannotationX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateglobalannotationX

`func (o *EzsigntemplateglobalannotationResponseCompound) SetIEzsigntemplateglobalannotationX(v int32)`

SetIEzsigntemplateglobalannotationX sets IEzsigntemplateglobalannotationX field to given value.


### GetIEzsigntemplateglobalannotationY

`func (o *EzsigntemplateglobalannotationResponseCompound) GetIEzsigntemplateglobalannotationY() int32`

GetIEzsigntemplateglobalannotationY returns the IEzsigntemplateglobalannotationY field if non-nil, zero value otherwise.

### GetIEzsigntemplateglobalannotationYOk

`func (o *EzsigntemplateglobalannotationResponseCompound) GetIEzsigntemplateglobalannotationYOk() (*int32, bool)`

GetIEzsigntemplateglobalannotationYOk returns a tuple with the IEzsigntemplateglobalannotationY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateglobalannotationY

`func (o *EzsigntemplateglobalannotationResponseCompound) SetIEzsigntemplateglobalannotationY(v int32)`

SetIEzsigntemplateglobalannotationY sets IEzsigntemplateglobalannotationY field to given value.


### GetIEzsigntemplateglobalannotationWidth

`func (o *EzsigntemplateglobalannotationResponseCompound) GetIEzsigntemplateglobalannotationWidth() int32`

GetIEzsigntemplateglobalannotationWidth returns the IEzsigntemplateglobalannotationWidth field if non-nil, zero value otherwise.

### GetIEzsigntemplateglobalannotationWidthOk

`func (o *EzsigntemplateglobalannotationResponseCompound) GetIEzsigntemplateglobalannotationWidthOk() (*int32, bool)`

GetIEzsigntemplateglobalannotationWidthOk returns a tuple with the IEzsigntemplateglobalannotationWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateglobalannotationWidth

`func (o *EzsigntemplateglobalannotationResponseCompound) SetIEzsigntemplateglobalannotationWidth(v int32)`

SetIEzsigntemplateglobalannotationWidth sets IEzsigntemplateglobalannotationWidth field to given value.


### GetIEzsigntemplateglobalannotationHeight

`func (o *EzsigntemplateglobalannotationResponseCompound) GetIEzsigntemplateglobalannotationHeight() int32`

GetIEzsigntemplateglobalannotationHeight returns the IEzsigntemplateglobalannotationHeight field if non-nil, zero value otherwise.

### GetIEzsigntemplateglobalannotationHeightOk

`func (o *EzsigntemplateglobalannotationResponseCompound) GetIEzsigntemplateglobalannotationHeightOk() (*int32, bool)`

GetIEzsigntemplateglobalannotationHeightOk returns a tuple with the IEzsigntemplateglobalannotationHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateglobalannotationHeight

`func (o *EzsigntemplateglobalannotationResponseCompound) SetIEzsigntemplateglobalannotationHeight(v int32)`

SetIEzsigntemplateglobalannotationHeight sets IEzsigntemplateglobalannotationHeight field to given value.


### GetIEzsigntemplateglobaldocumentpagePagenumber

`func (o *EzsigntemplateglobalannotationResponseCompound) GetIEzsigntemplateglobaldocumentpagePagenumber() int32`

GetIEzsigntemplateglobaldocumentpagePagenumber returns the IEzsigntemplateglobaldocumentpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsigntemplateglobaldocumentpagePagenumberOk

`func (o *EzsigntemplateglobalannotationResponseCompound) GetIEzsigntemplateglobaldocumentpagePagenumberOk() (*int32, bool)`

GetIEzsigntemplateglobaldocumentpagePagenumberOk returns a tuple with the IEzsigntemplateglobaldocumentpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateglobaldocumentpagePagenumber

`func (o *EzsigntemplateglobalannotationResponseCompound) SetIEzsigntemplateglobaldocumentpagePagenumber(v int32)`

SetIEzsigntemplateglobaldocumentpagePagenumber sets IEzsigntemplateglobaldocumentpagePagenumber field to given value.


### GetSEzsigntemplateglobalannotationDescription

`func (o *EzsigntemplateglobalannotationResponseCompound) GetSEzsigntemplateglobalannotationDescription() string`

GetSEzsigntemplateglobalannotationDescription returns the SEzsigntemplateglobalannotationDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplateglobalannotationDescriptionOk

`func (o *EzsigntemplateglobalannotationResponseCompound) GetSEzsigntemplateglobalannotationDescriptionOk() (*string, bool)`

GetSEzsigntemplateglobalannotationDescriptionOk returns a tuple with the SEzsigntemplateglobalannotationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateglobalannotationDescription

`func (o *EzsigntemplateglobalannotationResponseCompound) SetSEzsigntemplateglobalannotationDescription(v string)`

SetSEzsigntemplateglobalannotationDescription sets SEzsigntemplateglobalannotationDescription field to given value.


### GetSEzsigntemplateglobalannotationDefaulttext

`func (o *EzsigntemplateglobalannotationResponseCompound) GetSEzsigntemplateglobalannotationDefaulttext() string`

GetSEzsigntemplateglobalannotationDefaulttext returns the SEzsigntemplateglobalannotationDefaulttext field if non-nil, zero value otherwise.

### GetSEzsigntemplateglobalannotationDefaulttextOk

`func (o *EzsigntemplateglobalannotationResponseCompound) GetSEzsigntemplateglobalannotationDefaulttextOk() (*string, bool)`

GetSEzsigntemplateglobalannotationDefaulttextOk returns a tuple with the SEzsigntemplateglobalannotationDefaulttext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateglobalannotationDefaulttext

`func (o *EzsigntemplateglobalannotationResponseCompound) SetSEzsigntemplateglobalannotationDefaulttext(v string)`

SetSEzsigntemplateglobalannotationDefaulttext sets SEzsigntemplateglobalannotationDefaulttext field to given value.


### GetSEzsigntemplateglobalannotationDropdownvalues

`func (o *EzsigntemplateglobalannotationResponseCompound) GetSEzsigntemplateglobalannotationDropdownvalues() string`

GetSEzsigntemplateglobalannotationDropdownvalues returns the SEzsigntemplateglobalannotationDropdownvalues field if non-nil, zero value otherwise.

### GetSEzsigntemplateglobalannotationDropdownvaluesOk

`func (o *EzsigntemplateglobalannotationResponseCompound) GetSEzsigntemplateglobalannotationDropdownvaluesOk() (*string, bool)`

GetSEzsigntemplateglobalannotationDropdownvaluesOk returns a tuple with the SEzsigntemplateglobalannotationDropdownvalues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateglobalannotationDropdownvalues

`func (o *EzsigntemplateglobalannotationResponseCompound) SetSEzsigntemplateglobalannotationDropdownvalues(v string)`

SetSEzsigntemplateglobalannotationDropdownvalues sets SEzsigntemplateglobalannotationDropdownvalues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


