pipeline {
    agent any
    tools {
        go 'Golang1.15'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }

    stages {
        stage('setup workspace') {
            steps {
//                 sh 'export GOPATH=/var/jenkins_home/workspace/'
//                 sh 'export PATH=$GOPATH/bin:$PATH'
            }
        }

        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go mod tidy'
                sh 'ls'
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
                sh 'go get'
            }
        }

        stage('Test') {
                steps {
                    withEnv(["PATH+GO=${GOPATH}/bin"]){
                        sh 'go vet .'
                        echo 'Running linting'
                        sh 'golint .'
                        echo 'Running test'
                        sh 'go test -v'
                    }
                }
        }

        stage('Build And Deploy') {
            when {
                branch 'master'
            }
            steps {
                sh 'git branch'
                echo 'Compiling and building'
                sh 'go build -o main main.go'
                sshagent(['my-ssh-key']) {
                    sh 'ssh thai@192.168.1.10 pkill main'
                    sh 'scp ./main thai@192.168.1.10:~/app/'
//                     sh 'ssh thai@192.168.1.10 './main &'  ' netstat -ltnp | grep -w ":8000" ''
//                     sh(script: " ssh thai@192.168.1.10 'echo $1 | ~/app/main ",
//                      returnStdout: true )
                     sh '''
                        ssh thai@192.168.1.10 '~/app/main > ~/app/start.log &'
                     '''
                }
            }
        }

    }
}