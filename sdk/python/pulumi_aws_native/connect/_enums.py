# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ContactFlowState',
    'ContactFlowType',
    'EvaluationFormNumericQuestionPropertyValueAutomationLabel',
    'EvaluationFormQuestionQuestionType',
    'EvaluationFormScoringStrategyMode',
    'EvaluationFormScoringStrategyStatus',
    'EvaluationFormSingleSelectQuestionPropertiesDisplayAs',
    'EvaluationFormSingleSelectQuestionRuleCategoryAutomationCondition',
    'EvaluationFormStatus',
    'HoursOfOperationConfigDay',
    'InstanceIdentityManagementType',
    'InstanceStatus',
    'InstanceStorageConfigEncryptionType',
    'InstanceStorageConfigInstanceStorageResourceType',
    'InstanceStorageConfigStorageType',
    'IntegrationAssociationIntegrationType',
    'QueueStatus',
    'QueueType',
    'QuickConnectType',
    'RoutingProfileAgentAvailabilityTimer',
    'RoutingProfileBehaviorType',
    'RoutingProfileChannel',
    'RulePublishStatus',
    'RuleSendNotificationActionContentType',
    'RuleSendNotificationActionDeliveryMethod',
    'RuleTriggerEventSourceEventSourceName',
    'TaskTemplateFieldType',
    'TaskTemplateStatus',
    'TrafficDistributionGroupStatus',
    'UserPhoneType',
]


class ContactFlowState(str, Enum):
    """
    The state of the contact flow.
    """
    ACTIVE = "ACTIVE"
    ARCHIVED = "ARCHIVED"


class ContactFlowType(str, Enum):
    """
    The type of the contact flow.
    """
    CONTACT_FLOW = "CONTACT_FLOW"
    CUSTOMER_QUEUE = "CUSTOMER_QUEUE"
    CUSTOMER_HOLD = "CUSTOMER_HOLD"
    CUSTOMER_WHISPER = "CUSTOMER_WHISPER"
    AGENT_HOLD = "AGENT_HOLD"
    AGENT_WHISPER = "AGENT_WHISPER"
    OUTBOUND_WHISPER = "OUTBOUND_WHISPER"
    AGENT_TRANSFER = "AGENT_TRANSFER"
    QUEUE_TRANSFER = "QUEUE_TRANSFER"


class EvaluationFormNumericQuestionPropertyValueAutomationLabel(str, Enum):
    """
    The automation property label.
    """
    OVERALL_CUSTOMER_SENTIMENT_SCORE = "OVERALL_CUSTOMER_SENTIMENT_SCORE"
    OVERALL_AGENT_SENTIMENT_SCORE = "OVERALL_AGENT_SENTIMENT_SCORE"
    NON_TALK_TIME = "NON_TALK_TIME"
    NON_TALK_TIME_PERCENTAGE = "NON_TALK_TIME_PERCENTAGE"
    NUMBER_OF_INTERRUPTIONS = "NUMBER_OF_INTERRUPTIONS"
    CONTACT_DURATION = "CONTACT_DURATION"
    AGENT_INTERACTION_DURATION = "AGENT_INTERACTION_DURATION"
    CUSTOMER_HOLD_TIME = "CUSTOMER_HOLD_TIME"


class EvaluationFormQuestionQuestionType(str, Enum):
    """
    The type of the question.
    """
    NUMERIC = "NUMERIC"
    SINGLESELECT = "SINGLESELECT"
    TEXT = "TEXT"


class EvaluationFormScoringStrategyMode(str, Enum):
    """
    The scoring mode.
    """
    QUESTION_ONLY = "QUESTION_ONLY"
    SECTION_ONLY = "SECTION_ONLY"


class EvaluationFormScoringStrategyStatus(str, Enum):
    """
    The scoring status.
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class EvaluationFormSingleSelectQuestionPropertiesDisplayAs(str, Enum):
    """
    The display mode of the single-select question.
    """
    DROPDOWN = "DROPDOWN"
    RADIO = "RADIO"


class EvaluationFormSingleSelectQuestionRuleCategoryAutomationCondition(str, Enum):
    """
    The automation condition applied on contact categories.
    """
    PRESENT = "PRESENT"
    NOT_PRESENT = "NOT_PRESENT"


class EvaluationFormStatus(str, Enum):
    """
    The status of the evaluation form.
    """
    DRAFT = "DRAFT"
    ACTIVE = "ACTIVE"


class HoursOfOperationConfigDay(str, Enum):
    """
    The day that the hours of operation applies to.
    """
    SUNDAY = "SUNDAY"
    MONDAY = "MONDAY"
    TUESDAY = "TUESDAY"
    WEDNESDAY = "WEDNESDAY"
    THURSDAY = "THURSDAY"
    FRIDAY = "FRIDAY"
    SATURDAY = "SATURDAY"


class InstanceIdentityManagementType(str, Enum):
    """
    Specifies the type of directory integration for new instance.
    """
    SAML = "SAML"
    CONNECT_MANAGED = "CONNECT_MANAGED"
    EXISTING_DIRECTORY = "EXISTING_DIRECTORY"


class InstanceStatus(str, Enum):
    """
    Specifies the creation status of new instance.
    """
    CREATION_IN_PROGRESS = "CREATION_IN_PROGRESS"
    CREATION_FAILED = "CREATION_FAILED"
    ACTIVE = "ACTIVE"


class InstanceStorageConfigEncryptionType(str, Enum):
    """
    Specifies default encryption using AWS KMS-Managed Keys
    """
    KMS = "KMS"


class InstanceStorageConfigInstanceStorageResourceType(str, Enum):
    """
    Specifies the type of storage resource available for the instance
    """
    CHAT_TRANSCRIPTS = "CHAT_TRANSCRIPTS"
    CALL_RECORDINGS = "CALL_RECORDINGS"
    SCHEDULED_REPORTS = "SCHEDULED_REPORTS"
    MEDIA_STREAMS = "MEDIA_STREAMS"
    CONTACT_TRACE_RECORDS = "CONTACT_TRACE_RECORDS"
    AGENT_EVENTS = "AGENT_EVENTS"


class InstanceStorageConfigStorageType(str, Enum):
    """
    Specifies the storage type to be associated with the instance
    """
    S3 = "S3"
    KINESIS_VIDEO_STREAM = "KINESIS_VIDEO_STREAM"
    KINESIS_STREAM = "KINESIS_STREAM"
    KINESIS_FIREHOSE = "KINESIS_FIREHOSE"


class IntegrationAssociationIntegrationType(str, Enum):
    """
    Specifies the integration type to be associated with the instance
    """
    LEX_BOT = "LEX_BOT"
    LAMBDA_FUNCTION = "LAMBDA_FUNCTION"


class QueueStatus(str, Enum):
    """
    The status of the queue.
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class QueueType(str, Enum):
    """
    The type of queue.
    """
    STANDARD = "STANDARD"
    AGENT = "AGENT"


class QuickConnectType(str, Enum):
    """
    The type of quick connect. In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE).
    """
    PHONE_NUMBER = "PHONE_NUMBER"
    QUEUE = "QUEUE"
    USER = "USER"


class RoutingProfileAgentAvailabilityTimer(str, Enum):
    """
    Whether agents with this routing profile will have their routing order calculated based on longest idle time or time since their last inbound contact.
    """
    TIME_SINCE_LAST_ACTIVITY = "TIME_SINCE_LAST_ACTIVITY"
    TIME_SINCE_LAST_INBOUND = "TIME_SINCE_LAST_INBOUND"


class RoutingProfileBehaviorType(str, Enum):
    """
    Specifies the other channels that can be routed to an agent handling their current channel.
    """
    ROUTE_CURRENT_CHANNEL_ONLY = "ROUTE_CURRENT_CHANNEL_ONLY"
    ROUTE_ANY_CHANNEL = "ROUTE_ANY_CHANNEL"


class RoutingProfileChannel(str, Enum):
    """
    The channels that agents can handle in the Contact Control Panel (CCP).
    """
    VOICE = "VOICE"
    CHAT = "CHAT"
    TASK = "TASK"


class RulePublishStatus(str, Enum):
    """
    The publish status of a rule, either draft or published.
    """
    DRAFT = "DRAFT"
    PUBLISHED = "PUBLISHED"


class RuleSendNotificationActionContentType(str, Enum):
    """
    The type of content.
    """
    PLAIN_TEXT = "PLAIN_TEXT"


class RuleSendNotificationActionDeliveryMethod(str, Enum):
    """
    The means of delivery.
    """
    EMAIL = "EMAIL"


class RuleTriggerEventSourceEventSourceName(str, Enum):
    """
    The name of event source.
    """
    ON_CONTACT_EVALUATION_SUBMIT = "OnContactEvaluationSubmit"
    ON_POST_CALL_ANALYSIS_AVAILABLE = "OnPostCallAnalysisAvailable"
    ON_REAL_TIME_CALL_ANALYSIS_AVAILABLE = "OnRealTimeCallAnalysisAvailable"
    ON_POST_CHAT_ANALYSIS_AVAILABLE = "OnPostChatAnalysisAvailable"
    ON_ZENDESK_TICKET_CREATE = "OnZendeskTicketCreate"
    ON_ZENDESK_TICKET_STATUS_UPDATE = "OnZendeskTicketStatusUpdate"
    ON_SALESFORCE_CASE_CREATE = "OnSalesforceCaseCreate"
    ON_METRIC_DATA_UPDATE = "OnMetricDataUpdate"


class TaskTemplateFieldType(str, Enum):
    """
    The type of the task template's field
    """
    NAME = "NAME"
    DESCRIPTION = "DESCRIPTION"
    SCHEDULED_TIME = "SCHEDULED_TIME"
    QUICK_CONNECT = "QUICK_CONNECT"
    URL = "URL"
    NUMBER = "NUMBER"
    TEXT = "TEXT"
    TEXT_AREA = "TEXT_AREA"
    DATE_TIME = "DATE_TIME"
    BOOLEAN = "BOOLEAN"
    SINGLE_SELECT = "SINGLE_SELECT"
    EMAIL = "EMAIL"


class TaskTemplateStatus(str, Enum):
    """
    The status of the task template
    """
    ACTIVE = "ACTIVE"
    INACTIVE = "INACTIVE"


class TrafficDistributionGroupStatus(str, Enum):
    """
    The status of the traffic distribution group.
    """
    CREATION_IN_PROGRESS = "CREATION_IN_PROGRESS"
    ACTIVE = "ACTIVE"
    CREATION_FAILED = "CREATION_FAILED"
    PENDING_DELETION = "PENDING_DELETION"
    DELETION_FAILED = "DELETION_FAILED"
    UPDATE_IN_PROGRESS = "UPDATE_IN_PROGRESS"


class UserPhoneType(str, Enum):
    """
    The phone type.
    """
    SOFT_PHONE = "SOFT_PHONE"
    DESK_PHONE = "DESK_PHONE"
