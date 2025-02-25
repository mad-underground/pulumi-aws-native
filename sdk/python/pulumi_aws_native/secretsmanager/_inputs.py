# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'RotationScheduleHostedRotationLambdaArgs',
    'RotationScheduleRotationRulesArgs',
    'SecretGenerateSecretStringArgs',
    'SecretReplicaRegionArgs',
]

@pulumi.input_type
class RotationScheduleHostedRotationLambdaArgs:
    def __init__(__self__, *,
                 rotation_type: pulumi.Input[str],
                 exclude_characters: Optional[pulumi.Input[str]] = None,
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 master_secret_arn: Optional[pulumi.Input[str]] = None,
                 master_secret_kms_key_arn: Optional[pulumi.Input[str]] = None,
                 rotation_lambda_name: Optional[pulumi.Input[str]] = None,
                 runtime: Optional[pulumi.Input[str]] = None,
                 superuser_secret_arn: Optional[pulumi.Input[str]] = None,
                 superuser_secret_kms_key_arn: Optional[pulumi.Input[str]] = None,
                 vpc_security_group_ids: Optional[pulumi.Input[str]] = None,
                 vpc_subnet_ids: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "rotation_type", rotation_type)
        if exclude_characters is not None:
            pulumi.set(__self__, "exclude_characters", exclude_characters)
        if kms_key_arn is not None:
            pulumi.set(__self__, "kms_key_arn", kms_key_arn)
        if master_secret_arn is not None:
            pulumi.set(__self__, "master_secret_arn", master_secret_arn)
        if master_secret_kms_key_arn is not None:
            pulumi.set(__self__, "master_secret_kms_key_arn", master_secret_kms_key_arn)
        if rotation_lambda_name is not None:
            pulumi.set(__self__, "rotation_lambda_name", rotation_lambda_name)
        if runtime is not None:
            pulumi.set(__self__, "runtime", runtime)
        if superuser_secret_arn is not None:
            pulumi.set(__self__, "superuser_secret_arn", superuser_secret_arn)
        if superuser_secret_kms_key_arn is not None:
            pulumi.set(__self__, "superuser_secret_kms_key_arn", superuser_secret_kms_key_arn)
        if vpc_security_group_ids is not None:
            pulumi.set(__self__, "vpc_security_group_ids", vpc_security_group_ids)
        if vpc_subnet_ids is not None:
            pulumi.set(__self__, "vpc_subnet_ids", vpc_subnet_ids)

    @property
    @pulumi.getter(name="rotationType")
    def rotation_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "rotation_type")

    @rotation_type.setter
    def rotation_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "rotation_type", value)

    @property
    @pulumi.getter(name="excludeCharacters")
    def exclude_characters(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "exclude_characters")

    @exclude_characters.setter
    def exclude_characters(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclude_characters", value)

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kms_key_arn")

    @kms_key_arn.setter
    def kms_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_arn", value)

    @property
    @pulumi.getter(name="masterSecretArn")
    def master_secret_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "master_secret_arn")

    @master_secret_arn.setter
    def master_secret_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_secret_arn", value)

    @property
    @pulumi.getter(name="masterSecretKmsKeyArn")
    def master_secret_kms_key_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "master_secret_kms_key_arn")

    @master_secret_kms_key_arn.setter
    def master_secret_kms_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_secret_kms_key_arn", value)

    @property
    @pulumi.getter(name="rotationLambdaName")
    def rotation_lambda_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "rotation_lambda_name")

    @rotation_lambda_name.setter
    def rotation_lambda_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rotation_lambda_name", value)

    @property
    @pulumi.getter
    def runtime(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "runtime")

    @runtime.setter
    def runtime(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "runtime", value)

    @property
    @pulumi.getter(name="superuserSecretArn")
    def superuser_secret_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "superuser_secret_arn")

    @superuser_secret_arn.setter
    def superuser_secret_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "superuser_secret_arn", value)

    @property
    @pulumi.getter(name="superuserSecretKmsKeyArn")
    def superuser_secret_kms_key_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "superuser_secret_kms_key_arn")

    @superuser_secret_kms_key_arn.setter
    def superuser_secret_kms_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "superuser_secret_kms_key_arn", value)

    @property
    @pulumi.getter(name="vpcSecurityGroupIds")
    def vpc_security_group_ids(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vpc_security_group_ids")

    @vpc_security_group_ids.setter
    def vpc_security_group_ids(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_security_group_ids", value)

    @property
    @pulumi.getter(name="vpcSubnetIds")
    def vpc_subnet_ids(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vpc_subnet_ids")

    @vpc_subnet_ids.setter
    def vpc_subnet_ids(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_subnet_ids", value)


@pulumi.input_type
class RotationScheduleRotationRulesArgs:
    def __init__(__self__, *,
                 automatically_after_days: Optional[pulumi.Input[int]] = None,
                 duration: Optional[pulumi.Input[str]] = None,
                 schedule_expression: Optional[pulumi.Input[str]] = None):
        if automatically_after_days is not None:
            pulumi.set(__self__, "automatically_after_days", automatically_after_days)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if schedule_expression is not None:
            pulumi.set(__self__, "schedule_expression", schedule_expression)

    @property
    @pulumi.getter(name="automaticallyAfterDays")
    def automatically_after_days(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "automatically_after_days")

    @automatically_after_days.setter
    def automatically_after_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "automatically_after_days", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter(name="scheduleExpression")
    def schedule_expression(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "schedule_expression")

    @schedule_expression.setter
    def schedule_expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_expression", value)


@pulumi.input_type
class SecretGenerateSecretStringArgs:
    def __init__(__self__, *,
                 exclude_characters: Optional[pulumi.Input[str]] = None,
                 exclude_lowercase: Optional[pulumi.Input[bool]] = None,
                 exclude_numbers: Optional[pulumi.Input[bool]] = None,
                 exclude_punctuation: Optional[pulumi.Input[bool]] = None,
                 exclude_uppercase: Optional[pulumi.Input[bool]] = None,
                 generate_string_key: Optional[pulumi.Input[str]] = None,
                 include_space: Optional[pulumi.Input[bool]] = None,
                 password_length: Optional[pulumi.Input[int]] = None,
                 require_each_included_type: Optional[pulumi.Input[bool]] = None,
                 secret_string_template: Optional[pulumi.Input[str]] = None):
        """
        Generates a random password. We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support.
          *Required permissions:* ``secretsmanager:GetRandomPassword``. For more information, see [IAM policy actions for Secrets Manager](https://docs.aws.amazon.com/service-authorization/latest/reference/list_awssecretsmanager.html#awssecretsmanager-actions-as-permissions) and [Authentication and access control in Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html).
        :param pulumi.Input[str] exclude_characters: A string of the characters that you don't want in the password.
        :param pulumi.Input[bool] exclude_lowercase: Specifies whether to exclude lowercase letters from the password. If you don't include this switch, the password can contain lowercase letters.
        :param pulumi.Input[bool] exclude_numbers: Specifies whether to exclude numbers from the password. If you don't include this switch, the password can contain numbers.
        :param pulumi.Input[bool] exclude_punctuation: Specifies whether to exclude the following punctuation characters from the password: ``! " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \\ ] ^ _ ` { | } ~``. If you don't include this switch, the password can contain punctuation.
        :param pulumi.Input[bool] exclude_uppercase: Specifies whether to exclude uppercase letters from the password. If you don't include this switch, the password can contain uppercase letters.
        :param pulumi.Input[str] generate_string_key: The JSON key name for the key/value pair, where the value is the generated password. This pair is added to the JSON structure specified by the ``SecretStringTemplate`` parameter. If you specify this parameter, then you must also specify ``SecretStringTemplate``.
        :param pulumi.Input[bool] include_space: Specifies whether to include the space character. If you include this switch, the password can contain space characters.
        :param pulumi.Input[int] password_length: The length of the password. If you don't include this parameter, the default length is 32 characters.
        :param pulumi.Input[bool] require_each_included_type: Specifies whether to include at least one upper and lowercase letter, one number, and one punctuation. If you don't include this switch, the password contains at least one of every character type.
        :param pulumi.Input[str] secret_string_template: A template that the generated string must match. When you make a change to this property, a new secret version is created.
        """
        if exclude_characters is not None:
            pulumi.set(__self__, "exclude_characters", exclude_characters)
        if exclude_lowercase is not None:
            pulumi.set(__self__, "exclude_lowercase", exclude_lowercase)
        if exclude_numbers is not None:
            pulumi.set(__self__, "exclude_numbers", exclude_numbers)
        if exclude_punctuation is not None:
            pulumi.set(__self__, "exclude_punctuation", exclude_punctuation)
        if exclude_uppercase is not None:
            pulumi.set(__self__, "exclude_uppercase", exclude_uppercase)
        if generate_string_key is not None:
            pulumi.set(__self__, "generate_string_key", generate_string_key)
        if include_space is not None:
            pulumi.set(__self__, "include_space", include_space)
        if password_length is not None:
            pulumi.set(__self__, "password_length", password_length)
        if require_each_included_type is not None:
            pulumi.set(__self__, "require_each_included_type", require_each_included_type)
        if secret_string_template is not None:
            pulumi.set(__self__, "secret_string_template", secret_string_template)

    @property
    @pulumi.getter(name="excludeCharacters")
    def exclude_characters(self) -> Optional[pulumi.Input[str]]:
        """
        A string of the characters that you don't want in the password.
        """
        return pulumi.get(self, "exclude_characters")

    @exclude_characters.setter
    def exclude_characters(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclude_characters", value)

    @property
    @pulumi.getter(name="excludeLowercase")
    def exclude_lowercase(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to exclude lowercase letters from the password. If you don't include this switch, the password can contain lowercase letters.
        """
        return pulumi.get(self, "exclude_lowercase")

    @exclude_lowercase.setter
    def exclude_lowercase(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "exclude_lowercase", value)

    @property
    @pulumi.getter(name="excludeNumbers")
    def exclude_numbers(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to exclude numbers from the password. If you don't include this switch, the password can contain numbers.
        """
        return pulumi.get(self, "exclude_numbers")

    @exclude_numbers.setter
    def exclude_numbers(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "exclude_numbers", value)

    @property
    @pulumi.getter(name="excludePunctuation")
    def exclude_punctuation(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to exclude the following punctuation characters from the password: ``! " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \\ ] ^ _ ` { | } ~``. If you don't include this switch, the password can contain punctuation.
        """
        return pulumi.get(self, "exclude_punctuation")

    @exclude_punctuation.setter
    def exclude_punctuation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "exclude_punctuation", value)

    @property
    @pulumi.getter(name="excludeUppercase")
    def exclude_uppercase(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to exclude uppercase letters from the password. If you don't include this switch, the password can contain uppercase letters.
        """
        return pulumi.get(self, "exclude_uppercase")

    @exclude_uppercase.setter
    def exclude_uppercase(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "exclude_uppercase", value)

    @property
    @pulumi.getter(name="generateStringKey")
    def generate_string_key(self) -> Optional[pulumi.Input[str]]:
        """
        The JSON key name for the key/value pair, where the value is the generated password. This pair is added to the JSON structure specified by the ``SecretStringTemplate`` parameter. If you specify this parameter, then you must also specify ``SecretStringTemplate``.
        """
        return pulumi.get(self, "generate_string_key")

    @generate_string_key.setter
    def generate_string_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "generate_string_key", value)

    @property
    @pulumi.getter(name="includeSpace")
    def include_space(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to include the space character. If you include this switch, the password can contain space characters.
        """
        return pulumi.get(self, "include_space")

    @include_space.setter
    def include_space(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_space", value)

    @property
    @pulumi.getter(name="passwordLength")
    def password_length(self) -> Optional[pulumi.Input[int]]:
        """
        The length of the password. If you don't include this parameter, the default length is 32 characters.
        """
        return pulumi.get(self, "password_length")

    @password_length.setter
    def password_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "password_length", value)

    @property
    @pulumi.getter(name="requireEachIncludedType")
    def require_each_included_type(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to include at least one upper and lowercase letter, one number, and one punctuation. If you don't include this switch, the password contains at least one of every character type.
        """
        return pulumi.get(self, "require_each_included_type")

    @require_each_included_type.setter
    def require_each_included_type(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_each_included_type", value)

    @property
    @pulumi.getter(name="secretStringTemplate")
    def secret_string_template(self) -> Optional[pulumi.Input[str]]:
        """
        A template that the generated string must match. When you make a change to this property, a new secret version is created.
        """
        return pulumi.get(self, "secret_string_template")

    @secret_string_template.setter
    def secret_string_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_string_template", value)


@pulumi.input_type
class SecretReplicaRegionArgs:
    def __init__(__self__, *,
                 region: pulumi.Input[str],
                 kms_key_id: Optional[pulumi.Input[str]] = None):
        """
        Specifies a ``Region`` and the ``KmsKeyId`` for a replica secret.
        :param pulumi.Input[str] region: A string that represents a ``Region``, for example "us-east-1".
        :param pulumi.Input[str] kms_key_id: The ARN, key ID, or alias of the KMS key to encrypt the secret. If you don't include this field, Secrets Manager uses ``aws/secretsmanager``.
        """
        pulumi.set(__self__, "region", region)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        A string that represents a ``Region``, for example "us-east-1".
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN, key ID, or alias of the KMS key to encrypt the secret. If you don't include this field, Secrets Manager uses ``aws/secretsmanager``.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)


