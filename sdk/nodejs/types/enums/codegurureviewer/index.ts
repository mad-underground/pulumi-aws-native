// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const RepositoryAssociationType = {
    CodeCommit: "CodeCommit",
    Bitbucket: "Bitbucket",
    GitHubEnterpriseServer: "GitHubEnterpriseServer",
    S3Bucket: "S3Bucket",
} as const;

/**
 * The type of repository to be associated.
 */
export type RepositoryAssociationType = (typeof RepositoryAssociationType)[keyof typeof RepositoryAssociationType];
