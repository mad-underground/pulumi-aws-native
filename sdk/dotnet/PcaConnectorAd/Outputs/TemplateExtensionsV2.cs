// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.PcaConnectorAd.Outputs
{

    [OutputType]
    public sealed class TemplateExtensionsV2
    {
        public readonly Outputs.TemplateApplicationPolicies? ApplicationPolicies;
        public readonly Outputs.TemplateKeyUsage KeyUsage;

        [OutputConstructor]
        private TemplateExtensionsV2(
            Outputs.TemplateApplicationPolicies? applicationPolicies,

            Outputs.TemplateKeyUsage keyUsage)
        {
            ApplicationPolicies = applicationPolicies;
            KeyUsage = keyUsage;
        }
    }
}