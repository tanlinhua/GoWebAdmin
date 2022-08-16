package snow

import (
	"fmt"
	"testing"
)

func TestSnowflake(t *testing.T) {
	worker, err := NewWorker(0)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 5; i++ {
		id := worker.NextId() // 雪花算法生成唯一ID
		fmt.Println(id)
	}
}

/*
=== RUN   TestSnowflake
15931420048359424
15931420048359425
15931420048359426
15931420048359427
15931420048359428
--- PASS: TestSnowflake (0.00s)
*/
