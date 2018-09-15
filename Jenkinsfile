pipeline {
  agent any

  stages {
    stage('Build test container') {
      steps {
        sh '''
          docker build \
            --tag nx584-test:${BUILD_NUMBER} \
            --target test \
            .
        '''
      }
    }
    stage('Run tests') {
      steps {
        parallel(
          lint: {
            sh '''
              docker run nx584-test:${BUILD_NUMBER} \
                go tool vet -all -v .
            '''
          },
          test: {
            sh '''
              docker run nx584-test:${BUILD_NUMBER} \
                go test -v
            '''
          }
        )
      }
    }
    stage('Build production container') {
      steps {
        sh '''
          docker build \
            --tag nx584:${BUILD_NUMBER} \
            --target production \
            --cache-from=nx584-test:${BUILD_NUMBER} \
            .
        '''
      }
    }
  }
}
