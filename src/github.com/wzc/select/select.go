package selectrace

import (
	"fmt"
	"net/http"
	"time"
)

//Racer got who is faster url
//select关键字可以实现进程同步
func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

//SynchroniseRacer ping url Synchronously
func SynchroniseRacer(a, b string, timeout time.Duration) (winner string, e error) {
	//select允许在多通道上等待
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	//time.After()可以防止一个永远阻塞的代码
	case <-time.After(timeout):
		return "", fmt.Errorf("time out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	//内存的角度，chan struct{}是可用的最小内存
	//用make创建chan，因为使用var创建，默认值是nil，如果像它发送值它将永远阻塞，因为不能发送nil
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		//一旦完成get，则关闭channel
		close(ch)
	}()
	return ch
}
