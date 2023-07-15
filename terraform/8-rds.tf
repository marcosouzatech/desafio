resource "aws_db_instance" "desafio-db" {
  
  allocated_storage       = 10
  identifier              = "desafio"
  engine                  = "mysql"
  engine_version          = "5.7"
  instance_class          = "db.t3.micro"
  db_name                 = "desafio"
  username                = "desafio"
  password                = "M$la5WNtA5qT"
  parameter_group_name    = "default.mysql5.7"
  publicly_accessible     = false
  skip_final_snapshot     = true
}
 
output "Endpoint_string" {
  value = aws_db_instance.desafio-db.endpoint
 
}