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

__all__ = ['ReceiptFilterArgs', 'ReceiptFilter']

@pulumi.input_type
class ReceiptFilterArgs:
    def __init__(__self__, *,
                 filter: pulumi.Input['ReceiptFilterFilterArgs']):
        """
        The set of arguments for constructing a ReceiptFilter resource.
        """
        pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Input['ReceiptFilterFilterArgs']:
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: pulumi.Input['ReceiptFilterFilterArgs']):
        pulumi.set(self, "filter", value)


warnings.warn("""ReceiptFilter is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class ReceiptFilter(pulumi.CustomResource):
    warnings.warn("""ReceiptFilter is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 filter: Optional[pulumi.Input[pulumi.InputType['ReceiptFilterFilterArgs']]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::SES::ReceiptFilter

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReceiptFilterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::SES::ReceiptFilter

        :param str resource_name: The name of the resource.
        :param ReceiptFilterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReceiptFilterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 filter: Optional[pulumi.Input[pulumi.InputType['ReceiptFilterFilterArgs']]] = None,
                 __props__=None):
        pulumi.log.warn("""ReceiptFilter is deprecated: ReceiptFilter is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReceiptFilterArgs.__new__(ReceiptFilterArgs)

            if filter is None and not opts.urn:
                raise TypeError("Missing required property 'filter'")
            __props__.__dict__["filter"] = filter
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["filter"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ReceiptFilter, __self__).__init__(
            'aws-native:ses:ReceiptFilter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ReceiptFilter':
        """
        Get an existing ReceiptFilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ReceiptFilterArgs.__new__(ReceiptFilterArgs)

        __props__.__dict__["aws_id"] = None
        __props__.__dict__["filter"] = None
        return ReceiptFilter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output['outputs.ReceiptFilterFilter']:
        return pulumi.get(self, "filter")

