# EzsignfolderImportEzsigntemplatepackageV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiEzsigntemplatepackageID** | **int32** | The unique ID of the Ezsigntemplatepackage | 
**DtEzsigndocumentDuedate** | **string** | The maximum date and time at which the Ezsigndocument can be signed. | 
**AObjImportEzsigntemplatepackageRelation** | [**[]CustomImportEzsigntemplatepackageRelationRequest**](CustomImportEzsigntemplatepackageRelationRequest.md) |  | 

## Methods

### NewEzsignfolderImportEzsigntemplatepackageV2Request

`func NewEzsignfolderImportEzsigntemplatepackageV2Request(fkiEzsigntemplatepackageID int32, dtEzsigndocumentDuedate string, aObjImportEzsigntemplatepackageRelation []CustomImportEzsigntemplatepackageRelationRequest, ) *EzsignfolderImportEzsigntemplatepackageV2Request`

NewEzsignfolderImportEzsigntemplatepackageV2Request instantiates a new EzsignfolderImportEzsigntemplatepackageV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderImportEzsigntemplatepackageV2RequestWithDefaults

`func NewEzsignfolderImportEzsigntemplatepackageV2RequestWithDefaults() *EzsignfolderImportEzsigntemplatepackageV2Request`

NewEzsignfolderImportEzsigntemplatepackageV2RequestWithDefaults instantiates a new EzsignfolderImportEzsigntemplatepackageV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiEzsigntemplatepackageID

`func (o *EzsignfolderImportEzsigntemplatepackageV2Request) GetFkiEzsigntemplatepackageID() int32`

GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatepackageIDOk

`func (o *EzsignfolderImportEzsigntemplatepackageV2Request) GetFkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatepackageID

`func (o *EzsignfolderImportEzsigntemplatepackageV2Request) SetFkiEzsigntemplatepackageID(v int32)`

SetFkiEzsigntemplatepackageID sets FkiEzsigntemplatepackageID field to given value.


### GetDtEzsigndocumentDuedate

`func (o *EzsignfolderImportEzsigntemplatepackageV2Request) GetDtEzsigndocumentDuedate() string`

GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field if non-nil, zero value otherwise.

### GetDtEzsigndocumentDuedateOk

`func (o *EzsignfolderImportEzsigntemplatepackageV2Request) GetDtEzsigndocumentDuedateOk() (*string, bool)`

GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentDuedate

`func (o *EzsignfolderImportEzsigntemplatepackageV2Request) SetDtEzsigndocumentDuedate(v string)`

SetDtEzsigndocumentDuedate sets DtEzsigndocumentDuedate field to given value.


### GetAObjImportEzsigntemplatepackageRelation

`func (o *EzsignfolderImportEzsigntemplatepackageV2Request) GetAObjImportEzsigntemplatepackageRelation() []CustomImportEzsigntemplatepackageRelationRequest`

GetAObjImportEzsigntemplatepackageRelation returns the AObjImportEzsigntemplatepackageRelation field if non-nil, zero value otherwise.

### GetAObjImportEzsigntemplatepackageRelationOk

`func (o *EzsignfolderImportEzsigntemplatepackageV2Request) GetAObjImportEzsigntemplatepackageRelationOk() (*[]CustomImportEzsigntemplatepackageRelationRequest, bool)`

GetAObjImportEzsigntemplatepackageRelationOk returns a tuple with the AObjImportEzsigntemplatepackageRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjImportEzsigntemplatepackageRelation

`func (o *EzsignfolderImportEzsigntemplatepackageV2Request) SetAObjImportEzsigntemplatepackageRelation(v []CustomImportEzsigntemplatepackageRelationRequest)`

SetAObjImportEzsigntemplatepackageRelation sets AObjImportEzsigntemplatepackageRelation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


