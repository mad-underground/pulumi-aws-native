# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'FilterCriteriaArgs',
    'FilterDateFilterArgs',
    'FilterMapFilterArgs',
    'FilterNumberFilterArgs',
    'FilterPackageFilterArgs',
    'FilterPortRangeFilterArgs',
    'FilterStringFilterArgs',
]

@pulumi.input_type
class FilterCriteriaArgs:
    def __init__(__self__, *,
                 aws_account_id: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 component_id: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 component_type: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 ec2_instance_image_id: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 ec2_instance_subnet_id: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 ec2_instance_vpc_id: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 ecr_image_architecture: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 ecr_image_hash: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 ecr_image_pushed_at: Optional[pulumi.Input[Sequence[pulumi.Input['FilterDateFilterArgs']]]] = None,
                 ecr_image_registry: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 ecr_image_repository_name: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 ecr_image_tags: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 finding_arn: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 finding_status: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 finding_type: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 first_observed_at: Optional[pulumi.Input[Sequence[pulumi.Input['FilterDateFilterArgs']]]] = None,
                 inspector_score: Optional[pulumi.Input[Sequence[pulumi.Input['FilterNumberFilterArgs']]]] = None,
                 last_observed_at: Optional[pulumi.Input[Sequence[pulumi.Input['FilterDateFilterArgs']]]] = None,
                 network_protocol: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 port_range: Optional[pulumi.Input[Sequence[pulumi.Input['FilterPortRangeFilterArgs']]]] = None,
                 related_vulnerabilities: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 resource_id: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 resource_tags: Optional[pulumi.Input[Sequence[pulumi.Input['FilterMapFilterArgs']]]] = None,
                 resource_type: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 severity: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 title: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 updated_at: Optional[pulumi.Input[Sequence[pulumi.Input['FilterDateFilterArgs']]]] = None,
                 vendor_severity: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 vulnerability_id: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 vulnerability_source: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]] = None,
                 vulnerable_packages: Optional[pulumi.Input[Sequence[pulumi.Input['FilterPackageFilterArgs']]]] = None):
        if aws_account_id is not None:
            pulumi.set(__self__, "aws_account_id", aws_account_id)
        if component_id is not None:
            pulumi.set(__self__, "component_id", component_id)
        if component_type is not None:
            pulumi.set(__self__, "component_type", component_type)
        if ec2_instance_image_id is not None:
            pulumi.set(__self__, "ec2_instance_image_id", ec2_instance_image_id)
        if ec2_instance_subnet_id is not None:
            pulumi.set(__self__, "ec2_instance_subnet_id", ec2_instance_subnet_id)
        if ec2_instance_vpc_id is not None:
            pulumi.set(__self__, "ec2_instance_vpc_id", ec2_instance_vpc_id)
        if ecr_image_architecture is not None:
            pulumi.set(__self__, "ecr_image_architecture", ecr_image_architecture)
        if ecr_image_hash is not None:
            pulumi.set(__self__, "ecr_image_hash", ecr_image_hash)
        if ecr_image_pushed_at is not None:
            pulumi.set(__self__, "ecr_image_pushed_at", ecr_image_pushed_at)
        if ecr_image_registry is not None:
            pulumi.set(__self__, "ecr_image_registry", ecr_image_registry)
        if ecr_image_repository_name is not None:
            pulumi.set(__self__, "ecr_image_repository_name", ecr_image_repository_name)
        if ecr_image_tags is not None:
            pulumi.set(__self__, "ecr_image_tags", ecr_image_tags)
        if finding_arn is not None:
            pulumi.set(__self__, "finding_arn", finding_arn)
        if finding_status is not None:
            pulumi.set(__self__, "finding_status", finding_status)
        if finding_type is not None:
            pulumi.set(__self__, "finding_type", finding_type)
        if first_observed_at is not None:
            pulumi.set(__self__, "first_observed_at", first_observed_at)
        if inspector_score is not None:
            pulumi.set(__self__, "inspector_score", inspector_score)
        if last_observed_at is not None:
            pulumi.set(__self__, "last_observed_at", last_observed_at)
        if network_protocol is not None:
            pulumi.set(__self__, "network_protocol", network_protocol)
        if port_range is not None:
            pulumi.set(__self__, "port_range", port_range)
        if related_vulnerabilities is not None:
            pulumi.set(__self__, "related_vulnerabilities", related_vulnerabilities)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_tags is not None:
            pulumi.set(__self__, "resource_tags", resource_tags)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if vendor_severity is not None:
            pulumi.set(__self__, "vendor_severity", vendor_severity)
        if vulnerability_id is not None:
            pulumi.set(__self__, "vulnerability_id", vulnerability_id)
        if vulnerability_source is not None:
            pulumi.set(__self__, "vulnerability_source", vulnerability_source)
        if vulnerable_packages is not None:
            pulumi.set(__self__, "vulnerable_packages", vulnerable_packages)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "aws_account_id")

    @aws_account_id.setter
    def aws_account_id(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "aws_account_id", value)

    @property
    @pulumi.getter(name="componentId")
    def component_id(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "component_id")

    @component_id.setter
    def component_id(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "component_id", value)

    @property
    @pulumi.getter(name="componentType")
    def component_type(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "component_type")

    @component_type.setter
    def component_type(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "component_type", value)

    @property
    @pulumi.getter(name="ec2InstanceImageId")
    def ec2_instance_image_id(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "ec2_instance_image_id")

    @ec2_instance_image_id.setter
    def ec2_instance_image_id(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "ec2_instance_image_id", value)

    @property
    @pulumi.getter(name="ec2InstanceSubnetId")
    def ec2_instance_subnet_id(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "ec2_instance_subnet_id")

    @ec2_instance_subnet_id.setter
    def ec2_instance_subnet_id(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "ec2_instance_subnet_id", value)

    @property
    @pulumi.getter(name="ec2InstanceVpcId")
    def ec2_instance_vpc_id(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "ec2_instance_vpc_id")

    @ec2_instance_vpc_id.setter
    def ec2_instance_vpc_id(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "ec2_instance_vpc_id", value)

    @property
    @pulumi.getter(name="ecrImageArchitecture")
    def ecr_image_architecture(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "ecr_image_architecture")

    @ecr_image_architecture.setter
    def ecr_image_architecture(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "ecr_image_architecture", value)

    @property
    @pulumi.getter(name="ecrImageHash")
    def ecr_image_hash(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "ecr_image_hash")

    @ecr_image_hash.setter
    def ecr_image_hash(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "ecr_image_hash", value)

    @property
    @pulumi.getter(name="ecrImagePushedAt")
    def ecr_image_pushed_at(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterDateFilterArgs']]]]:
        return pulumi.get(self, "ecr_image_pushed_at")

    @ecr_image_pushed_at.setter
    def ecr_image_pushed_at(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterDateFilterArgs']]]]):
        pulumi.set(self, "ecr_image_pushed_at", value)

    @property
    @pulumi.getter(name="ecrImageRegistry")
    def ecr_image_registry(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "ecr_image_registry")

    @ecr_image_registry.setter
    def ecr_image_registry(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "ecr_image_registry", value)

    @property
    @pulumi.getter(name="ecrImageRepositoryName")
    def ecr_image_repository_name(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "ecr_image_repository_name")

    @ecr_image_repository_name.setter
    def ecr_image_repository_name(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "ecr_image_repository_name", value)

    @property
    @pulumi.getter(name="ecrImageTags")
    def ecr_image_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "ecr_image_tags")

    @ecr_image_tags.setter
    def ecr_image_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "ecr_image_tags", value)

    @property
    @pulumi.getter(name="findingArn")
    def finding_arn(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "finding_arn")

    @finding_arn.setter
    def finding_arn(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "finding_arn", value)

    @property
    @pulumi.getter(name="findingStatus")
    def finding_status(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "finding_status")

    @finding_status.setter
    def finding_status(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "finding_status", value)

    @property
    @pulumi.getter(name="findingType")
    def finding_type(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "finding_type")

    @finding_type.setter
    def finding_type(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "finding_type", value)

    @property
    @pulumi.getter(name="firstObservedAt")
    def first_observed_at(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterDateFilterArgs']]]]:
        return pulumi.get(self, "first_observed_at")

    @first_observed_at.setter
    def first_observed_at(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterDateFilterArgs']]]]):
        pulumi.set(self, "first_observed_at", value)

    @property
    @pulumi.getter(name="inspectorScore")
    def inspector_score(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterNumberFilterArgs']]]]:
        return pulumi.get(self, "inspector_score")

    @inspector_score.setter
    def inspector_score(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterNumberFilterArgs']]]]):
        pulumi.set(self, "inspector_score", value)

    @property
    @pulumi.getter(name="lastObservedAt")
    def last_observed_at(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterDateFilterArgs']]]]:
        return pulumi.get(self, "last_observed_at")

    @last_observed_at.setter
    def last_observed_at(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterDateFilterArgs']]]]):
        pulumi.set(self, "last_observed_at", value)

    @property
    @pulumi.getter(name="networkProtocol")
    def network_protocol(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "network_protocol")

    @network_protocol.setter
    def network_protocol(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "network_protocol", value)

    @property
    @pulumi.getter(name="portRange")
    def port_range(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterPortRangeFilterArgs']]]]:
        return pulumi.get(self, "port_range")

    @port_range.setter
    def port_range(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterPortRangeFilterArgs']]]]):
        pulumi.set(self, "port_range", value)

    @property
    @pulumi.getter(name="relatedVulnerabilities")
    def related_vulnerabilities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "related_vulnerabilities")

    @related_vulnerabilities.setter
    def related_vulnerabilities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "related_vulnerabilities", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceTags")
    def resource_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterMapFilterArgs']]]]:
        return pulumi.get(self, "resource_tags")

    @resource_tags.setter
    def resource_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterMapFilterArgs']]]]):
        pulumi.set(self, "resource_tags", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterDateFilterArgs']]]]:
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterDateFilterArgs']]]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="vendorSeverity")
    def vendor_severity(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "vendor_severity")

    @vendor_severity.setter
    def vendor_severity(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "vendor_severity", value)

    @property
    @pulumi.getter(name="vulnerabilityId")
    def vulnerability_id(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "vulnerability_id")

    @vulnerability_id.setter
    def vulnerability_id(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "vulnerability_id", value)

    @property
    @pulumi.getter(name="vulnerabilitySource")
    def vulnerability_source(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]:
        return pulumi.get(self, "vulnerability_source")

    @vulnerability_source.setter
    def vulnerability_source(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterStringFilterArgs']]]]):
        pulumi.set(self, "vulnerability_source", value)

    @property
    @pulumi.getter(name="vulnerablePackages")
    def vulnerable_packages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FilterPackageFilterArgs']]]]:
        return pulumi.get(self, "vulnerable_packages")

    @vulnerable_packages.setter
    def vulnerable_packages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FilterPackageFilterArgs']]]]):
        pulumi.set(self, "vulnerable_packages", value)


@pulumi.input_type
class FilterDateFilterArgs:
    def __init__(__self__, *,
                 end_inclusive: Optional[pulumi.Input[int]] = None,
                 start_inclusive: Optional[pulumi.Input[int]] = None):
        if end_inclusive is not None:
            pulumi.set(__self__, "end_inclusive", end_inclusive)
        if start_inclusive is not None:
            pulumi.set(__self__, "start_inclusive", start_inclusive)

    @property
    @pulumi.getter(name="endInclusive")
    def end_inclusive(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "end_inclusive")

    @end_inclusive.setter
    def end_inclusive(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "end_inclusive", value)

    @property
    @pulumi.getter(name="startInclusive")
    def start_inclusive(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "start_inclusive")

    @start_inclusive.setter
    def start_inclusive(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start_inclusive", value)


@pulumi.input_type
class FilterMapFilterArgs:
    def __init__(__self__, *,
                 comparison: pulumi.Input['FilterMapComparison'],
                 key: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "comparison", comparison)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def comparison(self) -> pulumi.Input['FilterMapComparison']:
        return pulumi.get(self, "comparison")

    @comparison.setter
    def comparison(self, value: pulumi.Input['FilterMapComparison']):
        pulumi.set(self, "comparison", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class FilterNumberFilterArgs:
    def __init__(__self__, *,
                 lower_inclusive: Optional[pulumi.Input[float]] = None,
                 upper_inclusive: Optional[pulumi.Input[float]] = None):
        if lower_inclusive is not None:
            pulumi.set(__self__, "lower_inclusive", lower_inclusive)
        if upper_inclusive is not None:
            pulumi.set(__self__, "upper_inclusive", upper_inclusive)

    @property
    @pulumi.getter(name="lowerInclusive")
    def lower_inclusive(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "lower_inclusive")

    @lower_inclusive.setter
    def lower_inclusive(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "lower_inclusive", value)

    @property
    @pulumi.getter(name="upperInclusive")
    def upper_inclusive(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "upper_inclusive")

    @upper_inclusive.setter
    def upper_inclusive(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "upper_inclusive", value)


@pulumi.input_type
class FilterPackageFilterArgs:
    def __init__(__self__, *,
                 architecture: Optional[pulumi.Input['FilterStringFilterArgs']] = None,
                 epoch: Optional[pulumi.Input['FilterNumberFilterArgs']] = None,
                 name: Optional[pulumi.Input['FilterStringFilterArgs']] = None,
                 release: Optional[pulumi.Input['FilterStringFilterArgs']] = None,
                 source_layer_hash: Optional[pulumi.Input['FilterStringFilterArgs']] = None,
                 version: Optional[pulumi.Input['FilterStringFilterArgs']] = None):
        if architecture is not None:
            pulumi.set(__self__, "architecture", architecture)
        if epoch is not None:
            pulumi.set(__self__, "epoch", epoch)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if release is not None:
            pulumi.set(__self__, "release", release)
        if source_layer_hash is not None:
            pulumi.set(__self__, "source_layer_hash", source_layer_hash)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def architecture(self) -> Optional[pulumi.Input['FilterStringFilterArgs']]:
        return pulumi.get(self, "architecture")

    @architecture.setter
    def architecture(self, value: Optional[pulumi.Input['FilterStringFilterArgs']]):
        pulumi.set(self, "architecture", value)

    @property
    @pulumi.getter
    def epoch(self) -> Optional[pulumi.Input['FilterNumberFilterArgs']]:
        return pulumi.get(self, "epoch")

    @epoch.setter
    def epoch(self, value: Optional[pulumi.Input['FilterNumberFilterArgs']]):
        pulumi.set(self, "epoch", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input['FilterStringFilterArgs']]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input['FilterStringFilterArgs']]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def release(self) -> Optional[pulumi.Input['FilterStringFilterArgs']]:
        return pulumi.get(self, "release")

    @release.setter
    def release(self, value: Optional[pulumi.Input['FilterStringFilterArgs']]):
        pulumi.set(self, "release", value)

    @property
    @pulumi.getter(name="sourceLayerHash")
    def source_layer_hash(self) -> Optional[pulumi.Input['FilterStringFilterArgs']]:
        return pulumi.get(self, "source_layer_hash")

    @source_layer_hash.setter
    def source_layer_hash(self, value: Optional[pulumi.Input['FilterStringFilterArgs']]):
        pulumi.set(self, "source_layer_hash", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input['FilterStringFilterArgs']]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input['FilterStringFilterArgs']]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class FilterPortRangeFilterArgs:
    def __init__(__self__, *,
                 begin_inclusive: Optional[pulumi.Input[int]] = None,
                 end_inclusive: Optional[pulumi.Input[int]] = None):
        if begin_inclusive is not None:
            pulumi.set(__self__, "begin_inclusive", begin_inclusive)
        if end_inclusive is not None:
            pulumi.set(__self__, "end_inclusive", end_inclusive)

    @property
    @pulumi.getter(name="beginInclusive")
    def begin_inclusive(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "begin_inclusive")

    @begin_inclusive.setter
    def begin_inclusive(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "begin_inclusive", value)

    @property
    @pulumi.getter(name="endInclusive")
    def end_inclusive(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "end_inclusive")

    @end_inclusive.setter
    def end_inclusive(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "end_inclusive", value)


@pulumi.input_type
class FilterStringFilterArgs:
    def __init__(__self__, *,
                 comparison: pulumi.Input['FilterStringComparison'],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "comparison", comparison)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def comparison(self) -> pulumi.Input['FilterStringComparison']:
        return pulumi.get(self, "comparison")

    @comparison.setter
    def comparison(self, value: pulumi.Input['FilterStringComparison']):
        pulumi.set(self, "comparison", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

