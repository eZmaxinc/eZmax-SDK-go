# EzsignfolderImportEzsigntemplatepackageV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiEzsigntemplatepackageID** | **int32** | The unique ID of the Ezsigntemplatepackage | 
**DtEzsigndocumentDuedate** | **string** | The maximum date and time at which the Ezsigndocument can be signed. | 
**AObjImportEzsigntemplatepackageRelation** | [**[]CustomImportEzsigntemplatepackageRelationRequest**](CustomImportEzsigntemplatepackageRelationRequest.md) |  | 

## Methods

### NewEzsignfolderImportEzsigntemplatepackageV1Request

`func NewEzsignfolderImportEzsigntemplatepackageV1Request(fkiEzsigntemplatepackageID int32, dtEzsigndocumentDuedate string, aObjImportEzsigntemplatepackageRelation []CustomImportEzsigntemplatepackageRelationRequest, ) *EzsignfolderImportEzsigntemplatepackageV1Request`

NewEzsignfolderImportEzsigntemplatepackageV1Request instantiates a new EzsignfolderImportEzsigntemplatepackageV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderImportEzsigntemplatepackageV1RequestWithDefaults

`func NewEzsignfolderImportEzsigntemplatepackageV1RequestWithDefaults() *EzsignfolderImportEzsigntemplatepackageV1Request`

NewEzsignfolderImportEzsigntemplatepackageV1RequestWithDefaults instantiates a new EzsignfolderImportEzsigntemplatepackageV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiEzsigntemplatepackageID

`func (o *EzsignfolderImportEzsigntemplatepackageV1Request) GetFkiEzsigntemplatepackageID() int32`

GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatepackageIDOk

`func (o *EzsignfolderImportEzsigntemplatepackageV1Request) GetFkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatepackageID

`func (o *EzsignfolderImportEzsigntemplatepackageV1Request) SetFkiEzsigntemplatepackageID(v int32)`

SetFkiEzsigntemplatepackageID sets FkiEzsigntemplatepackageID field to given value.


### GetDtEzsigndocumentDuedate

`func (o *EzsignfolderImportEzsigntemplatepackageV1Request) GetDtEzsigndocumentDuedate() string`

GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field if non-nil, zero value otherwise.

### GetDtEzsigndocumentDuedateOk

`func (o *EzsignfolderImportEzsigntemplatepackageV1Request) GetDtEzsigndocumentDuedateOk() (*string, bool)`

GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentDuedate

`func (o *EzsignfolderImportEzsigntemplatepackageV1Request) SetDtEzsigndocumentDuedate(v string)`

SetDtEzsigndocumentDuedate sets DtEzsigndocumentDuedate field to given value.


### GetAObjImportEzsigntemplatepackageRelation

`func (o *EzsignfolderImportEzsigntemplatepackageV1Request) GetAObjImportEzsigntemplatepackageRelation() []CustomImportEzsigntemplatepackageRelationRequest`

GetAObjImportEzsigntemplatepackageRelation returns the AObjImportEzsigntemplatepackageRelation field if non-nil, zero value otherwise.

### GetAObjImportEzsigntemplatepackageRelationOk

`func (o *EzsignfolderImportEzsigntemplatepackageV1Request) GetAObjImportEzsigntemplatepackageRelationOk() (*[]CustomImportEzsigntemplatepackageRelationRequest, bool)`

GetAObjImportEzsigntemplatepackageRelationOk returns a tuple with the AObjImportEzsigntemplatepackageRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjImportEzsigntemplatepackageRelation

`func (o *EzsignfolderImportEzsigntemplatepackageV1Request) SetAObjImportEzsigntemplatepackageRelation(v []CustomImportEzsigntemplatepackageRelationRequest)`

SetAObjImportEzsigntemplatepackageRelation sets AObjImportEzsigntemplatepackageRelation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


