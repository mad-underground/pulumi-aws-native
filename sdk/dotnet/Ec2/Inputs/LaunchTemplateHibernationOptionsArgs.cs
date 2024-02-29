// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    /// <summary>
    /// Specifies whether your instance is configured for hibernation. This parameter is valid only if the instance meets the [hibernation prerequisites](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html#hibernating-prerequisites). For more information, see [Hibernate Your Instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the *Amazon EC2 User Guide*.
    ///   ``HibernationOptions`` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html).
    /// </summary>
    public sealed class LaunchTemplateHibernationOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If you set this parameter to ``true``, the instance is enabled for hibernation.
        ///  Default: ``false``
        /// </summary>
        [Input("configured")]
        public Input<bool>? Configured { get; set; }

        public LaunchTemplateHibernationOptionsArgs()
        {
        }
        public static new LaunchTemplateHibernationOptionsArgs Empty => new LaunchTemplateHibernationOptionsArgs();
    }
}
