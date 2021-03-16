# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._inputs import *

__all__ = ['TaskDefinition']


class TaskDefinition(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_definitions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionContainerDefinitionArgs']]]]] = None,
                 cpu: Optional[pulumi.Input[str]] = None,
                 execution_role_arn: Optional[pulumi.Input[str]] = None,
                 family: Optional[pulumi.Input[str]] = None,
                 inference_accelerators: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionInferenceAcceleratorArgs']]]]] = None,
                 ipc_mode: Optional[pulumi.Input[str]] = None,
                 memory: Optional[pulumi.Input[str]] = None,
                 network_mode: Optional[pulumi.Input[str]] = None,
                 pid_mode: Optional[pulumi.Input[str]] = None,
                 placement_constraints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionTaskDefinitionPlacementConstraintArgs']]]]] = None,
                 proxy_configuration: Optional[pulumi.Input[pulumi.InputType['TaskDefinitionProxyConfigurationArgs']]] = None,
                 requires_compatibilities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 task_role_arn: Optional[pulumi.Input[str]] = None,
                 volumes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionVolumeArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionContainerDefinitionArgs']]]] container_definitions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-containerdefinitions
        :param pulumi.Input[str] cpu: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-cpu
        :param pulumi.Input[str] execution_role_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-executionrolearn
        :param pulumi.Input[str] family: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-family
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionInferenceAcceleratorArgs']]]] inference_accelerators: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-inferenceaccelerators
        :param pulumi.Input[str] ipc_mode: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-ipcmode
        :param pulumi.Input[str] memory: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-memory
        :param pulumi.Input[str] network_mode: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-networkmode
        :param pulumi.Input[str] pid_mode: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-pidmode
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionTaskDefinitionPlacementConstraintArgs']]]] placement_constraints: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-placementconstraints
        :param pulumi.Input[pulumi.InputType['TaskDefinitionProxyConfigurationArgs']] proxy_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-proxyconfiguration
        :param pulumi.Input[Sequence[pulumi.Input[str]]] requires_compatibilities: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-requirescompatibilities
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-tags
        :param pulumi.Input[str] task_role_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-taskrolearn
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionVolumeArgs']]]] volumes: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-volumes
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['container_definitions'] = container_definitions
            __props__['cpu'] = cpu
            __props__['execution_role_arn'] = execution_role_arn
            __props__['family'] = family
            __props__['inference_accelerators'] = inference_accelerators
            __props__['ipc_mode'] = ipc_mode
            __props__['memory'] = memory
            __props__['network_mode'] = network_mode
            __props__['pid_mode'] = pid_mode
            __props__['placement_constraints'] = placement_constraints
            __props__['proxy_configuration'] = proxy_configuration
            __props__['requires_compatibilities'] = requires_compatibilities
            __props__['tags'] = tags
            __props__['task_role_arn'] = task_role_arn
            __props__['volumes'] = volumes
            __props__['task_definition_arn'] = None
        super(TaskDefinition, __self__).__init__(
            'aws-native:ECS:TaskDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'TaskDefinition':
        """
        Get an existing TaskDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return TaskDefinition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="containerDefinitions")
    def container_definitions(self) -> pulumi.Output[Optional[Sequence['outputs.TaskDefinitionContainerDefinition']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-containerdefinitions
        """
        return pulumi.get(self, "container_definitions")

    @property
    @pulumi.getter
    def cpu(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-cpu
        """
        return pulumi.get(self, "cpu")

    @property
    @pulumi.getter(name="executionRoleArn")
    def execution_role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-executionrolearn
        """
        return pulumi.get(self, "execution_role_arn")

    @property
    @pulumi.getter
    def family(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-family
        """
        return pulumi.get(self, "family")

    @property
    @pulumi.getter(name="inferenceAccelerators")
    def inference_accelerators(self) -> pulumi.Output[Optional[Sequence['outputs.TaskDefinitionInferenceAccelerator']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-inferenceaccelerators
        """
        return pulumi.get(self, "inference_accelerators")

    @property
    @pulumi.getter(name="ipcMode")
    def ipc_mode(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-ipcmode
        """
        return pulumi.get(self, "ipc_mode")

    @property
    @pulumi.getter
    def memory(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-memory
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter(name="networkMode")
    def network_mode(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-networkmode
        """
        return pulumi.get(self, "network_mode")

    @property
    @pulumi.getter(name="pidMode")
    def pid_mode(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-pidmode
        """
        return pulumi.get(self, "pid_mode")

    @property
    @pulumi.getter(name="placementConstraints")
    def placement_constraints(self) -> pulumi.Output[Optional[Sequence['outputs.TaskDefinitionTaskDefinitionPlacementConstraint']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-placementconstraints
        """
        return pulumi.get(self, "placement_constraints")

    @property
    @pulumi.getter(name="proxyConfiguration")
    def proxy_configuration(self) -> pulumi.Output[Optional['outputs.TaskDefinitionProxyConfiguration']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-proxyconfiguration
        """
        return pulumi.get(self, "proxy_configuration")

    @property
    @pulumi.getter(name="requiresCompatibilities")
    def requires_compatibilities(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-requirescompatibilities
        """
        return pulumi.get(self, "requires_compatibilities")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="taskDefinitionArn")
    def task_definition_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "task_definition_arn")

    @property
    @pulumi.getter(name="taskRoleArn")
    def task_role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-taskrolearn
        """
        return pulumi.get(self, "task_role_arn")

    @property
    @pulumi.getter
    def volumes(self) -> pulumi.Output[Optional[Sequence['outputs.TaskDefinitionVolume']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-volumes
        """
        return pulumi.get(self, "volumes")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
