# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'CloudFormationProvisionedProductProvisioningParameter',
    'CloudFormationProvisionedProductProvisioningPreferences',
]

@pulumi.output_type
class CloudFormationProvisionedProductProvisioningParameter(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningparameter.html
    """
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningparameter.html
        :param str key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningparameter.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningparameter-key
        :param str value: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningparameter.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningparameter-value
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="Key")
    def key(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningparameter.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningparameter-key
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="Value")
    def value(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningparameter.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningparameter-value
        """
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CloudFormationProvisionedProductProvisioningPreferences(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html
    """
    def __init__(__self__, *,
                 stack_set_accounts: Optional[Sequence[str]] = None,
                 stack_set_failure_tolerance_count: Optional[int] = None,
                 stack_set_failure_tolerance_percentage: Optional[int] = None,
                 stack_set_max_concurrency_count: Optional[int] = None,
                 stack_set_max_concurrency_percentage: Optional[int] = None,
                 stack_set_operation_type: Optional[str] = None,
                 stack_set_regions: Optional[Sequence[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html
        :param Sequence[str] stack_set_accounts: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences-stacksetaccounts
        :param int stack_set_failure_tolerance_count: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences-stacksetfailuretolerancecount
        :param int stack_set_failure_tolerance_percentage: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences-stacksetfailuretolerancepercentage
        :param int stack_set_max_concurrency_count: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences-stacksetmaxconcurrencycount
        :param int stack_set_max_concurrency_percentage: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences-stacksetmaxconcurrencypercentage
        :param str stack_set_operation_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences-stacksetoperationtype
        :param Sequence[str] stack_set_regions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences-stacksetregions
        """
        if stack_set_accounts is not None:
            pulumi.set(__self__, "stack_set_accounts", stack_set_accounts)
        if stack_set_failure_tolerance_count is not None:
            pulumi.set(__self__, "stack_set_failure_tolerance_count", stack_set_failure_tolerance_count)
        if stack_set_failure_tolerance_percentage is not None:
            pulumi.set(__self__, "stack_set_failure_tolerance_percentage", stack_set_failure_tolerance_percentage)
        if stack_set_max_concurrency_count is not None:
            pulumi.set(__self__, "stack_set_max_concurrency_count", stack_set_max_concurrency_count)
        if stack_set_max_concurrency_percentage is not None:
            pulumi.set(__self__, "stack_set_max_concurrency_percentage", stack_set_max_concurrency_percentage)
        if stack_set_operation_type is not None:
            pulumi.set(__self__, "stack_set_operation_type", stack_set_operation_type)
        if stack_set_regions is not None:
            pulumi.set(__self__, "stack_set_regions", stack_set_regions)

    @property
    @pulumi.getter(name="StackSetAccounts")
    def stack_set_accounts(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences-stacksetaccounts
        """
        return pulumi.get(self, "stack_set_accounts")

    @property
    @pulumi.getter(name="StackSetFailureToleranceCount")
    def stack_set_failure_tolerance_count(self) -> Optional[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences-stacksetfailuretolerancecount
        """
        return pulumi.get(self, "stack_set_failure_tolerance_count")

    @property
    @pulumi.getter(name="StackSetFailureTolerancePercentage")
    def stack_set_failure_tolerance_percentage(self) -> Optional[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences-stacksetfailuretolerancepercentage
        """
        return pulumi.get(self, "stack_set_failure_tolerance_percentage")

    @property
    @pulumi.getter(name="StackSetMaxConcurrencyCount")
    def stack_set_max_concurrency_count(self) -> Optional[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences-stacksetmaxconcurrencycount
        """
        return pulumi.get(self, "stack_set_max_concurrency_count")

    @property
    @pulumi.getter(name="StackSetMaxConcurrencyPercentage")
    def stack_set_max_concurrency_percentage(self) -> Optional[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences-stacksetmaxconcurrencypercentage
        """
        return pulumi.get(self, "stack_set_max_concurrency_percentage")

    @property
    @pulumi.getter(name="StackSetOperationType")
    def stack_set_operation_type(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences-stacksetoperationtype
        """
        return pulumi.get(self, "stack_set_operation_type")

    @property
    @pulumi.getter(name="StackSetRegions")
    def stack_set_regions(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences-stacksetregions
        """
        return pulumi.get(self, "stack_set_regions")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


