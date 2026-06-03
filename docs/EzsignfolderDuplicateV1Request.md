# EzsignfolderDuplicateV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SEzsignfolderDescription** | **string** | The description of the Ezsignfolder | 
**AFkiEzsignfoldersignerassociationID** | **[]int32** |  | 
**AObjEzsigndocument** | [**[]CustomEzsigndocumentDuplicateRequest**](CustomEzsigndocumentDuplicateRequest.md) |  | 
**TEzsignfolderNote** | Pointer to **string** | Note about the Ezsignfolder | [optional] 

## Methods

### NewEzsignfolderDuplicateV1Request

`func NewEzsignfolderDuplicateV1Request(sEzsignfolderDescription string, aFkiEzsignfoldersignerassociationID []int32, aObjEzsigndocument []CustomEzsigndocumentDuplicateRequest, ) *EzsignfolderDuplicateV1Request`

NewEzsignfolderDuplicateV1Request instantiates a new EzsignfolderDuplicateV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderDuplicateV1RequestWithDefaults

`func NewEzsignfolderDuplicateV1RequestWithDefaults() *EzsignfolderDuplicateV1Request`

NewEzsignfolderDuplicateV1RequestWithDefaults instantiates a new EzsignfolderDuplicateV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSEzsignfolderDescription

`func (o *EzsignfolderDuplicateV1Request) GetSEzsignfolderDescription() string`

GetSEzsignfolderDescription returns the SEzsignfolderDescription field if non-nil, zero value otherwise.

### GetSEzsignfolderDescriptionOk

`func (o *EzsignfolderDuplicateV1Request) GetSEzsignfolderDescriptionOk() (*string, bool)`

GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfolderDescription

`func (o *EzsignfolderDuplicateV1Request) SetSEzsignfolderDescription(v string)`

SetSEzsignfolderDescription sets SEzsignfolderDescription field to given value.


### GetAFkiEzsignfoldersignerassociationID

`func (o *EzsignfolderDuplicateV1Request) GetAFkiEzsignfoldersignerassociationID() []int32`

GetAFkiEzsignfoldersignerassociationID returns the AFkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetAFkiEzsignfoldersignerassociationIDOk

`func (o *EzsignfolderDuplicateV1Request) GetAFkiEzsignfoldersignerassociationIDOk() (*[]int32, bool)`

GetAFkiEzsignfoldersignerassociationIDOk returns a tuple with the AFkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiEzsignfoldersignerassociationID

`func (o *EzsignfolderDuplicateV1Request) SetAFkiEzsignfoldersignerassociationID(v []int32)`

SetAFkiEzsignfoldersignerassociationID sets AFkiEzsignfoldersignerassociationID field to given value.


### GetAObjEzsigndocument

`func (o *EzsignfolderDuplicateV1Request) GetAObjEzsigndocument() []CustomEzsigndocumentDuplicateRequest`

GetAObjEzsigndocument returns the AObjEzsigndocument field if non-nil, zero value otherwise.

### GetAObjEzsigndocumentOk

`func (o *EzsignfolderDuplicateV1Request) GetAObjEzsigndocumentOk() (*[]CustomEzsigndocumentDuplicateRequest, bool)`

GetAObjEzsigndocumentOk returns a tuple with the AObjEzsigndocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigndocument

`func (o *EzsignfolderDuplicateV1Request) SetAObjEzsigndocument(v []CustomEzsigndocumentDuplicateRequest)`

SetAObjEzsigndocument sets AObjEzsigndocument field to given value.


### GetTEzsignfolderNote

`func (o *EzsignfolderDuplicateV1Request) GetTEzsignfolderNote() string`

GetTEzsignfolderNote returns the TEzsignfolderNote field if non-nil, zero value otherwise.

### GetTEzsignfolderNoteOk

`func (o *EzsignfolderDuplicateV1Request) GetTEzsignfolderNoteOk() (*string, bool)`

GetTEzsignfolderNoteOk returns a tuple with the TEzsignfolderNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfolderNote

`func (o *EzsignfolderDuplicateV1Request) SetTEzsignfolderNote(v string)`

SetTEzsignfolderNote sets TEzsignfolderNote field to given value.

### HasTEzsignfolderNote

`func (o *EzsignfolderDuplicateV1Request) HasTEzsignfolderNote() bool`

HasTEzsignfolderNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


