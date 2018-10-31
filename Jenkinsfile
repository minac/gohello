podTemplate(
    label: 'kubernetes',
    containers: [
        containerTemplate(name: 'maven', image: 'maven:alpine', ttyEnabled: true, command: 'cat'),
        containerTemplate(name: 'golang', image: 'golang:alpine', ttyEnabled: true, command: 'cat')
    ]
) {
    node('kubernetes') {
        container('maven') {
            stage('build') {
                sh 'mvn --version'
                slackSend channel: '#aws', color: 'good', message: 'Slack Message', teamDomain: 'carlymiguel', token: 'SBsVEshhLeHqrQTeuTVgeQtl'
            }
            stage('unit-test') {
                sh 'java -version'
            }
        }
        container('golang') {
            stage('deploy') {
                sh 'go version'
                sh 'go build hello.go'
            }
        }
    }
}