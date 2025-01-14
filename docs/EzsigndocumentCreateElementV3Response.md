# EzsigndocumentCreateElementV3Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 
**AObjMatchingtemplate** | [**[]EzsigndocumentMatchingtemplateV3Response**](EzsigndocumentMatchingtemplateV3Response.md) | An array of possibly matching template. | 

## Methods

### NewEzsigndocumentCreateElementV3Response

`func NewEzsigndocumentCreateElementV3Response(pkiEzsigndocumentID int32, aObjMatchingtemplate []EzsigndocumentMatchingtemplateV3Response, ) *EzsigndocumentCreateElementV3Response`

NewEzsigndocumentCreateElementV3Response instantiates a new EzsigndocumentCreateElementV3Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentCreateElementV3ResponseWithDefaults

`func NewEzsigndocumentCreateElementV3ResponseWithDefaults() *EzsigndocumentCreateElementV3Response`

NewEzsigndocumentCreateElementV3ResponseWithDefaults instantiates a new EzsigndocumentCreateElementV3Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigndocumentID

`func (o *EzsigndocumentCreateElementV3Response) GetPkiEzsigndocumentID() int32`

GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetPkiEzsigndocumentIDOk

`func (o *EzsigndocumentCreateElementV3Response) GetPkiEzsigndocumentIDOk() (*int32, bool)`

GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigndocumentID

`func (o *EzsigndocumentCreateElementV3Response) SetPkiEzsigndocumentID(v int32)`

SetPkiEzsigndocumentID sets PkiEzsigndocumentID field to given value.


### GetAObjMatchingtemplate

`func (o *EzsigndocumentCreateElementV3Response) GetAObjMatchingtemplate() []EzsigndocumentMatchingtemplateV3Response`

GetAObjMatchingtemplate returns the AObjMatchingtemplate field if non-nil, zero value otherwise.

### GetAObjMatchingtemplateOk

`func (o *EzsigndocumentCreateElementV3Response) GetAObjMatchingtemplateOk() (*[]EzsigndocumentMatchingtemplateV3Response, bool)`

GetAObjMatchingtemplateOk returns a tuple with the AObjMatchingtemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjMatchingtemplate

`func (o *EzsigndocumentCreateElementV3Response) SetAObjMatchingtemplate(v []EzsigndocumentMatchingtemplateV3Response)`

SetAObjMatchingtemplate sets AObjMatchingtemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


