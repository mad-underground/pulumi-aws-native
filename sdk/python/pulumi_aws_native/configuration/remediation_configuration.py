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
from ._inputs import *

__all__ = ['RemediationConfigurationArgs', 'RemediationConfiguration']

@pulumi.input_type
class RemediationConfigurationArgs:
    def __init__(__self__, *,
                 config_rule_name: pulumi.Input[str],
                 target_id: pulumi.Input[str],
                 target_type: pulumi.Input[str],
                 automatic: Optional[pulumi.Input[bool]] = None,
                 execution_controls: Optional[pulumi.Input['RemediationConfigurationExecutionControlsArgs']] = None,
                 maximum_automatic_attempts: Optional[pulumi.Input[int]] = None,
                 parameters: Optional[Any] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 retry_attempt_seconds: Optional[pulumi.Input[int]] = None,
                 target_version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RemediationConfiguration resource.
        :param Any parameters: Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Config::RemediationConfiguration` for more information about the expected schema for this property.
        """
        pulumi.set(__self__, "config_rule_name", config_rule_name)
        pulumi.set(__self__, "target_id", target_id)
        pulumi.set(__self__, "target_type", target_type)
        if automatic is not None:
            pulumi.set(__self__, "automatic", automatic)
        if execution_controls is not None:
            pulumi.set(__self__, "execution_controls", execution_controls)
        if maximum_automatic_attempts is not None:
            pulumi.set(__self__, "maximum_automatic_attempts", maximum_automatic_attempts)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if retry_attempt_seconds is not None:
            pulumi.set(__self__, "retry_attempt_seconds", retry_attempt_seconds)
        if target_version is not None:
            pulumi.set(__self__, "target_version", target_version)

    @property
    @pulumi.getter(name="configRuleName")
    def config_rule_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "config_rule_name")

    @config_rule_name.setter
    def config_rule_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "config_rule_name", value)

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "target_id")

    @target_id.setter
    def target_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_id", value)

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "target_type")

    @target_type.setter
    def target_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_type", value)

    @property
    @pulumi.getter
    def automatic(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "automatic")

    @automatic.setter
    def automatic(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "automatic", value)

    @property
    @pulumi.getter(name="executionControls")
    def execution_controls(self) -> Optional[pulumi.Input['RemediationConfigurationExecutionControlsArgs']]:
        return pulumi.get(self, "execution_controls")

    @execution_controls.setter
    def execution_controls(self, value: Optional[pulumi.Input['RemediationConfigurationExecutionControlsArgs']]):
        pulumi.set(self, "execution_controls", value)

    @property
    @pulumi.getter(name="maximumAutomaticAttempts")
    def maximum_automatic_attempts(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "maximum_automatic_attempts")

    @maximum_automatic_attempts.setter
    def maximum_automatic_attempts(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maximum_automatic_attempts", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Any]:
        """
        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Config::RemediationConfiguration` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[Any]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter(name="retryAttemptSeconds")
    def retry_attempt_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "retry_attempt_seconds")

    @retry_attempt_seconds.setter
    def retry_attempt_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retry_attempt_seconds", value)

    @property
    @pulumi.getter(name="targetVersion")
    def target_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "target_version")

    @target_version.setter
    def target_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_version", value)


warnings.warn("""RemediationConfiguration is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class RemediationConfiguration(pulumi.CustomResource):
    warnings.warn("""RemediationConfiguration is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 automatic: Optional[pulumi.Input[bool]] = None,
                 config_rule_name: Optional[pulumi.Input[str]] = None,
                 execution_controls: Optional[pulumi.Input[pulumi.InputType['RemediationConfigurationExecutionControlsArgs']]] = None,
                 maximum_automatic_attempts: Optional[pulumi.Input[int]] = None,
                 parameters: Optional[Any] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 retry_attempt_seconds: Optional[pulumi.Input[int]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 target_type: Optional[pulumi.Input[str]] = None,
                 target_version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Config::RemediationConfiguration

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param Any parameters: Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Config::RemediationConfiguration` for more information about the expected schema for this property.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RemediationConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Config::RemediationConfiguration

        :param str resource_name: The name of the resource.
        :param RemediationConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RemediationConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 automatic: Optional[pulumi.Input[bool]] = None,
                 config_rule_name: Optional[pulumi.Input[str]] = None,
                 execution_controls: Optional[pulumi.Input[pulumi.InputType['RemediationConfigurationExecutionControlsArgs']]] = None,
                 maximum_automatic_attempts: Optional[pulumi.Input[int]] = None,
                 parameters: Optional[Any] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 retry_attempt_seconds: Optional[pulumi.Input[int]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 target_type: Optional[pulumi.Input[str]] = None,
                 target_version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""RemediationConfiguration is deprecated: RemediationConfiguration is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RemediationConfigurationArgs.__new__(RemediationConfigurationArgs)

            __props__.__dict__["automatic"] = automatic
            if config_rule_name is None and not opts.urn:
                raise TypeError("Missing required property 'config_rule_name'")
            __props__.__dict__["config_rule_name"] = config_rule_name
            __props__.__dict__["execution_controls"] = execution_controls
            __props__.__dict__["maximum_automatic_attempts"] = maximum_automatic_attempts
            __props__.__dict__["parameters"] = parameters
            __props__.__dict__["resource_type"] = resource_type
            __props__.__dict__["retry_attempt_seconds"] = retry_attempt_seconds
            if target_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_id'")
            __props__.__dict__["target_id"] = target_id
            if target_type is None and not opts.urn:
                raise TypeError("Missing required property 'target_type'")
            __props__.__dict__["target_type"] = target_type
            __props__.__dict__["target_version"] = target_version
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["config_rule_name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(RemediationConfiguration, __self__).__init__(
            'aws-native:configuration:RemediationConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RemediationConfiguration':
        """
        Get an existing RemediationConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RemediationConfigurationArgs.__new__(RemediationConfigurationArgs)

        __props__.__dict__["automatic"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["config_rule_name"] = None
        __props__.__dict__["execution_controls"] = None
        __props__.__dict__["maximum_automatic_attempts"] = None
        __props__.__dict__["parameters"] = None
        __props__.__dict__["resource_type"] = None
        __props__.__dict__["retry_attempt_seconds"] = None
        __props__.__dict__["target_id"] = None
        __props__.__dict__["target_type"] = None
        __props__.__dict__["target_version"] = None
        return RemediationConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def automatic(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "automatic")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="configRuleName")
    def config_rule_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "config_rule_name")

    @property
    @pulumi.getter(name="executionControls")
    def execution_controls(self) -> pulumi.Output[Optional['outputs.RemediationConfigurationExecutionControls']]:
        return pulumi.get(self, "execution_controls")

    @property
    @pulumi.getter(name="maximumAutomaticAttempts")
    def maximum_automatic_attempts(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "maximum_automatic_attempts")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Any]]:
        """
        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Config::RemediationConfiguration` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="retryAttemptSeconds")
    def retry_attempt_seconds(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "retry_attempt_seconds")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "target_id")

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "target_type")

    @property
    @pulumi.getter(name="targetVersion")
    def target_version(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "target_version")

