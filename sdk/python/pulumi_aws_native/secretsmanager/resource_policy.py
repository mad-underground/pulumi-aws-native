# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ResourcePolicyArgs', 'ResourcePolicy']

@pulumi.input_type
class ResourcePolicyArgs:
    def __init__(__self__, *,
                 resource_policy: Any,
                 secret_id: pulumi.Input[str],
                 block_public_policy: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ResourcePolicy resource.
        :param Any resource_policy: Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SecretsManager::ResourcePolicy` for more information about the expected schema for this property.
        """
        pulumi.set(__self__, "resource_policy", resource_policy)
        pulumi.set(__self__, "secret_id", secret_id)
        if block_public_policy is not None:
            pulumi.set(__self__, "block_public_policy", block_public_policy)

    @property
    @pulumi.getter(name="resourcePolicy")
    def resource_policy(self) -> Any:
        """
        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SecretsManager::ResourcePolicy` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "resource_policy")

    @resource_policy.setter
    def resource_policy(self, value: Any):
        pulumi.set(self, "resource_policy", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_id", value)

    @property
    @pulumi.getter(name="blockPublicPolicy")
    def block_public_policy(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "block_public_policy")

    @block_public_policy.setter
    def block_public_policy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "block_public_policy", value)


warnings.warn("""ResourcePolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class ResourcePolicy(pulumi.CustomResource):
    warnings.warn("""ResourcePolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block_public_policy: Optional[pulumi.Input[bool]] = None,
                 resource_policy: Optional[Any] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::SecretsManager::ResourcePolicy

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param Any resource_policy: Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SecretsManager::ResourcePolicy` for more information about the expected schema for this property.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourcePolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::SecretsManager::ResourcePolicy

        :param str resource_name: The name of the resource.
        :param ResourcePolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourcePolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block_public_policy: Optional[pulumi.Input[bool]] = None,
                 resource_policy: Optional[Any] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""ResourcePolicy is deprecated: ResourcePolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourcePolicyArgs.__new__(ResourcePolicyArgs)

            __props__.__dict__["block_public_policy"] = block_public_policy
            if resource_policy is None and not opts.urn:
                raise TypeError("Missing required property 'resource_policy'")
            __props__.__dict__["resource_policy"] = resource_policy
            if secret_id is None and not opts.urn:
                raise TypeError("Missing required property 'secret_id'")
            __props__.__dict__["secret_id"] = secret_id
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["secret_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ResourcePolicy, __self__).__init__(
            'aws-native:secretsmanager:ResourcePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ResourcePolicy':
        """
        Get an existing ResourcePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ResourcePolicyArgs.__new__(ResourcePolicyArgs)

        __props__.__dict__["aws_id"] = None
        __props__.__dict__["block_public_policy"] = None
        __props__.__dict__["resource_policy"] = None
        __props__.__dict__["secret_id"] = None
        return ResourcePolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="blockPublicPolicy")
    def block_public_policy(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "block_public_policy")

    @property
    @pulumi.getter(name="resourcePolicy")
    def resource_policy(self) -> pulumi.Output[Any]:
        """
        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SecretsManager::ResourcePolicy` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "resource_policy")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "secret_id")

