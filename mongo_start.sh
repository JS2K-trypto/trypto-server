 docker run --name mongodb -v /home/opc/data:/data/db -d -p 27017:27017 mongo

# 재부팅시 도커 자동재시작
# docker update --restart=always mongodb