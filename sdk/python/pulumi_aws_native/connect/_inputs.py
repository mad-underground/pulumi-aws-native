# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'QuickConnectPhoneNumberQuickConnectConfigArgs',
    'QuickConnectQueueQuickConnectConfigArgs',
    'QuickConnectQuickConnectConfigArgs',
    'QuickConnectUserQuickConnectConfigArgs',
]

@pulumi.input_type
class QuickConnectPhoneNumberQuickConnectConfigArgs:
    def __init__(__self__, *,
                 phone_number: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-phonenumberquickconnectconfig.html
        :param pulumi.Input[str] phone_number: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-phonenumberquickconnectconfig.html#cfn-connect-quickconnect-phonenumberquickconnectconfig-phonenumber
        """
        pulumi.set(__self__, "phone_number", phone_number)

    @property
    @pulumi.getter(name="phoneNumber")
    def phone_number(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-phonenumberquickconnectconfig.html#cfn-connect-quickconnect-phonenumberquickconnectconfig-phonenumber
        """
        return pulumi.get(self, "phone_number")

    @phone_number.setter
    def phone_number(self, value: pulumi.Input[str]):
        pulumi.set(self, "phone_number", value)


@pulumi.input_type
class QuickConnectQueueQuickConnectConfigArgs:
    def __init__(__self__, *,
                 contact_flow_arn: pulumi.Input[str],
                 queue_arn: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-queuequickconnectconfig.html
        :param pulumi.Input[str] contact_flow_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-queuequickconnectconfig.html#cfn-connect-quickconnect-queuequickconnectconfig-contactflowarn
        :param pulumi.Input[str] queue_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-queuequickconnectconfig.html#cfn-connect-quickconnect-queuequickconnectconfig-queuearn
        """
        pulumi.set(__self__, "contact_flow_arn", contact_flow_arn)
        pulumi.set(__self__, "queue_arn", queue_arn)

    @property
    @pulumi.getter(name="contactFlowArn")
    def contact_flow_arn(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-queuequickconnectconfig.html#cfn-connect-quickconnect-queuequickconnectconfig-contactflowarn
        """
        return pulumi.get(self, "contact_flow_arn")

    @contact_flow_arn.setter
    def contact_flow_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "contact_flow_arn", value)

    @property
    @pulumi.getter(name="queueArn")
    def queue_arn(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-queuequickconnectconfig.html#cfn-connect-quickconnect-queuequickconnectconfig-queuearn
        """
        return pulumi.get(self, "queue_arn")

    @queue_arn.setter
    def queue_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "queue_arn", value)


@pulumi.input_type
class QuickConnectQuickConnectConfigArgs:
    def __init__(__self__, *,
                 quick_connect_type: pulumi.Input[str],
                 phone_config: Optional[pulumi.Input['QuickConnectPhoneNumberQuickConnectConfigArgs']] = None,
                 queue_config: Optional[pulumi.Input['QuickConnectQueueQuickConnectConfigArgs']] = None,
                 user_config: Optional[pulumi.Input['QuickConnectUserQuickConnectConfigArgs']] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html
        :param pulumi.Input[str] quick_connect_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-quickconnecttype
        :param pulumi.Input['QuickConnectPhoneNumberQuickConnectConfigArgs'] phone_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-phoneconfig
        :param pulumi.Input['QuickConnectQueueQuickConnectConfigArgs'] queue_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-queueconfig
        :param pulumi.Input['QuickConnectUserQuickConnectConfigArgs'] user_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-userconfig
        """
        pulumi.set(__self__, "quick_connect_type", quick_connect_type)
        if phone_config is not None:
            pulumi.set(__self__, "phone_config", phone_config)
        if queue_config is not None:
            pulumi.set(__self__, "queue_config", queue_config)
        if user_config is not None:
            pulumi.set(__self__, "user_config", user_config)

    @property
    @pulumi.getter(name="quickConnectType")
    def quick_connect_type(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-quickconnecttype
        """
        return pulumi.get(self, "quick_connect_type")

    @quick_connect_type.setter
    def quick_connect_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "quick_connect_type", value)

    @property
    @pulumi.getter(name="phoneConfig")
    def phone_config(self) -> Optional[pulumi.Input['QuickConnectPhoneNumberQuickConnectConfigArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-phoneconfig
        """
        return pulumi.get(self, "phone_config")

    @phone_config.setter
    def phone_config(self, value: Optional[pulumi.Input['QuickConnectPhoneNumberQuickConnectConfigArgs']]):
        pulumi.set(self, "phone_config", value)

    @property
    @pulumi.getter(name="queueConfig")
    def queue_config(self) -> Optional[pulumi.Input['QuickConnectQueueQuickConnectConfigArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-queueconfig
        """
        return pulumi.get(self, "queue_config")

    @queue_config.setter
    def queue_config(self, value: Optional[pulumi.Input['QuickConnectQueueQuickConnectConfigArgs']]):
        pulumi.set(self, "queue_config", value)

    @property
    @pulumi.getter(name="userConfig")
    def user_config(self) -> Optional[pulumi.Input['QuickConnectUserQuickConnectConfigArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-userconfig
        """
        return pulumi.get(self, "user_config")

    @user_config.setter
    def user_config(self, value: Optional[pulumi.Input['QuickConnectUserQuickConnectConfigArgs']]):
        pulumi.set(self, "user_config", value)


@pulumi.input_type
class QuickConnectUserQuickConnectConfigArgs:
    def __init__(__self__, *,
                 contact_flow_arn: pulumi.Input[str],
                 user_arn: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-userquickconnectconfig.html
        :param pulumi.Input[str] contact_flow_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-userquickconnectconfig.html#cfn-connect-quickconnect-userquickconnectconfig-contactflowarn
        :param pulumi.Input[str] user_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-userquickconnectconfig.html#cfn-connect-quickconnect-userquickconnectconfig-userarn
        """
        pulumi.set(__self__, "contact_flow_arn", contact_flow_arn)
        pulumi.set(__self__, "user_arn", user_arn)

    @property
    @pulumi.getter(name="contactFlowArn")
    def contact_flow_arn(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-userquickconnectconfig.html#cfn-connect-quickconnect-userquickconnectconfig-contactflowarn
        """
        return pulumi.get(self, "contact_flow_arn")

    @contact_flow_arn.setter
    def contact_flow_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "contact_flow_arn", value)

    @property
    @pulumi.getter(name="userArn")
    def user_arn(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-userquickconnectconfig.html#cfn-connect-quickconnect-userquickconnectconfig-userarn
        """
        return pulumi.get(self, "user_arn")

    @user_arn.setter
    def user_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_arn", value)

