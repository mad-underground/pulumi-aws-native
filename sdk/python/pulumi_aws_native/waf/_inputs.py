# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ByteMatchSetByteMatchTupleArgs',
    'ByteMatchSetFieldToMatchArgs',
    'IPSetIPSetDescriptorArgs',
    'RulePredicateArgs',
    'SizeConstraintSetFieldToMatchArgs',
    'SizeConstraintSetSizeConstraintArgs',
    'SqlInjectionMatchSetFieldToMatchArgs',
    'SqlInjectionMatchSetSqlInjectionMatchTupleArgs',
    'WebACLActivatedRuleArgs',
    'WebACLWafActionArgs',
    'XssMatchSetFieldToMatchArgs',
    'XssMatchSetXssMatchTupleArgs',
]

@pulumi.input_type
class ByteMatchSetByteMatchTupleArgs:
    def __init__(__self__, *,
                 field_to_match: pulumi.Input['ByteMatchSetFieldToMatchArgs'],
                 positional_constraint: pulumi.Input[str],
                 text_transformation: pulumi.Input[str],
                 target_string: Optional[pulumi.Input[str]] = None,
                 target_string_base64: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "field_to_match", field_to_match)
        pulumi.set(__self__, "positional_constraint", positional_constraint)
        pulumi.set(__self__, "text_transformation", text_transformation)
        if target_string is not None:
            pulumi.set(__self__, "target_string", target_string)
        if target_string_base64 is not None:
            pulumi.set(__self__, "target_string_base64", target_string_base64)

    @property
    @pulumi.getter(name="fieldToMatch")
    def field_to_match(self) -> pulumi.Input['ByteMatchSetFieldToMatchArgs']:
        return pulumi.get(self, "field_to_match")

    @field_to_match.setter
    def field_to_match(self, value: pulumi.Input['ByteMatchSetFieldToMatchArgs']):
        pulumi.set(self, "field_to_match", value)

    @property
    @pulumi.getter(name="positionalConstraint")
    def positional_constraint(self) -> pulumi.Input[str]:
        return pulumi.get(self, "positional_constraint")

    @positional_constraint.setter
    def positional_constraint(self, value: pulumi.Input[str]):
        pulumi.set(self, "positional_constraint", value)

    @property
    @pulumi.getter(name="textTransformation")
    def text_transformation(self) -> pulumi.Input[str]:
        return pulumi.get(self, "text_transformation")

    @text_transformation.setter
    def text_transformation(self, value: pulumi.Input[str]):
        pulumi.set(self, "text_transformation", value)

    @property
    @pulumi.getter(name="targetString")
    def target_string(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "target_string")

    @target_string.setter
    def target_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_string", value)

    @property
    @pulumi.getter(name="targetStringBase64")
    def target_string_base64(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "target_string_base64")

    @target_string_base64.setter
    def target_string_base64(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_string_base64", value)


@pulumi.input_type
class ByteMatchSetFieldToMatchArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 data: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "type", type)
        if data is not None:
            pulumi.set(__self__, "data", data)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)


@pulumi.input_type
class IPSetIPSetDescriptorArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class RulePredicateArgs:
    def __init__(__self__, *,
                 data_id: pulumi.Input[str],
                 negated: pulumi.Input[bool],
                 type: pulumi.Input[str]):
        pulumi.set(__self__, "data_id", data_id)
        pulumi.set(__self__, "negated", negated)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="dataId")
    def data_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "data_id")

    @data_id.setter
    def data_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_id", value)

    @property
    @pulumi.getter
    def negated(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "negated")

    @negated.setter
    def negated(self, value: pulumi.Input[bool]):
        pulumi.set(self, "negated", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class SizeConstraintSetFieldToMatchArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 data: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "type", type)
        if data is not None:
            pulumi.set(__self__, "data", data)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)


@pulumi.input_type
class SizeConstraintSetSizeConstraintArgs:
    def __init__(__self__, *,
                 comparison_operator: pulumi.Input[str],
                 field_to_match: pulumi.Input['SizeConstraintSetFieldToMatchArgs'],
                 size: pulumi.Input[int],
                 text_transformation: pulumi.Input[str]):
        pulumi.set(__self__, "comparison_operator", comparison_operator)
        pulumi.set(__self__, "field_to_match", field_to_match)
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "text_transformation", text_transformation)

    @property
    @pulumi.getter(name="comparisonOperator")
    def comparison_operator(self) -> pulumi.Input[str]:
        return pulumi.get(self, "comparison_operator")

    @comparison_operator.setter
    def comparison_operator(self, value: pulumi.Input[str]):
        pulumi.set(self, "comparison_operator", value)

    @property
    @pulumi.getter(name="fieldToMatch")
    def field_to_match(self) -> pulumi.Input['SizeConstraintSetFieldToMatchArgs']:
        return pulumi.get(self, "field_to_match")

    @field_to_match.setter
    def field_to_match(self, value: pulumi.Input['SizeConstraintSetFieldToMatchArgs']):
        pulumi.set(self, "field_to_match", value)

    @property
    @pulumi.getter
    def size(self) -> pulumi.Input[int]:
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: pulumi.Input[int]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="textTransformation")
    def text_transformation(self) -> pulumi.Input[str]:
        return pulumi.get(self, "text_transformation")

    @text_transformation.setter
    def text_transformation(self, value: pulumi.Input[str]):
        pulumi.set(self, "text_transformation", value)


@pulumi.input_type
class SqlInjectionMatchSetFieldToMatchArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 data: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "type", type)
        if data is not None:
            pulumi.set(__self__, "data", data)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)


@pulumi.input_type
class SqlInjectionMatchSetSqlInjectionMatchTupleArgs:
    def __init__(__self__, *,
                 field_to_match: pulumi.Input['SqlInjectionMatchSetFieldToMatchArgs'],
                 text_transformation: pulumi.Input[str]):
        pulumi.set(__self__, "field_to_match", field_to_match)
        pulumi.set(__self__, "text_transformation", text_transformation)

    @property
    @pulumi.getter(name="fieldToMatch")
    def field_to_match(self) -> pulumi.Input['SqlInjectionMatchSetFieldToMatchArgs']:
        return pulumi.get(self, "field_to_match")

    @field_to_match.setter
    def field_to_match(self, value: pulumi.Input['SqlInjectionMatchSetFieldToMatchArgs']):
        pulumi.set(self, "field_to_match", value)

    @property
    @pulumi.getter(name="textTransformation")
    def text_transformation(self) -> pulumi.Input[str]:
        return pulumi.get(self, "text_transformation")

    @text_transformation.setter
    def text_transformation(self, value: pulumi.Input[str]):
        pulumi.set(self, "text_transformation", value)


@pulumi.input_type
class WebACLActivatedRuleArgs:
    def __init__(__self__, *,
                 priority: pulumi.Input[int],
                 rule_id: pulumi.Input[str],
                 action: Optional[pulumi.Input['WebACLWafActionArgs']] = None):
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "rule_id", rule_id)
        if action is not None:
            pulumi.set(__self__, "action", action)

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[int]:
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: pulumi.Input[int]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input['WebACLWafActionArgs']]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input['WebACLWafActionArgs']]):
        pulumi.set(self, "action", value)


@pulumi.input_type
class WebACLWafActionArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str]):
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class XssMatchSetFieldToMatchArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 data: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "type", type)
        if data is not None:
            pulumi.set(__self__, "data", data)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)


@pulumi.input_type
class XssMatchSetXssMatchTupleArgs:
    def __init__(__self__, *,
                 field_to_match: pulumi.Input['XssMatchSetFieldToMatchArgs'],
                 text_transformation: pulumi.Input[str]):
        pulumi.set(__self__, "field_to_match", field_to_match)
        pulumi.set(__self__, "text_transformation", text_transformation)

    @property
    @pulumi.getter(name="fieldToMatch")
    def field_to_match(self) -> pulumi.Input['XssMatchSetFieldToMatchArgs']:
        return pulumi.get(self, "field_to_match")

    @field_to_match.setter
    def field_to_match(self, value: pulumi.Input['XssMatchSetFieldToMatchArgs']):
        pulumi.set(self, "field_to_match", value)

    @property
    @pulumi.getter(name="textTransformation")
    def text_transformation(self) -> pulumi.Input[str]:
        return pulumi.get(self, "text_transformation")

    @text_transformation.setter
    def text_transformation(self, value: pulumi.Input[str]):
        pulumi.set(self, "text_transformation", value)

