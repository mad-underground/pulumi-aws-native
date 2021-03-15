# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['OriginRequestPolicy']


class OriginRequestPolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 origin_request_policy_config: Optional[pulumi.Input[pulumi.InputType['OriginRequestPolicyOriginRequestPolicyConfigArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originrequestpolicy.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['OriginRequestPolicyOriginRequestPolicyConfigArgs']] origin_request_policy_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originrequestpolicy.html#cfn-cloudfront-originrequestpolicy-originrequestpolicyconfig
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

            if origin_request_policy_config is None and not opts.urn:
                raise TypeError("Missing required property 'origin_request_policy_config'")
            __props__['origin_request_policy_config'] = origin_request_policy_config
            __props__['id'] = None
            __props__['last_modified_time'] = None
        super(OriginRequestPolicy, __self__).__init__(
            'aws-native:CloudFront:OriginRequestPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'OriginRequestPolicy':
        """
        Get an existing OriginRequestPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return OriginRequestPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter(name="originRequestPolicyConfig")
    def origin_request_policy_config(self) -> pulumi.Output['outputs.OriginRequestPolicyOriginRequestPolicyConfig']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originrequestpolicy.html#cfn-cloudfront-originrequestpolicy-originrequestpolicyconfig
        """
        return pulumi.get(self, "origin_request_policy_config")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

