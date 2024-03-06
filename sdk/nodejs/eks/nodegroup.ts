// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::EKS::Nodegroup
 */
export class Nodegroup extends pulumi.CustomResource {
    /**
     * Get an existing Nodegroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Nodegroup {
        return new Nodegroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:eks:Nodegroup';

    /**
     * Returns true if the given object is an instance of Nodegroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Nodegroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Nodegroup.__pulumiType;
    }

    /**
     * The AMI type for your node group.
     */
    public readonly amiType!: pulumi.Output<string | undefined>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * The capacity type of your managed node group.
     */
    public readonly capacityType!: pulumi.Output<string | undefined>;
    /**
     * Name of the cluster to create the node group in.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * The root device disk size (in GiB) for your node group instances.
     */
    public readonly diskSize!: pulumi.Output<number | undefined>;
    /**
     * Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
     */
    public readonly forceUpdateEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specify the instance types for a node group.
     */
    public readonly instanceTypes!: pulumi.Output<string[] | undefined>;
    /**
     * The Kubernetes labels to be applied to the nodes in the node group when they are created.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * An object representing a node group's launch template specification.
     */
    public readonly launchTemplate!: pulumi.Output<outputs.eks.NodegroupLaunchTemplateSpecification | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
     */
    public readonly nodeRole!: pulumi.Output<string>;
    /**
     * The unique name to give your node group.
     */
    public readonly nodegroupName!: pulumi.Output<string | undefined>;
    /**
     * The AMI version of the Amazon EKS-optimized AMI to use with your node group.
     */
    public readonly releaseVersion!: pulumi.Output<string | undefined>;
    /**
     * The remote access (SSH) configuration to use with your node group.
     */
    public readonly remoteAccess!: pulumi.Output<outputs.eks.NodegroupRemoteAccess | undefined>;
    /**
     * The scaling configuration details for the Auto Scaling group that is created for your node group.
     */
    public readonly scalingConfig!: pulumi.Output<outputs.eks.NodegroupScalingConfig | undefined>;
    /**
     * The subnets to use for the Auto Scaling group that is created for your node group.
     */
    public readonly subnets!: pulumi.Output<string[]>;
    /**
     * The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Kubernetes taints to be applied to the nodes in the node group when they are created.
     */
    public readonly taints!: pulumi.Output<outputs.eks.NodegroupTaint[] | undefined>;
    /**
     * The node group update configuration.
     */
    public readonly updateConfig!: pulumi.Output<outputs.eks.NodegroupUpdateConfig | undefined>;
    /**
     * The Kubernetes version to use for your managed nodes.
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a Nodegroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodegroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            if ((!args || args.nodeRole === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeRole'");
            }
            if ((!args || args.subnets === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnets'");
            }
            resourceInputs["amiType"] = args ? args.amiType : undefined;
            resourceInputs["capacityType"] = args ? args.capacityType : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["diskSize"] = args ? args.diskSize : undefined;
            resourceInputs["forceUpdateEnabled"] = args ? args.forceUpdateEnabled : undefined;
            resourceInputs["instanceTypes"] = args ? args.instanceTypes : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["launchTemplate"] = args ? args.launchTemplate : undefined;
            resourceInputs["nodeRole"] = args ? args.nodeRole : undefined;
            resourceInputs["nodegroupName"] = args ? args.nodegroupName : undefined;
            resourceInputs["releaseVersion"] = args ? args.releaseVersion : undefined;
            resourceInputs["remoteAccess"] = args ? args.remoteAccess : undefined;
            resourceInputs["scalingConfig"] = args ? args.scalingConfig : undefined;
            resourceInputs["subnets"] = args ? args.subnets : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["taints"] = args ? args.taints : undefined;
            resourceInputs["updateConfig"] = args ? args.updateConfig : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["amiType"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["capacityType"] = undefined /*out*/;
            resourceInputs["clusterName"] = undefined /*out*/;
            resourceInputs["diskSize"] = undefined /*out*/;
            resourceInputs["forceUpdateEnabled"] = undefined /*out*/;
            resourceInputs["instanceTypes"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["launchTemplate"] = undefined /*out*/;
            resourceInputs["nodeRole"] = undefined /*out*/;
            resourceInputs["nodegroupName"] = undefined /*out*/;
            resourceInputs["releaseVersion"] = undefined /*out*/;
            resourceInputs["remoteAccess"] = undefined /*out*/;
            resourceInputs["scalingConfig"] = undefined /*out*/;
            resourceInputs["subnets"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["taints"] = undefined /*out*/;
            resourceInputs["updateConfig"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["amiType", "capacityType", "clusterName", "diskSize", "instanceTypes[*]", "nodeRole", "nodegroupName", "remoteAccess", "subnets[*]"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Nodegroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Nodegroup resource.
 */
export interface NodegroupArgs {
    /**
     * The AMI type for your node group.
     */
    amiType?: pulumi.Input<string>;
    /**
     * The capacity type of your managed node group.
     */
    capacityType?: pulumi.Input<string>;
    /**
     * Name of the cluster to create the node group in.
     */
    clusterName: pulumi.Input<string>;
    /**
     * The root device disk size (in GiB) for your node group instances.
     */
    diskSize?: pulumi.Input<number>;
    /**
     * Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
     */
    forceUpdateEnabled?: pulumi.Input<boolean>;
    /**
     * Specify the instance types for a node group.
     */
    instanceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Kubernetes labels to be applied to the nodes in the node group when they are created.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * An object representing a node group's launch template specification.
     */
    launchTemplate?: pulumi.Input<inputs.eks.NodegroupLaunchTemplateSpecificationArgs>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
     */
    nodeRole: pulumi.Input<string>;
    /**
     * The unique name to give your node group.
     */
    nodegroupName?: pulumi.Input<string>;
    /**
     * The AMI version of the Amazon EKS-optimized AMI to use with your node group.
     */
    releaseVersion?: pulumi.Input<string>;
    /**
     * The remote access (SSH) configuration to use with your node group.
     */
    remoteAccess?: pulumi.Input<inputs.eks.NodegroupRemoteAccessArgs>;
    /**
     * The scaling configuration details for the Auto Scaling group that is created for your node group.
     */
    scalingConfig?: pulumi.Input<inputs.eks.NodegroupScalingConfigArgs>;
    /**
     * The subnets to use for the Auto Scaling group that is created for your node group.
     */
    subnets: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Kubernetes taints to be applied to the nodes in the node group when they are created.
     */
    taints?: pulumi.Input<pulumi.Input<inputs.eks.NodegroupTaintArgs>[]>;
    /**
     * The node group update configuration.
     */
    updateConfig?: pulumi.Input<inputs.eks.NodegroupUpdateConfigArgs>;
    /**
     * The Kubernetes version to use for your managed nodes.
     */
    version?: pulumi.Input<string>;
}
