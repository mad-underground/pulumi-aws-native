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

__all__ = ['DeploymentArgs', 'Deployment']

@pulumi.input_type
class DeploymentArgs:
    def __init__(__self__, *,
                 target_arn: pulumi.Input[str],
                 components: Optional[Any] = None,
                 deployment_name: Optional[pulumi.Input[str]] = None,
                 deployment_policies: Optional[pulumi.Input['DeploymentPoliciesArgs']] = None,
                 iot_job_configuration: Optional[pulumi.Input['DeploymentIoTJobConfigurationArgs']] = None,
                 tags: Optional[Any] = None):
        """
        The set of arguments for constructing a Deployment resource.
        """
        pulumi.set(__self__, "target_arn", target_arn)
        if components is not None:
            pulumi.set(__self__, "components", components)
        if deployment_name is not None:
            pulumi.set(__self__, "deployment_name", deployment_name)
        if deployment_policies is not None:
            pulumi.set(__self__, "deployment_policies", deployment_policies)
        if iot_job_configuration is not None:
            pulumi.set(__self__, "iot_job_configuration", iot_job_configuration)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="targetArn")
    def target_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "target_arn")

    @target_arn.setter
    def target_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_arn", value)

    @property
    @pulumi.getter
    def components(self) -> Optional[Any]:
        return pulumi.get(self, "components")

    @components.setter
    def components(self, value: Optional[Any]):
        pulumi.set(self, "components", value)

    @property
    @pulumi.getter(name="deploymentName")
    def deployment_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "deployment_name")

    @deployment_name.setter
    def deployment_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deployment_name", value)

    @property
    @pulumi.getter(name="deploymentPolicies")
    def deployment_policies(self) -> Optional[pulumi.Input['DeploymentPoliciesArgs']]:
        return pulumi.get(self, "deployment_policies")

    @deployment_policies.setter
    def deployment_policies(self, value: Optional[pulumi.Input['DeploymentPoliciesArgs']]):
        pulumi.set(self, "deployment_policies", value)

    @property
    @pulumi.getter(name="iotJobConfiguration")
    def iot_job_configuration(self) -> Optional[pulumi.Input['DeploymentIoTJobConfigurationArgs']]:
        return pulumi.get(self, "iot_job_configuration")

    @iot_job_configuration.setter
    def iot_job_configuration(self, value: Optional[pulumi.Input['DeploymentIoTJobConfigurationArgs']]):
        pulumi.set(self, "iot_job_configuration", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[Any]):
        pulumi.set(self, "tags", value)


class Deployment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 components: Optional[Any] = None,
                 deployment_name: Optional[pulumi.Input[str]] = None,
                 deployment_policies: Optional[pulumi.Input[pulumi.InputType['DeploymentPoliciesArgs']]] = None,
                 iot_job_configuration: Optional[pulumi.Input[pulumi.InputType['DeploymentIoTJobConfigurationArgs']]] = None,
                 tags: Optional[Any] = None,
                 target_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource for Greengrass V2 deployment.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeploymentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for Greengrass V2 deployment.

        :param str resource_name: The name of the resource.
        :param DeploymentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeploymentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 components: Optional[Any] = None,
                 deployment_name: Optional[pulumi.Input[str]] = None,
                 deployment_policies: Optional[pulumi.Input[pulumi.InputType['DeploymentPoliciesArgs']]] = None,
                 iot_job_configuration: Optional[pulumi.Input[pulumi.InputType['DeploymentIoTJobConfigurationArgs']]] = None,
                 tags: Optional[Any] = None,
                 target_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeploymentArgs.__new__(DeploymentArgs)

            __props__.__dict__["components"] = components
            __props__.__dict__["deployment_name"] = deployment_name
            __props__.__dict__["deployment_policies"] = deployment_policies
            __props__.__dict__["iot_job_configuration"] = iot_job_configuration
            __props__.__dict__["tags"] = tags
            if target_arn is None and not opts.urn:
                raise TypeError("Missing required property 'target_arn'")
            __props__.__dict__["target_arn"] = target_arn
            __props__.__dict__["deployment_id"] = None
        super(Deployment, __self__).__init__(
            'aws-native:greengrassv2:Deployment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Deployment':
        """
        Get an existing Deployment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DeploymentArgs.__new__(DeploymentArgs)

        __props__.__dict__["components"] = None
        __props__.__dict__["deployment_id"] = None
        __props__.__dict__["deployment_name"] = None
        __props__.__dict__["deployment_policies"] = None
        __props__.__dict__["iot_job_configuration"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["target_arn"] = None
        return Deployment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def components(self) -> pulumi.Output[Optional[Any]]:
        return pulumi.get(self, "components")

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "deployment_id")

    @property
    @pulumi.getter(name="deploymentName")
    def deployment_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "deployment_name")

    @property
    @pulumi.getter(name="deploymentPolicies")
    def deployment_policies(self) -> pulumi.Output[Optional['outputs.DeploymentPolicies']]:
        return pulumi.get(self, "deployment_policies")

    @property
    @pulumi.getter(name="iotJobConfiguration")
    def iot_job_configuration(self) -> pulumi.Output[Optional['outputs.DeploymentIoTJobConfiguration']]:
        return pulumi.get(self, "iot_job_configuration")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Any]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetArn")
    def target_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "target_arn")

