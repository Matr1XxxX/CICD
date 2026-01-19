pipeline {
    agent any

    stages {

        stage('Checkout Code') {
            steps {
                echo 'Checking out source code'
                checkout scm
            }
        }

        stage('Build') {
            steps {
                echo 'Build stage'
                sh './hello.sh'
            }
        }

        stage('Unit Tests') {
            steps {
                echo 'Running unit tests (Shift Left)'
                sh './test.sh'
            }
        }

        stage('Security Scan') {
            steps {
                echo 'Running security checks (DevSecOps)'
                sh './security_scan.sh'
            }
        }
    }

    post {
        success {
            echo '✅ Pipeline SUCCESS - Code is secure and healthy'
        }
        failure {
            echo '❌ Pipeline FAILED - Stop the Line'
        }
    }
}

