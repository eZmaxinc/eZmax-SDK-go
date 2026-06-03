# EzsigntemplateannotationResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateannotationID** | **int32** | The unique ID of the Ezsigntemplateannotation | 
**FkiTextstylestaticID** | Pointer to **int32** | The unique ID of the Textstylestatic | [optional] 
**ObjTextstylestatic** | Pointer to [**TextstylestaticRequestCompound**](TextstylestaticRequestCompound.md) |  | [optional] 
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

## Methods

### NewEzsigntemplateannotationResponseCompound

`func NewEzsigntemplateannotationResponseCompound(pkiEzsigntemplateannotationID int32, eEzsigntemplateannotationHorizontalalignment FieldEEzsigntemplateannotationHorizontalalignment, eEzsigntemplateannotationVerticalalignment FieldEEzsigntemplateannotationVerticalalignment, eEzsigntemplateannotationType FieldEEzsigntemplateannotationType, iEzsigntemplateannotationX int32, iEzsigntemplateannotationY int32, iEzsigntemplateannotationWidth int32, iEzsigntemplateannotationHeight int32, iEzsigntemplatedocumentpagePagenumber int32, sEzsigntemplateannotationDescription string, sEzsigntemplateannotationDefaulttext string, sEzsigntemplateannotationDropdownvalues string, ) *EzsigntemplateannotationResponseCompound`

NewEzsigntemplateannotationResponseCompound instantiates a new EzsigntemplateannotationResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateannotationResponseCompoundWithDefaults

`func NewEzsigntemplateannotationResponseCompoundWithDefaults() *EzsigntemplateannotationResponseCompound`

NewEzsigntemplateannotationResponseCompoundWithDefaults instantiates a new EzsigntemplateannotationResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateannotationID

`func (o *EzsigntemplateannotationResponseCompound) GetPkiEzsigntemplateannotationID() int32`

GetPkiEzsigntemplateannotationID returns the PkiEzsigntemplateannotationID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateannotationIDOk

`func (o *EzsigntemplateannotationResponseCompound) GetPkiEzsigntemplateannotationIDOk() (*int32, bool)`

GetPkiEzsigntemplateannotationIDOk returns a tuple with the PkiEzsigntemplateannotationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateannotationID

`func (o *EzsigntemplateannotationResponseCompound) SetPkiEzsigntemplateannotationID(v int32)`

SetPkiEzsigntemplateannotationID sets PkiEzsigntemplateannotationID field to given value.


### GetFkiTextstylestaticID

`func (o *EzsigntemplateannotationResponseCompound) GetFkiTextstylestaticID() int32`

GetFkiTextstylestaticID returns the FkiTextstylestaticID field if non-nil, zero value otherwise.

### GetFkiTextstylestaticIDOk

`func (o *EzsigntemplateannotationResponseCompound) GetFkiTextstylestaticIDOk() (*int32, bool)`

GetFkiTextstylestaticIDOk returns a tuple with the FkiTextstylestaticID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTextstylestaticID

`func (o *EzsigntemplateannotationResponseCompound) SetFkiTextstylestaticID(v int32)`

SetFkiTextstylestaticID sets FkiTextstylestaticID field to given value.

### HasFkiTextstylestaticID

`func (o *EzsigntemplateannotationResponseCompound) HasFkiTextstylestaticID() bool`

HasFkiTextstylestaticID returns a boolean if a field has been set.

### GetObjTextstylestatic

`func (o *EzsigntemplateannotationResponseCompound) GetObjTextstylestatic() TextstylestaticRequestCompound`

GetObjTextstylestatic returns the ObjTextstylestatic field if non-nil, zero value otherwise.

### GetObjTextstylestaticOk

`func (o *EzsigntemplateannotationResponseCompound) GetObjTextstylestaticOk() (*TextstylestaticRequestCompound, bool)`

GetObjTextstylestaticOk returns a tuple with the ObjTextstylestatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjTextstylestatic

`func (o *EzsigntemplateannotationResponseCompound) SetObjTextstylestatic(v TextstylestaticRequestCompound)`

SetObjTextstylestatic sets ObjTextstylestatic field to given value.

### HasObjTextstylestatic

`func (o *EzsigntemplateannotationResponseCompound) HasObjTextstylestatic() bool`

HasObjTextstylestatic returns a boolean if a field has been set.

### GetEEzsigntemplateannotationHorizontalalignment

`func (o *EzsigntemplateannotationResponseCompound) GetEEzsigntemplateannotationHorizontalalignment() FieldEEzsigntemplateannotationHorizontalalignment`

GetEEzsigntemplateannotationHorizontalalignment returns the EEzsigntemplateannotationHorizontalalignment field if non-nil, zero value otherwise.

### GetEEzsigntemplateannotationHorizontalalignmentOk

`func (o *EzsigntemplateannotationResponseCompound) GetEEzsigntemplateannotationHorizontalalignmentOk() (*FieldEEzsigntemplateannotationHorizontalalignment, bool)`

GetEEzsigntemplateannotationHorizontalalignmentOk returns a tuple with the EEzsigntemplateannotationHorizontalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateannotationHorizontalalignment

`func (o *EzsigntemplateannotationResponseCompound) SetEEzsigntemplateannotationHorizontalalignment(v FieldEEzsigntemplateannotationHorizontalalignment)`

SetEEzsigntemplateannotationHorizontalalignment sets EEzsigntemplateannotationHorizontalalignment field to given value.


### GetEEzsigntemplateannotationVerticalalignment

`func (o *EzsigntemplateannotationResponseCompound) GetEEzsigntemplateannotationVerticalalignment() FieldEEzsigntemplateannotationVerticalalignment`

GetEEzsigntemplateannotationVerticalalignment returns the EEzsigntemplateannotationVerticalalignment field if non-nil, zero value otherwise.

### GetEEzsigntemplateannotationVerticalalignmentOk

`func (o *EzsigntemplateannotationResponseCompound) GetEEzsigntemplateannotationVerticalalignmentOk() (*FieldEEzsigntemplateannotationVerticalalignment, bool)`

GetEEzsigntemplateannotationVerticalalignmentOk returns a tuple with the EEzsigntemplateannotationVerticalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateannotationVerticalalignment

`func (o *EzsigntemplateannotationResponseCompound) SetEEzsigntemplateannotationVerticalalignment(v FieldEEzsigntemplateannotationVerticalalignment)`

SetEEzsigntemplateannotationVerticalalignment sets EEzsigntemplateannotationVerticalalignment field to given value.


### GetEEzsigntemplateannotationType

`func (o *EzsigntemplateannotationResponseCompound) GetEEzsigntemplateannotationType() FieldEEzsigntemplateannotationType`

GetEEzsigntemplateannotationType returns the EEzsigntemplateannotationType field if non-nil, zero value otherwise.

### GetEEzsigntemplateannotationTypeOk

`func (o *EzsigntemplateannotationResponseCompound) GetEEzsigntemplateannotationTypeOk() (*FieldEEzsigntemplateannotationType, bool)`

GetEEzsigntemplateannotationTypeOk returns a tuple with the EEzsigntemplateannotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateannotationType

`func (o *EzsigntemplateannotationResponseCompound) SetEEzsigntemplateannotationType(v FieldEEzsigntemplateannotationType)`

SetEEzsigntemplateannotationType sets EEzsigntemplateannotationType field to given value.


### GetIEzsigntemplateannotationX

`func (o *EzsigntemplateannotationResponseCompound) GetIEzsigntemplateannotationX() int32`

GetIEzsigntemplateannotationX returns the IEzsigntemplateannotationX field if non-nil, zero value otherwise.

### GetIEzsigntemplateannotationXOk

`func (o *EzsigntemplateannotationResponseCompound) GetIEzsigntemplateannotationXOk() (*int32, bool)`

GetIEzsigntemplateannotationXOk returns a tuple with the IEzsigntemplateannotationX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateannotationX

`func (o *EzsigntemplateannotationResponseCompound) SetIEzsigntemplateannotationX(v int32)`

SetIEzsigntemplateannotationX sets IEzsigntemplateannotationX field to given value.


### GetIEzsigntemplateannotationY

`func (o *EzsigntemplateannotationResponseCompound) GetIEzsigntemplateannotationY() int32`

GetIEzsigntemplateannotationY returns the IEzsigntemplateannotationY field if non-nil, zero value otherwise.

### GetIEzsigntemplateannotationYOk

`func (o *EzsigntemplateannotationResponseCompound) GetIEzsigntemplateannotationYOk() (*int32, bool)`

GetIEzsigntemplateannotationYOk returns a tuple with the IEzsigntemplateannotationY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateannotationY

`func (o *EzsigntemplateannotationResponseCompound) SetIEzsigntemplateannotationY(v int32)`

SetIEzsigntemplateannotationY sets IEzsigntemplateannotationY field to given value.


### GetIEzsigntemplateannotationWidth

`func (o *EzsigntemplateannotationResponseCompound) GetIEzsigntemplateannotationWidth() int32`

GetIEzsigntemplateannotationWidth returns the IEzsigntemplateannotationWidth field if non-nil, zero value otherwise.

### GetIEzsigntemplateannotationWidthOk

`func (o *EzsigntemplateannotationResponseCompound) GetIEzsigntemplateannotationWidthOk() (*int32, bool)`

GetIEzsigntemplateannotationWidthOk returns a tuple with the IEzsigntemplateannotationWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateannotationWidth

`func (o *EzsigntemplateannotationResponseCompound) SetIEzsigntemplateannotationWidth(v int32)`

SetIEzsigntemplateannotationWidth sets IEzsigntemplateannotationWidth field to given value.


### GetIEzsigntemplateannotationHeight

`func (o *EzsigntemplateannotationResponseCompound) GetIEzsigntemplateannotationHeight() int32`

GetIEzsigntemplateannotationHeight returns the IEzsigntemplateannotationHeight field if non-nil, zero value otherwise.

### GetIEzsigntemplateannotationHeightOk

`func (o *EzsigntemplateannotationResponseCompound) GetIEzsigntemplateannotationHeightOk() (*int32, bool)`

GetIEzsigntemplateannotationHeightOk returns a tuple with the IEzsigntemplateannotationHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateannotationHeight

`func (o *EzsigntemplateannotationResponseCompound) SetIEzsigntemplateannotationHeight(v int32)`

SetIEzsigntemplateannotationHeight sets IEzsigntemplateannotationHeight field to given value.


### GetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplateannotationResponseCompound) GetIEzsigntemplatedocumentpagePagenumber() int32`

GetIEzsigntemplatedocumentpagePagenumber returns the IEzsigntemplatedocumentpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagePagenumberOk

`func (o *EzsigntemplateannotationResponseCompound) GetIEzsigntemplatedocumentpagePagenumberOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagePagenumberOk returns a tuple with the IEzsigntemplatedocumentpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagePagenumber

`func (o *EzsigntemplateannotationResponseCompound) SetIEzsigntemplatedocumentpagePagenumber(v int32)`

SetIEzsigntemplatedocumentpagePagenumber sets IEzsigntemplatedocumentpagePagenumber field to given value.


### GetSEzsigntemplateannotationDescription

`func (o *EzsigntemplateannotationResponseCompound) GetSEzsigntemplateannotationDescription() string`

GetSEzsigntemplateannotationDescription returns the SEzsigntemplateannotationDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplateannotationDescriptionOk

`func (o *EzsigntemplateannotationResponseCompound) GetSEzsigntemplateannotationDescriptionOk() (*string, bool)`

GetSEzsigntemplateannotationDescriptionOk returns a tuple with the SEzsigntemplateannotationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateannotationDescription

`func (o *EzsigntemplateannotationResponseCompound) SetSEzsigntemplateannotationDescription(v string)`

SetSEzsigntemplateannotationDescription sets SEzsigntemplateannotationDescription field to given value.


### GetSEzsigntemplateannotationDefaulttext

`func (o *EzsigntemplateannotationResponseCompound) GetSEzsigntemplateannotationDefaulttext() string`

GetSEzsigntemplateannotationDefaulttext returns the SEzsigntemplateannotationDefaulttext field if non-nil, zero value otherwise.

### GetSEzsigntemplateannotationDefaulttextOk

`func (o *EzsigntemplateannotationResponseCompound) GetSEzsigntemplateannotationDefaulttextOk() (*string, bool)`

GetSEzsigntemplateannotationDefaulttextOk returns a tuple with the SEzsigntemplateannotationDefaulttext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateannotationDefaulttext

`func (o *EzsigntemplateannotationResponseCompound) SetSEzsigntemplateannotationDefaulttext(v string)`

SetSEzsigntemplateannotationDefaulttext sets SEzsigntemplateannotationDefaulttext field to given value.


### GetSEzsigntemplateannotationDropdownvalues

`func (o *EzsigntemplateannotationResponseCompound) GetSEzsigntemplateannotationDropdownvalues() string`

GetSEzsigntemplateannotationDropdownvalues returns the SEzsigntemplateannotationDropdownvalues field if non-nil, zero value otherwise.

### GetSEzsigntemplateannotationDropdownvaluesOk

`func (o *EzsigntemplateannotationResponseCompound) GetSEzsigntemplateannotationDropdownvaluesOk() (*string, bool)`

GetSEzsigntemplateannotationDropdownvaluesOk returns a tuple with the SEzsigntemplateannotationDropdownvalues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateannotationDropdownvalues

`func (o *EzsigntemplateannotationResponseCompound) SetSEzsigntemplateannotationDropdownvalues(v string)`

SetSEzsigntemplateannotationDropdownvalues sets SEzsigntemplateannotationDropdownvalues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


