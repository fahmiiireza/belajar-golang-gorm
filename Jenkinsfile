pipeline {
    agent any
 
    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', changelog: false, poll: false, url: 'https://github.com/Man4ct/belajar-golang-gorm.git'            }
        }
        stage('Build') {
            steps {
                // Build Docker images for each microservice
                dir('book-service') {
                    sh 'docker buildx build --platform=linux/amd64 -t book-service .'
                }
                dir('user-service') {
                    sh 'docker buildx build --platform=linux/amd64 -t user-service .'
                }
            }
        }
        stage('Tag') {
            steps {
                // Tag Docker images with ECR repository URL
                dir('book-service') {
                sh 'docker tag book-service:latest public.ecr.aws/p2c0c2f5/book-service:latest'
                }
                dir('user-service') {
                sh 'docker tag user-service:latest public.ecr.aws/p2c0c2f5/user-service:latest'
                }
            }
        }
        stage('Reauthenticate') {
            steps {
                // Reauthenticate Docker to ECR
                sh 'aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/p2c0c2f5'
            }
        }
        stage('Push to ECR') {
            steps {
                // Push Docker images to ECR
                dir('book-service') {
                sh 'docker push public.ecr.aws/p2c0c2f5/book-service:latest'
                }
                dir('user-service') {
                sh 'aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/p2c0c2f5'
                sh 'docker push public.ecr.aws/p2c0c2f5/user-service:latest'
                }
            }
        }
        stage('Update Kubernetes Deployment') {
            steps {
                // Use kubectl to update Kubernetes deployment
                dir('k8s') {
                sh 'kubectl apply -f user-service.yaml --validate=false'
                sh 'kubectl apply -f user-db.yaml --validate=false'
                sh 'kubectl apply -f book-service.yaml --validate=false'
                sh 'kubectl apply -f book-db.yaml --validate=false'
                }
            }
        }
        stage('Restart Pods') {
            steps {
                // Restart pods to apply changes
                sh 'kubectl rollout restart deployment user-service'
                sh 'kubectl rollout restart deployment book-service'
            }
        }
    }
}
