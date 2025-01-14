# EzsigntemplatedocumentpagerecognitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatedocumentpagerecognitionID** | **int32** | The unique ID of the Ezsigntemplatedocumentpagerecognition | 
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

### NewEzsigntemplatedocumentpagerecognitionResponse

`func NewEzsigntemplatedocumentpagerecognitionResponse(pkiEzsigntemplatedocumentpagerecognitionID int32, fkiEzsigntemplatedocumentpageID int32, eEzsigntemplatedocumentpagerecognitionOperator FieldEEzsigntemplatedocumentpagerecognitionOperator, eEzsigntemplatedocumentpagerecognitionSection FieldEEzsigntemplatedocumentpagerecognitionSection, tEzsigntemplatedocumentpagerecognitionText string, ) *EzsigntemplatedocumentpagerecognitionResponse`

NewEzsigntemplatedocumentpagerecognitionResponse instantiates a new EzsigntemplatedocumentpagerecognitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatedocumentpagerecognitionResponseWithDefaults

`func NewEzsigntemplatedocumentpagerecognitionResponseWithDefaults() *EzsigntemplatedocumentpagerecognitionResponse`

NewEzsigntemplatedocumentpagerecognitionResponseWithDefaults instantiates a new EzsigntemplatedocumentpagerecognitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatedocumentpagerecognitionID

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetPkiEzsigntemplatedocumentpagerecognitionID() int32`

GetPkiEzsigntemplatedocumentpagerecognitionID returns the PkiEzsigntemplatedocumentpagerecognitionID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatedocumentpagerecognitionIDOk

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetPkiEzsigntemplatedocumentpagerecognitionIDOk() (*int32, bool)`

GetPkiEzsigntemplatedocumentpagerecognitionIDOk returns a tuple with the PkiEzsigntemplatedocumentpagerecognitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatedocumentpagerecognitionID

`func (o *EzsigntemplatedocumentpagerecognitionResponse) SetPkiEzsigntemplatedocumentpagerecognitionID(v int32)`

SetPkiEzsigntemplatedocumentpagerecognitionID sets PkiEzsigntemplatedocumentpagerecognitionID field to given value.


### GetFkiEzsigntemplatedocumentpageID

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetFkiEzsigntemplatedocumentpageID() int32`

GetFkiEzsigntemplatedocumentpageID returns the FkiEzsigntemplatedocumentpageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatedocumentpageIDOk

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetFkiEzsigntemplatedocumentpageIDOk() (*int32, bool)`

GetFkiEzsigntemplatedocumentpageIDOk returns a tuple with the FkiEzsigntemplatedocumentpageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatedocumentpageID

`func (o *EzsigntemplatedocumentpagerecognitionResponse) SetFkiEzsigntemplatedocumentpageID(v int32)`

SetFkiEzsigntemplatedocumentpageID sets FkiEzsigntemplatedocumentpageID field to given value.


### GetEEzsigntemplatedocumentpagerecognitionOperator

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetEEzsigntemplatedocumentpagerecognitionOperator() FieldEEzsigntemplatedocumentpagerecognitionOperator`

GetEEzsigntemplatedocumentpagerecognitionOperator returns the EEzsigntemplatedocumentpagerecognitionOperator field if non-nil, zero value otherwise.

### GetEEzsigntemplatedocumentpagerecognitionOperatorOk

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetEEzsigntemplatedocumentpagerecognitionOperatorOk() (*FieldEEzsigntemplatedocumentpagerecognitionOperator, bool)`

GetEEzsigntemplatedocumentpagerecognitionOperatorOk returns a tuple with the EEzsigntemplatedocumentpagerecognitionOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatedocumentpagerecognitionOperator

`func (o *EzsigntemplatedocumentpagerecognitionResponse) SetEEzsigntemplatedocumentpagerecognitionOperator(v FieldEEzsigntemplatedocumentpagerecognitionOperator)`

SetEEzsigntemplatedocumentpagerecognitionOperator sets EEzsigntemplatedocumentpagerecognitionOperator field to given value.


### GetEEzsigntemplatedocumentpagerecognitionSection

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetEEzsigntemplatedocumentpagerecognitionSection() FieldEEzsigntemplatedocumentpagerecognitionSection`

GetEEzsigntemplatedocumentpagerecognitionSection returns the EEzsigntemplatedocumentpagerecognitionSection field if non-nil, zero value otherwise.

### GetEEzsigntemplatedocumentpagerecognitionSectionOk

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetEEzsigntemplatedocumentpagerecognitionSectionOk() (*FieldEEzsigntemplatedocumentpagerecognitionSection, bool)`

GetEEzsigntemplatedocumentpagerecognitionSectionOk returns a tuple with the EEzsigntemplatedocumentpagerecognitionSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatedocumentpagerecognitionSection

`func (o *EzsigntemplatedocumentpagerecognitionResponse) SetEEzsigntemplatedocumentpagerecognitionSection(v FieldEEzsigntemplatedocumentpagerecognitionSection)`

SetEEzsigntemplatedocumentpagerecognitionSection sets EEzsigntemplatedocumentpagerecognitionSection field to given value.


### GetIEzsigntemplatedocumentpagerecognitionSimilarpercentage

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetIEzsigntemplatedocumentpagerecognitionSimilarpercentage() int32`

GetIEzsigntemplatedocumentpagerecognitionSimilarpercentage returns the IEzsigntemplatedocumentpagerecognitionSimilarpercentage field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagerecognitionSimilarpercentageOk

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetIEzsigntemplatedocumentpagerecognitionSimilarpercentageOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagerecognitionSimilarpercentageOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionSimilarpercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagerecognitionSimilarpercentage

`func (o *EzsigntemplatedocumentpagerecognitionResponse) SetIEzsigntemplatedocumentpagerecognitionSimilarpercentage(v int32)`

SetIEzsigntemplatedocumentpagerecognitionSimilarpercentage sets IEzsigntemplatedocumentpagerecognitionSimilarpercentage field to given value.

### HasIEzsigntemplatedocumentpagerecognitionSimilarpercentage

`func (o *EzsigntemplatedocumentpagerecognitionResponse) HasIEzsigntemplatedocumentpagerecognitionSimilarpercentage() bool`

HasIEzsigntemplatedocumentpagerecognitionSimilarpercentage returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagerecognitionX

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetIEzsigntemplatedocumentpagerecognitionX() int32`

GetIEzsigntemplatedocumentpagerecognitionX returns the IEzsigntemplatedocumentpagerecognitionX field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagerecognitionXOk

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetIEzsigntemplatedocumentpagerecognitionXOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagerecognitionXOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagerecognitionX

`func (o *EzsigntemplatedocumentpagerecognitionResponse) SetIEzsigntemplatedocumentpagerecognitionX(v int32)`

SetIEzsigntemplatedocumentpagerecognitionX sets IEzsigntemplatedocumentpagerecognitionX field to given value.

### HasIEzsigntemplatedocumentpagerecognitionX

`func (o *EzsigntemplatedocumentpagerecognitionResponse) HasIEzsigntemplatedocumentpagerecognitionX() bool`

HasIEzsigntemplatedocumentpagerecognitionX returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagerecognitionY

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetIEzsigntemplatedocumentpagerecognitionY() int32`

GetIEzsigntemplatedocumentpagerecognitionY returns the IEzsigntemplatedocumentpagerecognitionY field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagerecognitionYOk

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetIEzsigntemplatedocumentpagerecognitionYOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagerecognitionYOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagerecognitionY

`func (o *EzsigntemplatedocumentpagerecognitionResponse) SetIEzsigntemplatedocumentpagerecognitionY(v int32)`

SetIEzsigntemplatedocumentpagerecognitionY sets IEzsigntemplatedocumentpagerecognitionY field to given value.

### HasIEzsigntemplatedocumentpagerecognitionY

`func (o *EzsigntemplatedocumentpagerecognitionResponse) HasIEzsigntemplatedocumentpagerecognitionY() bool`

HasIEzsigntemplatedocumentpagerecognitionY returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagerecognitionWidth

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetIEzsigntemplatedocumentpagerecognitionWidth() int32`

GetIEzsigntemplatedocumentpagerecognitionWidth returns the IEzsigntemplatedocumentpagerecognitionWidth field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagerecognitionWidthOk

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetIEzsigntemplatedocumentpagerecognitionWidthOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagerecognitionWidthOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagerecognitionWidth

`func (o *EzsigntemplatedocumentpagerecognitionResponse) SetIEzsigntemplatedocumentpagerecognitionWidth(v int32)`

SetIEzsigntemplatedocumentpagerecognitionWidth sets IEzsigntemplatedocumentpagerecognitionWidth field to given value.

### HasIEzsigntemplatedocumentpagerecognitionWidth

`func (o *EzsigntemplatedocumentpagerecognitionResponse) HasIEzsigntemplatedocumentpagerecognitionWidth() bool`

HasIEzsigntemplatedocumentpagerecognitionWidth returns a boolean if a field has been set.

### GetIEzsigntemplatedocumentpagerecognitionHeight

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetIEzsigntemplatedocumentpagerecognitionHeight() int32`

GetIEzsigntemplatedocumentpagerecognitionHeight returns the IEzsigntemplatedocumentpagerecognitionHeight field if non-nil, zero value otherwise.

### GetIEzsigntemplatedocumentpagerecognitionHeightOk

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetIEzsigntemplatedocumentpagerecognitionHeightOk() (*int32, bool)`

GetIEzsigntemplatedocumentpagerecognitionHeightOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatedocumentpagerecognitionHeight

`func (o *EzsigntemplatedocumentpagerecognitionResponse) SetIEzsigntemplatedocumentpagerecognitionHeight(v int32)`

SetIEzsigntemplatedocumentpagerecognitionHeight sets IEzsigntemplatedocumentpagerecognitionHeight field to given value.

### HasIEzsigntemplatedocumentpagerecognitionHeight

`func (o *EzsigntemplatedocumentpagerecognitionResponse) HasIEzsigntemplatedocumentpagerecognitionHeight() bool`

HasIEzsigntemplatedocumentpagerecognitionHeight returns a boolean if a field has been set.

### GetTEzsigntemplatedocumentpagerecognitionText

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetTEzsigntemplatedocumentpagerecognitionText() string`

GetTEzsigntemplatedocumentpagerecognitionText returns the TEzsigntemplatedocumentpagerecognitionText field if non-nil, zero value otherwise.

### GetTEzsigntemplatedocumentpagerecognitionTextOk

`func (o *EzsigntemplatedocumentpagerecognitionResponse) GetTEzsigntemplatedocumentpagerecognitionTextOk() (*string, bool)`

GetTEzsigntemplatedocumentpagerecognitionTextOk returns a tuple with the TEzsigntemplatedocumentpagerecognitionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsigntemplatedocumentpagerecognitionText

`func (o *EzsigntemplatedocumentpagerecognitionResponse) SetTEzsigntemplatedocumentpagerecognitionText(v string)`

SetTEzsigntemplatedocumentpagerecognitionText sets TEzsigntemplatedocumentpagerecognitionText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


