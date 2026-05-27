package worker

import (
	"ce191383/task_management/internal/entity"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func StartNotificationWorker(rdb *redis.Client) {
	ctx := context.Background()

	fmt.Println("Worker start...")

	for {
		result, err := rdb.BRPop(ctx, 0, "notification_queue").Result()
		if err != nil {
			fmt.Println("worker error:", err)
			continue
		}

		if len(result) < 2 {

			continue
		}

		var job entity.NotificationJob
		err = json.Unmarshal([]byte(result[1]), &job)
		if err != nil {
			fmt.Println("decode error:", err)
			continue
		}

		err = processJob(job)

		if err != nil {
			fmt.Println("job failded:", err)

			if job.Retry < 3 {
				job.Retry++
				fmt.Println("retry job:", job.Retry)

				data, _ := json.Marshal(job)
				rdb.LPush(ctx, "notification_queue", data)
			} else {
				fmt.Println("DEAD JOB:", job)
			}
		}
	}
}

func processJob(job entity.NotificationJob) error {
	if job.TaskID%2 == 0 {
		return errors.New("fake error processing job")
	}

	fmt.Println("Notification sent:")
	fmt.Println("taskID:", job.TaskID)
	fmt.Println("USerID:", job.UserID)
	fmt.Println("Message:", job.Message)

	return nil
}
