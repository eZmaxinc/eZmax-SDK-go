# WebhookEzsignFolderUnsent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjWebhook** | [**CustomWebhookResponse**](CustomWebhookResponse.md) |  | 
**AObjAttempt** | [**[]AttemptResponseCompound**](AttemptResponseCompound.md) | An array containing details of previous attempts that were made to deliver the message. The array is empty if it&#39;s the first attempt. | 
**ObjEzsignfolder** | [**EzsignfolderResponse**](EzsignfolderResponse.md) |  | 

## Methods

### NewWebhookEzsignFolderUnsent

`func NewWebhookEzsignFolderUnsent(objWebhook CustomWebhookResponse, aObjAttempt []AttemptResponseCompound, objEzsignfolder EzsignfolderResponse, ) *WebhookEzsignFolderUnsent`

NewWebhookEzsignFolderUnsent instantiates a new WebhookEzsignFolderUnsent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEzsignFolderUnsentWithDefaults

`func NewWebhookEzsignFolderUnsentWithDefaults() *WebhookEzsignFolderUnsent`

NewWebhookEzsignFolderUnsentWithDefaults instantiates a new WebhookEzsignFolderUnsent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjWebhook

`func (o *WebhookEzsignFolderUnsent) GetObjWebhook() CustomWebhookResponse`

GetObjWebhook returns the ObjWebhook field if non-nil, zero value otherwise.

### GetObjWebhookOk

`func (o *WebhookEzsignFolderUnsent) GetObjWebhookOk() (*CustomWebhookResponse, bool)`

GetObjWebhookOk returns a tuple with the ObjWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjWebhook

`func (o *WebhookEzsignFolderUnsent) SetObjWebhook(v CustomWebhookResponse)`

SetObjWebhook sets ObjWebhook field to given value.


### GetAObjAttempt

`func (o *WebhookEzsignFolderUnsent) GetAObjAttempt() []AttemptResponseCompound`

GetAObjAttempt returns the AObjAttempt field if non-nil, zero value otherwise.

### GetAObjAttemptOk

`func (o *WebhookEzsignFolderUnsent) GetAObjAttemptOk() (*[]AttemptResponseCompound, bool)`

GetAObjAttemptOk returns a tuple with the AObjAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAttempt

`func (o *WebhookEzsignFolderUnsent) SetAObjAttempt(v []AttemptResponseCompound)`

SetAObjAttempt sets AObjAttempt field to given value.


### GetObjEzsignfolder

`func (o *WebhookEzsignFolderUnsent) GetObjEzsignfolder() EzsignfolderResponse`

GetObjEzsignfolder returns the ObjEzsignfolder field if non-nil, zero value otherwise.

### GetObjEzsignfolderOk

`func (o *WebhookEzsignFolderUnsent) GetObjEzsignfolderOk() (*EzsignfolderResponse, bool)`

GetObjEzsignfolderOk returns a tuple with the ObjEzsignfolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignfolder

`func (o *WebhookEzsignFolderUnsent) SetObjEzsignfolder(v EzsignfolderResponse)`

SetObjEzsignfolder sets ObjEzsignfolder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


