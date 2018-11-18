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
    stage('static-analysis-checkstyle-linting-code-coverage') {
      failFast true
      parallel {
        stage('lint') {
          environment {
            foo = "bar"
          }
          steps {
            echo "static-analysis-checkstyle-linting-code-coverage Sonarqube?..."
            try {
              echo "Linting ember..."
              npm run clean
              npm run init
              npm run lint:jenkins
            }
            finally {
              checkstyle pattern: 'target/test/checkstyle/eslint-*.xml'
            }
          }
        }
        stage('dependencies-security-checks') {
          environment {
            foo = "bar"
          }
          steps {
            echo "running dependency checks with npm audit, security checks and getting missing ones..."
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
            echo "Building Cockpit..."
            cd embercli/cockpitapp
            npm run build:production
            cd ../..

            echo "Building Explore..."
            cd embercli/explore
            npm run build:production
            cd ../..

            echo "Building Planner..."
            cd embercli/planner
            npm run build:production
            cd ../..

            echo "Build admin and platform..."
            cd etc/release/jsOptimization
            echo "#############################################################"
            ./build.sh -n --all
            echo "############################################################# "

            echo "Cleanup..."
            npm run clean

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
            try {
              echo "Compile scala code"
              sbt ${SBT_OPTS} test:compile

              echo "Package application. For what?"
              sbt ${SBT_OPTS} stage
            }
            finally {
              echo "Cleanup generated artifacts (sbt target folder, node_modules) so they don't occupy space. Needed?"
              sh "sbt ${SBT_OPTS} clean"
            }
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
    stage('end-to-end-and-performance-tests') {
      steps {
        echo 'running end to end and performance tests...'
        container('worker') {
          sh 'node --version'
          sh 'npm --version'
          sh 'sbt ${SBT_OPTS} sbtVersion'
        }
      }
    }
    stage('deploy') {
      steps {
        echo 'running deployment or creating release on github...'
      }
      when { branch 'master' }
    }
  }
  post {
    success {
      slackSend color: "good", message: "Build Completed <@miguel> - ${env.BUILD_TAG} ${env.GIT_COMMIT} ${env.GIT_URL} ${env.GIT_BRANCH} traditional (<${env.BUILD_URL}|Open>) blue ocean (<${env.RUN_DISPLAY_URL}|Open>) ", channel: "#aws", teamDomain: "carlymiguel", tokenCredentialId: "slack-token"
    }

    failure {
      slackSend color: "danger", message: "Build Failed - ${env.JOB_NAME} ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)", channel: "#aws", teamDomain: "carlymiguel", tokenCredentialId: "slack-token"
      echo "creating JIRA issue for this..."
    }
  }
}
