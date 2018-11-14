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
    stage('lint-dependency-security') {
      failFast true
      parallel {
        stage('lint') {
          environment {
            foo = "bar"
          }
          steps {
            echo "linting..."
          }
        }
        stage('dependencies') {
          environment {
            foo = "bar"
          }
          steps {
            echo "running dependency checks, security checks and getting missing ones..."
          }
        }
      }
    }
    stage('build') {
      environment {
        repo = "https://github.com/minac/gohello.git"
      }
     failFast true
      parallel {
        stage('build-frontend') {
          environment {
            foo = "bar"
          }
          steps {
            echo "running build frontend..."
            container('maven') {
              sh 'mvn -version'
            }
            container('golang') {
              //checkout scm
              sh 'go version'
              k8sBuildGolang("hello")
            }
          }
          // when { changeset "**/*.js" }
          // when { changeRequest target: 'master' }
          // when { anyOf { branch 'master' ; branch 'staging'; environment name: 'DEPLOY_TO', value: 'production' } }
        }
        stage('build-backend') {
          environment {
            foo = "bar"
          }
          steps {
            echo "running build backend..."
          }
        }
      }
    }
    stage('unit-tests') {
      failFast true
      parallel {
        stage('unit-tests-frontend') {
          environment {
            foo = "bar"
          }
          steps {
            echo "running unit tests frontend..."
          }
        }
        stage('unit-tests-backend') {
          environment {
            foo = "bar"
          }
          steps {
            echo "running unit tests backend..."
          }
        }
      }
    }
    stage('end-to-end-tests') {
      steps {
        echo 'running end to end tests...'
        container('worker') {
          sh 'node --version'
          sh 'npm --version'
          sh 'sbt sbtVersion'
        }
      }
    }
    stage('deploy') {
      steps {
        echo 'running deployment/release...'
      }
      when { branch 'master' }
    }
  }
  post {
    success {
      slackSend color: "good", message: "Build Completed <@miguel> - ${env.BUILD_TAG} ${env.GIT_COMMIT} ${env.GIT_URL} ${env.GIT_BRANCH} (<${env.BUILD_URL}|Open>) blue ocean (<${env.RUN_DISPLAY_URL}|Open>) ", channel: "#aws", teamDomain: "carlymiguel", tokenCredentialId: "slack-token"
    }

    failure {
      slackSend color: "danger", message: "Build Failed - ${env.JOB_NAME} ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)", channel: "#aws", teamDomain: "carlymiguel", tokenCredentialId: "slack-token"
    }
  }
}
