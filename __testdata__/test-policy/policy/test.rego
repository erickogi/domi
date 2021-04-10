package terraform.aws.s3

has_field(obj, field) {
    obj[field]
}

warn[msg] {
    policyID := "DOMI-AWS-001"
    resource := input.resource.aws_s3_bucket[bucket]
    not has_field(resource, "tags")
    msg = sprintf("%s: No tags found: `%v`", [ policyID, bucket ])
}