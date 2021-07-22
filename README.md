# number-baseball
숫자 야구 게임을 TDD + Golang 으로 구현
최범균님의 [고랭(golang) 기초 연습 - 야구게임 01
](https://www.youtube.com/watch?v=Co2yAUJlm0c) 영상을 보고 작성했습니다.

## 규칙
- 문제와 답 모두 숫자 세개 입력
- 숫자는 같을 수 없음
- 같은 수 같은 자리 -> 스트라이크
- 같은 수 다른 자리 -> 볼
- 다른 수 -> nothing

### 요구사항 세분화
**Game**
* [X] 각 숫자는 1~9 이다. 
* [X] 세 자리의 숫자로 만든다.
* [X] 숫자는 중복되지 않는다.

**Numbers**
* [X] 다른 수 -> nothing
* [X] 같은 수 다른 자리 -> ball
* [X] 같은 수 같은 자리 -> strike
