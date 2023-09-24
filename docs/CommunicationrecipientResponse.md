# CommunicationrecipientResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCommunicationrecipientID** | **int32** | The unique ID of the Communicationrecipient. | 
**ECommunicationrecipientObjecttype** | Pointer to [**FieldECommunicationrecipientObjecttype**](FieldECommunicationrecipientObjecttype.md) |  | [optional] 
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
**ECommunicationrecipientType** | [**FieldECommunicationrecipientType**](FieldECommunicationrecipientType.md) |  | 
**ObjDescriptionstatic** | [**DescriptionstaticResponseCompound**](DescriptionstaticResponseCompound.md) |  | 
**ObjEmailstatic** | Pointer to [**EmailstaticResponseCompound**](EmailstaticResponseCompound.md) |  | [optional] 
**ObjPhonestatic** | Pointer to [**PhonestaticResponseCompound**](PhonestaticResponseCompound.md) |  | [optional] 

## Methods

### NewCommunicationrecipientResponse

`func NewCommunicationrecipientResponse(pkiCommunicationrecipientID int32, eCommunicationrecipientType FieldECommunicationrecipientType, objDescriptionstatic DescriptionstaticResponseCompound, ) *CommunicationrecipientResponse`

NewCommunicationrecipientResponse instantiates a new CommunicationrecipientResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationrecipientResponseWithDefaults

`func NewCommunicationrecipientResponseWithDefaults() *CommunicationrecipientResponse`

NewCommunicationrecipientResponseWithDefaults instantiates a new CommunicationrecipientResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCommunicationrecipientID

`func (o *CommunicationrecipientResponse) GetPkiCommunicationrecipientID() int32`

GetPkiCommunicationrecipientID returns the PkiCommunicationrecipientID field if non-nil, zero value otherwise.

### GetPkiCommunicationrecipientIDOk

`func (o *CommunicationrecipientResponse) GetPkiCommunicationrecipientIDOk() (*int32, bool)`

GetPkiCommunicationrecipientIDOk returns a tuple with the PkiCommunicationrecipientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCommunicationrecipientID

`func (o *CommunicationrecipientResponse) SetPkiCommunicationrecipientID(v int32)`

SetPkiCommunicationrecipientID sets PkiCommunicationrecipientID field to given value.


### GetECommunicationrecipientObjecttype

`func (o *CommunicationrecipientResponse) GetECommunicationrecipientObjecttype() FieldECommunicationrecipientObjecttype`

GetECommunicationrecipientObjecttype returns the ECommunicationrecipientObjecttype field if non-nil, zero value otherwise.

### GetECommunicationrecipientObjecttypeOk

`func (o *CommunicationrecipientResponse) GetECommunicationrecipientObjecttypeOk() (*FieldECommunicationrecipientObjecttype, bool)`

GetECommunicationrecipientObjecttypeOk returns a tuple with the ECommunicationrecipientObjecttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationrecipientObjecttype

`func (o *CommunicationrecipientResponse) SetECommunicationrecipientObjecttype(v FieldECommunicationrecipientObjecttype)`

SetECommunicationrecipientObjecttype sets ECommunicationrecipientObjecttype field to given value.

### HasECommunicationrecipientObjecttype

`func (o *CommunicationrecipientResponse) HasECommunicationrecipientObjecttype() bool`

HasECommunicationrecipientObjecttype returns a boolean if a field has been set.

### GetFkiAgentID

`func (o *CommunicationrecipientResponse) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *CommunicationrecipientResponse) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *CommunicationrecipientResponse) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *CommunicationrecipientResponse) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *CommunicationrecipientResponse) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *CommunicationrecipientResponse) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *CommunicationrecipientResponse) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *CommunicationrecipientResponse) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetFkiContactID

`func (o *CommunicationrecipientResponse) GetFkiContactID() int32`

GetFkiContactID returns the FkiContactID field if non-nil, zero value otherwise.

### GetFkiContactIDOk

`func (o *CommunicationrecipientResponse) GetFkiContactIDOk() (*int32, bool)`

GetFkiContactIDOk returns a tuple with the FkiContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiContactID

`func (o *CommunicationrecipientResponse) SetFkiContactID(v int32)`

SetFkiContactID sets FkiContactID field to given value.

### HasFkiContactID

`func (o *CommunicationrecipientResponse) HasFkiContactID() bool`

HasFkiContactID returns a boolean if a field has been set.

### GetFkiCustomerID

`func (o *CommunicationrecipientResponse) GetFkiCustomerID() int32`

GetFkiCustomerID returns the FkiCustomerID field if non-nil, zero value otherwise.

### GetFkiCustomerIDOk

`func (o *CommunicationrecipientResponse) GetFkiCustomerIDOk() (*int32, bool)`

GetFkiCustomerIDOk returns a tuple with the FkiCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomerID

`func (o *CommunicationrecipientResponse) SetFkiCustomerID(v int32)`

SetFkiCustomerID sets FkiCustomerID field to given value.

### HasFkiCustomerID

`func (o *CommunicationrecipientResponse) HasFkiCustomerID() bool`

HasFkiCustomerID returns a boolean if a field has been set.

### GetFkiEmployeeID

`func (o *CommunicationrecipientResponse) GetFkiEmployeeID() int32`

GetFkiEmployeeID returns the FkiEmployeeID field if non-nil, zero value otherwise.

### GetFkiEmployeeIDOk

`func (o *CommunicationrecipientResponse) GetFkiEmployeeIDOk() (*int32, bool)`

GetFkiEmployeeIDOk returns a tuple with the FkiEmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmployeeID

`func (o *CommunicationrecipientResponse) SetFkiEmployeeID(v int32)`

SetFkiEmployeeID sets FkiEmployeeID field to given value.

### HasFkiEmployeeID

`func (o *CommunicationrecipientResponse) HasFkiEmployeeID() bool`

HasFkiEmployeeID returns a boolean if a field has been set.

### GetFkiEzsignsignerID

`func (o *CommunicationrecipientResponse) GetFkiEzsignsignerID() int32`

GetFkiEzsignsignerID returns the FkiEzsignsignerID field if non-nil, zero value otherwise.

### GetFkiEzsignsignerIDOk

`func (o *CommunicationrecipientResponse) GetFkiEzsignsignerIDOk() (*int32, bool)`

GetFkiEzsignsignerIDOk returns a tuple with the FkiEzsignsignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsignerID

`func (o *CommunicationrecipientResponse) SetFkiEzsignsignerID(v int32)`

SetFkiEzsignsignerID sets FkiEzsignsignerID field to given value.

### HasFkiEzsignsignerID

`func (o *CommunicationrecipientResponse) HasFkiEzsignsignerID() bool`

HasFkiEzsignsignerID returns a boolean if a field has been set.

### GetFkiFranchiseofficeID

`func (o *CommunicationrecipientResponse) GetFkiFranchiseofficeID() int32`

GetFkiFranchiseofficeID returns the FkiFranchiseofficeID field if non-nil, zero value otherwise.

### GetFkiFranchiseofficeIDOk

`func (o *CommunicationrecipientResponse) GetFkiFranchiseofficeIDOk() (*int32, bool)`

GetFkiFranchiseofficeIDOk returns a tuple with the FkiFranchiseofficeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseofficeID

`func (o *CommunicationrecipientResponse) SetFkiFranchiseofficeID(v int32)`

SetFkiFranchiseofficeID sets FkiFranchiseofficeID field to given value.

### HasFkiFranchiseofficeID

`func (o *CommunicationrecipientResponse) HasFkiFranchiseofficeID() bool`

HasFkiFranchiseofficeID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *CommunicationrecipientResponse) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *CommunicationrecipientResponse) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *CommunicationrecipientResponse) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *CommunicationrecipientResponse) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiAgentincorporationID

`func (o *CommunicationrecipientResponse) GetFkiAgentincorporationID() int32`

GetFkiAgentincorporationID returns the FkiAgentincorporationID field if non-nil, zero value otherwise.

### GetFkiAgentincorporationIDOk

`func (o *CommunicationrecipientResponse) GetFkiAgentincorporationIDOk() (*int32, bool)`

GetFkiAgentincorporationIDOk returns a tuple with the FkiAgentincorporationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentincorporationID

`func (o *CommunicationrecipientResponse) SetFkiAgentincorporationID(v int32)`

SetFkiAgentincorporationID sets FkiAgentincorporationID field to given value.

### HasFkiAgentincorporationID

`func (o *CommunicationrecipientResponse) HasFkiAgentincorporationID() bool`

HasFkiAgentincorporationID returns a boolean if a field has been set.

### GetFkiAssistantID

`func (o *CommunicationrecipientResponse) GetFkiAssistantID() int32`

GetFkiAssistantID returns the FkiAssistantID field if non-nil, zero value otherwise.

### GetFkiAssistantIDOk

`func (o *CommunicationrecipientResponse) GetFkiAssistantIDOk() (*int32, bool)`

GetFkiAssistantIDOk returns a tuple with the FkiAssistantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAssistantID

`func (o *CommunicationrecipientResponse) SetFkiAssistantID(v int32)`

SetFkiAssistantID sets FkiAssistantID field to given value.

### HasFkiAssistantID

`func (o *CommunicationrecipientResponse) HasFkiAssistantID() bool`

HasFkiAssistantID returns a boolean if a field has been set.

### GetFkiExternalbrokerID

`func (o *CommunicationrecipientResponse) GetFkiExternalbrokerID() int32`

GetFkiExternalbrokerID returns the FkiExternalbrokerID field if non-nil, zero value otherwise.

### GetFkiExternalbrokerIDOk

`func (o *CommunicationrecipientResponse) GetFkiExternalbrokerIDOk() (*int32, bool)`

GetFkiExternalbrokerIDOk returns a tuple with the FkiExternalbrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiExternalbrokerID

`func (o *CommunicationrecipientResponse) SetFkiExternalbrokerID(v int32)`

SetFkiExternalbrokerID sets FkiExternalbrokerID field to given value.

### HasFkiExternalbrokerID

`func (o *CommunicationrecipientResponse) HasFkiExternalbrokerID() bool`

HasFkiExternalbrokerID returns a boolean if a field has been set.

### GetFkiEzcomagentID

`func (o *CommunicationrecipientResponse) GetFkiEzcomagentID() int32`

GetFkiEzcomagentID returns the FkiEzcomagentID field if non-nil, zero value otherwise.

### GetFkiEzcomagentIDOk

`func (o *CommunicationrecipientResponse) GetFkiEzcomagentIDOk() (*int32, bool)`

GetFkiEzcomagentIDOk returns a tuple with the FkiEzcomagentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzcomagentID

`func (o *CommunicationrecipientResponse) SetFkiEzcomagentID(v int32)`

SetFkiEzcomagentID sets FkiEzcomagentID field to given value.

### HasFkiEzcomagentID

`func (o *CommunicationrecipientResponse) HasFkiEzcomagentID() bool`

HasFkiEzcomagentID returns a boolean if a field has been set.

### GetFkiNotaryID

`func (o *CommunicationrecipientResponse) GetFkiNotaryID() int32`

GetFkiNotaryID returns the FkiNotaryID field if non-nil, zero value otherwise.

### GetFkiNotaryIDOk

`func (o *CommunicationrecipientResponse) GetFkiNotaryIDOk() (*int32, bool)`

GetFkiNotaryIDOk returns a tuple with the FkiNotaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiNotaryID

`func (o *CommunicationrecipientResponse) SetFkiNotaryID(v int32)`

SetFkiNotaryID sets FkiNotaryID field to given value.

### HasFkiNotaryID

`func (o *CommunicationrecipientResponse) HasFkiNotaryID() bool`

HasFkiNotaryID returns a boolean if a field has been set.

### GetFkiRewardmemberID

`func (o *CommunicationrecipientResponse) GetFkiRewardmemberID() int32`

GetFkiRewardmemberID returns the FkiRewardmemberID field if non-nil, zero value otherwise.

### GetFkiRewardmemberIDOk

`func (o *CommunicationrecipientResponse) GetFkiRewardmemberIDOk() (*int32, bool)`

GetFkiRewardmemberIDOk returns a tuple with the FkiRewardmemberID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiRewardmemberID

`func (o *CommunicationrecipientResponse) SetFkiRewardmemberID(v int32)`

SetFkiRewardmemberID sets FkiRewardmemberID field to given value.

### HasFkiRewardmemberID

`func (o *CommunicationrecipientResponse) HasFkiRewardmemberID() bool`

HasFkiRewardmemberID returns a boolean if a field has been set.

### GetFkiSupplierID

`func (o *CommunicationrecipientResponse) GetFkiSupplierID() int32`

GetFkiSupplierID returns the FkiSupplierID field if non-nil, zero value otherwise.

### GetFkiSupplierIDOk

`func (o *CommunicationrecipientResponse) GetFkiSupplierIDOk() (*int32, bool)`

GetFkiSupplierIDOk returns a tuple with the FkiSupplierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSupplierID

`func (o *CommunicationrecipientResponse) SetFkiSupplierID(v int32)`

SetFkiSupplierID sets FkiSupplierID field to given value.

### HasFkiSupplierID

`func (o *CommunicationrecipientResponse) HasFkiSupplierID() bool`

HasFkiSupplierID returns a boolean if a field has been set.

### GetECommunicationrecipientType

`func (o *CommunicationrecipientResponse) GetECommunicationrecipientType() FieldECommunicationrecipientType`

GetECommunicationrecipientType returns the ECommunicationrecipientType field if non-nil, zero value otherwise.

### GetECommunicationrecipientTypeOk

`func (o *CommunicationrecipientResponse) GetECommunicationrecipientTypeOk() (*FieldECommunicationrecipientType, bool)`

GetECommunicationrecipientTypeOk returns a tuple with the ECommunicationrecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationrecipientType

`func (o *CommunicationrecipientResponse) SetECommunicationrecipientType(v FieldECommunicationrecipientType)`

SetECommunicationrecipientType sets ECommunicationrecipientType field to given value.


### GetObjDescriptionstatic

`func (o *CommunicationrecipientResponse) GetObjDescriptionstatic() DescriptionstaticResponseCompound`

GetObjDescriptionstatic returns the ObjDescriptionstatic field if non-nil, zero value otherwise.

### GetObjDescriptionstaticOk

`func (o *CommunicationrecipientResponse) GetObjDescriptionstaticOk() (*DescriptionstaticResponseCompound, bool)`

GetObjDescriptionstaticOk returns a tuple with the ObjDescriptionstatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDescriptionstatic

`func (o *CommunicationrecipientResponse) SetObjDescriptionstatic(v DescriptionstaticResponseCompound)`

SetObjDescriptionstatic sets ObjDescriptionstatic field to given value.


### GetObjEmailstatic

`func (o *CommunicationrecipientResponse) GetObjEmailstatic() EmailstaticResponseCompound`

GetObjEmailstatic returns the ObjEmailstatic field if non-nil, zero value otherwise.

### GetObjEmailstaticOk

`func (o *CommunicationrecipientResponse) GetObjEmailstaticOk() (*EmailstaticResponseCompound, bool)`

GetObjEmailstaticOk returns a tuple with the ObjEmailstatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmailstatic

`func (o *CommunicationrecipientResponse) SetObjEmailstatic(v EmailstaticResponseCompound)`

SetObjEmailstatic sets ObjEmailstatic field to given value.

### HasObjEmailstatic

`func (o *CommunicationrecipientResponse) HasObjEmailstatic() bool`

HasObjEmailstatic returns a boolean if a field has been set.

### GetObjPhonestatic

`func (o *CommunicationrecipientResponse) GetObjPhonestatic() PhonestaticResponseCompound`

GetObjPhonestatic returns the ObjPhonestatic field if non-nil, zero value otherwise.

### GetObjPhonestaticOk

`func (o *CommunicationrecipientResponse) GetObjPhonestaticOk() (*PhonestaticResponseCompound, bool)`

GetObjPhonestaticOk returns a tuple with the ObjPhonestatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPhonestatic

`func (o *CommunicationrecipientResponse) SetObjPhonestatic(v PhonestaticResponseCompound)`

SetObjPhonestatic sets ObjPhonestatic field to given value.

### HasObjPhonestatic

`func (o *CommunicationrecipientResponse) HasObjPhonestatic() bool`

HasObjPhonestatic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


