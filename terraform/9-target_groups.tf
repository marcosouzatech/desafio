module "alb" {
  source  = "terraform-aws-modules/alb/aws"
  version = "~> 8.0"

  name = "eks-alb"

  load_balancer_type = "application"

  vpc_id             = "vpc-085a293c258e6518f"
  subnets            = ["subnet-02bf9e37016702058", "subnet-0003deb1e437176c7"]
  security_groups    = ["sg-0a06eb376c61457ae"]
  internal           = true
  target_groups = [
    {
      name_prefix      = "tg-"
      backend_protocol = "HTTP"
      backend_port     = 80
      target_type      = "instance"
      targets = {
        my_target1 = {
          target_id = "i-0380b5ad2576291c1"
          port = 30001
        }
        my_target2 = {
          target_id = "i-0380b5ad2576291c1"
          port = 30002
        }
        my_target3 = {
          target_id = "i-0380b5ad2576291c1"
          port = 30003
        }
        my_target4 = {
          target_id = "i-0380b5ad2576291c1"
          port = 30004
        }
      }
    }
  ]

  http_tcp_listeners = [
    {
      port               = 80
      protocol           = "HTTP"
    }
  ]

  tags = {
    Environment = "Test"
  }
}