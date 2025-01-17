# EzmaxinvoicingagentResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzmaxinvoicingagentID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicingagent | [optional] 
**FkiEzmaxinvoicingID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicing | [optional] 
**FkiBillingentityinternalID** | **int32** | The unique ID of the Billingentityinternal. | 
**SBillingentityinternalDescriptionX** | **string** | The description of the Billingentityinternal in the language of the requester | 
**FkiAgentID** | Pointer to **int32** | The unique ID of the Agent. | [optional] 
**FkiBrokerID** | Pointer to **int32** | The unique ID of the Broker. | [optional] 
**IEzmaxinvoicingagentSession** | **int32** | The number of sessions | 
**IEzmaxinvoicingagentCloned** | **int32** | The number of times this user was cloned | 
**IEzmaxinvoicingagentInvoice** | **int32** | The number of invoices | 
**IEzmaxinvoicingagentInscription** | **int32** | The number of inscriptions | 
**IEzmaxinvoicingagentInscriptionactive** | **int32** | The number of active inscriptions | 
**IEzmaxinvoicingagentSale** | **int32** | The number of sales | 
**IEzmaxinvoicingagentOtherincome** | **int32** | The number of otherincomes | 
**IEzmaxinvoicingagentCommissioncalculation** | **int32** | The number of commission calculations | 
**IEzmaxinvoicingagentEzsigndocument** | **int32** | The number of ezsign documents | 
**BEzmaxinvoicingagentEzsignaccount** | **bool** | Whether the agent has an eZsign account | 
**BEzmaxinvoicingagentBillableezmax** | **bool** | Whether it is billable for eZmax | 
**EEzmaxinvoicingagentVariationezmax** | [**FieldEEzmaxinvoicingagentVariationezmax**](FieldEEzmaxinvoicingagentVariationezmax.md) |  | 
**BEzmaxinvoicingagentBillableezsign** | **bool** | Whether it is billable for eZsign | 
**EEzmaxinvoicingagentVariationezsign** | [**FieldEEzmaxinvoicingagentVariationezsign**](FieldEEzmaxinvoicingagentVariationezsign.md) |  | 
**ObjContactName** | [**CustomContactNameResponse**](CustomContactNameResponse.md) |  | 

## Methods

### NewEzmaxinvoicingagentResponseCompound

`func NewEzmaxinvoicingagentResponseCompound(fkiBillingentityinternalID int32, sBillingentityinternalDescriptionX string, iEzmaxinvoicingagentSession int32, iEzmaxinvoicingagentCloned int32, iEzmaxinvoicingagentInvoice int32, iEzmaxinvoicingagentInscription int32, iEzmaxinvoicingagentInscriptionactive int32, iEzmaxinvoicingagentSale int32, iEzmaxinvoicingagentOtherincome int32, iEzmaxinvoicingagentCommissioncalculation int32, iEzmaxinvoicingagentEzsigndocument int32, bEzmaxinvoicingagentEzsignaccount bool, bEzmaxinvoicingagentBillableezmax bool, eEzmaxinvoicingagentVariationezmax FieldEEzmaxinvoicingagentVariationezmax, bEzmaxinvoicingagentBillableezsign bool, eEzmaxinvoicingagentVariationezsign FieldEEzmaxinvoicingagentVariationezsign, objContactName CustomContactNameResponse, ) *EzmaxinvoicingagentResponseCompound`

NewEzmaxinvoicingagentResponseCompound instantiates a new EzmaxinvoicingagentResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingagentResponseCompoundWithDefaults

`func NewEzmaxinvoicingagentResponseCompoundWithDefaults() *EzmaxinvoicingagentResponseCompound`

NewEzmaxinvoicingagentResponseCompoundWithDefaults instantiates a new EzmaxinvoicingagentResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicingagentID

`func (o *EzmaxinvoicingagentResponseCompound) GetPkiEzmaxinvoicingagentID() int32`

GetPkiEzmaxinvoicingagentID returns the PkiEzmaxinvoicingagentID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicingagentIDOk

`func (o *EzmaxinvoicingagentResponseCompound) GetPkiEzmaxinvoicingagentIDOk() (*int32, bool)`

GetPkiEzmaxinvoicingagentIDOk returns a tuple with the PkiEzmaxinvoicingagentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicingagentID

`func (o *EzmaxinvoicingagentResponseCompound) SetPkiEzmaxinvoicingagentID(v int32)`

SetPkiEzmaxinvoicingagentID sets PkiEzmaxinvoicingagentID field to given value.

### HasPkiEzmaxinvoicingagentID

`func (o *EzmaxinvoicingagentResponseCompound) HasPkiEzmaxinvoicingagentID() bool`

HasPkiEzmaxinvoicingagentID returns a boolean if a field has been set.

### GetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingagentResponseCompound) GetFkiEzmaxinvoicingID() int32`

GetFkiEzmaxinvoicingID returns the FkiEzmaxinvoicingID field if non-nil, zero value otherwise.

### GetFkiEzmaxinvoicingIDOk

`func (o *EzmaxinvoicingagentResponseCompound) GetFkiEzmaxinvoicingIDOk() (*int32, bool)`

GetFkiEzmaxinvoicingIDOk returns a tuple with the FkiEzmaxinvoicingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingagentResponseCompound) SetFkiEzmaxinvoicingID(v int32)`

SetFkiEzmaxinvoicingID sets FkiEzmaxinvoicingID field to given value.

### HasFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingagentResponseCompound) HasFkiEzmaxinvoicingID() bool`

HasFkiEzmaxinvoicingID returns a boolean if a field has been set.

### GetFkiBillingentityinternalID

`func (o *EzmaxinvoicingagentResponseCompound) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *EzmaxinvoicingagentResponseCompound) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *EzmaxinvoicingagentResponseCompound) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.


### GetSBillingentityinternalDescriptionX

`func (o *EzmaxinvoicingagentResponseCompound) GetSBillingentityinternalDescriptionX() string`

GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field if non-nil, zero value otherwise.

### GetSBillingentityinternalDescriptionXOk

`func (o *EzmaxinvoicingagentResponseCompound) GetSBillingentityinternalDescriptionXOk() (*string, bool)`

GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBillingentityinternalDescriptionX

`func (o *EzmaxinvoicingagentResponseCompound) SetSBillingentityinternalDescriptionX(v string)`

SetSBillingentityinternalDescriptionX sets SBillingentityinternalDescriptionX field to given value.


### GetFkiAgentID

`func (o *EzmaxinvoicingagentResponseCompound) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *EzmaxinvoicingagentResponseCompound) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *EzmaxinvoicingagentResponseCompound) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *EzmaxinvoicingagentResponseCompound) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *EzmaxinvoicingagentResponseCompound) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *EzmaxinvoicingagentResponseCompound) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *EzmaxinvoicingagentResponseCompound) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *EzmaxinvoicingagentResponseCompound) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetIEzmaxinvoicingagentSession

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentSession() int32`

GetIEzmaxinvoicingagentSession returns the IEzmaxinvoicingagentSession field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentSessionOk

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentSessionOk() (*int32, bool)`

GetIEzmaxinvoicingagentSessionOk returns a tuple with the IEzmaxinvoicingagentSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentSession

`func (o *EzmaxinvoicingagentResponseCompound) SetIEzmaxinvoicingagentSession(v int32)`

SetIEzmaxinvoicingagentSession sets IEzmaxinvoicingagentSession field to given value.


### GetIEzmaxinvoicingagentCloned

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentCloned() int32`

GetIEzmaxinvoicingagentCloned returns the IEzmaxinvoicingagentCloned field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentClonedOk

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentClonedOk() (*int32, bool)`

GetIEzmaxinvoicingagentClonedOk returns a tuple with the IEzmaxinvoicingagentCloned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentCloned

`func (o *EzmaxinvoicingagentResponseCompound) SetIEzmaxinvoicingagentCloned(v int32)`

SetIEzmaxinvoicingagentCloned sets IEzmaxinvoicingagentCloned field to given value.


### GetIEzmaxinvoicingagentInvoice

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentInvoice() int32`

GetIEzmaxinvoicingagentInvoice returns the IEzmaxinvoicingagentInvoice field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentInvoiceOk

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentInvoiceOk() (*int32, bool)`

GetIEzmaxinvoicingagentInvoiceOk returns a tuple with the IEzmaxinvoicingagentInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentInvoice

`func (o *EzmaxinvoicingagentResponseCompound) SetIEzmaxinvoicingagentInvoice(v int32)`

SetIEzmaxinvoicingagentInvoice sets IEzmaxinvoicingagentInvoice field to given value.


### GetIEzmaxinvoicingagentInscription

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentInscription() int32`

GetIEzmaxinvoicingagentInscription returns the IEzmaxinvoicingagentInscription field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentInscriptionOk

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentInscriptionOk() (*int32, bool)`

GetIEzmaxinvoicingagentInscriptionOk returns a tuple with the IEzmaxinvoicingagentInscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentInscription

`func (o *EzmaxinvoicingagentResponseCompound) SetIEzmaxinvoicingagentInscription(v int32)`

SetIEzmaxinvoicingagentInscription sets IEzmaxinvoicingagentInscription field to given value.


### GetIEzmaxinvoicingagentInscriptionactive

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentInscriptionactive() int32`

GetIEzmaxinvoicingagentInscriptionactive returns the IEzmaxinvoicingagentInscriptionactive field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentInscriptionactiveOk

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentInscriptionactiveOk() (*int32, bool)`

GetIEzmaxinvoicingagentInscriptionactiveOk returns a tuple with the IEzmaxinvoicingagentInscriptionactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentInscriptionactive

`func (o *EzmaxinvoicingagentResponseCompound) SetIEzmaxinvoicingagentInscriptionactive(v int32)`

SetIEzmaxinvoicingagentInscriptionactive sets IEzmaxinvoicingagentInscriptionactive field to given value.


### GetIEzmaxinvoicingagentSale

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentSale() int32`

GetIEzmaxinvoicingagentSale returns the IEzmaxinvoicingagentSale field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentSaleOk

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentSaleOk() (*int32, bool)`

GetIEzmaxinvoicingagentSaleOk returns a tuple with the IEzmaxinvoicingagentSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentSale

`func (o *EzmaxinvoicingagentResponseCompound) SetIEzmaxinvoicingagentSale(v int32)`

SetIEzmaxinvoicingagentSale sets IEzmaxinvoicingagentSale field to given value.


### GetIEzmaxinvoicingagentOtherincome

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentOtherincome() int32`

GetIEzmaxinvoicingagentOtherincome returns the IEzmaxinvoicingagentOtherincome field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentOtherincomeOk

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentOtherincomeOk() (*int32, bool)`

GetIEzmaxinvoicingagentOtherincomeOk returns a tuple with the IEzmaxinvoicingagentOtherincome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentOtherincome

`func (o *EzmaxinvoicingagentResponseCompound) SetIEzmaxinvoicingagentOtherincome(v int32)`

SetIEzmaxinvoicingagentOtherincome sets IEzmaxinvoicingagentOtherincome field to given value.


### GetIEzmaxinvoicingagentCommissioncalculation

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentCommissioncalculation() int32`

GetIEzmaxinvoicingagentCommissioncalculation returns the IEzmaxinvoicingagentCommissioncalculation field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentCommissioncalculationOk

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentCommissioncalculationOk() (*int32, bool)`

GetIEzmaxinvoicingagentCommissioncalculationOk returns a tuple with the IEzmaxinvoicingagentCommissioncalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentCommissioncalculation

`func (o *EzmaxinvoicingagentResponseCompound) SetIEzmaxinvoicingagentCommissioncalculation(v int32)`

SetIEzmaxinvoicingagentCommissioncalculation sets IEzmaxinvoicingagentCommissioncalculation field to given value.


### GetIEzmaxinvoicingagentEzsigndocument

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentEzsigndocument() int32`

GetIEzmaxinvoicingagentEzsigndocument returns the IEzmaxinvoicingagentEzsigndocument field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentEzsigndocumentOk

`func (o *EzmaxinvoicingagentResponseCompound) GetIEzmaxinvoicingagentEzsigndocumentOk() (*int32, bool)`

GetIEzmaxinvoicingagentEzsigndocumentOk returns a tuple with the IEzmaxinvoicingagentEzsigndocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentEzsigndocument

`func (o *EzmaxinvoicingagentResponseCompound) SetIEzmaxinvoicingagentEzsigndocument(v int32)`

SetIEzmaxinvoicingagentEzsigndocument sets IEzmaxinvoicingagentEzsigndocument field to given value.


### GetBEzmaxinvoicingagentEzsignaccount

`func (o *EzmaxinvoicingagentResponseCompound) GetBEzmaxinvoicingagentEzsignaccount() bool`

GetBEzmaxinvoicingagentEzsignaccount returns the BEzmaxinvoicingagentEzsignaccount field if non-nil, zero value otherwise.

### GetBEzmaxinvoicingagentEzsignaccountOk

`func (o *EzmaxinvoicingagentResponseCompound) GetBEzmaxinvoicingagentEzsignaccountOk() (*bool, bool)`

GetBEzmaxinvoicingagentEzsignaccountOk returns a tuple with the BEzmaxinvoicingagentEzsignaccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicingagentEzsignaccount

`func (o *EzmaxinvoicingagentResponseCompound) SetBEzmaxinvoicingagentEzsignaccount(v bool)`

SetBEzmaxinvoicingagentEzsignaccount sets BEzmaxinvoicingagentEzsignaccount field to given value.


### GetBEzmaxinvoicingagentBillableezmax

`func (o *EzmaxinvoicingagentResponseCompound) GetBEzmaxinvoicingagentBillableezmax() bool`

GetBEzmaxinvoicingagentBillableezmax returns the BEzmaxinvoicingagentBillableezmax field if non-nil, zero value otherwise.

### GetBEzmaxinvoicingagentBillableezmaxOk

`func (o *EzmaxinvoicingagentResponseCompound) GetBEzmaxinvoicingagentBillableezmaxOk() (*bool, bool)`

GetBEzmaxinvoicingagentBillableezmaxOk returns a tuple with the BEzmaxinvoicingagentBillableezmax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicingagentBillableezmax

`func (o *EzmaxinvoicingagentResponseCompound) SetBEzmaxinvoicingagentBillableezmax(v bool)`

SetBEzmaxinvoicingagentBillableezmax sets BEzmaxinvoicingagentBillableezmax field to given value.


### GetEEzmaxinvoicingagentVariationezmax

`func (o *EzmaxinvoicingagentResponseCompound) GetEEzmaxinvoicingagentVariationezmax() FieldEEzmaxinvoicingagentVariationezmax`

GetEEzmaxinvoicingagentVariationezmax returns the EEzmaxinvoicingagentVariationezmax field if non-nil, zero value otherwise.

### GetEEzmaxinvoicingagentVariationezmaxOk

`func (o *EzmaxinvoicingagentResponseCompound) GetEEzmaxinvoicingagentVariationezmaxOk() (*FieldEEzmaxinvoicingagentVariationezmax, bool)`

GetEEzmaxinvoicingagentVariationezmaxOk returns a tuple with the EEzmaxinvoicingagentVariationezmax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxinvoicingagentVariationezmax

`func (o *EzmaxinvoicingagentResponseCompound) SetEEzmaxinvoicingagentVariationezmax(v FieldEEzmaxinvoicingagentVariationezmax)`

SetEEzmaxinvoicingagentVariationezmax sets EEzmaxinvoicingagentVariationezmax field to given value.


### GetBEzmaxinvoicingagentBillableezsign

`func (o *EzmaxinvoicingagentResponseCompound) GetBEzmaxinvoicingagentBillableezsign() bool`

GetBEzmaxinvoicingagentBillableezsign returns the BEzmaxinvoicingagentBillableezsign field if non-nil, zero value otherwise.

### GetBEzmaxinvoicingagentBillableezsignOk

`func (o *EzmaxinvoicingagentResponseCompound) GetBEzmaxinvoicingagentBillableezsignOk() (*bool, bool)`

GetBEzmaxinvoicingagentBillableezsignOk returns a tuple with the BEzmaxinvoicingagentBillableezsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicingagentBillableezsign

`func (o *EzmaxinvoicingagentResponseCompound) SetBEzmaxinvoicingagentBillableezsign(v bool)`

SetBEzmaxinvoicingagentBillableezsign sets BEzmaxinvoicingagentBillableezsign field to given value.


### GetEEzmaxinvoicingagentVariationezsign

`func (o *EzmaxinvoicingagentResponseCompound) GetEEzmaxinvoicingagentVariationezsign() FieldEEzmaxinvoicingagentVariationezsign`

GetEEzmaxinvoicingagentVariationezsign returns the EEzmaxinvoicingagentVariationezsign field if non-nil, zero value otherwise.

### GetEEzmaxinvoicingagentVariationezsignOk

`func (o *EzmaxinvoicingagentResponseCompound) GetEEzmaxinvoicingagentVariationezsignOk() (*FieldEEzmaxinvoicingagentVariationezsign, bool)`

GetEEzmaxinvoicingagentVariationezsignOk returns a tuple with the EEzmaxinvoicingagentVariationezsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxinvoicingagentVariationezsign

`func (o *EzmaxinvoicingagentResponseCompound) SetEEzmaxinvoicingagentVariationezsign(v FieldEEzmaxinvoicingagentVariationezsign)`

SetEEzmaxinvoicingagentVariationezsign sets EEzmaxinvoicingagentVariationezsign field to given value.


### GetObjContactName

`func (o *EzmaxinvoicingagentResponseCompound) GetObjContactName() CustomContactNameResponse`

GetObjContactName returns the ObjContactName field if non-nil, zero value otherwise.

### GetObjContactNameOk

`func (o *EzmaxinvoicingagentResponseCompound) GetObjContactNameOk() (*CustomContactNameResponse, bool)`

GetObjContactNameOk returns a tuple with the ObjContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContactName

`func (o *EzmaxinvoicingagentResponseCompound) SetObjContactName(v CustomContactNameResponse)`

SetObjContactName sets ObjContactName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


