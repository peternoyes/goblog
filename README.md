# goblog
goblog is a very simple blogging platform written in Go and designed to run in the AWS Elastic Beanstalk and store the blog data S3. THe engine uses the blackfriday markdown library, gorilla mux, and the built in Go templating functionality. Development of goblog has been primarily an exercise to learn and experiment with Go and AWS to build my personal blog. [Here](http://peternoyes.elasticbeanstalk.com) is the blog in action.

## Getting Started
The configuration is set via environment variables. The blog data can either be stored in a local folder or in Amazon S3.

| Key                           | Value                         |
| ----------------------------- | ----------------------------- |
| GOBLOG_DATA                   | Local Path / S3 Bucket Name   |
| GOBLOG_DRIVER                 | file / aws                    |
| GOBLOG_REGION                 | AWS default region            |
| AWS_ACCESS_KEY_ID             | AWS access key                |
| AWS_SECRET_ACCESS_KEY         | AWS secret key                |

config.json needs to be located in the data folder. A sample configuration file is included in the posts/ folder in the repository.

Markdown files need to contain a set of key/value pairs in the top part of the file. See one of the sample markdown files for the syntax.

Finally, markdown files need to follow a precise naming convention:

YYYY-MM-DD-HH-mm-ss-Title.md