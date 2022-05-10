// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CertificateAdditionalExtensions struct {
	/* Immutable. Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error). */
	// +optional
	Critical *bool `json:"critical,omitempty"`

	/* Immutable. Required. The OID for this X.509 extension. */
	ObjectId CertificateObjectId `json:"objectId"`

	/* Immutable. Required. The value of this X.509 extension. */
	Value string `json:"value"`
}

type CertificateBaseKeyUsage struct {
	/* Immutable. The key may be used to sign certificates. */
	// +optional
	CertSign *bool `json:"certSign,omitempty"`

	/* Immutable. The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation". */
	// +optional
	ContentCommitment *bool `json:"contentCommitment,omitempty"`

	/* Immutable. The key may be used sign certificate revocation lists. */
	// +optional
	CrlSign *bool `json:"crlSign,omitempty"`

	/* Immutable. The key may be used to encipher data. */
	// +optional
	DataEncipherment *bool `json:"dataEncipherment,omitempty"`

	/* Immutable. The key may be used to decipher only. */
	// +optional
	DecipherOnly *bool `json:"decipherOnly,omitempty"`

	/* Immutable. The key may be used for digital signatures. */
	// +optional
	DigitalSignature *bool `json:"digitalSignature,omitempty"`

	/* Immutable. The key may be used to encipher only. */
	// +optional
	EncipherOnly *bool `json:"encipherOnly,omitempty"`

	/* Immutable. The key may be used in a key agreement protocol. */
	// +optional
	KeyAgreement *bool `json:"keyAgreement,omitempty"`

	/* Immutable. The key may be used to encipher other keys. */
	// +optional
	KeyEncipherment *bool `json:"keyEncipherment,omitempty"`
}

type CertificateCaOptions struct {
	/* Immutable. Optional. When true, the "CA" in Basic Constraints extension will be set to true. */
	// +optional
	IsCa *bool `json:"isCa,omitempty"`

	/* Immutable. Optional. Refers to the "path length constraint" in Basic Constraints extension. For a CA certificate, this value describes the depth of subordinate CA certificates that are allowed. If this value is less than 0, the request will fail. */
	// +optional
	MaxIssuerPathLength *int `json:"maxIssuerPathLength,omitempty"`

	/* Immutable. Optional. When true, the "CA" in Basic Constraints extension will be set to false. If both `is_ca` and `non_ca` are unset, the extension will be omitted from the CA certificate. */
	// +optional
	NonCa *bool `json:"nonCa,omitempty"`

	/* Immutable. Optional. When true, the "path length constraint" in Basic Constraints extension will be set to 0. if both max_issuer_path_length and zero_max_issuer_path_length are unset, the max path length will be omitted from the CA certificate. */
	// +optional
	ZeroMaxIssuerPathLength *bool `json:"zeroMaxIssuerPathLength,omitempty"`
}

type CertificateConfig struct {
	/* Immutable. Optional. The public key that corresponds to this config. This is, for example, used when issuing Certificates, but not when creating a self-signed CertificateAuthority or CertificateAuthority CSR. */
	// +optional
	PublicKey *CertificatePublicKey `json:"publicKey,omitempty"`

	/* Immutable. Required. Specifies some of the values in a certificate that are related to the subject. */
	SubjectConfig CertificateSubjectConfig `json:"subjectConfig"`

	/* Immutable. Required. Describes how some of the technical X.509 fields in a certificate should be populated. */
	X509Config CertificateX509Config `json:"x509Config"`
}

type CertificateExtendedKeyUsage struct {
	/* Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS. */
	// +optional
	ClientAuth *bool `json:"clientAuth,omitempty"`

	/* Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication". */
	// +optional
	CodeSigning *bool `json:"codeSigning,omitempty"`

	/* Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection". */
	// +optional
	EmailProtection *bool `json:"emailProtection,omitempty"`

	/* Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses". */
	// +optional
	OcspSigning *bool `json:"ocspSigning,omitempty"`

	/* Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS. */
	// +optional
	ServerAuth *bool `json:"serverAuth,omitempty"`

	/* Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time". */
	// +optional
	TimeStamping *bool `json:"timeStamping,omitempty"`
}

type CertificateKeyUsage struct {
	/* Immutable. Describes high-level ways in which a key may be used. */
	// +optional
	BaseKeyUsage *CertificateBaseKeyUsage `json:"baseKeyUsage,omitempty"`

	/* Immutable. Detailed scenarios in which a key may be used. */
	// +optional
	ExtendedKeyUsage *CertificateExtendedKeyUsage `json:"extendedKeyUsage,omitempty"`

	/* Immutable. Used to describe extended key usages that are not listed in the KeyUsage.ExtendedKeyUsageOptions message. */
	// +optional
	UnknownExtendedKeyUsages []CertificateUnknownExtendedKeyUsages `json:"unknownExtendedKeyUsages,omitempty"`
}

type CertificateObjectId struct {
	/* Immutable. Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath"`
}

type CertificatePolicyIds struct {
	/* Immutable. Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath"`
}

type CertificatePublicKey struct {
	/* Immutable. Required. The format of the public key. Possible values: KEY_FORMAT_UNSPECIFIED, PEM */
	Format string `json:"format"`

	/* Immutable. Required. A public key. The padding and encoding must match with the `KeyFormat` value specified for the `format` field. */
	Key string `json:"key"`
}

type CertificateSubject struct {
	/* Immutable. The "common name" of the subject. */
	// +optional
	CommonName *string `json:"commonName,omitempty"`

	/* Immutable. The country code of the subject. */
	// +optional
	CountryCode *string `json:"countryCode,omitempty"`

	/* Immutable. The locality or city of the subject. */
	// +optional
	Locality *string `json:"locality,omitempty"`

	/* Immutable. The organization of the subject. */
	// +optional
	Organization *string `json:"organization,omitempty"`

	/* Immutable. The organizational_unit of the subject. */
	// +optional
	OrganizationalUnit *string `json:"organizationalUnit,omitempty"`

	/* Immutable. The postal code of the subject. */
	// +optional
	PostalCode *string `json:"postalCode,omitempty"`

	/* Immutable. The province, territory, or regional state of the subject. */
	// +optional
	Province *string `json:"province,omitempty"`

	/* Immutable. The street address of the subject. */
	// +optional
	StreetAddress *string `json:"streetAddress,omitempty"`
}

type CertificateSubjectAltName struct {
	/* Immutable. Contains only valid, fully-qualified host names. */
	// +optional
	DnsNames []string `json:"dnsNames,omitempty"`

	/* Immutable. Contains only valid RFC 2822 E-mail addresses. */
	// +optional
	EmailAddresses []string `json:"emailAddresses,omitempty"`

	/* Immutable. Contains only valid 32-bit IPv4 addresses or RFC 4291 IPv6 addresses. */
	// +optional
	IpAddresses []string `json:"ipAddresses,omitempty"`

	/* Immutable. Contains only valid RFC 3986 URIs. */
	// +optional
	Uris []string `json:"uris,omitempty"`
}

type CertificateSubjectConfig struct {
	/* Immutable. Required. Contains distinguished name fields such as the common name, location and organization. */
	Subject CertificateSubject `json:"subject"`

	/* Immutable. Optional. The subject alternative name fields. */
	// +optional
	SubjectAltName *CertificateSubjectAltName `json:"subjectAltName,omitempty"`
}

type CertificateUnknownExtendedKeyUsages struct {
	/* Immutable. Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath"`
}

type CertificateX509Config struct {
	/* Immutable. Optional. Describes custom X.509 extensions. */
	// +optional
	AdditionalExtensions []CertificateAdditionalExtensions `json:"additionalExtensions,omitempty"`

	/* Immutable. Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate. */
	// +optional
	AiaOcspServers []string `json:"aiaOcspServers,omitempty"`

	/* Immutable. Optional. Describes options in this X509Parameters that are relevant in a CA certificate. */
	// +optional
	CaOptions *CertificateCaOptions `json:"caOptions,omitempty"`

	/* Immutable. Optional. Indicates the intended use for keys that correspond to a certificate. */
	// +optional
	KeyUsage *CertificateKeyUsage `json:"keyUsage,omitempty"`

	/* Immutable. Optional. Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4. */
	// +optional
	PolicyIds []CertificatePolicyIds `json:"policyIds,omitempty"`
}

type PrivateCACertificateSpec struct {
	/* Immutable. */
	CaPoolRef v1alpha1.ResourceRef `json:"caPoolRef"`

	/* Immutable. */
	// +optional
	CertificateAuthorityRef *v1alpha1.ResourceRef `json:"certificateAuthorityRef,omitempty"`

	/* Immutable. */
	// +optional
	CertificateTemplateRef *v1alpha1.ResourceRef `json:"certificateTemplateRef,omitempty"`

	/* Immutable. Immutable. A description of the certificate and key that does not require X.509 or ASN.1. */
	// +optional
	Config *CertificateConfig `json:"config,omitempty"`

	/* Immutable. Required. Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain. */
	Lifetime string `json:"lifetime"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/* Immutable. Immutable. A pem-encoded X.509 certificate signing request (CSR). */
	// +optional
	PemCsr *string `json:"pemCsr,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. Immutable. Specifies how the Certificate's identity fields are to be decided. If this is omitted, the `DEFAULT` subject mode will be used. Possible values: SUBJECT_REQUEST_MODE_UNSPECIFIED, DEFAULT, REFLECTED_SPIFFE */
	// +optional
	SubjectMode *string `json:"subjectMode,omitempty"`
}

type CertificateAdditionalExtensionsStatus struct {
	/* Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error). */
	Critical bool `json:"critical,omitempty"`

	/* Required. The OID for this X.509 extension. */
	ObjectId CertificateObjectIdStatus `json:"objectId,omitempty"`

	/* Required. The value of this X.509 extension. */
	Value string `json:"value,omitempty"`
}

type CertificateAuthorityKeyIdStatus struct {
	/* Optional. The value of this KeyId encoded in lowercase hexadecimal. This is most likely the 160 bit SHA-1 hash of the public key. */
	KeyId string `json:"keyId,omitempty"`
}

type CertificateBaseKeyUsageStatus struct {
	/* The key may be used to sign certificates. */
	CertSign bool `json:"certSign,omitempty"`

	/* The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation". */
	ContentCommitment bool `json:"contentCommitment,omitempty"`

	/* The key may be used sign certificate revocation lists. */
	CrlSign bool `json:"crlSign,omitempty"`

	/* The key may be used to encipher data. */
	DataEncipherment bool `json:"dataEncipherment,omitempty"`

	/* The key may be used to decipher only. */
	DecipherOnly bool `json:"decipherOnly,omitempty"`

	/* The key may be used for digital signatures. */
	DigitalSignature bool `json:"digitalSignature,omitempty"`

	/* The key may be used to encipher only. */
	EncipherOnly bool `json:"encipherOnly,omitempty"`

	/* The key may be used in a key agreement protocol. */
	KeyAgreement bool `json:"keyAgreement,omitempty"`

	/* The key may be used to encipher other keys. */
	KeyEncipherment bool `json:"keyEncipherment,omitempty"`
}

type CertificateCaOptionsStatus struct {
	/* Optional. Refers to the "CA" X.509 extension, which is a boolean value. When this value is missing, the extension will be omitted from the CA certificate. */
	IsCa bool `json:"isCa,omitempty"`

	/* Optional. Refers to the path length restriction X.509 extension. For a CA certificate, this value describes the depth of subordinate CA certificates that are allowed. If this value is less than 0, the request will fail. If this value is missing, the max path length will be omitted from the CA certificate. */
	MaxIssuerPathLength int `json:"maxIssuerPathLength,omitempty"`
}

type CertificateCertFingerprintStatus struct {
	/* The SHA 256 hash, encoded in hexadecimal, of the DER x509 certificate. */
	Sha256Hash string `json:"sha256Hash,omitempty"`
}

type CertificateCertificateDescriptionStatus struct {
	/* Describes lists of issuer CA certificate URLs that appear in the "Authority Information Access" extension in the certificate. */
	AiaIssuingCertificateUrls []string `json:"aiaIssuingCertificateUrls,omitempty"`

	/* Identifies the subject_key_id of the parent certificate, per https://tools.ietf.org/html/rfc5280#section-4.2.1.1 */
	AuthorityKeyId CertificateAuthorityKeyIdStatus `json:"authorityKeyId,omitempty"`

	/* The hash of the x.509 certificate. */
	CertFingerprint CertificateCertFingerprintStatus `json:"certFingerprint,omitempty"`

	/* Describes a list of locations to obtain CRL information, i.e. the DistributionPoint.fullName described by https://tools.ietf.org/html/rfc5280#section-4.2.1.13 */
	CrlDistributionPoints []string `json:"crlDistributionPoints,omitempty"`

	/* The public key that corresponds to an issued certificate. */
	PublicKey CertificatePublicKeyStatus `json:"publicKey,omitempty"`

	/* Describes some of the values in a certificate that are related to the subject and lifetime. */
	SubjectDescription CertificateSubjectDescriptionStatus `json:"subjectDescription,omitempty"`

	/* Provides a means of identifiying certificates that contain a particular public key, per https://tools.ietf.org/html/rfc5280#section-4.2.1.2. */
	SubjectKeyId CertificateSubjectKeyIdStatus `json:"subjectKeyId,omitempty"`

	/* Describes some of the technical X.509 fields in a certificate. */
	X509Description CertificateX509DescriptionStatus `json:"x509Description,omitempty"`
}

type CertificateCustomSansStatus struct {
	/* Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error). */
	Critical bool `json:"critical,omitempty"`

	/* Required. The OID for this X.509 extension. */
	ObjectId CertificateObjectIdStatus `json:"objectId,omitempty"`

	/* Required. The value of this X.509 extension. */
	Value string `json:"value,omitempty"`
}

type CertificateExtendedKeyUsageStatus struct {
	/* Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS. */
	ClientAuth bool `json:"clientAuth,omitempty"`

	/* Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication". */
	CodeSigning bool `json:"codeSigning,omitempty"`

	/* Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection". */
	EmailProtection bool `json:"emailProtection,omitempty"`

	/* Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses". */
	OcspSigning bool `json:"ocspSigning,omitempty"`

	/* Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS. */
	ServerAuth bool `json:"serverAuth,omitempty"`

	/* Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time". */
	TimeStamping bool `json:"timeStamping,omitempty"`
}

type CertificateKeyUsageStatus struct {
	/* Describes high-level ways in which a key may be used. */
	BaseKeyUsage CertificateBaseKeyUsageStatus `json:"baseKeyUsage,omitempty"`

	/* Detailed scenarios in which a key may be used. */
	ExtendedKeyUsage CertificateExtendedKeyUsageStatus `json:"extendedKeyUsage,omitempty"`

	/* Used to describe extended key usages that are not listed in the KeyUsage.ExtendedKeyUsageOptions message. */
	UnknownExtendedKeyUsages []CertificateUnknownExtendedKeyUsagesStatus `json:"unknownExtendedKeyUsages,omitempty"`
}

type CertificateObjectIdStatus struct {
	/* Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath,omitempty"`
}

type CertificatePolicyIdsStatus struct {
	/* Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath,omitempty"`
}

type CertificatePublicKeyStatus struct {
	/* Required. The format of the public key. Possible values: KEY_FORMAT_UNSPECIFIED, PEM */
	Format string `json:"format,omitempty"`

	/* Required. A public key. The padding and encoding must match with the `KeyFormat` value specified for the `format` field. */
	Key string `json:"key,omitempty"`
}

type CertificateRevocationDetailsStatus struct {
	/* Indicates why a Certificate was revoked. Possible values: REVOCATION_REASON_UNSPECIFIED, KEY_COMPROMISE, CERTIFICATE_AUTHORITY_COMPROMISE, AFFILIATION_CHANGED, SUPERSEDED, CESSATION_OF_OPERATION, CERTIFICATE_HOLD, PRIVILEGE_WITHDRAWN, ATTRIBUTE_AUTHORITY_COMPROMISE */
	RevocationState string `json:"revocationState,omitempty"`

	/* The time at which this Certificate was revoked. */
	RevocationTime string `json:"revocationTime,omitempty"`
}

type CertificateSubjectAltNameStatus struct {
	/* Contains additional subject alternative name values. */
	CustomSans []CertificateCustomSansStatus `json:"customSans,omitempty"`

	/* Contains only valid, fully-qualified host names. */
	DnsNames []string `json:"dnsNames,omitempty"`

	/* Contains only valid RFC 2822 E-mail addresses. */
	EmailAddresses []string `json:"emailAddresses,omitempty"`

	/* Contains only valid 32-bit IPv4 addresses or RFC 4291 IPv6 addresses. */
	IpAddresses []string `json:"ipAddresses,omitempty"`

	/* Contains only valid RFC 3986 URIs. */
	Uris []string `json:"uris,omitempty"`
}

type CertificateSubjectDescriptionStatus struct {
	/* The serial number encoded in lowercase hexadecimal. */
	HexSerialNumber string `json:"hexSerialNumber,omitempty"`

	/* For convenience, the actual lifetime of an issued certificate. */
	Lifetime string `json:"lifetime,omitempty"`

	/* The time after which the certificate is expired. Per RFC 5280, the validity period for a certificate is the period of time from not_before_time through not_after_time, inclusive. Corresponds to 'not_before_time' + 'lifetime' - 1 second. */
	NotAfterTime string `json:"notAfterTime,omitempty"`

	/* The time at which the certificate becomes valid. */
	NotBeforeTime string `json:"notBeforeTime,omitempty"`

	/* Contains distinguished name fields such as the common name, location and / organization. */
	Subject CertificateSubjectStatus `json:"subject,omitempty"`

	/* The subject alternative name fields. */
	SubjectAltName CertificateSubjectAltNameStatus `json:"subjectAltName,omitempty"`
}

type CertificateSubjectKeyIdStatus struct {
	/* Optional. The value of this KeyId encoded in lowercase hexadecimal. This is most likely the 160 bit SHA-1 hash of the public key. */
	KeyId string `json:"keyId,omitempty"`
}

type CertificateSubjectStatus struct {
	/* The "common name" of the subject. */
	CommonName string `json:"commonName,omitempty"`

	/* The country code of the subject. */
	CountryCode string `json:"countryCode,omitempty"`

	/* The locality or city of the subject. */
	Locality string `json:"locality,omitempty"`

	/* The organization of the subject. */
	Organization string `json:"organization,omitempty"`

	/* The organizational_unit of the subject. */
	OrganizationalUnit string `json:"organizationalUnit,omitempty"`

	/* The postal code of the subject. */
	PostalCode string `json:"postalCode,omitempty"`

	/* The province, territory, or regional state of the subject. */
	Province string `json:"province,omitempty"`

	/* The street address of the subject. */
	StreetAddress string `json:"streetAddress,omitempty"`
}

type CertificateUnknownExtendedKeyUsagesStatus struct {
	/* Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath,omitempty"`
}

type CertificateX509DescriptionStatus struct {
	/* Optional. Describes custom X.509 extensions. */
	AdditionalExtensions []CertificateAdditionalExtensionsStatus `json:"additionalExtensions,omitempty"`

	/* Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate. */
	AiaOcspServers []string `json:"aiaOcspServers,omitempty"`

	/* Optional. Describes options in this X509Parameters that are relevant in a CA certificate. */
	CaOptions CertificateCaOptionsStatus `json:"caOptions,omitempty"`

	/* Optional. Indicates the intended use for keys that correspond to a certificate. */
	KeyUsage CertificateKeyUsageStatus `json:"keyUsage,omitempty"`

	/* Optional. Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4. */
	PolicyIds []CertificatePolicyIdsStatus `json:"policyIds,omitempty"`
}

type PrivateCACertificateStatus struct {
	/* Conditions represent the latest available observations of the
	   PrivateCACertificate's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. A structured description of the issued X.509 certificate. */
	CertificateDescription CertificateCertificateDescriptionStatus `json:"certificateDescription,omitempty"`
	/* Output only. The time at which this Certificate was created. */
	CreateTime string `json:"createTime,omitempty"`
	/* Output only. The resource name of the issuing CertificateAuthority in the format `projects/* /locations/* /caPools/* /certificateAuthorities/*`. */
	IssuerCertificateAuthority string `json:"issuerCertificateAuthority,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* Output only. The pem-encoded, signed X.509 certificate. */
	PemCertificate string `json:"pemCertificate,omitempty"`
	/* Output only. The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246. */
	PemCertificateChain []string `json:"pemCertificateChain,omitempty"`
	/* Output only. Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present. */
	RevocationDetails CertificateRevocationDetailsStatus `json:"revocationDetails,omitempty"`
	/* Output only. The time at which this Certificate was updated. */
	UpdateTime string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PrivateCACertificate is the Schema for the privateca API
// +k8s:openapi-gen=true
type PrivateCACertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrivateCACertificateSpec   `json:"spec,omitempty"`
	Status PrivateCACertificateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PrivateCACertificateList contains a list of PrivateCACertificate
type PrivateCACertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateCACertificate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrivateCACertificate{}, &PrivateCACertificateList{})
}