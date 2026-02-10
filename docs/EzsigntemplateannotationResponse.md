# EzsigntemplateannotationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplateannotationID** | **int32** | The unique ID of the Ezsigntemplateannotation | 
**FkiEzsigntemplatedocumentpageID** | **int32** | The unique ID of the Ezsigntemplatedocumentpage | 
**FkiTextstylestaticID** | **int32** | The unique ID of the Textstylestatic | 
**EEzsigntemplateannotationHorizontalalignment** | [**FieldEEzsigntemplateannotationHorizontalalignment**](FieldEEzsigntemplateannotationHorizontalalignment.md) |  | 
**EEzsigntemplateannotationVerticalalignment** | [**FieldEEzsigntemplateannotationVerticalalignment**](FieldEEzsigntemplateannotationVerticalalignment.md) |  | 
**EEzsigntemplateannotationType** | [**FieldEEzsigntemplateannotationType**](FieldEEzsigntemplateannotationType.md) |  | 
**IEzsigntemplateannotationX** | **int32** | The x of the Ezsigntemplateannotation | 
**IEzsigntemplateannotationY** | **int32** | The y of the Ezsigntemplateannotation | 
**IEzsigntemplateannotationWidth** | **int32** | The width of the Ezsigntemplateannotation | 
**IEzsigntemplateannotationHeight** | **int32** | The height of the Ezsigntemplateannotation | 
**SEzsigntemplateannotationDescription** | **string** | The description of the Ezsigntemplateannotation | 
**SEzsigntemplateannotationDefaulttext** | **string** | The defaulttext of the Ezsigntemplateannotation | 
**SEzsigntemplateannotationnDropdownvalues** | **string** | The ndropdownvalues of the Ezsigntemplateannotation | 

## Methods

### NewEzsigntemplateannotationResponse

`func NewEzsigntemplateannotationResponse(pkiEzsigntemplateannotationID int32, fkiEzsigntemplatedocumentpageID int32, fkiTextstylestaticID int32, eEzsigntemplateannotationHorizontalalignment FieldEEzsigntemplateannotationHorizontalalignment, eEzsigntemplateannotationVerticalalignment FieldEEzsigntemplateannotationVerticalalignment, eEzsigntemplateannotationType FieldEEzsigntemplateannotationType, iEzsigntemplateannotationX int32, iEzsigntemplateannotationY int32, iEzsigntemplateannotationWidth int32, iEzsigntemplateannotationHeight int32, sEzsigntemplateannotationDescription string, sEzsigntemplateannotationDefaulttext string, sEzsigntemplateannotationnDropdownvalues string, ) *EzsigntemplateannotationResponse`

NewEzsigntemplateannotationResponse instantiates a new EzsigntemplateannotationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplateannotationResponseWithDefaults

`func NewEzsigntemplateannotationResponseWithDefaults() *EzsigntemplateannotationResponse`

NewEzsigntemplateannotationResponseWithDefaults instantiates a new EzsigntemplateannotationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplateannotationID

`func (o *EzsigntemplateannotationResponse) GetPkiEzsigntemplateannotationID() int32`

GetPkiEzsigntemplateannotationID returns the PkiEzsigntemplateannotationID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplateannotationIDOk

`func (o *EzsigntemplateannotationResponse) GetPkiEzsigntemplateannotationIDOk() (*int32, bool)`

GetPkiEzsigntemplateannotationIDOk returns a tuple with the PkiEzsigntemplateannotationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplateannotationID

`func (o *EzsigntemplateannotationResponse) SetPkiEzsigntemplateannotationID(v int32)`

SetPkiEzsigntemplateannotationID sets PkiEzsigntemplateannotationID field to given value.


### GetFkiEzsigntemplatedocumentpageID

`func (o *EzsigntemplateannotationResponse) GetFkiEzsigntemplatedocumentpageID() int32`

GetFkiEzsigntemplatedocumentpageID returns the FkiEzsigntemplatedocumentpageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatedocumentpageIDOk

`func (o *EzsigntemplateannotationResponse) GetFkiEzsigntemplatedocumentpageIDOk() (*int32, bool)`

GetFkiEzsigntemplatedocumentpageIDOk returns a tuple with the FkiEzsigntemplatedocumentpageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatedocumentpageID

`func (o *EzsigntemplateannotationResponse) SetFkiEzsigntemplatedocumentpageID(v int32)`

SetFkiEzsigntemplatedocumentpageID sets FkiEzsigntemplatedocumentpageID field to given value.


### GetFkiTextstylestaticID

`func (o *EzsigntemplateannotationResponse) GetFkiTextstylestaticID() int32`

GetFkiTextstylestaticID returns the FkiTextstylestaticID field if non-nil, zero value otherwise.

### GetFkiTextstylestaticIDOk

`func (o *EzsigntemplateannotationResponse) GetFkiTextstylestaticIDOk() (*int32, bool)`

GetFkiTextstylestaticIDOk returns a tuple with the FkiTextstylestaticID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTextstylestaticID

`func (o *EzsigntemplateannotationResponse) SetFkiTextstylestaticID(v int32)`

SetFkiTextstylestaticID sets FkiTextstylestaticID field to given value.


### GetEEzsigntemplateannotationHorizontalalignment

`func (o *EzsigntemplateannotationResponse) GetEEzsigntemplateannotationHorizontalalignment() FieldEEzsigntemplateannotationHorizontalalignment`

GetEEzsigntemplateannotationHorizontalalignment returns the EEzsigntemplateannotationHorizontalalignment field if non-nil, zero value otherwise.

### GetEEzsigntemplateannotationHorizontalalignmentOk

`func (o *EzsigntemplateannotationResponse) GetEEzsigntemplateannotationHorizontalalignmentOk() (*FieldEEzsigntemplateannotationHorizontalalignment, bool)`

GetEEzsigntemplateannotationHorizontalalignmentOk returns a tuple with the EEzsigntemplateannotationHorizontalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateannotationHorizontalalignment

`func (o *EzsigntemplateannotationResponse) SetEEzsigntemplateannotationHorizontalalignment(v FieldEEzsigntemplateannotationHorizontalalignment)`

SetEEzsigntemplateannotationHorizontalalignment sets EEzsigntemplateannotationHorizontalalignment field to given value.


### GetEEzsigntemplateannotationVerticalalignment

`func (o *EzsigntemplateannotationResponse) GetEEzsigntemplateannotationVerticalalignment() FieldEEzsigntemplateannotationVerticalalignment`

GetEEzsigntemplateannotationVerticalalignment returns the EEzsigntemplateannotationVerticalalignment field if non-nil, zero value otherwise.

### GetEEzsigntemplateannotationVerticalalignmentOk

`func (o *EzsigntemplateannotationResponse) GetEEzsigntemplateannotationVerticalalignmentOk() (*FieldEEzsigntemplateannotationVerticalalignment, bool)`

GetEEzsigntemplateannotationVerticalalignmentOk returns a tuple with the EEzsigntemplateannotationVerticalalignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateannotationVerticalalignment

`func (o *EzsigntemplateannotationResponse) SetEEzsigntemplateannotationVerticalalignment(v FieldEEzsigntemplateannotationVerticalalignment)`

SetEEzsigntemplateannotationVerticalalignment sets EEzsigntemplateannotationVerticalalignment field to given value.


### GetEEzsigntemplateannotationType

`func (o *EzsigntemplateannotationResponse) GetEEzsigntemplateannotationType() FieldEEzsigntemplateannotationType`

GetEEzsigntemplateannotationType returns the EEzsigntemplateannotationType field if non-nil, zero value otherwise.

### GetEEzsigntemplateannotationTypeOk

`func (o *EzsigntemplateannotationResponse) GetEEzsigntemplateannotationTypeOk() (*FieldEEzsigntemplateannotationType, bool)`

GetEEzsigntemplateannotationTypeOk returns a tuple with the EEzsigntemplateannotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplateannotationType

`func (o *EzsigntemplateannotationResponse) SetEEzsigntemplateannotationType(v FieldEEzsigntemplateannotationType)`

SetEEzsigntemplateannotationType sets EEzsigntemplateannotationType field to given value.


### GetIEzsigntemplateannotationX

`func (o *EzsigntemplateannotationResponse) GetIEzsigntemplateannotationX() int32`

GetIEzsigntemplateannotationX returns the IEzsigntemplateannotationX field if non-nil, zero value otherwise.

### GetIEzsigntemplateannotationXOk

`func (o *EzsigntemplateannotationResponse) GetIEzsigntemplateannotationXOk() (*int32, bool)`

GetIEzsigntemplateannotationXOk returns a tuple with the IEzsigntemplateannotationX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateannotationX

`func (o *EzsigntemplateannotationResponse) SetIEzsigntemplateannotationX(v int32)`

SetIEzsigntemplateannotationX sets IEzsigntemplateannotationX field to given value.


### GetIEzsigntemplateannotationY

`func (o *EzsigntemplateannotationResponse) GetIEzsigntemplateannotationY() int32`

GetIEzsigntemplateannotationY returns the IEzsigntemplateannotationY field if non-nil, zero value otherwise.

### GetIEzsigntemplateannotationYOk

`func (o *EzsigntemplateannotationResponse) GetIEzsigntemplateannotationYOk() (*int32, bool)`

GetIEzsigntemplateannotationYOk returns a tuple with the IEzsigntemplateannotationY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateannotationY

`func (o *EzsigntemplateannotationResponse) SetIEzsigntemplateannotationY(v int32)`

SetIEzsigntemplateannotationY sets IEzsigntemplateannotationY field to given value.


### GetIEzsigntemplateannotationWidth

`func (o *EzsigntemplateannotationResponse) GetIEzsigntemplateannotationWidth() int32`

GetIEzsigntemplateannotationWidth returns the IEzsigntemplateannotationWidth field if non-nil, zero value otherwise.

### GetIEzsigntemplateannotationWidthOk

`func (o *EzsigntemplateannotationResponse) GetIEzsigntemplateannotationWidthOk() (*int32, bool)`

GetIEzsigntemplateannotationWidthOk returns a tuple with the IEzsigntemplateannotationWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateannotationWidth

`func (o *EzsigntemplateannotationResponse) SetIEzsigntemplateannotationWidth(v int32)`

SetIEzsigntemplateannotationWidth sets IEzsigntemplateannotationWidth field to given value.


### GetIEzsigntemplateannotationHeight

`func (o *EzsigntemplateannotationResponse) GetIEzsigntemplateannotationHeight() int32`

GetIEzsigntemplateannotationHeight returns the IEzsigntemplateannotationHeight field if non-nil, zero value otherwise.

### GetIEzsigntemplateannotationHeightOk

`func (o *EzsigntemplateannotationResponse) GetIEzsigntemplateannotationHeightOk() (*int32, bool)`

GetIEzsigntemplateannotationHeightOk returns a tuple with the IEzsigntemplateannotationHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplateannotationHeight

`func (o *EzsigntemplateannotationResponse) SetIEzsigntemplateannotationHeight(v int32)`

SetIEzsigntemplateannotationHeight sets IEzsigntemplateannotationHeight field to given value.


### GetSEzsigntemplateannotationDescription

`func (o *EzsigntemplateannotationResponse) GetSEzsigntemplateannotationDescription() string`

GetSEzsigntemplateannotationDescription returns the SEzsigntemplateannotationDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplateannotationDescriptionOk

`func (o *EzsigntemplateannotationResponse) GetSEzsigntemplateannotationDescriptionOk() (*string, bool)`

GetSEzsigntemplateannotationDescriptionOk returns a tuple with the SEzsigntemplateannotationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateannotationDescription

`func (o *EzsigntemplateannotationResponse) SetSEzsigntemplateannotationDescription(v string)`

SetSEzsigntemplateannotationDescription sets SEzsigntemplateannotationDescription field to given value.


### GetSEzsigntemplateannotationDefaulttext

`func (o *EzsigntemplateannotationResponse) GetSEzsigntemplateannotationDefaulttext() string`

GetSEzsigntemplateannotationDefaulttext returns the SEzsigntemplateannotationDefaulttext field if non-nil, zero value otherwise.

### GetSEzsigntemplateannotationDefaulttextOk

`func (o *EzsigntemplateannotationResponse) GetSEzsigntemplateannotationDefaulttextOk() (*string, bool)`

GetSEzsigntemplateannotationDefaulttextOk returns a tuple with the SEzsigntemplateannotationDefaulttext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateannotationDefaulttext

`func (o *EzsigntemplateannotationResponse) SetSEzsigntemplateannotationDefaulttext(v string)`

SetSEzsigntemplateannotationDefaulttext sets SEzsigntemplateannotationDefaulttext field to given value.


### GetSEzsigntemplateannotationnDropdownvalues

`func (o *EzsigntemplateannotationResponse) GetSEzsigntemplateannotationnDropdownvalues() string`

GetSEzsigntemplateannotationnDropdownvalues returns the SEzsigntemplateannotationnDropdownvalues field if non-nil, zero value otherwise.

### GetSEzsigntemplateannotationnDropdownvaluesOk

`func (o *EzsigntemplateannotationResponse) GetSEzsigntemplateannotationnDropdownvaluesOk() (*string, bool)`

GetSEzsigntemplateannotationnDropdownvaluesOk returns a tuple with the SEzsigntemplateannotationnDropdownvalues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplateannotationnDropdownvalues

`func (o *EzsigntemplateannotationResponse) SetSEzsigntemplateannotationnDropdownvalues(v string)`

SetSEzsigntemplateannotationnDropdownvalues sets SEzsigntemplateannotationnDropdownvalues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


