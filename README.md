<div align="center">

</br>
<!-- logo -->
<img src="https://private-user-images.githubusercontent.com/83758624/377772070-d937bb49-c8c2-4e93-baa6-4e4b6b413d2e.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MjkyMzc3ODMsIm5iZiI6MTcyOTIzNzQ4MywicGF0aCI6Ii84Mzc1ODYyNC8zNzc3NzIwNzAtZDkzN2JiNDktYzhjMi00ZTkzLWJhYTYtNGU0YjZiNDEzZDJlLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDEwMTglMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQxMDE4VDA3NDQ0M1omWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPWY4OGVlZjdlZTllYTdiYWIyNDU2OTQ4YmMwZmU5NDVkZjIyNTcwZGE4MTZmMWUwMjQ4MTQ1ODNkZGQ1YTc2MGUmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.JPkCUf9kIlYfJt5vb-M1ESlS6sUPETl6Fbt4oIT4yIc" width="400"/>

### ✅ Gateway module ✅
</div> 

## 📝 소개
Nginx, ApiSix와 같은 게이트웨이 역할을 할 수 있는 web-server입니다.

간단하게 정해진 yaml형태를 통해서 라우팅 처리를 하는 서버를 자동으로 배포 및 Request를 처리하게 됩니다.

## 🗂️ yaml 규격

다음과 같은 `yaml` 형태를 지원합니다.
- 기본적인 `yaml`의 경로는 `./deploy.yaml` 파일 경로를 따르게 됩니다.
- `path`에 대한 `flag`를 자유롭게 지원합니다.

</br>

<div align="center">
1️⃣
</div>

```
app:
  name: "<Gatew Moudle Name>"
  version: v1
  port: "<Gatway Port>"
```

> 기본적인 module의 정보 입니다.
- name : 원하는 module의 이름을 지정합니다.
- version : version 관리를 합니다.
- port : web-server의 port를 설정합니다.

</br>

<div align="center">
2️⃣
</div>

```
http:
  base_url: "<Internal Service Base URL>"
  router:
    - method: "<Method>"
      get_type: "<Type>"
      path: "<Optional path>"
      header:
        Content-Type: "application/json"
        Accept: "application/json"
    - method: "<Method>"
      get_type: "<Type>"
      path: "<Optional path>"
      variable:
        - "address"
        - "tid"
```

> 라우팅하고자 하는 외부 서버의 정보를 입력하게 됩니다.
- base_url : 외부 서버의 기본 url를 적어줍니다.
- method : 외부 서버의 path가 지원하는 method 타입을 지정합니다.
  - `GET`, `POST`, `PUT`, `DELETE` 를 지원합니다.
- get_type : `GET`요청에 대해서 `url`, `query` 두가지를 지원합니다.
- path : 추가적인 path를 입력해 줍니다.
- header : 요청에 대해서 담고자 하는 정보를 추가합니다.
- variable : `GET` 요청의 `query` 타입에 대해서 맵핑고자 하는 키 값을 입력합니다.

</br>

<div align="center">
3️⃣
</div>

```
kafka:
  url: "<Kafka Producer Url>"
  client_id: "<Producer client id>"
  acks: "<producer acks>"
  topic: "<kafka topic>"
  batch_time: "<Request Produce Batch Time>"
```

>producing하고자 하는 kafka의 정보를 입력해 줍니다.
- url : producer의 접속 정보를 입력해줍니다.
- client_id : producer의 client_id를 입력합니다.
- acks : 원하는 acks 단계를 설정합니다. 빈 값인경우 `all`로 설정됩니다.
- topic : producing하고자 하는 topic을 넣어줍니다.
- batch_time : 주기적으로 produce하는 주기를 설정합니다.

## ⚙ 기술 스택
> module에서 사용하는 디펜던씨를 정리합니다.

`fx` : dependency injection를 활용하기 위해서 사용하였습니다.

`zap` : logging 기능을 위해서 사용하였습니다.

`resty` : builder 패턴을 활용하며 api를 전송하며, api에 대한 결과를 tracing하기 위해 사용 되었습니다.

`confluent-kafka-go` : api에 대한 전송량을 message 형태로 전송하기 위해서 사용되었습니다.
- 해당 모듈에서는 producing만 진행을 합니다. consumer는 구성이 되어 있지 않습니다.

`sonic` : web-server 특성상 성능적인 이점을 최대한 챙겨가기 위해서 직렬화 및 역직렬화에 적용을 하였습니다.

`fiber` : web-server의 api를 구성하며 빠른 성능을 보장하기 위해서 사용 되었습니다.

`gopkg.in/yaml.v2` : yaml 형태를 환경변수로써 관리하기 위해서 사용되었습니다.

<br />

## 🛠️ 프로젝트 아키텍쳐
<div align="center">

<img src="https://private-user-images.githubusercontent.com/83758624/377785076-e9bb0861-6b91-4d79-8384-b8791f1a69fc.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MjkyNDAwMDAsIm5iZiI6MTcyOTIzOTcwMCwicGF0aCI6Ii84Mzc1ODYyNC8zNzc3ODUwNzYtZTliYjA4NjEtNmI5MS00ZDc5LTgzODQtYjg3OTFmMWE2OWZjLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDEwMTglMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQxMDE4VDA4MjE0MFomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTg1ZWMzNTA2NWJkMmJmMWQ4ZmY1NTg4OTAyNDVhZTdlMWQ0YjYwZmU4OWMxN2I3MjEzZTMwOGM0MGFlZGJmZGYmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.7difsd042yyeQXay4-7onG56yoi286QBdtvHYui8W4w"/>

</div>

> 해당 아키텍처에서 consumer를 추가하여, 비동기로 들어오는 메시지를 Lock 제어를 통해 DB에 저장을 하는 아키텍처로써 돌아 가게 됩니다.
> >Consumer는 해당 module에 작성이 되어 있지 않으니, 따로 구성이 되어야 합니다.
