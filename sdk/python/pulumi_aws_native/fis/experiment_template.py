# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['ExperimentTemplateArgs', 'ExperimentTemplate']

@pulumi.input_type
class ExperimentTemplateArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 role_arn: pulumi.Input[str],
                 stop_conditions: pulumi.Input[Sequence[pulumi.Input['ExperimentTemplateStopConditionArgs']]],
                 tags: pulumi.Input[Mapping[str, pulumi.Input[str]]],
                 targets: pulumi.Input[Mapping[str, pulumi.Input['ExperimentTemplateTargetArgs']]],
                 actions: Optional[pulumi.Input[Mapping[str, pulumi.Input['ExperimentTemplateActionArgs']]]] = None,
                 experiment_options: Optional[pulumi.Input['ExperimentTemplateExperimentOptionsArgs']] = None,
                 log_configuration: Optional[pulumi.Input['ExperimentTemplateLogConfigurationArgs']] = None):
        """
        The set of arguments for constructing a ExperimentTemplate resource.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "role_arn", role_arn)
        pulumi.set(__self__, "stop_conditions", stop_conditions)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "targets", targets)
        if actions is not None:
            pulumi.set(__self__, "actions", actions)
        if experiment_options is not None:
            pulumi.set(__self__, "experiment_options", experiment_options)
        if log_configuration is not None:
            pulumi.set(__self__, "log_configuration", log_configuration)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="stopConditions")
    def stop_conditions(self) -> pulumi.Input[Sequence[pulumi.Input['ExperimentTemplateStopConditionArgs']]]:
        return pulumi.get(self, "stop_conditions")

    @stop_conditions.setter
    def stop_conditions(self, value: pulumi.Input[Sequence[pulumi.Input['ExperimentTemplateStopConditionArgs']]]):
        pulumi.set(self, "stop_conditions", value)

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Input[Mapping[str, pulumi.Input['ExperimentTemplateTargetArgs']]]:
        return pulumi.get(self, "targets")

    @targets.setter
    def targets(self, value: pulumi.Input[Mapping[str, pulumi.Input['ExperimentTemplateTargetArgs']]]):
        pulumi.set(self, "targets", value)

    @property
    @pulumi.getter
    def actions(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['ExperimentTemplateActionArgs']]]]:
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['ExperimentTemplateActionArgs']]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter(name="experimentOptions")
    def experiment_options(self) -> Optional[pulumi.Input['ExperimentTemplateExperimentOptionsArgs']]:
        return pulumi.get(self, "experiment_options")

    @experiment_options.setter
    def experiment_options(self, value: Optional[pulumi.Input['ExperimentTemplateExperimentOptionsArgs']]):
        pulumi.set(self, "experiment_options", value)

    @property
    @pulumi.getter(name="logConfiguration")
    def log_configuration(self) -> Optional[pulumi.Input['ExperimentTemplateLogConfigurationArgs']]:
        return pulumi.get(self, "log_configuration")

    @log_configuration.setter
    def log_configuration(self, value: Optional[pulumi.Input['ExperimentTemplateLogConfigurationArgs']]):
        pulumi.set(self, "log_configuration", value)


class ExperimentTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['ExperimentTemplateActionArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 experiment_options: Optional[pulumi.Input[pulumi.InputType['ExperimentTemplateExperimentOptionsArgs']]] = None,
                 log_configuration: Optional[pulumi.Input[pulumi.InputType['ExperimentTemplateLogConfigurationArgs']]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stop_conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExperimentTemplateStopConditionArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 targets: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['ExperimentTemplateTargetArgs']]]]] = None,
                 __props__=None):
        """
        Resource schema for AWS::FIS::ExperimentTemplate

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ExperimentTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::FIS::ExperimentTemplate

        :param str resource_name: The name of the resource.
        :param ExperimentTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExperimentTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['ExperimentTemplateActionArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 experiment_options: Optional[pulumi.Input[pulumi.InputType['ExperimentTemplateExperimentOptionsArgs']]] = None,
                 log_configuration: Optional[pulumi.Input[pulumi.InputType['ExperimentTemplateLogConfigurationArgs']]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stop_conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExperimentTemplateStopConditionArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 targets: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['ExperimentTemplateTargetArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ExperimentTemplateArgs.__new__(ExperimentTemplateArgs)

            __props__.__dict__["actions"] = actions
            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            __props__.__dict__["experiment_options"] = experiment_options
            __props__.__dict__["log_configuration"] = log_configuration
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            if stop_conditions is None and not opts.urn:
                raise TypeError("Missing required property 'stop_conditions'")
            __props__.__dict__["stop_conditions"] = stop_conditions
            if tags is None and not opts.urn:
                raise TypeError("Missing required property 'tags'")
            __props__.__dict__["tags"] = tags
            if targets is None and not opts.urn:
                raise TypeError("Missing required property 'targets'")
            __props__.__dict__["targets"] = targets
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["tags.*"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ExperimentTemplate, __self__).__init__(
            'aws-native:fis:ExperimentTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ExperimentTemplate':
        """
        Get an existing ExperimentTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ExperimentTemplateArgs.__new__(ExperimentTemplateArgs)

        __props__.__dict__["actions"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["experiment_options"] = None
        __props__.__dict__["log_configuration"] = None
        __props__.__dict__["role_arn"] = None
        __props__.__dict__["stop_conditions"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["targets"] = None
        return ExperimentTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Output[Optional[Mapping[str, 'outputs.ExperimentTemplateAction']]]:
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="experimentOptions")
    def experiment_options(self) -> pulumi.Output[Optional['outputs.ExperimentTemplateExperimentOptions']]:
        return pulumi.get(self, "experiment_options")

    @property
    @pulumi.getter(name="logConfiguration")
    def log_configuration(self) -> pulumi.Output[Optional['outputs.ExperimentTemplateLogConfiguration']]:
        return pulumi.get(self, "log_configuration")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="stopConditions")
    def stop_conditions(self) -> pulumi.Output[Sequence['outputs.ExperimentTemplateStopCondition']]:
        return pulumi.get(self, "stop_conditions")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Output[Mapping[str, 'outputs.ExperimentTemplateTarget']]:
        return pulumi.get(self, "targets")

