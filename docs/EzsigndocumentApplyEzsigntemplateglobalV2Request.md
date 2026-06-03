# EzsigndocumentApplyEzsigntemplateglobalV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiEzsigntemplateglobalID** | **int32** | The unique ID of the Ezsigntemplateglobal | 
**ASEzsigntemplateglobalsigner** | **[]string** |  | 
**AFkiEzsignfoldersignerassociationID** | **[]int32** |  | 
**ASEzsigntemplateglobalannotationDescription** | Pointer to **[]string** |  | [optional] 
**ASEzsigntemplateglobalannotationDefaulttext** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEzsigndocumentApplyEzsigntemplateglobalV2Request

`func NewEzsigndocumentApplyEzsigntemplateglobalV2Request(fkiEzsigntemplateglobalID int32, aSEzsigntemplateglobalsigner []string, aFkiEzsignfoldersignerassociationID []int32, ) *EzsigndocumentApplyEzsigntemplateglobalV2Request`

NewEzsigndocumentApplyEzsigntemplateglobalV2Request instantiates a new EzsigndocumentApplyEzsigntemplateglobalV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentApplyEzsigntemplateglobalV2RequestWithDefaults

`func NewEzsigndocumentApplyEzsigntemplateglobalV2RequestWithDefaults() *EzsigndocumentApplyEzsigntemplateglobalV2Request`

NewEzsigndocumentApplyEzsigntemplateglobalV2RequestWithDefaults instantiates a new EzsigndocumentApplyEzsigntemplateglobalV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiEzsigntemplateglobalID

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) GetFkiEzsigntemplateglobalID() int32`

GetFkiEzsigntemplateglobalID returns the FkiEzsigntemplateglobalID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateglobalIDOk

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) GetFkiEzsigntemplateglobalIDOk() (*int32, bool)`

GetFkiEzsigntemplateglobalIDOk returns a tuple with the FkiEzsigntemplateglobalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateglobalID

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) SetFkiEzsigntemplateglobalID(v int32)`

SetFkiEzsigntemplateglobalID sets FkiEzsigntemplateglobalID field to given value.


### GetASEzsigntemplateglobalsigner

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) GetASEzsigntemplateglobalsigner() []string`

GetASEzsigntemplateglobalsigner returns the ASEzsigntemplateglobalsigner field if non-nil, zero value otherwise.

### GetASEzsigntemplateglobalsignerOk

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) GetASEzsigntemplateglobalsignerOk() (*[]string, bool)`

GetASEzsigntemplateglobalsignerOk returns a tuple with the ASEzsigntemplateglobalsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASEzsigntemplateglobalsigner

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) SetASEzsigntemplateglobalsigner(v []string)`

SetASEzsigntemplateglobalsigner sets ASEzsigntemplateglobalsigner field to given value.


### GetAFkiEzsignfoldersignerassociationID

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) GetAFkiEzsignfoldersignerassociationID() []int32`

GetAFkiEzsignfoldersignerassociationID returns the AFkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetAFkiEzsignfoldersignerassociationIDOk

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) GetAFkiEzsignfoldersignerassociationIDOk() (*[]int32, bool)`

GetAFkiEzsignfoldersignerassociationIDOk returns a tuple with the AFkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiEzsignfoldersignerassociationID

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) SetAFkiEzsignfoldersignerassociationID(v []int32)`

SetAFkiEzsignfoldersignerassociationID sets AFkiEzsignfoldersignerassociationID field to given value.


### GetASEzsigntemplateglobalannotationDescription

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) GetASEzsigntemplateglobalannotationDescription() []string`

GetASEzsigntemplateglobalannotationDescription returns the ASEzsigntemplateglobalannotationDescription field if non-nil, zero value otherwise.

### GetASEzsigntemplateglobalannotationDescriptionOk

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) GetASEzsigntemplateglobalannotationDescriptionOk() (*[]string, bool)`

GetASEzsigntemplateglobalannotationDescriptionOk returns a tuple with the ASEzsigntemplateglobalannotationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASEzsigntemplateglobalannotationDescription

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) SetASEzsigntemplateglobalannotationDescription(v []string)`

SetASEzsigntemplateglobalannotationDescription sets ASEzsigntemplateglobalannotationDescription field to given value.

### HasASEzsigntemplateglobalannotationDescription

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) HasASEzsigntemplateglobalannotationDescription() bool`

HasASEzsigntemplateglobalannotationDescription returns a boolean if a field has been set.

### GetASEzsigntemplateglobalannotationDefaulttext

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) GetASEzsigntemplateglobalannotationDefaulttext() []string`

GetASEzsigntemplateglobalannotationDefaulttext returns the ASEzsigntemplateglobalannotationDefaulttext field if non-nil, zero value otherwise.

### GetASEzsigntemplateglobalannotationDefaulttextOk

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) GetASEzsigntemplateglobalannotationDefaulttextOk() (*[]string, bool)`

GetASEzsigntemplateglobalannotationDefaulttextOk returns a tuple with the ASEzsigntemplateglobalannotationDefaulttext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASEzsigntemplateglobalannotationDefaulttext

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) SetASEzsigntemplateglobalannotationDefaulttext(v []string)`

SetASEzsigntemplateglobalannotationDefaulttext sets ASEzsigntemplateglobalannotationDefaulttext field to given value.

### HasASEzsigntemplateglobalannotationDefaulttext

`func (o *EzsigndocumentApplyEzsigntemplateglobalV2Request) HasASEzsigntemplateglobalannotationDefaulttext() bool`

HasASEzsigntemplateglobalannotationDefaulttext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


