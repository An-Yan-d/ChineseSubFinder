package backend

import (
	"github.com/An-Yan-d/ChineseSubFinder/pkg/types/task_queue"
)

type ReplyAllJobs struct {
	AllJobs []task_queue.OneJob `json:"all_jobs"`
}
