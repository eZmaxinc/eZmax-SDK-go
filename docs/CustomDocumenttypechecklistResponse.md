# CustomDocumenttypechecklistResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDocumenttypechecklistID** | **int32** | The unique ID of the Documenttypechecklist | 
**FkiDocumenttypeID** | Pointer to **int32** | The unique ID of the Documenttype | [optional] 
**FkiAdjustmentID** | Pointer to **int32** | The unique ID of the Adjustment | [optional] 
**FkiAgentID** | Pointer to **int32** | The unique ID of the Agent. | [optional] 
**FkiBankaccountID** | Pointer to **int32** | The unique ID of the Bankaccount | [optional] 
**FkiBrokerID** | Pointer to **int32** | The unique ID of the Broker. | [optional] 
**FkiBuyercontractID** | Pointer to **int32** | The unique ID of the Buyercontract | [optional] 
**FkiCommissionadvanceID** | Pointer to **int32** | The unique ID of the Commissionadvance | [optional] 
**FkiCommunicationID** | Pointer to **int32** | The unique ID of the Communication. | [optional] 
**FkiCustomerID** | Pointer to **int32** | The unique ID of the Customer. | [optional] 
**FkiCustomertemplateID** | Pointer to **int32** | The unique ID of the Customertemplate | [optional] 
**FkiDepositID** | Pointer to **int32** | The unique ID of the Deposit | [optional] 
**FkiDeposittransitchequeID** | Pointer to **int32** | The unique ID of the Deposittransitcheque | [optional] 
**FkiDisclosureID** | Pointer to **int32** | The unique ID of the Disclosure | [optional] 
**FkiElectronicfundstransferID** | Pointer to **int32** | The unique ID of the Electronicfundstransfer | [optional] 
**FkiEmployeeID** | Pointer to **int32** | The unique ID of the Employee. | [optional] 
**FkiExternalbrokerID** | Pointer to **int32** | The unique ID of the Externalbroker. | [optional] 
**FkiEzsignfolderID** | Pointer to **int32** | The unique ID of the Ezsignfolder | [optional] 
**FkiFolderID** | Pointer to **int32** | The unique ID of the Folder | [optional] 
**FkiInscriptionID** | Pointer to **int32** | The unique ID of the Inscription. | [optional] 
**FkiInscriptionnotauthenticatedID** | Pointer to **int32** | The unique ID of the Inscriptionnotauthenticated. | [optional] 
**FkiInscriptiontempID** | Pointer to **int32** | The unique ID of the Inscriptiontemp | [optional] 
**FkiInvoiceID** | Pointer to **int32** | The unique ID of the Invoice. | [optional] 
**FkiLeadID** | Pointer to **int32** | The unique ID of the Lead | [optional] 
**FkiNotaryID** | Pointer to **int32** | The unique ID of the Notary. | [optional] 
**FkiOfficetaxreportID** | Pointer to **int32** | The unique ID of the Officetaxreport | [optional] 
**FkiOtherincomeID** | Pointer to **int32** | The unique ID of the Otherincome | [optional] 
**FkiPaymentpreparationID** | Pointer to **int32** | The unique ID of the Paymentpreparation | [optional] 
**FkiPurchaseID** | Pointer to **int32** | The unique ID of the Purchase | [optional] 
**FkiReconciliationID** | Pointer to **int32** | The unique ID of the Reconciliation | [optional] 
**FkiRejectedoffertopurchaseID** | Pointer to **int32** | The unique ID of the Rejectedoffertopurchase | [optional] 
**FkiSalaryID** | Pointer to **int32** | The unique ID of the Salary | [optional] 
**FkiSupplierID** | Pointer to **int32** | The unique ID of the Supplier. | [optional] 
**FkiTranqcontractID** | Pointer to **int32** | The unique ID of the Tranqcontract | [optional] 
**EDocumenttypechecklistStatus** | [**FieldEDocumenttypechecklistStatus**](FieldEDocumenttypechecklistStatus.md) |  | 
**TDocumenttypechecklistComment** | Pointer to **string** | The comment of the Documenttypechecklist | [optional] 

## Methods

### NewCustomDocumenttypechecklistResponse

`func NewCustomDocumenttypechecklistResponse(pkiDocumenttypechecklistID int32, eDocumenttypechecklistStatus FieldEDocumenttypechecklistStatus, ) *CustomDocumenttypechecklistResponse`

NewCustomDocumenttypechecklistResponse instantiates a new CustomDocumenttypechecklistResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDocumenttypechecklistResponseWithDefaults

`func NewCustomDocumenttypechecklistResponseWithDefaults() *CustomDocumenttypechecklistResponse`

NewCustomDocumenttypechecklistResponseWithDefaults instantiates a new CustomDocumenttypechecklistResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDocumenttypechecklistID

`func (o *CustomDocumenttypechecklistResponse) GetPkiDocumenttypechecklistID() int32`

GetPkiDocumenttypechecklistID returns the PkiDocumenttypechecklistID field if non-nil, zero value otherwise.

### GetPkiDocumenttypechecklistIDOk

`func (o *CustomDocumenttypechecklistResponse) GetPkiDocumenttypechecklistIDOk() (*int32, bool)`

GetPkiDocumenttypechecklistIDOk returns a tuple with the PkiDocumenttypechecklistID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDocumenttypechecklistID

`func (o *CustomDocumenttypechecklistResponse) SetPkiDocumenttypechecklistID(v int32)`

SetPkiDocumenttypechecklistID sets PkiDocumenttypechecklistID field to given value.


### GetFkiDocumenttypeID

`func (o *CustomDocumenttypechecklistResponse) GetFkiDocumenttypeID() int32`

GetFkiDocumenttypeID returns the FkiDocumenttypeID field if non-nil, zero value otherwise.

### GetFkiDocumenttypeIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiDocumenttypeIDOk() (*int32, bool)`

GetFkiDocumenttypeIDOk returns a tuple with the FkiDocumenttypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDocumenttypeID

`func (o *CustomDocumenttypechecklistResponse) SetFkiDocumenttypeID(v int32)`

SetFkiDocumenttypeID sets FkiDocumenttypeID field to given value.

### HasFkiDocumenttypeID

`func (o *CustomDocumenttypechecklistResponse) HasFkiDocumenttypeID() bool`

HasFkiDocumenttypeID returns a boolean if a field has been set.

### GetFkiAdjustmentID

`func (o *CustomDocumenttypechecklistResponse) GetFkiAdjustmentID() int32`

GetFkiAdjustmentID returns the FkiAdjustmentID field if non-nil, zero value otherwise.

### GetFkiAdjustmentIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiAdjustmentIDOk() (*int32, bool)`

GetFkiAdjustmentIDOk returns a tuple with the FkiAdjustmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAdjustmentID

`func (o *CustomDocumenttypechecklistResponse) SetFkiAdjustmentID(v int32)`

SetFkiAdjustmentID sets FkiAdjustmentID field to given value.

### HasFkiAdjustmentID

`func (o *CustomDocumenttypechecklistResponse) HasFkiAdjustmentID() bool`

HasFkiAdjustmentID returns a boolean if a field has been set.

### GetFkiAgentID

`func (o *CustomDocumenttypechecklistResponse) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *CustomDocumenttypechecklistResponse) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *CustomDocumenttypechecklistResponse) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBankaccountID

`func (o *CustomDocumenttypechecklistResponse) GetFkiBankaccountID() int32`

GetFkiBankaccountID returns the FkiBankaccountID field if non-nil, zero value otherwise.

### GetFkiBankaccountIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiBankaccountIDOk() (*int32, bool)`

GetFkiBankaccountIDOk returns a tuple with the FkiBankaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBankaccountID

`func (o *CustomDocumenttypechecklistResponse) SetFkiBankaccountID(v int32)`

SetFkiBankaccountID sets FkiBankaccountID field to given value.

### HasFkiBankaccountID

`func (o *CustomDocumenttypechecklistResponse) HasFkiBankaccountID() bool`

HasFkiBankaccountID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *CustomDocumenttypechecklistResponse) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *CustomDocumenttypechecklistResponse) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *CustomDocumenttypechecklistResponse) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetFkiBuyercontractID

`func (o *CustomDocumenttypechecklistResponse) GetFkiBuyercontractID() int32`

GetFkiBuyercontractID returns the FkiBuyercontractID field if non-nil, zero value otherwise.

### GetFkiBuyercontractIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiBuyercontractIDOk() (*int32, bool)`

GetFkiBuyercontractIDOk returns a tuple with the FkiBuyercontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBuyercontractID

`func (o *CustomDocumenttypechecklistResponse) SetFkiBuyercontractID(v int32)`

SetFkiBuyercontractID sets FkiBuyercontractID field to given value.

### HasFkiBuyercontractID

`func (o *CustomDocumenttypechecklistResponse) HasFkiBuyercontractID() bool`

HasFkiBuyercontractID returns a boolean if a field has been set.

### GetFkiCommissionadvanceID

`func (o *CustomDocumenttypechecklistResponse) GetFkiCommissionadvanceID() int32`

GetFkiCommissionadvanceID returns the FkiCommissionadvanceID field if non-nil, zero value otherwise.

### GetFkiCommissionadvanceIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiCommissionadvanceIDOk() (*int32, bool)`

GetFkiCommissionadvanceIDOk returns a tuple with the FkiCommissionadvanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCommissionadvanceID

`func (o *CustomDocumenttypechecklistResponse) SetFkiCommissionadvanceID(v int32)`

SetFkiCommissionadvanceID sets FkiCommissionadvanceID field to given value.

### HasFkiCommissionadvanceID

`func (o *CustomDocumenttypechecklistResponse) HasFkiCommissionadvanceID() bool`

HasFkiCommissionadvanceID returns a boolean if a field has been set.

### GetFkiCommunicationID

`func (o *CustomDocumenttypechecklistResponse) GetFkiCommunicationID() int32`

GetFkiCommunicationID returns the FkiCommunicationID field if non-nil, zero value otherwise.

### GetFkiCommunicationIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiCommunicationIDOk() (*int32, bool)`

GetFkiCommunicationIDOk returns a tuple with the FkiCommunicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCommunicationID

`func (o *CustomDocumenttypechecklistResponse) SetFkiCommunicationID(v int32)`

SetFkiCommunicationID sets FkiCommunicationID field to given value.

### HasFkiCommunicationID

`func (o *CustomDocumenttypechecklistResponse) HasFkiCommunicationID() bool`

HasFkiCommunicationID returns a boolean if a field has been set.

### GetFkiCustomerID

`func (o *CustomDocumenttypechecklistResponse) GetFkiCustomerID() int32`

GetFkiCustomerID returns the FkiCustomerID field if non-nil, zero value otherwise.

### GetFkiCustomerIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiCustomerIDOk() (*int32, bool)`

GetFkiCustomerIDOk returns a tuple with the FkiCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomerID

`func (o *CustomDocumenttypechecklistResponse) SetFkiCustomerID(v int32)`

SetFkiCustomerID sets FkiCustomerID field to given value.

### HasFkiCustomerID

`func (o *CustomDocumenttypechecklistResponse) HasFkiCustomerID() bool`

HasFkiCustomerID returns a boolean if a field has been set.

### GetFkiCustomertemplateID

`func (o *CustomDocumenttypechecklistResponse) GetFkiCustomertemplateID() int32`

GetFkiCustomertemplateID returns the FkiCustomertemplateID field if non-nil, zero value otherwise.

### GetFkiCustomertemplateIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiCustomertemplateIDOk() (*int32, bool)`

GetFkiCustomertemplateIDOk returns a tuple with the FkiCustomertemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomertemplateID

`func (o *CustomDocumenttypechecklistResponse) SetFkiCustomertemplateID(v int32)`

SetFkiCustomertemplateID sets FkiCustomertemplateID field to given value.

### HasFkiCustomertemplateID

`func (o *CustomDocumenttypechecklistResponse) HasFkiCustomertemplateID() bool`

HasFkiCustomertemplateID returns a boolean if a field has been set.

### GetFkiDepositID

`func (o *CustomDocumenttypechecklistResponse) GetFkiDepositID() int32`

GetFkiDepositID returns the FkiDepositID field if non-nil, zero value otherwise.

### GetFkiDepositIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiDepositIDOk() (*int32, bool)`

GetFkiDepositIDOk returns a tuple with the FkiDepositID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepositID

`func (o *CustomDocumenttypechecklistResponse) SetFkiDepositID(v int32)`

SetFkiDepositID sets FkiDepositID field to given value.

### HasFkiDepositID

`func (o *CustomDocumenttypechecklistResponse) HasFkiDepositID() bool`

HasFkiDepositID returns a boolean if a field has been set.

### GetFkiDeposittransitchequeID

`func (o *CustomDocumenttypechecklistResponse) GetFkiDeposittransitchequeID() int32`

GetFkiDeposittransitchequeID returns the FkiDeposittransitchequeID field if non-nil, zero value otherwise.

### GetFkiDeposittransitchequeIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiDeposittransitchequeIDOk() (*int32, bool)`

GetFkiDeposittransitchequeIDOk returns a tuple with the FkiDeposittransitchequeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDeposittransitchequeID

`func (o *CustomDocumenttypechecklistResponse) SetFkiDeposittransitchequeID(v int32)`

SetFkiDeposittransitchequeID sets FkiDeposittransitchequeID field to given value.

### HasFkiDeposittransitchequeID

`func (o *CustomDocumenttypechecklistResponse) HasFkiDeposittransitchequeID() bool`

HasFkiDeposittransitchequeID returns a boolean if a field has been set.

### GetFkiDisclosureID

`func (o *CustomDocumenttypechecklistResponse) GetFkiDisclosureID() int32`

GetFkiDisclosureID returns the FkiDisclosureID field if non-nil, zero value otherwise.

### GetFkiDisclosureIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiDisclosureIDOk() (*int32, bool)`

GetFkiDisclosureIDOk returns a tuple with the FkiDisclosureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDisclosureID

`func (o *CustomDocumenttypechecklistResponse) SetFkiDisclosureID(v int32)`

SetFkiDisclosureID sets FkiDisclosureID field to given value.

### HasFkiDisclosureID

`func (o *CustomDocumenttypechecklistResponse) HasFkiDisclosureID() bool`

HasFkiDisclosureID returns a boolean if a field has been set.

### GetFkiElectronicfundstransferID

`func (o *CustomDocumenttypechecklistResponse) GetFkiElectronicfundstransferID() int32`

GetFkiElectronicfundstransferID returns the FkiElectronicfundstransferID field if non-nil, zero value otherwise.

### GetFkiElectronicfundstransferIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiElectronicfundstransferIDOk() (*int32, bool)`

GetFkiElectronicfundstransferIDOk returns a tuple with the FkiElectronicfundstransferID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiElectronicfundstransferID

`func (o *CustomDocumenttypechecklistResponse) SetFkiElectronicfundstransferID(v int32)`

SetFkiElectronicfundstransferID sets FkiElectronicfundstransferID field to given value.

### HasFkiElectronicfundstransferID

`func (o *CustomDocumenttypechecklistResponse) HasFkiElectronicfundstransferID() bool`

HasFkiElectronicfundstransferID returns a boolean if a field has been set.

### GetFkiEmployeeID

`func (o *CustomDocumenttypechecklistResponse) GetFkiEmployeeID() int32`

GetFkiEmployeeID returns the FkiEmployeeID field if non-nil, zero value otherwise.

### GetFkiEmployeeIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiEmployeeIDOk() (*int32, bool)`

GetFkiEmployeeIDOk returns a tuple with the FkiEmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmployeeID

`func (o *CustomDocumenttypechecklistResponse) SetFkiEmployeeID(v int32)`

SetFkiEmployeeID sets FkiEmployeeID field to given value.

### HasFkiEmployeeID

`func (o *CustomDocumenttypechecklistResponse) HasFkiEmployeeID() bool`

HasFkiEmployeeID returns a boolean if a field has been set.

### GetFkiExternalbrokerID

`func (o *CustomDocumenttypechecklistResponse) GetFkiExternalbrokerID() int32`

GetFkiExternalbrokerID returns the FkiExternalbrokerID field if non-nil, zero value otherwise.

### GetFkiExternalbrokerIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiExternalbrokerIDOk() (*int32, bool)`

GetFkiExternalbrokerIDOk returns a tuple with the FkiExternalbrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiExternalbrokerID

`func (o *CustomDocumenttypechecklistResponse) SetFkiExternalbrokerID(v int32)`

SetFkiExternalbrokerID sets FkiExternalbrokerID field to given value.

### HasFkiExternalbrokerID

`func (o *CustomDocumenttypechecklistResponse) HasFkiExternalbrokerID() bool`

HasFkiExternalbrokerID returns a boolean if a field has been set.

### GetFkiEzsignfolderID

`func (o *CustomDocumenttypechecklistResponse) GetFkiEzsignfolderID() int32`

GetFkiEzsignfolderID returns the FkiEzsignfolderID field if non-nil, zero value otherwise.

### GetFkiEzsignfolderIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiEzsignfolderIDOk() (*int32, bool)`

GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfolderID

`func (o *CustomDocumenttypechecklistResponse) SetFkiEzsignfolderID(v int32)`

SetFkiEzsignfolderID sets FkiEzsignfolderID field to given value.

### HasFkiEzsignfolderID

`func (o *CustomDocumenttypechecklistResponse) HasFkiEzsignfolderID() bool`

HasFkiEzsignfolderID returns a boolean if a field has been set.

### GetFkiFolderID

`func (o *CustomDocumenttypechecklistResponse) GetFkiFolderID() int32`

GetFkiFolderID returns the FkiFolderID field if non-nil, zero value otherwise.

### GetFkiFolderIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiFolderIDOk() (*int32, bool)`

GetFkiFolderIDOk returns a tuple with the FkiFolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFolderID

`func (o *CustomDocumenttypechecklistResponse) SetFkiFolderID(v int32)`

SetFkiFolderID sets FkiFolderID field to given value.

### HasFkiFolderID

`func (o *CustomDocumenttypechecklistResponse) HasFkiFolderID() bool`

HasFkiFolderID returns a boolean if a field has been set.

### GetFkiInscriptionID

`func (o *CustomDocumenttypechecklistResponse) GetFkiInscriptionID() int32`

GetFkiInscriptionID returns the FkiInscriptionID field if non-nil, zero value otherwise.

### GetFkiInscriptionIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiInscriptionIDOk() (*int32, bool)`

GetFkiInscriptionIDOk returns a tuple with the FkiInscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionID

`func (o *CustomDocumenttypechecklistResponse) SetFkiInscriptionID(v int32)`

SetFkiInscriptionID sets FkiInscriptionID field to given value.

### HasFkiInscriptionID

`func (o *CustomDocumenttypechecklistResponse) HasFkiInscriptionID() bool`

HasFkiInscriptionID returns a boolean if a field has been set.

### GetFkiInscriptionnotauthenticatedID

`func (o *CustomDocumenttypechecklistResponse) GetFkiInscriptionnotauthenticatedID() int32`

GetFkiInscriptionnotauthenticatedID returns the FkiInscriptionnotauthenticatedID field if non-nil, zero value otherwise.

### GetFkiInscriptionnotauthenticatedIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiInscriptionnotauthenticatedIDOk() (*int32, bool)`

GetFkiInscriptionnotauthenticatedIDOk returns a tuple with the FkiInscriptionnotauthenticatedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionnotauthenticatedID

`func (o *CustomDocumenttypechecklistResponse) SetFkiInscriptionnotauthenticatedID(v int32)`

SetFkiInscriptionnotauthenticatedID sets FkiInscriptionnotauthenticatedID field to given value.

### HasFkiInscriptionnotauthenticatedID

`func (o *CustomDocumenttypechecklistResponse) HasFkiInscriptionnotauthenticatedID() bool`

HasFkiInscriptionnotauthenticatedID returns a boolean if a field has been set.

### GetFkiInscriptiontempID

`func (o *CustomDocumenttypechecklistResponse) GetFkiInscriptiontempID() int32`

GetFkiInscriptiontempID returns the FkiInscriptiontempID field if non-nil, zero value otherwise.

### GetFkiInscriptiontempIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiInscriptiontempIDOk() (*int32, bool)`

GetFkiInscriptiontempIDOk returns a tuple with the FkiInscriptiontempID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptiontempID

`func (o *CustomDocumenttypechecklistResponse) SetFkiInscriptiontempID(v int32)`

SetFkiInscriptiontempID sets FkiInscriptiontempID field to given value.

### HasFkiInscriptiontempID

`func (o *CustomDocumenttypechecklistResponse) HasFkiInscriptiontempID() bool`

HasFkiInscriptiontempID returns a boolean if a field has been set.

### GetFkiInvoiceID

`func (o *CustomDocumenttypechecklistResponse) GetFkiInvoiceID() int32`

GetFkiInvoiceID returns the FkiInvoiceID field if non-nil, zero value otherwise.

### GetFkiInvoiceIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiInvoiceIDOk() (*int32, bool)`

GetFkiInvoiceIDOk returns a tuple with the FkiInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInvoiceID

`func (o *CustomDocumenttypechecklistResponse) SetFkiInvoiceID(v int32)`

SetFkiInvoiceID sets FkiInvoiceID field to given value.

### HasFkiInvoiceID

`func (o *CustomDocumenttypechecklistResponse) HasFkiInvoiceID() bool`

HasFkiInvoiceID returns a boolean if a field has been set.

### GetFkiLeadID

`func (o *CustomDocumenttypechecklistResponse) GetFkiLeadID() int32`

GetFkiLeadID returns the FkiLeadID field if non-nil, zero value otherwise.

### GetFkiLeadIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiLeadIDOk() (*int32, bool)`

GetFkiLeadIDOk returns a tuple with the FkiLeadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLeadID

`func (o *CustomDocumenttypechecklistResponse) SetFkiLeadID(v int32)`

SetFkiLeadID sets FkiLeadID field to given value.

### HasFkiLeadID

`func (o *CustomDocumenttypechecklistResponse) HasFkiLeadID() bool`

HasFkiLeadID returns a boolean if a field has been set.

### GetFkiNotaryID

`func (o *CustomDocumenttypechecklistResponse) GetFkiNotaryID() int32`

GetFkiNotaryID returns the FkiNotaryID field if non-nil, zero value otherwise.

### GetFkiNotaryIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiNotaryIDOk() (*int32, bool)`

GetFkiNotaryIDOk returns a tuple with the FkiNotaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiNotaryID

`func (o *CustomDocumenttypechecklistResponse) SetFkiNotaryID(v int32)`

SetFkiNotaryID sets FkiNotaryID field to given value.

### HasFkiNotaryID

`func (o *CustomDocumenttypechecklistResponse) HasFkiNotaryID() bool`

HasFkiNotaryID returns a boolean if a field has been set.

### GetFkiOfficetaxreportID

`func (o *CustomDocumenttypechecklistResponse) GetFkiOfficetaxreportID() int32`

GetFkiOfficetaxreportID returns the FkiOfficetaxreportID field if non-nil, zero value otherwise.

### GetFkiOfficetaxreportIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiOfficetaxreportIDOk() (*int32, bool)`

GetFkiOfficetaxreportIDOk returns a tuple with the FkiOfficetaxreportID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiOfficetaxreportID

`func (o *CustomDocumenttypechecklistResponse) SetFkiOfficetaxreportID(v int32)`

SetFkiOfficetaxreportID sets FkiOfficetaxreportID field to given value.

### HasFkiOfficetaxreportID

`func (o *CustomDocumenttypechecklistResponse) HasFkiOfficetaxreportID() bool`

HasFkiOfficetaxreportID returns a boolean if a field has been set.

### GetFkiOtherincomeID

`func (o *CustomDocumenttypechecklistResponse) GetFkiOtherincomeID() int32`

GetFkiOtherincomeID returns the FkiOtherincomeID field if non-nil, zero value otherwise.

### GetFkiOtherincomeIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiOtherincomeIDOk() (*int32, bool)`

GetFkiOtherincomeIDOk returns a tuple with the FkiOtherincomeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiOtherincomeID

`func (o *CustomDocumenttypechecklistResponse) SetFkiOtherincomeID(v int32)`

SetFkiOtherincomeID sets FkiOtherincomeID field to given value.

### HasFkiOtherincomeID

`func (o *CustomDocumenttypechecklistResponse) HasFkiOtherincomeID() bool`

HasFkiOtherincomeID returns a boolean if a field has been set.

### GetFkiPaymentpreparationID

`func (o *CustomDocumenttypechecklistResponse) GetFkiPaymentpreparationID() int32`

GetFkiPaymentpreparationID returns the FkiPaymentpreparationID field if non-nil, zero value otherwise.

### GetFkiPaymentpreparationIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiPaymentpreparationIDOk() (*int32, bool)`

GetFkiPaymentpreparationIDOk returns a tuple with the FkiPaymentpreparationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPaymentpreparationID

`func (o *CustomDocumenttypechecklistResponse) SetFkiPaymentpreparationID(v int32)`

SetFkiPaymentpreparationID sets FkiPaymentpreparationID field to given value.

### HasFkiPaymentpreparationID

`func (o *CustomDocumenttypechecklistResponse) HasFkiPaymentpreparationID() bool`

HasFkiPaymentpreparationID returns a boolean if a field has been set.

### GetFkiPurchaseID

`func (o *CustomDocumenttypechecklistResponse) GetFkiPurchaseID() int32`

GetFkiPurchaseID returns the FkiPurchaseID field if non-nil, zero value otherwise.

### GetFkiPurchaseIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiPurchaseIDOk() (*int32, bool)`

GetFkiPurchaseIDOk returns a tuple with the FkiPurchaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPurchaseID

`func (o *CustomDocumenttypechecklistResponse) SetFkiPurchaseID(v int32)`

SetFkiPurchaseID sets FkiPurchaseID field to given value.

### HasFkiPurchaseID

`func (o *CustomDocumenttypechecklistResponse) HasFkiPurchaseID() bool`

HasFkiPurchaseID returns a boolean if a field has been set.

### GetFkiReconciliationID

`func (o *CustomDocumenttypechecklistResponse) GetFkiReconciliationID() int32`

GetFkiReconciliationID returns the FkiReconciliationID field if non-nil, zero value otherwise.

### GetFkiReconciliationIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiReconciliationIDOk() (*int32, bool)`

GetFkiReconciliationIDOk returns a tuple with the FkiReconciliationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiReconciliationID

`func (o *CustomDocumenttypechecklistResponse) SetFkiReconciliationID(v int32)`

SetFkiReconciliationID sets FkiReconciliationID field to given value.

### HasFkiReconciliationID

`func (o *CustomDocumenttypechecklistResponse) HasFkiReconciliationID() bool`

HasFkiReconciliationID returns a boolean if a field has been set.

### GetFkiRejectedoffertopurchaseID

`func (o *CustomDocumenttypechecklistResponse) GetFkiRejectedoffertopurchaseID() int32`

GetFkiRejectedoffertopurchaseID returns the FkiRejectedoffertopurchaseID field if non-nil, zero value otherwise.

### GetFkiRejectedoffertopurchaseIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiRejectedoffertopurchaseIDOk() (*int32, bool)`

GetFkiRejectedoffertopurchaseIDOk returns a tuple with the FkiRejectedoffertopurchaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiRejectedoffertopurchaseID

`func (o *CustomDocumenttypechecklistResponse) SetFkiRejectedoffertopurchaseID(v int32)`

SetFkiRejectedoffertopurchaseID sets FkiRejectedoffertopurchaseID field to given value.

### HasFkiRejectedoffertopurchaseID

`func (o *CustomDocumenttypechecklistResponse) HasFkiRejectedoffertopurchaseID() bool`

HasFkiRejectedoffertopurchaseID returns a boolean if a field has been set.

### GetFkiSalaryID

`func (o *CustomDocumenttypechecklistResponse) GetFkiSalaryID() int32`

GetFkiSalaryID returns the FkiSalaryID field if non-nil, zero value otherwise.

### GetFkiSalaryIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiSalaryIDOk() (*int32, bool)`

GetFkiSalaryIDOk returns a tuple with the FkiSalaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSalaryID

`func (o *CustomDocumenttypechecklistResponse) SetFkiSalaryID(v int32)`

SetFkiSalaryID sets FkiSalaryID field to given value.

### HasFkiSalaryID

`func (o *CustomDocumenttypechecklistResponse) HasFkiSalaryID() bool`

HasFkiSalaryID returns a boolean if a field has been set.

### GetFkiSupplierID

`func (o *CustomDocumenttypechecklistResponse) GetFkiSupplierID() int32`

GetFkiSupplierID returns the FkiSupplierID field if non-nil, zero value otherwise.

### GetFkiSupplierIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiSupplierIDOk() (*int32, bool)`

GetFkiSupplierIDOk returns a tuple with the FkiSupplierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSupplierID

`func (o *CustomDocumenttypechecklistResponse) SetFkiSupplierID(v int32)`

SetFkiSupplierID sets FkiSupplierID field to given value.

### HasFkiSupplierID

`func (o *CustomDocumenttypechecklistResponse) HasFkiSupplierID() bool`

HasFkiSupplierID returns a boolean if a field has been set.

### GetFkiTranqcontractID

`func (o *CustomDocumenttypechecklistResponse) GetFkiTranqcontractID() int32`

GetFkiTranqcontractID returns the FkiTranqcontractID field if non-nil, zero value otherwise.

### GetFkiTranqcontractIDOk

`func (o *CustomDocumenttypechecklistResponse) GetFkiTranqcontractIDOk() (*int32, bool)`

GetFkiTranqcontractIDOk returns a tuple with the FkiTranqcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTranqcontractID

`func (o *CustomDocumenttypechecklistResponse) SetFkiTranqcontractID(v int32)`

SetFkiTranqcontractID sets FkiTranqcontractID field to given value.

### HasFkiTranqcontractID

`func (o *CustomDocumenttypechecklistResponse) HasFkiTranqcontractID() bool`

HasFkiTranqcontractID returns a boolean if a field has been set.

### GetEDocumenttypechecklistStatus

`func (o *CustomDocumenttypechecklistResponse) GetEDocumenttypechecklistStatus() FieldEDocumenttypechecklistStatus`

GetEDocumenttypechecklistStatus returns the EDocumenttypechecklistStatus field if non-nil, zero value otherwise.

### GetEDocumenttypechecklistStatusOk

`func (o *CustomDocumenttypechecklistResponse) GetEDocumenttypechecklistStatusOk() (*FieldEDocumenttypechecklistStatus, bool)`

GetEDocumenttypechecklistStatusOk returns a tuple with the EDocumenttypechecklistStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypechecklistStatus

`func (o *CustomDocumenttypechecklistResponse) SetEDocumenttypechecklistStatus(v FieldEDocumenttypechecklistStatus)`

SetEDocumenttypechecklistStatus sets EDocumenttypechecklistStatus field to given value.


### GetTDocumenttypechecklistComment

`func (o *CustomDocumenttypechecklistResponse) GetTDocumenttypechecklistComment() string`

GetTDocumenttypechecklistComment returns the TDocumenttypechecklistComment field if non-nil, zero value otherwise.

### GetTDocumenttypechecklistCommentOk

`func (o *CustomDocumenttypechecklistResponse) GetTDocumenttypechecklistCommentOk() (*string, bool)`

GetTDocumenttypechecklistCommentOk returns a tuple with the TDocumenttypechecklistComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTDocumenttypechecklistComment

`func (o *CustomDocumenttypechecklistResponse) SetTDocumenttypechecklistComment(v string)`

SetTDocumenttypechecklistComment sets TDocumenttypechecklistComment field to given value.

### HasTDocumenttypechecklistComment

`func (o *CustomDocumenttypechecklistResponse) HasTDocumenttypechecklistComment() bool`

HasTDocumenttypechecklistComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


