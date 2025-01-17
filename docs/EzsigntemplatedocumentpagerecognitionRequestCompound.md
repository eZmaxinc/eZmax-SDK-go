# EzsigntemplatedocumentpagerecognitionRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatedocumentpagerecognitionID** | Pointer to **int32** | The unique ID of the Ezsigntemplatedocumentpagerecognition | [optional] 
**FkiEzsigntemplatedocumentpageID** | **int32** | The unique ID of the Ezsigntemplatedocumentpage | 
**EEzsigntemplatedocumentpagerecognitionOperator** | [**FieldEEzsigntemplatedocumentpagerecognitionOperator**](FieldEEzsigntemplatedocumentpagerecognitionOperator.md) |  | 
**EEzsigntemplatedocumentpagerecognitionSection** | [**FieldEEzsigntemplatedocumentpagerecognitionSection**](FieldEEzsigntemplatedocumentpagerecognitionSection.md) |  | 
**IEzsigntemplatedocumentpagerecognitionSimilarpercentage** | Pointer to **int32** | The similarpercentage of the Ezsigntemplatedocumentpagerecognition | [optional] 
**IEzsigntemplatedocumentpagerecognitionX** | Pointer to **int32** | The x of the Ezsigntemplatedocumentpagerecognition | [optional] 
**IEzsigntemplatedocumentpagerecognitionY** | Pointer to **int32** | The y of the Ezsigntemplatedocumentpagerecognition | [optional] 
**IEzsigntemplatedocumentpagerecognitionWidth** | Pointer to **int32** | The width of the Ezsigntemplatedocumentpagerecognition | [optional] 
**IEzsigntemplatedocumentpagerecognitionHeight** | Pointer to **int32** | The height of the Ezsigntemplatedocumentpagerecognition | [optional] 
**TEzsigntemplatedocumentpagerecognitionText** | **string** | The text of the Ezsigntemplatedocumentpagerecognition | 

## Methods

### NewEzsigntemplatedocumentpagerecognitionRequestCompound

`func NewEzsigntemplatedocumentpagerecognitionRequestCompound(fkiEzsigntemplatedocumentpageID int32, eEzsigntemplatedocumentpagerecognitionOperator FieldEEzsigntemplatedocumentpagerecognitionOperator, eEzsigntemplatedocumentpagerecognitionSection FieldEEzsigntemplatedocumentpagerecognitionSection, tEzsigntemplatedocumentpagerecognitionText string, ) *EzsigntemplatedocumentpagerecognitionRequestCompound`

NewEzsigntemplatedocumentpagerecognitionRequestCompound instantiates a new EzsigntemplatedocumentpagerecognitionRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatedocumentpagerecognitionRequestCompoundWithDefaults

`func NewEzsigntemplatedocumentpagerecognitionRequestCompoundWithDefaults() *EzsigntemplatedocumentpagerecognitionRequestCompound`

NewEzsigntemplatedocumentpagerecognitionRequestCompoundWithDefaults instantiates a new EzsigntemplatedocumentpagerecognitionRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatedocumentpagerecognitionID

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetPkiEzsigntemplatedocumentpagerecognitionID() int32`

GetPkiEzsigntemplatedocumentpagerecognitionID returns the PkiEzsigntemplatedocumentpagerecognitionID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatedocumentpagerecognitionIDOk

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetPkiEzsigntemplatedocumentpagerecognitionIDOk() (*int32, bool)`

GetPkiEzsigntemplatedocumentpagerecognitionIDOk returns a tuple with the PkiEzsigntemplatedocumentpagerecognitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatedocumentpagerecognitionID

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) SetPkiEzsigntemplatedocumentpagerecognitionID(v int32)`

SetPkiEzsigntemplatedocumentpagerecognitionID sets PkiEzsigntemplatedocumentpagerecognitionID field to given value.

### HasPkiEzsigntemplatedocumentpagerecognitionID

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) HasPkiEzsigntemplatedocumentpagerecognitionID() bool`

HasPkiEzsigntemplatedocumentpagerecognitionID returns a boolean if a field has been set.

### GetFkiEzsigntemplatedocumentpageID

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetFkiEzsigntemplatedocumentpageID() int32`

GetFkiEzsigntemplatedocumentpageID returns the FkiEzsigntemplatedocumentpageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatedocumentpageIDOk

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetFkiEzsigntemplatedocumentpageIDOk() (*int32, bool)`

GetFkiEzsigntemplatedocumentpageIDOk returns a tuple with the FkiEzsigntemplatedocumentpageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatedocumentpageID

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) SetFkiEzsigntemplatedocumentpageID(v int32)`

SetFkiEzsigntemplatedocumentpageID sets FkiEzsigntemplatedocumentpageID field to given value.


### GetEEzsigntemplatedocumentpagerecognitionOperator

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetEEzsigntemplatedocumentpagerecognitionOperator() FieldEEzsigntemplatedocumentpagerecognitionOperator`

GetEEzsigntemplatedocumentpagerecognitionOperator returns the EEzsigntemplatedocumentpagerecognitionOperator field if non-nil, zero value otherwise.

### GetEEzsigntemplatedocumentpagerecognitionOperatorOk

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetEEzsigntemplatedocumentpagerecognitionOperatorOk() (*FieldEEzsigntemplatedocumentpagerecognitionOperator, bool)`

GetEEzsigntemplatedocumentpagerecognitionOperatorOk returns a tuple with the EEzsigntemplatedocumentpagerecognitionOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatedocumentpagerecognitionOperator

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) SetEEzsigntemplatedocumentpagerecognitionOperator(v FieldEEzsigntemplatedocumentpagerecognitionOperator)`

SetEEzsigntemplatedocumentpagerecognitionOperator sets EEzsigntemplatedocumentpagerecognitionOperator field to given value.


### GetEEzsigntemplatedocumentpagerecognitionSection

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetEEzsigntemplatedocumentpagerecognitionSection() FieldEEzsigntemplatedocumentpagerecognitionSection`

GetEEzsigntemplatedocumentpagerecognitionSection returns the EEzsigntemplatedocumentpagerecognitionSection field if non-nil, zero value otherwise.

### GetEEzsigntemplatedocumentpagerecognitionSectionOk

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetEEzsigntemplatedocumentpagerecognitionSectionOk() (*FieldEEzsigntemplatedocumentpagerecognitionSection, bool)`

GetEEzsigntemplatedocumentpagerecognitionSectionOk returns a tuple with the EEzsigntemplatedocumentpagerecognitionSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatedocumentpagerecognitionSection

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) SetEEzsigntemplatedocumentpagerecognitionSection(v FieldEEzsigntemplatedocumentpagerecognitionSection)`

SetEEzsigntemplatedocumentpagerecognitionSection sets EEzsigntemplatedocumentpagerecognitionSection field to given value.


### GetIEzsigntemplatedocumentpagerecognitionSimilarpercentage

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetIEzsigntemplatedocumentpagerecognitionSimilarpercentage() int32`

GetIEzsigntemplatedocumentpagerecognitionSimilarpercentage returns the IEzsigntemplatedocumentpagerecognitionSimilarpercentage field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagerecognitionSimilarpercentageOk

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetIEzsigntemplatedocumentpagerecognitionSimilarpercentageOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagerecognitionSimilarpercentageOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionSimilarpercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagerecognitionSimilarpercentage

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) SetIEzsigntemplatedocumentpagerecognitionSimilarpercentage(v int32)`

SetIEzsigntemplatedocumentpagerecognitionSimilarpercentage sets IEzsigntemplatedocumentpagerecognitionSimilarpercentage field to given value.

### HasIEzsigntemplatedocumentpagerecognitionSimilarpercentage

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) HasIEzsigntemplatedocumentpagerecognitionSimilarpercentage() bool`

HasIEzsigntemplatedocumentpagerecognitionSimilarpercentage returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagerecognitionX

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetIEzsigntemplatedocumentpagerecognitionX() int32`

GetIEzsigntemplatedocumentpagerecognitionX returns the IEzsigntemplatedocumentpagerecognitionX field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagerecognitionXOk

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetIEzsigntemplatedocumentpagerecognitionXOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagerecognitionXOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagerecognitionX

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) SetIEzsigntemplatedocumentpagerecognitionX(v int32)`

SetIEzsigntemplatedocumentpagerecognitionX sets IEzsigntemplatedocumentpagerecognitionX field to given value.

### HasIEzsigntemplatedocumentpagerecognitionX

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) HasIEzsigntemplatedocumentpagerecognitionX() bool`

HasIEzsigntemplatedocumentpagerecognitionX returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagerecognitionY

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetIEzsigntemplatedocumentpagerecognitionY() int32`

GetIEzsigntemplatedocumentpagerecognitionY returns the IEzsigntemplatedocumentpagerecognitionY field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagerecognitionYOk

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetIEzsigntemplatedocumentpagerecognitionYOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagerecognitionYOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagerecognitionY

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) SetIEzsigntemplatedocumentpagerecognitionY(v int32)`

SetIEzsigntemplatedocumentpagerecognitionY sets IEzsigntemplatedocumentpagerecognitionY field to given value.

### HasIEzsigntemplatedocumentpagerecognitionY

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) HasIEzsigntemplatedocumentpagerecognitionY() bool`

HasIEzsigntemplatedocumentpagerecognitionY returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagerecognitionWidth

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetIEzsigntemplatedocumentpagerecognitionWidth() int32`

GetIEzsigntemplatedocumentpagerecognitionWidth returns the IEzsigntemplatedocumentpagerecognitionWidth field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagerecognitionWidthOk

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetIEzsigntemplatedocumentpagerecognitionWidthOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagerecognitionWidthOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagerecognitionWidth

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) SetIEzsigntemplatedocumentpagerecognitionWidth(v int32)`

SetIEzsigntemplatedocumentpagerecognitionWidth sets IEzsigntemplatedocumentpagerecognitionWidth field to given value.

### HasIEzsigntemplatedocumentpagerecognitionWidth

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) HasIEzsigntemplatedocumentpagerecognitionWidth() bool`

HasIEzsigntemplatedocumentpagerecognitionWidth returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagerecognitionHeight

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetIEzsigntemplatedocumentpagerecognitionHeight() int32`

GetIEzsigntemplatedocumentpagerecognitionHeight returns the IEzsigntemplatedocumentpagerecognitionHeight field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagerecognitionHeightOk

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetIEzsigntemplatedocumentpagerecognitionHeightOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagerecognitionHeightOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagerecognitionHeight

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) SetIEzsigntemplatedocumentpagerecognitionHeight(v int32)`

SetIEzsigntemplatedocumentpagerecognitionHeight sets IEzsigntemplatedocumentpagerecognitionHeight field to given value.

### HasIEzsigntemplatedocumentpagerecognitionHeight

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) HasIEzsigntemplatedocumentpagerecognitionHeight() bool`

HasIEzsigntemplatedocumentpagerecognitionHeight returns a boolean if a field has been set.

### GetTEzsigntemplatedocumentpagerecognitionText

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetTEzsigntemplatedocumentpagerecognitionText() string`

GetTEzsigntemplatedocumentpagerecognitionText returns the TEzsigntemplatedocumentpagerecognitionText field if non-nil, zero value otherwise.

### GetTEzsigntemplatedocumentpagerecognitionTextOk

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) GetTEzsigntemplatedocumentpagerecognitionTextOk() (*string, bool)`

GetTEzsigntemplatedocumentpagerecognitionTextOk returns a tuple with the TEzsigntemplatedocumentpagerecognitionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplatedocumentpagerecognitionText

`func (o *EzsigntemplatedocumentpagerecognitionRequestCompound) SetTEzsigntemplatedocumentpagerecognitionText(v string)`

SetTEzsigntemplatedocumentpagerecognitionText sets TEzsigntemplatedocumentpagerecognitionText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


