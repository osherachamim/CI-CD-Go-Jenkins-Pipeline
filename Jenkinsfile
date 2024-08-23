pipeline {
  agent any
  stages {
    stage('Dev') {
      steps {
        sh '''GO111MODULE=off go test ./...
'''
      }
    }

  }
}