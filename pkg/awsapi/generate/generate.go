package generate

//go:generate ../../../build/scripts/generate-aws-interfaces.sh sts STS
//go:generate ../../../build/scripts/generate-aws-interfaces.sh cloudformation CloudFormation
//go:generate ../../../build/scripts/generate-aws-interfaces.sh elasticloadbalancing ELB
//go:generate ../../../build/scripts/generate-aws-interfaces.sh elasticloadbalancingv2 ELBV2
//go:generate ../../../build/scripts/generate-aws-interfaces.sh ssm SSM
