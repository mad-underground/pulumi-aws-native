# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'CertificateAuthorityCrlConfigurationArgs',
    'CertificateAuthorityRevocationConfigurationArgs',
    'CertificateAuthoritySubjectArgs',
    'CertificateValidityArgs',
]

@pulumi.input_type
class CertificateAuthorityCrlConfigurationArgs:
    def __init__(__self__, *,
                 custom_cname: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 expiration_in_days: Optional[pulumi.Input[int]] = None,
                 s3_bucket_name: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html
        :param pulumi.Input[str] custom_cname: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html#cfn-acmpca-certificateauthority-crlconfiguration-customcname
        :param pulumi.Input[bool] enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html#cfn-acmpca-certificateauthority-crlconfiguration-enabled
        :param pulumi.Input[int] expiration_in_days: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html#cfn-acmpca-certificateauthority-crlconfiguration-expirationindays
        :param pulumi.Input[str] s3_bucket_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html#cfn-acmpca-certificateauthority-crlconfiguration-s3bucketname
        """
        if custom_cname is not None:
            pulumi.set(__self__, "custom_cname", custom_cname)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if expiration_in_days is not None:
            pulumi.set(__self__, "expiration_in_days", expiration_in_days)
        if s3_bucket_name is not None:
            pulumi.set(__self__, "s3_bucket_name", s3_bucket_name)

    @property
    @pulumi.getter(name="customCname")
    def custom_cname(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html#cfn-acmpca-certificateauthority-crlconfiguration-customcname
        """
        return pulumi.get(self, "custom_cname")

    @custom_cname.setter
    def custom_cname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_cname", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html#cfn-acmpca-certificateauthority-crlconfiguration-enabled
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="expirationInDays")
    def expiration_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html#cfn-acmpca-certificateauthority-crlconfiguration-expirationindays
        """
        return pulumi.get(self, "expiration_in_days")

    @expiration_in_days.setter
    def expiration_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expiration_in_days", value)

    @property
    @pulumi.getter(name="s3BucketName")
    def s3_bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html#cfn-acmpca-certificateauthority-crlconfiguration-s3bucketname
        """
        return pulumi.get(self, "s3_bucket_name")

    @s3_bucket_name.setter
    def s3_bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_bucket_name", value)


@pulumi.input_type
class CertificateAuthorityRevocationConfigurationArgs:
    def __init__(__self__, *,
                 crl_configuration: Optional[pulumi.Input['CertificateAuthorityCrlConfigurationArgs']] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-revocationconfiguration.html
        :param pulumi.Input['CertificateAuthorityCrlConfigurationArgs'] crl_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-revocationconfiguration.html#cfn-acmpca-certificateauthority-revocationconfiguration-crlconfiguration
        """
        if crl_configuration is not None:
            pulumi.set(__self__, "crl_configuration", crl_configuration)

    @property
    @pulumi.getter(name="crlConfiguration")
    def crl_configuration(self) -> Optional[pulumi.Input['CertificateAuthorityCrlConfigurationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-revocationconfiguration.html#cfn-acmpca-certificateauthority-revocationconfiguration-crlconfiguration
        """
        return pulumi.get(self, "crl_configuration")

    @crl_configuration.setter
    def crl_configuration(self, value: Optional[pulumi.Input['CertificateAuthorityCrlConfigurationArgs']]):
        pulumi.set(self, "crl_configuration", value)


@pulumi.input_type
class CertificateAuthoritySubjectArgs:
    def __init__(__self__, *,
                 common_name: Optional[pulumi.Input[str]] = None,
                 country: Optional[pulumi.Input[str]] = None,
                 distinguished_name_qualifier: Optional[pulumi.Input[str]] = None,
                 generation_qualifier: Optional[pulumi.Input[str]] = None,
                 given_name: Optional[pulumi.Input[str]] = None,
                 initials: Optional[pulumi.Input[str]] = None,
                 locality: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 organizational_unit: Optional[pulumi.Input[str]] = None,
                 pseudonym: Optional[pulumi.Input[str]] = None,
                 serial_number: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 surname: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html
        :param pulumi.Input[str] common_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-commonname
        :param pulumi.Input[str] country: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-country
        :param pulumi.Input[str] distinguished_name_qualifier: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-distinguishednamequalifier
        :param pulumi.Input[str] generation_qualifier: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-generationqualifier
        :param pulumi.Input[str] given_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-givenname
        :param pulumi.Input[str] initials: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-initials
        :param pulumi.Input[str] locality: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-locality
        :param pulumi.Input[str] organization: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-organization
        :param pulumi.Input[str] organizational_unit: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-organizationalunit
        :param pulumi.Input[str] pseudonym: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-pseudonym
        :param pulumi.Input[str] serial_number: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-serialnumber
        :param pulumi.Input[str] state: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-state
        :param pulumi.Input[str] surname: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-surname
        :param pulumi.Input[str] title: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-title
        """
        if common_name is not None:
            pulumi.set(__self__, "common_name", common_name)
        if country is not None:
            pulumi.set(__self__, "country", country)
        if distinguished_name_qualifier is not None:
            pulumi.set(__self__, "distinguished_name_qualifier", distinguished_name_qualifier)
        if generation_qualifier is not None:
            pulumi.set(__self__, "generation_qualifier", generation_qualifier)
        if given_name is not None:
            pulumi.set(__self__, "given_name", given_name)
        if initials is not None:
            pulumi.set(__self__, "initials", initials)
        if locality is not None:
            pulumi.set(__self__, "locality", locality)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if organizational_unit is not None:
            pulumi.set(__self__, "organizational_unit", organizational_unit)
        if pseudonym is not None:
            pulumi.set(__self__, "pseudonym", pseudonym)
        if serial_number is not None:
            pulumi.set(__self__, "serial_number", serial_number)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if surname is not None:
            pulumi.set(__self__, "surname", surname)
        if title is not None:
            pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-commonname
        """
        return pulumi.get(self, "common_name")

    @common_name.setter
    def common_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "common_name", value)

    @property
    @pulumi.getter
    def country(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-country
        """
        return pulumi.get(self, "country")

    @country.setter
    def country(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country", value)

    @property
    @pulumi.getter(name="distinguishedNameQualifier")
    def distinguished_name_qualifier(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-distinguishednamequalifier
        """
        return pulumi.get(self, "distinguished_name_qualifier")

    @distinguished_name_qualifier.setter
    def distinguished_name_qualifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "distinguished_name_qualifier", value)

    @property
    @pulumi.getter(name="generationQualifier")
    def generation_qualifier(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-generationqualifier
        """
        return pulumi.get(self, "generation_qualifier")

    @generation_qualifier.setter
    def generation_qualifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "generation_qualifier", value)

    @property
    @pulumi.getter(name="givenName")
    def given_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-givenname
        """
        return pulumi.get(self, "given_name")

    @given_name.setter
    def given_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "given_name", value)

    @property
    @pulumi.getter
    def initials(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-initials
        """
        return pulumi.get(self, "initials")

    @initials.setter
    def initials(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "initials", value)

    @property
    @pulumi.getter
    def locality(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-locality
        """
        return pulumi.get(self, "locality")

    @locality.setter
    def locality(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "locality", value)

    @property
    @pulumi.getter
    def organization(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-organization
        """
        return pulumi.get(self, "organization")

    @organization.setter
    def organization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization", value)

    @property
    @pulumi.getter(name="organizationalUnit")
    def organizational_unit(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-organizationalunit
        """
        return pulumi.get(self, "organizational_unit")

    @organizational_unit.setter
    def organizational_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organizational_unit", value)

    @property
    @pulumi.getter
    def pseudonym(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-pseudonym
        """
        return pulumi.get(self, "pseudonym")

    @pseudonym.setter
    def pseudonym(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pseudonym", value)

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-serialnumber
        """
        return pulumi.get(self, "serial_number")

    @serial_number.setter
    def serial_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "serial_number", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-state
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def surname(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-surname
        """
        return pulumi.get(self, "surname")

    @surname.setter
    def surname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "surname", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-subject.html#cfn-acmpca-certificateauthority-subject-title
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)


@pulumi.input_type
class CertificateValidityArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 value: pulumi.Input[int]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-validity.html
        :param pulumi.Input[str] type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-validity.html#cfn-acmpca-certificate-validity-type
        :param pulumi.Input[int] value: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-validity.html#cfn-acmpca-certificate-validity-value
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-validity.html#cfn-acmpca-certificate-validity-type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-validity.html#cfn-acmpca-certificate-validity-value
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[int]):
        pulumi.set(self, "value", value)

