# WebhookEzsignEzsignsignerAcceptclause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjWebhook** | [**CustomWebhookResponse**](CustomWebhookResponse.md) |  | 
**AObjAttempt** | [**[]AttemptResponse**](AttemptResponse.md) | An array containing details of previous attempts that were made to deliver the message. The array is empty if it&#39;s the first attempt. | 
**ObjEzsignfolder** | Pointer to [**EzsignfolderResponse**](EzsignfolderResponse.md) |  | [optional] 
**ObjEzsignfoldersignerassociation** | [**EzsignfoldersignerassociationResponseCompound**](EzsignfoldersignerassociationResponseCompound.md) |  | 

## Methods

### NewWebhookEzsignEzsignsignerAcceptclause

`func NewWebhookEzsignEzsignsignerAcceptclause(objWebhook CustomWebhookResponse, aObjAttempt []AttemptResponseCompound, objEzsignfoldersignerassociation EzsignfoldersignerassociationResponseCompound, ) *WebhookEzsignEzsignsignerAcceptclause`

NewWebhookEzsignEzsignsignerAcceptclause instantiates a new WebhookEzsignEzsignsignerAcceptclause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEzsignEzsignsignerAcceptclauseWithDefaults

`func NewWebhookEzsignEzsignsignerAcceptclauseWithDefaults() *WebhookEzsignEzsignsignerAcceptclause`

NewWebhookEzsignEzsignsignerAcceptclauseWithDefaults instantiates a new WebhookEzsignEzsignsignerAcceptclause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjWebhook

`func (o *WebhookEzsignEzsignsignerAcceptclause) GetObjWebhook() CustomWebhookResponse`

GetObjWebhook returns the ObjWebhook field if non-nil, zero value otherwise.

### GetObjWebhookOk

`func (o *WebhookEzsignEzsignsignerAcceptclause) GetObjWebhookOk() (*CustomWebhookResponse, bool)`

GetObjWebhookOk returns a tuple with the ObjWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjWebhook

`func (o *WebhookEzsignEzsignsignerAcceptclause) SetObjWebhook(v CustomWebhookResponse)`

SetObjWebhook sets ObjWebhook field to given value.


### GetAObjAttempt

`func (o *WebhookEzsignEzsignsignerAcceptclause) GetAObjAttempt() []AttemptResponseCompound`

GetAObjAttempt returns the AObjAttempt field if non-nil, zero value otherwise.

### GetAObjAttemptOk

`func (o *WebhookEzsignEzsignsignerAcceptclause) GetAObjAttemptOk() (*[]AttemptResponseCompound, bool)`

GetAObjAttemptOk returns a tuple with the AObjAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAttempt

`func (o *WebhookEzsignEzsignsignerAcceptclause) SetAObjAttempt(v []AttemptResponseCompound)`

SetAObjAttempt sets AObjAttempt field to given value.


### GetObjEzsignfolder

`func (o *WebhookEzsignEzsignsignerAcceptclause) GetObjEzsignfolder() EzsignfolderResponse`

GetObjEzsignfolder returns the ObjEzsignfolder field if non-nil, zero value otherwise.

### GetObjEzsignfolderOk

`func (o *WebhookEzsignEzsignsignerAcceptclause) GetObjEzsignfolderOk() (*EzsignfolderResponse, bool)`

GetObjEzsignfolderOk returns a tuple with the ObjEzsignfolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignfolder

`func (o *WebhookEzsignEzsignsignerAcceptclause) SetObjEzsignfolder(v EzsignfolderResponse)`

SetObjEzsignfolder sets ObjEzsignfolder field to given value.

### HasObjEzsignfolder

`func (o *WebhookEzsignEzsignsignerAcceptclause) HasObjEzsignfolder() bool`

HasObjEzsignfolder returns a boolean if a field has been set.

### GetObjEzsignfoldersignerassociation

`func (o *WebhookEzsignEzsignsignerAcceptclause) GetObjEzsignfoldersignerassociation() EzsignfoldersignerassociationResponseCompound`

GetObjEzsignfoldersignerassociation returns the ObjEzsignfoldersignerassociation field if non-nil, zero value otherwise.

### GetObjEzsignfoldersignerassociationOk

`func (o *WebhookEzsignEzsignsignerAcceptclause) GetObjEzsignfoldersignerassociationOk() (*EzsignfoldersignerassociationResponseCompound, bool)`

GetObjEzsignfoldersignerassociationOk returns a tuple with the ObjEzsignfoldersignerassociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsignfoldersignerassociation

`func (o *WebhookEzsignEzsignsignerAcceptclause) SetObjEzsignfoldersignerassociation(v EzsignfoldersignerassociationResponseCompound)`

SetObjEzsignfoldersignerassociation sets ObjEzsignfoldersignerassociation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


