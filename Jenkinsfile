pipeline {
  agent any
  stages {
    stage('first') {
      steps {
        echo 'first step'
        sleep 5
      }
    }
    stage('second') {
      steps {
        echo 'second'
      }
    }
  }
  environment {
    mykey = 'myvalue'
  }
}