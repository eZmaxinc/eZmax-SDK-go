# EzsignfolderImportEzsigntemplatepackageV3Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiEzsigntemplatepackageID** | **int32** | The unique ID of the Ezsigntemplatepackage | 
**DtEzsigndocumentDuedate** | **string** | The maximum date and time at which the Ezsigndocument can be signed. | 
**AObjImportEzsigntemplatepackageRelation** | [**[]CustomImportEzsigntemplatepackageRelationRequest**](CustomImportEzsigntemplatepackageRelationRequest.md) |  | 
**ASEzsigntemplateannotationDescription** | **[]string** |  | 
**ASEzsigntemplateannotationDefaulttext** | **[]string** |  | 

## Methods

### NewEzsignfolderImportEzsigntemplatepackageV3Request

`func NewEzsignfolderImportEzsigntemplatepackageV3Request(fkiEzsigntemplatepackageID int32, dtEzsigndocumentDuedate string, aObjImportEzsigntemplatepackageRelation []CustomImportEzsigntemplatepackageRelationRequest, aSEzsigntemplateannotationDescription []string, aSEzsigntemplateannotationDefaulttext []string, ) *EzsignfolderImportEzsigntemplatepackageV3Request`

NewEzsignfolderImportEzsigntemplatepackageV3Request instantiates a new EzsignfolderImportEzsigntemplatepackageV3Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderImportEzsigntemplatepackageV3RequestWithDefaults

`func NewEzsignfolderImportEzsigntemplatepackageV3RequestWithDefaults() *EzsignfolderImportEzsigntemplatepackageV3Request`

NewEzsignfolderImportEzsigntemplatepackageV3RequestWithDefaults instantiates a new EzsignfolderImportEzsigntemplatepackageV3Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiEzsigntemplatepackageID

`func (o *EzsignfolderImportEzsigntemplatepackageV3Request) GetFkiEzsigntemplatepackageID() int32`

GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatepackageIDOk

`func (o *EzsignfolderImportEzsigntemplatepackageV3Request) GetFkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatepackageID

`func (o *EzsignfolderImportEzsigntemplatepackageV3Request) SetFkiEzsigntemplatepackageID(v int32)`

SetFkiEzsigntemplatepackageID sets FkiEzsigntemplatepackageID field to given value.


### GetDtEzsigndocumentDuedate

`func (o *EzsignfolderImportEzsigntemplatepackageV3Request) GetDtEzsigndocumentDuedate() string`

GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field if non-nil, zero value otherwise.

### GetDtEzsigndocumentDuedateOk

`func (o *EzsignfolderImportEzsigntemplatepackageV3Request) GetDtEzsigndocumentDuedateOk() (*string, bool)`

GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentDuedate

`func (o *EzsignfolderImportEzsigntemplatepackageV3Request) SetDtEzsigndocumentDuedate(v string)`

SetDtEzsigndocumentDuedate sets DtEzsigndocumentDuedate field to given value.


### GetAObjImportEzsigntemplatepackageRelation

`func (o *EzsignfolderImportEzsigntemplatepackageV3Request) GetAObjImportEzsigntemplatepackageRelation() []CustomImportEzsigntemplatepackageRelationRequest`

GetAObjImportEzsigntemplatepackageRelation returns the AObjImportEzsigntemplatepackageRelation field if non-nil, zero value otherwise.

### GetAObjImportEzsigntemplatepackageRelationOk

`func (o *EzsignfolderImportEzsigntemplatepackageV3Request) GetAObjImportEzsigntemplatepackageRelationOk() (*[]CustomImportEzsigntemplatepackageRelationRequest, bool)`

GetAObjImportEzsigntemplatepackageRelationOk returns a tuple with the AObjImportEzsigntemplatepackageRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjImportEzsigntemplatepackageRelation

`func (o *EzsignfolderImportEzsigntemplatepackageV3Request) SetAObjImportEzsigntemplatepackageRelation(v []CustomImportEzsigntemplatepackageRelationRequest)`

SetAObjImportEzsigntemplatepackageRelation sets AObjImportEzsigntemplatepackageRelation field to given value.


### GetASEzsigntemplateannotationDescription

`func (o *EzsignfolderImportEzsigntemplatepackageV3Request) GetASEzsigntemplateannotationDescription() []string`

GetASEzsigntemplateannotationDescription returns the ASEzsigntemplateannotationDescription field if non-nil, zero value otherwise.

### GetASEzsigntemplateannotationDescriptionOk

`func (o *EzsignfolderImportEzsigntemplatepackageV3Request) GetASEzsigntemplateannotationDescriptionOk() (*[]string, bool)`

GetASEzsigntemplateannotationDescriptionOk returns a tuple with the ASEzsigntemplateannotationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASEzsigntemplateannotationDescription

`func (o *EzsignfolderImportEzsigntemplatepackageV3Request) SetASEzsigntemplateannotationDescription(v []string)`

SetASEzsigntemplateannotationDescription sets ASEzsigntemplateannotationDescription field to given value.


### GetASEzsigntemplateannotationDefaulttext

`func (o *EzsignfolderImportEzsigntemplatepackageV3Request) GetASEzsigntemplateannotationDefaulttext() []string`

GetASEzsigntemplateannotationDefaulttext returns the ASEzsigntemplateannotationDefaulttext field if non-nil, zero value otherwise.

### GetASEzsigntemplateannotationDefaulttextOk

`func (o *EzsignfolderImportEzsigntemplatepackageV3Request) GetASEzsigntemplateannotationDefaulttextOk() (*[]string, bool)`

GetASEzsigntemplateannotationDefaulttextOk returns a tuple with the ASEzsigntemplateannotationDefaulttext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASEzsigntemplateannotationDefaulttext

`func (o *EzsignfolderImportEzsigntemplatepackageV3Request) SetASEzsigntemplateannotationDefaulttext(v []string)`

SetASEzsigntemplateannotationDefaulttext sets ASEzsigntemplateannotationDefaulttext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


