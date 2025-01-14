# EzsigndocumentCreateObjectV3ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AObjEzsigndocument** | [**[]EzsigndocumentCreateElementV3Response**](EzsigndocumentCreateElementV3Response.md) | An array of objets that contain unique IDs representing the object that were requested to be created and possibly matching template IDs.  They are returned in the same order as the array containing the objects to be created that was sent in the request. | 

## Methods

### NewEzsigndocumentCreateObjectV3ResponseMPayload

`func NewEzsigndocumentCreateObjectV3ResponseMPayload(aObjEzsigndocument []EzsigndocumentCreateElementV3Response, ) *EzsigndocumentCreateObjectV3ResponseMPayload`

NewEzsigndocumentCreateObjectV3ResponseMPayload instantiates a new EzsigndocumentCreateObjectV3ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentCreateObjectV3ResponseMPayloadWithDefaults

`func NewEzsigndocumentCreateObjectV3ResponseMPayloadWithDefaults() *EzsigndocumentCreateObjectV3ResponseMPayload`

NewEzsigndocumentCreateObjectV3ResponseMPayloadWithDefaults instantiates a new EzsigndocumentCreateObjectV3ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAObjEzsigndocument

`func (o *EzsigndocumentCreateObjectV3ResponseMPayload) GetAObjEzsigndocument() []EzsigndocumentCreateElementV3Response`

GetAObjEzsigndocument returns the AObjEzsigndocument field if non-nil, zero value otherwise.

### GetAObjEzsigndocumentOk

`func (o *EzsigndocumentCreateObjectV3ResponseMPayload) GetAObjEzsigndocumentOk() (*[]EzsigndocumentCreateElementV3Response, bool)`

GetAObjEzsigndocumentOk returns a tuple with the AObjEzsigndocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigndocument

`func (o *EzsigndocumentCreateObjectV3ResponseMPayload) SetAObjEzsigndocument(v []EzsigndocumentCreateElementV3Response)`

SetAObjEzsigndocument sets AObjEzsigndocument field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


