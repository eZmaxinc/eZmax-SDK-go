# CustomCommunicationrecipientsrecipientResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiAgentID** | Pointer to **int32** | The unique ID of the Agent. | [optional] 
**FkiBrokerID** | Pointer to **int32** | The unique ID of the Broker. | [optional] 
**FkiContactID** | Pointer to **int32** | The unique ID of the Contact | [optional] 
**FkiCustomerID** | Pointer to **int32** | The unique ID of the Customer. | [optional] 
**FkiEmployeeID** | Pointer to **int32** | The unique ID of the Employee. | [optional] 
**FkiEzsignsignerID** | Pointer to **int32** | The unique ID of the Ezsignsigner | [optional] 
**FkiFranchiseofficeID** | Pointer to **int32** | The unique ID of the Franchisereoffice | [optional] 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiAgentincorporationID** | Pointer to **int32** | The unique ID of the Agentincorporation. | [optional] 
**FkiAssistantID** | Pointer to **int32** | The unique ID of the Assistant. | [optional] 
**FkiExternalbrokerID** | Pointer to **int32** | The unique ID of the Externalbroker. | [optional] 
**FkiEzcomagentID** | Pointer to **int32** | The unique ID of the Ezcomagent. | [optional] 
**FkiNotaryID** | Pointer to **int32** | The unique ID of the Notary. | [optional] 
**FkiRewardmemberID** | Pointer to **int32** | The unique ID of the Rewardmember. | [optional] 
**FkiSupplierID** | Pointer to **int32** | The unique ID of the Supplier. | [optional] 
**ECommunicationrecipientsrecipientObjecttype** | **string** |  | 
**ObjContactName** | [**CustomContactNameResponse**](CustomContactNameResponse.md) |  | 
**ObjEmail** | Pointer to [**EmailResponseCompound**](EmailResponseCompound.md) |  | [optional] 
**ObjPhoneFax** | Pointer to [**PhoneResponseCompound**](PhoneResponseCompound.md) |  | [optional] 
**ObjPhoneSMS** | Pointer to [**PhoneResponseCompound**](PhoneResponseCompound.md) |  | [optional] 

## Methods

### NewCustomCommunicationrecipientsrecipientResponse

`func NewCustomCommunicationrecipientsrecipientResponse(eCommunicationrecipientsrecipientObjecttype string, objContactName CustomContactNameResponse, ) *CustomCommunicationrecipientsrecipientResponse`

NewCustomCommunicationrecipientsrecipientResponse instantiates a new CustomCommunicationrecipientsrecipientResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomCommunicationrecipientsrecipientResponseWithDefaults

`func NewCustomCommunicationrecipientsrecipientResponseWithDefaults() *CustomCommunicationrecipientsrecipientResponse`

NewCustomCommunicationrecipientsrecipientResponseWithDefaults instantiates a new CustomCommunicationrecipientsrecipientResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiAgentID

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *CustomCommunicationrecipientsrecipientResponse) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *CustomCommunicationrecipientsrecipientResponse) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *CustomCommunicationrecipientsrecipientResponse) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *CustomCommunicationrecipientsrecipientResponse) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetFkiContactID

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiContactID() int32`

GetFkiContactID returns the FkiContactID field if non-nil, zero value otherwise.

### GetFkiContactIDOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiContactIDOk() (*int32, bool)`

GetFkiContactIDOk returns a tuple with the FkiContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiContactID

`func (o *CustomCommunicationrecipientsrecipientResponse) SetFkiContactID(v int32)`

SetFkiContactID sets FkiContactID field to given value.

### HasFkiContactID

`func (o *CustomCommunicationrecipientsrecipientResponse) HasFkiContactID() bool`

HasFkiContactID returns a boolean if a field has been set.

### GetFkiCustomerID

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiCustomerID() int32`

GetFkiCustomerID returns the FkiCustomerID field if non-nil, zero value otherwise.

### GetFkiCustomerIDOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiCustomerIDOk() (*int32, bool)`

GetFkiCustomerIDOk returns a tuple with the FkiCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomerID

`func (o *CustomCommunicationrecipientsrecipientResponse) SetFkiCustomerID(v int32)`

SetFkiCustomerID sets FkiCustomerID field to given value.

### HasFkiCustomerID

`func (o *CustomCommunicationrecipientsrecipientResponse) HasFkiCustomerID() bool`

HasFkiCustomerID returns a boolean if a field has been set.

### GetFkiEmployeeID

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiEmployeeID() int32`

GetFkiEmployeeID returns the FkiEmployeeID field if non-nil, zero value otherwise.

### GetFkiEmployeeIDOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiEmployeeIDOk() (*int32, bool)`

GetFkiEmployeeIDOk returns a tuple with the FkiEmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmployeeID

`func (o *CustomCommunicationrecipientsrecipientResponse) SetFkiEmployeeID(v int32)`

SetFkiEmployeeID sets FkiEmployeeID field to given value.

### HasFkiEmployeeID

`func (o *CustomCommunicationrecipientsrecipientResponse) HasFkiEmployeeID() bool`

HasFkiEmployeeID returns a boolean if a field has been set.

### GetFkiEzsignsignerID

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiEzsignsignerID() int32`

GetFkiEzsignsignerID returns the FkiEzsignsignerID field if non-nil, zero value otherwise.

### GetFkiEzsignsignerIDOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiEzsignsignerIDOk() (*int32, bool)`

GetFkiEzsignsignerIDOk returns a tuple with the FkiEzsignsignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsignerID

`func (o *CustomCommunicationrecipientsrecipientResponse) SetFkiEzsignsignerID(v int32)`

SetFkiEzsignsignerID sets FkiEzsignsignerID field to given value.

### HasFkiEzsignsignerID

`func (o *CustomCommunicationrecipientsrecipientResponse) HasFkiEzsignsignerID() bool`

HasFkiEzsignsignerID returns a boolean if a field has been set.

### GetFkiFranchiseofficeID

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiFranchiseofficeID() int32`

GetFkiFranchiseofficeID returns the FkiFranchiseofficeID field if non-nil, zero value otherwise.

### GetFkiFranchiseofficeIDOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiFranchiseofficeIDOk() (*int32, bool)`

GetFkiFranchiseofficeIDOk returns a tuple with the FkiFranchiseofficeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseofficeID

`func (o *CustomCommunicationrecipientsrecipientResponse) SetFkiFranchiseofficeID(v int32)`

SetFkiFranchiseofficeID sets FkiFranchiseofficeID field to given value.

### HasFkiFranchiseofficeID

`func (o *CustomCommunicationrecipientsrecipientResponse) HasFkiFranchiseofficeID() bool`

HasFkiFranchiseofficeID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *CustomCommunicationrecipientsrecipientResponse) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *CustomCommunicationrecipientsrecipientResponse) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiAgentincorporationID

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiAgentincorporationID() int32`

GetFkiAgentincorporationID returns the FkiAgentincorporationID field if non-nil, zero value otherwise.

### GetFkiAgentincorporationIDOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiAgentincorporationIDOk() (*int32, bool)`

GetFkiAgentincorporationIDOk returns a tuple with the FkiAgentincorporationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentincorporationID

`func (o *CustomCommunicationrecipientsrecipientResponse) SetFkiAgentincorporationID(v int32)`

SetFkiAgentincorporationID sets FkiAgentincorporationID field to given value.

### HasFkiAgentincorporationID

`func (o *CustomCommunicationrecipientsrecipientResponse) HasFkiAgentincorporationID() bool`

HasFkiAgentincorporationID returns a boolean if a field has been set.

### GetFkiAssistantID

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiAssistantID() int32`

GetFkiAssistantID returns the FkiAssistantID field if non-nil, zero value otherwise.

### GetFkiAssistantIDOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiAssistantIDOk() (*int32, bool)`

GetFkiAssistantIDOk returns a tuple with the FkiAssistantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAssistantID

`func (o *CustomCommunicationrecipientsrecipientResponse) SetFkiAssistantID(v int32)`

SetFkiAssistantID sets FkiAssistantID field to given value.

### HasFkiAssistantID

`func (o *CustomCommunicationrecipientsrecipientResponse) HasFkiAssistantID() bool`

HasFkiAssistantID returns a boolean if a field has been set.

### GetFkiExternalbrokerID

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiExternalbrokerID() int32`

GetFkiExternalbrokerID returns the FkiExternalbrokerID field if non-nil, zero value otherwise.

### GetFkiExternalbrokerIDOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiExternalbrokerIDOk() (*int32, bool)`

GetFkiExternalbrokerIDOk returns a tuple with the FkiExternalbrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiExternalbrokerID

`func (o *CustomCommunicationrecipientsrecipientResponse) SetFkiExternalbrokerID(v int32)`

SetFkiExternalbrokerID sets FkiExternalbrokerID field to given value.

### HasFkiExternalbrokerID

`func (o *CustomCommunicationrecipientsrecipientResponse) HasFkiExternalbrokerID() bool`

HasFkiExternalbrokerID returns a boolean if a field has been set.

### GetFkiEzcomagentID

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiEzcomagentID() int32`

GetFkiEzcomagentID returns the FkiEzcomagentID field if non-nil, zero value otherwise.

### GetFkiEzcomagentIDOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiEzcomagentIDOk() (*int32, bool)`

GetFkiEzcomagentIDOk returns a tuple with the FkiEzcomagentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzcomagentID

`func (o *CustomCommunicationrecipientsrecipientResponse) SetFkiEzcomagentID(v int32)`

SetFkiEzcomagentID sets FkiEzcomagentID field to given value.

### HasFkiEzcomagentID

`func (o *CustomCommunicationrecipientsrecipientResponse) HasFkiEzcomagentID() bool`

HasFkiEzcomagentID returns a boolean if a field has been set.

### GetFkiNotaryID

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiNotaryID() int32`

GetFkiNotaryID returns the FkiNotaryID field if non-nil, zero value otherwise.

### GetFkiNotaryIDOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiNotaryIDOk() (*int32, bool)`

GetFkiNotaryIDOk returns a tuple with the FkiNotaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiNotaryID

`func (o *CustomCommunicationrecipientsrecipientResponse) SetFkiNotaryID(v int32)`

SetFkiNotaryID sets FkiNotaryID field to given value.

### HasFkiNotaryID

`func (o *CustomCommunicationrecipientsrecipientResponse) HasFkiNotaryID() bool`

HasFkiNotaryID returns a boolean if a field has been set.

### GetFkiRewardmemberID

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiRewardmemberID() int32`

GetFkiRewardmemberID returns the FkiRewardmemberID field if non-nil, zero value otherwise.

### GetFkiRewardmemberIDOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiRewardmemberIDOk() (*int32, bool)`

GetFkiRewardmemberIDOk returns a tuple with the FkiRewardmemberID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiRewardmemberID

`func (o *CustomCommunicationrecipientsrecipientResponse) SetFkiRewardmemberID(v int32)`

SetFkiRewardmemberID sets FkiRewardmemberID field to given value.

### HasFkiRewardmemberID

`func (o *CustomCommunicationrecipientsrecipientResponse) HasFkiRewardmemberID() bool`

HasFkiRewardmemberID returns a boolean if a field has been set.

### GetFkiSupplierID

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiSupplierID() int32`

GetFkiSupplierID returns the FkiSupplierID field if non-nil, zero value otherwise.

### GetFkiSupplierIDOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetFkiSupplierIDOk() (*int32, bool)`

GetFkiSupplierIDOk returns a tuple with the FkiSupplierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSupplierID

`func (o *CustomCommunicationrecipientsrecipientResponse) SetFkiSupplierID(v int32)`

SetFkiSupplierID sets FkiSupplierID field to given value.

### HasFkiSupplierID

`func (o *CustomCommunicationrecipientsrecipientResponse) HasFkiSupplierID() bool`

HasFkiSupplierID returns a boolean if a field has been set.

### GetECommunicationrecipientsrecipientObjecttype

`func (o *CustomCommunicationrecipientsrecipientResponse) GetECommunicationrecipientsrecipientObjecttype() string`

GetECommunicationrecipientsrecipientObjecttype returns the ECommunicationrecipientsrecipientObjecttype field if non-nil, zero value otherwise.

### GetECommunicationrecipientsrecipientObjecttypeOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetECommunicationrecipientsrecipientObjecttypeOk() (*string, bool)`

GetECommunicationrecipientsrecipientObjecttypeOk returns a tuple with the ECommunicationrecipientsrecipientObjecttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationrecipientsrecipientObjecttype

`func (o *CustomCommunicationrecipientsrecipientResponse) SetECommunicationrecipientsrecipientObjecttype(v string)`

SetECommunicationrecipientsrecipientObjecttype sets ECommunicationrecipientsrecipientObjecttype field to given value.


### GetObjContactName

`func (o *CustomCommunicationrecipientsrecipientResponse) GetObjContactName() CustomContactNameResponse`

GetObjContactName returns the ObjContactName field if non-nil, zero value otherwise.

### GetObjContactNameOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetObjContactNameOk() (*CustomContactNameResponse, bool)`

GetObjContactNameOk returns a tuple with the ObjContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContactName

`func (o *CustomCommunicationrecipientsrecipientResponse) SetObjContactName(v CustomContactNameResponse)`

SetObjContactName sets ObjContactName field to given value.


### GetObjEmail

`func (o *CustomCommunicationrecipientsrecipientResponse) GetObjEmail() EmailResponseCompound`

GetObjEmail returns the ObjEmail field if non-nil, zero value otherwise.

### GetObjEmailOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetObjEmailOk() (*EmailResponseCompound, bool)`

GetObjEmailOk returns a tuple with the ObjEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmail

`func (o *CustomCommunicationrecipientsrecipientResponse) SetObjEmail(v EmailResponseCompound)`

SetObjEmail sets ObjEmail field to given value.

### HasObjEmail

`func (o *CustomCommunicationrecipientsrecipientResponse) HasObjEmail() bool`

HasObjEmail returns a boolean if a field has been set.

### GetObjPhoneFax

`func (o *CustomCommunicationrecipientsrecipientResponse) GetObjPhoneFax() PhoneResponseCompound`

GetObjPhoneFax returns the ObjPhoneFax field if non-nil, zero value otherwise.

### GetObjPhoneFaxOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetObjPhoneFaxOk() (*PhoneResponseCompound, bool)`

GetObjPhoneFaxOk returns a tuple with the ObjPhoneFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPhoneFax

`func (o *CustomCommunicationrecipientsrecipientResponse) SetObjPhoneFax(v PhoneResponseCompound)`

SetObjPhoneFax sets ObjPhoneFax field to given value.

### HasObjPhoneFax

`func (o *CustomCommunicationrecipientsrecipientResponse) HasObjPhoneFax() bool`

HasObjPhoneFax returns a boolean if a field has been set.

### GetObjPhoneSMS

`func (o *CustomCommunicationrecipientsrecipientResponse) GetObjPhoneSMS() PhoneResponseCompound`

GetObjPhoneSMS returns the ObjPhoneSMS field if non-nil, zero value otherwise.

### GetObjPhoneSMSOk

`func (o *CustomCommunicationrecipientsrecipientResponse) GetObjPhoneSMSOk() (*PhoneResponseCompound, bool)`

GetObjPhoneSMSOk returns a tuple with the ObjPhoneSMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPhoneSMS

`func (o *CustomCommunicationrecipientsrecipientResponse) SetObjPhoneSMS(v PhoneResponseCompound)`

SetObjPhoneSMS sets ObjPhoneSMS field to given value.

### HasObjPhoneSMS

`func (o *CustomCommunicationrecipientsrecipientResponse) HasObjPhoneSMS() bool`

HasObjPhoneSMS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


