package todo

import "time"

type Task struct {
	Title       string     `json:"title"`
	Text        string     `json:"text"`
	IsDone      bool       `json:"is_done"`
	CreatedAt   time.Time  `json:"created_at"`
	ComplitedAt *time.Time `json:"complited_at,omitempty"`
}

func NewTask(title string, text string) *Task {
	return &Task{
		Title:     title,
		Text:      text,
		IsDone:    false,
		CreatedAt: time.Now(),
	}
}

func (t *Task) Complite() {
	t.IsDone = true
	now := time.Now()
	t.ComplitedAt = &now

}
func (t *Task) Uncomplite() {
	t.IsDone = false
	t.ComplitedAt = nil
}
