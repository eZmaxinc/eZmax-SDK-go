# EzsigndocumentGetWordsPositionsV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EGet** | **string** | Specify if you want to retrieve *All* words or specific *Words* from the document. If you specify *Words*, you must send the list of words to search for in *a_sWord*. | 
**BWordCaseSensitive** | **bool** | IF *true*, words will be searched case-sensitive and results will be returned case-sensitive. IF *false*, words will be searched case-insensitive and results will be returned case-insensitive. | 
**ASWord** | Pointer to **[]string** | Array of words to find in the document | [optional] 

## Methods

### NewEzsigndocumentGetWordsPositionsV1Request

`func NewEzsigndocumentGetWordsPositionsV1Request(eGet string, bWordCaseSensitive bool, ) *EzsigndocumentGetWordsPositionsV1Request`

NewEzsigndocumentGetWordsPositionsV1Request instantiates a new EzsigndocumentGetWordsPositionsV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentGetWordsPositionsV1RequestWithDefaults

`func NewEzsigndocumentGetWordsPositionsV1RequestWithDefaults() *EzsigndocumentGetWordsPositionsV1Request`

NewEzsigndocumentGetWordsPositionsV1RequestWithDefaults instantiates a new EzsigndocumentGetWordsPositionsV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEGet

`func (o *EzsigndocumentGetWordsPositionsV1Request) GetEGet() string`

GetEGet returns the EGet field if non-nil, zero value otherwise.

### GetEGetOk

`func (o *EzsigndocumentGetWordsPositionsV1Request) GetEGetOk() (*string, bool)`

GetEGetOk returns a tuple with the EGet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEGet

`func (o *EzsigndocumentGetWordsPositionsV1Request) SetEGet(v string)`

SetEGet sets EGet field to given value.


### GetBWordCaseSensitive

`func (o *EzsigndocumentGetWordsPositionsV1Request) GetBWordCaseSensitive() bool`

GetBWordCaseSensitive returns the BWordCaseSensitive field if non-nil, zero value otherwise.

### GetBWordCaseSensitiveOk

`func (o *EzsigndocumentGetWordsPositionsV1Request) GetBWordCaseSensitiveOk() (*bool, bool)`

GetBWordCaseSensitiveOk returns a tuple with the BWordCaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBWordCaseSensitive

`func (o *EzsigndocumentGetWordsPositionsV1Request) SetBWordCaseSensitive(v bool)`

SetBWordCaseSensitive sets BWordCaseSensitive field to given value.


### GetASWord

`func (o *EzsigndocumentGetWordsPositionsV1Request) GetASWord() []string`

GetASWord returns the ASWord field if non-nil, zero value otherwise.

### GetASWordOk

`func (o *EzsigndocumentGetWordsPositionsV1Request) GetASWordOk() (*[]string, bool)`

GetASWordOk returns a tuple with the ASWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASWord

`func (o *EzsigndocumentGetWordsPositionsV1Request) SetASWord(v []string)`

SetASWord sets ASWord field to given value.

### HasASWord

`func (o *EzsigndocumentGetWordsPositionsV1Request) HasASWord() bool`

HasASWord returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


