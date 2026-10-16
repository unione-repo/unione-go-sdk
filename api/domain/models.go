// Models of the Domain section.

package domain

import "github.com/unione-repo/unione-go-sdk/common"

// DkimStatus — Only domains with "active" DKIM record are allowed as sender domains.
type DkimStatus string

// DkimStatus values.
const (
	DkimStatusActive   DkimStatus = "active"
	DkimStatusInactive DkimStatus = "inactive"
)

// DomainGetDnsRecordsRequest
type DomainGetDnsRecordsRequest struct {
	// Domain to get DNS records for.
	Domain string `json:"domain"`
}

// DomainGetDnsRecordsResponse
type DomainGetDnsRecordsResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Domain to get DNS records for.
	Domain string `json:"domain,omitempty"`

	// Record to be added "as is" to verify ownership of this domain.
	VerificationRecord string `json:"verification-record,omitempty"`

	// DKIM signature for the domain. This property contains only the key part, you should prepend it with
	// "v=DKIM1; k=rsa; p=" for the TXT record to be valid.
	DKIM string `json:"dkim,omitempty"`

	// Target of the CNAME record for the "_dmarc" host of the domain.
	Dmarc string `json:"dmarc,omitempty"`
}

// DomainValidateVerificationRecordRequest
type DomainValidateVerificationRecordRequest struct {
	// Domain to validate verification record for.
	Domain string `json:"domain"`
}

// DomainValidateVerificationRecordResponse
type DomainValidateVerificationRecordResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Debug message.
	Message string `json:"message,omitempty"`
}

// DomainValidateDkimRequest
type DomainValidateDkimRequest struct {
	// Domain to validate DKIM record for.
	Domain string `json:"domain"`
}

// DomainValidateDkimResponse
type DomainValidateDkimResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Debug message.
	Message string `json:"message,omitempty"`
}

// DomainListRequest
type DomainListRequest struct {
	// Optional domain name if you want to get the status of single domain only.
	Domain *string `json:"domain,omitempty"`

	// Limits the number of domains returned. Default is 50.
	// Default: `50`.
	Limit *int `json:"limit,omitempty"`

	// Index of the first domain in the list, starting with 0.
	// Default: `0`.
	Offset *int `json:"offset,omitempty"`
}

// DomainListResponse
type DomainListResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Array of objects describing domains.
	Domains []DomainListResponseDomainsItem `json:"domains,omitempty"`
}

// DomainListResponseDomainsItem
type DomainListResponseDomainsItem struct {
	// Domain name.
	Domain string `json:"domain,omitempty"`

	// Object describing verification record value and status.
	VerificationRecord *DomainListResponseDomainsItemVerificationRecord `json:"verification-record,omitempty"`

	// Object describing DKIM record value and status.
	DKIM *DomainListResponseDomainsItemDkim `json:"dkim,omitempty"`
}

// DomainListResponseDomainsItemVerificationRecord — Object describing verification record value and status.
type DomainListResponseDomainsItemVerificationRecord struct {
	// Record to be added "as is" to verify ownership of this domain.
	Value string `json:"value,omitempty"`

	// Only domains with "confirmed" verification record are allowed as sender domains.
	Status common.VerificationRecordStatus `json:"status,omitempty"`
}

// DomainListResponseDomainsItemDkim — Object describing DKIM record value and status.
type DomainListResponseDomainsItemDkim struct {
	// DKIM signature for the domain. This property contains only the key part, you should prepend it with
	// "k=rsa, p=" part for the record to be valid (see
	// [domain/get-dns-records](https://docs.unione.io/en/web-api-ref#domain-get-dns-records)).
	Key string `json:"key,omitempty"`

	// Only domains with "active" DKIM record are allowed as sender domains.
	Status DkimStatus `json:"status,omitempty"`
}

// DomainGetSandboxDomainResponse
type DomainGetSandboxDomainResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Sandbox domain name.
	Domain string `json:"domain,omitempty"`
}

// DomainDeleteRequest
type DomainDeleteRequest struct {
	// Domain name to delete.
	Domain string `json:"domain"`
}

// DomainDeleteResponse
type DomainDeleteResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Domain deleted
	Message string `json:"message,omitempty"`
}
