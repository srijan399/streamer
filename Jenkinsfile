pipeline {
    agent any
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
            githubNotify context: 'jenkins-ci', status: 'SUCCESS', credentialsId: 'github-token'
        }
        failure {
            githubNotify context: 'jenkins-ci', status: 'FAILURE', credentialsId: 'github-token'
        }
    }
}