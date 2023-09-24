# EzsignbulksenddocumentmappingResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignbulksenddocumentmappingID** | **int32** | The unique ID of the Ezsignbulksenddocumentmapping. | 
**FkiEzsignbulksendID** | **int32** | The unique ID of the Ezsignbulksend | 
**FkiEzsigntemplatepackageID** | Pointer to **int32** | The unique ID of the Ezsigntemplatepackage | [optional] 
**FkiEzsigntemplateID** | Pointer to **int32** | The unique ID of the Ezsigntemplate | [optional] 
**IEzsignbulksenddocumentmappingOrder** | **int32** | The order in which the Ezsigntemplate or Ezsigntemplatepackage will be presented to the signatory in the Ezsignfolder. | 
**ObjEzsigntemplate** | Pointer to [**EzsigntemplateResponseCompound**](EzsigntemplateResponseCompound.md) |  | [optional] 
**ObjEzsigntemplatepackage** | Pointer to [**EzsigntemplatepackageResponseCompound**](EzsigntemplatepackageResponseCompound.md) |  | [optional] 

## Methods

### NewEzsignbulksenddocumentmappingResponseCompound

`func NewEzsignbulksenddocumentmappingResponseCompound(pkiEzsignbulksenddocumentmappingID int32, fkiEzsignbulksendID int32, iEzsignbulksenddocumentmappingOrder int32, ) *EzsignbulksenddocumentmappingResponseCompound`

NewEzsignbulksenddocumentmappingResponseCompound instantiates a new EzsignbulksenddocumentmappingResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignbulksenddocumentmappingResponseCompoundWithDefaults

`func NewEzsignbulksenddocumentmappingResponseCompoundWithDefaults() *EzsignbulksenddocumentmappingResponseCompound`

NewEzsignbulksenddocumentmappingResponseCompoundWithDefaults instantiates a new EzsignbulksenddocumentmappingResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignbulksenddocumentmappingID

`func (o *EzsignbulksenddocumentmappingResponseCompound) GetPkiEzsignbulksenddocumentmappingID() int32`

GetPkiEzsignbulksenddocumentmappingID returns the PkiEzsignbulksenddocumentmappingID field if non-nil, zero value otherwise.

### GetPkiEzsignbulksenddocumentmappingIDOk

`func (o *EzsignbulksenddocumentmappingResponseCompound) GetPkiEzsignbulksenddocumentmappingIDOk() (*int32, bool)`

GetPkiEzsignbulksenddocumentmappingIDOk returns a tuple with the PkiEzsignbulksenddocumentmappingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignbulksenddocumentmappingID

`func (o *EzsignbulksenddocumentmappingResponseCompound) SetPkiEzsignbulksenddocumentmappingID(v int32)`

SetPkiEzsignbulksenddocumentmappingID sets PkiEzsignbulksenddocumentmappingID field to given value.


### GetFkiEzsignbulksendID

`func (o *EzsignbulksenddocumentmappingResponseCompound) GetFkiEzsignbulksendID() int32`

GetFkiEzsignbulksendID returns the FkiEzsignbulksendID field if non-nil, zero value otherwise.

### GetFkiEzsignbulksendIDOk

`func (o *EzsignbulksenddocumentmappingResponseCompound) GetFkiEzsignbulksendIDOk() (*int32, bool)`

GetFkiEzsignbulksendIDOk returns a tuple with the FkiEzsignbulksendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignbulksendID

`func (o *EzsignbulksenddocumentmappingResponseCompound) SetFkiEzsignbulksendID(v int32)`

SetFkiEzsignbulksendID sets FkiEzsignbulksendID field to given value.


### GetFkiEzsigntemplatepackageID

`func (o *EzsignbulksenddocumentmappingResponseCompound) GetFkiEzsigntemplatepackageID() int32`

GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatepackageIDOk

`func (o *EzsignbulksenddocumentmappingResponseCompound) GetFkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatepackageID

`func (o *EzsignbulksenddocumentmappingResponseCompound) SetFkiEzsigntemplatepackageID(v int32)`

SetFkiEzsigntemplatepackageID sets FkiEzsigntemplatepackageID field to given value.

### HasFkiEzsigntemplatepackageID

`func (o *EzsignbulksenddocumentmappingResponseCompound) HasFkiEzsigntemplatepackageID() bool`

HasFkiEzsigntemplatepackageID returns a boolean if a field has been set.

### GetFkiEzsigntemplateID

`func (o *EzsignbulksenddocumentmappingResponseCompound) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsignbulksenddocumentmappingResponseCompound) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsignbulksenddocumentmappingResponseCompound) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.

### HasFkiEzsigntemplateID

`func (o *EzsignbulksenddocumentmappingResponseCompound) HasFkiEzsigntemplateID() bool`

HasFkiEzsigntemplateID returns a boolean if a field has been set.

### GetIEzsignbulksenddocumentmappingOrder

`func (o *EzsignbulksenddocumentmappingResponseCompound) GetIEzsignbulksenddocumentmappingOrder() int32`

GetIEzsignbulksenddocumentmappingOrder returns the IEzsignbulksenddocumentmappingOrder field if non-nil, zero value otherwise.

### GetIEzsignbulksenddocumentmappingOrderOk

`func (o *EzsignbulksenddocumentmappingResponseCompound) GetIEzsignbulksenddocumentmappingOrderOk() (*int32, bool)`

GetIEzsignbulksenddocumentmappingOrderOk returns a tuple with the IEzsignbulksenddocumentmappingOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignbulksenddocumentmappingOrder

`func (o *EzsignbulksenddocumentmappingResponseCompound) SetIEzsignbulksenddocumentmappingOrder(v int32)`

SetIEzsignbulksenddocumentmappingOrder sets IEzsignbulksenddocumentmappingOrder field to given value.


### GetObjEzsigntemplate

`func (o *EzsignbulksenddocumentmappingResponseCompound) GetObjEzsigntemplate() EzsigntemplateResponseCompound`

GetObjEzsigntemplate returns the ObjEzsigntemplate field if non-nil, zero value otherwise.

### GetObjEzsigntemplateOk

`func (o *EzsignbulksenddocumentmappingResponseCompound) GetObjEzsigntemplateOk() (*EzsigntemplateResponseCompound, bool)`

GetObjEzsigntemplateOk returns a tuple with the ObjEzsigntemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsigntemplate

`func (o *EzsignbulksenddocumentmappingResponseCompound) SetObjEzsigntemplate(v EzsigntemplateResponseCompound)`

SetObjEzsigntemplate sets ObjEzsigntemplate field to given value.

### HasObjEzsigntemplate

`func (o *EzsignbulksenddocumentmappingResponseCompound) HasObjEzsigntemplate() bool`

HasObjEzsigntemplate returns a boolean if a field has been set.

### GetObjEzsigntemplatepackage

`func (o *EzsignbulksenddocumentmappingResponseCompound) GetObjEzsigntemplatepackage() EzsigntemplatepackageResponseCompound`

GetObjEzsigntemplatepackage returns the ObjEzsigntemplatepackage field if non-nil, zero value otherwise.

### GetObjEzsigntemplatepackageOk

`func (o *EzsignbulksenddocumentmappingResponseCompound) GetObjEzsigntemplatepackageOk() (*EzsigntemplatepackageResponseCompound, bool)`

GetObjEzsigntemplatepackageOk returns a tuple with the ObjEzsigntemplatepackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsigntemplatepackage

`func (o *EzsignbulksenddocumentmappingResponseCompound) SetObjEzsigntemplatepackage(v EzsigntemplatepackageResponseCompound)`

SetObjEzsigntemplatepackage sets ObjEzsigntemplatepackage field to given value.

### HasObjEzsigntemplatepackage

`func (o *EzsignbulksenddocumentmappingResponseCompound) HasObjEzsigntemplatepackage() bool`

HasObjEzsigntemplatepackage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


