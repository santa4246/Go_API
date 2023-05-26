# go_project
Golang API 서버 구축

## 개요

이 프로젝트는 Golang을 사용하여 간단한 API 서버를 구축하는 예시입니다. API 서버는 사용자 관리를 위한 기본적인 기능을 제공합니다.

## 기능

- 사용자 등록: 사용자의 이름, 이메일, 비밀번호를 입력받아 등록합니다.
- 사용자 인증: 입력된 이메일과 비밀번호를 확인하여 사용자를 인증합니다.
- 사용자 정보 조회: 등록된 사용자의 정보를 조회합니다.

## 설치 및 실행 방법

1. Golang 설치
2. 프로젝트 클론
```bash
git clone https://github.com/santa4246/go_project.git
```

3. 프로젝트 폴더로 이동 & go mod download
```bash
cd go_project
go mod download
```

4. 서버 실행
```bash
go run main.go
```