// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.KinesisAnalyticsV2.Outputs
{

    [OutputType]
    public sealed class ApplicationEnvironmentProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-environmentproperties.html#cfn-kinesisanalyticsv2-application-environmentproperties-propertygroups
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationPropertyGroup> PropertyGroups;

        [OutputConstructor]
        private ApplicationEnvironmentProperties(ImmutableArray<Outputs.ApplicationPropertyGroup> PropertyGroups)
        {
            this.PropertyGroups = PropertyGroups;
        }
    }
}