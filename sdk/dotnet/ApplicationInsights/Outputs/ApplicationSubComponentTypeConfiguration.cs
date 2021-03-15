// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationInsights.Outputs
{

    [OutputType]
    public sealed class ApplicationSubComponentTypeConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-subcomponenttypeconfiguration.html#cfn-applicationinsights-application-subcomponenttypeconfiguration-subcomponentconfigurationdetails
        /// </summary>
        public readonly Outputs.ApplicationSubComponentConfigurationDetails SubComponentConfigurationDetails;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-subcomponenttypeconfiguration.html#cfn-applicationinsights-application-subcomponenttypeconfiguration-subcomponenttype
        /// </summary>
        public readonly string SubComponentType;

        [OutputConstructor]
        private ApplicationSubComponentTypeConfiguration(
            Outputs.ApplicationSubComponentConfigurationDetails subComponentConfigurationDetails,

            string subComponentType)
        {
            SubComponentConfigurationDetails = subComponentConfigurationDetails;
            SubComponentType = subComponentType;
        }
    }
}
