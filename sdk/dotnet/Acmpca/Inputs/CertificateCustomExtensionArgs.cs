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
    /// Specifies the X.509 extension information for a certificate.
    ///  Extensions present in ``CustomExtensions`` follow the ``ApiPassthrough`` [template rules](https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html#template-order-of-operations).
    /// </summary>
    public sealed class CertificateCustomExtensionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the critical flag of the X.509 extension.
        /// </summary>
        [Input("critical")]
        public Input<bool>? Critical { get; set; }

        /// <summary>
        /// Specifies the object identifier (OID) of the X.509 extension. For more information, see the [Global OID reference database.](https://docs.aws.amazon.com/https://oidref.com/2.5.29)
        /// </summary>
        [Input("objectIdentifier", required: true)]
        public Input<string> ObjectIdentifier { get; set; } = null!;

        /// <summary>
        /// Specifies the base64-encoded value of the X.509 extension.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public CertificateCustomExtensionArgs()
        {
        }
        public static new CertificateCustomExtensionArgs Empty => new CertificateCustomExtensionArgs();
    }
}
