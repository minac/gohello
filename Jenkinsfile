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
          sh 'mvn -version'
          slackSend channel: '#aws', color: 'good', message: 'Slack Message', teamDomain: 'carlymiguel', token: 'SBsVEshhLeHqrQTeuTVgeQtl'
        }
        container('busybox') {
          sh '/bin/busybox echo "being busy"'
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
        container('busybox') {
          sh 'echo "test"'
        }
      }
    }
    stage('release') { 
      steps {
        container('busybox') {
          sh 'echo "release"'
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
