# CustomInscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiInscriptionID** | **int32** | The unique ID of the Inscription. | 
**FkiAddressID** | **int32** | The unique ID of the Address | 
**ObjAddress** | [**AddressResponseCompound**](AddressResponseCompound.md) |  | 

## Methods

### NewCustomInscription

`func NewCustomInscription(pkiInscriptionID int32, fkiAddressID int32, objAddress AddressResponseCompound, ) *CustomInscription`

NewCustomInscription instantiates a new CustomInscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomInscriptionWithDefaults

`func NewCustomInscriptionWithDefaults() *CustomInscription`

NewCustomInscriptionWithDefaults instantiates a new CustomInscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiInscriptionID

`func (o *CustomInscription) GetPkiInscriptionID() int32`

GetPkiInscriptionID returns the PkiInscriptionID field if non-nil, zero value otherwise.

### GetPkiInscriptionIDOk

`func (o *CustomInscription) GetPkiInscriptionIDOk() (*int32, bool)`

GetPkiInscriptionIDOk returns a tuple with the PkiInscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiInscriptionID

`func (o *CustomInscription) SetPkiInscriptionID(v int32)`

SetPkiInscriptionID sets PkiInscriptionID field to given value.


### GetFkiAddressID

`func (o *CustomInscription) GetFkiAddressID() int32`

GetFkiAddressID returns the FkiAddressID field if non-nil, zero value otherwise.

### GetFkiAddressIDOk

`func (o *CustomInscription) GetFkiAddressIDOk() (*int32, bool)`

GetFkiAddressIDOk returns a tuple with the FkiAddressID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAddressID

`func (o *CustomInscription) SetFkiAddressID(v int32)`

SetFkiAddressID sets FkiAddressID field to given value.


### GetObjAddress

`func (o *CustomInscription) GetObjAddress() AddressResponseCompound`

GetObjAddress returns the ObjAddress field if non-nil, zero value otherwise.

### GetObjAddressOk

`func (o *CustomInscription) GetObjAddressOk() (*AddressResponseCompound, bool)`

GetObjAddressOk returns a tuple with the ObjAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAddress

`func (o *CustomInscription) SetObjAddress(v AddressResponseCompound)`

SetObjAddress sets ObjAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


