// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SES.Outputs
{

    /// <summary>
    /// Used to enable or disable DKIM authentication for an email identity.
    /// </summary>
    [OutputType]
    public sealed class EmailIdentityDkimAttributes
    {
        /// <summary>
        /// Sets the DKIM signing configuration for the identity. When you set this value true, then the messages that are sent from the identity are signed using DKIM. If you set this value to false, your messages are sent without DKIM signing.
        /// </summary>
        public readonly bool? SigningEnabled;

        [OutputConstructor]
        private EmailIdentityDkimAttributes(bool? signingEnabled)
        {
            SigningEnabled = signingEnabled;
        }
    }
}