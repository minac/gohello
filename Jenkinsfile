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
          myGlobalFunction("myinput")
          sh 'echo MAVEN_CONTAINER_ENV_VAR = ${CONTAINER_ENV_VAR}'
          sh 'mvn -version'
          slackSend channel: '#aws', color: 'good', message: 'Slack Message', teamDomain: 'carlymiguel', token: 'SBsVEshhLeHqrQTeuTVgeQtl'
        }
        container('busybox') {
          sh 'echo BUSYBOX_CONTAINER_ENV_VAR = ${CONTAINER_ENV_VAR}'
          sh '/bin/busybox'
        }
        container('golang') {
          checkout scm
          sh 'go version'
          sh 'go build hello.go'
        }
      }
    }
    stage('test') { 
      steps {
        container('busybox') {
          sh 'echo 'test'
        }
      }
    }
    stage('release') { 
      steps {
        container('busybox') {
          sh 'echo 'release'
        }
      }
    }
  }
  post {
    failure {
      echo 'Booo!'
    }
  }
}
