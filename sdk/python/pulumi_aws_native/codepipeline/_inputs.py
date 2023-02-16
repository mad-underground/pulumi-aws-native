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
    'CustomActionTypeArtifactDetailsArgs',
    'CustomActionTypeConfigurationPropertiesArgs',
    'CustomActionTypeSettingsArgs',
    'CustomActionTypeTagArgs',
    'PipelineActionDeclarationArgs',
    'PipelineActionTypeIdArgs',
    'PipelineArtifactStoreMapArgs',
    'PipelineArtifactStoreArgs',
    'PipelineBlockerDeclarationArgs',
    'PipelineEncryptionKeyArgs',
    'PipelineInputArtifactArgs',
    'PipelineOutputArtifactArgs',
    'PipelineStageDeclarationArgs',
    'PipelineStageTransitionArgs',
    'PipelineTagArgs',
    'WebhookAuthConfigurationArgs',
    'WebhookFilterRuleArgs',
]

@pulumi.input_type
class CustomActionTypeArtifactDetailsArgs:
    def __init__(__self__, *,
                 maximum_count: pulumi.Input[int],
                 minimum_count: pulumi.Input[int]):
        """
        Returns information about the details of an artifact.
        :param pulumi.Input[int] maximum_count: The maximum number of artifacts allowed for the action type.
        :param pulumi.Input[int] minimum_count: The minimum number of artifacts allowed for the action type.
        """
        pulumi.set(__self__, "maximum_count", maximum_count)
        pulumi.set(__self__, "minimum_count", minimum_count)

    @property
    @pulumi.getter(name="maximumCount")
    def maximum_count(self) -> pulumi.Input[int]:
        """
        The maximum number of artifacts allowed for the action type.
        """
        return pulumi.get(self, "maximum_count")

    @maximum_count.setter
    def maximum_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "maximum_count", value)

    @property
    @pulumi.getter(name="minimumCount")
    def minimum_count(self) -> pulumi.Input[int]:
        """
        The minimum number of artifacts allowed for the action type.
        """
        return pulumi.get(self, "minimum_count")

    @minimum_count.setter
    def minimum_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "minimum_count", value)


@pulumi.input_type
class CustomActionTypeConfigurationPropertiesArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[bool],
                 name: pulumi.Input[str],
                 required: pulumi.Input[bool],
                 secret: pulumi.Input[bool],
                 description: Optional[pulumi.Input[str]] = None,
                 queryable: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The configuration properties for the custom action.
        :param pulumi.Input[bool] key: Whether the configuration property is a key.
        :param pulumi.Input[str] name: The name of the action configuration property.
        :param pulumi.Input[bool] required: Whether the configuration property is a required value.
        :param pulumi.Input[bool] secret: Whether the configuration property is secret. Secrets are hidden from all calls except for GetJobDetails, GetThirdPartyJobDetails, PollForJobs, and PollForThirdPartyJobs.
        :param pulumi.Input[str] description: The description of the action configuration property that is displayed to users. 
        :param pulumi.Input[bool] queryable: Indicates that the property is used with PollForJobs. When creating a custom action, an action can have up to one queryable property. If it has one, that property must be both required and not secret.If you create a pipeline with a custom action type, and that custom action contains a queryable property, the value for that configuration property is subject to other restrictions. The value must be less than or equal to twenty (20) characters. The value can contain only alphanumeric characters, underscores, and hyphens. 
        :param pulumi.Input[str] type: The type of the configuration property.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "required", required)
        pulumi.set(__self__, "secret", secret)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if queryable is not None:
            pulumi.set(__self__, "queryable", queryable)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[bool]:
        """
        Whether the configuration property is a key.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[bool]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the action configuration property.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def required(self) -> pulumi.Input[bool]:
        """
        Whether the configuration property is a required value.
        """
        return pulumi.get(self, "required")

    @required.setter
    def required(self, value: pulumi.Input[bool]):
        pulumi.set(self, "required", value)

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Input[bool]:
        """
        Whether the configuration property is secret. Secrets are hidden from all calls except for GetJobDetails, GetThirdPartyJobDetails, PollForJobs, and PollForThirdPartyJobs.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: pulumi.Input[bool]):
        pulumi.set(self, "secret", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the action configuration property that is displayed to users. 
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def queryable(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates that the property is used with PollForJobs. When creating a custom action, an action can have up to one queryable property. If it has one, that property must be both required and not secret.If you create a pipeline with a custom action type, and that custom action contains a queryable property, the value for that configuration property is subject to other restrictions. The value must be less than or equal to twenty (20) characters. The value can contain only alphanumeric characters, underscores, and hyphens. 
        """
        return pulumi.get(self, "queryable")

    @queryable.setter
    def queryable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "queryable", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the configuration property.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class CustomActionTypeSettingsArgs:
    def __init__(__self__, *,
                 entity_url_template: Optional[pulumi.Input[str]] = None,
                 execution_url_template: Optional[pulumi.Input[str]] = None,
                 revision_url_template: Optional[pulumi.Input[str]] = None,
                 third_party_configuration_url: Optional[pulumi.Input[str]] = None):
        """
        Settings is a property of the AWS::CodePipeline::CustomActionType resource that provides URLs that users can access to view information about the CodePipeline custom action. 
        :param pulumi.Input[str] entity_url_template: The URL returned to the AWS CodePipeline console that provides a deep link to the resources of the external system, such as the configuration page for an AWS CodeDeploy deployment group. This link is provided as part of the action display in the pipeline. 
        :param pulumi.Input[str] execution_url_template: The URL returned to the AWS CodePipeline console that contains a link to the top-level landing page for the external system, such as the console page for AWS CodeDeploy. This link is shown on the pipeline view page in the AWS CodePipeline console and provides a link to the execution entity of the external action. 
        :param pulumi.Input[str] revision_url_template: The URL returned to the AWS CodePipeline console that contains a link to the page where customers can update or change the configuration of the external action. 
        :param pulumi.Input[str] third_party_configuration_url: The URL of a sign-up page where users can sign up for an external service and perform initial configuration of the action provided by that service.
        """
        if entity_url_template is not None:
            pulumi.set(__self__, "entity_url_template", entity_url_template)
        if execution_url_template is not None:
            pulumi.set(__self__, "execution_url_template", execution_url_template)
        if revision_url_template is not None:
            pulumi.set(__self__, "revision_url_template", revision_url_template)
        if third_party_configuration_url is not None:
            pulumi.set(__self__, "third_party_configuration_url", third_party_configuration_url)

    @property
    @pulumi.getter(name="entityUrlTemplate")
    def entity_url_template(self) -> Optional[pulumi.Input[str]]:
        """
        The URL returned to the AWS CodePipeline console that provides a deep link to the resources of the external system, such as the configuration page for an AWS CodeDeploy deployment group. This link is provided as part of the action display in the pipeline. 
        """
        return pulumi.get(self, "entity_url_template")

    @entity_url_template.setter
    def entity_url_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_url_template", value)

    @property
    @pulumi.getter(name="executionUrlTemplate")
    def execution_url_template(self) -> Optional[pulumi.Input[str]]:
        """
        The URL returned to the AWS CodePipeline console that contains a link to the top-level landing page for the external system, such as the console page for AWS CodeDeploy. This link is shown on the pipeline view page in the AWS CodePipeline console and provides a link to the execution entity of the external action. 
        """
        return pulumi.get(self, "execution_url_template")

    @execution_url_template.setter
    def execution_url_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "execution_url_template", value)

    @property
    @pulumi.getter(name="revisionUrlTemplate")
    def revision_url_template(self) -> Optional[pulumi.Input[str]]:
        """
        The URL returned to the AWS CodePipeline console that contains a link to the page where customers can update or change the configuration of the external action. 
        """
        return pulumi.get(self, "revision_url_template")

    @revision_url_template.setter
    def revision_url_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "revision_url_template", value)

    @property
    @pulumi.getter(name="thirdPartyConfigurationUrl")
    def third_party_configuration_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of a sign-up page where users can sign up for an external service and perform initial configuration of the action provided by that service.
        """
        return pulumi.get(self, "third_party_configuration_url")

    @third_party_configuration_url.setter
    def third_party_configuration_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "third_party_configuration_url", value)


@pulumi.input_type
class CustomActionTypeTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class PipelineActionDeclarationArgs:
    def __init__(__self__, *,
                 action_type_id: pulumi.Input['PipelineActionTypeIdArgs'],
                 name: pulumi.Input[str],
                 configuration: Optional[Any] = None,
                 input_artifacts: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineInputArtifactArgs']]]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 output_artifacts: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineOutputArtifactArgs']]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 run_order: Optional[pulumi.Input[int]] = None):
        pulumi.set(__self__, "action_type_id", action_type_id)
        pulumi.set(__self__, "name", name)
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if input_artifacts is not None:
            pulumi.set(__self__, "input_artifacts", input_artifacts)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if output_artifacts is not None:
            pulumi.set(__self__, "output_artifacts", output_artifacts)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if run_order is not None:
            pulumi.set(__self__, "run_order", run_order)

    @property
    @pulumi.getter(name="actionTypeId")
    def action_type_id(self) -> pulumi.Input['PipelineActionTypeIdArgs']:
        return pulumi.get(self, "action_type_id")

    @action_type_id.setter
    def action_type_id(self, value: pulumi.Input['PipelineActionTypeIdArgs']):
        pulumi.set(self, "action_type_id", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def configuration(self) -> Optional[Any]:
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: Optional[Any]):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter(name="inputArtifacts")
    def input_artifacts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineInputArtifactArgs']]]]:
        return pulumi.get(self, "input_artifacts")

    @input_artifacts.setter
    def input_artifacts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineInputArtifactArgs']]]]):
        pulumi.set(self, "input_artifacts", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="outputArtifacts")
    def output_artifacts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineOutputArtifactArgs']]]]:
        return pulumi.get(self, "output_artifacts")

    @output_artifacts.setter
    def output_artifacts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineOutputArtifactArgs']]]]):
        pulumi.set(self, "output_artifacts", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="runOrder")
    def run_order(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "run_order")

    @run_order.setter
    def run_order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "run_order", value)


@pulumi.input_type
class PipelineActionTypeIdArgs:
    def __init__(__self__, *,
                 category: pulumi.Input[str],
                 owner: pulumi.Input[str],
                 provider: pulumi.Input[str],
                 version: pulumi.Input[str]):
        pulumi.set(__self__, "category", category)
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "provider", provider)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def category(self) -> pulumi.Input[str]:
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: pulumi.Input[str]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Input[str]:
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: pulumi.Input[str]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter
    def provider(self) -> pulumi.Input[str]:
        return pulumi.get(self, "provider")

    @provider.setter
    def provider(self, value: pulumi.Input[str]):
        pulumi.set(self, "provider", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[str]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class PipelineArtifactStoreMapArgs:
    def __init__(__self__, *,
                 artifact_store: pulumi.Input['PipelineArtifactStoreArgs'],
                 region: pulumi.Input[str]):
        pulumi.set(__self__, "artifact_store", artifact_store)
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="artifactStore")
    def artifact_store(self) -> pulumi.Input['PipelineArtifactStoreArgs']:
        return pulumi.get(self, "artifact_store")

    @artifact_store.setter
    def artifact_store(self, value: pulumi.Input['PipelineArtifactStoreArgs']):
        pulumi.set(self, "artifact_store", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class PipelineArtifactStoreArgs:
    def __init__(__self__, *,
                 location: pulumi.Input[str],
                 type: pulumi.Input[str],
                 encryption_key: Optional[pulumi.Input['PipelineEncryptionKeyArgs']] = None):
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "type", type)
        if encryption_key is not None:
            pulumi.set(__self__, "encryption_key", encryption_key)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="encryptionKey")
    def encryption_key(self) -> Optional[pulumi.Input['PipelineEncryptionKeyArgs']]:
        return pulumi.get(self, "encryption_key")

    @encryption_key.setter
    def encryption_key(self, value: Optional[pulumi.Input['PipelineEncryptionKeyArgs']]):
        pulumi.set(self, "encryption_key", value)


@pulumi.input_type
class PipelineBlockerDeclarationArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 type: pulumi.Input[str]):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class PipelineEncryptionKeyArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 type: pulumi.Input[str]):
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class PipelineInputArtifactArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str]):
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PipelineOutputArtifactArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str]):
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PipelineStageDeclarationArgs:
    def __init__(__self__, *,
                 actions: pulumi.Input[Sequence[pulumi.Input['PipelineActionDeclarationArgs']]],
                 name: pulumi.Input[str],
                 blockers: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineBlockerDeclarationArgs']]]] = None):
        pulumi.set(__self__, "actions", actions)
        pulumi.set(__self__, "name", name)
        if blockers is not None:
            pulumi.set(__self__, "blockers", blockers)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Input[Sequence[pulumi.Input['PipelineActionDeclarationArgs']]]:
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: pulumi.Input[Sequence[pulumi.Input['PipelineActionDeclarationArgs']]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def blockers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineBlockerDeclarationArgs']]]]:
        return pulumi.get(self, "blockers")

    @blockers.setter
    def blockers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineBlockerDeclarationArgs']]]]):
        pulumi.set(self, "blockers", value)


@pulumi.input_type
class PipelineStageTransitionArgs:
    def __init__(__self__, *,
                 reason: pulumi.Input[str],
                 stage_name: pulumi.Input[str]):
        pulumi.set(__self__, "reason", reason)
        pulumi.set(__self__, "stage_name", stage_name)

    @property
    @pulumi.getter
    def reason(self) -> pulumi.Input[str]:
        return pulumi.get(self, "reason")

    @reason.setter
    def reason(self, value: pulumi.Input[str]):
        pulumi.set(self, "reason", value)

    @property
    @pulumi.getter(name="stageName")
    def stage_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "stage_name")

    @stage_name.setter
    def stage_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "stage_name", value)


@pulumi.input_type
class PipelineTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class WebhookAuthConfigurationArgs:
    def __init__(__self__, *,
                 allowed_ip_range: Optional[pulumi.Input[str]] = None,
                 secret_token: Optional[pulumi.Input[str]] = None):
        if allowed_ip_range is not None:
            pulumi.set(__self__, "allowed_ip_range", allowed_ip_range)
        if secret_token is not None:
            pulumi.set(__self__, "secret_token", secret_token)

    @property
    @pulumi.getter(name="allowedIPRange")
    def allowed_ip_range(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "allowed_ip_range")

    @allowed_ip_range.setter
    def allowed_ip_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allowed_ip_range", value)

    @property
    @pulumi.getter(name="secretToken")
    def secret_token(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secret_token")

    @secret_token.setter
    def secret_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_token", value)


@pulumi.input_type
class WebhookFilterRuleArgs:
    def __init__(__self__, *,
                 json_path: pulumi.Input[str],
                 match_equals: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "json_path", json_path)
        if match_equals is not None:
            pulumi.set(__self__, "match_equals", match_equals)

    @property
    @pulumi.getter(name="jsonPath")
    def json_path(self) -> pulumi.Input[str]:
        return pulumi.get(self, "json_path")

    @json_path.setter
    def json_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "json_path", value)

    @property
    @pulumi.getter(name="matchEquals")
    def match_equals(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "match_equals")

    @match_equals.setter
    def match_equals(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_equals", value)


