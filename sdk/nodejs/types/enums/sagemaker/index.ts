// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AppResourceSpecInstanceType = {
    System: "system",
    MlT3Micro: "ml.t3.micro",
    MlT3Small: "ml.t3.small",
    MlT3Medium: "ml.t3.medium",
    MlT3Large: "ml.t3.large",
    MlT3Xlarge: "ml.t3.xlarge",
    MlT32xlarge: "ml.t3.2xlarge",
    MlM5Large: "ml.m5.large",
    MlM5Xlarge: "ml.m5.xlarge",
    MlM52xlarge: "ml.m5.2xlarge",
    MlM54xlarge: "ml.m5.4xlarge",
    MlM58xlarge: "ml.m5.8xlarge",
    MlM512xlarge: "ml.m5.12xlarge",
    MlM516xlarge: "ml.m5.16xlarge",
    MlM524xlarge: "ml.m5.24xlarge",
    MlC5Large: "ml.c5.large",
    MlC5Xlarge: "ml.c5.xlarge",
    MlC52xlarge: "ml.c5.2xlarge",
    MlC54xlarge: "ml.c5.4xlarge",
    MlC59xlarge: "ml.c5.9xlarge",
    MlC512xlarge: "ml.c5.12xlarge",
    MlC518xlarge: "ml.c5.18xlarge",
    MlC524xlarge: "ml.c5.24xlarge",
    MlP32xlarge: "ml.p3.2xlarge",
    MlP38xlarge: "ml.p3.8xlarge",
    MlP316xlarge: "ml.p3.16xlarge",
    MlG4dnXlarge: "ml.g4dn.xlarge",
    MlG4dn2xlarge: "ml.g4dn.2xlarge",
    MlG4dn4xlarge: "ml.g4dn.4xlarge",
    MlG4dn8xlarge: "ml.g4dn.8xlarge",
    MlG4dn12xlarge: "ml.g4dn.12xlarge",
    MlG4dn16xlarge: "ml.g4dn.16xlarge",
    MlR5Large: "ml.r5.large",
    MlR5Xlarge: "ml.r5.xlarge",
    MlR52xlarge: "ml.r5.2xlarge",
    MlR54xlarge: "ml.r5.4xlarge",
    MlR58xlarge: "ml.r5.8xlarge",
    MlR512xlarge: "ml.r5.12xlarge",
    MlR516xlarge: "ml.r5.16xlarge",
    MlR524xlarge: "ml.r5.24xlarge",
    MlP3dn24xlarge: "ml.p3dn.24xlarge",
    MlM5dLarge: "ml.m5d.large",
    MlM5dXlarge: "ml.m5d.xlarge",
    MlM5d2xlarge: "ml.m5d.2xlarge",
    MlM5d4xlarge: "ml.m5d.4xlarge",
    MlM5d8xlarge: "ml.m5d.8xlarge",
    MlM5d12xlarge: "ml.m5d.12xlarge",
    MlM5d16xlarge: "ml.m5d.16xlarge",
    MlM5d24xlarge: "ml.m5d.24xlarge",
    MlG5Xlarge: "ml.g5.xlarge",
    MlG52xlarge: "ml.g5.2xlarge",
    MlG54xlarge: "ml.g5.4xlarge",
    MlG58xlarge: "ml.g5.8xlarge",
    MlG512xlarge: "ml.g5.12xlarge",
    MlG516xlarge: "ml.g5.16xlarge",
    MlG524xlarge: "ml.g5.24xlarge",
    MlG548xlarge: "ml.g5.48xlarge",
    MlP4d24xlarge: "ml.p4d.24xlarge",
    MlP4de24xlarge: "ml.p4de.24xlarge",
    MlGeospatialInteractive: "ml.geospatial.interactive",
} as const;

/**
 * The instance type that the image version runs on.
 */
export type AppResourceSpecInstanceType = (typeof AppResourceSpecInstanceType)[keyof typeof AppResourceSpecInstanceType];

export const AppType = {
    JupyterServer: "JupyterServer",
    KernelGateway: "KernelGateway",
    RStudioServerPro: "RStudioServerPro",
    RSessionGateway: "RSessionGateway",
    Canvas: "Canvas",
} as const;

/**
 * The type of app.
 */
export type AppType = (typeof AppType)[keyof typeof AppType];

export const DataQualityJobDefinitionBatchTransformInputS3DataDistributionType = {
    FullyReplicated: "FullyReplicated",
    ShardedByS3Key: "ShardedByS3Key",
} as const;

/**
 * Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
 */
export type DataQualityJobDefinitionBatchTransformInputS3DataDistributionType = (typeof DataQualityJobDefinitionBatchTransformInputS3DataDistributionType)[keyof typeof DataQualityJobDefinitionBatchTransformInputS3DataDistributionType];

export const DataQualityJobDefinitionBatchTransformInputS3InputMode = {
    Pipe: "Pipe",
    File: "File",
} as const;

/**
 * Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
 */
export type DataQualityJobDefinitionBatchTransformInputS3InputMode = (typeof DataQualityJobDefinitionBatchTransformInputS3InputMode)[keyof typeof DataQualityJobDefinitionBatchTransformInputS3InputMode];

export const DataQualityJobDefinitionEndpointInputS3DataDistributionType = {
    FullyReplicated: "FullyReplicated",
    ShardedByS3Key: "ShardedByS3Key",
} as const;

/**
 * Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
 */
export type DataQualityJobDefinitionEndpointInputS3DataDistributionType = (typeof DataQualityJobDefinitionEndpointInputS3DataDistributionType)[keyof typeof DataQualityJobDefinitionEndpointInputS3DataDistributionType];

export const DataQualityJobDefinitionEndpointInputS3InputMode = {
    Pipe: "Pipe",
    File: "File",
} as const;

/**
 * Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
 */
export type DataQualityJobDefinitionEndpointInputS3InputMode = (typeof DataQualityJobDefinitionEndpointInputS3InputMode)[keyof typeof DataQualityJobDefinitionEndpointInputS3InputMode];

export const DataQualityJobDefinitionS3OutputS3UploadMode = {
    Continuous: "Continuous",
    EndOfJob: "EndOfJob",
} as const;

/**
 * Whether to upload the results of the monitoring job continuously or after the job completes.
 */
export type DataQualityJobDefinitionS3OutputS3UploadMode = (typeof DataQualityJobDefinitionS3OutputS3UploadMode)[keyof typeof DataQualityJobDefinitionS3OutputS3UploadMode];

export const DomainAppNetworkAccessType = {
    PublicInternetOnly: "PublicInternetOnly",
    VpcOnly: "VpcOnly",
} as const;

/**
 * Specifies the VPC used for non-EFS traffic. The default value is PublicInternetOnly.
 */
export type DomainAppNetworkAccessType = (typeof DomainAppNetworkAccessType)[keyof typeof DomainAppNetworkAccessType];

export const DomainAppSecurityGroupManagement = {
    Service: "Service",
    Customer: "Customer",
} as const;

/**
 * The entity that creates and manages the required security groups for inter-app communication in VPCOnly mode. Required when CreateDomain.AppNetworkAccessType is VPCOnly and DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is provided.
 */
export type DomainAppSecurityGroupManagement = (typeof DomainAppSecurityGroupManagement)[keyof typeof DomainAppSecurityGroupManagement];

export const DomainAuthMode = {
    Sso: "SSO",
    Iam: "IAM",
} as const;

/**
 * The mode of authentication that members use to access the domain.
 */
export type DomainAuthMode = (typeof DomainAuthMode)[keyof typeof DomainAuthMode];

export const DomainRStudioServerProAppSettingsAccessStatus = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

/**
 * Indicates whether the current user has access to the RStudioServerPro app.
 */
export type DomainRStudioServerProAppSettingsAccessStatus = (typeof DomainRStudioServerProAppSettingsAccessStatus)[keyof typeof DomainRStudioServerProAppSettingsAccessStatus];

export const DomainRStudioServerProAppSettingsUserGroup = {
    RStudioAdmin: "R_STUDIO_ADMIN",
    RStudioUser: "R_STUDIO_USER",
} as const;

/**
 * The level of permissions that the user has within the RStudioServerPro app. This value defaults to User. The Admin value allows the user access to the RStudio Administrative Dashboard.
 */
export type DomainRStudioServerProAppSettingsUserGroup = (typeof DomainRStudioServerProAppSettingsUserGroup)[keyof typeof DomainRStudioServerProAppSettingsUserGroup];

export const DomainResourceSpecInstanceType = {
    System: "system",
    MlT3Micro: "ml.t3.micro",
    MlT3Small: "ml.t3.small",
    MlT3Medium: "ml.t3.medium",
    MlT3Large: "ml.t3.large",
    MlT3Xlarge: "ml.t3.xlarge",
    MlT32xlarge: "ml.t3.2xlarge",
    MlM5Large: "ml.m5.large",
    MlM5Xlarge: "ml.m5.xlarge",
    MlM52xlarge: "ml.m5.2xlarge",
    MlM54xlarge: "ml.m5.4xlarge",
    MlM58xlarge: "ml.m5.8xlarge",
    MlM512xlarge: "ml.m5.12xlarge",
    MlM516xlarge: "ml.m5.16xlarge",
    MlM524xlarge: "ml.m5.24xlarge",
    MlC5Large: "ml.c5.large",
    MlC5Xlarge: "ml.c5.xlarge",
    MlC52xlarge: "ml.c5.2xlarge",
    MlC54xlarge: "ml.c5.4xlarge",
    MlC59xlarge: "ml.c5.9xlarge",
    MlC512xlarge: "ml.c5.12xlarge",
    MlC518xlarge: "ml.c5.18xlarge",
    MlC524xlarge: "ml.c5.24xlarge",
    MlP32xlarge: "ml.p3.2xlarge",
    MlP38xlarge: "ml.p3.8xlarge",
    MlP316xlarge: "ml.p3.16xlarge",
    MlG4dnXlarge: "ml.g4dn.xlarge",
    MlG4dn2xlarge: "ml.g4dn.2xlarge",
    MlG4dn4xlarge: "ml.g4dn.4xlarge",
    MlG4dn8xlarge: "ml.g4dn.8xlarge",
    MlG4dn12xlarge: "ml.g4dn.12xlarge",
    MlG4dn16xlarge: "ml.g4dn.16xlarge",
    MlR5Large: "ml.r5.large",
    MlR5Xlarge: "ml.r5.xlarge",
    MlR52xlarge: "ml.r5.2xlarge",
    MlR54xlarge: "ml.r5.4xlarge",
    MlR58xlarge: "ml.r5.8xlarge",
    MlR512xlarge: "ml.r5.12xlarge",
    MlR516xlarge: "ml.r5.16xlarge",
    MlR524xlarge: "ml.r5.24xlarge",
    MlP3dn24xlarge: "ml.p3dn.24xlarge",
    MlM5dLarge: "ml.m5d.large",
    MlM5dXlarge: "ml.m5d.xlarge",
    MlM5d2xlarge: "ml.m5d.2xlarge",
    MlM5d4xlarge: "ml.m5d.4xlarge",
    MlM5d8xlarge: "ml.m5d.8xlarge",
    MlM5d12xlarge: "ml.m5d.12xlarge",
    MlM5d16xlarge: "ml.m5d.16xlarge",
    MlM5d24xlarge: "ml.m5d.24xlarge",
    MlG5Xlarge: "ml.g5.xlarge",
    MlG52xlarge: "ml.g5.2xlarge",
    MlG54xlarge: "ml.g5.4xlarge",
    MlG58xlarge: "ml.g5.8xlarge",
    MlG512xlarge: "ml.g5.12xlarge",
    MlG516xlarge: "ml.g5.16xlarge",
    MlG524xlarge: "ml.g5.24xlarge",
    MlG548xlarge: "ml.g5.48xlarge",
    MlP4d24xlarge: "ml.p4d.24xlarge",
    MlP4de24xlarge: "ml.p4de.24xlarge",
    MlGeospatialInteractive: "ml.geospatial.interactive",
} as const;

/**
 * The instance type that the image version runs on.
 */
export type DomainResourceSpecInstanceType = (typeof DomainResourceSpecInstanceType)[keyof typeof DomainResourceSpecInstanceType];

export const DomainSharingSettingsNotebookOutputOption = {
    Allowed: "Allowed",
    Disabled: "Disabled",
} as const;

/**
 * Whether to include the notebook cell output when sharing the notebook. The default is Disabled.
 */
export type DomainSharingSettingsNotebookOutputOption = (typeof DomainSharingSettingsNotebookOutputOption)[keyof typeof DomainSharingSettingsNotebookOutputOption];

export const FeatureGroupFeatureDefinitionFeatureType = {
    Integral: "Integral",
    Fractional: "Fractional",
    String: "String",
} as const;

export type FeatureGroupFeatureDefinitionFeatureType = (typeof FeatureGroupFeatureDefinitionFeatureType)[keyof typeof FeatureGroupFeatureDefinitionFeatureType];

export const FeatureGroupTableFormat = {
    Iceberg: "Iceberg",
    Glue: "Glue",
} as const;

/**
 * Format for the offline store feature group. Iceberg is the optimal format for feature groups shared between offline and online stores.
 */
export type FeatureGroupTableFormat = (typeof FeatureGroupTableFormat)[keyof typeof FeatureGroupTableFormat];

export const ImageVersionJobType = {
    Training: "TRAINING",
    Inference: "INFERENCE",
    NotebookKernel: "NOTEBOOK_KERNEL",
} as const;

/**
 * Indicates SageMaker job type compatibility.
 */
export type ImageVersionJobType = (typeof ImageVersionJobType)[keyof typeof ImageVersionJobType];

export const ImageVersionProcessor = {
    Cpu: "CPU",
    Gpu: "GPU",
} as const;

/**
 * Indicates CPU or GPU compatibility.
 */
export type ImageVersionProcessor = (typeof ImageVersionProcessor)[keyof typeof ImageVersionProcessor];

export const ImageVersionVendorGuidance = {
    NotProvided: "NOT_PROVIDED",
    Stable: "STABLE",
    ToBeArchived: "TO_BE_ARCHIVED",
    Archived: "ARCHIVED",
} as const;

/**
 * The availability of the image version specified by the maintainer.
 */
export type ImageVersionVendorGuidance = (typeof ImageVersionVendorGuidance)[keyof typeof ImageVersionVendorGuidance];

export const InferenceExperimentDesiredState = {
    Running: "Running",
    Completed: "Completed",
    Cancelled: "Cancelled",
} as const;

/**
 * The desired state of the experiment after starting or stopping operation.
 */
export type InferenceExperimentDesiredState = (typeof InferenceExperimentDesiredState)[keyof typeof InferenceExperimentDesiredState];

export const InferenceExperimentEndpointMetadataEndpointStatus = {
    Creating: "Creating",
    Updating: "Updating",
    SystemUpdating: "SystemUpdating",
    RollingBack: "RollingBack",
    InService: "InService",
    OutOfService: "OutOfService",
    Deleting: "Deleting",
    Failed: "Failed",
} as const;

/**
 * The status of the endpoint. For possible values of the status of an endpoint.
 */
export type InferenceExperimentEndpointMetadataEndpointStatus = (typeof InferenceExperimentEndpointMetadataEndpointStatus)[keyof typeof InferenceExperimentEndpointMetadataEndpointStatus];

export const InferenceExperimentModelInfrastructureConfigInfrastructureType = {
    RealTimeInference: "RealTimeInference",
} as const;

/**
 * The type of the inference experiment that you want to run.
 */
export type InferenceExperimentModelInfrastructureConfigInfrastructureType = (typeof InferenceExperimentModelInfrastructureConfigInfrastructureType)[keyof typeof InferenceExperimentModelInfrastructureConfigInfrastructureType];

export const InferenceExperimentStatus = {
    Creating: "Creating",
    Created: "Created",
    Updating: "Updating",
    Starting: "Starting",
    Stopping: "Stopping",
    Running: "Running",
    Completed: "Completed",
    Cancelled: "Cancelled",
} as const;

/**
 * The status of the inference experiment.
 */
export type InferenceExperimentStatus = (typeof InferenceExperimentStatus)[keyof typeof InferenceExperimentStatus];

export const InferenceExperimentType = {
    ShadowMode: "ShadowMode",
} as const;

/**
 * The type of the inference experiment that you want to run.
 */
export type InferenceExperimentType = (typeof InferenceExperimentType)[keyof typeof InferenceExperimentType];

export const ModelBiasJobDefinitionBatchTransformInputS3DataDistributionType = {
    FullyReplicated: "FullyReplicated",
    ShardedByS3Key: "ShardedByS3Key",
} as const;

/**
 * Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
 */
export type ModelBiasJobDefinitionBatchTransformInputS3DataDistributionType = (typeof ModelBiasJobDefinitionBatchTransformInputS3DataDistributionType)[keyof typeof ModelBiasJobDefinitionBatchTransformInputS3DataDistributionType];

export const ModelBiasJobDefinitionBatchTransformInputS3InputMode = {
    Pipe: "Pipe",
    File: "File",
} as const;

/**
 * Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
 */
export type ModelBiasJobDefinitionBatchTransformInputS3InputMode = (typeof ModelBiasJobDefinitionBatchTransformInputS3InputMode)[keyof typeof ModelBiasJobDefinitionBatchTransformInputS3InputMode];

export const ModelBiasJobDefinitionEndpointInputS3DataDistributionType = {
    FullyReplicated: "FullyReplicated",
    ShardedByS3Key: "ShardedByS3Key",
} as const;

/**
 * Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
 */
export type ModelBiasJobDefinitionEndpointInputS3DataDistributionType = (typeof ModelBiasJobDefinitionEndpointInputS3DataDistributionType)[keyof typeof ModelBiasJobDefinitionEndpointInputS3DataDistributionType];

export const ModelBiasJobDefinitionEndpointInputS3InputMode = {
    Pipe: "Pipe",
    File: "File",
} as const;

/**
 * Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
 */
export type ModelBiasJobDefinitionEndpointInputS3InputMode = (typeof ModelBiasJobDefinitionEndpointInputS3InputMode)[keyof typeof ModelBiasJobDefinitionEndpointInputS3InputMode];

export const ModelBiasJobDefinitionS3OutputS3UploadMode = {
    Continuous: "Continuous",
    EndOfJob: "EndOfJob",
} as const;

/**
 * Whether to upload the results of the monitoring job continuously or after the job completes.
 */
export type ModelBiasJobDefinitionS3OutputS3UploadMode = (typeof ModelBiasJobDefinitionS3OutputS3UploadMode)[keyof typeof ModelBiasJobDefinitionS3OutputS3UploadMode];

export const ModelCardBarChartMetricType = {
    BarChart: "bar_chart",
} as const;

export type ModelCardBarChartMetricType = (typeof ModelCardBarChartMetricType)[keyof typeof ModelCardBarChartMetricType];

export const ModelCardLinearGraphMetricType = {
    LinearGraph: "linear_graph",
} as const;

export type ModelCardLinearGraphMetricType = (typeof ModelCardLinearGraphMetricType)[keyof typeof ModelCardLinearGraphMetricType];

export const ModelCardMatrixMetricType = {
    Matrix: "matrix",
} as const;

export type ModelCardMatrixMetricType = (typeof ModelCardMatrixMetricType)[keyof typeof ModelCardMatrixMetricType];

export const ModelCardModelPackageDetailsModelApprovalStatus = {
    Approved: "Approved",
    Rejected: "Rejected",
    PendingManualApproval: "PendingManualApproval",
} as const;

/**
 * Current approval status of model package
 */
export type ModelCardModelPackageDetailsModelApprovalStatus = (typeof ModelCardModelPackageDetailsModelApprovalStatus)[keyof typeof ModelCardModelPackageDetailsModelApprovalStatus];

export const ModelCardModelPackageDetailsModelPackageStatus = {
    Pending: "Pending",
    InProgress: "InProgress",
    Completed: "Completed",
    Failed: "Failed",
    Deleting: "Deleting",
} as const;

/**
 * Current status of model package
 */
export type ModelCardModelPackageDetailsModelPackageStatus = (typeof ModelCardModelPackageDetailsModelPackageStatus)[keyof typeof ModelCardModelPackageDetailsModelPackageStatus];

export const ModelCardObjectiveFunctionFunctionPropertiesFunction = {
    Maximize: "Maximize",
    Minimize: "Minimize",
} as const;

export type ModelCardObjectiveFunctionFunctionPropertiesFunction = (typeof ModelCardObjectiveFunctionFunctionPropertiesFunction)[keyof typeof ModelCardObjectiveFunctionFunctionPropertiesFunction];

export const ModelCardProcessingStatus = {
    UnsetValue: "UnsetValue",
    DeleteInProgress: "DeleteInProgress",
    DeletePending: "DeletePending",
    ContentDeleted: "ContentDeleted",
    ExportJobsDeleted: "ExportJobsDeleted",
    DeleteCompleted: "DeleteCompleted",
    DeleteFailed: "DeleteFailed",
} as const;

/**
 * The processing status of model card deletion. The ModelCardProcessingStatus updates throughout the different deletion steps.
 */
export type ModelCardProcessingStatus = (typeof ModelCardProcessingStatus)[keyof typeof ModelCardProcessingStatus];

export const ModelCardRiskRating = {
    High: "High",
    Medium: "Medium",
    Low: "Low",
    Unknown: "Unknown",
} as const;

/**
 * Risk rating of model.
 */
export type ModelCardRiskRating = (typeof ModelCardRiskRating)[keyof typeof ModelCardRiskRating];

export const ModelCardSimpleMetricType = {
    Number: "number",
    String: "string",
    Boolean: "boolean",
} as const;

export type ModelCardSimpleMetricType = (typeof ModelCardSimpleMetricType)[keyof typeof ModelCardSimpleMetricType];

export const ModelCardStatus = {
    Draft: "Draft",
    PendingReview: "PendingReview",
    Approved: "Approved",
    Archived: "Archived",
} as const;

/**
 * The approval status of the model card within your organization. Different organizations might have different criteria for model card review and approval.
 */
export type ModelCardStatus = (typeof ModelCardStatus)[keyof typeof ModelCardStatus];

export const ModelExplainabilityJobDefinitionBatchTransformInputS3DataDistributionType = {
    FullyReplicated: "FullyReplicated",
    ShardedByS3Key: "ShardedByS3Key",
} as const;

/**
 * Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
 */
export type ModelExplainabilityJobDefinitionBatchTransformInputS3DataDistributionType = (typeof ModelExplainabilityJobDefinitionBatchTransformInputS3DataDistributionType)[keyof typeof ModelExplainabilityJobDefinitionBatchTransformInputS3DataDistributionType];

export const ModelExplainabilityJobDefinitionBatchTransformInputS3InputMode = {
    Pipe: "Pipe",
    File: "File",
} as const;

/**
 * Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
 */
export type ModelExplainabilityJobDefinitionBatchTransformInputS3InputMode = (typeof ModelExplainabilityJobDefinitionBatchTransformInputS3InputMode)[keyof typeof ModelExplainabilityJobDefinitionBatchTransformInputS3InputMode];

export const ModelExplainabilityJobDefinitionEndpointInputS3DataDistributionType = {
    FullyReplicated: "FullyReplicated",
    ShardedByS3Key: "ShardedByS3Key",
} as const;

/**
 * Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
 */
export type ModelExplainabilityJobDefinitionEndpointInputS3DataDistributionType = (typeof ModelExplainabilityJobDefinitionEndpointInputS3DataDistributionType)[keyof typeof ModelExplainabilityJobDefinitionEndpointInputS3DataDistributionType];

export const ModelExplainabilityJobDefinitionEndpointInputS3InputMode = {
    Pipe: "Pipe",
    File: "File",
} as const;

/**
 * Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
 */
export type ModelExplainabilityJobDefinitionEndpointInputS3InputMode = (typeof ModelExplainabilityJobDefinitionEndpointInputS3InputMode)[keyof typeof ModelExplainabilityJobDefinitionEndpointInputS3InputMode];

export const ModelExplainabilityJobDefinitionS3OutputS3UploadMode = {
    Continuous: "Continuous",
    EndOfJob: "EndOfJob",
} as const;

/**
 * Whether to upload the results of the monitoring job continuously or after the job completes.
 */
export type ModelExplainabilityJobDefinitionS3OutputS3UploadMode = (typeof ModelExplainabilityJobDefinitionS3OutputS3UploadMode)[keyof typeof ModelExplainabilityJobDefinitionS3OutputS3UploadMode];

export const ModelPackageGroupStatus = {
    Pending: "Pending",
    InProgress: "InProgress",
    Completed: "Completed",
    Failed: "Failed",
    Deleting: "Deleting",
    DeleteFailed: "DeleteFailed",
} as const;

/**
 * The status of a modelpackage group job.
 */
export type ModelPackageGroupStatus = (typeof ModelPackageGroupStatus)[keyof typeof ModelPackageGroupStatus];

export const ModelPackageModelApprovalStatus = {
    Approved: "Approved",
    Rejected: "Rejected",
    PendingManualApproval: "PendingManualApproval",
} as const;

/**
 * The approval status of the model package.
 */
export type ModelPackageModelApprovalStatus = (typeof ModelPackageModelApprovalStatus)[keyof typeof ModelPackageModelApprovalStatus];

export const ModelPackageS3DataSourceS3DataType = {
    ManifestFile: "ManifestFile",
    S3Prefix: "S3Prefix",
    AugmentedManifestFile: "AugmentedManifestFile",
} as const;

/**
 * The S3 Data Source Type
 */
export type ModelPackageS3DataSourceS3DataType = (typeof ModelPackageS3DataSourceS3DataType)[keyof typeof ModelPackageS3DataSourceS3DataType];

export const ModelPackageStatus = {
    Pending: "Pending",
    Deleting: "Deleting",
    InProgress: "InProgress",
    Completed: "Completed",
    Failed: "Failed",
} as const;

/**
 * The current status of the model package.
 */
export type ModelPackageStatus = (typeof ModelPackageStatus)[keyof typeof ModelPackageStatus];

export const ModelPackageStatusItemStatus = {
    NotStarted: "NotStarted",
    Failed: "Failed",
    InProgress: "InProgress",
    Completed: "Completed",
} as const;

/**
 * The current status.
 */
export type ModelPackageStatusItemStatus = (typeof ModelPackageStatusItemStatus)[keyof typeof ModelPackageStatusItemStatus];

export const ModelPackageTransformInputCompressionType = {
    None: "None",
    Gzip: "Gzip",
} as const;

/**
 * If your transform data is compressed, specify the compression type. Amazon SageMaker automatically decompresses the data for the transform job accordingly. The default value is None.
 */
export type ModelPackageTransformInputCompressionType = (typeof ModelPackageTransformInputCompressionType)[keyof typeof ModelPackageTransformInputCompressionType];

export const ModelPackageTransformInputSplitType = {
    None: "None",
    TFRecord: "TFRecord",
    Line: "Line",
    RecordIO: "RecordIO",
} as const;

/**
 * The method to use to split the transform job's data files into smaller batches. 
 */
export type ModelPackageTransformInputSplitType = (typeof ModelPackageTransformInputSplitType)[keyof typeof ModelPackageTransformInputSplitType];

export const ModelPackageTransformJobDefinitionBatchStrategy = {
    MultiRecord: "MultiRecord",
    SingleRecord: "SingleRecord",
} as const;

/**
 * A string that determines the number of records included in a single mini-batch.
 */
export type ModelPackageTransformJobDefinitionBatchStrategy = (typeof ModelPackageTransformJobDefinitionBatchStrategy)[keyof typeof ModelPackageTransformJobDefinitionBatchStrategy];

export const ModelPackageTransformOutputAssembleWith = {
    None: "None",
    Line: "Line",
} as const;

/**
 * Defines how to assemble the results of the transform job as a single S3 object.
 */
export type ModelPackageTransformOutputAssembleWith = (typeof ModelPackageTransformOutputAssembleWith)[keyof typeof ModelPackageTransformOutputAssembleWith];

export const ModelQualityJobDefinitionBatchTransformInputS3DataDistributionType = {
    FullyReplicated: "FullyReplicated",
    ShardedByS3Key: "ShardedByS3Key",
} as const;

/**
 * Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
 */
export type ModelQualityJobDefinitionBatchTransformInputS3DataDistributionType = (typeof ModelQualityJobDefinitionBatchTransformInputS3DataDistributionType)[keyof typeof ModelQualityJobDefinitionBatchTransformInputS3DataDistributionType];

export const ModelQualityJobDefinitionBatchTransformInputS3InputMode = {
    Pipe: "Pipe",
    File: "File",
} as const;

/**
 * Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
 */
export type ModelQualityJobDefinitionBatchTransformInputS3InputMode = (typeof ModelQualityJobDefinitionBatchTransformInputS3InputMode)[keyof typeof ModelQualityJobDefinitionBatchTransformInputS3InputMode];

export const ModelQualityJobDefinitionEndpointInputS3DataDistributionType = {
    FullyReplicated: "FullyReplicated",
    ShardedByS3Key: "ShardedByS3Key",
} as const;

/**
 * Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
 */
export type ModelQualityJobDefinitionEndpointInputS3DataDistributionType = (typeof ModelQualityJobDefinitionEndpointInputS3DataDistributionType)[keyof typeof ModelQualityJobDefinitionEndpointInputS3DataDistributionType];

export const ModelQualityJobDefinitionEndpointInputS3InputMode = {
    Pipe: "Pipe",
    File: "File",
} as const;

/**
 * Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
 */
export type ModelQualityJobDefinitionEndpointInputS3InputMode = (typeof ModelQualityJobDefinitionEndpointInputS3InputMode)[keyof typeof ModelQualityJobDefinitionEndpointInputS3InputMode];

export const ModelQualityJobDefinitionProblemType = {
    BinaryClassification: "BinaryClassification",
    MulticlassClassification: "MulticlassClassification",
    Regression: "Regression",
} as const;

/**
 * The status of the monitoring job.
 */
export type ModelQualityJobDefinitionProblemType = (typeof ModelQualityJobDefinitionProblemType)[keyof typeof ModelQualityJobDefinitionProblemType];

export const ModelQualityJobDefinitionS3OutputS3UploadMode = {
    Continuous: "Continuous",
    EndOfJob: "EndOfJob",
} as const;

/**
 * Whether to upload the results of the monitoring job continuously or after the job completes.
 */
export type ModelQualityJobDefinitionS3OutputS3UploadMode = (typeof ModelQualityJobDefinitionS3OutputS3UploadMode)[keyof typeof ModelQualityJobDefinitionS3OutputS3UploadMode];

export const MonitoringScheduleBatchTransformInputS3DataDistributionType = {
    FullyReplicated: "FullyReplicated",
    ShardedByS3Key: "ShardedByS3Key",
} as const;

/**
 * Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
 */
export type MonitoringScheduleBatchTransformInputS3DataDistributionType = (typeof MonitoringScheduleBatchTransformInputS3DataDistributionType)[keyof typeof MonitoringScheduleBatchTransformInputS3DataDistributionType];

export const MonitoringScheduleBatchTransformInputS3InputMode = {
    Pipe: "Pipe",
    File: "File",
} as const;

/**
 * Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
 */
export type MonitoringScheduleBatchTransformInputS3InputMode = (typeof MonitoringScheduleBatchTransformInputS3InputMode)[keyof typeof MonitoringScheduleBatchTransformInputS3InputMode];

export const MonitoringScheduleEndpointInputS3DataDistributionType = {
    FullyReplicated: "FullyReplicated",
    ShardedByS3Key: "ShardedByS3Key",
} as const;

/**
 * Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
 */
export type MonitoringScheduleEndpointInputS3DataDistributionType = (typeof MonitoringScheduleEndpointInputS3DataDistributionType)[keyof typeof MonitoringScheduleEndpointInputS3DataDistributionType];

export const MonitoringScheduleEndpointInputS3InputMode = {
    Pipe: "Pipe",
    File: "File",
} as const;

/**
 * Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
 */
export type MonitoringScheduleEndpointInputS3InputMode = (typeof MonitoringScheduleEndpointInputS3InputMode)[keyof typeof MonitoringScheduleEndpointInputS3InputMode];

export const MonitoringScheduleMonitoringExecutionSummaryMonitoringExecutionStatus = {
    Pending: "Pending",
    Completed: "Completed",
    CompletedWithViolations: "CompletedWithViolations",
    InProgress: "InProgress",
    Failed: "Failed",
    Stopping: "Stopping",
    Stopped: "Stopped",
} as const;

/**
 * The status of the monitoring job.
 */
export type MonitoringScheduleMonitoringExecutionSummaryMonitoringExecutionStatus = (typeof MonitoringScheduleMonitoringExecutionSummaryMonitoringExecutionStatus)[keyof typeof MonitoringScheduleMonitoringExecutionSummaryMonitoringExecutionStatus];

export const MonitoringScheduleMonitoringType = {
    DataQuality: "DataQuality",
    ModelQuality: "ModelQuality",
    ModelBias: "ModelBias",
    ModelExplainability: "ModelExplainability",
} as const;

/**
 * The type of monitoring job.
 */
export type MonitoringScheduleMonitoringType = (typeof MonitoringScheduleMonitoringType)[keyof typeof MonitoringScheduleMonitoringType];

export const MonitoringScheduleS3OutputS3UploadMode = {
    Continuous: "Continuous",
    EndOfJob: "EndOfJob",
} as const;

/**
 * Whether to upload the results of the monitoring job continuously or after the job completes.
 */
export type MonitoringScheduleS3OutputS3UploadMode = (typeof MonitoringScheduleS3OutputS3UploadMode)[keyof typeof MonitoringScheduleS3OutputS3UploadMode];

export const MonitoringScheduleStatus = {
    Pending: "Pending",
    Failed: "Failed",
    Scheduled: "Scheduled",
    Stopped: "Stopped",
} as const;

/**
 * The status of a schedule job.
 */
export type MonitoringScheduleStatus = (typeof MonitoringScheduleStatus)[keyof typeof MonitoringScheduleStatus];

export const ProjectStatus = {
    Pending: "Pending",
    CreateInProgress: "CreateInProgress",
    CreateCompleted: "CreateCompleted",
    CreateFailed: "CreateFailed",
    DeleteInProgress: "DeleteInProgress",
    DeleteFailed: "DeleteFailed",
    DeleteCompleted: "DeleteCompleted",
} as const;

/**
 * The status of a project.
 */
export type ProjectStatus = (typeof ProjectStatus)[keyof typeof ProjectStatus];

export const SpaceResourceSpecInstanceType = {
    System: "system",
    MlT3Micro: "ml.t3.micro",
    MlT3Small: "ml.t3.small",
    MlT3Medium: "ml.t3.medium",
    MlT3Large: "ml.t3.large",
    MlT3Xlarge: "ml.t3.xlarge",
    MlT32xlarge: "ml.t3.2xlarge",
    MlM5Large: "ml.m5.large",
    MlM5Xlarge: "ml.m5.xlarge",
    MlM52xlarge: "ml.m5.2xlarge",
    MlM54xlarge: "ml.m5.4xlarge",
    MlM58xlarge: "ml.m5.8xlarge",
    MlM512xlarge: "ml.m5.12xlarge",
    MlM516xlarge: "ml.m5.16xlarge",
    MlM524xlarge: "ml.m5.24xlarge",
    MlC5Large: "ml.c5.large",
    MlC5Xlarge: "ml.c5.xlarge",
    MlC52xlarge: "ml.c5.2xlarge",
    MlC54xlarge: "ml.c5.4xlarge",
    MlC59xlarge: "ml.c5.9xlarge",
    MlC512xlarge: "ml.c5.12xlarge",
    MlC518xlarge: "ml.c5.18xlarge",
    MlC524xlarge: "ml.c5.24xlarge",
    MlP32xlarge: "ml.p3.2xlarge",
    MlP38xlarge: "ml.p3.8xlarge",
    MlP316xlarge: "ml.p3.16xlarge",
    MlG4dnXlarge: "ml.g4dn.xlarge",
    MlG4dn2xlarge: "ml.g4dn.2xlarge",
    MlG4dn4xlarge: "ml.g4dn.4xlarge",
    MlG4dn8xlarge: "ml.g4dn.8xlarge",
    MlG4dn12xlarge: "ml.g4dn.12xlarge",
    MlG4dn16xlarge: "ml.g4dn.16xlarge",
    MlR5Large: "ml.r5.large",
    MlR5Xlarge: "ml.r5.xlarge",
    MlR52xlarge: "ml.r5.2xlarge",
    MlR54xlarge: "ml.r5.4xlarge",
    MlR58xlarge: "ml.r5.8xlarge",
    MlR512xlarge: "ml.r5.12xlarge",
    MlR516xlarge: "ml.r5.16xlarge",
    MlR524xlarge: "ml.r5.24xlarge",
    MlP3dn24xlarge: "ml.p3dn.24xlarge",
    MlM5dLarge: "ml.m5d.large",
    MlM5dXlarge: "ml.m5d.xlarge",
    MlM5d2xlarge: "ml.m5d.2xlarge",
    MlM5d4xlarge: "ml.m5d.4xlarge",
    MlM5d8xlarge: "ml.m5d.8xlarge",
    MlM5d12xlarge: "ml.m5d.12xlarge",
    MlM5d16xlarge: "ml.m5d.16xlarge",
    MlM5d24xlarge: "ml.m5d.24xlarge",
    MlG5Xlarge: "ml.g5.xlarge",
    MlG52xlarge: "ml.g5.2xlarge",
    MlG54xlarge: "ml.g5.4xlarge",
    MlG58xlarge: "ml.g5.8xlarge",
    MlG512xlarge: "ml.g5.12xlarge",
    MlG516xlarge: "ml.g5.16xlarge",
    MlG524xlarge: "ml.g5.24xlarge",
    MlG548xlarge: "ml.g5.48xlarge",
    MlP4d24xlarge: "ml.p4d.24xlarge",
    MlP4de24xlarge: "ml.p4de.24xlarge",
    MlGeospatialInteractive: "ml.geospatial.interactive",
} as const;

/**
 * The instance type that the image version runs on.
 */
export type SpaceResourceSpecInstanceType = (typeof SpaceResourceSpecInstanceType)[keyof typeof SpaceResourceSpecInstanceType];

export const UserProfileRStudioServerProAppSettingsAccessStatus = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

/**
 * Indicates whether the current user has access to the RStudioServerPro app.
 */
export type UserProfileRStudioServerProAppSettingsAccessStatus = (typeof UserProfileRStudioServerProAppSettingsAccessStatus)[keyof typeof UserProfileRStudioServerProAppSettingsAccessStatus];

export const UserProfileRStudioServerProAppSettingsUserGroup = {
    RStudioAdmin: "R_STUDIO_ADMIN",
    RStudioUser: "R_STUDIO_USER",
} as const;

/**
 * The level of permissions that the user has within the RStudioServerPro app. This value defaults to User. The Admin value allows the user access to the RStudio Administrative Dashboard.
 */
export type UserProfileRStudioServerProAppSettingsUserGroup = (typeof UserProfileRStudioServerProAppSettingsUserGroup)[keyof typeof UserProfileRStudioServerProAppSettingsUserGroup];

export const UserProfileResourceSpecInstanceType = {
    System: "system",
    MlT3Micro: "ml.t3.micro",
    MlT3Small: "ml.t3.small",
    MlT3Medium: "ml.t3.medium",
    MlT3Large: "ml.t3.large",
    MlT3Xlarge: "ml.t3.xlarge",
    MlT32xlarge: "ml.t3.2xlarge",
    MlM5Large: "ml.m5.large",
    MlM5Xlarge: "ml.m5.xlarge",
    MlM52xlarge: "ml.m5.2xlarge",
    MlM54xlarge: "ml.m5.4xlarge",
    MlM58xlarge: "ml.m5.8xlarge",
    MlM512xlarge: "ml.m5.12xlarge",
    MlM516xlarge: "ml.m5.16xlarge",
    MlM524xlarge: "ml.m5.24xlarge",
    MlC5Large: "ml.c5.large",
    MlC5Xlarge: "ml.c5.xlarge",
    MlC52xlarge: "ml.c5.2xlarge",
    MlC54xlarge: "ml.c5.4xlarge",
    MlC59xlarge: "ml.c5.9xlarge",
    MlC512xlarge: "ml.c5.12xlarge",
    MlC518xlarge: "ml.c5.18xlarge",
    MlC524xlarge: "ml.c5.24xlarge",
    MlP32xlarge: "ml.p3.2xlarge",
    MlP38xlarge: "ml.p3.8xlarge",
    MlP316xlarge: "ml.p3.16xlarge",
    MlG4dnXlarge: "ml.g4dn.xlarge",
    MlG4dn2xlarge: "ml.g4dn.2xlarge",
    MlG4dn4xlarge: "ml.g4dn.4xlarge",
    MlG4dn8xlarge: "ml.g4dn.8xlarge",
    MlG4dn12xlarge: "ml.g4dn.12xlarge",
    MlG4dn16xlarge: "ml.g4dn.16xlarge",
    MlR5Large: "ml.r5.large",
    MlR5Xlarge: "ml.r5.xlarge",
    MlR52xlarge: "ml.r5.2xlarge",
    MlR54xlarge: "ml.r5.4xlarge",
    MlR58xlarge: "ml.r5.8xlarge",
    MlR512xlarge: "ml.r5.12xlarge",
    MlR516xlarge: "ml.r5.16xlarge",
    MlR524xlarge: "ml.r5.24xlarge",
    MlP3dn24xlarge: "ml.p3dn.24xlarge",
    MlM5dLarge: "ml.m5d.large",
    MlM5dXlarge: "ml.m5d.xlarge",
    MlM5d2xlarge: "ml.m5d.2xlarge",
    MlM5d4xlarge: "ml.m5d.4xlarge",
    MlM5d8xlarge: "ml.m5d.8xlarge",
    MlM5d12xlarge: "ml.m5d.12xlarge",
    MlM5d16xlarge: "ml.m5d.16xlarge",
    MlM5d24xlarge: "ml.m5d.24xlarge",
    MlG5Xlarge: "ml.g5.xlarge",
    MlG52xlarge: "ml.g5.2xlarge",
    MlG54xlarge: "ml.g5.4xlarge",
    MlG58xlarge: "ml.g5.8xlarge",
    MlG512xlarge: "ml.g5.12xlarge",
    MlG516xlarge: "ml.g5.16xlarge",
    MlG524xlarge: "ml.g5.24xlarge",
    MlG548xlarge: "ml.g5.48xlarge",
    MlP4d24xlarge: "ml.p4d.24xlarge",
    MlP4de24xlarge: "ml.p4de.24xlarge",
    MlGeospatialInteractive: "ml.geospatial.interactive",
} as const;

/**
 * The instance type that the image version runs on.
 */
export type UserProfileResourceSpecInstanceType = (typeof UserProfileResourceSpecInstanceType)[keyof typeof UserProfileResourceSpecInstanceType];

export const UserProfileSharingSettingsNotebookOutputOption = {
    Allowed: "Allowed",
    Disabled: "Disabled",
} as const;

/**
 * Whether to include the notebook cell output when sharing the notebook. The default is Disabled.
 */
export type UserProfileSharingSettingsNotebookOutputOption = (typeof UserProfileSharingSettingsNotebookOutputOption)[keyof typeof UserProfileSharingSettingsNotebookOutputOption];
