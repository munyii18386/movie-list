
export GOOS="linux"
# build GO executable
GOOS=linux go build

docker build -t lmburu/database .

docker push lmburu/database

# delete GO executable
go clean

ssh -T ec2-user@ec2-13-59-0-236.us-east-2.compute.amazonaws.com  < database-deploy.sh