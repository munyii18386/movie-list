export GOOS="linux"
# build GO executable
GOOS=linux go build

docker build -t lmburu/database .

docker push lmburu/database

# delete GO executable
go clean

sh database-deploy.sh