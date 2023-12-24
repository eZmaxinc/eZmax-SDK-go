# AttachmentResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiAttachmentID** | **int32** | The unique ID of the Attachment. | 
**FkiComputerID** | Pointer to **int32** | The unique ID of the Computer | [optional] 
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
**FkiInscriptionchecklistID** | Pointer to **int32** | The unique ID of the Inscriptionchecklist | [optional] 
**FkiFolderID** | Pointer to **int32** | The unique ID of the Folder | [optional] 
**FkiRejectedoffertopurchaseID** | Pointer to **int32** | The unique ID of the Rejectedoffertopurchase | [optional] 
**FkiDisclosureID** | Pointer to **int32** | The unique ID of the Disclosure | [optional] 
**FkiReconciliationID** | Pointer to **int32** | The unique ID of the Reconciliation | [optional] 
**FkiEzsigndocumentIDReference** | Pointer to **int32** | The unique ID of the Ezsigndocument | [optional] 
**EAttachmentDocumenttype** | [**FieldEAttachmentDocumenttype**](FieldEAttachmentDocumenttype.md) |  | 
**SAttachmentName** | **string** | The name of the Attachment | 
**EAttachmentPrivacy** | [**FieldEAttachmentPrivacy**](FieldEAttachmentPrivacy.md) |  | 
**FkiUserIDSpecific** | Pointer to **int32** | The unique ID of the User | [optional] 
**EAttachmentType** | [**FieldEAttachmentType**](FieldEAttachmentType.md) |  | 
**IAttachmentSize** | **int32** | The size of the Attachment | 
**IAttachmentEDMmoduleflag** | Pointer to **int32** | The edmmoduleflag of the Attachment | [optional] 
**SAttachmentMD5** | **string** | The md5 of the Attachment | 
**BAttachmentDeleted** | **bool** | Whether if it&#39;s deleted | 
**BAttachmentValid** | **bool** | Whether if it&#39;s valid | 
**EAttachmentVerified** | [**FieldEAttachmentVerified**](FieldEAttachmentVerified.md) |  | 
**TAttachmentRejectioncomment** | Pointer to **string** | The rejectioncomment of the Attachment | [optional] 
**FkiUserIDOwner** | Pointer to **int32** | The unique ID of the User | [optional] 
**ObjAudit** | Pointer to [**CommonAudit**](CommonAudit.md) |  | [optional] 

## Methods

### NewAttachmentResponseCompound

`func NewAttachmentResponseCompound(pkiAttachmentID int32, eAttachmentDocumenttype FieldEAttachmentDocumenttype, sAttachmentName string, eAttachmentPrivacy FieldEAttachmentPrivacy, eAttachmentType FieldEAttachmentType, iAttachmentSize int32, sAttachmentMD5 string, bAttachmentDeleted bool, bAttachmentValid bool, eAttachmentVerified FieldEAttachmentVerified, ) *AttachmentResponseCompound`

NewAttachmentResponseCompound instantiates a new AttachmentResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentResponseCompoundWithDefaults

`func NewAttachmentResponseCompoundWithDefaults() *AttachmentResponseCompound`

NewAttachmentResponseCompoundWithDefaults instantiates a new AttachmentResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiAttachmentID

`func (o *AttachmentResponseCompound) GetPkiAttachmentID() int32`

GetPkiAttachmentID returns the PkiAttachmentID field if non-nil, zero value otherwise.

### GetPkiAttachmentIDOk

`func (o *AttachmentResponseCompound) GetPkiAttachmentIDOk() (*int32, bool)`

GetPkiAttachmentIDOk returns a tuple with the PkiAttachmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiAttachmentID

`func (o *AttachmentResponseCompound) SetPkiAttachmentID(v int32)`

SetPkiAttachmentID sets PkiAttachmentID field to given value.


### GetFkiComputerID

`func (o *AttachmentResponseCompound) GetFkiComputerID() int32`

GetFkiComputerID returns the FkiComputerID field if non-nil, zero value otherwise.

### GetFkiComputerIDOk

`func (o *AttachmentResponseCompound) GetFkiComputerIDOk() (*int32, bool)`

GetFkiComputerIDOk returns a tuple with the FkiComputerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiComputerID

`func (o *AttachmentResponseCompound) SetFkiComputerID(v int32)`

SetFkiComputerID sets FkiComputerID field to given value.

### HasFkiComputerID

`func (o *AttachmentResponseCompound) HasFkiComputerID() bool`

HasFkiComputerID returns a boolean if a field has been set.

### GetFkiAdjustmentID

`func (o *AttachmentResponseCompound) GetFkiAdjustmentID() int32`

GetFkiAdjustmentID returns the FkiAdjustmentID field if non-nil, zero value otherwise.

### GetFkiAdjustmentIDOk

`func (o *AttachmentResponseCompound) GetFkiAdjustmentIDOk() (*int32, bool)`

GetFkiAdjustmentIDOk returns a tuple with the FkiAdjustmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAdjustmentID

`func (o *AttachmentResponseCompound) SetFkiAdjustmentID(v int32)`

SetFkiAdjustmentID sets FkiAdjustmentID field to given value.

### HasFkiAdjustmentID

`func (o *AttachmentResponseCompound) HasFkiAdjustmentID() bool`

HasFkiAdjustmentID returns a boolean if a field has been set.

### GetFkiAgentID

`func (o *AttachmentResponseCompound) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *AttachmentResponseCompound) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *AttachmentResponseCompound) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *AttachmentResponseCompound) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBankaccountID

`func (o *AttachmentResponseCompound) GetFkiBankaccountID() int32`

GetFkiBankaccountID returns the FkiBankaccountID field if non-nil, zero value otherwise.

### GetFkiBankaccountIDOk

`func (o *AttachmentResponseCompound) GetFkiBankaccountIDOk() (*int32, bool)`

GetFkiBankaccountIDOk returns a tuple with the FkiBankaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBankaccountID

`func (o *AttachmentResponseCompound) SetFkiBankaccountID(v int32)`

SetFkiBankaccountID sets FkiBankaccountID field to given value.

### HasFkiBankaccountID

`func (o *AttachmentResponseCompound) HasFkiBankaccountID() bool`

HasFkiBankaccountID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *AttachmentResponseCompound) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *AttachmentResponseCompound) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *AttachmentResponseCompound) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *AttachmentResponseCompound) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetFkiCommissionadvanceID

`func (o *AttachmentResponseCompound) GetFkiCommissionadvanceID() int32`

GetFkiCommissionadvanceID returns the FkiCommissionadvanceID field if non-nil, zero value otherwise.

### GetFkiCommissionadvanceIDOk

`func (o *AttachmentResponseCompound) GetFkiCommissionadvanceIDOk() (*int32, bool)`

GetFkiCommissionadvanceIDOk returns a tuple with the FkiCommissionadvanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCommissionadvanceID

`func (o *AttachmentResponseCompound) SetFkiCommissionadvanceID(v int32)`

SetFkiCommissionadvanceID sets FkiCommissionadvanceID field to given value.

### HasFkiCommissionadvanceID

`func (o *AttachmentResponseCompound) HasFkiCommissionadvanceID() bool`

HasFkiCommissionadvanceID returns a boolean if a field has been set.

### GetFkiCommunicationID

`func (o *AttachmentResponseCompound) GetFkiCommunicationID() int32`

GetFkiCommunicationID returns the FkiCommunicationID field if non-nil, zero value otherwise.

### GetFkiCommunicationIDOk

`func (o *AttachmentResponseCompound) GetFkiCommunicationIDOk() (*int32, bool)`

GetFkiCommunicationIDOk returns a tuple with the FkiCommunicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCommunicationID

`func (o *AttachmentResponseCompound) SetFkiCommunicationID(v int32)`

SetFkiCommunicationID sets FkiCommunicationID field to given value.

### HasFkiCommunicationID

`func (o *AttachmentResponseCompound) HasFkiCommunicationID() bool`

HasFkiCommunicationID returns a boolean if a field has been set.

### GetFkiCustomerID

`func (o *AttachmentResponseCompound) GetFkiCustomerID() int32`

GetFkiCustomerID returns the FkiCustomerID field if non-nil, zero value otherwise.

### GetFkiCustomerIDOk

`func (o *AttachmentResponseCompound) GetFkiCustomerIDOk() (*int32, bool)`

GetFkiCustomerIDOk returns a tuple with the FkiCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomerID

`func (o *AttachmentResponseCompound) SetFkiCustomerID(v int32)`

SetFkiCustomerID sets FkiCustomerID field to given value.

### HasFkiCustomerID

`func (o *AttachmentResponseCompound) HasFkiCustomerID() bool`

HasFkiCustomerID returns a boolean if a field has been set.

### GetFkiCustomertemplateID

`func (o *AttachmentResponseCompound) GetFkiCustomertemplateID() int32`

GetFkiCustomertemplateID returns the FkiCustomertemplateID field if non-nil, zero value otherwise.

### GetFkiCustomertemplateIDOk

`func (o *AttachmentResponseCompound) GetFkiCustomertemplateIDOk() (*int32, bool)`

GetFkiCustomertemplateIDOk returns a tuple with the FkiCustomertemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomertemplateID

`func (o *AttachmentResponseCompound) SetFkiCustomertemplateID(v int32)`

SetFkiCustomertemplateID sets FkiCustomertemplateID field to given value.

### HasFkiCustomertemplateID

`func (o *AttachmentResponseCompound) HasFkiCustomertemplateID() bool`

HasFkiCustomertemplateID returns a boolean if a field has been set.

### GetFkiDepositID

`func (o *AttachmentResponseCompound) GetFkiDepositID() int32`

GetFkiDepositID returns the FkiDepositID field if non-nil, zero value otherwise.

### GetFkiDepositIDOk

`func (o *AttachmentResponseCompound) GetFkiDepositIDOk() (*int32, bool)`

GetFkiDepositIDOk returns a tuple with the FkiDepositID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepositID

`func (o *AttachmentResponseCompound) SetFkiDepositID(v int32)`

SetFkiDepositID sets FkiDepositID field to given value.

### HasFkiDepositID

`func (o *AttachmentResponseCompound) HasFkiDepositID() bool`

HasFkiDepositID returns a boolean if a field has been set.

### GetFkiDeposittransitchequeID

`func (o *AttachmentResponseCompound) GetFkiDeposittransitchequeID() int32`

GetFkiDeposittransitchequeID returns the FkiDeposittransitchequeID field if non-nil, zero value otherwise.

### GetFkiDeposittransitchequeIDOk

`func (o *AttachmentResponseCompound) GetFkiDeposittransitchequeIDOk() (*int32, bool)`

GetFkiDeposittransitchequeIDOk returns a tuple with the FkiDeposittransitchequeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDeposittransitchequeID

`func (o *AttachmentResponseCompound) SetFkiDeposittransitchequeID(v int32)`

SetFkiDeposittransitchequeID sets FkiDeposittransitchequeID field to given value.

### HasFkiDeposittransitchequeID

`func (o *AttachmentResponseCompound) HasFkiDeposittransitchequeID() bool`

HasFkiDeposittransitchequeID returns a boolean if a field has been set.

### GetFkiElectronicfundstransferID

`func (o *AttachmentResponseCompound) GetFkiElectronicfundstransferID() int32`

GetFkiElectronicfundstransferID returns the FkiElectronicfundstransferID field if non-nil, zero value otherwise.

### GetFkiElectronicfundstransferIDOk

`func (o *AttachmentResponseCompound) GetFkiElectronicfundstransferIDOk() (*int32, bool)`

GetFkiElectronicfundstransferIDOk returns a tuple with the FkiElectronicfundstransferID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiElectronicfundstransferID

`func (o *AttachmentResponseCompound) SetFkiElectronicfundstransferID(v int32)`

SetFkiElectronicfundstransferID sets FkiElectronicfundstransferID field to given value.

### HasFkiElectronicfundstransferID

`func (o *AttachmentResponseCompound) HasFkiElectronicfundstransferID() bool`

HasFkiElectronicfundstransferID returns a boolean if a field has been set.

### GetFkiEmployeeID

`func (o *AttachmentResponseCompound) GetFkiEmployeeID() int32`

GetFkiEmployeeID returns the FkiEmployeeID field if non-nil, zero value otherwise.

### GetFkiEmployeeIDOk

`func (o *AttachmentResponseCompound) GetFkiEmployeeIDOk() (*int32, bool)`

GetFkiEmployeeIDOk returns a tuple with the FkiEmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmployeeID

`func (o *AttachmentResponseCompound) SetFkiEmployeeID(v int32)`

SetFkiEmployeeID sets FkiEmployeeID field to given value.

### HasFkiEmployeeID

`func (o *AttachmentResponseCompound) HasFkiEmployeeID() bool`

HasFkiEmployeeID returns a boolean if a field has been set.

### GetFkiExternalbrokerID

`func (o *AttachmentResponseCompound) GetFkiExternalbrokerID() int32`

GetFkiExternalbrokerID returns the FkiExternalbrokerID field if non-nil, zero value otherwise.

### GetFkiExternalbrokerIDOk

`func (o *AttachmentResponseCompound) GetFkiExternalbrokerIDOk() (*int32, bool)`

GetFkiExternalbrokerIDOk returns a tuple with the FkiExternalbrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiExternalbrokerID

`func (o *AttachmentResponseCompound) SetFkiExternalbrokerID(v int32)`

SetFkiExternalbrokerID sets FkiExternalbrokerID field to given value.

### HasFkiExternalbrokerID

`func (o *AttachmentResponseCompound) HasFkiExternalbrokerID() bool`

HasFkiExternalbrokerID returns a boolean if a field has been set.

### GetFkiEzcomadvanceserverID

`func (o *AttachmentResponseCompound) GetFkiEzcomadvanceserverID() int32`

GetFkiEzcomadvanceserverID returns the FkiEzcomadvanceserverID field if non-nil, zero value otherwise.

### GetFkiEzcomadvanceserverIDOk

`func (o *AttachmentResponseCompound) GetFkiEzcomadvanceserverIDOk() (*int32, bool)`

GetFkiEzcomadvanceserverIDOk returns a tuple with the FkiEzcomadvanceserverID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzcomadvanceserverID

`func (o *AttachmentResponseCompound) SetFkiEzcomadvanceserverID(v int32)`

SetFkiEzcomadvanceserverID sets FkiEzcomadvanceserverID field to given value.

### HasFkiEzcomadvanceserverID

`func (o *AttachmentResponseCompound) HasFkiEzcomadvanceserverID() bool`

HasFkiEzcomadvanceserverID returns a boolean if a field has been set.

### GetFkiEzcomcompanyID

`func (o *AttachmentResponseCompound) GetFkiEzcomcompanyID() int32`

GetFkiEzcomcompanyID returns the FkiEzcomcompanyID field if non-nil, zero value otherwise.

### GetFkiEzcomcompanyIDOk

`func (o *AttachmentResponseCompound) GetFkiEzcomcompanyIDOk() (*int32, bool)`

GetFkiEzcomcompanyIDOk returns a tuple with the FkiEzcomcompanyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzcomcompanyID

`func (o *AttachmentResponseCompound) SetFkiEzcomcompanyID(v int32)`

SetFkiEzcomcompanyID sets FkiEzcomcompanyID field to given value.

### HasFkiEzcomcompanyID

`func (o *AttachmentResponseCompound) HasFkiEzcomcompanyID() bool`

HasFkiEzcomcompanyID returns a boolean if a field has been set.

### GetFkiEzsigndocumentID

`func (o *AttachmentResponseCompound) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *AttachmentResponseCompound) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *AttachmentResponseCompound) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.

### HasFkiEzsigndocumentID

`func (o *AttachmentResponseCompound) HasFkiEzsigndocumentID() bool`

HasFkiEzsigndocumentID returns a boolean if a field has been set.

### GetFkiGhacqcontractID

`func (o *AttachmentResponseCompound) GetFkiGhacqcontractID() int32`

GetFkiGhacqcontractID returns the FkiGhacqcontractID field if non-nil, zero value otherwise.

### GetFkiGhacqcontractIDOk

`func (o *AttachmentResponseCompound) GetFkiGhacqcontractIDOk() (*int32, bool)`

GetFkiGhacqcontractIDOk returns a tuple with the FkiGhacqcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGhacqcontractID

`func (o *AttachmentResponseCompound) SetFkiGhacqcontractID(v int32)`

SetFkiGhacqcontractID sets FkiGhacqcontractID field to given value.

### HasFkiGhacqcontractID

`func (o *AttachmentResponseCompound) HasFkiGhacqcontractID() bool`

HasFkiGhacqcontractID returns a boolean if a field has been set.

### GetFkiInscriptionID

`func (o *AttachmentResponseCompound) GetFkiInscriptionID() int32`

GetFkiInscriptionID returns the FkiInscriptionID field if non-nil, zero value otherwise.

### GetFkiInscriptionIDOk

`func (o *AttachmentResponseCompound) GetFkiInscriptionIDOk() (*int32, bool)`

GetFkiInscriptionIDOk returns a tuple with the FkiInscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionID

`func (o *AttachmentResponseCompound) SetFkiInscriptionID(v int32)`

SetFkiInscriptionID sets FkiInscriptionID field to given value.

### HasFkiInscriptionID

`func (o *AttachmentResponseCompound) HasFkiInscriptionID() bool`

HasFkiInscriptionID returns a boolean if a field has been set.

### GetFkiInscriptiontempID

`func (o *AttachmentResponseCompound) GetFkiInscriptiontempID() int32`

GetFkiInscriptiontempID returns the FkiInscriptiontempID field if non-nil, zero value otherwise.

### GetFkiInscriptiontempIDOk

`func (o *AttachmentResponseCompound) GetFkiInscriptiontempIDOk() (*int32, bool)`

GetFkiInscriptiontempIDOk returns a tuple with the FkiInscriptiontempID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptiontempID

`func (o *AttachmentResponseCompound) SetFkiInscriptiontempID(v int32)`

SetFkiInscriptiontempID sets FkiInscriptiontempID field to given value.

### HasFkiInscriptiontempID

`func (o *AttachmentResponseCompound) HasFkiInscriptiontempID() bool`

HasFkiInscriptiontempID returns a boolean if a field has been set.

### GetFkiInscriptionnotauthenticatedID

`func (o *AttachmentResponseCompound) GetFkiInscriptionnotauthenticatedID() int32`

GetFkiInscriptionnotauthenticatedID returns the FkiInscriptionnotauthenticatedID field if non-nil, zero value otherwise.

### GetFkiInscriptionnotauthenticatedIDOk

`func (o *AttachmentResponseCompound) GetFkiInscriptionnotauthenticatedIDOk() (*int32, bool)`

GetFkiInscriptionnotauthenticatedIDOk returns a tuple with the FkiInscriptionnotauthenticatedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionnotauthenticatedID

`func (o *AttachmentResponseCompound) SetFkiInscriptionnotauthenticatedID(v int32)`

SetFkiInscriptionnotauthenticatedID sets FkiInscriptionnotauthenticatedID field to given value.

### HasFkiInscriptionnotauthenticatedID

`func (o *AttachmentResponseCompound) HasFkiInscriptionnotauthenticatedID() bool`

HasFkiInscriptionnotauthenticatedID returns a boolean if a field has been set.

### GetFkiInvoiceID

`func (o *AttachmentResponseCompound) GetFkiInvoiceID() int32`

GetFkiInvoiceID returns the FkiInvoiceID field if non-nil, zero value otherwise.

### GetFkiInvoiceIDOk

`func (o *AttachmentResponseCompound) GetFkiInvoiceIDOk() (*int32, bool)`

GetFkiInvoiceIDOk returns a tuple with the FkiInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInvoiceID

`func (o *AttachmentResponseCompound) SetFkiInvoiceID(v int32)`

SetFkiInvoiceID sets FkiInvoiceID field to given value.

### HasFkiInvoiceID

`func (o *AttachmentResponseCompound) HasFkiInvoiceID() bool`

HasFkiInvoiceID returns a boolean if a field has been set.

### GetFkiBuyercontractID

`func (o *AttachmentResponseCompound) GetFkiBuyercontractID() int32`

GetFkiBuyercontractID returns the FkiBuyercontractID field if non-nil, zero value otherwise.

### GetFkiBuyercontractIDOk

`func (o *AttachmentResponseCompound) GetFkiBuyercontractIDOk() (*int32, bool)`

GetFkiBuyercontractIDOk returns a tuple with the FkiBuyercontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBuyercontractID

`func (o *AttachmentResponseCompound) SetFkiBuyercontractID(v int32)`

SetFkiBuyercontractID sets FkiBuyercontractID field to given value.

### HasFkiBuyercontractID

`func (o *AttachmentResponseCompound) HasFkiBuyercontractID() bool`

HasFkiBuyercontractID returns a boolean if a field has been set.

### GetFkiFranchisebrokerID

`func (o *AttachmentResponseCompound) GetFkiFranchisebrokerID() int32`

GetFkiFranchisebrokerID returns the FkiFranchisebrokerID field if non-nil, zero value otherwise.

### GetFkiFranchisebrokerIDOk

`func (o *AttachmentResponseCompound) GetFkiFranchisebrokerIDOk() (*int32, bool)`

GetFkiFranchisebrokerIDOk returns a tuple with the FkiFranchisebrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisebrokerID

`func (o *AttachmentResponseCompound) SetFkiFranchisebrokerID(v int32)`

SetFkiFranchisebrokerID sets FkiFranchisebrokerID field to given value.

### HasFkiFranchisebrokerID

`func (o *AttachmentResponseCompound) HasFkiFranchisebrokerID() bool`

HasFkiFranchisebrokerID returns a boolean if a field has been set.

### GetFkiFranchiseagenceID

`func (o *AttachmentResponseCompound) GetFkiFranchiseagenceID() int32`

GetFkiFranchiseagenceID returns the FkiFranchiseagenceID field if non-nil, zero value otherwise.

### GetFkiFranchiseagenceIDOk

`func (o *AttachmentResponseCompound) GetFkiFranchiseagenceIDOk() (*int32, bool)`

GetFkiFranchiseagenceIDOk returns a tuple with the FkiFranchiseagenceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseagenceID

`func (o *AttachmentResponseCompound) SetFkiFranchiseagenceID(v int32)`

SetFkiFranchiseagenceID sets FkiFranchiseagenceID field to given value.

### HasFkiFranchiseagenceID

`func (o *AttachmentResponseCompound) HasFkiFranchiseagenceID() bool`

HasFkiFranchiseagenceID returns a boolean if a field has been set.

### GetFkiFranchiseofficeID

`func (o *AttachmentResponseCompound) GetFkiFranchiseofficeID() int32`

GetFkiFranchiseofficeID returns the FkiFranchiseofficeID field if non-nil, zero value otherwise.

### GetFkiFranchiseofficeIDOk

`func (o *AttachmentResponseCompound) GetFkiFranchiseofficeIDOk() (*int32, bool)`

GetFkiFranchiseofficeIDOk returns a tuple with the FkiFranchiseofficeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseofficeID

`func (o *AttachmentResponseCompound) SetFkiFranchiseofficeID(v int32)`

SetFkiFranchiseofficeID sets FkiFranchiseofficeID field to given value.

### HasFkiFranchiseofficeID

`func (o *AttachmentResponseCompound) HasFkiFranchiseofficeID() bool`

HasFkiFranchiseofficeID returns a boolean if a field has been set.

### GetFkiFranchisefranchiseID

`func (o *AttachmentResponseCompound) GetFkiFranchisefranchiseID() int32`

GetFkiFranchisefranchiseID returns the FkiFranchisefranchiseID field if non-nil, zero value otherwise.

### GetFkiFranchisefranchiseIDOk

`func (o *AttachmentResponseCompound) GetFkiFranchisefranchiseIDOk() (*int32, bool)`

GetFkiFranchisefranchiseIDOk returns a tuple with the FkiFranchisefranchiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisefranchiseID

`func (o *AttachmentResponseCompound) SetFkiFranchisefranchiseID(v int32)`

SetFkiFranchisefranchiseID sets FkiFranchisefranchiseID field to given value.

### HasFkiFranchisefranchiseID

`func (o *AttachmentResponseCompound) HasFkiFranchisefranchiseID() bool`

HasFkiFranchisefranchiseID returns a boolean if a field has been set.

### GetFkiFranchisecomplaintID

`func (o *AttachmentResponseCompound) GetFkiFranchisecomplaintID() int32`

GetFkiFranchisecomplaintID returns the FkiFranchisecomplaintID field if non-nil, zero value otherwise.

### GetFkiFranchisecomplaintIDOk

`func (o *AttachmentResponseCompound) GetFkiFranchisecomplaintIDOk() (*int32, bool)`

GetFkiFranchisecomplaintIDOk returns a tuple with the FkiFranchisecomplaintID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisecomplaintID

`func (o *AttachmentResponseCompound) SetFkiFranchisecomplaintID(v int32)`

SetFkiFranchisecomplaintID sets FkiFranchisecomplaintID field to given value.

### HasFkiFranchisecomplaintID

`func (o *AttachmentResponseCompound) HasFkiFranchisecomplaintID() bool`

HasFkiFranchisecomplaintID returns a boolean if a field has been set.

### GetFkiLeadID

`func (o *AttachmentResponseCompound) GetFkiLeadID() int32`

GetFkiLeadID returns the FkiLeadID field if non-nil, zero value otherwise.

### GetFkiLeadIDOk

`func (o *AttachmentResponseCompound) GetFkiLeadIDOk() (*int32, bool)`

GetFkiLeadIDOk returns a tuple with the FkiLeadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLeadID

`func (o *AttachmentResponseCompound) SetFkiLeadID(v int32)`

SetFkiLeadID sets FkiLeadID field to given value.

### HasFkiLeadID

`func (o *AttachmentResponseCompound) HasFkiLeadID() bool`

HasFkiLeadID returns a boolean if a field has been set.

### GetFkiMarketingprogramID

`func (o *AttachmentResponseCompound) GetFkiMarketingprogramID() int32`

GetFkiMarketingprogramID returns the FkiMarketingprogramID field if non-nil, zero value otherwise.

### GetFkiMarketingprogramIDOk

`func (o *AttachmentResponseCompound) GetFkiMarketingprogramIDOk() (*int32, bool)`

GetFkiMarketingprogramIDOk returns a tuple with the FkiMarketingprogramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiMarketingprogramID

`func (o *AttachmentResponseCompound) SetFkiMarketingprogramID(v int32)`

SetFkiMarketingprogramID sets FkiMarketingprogramID field to given value.

### HasFkiMarketingprogramID

`func (o *AttachmentResponseCompound) HasFkiMarketingprogramID() bool`

HasFkiMarketingprogramID returns a boolean if a field has been set.

### GetFkiMarketingfollowID

`func (o *AttachmentResponseCompound) GetFkiMarketingfollowID() int32`

GetFkiMarketingfollowID returns the FkiMarketingfollowID field if non-nil, zero value otherwise.

### GetFkiMarketingfollowIDOk

`func (o *AttachmentResponseCompound) GetFkiMarketingfollowIDOk() (*int32, bool)`

GetFkiMarketingfollowIDOk returns a tuple with the FkiMarketingfollowID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiMarketingfollowID

`func (o *AttachmentResponseCompound) SetFkiMarketingfollowID(v int32)`

SetFkiMarketingfollowID sets FkiMarketingfollowID field to given value.

### HasFkiMarketingfollowID

`func (o *AttachmentResponseCompound) HasFkiMarketingfollowID() bool`

HasFkiMarketingfollowID returns a boolean if a field has been set.

### GetFkiNotaryID

`func (o *AttachmentResponseCompound) GetFkiNotaryID() int32`

GetFkiNotaryID returns the FkiNotaryID field if non-nil, zero value otherwise.

### GetFkiNotaryIDOk

`func (o *AttachmentResponseCompound) GetFkiNotaryIDOk() (*int32, bool)`

GetFkiNotaryIDOk returns a tuple with the FkiNotaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiNotaryID

`func (o *AttachmentResponseCompound) SetFkiNotaryID(v int32)`

SetFkiNotaryID sets FkiNotaryID field to given value.

### HasFkiNotaryID

`func (o *AttachmentResponseCompound) HasFkiNotaryID() bool`

HasFkiNotaryID returns a boolean if a field has been set.

### GetFkiOfficetaxreportID

`func (o *AttachmentResponseCompound) GetFkiOfficetaxreportID() int32`

GetFkiOfficetaxreportID returns the FkiOfficetaxreportID field if non-nil, zero value otherwise.

### GetFkiOfficetaxreportIDOk

`func (o *AttachmentResponseCompound) GetFkiOfficetaxreportIDOk() (*int32, bool)`

GetFkiOfficetaxreportIDOk returns a tuple with the FkiOfficetaxreportID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiOfficetaxreportID

`func (o *AttachmentResponseCompound) SetFkiOfficetaxreportID(v int32)`

SetFkiOfficetaxreportID sets FkiOfficetaxreportID field to given value.

### HasFkiOfficetaxreportID

`func (o *AttachmentResponseCompound) HasFkiOfficetaxreportID() bool`

HasFkiOfficetaxreportID returns a boolean if a field has been set.

### GetFkiOtherincomeID

`func (o *AttachmentResponseCompound) GetFkiOtherincomeID() int32`

GetFkiOtherincomeID returns the FkiOtherincomeID field if non-nil, zero value otherwise.

### GetFkiOtherincomeIDOk

`func (o *AttachmentResponseCompound) GetFkiOtherincomeIDOk() (*int32, bool)`

GetFkiOtherincomeIDOk returns a tuple with the FkiOtherincomeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiOtherincomeID

`func (o *AttachmentResponseCompound) SetFkiOtherincomeID(v int32)`

SetFkiOtherincomeID sets FkiOtherincomeID field to given value.

### HasFkiOtherincomeID

`func (o *AttachmentResponseCompound) HasFkiOtherincomeID() bool`

HasFkiOtherincomeID returns a boolean if a field has been set.

### GetFkiPaymentpreparationID

`func (o *AttachmentResponseCompound) GetFkiPaymentpreparationID() int32`

GetFkiPaymentpreparationID returns the FkiPaymentpreparationID field if non-nil, zero value otherwise.

### GetFkiPaymentpreparationIDOk

`func (o *AttachmentResponseCompound) GetFkiPaymentpreparationIDOk() (*int32, bool)`

GetFkiPaymentpreparationIDOk returns a tuple with the FkiPaymentpreparationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPaymentpreparationID

`func (o *AttachmentResponseCompound) SetFkiPaymentpreparationID(v int32)`

SetFkiPaymentpreparationID sets FkiPaymentpreparationID field to given value.

### HasFkiPaymentpreparationID

`func (o *AttachmentResponseCompound) HasFkiPaymentpreparationID() bool`

HasFkiPaymentpreparationID returns a boolean if a field has been set.

### GetFkiPurchaseID

`func (o *AttachmentResponseCompound) GetFkiPurchaseID() int32`

GetFkiPurchaseID returns the FkiPurchaseID field if non-nil, zero value otherwise.

### GetFkiPurchaseIDOk

`func (o *AttachmentResponseCompound) GetFkiPurchaseIDOk() (*int32, bool)`

GetFkiPurchaseIDOk returns a tuple with the FkiPurchaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPurchaseID

`func (o *AttachmentResponseCompound) SetFkiPurchaseID(v int32)`

SetFkiPurchaseID sets FkiPurchaseID field to given value.

### HasFkiPurchaseID

`func (o *AttachmentResponseCompound) HasFkiPurchaseID() bool`

HasFkiPurchaseID returns a boolean if a field has been set.

### GetFkiSalaryID

`func (o *AttachmentResponseCompound) GetFkiSalaryID() int32`

GetFkiSalaryID returns the FkiSalaryID field if non-nil, zero value otherwise.

### GetFkiSalaryIDOk

`func (o *AttachmentResponseCompound) GetFkiSalaryIDOk() (*int32, bool)`

GetFkiSalaryIDOk returns a tuple with the FkiSalaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSalaryID

`func (o *AttachmentResponseCompound) SetFkiSalaryID(v int32)`

SetFkiSalaryID sets FkiSalaryID field to given value.

### HasFkiSalaryID

`func (o *AttachmentResponseCompound) HasFkiSalaryID() bool`

HasFkiSalaryID returns a boolean if a field has been set.

### GetFkiSupplierID

`func (o *AttachmentResponseCompound) GetFkiSupplierID() int32`

GetFkiSupplierID returns the FkiSupplierID field if non-nil, zero value otherwise.

### GetFkiSupplierIDOk

`func (o *AttachmentResponseCompound) GetFkiSupplierIDOk() (*int32, bool)`

GetFkiSupplierIDOk returns a tuple with the FkiSupplierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSupplierID

`func (o *AttachmentResponseCompound) SetFkiSupplierID(v int32)`

SetFkiSupplierID sets FkiSupplierID field to given value.

### HasFkiSupplierID

`func (o *AttachmentResponseCompound) HasFkiSupplierID() bool`

HasFkiSupplierID returns a boolean if a field has been set.

### GetFkiTranqcontractID

`func (o *AttachmentResponseCompound) GetFkiTranqcontractID() int32`

GetFkiTranqcontractID returns the FkiTranqcontractID field if non-nil, zero value otherwise.

### GetFkiTranqcontractIDOk

`func (o *AttachmentResponseCompound) GetFkiTranqcontractIDOk() (*int32, bool)`

GetFkiTranqcontractIDOk returns a tuple with the FkiTranqcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTranqcontractID

`func (o *AttachmentResponseCompound) SetFkiTranqcontractID(v int32)`

SetFkiTranqcontractID sets FkiTranqcontractID field to given value.

### HasFkiTranqcontractID

`func (o *AttachmentResponseCompound) HasFkiTranqcontractID() bool`

HasFkiTranqcontractID returns a boolean if a field has been set.

### GetFkiTemplateID

`func (o *AttachmentResponseCompound) GetFkiTemplateID() int32`

GetFkiTemplateID returns the FkiTemplateID field if non-nil, zero value otherwise.

### GetFkiTemplateIDOk

`func (o *AttachmentResponseCompound) GetFkiTemplateIDOk() (*int32, bool)`

GetFkiTemplateIDOk returns a tuple with the FkiTemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTemplateID

`func (o *AttachmentResponseCompound) SetFkiTemplateID(v int32)`

SetFkiTemplateID sets FkiTemplateID field to given value.

### HasFkiTemplateID

`func (o *AttachmentResponseCompound) HasFkiTemplateID() bool`

HasFkiTemplateID returns a boolean if a field has been set.

### GetFkiInscriptionchecklistID

`func (o *AttachmentResponseCompound) GetFkiInscriptionchecklistID() int32`

GetFkiInscriptionchecklistID returns the FkiInscriptionchecklistID field if non-nil, zero value otherwise.

### GetFkiInscriptionchecklistIDOk

`func (o *AttachmentResponseCompound) GetFkiInscriptionchecklistIDOk() (*int32, bool)`

GetFkiInscriptionchecklistIDOk returns a tuple with the FkiInscriptionchecklistID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionchecklistID

`func (o *AttachmentResponseCompound) SetFkiInscriptionchecklistID(v int32)`

SetFkiInscriptionchecklistID sets FkiInscriptionchecklistID field to given value.

### HasFkiInscriptionchecklistID

`func (o *AttachmentResponseCompound) HasFkiInscriptionchecklistID() bool`

HasFkiInscriptionchecklistID returns a boolean if a field has been set.

### GetFkiFolderID

`func (o *AttachmentResponseCompound) GetFkiFolderID() int32`

GetFkiFolderID returns the FkiFolderID field if non-nil, zero value otherwise.

### GetFkiFolderIDOk

`func (o *AttachmentResponseCompound) GetFkiFolderIDOk() (*int32, bool)`

GetFkiFolderIDOk returns a tuple with the FkiFolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFolderID

`func (o *AttachmentResponseCompound) SetFkiFolderID(v int32)`

SetFkiFolderID sets FkiFolderID field to given value.

### HasFkiFolderID

`func (o *AttachmentResponseCompound) HasFkiFolderID() bool`

HasFkiFolderID returns a boolean if a field has been set.

### GetFkiRejectedoffertopurchaseID

`func (o *AttachmentResponseCompound) GetFkiRejectedoffertopurchaseID() int32`

GetFkiRejectedoffertopurchaseID returns the FkiRejectedoffertopurchaseID field if non-nil, zero value otherwise.

### GetFkiRejectedoffertopurchaseIDOk

`func (o *AttachmentResponseCompound) GetFkiRejectedoffertopurchaseIDOk() (*int32, bool)`

GetFkiRejectedoffertopurchaseIDOk returns a tuple with the FkiRejectedoffertopurchaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiRejectedoffertopurchaseID

`func (o *AttachmentResponseCompound) SetFkiRejectedoffertopurchaseID(v int32)`

SetFkiRejectedoffertopurchaseID sets FkiRejectedoffertopurchaseID field to given value.

### HasFkiRejectedoffertopurchaseID

`func (o *AttachmentResponseCompound) HasFkiRejectedoffertopurchaseID() bool`

HasFkiRejectedoffertopurchaseID returns a boolean if a field has been set.

### GetFkiDisclosureID

`func (o *AttachmentResponseCompound) GetFkiDisclosureID() int32`

GetFkiDisclosureID returns the FkiDisclosureID field if non-nil, zero value otherwise.

### GetFkiDisclosureIDOk

`func (o *AttachmentResponseCompound) GetFkiDisclosureIDOk() (*int32, bool)`

GetFkiDisclosureIDOk returns a tuple with the FkiDisclosureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDisclosureID

`func (o *AttachmentResponseCompound) SetFkiDisclosureID(v int32)`

SetFkiDisclosureID sets FkiDisclosureID field to given value.

### HasFkiDisclosureID

`func (o *AttachmentResponseCompound) HasFkiDisclosureID() bool`

HasFkiDisclosureID returns a boolean if a field has been set.

### GetFkiReconciliationID

`func (o *AttachmentResponseCompound) GetFkiReconciliationID() int32`

GetFkiReconciliationID returns the FkiReconciliationID field if non-nil, zero value otherwise.

### GetFkiReconciliationIDOk

`func (o *AttachmentResponseCompound) GetFkiReconciliationIDOk() (*int32, bool)`

GetFkiReconciliationIDOk returns a tuple with the FkiReconciliationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiReconciliationID

`func (o *AttachmentResponseCompound) SetFkiReconciliationID(v int32)`

SetFkiReconciliationID sets FkiReconciliationID field to given value.

### HasFkiReconciliationID

`func (o *AttachmentResponseCompound) HasFkiReconciliationID() bool`

HasFkiReconciliationID returns a boolean if a field has been set.

### GetFkiEzsigndocumentIDReference

`func (o *AttachmentResponseCompound) GetFkiEzsigndocumentIDReference() int32`

GetFkiEzsigndocumentIDReference returns the FkiEzsigndocumentIDReference field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDReferenceOk

`func (o *AttachmentResponseCompound) GetFkiEzsigndocumentIDReferenceOk() (*int32, bool)`

GetFkiEzsigndocumentIDReferenceOk returns a tuple with the FkiEzsigndocumentIDReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentIDReference

`func (o *AttachmentResponseCompound) SetFkiEzsigndocumentIDReference(v int32)`

SetFkiEzsigndocumentIDReference sets FkiEzsigndocumentIDReference field to given value.

### HasFkiEzsigndocumentIDReference

`func (o *AttachmentResponseCompound) HasFkiEzsigndocumentIDReference() bool`

HasFkiEzsigndocumentIDReference returns a boolean if a field has been set.

### GetEAttachmentDocumenttype

`func (o *AttachmentResponseCompound) GetEAttachmentDocumenttype() FieldEAttachmentDocumenttype`

GetEAttachmentDocumenttype returns the EAttachmentDocumenttype field if non-nil, zero value otherwise.

### GetEAttachmentDocumenttypeOk

`func (o *AttachmentResponseCompound) GetEAttachmentDocumenttypeOk() (*FieldEAttachmentDocumenttype, bool)`

GetEAttachmentDocumenttypeOk returns a tuple with the EAttachmentDocumenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentDocumenttype

`func (o *AttachmentResponseCompound) SetEAttachmentDocumenttype(v FieldEAttachmentDocumenttype)`

SetEAttachmentDocumenttype sets EAttachmentDocumenttype field to given value.


### GetSAttachmentName

`func (o *AttachmentResponseCompound) GetSAttachmentName() string`

GetSAttachmentName returns the SAttachmentName field if non-nil, zero value otherwise.

### GetSAttachmentNameOk

`func (o *AttachmentResponseCompound) GetSAttachmentNameOk() (*string, bool)`

GetSAttachmentNameOk returns a tuple with the SAttachmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentName

`func (o *AttachmentResponseCompound) SetSAttachmentName(v string)`

SetSAttachmentName sets SAttachmentName field to given value.


### GetEAttachmentPrivacy

`func (o *AttachmentResponseCompound) GetEAttachmentPrivacy() FieldEAttachmentPrivacy`

GetEAttachmentPrivacy returns the EAttachmentPrivacy field if non-nil, zero value otherwise.

### GetEAttachmentPrivacyOk

`func (o *AttachmentResponseCompound) GetEAttachmentPrivacyOk() (*FieldEAttachmentPrivacy, bool)`

GetEAttachmentPrivacyOk returns a tuple with the EAttachmentPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentPrivacy

`func (o *AttachmentResponseCompound) SetEAttachmentPrivacy(v FieldEAttachmentPrivacy)`

SetEAttachmentPrivacy sets EAttachmentPrivacy field to given value.


### GetFkiUserIDSpecific

`func (o *AttachmentResponseCompound) GetFkiUserIDSpecific() int32`

GetFkiUserIDSpecific returns the FkiUserIDSpecific field if non-nil, zero value otherwise.

### GetFkiUserIDSpecificOk

`func (o *AttachmentResponseCompound) GetFkiUserIDSpecificOk() (*int32, bool)`

GetFkiUserIDSpecificOk returns a tuple with the FkiUserIDSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDSpecific

`func (o *AttachmentResponseCompound) SetFkiUserIDSpecific(v int32)`

SetFkiUserIDSpecific sets FkiUserIDSpecific field to given value.

### HasFkiUserIDSpecific

`func (o *AttachmentResponseCompound) HasFkiUserIDSpecific() bool`

HasFkiUserIDSpecific returns a boolean if a field has been set.

### GetEAttachmentType

`func (o *AttachmentResponseCompound) GetEAttachmentType() FieldEAttachmentType`

GetEAttachmentType returns the EAttachmentType field if non-nil, zero value otherwise.

### GetEAttachmentTypeOk

`func (o *AttachmentResponseCompound) GetEAttachmentTypeOk() (*FieldEAttachmentType, bool)`

GetEAttachmentTypeOk returns a tuple with the EAttachmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentType

`func (o *AttachmentResponseCompound) SetEAttachmentType(v FieldEAttachmentType)`

SetEAttachmentType sets EAttachmentType field to given value.


### GetIAttachmentSize

`func (o *AttachmentResponseCompound) GetIAttachmentSize() int32`

GetIAttachmentSize returns the IAttachmentSize field if non-nil, zero value otherwise.

### GetIAttachmentSizeOk

`func (o *AttachmentResponseCompound) GetIAttachmentSizeOk() (*int32, bool)`

GetIAttachmentSizeOk returns a tuple with the IAttachmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAttachmentSize

`func (o *AttachmentResponseCompound) SetIAttachmentSize(v int32)`

SetIAttachmentSize sets IAttachmentSize field to given value.


### GetIAttachmentEDMmoduleflag

`func (o *AttachmentResponseCompound) GetIAttachmentEDMmoduleflag() int32`

GetIAttachmentEDMmoduleflag returns the IAttachmentEDMmoduleflag field if non-nil, zero value otherwise.

### GetIAttachmentEDMmoduleflagOk

`func (o *AttachmentResponseCompound) GetIAttachmentEDMmoduleflagOk() (*int32, bool)`

GetIAttachmentEDMmoduleflagOk returns a tuple with the IAttachmentEDMmoduleflag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAttachmentEDMmoduleflag

`func (o *AttachmentResponseCompound) SetIAttachmentEDMmoduleflag(v int32)`

SetIAttachmentEDMmoduleflag sets IAttachmentEDMmoduleflag field to given value.

### HasIAttachmentEDMmoduleflag

`func (o *AttachmentResponseCompound) HasIAttachmentEDMmoduleflag() bool`

HasIAttachmentEDMmoduleflag returns a boolean if a field has been set.

### GetSAttachmentMD5

`func (o *AttachmentResponseCompound) GetSAttachmentMD5() string`

GetSAttachmentMD5 returns the SAttachmentMD5 field if non-nil, zero value otherwise.

### GetSAttachmentMD5Ok

`func (o *AttachmentResponseCompound) GetSAttachmentMD5Ok() (*string, bool)`

GetSAttachmentMD5Ok returns a tuple with the SAttachmentMD5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentMD5

`func (o *AttachmentResponseCompound) SetSAttachmentMD5(v string)`

SetSAttachmentMD5 sets SAttachmentMD5 field to given value.


### GetBAttachmentDeleted

`func (o *AttachmentResponseCompound) GetBAttachmentDeleted() bool`

GetBAttachmentDeleted returns the BAttachmentDeleted field if non-nil, zero value otherwise.

### GetBAttachmentDeletedOk

`func (o *AttachmentResponseCompound) GetBAttachmentDeletedOk() (*bool, bool)`

GetBAttachmentDeletedOk returns a tuple with the BAttachmentDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAttachmentDeleted

`func (o *AttachmentResponseCompound) SetBAttachmentDeleted(v bool)`

SetBAttachmentDeleted sets BAttachmentDeleted field to given value.


### GetBAttachmentValid

`func (o *AttachmentResponseCompound) GetBAttachmentValid() bool`

GetBAttachmentValid returns the BAttachmentValid field if non-nil, zero value otherwise.

### GetBAttachmentValidOk

`func (o *AttachmentResponseCompound) GetBAttachmentValidOk() (*bool, bool)`

GetBAttachmentValidOk returns a tuple with the BAttachmentValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAttachmentValid

`func (o *AttachmentResponseCompound) SetBAttachmentValid(v bool)`

SetBAttachmentValid sets BAttachmentValid field to given value.


### GetEAttachmentVerified

`func (o *AttachmentResponseCompound) GetEAttachmentVerified() FieldEAttachmentVerified`

GetEAttachmentVerified returns the EAttachmentVerified field if non-nil, zero value otherwise.

### GetEAttachmentVerifiedOk

`func (o *AttachmentResponseCompound) GetEAttachmentVerifiedOk() (*FieldEAttachmentVerified, bool)`

GetEAttachmentVerifiedOk returns a tuple with the EAttachmentVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentVerified

`func (o *AttachmentResponseCompound) SetEAttachmentVerified(v FieldEAttachmentVerified)`

SetEAttachmentVerified sets EAttachmentVerified field to given value.


### GetTAttachmentRejectioncomment

`func (o *AttachmentResponseCompound) GetTAttachmentRejectioncomment() string`

GetTAttachmentRejectioncomment returns the TAttachmentRejectioncomment field if non-nil, zero value otherwise.

### GetTAttachmentRejectioncommentOk

`func (o *AttachmentResponseCompound) GetTAttachmentRejectioncommentOk() (*string, bool)`

GetTAttachmentRejectioncommentOk returns a tuple with the TAttachmentRejectioncomment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTAttachmentRejectioncomment

`func (o *AttachmentResponseCompound) SetTAttachmentRejectioncomment(v string)`

SetTAttachmentRejectioncomment sets TAttachmentRejectioncomment field to given value.

### HasTAttachmentRejectioncomment

`func (o *AttachmentResponseCompound) HasTAttachmentRejectioncomment() bool`

HasTAttachmentRejectioncomment returns a boolean if a field has been set.

### GetFkiUserIDOwner

`func (o *AttachmentResponseCompound) GetFkiUserIDOwner() int32`

GetFkiUserIDOwner returns the FkiUserIDOwner field if non-nil, zero value otherwise.

### GetFkiUserIDOwnerOk

`func (o *AttachmentResponseCompound) GetFkiUserIDOwnerOk() (*int32, bool)`

GetFkiUserIDOwnerOk returns a tuple with the FkiUserIDOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDOwner

`func (o *AttachmentResponseCompound) SetFkiUserIDOwner(v int32)`

SetFkiUserIDOwner sets FkiUserIDOwner field to given value.

### HasFkiUserIDOwner

`func (o *AttachmentResponseCompound) HasFkiUserIDOwner() bool`

HasFkiUserIDOwner returns a boolean if a field has been set.

### GetObjAudit

`func (o *AttachmentResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *AttachmentResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *AttachmentResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.

### HasObjAudit

`func (o *AttachmentResponseCompound) HasObjAudit() bool`

HasObjAudit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


