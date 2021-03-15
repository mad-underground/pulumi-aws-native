# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'AnalyzerArchiveRuleArgs',
    'AnalyzerFilterArgs',
]

@pulumi.input_type
class AnalyzerArchiveRuleArgs:
    def __init__(__self__, *,
                 filter: pulumi.Input[Sequence[pulumi.Input['AnalyzerFilterArgs']]],
                 rule_name: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-archiverule.html
        :param pulumi.Input[Sequence[pulumi.Input['AnalyzerFilterArgs']]] filter: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-archiverule.html#cfn-accessanalyzer-analyzer-archiverule-filter
        :param pulumi.Input[str] rule_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-archiverule.html#cfn-accessanalyzer-analyzer-archiverule-rulename
        """
        pulumi.set(__self__, "filter", filter)
        pulumi.set(__self__, "rule_name", rule_name)

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Input[Sequence[pulumi.Input['AnalyzerFilterArgs']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-archiverule.html#cfn-accessanalyzer-analyzer-archiverule-filter
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: pulumi.Input[Sequence[pulumi.Input['AnalyzerFilterArgs']]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-archiverule.html#cfn-accessanalyzer-analyzer-archiverule-rulename
        """
        return pulumi.get(self, "rule_name")

    @rule_name.setter
    def rule_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_name", value)


@pulumi.input_type
class AnalyzerFilterArgs:
    def __init__(__self__, *,
                 property: pulumi.Input[str],
                 contains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 eq: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 exists: Optional[pulumi.Input[bool]] = None,
                 neq: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html
        :param pulumi.Input[str] property: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html#cfn-accessanalyzer-analyzer-filter-property
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contains: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html#cfn-accessanalyzer-analyzer-filter-contains
        :param pulumi.Input[Sequence[pulumi.Input[str]]] eq: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html#cfn-accessanalyzer-analyzer-filter-eq
        :param pulumi.Input[bool] exists: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html#cfn-accessanalyzer-analyzer-filter-exists
        :param pulumi.Input[Sequence[pulumi.Input[str]]] neq: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html#cfn-accessanalyzer-analyzer-filter-neq
        """
        pulumi.set(__self__, "property", property)
        if contains is not None:
            pulumi.set(__self__, "contains", contains)
        if eq is not None:
            pulumi.set(__self__, "eq", eq)
        if exists is not None:
            pulumi.set(__self__, "exists", exists)
        if neq is not None:
            pulumi.set(__self__, "neq", neq)

    @property
    @pulumi.getter
    def contains(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html#cfn-accessanalyzer-analyzer-filter-contains
        """
        return pulumi.get(self, "contains")

    @contains.setter
    def contains(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "contains", value)

    @property
    @pulumi.getter
    def eq(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html#cfn-accessanalyzer-analyzer-filter-eq
        """
        return pulumi.get(self, "eq")

    @eq.setter
    def eq(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "eq", value)

    @property
    @pulumi.getter
    def exists(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html#cfn-accessanalyzer-analyzer-filter-exists
        """
        return pulumi.get(self, "exists")

    @exists.setter
    def exists(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "exists", value)

    @property
    @pulumi.getter
    def neq(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html#cfn-accessanalyzer-analyzer-filter-neq
        """
        return pulumi.get(self, "neq")

    @neq.setter
    def neq(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "neq", value)

    @property
    @pulumi.getter
    def property(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html#cfn-accessanalyzer-analyzer-filter-property
        """
        return pulumi.get(self, "property")

    @property.setter
    def property(self, value: pulumi.Input[str]):
        pulumi.set(self, "property", value)


