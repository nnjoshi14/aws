# ECS commands guide #

## Create and push image to ECR ##

### ECR Repository ###
* Create repository # aws ecr create-repository --repository-name hello-repository --region ap-south-1 

### Docker Image ###
* Create Docker Image # vim Dockerfile
* Build docker Imange # docker build -t hello-world .
* Test image locally  # docker run -t -i -p 80:80 hello-world
* Tag repo to image   # docker tag hello-world aws_account_id.dkr.ecr.ap-south-1.amazonaws.com/hello-repository
* Login to the repo   # aws ecr get-login-password | docker login --username AWS --password-stdin aws_account_id.dkr.ecr.ap-south-1.amazonaws.com
* Push the image      # docker push aws_account_id.dkr.ecr.ap-south-1.amazonaws.com


## Cleanup ##
### ECR repostiroty ###
* Force delete repo   # aws ecr delete-repository --repository-name hello-repository --region hello-repository --force

