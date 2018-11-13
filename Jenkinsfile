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
  }
  stages {
    stage('build') {
      environment {
        repo = "https://github.com/minac/gohello.git"
      }
      steps {
        container('maven') {
          sh 'mvn -version'
        }
        container('golang') {
          //checkout scm
          sh 'go version'
          k8sBuildGolang("hello.go")
        }
      }
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
          sh '"$(which node)" --version'
          sh '"$(which npm)" --version'
          sh 'sbt sbtVersion'
        }
      }
    }
    stage('release') {
      steps {
        echo 'This would release it.'
      }
      when { branch 'staging' }
    }
  }
  post {
    success {
      slackSend color: "good", message: "Build Completed - ${env.JOB_NAME} ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)", channel: "#aws", teamDomain: "carlymiguel", tokenCredentialId: slack-token
      // 'SBsVEshhLeHqrQTeuTVgeQtl'
    }

    failure {
      slackSend color: "danger", message: "Build Failed - ${env.JOB_NAME} ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)", channel: "#aws", teamDomain: "carlymiguel", tokenCredentialId: slack-token
    }
  }
}
