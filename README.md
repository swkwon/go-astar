# A* 알고리즘
A* 길찾기 알고리즘을 Go로 구현하였습니다.
# 패키지 설치
```
$ go get github.com/swkwon/go-astar@latest
```
# 시작하기
```
	if m1, err := astar.MakeDefaultMap("map-1", 10, 10, astar.Coordinate{X: 3, Y: 2}, astar.Coordinate{X: 6, Y: 9},
		[]astar.Coordinate{
			{X: 2, Y: 4},
			{X: 3, Y: 4},
			{X: 4, Y: 4},
			{X: 5, Y: 4},
			{X: 5, Y: 8},
			{X: 6, Y: 8},
			{X: 7, Y: 8},
			{X: 8, Y: 8},
		}); err == nil {
		m1.Find()
		m1.Print()
	} else {
		panic(err)
	}
```
* `MakeDefaultMap` 메서드를 이용하여 2차원 맵 정보를 생성합니다. 
* `Find` 메서드를 이용하여 길을 찾습니다.
* `Print` 메서드를 이용하여 결과를 출력합니다.
# 예제 실행방법
```
$ go install github.com/swkwon/go-astar/cmd/example@latest

$ example.exe
```
예제는 `GOBIN` 폴더에 설치됩니다. GOBIN 환경변수 설정이 안되어 있을 경우 설치되지 않습니다. 예제를 실행하면 두가지 예제에 대해 출력됩니다.
```
start a* algorithm

map-1
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 X 0 0 0 0 0
0 0 S 0 X 0 0 0 0 0
0 0 0 . X 0 0 0 . 0
0 0 0 . X 0 0 . X .
0 0 0 0 . . . 0 X E
0 0 0 0 0 0 0 0 X 0
0 0 0 0 0 0 0 0 X 0
0 0 0 0 0 0 0 0 0 0

map-2
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 X 0 0 0 0 0
0 0 S 0 X 0 0 0 0 0
0 0 0 . X 0 0 0 0 0
0 0 0 . X 0 0 0 X X
0 0 0 0 . . . 0 X E
0 0 0 0 0 0 0 . X .
0 0 0 0 0 0 0 . X .
0 0 0 0 0 0 0 0 . 0
end a* algorithm
```
## Key
map의 key 의미는 아래와 같습니다.
* `0` - 갈 수 있는 길
* `X` - 막힌길
* `S` - 시작점
* `E` - 종착점
* `.` - 이동한 길
## 비용 계산
상하좌우 이동할 때는 `10`, 대각선 이동할 경우 `14`의 이동 비용이 발생합니다.