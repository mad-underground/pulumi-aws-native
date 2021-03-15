# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'StackSetAutoDeploymentArgs',
    'StackSetDeploymentTargetsArgs',
    'StackSetOperationPreferencesArgs',
    'StackSetParameterArgs',
    'StackSetStackInstancesArgs',
]

@pulumi.input_type
class StackSetAutoDeploymentArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 retain_stacks_on_account_removal: Optional[pulumi.Input[bool]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-autodeployment.html
        :param pulumi.Input[bool] enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-autodeployment.html#cfn-cloudformation-stackset-autodeployment-enabled
        :param pulumi.Input[bool] retain_stacks_on_account_removal: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-autodeployment.html#cfn-cloudformation-stackset-autodeployment-retainstacksonaccountremoval
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if retain_stacks_on_account_removal is not None:
            pulumi.set(__self__, "retain_stacks_on_account_removal", retain_stacks_on_account_removal)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-autodeployment.html#cfn-cloudformation-stackset-autodeployment-enabled
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="retainStacksOnAccountRemoval")
    def retain_stacks_on_account_removal(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-autodeployment.html#cfn-cloudformation-stackset-autodeployment-retainstacksonaccountremoval
        """
        return pulumi.get(self, "retain_stacks_on_account_removal")

    @retain_stacks_on_account_removal.setter
    def retain_stacks_on_account_removal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "retain_stacks_on_account_removal", value)


@pulumi.input_type
class StackSetDeploymentTargetsArgs:
    def __init__(__self__, *,
                 accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 organizational_unit_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-deploymenttargets.html
        :param pulumi.Input[Sequence[pulumi.Input[str]]] accounts: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-deploymenttargets.html#cfn-cloudformation-stackset-deploymenttargets-accounts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] organizational_unit_ids: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-deploymenttargets.html#cfn-cloudformation-stackset-deploymenttargets-organizationalunitids
        """
        if accounts is not None:
            pulumi.set(__self__, "accounts", accounts)
        if organizational_unit_ids is not None:
            pulumi.set(__self__, "organizational_unit_ids", organizational_unit_ids)

    @property
    @pulumi.getter
    def accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-deploymenttargets.html#cfn-cloudformation-stackset-deploymenttargets-accounts
        """
        return pulumi.get(self, "accounts")

    @accounts.setter
    def accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "accounts", value)

    @property
    @pulumi.getter(name="organizationalUnitIds")
    def organizational_unit_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-deploymenttargets.html#cfn-cloudformation-stackset-deploymenttargets-organizationalunitids
        """
        return pulumi.get(self, "organizational_unit_ids")

    @organizational_unit_ids.setter
    def organizational_unit_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "organizational_unit_ids", value)


@pulumi.input_type
class StackSetOperationPreferencesArgs:
    def __init__(__self__, *,
                 failure_tolerance_count: Optional[pulumi.Input[int]] = None,
                 failure_tolerance_percentage: Optional[pulumi.Input[int]] = None,
                 max_concurrent_count: Optional[pulumi.Input[int]] = None,
                 max_concurrent_percentage: Optional[pulumi.Input[int]] = None,
                 region_order: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html
        :param pulumi.Input[int] failure_tolerance_count: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-failuretolerancecount
        :param pulumi.Input[int] failure_tolerance_percentage: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-failuretolerancepercentage
        :param pulumi.Input[int] max_concurrent_count: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-maxconcurrentcount
        :param pulumi.Input[int] max_concurrent_percentage: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-maxconcurrentpercentage
        :param pulumi.Input[Sequence[pulumi.Input[str]]] region_order: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-regionorder
        """
        if failure_tolerance_count is not None:
            pulumi.set(__self__, "failure_tolerance_count", failure_tolerance_count)
        if failure_tolerance_percentage is not None:
            pulumi.set(__self__, "failure_tolerance_percentage", failure_tolerance_percentage)
        if max_concurrent_count is not None:
            pulumi.set(__self__, "max_concurrent_count", max_concurrent_count)
        if max_concurrent_percentage is not None:
            pulumi.set(__self__, "max_concurrent_percentage", max_concurrent_percentage)
        if region_order is not None:
            pulumi.set(__self__, "region_order", region_order)

    @property
    @pulumi.getter(name="failureToleranceCount")
    def failure_tolerance_count(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-failuretolerancecount
        """
        return pulumi.get(self, "failure_tolerance_count")

    @failure_tolerance_count.setter
    def failure_tolerance_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "failure_tolerance_count", value)

    @property
    @pulumi.getter(name="failureTolerancePercentage")
    def failure_tolerance_percentage(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-failuretolerancepercentage
        """
        return pulumi.get(self, "failure_tolerance_percentage")

    @failure_tolerance_percentage.setter
    def failure_tolerance_percentage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "failure_tolerance_percentage", value)

    @property
    @pulumi.getter(name="maxConcurrentCount")
    def max_concurrent_count(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-maxconcurrentcount
        """
        return pulumi.get(self, "max_concurrent_count")

    @max_concurrent_count.setter
    def max_concurrent_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_concurrent_count", value)

    @property
    @pulumi.getter(name="maxConcurrentPercentage")
    def max_concurrent_percentage(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-maxconcurrentpercentage
        """
        return pulumi.get(self, "max_concurrent_percentage")

    @max_concurrent_percentage.setter
    def max_concurrent_percentage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_concurrent_percentage", value)

    @property
    @pulumi.getter(name="regionOrder")
    def region_order(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-regionorder
        """
        return pulumi.get(self, "region_order")

    @region_order.setter
    def region_order(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "region_order", value)


@pulumi.input_type
class StackSetParameterArgs:
    def __init__(__self__, *,
                 parameter_key: pulumi.Input[str],
                 parameter_value: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-parameter.html
        :param pulumi.Input[str] parameter_key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-parameter.html#cfn-cloudformation-stackset-parameter-parameterkey
        :param pulumi.Input[str] parameter_value: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-parameter.html#cfn-cloudformation-stackset-parameter-parametervalue
        """
        pulumi.set(__self__, "parameter_key", parameter_key)
        pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="parameterKey")
    def parameter_key(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-parameter.html#cfn-cloudformation-stackset-parameter-parameterkey
        """
        return pulumi.get(self, "parameter_key")

    @parameter_key.setter
    def parameter_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "parameter_key", value)

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-parameter.html#cfn-cloudformation-stackset-parameter-parametervalue
        """
        return pulumi.get(self, "parameter_value")

    @parameter_value.setter
    def parameter_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "parameter_value", value)


@pulumi.input_type
class StackSetStackInstancesArgs:
    def __init__(__self__, *,
                 deployment_targets: pulumi.Input['StackSetDeploymentTargetsArgs'],
                 regions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 parameter_overrides: Optional[pulumi.Input[Sequence[pulumi.Input['StackSetParameterArgs']]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-stackinstances.html
        :param pulumi.Input['StackSetDeploymentTargetsArgs'] deployment_targets: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-stackinstances.html#cfn-cloudformation-stackset-stackinstances-deploymenttargets
        :param pulumi.Input[Sequence[pulumi.Input[str]]] regions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-stackinstances.html#cfn-cloudformation-stackset-stackinstances-regions
        :param pulumi.Input[Sequence[pulumi.Input['StackSetParameterArgs']]] parameter_overrides: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-stackinstances.html#cfn-cloudformation-stackset-stackinstances-parameteroverrides
        """
        pulumi.set(__self__, "deployment_targets", deployment_targets)
        pulumi.set(__self__, "regions", regions)
        if parameter_overrides is not None:
            pulumi.set(__self__, "parameter_overrides", parameter_overrides)

    @property
    @pulumi.getter(name="deploymentTargets")
    def deployment_targets(self) -> pulumi.Input['StackSetDeploymentTargetsArgs']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-stackinstances.html#cfn-cloudformation-stackset-stackinstances-deploymenttargets
        """
        return pulumi.get(self, "deployment_targets")

    @deployment_targets.setter
    def deployment_targets(self, value: pulumi.Input['StackSetDeploymentTargetsArgs']):
        pulumi.set(self, "deployment_targets", value)

    @property
    @pulumi.getter
    def regions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-stackinstances.html#cfn-cloudformation-stackset-stackinstances-regions
        """
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="parameterOverrides")
    def parameter_overrides(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StackSetParameterArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-stackinstances.html#cfn-cloudformation-stackset-stackinstances-parameteroverrides
        """
        return pulumi.get(self, "parameter_overrides")

    @parameter_overrides.setter
    def parameter_overrides(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StackSetParameterArgs']]]]):
        pulumi.set(self, "parameter_overrides", value)


