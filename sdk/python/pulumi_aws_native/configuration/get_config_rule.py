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

__all__ = [
    'GetConfigRuleResult',
    'AwaitableGetConfigRuleResult',
    'get_config_rule',
    'get_config_rule_output',
]

@pulumi.output_type
class GetConfigRuleResult:
    def __init__(__self__, arn=None, compliance=None, config_rule_id=None, description=None, evaluation_modes=None, input_parameters=None, maximum_execution_frequency=None, scope=None, source=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if compliance and not isinstance(compliance, dict):
            raise TypeError("Expected argument 'compliance' to be a dict")
        pulumi.set(__self__, "compliance", compliance)
        if config_rule_id and not isinstance(config_rule_id, str):
            raise TypeError("Expected argument 'config_rule_id' to be a str")
        pulumi.set(__self__, "config_rule_id", config_rule_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if evaluation_modes and not isinstance(evaluation_modes, list):
            raise TypeError("Expected argument 'evaluation_modes' to be a list")
        pulumi.set(__self__, "evaluation_modes", evaluation_modes)
        if input_parameters and not isinstance(input_parameters, str):
            raise TypeError("Expected argument 'input_parameters' to be a str")
        pulumi.set(__self__, "input_parameters", input_parameters)
        if maximum_execution_frequency and not isinstance(maximum_execution_frequency, str):
            raise TypeError("Expected argument 'maximum_execution_frequency' to be a str")
        pulumi.set(__self__, "maximum_execution_frequency", maximum_execution_frequency)
        if scope and not isinstance(scope, dict):
            raise TypeError("Expected argument 'scope' to be a dict")
        pulumi.set(__self__, "scope", scope)
        if source and not isinstance(source, dict):
            raise TypeError("Expected argument 'source' to be a dict")
        pulumi.set(__self__, "source", source)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def compliance(self) -> Optional['outputs.ComplianceProperties']:
        """
        Indicates whether an AWS resource or CC rule is compliant and provides the number of contributors that affect the compliance.
        """
        return pulumi.get(self, "compliance")

    @property
    @pulumi.getter(name="configRuleId")
    def config_rule_id(self) -> Optional[str]:
        return pulumi.get(self, "config_rule_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description that you provide for the CC rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="evaluationModes")
    def evaluation_modes(self) -> Optional[Sequence['outputs.ConfigRuleEvaluationModeConfiguration']]:
        """
        The modes the CC rule can be evaluated in. The valid values are distinct objects. By default, the value is Detective evaluation mode only.
        """
        return pulumi.get(self, "evaluation_modes")

    @property
    @pulumi.getter(name="inputParameters")
    def input_parameters(self) -> Optional[str]:
        """
        A string, in JSON format, that is passed to the CC rule Lambda function.
        """
        return pulumi.get(self, "input_parameters")

    @property
    @pulumi.getter(name="maximumExecutionFrequency")
    def maximum_execution_frequency(self) -> Optional[str]:
        """
        The maximum frequency with which CC runs evaluations for a rule. You can specify a value for ``MaximumExecutionFrequency`` when:
          +  You are using an AWS managed rule that is triggered at a periodic frequency.
          +  Your custom rule is triggered when CC delivers the configuration snapshot. For more information, see [ConfigSnapshotDeliveryProperties](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html).
          
          By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the ``MaximumExecutionFrequency`` parameter.
        """
        return pulumi.get(self, "maximum_execution_frequency")

    @property
    @pulumi.getter
    def scope(self) -> Optional['outputs.ConfigRuleScope']:
        """
        Defines which resources can trigger an evaluation for the rule. The scope can include one or more resource types, a combination of one resource type and one resource ID, or a combination of a tag key and value. Specify a scope to constrain the resources that can trigger an evaluation for the rule. If you do not specify a scope, evaluations are triggered when any resource in the recording group changes.
          The scope can be empty.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def source(self) -> Optional['outputs.ConfigRuleSource']:
        """
        Provides the rule owner (```` for managed rules, ``CUSTOM_POLICY`` for Custom Policy rules, and ``CUSTOM_LAMBDA`` for Custom Lambda rules), the rule identifier, and the notifications that cause the function to evaluate your AWS resources.
        """
        return pulumi.get(self, "source")


class AwaitableGetConfigRuleResult(GetConfigRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConfigRuleResult(
            arn=self.arn,
            compliance=self.compliance,
            config_rule_id=self.config_rule_id,
            description=self.description,
            evaluation_modes=self.evaluation_modes,
            input_parameters=self.input_parameters,
            maximum_execution_frequency=self.maximum_execution_frequency,
            scope=self.scope,
            source=self.source)


def get_config_rule(config_rule_name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConfigRuleResult:
    """
    You must first create and start the CC configuration recorder in order to create CC managed rules with CFNlong. For more information, see [Managing the Configuration Recorder](https://docs.aws.amazon.com/config/latest/developerguide/stop-start-recorder.html).
     Adds or updates an CC rule to evaluate if your AWS resources comply with your desired configurations. For information on how many CC rules you can have per account, see [Service Limits](https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html) in the *Developer Guide*.
     There are two types of rules: *Managed Rules* and *Custom Rules*. You can use the ``ConfigRule`` resource to create both CC Managed Rules and CC Custom Rules.
     CC Managed Rules are predefined, customizable rules created by CC. For a list of managed rules, see [List of Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html). If you are adding an CC managed rule, you must specify the rule's identifi


    :param str config_rule_name: A name for the CC rule. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the rule name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
    """
    __args__ = dict()
    __args__['configRuleName'] = config_rule_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:configuration:getConfigRule', __args__, opts=opts, typ=GetConfigRuleResult).value

    return AwaitableGetConfigRuleResult(
        arn=pulumi.get(__ret__, 'arn'),
        compliance=pulumi.get(__ret__, 'compliance'),
        config_rule_id=pulumi.get(__ret__, 'config_rule_id'),
        description=pulumi.get(__ret__, 'description'),
        evaluation_modes=pulumi.get(__ret__, 'evaluation_modes'),
        input_parameters=pulumi.get(__ret__, 'input_parameters'),
        maximum_execution_frequency=pulumi.get(__ret__, 'maximum_execution_frequency'),
        scope=pulumi.get(__ret__, 'scope'),
        source=pulumi.get(__ret__, 'source'))


@_utilities.lift_output_func(get_config_rule)
def get_config_rule_output(config_rule_name: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConfigRuleResult]:
    """
    You must first create and start the CC configuration recorder in order to create CC managed rules with CFNlong. For more information, see [Managing the Configuration Recorder](https://docs.aws.amazon.com/config/latest/developerguide/stop-start-recorder.html).
     Adds or updates an CC rule to evaluate if your AWS resources comply with your desired configurations. For information on how many CC rules you can have per account, see [Service Limits](https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html) in the *Developer Guide*.
     There are two types of rules: *Managed Rules* and *Custom Rules*. You can use the ``ConfigRule`` resource to create both CC Managed Rules and CC Custom Rules.
     CC Managed Rules are predefined, customizable rules created by CC. For a list of managed rules, see [List of Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html). If you are adding an CC managed rule, you must specify the rule's identifi


    :param str config_rule_name: A name for the CC rule. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the rule name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
    """
    ...
