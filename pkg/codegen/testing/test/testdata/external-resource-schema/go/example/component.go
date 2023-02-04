// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	storagev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/storage/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Component struct {
	pulumi.CustomResourceState

	Provider       kubernetes.ProviderOutput       `pulumi:"provider"`
	SecurityGroup  ec2.SecurityGroupOutput         `pulumi:"securityGroup"`
	StorageClasses storagev1.StorageClassMapOutput `pulumi:"storageClasses"`
}

// NewComponent registers a new resource with the given unique name, arguments, and options.
func NewComponent(ctx *pulumi.Context,
	name string, args *ComponentArgs, opts ...pulumi.ResourceOption) (*Component, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RequiredMetadata == nil {
		return nil, errors.New("invalid value for required argument 'RequiredMetadata'")
	}
	if args.RequiredMetadataArray == nil {
		return nil, errors.New("invalid value for required argument 'RequiredMetadataArray'")
	}
	if args.RequiredMetadataMap == nil {
		return nil, errors.New("invalid value for required argument 'RequiredMetadataMap'")
	}
	var resource Component
	err := ctx.RegisterResource("example::Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComponent gets an existing Component resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComponent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComponentState, opts ...pulumi.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.ReadResource("example::Component", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Component resources.
type componentState struct {
}

type ComponentState struct {
}

func (ComponentState) ElementType() reflect.Type {
	return reflect.TypeOf((*componentState)(nil)).Elem()
}

type componentArgs struct {
	Metadata              *metav1.ObjectMeta           `pulumi:"metadata"`
	MetadataArray         []metav1.ObjectMeta          `pulumi:"metadataArray"`
	MetadataMap           map[string]metav1.ObjectMeta `pulumi:"metadataMap"`
	RequiredMetadata      metav1.ObjectMeta            `pulumi:"requiredMetadata"`
	RequiredMetadataArray []metav1.ObjectMeta          `pulumi:"requiredMetadataArray"`
	RequiredMetadataMap   map[string]metav1.ObjectMeta `pulumi:"requiredMetadataMap"`
}

// The set of arguments for constructing a Component resource.
type ComponentArgs struct {
	Metadata              metav1.ObjectMetaPtrInput
	MetadataArray         metav1.ObjectMetaArrayInput
	MetadataMap           metav1.ObjectMetaMapInput
	RequiredMetadata      metav1.ObjectMetaInput
	RequiredMetadataArray metav1.ObjectMetaArrayInput
	RequiredMetadataMap   metav1.ObjectMetaMapInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type ComponentInput interface {
	pulumi.Input

	ToComponentOutput() ComponentOutput
	ToComponentOutputWithContext(ctx context.Context) ComponentOutput
}

func (*Component) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (i *Component) ToComponentOutput() ComponentOutput {
	return i.ToComponentOutputWithContext(context.Background())
}

func (i *Component) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentOutput)
}

// ComponentArrayInput is an input type that accepts ComponentArray and ComponentArrayOutput values.
// You can construct a concrete instance of `ComponentArrayInput` via:
//
//	ComponentArray{ ComponentArgs{...} }
type ComponentArrayInput interface {
	pulumi.Input

	ToComponentArrayOutput() ComponentArrayOutput
	ToComponentArrayOutputWithContext(context.Context) ComponentArrayOutput
}

type ComponentArray []ComponentInput

func (ComponentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Component)(nil)).Elem()
}

func (i ComponentArray) ToComponentArrayOutput() ComponentArrayOutput {
	return i.ToComponentArrayOutputWithContext(context.Background())
}

func (i ComponentArray) ToComponentArrayOutputWithContext(ctx context.Context) ComponentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentArrayOutput)
}

// ComponentMapInput is an input type that accepts ComponentMap and ComponentMapOutput values.
// You can construct a concrete instance of `ComponentMapInput` via:
//
//	ComponentMap{ "key": ComponentArgs{...} }
type ComponentMapInput interface {
	pulumi.Input

	ToComponentMapOutput() ComponentMapOutput
	ToComponentMapOutputWithContext(context.Context) ComponentMapOutput
}

type ComponentMap map[string]ComponentInput

func (ComponentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Component)(nil)).Elem()
}

func (i ComponentMap) ToComponentMapOutput() ComponentMapOutput {
	return i.ToComponentMapOutputWithContext(context.Background())
}

func (i ComponentMap) ToComponentMapOutputWithContext(ctx context.Context) ComponentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentMapOutput)
}

type ComponentOutput struct{ *pulumi.OutputState }

func (ComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (o ComponentOutput) ToComponentOutput() ComponentOutput {
	return o
}

func (o ComponentOutput) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return o
}

func (o ComponentOutput) Provider() kubernetes.ProviderOutput {
	return o.ApplyT(func(v *Component) kubernetes.ProviderOutput { return v.Provider }).(kubernetes.ProviderOutput)
}

func (o ComponentOutput) SecurityGroup() ec2.SecurityGroupOutput {
	return o.ApplyT(func(v *Component) ec2.SecurityGroupOutput { return v.SecurityGroup }).(ec2.SecurityGroupOutput)
}

func (o ComponentOutput) StorageClasses() storagev1.StorageClassMapOutput {
	return o.ApplyT(func(v *Component) storagev1.StorageClassMapOutput { return v.StorageClasses }).(storagev1.StorageClassMapOutput)
}

type ComponentArrayOutput struct{ *pulumi.OutputState }

func (ComponentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Component)(nil)).Elem()
}

func (o ComponentArrayOutput) ToComponentArrayOutput() ComponentArrayOutput {
	return o
}

func (o ComponentArrayOutput) ToComponentArrayOutputWithContext(ctx context.Context) ComponentArrayOutput {
	return o
}

func (o ComponentArrayOutput) Index(i pulumi.IntInput) ComponentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Component {
		return vs[0].([]*Component)[vs[1].(int)]
	}).(ComponentOutput)
}

type ComponentMapOutput struct{ *pulumi.OutputState }

func (ComponentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Component)(nil)).Elem()
}

func (o ComponentMapOutput) ToComponentMapOutput() ComponentMapOutput {
	return o
}

func (o ComponentMapOutput) ToComponentMapOutputWithContext(ctx context.Context) ComponentMapOutput {
	return o
}

func (o ComponentMapOutput) MapIndex(k pulumi.StringInput) ComponentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Component {
		return vs[0].(map[string]*Component)[vs[1].(string)]
	}).(ComponentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComponentInput)(nil)).Elem(), &Component{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComponentArrayInput)(nil)).Elem(), ComponentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComponentMapInput)(nil)).Elem(), ComponentMap{})
	pulumi.RegisterOutputType(ComponentOutput{})
	pulumi.RegisterOutputType(ComponentArrayOutput{})
	pulumi.RegisterOutputType(ComponentMapOutput{})
}
