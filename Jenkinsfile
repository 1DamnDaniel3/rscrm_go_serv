pipeline {
    agent any
    environment {
        PROJECT_DIR = '/srv/developers/src/rscrm'
    }
    stages {
        stage('Deploy') {
            steps {
                dir("${env.WORKSPACE}") {
                    sh """
                    cd ${PROJECT_DIR}/rscrm_go_serv
                    git reset --hard
                    git clean -fd
                    git pull
                    cd ${PROJECT_DIR}
                    docker compose up -d --build backend
                    """
                }
            }
        }
    }
}
