// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lightsail::Database
type Database struct {
	pulumi.CustomResourceState

	// The Availability Zone in which to create your new database. Use the us-east-2a case-sensitive format.
	AvailabilityZone pulumi.StringPtrOutput `pulumi:"availabilityZone"`
	// When true, enables automated backup retention for your database. Updates are applied during the next maintenance window because this can result in an outage.
	BackupRetention pulumi.BoolPtrOutput `pulumi:"backupRetention"`
	// Indicates the certificate that needs to be associated with the database.
	CaCertificateIdentifier pulumi.StringPtrOutput `pulumi:"caCertificateIdentifier"`
	DatabaseArn             pulumi.StringOutput    `pulumi:"databaseArn"`
	// The name of the database to create when the Lightsail database resource is created. For MySQL, if this parameter isn't specified, no database is created in the database resource. For PostgreSQL, if this parameter isn't specified, a database named postgres is created in the database resource.
	MasterDatabaseName pulumi.StringOutput `pulumi:"masterDatabaseName"`
	// The password for the master user. The password can include any printable ASCII character except "/", """, or "@". It cannot contain spaces.
	MasterUserPassword pulumi.StringPtrOutput `pulumi:"masterUserPassword"`
	// The name for the master user.
	MasterUsername pulumi.StringOutput `pulumi:"masterUsername"`
	// The daily time range during which automated backups are created for your new database if automated backups are enabled.
	PreferredBackupWindow pulumi.StringPtrOutput `pulumi:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur on your new database.
	PreferredMaintenanceWindow pulumi.StringPtrOutput `pulumi:"preferredMaintenanceWindow"`
	// Specifies the accessibility options for your new database. A value of true specifies a database that is available to resources outside of your Lightsail account. A value of false specifies a database that is available only to your Lightsail resources in the same region as your database.
	PubliclyAccessible pulumi.BoolPtrOutput `pulumi:"publiclyAccessible"`
	// The blueprint ID for your new database. A blueprint describes the major engine version of a database.
	RelationalDatabaseBlueprintId pulumi.StringOutput `pulumi:"relationalDatabaseBlueprintId"`
	// The bundle ID for your new database. A bundle describes the performance specifications for your database.
	RelationalDatabaseBundleId pulumi.StringOutput `pulumi:"relationalDatabaseBundleId"`
	// The name to use for your new Lightsail database resource.
	RelationalDatabaseName pulumi.StringOutput `pulumi:"relationalDatabaseName"`
	// Update one or more parameters of the relational database.
	RelationalDatabaseParameters DatabaseRelationalDatabaseParameterArrayOutput `pulumi:"relationalDatabaseParameters"`
	// When true, the master user password is changed to a new strong password generated by Lightsail. Use the get relational database master user password operation to get the new password.
	RotateMasterUserPassword pulumi.BoolPtrOutput `pulumi:"rotateMasterUserPassword"`
	// An array of key-value pairs to apply to this resource.
	Tags DatabaseTagArrayOutput `pulumi:"tags"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MasterDatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'MasterDatabaseName'")
	}
	if args.MasterUsername == nil {
		return nil, errors.New("invalid value for required argument 'MasterUsername'")
	}
	if args.RelationalDatabaseBlueprintId == nil {
		return nil, errors.New("invalid value for required argument 'RelationalDatabaseBlueprintId'")
	}
	if args.RelationalDatabaseBundleId == nil {
		return nil, errors.New("invalid value for required argument 'RelationalDatabaseBundleId'")
	}
	if args.RelationalDatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'RelationalDatabaseName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Database
	err := ctx.RegisterResource("aws-native:lightsail:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("aws-native:lightsail:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
}

type DatabaseState struct {
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// The Availability Zone in which to create your new database. Use the us-east-2a case-sensitive format.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// When true, enables automated backup retention for your database. Updates are applied during the next maintenance window because this can result in an outage.
	BackupRetention *bool `pulumi:"backupRetention"`
	// Indicates the certificate that needs to be associated with the database.
	CaCertificateIdentifier *string `pulumi:"caCertificateIdentifier"`
	// The name of the database to create when the Lightsail database resource is created. For MySQL, if this parameter isn't specified, no database is created in the database resource. For PostgreSQL, if this parameter isn't specified, a database named postgres is created in the database resource.
	MasterDatabaseName string `pulumi:"masterDatabaseName"`
	// The password for the master user. The password can include any printable ASCII character except "/", """, or "@". It cannot contain spaces.
	MasterUserPassword *string `pulumi:"masterUserPassword"`
	// The name for the master user.
	MasterUsername string `pulumi:"masterUsername"`
	// The daily time range during which automated backups are created for your new database if automated backups are enabled.
	PreferredBackupWindow *string `pulumi:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur on your new database.
	PreferredMaintenanceWindow *string `pulumi:"preferredMaintenanceWindow"`
	// Specifies the accessibility options for your new database. A value of true specifies a database that is available to resources outside of your Lightsail account. A value of false specifies a database that is available only to your Lightsail resources in the same region as your database.
	PubliclyAccessible *bool `pulumi:"publiclyAccessible"`
	// The blueprint ID for your new database. A blueprint describes the major engine version of a database.
	RelationalDatabaseBlueprintId string `pulumi:"relationalDatabaseBlueprintId"`
	// The bundle ID for your new database. A bundle describes the performance specifications for your database.
	RelationalDatabaseBundleId string `pulumi:"relationalDatabaseBundleId"`
	// The name to use for your new Lightsail database resource.
	RelationalDatabaseName string `pulumi:"relationalDatabaseName"`
	// Update one or more parameters of the relational database.
	RelationalDatabaseParameters []DatabaseRelationalDatabaseParameter `pulumi:"relationalDatabaseParameters"`
	// When true, the master user password is changed to a new strong password generated by Lightsail. Use the get relational database master user password operation to get the new password.
	RotateMasterUserPassword *bool `pulumi:"rotateMasterUserPassword"`
	// An array of key-value pairs to apply to this resource.
	Tags []DatabaseTag `pulumi:"tags"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// The Availability Zone in which to create your new database. Use the us-east-2a case-sensitive format.
	AvailabilityZone pulumi.StringPtrInput
	// When true, enables automated backup retention for your database. Updates are applied during the next maintenance window because this can result in an outage.
	BackupRetention pulumi.BoolPtrInput
	// Indicates the certificate that needs to be associated with the database.
	CaCertificateIdentifier pulumi.StringPtrInput
	// The name of the database to create when the Lightsail database resource is created. For MySQL, if this parameter isn't specified, no database is created in the database resource. For PostgreSQL, if this parameter isn't specified, a database named postgres is created in the database resource.
	MasterDatabaseName pulumi.StringInput
	// The password for the master user. The password can include any printable ASCII character except "/", """, or "@". It cannot contain spaces.
	MasterUserPassword pulumi.StringPtrInput
	// The name for the master user.
	MasterUsername pulumi.StringInput
	// The daily time range during which automated backups are created for your new database if automated backups are enabled.
	PreferredBackupWindow pulumi.StringPtrInput
	// The weekly time range during which system maintenance can occur on your new database.
	PreferredMaintenanceWindow pulumi.StringPtrInput
	// Specifies the accessibility options for your new database. A value of true specifies a database that is available to resources outside of your Lightsail account. A value of false specifies a database that is available only to your Lightsail resources in the same region as your database.
	PubliclyAccessible pulumi.BoolPtrInput
	// The blueprint ID for your new database. A blueprint describes the major engine version of a database.
	RelationalDatabaseBlueprintId pulumi.StringInput
	// The bundle ID for your new database. A bundle describes the performance specifications for your database.
	RelationalDatabaseBundleId pulumi.StringInput
	// The name to use for your new Lightsail database resource.
	RelationalDatabaseName pulumi.StringInput
	// Update one or more parameters of the relational database.
	RelationalDatabaseParameters DatabaseRelationalDatabaseParameterArrayInput
	// When true, the master user password is changed to a new strong password generated by Lightsail. Use the get relational database master user password operation to get the new password.
	RotateMasterUserPassword pulumi.BoolPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags DatabaseTagArrayInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

// The Availability Zone in which to create your new database. Use the us-east-2a case-sensitive format.
func (o DatabaseOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// When true, enables automated backup retention for your database. Updates are applied during the next maintenance window because this can result in an outage.
func (o DatabaseOutput) BackupRetention() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.BoolPtrOutput { return v.BackupRetention }).(pulumi.BoolPtrOutput)
}

// Indicates the certificate that needs to be associated with the database.
func (o DatabaseOutput) CaCertificateIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.CaCertificateIdentifier }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) DatabaseArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.DatabaseArn }).(pulumi.StringOutput)
}

// The name of the database to create when the Lightsail database resource is created. For MySQL, if this parameter isn't specified, no database is created in the database resource. For PostgreSQL, if this parameter isn't specified, a database named postgres is created in the database resource.
func (o DatabaseOutput) MasterDatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.MasterDatabaseName }).(pulumi.StringOutput)
}

// The password for the master user. The password can include any printable ASCII character except "/", """, or "@". It cannot contain spaces.
func (o DatabaseOutput) MasterUserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.MasterUserPassword }).(pulumi.StringPtrOutput)
}

// The name for the master user.
func (o DatabaseOutput) MasterUsername() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.MasterUsername }).(pulumi.StringOutput)
}

// The daily time range during which automated backups are created for your new database if automated backups are enabled.
func (o DatabaseOutput) PreferredBackupWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.PreferredBackupWindow }).(pulumi.StringPtrOutput)
}

// The weekly time range during which system maintenance can occur on your new database.
func (o DatabaseOutput) PreferredMaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.PreferredMaintenanceWindow }).(pulumi.StringPtrOutput)
}

// Specifies the accessibility options for your new database. A value of true specifies a database that is available to resources outside of your Lightsail account. A value of false specifies a database that is available only to your Lightsail resources in the same region as your database.
func (o DatabaseOutput) PubliclyAccessible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.BoolPtrOutput { return v.PubliclyAccessible }).(pulumi.BoolPtrOutput)
}

// The blueprint ID for your new database. A blueprint describes the major engine version of a database.
func (o DatabaseOutput) RelationalDatabaseBlueprintId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.RelationalDatabaseBlueprintId }).(pulumi.StringOutput)
}

// The bundle ID for your new database. A bundle describes the performance specifications for your database.
func (o DatabaseOutput) RelationalDatabaseBundleId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.RelationalDatabaseBundleId }).(pulumi.StringOutput)
}

// The name to use for your new Lightsail database resource.
func (o DatabaseOutput) RelationalDatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.RelationalDatabaseName }).(pulumi.StringOutput)
}

// Update one or more parameters of the relational database.
func (o DatabaseOutput) RelationalDatabaseParameters() DatabaseRelationalDatabaseParameterArrayOutput {
	return o.ApplyT(func(v *Database) DatabaseRelationalDatabaseParameterArrayOutput {
		return v.RelationalDatabaseParameters
	}).(DatabaseRelationalDatabaseParameterArrayOutput)
}

// When true, the master user password is changed to a new strong password generated by Lightsail. Use the get relational database master user password operation to get the new password.
func (o DatabaseOutput) RotateMasterUserPassword() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.BoolPtrOutput { return v.RotateMasterUserPassword }).(pulumi.BoolPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o DatabaseOutput) Tags() DatabaseTagArrayOutput {
	return o.ApplyT(func(v *Database) DatabaseTagArrayOutput { return v.Tags }).(DatabaseTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInput)(nil)).Elem(), &Database{})
	pulumi.RegisterOutputType(DatabaseOutput{})
}
