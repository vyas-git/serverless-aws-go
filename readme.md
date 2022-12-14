# Go Serverless api aws lambda

**Serverless Golang REST API with AWS Lambda. Deploying a Golang API to AWS Lambda and handling requests through AWS API Gateway**

## AWS Setup & Configure 

* install aws cli https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html



* create Iam USER and attach AdministratorAccess policy https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html


* aws configuration run below command \
` aws configure `


	AWS Access Key ID (None): aws_access_key_id \
	AWS Secret Access Key (None): aws_secret_access_key \
	Default region name (None): us-east-2 \
	Default output format (None): json


* Create File: ./tmp/trust-policy.json and copy paste the below json which will  allow lambda services to assume the lambda-function-executor role \n

`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Service": "lambda.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
        }
    ]
}`


* Create Role with trust policy 

`aws iam create-role --role-name lambda-function-executor \
--assume-role-policy-document file://./tmp/trust-policy.json`


* Attach the policy AWSLambdaBasicExecutionRole using the aws iam attach-role-policy command as below

`
aws iam attach-role-policy --role-name lambda-function-executor \
--policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

`

* Now, let us deploy the actual Lambda function to AWS using the command aws lambda create-function

`functionName=gettingstarted #you can give your preferred name
aws lambda create-function --function-name $functionName --runtime go1.x \
--role $rolearn \
--handler main --zip-file fileb://./output/function.zip
`


* check list of functions 
` aws lambda list-functions `
