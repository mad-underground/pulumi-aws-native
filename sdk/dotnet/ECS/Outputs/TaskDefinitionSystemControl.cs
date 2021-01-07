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
    public sealed class TaskDefinitionSystemControl
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-systemcontrol.html#cfn-ecs-taskdefinition-systemcontrol-namespace
        /// </summary>
        public readonly string? Namespace;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-systemcontrol.html#cfn-ecs-taskdefinition-systemcontrol-value
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private TaskDefinitionSystemControl(
            string? Namespace,

            string? Value)
        {
            this.Namespace = Namespace;
            this.Value = Value;
        }
    }
}
