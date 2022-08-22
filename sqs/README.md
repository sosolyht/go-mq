## SQS 함수 종류

```Text
CreateQueue : 큐 생성

ListQueues : 큐 목록 출력

DeleteQueue : 큐 삭제

SendMessage : 큐에 메세지 추가

SendMessageBatch : 큐에 여러 메세지 추가

ReceiveMessage : 큐에서 메세지를 꺼내서 보기

ChangeMessageVisibility : 메세지의 보기 제한 시간 변경

ChangeMessageVisibilityBatch : 여러 메세지의 보기 제한 시간 변경

DeleteMessage : 큐에서 메세지 삭제

DeleteMesasgeBatch : 큐에서 여러 메세지 삭제

SetQueueAttributes : 큐 설정 변경 (지연 전송 시간, 최대 메세지 크기, 메세지 보관 기간, 접근 정책, 짧은 폴링/긴 폴링 시간, 처리 실패 큐)

GetQueueAttributes : 큐 설정 확인 (메세지 개수, 큐 생성 시간, 최종 큐 변경 시간, 큐 ARN, 지연 전송 시간, 최대 메세지 크기, 메세지 보관 기간, 접근 정책, 짧은 폴링/긴 폴링 시간, 처리 실패 큐)

GetQueueUrl : 큐 엔드포인트 URL 확인

AddPermission : 다른 AWS 계정에 대한 접근 권한 설정

RemovePermission : 다른 AWS 계정에 대한 접근 권한 삭제

```