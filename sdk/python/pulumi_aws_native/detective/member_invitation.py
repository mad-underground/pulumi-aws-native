# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['MemberInvitation']


class MemberInvitation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 graph_arn: Optional[pulumi.Input[str]] = None,
                 member_email_address: Optional[pulumi.Input[str]] = None,
                 member_id: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] graph_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-grapharn
        :param pulumi.Input[str] member_email_address: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-memberemailaddress
        :param pulumi.Input[str] member_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-memberid
        :param pulumi.Input[str] message: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-message
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

            if graph_arn is None and not opts.urn:
                raise TypeError("Missing required property 'graph_arn'")
            __props__['graph_arn'] = graph_arn
            if member_email_address is None and not opts.urn:
                raise TypeError("Missing required property 'member_email_address'")
            __props__['member_email_address'] = member_email_address
            if member_id is None and not opts.urn:
                raise TypeError("Missing required property 'member_id'")
            __props__['member_id'] = member_id
            __props__['message'] = message
        super(MemberInvitation, __self__).__init__(
            'aws-native:Detective:MemberInvitation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'MemberInvitation':
        """
        Get an existing MemberInvitation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return MemberInvitation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="graphArn")
    def graph_arn(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-grapharn
        """
        return pulumi.get(self, "graph_arn")

    @property
    @pulumi.getter(name="memberEmailAddress")
    def member_email_address(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-memberemailaddress
        """
        return pulumi.get(self, "member_email_address")

    @property
    @pulumi.getter(name="memberId")
    def member_id(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-memberid
        """
        return pulumi.get(self, "member_id")

    @property
    @pulumi.getter
    def message(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-message
        """
        return pulumi.get(self, "message")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
