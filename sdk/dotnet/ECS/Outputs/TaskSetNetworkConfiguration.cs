// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ECS.Outputs
{

    [OutputType]
    public sealed class TaskSetNetworkConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskset-networkconfiguration.html#cfn-ecs-taskset-networkconfiguration-awsvpcconfiguration
        /// </summary>
        public readonly Outputs.TaskSetAwsVpcConfiguration? AwsVpcConfiguration;

        [OutputConstructor]
        private TaskSetNetworkConfiguration(Outputs.TaskSetAwsVpcConfiguration? AwsVpcConfiguration)
        {
            this.AwsVpcConfiguration = AwsVpcConfiguration;
        }
    }
}