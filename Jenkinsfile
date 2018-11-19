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
    stage('setup') {
      failFast true
      parallel {
        stage('lint') {
          environment {
            foo = "bar"
          }
          steps {
            container('worker') {
              echo "static-analysis-checkstyle-linting-code-coverage Sonarqube?..."
              echo "Linting ember..."
              sh """
                STARTTIME=$(date -u +%s)

                echo "this next line will probably not be required because the containers are always fresh"
                npm run clean
                #is this necessary after clean? npm run init
                mkdir -p target/test/checkstyle/
                npm run lint

                ENDTIME=$((date -u +%s))
                echo "Task finished in $(($ENDTIME - $STARTTIME)) seconds."
              """
            }
          }
          post {
            always {
              container('worker') {
                sh """
                  STARTTIME=$(date -u +%s)

                  echo "Creating reports from linting..."
                  recordIssues enabledForFailure: true, tools: [[pattern: 'target/test/checkstyle/eslint-*.xml', tool: [$class: 'CheckStyle']]]
                  # sh "checkstyle pattern: 'target/test/checkstyle/eslint-*.xml'"

                  ENDTIME=$((date -u +%s))
                  echo "Task finished in $(($ENDTIME - $STARTTIME)) seconds."
                """
              }
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
            container('worker') {
              echo "running build frontend..."
              echo "Building Cockpit..."
              sh """
                STARTTIME=$(date -u +%s)

                cd embercli/cockpitapp
                npm run build:production
                cd ../..

                ENDTIME=$((date -u +%s))
                echo "Task finished in $(($ENDTIME - $STARTTIME)) seconds."
              """

              echo "Building Explore..."
              sh """
                STARTTIME=$(date -u +%s)

                cd embercli/explore
                npm run build:production
                cd ../..

                ENDTIME=$((date -u +%s))
                echo "Task finished in $(($ENDTIME - $STARTTIME)) seconds."
              """

              echo "Building Planner..."
              sh """
                STARTTIME=$(date -u +%s)

                cd embercli/planner
                npm run build:production
                cd ../..

                ENDTIME=$((date -u +%s))
                echo "Task finished in $(($ENDTIME - $STARTTIME)) seconds."
              """

              echo "Build admin and platform..."
              sh """
                STARTTIME=$(date -u +%s)

                cd etc/release/jsOptimization

                echo "Minimizing platform JS scripts ..."
                start_time_platform=`date +%s`
                r.js -o optimizationSettings-platform.js
                cp -v build/platform/main.js ../../../public/javascripts/platform/nezasa-platform.min.js
                echo component run time: $(expr `date +%s` - $start_time_platform) s

                echo "Minimizing admin JS scripts ..."
                start_time_admin=`date +%s`
                r.js -o optimizationSettings-admin.js
                cp -v build/admin/main.js ../../../public/javascripts/admin/nezasa-admin.min.js
                echo component run time: $(expr `date +%s` - $start_time_admin) s

                cd -

                ENDTIME=$((date -u +%s))
                echo "Task finished in $(($ENDTIME - $STARTTIME)) seconds."
              """
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
            container('worker') {
              sh """
                STARTTIME=$(date -u +%s)
                echo "running build backend..."

                echo "Compile scala code"
                // not needed because test:compile does both
                //sbt ${SBT_OPTS} compile"
                sbt ${SBT_OPTS} test:compile

                echo "Package application. For what?"
                sbt ${SBT_OPTS} stage

                ENDTIME=$((date -u +%s))
                echo "Task finished in $(($ENDTIME - $STARTTIME)) seconds."
              """
            }
          }
          post {
            always {
              container('worker') {
                sh """
                  STARTTIME=$(date -u +%s)

                  echo "Cleanup NPM..."
                  npm run clean

                  echo "Cleanup generated artifacts (sbt target folder, node_modules) so they don't occupy space."
                  sbt ${SBT_OPTS} clean

                  ENDTIME=$((date -u +%s))
                  echo "Task finished in $(($ENDTIME - $STARTTIME)) seconds."
                """
              }
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
          sh """
            STARTTIME=$(date -u +%s)

            node --version
            npm --version
            sbt ${SBT_OPTS} sbtVersion

            ENDTIME=$((date -u +%s))
            echo "Task finished in $(($ENDTIME - $STARTTIME)) seconds."
          """
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
