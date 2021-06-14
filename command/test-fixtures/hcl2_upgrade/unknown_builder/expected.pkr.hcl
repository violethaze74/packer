packer {
  required_version = ">= 1.6.0"
}

variable "aws_access_key" {
  type    = string
  default = ""
}

variable "aws_region" {
  type = string
}

variable "aws_secret_key" {
  type    = string
  default = ""
}

source "potatoes" "autogenerated_1" {
  access_key      = "${var.aws_access_key}"
  ami_description = "Ubuntu 16.04 LTS - expand root partition"
  ami_name        = "ubuntu-16-04-test"
  encrypt_boot    = true
  launch_block_device_mappings {
    delete_on_termination = true
    device_name           = "/dev/sda1"
    volume_size           = 48
    volume_type           = "gp2"
  }
  region              = "${var.aws_region}"
  secret_key          = "${var.aws_secret_key}"
  source_ami          = "ami1234567"
  spot_instance_types = ["t2.small", "t2.medium", "t2.large"]
  spot_price          = "0.0075"
  ssh_interface       = "session_manager"
  ssh_username        = "ubuntu"
  temporary_iam_instance_profile_policy_document {
    Statement {
      Action   = ["*"]
      Effect   = "Allow"
      Resource = ["*"]
    }
    Version = "2012-10-17"
  }
}

build {
  sources = ["source.potatoes.autogenerated_1"]

}