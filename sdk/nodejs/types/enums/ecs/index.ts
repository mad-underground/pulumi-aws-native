// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const CapacityProviderAutoScalingGroupProviderManagedDraining = {
    Disabled: "DISABLED",
    Enabled: "ENABLED",
} as const;

export type CapacityProviderAutoScalingGroupProviderManagedDraining = (typeof CapacityProviderAutoScalingGroupProviderManagedDraining)[keyof typeof CapacityProviderAutoScalingGroupProviderManagedDraining];

export const CapacityProviderAutoScalingGroupProviderManagedTerminationProtection = {
    Disabled: "DISABLED",
    Enabled: "ENABLED",
} as const;

export type CapacityProviderAutoScalingGroupProviderManagedTerminationProtection = (typeof CapacityProviderAutoScalingGroupProviderManagedTerminationProtection)[keyof typeof CapacityProviderAutoScalingGroupProviderManagedTerminationProtection];

export const CapacityProviderManagedScalingStatus = {
    Disabled: "DISABLED",
    Enabled: "ENABLED",
} as const;

export type CapacityProviderManagedScalingStatus = (typeof CapacityProviderManagedScalingStatus)[keyof typeof CapacityProviderManagedScalingStatus];

export const ClusterCapacityProviderAssociationsCapacityProvider = {
    Fargate: "FARGATE",
    FargateSpot: "FARGATE_SPOT",
} as const;

/**
 * If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.
 */
export type ClusterCapacityProviderAssociationsCapacityProvider = (typeof ClusterCapacityProviderAssociationsCapacityProvider)[keyof typeof ClusterCapacityProviderAssociationsCapacityProvider];

export const ClusterCapacityProviderAssociationsCapacityProvider0 = {
    Fargate: "FARGATE",
    FargateSpot: "FARGATE_SPOT",
} as const;

/**
 * If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.
 */
export type ClusterCapacityProviderAssociationsCapacityProvider0 = (typeof ClusterCapacityProviderAssociationsCapacityProvider0)[keyof typeof ClusterCapacityProviderAssociationsCapacityProvider0];

export const ServiceAwsVpcConfigurationAssignPublicIp = {
    Disabled: "DISABLED",
    Enabled: "ENABLED",
} as const;

export type ServiceAwsVpcConfigurationAssignPublicIp = (typeof ServiceAwsVpcConfigurationAssignPublicIp)[keyof typeof ServiceAwsVpcConfigurationAssignPublicIp];

export const ServiceDeploymentControllerType = {
    CodeDeploy: "CODE_DEPLOY",
    Ecs: "ECS",
    External: "EXTERNAL",
} as const;

export type ServiceDeploymentControllerType = (typeof ServiceDeploymentControllerType)[keyof typeof ServiceDeploymentControllerType];

export const ServiceEbsTagSpecificationPropagateTags = {
    Service: "SERVICE",
    TaskDefinition: "TASK_DEFINITION",
} as const;

export type ServiceEbsTagSpecificationPropagateTags = (typeof ServiceEbsTagSpecificationPropagateTags)[keyof typeof ServiceEbsTagSpecificationPropagateTags];

export const ServiceLaunchType = {
    Ec2: "EC2",
    Fargate: "FARGATE",
    External: "EXTERNAL",
} as const;

export type ServiceLaunchType = (typeof ServiceLaunchType)[keyof typeof ServiceLaunchType];

export const ServicePlacementConstraintType = {
    DistinctInstance: "distinctInstance",
    MemberOf: "memberOf",
} as const;

export type ServicePlacementConstraintType = (typeof ServicePlacementConstraintType)[keyof typeof ServicePlacementConstraintType];

export const ServicePlacementStrategyType = {
    Binpack: "binpack",
    Random: "random",
    Spread: "spread",
} as const;

export type ServicePlacementStrategyType = (typeof ServicePlacementStrategyType)[keyof typeof ServicePlacementStrategyType];

export const ServicePropagateTags = {
    Service: "SERVICE",
    TaskDefinition: "TASK_DEFINITION",
} as const;

export type ServicePropagateTags = (typeof ServicePropagateTags)[keyof typeof ServicePropagateTags];

export const ServiceSchedulingStrategy = {
    Daemon: "DAEMON",
    Replica: "REPLICA",
} as const;

export type ServiceSchedulingStrategy = (typeof ServiceSchedulingStrategy)[keyof typeof ServiceSchedulingStrategy];

export const TaskDefinitionAuthorizationConfigIam = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

/**
 * Determines whether to use the Amazon ECS task role defined in a task definition when mounting the Amazon EFS file system. If it is turned on, transit encryption must be turned on in the ``EFSVolumeConfiguration``. If this parameter is omitted, the default value of ``DISABLED`` is used. For more information, see [Using Amazon EFS access points](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/efs-volumes.html#efs-volume-accesspoints) in the *Amazon Elastic Container Service Developer Guide*.
 */
export type TaskDefinitionAuthorizationConfigIam = (typeof TaskDefinitionAuthorizationConfigIam)[keyof typeof TaskDefinitionAuthorizationConfigIam];

export const TaskDefinitionEfsVolumeConfigurationTransitEncryption = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

/**
 * Determines whether to use encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server. Transit encryption must be turned on if Amazon EFS IAM authorization is used. If this parameter is omitted, the default value of ``DISABLED`` is used. For more information, see [Encrypting data in transit](https://docs.aws.amazon.com/efs/latest/ug/encryption-in-transit.html) in the *Amazon Elastic File System User Guide*.
 */
export type TaskDefinitionEfsVolumeConfigurationTransitEncryption = (typeof TaskDefinitionEfsVolumeConfigurationTransitEncryption)[keyof typeof TaskDefinitionEfsVolumeConfigurationTransitEncryption];

export const TaskDefinitionPortMappingAppProtocol = {
    Http: "http",
    Http2: "http2",
    Grpc: "grpc",
} as const;

/**
 * The application protocol that's used for the port mapping. This parameter only applies to Service Connect. We recommend that you set this parameter to be consistent with the protocol that your application uses. If you set this parameter, Amazon ECS adds protocol-specific connection handling to the Service Connect proxy. If you set this parameter, Amazon ECS adds protocol-specific telemetry in the Amazon ECS console and CloudWatch.
 *  If you don't set a value for this parameter, then TCP is used. However, Amazon ECS doesn't add protocol-specific telemetry for TCP.
 *   ``appProtocol`` is immutable in a Service Connect service. Updating this field requires a service deletion and redeployment.
 *  Tasks that run in a namespace can use short names to connect to services in the namespace. Tasks can connect to services across all of the clusters in the namespace. Tasks connect through a managed proxy container that collects logs and metrics for increased visibility. Only the tasks that Amazon ECS se
 */
export type TaskDefinitionPortMappingAppProtocol = (typeof TaskDefinitionPortMappingAppProtocol)[keyof typeof TaskDefinitionPortMappingAppProtocol];

export const TaskSetAwsVpcConfigurationAssignPublicIp = {
    Disabled: "DISABLED",
    Enabled: "ENABLED",
} as const;

/**
 * Whether the task's elastic network interface receives a public IP address. The default value is DISABLED.
 */
export type TaskSetAwsVpcConfigurationAssignPublicIp = (typeof TaskSetAwsVpcConfigurationAssignPublicIp)[keyof typeof TaskSetAwsVpcConfigurationAssignPublicIp];

export const TaskSetLaunchType = {
    Ec2: "EC2",
    Fargate: "FARGATE",
} as const;

/**
 * The launch type that new tasks in the task set will use. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html in the Amazon Elastic Container Service Developer Guide. 
 */
export type TaskSetLaunchType = (typeof TaskSetLaunchType)[keyof typeof TaskSetLaunchType];

export const TaskSetScaleUnit = {
    Percent: "PERCENT",
} as const;

/**
 * The unit of measure for the scale value.
 */
export type TaskSetScaleUnit = (typeof TaskSetScaleUnit)[keyof typeof TaskSetScaleUnit];
