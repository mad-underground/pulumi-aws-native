// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.SSM.Outputs
{

    [OutputType]
    public sealed class MaintenanceWindowTaskProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-logginginfo
        /// </summary>
        public readonly Outputs.MaintenanceWindowTaskLoggingInfo? LoggingInfo;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-maxconcurrency
        /// </summary>
        public readonly string MaxConcurrency;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-maxerrors
        /// </summary>
        public readonly string MaxErrors;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-priority
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-servicerolearn
        /// </summary>
        public readonly string? ServiceRoleArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-targets
        /// </summary>
        public readonly ImmutableArray<Outputs.MaintenanceWindowTaskTarget> Targets;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-taskarn
        /// </summary>
        public readonly string TaskArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters
        /// </summary>
        public readonly Outputs.MaintenanceWindowTaskTaskInvocationParameters? TaskInvocationParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-taskparameters
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? TaskParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-tasktype
        /// </summary>
        public readonly string TaskType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-windowid
        /// </summary>
        public readonly string WindowId;

        [OutputConstructor]
        private MaintenanceWindowTaskProperties(
            string? Description,

            Outputs.MaintenanceWindowTaskLoggingInfo? LoggingInfo,

            string MaxConcurrency,

            string MaxErrors,

            string? Name,

            int Priority,

            string? ServiceRoleArn,

            ImmutableArray<Outputs.MaintenanceWindowTaskTarget> Targets,

            string TaskArn,

            Outputs.MaintenanceWindowTaskTaskInvocationParameters? TaskInvocationParameters,

            Union<System.Text.Json.JsonElement, string>? TaskParameters,

            string TaskType,

            string WindowId)
        {
            this.Description = Description;
            this.LoggingInfo = LoggingInfo;
            this.MaxConcurrency = MaxConcurrency;
            this.MaxErrors = MaxErrors;
            this.Name = Name;
            this.Priority = Priority;
            this.ServiceRoleArn = ServiceRoleArn;
            this.Targets = Targets;
            this.TaskArn = TaskArn;
            this.TaskInvocationParameters = TaskInvocationParameters;
            this.TaskParameters = TaskParameters;
            this.TaskType = TaskType;
            this.WindowId = WindowId;
        }
    }
}