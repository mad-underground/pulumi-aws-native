// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class InstanceLicenseSpecificationArgs : global::Pulumi.ResourceArgs
    {
        [Input("licenseConfigurationArn", required: true)]
        public Input<string> LicenseConfigurationArn { get; set; } = null!;

        public InstanceLicenseSpecificationArgs()
        {
        }
        public static new InstanceLicenseSpecificationArgs Empty => new InstanceLicenseSpecificationArgs();
    }
}
