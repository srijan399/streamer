pipeline {
    agent any
    environment {
        GITHUB_CONTEXT = 'jenkins-ci'
    }
    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/srijan399/streamer.git'
            }
        }
        stage('Build') {
            steps {
                sh 'go build ./...'
            }
        }
        stage('Test') {
            steps {
                sh 'go test -v ./...'
            }
        }
    }
    post {
        success {
            publishChecks name: env.GITHUB_CONTEXT, status: 'COMPLETED', conclusion: 'SUCCESS', summary: 'Build and tests passed'
        }
        failure {
            publishChecks name: env.GITHUB_CONTEXT, status: 'COMPLETED', conclusion: 'FAILURE', summary: 'Build or tests failed'
        }
    }
}