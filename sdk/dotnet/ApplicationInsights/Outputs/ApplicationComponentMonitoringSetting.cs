// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationInsights.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentmonitoringsetting.html
    /// </summary>
    [OutputType]
    public sealed class ApplicationComponentMonitoringSetting
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentmonitoringsetting.html#cfn-applicationinsights-application-componentmonitoringsetting-componentarn
        /// </summary>
        public readonly string? ComponentARN;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentmonitoringsetting.html#cfn-applicationinsights-application-componentmonitoringsetting-componentconfigurationmode
        /// </summary>
        public readonly string? ComponentConfigurationMode;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentmonitoringsetting.html#cfn-applicationinsights-application-componentmonitoringsetting-componentname
        /// </summary>
        public readonly string? ComponentName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentmonitoringsetting.html#cfn-applicationinsights-application-componentmonitoringsetting-customcomponentconfiguration
        /// </summary>
        public readonly Outputs.ApplicationComponentConfiguration? CustomComponentConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentmonitoringsetting.html#cfn-applicationinsights-application-componentmonitoringsetting-defaultoverwritecomponentconfiguration
        /// </summary>
        public readonly Outputs.ApplicationComponentConfiguration? DefaultOverwriteComponentConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentmonitoringsetting.html#cfn-applicationinsights-application-componentmonitoringsetting-tier
        /// </summary>
        public readonly string? Tier;

        [OutputConstructor]
        private ApplicationComponentMonitoringSetting(
            string? componentARN,

            string? componentConfigurationMode,

            string? componentName,

            Outputs.ApplicationComponentConfiguration? customComponentConfiguration,

            Outputs.ApplicationComponentConfiguration? defaultOverwriteComponentConfiguration,

            string? tier)
        {
            ComponentARN = componentARN;
            ComponentConfigurationMode = componentConfigurationMode;
            ComponentName = componentName;
            CustomComponentConfiguration = customComponentConfiguration;
            DefaultOverwriteComponentConfiguration = defaultOverwriteComponentConfiguration;
            Tier = tier;
        }
    }
}
