# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ProfilePermission']


class ProfilePermission(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 profile_name: Optional[pulumi.Input[str]] = None,
                 profile_version: Optional[pulumi.Input[str]] = None,
                 statement_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-action
        :param pulumi.Input[str] principal: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-principal
        :param pulumi.Input[str] profile_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-profilename
        :param pulumi.Input[str] profile_version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-profileversion
        :param pulumi.Input[str] statement_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-statementid
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

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__['action'] = action
            if principal is None and not opts.urn:
                raise TypeError("Missing required property 'principal'")
            __props__['principal'] = principal
            if profile_name is None and not opts.urn:
                raise TypeError("Missing required property 'profile_name'")
            __props__['profile_name'] = profile_name
            __props__['profile_version'] = profile_version
            if statement_id is None and not opts.urn:
                raise TypeError("Missing required property 'statement_id'")
            __props__['statement_id'] = statement_id
        super(ProfilePermission, __self__).__init__(
            'aws-native:Signer:ProfilePermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ProfilePermission':
        """
        Get an existing ProfilePermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ProfilePermission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-action
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def principal(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-principal
        """
        return pulumi.get(self, "principal")

    @property
    @pulumi.getter(name="profileName")
    def profile_name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-profilename
        """
        return pulumi.get(self, "profile_name")

    @property
    @pulumi.getter(name="profileVersion")
    def profile_version(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-profileversion
        """
        return pulumi.get(self, "profile_version")

    @property
    @pulumi.getter(name="statementId")
    def statement_id(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-statementid
        """
        return pulumi.get(self, "statement_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
