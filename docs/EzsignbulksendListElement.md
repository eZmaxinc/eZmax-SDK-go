# EzsignbulksendListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignbulksendID** | **int32** | The unique ID of the Ezsignbulksend | 
**FkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**SEzsignbulksendDescription** | **string** | The description of the Ezsignbulksend | 
**SEzsignfoldertypeNameX** | **string** | The name of the Ezsignfoldertype in the language of the requester | 
**BEzsignbulksendNeedvalidation** | **bool** | Whether the Ezsigntemplatepackage was automatically modified and needs a manual validation | 
**IEzsignbulksendtransmission** | **int32** | The total number of Ezsignbulksendtransmissions in the Ezsignbulksend | 
**IEzsignfolder** | **int32** | The total number of Ezsignfolders in the Ezsignbulksend | 
**IEzsigndocument** | **int32** | The total number of Ezsigndocuments in the Ezsignbulksend | 
**IEzsignsignature** | **int32** | The total number of Ezsignsignature in the Ezsignbulksend | 
**IEzsignsignatureSigned** | **int32** | The total number of already signed Ezsignsignature blocks in the Ezsignbulksend | 
**BEzsignbulksendIsactive** | **bool** | Whether the Ezsignbulksend is active or not | 

## Methods

### NewEzsignbulksendListElement

`func NewEzsignbulksendListElement(pkiEzsignbulksendID int32, fkiEzsignfoldertypeID int32, sEzsignbulksendDescription string, sEzsignfoldertypeNameX string, bEzsignbulksendNeedvalidation bool, iEzsignbulksendtransmission int32, iEzsignfolder int32, iEzsigndocument int32, iEzsignsignature int32, iEzsignsignatureSigned int32, bEzsignbulksendIsactive bool, ) *EzsignbulksendListElement`

NewEzsignbulksendListElement instantiates a new EzsignbulksendListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignbulksendListElementWithDefaults

`func NewEzsignbulksendListElementWithDefaults() *EzsignbulksendListElement`

NewEzsignbulksendListElementWithDefaults instantiates a new EzsignbulksendListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignbulksendID

`func (o *EzsignbulksendListElement) GetPkiEzsignbulksendID() int32`

GetPkiEzsignbulksendID returns the PkiEzsignbulksendID field if non-nil, zero value otherwise.

### GetPkiEzsignbulksendIDOk

`func (o *EzsignbulksendListElement) GetPkiEzsignbulksendIDOk() (*int32, bool)`

GetPkiEzsignbulksendIDOk returns a tuple with the PkiEzsignbulksendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignbulksendID

`func (o *EzsignbulksendListElement) SetPkiEzsignbulksendID(v int32)`

SetPkiEzsignbulksendID sets PkiEzsignbulksendID field to given value.


### GetFkiEzsignfoldertypeID

`func (o *EzsignbulksendListElement) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsignbulksendListElement) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsignbulksendListElement) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### GetSEzsignbulksendDescription

`func (o *EzsignbulksendListElement) GetSEzsignbulksendDescription() string`

GetSEzsignbulksendDescription returns the SEzsignbulksendDescription field if non-nil, zero value otherwise.

### GetSEzsignbulksendDescriptionOk

`func (o *EzsignbulksendListElement) GetSEzsignbulksendDescriptionOk() (*string, bool)`

GetSEzsignbulksendDescriptionOk returns a tuple with the SEzsignbulksendDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignbulksendDescription

`func (o *EzsignbulksendListElement) SetSEzsignbulksendDescription(v string)`

SetSEzsignbulksendDescription sets SEzsignbulksendDescription field to given value.


### GetSEzsignfoldertypeNameX

`func (o *EzsignbulksendListElement) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *EzsignbulksendListElement) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *EzsignbulksendListElement) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.


### GetBEzsignbulksendNeedvalidation

`func (o *EzsignbulksendListElement) GetBEzsignbulksendNeedvalidation() bool`

GetBEzsignbulksendNeedvalidation returns the BEzsignbulksendNeedvalidation field if non-nil, zero value otherwise.

### GetBEzsignbulksendNeedvalidationOk

`func (o *EzsignbulksendListElement) GetBEzsignbulksendNeedvalidationOk() (*bool, bool)`

GetBEzsignbulksendNeedvalidationOk returns a tuple with the BEzsignbulksendNeedvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignbulksendNeedvalidation

`func (o *EzsignbulksendListElement) SetBEzsignbulksendNeedvalidation(v bool)`

SetBEzsignbulksendNeedvalidation sets BEzsignbulksendNeedvalidation field to given value.


### GetIEzsignbulksendtransmission

`func (o *EzsignbulksendListElement) GetIEzsignbulksendtransmission() int32`

GetIEzsignbulksendtransmission returns the IEzsignbulksendtransmission field if non-nil, zero value otherwise.

### GetIEzsignbulksendtransmissionOk

`func (o *EzsignbulksendListElement) GetIEzsignbulksendtransmissionOk() (*int32, bool)`

GetIEzsignbulksendtransmissionOk returns a tuple with the IEzsignbulksendtransmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignbulksendtransmission

`func (o *EzsignbulksendListElement) SetIEzsignbulksendtransmission(v int32)`

SetIEzsignbulksendtransmission sets IEzsignbulksendtransmission field to given value.


### GetIEzsignfolder

`func (o *EzsignbulksendListElement) GetIEzsignfolder() int32`

GetIEzsignfolder returns the IEzsignfolder field if non-nil, zero value otherwise.

### GetIEzsignfolderOk

`func (o *EzsignbulksendListElement) GetIEzsignfolderOk() (*int32, bool)`

GetIEzsignfolderOk returns a tuple with the IEzsignfolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfolder

`func (o *EzsignbulksendListElement) SetIEzsignfolder(v int32)`

SetIEzsignfolder sets IEzsignfolder field to given value.


### GetIEzsigndocument

`func (o *EzsignbulksendListElement) GetIEzsigndocument() int32`

GetIEzsigndocument returns the IEzsigndocument field if non-nil, zero value otherwise.

### GetIEzsigndocumentOk

`func (o *EzsignbulksendListElement) GetIEzsigndocumentOk() (*int32, bool)`

GetIEzsigndocumentOk returns a tuple with the IEzsigndocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndocument

`func (o *EzsignbulksendListElement) SetIEzsigndocument(v int32)`

SetIEzsigndocument sets IEzsigndocument field to given value.


### GetIEzsignsignature

`func (o *EzsignbulksendListElement) GetIEzsignsignature() int32`

GetIEzsignsignature returns the IEzsignsignature field if non-nil, zero value otherwise.

### GetIEzsignsignatureOk

`func (o *EzsignbulksendListElement) GetIEzsignsignatureOk() (*int32, bool)`

GetIEzsignsignatureOk returns a tuple with the IEzsignsignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignature

`func (o *EzsignbulksendListElement) SetIEzsignsignature(v int32)`

SetIEzsignsignature sets IEzsignsignature field to given value.


### GetIEzsignsignatureSigned

`func (o *EzsignbulksendListElement) GetIEzsignsignatureSigned() int32`

GetIEzsignsignatureSigned returns the IEzsignsignatureSigned field if non-nil, zero value otherwise.

### GetIEzsignsignatureSignedOk

`func (o *EzsignbulksendListElement) GetIEzsignsignatureSignedOk() (*int32, bool)`

GetIEzsignsignatureSignedOk returns a tuple with the IEzsignsignatureSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureSigned

`func (o *EzsignbulksendListElement) SetIEzsignsignatureSigned(v int32)`

SetIEzsignsignatureSigned sets IEzsignsignatureSigned field to given value.


### GetBEzsignbulksendIsactive

`func (o *EzsignbulksendListElement) GetBEzsignbulksendIsactive() bool`

GetBEzsignbulksendIsactive returns the BEzsignbulksendIsactive field if non-nil, zero value otherwise.

### GetBEzsignbulksendIsactiveOk

`func (o *EzsignbulksendListElement) GetBEzsignbulksendIsactiveOk() (*bool, bool)`

GetBEzsignbulksendIsactiveOk returns a tuple with the BEzsignbulksendIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignbulksendIsactive

`func (o *EzsignbulksendListElement) SetBEzsignbulksendIsactive(v bool)`

SetBEzsignbulksendIsactive sets BEzsignbulksendIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


