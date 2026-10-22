# DocumenttypechecklistResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDocumenttypechecklistID** | **int32** | The unique ID of the Documenttypechecklist | 
**FkiDocumenttypeID** | **int32** | The unique ID of the Documenttype | 
**SDocumenttypeNameX** | Pointer to **string** | The name of the Documenttype in the language of the requester | [optional] 
**EDocumenttypechecklistStatus** | [**FieldEDocumenttypechecklistStatus**](FieldEDocumenttypechecklistStatus.md) |  | 
**TDocumenttypechecklistComment** | **string** | The comment of the Documenttypechecklist | 
**FkiAdjustmentID** | Pointer to **int32** | The unique ID of the Adjustment | [optional] 
**FkiAgentID** | Pointer to **int32** | The unique ID of the Agent. | [optional] 
**FkiBankaccountID** | Pointer to **int32** | The unique ID of the Bankaccount | [optional] 
**FkiBrokerID** | Pointer to **int32** | The unique ID of the Broker. | [optional] 
**FkiCommissionadvanceID** | Pointer to **int32** | The unique ID of the Commissionadvance | [optional] 
**FkiCommunicationID** | Pointer to **int32** | The unique ID of the Communication. | [optional] 
**FkiCustomerID** | Pointer to **int32** | The unique ID of the Customer. | [optional] 
**FkiCustomertemplateID** | Pointer to **int32** | The unique ID of the Customertemplate | [optional] 
**FkiDepositID** | Pointer to **int32** | The unique ID of the Deposit | [optional] 
**FkiDeposittransitchequeID** | Pointer to **int32** | The unique ID of the Deposittransitcheque | [optional] 
**FkiElectronicfundstransferID** | Pointer to **int32** | The unique ID of the Electronicfundstransfer | [optional] 
**FkiEmployeeID** | Pointer to **int32** | The unique ID of the Employee. | [optional] 
**FkiExternalbrokerID** | Pointer to **int32** | The unique ID of the Externalbroker. | [optional] 
**FkiEzcomadvanceserverID** | Pointer to **int32** | The unique ID of the Ezcomadvanceserver | [optional] 
**FkiEzcomcompanyID** | Pointer to **int32** | The unique ID of the Ezcomcompany | [optional] 
**FkiEzsigndocumentID** | Pointer to **int32** | The unique ID of the Ezsigndocument | [optional] 
**FkiGhacqcontractID** | Pointer to **int32** | The unique ID of the Ghacqcontract | [optional] 
**FkiInscriptionID** | Pointer to **int32** | The unique ID of the Inscription. | [optional] 
**FkiInscriptiontempID** | Pointer to **int32** | The unique ID of the Inscriptiontemp | [optional] 
**FkiInscriptionnotauthenticatedID** | Pointer to **int32** | The unique ID of the Inscriptionnotauthenticated. | [optional] 
**FkiInvoiceID** | Pointer to **int32** | The unique ID of the Invoice. | [optional] 
**FkiBuyercontractID** | Pointer to **int32** | The unique ID of the Buyercontract | [optional] 
**FkiFranchisebrokerID** | Pointer to **int32** | The unique ID of the Franchisebroker | [optional] 
**FkiFranchiseagenceID** | Pointer to **int32** | The unique ID of the Franchiseagence | [optional] 
**FkiFranchiseofficeID** | Pointer to **int32** | The unique ID of the Franchisereoffice | [optional] 
**FkiFranchisefranchiseID** | Pointer to **int32** | The unique ID of the Franchisefranchise | [optional] 
**FkiFranchisecomplaintID** | Pointer to **int32** | The unique ID of the Franchisecomplaint | [optional] 
**FkiLeadID** | Pointer to **int32** | The unique ID of the Lead | [optional] 
**FkiMarketingprogramID** | Pointer to **int32** | The unique ID of the Marketingprogram | [optional] 
**FkiMarketingfollowID** | Pointer to **int32** | The unique ID of the Marketingfollow | [optional] 
**FkiNotaryID** | Pointer to **int32** | The unique ID of the Notary. | [optional] 
**FkiOfficetaxreportID** | Pointer to **int32** | The unique ID of the Officetaxreport | [optional] 
**FkiOtherincomeID** | Pointer to **int32** | The unique ID of the Otherincome | [optional] 
**FkiPaymentpreparationID** | Pointer to **int32** | The unique ID of the Paymentpreparation | [optional] 
**FkiPurchaseID** | Pointer to **int32** | The unique ID of the Purchase | [optional] 
**FkiSalaryID** | Pointer to **int32** | The unique ID of the Salary | [optional] 
**FkiSupplierID** | Pointer to **int32** | The unique ID of the Supplier. | [optional] 
**FkiTranqcontractID** | Pointer to **int32** | The unique ID of the Tranqcontract | [optional] 
**FkiTemplateID** | Pointer to **int32** | The unique ID of the Template | [optional] 

## Methods

### NewDocumenttypechecklistResponse

`func NewDocumenttypechecklistResponse(pkiDocumenttypechecklistID int32, fkiDocumenttypeID int32, eDocumenttypechecklistStatus FieldEDocumenttypechecklistStatus, tDocumenttypechecklistComment string, ) *DocumenttypechecklistResponse`

NewDocumenttypechecklistResponse instantiates a new DocumenttypechecklistResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumenttypechecklistResponseWithDefaults

`func NewDocumenttypechecklistResponseWithDefaults() *DocumenttypechecklistResponse`

NewDocumenttypechecklistResponseWithDefaults instantiates a new DocumenttypechecklistResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDocumenttypechecklistID

`func (o *DocumenttypechecklistResponse) GetPkiDocumenttypechecklistID() int32`

GetPkiDocumenttypechecklistID returns the PkiDocumenttypechecklistID field if non-nil, zero value otherwise.

### GetPkiDocumenttypechecklistIDOk

`func (o *DocumenttypechecklistResponse) GetPkiDocumenttypechecklistIDOk() (*int32, bool)`

GetPkiDocumenttypechecklistIDOk returns a tuple with the PkiDocumenttypechecklistID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDocumenttypechecklistID

`func (o *DocumenttypechecklistResponse) SetPkiDocumenttypechecklistID(v int32)`

SetPkiDocumenttypechecklistID sets PkiDocumenttypechecklistID field to given value.


### GetFkiDocumenttypeID

`func (o *DocumenttypechecklistResponse) GetFkiDocumenttypeID() int32`

GetFkiDocumenttypeID returns the FkiDocumenttypeID field if non-nil, zero value otherwise.

### GetFkiDocumenttypeIDOk

`func (o *DocumenttypechecklistResponse) GetFkiDocumenttypeIDOk() (*int32, bool)`

GetFkiDocumenttypeIDOk returns a tuple with the FkiDocumenttypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDocumenttypeID

`func (o *DocumenttypechecklistResponse) SetFkiDocumenttypeID(v int32)`

SetFkiDocumenttypeID sets FkiDocumenttypeID field to given value.


### GetSDocumenttypeNameX

`func (o *DocumenttypechecklistResponse) GetSDocumenttypeNameX() string`

GetSDocumenttypeNameX returns the SDocumenttypeNameX field if non-nil, zero value otherwise.

### GetSDocumenttypeNameXOk

`func (o *DocumenttypechecklistResponse) GetSDocumenttypeNameXOk() (*string, bool)`

GetSDocumenttypeNameXOk returns a tuple with the SDocumenttypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDocumenttypeNameX

`func (o *DocumenttypechecklistResponse) SetSDocumenttypeNameX(v string)`

SetSDocumenttypeNameX sets SDocumenttypeNameX field to given value.

### HasSDocumenttypeNameX

`func (o *DocumenttypechecklistResponse) HasSDocumenttypeNameX() bool`

HasSDocumenttypeNameX returns a boolean if a field has been set.

### GetEDocumenttypechecklistStatus

`func (o *DocumenttypechecklistResponse) GetEDocumenttypechecklistStatus() FieldEDocumenttypechecklistStatus`

GetEDocumenttypechecklistStatus returns the EDocumenttypechecklistStatus field if non-nil, zero value otherwise.

### GetEDocumenttypechecklistStatusOk

`func (o *DocumenttypechecklistResponse) GetEDocumenttypechecklistStatusOk() (*FieldEDocumenttypechecklistStatus, bool)`

GetEDocumenttypechecklistStatusOk returns a tuple with the EDocumenttypechecklistStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypechecklistStatus

`func (o *DocumenttypechecklistResponse) SetEDocumenttypechecklistStatus(v FieldEDocumenttypechecklistStatus)`

SetEDocumenttypechecklistStatus sets EDocumenttypechecklistStatus field to given value.


### GetTDocumenttypechecklistComment

`func (o *DocumenttypechecklistResponse) GetTDocumenttypechecklistComment() string`

GetTDocumenttypechecklistComment returns the TDocumenttypechecklistComment field if non-nil, zero value otherwise.

### GetTDocumenttypechecklistCommentOk

`func (o *DocumenttypechecklistResponse) GetTDocumenttypechecklistCommentOk() (*string, bool)`

GetTDocumenttypechecklistCommentOk returns a tuple with the TDocumenttypechecklistComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTDocumenttypechecklistComment

`func (o *DocumenttypechecklistResponse) SetTDocumenttypechecklistComment(v string)`

SetTDocumenttypechecklistComment sets TDocumenttypechecklistComment field to given value.


### GetFkiAdjustmentID

`func (o *DocumenttypechecklistResponse) GetFkiAdjustmentID() int32`

GetFkiAdjustmentID returns the FkiAdjustmentID field if non-nil, zero value otherwise.

### GetFkiAdjustmentIDOk

`func (o *DocumenttypechecklistResponse) GetFkiAdjustmentIDOk() (*int32, bool)`

GetFkiAdjustmentIDOk returns a tuple with the FkiAdjustmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAdjustmentID

`func (o *DocumenttypechecklistResponse) SetFkiAdjustmentID(v int32)`

SetFkiAdjustmentID sets FkiAdjustmentID field to given value.

### HasFkiAdjustmentID

`func (o *DocumenttypechecklistResponse) HasFkiAdjustmentID() bool`

HasFkiAdjustmentID returns a boolean if a field has been set.

### GetFkiAgentID

`func (o *DocumenttypechecklistResponse) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *DocumenttypechecklistResponse) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *DocumenttypechecklistResponse) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *DocumenttypechecklistResponse) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBankaccountID

`func (o *DocumenttypechecklistResponse) GetFkiBankaccountID() int32`

GetFkiBankaccountID returns the FkiBankaccountID field if non-nil, zero value otherwise.

### GetFkiBankaccountIDOk

`func (o *DocumenttypechecklistResponse) GetFkiBankaccountIDOk() (*int32, bool)`

GetFkiBankaccountIDOk returns a tuple with the FkiBankaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBankaccountID

`func (o *DocumenttypechecklistResponse) SetFkiBankaccountID(v int32)`

SetFkiBankaccountID sets FkiBankaccountID field to given value.

### HasFkiBankaccountID

`func (o *DocumenttypechecklistResponse) HasFkiBankaccountID() bool`

HasFkiBankaccountID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *DocumenttypechecklistResponse) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *DocumenttypechecklistResponse) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *DocumenttypechecklistResponse) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *DocumenttypechecklistResponse) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetFkiCommissionadvanceID

`func (o *DocumenttypechecklistResponse) GetFkiCommissionadvanceID() int32`

GetFkiCommissionadvanceID returns the FkiCommissionadvanceID field if non-nil, zero value otherwise.

### GetFkiCommissionadvanceIDOk

`func (o *DocumenttypechecklistResponse) GetFkiCommissionadvanceIDOk() (*int32, bool)`

GetFkiCommissionadvanceIDOk returns a tuple with the FkiCommissionadvanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCommissionadvanceID

`func (o *DocumenttypechecklistResponse) SetFkiCommissionadvanceID(v int32)`

SetFkiCommissionadvanceID sets FkiCommissionadvanceID field to given value.

### HasFkiCommissionadvanceID

`func (o *DocumenttypechecklistResponse) HasFkiCommissionadvanceID() bool`

HasFkiCommissionadvanceID returns a boolean if a field has been set.

### GetFkiCommunicationID

`func (o *DocumenttypechecklistResponse) GetFkiCommunicationID() int32`

GetFkiCommunicationID returns the FkiCommunicationID field if non-nil, zero value otherwise.

### GetFkiCommunicationIDOk

`func (o *DocumenttypechecklistResponse) GetFkiCommunicationIDOk() (*int32, bool)`

GetFkiCommunicationIDOk returns a tuple with the FkiCommunicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCommunicationID

`func (o *DocumenttypechecklistResponse) SetFkiCommunicationID(v int32)`

SetFkiCommunicationID sets FkiCommunicationID field to given value.

### HasFkiCommunicationID

`func (o *DocumenttypechecklistResponse) HasFkiCommunicationID() bool`

HasFkiCommunicationID returns a boolean if a field has been set.

### GetFkiCustomerID

`func (o *DocumenttypechecklistResponse) GetFkiCustomerID() int32`

GetFkiCustomerID returns the FkiCustomerID field if non-nil, zero value otherwise.

### GetFkiCustomerIDOk

`func (o *DocumenttypechecklistResponse) GetFkiCustomerIDOk() (*int32, bool)`

GetFkiCustomerIDOk returns a tuple with the FkiCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomerID

`func (o *DocumenttypechecklistResponse) SetFkiCustomerID(v int32)`

SetFkiCustomerID sets FkiCustomerID field to given value.

### HasFkiCustomerID

`func (o *DocumenttypechecklistResponse) HasFkiCustomerID() bool`

HasFkiCustomerID returns a boolean if a field has been set.

### GetFkiCustomertemplateID

`func (o *DocumenttypechecklistResponse) GetFkiCustomertemplateID() int32`

GetFkiCustomertemplateID returns the FkiCustomertemplateID field if non-nil, zero value otherwise.

### GetFkiCustomertemplateIDOk

`func (o *DocumenttypechecklistResponse) GetFkiCustomertemplateIDOk() (*int32, bool)`

GetFkiCustomertemplateIDOk returns a tuple with the FkiCustomertemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomertemplateID

`func (o *DocumenttypechecklistResponse) SetFkiCustomertemplateID(v int32)`

SetFkiCustomertemplateID sets FkiCustomertemplateID field to given value.

### HasFkiCustomertemplateID

`func (o *DocumenttypechecklistResponse) HasFkiCustomertemplateID() bool`

HasFkiCustomertemplateID returns a boolean if a field has been set.

### GetFkiDepositID

`func (o *DocumenttypechecklistResponse) GetFkiDepositID() int32`

GetFkiDepositID returns the FkiDepositID field if non-nil, zero value otherwise.

### GetFkiDepositIDOk

`func (o *DocumenttypechecklistResponse) GetFkiDepositIDOk() (*int32, bool)`

GetFkiDepositIDOk returns a tuple with the FkiDepositID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepositID

`func (o *DocumenttypechecklistResponse) SetFkiDepositID(v int32)`

SetFkiDepositID sets FkiDepositID field to given value.

### HasFkiDepositID

`func (o *DocumenttypechecklistResponse) HasFkiDepositID() bool`

HasFkiDepositID returns a boolean if a field has been set.

### GetFkiDeposittransitchequeID

`func (o *DocumenttypechecklistResponse) GetFkiDeposittransitchequeID() int32`

GetFkiDeposittransitchequeID returns the FkiDeposittransitchequeID field if non-nil, zero value otherwise.

### GetFkiDeposittransitchequeIDOk

`func (o *DocumenttypechecklistResponse) GetFkiDeposittransitchequeIDOk() (*int32, bool)`

GetFkiDeposittransitchequeIDOk returns a tuple with the FkiDeposittransitchequeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDeposittransitchequeID

`func (o *DocumenttypechecklistResponse) SetFkiDeposittransitchequeID(v int32)`

SetFkiDeposittransitchequeID sets FkiDeposittransitchequeID field to given value.

### HasFkiDeposittransitchequeID

`func (o *DocumenttypechecklistResponse) HasFkiDeposittransitchequeID() bool`

HasFkiDeposittransitchequeID returns a boolean if a field has been set.

### GetFkiElectronicfundstransferID

`func (o *DocumenttypechecklistResponse) GetFkiElectronicfundstransferID() int32`

GetFkiElectronicfundstransferID returns the FkiElectronicfundstransferID field if non-nil, zero value otherwise.

### GetFkiElectronicfundstransferIDOk

`func (o *DocumenttypechecklistResponse) GetFkiElectronicfundstransferIDOk() (*int32, bool)`

GetFkiElectronicfundstransferIDOk returns a tuple with the FkiElectronicfundstransferID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiElectronicfundstransferID

`func (o *DocumenttypechecklistResponse) SetFkiElectronicfundstransferID(v int32)`

SetFkiElectronicfundstransferID sets FkiElectronicfundstransferID field to given value.

### HasFkiElectronicfundstransferID

`func (o *DocumenttypechecklistResponse) HasFkiElectronicfundstransferID() bool`

HasFkiElectronicfundstransferID returns a boolean if a field has been set.

### GetFkiEmployeeID

`func (o *DocumenttypechecklistResponse) GetFkiEmployeeID() int32`

GetFkiEmployeeID returns the FkiEmployeeID field if non-nil, zero value otherwise.

### GetFkiEmployeeIDOk

`func (o *DocumenttypechecklistResponse) GetFkiEmployeeIDOk() (*int32, bool)`

GetFkiEmployeeIDOk returns a tuple with the FkiEmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmployeeID

`func (o *DocumenttypechecklistResponse) SetFkiEmployeeID(v int32)`

SetFkiEmployeeID sets FkiEmployeeID field to given value.

### HasFkiEmployeeID

`func (o *DocumenttypechecklistResponse) HasFkiEmployeeID() bool`

HasFkiEmployeeID returns a boolean if a field has been set.

### GetFkiExternalbrokerID

`func (o *DocumenttypechecklistResponse) GetFkiExternalbrokerID() int32`

GetFkiExternalbrokerID returns the FkiExternalbrokerID field if non-nil, zero value otherwise.

### GetFkiExternalbrokerIDOk

`func (o *DocumenttypechecklistResponse) GetFkiExternalbrokerIDOk() (*int32, bool)`

GetFkiExternalbrokerIDOk returns a tuple with the FkiExternalbrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiExternalbrokerID

`func (o *DocumenttypechecklistResponse) SetFkiExternalbrokerID(v int32)`

SetFkiExternalbrokerID sets FkiExternalbrokerID field to given value.

### HasFkiExternalbrokerID

`func (o *DocumenttypechecklistResponse) HasFkiExternalbrokerID() bool`

HasFkiExternalbrokerID returns a boolean if a field has been set.

### GetFkiEzcomadvanceserverID

`func (o *DocumenttypechecklistResponse) GetFkiEzcomadvanceserverID() int32`

GetFkiEzcomadvanceserverID returns the FkiEzcomadvanceserverID field if non-nil, zero value otherwise.

### GetFkiEzcomadvanceserverIDOk

`func (o *DocumenttypechecklistResponse) GetFkiEzcomadvanceserverIDOk() (*int32, bool)`

GetFkiEzcomadvanceserverIDOk returns a tuple with the FkiEzcomadvanceserverID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzcomadvanceserverID

`func (o *DocumenttypechecklistResponse) SetFkiEzcomadvanceserverID(v int32)`

SetFkiEzcomadvanceserverID sets FkiEzcomadvanceserverID field to given value.

### HasFkiEzcomadvanceserverID

`func (o *DocumenttypechecklistResponse) HasFkiEzcomadvanceserverID() bool`

HasFkiEzcomadvanceserverID returns a boolean if a field has been set.

### GetFkiEzcomcompanyID

`func (o *DocumenttypechecklistResponse) GetFkiEzcomcompanyID() int32`

GetFkiEzcomcompanyID returns the FkiEzcomcompanyID field if non-nil, zero value otherwise.

### GetFkiEzcomcompanyIDOk

`func (o *DocumenttypechecklistResponse) GetFkiEzcomcompanyIDOk() (*int32, bool)`

GetFkiEzcomcompanyIDOk returns a tuple with the FkiEzcomcompanyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzcomcompanyID

`func (o *DocumenttypechecklistResponse) SetFkiEzcomcompanyID(v int32)`

SetFkiEzcomcompanyID sets FkiEzcomcompanyID field to given value.

### HasFkiEzcomcompanyID

`func (o *DocumenttypechecklistResponse) HasFkiEzcomcompanyID() bool`

HasFkiEzcomcompanyID returns a boolean if a field has been set.

### GetFkiEzsigndocumentID

`func (o *DocumenttypechecklistResponse) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *DocumenttypechecklistResponse) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *DocumenttypechecklistResponse) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.

### HasFkiEzsigndocumentID

`func (o *DocumenttypechecklistResponse) HasFkiEzsigndocumentID() bool`

HasFkiEzsigndocumentID returns a boolean if a field has been set.

### GetFkiGhacqcontractID

`func (o *DocumenttypechecklistResponse) GetFkiGhacqcontractID() int32`

GetFkiGhacqcontractID returns the FkiGhacqcontractID field if non-nil, zero value otherwise.

### GetFkiGhacqcontractIDOk

`func (o *DocumenttypechecklistResponse) GetFkiGhacqcontractIDOk() (*int32, bool)`

GetFkiGhacqcontractIDOk returns a tuple with the FkiGhacqcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGhacqcontractID

`func (o *DocumenttypechecklistResponse) SetFkiGhacqcontractID(v int32)`

SetFkiGhacqcontractID sets FkiGhacqcontractID field to given value.

### HasFkiGhacqcontractID

`func (o *DocumenttypechecklistResponse) HasFkiGhacqcontractID() bool`

HasFkiGhacqcontractID returns a boolean if a field has been set.

### GetFkiInscriptionID

`func (o *DocumenttypechecklistResponse) GetFkiInscriptionID() int32`

GetFkiInscriptionID returns the FkiInscriptionID field if non-nil, zero value otherwise.

### GetFkiInscriptionIDOk

`func (o *DocumenttypechecklistResponse) GetFkiInscriptionIDOk() (*int32, bool)`

GetFkiInscriptionIDOk returns a tuple with the FkiInscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionID

`func (o *DocumenttypechecklistResponse) SetFkiInscriptionID(v int32)`

SetFkiInscriptionID sets FkiInscriptionID field to given value.

### HasFkiInscriptionID

`func (o *DocumenttypechecklistResponse) HasFkiInscriptionID() bool`

HasFkiInscriptionID returns a boolean if a field has been set.

### GetFkiInscriptiontempID

`func (o *DocumenttypechecklistResponse) GetFkiInscriptiontempID() int32`

GetFkiInscriptiontempID returns the FkiInscriptiontempID field if non-nil, zero value otherwise.

### GetFkiInscriptiontempIDOk

`func (o *DocumenttypechecklistResponse) GetFkiInscriptiontempIDOk() (*int32, bool)`

GetFkiInscriptiontempIDOk returns a tuple with the FkiInscriptiontempID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptiontempID

`func (o *DocumenttypechecklistResponse) SetFkiInscriptiontempID(v int32)`

SetFkiInscriptiontempID sets FkiInscriptiontempID field to given value.

### HasFkiInscriptiontempID

`func (o *DocumenttypechecklistResponse) HasFkiInscriptiontempID() bool`

HasFkiInscriptiontempID returns a boolean if a field has been set.

### GetFkiInscriptionnotauthenticatedID

`func (o *DocumenttypechecklistResponse) GetFkiInscriptionnotauthenticatedID() int32`

GetFkiInscriptionnotauthenticatedID returns the FkiInscriptionnotauthenticatedID field if non-nil, zero value otherwise.

### GetFkiInscriptionnotauthenticatedIDOk

`func (o *DocumenttypechecklistResponse) GetFkiInscriptionnotauthenticatedIDOk() (*int32, bool)`

GetFkiInscriptionnotauthenticatedIDOk returns a tuple with the FkiInscriptionnotauthenticatedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionnotauthenticatedID

`func (o *DocumenttypechecklistResponse) SetFkiInscriptionnotauthenticatedID(v int32)`

SetFkiInscriptionnotauthenticatedID sets FkiInscriptionnotauthenticatedID field to given value.

### HasFkiInscriptionnotauthenticatedID

`func (o *DocumenttypechecklistResponse) HasFkiInscriptionnotauthenticatedID() bool`

HasFkiInscriptionnotauthenticatedID returns a boolean if a field has been set.

### GetFkiInvoiceID

`func (o *DocumenttypechecklistResponse) GetFkiInvoiceID() int32`

GetFkiInvoiceID returns the FkiInvoiceID field if non-nil, zero value otherwise.

### GetFkiInvoiceIDOk

`func (o *DocumenttypechecklistResponse) GetFkiInvoiceIDOk() (*int32, bool)`

GetFkiInvoiceIDOk returns a tuple with the FkiInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInvoiceID

`func (o *DocumenttypechecklistResponse) SetFkiInvoiceID(v int32)`

SetFkiInvoiceID sets FkiInvoiceID field to given value.

### HasFkiInvoiceID

`func (o *DocumenttypechecklistResponse) HasFkiInvoiceID() bool`

HasFkiInvoiceID returns a boolean if a field has been set.

### GetFkiBuyercontractID

`func (o *DocumenttypechecklistResponse) GetFkiBuyercontractID() int32`

GetFkiBuyercontractID returns the FkiBuyercontractID field if non-nil, zero value otherwise.

### GetFkiBuyercontractIDOk

`func (o *DocumenttypechecklistResponse) GetFkiBuyercontractIDOk() (*int32, bool)`

GetFkiBuyercontractIDOk returns a tuple with the FkiBuyercontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBuyercontractID

`func (o *DocumenttypechecklistResponse) SetFkiBuyercontractID(v int32)`

SetFkiBuyercontractID sets FkiBuyercontractID field to given value.

### HasFkiBuyercontractID

`func (o *DocumenttypechecklistResponse) HasFkiBuyercontractID() bool`

HasFkiBuyercontractID returns a boolean if a field has been set.

### GetFkiFranchisebrokerID

`func (o *DocumenttypechecklistResponse) GetFkiFranchisebrokerID() int32`

GetFkiFranchisebrokerID returns the FkiFranchisebrokerID field if non-nil, zero value otherwise.

### GetFkiFranchisebrokerIDOk

`func (o *DocumenttypechecklistResponse) GetFkiFranchisebrokerIDOk() (*int32, bool)`

GetFkiFranchisebrokerIDOk returns a tuple with the FkiFranchisebrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisebrokerID

`func (o *DocumenttypechecklistResponse) SetFkiFranchisebrokerID(v int32)`

SetFkiFranchisebrokerID sets FkiFranchisebrokerID field to given value.

### HasFkiFranchisebrokerID

`func (o *DocumenttypechecklistResponse) HasFkiFranchisebrokerID() bool`

HasFkiFranchisebrokerID returns a boolean if a field has been set.

### GetFkiFranchiseagenceID

`func (o *DocumenttypechecklistResponse) GetFkiFranchiseagenceID() int32`

GetFkiFranchiseagenceID returns the FkiFranchiseagenceID field if non-nil, zero value otherwise.

### GetFkiFranchiseagenceIDOk

`func (o *DocumenttypechecklistResponse) GetFkiFranchiseagenceIDOk() (*int32, bool)`

GetFkiFranchiseagenceIDOk returns a tuple with the FkiFranchiseagenceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseagenceID

`func (o *DocumenttypechecklistResponse) SetFkiFranchiseagenceID(v int32)`

SetFkiFranchiseagenceID sets FkiFranchiseagenceID field to given value.

### HasFkiFranchiseagenceID

`func (o *DocumenttypechecklistResponse) HasFkiFranchiseagenceID() bool`

HasFkiFranchiseagenceID returns a boolean if a field has been set.

### GetFkiFranchiseofficeID

`func (o *DocumenttypechecklistResponse) GetFkiFranchiseofficeID() int32`

GetFkiFranchiseofficeID returns the FkiFranchiseofficeID field if non-nil, zero value otherwise.

### GetFkiFranchiseofficeIDOk

`func (o *DocumenttypechecklistResponse) GetFkiFranchiseofficeIDOk() (*int32, bool)`

GetFkiFranchiseofficeIDOk returns a tuple with the FkiFranchiseofficeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseofficeID

`func (o *DocumenttypechecklistResponse) SetFkiFranchiseofficeID(v int32)`

SetFkiFranchiseofficeID sets FkiFranchiseofficeID field to given value.

### HasFkiFranchiseofficeID

`func (o *DocumenttypechecklistResponse) HasFkiFranchiseofficeID() bool`

HasFkiFranchiseofficeID returns a boolean if a field has been set.

### GetFkiFranchisefranchiseID

`func (o *DocumenttypechecklistResponse) GetFkiFranchisefranchiseID() int32`

GetFkiFranchisefranchiseID returns the FkiFranchisefranchiseID field if non-nil, zero value otherwise.

### GetFkiFranchisefranchiseIDOk

`func (o *DocumenttypechecklistResponse) GetFkiFranchisefranchiseIDOk() (*int32, bool)`

GetFkiFranchisefranchiseIDOk returns a tuple with the FkiFranchisefranchiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisefranchiseID

`func (o *DocumenttypechecklistResponse) SetFkiFranchisefranchiseID(v int32)`

SetFkiFranchisefranchiseID sets FkiFranchisefranchiseID field to given value.

### HasFkiFranchisefranchiseID

`func (o *DocumenttypechecklistResponse) HasFkiFranchisefranchiseID() bool`

HasFkiFranchisefranchiseID returns a boolean if a field has been set.

### GetFkiFranchisecomplaintID

`func (o *DocumenttypechecklistResponse) GetFkiFranchisecomplaintID() int32`

GetFkiFranchisecomplaintID returns the FkiFranchisecomplaintID field if non-nil, zero value otherwise.

### GetFkiFranchisecomplaintIDOk

`func (o *DocumenttypechecklistResponse) GetFkiFranchisecomplaintIDOk() (*int32, bool)`

GetFkiFranchisecomplaintIDOk returns a tuple with the FkiFranchisecomplaintID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisecomplaintID

`func (o *DocumenttypechecklistResponse) SetFkiFranchisecomplaintID(v int32)`

SetFkiFranchisecomplaintID sets FkiFranchisecomplaintID field to given value.

### HasFkiFranchisecomplaintID

`func (o *DocumenttypechecklistResponse) HasFkiFranchisecomplaintID() bool`

HasFkiFranchisecomplaintID returns a boolean if a field has been set.

### GetFkiLeadID

`func (o *DocumenttypechecklistResponse) GetFkiLeadID() int32`

GetFkiLeadID returns the FkiLeadID field if non-nil, zero value otherwise.

### GetFkiLeadIDOk

`func (o *DocumenttypechecklistResponse) GetFkiLeadIDOk() (*int32, bool)`

GetFkiLeadIDOk returns a tuple with the FkiLeadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLeadID

`func (o *DocumenttypechecklistResponse) SetFkiLeadID(v int32)`

SetFkiLeadID sets FkiLeadID field to given value.

### HasFkiLeadID

`func (o *DocumenttypechecklistResponse) HasFkiLeadID() bool`

HasFkiLeadID returns a boolean if a field has been set.

### GetFkiMarketingprogramID

`func (o *DocumenttypechecklistResponse) GetFkiMarketingprogramID() int32`

GetFkiMarketingprogramID returns the FkiMarketingprogramID field if non-nil, zero value otherwise.

### GetFkiMarketingprogramIDOk

`func (o *DocumenttypechecklistResponse) GetFkiMarketingprogramIDOk() (*int32, bool)`

GetFkiMarketingprogramIDOk returns a tuple with the FkiMarketingprogramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiMarketingprogramID

`func (o *DocumenttypechecklistResponse) SetFkiMarketingprogramID(v int32)`

SetFkiMarketingprogramID sets FkiMarketingprogramID field to given value.

### HasFkiMarketingprogramID

`func (o *DocumenttypechecklistResponse) HasFkiMarketingprogramID() bool`

HasFkiMarketingprogramID returns a boolean if a field has been set.

### GetFkiMarketingfollowID

`func (o *DocumenttypechecklistResponse) GetFkiMarketingfollowID() int32`

GetFkiMarketingfollowID returns the FkiMarketingfollowID field if non-nil, zero value otherwise.

### GetFkiMarketingfollowIDOk

`func (o *DocumenttypechecklistResponse) GetFkiMarketingfollowIDOk() (*int32, bool)`

GetFkiMarketingfollowIDOk returns a tuple with the FkiMarketingfollowID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiMarketingfollowID

`func (o *DocumenttypechecklistResponse) SetFkiMarketingfollowID(v int32)`

SetFkiMarketingfollowID sets FkiMarketingfollowID field to given value.

### HasFkiMarketingfollowID

`func (o *DocumenttypechecklistResponse) HasFkiMarketingfollowID() bool`

HasFkiMarketingfollowID returns a boolean if a field has been set.

### GetFkiNotaryID

`func (o *DocumenttypechecklistResponse) GetFkiNotaryID() int32`

GetFkiNotaryID returns the FkiNotaryID field if non-nil, zero value otherwise.

### GetFkiNotaryIDOk

`func (o *DocumenttypechecklistResponse) GetFkiNotaryIDOk() (*int32, bool)`

GetFkiNotaryIDOk returns a tuple with the FkiNotaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiNotaryID

`func (o *DocumenttypechecklistResponse) SetFkiNotaryID(v int32)`

SetFkiNotaryID sets FkiNotaryID field to given value.

### HasFkiNotaryID

`func (o *DocumenttypechecklistResponse) HasFkiNotaryID() bool`

HasFkiNotaryID returns a boolean if a field has been set.

### GetFkiOfficetaxreportID

`func (o *DocumenttypechecklistResponse) GetFkiOfficetaxreportID() int32`

GetFkiOfficetaxreportID returns the FkiOfficetaxreportID field if non-nil, zero value otherwise.

### GetFkiOfficetaxreportIDOk

`func (o *DocumenttypechecklistResponse) GetFkiOfficetaxreportIDOk() (*int32, bool)`

GetFkiOfficetaxreportIDOk returns a tuple with the FkiOfficetaxreportID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiOfficetaxreportID

`func (o *DocumenttypechecklistResponse) SetFkiOfficetaxreportID(v int32)`

SetFkiOfficetaxreportID sets FkiOfficetaxreportID field to given value.

### HasFkiOfficetaxreportID

`func (o *DocumenttypechecklistResponse) HasFkiOfficetaxreportID() bool`

HasFkiOfficetaxreportID returns a boolean if a field has been set.

### GetFkiOtherincomeID

`func (o *DocumenttypechecklistResponse) GetFkiOtherincomeID() int32`

GetFkiOtherincomeID returns the FkiOtherincomeID field if non-nil, zero value otherwise.

### GetFkiOtherincomeIDOk

`func (o *DocumenttypechecklistResponse) GetFkiOtherincomeIDOk() (*int32, bool)`

GetFkiOtherincomeIDOk returns a tuple with the FkiOtherincomeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiOtherincomeID

`func (o *DocumenttypechecklistResponse) SetFkiOtherincomeID(v int32)`

SetFkiOtherincomeID sets FkiOtherincomeID field to given value.

### HasFkiOtherincomeID

`func (o *DocumenttypechecklistResponse) HasFkiOtherincomeID() bool`

HasFkiOtherincomeID returns a boolean if a field has been set.

### GetFkiPaymentpreparationID

`func (o *DocumenttypechecklistResponse) GetFkiPaymentpreparationID() int32`

GetFkiPaymentpreparationID returns the FkiPaymentpreparationID field if non-nil, zero value otherwise.

### GetFkiPaymentpreparationIDOk

`func (o *DocumenttypechecklistResponse) GetFkiPaymentpreparationIDOk() (*int32, bool)`

GetFkiPaymentpreparationIDOk returns a tuple with the FkiPaymentpreparationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPaymentpreparationID

`func (o *DocumenttypechecklistResponse) SetFkiPaymentpreparationID(v int32)`

SetFkiPaymentpreparationID sets FkiPaymentpreparationID field to given value.

### HasFkiPaymentpreparationID

`func (o *DocumenttypechecklistResponse) HasFkiPaymentpreparationID() bool`

HasFkiPaymentpreparationID returns a boolean if a field has been set.

### GetFkiPurchaseID

`func (o *DocumenttypechecklistResponse) GetFkiPurchaseID() int32`

GetFkiPurchaseID returns the FkiPurchaseID field if non-nil, zero value otherwise.

### GetFkiPurchaseIDOk

`func (o *DocumenttypechecklistResponse) GetFkiPurchaseIDOk() (*int32, bool)`

GetFkiPurchaseIDOk returns a tuple with the FkiPurchaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPurchaseID

`func (o *DocumenttypechecklistResponse) SetFkiPurchaseID(v int32)`

SetFkiPurchaseID sets FkiPurchaseID field to given value.

### HasFkiPurchaseID

`func (o *DocumenttypechecklistResponse) HasFkiPurchaseID() bool`

HasFkiPurchaseID returns a boolean if a field has been set.

### GetFkiSalaryID

`func (o *DocumenttypechecklistResponse) GetFkiSalaryID() int32`

GetFkiSalaryID returns the FkiSalaryID field if non-nil, zero value otherwise.

### GetFkiSalaryIDOk

`func (o *DocumenttypechecklistResponse) GetFkiSalaryIDOk() (*int32, bool)`

GetFkiSalaryIDOk returns a tuple with the FkiSalaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSalaryID

`func (o *DocumenttypechecklistResponse) SetFkiSalaryID(v int32)`

SetFkiSalaryID sets FkiSalaryID field to given value.

### HasFkiSalaryID

`func (o *DocumenttypechecklistResponse) HasFkiSalaryID() bool`

HasFkiSalaryID returns a boolean if a field has been set.

### GetFkiSupplierID

`func (o *DocumenttypechecklistResponse) GetFkiSupplierID() int32`

GetFkiSupplierID returns the FkiSupplierID field if non-nil, zero value otherwise.

### GetFkiSupplierIDOk

`func (o *DocumenttypechecklistResponse) GetFkiSupplierIDOk() (*int32, bool)`

GetFkiSupplierIDOk returns a tuple with the FkiSupplierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSupplierID

`func (o *DocumenttypechecklistResponse) SetFkiSupplierID(v int32)`

SetFkiSupplierID sets FkiSupplierID field to given value.

### HasFkiSupplierID

`func (o *DocumenttypechecklistResponse) HasFkiSupplierID() bool`

HasFkiSupplierID returns a boolean if a field has been set.

### GetFkiTranqcontractID

`func (o *DocumenttypechecklistResponse) GetFkiTranqcontractID() int32`

GetFkiTranqcontractID returns the FkiTranqcontractID field if non-nil, zero value otherwise.

### GetFkiTranqcontractIDOk

`func (o *DocumenttypechecklistResponse) GetFkiTranqcontractIDOk() (*int32, bool)`

GetFkiTranqcontractIDOk returns a tuple with the FkiTranqcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTranqcontractID

`func (o *DocumenttypechecklistResponse) SetFkiTranqcontractID(v int32)`

SetFkiTranqcontractID sets FkiTranqcontractID field to given value.

### HasFkiTranqcontractID

`func (o *DocumenttypechecklistResponse) HasFkiTranqcontractID() bool`

HasFkiTranqcontractID returns a boolean if a field has been set.

### GetFkiTemplateID

`func (o *DocumenttypechecklistResponse) GetFkiTemplateID() int32`

GetFkiTemplateID returns the FkiTemplateID field if non-nil, zero value otherwise.

### GetFkiTemplateIDOk

`func (o *DocumenttypechecklistResponse) GetFkiTemplateIDOk() (*int32, bool)`

GetFkiTemplateIDOk returns a tuple with the FkiTemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTemplateID

`func (o *DocumenttypechecklistResponse) SetFkiTemplateID(v int32)`

SetFkiTemplateID sets FkiTemplateID field to given value.

### HasFkiTemplateID

`func (o *DocumenttypechecklistResponse) HasFkiTemplateID() bool`

HasFkiTemplateID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


