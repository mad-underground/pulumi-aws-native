// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::ECR::PublicRepository resource specifies an Amazon Elastic Container Public Registry (Amazon Public ECR) repository, where users can push and pull Docker images. For more information, see https://docs.aws.amazon.com/AmazonECR
 */
export function getPublicRepository(args: GetPublicRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetPublicRepositoryResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ecr:getPublicRepository", {
        "repositoryName": args.repositoryName,
    }, opts);
}

export interface GetPublicRepositoryArgs {
    /**
     * The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.
     */
    repositoryName: string;
}

export interface GetPublicRepositoryResult {
    readonly arn?: string;
    /**
     * The CatalogData property type specifies Catalog data for ECR Public Repository. For information about Catalog Data, see <link>
     */
    readonly repositoryCatalogData?: outputs.ecr.RepositoryCatalogDataProperties;
    /**
     * The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. 
     */
    readonly repositoryPolicyText?: any;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.ecr.PublicRepositoryTag[];
}

export function getPublicRepositoryOutput(args: GetPublicRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPublicRepositoryResult> {
    return pulumi.output(args).apply(a => getPublicRepository(a, opts))
}

export interface GetPublicRepositoryOutputArgs {
    /**
     * The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.
     */
    repositoryName: pulumi.Input<string>;
}
