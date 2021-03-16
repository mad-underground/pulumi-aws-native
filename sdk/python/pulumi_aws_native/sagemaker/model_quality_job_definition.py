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

__all__ = ['ModelQualityJobDefinition']


class ModelQualityJobDefinition(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_definition_name: Optional[pulumi.Input[str]] = None,
                 job_resources: Optional[pulumi.Input[pulumi.InputType['ModelQualityJobDefinitionMonitoringResourcesArgs']]] = None,
                 model_quality_app_specification: Optional[pulumi.Input[pulumi.InputType['ModelQualityJobDefinitionModelQualityAppSpecificationArgs']]] = None,
                 model_quality_baseline_config: Optional[pulumi.Input[pulumi.InputType['ModelQualityJobDefinitionModelQualityBaselineConfigArgs']]] = None,
                 model_quality_job_input: Optional[pulumi.Input[pulumi.InputType['ModelQualityJobDefinitionModelQualityJobInputArgs']]] = None,
                 model_quality_job_output_config: Optional[pulumi.Input[pulumi.InputType['ModelQualityJobDefinitionMonitoringOutputConfigArgs']]] = None,
                 network_config: Optional[pulumi.Input[pulumi.InputType['ModelQualityJobDefinitionNetworkConfigArgs']]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stopping_condition: Optional[pulumi.Input[pulumi.InputType['ModelQualityJobDefinitionStoppingConditionArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] job_definition_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-jobdefinitionname
        :param pulumi.Input[pulumi.InputType['ModelQualityJobDefinitionMonitoringResourcesArgs']] job_resources: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-jobresources
        :param pulumi.Input[pulumi.InputType['ModelQualityJobDefinitionModelQualityAppSpecificationArgs']] model_quality_app_specification: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityappspecification
        :param pulumi.Input[pulumi.InputType['ModelQualityJobDefinitionModelQualityBaselineConfigArgs']] model_quality_baseline_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualitybaselineconfig
        :param pulumi.Input[pulumi.InputType['ModelQualityJobDefinitionModelQualityJobInputArgs']] model_quality_job_input: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityjobinput
        :param pulumi.Input[pulumi.InputType['ModelQualityJobDefinitionMonitoringOutputConfigArgs']] model_quality_job_output_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityjoboutputconfig
        :param pulumi.Input[pulumi.InputType['ModelQualityJobDefinitionNetworkConfigArgs']] network_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-networkconfig
        :param pulumi.Input[str] role_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-rolearn
        :param pulumi.Input[pulumi.InputType['ModelQualityJobDefinitionStoppingConditionArgs']] stopping_condition: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-stoppingcondition
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-tags
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

            __props__['job_definition_name'] = job_definition_name
            if job_resources is None and not opts.urn:
                raise TypeError("Missing required property 'job_resources'")
            __props__['job_resources'] = job_resources
            if model_quality_app_specification is None and not opts.urn:
                raise TypeError("Missing required property 'model_quality_app_specification'")
            __props__['model_quality_app_specification'] = model_quality_app_specification
            __props__['model_quality_baseline_config'] = model_quality_baseline_config
            if model_quality_job_input is None and not opts.urn:
                raise TypeError("Missing required property 'model_quality_job_input'")
            __props__['model_quality_job_input'] = model_quality_job_input
            if model_quality_job_output_config is None and not opts.urn:
                raise TypeError("Missing required property 'model_quality_job_output_config'")
            __props__['model_quality_job_output_config'] = model_quality_job_output_config
            __props__['network_config'] = network_config
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__['role_arn'] = role_arn
            __props__['stopping_condition'] = stopping_condition
            __props__['tags'] = tags
            __props__['creation_time'] = None
            __props__['job_definition_arn'] = None
        super(ModelQualityJobDefinition, __self__).__init__(
            'aws-native:SageMaker:ModelQualityJobDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ModelQualityJobDefinition':
        """
        Get an existing ModelQualityJobDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ModelQualityJobDefinition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="jobDefinitionArn")
    def job_definition_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "job_definition_arn")

    @property
    @pulumi.getter(name="jobDefinitionName")
    def job_definition_name(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-jobdefinitionname
        """
        return pulumi.get(self, "job_definition_name")

    @property
    @pulumi.getter(name="jobResources")
    def job_resources(self) -> pulumi.Output['outputs.ModelQualityJobDefinitionMonitoringResources']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-jobresources
        """
        return pulumi.get(self, "job_resources")

    @property
    @pulumi.getter(name="modelQualityAppSpecification")
    def model_quality_app_specification(self) -> pulumi.Output['outputs.ModelQualityJobDefinitionModelQualityAppSpecification']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityappspecification
        """
        return pulumi.get(self, "model_quality_app_specification")

    @property
    @pulumi.getter(name="modelQualityBaselineConfig")
    def model_quality_baseline_config(self) -> pulumi.Output[Optional['outputs.ModelQualityJobDefinitionModelQualityBaselineConfig']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualitybaselineconfig
        """
        return pulumi.get(self, "model_quality_baseline_config")

    @property
    @pulumi.getter(name="modelQualityJobInput")
    def model_quality_job_input(self) -> pulumi.Output['outputs.ModelQualityJobDefinitionModelQualityJobInput']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityjobinput
        """
        return pulumi.get(self, "model_quality_job_input")

    @property
    @pulumi.getter(name="modelQualityJobOutputConfig")
    def model_quality_job_output_config(self) -> pulumi.Output['outputs.ModelQualityJobDefinitionMonitoringOutputConfig']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityjoboutputconfig
        """
        return pulumi.get(self, "model_quality_job_output_config")

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> pulumi.Output[Optional['outputs.ModelQualityJobDefinitionNetworkConfig']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-networkconfig
        """
        return pulumi.get(self, "network_config")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-rolearn
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="stoppingCondition")
    def stopping_condition(self) -> pulumi.Output[Optional['outputs.ModelQualityJobDefinitionStoppingCondition']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-stoppingcondition
        """
        return pulumi.get(self, "stopping_condition")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-tags
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
