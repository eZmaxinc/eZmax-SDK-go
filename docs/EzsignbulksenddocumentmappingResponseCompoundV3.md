# EzsignbulksenddocumentmappingResponseCompoundV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignbulksenddocumentmappingID** | **int32** | The unique ID of the Ezsignbulksenddocumentmapping. | 
**FkiEzsignbulksendID** | **int32** | The unique ID of the Ezsignbulksend | 
**FkiEzsigntemplatepackageID** | Pointer to **int32** | The unique ID of the Ezsigntemplatepackage | [optional] 
**FkiEzsigntemplateID** | Pointer to **int32** | The unique ID of the Ezsigntemplate | [optional] 
**IEzsignbulksenddocumentmappingOrder** | **int32** | The order in which the Ezsigntemplate or Ezsigntemplatepackage will be presented to the signatory in the Ezsignfolder. | 
**ObjEzsigntemplate** | Pointer to [**EzsigntemplateResponseCompoundV4**](EzsigntemplateResponseCompoundV4.md) |  | [optional] 
**ObjEzsigntemplatepackage** | Pointer to [**EzsigntemplatepackageResponseCompoundV3**](EzsigntemplatepackageResponseCompoundV3.md) |  | [optional] 

## Methods

### NewEzsignbulksenddocumentmappingResponseCompoundV3

`func NewEzsignbulksenddocumentmappingResponseCompoundV3(pkiEzsignbulksenddocumentmappingID int32, fkiEzsignbulksendID int32, iEzsignbulksenddocumentmappingOrder int32, ) *EzsignbulksenddocumentmappingResponseCompoundV3`

NewEzsignbulksenddocumentmappingResponseCompoundV3 instantiates a new EzsignbulksenddocumentmappingResponseCompoundV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignbulksenddocumentmappingResponseCompoundV3WithDefaults

`func NewEzsignbulksenddocumentmappingResponseCompoundV3WithDefaults() *EzsignbulksenddocumentmappingResponseCompoundV3`

NewEzsignbulksenddocumentmappingResponseCompoundV3WithDefaults instantiates a new EzsignbulksenddocumentmappingResponseCompoundV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignbulksenddocumentmappingID

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) GetPkiEzsignbulksenddocumentmappingID() int32`

GetPkiEzsignbulksenddocumentmappingID returns the PkiEzsignbulksenddocumentmappingID field if non-nil, zero value otherwise.

### GetPkiEzsignbulksenddocumentmappingIDOk

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) GetPkiEzsignbulksenddocumentmappingIDOk() (*int32, bool)`

GetPkiEzsignbulksenddocumentmappingIDOk returns a tuple with the PkiEzsignbulksenddocumentmappingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignbulksenddocumentmappingID

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) SetPkiEzsignbulksenddocumentmappingID(v int32)`

SetPkiEzsignbulksenddocumentmappingID sets PkiEzsignbulksenddocumentmappingID field to given value.


### GetFkiEzsignbulksendID

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) GetFkiEzsignbulksendID() int32`

GetFkiEzsignbulksendID returns the FkiEzsignbulksendID field if non-nil, zero value otherwise.

### GetFkiEzsignbulksendIDOk

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) GetFkiEzsignbulksendIDOk() (*int32, bool)`

GetFkiEzsignbulksendIDOk returns a tuple with the FkiEzsignbulksendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignbulksendID

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) SetFkiEzsignbulksendID(v int32)`

SetFkiEzsignbulksendID sets FkiEzsignbulksendID field to given value.


### GetFkiEzsigntemplatepackageID

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) GetFkiEzsigntemplatepackageID() int32`

GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatepackageIDOk

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) GetFkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatepackageID

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) SetFkiEzsigntemplatepackageID(v int32)`

SetFkiEzsigntemplatepackageID sets FkiEzsigntemplatepackageID field to given value.

### HasFkiEzsigntemplatepackageID

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) HasFkiEzsigntemplatepackageID() bool`

HasFkiEzsigntemplatepackageID returns a boolean if a field has been set.

### GetFkiEzsigntemplateID

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.

### HasFkiEzsigntemplateID

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) HasFkiEzsigntemplateID() bool`

HasFkiEzsigntemplateID returns a boolean if a field has been set.

### GetIEzsignbulksenddocumentmappingOrder

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) GetIEzsignbulksenddocumentmappingOrder() int32`

GetIEzsignbulksenddocumentmappingOrder returns the IEzsignbulksenddocumentmappingOrder field if non-nil, zero value otherwise.

### GetIEzsignbulksenddocumentmappingOrderOk

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) GetIEzsignbulksenddocumentmappingOrderOk() (*int32, bool)`

GetIEzsignbulksenddocumentmappingOrderOk returns a tuple with the IEzsignbulksenddocumentmappingOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignbulksenddocumentmappingOrder

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) SetIEzsignbulksenddocumentmappingOrder(v int32)`

SetIEzsignbulksenddocumentmappingOrder sets IEzsignbulksenddocumentmappingOrder field to given value.


### GetObjEzsigntemplate

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) GetObjEzsigntemplate() EzsigntemplateResponseCompoundV4`

GetObjEzsigntemplate returns the ObjEzsigntemplate field if non-nil, zero value otherwise.

### GetObjEzsigntemplateOk

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) GetObjEzsigntemplateOk() (*EzsigntemplateResponseCompoundV4, bool)`

GetObjEzsigntemplateOk returns a tuple with the ObjEzsigntemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsigntemplate

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) SetObjEzsigntemplate(v EzsigntemplateResponseCompoundV4)`

SetObjEzsigntemplate sets ObjEzsigntemplate field to given value.

### HasObjEzsigntemplate

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) HasObjEzsigntemplate() bool`

HasObjEzsigntemplate returns a boolean if a field has been set.

### GetObjEzsigntemplatepackage

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) GetObjEzsigntemplatepackage() EzsigntemplatepackageResponseCompoundV3`

GetObjEzsigntemplatepackage returns the ObjEzsigntemplatepackage field if non-nil, zero value otherwise.

### GetObjEzsigntemplatepackageOk

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) GetObjEzsigntemplatepackageOk() (*EzsigntemplatepackageResponseCompoundV3, bool)`

GetObjEzsigntemplatepackageOk returns a tuple with the ObjEzsigntemplatepackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsigntemplatepackage

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) SetObjEzsigntemplatepackage(v EzsigntemplatepackageResponseCompoundV3)`

SetObjEzsigntemplatepackage sets ObjEzsigntemplatepackage field to given value.

### HasObjEzsigntemplatepackage

`func (o *EzsignbulksenddocumentmappingResponseCompoundV3) HasObjEzsigntemplatepackage() bool`

HasObjEzsigntemplatepackage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


