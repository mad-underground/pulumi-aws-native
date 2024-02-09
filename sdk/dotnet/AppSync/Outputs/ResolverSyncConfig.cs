// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Outputs
{

    /// <summary>
    /// Describes a Sync configuration for a resolver.
    ///  Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is invoked.
    /// </summary>
    [OutputType]
    public sealed class ResolverSyncConfig
    {
        /// <summary>
        /// The Conflict Detection strategy to use.
        ///   +   *VERSION*: Detect conflicts based on object versions for this resolver.
        ///   +   *NONE*: Do not detect conflicts when invoking this resolver.
        /// </summary>
        public readonly string ConflictDetection;
        /// <summary>
        /// The Conflict Resolution strategy to perform in the event of a conflict.
        ///   +   *OPTIMISTIC_CONCURRENCY*: Resolve conflicts by rejecting mutations when versions don't match the latest version at the server.
        ///   +   *AUTOMERGE*: Resolve conflicts with the Automerge conflict resolution strategy.
        ///   +   *LAMBDA*: Resolve conflicts with an LAMlong function supplied in the ``LambdaConflictHandlerConfig``.
        /// </summary>
        public readonly string? ConflictHandler;
        /// <summary>
        /// The ``LambdaConflictHandlerConfig`` when configuring ``LAMBDA`` as the Conflict Handler.
        /// </summary>
        public readonly Outputs.ResolverLambdaConflictHandlerConfig? LambdaConflictHandlerConfig;

        [OutputConstructor]
        private ResolverSyncConfig(
            string conflictDetection,

            string? conflictHandler,

            Outputs.ResolverLambdaConflictHandlerConfig? lambdaConflictHandlerConfig)
        {
            ConflictDetection = conflictDetection;
            ConflictHandler = conflictHandler;
            LambdaConflictHandlerConfig = lambdaConflictHandlerConfig;
        }
    }
}
