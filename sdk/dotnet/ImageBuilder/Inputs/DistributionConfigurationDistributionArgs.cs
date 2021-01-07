// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ImageBuilder.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html
    /// </summary>
    public sealed class DistributionConfigurationDistributionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-amidistributionconfiguration
        /// </summary>
        [Input("AmiDistributionConfiguration")]
        public InputUnion<System.Text.Json.JsonElement, string>? AmiDistributionConfiguration { get; set; }

        [Input("LicenseConfigurationArns")]
        private InputList<string>? _LicenseConfigurationArns;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-licenseconfigurationarns
        /// </summary>
        public InputList<string> LicenseConfigurationArns
        {
            get => _LicenseConfigurationArns ?? (_LicenseConfigurationArns = new InputList<string>());
            set => _LicenseConfigurationArns = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-region
        /// </summary>
        [Input("Region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public DistributionConfigurationDistributionArgs()
        {
        }
    }
}