# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'ApiBodyS3Location',
    'ApiCors',
    'ApiGatewayManagedOverridesAccessLogSettings',
    'ApiGatewayManagedOverridesIntegrationOverrides',
    'ApiGatewayManagedOverridesRouteOverrides',
    'ApiGatewayManagedOverridesRouteSettings',
    'ApiGatewayManagedOverridesStageOverrides',
    'AuthorizerJWTConfiguration',
    'DomainNameConfiguration',
    'DomainNameMutualTlsAuthentication',
    'IntegrationTlsConfig',
    'StageAccessLogSettings',
    'StageRouteSettings',
]

@pulumi.output_type
class ApiBodyS3Location(dict):
    def __init__(__self__, *,
                 bucket: Optional[str] = None,
                 etag: Optional[str] = None,
                 key: Optional[str] = None,
                 version: Optional[str] = None):
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[str]:
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        return pulumi.get(self, "version")


@pulumi.output_type
class ApiCors(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowCredentials":
            suggest = "allow_credentials"
        elif key == "allowHeaders":
            suggest = "allow_headers"
        elif key == "allowMethods":
            suggest = "allow_methods"
        elif key == "allowOrigins":
            suggest = "allow_origins"
        elif key == "exposeHeaders":
            suggest = "expose_headers"
        elif key == "maxAge":
            suggest = "max_age"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApiCors. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApiCors.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApiCors.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allow_credentials: Optional[bool] = None,
                 allow_headers: Optional[Sequence[str]] = None,
                 allow_methods: Optional[Sequence[str]] = None,
                 allow_origins: Optional[Sequence[str]] = None,
                 expose_headers: Optional[Sequence[str]] = None,
                 max_age: Optional[int] = None):
        if allow_credentials is not None:
            pulumi.set(__self__, "allow_credentials", allow_credentials)
        if allow_headers is not None:
            pulumi.set(__self__, "allow_headers", allow_headers)
        if allow_methods is not None:
            pulumi.set(__self__, "allow_methods", allow_methods)
        if allow_origins is not None:
            pulumi.set(__self__, "allow_origins", allow_origins)
        if expose_headers is not None:
            pulumi.set(__self__, "expose_headers", expose_headers)
        if max_age is not None:
            pulumi.set(__self__, "max_age", max_age)

    @property
    @pulumi.getter(name="allowCredentials")
    def allow_credentials(self) -> Optional[bool]:
        return pulumi.get(self, "allow_credentials")

    @property
    @pulumi.getter(name="allowHeaders")
    def allow_headers(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "allow_headers")

    @property
    @pulumi.getter(name="allowMethods")
    def allow_methods(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "allow_methods")

    @property
    @pulumi.getter(name="allowOrigins")
    def allow_origins(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "allow_origins")

    @property
    @pulumi.getter(name="exposeHeaders")
    def expose_headers(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "expose_headers")

    @property
    @pulumi.getter(name="maxAge")
    def max_age(self) -> Optional[int]:
        return pulumi.get(self, "max_age")


@pulumi.output_type
class ApiGatewayManagedOverridesAccessLogSettings(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "destinationArn":
            suggest = "destination_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApiGatewayManagedOverridesAccessLogSettings. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApiGatewayManagedOverridesAccessLogSettings.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApiGatewayManagedOverridesAccessLogSettings.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 destination_arn: Optional[str] = None,
                 format: Optional[str] = None):
        if destination_arn is not None:
            pulumi.set(__self__, "destination_arn", destination_arn)
        if format is not None:
            pulumi.set(__self__, "format", format)

    @property
    @pulumi.getter(name="destinationArn")
    def destination_arn(self) -> Optional[str]:
        return pulumi.get(self, "destination_arn")

    @property
    @pulumi.getter
    def format(self) -> Optional[str]:
        return pulumi.get(self, "format")


@pulumi.output_type
class ApiGatewayManagedOverridesIntegrationOverrides(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "integrationMethod":
            suggest = "integration_method"
        elif key == "payloadFormatVersion":
            suggest = "payload_format_version"
        elif key == "timeoutInMillis":
            suggest = "timeout_in_millis"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApiGatewayManagedOverridesIntegrationOverrides. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApiGatewayManagedOverridesIntegrationOverrides.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApiGatewayManagedOverridesIntegrationOverrides.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 description: Optional[str] = None,
                 integration_method: Optional[str] = None,
                 payload_format_version: Optional[str] = None,
                 timeout_in_millis: Optional[int] = None):
        if description is not None:
            pulumi.set(__self__, "description", description)
        if integration_method is not None:
            pulumi.set(__self__, "integration_method", integration_method)
        if payload_format_version is not None:
            pulumi.set(__self__, "payload_format_version", payload_format_version)
        if timeout_in_millis is not None:
            pulumi.set(__self__, "timeout_in_millis", timeout_in_millis)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="integrationMethod")
    def integration_method(self) -> Optional[str]:
        return pulumi.get(self, "integration_method")

    @property
    @pulumi.getter(name="payloadFormatVersion")
    def payload_format_version(self) -> Optional[str]:
        return pulumi.get(self, "payload_format_version")

    @property
    @pulumi.getter(name="timeoutInMillis")
    def timeout_in_millis(self) -> Optional[int]:
        return pulumi.get(self, "timeout_in_millis")


@pulumi.output_type
class ApiGatewayManagedOverridesRouteOverrides(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "authorizationScopes":
            suggest = "authorization_scopes"
        elif key == "authorizationType":
            suggest = "authorization_type"
        elif key == "authorizerId":
            suggest = "authorizer_id"
        elif key == "operationName":
            suggest = "operation_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApiGatewayManagedOverridesRouteOverrides. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApiGatewayManagedOverridesRouteOverrides.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApiGatewayManagedOverridesRouteOverrides.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 authorization_scopes: Optional[Sequence[str]] = None,
                 authorization_type: Optional[str] = None,
                 authorizer_id: Optional[str] = None,
                 operation_name: Optional[str] = None,
                 target: Optional[str] = None):
        if authorization_scopes is not None:
            pulumi.set(__self__, "authorization_scopes", authorization_scopes)
        if authorization_type is not None:
            pulumi.set(__self__, "authorization_type", authorization_type)
        if authorizer_id is not None:
            pulumi.set(__self__, "authorizer_id", authorizer_id)
        if operation_name is not None:
            pulumi.set(__self__, "operation_name", operation_name)
        if target is not None:
            pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter(name="authorizationScopes")
    def authorization_scopes(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "authorization_scopes")

    @property
    @pulumi.getter(name="authorizationType")
    def authorization_type(self) -> Optional[str]:
        return pulumi.get(self, "authorization_type")

    @property
    @pulumi.getter(name="authorizerId")
    def authorizer_id(self) -> Optional[str]:
        return pulumi.get(self, "authorizer_id")

    @property
    @pulumi.getter(name="operationName")
    def operation_name(self) -> Optional[str]:
        return pulumi.get(self, "operation_name")

    @property
    @pulumi.getter
    def target(self) -> Optional[str]:
        return pulumi.get(self, "target")


@pulumi.output_type
class ApiGatewayManagedOverridesRouteSettings(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dataTraceEnabled":
            suggest = "data_trace_enabled"
        elif key == "detailedMetricsEnabled":
            suggest = "detailed_metrics_enabled"
        elif key == "loggingLevel":
            suggest = "logging_level"
        elif key == "throttlingBurstLimit":
            suggest = "throttling_burst_limit"
        elif key == "throttlingRateLimit":
            suggest = "throttling_rate_limit"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApiGatewayManagedOverridesRouteSettings. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApiGatewayManagedOverridesRouteSettings.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApiGatewayManagedOverridesRouteSettings.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 data_trace_enabled: Optional[bool] = None,
                 detailed_metrics_enabled: Optional[bool] = None,
                 logging_level: Optional[str] = None,
                 throttling_burst_limit: Optional[int] = None,
                 throttling_rate_limit: Optional[float] = None):
        if data_trace_enabled is not None:
            pulumi.set(__self__, "data_trace_enabled", data_trace_enabled)
        if detailed_metrics_enabled is not None:
            pulumi.set(__self__, "detailed_metrics_enabled", detailed_metrics_enabled)
        if logging_level is not None:
            pulumi.set(__self__, "logging_level", logging_level)
        if throttling_burst_limit is not None:
            pulumi.set(__self__, "throttling_burst_limit", throttling_burst_limit)
        if throttling_rate_limit is not None:
            pulumi.set(__self__, "throttling_rate_limit", throttling_rate_limit)

    @property
    @pulumi.getter(name="dataTraceEnabled")
    def data_trace_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "data_trace_enabled")

    @property
    @pulumi.getter(name="detailedMetricsEnabled")
    def detailed_metrics_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "detailed_metrics_enabled")

    @property
    @pulumi.getter(name="loggingLevel")
    def logging_level(self) -> Optional[str]:
        return pulumi.get(self, "logging_level")

    @property
    @pulumi.getter(name="throttlingBurstLimit")
    def throttling_burst_limit(self) -> Optional[int]:
        return pulumi.get(self, "throttling_burst_limit")

    @property
    @pulumi.getter(name="throttlingRateLimit")
    def throttling_rate_limit(self) -> Optional[float]:
        return pulumi.get(self, "throttling_rate_limit")


@pulumi.output_type
class ApiGatewayManagedOverridesStageOverrides(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accessLogSettings":
            suggest = "access_log_settings"
        elif key == "autoDeploy":
            suggest = "auto_deploy"
        elif key == "defaultRouteSettings":
            suggest = "default_route_settings"
        elif key == "routeSettings":
            suggest = "route_settings"
        elif key == "stageVariables":
            suggest = "stage_variables"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApiGatewayManagedOverridesStageOverrides. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApiGatewayManagedOverridesStageOverrides.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApiGatewayManagedOverridesStageOverrides.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 access_log_settings: Optional['outputs.ApiGatewayManagedOverridesAccessLogSettings'] = None,
                 auto_deploy: Optional[bool] = None,
                 default_route_settings: Optional['outputs.ApiGatewayManagedOverridesRouteSettings'] = None,
                 description: Optional[str] = None,
                 route_settings: Optional[Any] = None,
                 stage_variables: Optional[Any] = None):
        if access_log_settings is not None:
            pulumi.set(__self__, "access_log_settings", access_log_settings)
        if auto_deploy is not None:
            pulumi.set(__self__, "auto_deploy", auto_deploy)
        if default_route_settings is not None:
            pulumi.set(__self__, "default_route_settings", default_route_settings)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if route_settings is not None:
            pulumi.set(__self__, "route_settings", route_settings)
        if stage_variables is not None:
            pulumi.set(__self__, "stage_variables", stage_variables)

    @property
    @pulumi.getter(name="accessLogSettings")
    def access_log_settings(self) -> Optional['outputs.ApiGatewayManagedOverridesAccessLogSettings']:
        return pulumi.get(self, "access_log_settings")

    @property
    @pulumi.getter(name="autoDeploy")
    def auto_deploy(self) -> Optional[bool]:
        return pulumi.get(self, "auto_deploy")

    @property
    @pulumi.getter(name="defaultRouteSettings")
    def default_route_settings(self) -> Optional['outputs.ApiGatewayManagedOverridesRouteSettings']:
        return pulumi.get(self, "default_route_settings")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="routeSettings")
    def route_settings(self) -> Optional[Any]:
        return pulumi.get(self, "route_settings")

    @property
    @pulumi.getter(name="stageVariables")
    def stage_variables(self) -> Optional[Any]:
        return pulumi.get(self, "stage_variables")


@pulumi.output_type
class AuthorizerJWTConfiguration(dict):
    def __init__(__self__, *,
                 audience: Optional[Sequence[str]] = None,
                 issuer: Optional[str] = None):
        if audience is not None:
            pulumi.set(__self__, "audience", audience)
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)

    @property
    @pulumi.getter
    def audience(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "audience")

    @property
    @pulumi.getter
    def issuer(self) -> Optional[str]:
        return pulumi.get(self, "issuer")


@pulumi.output_type
class DomainNameConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "certificateArn":
            suggest = "certificate_arn"
        elif key == "certificateName":
            suggest = "certificate_name"
        elif key == "endpointType":
            suggest = "endpoint_type"
        elif key == "ownershipVerificationCertificateArn":
            suggest = "ownership_verification_certificate_arn"
        elif key == "securityPolicy":
            suggest = "security_policy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DomainNameConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DomainNameConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DomainNameConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 certificate_arn: Optional[str] = None,
                 certificate_name: Optional[str] = None,
                 endpoint_type: Optional[str] = None,
                 ownership_verification_certificate_arn: Optional[str] = None,
                 security_policy: Optional[str] = None):
        if certificate_arn is not None:
            pulumi.set(__self__, "certificate_arn", certificate_arn)
        if certificate_name is not None:
            pulumi.set(__self__, "certificate_name", certificate_name)
        if endpoint_type is not None:
            pulumi.set(__self__, "endpoint_type", endpoint_type)
        if ownership_verification_certificate_arn is not None:
            pulumi.set(__self__, "ownership_verification_certificate_arn", ownership_verification_certificate_arn)
        if security_policy is not None:
            pulumi.set(__self__, "security_policy", security_policy)

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> Optional[str]:
        return pulumi.get(self, "certificate_arn")

    @property
    @pulumi.getter(name="certificateName")
    def certificate_name(self) -> Optional[str]:
        return pulumi.get(self, "certificate_name")

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> Optional[str]:
        return pulumi.get(self, "endpoint_type")

    @property
    @pulumi.getter(name="ownershipVerificationCertificateArn")
    def ownership_verification_certificate_arn(self) -> Optional[str]:
        return pulumi.get(self, "ownership_verification_certificate_arn")

    @property
    @pulumi.getter(name="securityPolicy")
    def security_policy(self) -> Optional[str]:
        return pulumi.get(self, "security_policy")


@pulumi.output_type
class DomainNameMutualTlsAuthentication(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "truststoreUri":
            suggest = "truststore_uri"
        elif key == "truststoreVersion":
            suggest = "truststore_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DomainNameMutualTlsAuthentication. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DomainNameMutualTlsAuthentication.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DomainNameMutualTlsAuthentication.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 truststore_uri: Optional[str] = None,
                 truststore_version: Optional[str] = None):
        if truststore_uri is not None:
            pulumi.set(__self__, "truststore_uri", truststore_uri)
        if truststore_version is not None:
            pulumi.set(__self__, "truststore_version", truststore_version)

    @property
    @pulumi.getter(name="truststoreUri")
    def truststore_uri(self) -> Optional[str]:
        return pulumi.get(self, "truststore_uri")

    @property
    @pulumi.getter(name="truststoreVersion")
    def truststore_version(self) -> Optional[str]:
        return pulumi.get(self, "truststore_version")


@pulumi.output_type
class IntegrationTlsConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "serverNameToVerify":
            suggest = "server_name_to_verify"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IntegrationTlsConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IntegrationTlsConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IntegrationTlsConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 server_name_to_verify: Optional[str] = None):
        if server_name_to_verify is not None:
            pulumi.set(__self__, "server_name_to_verify", server_name_to_verify)

    @property
    @pulumi.getter(name="serverNameToVerify")
    def server_name_to_verify(self) -> Optional[str]:
        return pulumi.get(self, "server_name_to_verify")


@pulumi.output_type
class StageAccessLogSettings(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "destinationArn":
            suggest = "destination_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StageAccessLogSettings. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StageAccessLogSettings.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StageAccessLogSettings.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 destination_arn: Optional[str] = None,
                 format: Optional[str] = None):
        if destination_arn is not None:
            pulumi.set(__self__, "destination_arn", destination_arn)
        if format is not None:
            pulumi.set(__self__, "format", format)

    @property
    @pulumi.getter(name="destinationArn")
    def destination_arn(self) -> Optional[str]:
        return pulumi.get(self, "destination_arn")

    @property
    @pulumi.getter
    def format(self) -> Optional[str]:
        return pulumi.get(self, "format")


@pulumi.output_type
class StageRouteSettings(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dataTraceEnabled":
            suggest = "data_trace_enabled"
        elif key == "detailedMetricsEnabled":
            suggest = "detailed_metrics_enabled"
        elif key == "loggingLevel":
            suggest = "logging_level"
        elif key == "throttlingBurstLimit":
            suggest = "throttling_burst_limit"
        elif key == "throttlingRateLimit":
            suggest = "throttling_rate_limit"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StageRouteSettings. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StageRouteSettings.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StageRouteSettings.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 data_trace_enabled: Optional[bool] = None,
                 detailed_metrics_enabled: Optional[bool] = None,
                 logging_level: Optional[str] = None,
                 throttling_burst_limit: Optional[int] = None,
                 throttling_rate_limit: Optional[float] = None):
        if data_trace_enabled is not None:
            pulumi.set(__self__, "data_trace_enabled", data_trace_enabled)
        if detailed_metrics_enabled is not None:
            pulumi.set(__self__, "detailed_metrics_enabled", detailed_metrics_enabled)
        if logging_level is not None:
            pulumi.set(__self__, "logging_level", logging_level)
        if throttling_burst_limit is not None:
            pulumi.set(__self__, "throttling_burst_limit", throttling_burst_limit)
        if throttling_rate_limit is not None:
            pulumi.set(__self__, "throttling_rate_limit", throttling_rate_limit)

    @property
    @pulumi.getter(name="dataTraceEnabled")
    def data_trace_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "data_trace_enabled")

    @property
    @pulumi.getter(name="detailedMetricsEnabled")
    def detailed_metrics_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "detailed_metrics_enabled")

    @property
    @pulumi.getter(name="loggingLevel")
    def logging_level(self) -> Optional[str]:
        return pulumi.get(self, "logging_level")

    @property
    @pulumi.getter(name="throttlingBurstLimit")
    def throttling_burst_limit(self) -> Optional[int]:
        return pulumi.get(self, "throttling_burst_limit")

    @property
    @pulumi.getter(name="throttlingRateLimit")
    def throttling_rate_limit(self) -> Optional[float]:
        return pulumi.get(self, "throttling_rate_limit")


