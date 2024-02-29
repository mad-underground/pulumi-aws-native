// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Acmpca.Inputs
{

    /// <summary>
    /// Contains information about the certificate subject. The ``Subject`` field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The ``Subject``must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate.
    /// </summary>
    public sealed class CertificateSubjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// For CA and end-entity certificates in a private PKI, the common name (CN) can be any string within the length limit.
        ///  Note: In publicly trusted certificates, the common name must be a fully qualified domain name (FQDN) associated with the certificate subject.
        /// </summary>
        [Input("commonName")]
        public Input<string>? CommonName { get; set; }

        /// <summary>
        /// Two-digit code that specifies the country in which the certificate subject located.
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        [Input("customAttributes")]
        private InputList<Inputs.CertificateCustomAttributeArgs>? _customAttributes;

        /// <summary>
        /// Contains a sequence of one or more X.500 relative distinguished names (RDNs), each of which consists of an object identifier (OID) and a value. For more information, see NIST’s definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier).
        ///   Custom attributes cannot be used in combination with standard attributes.
        /// </summary>
        public InputList<Inputs.CertificateCustomAttributeArgs> CustomAttributes
        {
            get => _customAttributes ?? (_customAttributes = new InputList<Inputs.CertificateCustomAttributeArgs>());
            set => _customAttributes = value;
        }

        /// <summary>
        /// Disambiguating information for the certificate subject.
        /// </summary>
        [Input("distinguishedNameQualifier")]
        public Input<string>? DistinguishedNameQualifier { get; set; }

        /// <summary>
        /// Typically a qualifier appended to the name of an individual. Examples include Jr. for junior, Sr. for senior, and III for third.
        /// </summary>
        [Input("generationQualifier")]
        public Input<string>? GenerationQualifier { get; set; }

        /// <summary>
        /// First name.
        /// </summary>
        [Input("givenName")]
        public Input<string>? GivenName { get; set; }

        /// <summary>
        /// Concatenation that typically contains the first letter of the *GivenName*, the first letter of the middle name if one exists, and the first letter of the *Surname*.
        /// </summary>
        [Input("initials")]
        public Input<string>? Initials { get; set; }

        /// <summary>
        /// The locality (such as a city or town) in which the certificate subject is located.
        /// </summary>
        [Input("locality")]
        public Input<string>? Locality { get; set; }

        /// <summary>
        /// Legal name of the organization with which the certificate subject is affiliated.
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.
        /// </summary>
        [Input("organizationalUnit")]
        public Input<string>? OrganizationalUnit { get; set; }

        /// <summary>
        /// Typically a shortened version of a longer *GivenName*. For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
        /// </summary>
        [Input("pseudonym")]
        public Input<string>? Pseudonym { get; set; }

        /// <summary>
        /// The certificate serial number.
        /// </summary>
        [Input("serialNumber")]
        public Input<string>? SerialNumber { get; set; }

        /// <summary>
        /// State in which the subject of the certificate is located.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Family name. In the US and the UK, for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first.
        /// </summary>
        [Input("surname")]
        public Input<string>? Surname { get; set; }

        /// <summary>
        /// A title such as Mr. or Ms., which is pre-pended to the name to refer formally to the certificate subject.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public CertificateSubjectArgs()
        {
        }
        public static new CertificateSubjectArgs Empty => new CertificateSubjectArgs();
    }
}
