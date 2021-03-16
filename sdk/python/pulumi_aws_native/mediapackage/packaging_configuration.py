# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._inputs import *

__all__ = ['PackagingConfiguration']


class PackagingConfiguration(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cmaf_package: Optional[pulumi.Input[pulumi.InputType['PackagingConfigurationCmafPackageArgs']]] = None,
                 dash_package: Optional[pulumi.Input[pulumi.InputType['PackagingConfigurationDashPackageArgs']]] = None,
                 hls_package: Optional[pulumi.Input[pulumi.InputType['PackagingConfigurationHlsPackageArgs']]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 mss_package: Optional[pulumi.Input[pulumi.InputType['PackagingConfigurationMssPackageArgs']]] = None,
                 packaging_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PackagingConfigurationCmafPackageArgs']] cmaf_package: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-cmafpackage
        :param pulumi.Input[pulumi.InputType['PackagingConfigurationDashPackageArgs']] dash_package: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-dashpackage
        :param pulumi.Input[pulumi.InputType['PackagingConfigurationHlsPackageArgs']] hls_package: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-hlspackage
        :param pulumi.Input[str] id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-id
        :param pulumi.Input[pulumi.InputType['PackagingConfigurationMssPackageArgs']] mss_package: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-msspackage
        :param pulumi.Input[str] packaging_group_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-packaginggroupid
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-tags
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

            __props__['cmaf_package'] = cmaf_package
            __props__['dash_package'] = dash_package
            __props__['hls_package'] = hls_package
            if id is None and not opts.urn:
                raise TypeError("Missing required property 'id'")
            __props__['id'] = id
            __props__['mss_package'] = mss_package
            if packaging_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'packaging_group_id'")
            __props__['packaging_group_id'] = packaging_group_id
            __props__['tags'] = tags
            __props__['arn'] = None
        super(PackagingConfiguration, __self__).__init__(
            'aws-native:MediaPackage:PackagingConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PackagingConfiguration':
        """
        Get an existing PackagingConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return PackagingConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="cmafPackage")
    def cmaf_package(self) -> pulumi.Output[Optional['outputs.PackagingConfigurationCmafPackage']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-cmafpackage
        """
        return pulumi.get(self, "cmaf_package")

    @property
    @pulumi.getter(name="dashPackage")
    def dash_package(self) -> pulumi.Output[Optional['outputs.PackagingConfigurationDashPackage']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-dashpackage
        """
        return pulumi.get(self, "dash_package")

    @property
    @pulumi.getter(name="hlsPackage")
    def hls_package(self) -> pulumi.Output[Optional['outputs.PackagingConfigurationHlsPackage']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-hlspackage
        """
        return pulumi.get(self, "hls_package")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="mssPackage")
    def mss_package(self) -> pulumi.Output[Optional['outputs.PackagingConfigurationMssPackage']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-msspackage
        """
        return pulumi.get(self, "mss_package")

    @property
    @pulumi.getter(name="packagingGroupId")
    def packaging_group_id(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-packaginggroupid
        """
        return pulumi.get(self, "packaging_group_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-tags
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
