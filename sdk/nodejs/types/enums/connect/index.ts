// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ContactFlowModuleState = {
    Active: "ACTIVE",
    Archived: "ARCHIVED",
} as const;

/**
 * The state of the contact flow module.
 */
export type ContactFlowModuleState = (typeof ContactFlowModuleState)[keyof typeof ContactFlowModuleState];

export const ContactFlowModuleStatus = {
    Published: "PUBLISHED",
    Saved: "SAVED",
} as const;

/**
 * The status of the contact flow module.
 */
export type ContactFlowModuleStatus = (typeof ContactFlowModuleStatus)[keyof typeof ContactFlowModuleStatus];

export const ContactFlowState = {
    Active: "ACTIVE",
    Archived: "ARCHIVED",
} as const;

/**
 * The state of the contact flow.
 */
export type ContactFlowState = (typeof ContactFlowState)[keyof typeof ContactFlowState];

export const ContactFlowType = {
    ContactFlow: "CONTACT_FLOW",
    CustomerQueue: "CUSTOMER_QUEUE",
    CustomerHold: "CUSTOMER_HOLD",
    CustomerWhisper: "CUSTOMER_WHISPER",
    AgentHold: "AGENT_HOLD",
    AgentWhisper: "AGENT_WHISPER",
    OutboundWhisper: "OUTBOUND_WHISPER",
    AgentTransfer: "AGENT_TRANSFER",
    QueueTransfer: "QUEUE_TRANSFER",
} as const;

/**
 * The type of the contact flow.
 */
export type ContactFlowType = (typeof ContactFlowType)[keyof typeof ContactFlowType];

export const HoursOfOperationConfigDay = {
    Sunday: "SUNDAY",
    Monday: "MONDAY",
    Tuesday: "TUESDAY",
    Wednesday: "WEDNESDAY",
    Thursday: "THURSDAY",
    Friday: "FRIDAY",
    Saturday: "SATURDAY",
} as const;

/**
 * The day that the hours of operation applies to.
 */
export type HoursOfOperationConfigDay = (typeof HoursOfOperationConfigDay)[keyof typeof HoursOfOperationConfigDay];

export const InstanceIdentityManagementType = {
    Saml: "SAML",
    ConnectManaged: "CONNECT_MANAGED",
    ExistingDirectory: "EXISTING_DIRECTORY",
} as const;

/**
 * Specifies the type of directory integration for new instance.
 */
export type InstanceIdentityManagementType = (typeof InstanceIdentityManagementType)[keyof typeof InstanceIdentityManagementType];

export const InstanceStatus = {
    CreationInProgress: "CREATION_IN_PROGRESS",
    CreationFailed: "CREATION_FAILED",
    Active: "ACTIVE",
} as const;

/**
 * Specifies the creation status of new instance.
 */
export type InstanceStatus = (typeof InstanceStatus)[keyof typeof InstanceStatus];

export const InstanceStorageConfigEncryptionType = {
    Kms: "KMS",
} as const;

/**
 * Specifies default encryption using AWS KMS-Managed Keys
 */
export type InstanceStorageConfigEncryptionType = (typeof InstanceStorageConfigEncryptionType)[keyof typeof InstanceStorageConfigEncryptionType];

export const InstanceStorageConfigInstanceStorageResourceType = {
    ChatTranscripts: "CHAT_TRANSCRIPTS",
    CallRecordings: "CALL_RECORDINGS",
    ScheduledReports: "SCHEDULED_REPORTS",
    MediaStreams: "MEDIA_STREAMS",
    ContactTraceRecords: "CONTACT_TRACE_RECORDS",
    AgentEvents: "AGENT_EVENTS",
} as const;

/**
 * Specifies the type of storage resource available for the instance
 */
export type InstanceStorageConfigInstanceStorageResourceType = (typeof InstanceStorageConfigInstanceStorageResourceType)[keyof typeof InstanceStorageConfigInstanceStorageResourceType];

export const InstanceStorageConfigStorageType = {
    S3: "S3",
    KinesisVideoStream: "KINESIS_VIDEO_STREAM",
    KinesisStream: "KINESIS_STREAM",
    KinesisFirehose: "KINESIS_FIREHOSE",
} as const;

/**
 * Specifies the storage type to be associated with the instance
 */
export type InstanceStorageConfigStorageType = (typeof InstanceStorageConfigStorageType)[keyof typeof InstanceStorageConfigStorageType];

export const QuickConnectType = {
    PhoneNumber: "PHONE_NUMBER",
    Queue: "QUEUE",
    User: "USER",
} as const;

/**
 * The type of quick connect. In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE).
 */
export type QuickConnectType = (typeof QuickConnectType)[keyof typeof QuickConnectType];

export const RulePublishStatus = {
    Draft: "DRAFT",
    Published: "PUBLISHED",
} as const;

/**
 * The publish status of a rule, either draft or published.
 */
export type RulePublishStatus = (typeof RulePublishStatus)[keyof typeof RulePublishStatus];

export const RuleSendNotificationActionContentType = {
    PlainText: "PLAIN_TEXT",
} as const;

/**
 * The type of content.
 */
export type RuleSendNotificationActionContentType = (typeof RuleSendNotificationActionContentType)[keyof typeof RuleSendNotificationActionContentType];

export const RuleSendNotificationActionDeliveryMethod = {
    Email: "EMAIL",
} as const;

/**
 * The means of delivery.
 */
export type RuleSendNotificationActionDeliveryMethod = (typeof RuleSendNotificationActionDeliveryMethod)[keyof typeof RuleSendNotificationActionDeliveryMethod];

export const RuleTriggerEventSourceEventSourceName = {
    OnPostCallAnalysisAvailable: "OnPostCallAnalysisAvailable",
    OnRealTimeCallAnalysisAvailable: "OnRealTimeCallAnalysisAvailable",
    OnPostChatAnalysisAvailable: "OnPostChatAnalysisAvailable",
    OnZendeskTicketCreate: "OnZendeskTicketCreate",
    OnZendeskTicketStatusUpdate: "OnZendeskTicketStatusUpdate",
    OnSalesforceCaseCreate: "OnSalesforceCaseCreate",
} as const;

/**
 * The name of event source.
 */
export type RuleTriggerEventSourceEventSourceName = (typeof RuleTriggerEventSourceEventSourceName)[keyof typeof RuleTriggerEventSourceEventSourceName];

export const TaskTemplateFieldType = {
    Name: "NAME",
    Description: "DESCRIPTION",
    ScheduledTime: "SCHEDULED_TIME",
    QuickConnect: "QUICK_CONNECT",
    Url: "URL",
    Number: "NUMBER",
    Text: "TEXT",
    TextArea: "TEXT_AREA",
    DateTime: "DATE_TIME",
    Boolean: "BOOLEAN",
    SingleSelect: "SINGLE_SELECT",
    Email: "EMAIL",
} as const;

/**
 * The type of the task template's field
 */
export type TaskTemplateFieldType = (typeof TaskTemplateFieldType)[keyof typeof TaskTemplateFieldType];

export const TaskTemplateStatus = {
    Active: "ACTIVE",
    Inactive: "INACTIVE",
} as const;

/**
 * The status of the task template
 */
export type TaskTemplateStatus = (typeof TaskTemplateStatus)[keyof typeof TaskTemplateStatus];

export const UserPhoneType = {
    SoftPhone: "SOFT_PHONE",
    DeskPhone: "DESK_PHONE",
} as const;

/**
 * The phone type.
 */
export type UserPhoneType = (typeof UserPhoneType)[keyof typeof UserPhoneType];
