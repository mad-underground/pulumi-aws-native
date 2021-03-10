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

__all__ = ['Assessment']


class Assessment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assessment_reports_destination: Optional[pulumi.Input[pulumi.InputType['AssessmentAssessmentReportsDestinationArgs']]] = None,
                 aws_account: Optional[pulumi.Input[pulumi.InputType['AssessmentAWSAccountArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 framework_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[pulumi.InputType['AssessmentRolesArgs']]] = None,
                 scope: Optional[pulumi.Input[pulumi.InputType['AssessmentScopeArgs']]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[pulumi.InputType['AssessmentTagsArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AssessmentAssessmentReportsDestinationArgs']] assessment_reports_destination: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-assessmentreportsdestination
        :param pulumi.Input[pulumi.InputType['AssessmentAWSAccountArgs']] aws_account: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-awsaccount
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-description
        :param pulumi.Input[str] framework_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-frameworkid
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-name
        :param pulumi.Input[pulumi.InputType['AssessmentRolesArgs']] roles: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-roles
        :param pulumi.Input[pulumi.InputType['AssessmentScopeArgs']] scope: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-scope
        :param pulumi.Input[str] status: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-status
        :param pulumi.Input[pulumi.InputType['AssessmentTagsArgs']] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-tags
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

            __props__['assessment_reports_destination'] = assessment_reports_destination
            __props__['aws_account'] = aws_account
            __props__['description'] = description
            __props__['framework_id'] = framework_id
            __props__['name'] = name
            __props__['roles'] = roles
            __props__['scope'] = scope
            __props__['status'] = status
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['assessment_id'] = None
            __props__['creation_time'] = None
            __props__['delegations'] = None
        super(Assessment, __self__).__init__(
            'aws-native:AuditManager:Assessment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Assessment':
        """
        Get an existing Assessment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Assessment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="assessmentId")
    def assessment_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "assessment_id")

    @property
    @pulumi.getter(name="assessmentReportsDestination")
    def assessment_reports_destination(self) -> pulumi.Output[Optional['outputs.AssessmentAssessmentReportsDestination']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-assessmentreportsdestination
        """
        return pulumi.get(self, "assessment_reports_destination")

    @property
    @pulumi.getter(name="awsAccount")
    def aws_account(self) -> pulumi.Output[Optional['outputs.AssessmentAWSAccount']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-awsaccount
        """
        return pulumi.get(self, "aws_account")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[float]:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def delegations(self) -> pulumi.Output['outputs.AssessmentDelegations']:
        return pulumi.get(self, "delegations")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="frameworkId")
    def framework_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "framework_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Output[Optional['outputs.AssessmentRoles']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-roles
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[Optional['outputs.AssessmentScope']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-scope
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional['outputs.AssessmentTags']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-auditmanager-assessment.html#cfn-auditmanager-assessment-tags
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

