// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ECS.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html
    /// </summary>
    public sealed class TaskDefinitionPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("ContainerDefinitions")]
        private InputList<Inputs.TaskDefinitionContainerDefinitionArgs>? _ContainerDefinitions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-containerdefinitions
        /// </summary>
        public InputList<Inputs.TaskDefinitionContainerDefinitionArgs> ContainerDefinitions
        {
            get => _ContainerDefinitions ?? (_ContainerDefinitions = new InputList<Inputs.TaskDefinitionContainerDefinitionArgs>());
            set => _ContainerDefinitions = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-cpu
        /// </summary>
        [Input("Cpu")]
        public Input<string>? Cpu { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-executionrolearn
        /// </summary>
        [Input("ExecutionRoleArn")]
        public Input<string>? ExecutionRoleArn { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-family
        /// </summary>
        [Input("Family")]
        public Input<string>? Family { get; set; }

        [Input("InferenceAccelerators")]
        private InputList<Inputs.TaskDefinitionInferenceAcceleratorArgs>? _InferenceAccelerators;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-inferenceaccelerators
        /// </summary>
        public InputList<Inputs.TaskDefinitionInferenceAcceleratorArgs> InferenceAccelerators
        {
            get => _InferenceAccelerators ?? (_InferenceAccelerators = new InputList<Inputs.TaskDefinitionInferenceAcceleratorArgs>());
            set => _InferenceAccelerators = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-ipcmode
        /// </summary>
        [Input("IpcMode")]
        public Input<string>? IpcMode { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-memory
        /// </summary>
        [Input("Memory")]
        public Input<string>? Memory { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-networkmode
        /// </summary>
        [Input("NetworkMode")]
        public Input<string>? NetworkMode { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-pidmode
        /// </summary>
        [Input("PidMode")]
        public Input<string>? PidMode { get; set; }

        [Input("PlacementConstraints")]
        private InputList<Inputs.TaskDefinitionTaskDefinitionPlacementConstraintArgs>? _PlacementConstraints;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-placementconstraints
        /// </summary>
        public InputList<Inputs.TaskDefinitionTaskDefinitionPlacementConstraintArgs> PlacementConstraints
        {
            get => _PlacementConstraints ?? (_PlacementConstraints = new InputList<Inputs.TaskDefinitionTaskDefinitionPlacementConstraintArgs>());
            set => _PlacementConstraints = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-proxyconfiguration
        /// </summary>
        [Input("ProxyConfiguration")]
        public Input<Inputs.TaskDefinitionProxyConfigurationArgs>? ProxyConfiguration { get; set; }

        [Input("RequiresCompatibilities")]
        private InputList<string>? _RequiresCompatibilities;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-requirescompatibilities
        /// </summary>
        public InputList<string> RequiresCompatibilities
        {
            get => _RequiresCompatibilities ?? (_RequiresCompatibilities = new InputList<string>());
            set => _RequiresCompatibilities = value;
        }

        [Input("Tags")]
        private InputList<Pulumi.Cloudformation.Inputs.TagArgs>? _Tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-tags
        /// </summary>
        public InputList<Pulumi.Cloudformation.Inputs.TagArgs> Tags
        {
            get => _Tags ?? (_Tags = new InputList<Pulumi.Cloudformation.Inputs.TagArgs>());
            set => _Tags = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-taskrolearn
        /// </summary>
        [Input("TaskRoleArn")]
        public Input<string>? TaskRoleArn { get; set; }

        [Input("Volumes")]
        private InputList<Inputs.TaskDefinitionVolumeArgs>? _Volumes;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-volumes
        /// </summary>
        public InputList<Inputs.TaskDefinitionVolumeArgs> Volumes
        {
            get => _Volumes ?? (_Volumes = new InputList<Inputs.TaskDefinitionVolumeArgs>());
            set => _Volumes = value;
        }

        public TaskDefinitionPropertiesArgs()
        {
        }
    }
}