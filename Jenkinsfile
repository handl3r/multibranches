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
                sh git branch
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            }
        }

        stage('Test') {
                steps {
                    withEnv(["PATH+GO=${GOPATH}/bin"]){
                        sh git branch
                        echo 'Running vetting'
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
                branch 'develop'
            }
            steps {
                sh git branch
                echo 'Compiling and building'
                sh 'go build'
                sh './testJ'
            }
        }

    }
}