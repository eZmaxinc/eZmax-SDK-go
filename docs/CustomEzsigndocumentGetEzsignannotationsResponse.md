# CustomEzsigndocumentGetEzsignannotationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 
**SEzsigndocumentName** | **string** | The name of the document that will be presented to Ezsignfoldersignerassociations | 
**AObjEzsignannotation** | [**[]EzsignannotationResponseCompound**](EzsignannotationResponseCompound.md) |  | 

## Methods

### NewCustomEzsigndocumentGetEzsignannotationsResponse

`func NewCustomEzsigndocumentGetEzsignannotationsResponse(pkiEzsigndocumentID int32, sEzsigndocumentName string, aObjEzsignannotation []EzsignannotationResponseCompound, ) *CustomEzsigndocumentGetEzsignannotationsResponse`

NewCustomEzsigndocumentGetEzsignannotationsResponse instantiates a new CustomEzsigndocumentGetEzsignannotationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzsigndocumentGetEzsignannotationsResponseWithDefaults

`func NewCustomEzsigndocumentGetEzsignannotationsResponseWithDefaults() *CustomEzsigndocumentGetEzsignannotationsResponse`

NewCustomEzsigndocumentGetEzsignannotationsResponseWithDefaults instantiates a new CustomEzsigndocumentGetEzsignannotationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigndocumentID

`func (o *CustomEzsigndocumentGetEzsignannotationsResponse) GetPkiEzsigndocumentID() int32`

GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetPkiEzsigndocumentIDOk

`func (o *CustomEzsigndocumentGetEzsignannotationsResponse) GetPkiEzsigndocumentIDOk() (*int32, bool)`

GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigndocumentID

`func (o *CustomEzsigndocumentGetEzsignannotationsResponse) SetPkiEzsigndocumentID(v int32)`

SetPkiEzsigndocumentID sets PkiEzsigndocumentID field to given value.


### GetSEzsigndocumentName

`func (o *CustomEzsigndocumentGetEzsignannotationsResponse) GetSEzsigndocumentName() string`

GetSEzsigndocumentName returns the SEzsigndocumentName field if non-nil, zero value otherwise.

### GetSEzsigndocumentNameOk

`func (o *CustomEzsigndocumentGetEzsignannotationsResponse) GetSEzsigndocumentNameOk() (*string, bool)`

GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentName

`func (o *CustomEzsigndocumentGetEzsignannotationsResponse) SetSEzsigndocumentName(v string)`

SetSEzsigndocumentName sets SEzsigndocumentName field to given value.


### GetAObjEzsignannotation

`func (o *CustomEzsigndocumentGetEzsignannotationsResponse) GetAObjEzsignannotation() []EzsignannotationResponseCompound`

GetAObjEzsignannotation returns the AObjEzsignannotation field if non-nil, zero value otherwise.

### GetAObjEzsignannotationOk

`func (o *CustomEzsigndocumentGetEzsignannotationsResponse) GetAObjEzsignannotationOk() (*[]EzsignannotationResponseCompound, bool)`

GetAObjEzsignannotationOk returns a tuple with the AObjEzsignannotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignannotation

`func (o *CustomEzsigndocumentGetEzsignannotationsResponse) SetAObjEzsignannotation(v []EzsignannotationResponseCompound)`

SetAObjEzsignannotation sets AObjEzsignannotation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


