# EzsigntemplatedocumentpagerecognitionRequest

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

### NewEzsigntemplatedocumentpagerecognitionRequest

`func NewEzsigntemplatedocumentpagerecognitionRequest(fkiEzsigntemplatedocumentpageID int32, eEzsigntemplatedocumentpagerecognitionOperator FieldEEzsigntemplatedocumentpagerecognitionOperator, eEzsigntemplatedocumentpagerecognitionSection FieldEEzsigntemplatedocumentpagerecognitionSection, tEzsigntemplatedocumentpagerecognitionText string, ) *EzsigntemplatedocumentpagerecognitionRequest`

NewEzsigntemplatedocumentpagerecognitionRequest instantiates a new EzsigntemplatedocumentpagerecognitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatedocumentpagerecognitionRequestWithDefaults

`func NewEzsigntemplatedocumentpagerecognitionRequestWithDefaults() *EzsigntemplatedocumentpagerecognitionRequest`

NewEzsigntemplatedocumentpagerecognitionRequestWithDefaults instantiates a new EzsigntemplatedocumentpagerecognitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatedocumentpagerecognitionID

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetPkiEzsigntemplatedocumentpagerecognitionID() int32`

GetPkiEzsigntemplatedocumentpagerecognitionID returns the PkiEzsigntemplatedocumentpagerecognitionID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatedocumentpagerecognitionIDOk

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetPkiEzsigntemplatedocumentpagerecognitionIDOk() (*int32, bool)`

GetPkiEzsigntemplatedocumentpagerecognitionIDOk returns a tuple with the PkiEzsigntemplatedocumentpagerecognitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatedocumentpagerecognitionID

`func (o *EzsigntemplatedocumentpagerecognitionRequest) SetPkiEzsigntemplatedocumentpagerecognitionID(v int32)`

SetPkiEzsigntemplatedocumentpagerecognitionID sets PkiEzsigntemplatedocumentpagerecognitionID field to given value.

### HasPkiEzsigntemplatedocumentpagerecognitionID

`func (o *EzsigntemplatedocumentpagerecognitionRequest) HasPkiEzsigntemplatedocumentpagerecognitionID() bool`

HasPkiEzsigntemplatedocumentpagerecognitionID returns a boolean if a field has been set.

### GetFkiEzsigntemplatedocumentpageID

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetFkiEzsigntemplatedocumentpageID() int32`

GetFkiEzsigntemplatedocumentpageID returns the FkiEzsigntemplatedocumentpageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatedocumentpageIDOk

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetFkiEzsigntemplatedocumentpageIDOk() (*int32, bool)`

GetFkiEzsigntemplatedocumentpageIDOk returns a tuple with the FkiEzsigntemplatedocumentpageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatedocumentpageID

`func (o *EzsigntemplatedocumentpagerecognitionRequest) SetFkiEzsigntemplatedocumentpageID(v int32)`

SetFkiEzsigntemplatedocumentpageID sets FkiEzsigntemplatedocumentpageID field to given value.


### GetEEzsigntemplatedocumentpagerecognitionOperator

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetEEzsigntemplatedocumentpagerecognitionOperator() FieldEEzsigntemplatedocumentpagerecognitionOperator`

GetEEzsigntemplatedocumentpagerecognitionOperator returns the EEzsigntemplatedocumentpagerecognitionOperator field if non-nil, zero value otherwise.

### GetEEzsigntemplatedocumentpagerecognitionOperatorOk

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetEEzsigntemplatedocumentpagerecognitionOperatorOk() (*FieldEEzsigntemplatedocumentpagerecognitionOperator, bool)`

GetEEzsigntemplatedocumentpagerecognitionOperatorOk returns a tuple with the EEzsigntemplatedocumentpagerecognitionOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatedocumentpagerecognitionOperator

`func (o *EzsigntemplatedocumentpagerecognitionRequest) SetEEzsigntemplatedocumentpagerecognitionOperator(v FieldEEzsigntemplatedocumentpagerecognitionOperator)`

SetEEzsigntemplatedocumentpagerecognitionOperator sets EEzsigntemplatedocumentpagerecognitionOperator field to given value.


### GetEEzsigntemplatedocumentpagerecognitionSection

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetEEzsigntemplatedocumentpagerecognitionSection() FieldEEzsigntemplatedocumentpagerecognitionSection`

GetEEzsigntemplatedocumentpagerecognitionSection returns the EEzsigntemplatedocumentpagerecognitionSection field if non-nil, zero value otherwise.

### GetEEzsigntemplatedocumentpagerecognitionSectionOk

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetEEzsigntemplatedocumentpagerecognitionSectionOk() (*FieldEEzsigntemplatedocumentpagerecognitionSection, bool)`

GetEEzsigntemplatedocumentpagerecognitionSectionOk returns a tuple with the EEzsigntemplatedocumentpagerecognitionSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatedocumentpagerecognitionSection

`func (o *EzsigntemplatedocumentpagerecognitionRequest) SetEEzsigntemplatedocumentpagerecognitionSection(v FieldEEzsigntemplatedocumentpagerecognitionSection)`

SetEEzsigntemplatedocumentpagerecognitionSection sets EEzsigntemplatedocumentpagerecognitionSection field to given value.


### GetIEzsigntemplatedocumentpagerecognitionSimilarpercentage

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionSimilarpercentage() int32`

GetIEzsigntemplatedocumentpagerecognitionSimilarpercentage returns the IEzsigntemplatedocumentpagerecognitionSimilarpercentage field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagerecognitionSimilarpercentageOk

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionSimilarpercentageOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagerecognitionSimilarpercentageOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionSimilarpercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagerecognitionSimilarpercentage

`func (o *EzsigntemplatedocumentpagerecognitionRequest) SetIEzsigntemplatedocumentpagerecognitionSimilarpercentage(v int32)`

SetIEzsigntemplatedocumentpagerecognitionSimilarpercentage sets IEzsigntemplatedocumentpagerecognitionSimilarpercentage field to given value.

### HasIEzsigntemplatedocumentpagerecognitionSimilarpercentage

`func (o *EzsigntemplatedocumentpagerecognitionRequest) HasIEzsigntemplatedocumentpagerecognitionSimilarpercentage() bool`

HasIEzsigntemplatedocumentpagerecognitionSimilarpercentage returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagerecognitionX

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionX() int32`

GetIEzsigntemplatedocumentpagerecognitionX returns the IEzsigntemplatedocumentpagerecognitionX field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagerecognitionXOk

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionXOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagerecognitionXOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagerecognitionX

`func (o *EzsigntemplatedocumentpagerecognitionRequest) SetIEzsigntemplatedocumentpagerecognitionX(v int32)`

SetIEzsigntemplatedocumentpagerecognitionX sets IEzsigntemplatedocumentpagerecognitionX field to given value.

### HasIEzsigntemplatedocumentpagerecognitionX

`func (o *EzsigntemplatedocumentpagerecognitionRequest) HasIEzsigntemplatedocumentpagerecognitionX() bool`

HasIEzsigntemplatedocumentpagerecognitionX returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagerecognitionY

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionY() int32`

GetIEzsigntemplatedocumentpagerecognitionY returns the IEzsigntemplatedocumentpagerecognitionY field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagerecognitionYOk

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionYOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagerecognitionYOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagerecognitionY

`func (o *EzsigntemplatedocumentpagerecognitionRequest) SetIEzsigntemplatedocumentpagerecognitionY(v int32)`

SetIEzsigntemplatedocumentpagerecognitionY sets IEzsigntemplatedocumentpagerecognitionY field to given value.

### HasIEzsigntemplatedocumentpagerecognitionY

`func (o *EzsigntemplatedocumentpagerecognitionRequest) HasIEzsigntemplatedocumentpagerecognitionY() bool`

HasIEzsigntemplatedocumentpagerecognitionY returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagerecognitionWidth

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionWidth() int32`

GetIEzsigntemplatedocumentpagerecognitionWidth returns the IEzsigntemplatedocumentpagerecognitionWidth field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagerecognitionWidthOk

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionWidthOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagerecognitionWidthOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagerecognitionWidth

`func (o *EzsigntemplatedocumentpagerecognitionRequest) SetIEzsigntemplatedocumentpagerecognitionWidth(v int32)`

SetIEzsigntemplatedocumentpagerecognitionWidth sets IEzsigntemplatedocumentpagerecognitionWidth field to given value.

### HasIEzsigntemplatedocumentpagerecognitionWidth

`func (o *EzsigntemplatedocumentpagerecognitionRequest) HasIEzsigntemplatedocumentpagerecognitionWidth() bool`

HasIEzsigntemplatedocumentpagerecognitionWidth returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagerecognitionHeight

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionHeight() int32`

GetIEzsigntemplatedocumentpagerecognitionHeight returns the IEzsigntemplatedocumentpagerecognitionHeight field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagerecognitionHeightOk

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionHeightOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagerecognitionHeightOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagerecognitionHeight

`func (o *EzsigntemplatedocumentpagerecognitionRequest) SetIEzsigntemplatedocumentpagerecognitionHeight(v int32)`

SetIEzsigntemplatedocumentpagerecognitionHeight sets IEzsigntemplatedocumentpagerecognitionHeight field to given value.

### HasIEzsigntemplatedocumentpagerecognitionHeight

`func (o *EzsigntemplatedocumentpagerecognitionRequest) HasIEzsigntemplatedocumentpagerecognitionHeight() bool`

HasIEzsigntemplatedocumentpagerecognitionHeight returns a boolean if a field has been set.

### GetTEzsigntemplatedocumentpagerecognitionText

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetTEzsigntemplatedocumentpagerecognitionText() string`

GetTEzsigntemplatedocumentpagerecognitionText returns the TEzsigntemplatedocumentpagerecognitionText field if non-nil, zero value otherwise.

### GetTEzsigntemplatedocumentpagerecognitionTextOk

`func (o *EzsigntemplatedocumentpagerecognitionRequest) GetTEzsigntemplatedocumentpagerecognitionTextOk() (*string, bool)`

GetTEzsigntemplatedocumentpagerecognitionTextOk returns a tuple with the TEzsigntemplatedocumentpagerecognitionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplatedocumentpagerecognitionText

`func (o *EzsigntemplatedocumentpagerecognitionRequest) SetTEzsigntemplatedocumentpagerecognitionText(v string)`

SetTEzsigntemplatedocumentpagerecognitionText sets TEzsigntemplatedocumentpagerecognitionText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


