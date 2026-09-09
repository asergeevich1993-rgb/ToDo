package todo

import (
	"errors"
	"fmt"
	"sync"
)

type List struct {
	tasks map[string]*Task
	mtx   sync.RWMutex
}

func NewList() *List {
	return &List{
		tasks: make(map[string]*Task),
	}
}

func (l *List) Add(task *Task) error {
	defer l.mtx.Unlock()
	l.mtx.Lock()
	_, ok := l.tasks[task.Title]
	if ok {
		return errors.New("такая задача уже есть")
	}
	l.tasks[task.Title] = task

	return nil
}

func (l *List) Done(title string) error {
	defer l.mtx.Unlock()
	l.mtx.Lock()
	task, ok := l.tasks[title]
	if !ok {
		return fmt.Errorf("такой задачи нет")
	}
	task.Complite()

	return nil

}
func (l *List) UnDone(title string) error {

	defer l.mtx.Unlock()
	l.mtx.Lock()
	task, ok := l.tasks[title]
	if !ok {
		return errors.New("такой задачи нет")
	}
	task.Uncomplite()

	return nil
}

func (l *List) Show() []Task {
	defer l.mtx.RUnlock()
	l.mtx.RLock()

	result := make([]Task, 0, len(l.tasks))
	for _, t := range l.tasks {
		result = append(result, *t)
	}

	return result
}
func (l *List) ShowComplitedUnComplited(done bool) map[string]Task {
	defer l.mtx.RUnlock()
	l.mtx.RLock()

	result := make(map[string]Task, len(l.tasks))
	for k, v := range l.tasks {
		if v.IsDone == done {
			result[k] = *v
		}
	}
	return result

}

func (l *List) Del(title string) error {
	defer l.mtx.Unlock()
	l.mtx.Lock()

	_, ok := l.tasks[title]
	if !ok {
		return fmt.Errorf("такой задачи нет")
	}
	delete(l.tasks, title)
	return nil

}
