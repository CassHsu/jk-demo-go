node {
    stage('Prepare') {
        sh "printenv"
        echo "1.Prepare Stage"
        checkout scm
        script {
            build_tag = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
        }
    }
    stage('Test') {
      echo "2.Test Stage"
    }
    stage('Build') {
        echo "3.Build Docker Image Stage"
        sh "docker build -t 172.17.8.111:8595/sv-jk-demo-go:${build_tag} ."
    }
    stage('Push') {
        echo "4.Push Docker Image Stage"
        withCredentials([usernamePassword(credentialsId: 'sv-rigistery', passwordVariable: 'password', usernameVariable: 'user')]) {
            sh "docker login -u ${user} -p ${password}"
            sh "docker push 172.17.8.111:8595/sv-jk-demo-go:${build_tag}"
        }
    }
    stage('Deploy') {
        echo "5. Deploy Stage"
        sh "sed -i 's/<BUILD_TAG>/${build_tag}/' k8s.yaml"
        sh "sed -i 's/<BRANCH_NAME>/${env.BRANCH_NAME}/' k8s.yaml"
        sh "kubectl apply -f k8s.yaml --record"
    }
}