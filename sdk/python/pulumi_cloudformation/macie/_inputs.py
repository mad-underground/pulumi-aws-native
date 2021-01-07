# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'CustomDataIdentifierPropertiesArgs',
    'FindingsFilterCriterionArgs',
    'FindingsFilterFindingCriteriaArgs',
    'FindingsFilterPropertiesArgs',
    'SessionPropertiesArgs',
]

@pulumi.input_type
class CustomDataIdentifierPropertiesArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 regex: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 ignore_words: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 keywords: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 maximum_match_distance: Optional[pulumi.Input[int]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-name
        :param pulumi.Input[str] regex: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-regex
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-description
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ignore_words: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-ignorewords
        :param pulumi.Input[Sequence[pulumi.Input[str]]] keywords: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-keywords
        :param pulumi.Input[int] maximum_match_distance: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-maximummatchdistance
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "regex", regex)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ignore_words is not None:
            pulumi.set(__self__, "ignore_words", ignore_words)
        if keywords is not None:
            pulumi.set(__self__, "keywords", keywords)
        if maximum_match_distance is not None:
            pulumi.set(__self__, "maximum_match_distance", maximum_match_distance)

    @property
    @pulumi.getter(name="Name")
    def name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="Regex")
    def regex(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-regex
        """
        return pulumi.get(self, "regex")

    @regex.setter
    def regex(self, value: pulumi.Input[str]):
        pulumi.set(self, "regex", value)

    @property
    @pulumi.getter(name="Description")
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="IgnoreWords")
    def ignore_words(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-ignorewords
        """
        return pulumi.get(self, "ignore_words")

    @ignore_words.setter
    def ignore_words(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ignore_words", value)

    @property
    @pulumi.getter(name="Keywords")
    def keywords(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-keywords
        """
        return pulumi.get(self, "keywords")

    @keywords.setter
    def keywords(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "keywords", value)

    @property
    @pulumi.getter(name="MaximumMatchDistance")
    def maximum_match_distance(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-customdataidentifier.html#cfn-macie-customdataidentifier-maximummatchdistance
        """
        return pulumi.get(self, "maximum_match_distance")

    @maximum_match_distance.setter
    def maximum_match_distance(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maximum_match_distance", value)


@pulumi.input_type
class FindingsFilterCriterionArgs:
    def __init__(__self__):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-criterion.html
        """
        pass


@pulumi.input_type
class FindingsFilterFindingCriteriaArgs:
    def __init__(__self__, *,
                 criterion: Optional[pulumi.Input['FindingsFilterCriterionArgs']] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-findingcriteria.html
        :param pulumi.Input['FindingsFilterCriterionArgs'] criterion: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-findingcriteria.html#cfn-macie-findingsfilter-findingcriteria-criterion
        """
        if criterion is not None:
            pulumi.set(__self__, "criterion", criterion)

    @property
    @pulumi.getter(name="Criterion")
    def criterion(self) -> Optional[pulumi.Input['FindingsFilterCriterionArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-findingcriteria.html#cfn-macie-findingsfilter-findingcriteria-criterion
        """
        return pulumi.get(self, "criterion")

    @criterion.setter
    def criterion(self, value: Optional[pulumi.Input['FindingsFilterCriterionArgs']]):
        pulumi.set(self, "criterion", value)


@pulumi.input_type
class FindingsFilterPropertiesArgs:
    def __init__(__self__, *,
                 finding_criteria: pulumi.Input['FindingsFilterFindingCriteriaArgs'],
                 name: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-findingsfilter.html
        :param pulumi.Input['FindingsFilterFindingCriteriaArgs'] finding_criteria: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-findingsfilter.html#cfn-macie-findingsfilter-findingcriteria
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-findingsfilter.html#cfn-macie-findingsfilter-name
        :param pulumi.Input[str] action: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-findingsfilter.html#cfn-macie-findingsfilter-action
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-findingsfilter.html#cfn-macie-findingsfilter-description
        :param pulumi.Input[int] position: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-findingsfilter.html#cfn-macie-findingsfilter-position
        """
        pulumi.set(__self__, "finding_criteria", finding_criteria)
        pulumi.set(__self__, "name", name)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if position is not None:
            pulumi.set(__self__, "position", position)

    @property
    @pulumi.getter(name="FindingCriteria")
    def finding_criteria(self) -> pulumi.Input['FindingsFilterFindingCriteriaArgs']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-findingsfilter.html#cfn-macie-findingsfilter-findingcriteria
        """
        return pulumi.get(self, "finding_criteria")

    @finding_criteria.setter
    def finding_criteria(self, value: pulumi.Input['FindingsFilterFindingCriteriaArgs']):
        pulumi.set(self, "finding_criteria", value)

    @property
    @pulumi.getter(name="Name")
    def name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-findingsfilter.html#cfn-macie-findingsfilter-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="Action")
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-findingsfilter.html#cfn-macie-findingsfilter-action
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="Description")
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-findingsfilter.html#cfn-macie-findingsfilter-description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="Position")
    def position(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-findingsfilter.html#cfn-macie-findingsfilter-position
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "position", value)


@pulumi.input_type
class SessionPropertiesArgs:
    def __init__(__self__, *,
                 finding_publishing_frequency: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html
        :param pulumi.Input[str] finding_publishing_frequency: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html#cfn-macie-session-findingpublishingfrequency
        :param pulumi.Input[str] status: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html#cfn-macie-session-status
        """
        if finding_publishing_frequency is not None:
            pulumi.set(__self__, "finding_publishing_frequency", finding_publishing_frequency)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="FindingPublishingFrequency")
    def finding_publishing_frequency(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html#cfn-macie-session-findingpublishingfrequency
        """
        return pulumi.get(self, "finding_publishing_frequency")

    @finding_publishing_frequency.setter
    def finding_publishing_frequency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "finding_publishing_frequency", value)

    @property
    @pulumi.getter(name="Status")
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html#cfn-macie-session-status
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

