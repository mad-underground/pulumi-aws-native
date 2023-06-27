# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = ['OrganizationArgs', 'Organization']

@pulumi.input_type
class OrganizationArgs:
    def __init__(__self__, *,
                 feature_set: Optional[pulumi.Input['OrganizationFeatureSet']] = None):
        """
        The set of arguments for constructing a Organization resource.
        :param pulumi.Input['OrganizationFeatureSet'] feature_set: Specifies the feature set supported by the new organization. Each feature set supports different levels of functionality.
        """
        if feature_set is not None:
            pulumi.set(__self__, "feature_set", feature_set)

    @property
    @pulumi.getter(name="featureSet")
    def feature_set(self) -> Optional[pulumi.Input['OrganizationFeatureSet']]:
        """
        Specifies the feature set supported by the new organization. Each feature set supports different levels of functionality.
        """
        return pulumi.get(self, "feature_set")

    @feature_set.setter
    def feature_set(self, value: Optional[pulumi.Input['OrganizationFeatureSet']]):
        pulumi.set(self, "feature_set", value)


class Organization(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 feature_set: Optional[pulumi.Input['OrganizationFeatureSet']] = None,
                 __props__=None):
        """
        Resource schema for AWS::Organizations::Organization

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['OrganizationFeatureSet'] feature_set: Specifies the feature set supported by the new organization. Each feature set supports different levels of functionality.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OrganizationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::Organizations::Organization

        :param str resource_name: The name of the resource.
        :param OrganizationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 feature_set: Optional[pulumi.Input['OrganizationFeatureSet']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrganizationArgs.__new__(OrganizationArgs)

            __props__.__dict__["feature_set"] = feature_set
            __props__.__dict__["arn"] = None
            __props__.__dict__["management_account_arn"] = None
            __props__.__dict__["management_account_email"] = None
            __props__.__dict__["management_account_id"] = None
            __props__.__dict__["root_id"] = None
        super(Organization, __self__).__init__(
            'aws-native:organizations:Organization',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Organization':
        """
        Get an existing Organization resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = OrganizationArgs.__new__(OrganizationArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["feature_set"] = None
        __props__.__dict__["management_account_arn"] = None
        __props__.__dict__["management_account_email"] = None
        __props__.__dict__["management_account_id"] = None
        __props__.__dict__["root_id"] = None
        return Organization(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of an organization.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="featureSet")
    def feature_set(self) -> pulumi.Output[Optional['OrganizationFeatureSet']]:
        """
        Specifies the feature set supported by the new organization. Each feature set supports different levels of functionality.
        """
        return pulumi.get(self, "feature_set")

    @property
    @pulumi.getter(name="managementAccountArn")
    def management_account_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the account that is designated as the management account for the organization.
        """
        return pulumi.get(self, "management_account_arn")

    @property
    @pulumi.getter(name="managementAccountEmail")
    def management_account_email(self) -> pulumi.Output[str]:
        """
        The email address that is associated with the AWS account that is designated as the management account for the organization.
        """
        return pulumi.get(self, "management_account_email")

    @property
    @pulumi.getter(name="managementAccountId")
    def management_account_id(self) -> pulumi.Output[str]:
        """
        The unique identifier (ID) of the management account of an organization.
        """
        return pulumi.get(self, "management_account_id")

    @property
    @pulumi.getter(name="rootId")
    def root_id(self) -> pulumi.Output[str]:
        """
        The unique identifier (ID) for the root.
        """
        return pulumi.get(self, "root_id")

