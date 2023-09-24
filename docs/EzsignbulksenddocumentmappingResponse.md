# EzsignbulksenddocumentmappingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignbulksenddocumentmappingID** | **int32** | The unique ID of the Ezsignbulksenddocumentmapping. | 
**FkiEzsignbulksendID** | **int32** | The unique ID of the Ezsignbulksend | 
**FkiEzsigntemplatepackageID** | Pointer to **int32** | The unique ID of the Ezsigntemplatepackage | [optional] 
**FkiEzsigntemplateID** | Pointer to **int32** | The unique ID of the Ezsigntemplate | [optional] 
**IEzsignbulksenddocumentmappingOrder** | **int32** | The order in which the Ezsigntemplate or Ezsigntemplatepackage will be presented to the signatory in the Ezsignfolder. | 

## Methods

### NewEzsignbulksenddocumentmappingResponse

`func NewEzsignbulksenddocumentmappingResponse(pkiEzsignbulksenddocumentmappingID int32, fkiEzsignbulksendID int32, iEzsignbulksenddocumentmappingOrder int32, ) *EzsignbulksenddocumentmappingResponse`

NewEzsignbulksenddocumentmappingResponse instantiates a new EzsignbulksenddocumentmappingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignbulksenddocumentmappingResponseWithDefaults

`func NewEzsignbulksenddocumentmappingResponseWithDefaults() *EzsignbulksenddocumentmappingResponse`

NewEzsignbulksenddocumentmappingResponseWithDefaults instantiates a new EzsignbulksenddocumentmappingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignbulksenddocumentmappingID

`func (o *EzsignbulksenddocumentmappingResponse) GetPkiEzsignbulksenddocumentmappingID() int32`

GetPkiEzsignbulksenddocumentmappingID returns the PkiEzsignbulksenddocumentmappingID field if non-nil, zero value otherwise.

### GetPkiEzsignbulksenddocumentmappingIDOk

`func (o *EzsignbulksenddocumentmappingResponse) GetPkiEzsignbulksenddocumentmappingIDOk() (*int32, bool)`

GetPkiEzsignbulksenddocumentmappingIDOk returns a tuple with the PkiEzsignbulksenddocumentmappingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignbulksenddocumentmappingID

`func (o *EzsignbulksenddocumentmappingResponse) SetPkiEzsignbulksenddocumentmappingID(v int32)`

SetPkiEzsignbulksenddocumentmappingID sets PkiEzsignbulksenddocumentmappingID field to given value.


### GetFkiEzsignbulksendID

`func (o *EzsignbulksenddocumentmappingResponse) GetFkiEzsignbulksendID() int32`

GetFkiEzsignbulksendID returns the FkiEzsignbulksendID field if non-nil, zero value otherwise.

### GetFkiEzsignbulksendIDOk

`func (o *EzsignbulksenddocumentmappingResponse) GetFkiEzsignbulksendIDOk() (*int32, bool)`

GetFkiEzsignbulksendIDOk returns a tuple with the FkiEzsignbulksendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignbulksendID

`func (o *EzsignbulksenddocumentmappingResponse) SetFkiEzsignbulksendID(v int32)`

SetFkiEzsignbulksendID sets FkiEzsignbulksendID field to given value.


### GetFkiEzsigntemplatepackageID

`func (o *EzsignbulksenddocumentmappingResponse) GetFkiEzsigntemplatepackageID() int32`

GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatepackageIDOk

`func (o *EzsignbulksenddocumentmappingResponse) GetFkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatepackageID

`func (o *EzsignbulksenddocumentmappingResponse) SetFkiEzsigntemplatepackageID(v int32)`

SetFkiEzsigntemplatepackageID sets FkiEzsigntemplatepackageID field to given value.

### HasFkiEzsigntemplatepackageID

`func (o *EzsignbulksenddocumentmappingResponse) HasFkiEzsigntemplatepackageID() bool`

HasFkiEzsigntemplatepackageID returns a boolean if a field has been set.

### GetFkiEzsigntemplateID

`func (o *EzsignbulksenddocumentmappingResponse) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsignbulksenddocumentmappingResponse) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsignbulksenddocumentmappingResponse) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.

### HasFkiEzsigntemplateID

`func (o *EzsignbulksenddocumentmappingResponse) HasFkiEzsigntemplateID() bool`

HasFkiEzsigntemplateID returns a boolean if a field has been set.

### GetIEzsignbulksenddocumentmappingOrder

`func (o *EzsignbulksenddocumentmappingResponse) GetIEzsignbulksenddocumentmappingOrder() int32`

GetIEzsignbulksenddocumentmappingOrder returns the IEzsignbulksenddocumentmappingOrder field if non-nil, zero value otherwise.

### GetIEzsignbulksenddocumentmappingOrderOk

`func (o *EzsignbulksenddocumentmappingResponse) GetIEzsignbulksenddocumentmappingOrderOk() (*int32, bool)`

GetIEzsignbulksenddocumentmappingOrderOk returns a tuple with the IEzsignbulksenddocumentmappingOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignbulksenddocumentmappingOrder

`func (o *EzsignbulksenddocumentmappingResponse) SetIEzsignbulksenddocumentmappingOrder(v int32)`

SetIEzsignbulksenddocumentmappingOrder sets IEzsignbulksenddocumentmappingOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


