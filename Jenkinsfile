pipeline {
  agent {
    kubernetes {
      label 'mypod'
      yamlFile 'KubernetesPod.yaml'
    }
  }
  options {
    buildDiscarder logRotator(numToKeepStr: '5')
    disableConcurrentBuilds()
  }
  environment {
    domain = "acme.com"
  }
  stages {
    stage('build') {
      steps {
        container('maven') {
          script { currentBuild.displayName = "${env.BUILD_NUMBER}"}
          sh 'echo ${env.BUILD_NUMBER}'
          sh 'mvn -version'
          slackSend channel: '#aws', color: 'good', message: 'Slack Message', teamDomain: 'carlymiguel', token: 'SBsVEshhLeHqrQTeuTVgeQtl'
        }
        container('golang') {
          //checkout scm
          sh 'go version'
          k8sBuildGolang("hello.go")
        }
      }
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
    failure {
      echo 'Booo!'
    }
  }
}
