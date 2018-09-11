#!/usr/bin/env groovy
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
        echo 'second step'
      }
    }
  }
  environment {
    mykey = 'myvalue'
  }
}
