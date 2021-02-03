# FranchisereferalincomeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiFranchisebrokerID** | **int32** | The unique ID of the Franchisebroker | 
**FkiFranchisereferalincomeprogramID** | **int32** | The unique ID of the Franchisereferalincomeprogram | 
**FkiPeriodID** | **int32** | The unique ID of the Period | 
**DFranchisereferalincomeLoan** | **string** | The loan amount | 
**DFranchisereferalincomeFranchiseamount** | **string** | The amount that will be given to the franchise | 
**DFranchisereferalincomeFranchisoramount** | **string** | The amount that will be kept by the franchisor | 
**DFranchisereferalincomeAgentamount** | **string** | The amount that will be given to the agent | 
**DtFranchisereferalincomeDisbursed** | **string** | The date the amounts were disbursed | 
**TFranchisereferalincomeComment** | **string** | A comment about the transaction | 
**FkiFranchiseofficeID** | **int32** | The unique ID of the Franchisereoffice | 
**SFranchisereferalincomeRemoteid** | **string** |  | 

## Methods

### NewFranchisereferalincomeRequest

`func NewFranchisereferalincomeRequest(fkiFranchisebrokerID int32, fkiFranchisereferalincomeprogramID int32, fkiPeriodID int32, dFranchisereferalincomeLoan string, dFranchisereferalincomeFranchiseamount string, dFranchisereferalincomeFranchisoramount string, dFranchisereferalincomeAgentamount string, dtFranchisereferalincomeDisbursed string, tFranchisereferalincomeComment string, fkiFranchiseofficeID int32, sFranchisereferalincomeRemoteid string, ) *FranchisereferalincomeRequest`

NewFranchisereferalincomeRequest instantiates a new FranchisereferalincomeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFranchisereferalincomeRequestWithDefaults

`func NewFranchisereferalincomeRequestWithDefaults() *FranchisereferalincomeRequest`

NewFranchisereferalincomeRequestWithDefaults instantiates a new FranchisereferalincomeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiFranchisebrokerID

`func (o *FranchisereferalincomeRequest) GetFkiFranchisebrokerID() int32`

GetFkiFranchisebrokerID returns the FkiFranchisebrokerID field if non-nil, zero value otherwise.

### GetFkiFranchisebrokerIDOk

`func (o *FranchisereferalincomeRequest) GetFkiFranchisebrokerIDOk() (*int32, bool)`

GetFkiFranchisebrokerIDOk returns a tuple with the FkiFranchisebrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisebrokerID

`func (o *FranchisereferalincomeRequest) SetFkiFranchisebrokerID(v int32)`

SetFkiFranchisebrokerID sets FkiFranchisebrokerID field to given value.


### GetFkiFranchisereferalincomeprogramID

`func (o *FranchisereferalincomeRequest) GetFkiFranchisereferalincomeprogramID() int32`

GetFkiFranchisereferalincomeprogramID returns the FkiFranchisereferalincomeprogramID field if non-nil, zero value otherwise.

### GetFkiFranchisereferalincomeprogramIDOk

`func (o *FranchisereferalincomeRequest) GetFkiFranchisereferalincomeprogramIDOk() (*int32, bool)`

GetFkiFranchisereferalincomeprogramIDOk returns a tuple with the FkiFranchisereferalincomeprogramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisereferalincomeprogramID

`func (o *FranchisereferalincomeRequest) SetFkiFranchisereferalincomeprogramID(v int32)`

SetFkiFranchisereferalincomeprogramID sets FkiFranchisereferalincomeprogramID field to given value.


### GetFkiPeriodID

`func (o *FranchisereferalincomeRequest) GetFkiPeriodID() int32`

GetFkiPeriodID returns the FkiPeriodID field if non-nil, zero value otherwise.

### GetFkiPeriodIDOk

`func (o *FranchisereferalincomeRequest) GetFkiPeriodIDOk() (*int32, bool)`

GetFkiPeriodIDOk returns a tuple with the FkiPeriodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPeriodID

`func (o *FranchisereferalincomeRequest) SetFkiPeriodID(v int32)`

SetFkiPeriodID sets FkiPeriodID field to given value.


### GetDFranchisereferalincomeLoan

`func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeLoan() string`

GetDFranchisereferalincomeLoan returns the DFranchisereferalincomeLoan field if non-nil, zero value otherwise.

### GetDFranchisereferalincomeLoanOk

`func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeLoanOk() (*string, bool)`

GetDFranchisereferalincomeLoanOk returns a tuple with the DFranchisereferalincomeLoan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDFranchisereferalincomeLoan

`func (o *FranchisereferalincomeRequest) SetDFranchisereferalincomeLoan(v string)`

SetDFranchisereferalincomeLoan sets DFranchisereferalincomeLoan field to given value.


### GetDFranchisereferalincomeFranchiseamount

`func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeFranchiseamount() string`

GetDFranchisereferalincomeFranchiseamount returns the DFranchisereferalincomeFranchiseamount field if non-nil, zero value otherwise.

### GetDFranchisereferalincomeFranchiseamountOk

`func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeFranchiseamountOk() (*string, bool)`

GetDFranchisereferalincomeFranchiseamountOk returns a tuple with the DFranchisereferalincomeFranchiseamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDFranchisereferalincomeFranchiseamount

`func (o *FranchisereferalincomeRequest) SetDFranchisereferalincomeFranchiseamount(v string)`

SetDFranchisereferalincomeFranchiseamount sets DFranchisereferalincomeFranchiseamount field to given value.


### GetDFranchisereferalincomeFranchisoramount

`func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeFranchisoramount() string`

GetDFranchisereferalincomeFranchisoramount returns the DFranchisereferalincomeFranchisoramount field if non-nil, zero value otherwise.

### GetDFranchisereferalincomeFranchisoramountOk

`func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeFranchisoramountOk() (*string, bool)`

GetDFranchisereferalincomeFranchisoramountOk returns a tuple with the DFranchisereferalincomeFranchisoramount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDFranchisereferalincomeFranchisoramount

`func (o *FranchisereferalincomeRequest) SetDFranchisereferalincomeFranchisoramount(v string)`

SetDFranchisereferalincomeFranchisoramount sets DFranchisereferalincomeFranchisoramount field to given value.


### GetDFranchisereferalincomeAgentamount

`func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeAgentamount() string`

GetDFranchisereferalincomeAgentamount returns the DFranchisereferalincomeAgentamount field if non-nil, zero value otherwise.

### GetDFranchisereferalincomeAgentamountOk

`func (o *FranchisereferalincomeRequest) GetDFranchisereferalincomeAgentamountOk() (*string, bool)`

GetDFranchisereferalincomeAgentamountOk returns a tuple with the DFranchisereferalincomeAgentamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDFranchisereferalincomeAgentamount

`func (o *FranchisereferalincomeRequest) SetDFranchisereferalincomeAgentamount(v string)`

SetDFranchisereferalincomeAgentamount sets DFranchisereferalincomeAgentamount field to given value.


### GetDtFranchisereferalincomeDisbursed

`func (o *FranchisereferalincomeRequest) GetDtFranchisereferalincomeDisbursed() string`

GetDtFranchisereferalincomeDisbursed returns the DtFranchisereferalincomeDisbursed field if non-nil, zero value otherwise.

### GetDtFranchisereferalincomeDisbursedOk

`func (o *FranchisereferalincomeRequest) GetDtFranchisereferalincomeDisbursedOk() (*string, bool)`

GetDtFranchisereferalincomeDisbursedOk returns a tuple with the DtFranchisereferalincomeDisbursed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtFranchisereferalincomeDisbursed

`func (o *FranchisereferalincomeRequest) SetDtFranchisereferalincomeDisbursed(v string)`

SetDtFranchisereferalincomeDisbursed sets DtFranchisereferalincomeDisbursed field to given value.


### GetTFranchisereferalincomeComment

`func (o *FranchisereferalincomeRequest) GetTFranchisereferalincomeComment() string`

GetTFranchisereferalincomeComment returns the TFranchisereferalincomeComment field if non-nil, zero value otherwise.

### GetTFranchisereferalincomeCommentOk

`func (o *FranchisereferalincomeRequest) GetTFranchisereferalincomeCommentOk() (*string, bool)`

GetTFranchisereferalincomeCommentOk returns a tuple with the TFranchisereferalincomeComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTFranchisereferalincomeComment

`func (o *FranchisereferalincomeRequest) SetTFranchisereferalincomeComment(v string)`

SetTFranchisereferalincomeComment sets TFranchisereferalincomeComment field to given value.


### GetFkiFranchiseofficeID

`func (o *FranchisereferalincomeRequest) GetFkiFranchiseofficeID() int32`

GetFkiFranchiseofficeID returns the FkiFranchiseofficeID field if non-nil, zero value otherwise.

### GetFkiFranchiseofficeIDOk

`func (o *FranchisereferalincomeRequest) GetFkiFranchiseofficeIDOk() (*int32, bool)`

GetFkiFranchiseofficeIDOk returns a tuple with the FkiFranchiseofficeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseofficeID

`func (o *FranchisereferalincomeRequest) SetFkiFranchiseofficeID(v int32)`

SetFkiFranchiseofficeID sets FkiFranchiseofficeID field to given value.


### GetSFranchisereferalincomeRemoteid

`func (o *FranchisereferalincomeRequest) GetSFranchisereferalincomeRemoteid() string`

GetSFranchisereferalincomeRemoteid returns the SFranchisereferalincomeRemoteid field if non-nil, zero value otherwise.

### GetSFranchisereferalincomeRemoteidOk

`func (o *FranchisereferalincomeRequest) GetSFranchisereferalincomeRemoteidOk() (*string, bool)`

GetSFranchisereferalincomeRemoteidOk returns a tuple with the SFranchisereferalincomeRemoteid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSFranchisereferalincomeRemoteid

`func (o *FranchisereferalincomeRequest) SetSFranchisereferalincomeRemoteid(v string)`

SetSFranchisereferalincomeRemoteid sets SFranchisereferalincomeRemoteid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


