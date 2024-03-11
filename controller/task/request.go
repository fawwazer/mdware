package task

type TaskRequest struct {
	Kegiatan string `json:"taskname" form:"taskname" validate:"required"`
}
