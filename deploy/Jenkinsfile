pipeline {
  agent any
  stages {
    stage('检出') {
      steps {
        script {
          if (env.GIT_TAG){
            DOCKER_IMAGE_VERSION = "${GIT_TAG}-${GIT_COMMIT_SHORT}-${CI_BUILD_NUMBER}"
          }
        }

        checkout([$class: 'GitSCM',
        branches: [[name: GIT_BUILD_REF]],
        userRemoteConfigs: [[
          url: GIT_REPO_URL,
          credentialsId: CREDENTIALS_ID
        ]]])
      }
    }

    stage('环境配置') {
      steps {
        sh 'rm -rf /root/programs/go'
        dir('/root/.cache/downloads') {
          sh 'wget -nc "https://golang.google.cn/dl/go1.20.3.linux-amd64.tar.gz" -O go-linux-amd64-1.20.3.tar.gz | true'
          sh 'tar -zxvf go-linux-amd64-1.20.3.tar.gz -C /root/programs'
        }

        sh 'go version'
      }
    }

    stage('单元测试') {
      post {
        always {
          junit '*.xml'
        }

      }
      steps {
        sh 'go get -u github.com/jstemmer/go-junit-report'
        sh 'go test -v ./... | go-junit-report > test.xml'
      }
    }

    stage('构建镜像并推送到制品库') {
      steps {
        sh "docker build -t ${CODING_DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION} ."
        useCustomStepPlugin(key: 'SYSTEM:artifact_docker_push', version: 'latest', params: [image:"${CODING_DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION}",repo:"${DOCKER_REPO_NAME}"])
      }
    }

    stage('人工确认') {
      steps {
        input(message: '是否发布到远端服务', submitter: 'BGTEoiwHWL')
      }
    }

    stage('部署到远端服务') {
      steps {
        script {
          def remoteConfig = [:]
          remoteConfig.name = "my-remote-server"
          remoteConfig.host = "${REMOTE_HOST}"
          remoteConfig.port = "${REMOTE_SSH_PORT}".toInteger()
          remoteConfig.user = "${REMOTE_USER_NAME}"
          remoteConfig.password = "${REMOTE_USER_PASSWD}"
          remoteConfig.allowAnyHosts = true

          withCredentials([
            usernamePassword(
              credentialsId: "${CODING_ARTIFACTS_CREDENTIALS_ID}",
              usernameVariable: 'CODING_DOCKER_REG_USERNAME',
              passwordVariable: 'CODING_DOCKER_REG_PASSWORD'
            )
          ]) {

            // 请确保远端环境中有 Docker 环境
            sshCommand(
              remote: remoteConfig,
              command: "docker login -u ${CODING_DOCKER_REG_USERNAME} -p ${CODING_DOCKER_REG_PASSWORD} ${CODING_DOCKER_REG_HOST}",
              sudo: true,
            )

            sshCommand(
              remote: remoteConfig,
              command: "docker rm -f ${DOCKER_IMAGE_NAME} | true",
              sudo: true,
            )

            // DOCKER_IMAGE_VERSION 中涉及到 GIT_LOCAL_BRANCH / GIT_TAG / GIT_COMMIT 的环境变量的使用
            // 需要在本地完成拼接后，再传入到远端服务器中使用
            DOCKER_IMAGE_URL = sh(
              script: "echo ${CODING_DOCKER_REG_HOST}/${CODING_DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_VERSION}",
              returnStdout: true
            )

            sshCommand(
              remote: remoteConfig,
              command: "docker run -d -p 8000:8000 -p 9000:9000 --name ${DOCKER_IMAGE_NAME} ${DOCKER_IMAGE_URL}",
              sudo: true,
            )

            echo "部署成功，请到 http://${REMOTE_HOST}:8000 预览效果"
          }
        }

      }
    }

  }
  environment {
    CODING_DOCKER_REG_HOST = "${CCI_CURRENT_TEAM}-docker.pkg.${CCI_CURRENT_DOMAIN}"
    CODING_DOCKER_IMAGE_NAME = "${PROJECT_NAME.toLowerCase()}/${DOCKER_REPO_NAME}/${DOCKER_IMAGE_NAME}"
  }
}