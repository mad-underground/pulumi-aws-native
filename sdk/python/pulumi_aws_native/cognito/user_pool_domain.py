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

__all__ = ['UserPoolDomainArgs', 'UserPoolDomain']

@pulumi.input_type
class UserPoolDomainArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 user_pool_id: pulumi.Input[str],
                 custom_domain_config: Optional[pulumi.Input['UserPoolDomainCustomDomainConfigTypeArgs']] = None):
        """
        The set of arguments for constructing a UserPoolDomain resource.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "user_pool_id", user_pool_id)
        if custom_domain_config is not None:
            pulumi.set(__self__, "custom_domain_config", custom_domain_config)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="userPoolId")
    def user_pool_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "user_pool_id")

    @user_pool_id.setter
    def user_pool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_pool_id", value)

    @property
    @pulumi.getter(name="customDomainConfig")
    def custom_domain_config(self) -> Optional[pulumi.Input['UserPoolDomainCustomDomainConfigTypeArgs']]:
        return pulumi.get(self, "custom_domain_config")

    @custom_domain_config.setter
    def custom_domain_config(self, value: Optional[pulumi.Input['UserPoolDomainCustomDomainConfigTypeArgs']]):
        pulumi.set(self, "custom_domain_config", value)


class UserPoolDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_domain_config: Optional[pulumi.Input[pulumi.InputType['UserPoolDomainCustomDomainConfigTypeArgs']]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 user_pool_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Cognito::UserPoolDomain

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserPoolDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Cognito::UserPoolDomain

        :param str resource_name: The name of the resource.
        :param UserPoolDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserPoolDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_domain_config: Optional[pulumi.Input[pulumi.InputType['UserPoolDomainCustomDomainConfigTypeArgs']]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 user_pool_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserPoolDomainArgs.__new__(UserPoolDomainArgs)

            __props__.__dict__["custom_domain_config"] = custom_domain_config
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            if user_pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_pool_id'")
            __props__.__dict__["user_pool_id"] = user_pool_id
            __props__.__dict__["aws_id"] = None
            __props__.__dict__["cloud_front_distribution"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["domain", "user_pool_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(UserPoolDomain, __self__).__init__(
            'aws-native:cognito:UserPoolDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'UserPoolDomain':
        """
        Get an existing UserPoolDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = UserPoolDomainArgs.__new__(UserPoolDomainArgs)

        __props__.__dict__["aws_id"] = None
        __props__.__dict__["cloud_front_distribution"] = None
        __props__.__dict__["custom_domain_config"] = None
        __props__.__dict__["domain"] = None
        __props__.__dict__["user_pool_id"] = None
        return UserPoolDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="cloudFrontDistribution")
    def cloud_front_distribution(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cloud_front_distribution")

    @property
    @pulumi.getter(name="customDomainConfig")
    def custom_domain_config(self) -> pulumi.Output[Optional['outputs.UserPoolDomainCustomDomainConfigType']]:
        return pulumi.get(self, "custom_domain_config")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="userPoolId")
    def user_pool_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "user_pool_id")

