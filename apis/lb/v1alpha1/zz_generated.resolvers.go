/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1"
	v1alpha11 "github.com/crossplane-contrib/provider-tf-aws/apis/s3/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this LBListener.
func (mg *LBListener) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DefaultAction); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.DefaultAction[i3].Forward); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.DefaultAction[i3].Forward[i4].TargetGroup); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultAction[i3].Forward[i4].TargetGroup[i5].Arn),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.DefaultAction[i3].Forward[i4].TargetGroup[i5].ArnRef,
					Selector:     mg.Spec.ForProvider.DefaultAction[i3].Forward[i4].TargetGroup[i5].ArnSelector,
					To: reference.To{
						List:    &LBTargetGroupList{},
						Managed: &LBTargetGroup{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.DefaultAction[i3].Forward[i4].TargetGroup[i5].Arn")
				}
				mg.Spec.ForProvider.DefaultAction[i3].Forward[i4].TargetGroup[i5].Arn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.DefaultAction[i3].Forward[i4].TargetGroup[i5].ArnRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.DefaultAction); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArnRef,
			Selector:     mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArnSelector,
			To: reference.To{
				List:    &LBTargetGroupList{},
				Managed: &LBTargetGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArn")
		}
		mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArnRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancerArn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LoadBalancerArnRef,
		Selector:     mg.Spec.ForProvider.LoadBalancerArnSelector,
		To: reference.To{
			List:    &LoadBalancerList{},
			Managed: &LoadBalancer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancerArn")
	}
	mg.Spec.ForProvider.LoadBalancerArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LBTargetGroup.
func (mg *LBTargetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VpcID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VpcIDRef,
		Selector:     mg.Spec.ForProvider.VpcIDSelector,
		To: reference.To{
			List:    &v1alpha1.VPCList{},
			Managed: &v1alpha1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VpcID")
	}
	mg.Spec.ForProvider.VpcID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VpcIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LBTargetGroupAttachment.
func (mg *LBTargetGroupAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TargetGroupArn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TargetGroupArnRef,
		Selector:     mg.Spec.ForProvider.TargetGroupArnSelector,
		To: reference.To{
			List:    &LBTargetGroupList{},
			Managed: &LBTargetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetGroupArn")
	}
	mg.Spec.ForProvider.TargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetGroupArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LoadBalancer.
func (mg *LoadBalancer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.AccessLogs); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccessLogs[i3].Bucket),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.AccessLogs[i3].BucketRef,
			Selector:     mg.Spec.ForProvider.AccessLogs[i3].BucketSelector,
			To: reference.To{
				List:    &v1alpha11.BucketList{},
				Managed: &v1alpha11.Bucket{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.AccessLogs[i3].Bucket")
		}
		mg.Spec.ForProvider.AccessLogs[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.AccessLogs[i3].BucketRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroups),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupsSelector,
		To: reference.To{
			List:    &v1alpha1.SecurityGroupList{},
			Managed: &v1alpha1.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroups")
	}
	mg.Spec.ForProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupsRefs = mrsp.ResolvedReferences

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SubnetMapping); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetMapping[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.SubnetMapping[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.SubnetMapping[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1alpha1.SubnetList{},
				Managed: &v1alpha1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SubnetMapping[i3].SubnetID")
		}
		mg.Spec.ForProvider.SubnetMapping[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SubnetMapping[i3].SubnetIDRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Subnets),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetsRefs,
		Selector:      mg.Spec.ForProvider.SubnetsSelector,
		To: reference.To{
			List:    &v1alpha1.SubnetList{},
			Managed: &v1alpha1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Subnets")
	}
	mg.Spec.ForProvider.Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetsRefs = mrsp.ResolvedReferences

	return nil
}
