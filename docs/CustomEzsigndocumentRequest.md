# CustomEzsigndocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 
**AObjEzsigndocumentdependency** | [**[]EzsigndocumentdependencyRequest**](EzsigndocumentdependencyRequest.md) |  | 

## Methods

### NewCustomEzsigndocumentRequest

`func NewCustomEzsigndocumentRequest(pkiEzsigndocumentID int32, aObjEzsigndocumentdependency []EzsigndocumentdependencyRequestCompound, ) *CustomEzsigndocumentRequest`

NewCustomEzsigndocumentRequest instantiates a new CustomEzsigndocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzsigndocumentRequestWithDefaults

`func NewCustomEzsigndocumentRequestWithDefaults() *CustomEzsigndocumentRequest`

NewCustomEzsigndocumentRequestWithDefaults instantiates a new CustomEzsigndocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigndocumentID

`func (o *CustomEzsigndocumentRequest) GetPkiEzsigndocumentID() int32`

GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetPkiEzsigndocumentIDOk

`func (o *CustomEzsigndocumentRequest) GetPkiEzsigndocumentIDOk() (*int32, bool)`

GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigndocumentID

`func (o *CustomEzsigndocumentRequest) SetPkiEzsigndocumentID(v int32)`

SetPkiEzsigndocumentID sets PkiEzsigndocumentID field to given value.


### GetAObjEzsigndocumentdependency

`func (o *CustomEzsigndocumentRequest) GetAObjEzsigndocumentdependency() []EzsigndocumentdependencyRequestCompound`

GetAObjEzsigndocumentdependency returns the AObjEzsigndocumentdependency field if non-nil, zero value otherwise.

### GetAObjEzsigndocumentdependencyOk

`func (o *CustomEzsigndocumentRequest) GetAObjEzsigndocumentdependencyOk() (*[]EzsigndocumentdependencyRequestCompound, bool)`

GetAObjEzsigndocumentdependencyOk returns a tuple with the AObjEzsigndocumentdependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigndocumentdependency

`func (o *CustomEzsigndocumentRequest) SetAObjEzsigndocumentdependency(v []EzsigndocumentdependencyRequestCompound)`

SetAObjEzsigndocumentdependency sets AObjEzsigndocumentdependency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


