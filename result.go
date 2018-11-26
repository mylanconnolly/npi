package npi

type listResult struct {
	ResultCount int      `json:"result_count"`
	Results     []Result `json:"results"`
}

// Result is used to represent an NPI lookup result.
type Result struct {
	Number           int       `json:"number"`
	CreatedEpoch     int       `json:"created_epoch"`
	LastUpdatedEpoch int       `json:"last_updated_epoch"`
	EnumerationType  string    `json:"enumeration_type"`
	Basic            Basic     `json:"basic"`
	Addresses        []Address `json:"addresses"`
	// TODO: Implement addresses and taxonomies
}

// Basic includes basic result information.
type Basic struct {
	Status          string `json:"status"`
	Credential      string `json:"credential"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	MiddleName      string `json:"middle_name"`
	Name            string `json:"name"`
	Gender          string `json:"gender"`
	SoleProprietor  string `json:"sole_proprietor"`  // Should this get changed to a bool?
	LastUpdated     string `json:"last_updated"`     // Should this get changed to a time.Time?
	EnumerationDate string `json:"enumeration_date"` // Should this get changed to a time.Time?
}

// Identifier is used to represent identifiers for the result.
type Identifier struct {
	Code        string `json:"code"`
	Issuer      string `json:"issuer"`
	State       string `json:"state"`
	Identifier  string `json:"identifier"`
	Description string `json:"desc"`
}

// Address is used to represent an address in the NPI record.
type Address struct {
	Address1    string `json:"address_1"`
	Address2    string `json:"address_2"`
	City        string `json:"city"`
	State       string `json:"state"`
	ZIP         string `json:"postal_code"`
	Phone       string `json:"telephone_number"`
	Fax         string `json:"fax_number"`
	AddressType string `json:"address_type"`
	Purpose     string `json:"address_purpose"`
}

// GetAddress is used to fetch the address with the given purpose, as well as a
// boolean flag to denote whether it was found or not.
func (r Result) GetAddress(purpose string) (Address, bool) {
	for _, a := range r.Addresses {
		if a.Purpose == purpose {
			return a, true
		}
	}
	return Address{}, false
}
