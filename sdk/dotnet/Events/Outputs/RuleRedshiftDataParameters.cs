// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Outputs
{

    [OutputType]
    public sealed class RuleRedshiftDataParameters
    {
        public readonly string Database;
        public readonly string? DbUser;
        public readonly string? SecretManagerArn;
        public readonly string Sql;
        public readonly ImmutableArray<string> Sqls;
        public readonly string? StatementName;
        public readonly bool? WithEvent;

        [OutputConstructor]
        private RuleRedshiftDataParameters(
            string database,

            string? dbUser,

            string? secretManagerArn,

            string sql,

            ImmutableArray<string> sqls,

            string? statementName,

            bool? withEvent)
        {
            Database = database;
            DbUser = dbUser;
            SecretManagerArn = secretManagerArn;
            Sql = sql;
            Sqls = sqls;
            StatementName = statementName;
            WithEvent = withEvent;
        }
    }
}
