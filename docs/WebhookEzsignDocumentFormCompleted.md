# WebhookEzsignDocumentFormCompleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjWebhook** | [**CustomWebhookResponse**](CustomWebhookResponse.md) |  | 
**AObjAttempt** | [**[]AttemptResponseCompound**](AttemptResponseCompound.md) | An array containing details of previous attempts that were made to deliver the message. The array is empty if it&#39;s the first attempt. | 
**ObjEzsigndocument** | [**EzsigndocumentResponse**](EzsigndocumentResponse.md) |  | 

## Methods

### NewWebhookEzsignDocumentFormCompleted

`func NewWebhookEzsignDocumentFormCompleted(objWebhook CustomWebhookResponse, aObjAttempt []AttemptResponseCompound, objEzsigndocument EzsigndocumentResponse, ) *WebhookEzsignDocumentFormCompleted`

NewWebhookEzsignDocumentFormCompleted instantiates a new WebhookEzsignDocumentFormCompleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEzsignDocumentFormCompletedWithDefaults

`func NewWebhookEzsignDocumentFormCompletedWithDefaults() *WebhookEzsignDocumentFormCompleted`

NewWebhookEzsignDocumentFormCompletedWithDefaults instantiates a new WebhookEzsignDocumentFormCompleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjWebhook

`func (o *WebhookEzsignDocumentFormCompleted) GetObjWebhook() CustomWebhookResponse`

GetObjWebhook returns the ObjWebhook field if non-nil, zero value otherwise.

### GetObjWebhookOk

`func (o *WebhookEzsignDocumentFormCompleted) GetObjWebhookOk() (*CustomWebhookResponse, bool)`

GetObjWebhookOk returns a tuple with the ObjWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjWebhook

`func (o *WebhookEzsignDocumentFormCompleted) SetObjWebhook(v CustomWebhookResponse)`

SetObjWebhook sets ObjWebhook field to given value.


### GetAObjAttempt

`func (o *WebhookEzsignDocumentFormCompleted) GetAObjAttempt() []AttemptResponseCompound`

GetAObjAttempt returns the AObjAttempt field if non-nil, zero value otherwise.

### GetAObjAttemptOk

`func (o *WebhookEzsignDocumentFormCompleted) GetAObjAttemptOk() (*[]AttemptResponseCompound, bool)`

GetAObjAttemptOk returns a tuple with the AObjAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAttempt

`func (o *WebhookEzsignDocumentFormCompleted) SetAObjAttempt(v []AttemptResponseCompound)`

SetAObjAttempt sets AObjAttempt field to given value.


### GetObjEzsigndocument

`func (o *WebhookEzsignDocumentFormCompleted) GetObjEzsigndocument() EzsigndocumentResponse`

GetObjEzsigndocument returns the ObjEzsigndocument field if non-nil, zero value otherwise.

### GetObjEzsigndocumentOk

`func (o *WebhookEzsignDocumentFormCompleted) GetObjEzsigndocumentOk() (*EzsigndocumentResponse, bool)`

GetObjEzsigndocumentOk returns a tuple with the ObjEzsigndocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsigndocument

`func (o *WebhookEzsignDocumentFormCompleted) SetObjEzsigndocument(v EzsigndocumentResponse)`

SetObjEzsigndocument sets ObjEzsigndocument field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


