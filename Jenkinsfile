pipeline {
    agent any
    environment {
        GITHUB_CONTEXT = 'jenkins-ci'
        CREDENTIAL = 'e5d3a2a4-4b1c-49b3-963b-db73d93c21a9'
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
            githubNotify context: env.GITHUB_CONTEXT, status: 'SUCCESS', credentialsId: env.CREDENTIAL
        }
        failure {
            githubNotify context: env.GITHUB_CONTEXT, status: 'FAILURE', credentialsId: env.CREDENTIAL
        }
    }
}
