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

__all__ = ['OrganizationConformancePack']


class OrganizationConformancePack(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 conformance_pack_input_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrganizationConformancePackConformancePackInputParameterArgs']]]]] = None,
                 delivery_s3_bucket: Optional[pulumi.Input[str]] = None,
                 delivery_s3_key_prefix: Optional[pulumi.Input[str]] = None,
                 excluded_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 organization_conformance_pack_name: Optional[pulumi.Input[str]] = None,
                 template_body: Optional[pulumi.Input[str]] = None,
                 template_s3_uri: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrganizationConformancePackConformancePackInputParameterArgs']]]] conformance_pack_input_parameters: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-conformancepackinputparameters
        :param pulumi.Input[str] delivery_s3_bucket: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-deliverys3bucket
        :param pulumi.Input[str] delivery_s3_key_prefix: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-deliverys3keyprefix
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_accounts: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-excludedaccounts
        :param pulumi.Input[str] organization_conformance_pack_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-organizationconformancepackname
        :param pulumi.Input[str] template_body: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-templatebody
        :param pulumi.Input[str] template_s3_uri: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-templates3uri
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

            __props__['conformance_pack_input_parameters'] = conformance_pack_input_parameters
            __props__['delivery_s3_bucket'] = delivery_s3_bucket
            __props__['delivery_s3_key_prefix'] = delivery_s3_key_prefix
            __props__['excluded_accounts'] = excluded_accounts
            if organization_conformance_pack_name is None and not opts.urn:
                raise TypeError("Missing required property 'organization_conformance_pack_name'")
            __props__['organization_conformance_pack_name'] = organization_conformance_pack_name
            __props__['template_body'] = template_body
            __props__['template_s3_uri'] = template_s3_uri
        super(OrganizationConformancePack, __self__).__init__(
            'aws-native:Configuration:OrganizationConformancePack',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'OrganizationConformancePack':
        """
        Get an existing OrganizationConformancePack resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return OrganizationConformancePack(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ConformancePackInputParameters")
    def conformance_pack_input_parameters(self) -> pulumi.Output[Optional[Sequence['outputs.OrganizationConformancePackConformancePackInputParameter']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-conformancepackinputparameters
        """
        return pulumi.get(self, "conformance_pack_input_parameters")

    @property
    @pulumi.getter(name="DeliveryS3Bucket")
    def delivery_s3_bucket(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-deliverys3bucket
        """
        return pulumi.get(self, "delivery_s3_bucket")

    @property
    @pulumi.getter(name="DeliveryS3KeyPrefix")
    def delivery_s3_key_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-deliverys3keyprefix
        """
        return pulumi.get(self, "delivery_s3_key_prefix")

    @property
    @pulumi.getter(name="ExcludedAccounts")
    def excluded_accounts(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-excludedaccounts
        """
        return pulumi.get(self, "excluded_accounts")

    @property
    @pulumi.getter(name="OrganizationConformancePackName")
    def organization_conformance_pack_name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-organizationconformancepackname
        """
        return pulumi.get(self, "organization_conformance_pack_name")

    @property
    @pulumi.getter(name="TemplateBody")
    def template_body(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-templatebody
        """
        return pulumi.get(self, "template_body")

    @property
    @pulumi.getter(name="TemplateS3Uri")
    def template_s3_uri(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-templates3uri
        """
        return pulumi.get(self, "template_s3_uri")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

