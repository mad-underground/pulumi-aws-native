# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'BudgetBudgetDataArgs',
    'BudgetCostTypesArgs',
    'BudgetNotificationArgs',
    'BudgetNotificationWithSubscribersArgs',
    'BudgetPropertiesArgs',
    'BudgetSpendArgs',
    'BudgetSubscriberArgs',
    'BudgetTimePeriodArgs',
]

@pulumi.input_type
class BudgetBudgetDataArgs:
    def __init__(__self__, *,
                 budget_type: pulumi.Input[str],
                 time_unit: pulumi.Input[str],
                 budget_limit: Optional[pulumi.Input['BudgetSpendArgs']] = None,
                 budget_name: Optional[pulumi.Input[str]] = None,
                 cost_filters: Optional[pulumi.Input[Union[Any, str]]] = None,
                 cost_types: Optional[pulumi.Input['BudgetCostTypesArgs']] = None,
                 planned_budget_limits: Optional[pulumi.Input[Union[Any, str]]] = None,
                 time_period: Optional[pulumi.Input['BudgetTimePeriodArgs']] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html
        :param pulumi.Input[str] budget_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-budgettype
        :param pulumi.Input[str] time_unit: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-timeunit
        :param pulumi.Input['BudgetSpendArgs'] budget_limit: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-budgetlimit
        :param pulumi.Input[str] budget_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-budgetname
        :param pulumi.Input[Union[Any, str]] cost_filters: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-costfilters
        :param pulumi.Input['BudgetCostTypesArgs'] cost_types: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-costtypes
        :param pulumi.Input[Union[Any, str]] planned_budget_limits: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-plannedbudgetlimits
        :param pulumi.Input['BudgetTimePeriodArgs'] time_period: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-timeperiod
        """
        pulumi.set(__self__, "budget_type", budget_type)
        pulumi.set(__self__, "time_unit", time_unit)
        if budget_limit is not None:
            pulumi.set(__self__, "budget_limit", budget_limit)
        if budget_name is not None:
            pulumi.set(__self__, "budget_name", budget_name)
        if cost_filters is not None:
            pulumi.set(__self__, "cost_filters", cost_filters)
        if cost_types is not None:
            pulumi.set(__self__, "cost_types", cost_types)
        if planned_budget_limits is not None:
            pulumi.set(__self__, "planned_budget_limits", planned_budget_limits)
        if time_period is not None:
            pulumi.set(__self__, "time_period", time_period)

    @property
    @pulumi.getter(name="BudgetType")
    def budget_type(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-budgettype
        """
        return pulumi.get(self, "budget_type")

    @budget_type.setter
    def budget_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "budget_type", value)

    @property
    @pulumi.getter(name="TimeUnit")
    def time_unit(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-timeunit
        """
        return pulumi.get(self, "time_unit")

    @time_unit.setter
    def time_unit(self, value: pulumi.Input[str]):
        pulumi.set(self, "time_unit", value)

    @property
    @pulumi.getter(name="BudgetLimit")
    def budget_limit(self) -> Optional[pulumi.Input['BudgetSpendArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-budgetlimit
        """
        return pulumi.get(self, "budget_limit")

    @budget_limit.setter
    def budget_limit(self, value: Optional[pulumi.Input['BudgetSpendArgs']]):
        pulumi.set(self, "budget_limit", value)

    @property
    @pulumi.getter(name="BudgetName")
    def budget_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-budgetname
        """
        return pulumi.get(self, "budget_name")

    @budget_name.setter
    def budget_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "budget_name", value)

    @property
    @pulumi.getter(name="CostFilters")
    def cost_filters(self) -> Optional[pulumi.Input[Union[Any, str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-costfilters
        """
        return pulumi.get(self, "cost_filters")

    @cost_filters.setter
    def cost_filters(self, value: Optional[pulumi.Input[Union[Any, str]]]):
        pulumi.set(self, "cost_filters", value)

    @property
    @pulumi.getter(name="CostTypes")
    def cost_types(self) -> Optional[pulumi.Input['BudgetCostTypesArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-costtypes
        """
        return pulumi.get(self, "cost_types")

    @cost_types.setter
    def cost_types(self, value: Optional[pulumi.Input['BudgetCostTypesArgs']]):
        pulumi.set(self, "cost_types", value)

    @property
    @pulumi.getter(name="PlannedBudgetLimits")
    def planned_budget_limits(self) -> Optional[pulumi.Input[Union[Any, str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-plannedbudgetlimits
        """
        return pulumi.get(self, "planned_budget_limits")

    @planned_budget_limits.setter
    def planned_budget_limits(self, value: Optional[pulumi.Input[Union[Any, str]]]):
        pulumi.set(self, "planned_budget_limits", value)

    @property
    @pulumi.getter(name="TimePeriod")
    def time_period(self) -> Optional[pulumi.Input['BudgetTimePeriodArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-timeperiod
        """
        return pulumi.get(self, "time_period")

    @time_period.setter
    def time_period(self, value: Optional[pulumi.Input['BudgetTimePeriodArgs']]):
        pulumi.set(self, "time_period", value)


@pulumi.input_type
class BudgetCostTypesArgs:
    def __init__(__self__, *,
                 include_credit: Optional[pulumi.Input[bool]] = None,
                 include_discount: Optional[pulumi.Input[bool]] = None,
                 include_other_subscription: Optional[pulumi.Input[bool]] = None,
                 include_recurring: Optional[pulumi.Input[bool]] = None,
                 include_refund: Optional[pulumi.Input[bool]] = None,
                 include_subscription: Optional[pulumi.Input[bool]] = None,
                 include_support: Optional[pulumi.Input[bool]] = None,
                 include_tax: Optional[pulumi.Input[bool]] = None,
                 include_upfront: Optional[pulumi.Input[bool]] = None,
                 use_amortized: Optional[pulumi.Input[bool]] = None,
                 use_blended: Optional[pulumi.Input[bool]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html
        :param pulumi.Input[bool] include_credit: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includecredit
        :param pulumi.Input[bool] include_discount: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includediscount
        :param pulumi.Input[bool] include_other_subscription: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includeothersubscription
        :param pulumi.Input[bool] include_recurring: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includerecurring
        :param pulumi.Input[bool] include_refund: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includerefund
        :param pulumi.Input[bool] include_subscription: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includesubscription
        :param pulumi.Input[bool] include_support: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includesupport
        :param pulumi.Input[bool] include_tax: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includetax
        :param pulumi.Input[bool] include_upfront: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includeupfront
        :param pulumi.Input[bool] use_amortized: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-useamortized
        :param pulumi.Input[bool] use_blended: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-useblended
        """
        if include_credit is not None:
            pulumi.set(__self__, "include_credit", include_credit)
        if include_discount is not None:
            pulumi.set(__self__, "include_discount", include_discount)
        if include_other_subscription is not None:
            pulumi.set(__self__, "include_other_subscription", include_other_subscription)
        if include_recurring is not None:
            pulumi.set(__self__, "include_recurring", include_recurring)
        if include_refund is not None:
            pulumi.set(__self__, "include_refund", include_refund)
        if include_subscription is not None:
            pulumi.set(__self__, "include_subscription", include_subscription)
        if include_support is not None:
            pulumi.set(__self__, "include_support", include_support)
        if include_tax is not None:
            pulumi.set(__self__, "include_tax", include_tax)
        if include_upfront is not None:
            pulumi.set(__self__, "include_upfront", include_upfront)
        if use_amortized is not None:
            pulumi.set(__self__, "use_amortized", use_amortized)
        if use_blended is not None:
            pulumi.set(__self__, "use_blended", use_blended)

    @property
    @pulumi.getter(name="IncludeCredit")
    def include_credit(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includecredit
        """
        return pulumi.get(self, "include_credit")

    @include_credit.setter
    def include_credit(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_credit", value)

    @property
    @pulumi.getter(name="IncludeDiscount")
    def include_discount(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includediscount
        """
        return pulumi.get(self, "include_discount")

    @include_discount.setter
    def include_discount(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_discount", value)

    @property
    @pulumi.getter(name="IncludeOtherSubscription")
    def include_other_subscription(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includeothersubscription
        """
        return pulumi.get(self, "include_other_subscription")

    @include_other_subscription.setter
    def include_other_subscription(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_other_subscription", value)

    @property
    @pulumi.getter(name="IncludeRecurring")
    def include_recurring(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includerecurring
        """
        return pulumi.get(self, "include_recurring")

    @include_recurring.setter
    def include_recurring(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_recurring", value)

    @property
    @pulumi.getter(name="IncludeRefund")
    def include_refund(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includerefund
        """
        return pulumi.get(self, "include_refund")

    @include_refund.setter
    def include_refund(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_refund", value)

    @property
    @pulumi.getter(name="IncludeSubscription")
    def include_subscription(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includesubscription
        """
        return pulumi.get(self, "include_subscription")

    @include_subscription.setter
    def include_subscription(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_subscription", value)

    @property
    @pulumi.getter(name="IncludeSupport")
    def include_support(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includesupport
        """
        return pulumi.get(self, "include_support")

    @include_support.setter
    def include_support(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_support", value)

    @property
    @pulumi.getter(name="IncludeTax")
    def include_tax(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includetax
        """
        return pulumi.get(self, "include_tax")

    @include_tax.setter
    def include_tax(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_tax", value)

    @property
    @pulumi.getter(name="IncludeUpfront")
    def include_upfront(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includeupfront
        """
        return pulumi.get(self, "include_upfront")

    @include_upfront.setter
    def include_upfront(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_upfront", value)

    @property
    @pulumi.getter(name="UseAmortized")
    def use_amortized(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-useamortized
        """
        return pulumi.get(self, "use_amortized")

    @use_amortized.setter
    def use_amortized(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_amortized", value)

    @property
    @pulumi.getter(name="UseBlended")
    def use_blended(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-useblended
        """
        return pulumi.get(self, "use_blended")

    @use_blended.setter
    def use_blended(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_blended", value)


@pulumi.input_type
class BudgetNotificationArgs:
    def __init__(__self__, *,
                 comparison_operator: pulumi.Input[str],
                 notification_type: pulumi.Input[str],
                 threshold: pulumi.Input[float],
                 threshold_type: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html
        :param pulumi.Input[str] comparison_operator: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-comparisonoperator
        :param pulumi.Input[str] notification_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-notificationtype
        :param pulumi.Input[float] threshold: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-threshold
        :param pulumi.Input[str] threshold_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-thresholdtype
        """
        pulumi.set(__self__, "comparison_operator", comparison_operator)
        pulumi.set(__self__, "notification_type", notification_type)
        pulumi.set(__self__, "threshold", threshold)
        if threshold_type is not None:
            pulumi.set(__self__, "threshold_type", threshold_type)

    @property
    @pulumi.getter(name="ComparisonOperator")
    def comparison_operator(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-comparisonoperator
        """
        return pulumi.get(self, "comparison_operator")

    @comparison_operator.setter
    def comparison_operator(self, value: pulumi.Input[str]):
        pulumi.set(self, "comparison_operator", value)

    @property
    @pulumi.getter(name="NotificationType")
    def notification_type(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-notificationtype
        """
        return pulumi.get(self, "notification_type")

    @notification_type.setter
    def notification_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "notification_type", value)

    @property
    @pulumi.getter(name="Threshold")
    def threshold(self) -> pulumi.Input[float]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-threshold
        """
        return pulumi.get(self, "threshold")

    @threshold.setter
    def threshold(self, value: pulumi.Input[float]):
        pulumi.set(self, "threshold", value)

    @property
    @pulumi.getter(name="ThresholdType")
    def threshold_type(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-thresholdtype
        """
        return pulumi.get(self, "threshold_type")

    @threshold_type.setter
    def threshold_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "threshold_type", value)


@pulumi.input_type
class BudgetNotificationWithSubscribersArgs:
    def __init__(__self__, *,
                 notification: pulumi.Input['BudgetNotificationArgs'],
                 subscribers: pulumi.Input[Sequence[pulumi.Input['BudgetSubscriberArgs']]]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notificationwithsubscribers.html
        :param pulumi.Input['BudgetNotificationArgs'] notification: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notificationwithsubscribers.html#cfn-budgets-budget-notificationwithsubscribers-notification
        :param pulumi.Input[Sequence[pulumi.Input['BudgetSubscriberArgs']]] subscribers: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notificationwithsubscribers.html#cfn-budgets-budget-notificationwithsubscribers-subscribers
        """
        pulumi.set(__self__, "notification", notification)
        pulumi.set(__self__, "subscribers", subscribers)

    @property
    @pulumi.getter(name="Notification")
    def notification(self) -> pulumi.Input['BudgetNotificationArgs']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notificationwithsubscribers.html#cfn-budgets-budget-notificationwithsubscribers-notification
        """
        return pulumi.get(self, "notification")

    @notification.setter
    def notification(self, value: pulumi.Input['BudgetNotificationArgs']):
        pulumi.set(self, "notification", value)

    @property
    @pulumi.getter(name="Subscribers")
    def subscribers(self) -> pulumi.Input[Sequence[pulumi.Input['BudgetSubscriberArgs']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notificationwithsubscribers.html#cfn-budgets-budget-notificationwithsubscribers-subscribers
        """
        return pulumi.get(self, "subscribers")

    @subscribers.setter
    def subscribers(self, value: pulumi.Input[Sequence[pulumi.Input['BudgetSubscriberArgs']]]):
        pulumi.set(self, "subscribers", value)


@pulumi.input_type
class BudgetPropertiesArgs:
    def __init__(__self__, *,
                 budget: pulumi.Input['BudgetBudgetDataArgs'],
                 notifications_with_subscribers: Optional[pulumi.Input[Sequence[pulumi.Input['BudgetNotificationWithSubscribersArgs']]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html
        :param pulumi.Input['BudgetBudgetDataArgs'] budget: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-budget
        :param pulumi.Input[Sequence[pulumi.Input['BudgetNotificationWithSubscribersArgs']]] notifications_with_subscribers: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-notificationswithsubscribers
        """
        pulumi.set(__self__, "budget", budget)
        if notifications_with_subscribers is not None:
            pulumi.set(__self__, "notifications_with_subscribers", notifications_with_subscribers)

    @property
    @pulumi.getter(name="Budget")
    def budget(self) -> pulumi.Input['BudgetBudgetDataArgs']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-budget
        """
        return pulumi.get(self, "budget")

    @budget.setter
    def budget(self, value: pulumi.Input['BudgetBudgetDataArgs']):
        pulumi.set(self, "budget", value)

    @property
    @pulumi.getter(name="NotificationsWithSubscribers")
    def notifications_with_subscribers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BudgetNotificationWithSubscribersArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-notificationswithsubscribers
        """
        return pulumi.get(self, "notifications_with_subscribers")

    @notifications_with_subscribers.setter
    def notifications_with_subscribers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BudgetNotificationWithSubscribersArgs']]]]):
        pulumi.set(self, "notifications_with_subscribers", value)


@pulumi.input_type
class BudgetSpendArgs:
    def __init__(__self__, *,
                 amount: pulumi.Input[float],
                 unit: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-spend.html
        :param pulumi.Input[float] amount: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-spend.html#cfn-budgets-budget-spend-amount
        :param pulumi.Input[str] unit: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-spend.html#cfn-budgets-budget-spend-unit
        """
        pulumi.set(__self__, "amount", amount)
        pulumi.set(__self__, "unit", unit)

    @property
    @pulumi.getter(name="Amount")
    def amount(self) -> pulumi.Input[float]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-spend.html#cfn-budgets-budget-spend-amount
        """
        return pulumi.get(self, "amount")

    @amount.setter
    def amount(self, value: pulumi.Input[float]):
        pulumi.set(self, "amount", value)

    @property
    @pulumi.getter(name="Unit")
    def unit(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-spend.html#cfn-budgets-budget-spend-unit
        """
        return pulumi.get(self, "unit")

    @unit.setter
    def unit(self, value: pulumi.Input[str]):
        pulumi.set(self, "unit", value)


@pulumi.input_type
class BudgetSubscriberArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 subscription_type: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-subscriber.html
        :param pulumi.Input[str] address: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-subscriber.html#cfn-budgets-budget-subscriber-address
        :param pulumi.Input[str] subscription_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-subscriber.html#cfn-budgets-budget-subscriber-subscriptiontype
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "subscription_type", subscription_type)

    @property
    @pulumi.getter(name="Address")
    def address(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-subscriber.html#cfn-budgets-budget-subscriber-address
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="SubscriptionType")
    def subscription_type(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-subscriber.html#cfn-budgets-budget-subscriber-subscriptiontype
        """
        return pulumi.get(self, "subscription_type")

    @subscription_type.setter
    def subscription_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "subscription_type", value)


@pulumi.input_type
class BudgetTimePeriodArgs:
    def __init__(__self__, *,
                 end: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-timeperiod.html
        :param pulumi.Input[str] end: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-timeperiod.html#cfn-budgets-budget-timeperiod-end
        :param pulumi.Input[str] start: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-timeperiod.html#cfn-budgets-budget-timeperiod-start
        """
        if end is not None:
            pulumi.set(__self__, "end", end)
        if start is not None:
            pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter(name="End")
    def end(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-timeperiod.html#cfn-budgets-budget-timeperiod-end
        """
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter(name="Start")
    def start(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-timeperiod.html#cfn-budgets-budget-timeperiod-start
        """
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start", value)

