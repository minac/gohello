#!groovy
import java.text.SimpleDateFormat

pipeline {
  agent {
    kubernetes {
      label 'mypod'
      yamlFile 'KubernetesPod.yaml'
      idleMinutes 10
    }
  }
  options {
    buildDiscarder logRotator(numToKeepStr: '5')
    disableConcurrentBuilds()
    retry(2)
    timeout(time: 2, unit: 'HOURS')
    timestamps()
  }
  environment {
    domain = "curiousellie.com"
    repo = "https://github.com/minac/gohello.git"
  }
  stages {
    stage('build') {
      environment {
        SLACK_TOKEN = credentials('slack-token')
      }
      steps {
        container('maven') {
          sh 'mvn -version'
          slackSend channel: '#aws', color: 'good', message: "STARTED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})", teamDomain: 'carlymiguel', token: "${env.SLACK_TOKEN}"
          // 'SBsVEshhLeHqrQTeuTVgeQtl'
        }
        container('golang') {
          //checkout scm
          sh 'go version'
          k8sBuildGolang("hello.go")
        }
      }
      when { branch 'master' }
      // when { changeset "**/*.js" }
      // when { changeRequest target: 'master' }
      // when { anyOf { branch 'master'; branch 'staging'; environment name: 'DEPLOY_TO', value: 'production' } }
    }
    stage('test') {
      steps {
        container('golang') {
          checkout scm
        }
        container('worker') {
          sh 'node --version'
          sh 'npm --version'
          sh 'sbt sbtVersion'
        }
      }
    }
    stage('release') {
      steps {
        echo 'This would release it.'
      }
    }
  }
  post {
    success {
      slackSend channel: '#aws', color: 'good', message: "SUCCESSFUL: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})", teamDomain: 'carlymiguel', token: "${env.SLACK_TOKEN}"
    }

    failure {
      slackSend channel: '#aws', color: 'good', message: "FAILED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})", teamDomain: 'carlymiguel', token: "${env.SLACK_TOKEN}"
    }
  }
}
