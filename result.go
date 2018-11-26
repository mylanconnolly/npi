package npi

type listResult struct {
	ResultCount int      `json:"result_count"`
	Results     []Result `json:"results"`
}

// Result is used to represent an NPI lookup result.
type Result struct {
	Number           int    `json:"number"`
	CreatedEpoch     int    `json:"created_epoch"`
	LastUpdatedEpoch int    `json:"last_updated_epoch"`
	EnumerationType  string `json:"enumeration_type"`
	Basic            Basic  `json:"basic"`
	// TODO: Implement addresses and taxonomies
}

// Basic includes basic result information.
type Basic struct {
	Status          string `json:"status"`
	Credential      string `json:"credential"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	MiddleNmae      string `json:"middle_name"`
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
