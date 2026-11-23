# EzmaxpartnerproductstageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzmaxpartnerproductstageID** | **int32** | The unique ID of the Ezmaxpartnerproductstage | 
**FkiEzmaxpartnerproductID** | **int32** | The unique ID of the Ezmaxpartnerproduct | 
**EEzmaxpartnerproductstageCode** | [**FieldEEzmaxpartnerproductstageCode**](FieldEEzmaxpartnerproductstageCode.md) |  | 
**EEzmaxpartnerproductstageInstallationtype** | [**FieldEEzmaxpartnerproductstageInstallationtype**](FieldEEzmaxpartnerproductstageInstallationtype.md) |  | 
**BEzmaxpartnerproductstagePublished** | **bool** | Whether the Ezmaxpartnerproducstage is published or not | 
**BEzmaxpartnerproductstageInstalled** | **bool** | Whether the Ezmaxpartnerproductstage is installed or not | 
**ObjEzmaxpartnerproductstagesubscription** | Pointer to [**EzmaxpartnerproductstagesubscriptionResponse**](EzmaxpartnerproductstagesubscriptionResponse.md) |  | [optional] 

## Methods

### NewEzmaxpartnerproductstageResponse

`func NewEzmaxpartnerproductstageResponse(pkiEzmaxpartnerproductstageID int32, fkiEzmaxpartnerproductID int32, eEzmaxpartnerproductstageCode FieldEEzmaxpartnerproductstageCode, eEzmaxpartnerproductstageInstallationtype FieldEEzmaxpartnerproductstageInstallationtype, bEzmaxpartnerproductstagePublished bool, bEzmaxpartnerproductstageInstalled bool, ) *EzmaxpartnerproductstageResponse`

NewEzmaxpartnerproductstageResponse instantiates a new EzmaxpartnerproductstageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxpartnerproductstageResponseWithDefaults

`func NewEzmaxpartnerproductstageResponseWithDefaults() *EzmaxpartnerproductstageResponse`

NewEzmaxpartnerproductstageResponseWithDefaults instantiates a new EzmaxpartnerproductstageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxpartnerproductstageID

`func (o *EzmaxpartnerproductstageResponse) GetPkiEzmaxpartnerproductstageID() int32`

GetPkiEzmaxpartnerproductstageID returns the PkiEzmaxpartnerproductstageID field if non-nil, zero value otherwise.

### GetPkiEzmaxpartnerproductstageIDOk

`func (o *EzmaxpartnerproductstageResponse) GetPkiEzmaxpartnerproductstageIDOk() (*int32, bool)`

GetPkiEzmaxpartnerproductstageIDOk returns a tuple with the PkiEzmaxpartnerproductstageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxpartnerproductstageID

`func (o *EzmaxpartnerproductstageResponse) SetPkiEzmaxpartnerproductstageID(v int32)`

SetPkiEzmaxpartnerproductstageID sets PkiEzmaxpartnerproductstageID field to given value.


### GetFkiEzmaxpartnerproductID

`func (o *EzmaxpartnerproductstageResponse) GetFkiEzmaxpartnerproductID() int32`

GetFkiEzmaxpartnerproductID returns the FkiEzmaxpartnerproductID field if non-nil, zero value otherwise.

### GetFkiEzmaxpartnerproductIDOk

`func (o *EzmaxpartnerproductstageResponse) GetFkiEzmaxpartnerproductIDOk() (*int32, bool)`

GetFkiEzmaxpartnerproductIDOk returns a tuple with the FkiEzmaxpartnerproductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxpartnerproductID

`func (o *EzmaxpartnerproductstageResponse) SetFkiEzmaxpartnerproductID(v int32)`

SetFkiEzmaxpartnerproductID sets FkiEzmaxpartnerproductID field to given value.


### GetEEzmaxpartnerproductstageCode

`func (o *EzmaxpartnerproductstageResponse) GetEEzmaxpartnerproductstageCode() FieldEEzmaxpartnerproductstageCode`

GetEEzmaxpartnerproductstageCode returns the EEzmaxpartnerproductstageCode field if non-nil, zero value otherwise.

### GetEEzmaxpartnerproductstageCodeOk

`func (o *EzmaxpartnerproductstageResponse) GetEEzmaxpartnerproductstageCodeOk() (*FieldEEzmaxpartnerproductstageCode, bool)`

GetEEzmaxpartnerproductstageCodeOk returns a tuple with the EEzmaxpartnerproductstageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxpartnerproductstageCode

`func (o *EzmaxpartnerproductstageResponse) SetEEzmaxpartnerproductstageCode(v FieldEEzmaxpartnerproductstageCode)`

SetEEzmaxpartnerproductstageCode sets EEzmaxpartnerproductstageCode field to given value.


### GetEEzmaxpartnerproductstageInstallationtype

`func (o *EzmaxpartnerproductstageResponse) GetEEzmaxpartnerproductstageInstallationtype() FieldEEzmaxpartnerproductstageInstallationtype`

GetEEzmaxpartnerproductstageInstallationtype returns the EEzmaxpartnerproductstageInstallationtype field if non-nil, zero value otherwise.

### GetEEzmaxpartnerproductstageInstallationtypeOk

`func (o *EzmaxpartnerproductstageResponse) GetEEzmaxpartnerproductstageInstallationtypeOk() (*FieldEEzmaxpartnerproductstageInstallationtype, bool)`

GetEEzmaxpartnerproductstageInstallationtypeOk returns a tuple with the EEzmaxpartnerproductstageInstallationtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxpartnerproductstageInstallationtype

`func (o *EzmaxpartnerproductstageResponse) SetEEzmaxpartnerproductstageInstallationtype(v FieldEEzmaxpartnerproductstageInstallationtype)`

SetEEzmaxpartnerproductstageInstallationtype sets EEzmaxpartnerproductstageInstallationtype field to given value.


### GetBEzmaxpartnerproductstagePublished

`func (o *EzmaxpartnerproductstageResponse) GetBEzmaxpartnerproductstagePublished() bool`

GetBEzmaxpartnerproductstagePublished returns the BEzmaxpartnerproductstagePublished field if non-nil, zero value otherwise.

### GetBEzmaxpartnerproductstagePublishedOk

`func (o *EzmaxpartnerproductstageResponse) GetBEzmaxpartnerproductstagePublishedOk() (*bool, bool)`

GetBEzmaxpartnerproductstagePublishedOk returns a tuple with the BEzmaxpartnerproductstagePublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxpartnerproductstagePublished

`func (o *EzmaxpartnerproductstageResponse) SetBEzmaxpartnerproductstagePublished(v bool)`

SetBEzmaxpartnerproductstagePublished sets BEzmaxpartnerproductstagePublished field to given value.


### GetBEzmaxpartnerproductstageInstalled

`func (o *EzmaxpartnerproductstageResponse) GetBEzmaxpartnerproductstageInstalled() bool`

GetBEzmaxpartnerproductstageInstalled returns the BEzmaxpartnerproductstageInstalled field if non-nil, zero value otherwise.

### GetBEzmaxpartnerproductstageInstalledOk

`func (o *EzmaxpartnerproductstageResponse) GetBEzmaxpartnerproductstageInstalledOk() (*bool, bool)`

GetBEzmaxpartnerproductstageInstalledOk returns a tuple with the BEzmaxpartnerproductstageInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxpartnerproductstageInstalled

`func (o *EzmaxpartnerproductstageResponse) SetBEzmaxpartnerproductstageInstalled(v bool)`

SetBEzmaxpartnerproductstageInstalled sets BEzmaxpartnerproductstageInstalled field to given value.


### GetObjEzmaxpartnerproductstagesubscription

`func (o *EzmaxpartnerproductstageResponse) GetObjEzmaxpartnerproductstagesubscription() EzmaxpartnerproductstagesubscriptionResponse`

GetObjEzmaxpartnerproductstagesubscription returns the ObjEzmaxpartnerproductstagesubscription field if non-nil, zero value otherwise.

### GetObjEzmaxpartnerproductstagesubscriptionOk

`func (o *EzmaxpartnerproductstageResponse) GetObjEzmaxpartnerproductstagesubscriptionOk() (*EzmaxpartnerproductstagesubscriptionResponse, bool)`

GetObjEzmaxpartnerproductstagesubscriptionOk returns a tuple with the ObjEzmaxpartnerproductstagesubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzmaxpartnerproductstagesubscription

`func (o *EzmaxpartnerproductstageResponse) SetObjEzmaxpartnerproductstagesubscription(v EzmaxpartnerproductstagesubscriptionResponse)`

SetObjEzmaxpartnerproductstagesubscription sets ObjEzmaxpartnerproductstagesubscription field to given value.

### HasObjEzmaxpartnerproductstagesubscription

`func (o *EzmaxpartnerproductstageResponse) HasObjEzmaxpartnerproductstagesubscription() bool`

HasObjEzmaxpartnerproductstagesubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


