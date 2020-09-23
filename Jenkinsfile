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
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            }
        }

        stage('Test') {
                steps {
                    withEnv(["PATH+GO=${GOPATH}/bin"]){
                        sh 'pwd'
                        echo 'Running vetting here'
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
                echo 'Compiling and building'
                sh 'go build -o main main.go'
                sh './main'
            }
        }

    }
}