package wallet

// AddressSize of array used to store Bitcoin addresses.
const AddressSize = 20

// Address represents a Bitcoin address.
type Address [AddressSize]byte
